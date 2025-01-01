package Button

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/BaseButton"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
type Instance [1]classdb.Button
type Any interface {
	gd.IsClass
	AsButton() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Button

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Button"))
	return Instance{classdb.Button(object)}
}

func (self Instance) Text() string {
	return string(class(self).GetText().String())
}

func (self Instance) SetText(value string) {
	class(self).SetText(gd.NewString(value))
}

func (self Instance) Icon() objects.Texture2D {
	return objects.Texture2D(class(self).GetButtonIcon())
}

func (self Instance) SetIcon(value objects.Texture2D) {
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

func (self Instance) TextOverrunBehavior() classdb.TextServerOverrunBehavior {
	return classdb.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior())
}

func (self Instance) SetTextOverrunBehavior(value classdb.TextServerOverrunBehavior) {
	class(self).SetTextOverrunBehavior(value)
}

func (self Instance) AutowrapMode() classdb.TextServerAutowrapMode {
	return classdb.TextServerAutowrapMode(class(self).GetAutowrapMode())
}

func (self Instance) SetAutowrapMode(value classdb.TextServerAutowrapMode) {
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

func (self Instance) TextDirection() classdb.ControlTextDirection {
	return classdb.ControlTextDirection(class(self).GetTextDirection())
}

func (self Instance) SetTextDirection(value classdb.ControlTextDirection) {
	class(self).SetTextDirection(value)
}

func (self Instance) Language() string {
	return string(class(self).GetLanguage().String())
}

func (self Instance) SetLanguage(value string) {
	class(self).SetLanguage(gd.NewString(value))
}

//go:nosplit
func (self class) SetText(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior classdb.TextServerOverrunBehavior) {
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextOverrunBehavior() classdb.TextServerOverrunBehavior {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutowrapMode(autowrap_mode classdb.TextServerAutowrapMode) {
	var frame = callframe.New()
	callframe.Arg(frame, autowrap_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutowrapMode() classdb.TextServerAutowrapMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerAutowrapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextDirection(direction classdb.ControlTextDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextDirection() classdb.ControlTextDirection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetButtonIcon(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_button_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetButtonIcon() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_button_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlat(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_flat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlat() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_is_flat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClipText(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_clip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetClipText() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_clip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextAlignment(alignment HorizontalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_text_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextAlignment() HorizontalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_text_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIconAlignment(icon_alignment HorizontalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, icon_alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_icon_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetIconAlignment() HorizontalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_icon_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticalIconAlignment(vertical_icon_alignment VerticalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, vertical_icon_alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_vertical_icon_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticalIconAlignment() VerticalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[VerticalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_get_vertical_icon_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExpandIcon(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_set_expand_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsExpandIcon() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Button.Bind_is_expand_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
		return gd.VirtualByName(self.AsBaseButton(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsBaseButton(), name)
	}
}
func init() {
	classdb.Register("Button", func(ptr gd.Object) any { return [1]classdb.Button{classdb.Button(ptr)} })
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
