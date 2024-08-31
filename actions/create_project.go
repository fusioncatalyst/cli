package actions

import (
	"fmt"
	"github.com/fusioncatalyst/cli/api"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/urfave/cli/v2"
)

func CreateProjectAction(cCtx *cli.Context) error {
	returnID := cCtx.Bool("return-id")
	projectName := cCtx.String("project-name")

	// Call the FusionCatalyst API to convert JSON to schema
	apiClient := api.NewFCApiClient(utils.GetFCHost())
	result, e := apiClient.CallPrivateNewProject(projectName)
	if e != nil {
		return e
	}
	if returnID {
		fmt.Println(result.ID)
	}

	return nil
}
