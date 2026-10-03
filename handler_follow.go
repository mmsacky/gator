package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mmsacky/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return errors.New("the follow handler expects a single argument, the URL")
	}

	ctx := context.Background()

	feedURL := cmd.args[0]
	now := time.Now()
	username := s.cfg.UserName

	user, err := s.db.GetUser(ctx, username)
	if err != nil {
		return fmt.Errorf("error retrieving user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error retrieving feed: %w", err)
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollowRecord, err := s.db.CreateFeedFollow(ctx, newFeedFollow)
	if err != nil {
		return fmt.Errorf("error creating new feed follow record: %w", err)
	}

	fmt.Println("Feed Name:", feedFollowRecord.FeedName)
	fmt.Println("User:", feedFollowRecord.UserName)

	return nil
}
