package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
	"github.com/google/uuid"
)

func HandlerBrowse(state *config.Config, command types.Command, args ...interface{}) error {
	if DatabaseInstance, ok := args[0].(*database.Queries); ok {
		limit := 2
		if len(command.Tokens) > 1 {
			parsedLimit, _ := strconv.Atoi(command.Tokens[1])
			limit = parsedLimit
		}

		userID, err := DatabaseInstance.GetIdByName(context.Background(), state.Username)
		if err != nil {
			return err
		}

		Posts, err := DatabaseInstance.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
			UserID: uuid.NullUUID{UUID: userID, Valid: true},
			Limit:  int32(limit),
		})
		if err != nil {
			return err
		}

		for i, post := range Posts {
			fmt.Printf("Feed %d\nTitle: %s\nDescription: %s\nURL: %s\n", i, post.Title.String, post.Description.String, post.Url.String)
		}

		return nil
	}
	return errors.New("no database instance provided in arguments")
}
