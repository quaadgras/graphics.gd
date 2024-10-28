package GraphNode

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/GraphElement"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[GraphNode] allows to create nodes for a [GraphEdit] graph with customizable content based on its child controls. [GraphNode] is derived from [Container] and it is responsible for placing its children on screen. This works similar to [VBoxContainer]. Children, in turn, provide [GraphNode] with so-called slots, each of which can have a connection port on either side.
Each [GraphNode] slot is defined by its index and can provide the node with up to two ports: one on the left, and one on the right. By convention the left port is also referred to as the [b]input port[/b] and the right port is referred to as the [b]output port[/b]. Each port can be enabled and configured individually, using different type and color. The type is an arbitrary value that you can define using your own considerations. The parent [GraphEdit] will receive this information on each connect and disconnect request.
Slots can be configured in the Inspector dock once you add at least one child [Control]. The properties are grouped by each slot's index in the "Slot" section.
[b]Note:[/b] While GraphNode is set up using slots and slot indices, connections are made between the ports which are enabled. Because of that [GraphEdit] uses the port's index and not the slot's index. You can use [method get_input_port_slot] and [method get_output_port_slot] to get the slot index from the port index.
	// GraphNode methods that can be overridden by a [Class] that extends it.
	type GraphNode interface {
		DrawPort(slot_index int, position gd.Vector2i, left bool, color gd.Color) 
	}

*/
type Go [1]classdb.GraphNode
func (Go) _draw_port(impl func(ptr unsafe.Pointer, slot_index int, position gd.Vector2i, left bool, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var slot_index = gd.UnsafeGet[gd.Int](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var left = gd.UnsafeGet[bool](p_args,2)
		var color = gd.UnsafeGet[gd.Color](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(slot_index), position, left, color)
	}
}

/*
Returns the [HBoxContainer] used for the title bar, only containing a [Label] for displaying the title by default. This can be used to add custom controls to the title bar such as option or close buttons.
*/
func (self Go) GetTitlebarHbox() gdclass.HBoxContainer {
	return gdclass.HBoxContainer(class(self).GetTitlebarHbox())
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
func (self Go) SetSlot(slot_index int, enable_left_port bool, type_left int, color_left gd.Color, enable_right_port bool, type_right int, color_right gd.Color) {
	class(self).SetSlot(gd.Int(slot_index), enable_left_port, gd.Int(type_left), color_left, enable_right_port, gd.Int(type_right), color_right, ([1]gdclass.Texture2D{}[0]), ([1]gdclass.Texture2D{}[0]), true)
}

/*
Disables the slot with the given [param slot_index]. This will remove the corresponding input and output port from the GraphNode.
*/
func (self Go) ClearSlot(slot_index int) {
	class(self).ClearSlot(gd.Int(slot_index))
}

/*
Disables all slots of the GraphNode. This will remove all input/output ports from the GraphNode.
*/
func (self Go) ClearAllSlots() {
	class(self).ClearAllSlots()
}

/*
Returns [code]true[/code] if left (input) side of the slot with the given [param slot_index] is enabled.
*/
func (self Go) IsSlotEnabledLeft(slot_index int) bool {
	return bool(class(self).IsSlotEnabledLeft(gd.Int(slot_index)))
}

/*
Toggles the left (input) side of the slot with the given [param slot_index]. If [param enable] is [code]true[/code], a port will appear on the left side and the slot will be able to be connected from this side.
*/
func (self Go) SetSlotEnabledLeft(slot_index int, enable bool) {
	class(self).SetSlotEnabledLeft(gd.Int(slot_index), enable)
}

/*
Sets the left (input) type of the slot with the given [param slot_index] to [param type]. If the value is negative, all connections will be disallowed to be created via user inputs.
*/
func (self Go) SetSlotTypeLeft(slot_index int, atype int) {
	class(self).SetSlotTypeLeft(gd.Int(slot_index), gd.Int(atype))
}

/*
Returns the left (input) type of the slot with the given [param slot_index].
*/
func (self Go) GetSlotTypeLeft(slot_index int) int {
	return int(int(class(self).GetSlotTypeLeft(gd.Int(slot_index))))
}

/*
Sets the [Color] of the left (input) side of the slot with the given [param slot_index] to [param color].
*/
func (self Go) SetSlotColorLeft(slot_index int, color gd.Color) {
	class(self).SetSlotColorLeft(gd.Int(slot_index), color)
}

/*
Returns the left (input) [Color] of the slot with the given [param slot_index].
*/
func (self Go) GetSlotColorLeft(slot_index int) gd.Color {
	return gd.Color(class(self).GetSlotColorLeft(gd.Int(slot_index)))
}

/*
Sets the custom [Texture2D] of the left (input) side of the slot with the given [param slot_index] to [param custom_icon].
*/
func (self Go) SetSlotCustomIconLeft(slot_index int, custom_icon gdclass.Texture2D) {
	class(self).SetSlotCustomIconLeft(gd.Int(slot_index), custom_icon)
}

/*
Returns the left (input) custom [Texture2D] of the slot with the given [param slot_index].
*/
func (self Go) GetSlotCustomIconLeft(slot_index int) gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetSlotCustomIconLeft(gd.Int(slot_index)))
}

