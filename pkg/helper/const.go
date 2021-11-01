package helper

import (
	"os"
	"path/filepath"
)

var (
	ex, _    = os.Executable()
	Basepath = filepath.Dir(ex)
)
