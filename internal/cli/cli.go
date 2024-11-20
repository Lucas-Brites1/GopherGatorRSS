package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	utils "github.com/Lucas-Brites1/RSSGopher/internal/cli/utils"
	config "github.com/Lucas-Brites1/RSSGopher/internal/config"
	"github.com/Lucas-Brites1/RSSGopher/internal/database"
	handler "github.com/Lucas-Brites1/RSSGopher/internal/handler"
	types "github.com/Lucas-Brites1/RSSGopher/internal/types"
)

type Cli struct {
	Reader   utils.Reader
	PConfig  *config.Config
	Database *database.Queries
	Commands *Commands
	Running  bool
}

func NewCli(cfg *config.Config, dbQuery *database.Queries) *Cli {
	cli := &Cli{
		Reader:   utils.Reader{},
		PConfig:  cfg,
		Database: dbQuery,
		Commands: NewCommands(),
		Running:  true,
	}

	cli.Commands.CreateCommands([]struct {
		name string
		cb   func(state *config.Config, command types.Command, args ...interface{}) error
	}{
		{
			name: "login",
			cb:   handler.HandlerLogin,
		},
		{
			name: "register",
			cb:   handler.HandlerRegister,
		},
		{
			name: "reset",
			cb:   handler.HandlerReset,
		},
		{
			name: "users",
			cb:   handler.HandlerGetUsers,
		},
		{
			name: "agg",
			cb:   handler.HandlerAggregate,
		},
		{
			name: "addfeed",
			cb:   handler.HandlerAddFeed,
		},
		{
			name: "feeds",
			cb:   handler.HandlerGetFeeds,
		},
		{
			name: "follow",
			cb:   handler.HandlerFollow,
		},
		{
			name: "following",
			cb:   handler.HandlerFollowing,
		},
		{
			name: "unfollow",
			cb:   handler.HandlerUnfollow,
		},
		{
			name: "browse",
			cb:   handler.HandlerBrowse,
		},
		{
			name: "clear",
			cb:   cli.Clear,
		},
		{
			name: "exit",
			cb:   cli.Exit,
		},
	})

	return cli
}

func (cli *Cli) Run() {
	for cli.Running {
		cli.Reader.Input()
		tokens := cli.Reader.Tokenize()

		command := types.Command{
			Name:   tokens[0],
			Tokens: tokens,
		}

		err := cli.Commands.execute(cli.PConfig, command, cli.Database)
		if err != nil {
			log.Println("Error: ", err)
		} else {
			if command.Name != "clear" {
				fmt.Printf("command executed successfully: %s\n", command.Name)
			}
		}
	}
}

func (c *Cli) Clear(state *config.Config, command types.Command, args ...interface{}) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

func (c *Cli) Exit(state *config.Config, command types.Command, args ...interface{}) error {
	c.Running = false
	return nil
}
