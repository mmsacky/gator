package main

import (
	"errors"
	"fmt"

	"github.com/mmsacky/gator/internal/config"
	"github.com/mmsacky/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {

	if command, exists := c.commandMap[cmd.name]; exists {
		err := command(s, cmd)
		if err != nil {
			return fmt.Errorf("unable to run command: %w", err)
		}
	} else {
		return errors.New("command doesn't exist")
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}
