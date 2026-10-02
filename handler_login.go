package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mmsacky/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}

	context := context.Background()
	username := cmd.args[0]

	user, _ := s.db.GetUser(context, username)

	if user == (database.User{}) {
		return errors.New("this user doesn't exist in the database")
	}

	s.cfg.SetUser(username)

	fmt.Printf("%s has been set as the current username\n", username)

	return nil
}
