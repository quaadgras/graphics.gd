// Package GraphNode provides methods for working with GraphNode object instances.
package GraphNode

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/GraphElement"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector2i"

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
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
[GraphNode] allows to create nodes for a [GraphEdit] graph with customizable content based on its child controls. [GraphNode] is derived from [Container] and it is responsible for placing its children on screen. This works similar to [VBoxContainer]. Children, in turn, provide [GraphNode] with so-called slots, each of which can have a connection port on either side.
Each [GraphNode] slot is defined by its index and can provide the node with up to two ports: one on the left, and one on the right. By convention the left port is also referred to as the [b]input port[/b] and the right port is referred to as the [b]output port[/b]. Each port can be enabled and configured individually, using different type and color. The type is an arbitrary value that you can define using your own considerations. The parent [GraphEdit] will receive this information on each connect and disconnect request.
Slots can be configured in the Inspector dock once you add at least one child [Control]. The properties are grouped by each slot's index in the "Slot" section.
[b]Note:[/b] While GraphNode is set up using slots and slot indices, connections are made between the ports which are enabled. Because of that [GraphEdit] uses the port's index and not the slot's index. You can use [method get_input_port_slot] and [method get_output_port_slot] to get the slot index from the port index.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=GraphNode)
*/
type Instance [1]gdclass.GraphNode

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGraphNode() Instance
}
type Interface interface {
	DrawPort(slot_index int, position Vector2i.XY, left bool, color Color.RGBA)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) DrawPort(slot_index int, position Vector2i.XY, left bool, color Color.RGBA) {
	return
}
func (Instance) _draw_port(impl func(ptr unsafe.Pointer, slot_index int, position Vector2i.XY, left bool, color Color.RGBA)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var slot_index = gd.UnsafeGet[int64](p_args, 0)

		var position = gd.UnsafeGet[Vector2i.XY](p_args, 1)

		var left = gd.UnsafeGet[bool](p_args, 2)

		var color = gd.UnsafeGet[Color.RGBA](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(slot_index), position, left, color)
	}
}

/*
Returns the [HBoxContainer] used for the title bar, only containing a [Label] for displaying the title by default. This can be used to add custom controls to the title bar such as option or close buttons.
*/
func (self Instance) GetTitlebarHbox() [1]gdclass.HBoxContainer { //gd:GraphNode.get_titlebar_hbox
	return [1]gdclass.HBoxContainer(class(self).GetTitlebarHbox())
}

/*
Sets properties of the slot with the given [param slot_index].
If [param enable_left_port]/[param enable_right_port] is [code]true[/code], a port will appear and the slot will be able to be connected from this side.
With [param type_left]/[param type_right] an arbitrary type can be assigned to each port. Two ports can be connected if they share the same type, or if the connection between their types is allowed in the parent [GraphEdit] (see [method GraphEdit.add_valid_connection_type]). Keep in mind that the [GraphEdit] has the final say in accepting the connection. Type compatibility simply allows the [signal GraphEdit.connection_request] signal to be emitted.
Ports can be further customized using [param color_left]/[param color_right] and [param custom_icon_left]/[param custom_icon_right]. The color parameter adds a tint to the icon. The custom icon can be used to override the default port dot.
Additionally, [param draw_stylebox] can be used to enable or disable drawing of the background stylebox for each slot. See [theme_item slot].
Individual properties can also be set using one of the [code]set_slot_*[/code] methods.
[b]Note:[/b] This method only sets properties of the slot. To create the slot itself, add a [Control]-derived child to the GraphNode.
*/
func (self Instance) SetSlot(slot_index int, enable_left_port bool, type_left int, color_left Color.RGBA, enable_right_port bool, type_right int, color_right Color.RGBA) { //gd:GraphNode.set_slot
	class(self).SetSlot(int64(slot_index), enable_left_port, int64(type_left), Color.RGBA(color_left), enable_right_port, int64(type_right), Color.RGBA(color_right), [1][1]gdclass.Texture2D{}[0], [1][1]gdclass.Texture2D{}[0], true)
}

