package cli

import (
	"errors"
	"fmt"

	config "github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
)

type Commands struct {
	RegisteredCommands map[string]func(state *config.Config, command types.Command, args ...interface{}) error
}

func (c *Commands) register(name string, f func(state *config.Config, command types.Command, args ...interface{}) error) error {
	_, exist := c.RegisteredCommands[name]
	if exist {
		return errors.New("this command already exists")
	}

	c.RegisteredCommands[name] = f
	return nil
}

func (c *Commands) execute(state *config.Config, command types.Command, args ...interface{}) error {
	handler, exists := c.RegisteredCommands[command.Name]
	if !exists {
		return fmt.Errorf("command '%s' not found", command.Name)
	}

	return handler(state, command, args...)
}

func NewCommands() *Commands {
	return &Commands{
		RegisteredCommands: make(map[string]func(state *config.Config, command types.Command, args ...interface{}) error),
	}
}

func (c *Commands) CreateCommands(commands []struct {
	name string
	cb   func(state *config.Config, command types.Command, args ...interface{}) error
}) {
	for _, command := range commands {
		err := c.register(command.name, command.cb)
		if err != nil {
			fmt.Printf("error trying to create command: %v", command.name)
		}
	}
}
