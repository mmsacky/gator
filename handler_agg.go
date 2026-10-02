package main

import (
	"context"
	"fmt"
)

const rssFeedURL = "https://www.wagslane.dev/index.xml"

func aggregator(s *state, cmd command) error {

	ctx := context.Background()

	rssFeed, err := fetchFeed(ctx, rssFeedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Printf("%+v\n", rssFeed)

	return nil
}
