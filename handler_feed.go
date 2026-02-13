package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"os"
	"time"

	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		fmt.Println("You need a name and URL to continue")
		os.Exit(1)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Url:       url,
		UserID:    user.ID,
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		fmt.Printf("Error saving feed to the db: %v", err)
		os.Exit(1)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		fmt.Printf("Error saving feed follow to the db: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Feed: %v has been added\n", name)
	return nil
}

func handleListFeed(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())

	if err != nil {
		os.Exit(1)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name: %v\n", feed.Name)
		fmt.Printf("Feed URL: %v\n", feed.Url)

		user, err := s.db.GetUserByID(context.Background(), feed.UserID)

		if err != nil {
			os.Exit(1)
		}

		fmt.Printf("Feed Owner: %v\n", user.Name)

	}
	return nil
}