/*
Disables the slot with the given [param slot_index]. This will remove the corresponding input and output port from the GraphNode.
*/
func (self Instance) ClearSlot(slot_index int) { //gd:GraphNode.clear_slot
	class(self).ClearSlot(int64(slot_index))
}

/*
Disables all slots of the GraphNode. This will remove all input/output ports from the GraphNode.
*/
func (self Instance) ClearAllSlots() { //gd:GraphNode.clear_all_slots
	class(self).ClearAllSlots()
}

/*
Returns [code]true[/code] if left (input) side of the slot with the given [param slot_index] is enabled.
*/
func (self Instance) IsSlotEnabledLeft(slot_index int) bool { //gd:GraphNode.is_slot_enabled_left
	return bool(class(self).IsSlotEnabledLeft(int64(slot_index)))
}

/*
Toggles the left (input) side of the slot with the given [param slot_index]. If [param enable] is [code]true[/code], a port will appear on the left side and the slot will be able to be connected from this side.
*/
func (self Instance) SetSlotEnabledLeft(slot_index int, enable bool) { //gd:GraphNode.set_slot_enabled_left
	class(self).SetSlotEnabledLeft(int64(slot_index), enable)
}

/*
Sets the left (input) type of the slot with the given [param slot_index] to [param type]. If the value is negative, all connections will be disallowed to be created via user inputs.
*/
func (self Instance) SetSlotTypeLeft(slot_index int, atype int) { //gd:GraphNode.set_slot_type_left
	class(self).SetSlotTypeLeft(int64(slot_index), int64(atype))
}

/*
Returns the left (input) type of the slot with the given [param slot_index].
*/
func (self Instance) GetSlotTypeLeft(slot_index int) int { //gd:GraphNode.get_slot_type_left
	return int(int(class(self).GetSlotTypeLeft(int64(slot_index))))
}

/*
Sets the [Color] of the left (input) side of the slot with the given [param slot_index] to [param color].
*/
func (self Instance) SetSlotColorLeft(slot_index int, color Color.RGBA) { //gd:GraphNode.set_slot_color_left
	class(self).SetSlotColorLeft(int64(slot_index), Color.RGBA(color))
}

/*
Returns the left (input) [Color] of the slot with the given [param slot_index].
*/
func (self Instance) GetSlotColorLeft(slot_index int) Color.RGBA { //gd:GraphNode.get_slot_color_left
	return Color.RGBA(class(self).GetSlotColorLeft(int64(slot_index)))
}

/*
Sets the custom [Texture2D] of the left (input) side of the slot with the given [param slot_index] to [param custom_icon].
*/
func (self Instance) SetSlotCustomIconLeft(slot_index int, custom_icon [1]gdclass.Texture2D) { //gd:GraphNode.set_slot_custom_icon_left
	class(self).SetSlotCustomIconLeft(int64(slot_index), custom_icon)
}

/*
Returns the left (input) custom [Texture2D] of the slot with the given [param slot_index].
*/
func (self Instance) GetSlotCustomIconLeft(slot_index int) [1]gdclass.Texture2D { //gd:GraphNode.get_slot_custom_icon_left
	return [1]gdclass.Texture2D(class(self).GetSlotCustomIconLeft(int64(slot_index)))
}

/*
Returns [code]true[/code] if right (output) side of the slot with the given [param slot_index] is enabled.
*/
func (self Instance) IsSlotEnabledRight(slot_index int) bool { //gd:GraphNode.is_slot_enabled_right
	return bool(class(self).IsSlotEnabledRight(int64(slot_index)))
}

/*
Toggles the right (output) side of the slot with the given [param slot_index]. If [param enable] is [code]true[/code], a port will appear on the right side and the slot will be able to be connected from this side.
*/
func (self Instance) SetSlotEnabledRight(slot_index int, enable bool) { //gd:GraphNode.set_slot_enabled_right
	class(self).SetSlotEnabledRight(int64(slot_index), enable)
}

