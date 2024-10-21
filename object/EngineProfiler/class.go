package EngineProfiler

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
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
		Tick(frame_time gd.Float, process_time gd.Float, physics_time gd.Float, physics_frame_time gd.Float) 
	}

*/
type Simple [1]classdb.EngineProfiler
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EngineProfiler
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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


//go:nosplit
func (self class) AsEngineProfiler() Expert { return self[0].AsEngineProfiler() }


//go:nosplit
func (self Simple) AsEngineProfiler() Simple { return self[0].AsEngineProfiler() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_toggle": return reflect.ValueOf(self._toggle);
	case "_add_frame": return reflect.ValueOf(self._add_frame);
	case "_tick": return reflect.ValueOf(self._tick);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EngineProfiler", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
