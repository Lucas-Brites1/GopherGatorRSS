package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
)

func HandlerFollowing(state *config.Config, command types.Command, args ...interface{}) error {
	if DatabaseInstance, ok := args[0].(*database.Queries); ok {
		user, _ := DatabaseInstance.GetUser(context.Background(), state.Username)
		followedFeeds, err := DatabaseInstance.GetFeedFollowsForUser(context.Background(), user.ID)
		if err != nil {
			return errors.New("error while trying to get followed feeds")
		}
		for i, feed := range followedFeeds {
			fmt.Printf("Feed %d\nName: %s\nURL: %s\n", i, feed.FeedName, feed.FeedUrl)
		}
		return nil
	}
	return errors.New("no database instance provided in arguments")
}
