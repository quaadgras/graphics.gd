package upx

import (
	"runtime.link/api"
	"runtime.link/api/cmdl"
)

var CMD = api.Import[struct {
	Compress func(string) error `cmdl:"-1 %v"`
}](cmdl.API, "upx", nil)
