package main

import (
	"context"
	"fmt"
	"time"

	"gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	//takes a single url argument and creates a new feed follow record for the current user. It should print the name of the feed and the current user once the record is created
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed url>", cmd.Name)
	}
	url := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}
	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Printf("Successfully followed feed '%s' by user '%s'\n", follow.FeedName, follow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	//takes no arguments and prints out all the feeds that the current user is following. It should print the feed name and url for each feed.
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}
	if len(follows) == 0 {
		fmt.Println("You are not following any feeds.")
		return nil
	}
	fmt.Println("You are following these feeds:")
	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	return nil
}
