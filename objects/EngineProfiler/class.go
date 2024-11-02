package EngineProfiler

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/variant/Array"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This class can be used to implement custom profilers that are able to interact with the engine and editor debugger.
See [EngineDebugger] and [EditorDebuggerPlugin] for more information.

	// EngineProfiler methods that can be overridden by a [Class] that extends it.
	type EngineProfiler interface {
		//Called when the profiler is enabled/disabled, along with a set of [param options].
		Toggle(enable bool, options Array.Any)
		//Called when data is added to profiler using [method EngineDebugger.profiler_add_frame_data].
		AddFrame(data Array.Any)
		//Called once every engine iteration when the profiler is active with information about the current frame. All time values are in seconds. Lower values represent faster processing times and are therefore considered better.
		Tick(frame_time Float.X, process_time Float.X, physics_time Float.X, physics_frame_time Float.X)
	}
*/
type Instance [1]classdb.EngineProfiler

/*
Called when the profiler is enabled/disabled, along with a set of [param options].
*/
func (Instance) _toggle(impl func(ptr unsafe.Pointer, enable bool, options Array.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var enable = gd.UnsafeGet[bool](p_args, 0)
		var options = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(options)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, enable, options)
	}
}

/*
Called when data is added to profiler using [method EngineDebugger.profiler_add_frame_data].
*/
func (Instance) _add_frame(impl func(ptr unsafe.Pointer, data Array.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var data = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(data)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, data)
	}
}

/*
Called once every engine iteration when the profiler is active with information about the current frame. All time values are in seconds. Lower values represent faster processing times and are therefore considered better.
*/
func (Instance) _tick(impl func(ptr unsafe.Pointer, frame_time Float.X, process_time Float.X, physics_time Float.X, physics_frame_time Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var frame_time = gd.UnsafeGet[gd.Float](p_args, 0)
		var process_time = gd.UnsafeGet[gd.Float](p_args, 1)
		var physics_time = gd.UnsafeGet[gd.Float](p_args, 2)
		var physics_frame_time = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(frame_time), Float.X(process_time), Float.X(physics_time), Float.X(physics_frame_time))
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EngineProfiler

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EngineProfiler"))
	return Instance{classdb.EngineProfiler(object)}
}

/*
Called when the profiler is enabled/disabled, along with a set of [param options].
*/
func (class) _toggle(impl func(ptr unsafe.Pointer, enable bool, options gd.Array)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var enable = gd.UnsafeGet[bool](p_args, 0)
		var options = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, enable, options)
	}
}

/*
Called when data is added to profiler using [method EngineDebugger.profiler_add_frame_data].
*/
func (class) _add_frame(impl func(ptr unsafe.Pointer, data gd.Array)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var data = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, data)
	}
}

/*
Called once every engine iteration when the profiler is active with information about the current frame. All time values are in seconds. Lower values represent faster processing times and are therefore considered better.
*/
func (class) _tick(impl func(ptr unsafe.Pointer, frame_time gd.Float, process_time gd.Float, physics_time gd.Float, physics_frame_time gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var frame_time = gd.UnsafeGet[gd.Float](p_args, 0)
		var process_time = gd.UnsafeGet[gd.Float](p_args, 1)
		var physics_time = gd.UnsafeGet[gd.Float](p_args, 2)
		var physics_frame_time = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, frame_time, process_time, physics_time, physics_frame_time)
	}
}

func (self class) AsEngineProfiler() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEngineProfiler() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_toggle":
		return reflect.ValueOf(self._toggle)
	case "_add_frame":
		return reflect.ValueOf(self._add_frame)
	case "_tick":
		return reflect.ValueOf(self._tick)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_toggle":
		return reflect.ValueOf(self._toggle)
	case "_add_frame":
		return reflect.ValueOf(self._add_frame)
	case "_tick":
		return reflect.ValueOf(self._tick)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("EngineProfiler", func(ptr gd.Object) any { return classdb.EngineProfiler(ptr) })
}
