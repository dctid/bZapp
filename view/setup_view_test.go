package view

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func shutdown() {
}

func setup() {
	// Set the running dir to the project root to support reading test json files
	os.Chdir("..")
}

