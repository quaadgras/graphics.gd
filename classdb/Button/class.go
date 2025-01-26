// Package Button provides methods for working with Button object instances.
package Button

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
import "graphics.gd/classdb/BaseButton"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

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
[Button] is the standard themed button. It can contain text and an icon, and it will display them according to the current [Theme].
[b]Example of creating a button and assigning an action when pressed by code:[/b]
[codeblocks]
[gdscript]
func _ready():

	var button = Button.new()
	button.text = "Click me"
	button.pressed.connect(self._button_pressed)
	add_child(button)

func _button_pressed():

	print("Hello world!")

[/gdscript]
[csharp]
public override void _Ready()

	{
	    var button = new Button();
	    button.Text = "Click me";
	    button.Pressed += ButtonPressed;
	    AddChild(button);
	}

private void ButtonPressed()

	{
	    GD.Print("Hello world!");
	}

[/csharp]
[/codeblocks]
See also [BaseButton] which contains common properties and methods associated with this node.
[b]Note:[/b] Buttons do not interpret touch input and therefore don't support multitouch, since mouse emulation can only press one button at a given time. Use [TouchScreenButton] for buttons that trigger gameplay movement or actions.
*/
type Instance [1]gdclass.Button

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsButton() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Button

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Button"))
	casted := Instance{*(*gdclass.Button)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Text() string {
	return string(class(self).GetText().String())
}

func (self Instance) SetText(value string) {
	class(self).SetText(gd.NewString(value))
}

func (self Instance) Icon() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetButtonIcon())
}

func (self Instance) SetIcon(value [1]gdclass.Texture2D) {
	class(self).SetButtonIcon(value)
}

func (self Instance) Flat() bool {
	return bool(class(self).IsFlat())
}

func (self Instance) SetFlat(value bool) {
	class(self).SetFlat(value)
}

func (self Instance) Alignment() HorizontalAlignment {
	return HorizontalAlignment(class(self).GetTextAlignment())
}

func (self Instance) SetAlignment(value HorizontalAlignment) {
	class(self).SetTextAlignment(value)
}

func (self Instance) TextOverrunBehavior() gdclass.TextServerOverrunBehavior {
	return gdclass.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior())
}

func (self Instance) SetTextOverrunBehavior(value gdclass.TextServerOverrunBehavior) {
	class(self).SetTextOverrunBehavior(value)
}

func (self Instance) AutowrapMode() gdclass.TextServerAutowrapMode {
	return gdclass.TextServerAutowrapMode(class(self).GetAutowrapMode())
}

func (self Instance) SetAutowrapMode(value gdclass.TextServerAutowrapMode) {
	class(self).SetAutowrapMode(value)
}

func (self Instance) ClipText() bool {
	return bool(class(self).GetClipText())
}

func (self Instance) SetClipText(value bool) {
	class(self).SetClipText(value)
}

func (self Instance) IconAlignment() HorizontalAlignment {
	return HorizontalAlignment(class(self).GetIconAlignment())
}

func (self Instance) SetIconAlignment(value HorizontalAlignment) {
	class(self).SetIconAlignment(value)
}

func (self Instance) VerticalIconAlignment() VerticalAlignment {
	return VerticalAlignment(class(self).GetVerticalIconAlignment())
}

func (self Instance) SetVerticalIconAlignment(value VerticalAlignment) {
	class(self).SetVerticalIconAlignment(value)
}

func (self Instance) ExpandIcon() bool {
	return bool(class(self).IsExpandIcon())
}

func (self Instance) SetExpandIcon(value bool) {
	class(self).SetExpandIcon(value)
}

func (self Instance) TextDirection() gdclass.ControlTextDirection {
	return gdclass.ControlTextDirection(class(self).GetTextDirection())
}

func (self Instance) SetTextDirection(value gdclass.ControlTextDirection) {
	class(self).SetTextDirection(value)
}

func (self Instance) Language() string {
	return string(class(self).GetLanguage().String())
}

func (self Instance) SetLanguage(value string) {
	class(self).SetLanguage(gd.NewString(value))
}

