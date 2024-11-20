package handlerUtils

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	"github.com/Lucas-Brites1/RSSGopher/internal/rss"
	"github.com/google/uuid"
)

func ScrapeFeeds(Database *database.Queries) error {
	nextFeed, err := Database.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	id := nextFeed.ID
	Feed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, item := range Feed.Channel.Item {
		layout := "Mon, 02 Jan 2006 15:04:05 GMT"
		pubDate, _ := time.Parse(layout, item.PubDate)
		//fmt.Printf("\n> Feed %d\nTitle: %s\nDescription: %s\nLink: %s\nDate: %s\n", i, item.Title, item.Description, item.Link, item.PubDate)
		_, err := Database.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       sql.NullString{String: item.Title, Valid: true},
			Description: sql.NullString{String: item.Description, Valid: true},
			Url:         sql.NullString{String: item.Link, Valid: true},
			FeedID:      uuid.NullUUID{UUID: id, Valid: true},
			PublishedAt: pubDate,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		if err != nil {
			return fmt.Errorf("error trying to create new post %v", err)
		}
	}

	err = Database.MarkFeedFetch(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}

	return nil
}
