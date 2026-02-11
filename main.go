package main

import (
	"gator/internal/config"
	"log"
	"os"
)

type state struct {
	conPointer *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		conPointer: &cfg,
	}

	cmds := commands{
		arrOfCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("You need to pass in enough args like login or register")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})

	if err != nil {
		log.Fatal(err)
	}
}