/*
Returns [code]true[/code] if right (output) side of the slot with the given [param slot_index] is enabled.
*/
func (self Go) IsSlotEnabledRight(slot_index int) bool {
	return bool(class(self).IsSlotEnabledRight(gd.Int(slot_index)))
}

/*
Toggles the right (output) side of the slot with the given [param slot_index]. If [param enable] is [code]true[/code], a port will appear on the right side and the slot will be able to be connected from this side.
*/
func (self Go) SetSlotEnabledRight(slot_index int, enable bool) {
	class(self).SetSlotEnabledRight(gd.Int(slot_index), enable)
}

/*
Sets the right (output) type of the slot with the given [param slot_index] to [param type]. If the value is negative, all connections will be disallowed to be created via user inputs.
*/
func (self Go) SetSlotTypeRight(slot_index int, atype int) {
	class(self).SetSlotTypeRight(gd.Int(slot_index), gd.Int(atype))
}

/*
Returns the right (output) type of the slot with the given [param slot_index].
*/
func (self Go) GetSlotTypeRight(slot_index int) int {
	return int(int(class(self).GetSlotTypeRight(gd.Int(slot_index))))
}

/*
Sets the [Color] of the right (output) side of the slot with the given [param slot_index] to [param color].
*/
func (self Go) SetSlotColorRight(slot_index int, color gd.Color) {
	class(self).SetSlotColorRight(gd.Int(slot_index), color)
}

/*
Returns the right (output) [Color] of the slot with the given [param slot_index].
*/
func (self Go) GetSlotColorRight(slot_index int) gd.Color {
	return gd.Color(class(self).GetSlotColorRight(gd.Int(slot_index)))
}

/*
Sets the custom [Texture2D] of the right (output) side of the slot with the given [param slot_index] to [param custom_icon].
*/
func (self Go) SetSlotCustomIconRight(slot_index int, custom_icon gdclass.Texture2D) {
	class(self).SetSlotCustomIconRight(gd.Int(slot_index), custom_icon)
}

/*
Returns the right (output) custom [Texture2D] of the slot with the given [param slot_index].
*/
func (self Go) GetSlotCustomIconRight(slot_index int) gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetSlotCustomIconRight(gd.Int(slot_index)))
}

/*
Returns true if the background [StyleBox] of the slot with the given [param slot_index] is drawn.
*/
func (self Go) IsSlotDrawStylebox(slot_index int) bool {
	return bool(class(self).IsSlotDrawStylebox(gd.Int(slot_index)))
}

/*
Toggles the background [StyleBox] of the slot with the given [param slot_index].
*/
func (self Go) SetSlotDrawStylebox(slot_index int, enable bool) {
	class(self).SetSlotDrawStylebox(gd.Int(slot_index), enable)
}

/*
Returns the number of slots with an enabled input port.
*/
func (self Go) GetInputPortCount() int {
	return int(int(class(self).GetInputPortCount()))
}

/*
Returns the position of the input port with the given [param port_idx].
*/
func (self Go) GetInputPortPosition(port_idx int) gd.Vector2 {
	return gd.Vector2(class(self).GetInputPortPosition(gd.Int(port_idx)))
}

/*
Returns the type of the input port with the given [param port_idx].
*/
func (self Go) GetInputPortType(port_idx int) int {
	return int(int(class(self).GetInputPortType(gd.Int(port_idx))))
}

