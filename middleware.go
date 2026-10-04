package main

import (
	"context"
	"fmt"

	"github.com/mmsacky/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {

		user, err := s.db.GetUser(context.Background(), s.cfg.UserName)
		if err != nil {
			return fmt.Errorf("error retrieving database user: %w", err)
		}

		err = handler(s, cmd, user)
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}

		return nil
	}
}
