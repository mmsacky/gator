package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}

	username := cmd.args[0]

	s.ConfigPTR.SetUser(username)

	fmt.Printf("%s has been set as the current username", username)

	return nil
}
