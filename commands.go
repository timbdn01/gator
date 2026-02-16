package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}

func (c *commands) reset() {
	c.registeredCommands = make(map[string]func(*state, command) error)
}

func (c *commands) users() []string {
	keys := make([]string, 0, len(c.registeredCommands))
	for k := range c.registeredCommands {
		keys = append(keys, k)
	}
	return keys
}


