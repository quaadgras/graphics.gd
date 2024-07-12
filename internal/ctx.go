//go:build !generate

package gd

import (
	"runtime.link/mmm"
)

// Lifetime for ownership and a reference to the Godot API, apart from
// its use as an ordinary [context.Lifetime] to signal cancellation, this
// value is not safe to use concurrently. Each goroutine should create
// its own [Lifetime] and use that instead.
//
//	newctx := gd.NewLifetime(oldctx.API())
//
// When a [Lifetime] is freed, it will free all of the objects that were
// created using it. A [Lifetime] should not be used after free, as it
// will be recycled and will cause values to be unexpectedly freed.
//
// When the context has been passed in as a function argument, always
// assume that the [Lifetime] will be freed when the function returns.
// Classes can be moved between contexts using their [KeepAlive] method.
type Lifetime struct {
	mmm.Lifetime

	API *API
}

type Context = Lifetime

// NewContext returns a new [Lifetime] with the given API, the [Lifetime]
// will need to be freed manually by calling [Lifetime.Free].
func NewContext(api *API) Lifetime {
	return newContext(api)
}

// NewLifetime returns a new [Lifetime] with the given API, the [Lifetime]
// will need to be freed manually by calling [Lifetime.Free].
func NewLifetime(api *API) Lifetime {
	return newContext(api)
}

type Registrable interface {
	Register(Lifetime)
}

func (godot Lifetime) Register(spec Registrable) {
	spec.Register(godot)
}

func newContext(api *API) Lifetime {
	var ctx Lifetime
	ctx.Lifetime = mmm.NewLifetime()
	ctx.API = api
	return ctx
}

// Version returns the version of the Godot API that we are linked in to.
func (godot Lifetime) Version() Version {
	return godot.API.GetGodotVersion()
}

// GetLibraryPath returns the path to the library that was loaded.
func (godot Lifetime) GetLibraryPath() string {
	return godot.API.GetLibraryPath(godot, godot.API.ExtensionToken).String()
}

// String returns a [String] from a standard UTF8 Go string.
func (godot Lifetime) String(s string) String {
	return godot.API.Strings.New(godot, s)
}

// StringName returns a [StringName] from a standard UTF8 Go string.
func (godot Lifetime) StringName(s string) StringName {
	return godot.API.StringNames.New(godot, s)
}

func (godot Lifetime) Free() {
	godot.Lifetime.End()
}
