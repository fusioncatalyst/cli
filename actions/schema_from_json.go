package actions

import (
	"encoding/json"
	"fmt"
	"github.com/fusioncatalyst/cli/api"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
)

func SchemaFromJsonAction(cCtx *cli.Context) error {
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
	apiResponse := apiClient.CallPublicConvertor(stringContent)

	// Unmarshal the raw JSON into an interface{}
	var data interface{}
	err = json.Unmarshal([]byte(apiResponse.Response), &data)
	if err != nil {
		log.Fatalf("Error unmarshaling raw JSON: %v", err)
	}

	// Marshal the interface{} back to JSON with indentation
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	fmt.Println(string(prettyJSON))

	return nil
}