//go:nosplit
func (self class) SetText(text gd.String) { //gd:Button.set_text
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetText() gd.String { //gd:Button.get_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior gdclass.TextServerOverrunBehavior) { //gd:Button.set_text_overrun_behavior
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextOverrunBehavior() gdclass.TextServerOverrunBehavior { //gd:Button.get_text_overrun_behavior
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutowrapMode(autowrap_mode gdclass.TextServerAutowrapMode) { //gd:Button.set_autowrap_mode
	var frame = callframe.New()
	callframe.Arg(frame, autowrap_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutowrapMode() gdclass.TextServerAutowrapMode { //gd:Button.get_autowrap_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerAutowrapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextDirection(direction gdclass.ControlTextDirection) { //gd:Button.set_text_direction
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextDirection() gdclass.ControlTextDirection { //gd:Button.get_text_direction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language gd.String) { //gd:Button.set_language
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() gd.String { //gd:Button.get_language
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetButtonIcon(texture [1]gdclass.Texture2D) { //gd:Button.set_button_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_button_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetButtonIcon() [1]gdclass.Texture2D { //gd:Button.get_button_icon
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_button_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlat(enabled bool) { //gd:Button.set_flat
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_flat, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlat() bool { //gd:Button.is_flat
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_is_flat, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClipText(enabled bool) { //gd:Button.set_clip_text
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_clip_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetClipText() bool { //gd:Button.get_clip_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_clip_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextAlignment(alignment HorizontalAlignment) { //gd:Button.set_text_alignment
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_text_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextAlignment() HorizontalAlignment { //gd:Button.get_text_alignment
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_text_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIconAlignment(icon_alignment HorizontalAlignment) { //gd:Button.set_icon_alignment
	var frame = callframe.New()
	callframe.Arg(frame, icon_alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_icon_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIconAlignment() HorizontalAlignment { //gd:Button.get_icon_alignment
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_icon_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticalIconAlignment(vertical_icon_alignment VerticalAlignment) { //gd:Button.set_vertical_icon_alignment
	var frame = callframe.New()
	callframe.Arg(frame, vertical_icon_alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_vertical_icon_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticalIconAlignment() VerticalAlignment { //gd:Button.get_vertical_icon_alignment
	var frame = callframe.New()
	var r_ret = callframe.Ret[VerticalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_vertical_icon_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExpandIcon(enabled bool) { //gd:Button.set_expand_icon
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_expand_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsExpandIcon() bool { //gd:Button.is_expand_icon
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_is_expand_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsButton() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsButton() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsBaseButton() BaseButton.Advanced {
	return *((*BaseButton.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBaseButton() BaseButton.Instance {
	return *((*BaseButton.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(BaseButton.Advanced(self.AsBaseButton()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(BaseButton.Instance(self.AsBaseButton()), name)
	}
}
func init() {
	gdclass.Register("Button", func(ptr gd.Object) any { return [1]gdclass.Button{*(*gdclass.Button)(unsafe.Pointer(&ptr))} })
}

type HorizontalAlignment int

const (
	/*Horizontal left alignment, usually for text-derived classes.*/
	HorizontalAlignmentLeft HorizontalAlignment = 0
	/*Horizontal center alignment, usually for text-derived classes.*/
	HorizontalAlignmentCenter HorizontalAlignment = 1
	/*Horizontal right alignment, usually for text-derived classes.*/
	HorizontalAlignmentRight HorizontalAlignment = 2
	/*Expand row to fit width, usually for text-derived classes.*/
	HorizontalAlignmentFill HorizontalAlignment = 3
)

type VerticalAlignment int

const (
	/*Vertical top alignment, usually for text-derived classes.*/
	VerticalAlignmentTop VerticalAlignment = 0
	/*Vertical center alignment, usually for text-derived classes.*/
	VerticalAlignmentCenter VerticalAlignment = 1
	/*Vertical bottom alignment, usually for text-derived classes.*/
	VerticalAlignmentBottom VerticalAlignment = 2
	/*Expand rows to fit height, usually for text-derived classes.*/
	VerticalAlignmentFill VerticalAlignment = 3
)
