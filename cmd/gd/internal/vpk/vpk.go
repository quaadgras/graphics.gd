package vpk

import (
	"runtime.link/api"
	"runtime.link/api/cmdl"
)

var CMD = api.Import[struct {
	Commands `cmdl:"%v"`

	Linux   Commands `cmdl:"[linux]"`
	Windows Commands `cmdl:"[win]"`
	OSX     Commands `cmdl:"[osx]"`
}](cmdl.API, "vpk", nil)

type Commands struct {
	Pack func(Package) error `cmdl:"pack %v"`
}

type Package struct {
	ID      string `cmdl:"--packId %v,omitempty"`
	Version string `cmdl:"--packVersion %v,omitempty"`
	Dir     string `cmdl:"--packDir %v,omitempty"`
	MainEXE string `cmdl:"--mainExe %v,omitempty"`
	Output  string `cmdl:"-o %v,omitempty"`
	Icon    string `cmdl:"--icon %v,omitempty"`
}
