package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mmsacky/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {

	if len(cmd.args) < 2 {
		return errors.New("the add feed handler expects a two arguments, a feed name and a feed url")
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]

	ctx := context.Background()
	currentUser := s.cfg.UserName

	dbUser, err := s.db.GetUser(ctx, currentUser)
	if err != nil {
		return fmt.Errorf("error retrieving database user: %w", err)
	}

	newFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    dbUser.ID,
	}

	newFeed, err := s.db.CreateFeed(ctx, newFeedParams)
	if err != nil {
		return fmt.Errorf("error creating RSS feed: %w", err)
	}

	fmt.Printf("%+v\n", newFeed)

	return nil
}
