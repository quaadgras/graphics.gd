//go:build !generate

package gd

import (
	"context"

	"runtime.link/mmm"
)

type godotContext = mmm.ContextWith[API]

type Context struct {
	godotContext
}

func NewContext(api *API) Context {
	return newContext(api)
}

func newContext(api *API) Context {
	var ctx Context
	ctx.godotContext = mmm.NewContextWith[API](context.Background(), api)
	return ctx
}

func (ctx Context) String(s string) String {
	var godot = ctx.API()
	var frame = godot.newFrame()
	ctx.API().Strings.New(frame.Back(), s)
	var str = frameGet[uintptr](frame)
	frame.free()
	return mmm.Make[API, String](ctx, godot, str)
}
