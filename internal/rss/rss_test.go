package rss_test

import (
	"context"
	"testing"

	"github.com/Lucas-Brites1/RSSGopher/internal/rss"
)

func TestPrintMessage(t *testing.T) {
	rss.FetchFeed(context.Background(), "https://pkg.go.dev/net/http#pkg-overview")
}