/*
Returns the [Color] of the input port with the given [param port_idx].
*/
func (self Go) GetInputPortColor(port_idx int) gd.Color {
	return gd.Color(class(self).GetInputPortColor(gd.Int(port_idx)))
}

/*
Returns the corresponding slot index of the input port with the given [param port_idx].
*/
func (self Go) GetInputPortSlot(port_idx int) int {
	return int(int(class(self).GetInputPortSlot(gd.Int(port_idx))))
}

/*
Returns the number of slots with an enabled output port.
*/
func (self Go) GetOutputPortCount() int {
	return int(int(class(self).GetOutputPortCount()))
}

/*
Returns the position of the output port with the given [param port_idx].
*/
func (self Go) GetOutputPortPosition(port_idx int) gd.Vector2 {
	return gd.Vector2(class(self).GetOutputPortPosition(gd.Int(port_idx)))
}

/*
Returns the type of the output port with the given [param port_idx].
*/
func (self Go) GetOutputPortType(port_idx int) int {
	return int(int(class(self).GetOutputPortType(gd.Int(port_idx))))
}

/*
Returns the [Color] of the output port with the given [param port_idx].
*/
func (self Go) GetOutputPortColor(port_idx int) gd.Color {
	return gd.Color(class(self).GetOutputPortColor(gd.Int(port_idx)))
}

