package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *state, _ command) error {

	ctx := context.Background()

	if err := s.db.ResetDB(ctx); err != nil {
		return fmt.Errorf("reset database %w", err)
	}

	log.Println("All user rows have been deleted")

	return nil
}
