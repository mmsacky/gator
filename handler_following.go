package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, _ command) error {

	ctx := context.Background()

	username := s.cfg.UserName

	user, err := s.db.GetUser(ctx, username)
	if err != nil {
		return fmt.Errorf("error retrieving user: %w", err)
	}

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
