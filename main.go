package main

import (
	"github.com/fusioncatalyst/cli/common"
	"log"
	"os"
)

func main() {
	app := common.GetAssembledApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
