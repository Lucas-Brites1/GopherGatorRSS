package handlerUtils

import (
	"context"
	"fmt"

	"github.com/Lucas-Brites1/RSSGopher/internal/database"
)

func UserExists(db *database.Queries, username string) (exists bool, err error) {
	user, err := db.GetUser(context.Background(), username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, fmt.Errorf("error fetching user: %v", err)
	}

	return user.Name != "", nil
}
