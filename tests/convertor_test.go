package tests

import (
	"bytes"
	"encoding/json"
	"github.com/fusioncatalyst/cli/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"github.com/xeipuuv/gojsonschema"
	"log"
	"os"
	"testing"
)

var exitFunc = func(code int) {
	os.Exit(code)
}

func TestPublicToolsConvertor(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(".env file not found in current directory")
	}

	app := common.GetAssembledApp()

	// Capture the output
	var output bytes.Buffer
	var errOutput bytes.Buffer
	app.Writer = &output
	app.ErrWriter = &errOutput
	app.ExitErrHandler = func(cCtx *cli.Context, err error) {
		//fmt.Println(err)
	}

	cli.OsExiter = exitFunc

	// Test 1: missing file path
	err = app.Run([]string{"cmd", "schema-from-json"})
	assert.Contains(t, err.Error(), "Missing")

	// Test 2: missing output parameter
	err = app.Run([]string{"cmd", "schema-from-json", "./testfiles/invalid.json"})
	assert.Contains(t, err.Error(), "Missing output file argument")

	// Test 3: invalid file path
	err = app.Run([]string{"cmd", "schema-from-json", "./testfiles/invalid.json", "./testifles/outputs/1.json"})
	assert.Contains(t, err.Error(), "Error opening JSON file")

	// Test 4: invalid JSON content
	err = app.Run([]string{"cmd", "schema-from-json", "./testfiles/json_but_not_json.json", "./testifles/outputs/1.json"})
	assert.Contains(t, err.Error(), "Invalid JSON content")

	// Test 5: JSON content
	err = app.Run([]string{"cmd", "schema-from-json", "./testfiles/valid_json_1.json", "./testfiles/outputs/1.json"})

	// Validate the output JSON schema
	outputFilePath := "./testfiles/outputs/1.json"
	err = validateJSONSchema(outputFilePath)
	assert.NoError(t, err)
}

// validateJSONSchema validates the JSON schema in the specified file
func validateJSONSchema(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var schema map[string]interface{}
	err = json.Unmarshal(data, &schema)
	if err != nil {
		return err
	}

	// Use a JSON Schema Draft validator to validate the schema itself
	loader := gojsonschema.NewStringLoader(string(data))
	_, err = gojsonschema.NewSchema(loader)
	if err != nil {
		return err
	}

	return nil
}
