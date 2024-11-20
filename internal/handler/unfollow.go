package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
	"github.com/google/uuid"
)

func HandlerUnfollow(state *config.Config, command types.Command, args ...interface{}) error {
	if len(command.Tokens) < 2 {
		return errors.New("url do feed não fornecida")
	}

	if DatabaseInstance, ok := args[0].(*database.Queries); ok {
		url := command.Tokens[1]
		userID, err := DatabaseInstance.GetIdByName(context.Background(), state.Username)
		if err != nil {
			return err
		}

		err = DatabaseInstance.Unfollow(context.Background(), database.UnfollowParams{
			Url:    url,
			UserID: uuid.NullUUID{UUID: userID, Valid: true},
		})
		if err != nil {
			return err
		}

		Feed, err := DatabaseInstance.GetFeedByURL(context.Background(), url)
		if err != nil {
			return err
		}
		fmt.Printf("user '%s' now is unfollowing feed '%s'\n", state.Username, Feed.Name)
		return nil
	}
	return errors.New("no database instance provided in arguments")
}
