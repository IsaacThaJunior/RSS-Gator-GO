package main

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("You need a name to continue")
	}
	name := cmd.Args[0]

	// Check DB if the user exists
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		fmt.Println("User does not exist")
		os.Exit(1)

	}

	err = s.conPointer.SetUser(name)
	if err != nil {
		return errors.New("Couldnt set the User")
	}
	fmt.Printf("The user: %v has been set\n", name)
	return nil
}

func handleRegisterUser(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("You need a name to continue")
	}
	name := cmd.Args[0]

	data, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		Name:      name,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			os.Exit(1)
			return errors.New("user already exists")
		}
		return err
	}

	err = s.conPointer.SetUser(name)
	if err != nil {
		return errors.New("Couldnt set the User")
	}

	fmt.Printf("The user: %v has been created\n", name)
	fmt.Printf("Here are the user details: %v\n", data)

	return nil
}
