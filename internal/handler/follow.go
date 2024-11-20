package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
	"github.com/google/uuid"
)

func HandlerFollow(state *config.Config, command types.Command, args ...interface{}) error {
	if len(command.Tokens) < 2 {
		return errors.New("url do feed não fornecida")
	}

	if DatabaseInstance, ok := args[0].(*database.Queries); ok {
		url := command.Tokens[1]

		if len(args) > 1 {
			if str, ok := args[1].(string); ok {
				fmt.Println("1")
				url = str
			}
		}

		feedID, err := DatabaseInstance.GetIdByURL(context.Background(), url)
		if err != nil {
			return errors.New("erro getidbyurl")
		}

		userID, err := DatabaseInstance.GetIdByName(context.Background(), state.Username)
		if err != nil {
			return errors.New("erro getidbyname")
		}

		_, err = DatabaseInstance.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
			ID:        uuid.New(),
			FeedID:    uuid.NullUUID{UUID: feedID, Valid: true},
			UserID:    uuid.NullUUID{UUID: userID, Valid: true},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		if err != nil {
			return errors.New("error trying to create feed follow")
		}

		feedName, err := DatabaseInstance.GetFeedById(context.Background(), feedID)
		if err != nil {
			return err
		}

		fmt.Printf("user '%s' now is following feed '%s'\n", state.Username, feedName)
		return nil
	}
	return errors.New("no database instance provided in arguments")
}
