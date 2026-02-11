package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("You need a name to continue")
	}
	name := cmd.Args[0]
	err := s.conPointer.SetUser(name)
	if err != nil {
		return errors.New("Couldnt set the User")
	}
	fmt.Printf("The user: %v has been set\n", name)
	return nil
}
