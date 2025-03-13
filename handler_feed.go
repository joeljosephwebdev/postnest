package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joeljosephwebdev/postnest.git/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	User, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("login error - %v", err)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    User.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create feed %w", err)
	}

	printFeed(feed)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("feedID:      %v\n", feed.ID)
	fmt.Printf("name:        %s\n", feed.Name)
	fmt.Printf("url:         %s\n", feed.Url)
	fmt.Printf("userID:      %v\n", feed.UserID)
	fmt.Printf("createdAt:   %v\n", feed.CreatedAt.Format(time.UnixDate))
	fmt.Printf("UpdatedAt:   %v\n", feed.UpdatedAt.Format(time.UnixDate))
}
