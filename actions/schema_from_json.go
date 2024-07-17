package actions

import (
	"encoding/json"
	"fmt"
	"github.com/fusioncatalyst/cli/api"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

func SchemaFromJsonAction(cCtx *cli.Context) error {
	jsonFilePath := cCtx.Args().Get(0)

	// Open the JSON outputFile
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Error opening JSON file: %v", err), 1)
	}
	defer jsonFile.Close()

	// Read the outputFile content
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
	apiResponse := apiClient.CallPublicConvertor(stringContent)

	// Unmarshal the raw JSON into an interface{}
	var data interface{}
	err = json.Unmarshal([]byte(apiResponse.Response), &data)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Error unmarshaling raw JSON: %v", err), 1)
	}

	// Marshal the interface{} back to JSON with indentation
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return cli.Exit(fmt.Sprintf("Error marshaling JSON: %v", err), 1)
	}

	// Write the JSON schema to an outputFile
	outputFile, err := os.Create(cCtx.Args().Get(1))
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create output file: %s", err), 1)
	}
	defer outputFile.Close()

	// Step 2: Write bytes to the outputFile
	_, err = outputFile.Write(prettyJSON)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to write to output file: %s", err), 1)
	}

	return nil
}
