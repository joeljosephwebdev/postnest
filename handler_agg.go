package main

import (
	"context"
	"fmt"
)

func agg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error aggregating feed: %w", err)
	}

	fmt.Println(feed)
	return nil
}
