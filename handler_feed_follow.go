package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joeljosephwebdev/postnest.git/internal/database"
)

func handlerFeedFollow(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user id for user %s: %v", s.cfg.CurrentUserName, err)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("feed not found %v", err)
	}

	feed_follow, err := createFeedFollow(s, user.ID, feed.ID)
	if err != nil {
		return err
	}

	fmt.Printf("User: %s\nFeed: %s\n", user.Name, feed.Name)
	printFeedFollow(feed_follow)

	return nil
}

func createFeedFollow(
	s *state,
	userID uuid.UUID,
	feedID uuid.UUID,
) (database.FeedsFollow, error) {

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return database.FeedsFollow{}, fmt.Errorf("unable to follow feed %v", err)
	}

	return feed_follow, nil
}

func printFeedFollow(feeds_follow database.FeedsFollow) {
	fmt.Printf("ID: %v\n", feeds_follow.ID)
	fmt.Printf("CreatedAt: %s\n", feeds_follow.CreatedAt.Format(time.UnixDate))
	fmt.Printf("UpdatedAt: %s\n", feeds_follow.UpdatedAt.Format(time.UnixDate))
	fmt.Printf("UserID: %v\n", feeds_follow.UserID)
	fmt.Printf("FeedID: %v\n", feeds_follow.FeedID)
}

func printFeedFollowForUserRow(ffRow database.GetFeedFollowsForUserRow) {
	fmt.Printf("ID: %v\n", ffRow.ID)
	fmt.Printf("Username: %s\n", ffRow.Username)
	fmt.Printf("Feed: %s\n", ffRow.Feedname)
	fmt.Printf("URL: %s\n", ffRow.Url)
	fmt.Printf("CreatedAt: %s\n", ffRow.CreatedAt.Format(time.UnixDate))
	fmt.Printf("UpdatedAt: %s\n", ffRow.UpdatedAt.Format(time.UnixDate))
	fmt.Printf("UserID: %v\n", ffRow.UserID)
	fmt.Printf("FeedID: %v\n", ffRow.FeedID)
}

func showFollowing(s *state, cmd command) error {
	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to retrieve followed feeds: %v", err)
	}

	for _, feed_follow := range feed_follows {
		printFeedFollowForUserRow(feed_follow)
		fmt.Println("===============================================")
	}

	return nil
}
