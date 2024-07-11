package tests

import (
	"bytes"
	"fmt"
	"github.com/fusioncatalyst/cli/common"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"os"
	"testing"
)

var exitFunc = func(code int) {
	os.Exit(code)
}

func TestPublicToolsConvertor(t *testing.T) {
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

	// Run the test case
	err := app.Run([]string{"cmd", "schema-from-json"})
	fmt.Println(err)
	assert.Contains(t, err.Error(), "Missing")
}
