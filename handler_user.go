package main

import (
	"context"
	"fmt"
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
	printUser(User)
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
	printUser(User)

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	Users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	for _, user := range Users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
