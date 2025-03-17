package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joeljosephwebdev/postnest.git/internal/database"
)

func agg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %s <time_string>", cmd.Name)
	}

	interval, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to parse time duration %w", err)
	}

	log.Printf("Collecting feeds every %s...\n", interval)
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		ScrapeFeeds(s)
	}
}

func ScrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("unable to get next feed: ", err)
		return
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		log.Printf("failed to mark %s feed as fetched: %v\n", nextFeed.Name, err)
		return
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		log.Println("unable to aggregate feed: ", err)
		return
	}

	for _, item := range feed.Channel.Item {
		if item.Title == "" {
			continue
		}
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: item.PubDate,
			FeedID:      nextFeed.ID,
		}
		_, err := s.db.CreatePost(context.Background(), params)
		if err != nil && !strings.Contains(err.Error(), "duplicate") {
			log.Printf("failed to save post %s: %v", item.Title, err)
		}
	}
	log.Printf("Feed %s collected, %v posts saved!", nextFeed.Name, len(feed.Channel.Item))
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2

	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage %s <row_limit>[optional]", cmd.Name)
	}

	if len(cmd.Args) == 1 {
		input := cmd.Args[0]
		temp_limit, err := strconv.Atoi(input)
		if err != nil {
			return fmt.Errorf("invalid row limit: %w", err)
		}
		limit = int32(temp_limit)
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to retrieve posts: %w", err)
	}

	for _, post := range posts {
		feed, err := s.db.GetFeedByID(context.Background(), post.FeedID)
		if err != nil {
			return fmt.Errorf("feed not found: %w", err)
		}
		fmt.Printf("%s from %s\n", post.PublishedAt, feed.Name)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
