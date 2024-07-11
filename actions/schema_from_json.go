package actions

import (
	"fmt"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
)

func SchemaFromJsonAction(cCtx *cli.Context) error {
	if cCtx.NArg() == 0 {
		return cli.Exit("Missing JSON file argument", 1)
	}

	jsonFilePath := cCtx.Args().Get(0)
	fmt.Printf("Processing JSON schema from file: %s\n", jsonFilePath)
	// Add your logic to process the JSON file here

	// Open the JSON file
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()

	// Read the file content
	byteContent, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}
	stringContent := string(byteContent)
	if !utils.IsValidJSON(stringContent) {
		log.Fatalf("Invalid JSON content in the file: %s", jsonFilePath)
	}

	return nil

}
