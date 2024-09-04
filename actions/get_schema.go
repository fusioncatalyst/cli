package actions

import (
	"encoding/json"
	"fmt"
	"github.com/fusioncatalyst/cli/api"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/urfave/cli/v2"
	"os"
)

func GetSchemaAction(cCtx *cli.Context) error {
	schemaID := cCtx.String("schema-id")
	toFile := cCtx.Bool("to-file")
	toFileWithName := cCtx.String("to-file-with-name")
	apiResponse := cCtx.Bool("api-response")

	apiClient := api.NewFCApiClient(utils.GetFCHost())
	result, e := apiClient.CallPrivateGetSchema(schemaID)

	if e != nil {
		return cli.Exit(e.Error(), 1)
	}

	schemaItself := utils.UnescapeJSONString(result.Schema)

	// If the user wants to write the schema to a file with default name, to do so
	if toFile || toFileWithName != "" {
		schemaStruct := utils.JSONStringToMap(schemaItself)
		beautifiedJSONSchema, err := json.MarshalIndent(schemaStruct, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
			return cli.Exit(e.Error(), 1)
		}

		filename := ""
		if toFile {
			filename = fmt.Sprintf("%s.json", result.Name)
		}

		if toFileWithName != "" {
			filename = toFileWithName
		}

		// Create a file to write the JSON data
		file, err := os.Create(filename)
		if err != nil {
			return cli.Exit(err, 1)
		}
		defer file.Close()

		_, err = file.Write(beautifiedJSONSchema)
		if err != nil {
			return cli.Exit(err, 1)
		}

		return nil

	}

	// If the user wants the API response, print it to the console
	if apiResponse {
		beautifiedAPIResponse, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
			return cli.Exit(e.Error(), 1)
		}

		fmt.Println(string(beautifiedAPIResponse))

		return nil
	}

	// Default option is to print the schema to the console
	beautifiedJSON, err := json.MarshalIndent(utils.JSONStringToMap(schemaItself), "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return cli.Exit(e.Error(), 1)
	}

	fmt.Println(string(beautifiedJSON))

	return nil
}
