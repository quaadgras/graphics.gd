package LinkButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/BaseButton"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A button that represents a link. This type of button is primarily used for interactions that cause a context change (like linking to a web page).
See also [BaseButton] which contains common properties and methods associated with this node.

*/
type Go [1]classdb.LinkButton
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.LinkButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("LinkButton"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Text() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetText(gc).String())
}

func (self Go) SetText(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetText(gc.String(value))
}

func (self Go) Underline() classdb.LinkButtonUnderlineMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.LinkButtonUnderlineMode(class(self).GetUnderlineMode())
}

func (self Go) SetUnderline(value classdb.LinkButtonUnderlineMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUnderlineMode(value)
}

func (self Go) Uri() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetUri(gc).String())
}

func (self Go) SetUri(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUri(gc.String(value))
}

func (self Go) TextDirection() classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.ControlTextDirection(class(self).GetTextDirection())
}

func (self Go) SetTextDirection(value classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextDirection(value)
}

func (self Go) Language() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetLanguage(gc).String())
}

func (self Go) SetLanguage(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLanguage(gc.String(value))
}

func (self Go) StructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride())
}

func (self Go) SetStructuredTextBidiOverride(value classdb.TextServerStructuredTextParser) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStructuredTextBidiOverride(value)
}

func (self Go) StructuredTextBidiOverrideOptions() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Array(class(self).GetStructuredTextBidiOverrideOptions(gc))
}

func (self Go) SetStructuredTextBidiOverrideOptions(value gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStructuredTextBidiOverrideOptions(value)
}

//go:nosplit
func (self class) SetText(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextDirection(direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextDirection() classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLanguage(language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLanguage(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUri(uri gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(uri))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_set_uri, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUri(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_get_uri, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUnderlineMode(underline_mode classdb.LinkButtonUnderlineMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, underline_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_set_underline_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUnderlineMode() classdb.LinkButtonUnderlineMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.LinkButtonUnderlineMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_get_underline_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser classdb.TextServerStructuredTextParser)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerStructuredTextParser](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(args gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(args))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LinkButton.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsLinkButton() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsLinkButton() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsBaseButton() BaseButton.GD { return *((*BaseButton.GD)(unsafe.Pointer(&self))) }
func (self Go) AsBaseButton() BaseButton.Go { return *((*BaseButton.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsBaseButton(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsBaseButton(), name)
	}
}
func init() {classdb.Register("LinkButton", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type UnderlineMode = classdb.LinkButtonUnderlineMode

const (
/*The LinkButton will always show an underline at the bottom of its text.*/
	UnderlineModeAlways UnderlineMode = 0
/*The LinkButton will show an underline at the bottom of its text when the mouse cursor is over it.*/
	UnderlineModeOnHover UnderlineMode = 1
/*The LinkButton will never show an underline at the bottom of its text.*/
	UnderlineModeNever UnderlineMode = 2
)
