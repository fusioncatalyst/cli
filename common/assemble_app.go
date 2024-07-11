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
				Name:   "schema-from-json",
				Usage:  "Create a JSON schema from the specified JSON file",
				Action: actions.SchemaFromJsonAction,
			},
		},
	}

	return app
}
