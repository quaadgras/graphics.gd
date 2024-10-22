package MainLoop

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
[MainLoop] is the abstract base class for a Godot project's game loop. It is inherited by [SceneTree], which is the default game loop implementation used in Godot projects, though it is also possible to write and use one's own [MainLoop] subclass instead of the scene tree.
Upon the application start, a [MainLoop] implementation must be provided to the OS; otherwise, the application will exit. This happens automatically (and a [SceneTree] is created) unless a [MainLoop] [Script] is provided from the command line (with e.g. [code]godot -s my_loop.gd[/code]) or the "Main Loop Type" project setting is overwritten.
Here is an example script implementing a simple [MainLoop]:
[codeblocks]
[gdscript]
class_name CustomMainLoop
extends MainLoop

var time_elapsed = 0

func _initialize():
    print("Initialized:")
    print("  Starting time: %s" % str(time_elapsed))

func _process(delta):
    time_elapsed += delta
    # Return true to end the main loop.
    return Input.get_mouse_button_mask() != 0 || Input.is_key_pressed(KEY_ESCAPE)

func _finalize():
    print("Finalized:")
    print("  End time: %s" % str(time_elapsed))
[/gdscript]
[csharp]
using Godot;

[GlobalClass]
public partial class CustomMainLoop : MainLoop
{
    private double _timeElapsed = 0;

    public override void _Initialize()
    {
        GD.Print("Initialized:");
        GD.Print($"  Starting Time: {_timeElapsed}");
    }

    public override bool _Process(double delta)
    {
        _timeElapsed += delta;
        // Return true to end the main loop.
        return Input.GetMouseButtonMask() != 0 || Input.IsKeyPressed(Key.Escape);
    }

    private void _Finalize()
    {
        GD.Print("Finalized:");
        GD.Print($"  End Time: {_timeElapsed}");
    }
}
[/csharp]
[/codeblocks]
	// MainLoop methods that can be overridden by a [Class] that extends it.
	type MainLoop interface {
		//Called once during initialization.
		Initialize() 
		//Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
		//If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
		PhysicsProcess(delta float64) bool
		//Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
		//If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
		Process(delta float64) bool
		//Called before the program exits.
		Finalize() 
	}

*/
type Go [1]classdb.MainLoop

/*
Called once during initialization.
*/
func (Go) _initialize(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
*/
func (Go) _physics_process(impl func(ptr unsafe.Pointer, delta float64) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var delta = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, float64(delta))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
*/
func (Go) _process(impl func(ptr unsafe.Pointer, delta float64) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var delta = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, float64(delta))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Called before the program exits.
*/
func (Go) _finalize(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MainLoop
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("MainLoop"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Called once during initialization.
*/
func (class) _initialize(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
*/
func (class) _physics_process(impl func(ptr unsafe.Pointer, delta gd.Float) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var delta = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, delta)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
*/
func (class) _process(impl func(ptr unsafe.Pointer, delta gd.Float) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var delta = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, delta)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called before the program exits.
*/
func (class) _finalize(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (self Go) OnOnRequestPermissionsResult(cb func(permission string, granted bool)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("on_request_permissions_result"), gc.Callable(cb), 0)
}


func (self class) AsMainLoop() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMainLoop() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_initialize": return reflect.ValueOf(self._initialize);
	case "_physics_process": return reflect.ValueOf(self._physics_process);
	case "_process": return reflect.ValueOf(self._process);
	case "_finalize": return reflect.ValueOf(self._finalize);
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_initialize": return reflect.ValueOf(self._initialize);
	case "_physics_process": return reflect.ValueOf(self._physics_process);
	case "_process": return reflect.ValueOf(self._process);
	case "_finalize": return reflect.ValueOf(self._finalize);
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("MainLoop", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
