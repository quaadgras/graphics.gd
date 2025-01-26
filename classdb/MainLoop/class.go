// Package MainLoop provides methods for working with MainLoop object instances.
package MainLoop

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any

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

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=MainLoop)
*/
type Instance [1]gdclass.MainLoop

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMainLoop() Instance
}
type Interface interface {
	//Called once during initialization.
	Initialize()
	//Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
	//If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
	PhysicsProcess(delta Float.X) bool
	//Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
	//If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
	Process(delta Float.X) bool
	//Called before the program exits.
	Finalize()
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) Initialize()                           { return }
func (self implementation) PhysicsProcess(delta Float.X) (_ bool) { return }
func (self implementation) Process(delta Float.X) (_ bool)        { return }
func (self implementation) Finalize()                             { return }

/*
Called once during initialization.
*/
func (Instance) _initialize(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
*/
func (Instance) _physics_process(impl func(ptr unsafe.Pointer, delta Float.X) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var delta = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, Float.X(delta))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
*/
func (Instance) _process(impl func(ptr unsafe.Pointer, delta Float.X) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var delta = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, Float.X(delta))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called before the program exits.
*/
func (Instance) _finalize(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MainLoop

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MainLoop"))
	casted := Instance{*(*gdclass.MainLoop)(unsafe.Pointer(&object))}
	return casted
}

/*
Called once during initialization.
*/
func (class) _initialize(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
*/
func (class) _physics_process(impl func(ptr unsafe.Pointer, delta gd.Float) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var delta = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, delta)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
*/
func (class) _process(impl func(ptr unsafe.Pointer, delta gd.Float) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var delta = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, delta)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called before the program exits.
*/
func (class) _finalize(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (self Instance) OnOnRequestPermissionsResult(cb func(permission string, granted bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("on_request_permissions_result"), gd.NewCallable(cb), 0)
}

func (self class) AsMainLoop() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMainLoop() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_physics_process":
		return reflect.ValueOf(self._physics_process)
	case "_process":
		return reflect.ValueOf(self._process)
	case "_finalize":
		return reflect.ValueOf(self._finalize)
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_physics_process":
		return reflect.ValueOf(self._physics_process)
	case "_process":
		return reflect.ValueOf(self._process)
	case "_finalize":
		return reflect.ValueOf(self._finalize)
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("MainLoop", func(ptr gd.Object) any { return [1]gdclass.MainLoop{*(*gdclass.MainLoop)(unsafe.Pointer(&ptr))} })
}