/*
Sets the right (output) type of the slot with the given [param slot_index] to [param type]. If the value is negative, all connections will be disallowed to be created via user inputs.
*/
func (self Instance) SetSlotTypeRight(slot_index int, atype int) { //gd:GraphNode.set_slot_type_right
	class(self).SetSlotTypeRight(int64(slot_index), int64(atype))
}

/*
Returns the right (output) type of the slot with the given [param slot_index].
*/
func (self Instance) GetSlotTypeRight(slot_index int) int { //gd:GraphNode.get_slot_type_right
	return int(int(class(self).GetSlotTypeRight(int64(slot_index))))
}

/*
Sets the [Color] of the right (output) side of the slot with the given [param slot_index] to [param color].
*/
func (self Instance) SetSlotColorRight(slot_index int, color Color.RGBA) { //gd:GraphNode.set_slot_color_right
	class(self).SetSlotColorRight(int64(slot_index), Color.RGBA(color))
}

/*
Returns the right (output) [Color] of the slot with the given [param slot_index].
*/
func (self Instance) GetSlotColorRight(slot_index int) Color.RGBA { //gd:GraphNode.get_slot_color_right
	return Color.RGBA(class(self).GetSlotColorRight(int64(slot_index)))
}

/*
Sets the custom [Texture2D] of the right (output) side of the slot with the given [param slot_index] to [param custom_icon].
*/
func (self Instance) SetSlotCustomIconRight(slot_index int, custom_icon [1]gdclass.Texture2D) { //gd:GraphNode.set_slot_custom_icon_right
	class(self).SetSlotCustomIconRight(int64(slot_index), custom_icon)
}

/*
Returns the right (output) custom [Texture2D] of the slot with the given [param slot_index].
*/
func (self Instance) GetSlotCustomIconRight(slot_index int) [1]gdclass.Texture2D { //gd:GraphNode.get_slot_custom_icon_right
	return [1]gdclass.Texture2D(class(self).GetSlotCustomIconRight(int64(slot_index)))
}

/*
Returns true if the background [StyleBox] of the slot with the given [param slot_index] is drawn.
*/
func (self Instance) IsSlotDrawStylebox(slot_index int) bool { //gd:GraphNode.is_slot_draw_stylebox
	return bool(class(self).IsSlotDrawStylebox(int64(slot_index)))
}

/*
Toggles the background [StyleBox] of the slot with the given [param slot_index].
*/
func (self Instance) SetSlotDrawStylebox(slot_index int, enable bool) { //gd:GraphNode.set_slot_draw_stylebox
	class(self).SetSlotDrawStylebox(int64(slot_index), enable)
}

/*
Returns the number of slots with an enabled input port.
*/
func (self Instance) GetInputPortCount() int { //gd:GraphNode.get_input_port_count
	return int(int(class(self).GetInputPortCount()))
}

/*
Returns the position of the input port with the given [param port_idx].
*/
func (self Instance) GetInputPortPosition(port_idx int) Vector2.XY { //gd:GraphNode.get_input_port_position
	return Vector2.XY(class(self).GetInputPortPosition(int64(port_idx)))
}

/*
Returns the type of the input port with the given [param port_idx].
*/
func (self Instance) GetInputPortType(port_idx int) int { //gd:GraphNode.get_input_port_type
	return int(int(class(self).GetInputPortType(int64(port_idx))))
}

/*
Returns the [Color] of the input port with the given [param port_idx].
*/
func (self Instance) GetInputPortColor(port_idx int) Color.RGBA { //gd:GraphNode.get_input_port_color
	return Color.RGBA(class(self).GetInputPortColor(int64(port_idx)))
}

/*
Returns the corresponding slot index of the input port with the given [param port_idx].
*/
func (self Instance) GetInputPortSlot(port_idx int) int { //gd:GraphNode.get_input_port_slot
	return int(int(class(self).GetInputPortSlot(int64(port_idx))))
}

