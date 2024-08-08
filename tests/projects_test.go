package tests

import (
	"bytes"
	"fmt"
	"github.com/fusioncatalyst/cli/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestProjectsCRUD(t *testing.T) {
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

	currentTimestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	newUniqueProjectName := fmt.Sprintf("testproject%s", currentTimestamp)

	// Test 1: create a new project
	err = app.Run([]string{"cmd", "create-project", newUniqueProjectName})
	assert.Nil(t, output.Bytes())

	// Test 2: trying to create project with exactly same name
	err = app.Run([]string{"cmd", "create-project", newUniqueProjectName})
	assert.Contains(t, err.Error(), "Error making fucioncatalyst server request. Server returned status: 409")

	// Test 3: list projects
	captureOutput := captureSucessfullClIActionOutput(app.Run, []string{"cmd", "list-projects"})
	assert.Contains(t, captureOutput, newUniqueProjectName)
}

func captureSucessfullClIActionOutput(f func(arguments []string) (err error), args []string) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f(args)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}
