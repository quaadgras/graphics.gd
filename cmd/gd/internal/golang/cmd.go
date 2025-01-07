package golang

import (
	"runtime.link/api"
	"runtime.link/api/cmdl"
)

var CMD = api.Import[struct {
	Env struct {
		GOROOT func() string `cmdl:"env GOROOT"`
	}
}](cmdl.API, "go", nil)
