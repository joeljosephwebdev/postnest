package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

type commandsList struct {
	commands map[string]func(*state, command) error
}

func (c *commandsList) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func (c *commandsList) run(s *state, cmd command) error {
	f, ok := c.commands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
