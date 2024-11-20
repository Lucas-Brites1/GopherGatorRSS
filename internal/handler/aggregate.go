package handler

import (
	"fmt"
	"time"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	handlerUtils "github.com/Lucas-Brites1/RSSGopher/internal/handler/utils"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
)

const time_between_reqs = "1m"

func HandlerAggregate(state *config.Config, command types.Command, args ...interface{}) error {
	if len(command.Tokens) < 2 {
		return fmt.Errorf("URL not provided. Use the command in the format: 'agg <URL>'")
	}

	timerBeetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}

	fmt.Println("Collecting feeds every " + timerBeetweenRequests.String())
	ticker := time.NewTicker(timerBeetweenRequests)
	go func() {
		for {
			err := handlerUtils.ScrapeFeeds(args[0].(*database.Queries))
			if err != nil {
				fmt.Println("error trying to scraping: ", err)
			}
			<-ticker.C
		}
	}()

	return nil
}
