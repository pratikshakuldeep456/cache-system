package main

import (
	"log"
	"os"
	"pratikshakuldeep456/cache-system/command"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cache system",
		Usage: "implement cache system",
		Commands: []*cli.Command{
			{
				Name:   "start",
				Usage:  "start the session",
				Action: command.StartCacheSystem,
			}}}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
