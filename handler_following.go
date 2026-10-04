package main

import (
	"context"
	"fmt"

	"github.com/mmsacky/gator/internal/database"
)

func handlerFollowing(s *state, _ command, user database.User) error {

	ctx := context.Background()

	followingFeeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("error retrieving your following feeds: %w", err)
	}

	if len(followingFeeds) == 0 {
		fmt.Println("You are not following any feeds")
		return nil
	}

	for _, feed := range followingFeeds {
		fmt.Println("Feed Name:", feed.FeedName)
	}

	return nil
}
