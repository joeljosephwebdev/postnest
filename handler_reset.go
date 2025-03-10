package main

import (
	"context"
	"fmt"
	"log"
)

func handleReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset users: %w", err)
	}
	log.Printf("User table reset!\n")
	return nil
}
