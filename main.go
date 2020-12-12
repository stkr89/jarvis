package main

import (
	"github.com/stkr89/jarvis/commands"
	"github.com/stkr89/jarvis/commands/apply"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Description: "Automate your routine tasks",
		Commands: []*cli.Command{
			{
				Name:   "verify",
				Usage:  "Verify automation config file",
				Action: commands.Verify(),
			},
			{
				Name:   "apply",
				Usage:  "Apply automation",
				Action: apply.Apply(),
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