/*
Returns the corresponding slot index of the output port with the given [param port_idx].
*/
func (self Go) GetOutputPortSlot(port_idx int) int {
	return int(int(class(self).GetOutputPortSlot(gd.Int(port_idx))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GraphNode
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GraphNode"))
	return Go{classdb.GraphNode(object)}
}

func (self Go) Title() string {
		return string(class(self).GetTitle().String())
}

func (self Go) SetTitle(value string) {
	class(self).SetTitle(gd.NewString(value))
}

func (self Go) IgnoreInvalidConnectionType() bool {
		return bool(class(self).IsIgnoringValidConnectionType())
}

func (self Go) SetIgnoreInvalidConnectionType(value bool) {
	class(self).SetIgnoreInvalidConnectionType(value)
}

func (class) _draw_port(impl func(ptr unsafe.Pointer, slot_index gd.Int, position gd.Vector2i, left bool, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var slot_index = gd.UnsafeGet[gd.Int](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var left = gd.UnsafeGet[bool](p_args,2)
		var color = gd.UnsafeGet[gd.Color](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, slot_index, position, left, color)
	}
}

//go:nosplit
func (self class) SetTitle(title gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(title))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTitle() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the [HBoxContainer] used for the title bar, only containing a [Label] for displaying the title by default. This can be used to add custom controls to the title bar such as option or close buttons.
*/
//go:nosplit
func (self class) GetTitlebarHbox() gdclass.HBoxContainer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_titlebar_hbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.HBoxContainer{classdb.HBoxContainer(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
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
func (self class) SetSlot(slot_index gd.Int, enable_left_port bool, type_left gd.Int, color_left gd.Color, enable_right_port bool, type_right gd.Int, color_right gd.Color, custom_icon_left gdclass.Texture2D, custom_icon_right gdclass.Texture2D, draw_stylebox bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable_left_port)
	callframe.Arg(frame, type_left)
	callframe.Arg(frame, color_left)
	callframe.Arg(frame, enable_right_port)
	callframe.Arg(frame, type_right)
	callframe.Arg(frame, color_right)
	callframe.Arg(frame, discreet.Get(custom_icon_left[0])[0])
	callframe.Arg(frame, discreet.Get(custom_icon_right[0])[0])
	callframe.Arg(frame, draw_stylebox)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disables the slot with the given [param slot_index]. This will remove the corresponding input and output port from the GraphNode.
*/
//go:nosplit
func (self class) ClearSlot(slot_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_clear_slot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disables all slots of the GraphNode. This will remove all input/output ports from the GraphNode.
*/
//go:nosplit
func (self class) ClearAllSlots()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_clear_all_slots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if left (input) side of the slot with the given [param slot_index] is enabled.
*/
//go:nosplit
func (self class) IsSlotEnabledLeft(slot_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_is_slot_enabled_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Toggles the left (input) side of the slot with the given [param slot_index]. If [param enable] is [code]true[/code], a port will appear on the left side and the slot will be able to be connected from this side.
*/
//go:nosplit
func (self class) SetSlotEnabledLeft(slot_index gd.Int, enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_enabled_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the left (input) type of the slot with the given [param slot_index] to [param type]. If the value is negative, all connections will be disallowed to be created via user inputs.
*/
//go:nosplit
func (self class) SetSlotTypeLeft(slot_index gd.Int, atype gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_type_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the left (input) type of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotTypeLeft(slot_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_type_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Color] of the left (input) side of the slot with the given [param slot_index] to [param color].
*/
//go:nosplit
func (self class) SetSlotColorLeft(slot_index gd.Int, color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_color_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the left (input) [Color] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotColorLeft(slot_index gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_color_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the custom [Texture2D] of the left (input) side of the slot with the given [param slot_index] to [param custom_icon].
*/
//go:nosplit
func (self class) SetSlotCustomIconLeft(slot_index gd.Int, custom_icon gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, discreet.Get(custom_icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_custom_icon_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the left (input) custom [Texture2D] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotCustomIconLeft(slot_index gd.Int) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_custom_icon_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if right (output) side of the slot with the given [param slot_index] is enabled.
*/
//go:nosplit
func (self class) IsSlotEnabledRight(slot_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_is_slot_enabled_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Toggles the right (output) side of the slot with the given [param slot_index]. If [param enable] is [code]true[/code], a port will appear on the right side and the slot will be able to be connected from this side.
*/
//go:nosplit
func (self class) SetSlotEnabledRight(slot_index gd.Int, enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_enabled_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the right (output) type of the slot with the given [param slot_index] to [param type]. If the value is negative, all connections will be disallowed to be created via user inputs.
*/
//go:nosplit
func (self class) SetSlotTypeRight(slot_index gd.Int, atype gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_type_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the right (output) type of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotTypeRight(slot_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_type_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Color] of the right (output) side of the slot with the given [param slot_index] to [param color].
*/
//go:nosplit
func (self class) SetSlotColorRight(slot_index gd.Int, color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_color_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the right (output) [Color] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotColorRight(slot_index gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_color_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the custom [Texture2D] of the right (output) side of the slot with the given [param slot_index] to [param custom_icon].
*/
//go:nosplit
func (self class) SetSlotCustomIconRight(slot_index gd.Int, custom_icon gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, discreet.Get(custom_icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_custom_icon_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the right (output) custom [Texture2D] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotCustomIconRight(slot_index gd.Int) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_slot_custom_icon_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns true if the background [StyleBox] of the slot with the given [param slot_index] is drawn.
*/
//go:nosplit
func (self class) IsSlotDrawStylebox(slot_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_is_slot_draw_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Toggles the background [StyleBox] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) SetSlotDrawStylebox(slot_index gd.Int, enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_slot_draw_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetIgnoreInvalidConnectionType(ignore bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_set_ignore_invalid_connection_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsIgnoringValidConnectionType() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_is_ignoring_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of slots with an enabled input port.
*/
//go:nosplit
func (self class) GetInputPortCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_input_port_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortPosition(port_idx gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_input_port_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the type of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortType(port_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_input_port_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Color] of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortColor(port_idx gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_input_port_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the corresponding slot index of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortSlot(port_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_input_port_slot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of slots with an enabled output port.
*/
//go:nosplit
func (self class) GetOutputPortCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_output_port_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortPosition(port_idx gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_output_port_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the type of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortType(port_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_output_port_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Color] of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortColor(port_idx gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_output_port_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the corresponding slot index of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortSlot(port_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphNode.Bind_get_output_port_slot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnSlotUpdated(cb func(slot_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("slot_updated"), gd.NewCallable(cb), 0)
}


func (self class) AsGraphNode() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGraphNode() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsGraphElement() GraphElement.GD { return *((*GraphElement.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGraphElement() GraphElement.Go { return *((*GraphElement.Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_draw_port": return reflect.ValueOf(self._draw_port);
	default: return gd.VirtualByName(self.AsGraphElement(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_draw_port": return reflect.ValueOf(self._draw_port);
	default: return gd.VirtualByName(self.AsGraphElement(), name)
	}
}
func init() {classdb.Register("GraphNode", func(ptr gd.Object) any { return classdb.GraphNode(ptr) })}
