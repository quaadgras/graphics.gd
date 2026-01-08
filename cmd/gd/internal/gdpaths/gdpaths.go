package gdpaths

import (
	"os"
	"os/user"
	"path/filepath"
)

var (
	Bin string
	Lib string
)

func init() {
	GDPATH := os.Getenv("GDPATH")
	whoami, err := user.Current()
	if GDPATH == "" && err == nil {
		GDPATH = filepath.Join(whoami.HomeDir, "gd")
	}
	Bin = filepath.Join(GDPATH, "bin")
	Lib = filepath.Join(GDPATH, "lib")
}
