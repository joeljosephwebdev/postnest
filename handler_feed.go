package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joeljosephwebdev/postnest.git/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create feed %w", err)
	}

	_, err = createFeedFollow(s, user.ID, feed.ID)
	if err != nil {
		return err
	}

	fmt.Printf(
		"feedID: %v\nname: %s\nurl: %s\ncreatedAt: %v\n",
		feed.ID,
		feed.Name,
		feed.Url,
		feed.CreatedAt.Format(time.UnixDate),
	)

	return nil
}

func printFeed(feed database.GetFeedsRow) {
	fmt.Printf("feedID:      %v\n", feed.ID)
	fmt.Printf("name:        %s\n", feed.Name)
	fmt.Printf("url:         %s\n", feed.Url)
	fmt.Printf("username:    %v\n", feed.Username)
	fmt.Printf("createdAt:   %v\n", feed.CreatedAt.Format(time.UnixDate))
	fmt.Printf("updatedAt:   %v\n", feed.UpdatedAt.Format(time.UnixDate))
	fmt.Println("================================================================")
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %v", err)
	}

	for _, feed := range feeds {
		printFeed(feed)
	}

	return nil
}
