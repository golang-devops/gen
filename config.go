package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/clipperhouse/typewriter"
)

type config struct {
	out        io.Writer
	customName string
	*typewriter.Config
}

var defaultConfig = config{
	out:        os.Stdout,
	customName: mustGetConfigCustomName(),
	Config:     &typewriter.Config{},
}

func mustGetConfigCustomName() string {
	homeDir := strings.TrimSpace(os.Getenv("HOME"))
	if homeDir == "" {
		panic("Home dir (env $HOME) is required")
	}
	return filepath.Join(homeDir, "_user_gen.go")
}

// keep in sync with imports.go
var stdImports = typewriter.NewImportSpecSet(
	typewriter.ImportSpec{Name: "_", Path: "github.com/clipperhouse/slice"},
	typewriter.ImportSpec{Name: "_", Path: "github.com/clipperhouse/stringer"},
)
