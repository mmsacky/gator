package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, _ command) error {

	ctx := context.Background()

	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("error encountered: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("There are no registered users")
		return nil
	}

	for _, user := range users {

		if user.Name == s.cfg.UserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}

	}

	return nil
}
