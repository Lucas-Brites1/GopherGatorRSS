package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	"github.com/Lucas-Brites1/RSSGopher/internal/types"
)

func HandlerReset(state *config.Config, command types.Command, args ...interface{}) error {
	if DatabaseInstance, ok := args[0].(*database.Queries); ok {
		err := DatabaseInstance.Reset(context.Background())
		if err != nil {
			return errors.New("failed to reset database")
		}
		fmt.Println("Database reseted successfully")
		return nil
	}
	return errors.New("no database instance provided in arguments")
}
