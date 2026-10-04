package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mmsacky/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 2 {
		return errors.New("the add feed handler expects two arguments, a feed name and a feed url")
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]

	ctx := context.Background()
	now := time.Now()

	newFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	}

	newFeed, err := s.db.CreateFeed(ctx, newFeedParams)
	if err != nil {
		return fmt.Errorf("error creating RSS feed: %w", err)
	}

	fmt.Printf("%+v\n", newFeed)

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    newFeed.ID,
	}

	feedFollowRecord, err := s.db.CreateFeedFollow(ctx, newFeedFollow)
	if err != nil {
		return fmt.Errorf("error creating new feed follow record: %w", err)
	}

	fmt.Println("You are now following:", feedFollowRecord.FeedName)

	return nil
}