/*
Returns the number of slots with an enabled output port.
*/
func (self Instance) GetOutputPortCount() int { //gd:GraphNode.get_output_port_count
	return int(int(class(self).GetOutputPortCount()))
}

/*
Returns the position of the output port with the given [param port_idx].
*/
func (self Instance) GetOutputPortPosition(port_idx int) Vector2.XY { //gd:GraphNode.get_output_port_position
	return Vector2.XY(class(self).GetOutputPortPosition(int64(port_idx)))
}

/*
Returns the type of the output port with the given [param port_idx].
*/
func (self Instance) GetOutputPortType(port_idx int) int { //gd:GraphNode.get_output_port_type
	return int(int(class(self).GetOutputPortType(int64(port_idx))))
}

/*
Returns the [Color] of the output port with the given [param port_idx].
*/
func (self Instance) GetOutputPortColor(port_idx int) Color.RGBA { //gd:GraphNode.get_output_port_color
	return Color.RGBA(class(self).GetOutputPortColor(int64(port_idx)))
}

/*
Returns the corresponding slot index of the output port with the given [param port_idx].
*/
func (self Instance) GetOutputPortSlot(port_idx int) int { //gd:GraphNode.get_output_port_slot
	return int(int(class(self).GetOutputPortSlot(int64(port_idx))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GraphNode

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GraphNode"))
	casted := Instance{*(*gdclass.GraphNode)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Title() string {
	return string(class(self).GetTitle().String())
}

func (self Instance) SetTitle(value string) {
	class(self).SetTitle(String.New(value))
}

func (self Instance) IgnoreInvalidConnectionType() bool {
	return bool(class(self).IsIgnoringValidConnectionType())
}

func (self Instance) SetIgnoreInvalidConnectionType(value bool) {
	class(self).SetIgnoreInvalidConnectionType(value)
}

func (class) _draw_port(impl func(ptr unsafe.Pointer, slot_index int64, position Vector2i.XY, left bool, color Color.RGBA)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var slot_index = gd.UnsafeGet[int64](p_args, 0)

		var position = gd.UnsafeGet[Vector2i.XY](p_args, 1)

		var left = gd.UnsafeGet[bool](p_args, 2)

		var color = gd.UnsafeGet[Color.RGBA](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, slot_index, position, left, color)
	}
}

//go:nosplit
func (self class) SetTitle(title String.Readable) { //gd:GraphNode.set_title
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTitle() String.Readable { //gd:GraphNode.get_title
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the [HBoxContainer] used for the title bar, only containing a [Label] for displaying the title by default. This can be used to add custom controls to the title bar such as option or close buttons.
*/
//go:nosplit
func (self class) GetTitlebarHbox() [1]gdclass.HBoxContainer { //gd:GraphNode.get_titlebar_hbox
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_titlebar_hbox, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.HBoxContainer{gd.PointerLifetimeBoundTo[gdclass.HBoxContainer](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets properties of the slot with the given [param slot_index].
If [param enable_left_port]/[param enable_right_port] is [code]true[/code], a port will appear and the slot will be able to be connected from this side.
With [param type_left]/[param type_right] an arbitrary type can be assigned to each port. Two ports can be connected if they share the same type, or if the connection between their types is allowed in the parent [GraphEdit] (see [method GraphEdit.add_valid_connection_type]). Keep in mind that the [GraphEdit] has the final say in accepting the connection. Type compatibility simply allows the [signal GraphEdit.connection_request] signal to be emitted.
Ports can be further customized using [param color_left]/[param color_right] and [param custom_icon_left]/[param custom_icon_right]. The color parameter adds a tint to the icon. The custom icon can be used to override the default port dot.
Additionally, [param draw_stylebox] can be used to enable or disable drawing of the background stylebox for each slot. See [theme_item slot].
Individual properties can also be set using one of the [code]set_slot_*[/code] methods.
[b]Note:[/b] This method only sets properties of the slot. To create the slot itself, add a [Control]-derived child to the GraphNode.
*/
//go:nosplit
func (self class) SetSlot(slot_index int64, enable_left_port bool, type_left int64, color_left Color.RGBA, enable_right_port bool, type_right int64, color_right Color.RGBA, custom_icon_left [1]gdclass.Texture2D, custom_icon_right [1]gdclass.Texture2D, draw_stylebox bool) { //gd:GraphNode.set_slot
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable_left_port)
	callframe.Arg(frame, type_left)
	callframe.Arg(frame, color_left)
	callframe.Arg(frame, enable_right_port)
	callframe.Arg(frame, type_right)
	callframe.Arg(frame, color_right)
	callframe.Arg(frame, pointers.Get(custom_icon_left[0])[0])
	callframe.Arg(frame, pointers.Get(custom_icon_right[0])[0])
	callframe.Arg(frame, draw_stylebox)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disables the slot with the given [param slot_index]. This will remove the corresponding input and output port from the GraphNode.
*/
//go:nosplit
func (self class) ClearSlot(slot_index int64) { //gd:GraphNode.clear_slot
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_clear_slot, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disables all slots of the GraphNode. This will remove all input/output ports from the GraphNode.
*/
//go:nosplit
func (self class) ClearAllSlots() { //gd:GraphNode.clear_all_slots
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_clear_all_slots, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if left (input) side of the slot with the given [param slot_index] is enabled.
*/
//go:nosplit
func (self class) IsSlotEnabledLeft(slot_index int64) bool { //gd:GraphNode.is_slot_enabled_left
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_is_slot_enabled_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Toggles the left (input) side of the slot with the given [param slot_index]. If [param enable] is [code]true[/code], a port will appear on the left side and the slot will be able to be connected from this side.
*/
//go:nosplit
func (self class) SetSlotEnabledLeft(slot_index int64, enable bool) { //gd:GraphNode.set_slot_enabled_left
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_enabled_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the left (input) type of the slot with the given [param slot_index] to [param type]. If the value is negative, all connections will be disallowed to be created via user inputs.
*/
//go:nosplit
func (self class) SetSlotTypeLeft(slot_index int64, atype int64) { //gd:GraphNode.set_slot_type_left
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_type_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the left (input) type of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotTypeLeft(slot_index int64) int64 { //gd:GraphNode.get_slot_type_left
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_type_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Color] of the left (input) side of the slot with the given [param slot_index] to [param color].
*/
//go:nosplit
func (self class) SetSlotColorLeft(slot_index int64, color Color.RGBA) { //gd:GraphNode.set_slot_color_left
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_color_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the left (input) [Color] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotColorLeft(slot_index int64) Color.RGBA { //gd:GraphNode.get_slot_color_left
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_color_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the custom [Texture2D] of the left (input) side of the slot with the given [param slot_index] to [param custom_icon].
*/
//go:nosplit
func (self class) SetSlotCustomIconLeft(slot_index int64, custom_icon [1]gdclass.Texture2D) { //gd:GraphNode.set_slot_custom_icon_left
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, pointers.Get(custom_icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_custom_icon_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the left (input) custom [Texture2D] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotCustomIconLeft(slot_index int64) [1]gdclass.Texture2D { //gd:GraphNode.get_slot_custom_icon_left
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_custom_icon_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if right (output) side of the slot with the given [param slot_index] is enabled.
*/
//go:nosplit
func (self class) IsSlotEnabledRight(slot_index int64) bool { //gd:GraphNode.is_slot_enabled_right
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_is_slot_enabled_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Toggles the right (output) side of the slot with the given [param slot_index]. If [param enable] is [code]true[/code], a port will appear on the right side and the slot will be able to be connected from this side.
*/
//go:nosplit
func (self class) SetSlotEnabledRight(slot_index int64, enable bool) { //gd:GraphNode.set_slot_enabled_right
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_enabled_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the right (output) type of the slot with the given [param slot_index] to [param type]. If the value is negative, all connections will be disallowed to be created via user inputs.
*/
//go:nosplit
func (self class) SetSlotTypeRight(slot_index int64, atype int64) { //gd:GraphNode.set_slot_type_right
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_type_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the right (output) type of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotTypeRight(slot_index int64) int64 { //gd:GraphNode.get_slot_type_right
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_type_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Color] of the right (output) side of the slot with the given [param slot_index] to [param color].
*/
//go:nosplit
func (self class) SetSlotColorRight(slot_index int64, color Color.RGBA) { //gd:GraphNode.set_slot_color_right
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_color_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the right (output) [Color] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotColorRight(slot_index int64) Color.RGBA { //gd:GraphNode.get_slot_color_right
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_color_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the custom [Texture2D] of the right (output) side of the slot with the given [param slot_index] to [param custom_icon].
*/
//go:nosplit
func (self class) SetSlotCustomIconRight(slot_index int64, custom_icon [1]gdclass.Texture2D) { //gd:GraphNode.set_slot_custom_icon_right
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, pointers.Get(custom_icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_custom_icon_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the right (output) custom [Texture2D] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotCustomIconRight(slot_index int64) [1]gdclass.Texture2D { //gd:GraphNode.get_slot_custom_icon_right
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_custom_icon_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns true if the background [StyleBox] of the slot with the given [param slot_index] is drawn.
*/
//go:nosplit
func (self class) IsSlotDrawStylebox(slot_index int64) bool { //gd:GraphNode.is_slot_draw_stylebox
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_is_slot_draw_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Toggles the background [StyleBox] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) SetSlotDrawStylebox(slot_index int64, enable bool) { //gd:GraphNode.set_slot_draw_stylebox
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_draw_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetIgnoreInvalidConnectionType(ignore bool) { //gd:GraphNode.set_ignore_invalid_connection_type
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_ignore_invalid_connection_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsIgnoringValidConnectionType() bool { //gd:GraphNode.is_ignoring_valid_connection_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_is_ignoring_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of slots with an enabled input port.
*/
//go:nosplit
func (self class) GetInputPortCount() int64 { //gd:GraphNode.get_input_port_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_input_port_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortPosition(port_idx int64) Vector2.XY { //gd:GraphNode.get_input_port_position
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_input_port_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the type of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortType(port_idx int64) int64 { //gd:GraphNode.get_input_port_type
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_input_port_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Color] of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortColor(port_idx int64) Color.RGBA { //gd:GraphNode.get_input_port_color
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_input_port_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the corresponding slot index of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortSlot(port_idx int64) int64 { //gd:GraphNode.get_input_port_slot
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_input_port_slot, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of slots with an enabled output port.
*/
//go:nosplit
func (self class) GetOutputPortCount() int64 { //gd:GraphNode.get_output_port_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_output_port_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortPosition(port_idx int64) Vector2.XY { //gd:GraphNode.get_output_port_position
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_output_port_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the type of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortType(port_idx int64) int64 { //gd:GraphNode.get_output_port_type
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_output_port_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Color] of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortColor(port_idx int64) Color.RGBA { //gd:GraphNode.get_output_port_color
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_output_port_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the corresponding slot index of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortSlot(port_idx int64) int64 { //gd:GraphNode.get_output_port_slot
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_output_port_slot, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnSlotUpdated(cb func(slot_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("slot_updated"), gd.NewCallable(cb), 0)
}

func (self class) AsGraphNode() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGraphNode() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGraphElement() GraphElement.Advanced {
	return *((*GraphElement.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGraphElement() GraphElement.Instance {
	return *((*GraphElement.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
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
	case "_draw_port":
		return reflect.ValueOf(self._draw_port)
	default:
		return gd.VirtualByName(GraphElement.Advanced(self.AsGraphElement()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_draw_port":
		return reflect.ValueOf(self._draw_port)
	default:
		return gd.VirtualByName(GraphElement.Instance(self.AsGraphElement()), name)
	}
}
func init() {
	gdclass.Register("GraphNode", func(ptr gd.Object) any { return [1]gdclass.GraphNode{*(*gdclass.GraphNode)(unsafe.Pointer(&ptr))} })
}
