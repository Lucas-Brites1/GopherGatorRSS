package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	handlerUtils "github.com/Lucas-Brites1/RSSGopher/internal/handler/utils"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
	"github.com/google/uuid"
)

func HandlerRegister(state *config.Config, command types.Command, args ...interface{}) error {
	if len(command.Tokens) < 2 {
		return errors.New("invalid number of tokens in login command. Expected username")
	}

	if len(args) > 0 {
		username := command.Tokens[1]
		if DatabaseInstance, ok := args[0].(*database.Queries); ok {
			exists, err := handlerUtils.UserExists(DatabaseInstance, username)
			if err != nil {
				return err
			}
			if exists {
				return fmt.Errorf("user already exists")
			}

			_, err = DatabaseInstance.CreateUser(context.Background(), database.CreateUserParams{
				ID:        uuid.New(),
				Name:      username,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})

			if err != nil {
				return fmt.Errorf("error while trying to create a new user: %v", err)
			}
			state.SetUser(username)
			fmt.Printf("User %s has been set. (handlerRegister)\n", username)
			return nil
		} else {
			return fmt.Errorf("expected first argument to be of type *database.Queries")
		}
	} else {
		return errors.New("no database instance provided in arguments")
	}
}
