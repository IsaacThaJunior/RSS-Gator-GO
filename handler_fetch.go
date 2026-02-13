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

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Url:    url,
		UserID: user.ID,
		ID:     uuid.New(),
		Name:   name,
	})

	if err != nil {
		fmt.Printf("Error saving feed to the db: %v", err)
		os.Exit(1)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
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

func handleAddFeedFollow(s *state, cmd command) error {
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

	// Get Logged in User
	user, err := s.db.GetUser(context.Background(), s.conPointer.CurrentUserName)
	if err != nil {
		fmt.Println("Not Logged in")
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

func handleGetUserFollow(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.conPointer.CurrentUserName)

	if err != nil {
		fmt.Println("Not Logged in")
		os.Exit(1)
	}
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
