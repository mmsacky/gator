package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mmsacky/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	limit := 2
	ctx := context.Background()

	if len(cmd.args) > 1 {
		return fmt.Errorf("browse accepts at most one limit argument")
	}

	if len(cmd.args) == 1 {

		argToInt, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("Failed to convert string: %w", err)
		}

		limit = argToInt
	}

	if limit < 1 {
		return fmt.Errorf("limit must be greater than zero")
	}

	getPostsForUserParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(ctx, getPostsForUserParams)
	if err != nil {
		return fmt.Errorf("error getting posts for user: %w", err)
	}

	for _, post := range posts {
		fmt.Println("Feed Name", post.FeedName)
		fmt.Println(post.Title)
		fmt.Println(post.Url)
		if post.PublishedAt.Valid {
			fmt.Println(post.PublishedAt.Time)
		} else {
			fmt.Println("Publication date unavailable")
		}
		fmt.Println()
	}

	return nil
}
