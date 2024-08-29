package actions

import (
	"fmt"
	"github.com/fusioncatalyst/cli/api"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

func NewSchemaAction(cCtx *cli.Context) error {
	patchToSchemaFile := cCtx.String("file")
	projectID := cCtx.String("project-id")
	schemaName := cCtx.String("schema-name")

	// Open the schema file
	schemaFile, err := os.Open(patchToSchemaFile)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Error opening schema file: %v", err), 1)
	}
	defer schemaFile.Close()

	// Read the outputFile content
	byteContent, err := io.ReadAll(schemaFile)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Error reading JSON file: %v", err), 1)
	}
	fileContent := string(byteContent)

	// Call the FusionCatalyst API to create schema file
	apiClient := api.NewFCApiClient(utils.GetFCHost())
	_, e := apiClient.CallPrivateNewJSONSchema(fileContent, schemaName, projectID)
	if e != nil {
		return cli.Exit(e.Error(), 1)
	}

	return nil
}
