package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mmsacky/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 1 {
		return errors.New("the unfollow handler expects a single argument, the URL")
	}
	ctx := context.Background()
	feedURL := cmd.args[0]

	feed, err := s.db.GetFeedByURL(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error retrieving feed: %w", err)
	}

	deleteFeedFollowParams := database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}

	err = s.db.DeleteFeedFollow(ctx, deleteFeedFollowParams)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}

	fmt.Printf("%s unfollowed successfully!\n", feed.Name)

	return nil
}
