package common

import (
	"github.com/fusioncatalyst/cli/actions"
	"github.com/urfave/cli/v2"
)

func GetAssembledApp() *cli.App {
	app := &cli.App{
		Name:    "fusioncatalyst CLI tool",
		Usage:   "Manage data contracts in your project",
		Authors: []*cli.Author{{Name: "fusioncatalyst.io", Email: "hi@fusioncatalyst.io"}},
		Version: "v0.0.1",
		Commands: []*cli.Command{
			{
				Name:      "schema-from-json",
				Usage:     "Create a JSON schema from the specified JSON file",
				Action:    actions.SchemaFromJsonAction,
				ArgsUsage: "[arg1] [arg2]",
				Before: func(c *cli.Context) error {
					if c.NArg() == 0 {
						return cli.Exit("Missing JSON file argument", 1)
					}

					if c.NArg() == 1 {
						return cli.Exit("Missing output file argument", 1)
					}
					return nil
				},
			},
			{
				Name:      "create-project",
				Usage:     "Create a new project",
				Action:    actions.CreateProjectAction,
				ArgsUsage: "[arg1]",
				Before: func(c *cli.Context) error {
					if c.NArg() == 0 {
						return cli.Exit("Missing project name", 1)
					}

					return nil
				},
			},
			{
				Name:   "list-projects",
				Usage:  "List all projects available to current team",
				Action: actions.ListProjectsAction,
			},
		},
	}

	return app
}
