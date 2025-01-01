package RichTextEffect

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A custom effect for a [RichTextLabel], which can be loaded in the [RichTextLabel] inspector or using [method RichTextLabel.install_effect].
[b]Note:[/b] For a [RichTextEffect] to be usable, a BBCode tag must be defined as a member variable called [code]bbcode[/code] in the script.
[codeblocks]
[gdscript skip-lint]
# The RichTextEffect will be usable like this: `[example]Some text[/example]`
var bbcode = "example"
[/gdscript]
[csharp skip-lint]
// The RichTextEffect will be usable like this: `[example]Some text[/example]`
string bbcode = "example";
[/csharp]
[/codeblocks]
[b]Note:[/b] As soon as a [RichTextLabel] contains at least one [RichTextEffect], it will continuously process the effect unless the project is paused. This may impact battery life negatively.

	// RichTextEffect methods that can be overridden by a [Class] that extends it.
	type RichTextEffect interface {
		//Override this method to modify properties in [param char_fx]. The method must return [code]true[/code] if the character could be transformed successfully. If the method returns [code]false[/code], it will skip transformation to avoid displaying broken text.
		ProcessCustomFx(char_fx objects.CharFXTransform) bool
	}
*/
type Instance [1]classdb.RichTextEffect
type Any interface {
	gd.IsClass
	AsRichTextEffect() Instance
}

/*
Override this method to modify properties in [param char_fx]. The method must return [code]true[/code] if the character could be transformed successfully. If the method returns [code]false[/code], it will skip transformation to avoid displaying broken text.
*/
func (Instance) _process_custom_fx(impl func(ptr unsafe.Pointer, char_fx objects.CharFXTransform) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var char_fx = objects.CharFXTransform{pointers.New[classdb.CharFXTransform]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(char_fx[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, char_fx)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RichTextEffect

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RichTextEffect"))
	return Instance{classdb.RichTextEffect(object)}
}

/*
Override this method to modify properties in [param char_fx]. The method must return [code]true[/code] if the character could be transformed successfully. If the method returns [code]false[/code], it will skip transformation to avoid displaying broken text.
*/
func (class) _process_custom_fx(impl func(ptr unsafe.Pointer, char_fx objects.CharFXTransform) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var char_fx = objects.CharFXTransform{pointers.New[classdb.CharFXTransform]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(char_fx[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, char_fx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsRichTextEffect() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRichTextEffect() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_process_custom_fx":
		return reflect.ValueOf(self._process_custom_fx)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_process_custom_fx":
		return reflect.ValueOf(self._process_custom_fx)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("RichTextEffect", func(ptr gd.Object) any { return [1]classdb.RichTextEffect{classdb.RichTextEffect(ptr)} })
}
