package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
)

func HandlerGetUsers(state *config.Config, command types.Command, args ...interface{}) error {
	if DatabaseInstance, ok := args[0].(*database.Queries); ok {
		users, err := DatabaseInstance.GetUsers(context.Background())
		if err != nil {
			return errors.New("error while trying to get users")
		}
		for _, user := range users {
			if state.Username == user.Name {
				fmt.Println("> " + user.Name)
				continue
			}
			fmt.Println("* " + user.Name)
		}
		return nil
	}
	return errors.New("no database instance provided in arguments")
}
