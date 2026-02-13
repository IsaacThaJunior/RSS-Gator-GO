package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"os"

	"github.com/google/uuid"
)



func handleAddFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		fmt.Println("You need a URL to continue")
		os.Exit(1)
	}

	url := cmd.Args[0]
	// Get url from db
	feed, err := s.db.GetFeedByUrl(context.Background(), url)

	if err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}

	data, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		fmt.Printf("Error saving feed to the db: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Here is the Feed Name: %v\n", data.FeedUrl)
	fmt.Printf("Here is the Current user: %v\n", data.UserName)
	return nil
}

func handleGetUserFollow(s *state, cmd command, user database.User) error {

	data, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}

	for _, feed := range data {
		fmt.Printf("Feed Name: %v\n", feed.FeedUrl)
		fmt.Printf("Feed Owner: %v\n", feed.UserName)
		fmt.Printf("Feed Owner: %v\n", feed.FeedName)

	}
	return nil
}
