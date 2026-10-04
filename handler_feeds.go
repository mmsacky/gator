package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, _ command) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error encountered: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("There are no RSS feeds")
		return nil
	}

	for _, feed := range feeds {
		fmt.Println()
		fmt.Println("Name:", feed.FeedName)
		fmt.Println("URL:", feed.Url)
		fmt.Println("User:", feed.UserName)
		fmt.Println()
	}

	return nil
}
