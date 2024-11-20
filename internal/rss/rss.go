package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"

	rsstypes "github.com/Lucas-Brites1/RSSGopher/internal/rss/rss_types"
)

func FetchFeed(ctx context.Context, feedURL string) (*rsstypes.RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gophergator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	Feed := &rsstypes.RSSFeed{}
	err = xml.Unmarshal(response, Feed)
	if err != nil {
		return nil, fmt.Errorf("error trying to deserialize XML: %v", err)
	}

	Feed.Channel.Title = html.UnescapeString(Feed.Channel.Title)
	Feed.Channel.Description = html.UnescapeString(Feed.Channel.Description)

	for _, item := range Feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
	return Feed, nil
}
