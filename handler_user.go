package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joeljosephwebdev/postnest.git/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("missing login credentials")
	}
	username := cmd.Args[0]

	User, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("login error - %v", err)
	}

	err = s.cfg.SetUser(User.Name)
	if err != nil {
		return err
	}

	fmt.Println("User switched successfully!")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("missing login credentials")
	}
	name := cmd.Args[0]

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	User, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	s.cfg.SetUser(name)
	fmt.Println("User created successfully!")
	log.Printf("\nID: %s\nName: %s\nCreated_at: %s\n", User.ID.String(), User.Name, User.CreatedAt.Format(time.RFC822))

	return nil
}
