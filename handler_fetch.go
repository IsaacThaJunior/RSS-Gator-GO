package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"os"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return err
	}

	fmt.Printf("Feed: %v\n", feed)
	return nil
}

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		fmt.Println("You need a name and URL to continue")
		os.Exit(1)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	// Get Logged in User
	user, err := s.db.GetUser(context.Background(), s.conPointer.CurrentUserName)
	if err != nil {
		fmt.Println("Not Logged in")
		os.Exit(1)
	}

	_, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Url:    url,
		UserID: user.ID,
		ID:     uuid.New(),
		Name:   name,
	})

	if err != nil {
		fmt.Printf("Error saving feed to the db: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Feed: %v has been added\n", name)
	return nil
}
