package main

import (
	"context"
	"errors"
	"fmt"
	"time"
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
		scrapeFeeds(s)
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

	for _, item := range RSSFeed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
