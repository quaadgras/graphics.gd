//go:build !generate

package gd

import (
	"context"
	"reflect"
	"strings"
	"unsafe"

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

// Create a new instance of the given class, which should be an uninitialised
// pointer to a value of that class. T must be a class from this package.
func Create[T PointerToClass](ctx Context, ptr T) T {
	var godot = ctx.API()
	var sname = ctx.StringName(strings.TrimPrefix(reflect.TypeOf(ptr).Elem().Name(), "class"))
	var fresh = godot.ClassDB.CreateObject((StringNamePtr)(unsafe.Pointer(&sname)))
	ptr.SetPointer(mmm.Make[API, Pointer](ctx, godot, fresh))
	return ptr
}

// String returns a [String] from a standard UTF8 Go string.
func (ctx Context) String(s string) String {
	var godot = ctx.API()
	var frame = godot.newFrame()
	ctx.API().Strings.New(frame.Back(), s)
	var str = frameGet[uintptr](frame)
	frame.free()
	return mmm.Make[API, String](ctx, godot, str)
}

// StringName returns a [StringName] from a standard UTF8 Go string.
func (ctx Context) StringName(s string) StringName {
	var godot = ctx.API()
	var frame = godot.newFrame()
	ctx.API().StringNames.New(frame.Back(), s)
	var str = frameGet[uintptr](frame)
	frame.free()
	return mmm.Make[API, StringName](ctx, godot, str)
}

// Sin returns the sine of the given angle.
func (ctx Context) Sin(x Float) Float {
	var godot = ctx.API()
	var frame = godot.newFrame()
	frameSet[Float](0, frame, x)
	godot.utility.sin(frame.Back(), frame.Args(), 1)
	var ret = frameGet[Float](frame)
	frame.free()
	return ret
}
