package GraphNode

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/GraphElement"
import "grow.graphics/gd/object/Container"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[GraphNode] allows to create nodes for a [GraphEdit] graph with customizable content based on its child controls. [GraphNode] is derived from [Container] and it is responsible for placing its children on screen. This works similar to [VBoxContainer]. Children, in turn, provide [GraphNode] with so-called slots, each of which can have a connection port on either side.
Each [GraphNode] slot is defined by its index and can provide the node with up to two ports: one on the left, and one on the right. By convention the left port is also referred to as the [b]input port[/b] and the right port is referred to as the [b]output port[/b]. Each port can be enabled and configured individually, using different type and color. The type is an arbitrary value that you can define using your own considerations. The parent [GraphEdit] will receive this information on each connect and disconnect request.
Slots can be configured in the Inspector dock once you add at least one child [Control]. The properties are grouped by each slot's index in the "Slot" section.
[b]Note:[/b] While GraphNode is set up using slots and slot indices, connections are made between the ports which are enabled. Because of that [GraphEdit] uses the port's index and not the slot's index. You can use [method get_input_port_slot] and [method get_output_port_slot] to get the slot index from the port index.
	// GraphNode methods that can be overridden by a [Class] that extends it.
	type GraphNode interface {
		DrawPort(slot_index gd.Int, position gd.Vector2i, left bool, color gd.Color) 
	}

