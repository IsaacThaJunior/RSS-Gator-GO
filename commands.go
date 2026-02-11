package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	arrOfCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.arrOfCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	command, exists := c.arrOfCommands[cmd.Name]
	if !exists {
		return fmt.Errorf("Command Not Found")
	}
	return command(s, cmd)
}
