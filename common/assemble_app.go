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
				Name:   "create-project",
				Usage:  "Create a new project",
				Action: actions.CreateProjectAction,
				Before: func(c *cli.Context) error {
					projectName := c.String("project-name")
					if projectName == "" {
						return cli.Exit("Missing required --project-name flag", 1)
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "project-name",
						Usage:    "The name of the project to create",
						Required: true,
					},
					&cli.BoolFlag{
						Name:     "return-id",
						Usage:    "If set, the UUID of the newly created schema will be returned",
						Required: false,
					},
				},
			},
			{
				Name:   "list-projects",
				Usage:  "List all projects available to current team",
				Action: actions.ListProjectsAction,
			},
			{
				Name:   "new-schema",
				Usage:  "Create a new schema from the specified JSON file and associate it with a project",
				Action: actions.NewSchemaAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "project-id",
						Usage:    "The UUID of the project to associate with this schema",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "file",
						Usage:    "Full path to the file containing the schema",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "schema-name",
						Usage:    "The name of the schema to reflect in the registry",
						Required: true,
					},
					&cli.BoolFlag{
						Name:     "return-id",
						Usage:    "If set, the UUID of the newly created schema will be returned",
						Required: false,
					},
				},
				Before: func(c *cli.Context) error {
					projectID := c.String("project-id")
					if projectID == "" {
						return cli.Exit("Missing required --project-id flag", 1)
					}

					schemaFile := c.String("file")
					if schemaFile == "" {
						return cli.Exit("Missing required --file flag", 1)
					}

					schemaName := c.String("schema-name")
					if schemaName == "" {
						return cli.Exit("Missing required --schema-name flag", 1)
					}
					return nil
				},
			},
		},
	}

	return app
}
