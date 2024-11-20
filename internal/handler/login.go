package handler

import (
	"errors"
	"fmt"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	handlerUtils "github.com/Lucas-Brites1/RSSGopher/internal/handler/utils"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
)

func HandlerLogin(state *config.Config, command types.Command, args ...interface{}) error {
	if len(command.Tokens) < 2 {
		return errors.New("invalid number of tokens in login command. Expected username")
	}

	username := command.Tokens[1]

	if DatabaseInstance, ok := args[0].(*database.Queries); ok {
		exists, err := handlerUtils.UserExists(DatabaseInstance, username)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("user %v - does not exist", username)
		}

		state.SetUser(username)
		fmt.Printf("User %s has been set. (handlerLogin)\n", username)
		return nil
	}

	state.SetUser(username)
	fmt.Printf("User %s has been set. (handlerLogin)\n", username)
	return nil
}
