package actions

import (
	"github.com/fusioncatalyst/cli/api"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
	"os"
)

func ListProjectsAction(cCtx *cli.Context) error {
	apiClient := api.NewFCApiClient(utils.GetFCHost())
	apiResponse, e := apiClient.CallPrivateListProjects()
	if e != nil {
		return e
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Project", "ID"})
	for _, project := range *apiResponse {
		t.AppendRow([]interface{}{project.Name, project.ID})
	}
	t.Render()

	//fmt.Println(apiResponse)

	return nil
}
