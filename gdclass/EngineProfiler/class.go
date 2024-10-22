package EngineProfiler

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class can be used to implement custom profilers that are able to interact with the engine and editor debugger.
See [EngineDebugger] and [EditorDebuggerPlugin] for more information.
	// EngineProfiler methods that can be overridden by a [Class] that extends it.
	type EngineProfiler interface {
		//Called when the profiler is enabled/disabled, along with a set of [param options].
		Toggle(enable bool, options gd.Array) 
		//Called when data is added to profiler using [method EngineDebugger.profiler_add_frame_data].
		AddFrame(data gd.Array) 
		//Called once every engine iteration when the profiler is active with information about the current frame. All time values are in seconds. Lower values represent faster processing times and are therefore considered better.
		Tick(frame_time float64, process_time float64, physics_time float64, physics_frame_time float64) 
	}

*/
type Go [1]classdb.EngineProfiler

/*
Called when the profiler is enabled/disabled, along with a set of [param options].
*/
func (Go) _toggle(impl func(ptr unsafe.Pointer, enable bool, options gd.Array) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var enable = gd.UnsafeGet[bool](p_args,0)
		var options = mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enable, options)
		gc.End()
	}
}

/*
Called when data is added to profiler using [method EngineDebugger.profiler_add_frame_data].
*/
func (Go) _add_frame(impl func(ptr unsafe.Pointer, data gd.Array) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var data = mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, data)
		gc.End()
	}
}

/*
Called once every engine iteration when the profiler is active with information about the current frame. All time values are in seconds. Lower values represent faster processing times and are therefore considered better.
*/
func (Go) _tick(impl func(ptr unsafe.Pointer, frame_time float64, process_time float64, physics_time float64, physics_frame_time float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var frame_time = gd.UnsafeGet[gd.Float](p_args,0)
		var process_time = gd.UnsafeGet[gd.Float](p_args,1)
		var physics_time = gd.UnsafeGet[gd.Float](p_args,2)
		var physics_frame_time = gd.UnsafeGet[gd.Float](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(frame_time), float64(process_time), float64(physics_time), float64(physics_frame_time))
		gc.End()
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EngineProfiler
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EngineProfiler"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Called when the profiler is enabled/disabled, along with a set of [param options].
*/
func (class) _toggle(impl func(ptr unsafe.Pointer, enable bool, options gd.Array) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var enable = gd.UnsafeGet[bool](p_args,0)
		var options = mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enable, options)
		ctx.End()
	}
}

/*
Called when data is added to profiler using [method EngineDebugger.profiler_add_frame_data].
*/
func (class) _add_frame(impl func(ptr unsafe.Pointer, data gd.Array) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var data = mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, data)
		ctx.End()
	}
}

/*
Called once every engine iteration when the profiler is active with information about the current frame. All time values are in seconds. Lower values represent faster processing times and are therefore considered better.
*/
func (class) _tick(impl func(ptr unsafe.Pointer, frame_time gd.Float, process_time gd.Float, physics_time gd.Float, physics_frame_time gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var frame_time = gd.UnsafeGet[gd.Float](p_args,0)
		var process_time = gd.UnsafeGet[gd.Float](p_args,1)
		var physics_time = gd.UnsafeGet[gd.Float](p_args,2)
		var physics_frame_time = gd.UnsafeGet[gd.Float](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, frame_time, process_time, physics_time, physics_frame_time)
		ctx.End()
	}
}

func (self class) AsEngineProfiler() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEngineProfiler() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_toggle": return reflect.ValueOf(self._toggle);
	case "_add_frame": return reflect.ValueOf(self._add_frame);
	case "_tick": return reflect.ValueOf(self._tick);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_toggle": return reflect.ValueOf(self._toggle);
	case "_add_frame": return reflect.ValueOf(self._add_frame);
	case "_tick": return reflect.ValueOf(self._tick);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EngineProfiler", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
