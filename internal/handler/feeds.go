package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
)

func HandlerGetFeeds(state *config.Config, command types.Command, args ...interface{}) error {
	if DatabaseInstance, ok := args[0].(*database.Queries); ok {
		feeds, err := DatabaseInstance.GetFeeds(context.Background())
		if err != nil {
			return errors.New("error while trying to get feeds")
		}

		for i, feed := range feeds {
			name, err := DatabaseInstance.GetNameByID(context.Background(), feed.UserID.UUID)
			if err != nil {
				continue
			}
			fmt.Printf("%d - %s| %s| %s\n", i, feed.Name, feed.Url, name)
		}
		return nil
	}
	return errors.New("no database instance provided in arguments")
}
