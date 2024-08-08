package actions

import (
	"github.com/fusioncatalyst/cli/api"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/urfave/cli/v2"
)

func CreateProjectAction(cCtx *cli.Context) error {
	projectName := cCtx.Args().Get(0)

	// Call the FusionCatalyst API to convert JSON to schema
	apiClient := api.NewFCApiClient(utils.GetFCHost())
	_, e := apiClient.CallPrivateNewProject(projectName)
	if e != nil {
		return e
	}

	return nil
}
