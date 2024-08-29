package utils

import (
	"bytes"
	"os"
)

func CaptureSucessfulClIActionOutput(f func(arguments []string) (err error), args []string) string {
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
