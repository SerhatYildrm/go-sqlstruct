package commands

import (
	"github.com/urfave/cli"
)

// GetAllCommands ...
func GetAllCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "mssql",
			Aliases: []string{"m"},
			Usage:   "<TABLENAME> <PATH>",
			Action: func(c *cli.Context) error {
				var err error
				tableName := c.Args().First()
				path := c.Args().Get(1)
				if path == "" {
					path = "."
				}
				if err = MSSQLExecute(tableName, path); err != nil {
					return err
				}
				return nil
			},
		},
	}
}
