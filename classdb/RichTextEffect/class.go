// Package RichTextEffect provides methods for working with RichTextEffect object instances.
package RichTextEffect

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=RichTextEffect)
*/
type Instance [1]gdclass.RichTextEffect

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRichTextEffect() Instance
}
type Interface interface {
	//Override this method to modify properties in [param char_fx]. The method must return [code]true[/code] if the character could be transformed successfully. If the method returns [code]false[/code], it will skip transformation to avoid displaying broken text.
	ProcessCustomFx(char_fx [1]gdclass.CharFXTransform) bool
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) ProcessCustomFx(char_fx [1]gdclass.CharFXTransform) (_ bool) { return }

/*
Override this method to modify properties in [param char_fx]. The method must return [code]true[/code] if the character could be transformed successfully. If the method returns [code]false[/code], it will skip transformation to avoid displaying broken text.
*/
func (Instance) _process_custom_fx(impl func(ptr unsafe.Pointer, char_fx [1]gdclass.CharFXTransform) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var char_fx = [1]gdclass.CharFXTransform{pointers.New[gdclass.CharFXTransform]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(char_fx[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, char_fx)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RichTextEffect

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RichTextEffect"))
	casted := Instance{*(*gdclass.RichTextEffect)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Override this method to modify properties in [param char_fx]. The method must return [code]true[/code] if the character could be transformed successfully. If the method returns [code]false[/code], it will skip transformation to avoid displaying broken text.
*/
func (class) _process_custom_fx(impl func(ptr unsafe.Pointer, char_fx [1]gdclass.CharFXTransform) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var char_fx = [1]gdclass.CharFXTransform{pointers.New[gdclass.CharFXTransform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_process_custom_fx":
		return reflect.ValueOf(self._process_custom_fx)
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_process_custom_fx":
		return reflect.ValueOf(self._process_custom_fx)
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("RichTextEffect", func(ptr gd.Object) any {
		return [1]gdclass.RichTextEffect{*(*gdclass.RichTextEffect)(unsafe.Pointer(&ptr))}
	})
}
