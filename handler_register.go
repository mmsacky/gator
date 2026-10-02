package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/mmsacky/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("the register handler expects a single argument, the username")
	}

	ctx := context.Background()

	username := cmd.args[0]

	newUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	_, err := s.db.GetUser(ctx, newUserParams.Name)

	if errors.Is(err, sql.ErrNoRows) {
		newUser, err := s.db.CreateUser(ctx, newUserParams)
		if err != nil {
			return fmt.Errorf("error creating new user %w", err)
		}

		s.cfg.SetUser(newUser.Name)

		fmt.Printf("New user %s was created and has been set as the current username\n", newUser.Name)

	} else if err == nil {
		return errors.New("user already exists")
	} else {
		return fmt.Errorf("error checking for existing user: %w", err)
	}

	return nil
}
