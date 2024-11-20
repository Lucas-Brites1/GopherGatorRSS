package main

import (
	"database/sql"
	"log"

	"github.com/Lucas-Brites1/RSSGopher/internal/cli"
	Config "github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	config, err := Config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	DatabaseQueries := database.New(db)
	if err != nil {
		log.Fatal(err)
	}
	CLI := cli.NewCli(config, DatabaseQueries)
	CLI.Run()
}
