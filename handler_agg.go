package main

import (
	"context"
	"fmt"
	"log"
	"time"
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
		log.Printf("- %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found!", nextFeed.Name, len(feed.Channel.Item))
}
