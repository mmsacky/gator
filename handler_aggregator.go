package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mmsacky/gator/internal/database"
)

func handlerAggregator(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return errors.New("the follow handler expects a single argument, the time between requests")
	}

	timeBetweenRequestsString := cmd.args[0]

	timeBetweenRequests, err := time.ParseDuration(timeBetweenRequestsString)
	if err != nil {
		return fmt.Errorf("error converting time string: %w", err)
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Println("Collecting feeds every ", timeBetweenRequests)
		if err := scrapeFeeds(s); err != nil {
			fmt.Println(err)
		}
	}

}

func scrapeFeeds(s *state) error {

	ctx := context.Background()

	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving next feed: %w", err)
	}

	err = s.db.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	RSSFeed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	for _, item := range RSSFeed.Channel.Item {
		now := time.Now()

		newPostDescription := sql.NullString{}

		if item.Description != "" {
			newPostDescription.String = item.Description
			newPostDescription.Valid = true
		}

		newPostPublishedAt := sql.NullTime{}

		if item.PubDate != "" {
			parsePublishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err == nil {
				newPostPublishedAt = sql.NullTime{
					Time:  parsePublishedAt,
					Valid: true,
				}

			}

			newPost := database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   now,
				UpdatedAt:   now,
				Title:       item.Title,
				Url:         item.Link,
				Description: newPostDescription,
				PublishedAt: newPostPublishedAt,
				FeedID:      nextFeed.ID,
			}

			_, err = s.db.CreatePost(ctx, newPost)
			if err != nil {
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) && pgErr.Code == "23505" {
					continue
				}
				return fmt.Errorf("unexpected database error: %w", err)
			}

		}

	}
	return nil
}
