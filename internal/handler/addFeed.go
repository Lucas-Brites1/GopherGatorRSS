package handler

import (
	"context"
	"time"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	handlerUtils "github.com/Lucas-Brites1/RSSGopher/internal/handler/utils"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
	"github.com/google/uuid"
)

func HandlerAddFeed(state *config.Config, command types.Command, args ...interface{}) error {
	if DatabaseInstance, ok := args[0].(*database.Queries); ok {
		user, err := DatabaseInstance.GetUser(context.Background(), state.Username)
		if err != nil {
			return err
		}

		userID := uuid.NullUUID{UUID: user.ID, Valid: true}

		FeedCreated, err := DatabaseInstance.CreateFeed(context.Background(), database.CreateFeedParams{
			ID:        uuid.New(),
			UserID:    userID,
			Name:      handlerUtils.GetName(command.Tokens),
			Url:       handlerUtils.GetURL(command.Tokens),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		HandlerFollow(state, command, DatabaseInstance, FeedCreated.Url)
		if err != nil {
			return err
		}
	}
	return nil
}
