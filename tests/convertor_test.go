package tests

import (
	"bytes"
	"fmt"
	"github.com/fusioncatalyst/cli/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
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
		fmt.Println(err)
	}

	cli.OsExiter = exitFunc

	// Test 1: missing file path
	err = app.Run([]string{"cmd", "schema-from-json"})
	assert.Contains(t, err.Error(), "Missing")

	// Test 2: invalid file path
	err = app.Run([]string{"cmd", "schema-from-json", "./testfiles/invalid.json"})
	assert.Contains(t, err.Error(), "Error opening JSON file")

	// Test 3: invalid JSON content
	err = app.Run([]string{"cmd", "schema-from-json", "./testfiles/json_but_not_json.json"})
	assert.Contains(t, err.Error(), "Invalid JSON content")

	// Test 4: JSON content
	err = app.Run([]string{"cmd", "schema-from-json", "./testfiles/valid_json_1.json"})
	//assert.Contains(t, err.Error(), "Invalid JSON content")
}
