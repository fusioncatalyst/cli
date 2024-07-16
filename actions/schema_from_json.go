package actions

import (
	"fmt"
	"github.com/fusioncatalyst/cli/api"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

func SchemaFromJsonAction(cCtx *cli.Context) error {
	if cCtx.NArg() == 0 {
		return cli.Exit("Missing JSON file argument", 1)
	}

	jsonFilePath := cCtx.Args().Get(0)

	// Open the JSON file
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Error opening JSON file: %v", err), 1)
	}
	defer jsonFile.Close()

	// Read the file content
	byteContent, err := io.ReadAll(jsonFile)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Error reading JSON file: %v", err), 1)
	}
	stringContent := string(byteContent)
	if !utils.IsValidJSON(stringContent) {
		return cli.Exit(fmt.Sprintf("Invalid JSON content in the file: %s", jsonFilePath), 1)
	}

	// Call the FusionCatalyst API to convert JSON to schema
	apiClient := api.NewFCApiClient(utils.GetFCHost())
	apiClient.CallPublicConvertor(stringContent)

	return nil
}