*/
type Simple [1]classdb.GraphNode
func (self Simple) SetTitle(title string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTitle(gc.String(title))
}
func (self Simple) GetTitle() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTitle(gc).String())
}
func (self Simple) GetTitlebarHbox() [1]classdb.HBoxContainer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.HBoxContainer(Expert(self).GetTitlebarHbox(gc))
}
func (self Simple) SetSlot(slot_index int, enable_left_port bool, type_left int, color_left gd.Color, enable_right_port bool, type_right int, color_right gd.Color, custom_icon_left [1]classdb.Texture2D, custom_icon_right [1]classdb.Texture2D, draw_stylebox bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlot(gd.Int(slot_index), enable_left_port, gd.Int(type_left), color_left, enable_right_port, gd.Int(type_right), color_right, custom_icon_left, custom_icon_right, draw_stylebox)
}
func (self Simple) ClearSlot(slot_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearSlot(gd.Int(slot_index))
}
func (self Simple) ClearAllSlots() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearAllSlots()
}
func (self Simple) IsSlotEnabledLeft(slot_index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSlotEnabledLeft(gd.Int(slot_index)))
}
func (self Simple) SetSlotEnabledLeft(slot_index int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlotEnabledLeft(gd.Int(slot_index), enable)
}
func (self Simple) SetSlotTypeLeft(slot_index int, atype int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlotTypeLeft(gd.Int(slot_index), gd.Int(atype))
}
func (self Simple) GetSlotTypeLeft(slot_index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSlotTypeLeft(gd.Int(slot_index))))
}
func (self Simple) SetSlotColorLeft(slot_index int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlotColorLeft(gd.Int(slot_index), color)
}
func (self Simple) GetSlotColorLeft(slot_index int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetSlotColorLeft(gd.Int(slot_index)))
}
func (self Simple) SetSlotCustomIconLeft(slot_index int, custom_icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlotCustomIconLeft(gd.Int(slot_index), custom_icon)
}
func (self Simple) GetSlotCustomIconLeft(slot_index int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetSlotCustomIconLeft(gc, gd.Int(slot_index)))
}
func (self Simple) IsSlotEnabledRight(slot_index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSlotEnabledRight(gd.Int(slot_index)))
}
func (self Simple) SetSlotEnabledRight(slot_index int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlotEnabledRight(gd.Int(slot_index), enable)
}
func (self Simple) SetSlotTypeRight(slot_index int, atype int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlotTypeRight(gd.Int(slot_index), gd.Int(atype))
}
func (self Simple) GetSlotTypeRight(slot_index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSlotTypeRight(gd.Int(slot_index))))
}
func (self Simple) SetSlotColorRight(slot_index int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlotColorRight(gd.Int(slot_index), color)
}
func (self Simple) GetSlotColorRight(slot_index int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetSlotColorRight(gd.Int(slot_index)))
}
func (self Simple) SetSlotCustomIconRight(slot_index int, custom_icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlotCustomIconRight(gd.Int(slot_index), custom_icon)
}
func (self Simple) GetSlotCustomIconRight(slot_index int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetSlotCustomIconRight(gc, gd.Int(slot_index)))
}
func (self Simple) IsSlotDrawStylebox(slot_index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSlotDrawStylebox(gd.Int(slot_index)))
}
func (self Simple) SetSlotDrawStylebox(slot_index int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlotDrawStylebox(gd.Int(slot_index), enable)
}
func (self Simple) SetIgnoreInvalidConnectionType(ignore bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIgnoreInvalidConnectionType(ignore)
}
func (self Simple) IsIgnoringValidConnectionType() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsIgnoringValidConnectionType())
}
func (self Simple) GetInputPortCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInputPortCount()))
}
func (self Simple) GetInputPortPosition(port_idx int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetInputPortPosition(gd.Int(port_idx)))
}
func (self Simple) GetInputPortType(port_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInputPortType(gd.Int(port_idx))))
}
func (self Simple) GetInputPortColor(port_idx int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetInputPortColor(gd.Int(port_idx)))
}
func (self Simple) GetInputPortSlot(port_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInputPortSlot(gd.Int(port_idx))))
}
func (self Simple) GetOutputPortCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOutputPortCount()))
}
func (self Simple) GetOutputPortPosition(port_idx int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetOutputPortPosition(gd.Int(port_idx)))
}
func (self Simple) GetOutputPortType(port_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOutputPortType(gd.Int(port_idx))))
}
func (self Simple) GetOutputPortColor(port_idx int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetOutputPortColor(gd.Int(port_idx)))
}
func (self Simple) GetOutputPortSlot(port_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOutputPortSlot(gd.Int(port_idx))))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GraphNode
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func (class) _draw_port(impl func(ptr unsafe.Pointer, slot_index gd.Int, position gd.Vector2i, left bool, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var slot_index = gd.UnsafeGet[gd.Int](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var left = gd.UnsafeGet[bool](p_args,2)
		var color = gd.UnsafeGet[gd.Color](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, slot_index, position, left, color)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetTitle(title gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTitle(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the [HBoxContainer] used for the title bar, only containing a [Label] for displaying the title by default. This can be used to add custom controls to the title bar such as option or close buttons.
*/
//go:nosplit
func (self class) GetTitlebarHbox(ctx gd.Lifetime) object.HBoxContainer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_titlebar_hbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.HBoxContainer
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
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
func (self class) SetSlot(slot_index gd.Int, enable_left_port bool, type_left gd.Int, color_left gd.Color, enable_right_port bool, type_right gd.Int, color_right gd.Color, custom_icon_left object.Texture2D, custom_icon_right object.Texture2D, draw_stylebox bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable_left_port)
	callframe.Arg(frame, type_left)
	callframe.Arg(frame, color_left)
	callframe.Arg(frame, enable_right_port)
	callframe.Arg(frame, type_right)
	callframe.Arg(frame, color_right)
	callframe.Arg(frame, mmm.Get(custom_icon_left[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(custom_icon_right[0].AsPointer())[0])
	callframe.Arg(frame, draw_stylebox)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_slot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disables the slot with the given [param slot_index]. This will remove the corresponding input and output port from the GraphNode.
*/
//go:nosplit
func (self class) ClearSlot(slot_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_clear_slot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disables all slots of the GraphNode. This will remove all input/output ports from the GraphNode.
*/
//go:nosplit
func (self class) ClearAllSlots()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_clear_all_slots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if left (input) side of the slot with the given [param slot_index] is enabled.
*/
//go:nosplit
func (self class) IsSlotEnabledLeft(slot_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_is_slot_enabled_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Toggles the left (input) side of the slot with the given [param slot_index]. If [param enable] is [code]true[/code], a port will appear on the left side and the slot will be able to be connected from this side.
*/
//go:nosplit
func (self class) SetSlotEnabledLeft(slot_index gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_slot_enabled_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the left (input) type of the slot with the given [param slot_index] to [param type]. If the value is negative, all connections will be disallowed to be created via user inputs.
*/
//go:nosplit
func (self class) SetSlotTypeLeft(slot_index gd.Int, atype gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_slot_type_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the left (input) type of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotTypeLeft(slot_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_slot_type_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Color] of the left (input) side of the slot with the given [param slot_index] to [param color].
*/
//go:nosplit
func (self class) SetSlotColorLeft(slot_index gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_slot_color_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the left (input) [Color] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotColorLeft(slot_index gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_slot_color_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the custom [Texture2D] of the left (input) side of the slot with the given [param slot_index] to [param custom_icon].
*/
//go:nosplit
func (self class) SetSlotCustomIconLeft(slot_index gd.Int, custom_icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, mmm.Get(custom_icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_slot_custom_icon_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the left (input) custom [Texture2D] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotCustomIconLeft(ctx gd.Lifetime, slot_index gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_slot_custom_icon_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if right (output) side of the slot with the given [param slot_index] is enabled.
*/
//go:nosplit
func (self class) IsSlotEnabledRight(slot_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_is_slot_enabled_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Toggles the right (output) side of the slot with the given [param slot_index]. If [param enable] is [code]true[/code], a port will appear on the right side and the slot will be able to be connected from this side.
*/
//go:nosplit
func (self class) SetSlotEnabledRight(slot_index gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_slot_enabled_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the right (output) type of the slot with the given [param slot_index] to [param type]. If the value is negative, all connections will be disallowed to be created via user inputs.
*/
//go:nosplit
func (self class) SetSlotTypeRight(slot_index gd.Int, atype gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_slot_type_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the right (output) type of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotTypeRight(slot_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_slot_type_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Color] of the right (output) side of the slot with the given [param slot_index] to [param color].
*/
//go:nosplit
func (self class) SetSlotColorRight(slot_index gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_slot_color_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the right (output) [Color] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotColorRight(slot_index gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_slot_color_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the custom [Texture2D] of the right (output) side of the slot with the given [param slot_index] to [param custom_icon].
*/
//go:nosplit
func (self class) SetSlotCustomIconRight(slot_index gd.Int, custom_icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, mmm.Get(custom_icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_slot_custom_icon_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the right (output) custom [Texture2D] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) GetSlotCustomIconRight(ctx gd.Lifetime, slot_index gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_slot_custom_icon_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns true if the background [StyleBox] of the slot with the given [param slot_index] is drawn.
*/
//go:nosplit
func (self class) IsSlotDrawStylebox(slot_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_is_slot_draw_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Toggles the background [StyleBox] of the slot with the given [param slot_index].
*/
//go:nosplit
func (self class) SetSlotDrawStylebox(slot_index gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slot_index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_slot_draw_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetIgnoreInvalidConnectionType(ignore bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_set_ignore_invalid_connection_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsIgnoringValidConnectionType() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_is_ignoring_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of slots with an enabled input port.
*/
//go:nosplit
func (self class) GetInputPortCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_input_port_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortPosition(port_idx gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_input_port_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the type of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortType(port_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_input_port_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Color] of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortColor(port_idx gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_input_port_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the corresponding slot index of the input port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetInputPortSlot(port_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_input_port_slot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of slots with an enabled output port.
*/
//go:nosplit
func (self class) GetOutputPortCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_output_port_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortPosition(port_idx gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_output_port_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the type of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortType(port_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_output_port_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Color] of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortColor(port_idx gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_output_port_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the corresponding slot index of the output port with the given [param port_idx].
*/
//go:nosplit
func (self class) GetOutputPortSlot(port_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphNode.Bind_get_output_port_slot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsGraphNode() Expert { return self[0].AsGraphNode() }


//go:nosplit
func (self Simple) AsGraphNode() Simple { return self[0].AsGraphNode() }


//go:nosplit
func (self class) AsGraphElement() GraphElement.Expert { return self[0].AsGraphElement() }


//go:nosplit
func (self Simple) AsGraphElement() GraphElement.Simple { return self[0].AsGraphElement() }


//go:nosplit
func (self class) AsContainer() Container.Expert { return self[0].AsContainer() }


//go:nosplit
func (self Simple) AsContainer() Container.Simple { return self[0].AsContainer() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_draw_port": return reflect.ValueOf(self._draw_port);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("GraphNode", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
