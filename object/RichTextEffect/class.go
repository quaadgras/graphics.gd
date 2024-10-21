package RichTextEffect

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
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
		ProcessCustomFx(char_fx object.CharFXTransform) bool
	}

*/
type Simple [1]classdb.RichTextEffect
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RichTextEffect
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Override this method to modify properties in [param char_fx]. The method must return [code]true[/code] if the character could be transformed successfully. If the method returns [code]false[/code], it will skip transformation to avoid displaying broken text.
*/
func (class) _process_custom_fx(impl func(ptr unsafe.Pointer, char_fx object.CharFXTransform) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var char_fx object.CharFXTransform
		char_fx[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, char_fx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}


//go:nosplit
func (self class) AsRichTextEffect() Expert { return self[0].AsRichTextEffect() }


//go:nosplit
func (self Simple) AsRichTextEffect() Simple { return self[0].AsRichTextEffect() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_process_custom_fx": return reflect.ValueOf(self._process_custom_fx);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RichTextEffect", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
