package RichTextEffect

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
		ProcessCustomFx(char_fx gdclass.CharFXTransform) bool
	}

*/
type Go [1]classdb.RichTextEffect

/*
Override this method to modify properties in [param char_fx]. The method must return [code]true[/code] if the character could be transformed successfully. If the method returns [code]false[/code], it will skip transformation to avoid displaying broken text.
*/
func (Go) _process_custom_fx(impl func(ptr unsafe.Pointer, char_fx gdclass.CharFXTransform) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var char_fx gdclass.CharFXTransform
		char_fx[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, char_fx)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RichTextEffect
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RichTextEffect"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Override this method to modify properties in [param char_fx]. The method must return [code]true[/code] if the character could be transformed successfully. If the method returns [code]false[/code], it will skip transformation to avoid displaying broken text.
*/
func (class) _process_custom_fx(impl func(ptr unsafe.Pointer, char_fx gdclass.CharFXTransform) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var char_fx gdclass.CharFXTransform
		char_fx[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, char_fx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (self class) AsRichTextEffect() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRichTextEffect() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_process_custom_fx": return reflect.ValueOf(self._process_custom_fx);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_process_custom_fx": return reflect.ValueOf(self._process_custom_fx);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("RichTextEffect", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
