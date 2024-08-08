package actions

import (
	"fmt"
	"github.com/fusioncatalyst/cli/api"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/urfave/cli/v2"
)

func ListProjectsAction(cCtx *cli.Context) error {
	apiClient := api.NewFCApiClient(utils.GetFCHost())
	apiResponse := apiClient.CallPrivateListProjects()
	fmt.Println(apiResponse)

	return nil
}
