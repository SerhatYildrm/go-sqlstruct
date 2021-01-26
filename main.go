package main

import (
	"os"

	"github.com/urfave/cli"

	"gocompare/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "gocompare"
	app.Usage = "Text"
	app.Version = "0.1"
	app.Commands = commands.GetAllCommands()
	app.Run(os.Args)
}
