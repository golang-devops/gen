package main

import (
	"io"
	"os"
	"os/user"
	"path/filepath"

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
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(curUser.HomeDir, "_user_gen.go")
}

// keep in sync with imports.go
var stdImports = typewriter.NewImportSpecSet(
	typewriter.ImportSpec{Name: "_", Path: "github.com/clipperhouse/slice"},
	typewriter.ImportSpec{Name: "_", Path: "github.com/clipperhouse/stringer"},
)
