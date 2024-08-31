package tests

import (
	"bytes"
	"fmt"
	"github.com/fusioncatalyst/cli/common"
	"github.com/fusioncatalyst/cli/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestSchemaCRUD(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(".env file not found in current directory")
	}

	testProjectID, testProjectIDExists := os.LookupEnv("TEST_DATA_TEST_PROJECT_ID")
	if !testProjectIDExists {
		panic("TEST_DATA_TEST_PROJECT_ID not set")
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

	// Calculate the patches to the test files
	_, filename, _, _ := runtime.Caller(0) // get the current file's path
	testValidFilePath := filepath.Join(filepath.Dir(filename), "testfiles", "valid_json_schema_example.json")
	testInvalidFilePath := filepath.Join(filepath.Dir(filename), "testfiles", "invalid_json_schema_example.json")

	currentTimestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	newUniqueSchemaName := fmt.Sprintf("schema%s", currentTimestamp)

	// Test 1: create a new schema
	e := utils.CaptureSucessfulClIActionOutput(app.Run, []string{"cmd", "new-schema",
		"--file", testValidFilePath,
		"--project-id", testProjectID,
		"--schema-name", newUniqueSchemaName})

	// Test 2: an attempt to create a new schema with same name again
	e = utils.CaptureSucessfulClIActionOutput(app.Run, []string{"cmd", "new-schema",
		"--file", testValidFilePath,
		"--project-id", testProjectID,
		"--schema-name", newUniqueSchemaName})
	assert.Contains(t, e, "Server returned status: 409")

	currentTimestamp = strconv.FormatInt(time.Now().UnixNano(), 10)
	newUniqueSchemaNameForInvalidSchemaTest := fmt.Sprintf("schema%s", currentTimestamp)

	// Test 3: test invalid schema
	e = utils.CaptureSucessfulClIActionOutput(app.Run, []string{"cmd", "new-schema",
		"--file", testInvalidFilePath,
		"--project-id", testProjectID,
		"--schema-name", newUniqueSchemaNameForInvalidSchemaTest})
	assert.Contains(t, e, "Invalid JSON schema")

	newUniqueSchemaNameAgain := fmt.Sprintf("schema%s", currentTimestamp)

	// Test 4: create a new schema and return its ID
	e = utils.CaptureSucessfulClIActionOutput(app.Run, []string{"cmd", "new-schema",
		"--file", testValidFilePath,
		"--project-id", testProjectID,
		"--schema-name", newUniqueSchemaNameAgain,
		"--return-id"})
	assert.NotEmpty(t, e)
}
