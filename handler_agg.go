package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("Not enough arguments")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])

	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		log.Printf("Collecting data every %v", timeBetweenRequests)

		scrapeFeed(s)
	}

}

func scrapeFeed(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		return err
	}

	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	data, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range data.Channel.Item {
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return err
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			Title: sql.NullString{
				String: item.Title,
				Valid:  item.Title != "",
			},
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: sql.NullTime{
				Time:  t,
				Valid: t.Unix() != 0,
			},
			FeedID:    feed.ID,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}

	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(data.Channel.Item))

	return nil
}
