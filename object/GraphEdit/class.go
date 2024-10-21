package GraphEdit

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[GraphEdit] provides tools for creation, manipulation, and display of various graphs. Its main purpose in the engine is to power the visual programming systems, such as visual shaders, but it is also available for use in user projects.
[GraphEdit] by itself is only an empty container, representing an infinite grid where [GraphNode]s can be placed. Each [GraphNode] represents a node in the graph, a single unit of data in the connected scheme. [GraphEdit], in turn, helps to control various interactions with nodes and between nodes. When the user attempts to connect, disconnect, or delete a [GraphNode], a signal is emitted in the [GraphEdit], but no action is taken by default. It is the responsibility of the programmer utilizing this control to implement the necessary logic to determine how each request should be handled.
[b]Performance:[/b] It is greatly advised to enable low-processor usage mode (see [member OS.low_processor_usage_mode]) when using GraphEdits.
[b]Note:[/b] Keep in mind that [method Node.get_children] will also return the connection layer node named [code]_connection_layer[/code] due to technical limitations. This behavior may change in future releases.
	// GraphEdit methods that can be overridden by a [Class] that extends it.
	type GraphEdit interface {
		//Returns whether the [param mouse_position] is in the input hot zone.
		//By default, a hot zone is a [Rect2] positioned such that its center is at [param in_node].[method GraphNode.get_input_port_position]([param in_port]) (For output's case, call [method GraphNode.get_output_port_position] instead). The hot zone's width is twice the Theme Property [code]port_grab_distance_horizontal[/code], and its height is twice the [code]port_grab_distance_vertical[/code].
		//Below is a sample code to help get started:
		//[codeblock]
		//func _is_in_input_hotzone(in_node, in_port, mouse_position):
		//    var port_size: Vector2 = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
		//    var port_pos: Vector2 = in_node.get_position() + in_node.get_input_port_position(in_port) - port_size / 2
		//    var rect = Rect2(port_pos, port_size)
		//
		//    return rect.has_point(mouse_position)
		//[/codeblock]
		IsInInputHotzone(in_node gd.Object, in_port gd.Int, mouse_position gd.Vector2) bool
		//Returns whether the [param mouse_position] is in the output hot zone. For more information on hot zones, see [method _is_in_input_hotzone].
		//Below is a sample code to help get started:
		//[codeblock]
		//func _is_in_output_hotzone(in_node, in_port, mouse_position):
		//    var port_size: Vector2 = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
		//    var port_pos: Vector2 = in_node.get_position() + in_node.get_output_port_position(in_port) - port_size / 2
		//    var rect = Rect2(port_pos, port_size)
		//
		//    return rect.has_point(mouse_position)
		//[/codeblock]
		IsInOutputHotzone(in_node gd.Object, in_port gd.Int, mouse_position gd.Vector2) bool
		//Virtual method which can be overridden to customize how connections are drawn.
		GetConnectionLine(from_position gd.Vector2, to_position gd.Vector2) gd.PackedVector2Array
		//This virtual method can be used to insert additional error detection while the user is dragging a connection over a valid port.
		//Return [code]true[/code] if the connection is indeed valid or return [code]false[/code] if the connection is impossible. If the connection is impossible, no snapping to the port and thus no connection request to that port will happen.
		//In this example a connection to same node is suppressed:
		//[codeblocks]
		//[gdscript]
		//func _is_node_hover_valid(from, from_port, to, to_port):
		//    return from != to
		//[/gdscript]
		//[csharp]
		//public override bool _IsNodeHoverValid(StringName fromNode, int fromPort, StringName toNode, int toPort)
		//{
		//    return fromNode != toNode;
		//}
		//[/csharp]
		//[/codeblocks]
		IsNodeHoverValid(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int) bool
	}

*/
type Simple [1]classdb.GraphEdit
func (Simple) _is_in_input_hotzone(impl func(ptr unsafe.Pointer, in_node gd.Object, in_port int, mouse_position gd.Vector2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var in_node gd.Object
		in_node.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var in_port = gd.UnsafeGet[gd.Int](p_args,1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, int(in_port), mouse_position)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _is_in_output_hotzone(impl func(ptr unsafe.Pointer, in_node gd.Object, in_port int, mouse_position gd.Vector2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var in_node gd.Object
		in_node.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var in_port = gd.UnsafeGet[gd.Int](p_args,1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, int(in_port), mouse_position)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_connection_line(impl func(ptr unsafe.Pointer, from_position gd.Vector2, to_position gd.Vector2) gd.PackedVector2Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var from_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		var to_position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_position, to_position)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _is_node_hover_valid(impl func(ptr unsafe.Pointer, from_node string, from_port int, to_node string, to_port int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var from_node = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var from_port = gd.UnsafeGet[gd.Int](p_args,1)
		var to_node = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2))
		var to_port = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_node.String(), int(from_port), to_node.String(), int(to_port))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (self Simple) ConnectNode(from_node string, from_port int, to_node string, to_port int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).ConnectNode(gc.StringName(from_node), gd.Int(from_port), gc.StringName(to_node), gd.Int(to_port)))
}
func (self Simple) IsNodeConnected(from_node string, from_port int, to_node string, to_port int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsNodeConnected(gc.StringName(from_node), gd.Int(from_port), gc.StringName(to_node), gd.Int(to_port)))
}
func (self Simple) DisconnectNode(from_node string, from_port int, to_node string, to_port int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DisconnectNode(gc.StringName(from_node), gd.Int(from_port), gc.StringName(to_node), gd.Int(to_port))
}
func (self Simple) SetConnectionActivity(from_node string, from_port int, to_node string, to_port int, amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConnectionActivity(gc.StringName(from_node), gd.Int(from_port), gc.StringName(to_node), gd.Int(to_port), gd.Float(amount))
}
func (self Simple) GetConnectionList() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetConnectionList(gc))
}
func (self Simple) GetClosestConnectionAtPoint(point gd.Vector2, max_distance float64) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetClosestConnectionAtPoint(gc, point, gd.Float(max_distance)))
}
func (self Simple) GetConnectionsIntersectingWithRect(rect gd.Rect2) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetConnectionsIntersectingWithRect(gc, rect))
}
func (self Simple) ClearConnections() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearConnections()
}
func (self Simple) ForceConnectionDragEnd() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceConnectionDragEnd()
}
func (self Simple) GetScrollOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetScrollOffset())
}
func (self Simple) SetScrollOffset(offset gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScrollOffset(offset)
}
func (self Simple) AddValidRightDisconnectType(atype int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddValidRightDisconnectType(gd.Int(atype))
}
func (self Simple) RemoveValidRightDisconnectType(atype int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveValidRightDisconnectType(gd.Int(atype))
}
func (self Simple) AddValidLeftDisconnectType(atype int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddValidLeftDisconnectType(gd.Int(atype))
}
func (self Simple) RemoveValidLeftDisconnectType(atype int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveValidLeftDisconnectType(gd.Int(atype))
}
func (self Simple) AddValidConnectionType(from_type int, to_type int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddValidConnectionType(gd.Int(from_type), gd.Int(to_type))
}
func (self Simple) RemoveValidConnectionType(from_type int, to_type int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveValidConnectionType(gd.Int(from_type), gd.Int(to_type))
}
func (self Simple) IsValidConnectionType(from_type int, to_type int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsValidConnectionType(gd.Int(from_type), gd.Int(to_type)))
}
func (self Simple) GetConnectionLine(from_node gd.Vector2, to_node gd.Vector2) gd.PackedVector2Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector2Array(Expert(self).GetConnectionLine(gc, from_node, to_node))
}
func (self Simple) AttachGraphElementToFrame(element string, frame_ string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AttachGraphElementToFrame(gc.StringName(element), gc.StringName(frame_))
}
func (self Simple) DetachGraphElementFromFrame(element string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DetachGraphElementFromFrame(gc.StringName(element))
}
func (self Simple) GetElementFrame(element string) [1]classdb.GraphFrame {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.GraphFrame(Expert(self).GetElementFrame(gc, gc.StringName(element)))
}
func (self Simple) GetAttachedNodesOfFrame(frame_ string) gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](Expert(self).GetAttachedNodesOfFrame(gc, gc.StringName(frame_)))
}
func (self Simple) SetPanningScheme(scheme classdb.GraphEditPanningScheme) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPanningScheme(scheme)
}
func (self Simple) GetPanningScheme() classdb.GraphEditPanningScheme {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.GraphEditPanningScheme(Expert(self).GetPanningScheme())
}
func (self Simple) SetZoom(zoom float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetZoom(gd.Float(zoom))
}
func (self Simple) GetZoom() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetZoom()))
}
func (self Simple) SetZoomMin(zoom_min float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetZoomMin(gd.Float(zoom_min))
}
func (self Simple) GetZoomMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetZoomMin()))
}
func (self Simple) SetZoomMax(zoom_max float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetZoomMax(gd.Float(zoom_max))
}
func (self Simple) GetZoomMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetZoomMax()))
}
func (self Simple) SetZoomStep(zoom_step float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetZoomStep(gd.Float(zoom_step))
}
func (self Simple) GetZoomStep() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetZoomStep()))
}
func (self Simple) SetShowGrid(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowGrid(enable)
}
func (self Simple) IsShowingGrid() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShowingGrid())
}
func (self Simple) SetGridPattern(pattern classdb.GraphEditGridPattern) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGridPattern(pattern)
}
func (self Simple) GetGridPattern() classdb.GraphEditGridPattern {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.GraphEditGridPattern(Expert(self).GetGridPattern())
}
func (self Simple) SetSnappingEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSnappingEnabled(enable)
}
func (self Simple) IsSnappingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSnappingEnabled())
}
func (self Simple) SetSnappingDistance(pixels int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSnappingDistance(gd.Int(pixels))
}
func (self Simple) GetSnappingDistance() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSnappingDistance()))
}
func (self Simple) SetConnectionLinesCurvature(curvature float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConnectionLinesCurvature(gd.Float(curvature))
}
func (self Simple) GetConnectionLinesCurvature() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetConnectionLinesCurvature()))
}
func (self Simple) SetConnectionLinesThickness(pixels float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConnectionLinesThickness(gd.Float(pixels))
}
func (self Simple) GetConnectionLinesThickness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetConnectionLinesThickness()))
}
func (self Simple) SetConnectionLinesAntialiased(pixels bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConnectionLinesAntialiased(pixels)
}
func (self Simple) IsConnectionLinesAntialiased() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsConnectionLinesAntialiased())
}
func (self Simple) SetMinimapSize(size gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinimapSize(size)
}
func (self Simple) GetMinimapSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetMinimapSize())
}
func (self Simple) SetMinimapOpacity(opacity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinimapOpacity(gd.Float(opacity))
}
func (self Simple) GetMinimapOpacity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMinimapOpacity()))
}
func (self Simple) SetMinimapEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinimapEnabled(enable)
}
func (self Simple) IsMinimapEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMinimapEnabled())
}
func (self Simple) SetShowMenu(hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowMenu(hidden)
}
func (self Simple) IsShowingMenu() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShowingMenu())
}
func (self Simple) SetShowZoomLabel(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowZoomLabel(enable)
}
func (self Simple) IsShowingZoomLabel() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShowingZoomLabel())
}
func (self Simple) SetShowGridButtons(hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowGridButtons(hidden)
}
func (self Simple) IsShowingGridButtons() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShowingGridButtons())
}
func (self Simple) SetShowZoomButtons(hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowZoomButtons(hidden)
}
func (self Simple) IsShowingZoomButtons() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShowingZoomButtons())
}
func (self Simple) SetShowMinimapButton(hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowMinimapButton(hidden)
}
func (self Simple) IsShowingMinimapButton() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShowingMinimapButton())
}
func (self Simple) SetShowArrangeButton(hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowArrangeButton(hidden)
}
func (self Simple) IsShowingArrangeButton() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShowingArrangeButton())
}
func (self Simple) SetRightDisconnects(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRightDisconnects(enable)
}
func (self Simple) IsRightDisconnectsEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRightDisconnectsEnabled())
}
func (self Simple) GetMenuHbox() [1]classdb.HBoxContainer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.HBoxContainer(Expert(self).GetMenuHbox(gc))
}
func (self Simple) ArrangeNodes() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ArrangeNodes()
}
func (self Simple) SetSelected(node [1]classdb.Node) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSelected(node)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GraphEdit
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns whether the [param mouse_position] is in the input hot zone.
By default, a hot zone is a [Rect2] positioned such that its center is at [param in_node].[method GraphNode.get_input_port_position]([param in_port]) (For output's case, call [method GraphNode.get_output_port_position] instead). The hot zone's width is twice the Theme Property [code]port_grab_distance_horizontal[/code], and its height is twice the [code]port_grab_distance_vertical[/code].
Below is a sample code to help get started:
[codeblock]
func _is_in_input_hotzone(in_node, in_port, mouse_position):
    var port_size: Vector2 = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
    var port_pos: Vector2 = in_node.get_position() + in_node.get_input_port_position(in_port) - port_size / 2
    var rect = Rect2(port_pos, port_size)

    return rect.has_point(mouse_position)
[/codeblock]
*/
func (class) _is_in_input_hotzone(impl func(ptr unsafe.Pointer, in_node gd.Object, in_port gd.Int, mouse_position gd.Vector2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var in_node gd.Object
		in_node.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var in_port = gd.UnsafeGet[gd.Int](p_args,1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, in_port, mouse_position)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns whether the [param mouse_position] is in the output hot zone. For more information on hot zones, see [method _is_in_input_hotzone].
Below is a sample code to help get started:
[codeblock]
func _is_in_output_hotzone(in_node, in_port, mouse_position):
    var port_size: Vector2 = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
    var port_pos: Vector2 = in_node.get_position() + in_node.get_output_port_position(in_port) - port_size / 2
    var rect = Rect2(port_pos, port_size)

    return rect.has_point(mouse_position)
[/codeblock]
*/
func (class) _is_in_output_hotzone(impl func(ptr unsafe.Pointer, in_node gd.Object, in_port gd.Int, mouse_position gd.Vector2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var in_node gd.Object
		in_node.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var in_port = gd.UnsafeGet[gd.Int](p_args,1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, in_port, mouse_position)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Virtual method which can be overridden to customize how connections are drawn.
*/
func (class) _get_connection_line(impl func(ptr unsafe.Pointer, from_position gd.Vector2, to_position gd.Vector2) gd.PackedVector2Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var from_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		var to_position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_position, to_position)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
This virtual method can be used to insert additional error detection while the user is dragging a connection over a valid port.
Return [code]true[/code] if the connection is indeed valid or return [code]false[/code] if the connection is impossible. If the connection is impossible, no snapping to the port and thus no connection request to that port will happen.
In this example a connection to same node is suppressed:
[codeblocks]
[gdscript]
func _is_node_hover_valid(from, from_port, to, to_port):
    return from != to
[/gdscript]
[csharp]
public override bool _IsNodeHoverValid(StringName fromNode, int fromPort, StringName toNode, int toPort)
{
    return fromNode != toNode;
}
[/csharp]
[/codeblocks]
*/
func (class) _is_node_hover_valid(impl func(ptr unsafe.Pointer, from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var from_node = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var from_port = gd.UnsafeGet[gd.Int](p_args,1)
		var to_node = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		var to_port = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_node, from_port, to_node, to_port)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Create a connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection already exists, no connection is created.
*/
//go:nosplit
func (self class) ConnectNode(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, mmm.Get(to_node))
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_connect_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [param from_port] of the [param from_node] [GraphNode] is connected to the [param to_port] of the [param to_node] [GraphNode].
*/
//go:nosplit
func (self class) IsNodeConnected(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, mmm.Get(to_node))
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_node_connected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection does not exist, no connection is removed.
*/
//go:nosplit
func (self class) DisconnectNode(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, mmm.Get(to_node))
	callframe.Arg(frame, to_port)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_disconnect_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the coloration of the connection between [param from_node]'s [param from_port] and [param to_node]'s [param to_port] with the color provided in the [theme_item activity] theme property. The color is linearly interpolated between the connection color and the activity color using [param amount] as weight.
*/
//go:nosplit
func (self class) SetConnectionActivity(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int, amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, mmm.Get(to_node))
	callframe.Arg(frame, to_port)
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_connection_activity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an [Array] containing the list of connections. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
//go:nosplit
func (self class) GetConnectionList(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_connection_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns the closest connection to the given point in screen space. If no connection is found within [param max_distance] pixels, an empty [Dictionary] is returned.
A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
For example, getting a connection at a given mouse position can be achieved like this:
[codeblocks]
[gdscript]
var connection = get_closest_connection_at_point(mouse_event.get_position())
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) GetClosestConnectionAtPoint(ctx gd.Lifetime, point gd.Vector2, max_distance gd.Float) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, max_distance)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_closest_connection_at_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an [Array] containing the list of connections that intersect with the given [Rect2]. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
//go:nosplit
func (self class) GetConnectionsIntersectingWithRect(ctx gd.Lifetime, rect gd.Rect2) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_connections_intersecting_with_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Removes all connections between nodes.
*/
//go:nosplit
func (self class) ClearConnections()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_clear_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Ends the creation of the current connection. In other words, if you are dragging a connection you can use this method to abort the process and remove the line that followed your cursor.
This is best used together with [signal connection_drag_started] and [signal connection_drag_ended] to add custom behavior like node addition through shortcuts.
[b]Note:[/b] This method suppresses any other connection request signals apart from [signal connection_drag_ended].
*/
//go:nosplit
func (self class) ForceConnectionDragEnd()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_force_connection_drag_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScrollOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScrollOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Allows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_right_disconnect_type].
*/
//go:nosplit
func (self class) AddValidRightDisconnectType(atype gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_add_valid_right_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disallows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_right_disconnect_type].
*/
//go:nosplit
func (self class) RemoveValidRightDisconnectType(atype gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_remove_valid_right_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Allows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_left_disconnect_type].
*/
//go:nosplit
func (self class) AddValidLeftDisconnectType(atype gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_add_valid_left_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disallows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_left_disconnect_type].
*/
//go:nosplit
func (self class) RemoveValidLeftDisconnectType(atype gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_remove_valid_left_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Allows the connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type] and [method remove_valid_connection_type].
*/
//go:nosplit
func (self class) AddValidConnectionType(from_type gd.Int, to_type gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_add_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disallows the connection between two different port types previously allowed by [method add_valid_connection_type]. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type].
*/
//go:nosplit
func (self class) RemoveValidConnectionType(from_type gd.Int, to_type gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_remove_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether it's possible to make a connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method add_valid_connection_type] and [method remove_valid_connection_type].
*/
//go:nosplit
func (self class) IsValidConnectionType(from_type gd.Int, to_type gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the points which would make up a connection between [param from_node] and [param to_node].
*/
//go:nosplit
func (self class) GetConnectionLine(ctx gd.Lifetime, from_node gd.Vector2, to_node gd.Vector2) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, to_node)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_connection_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Attaches the [param element] [GraphElement] to the [param frame] [GraphFrame].
*/
//go:nosplit
func (self class) AttachGraphElementToFrame(element gd.StringName, frame_ gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(element))
	callframe.Arg(frame, mmm.Get(frame_))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_attach_graph_element_to_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Detaches the [param element] [GraphElement] from the [GraphFrame] it is currently attached to.
*/
//go:nosplit
func (self class) DetachGraphElementFromFrame(element gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(element))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_detach_graph_element_from_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [GraphFrame] that contains the [GraphElement] with the given name.
*/
//go:nosplit
func (self class) GetElementFrame(ctx gd.Lifetime, element gd.StringName) object.GraphFrame {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(element))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_element_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.GraphFrame
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an array of node names that are attached to the [GraphFrame] with the given name.
*/
//go:nosplit
func (self class) GetAttachedNodesOfFrame(ctx gd.Lifetime, frame_ gd.StringName) gd.ArrayOf[gd.StringName] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(frame_))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_attached_nodes_of_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.StringName](ret)
}
//go:nosplit
func (self class) SetPanningScheme(scheme classdb.GraphEditPanningScheme)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scheme)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_panning_scheme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPanningScheme() classdb.GraphEditPanningScheme {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GraphEditPanningScheme](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_panning_scheme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZoom(zoom gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, zoom)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_zoom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZoom() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_zoom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZoomMin(zoom_min gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, zoom_min)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_zoom_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZoomMin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_zoom_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZoomMax(zoom_max gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, zoom_max)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_zoom_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZoomMax() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_zoom_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZoomStep(zoom_step gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, zoom_step)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_zoom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZoomStep() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_zoom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowGrid(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_show_grid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingGrid() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_showing_grid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGridPattern(pattern classdb.GraphEditGridPattern)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pattern)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_grid_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGridPattern() classdb.GraphEditGridPattern {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GraphEditGridPattern](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_grid_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSnappingEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_snapping_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSnappingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_snapping_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSnappingDistance(pixels gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_snapping_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSnappingDistance() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_snapping_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConnectionLinesCurvature(curvature gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, curvature)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_connection_lines_curvature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConnectionLinesCurvature() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_connection_lines_curvature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConnectionLinesThickness(pixels gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_connection_lines_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConnectionLinesThickness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_connection_lines_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConnectionLinesAntialiased(pixels bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_connection_lines_antialiased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsConnectionLinesAntialiased() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_connection_lines_antialiased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinimapSize(size gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_minimap_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinimapSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_minimap_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinimapOpacity(opacity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, opacity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_minimap_opacity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinimapOpacity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_minimap_opacity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinimapEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_minimap_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMinimapEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_minimap_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowMenu(hidden bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_show_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingMenu() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_showing_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowZoomLabel(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_show_zoom_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingZoomLabel() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_showing_zoom_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowGridButtons(hidden bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_show_grid_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingGridButtons() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_showing_grid_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowZoomButtons(hidden bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_show_zoom_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingZoomButtons() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_showing_zoom_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowMinimapButton(hidden bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_show_minimap_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingMinimapButton() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_showing_minimap_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowArrangeButton(hidden bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_show_arrange_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingArrangeButton() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_showing_arrange_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRightDisconnects(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_right_disconnects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRightDisconnectsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_is_right_disconnects_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the [HBoxContainer] that contains the zooming and grid snap controls in the top left of the graph. You can use this method to reposition the toolbar or to add your own custom controls to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetMenuHbox(ctx gd.Lifetime) object.HBoxContainer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_menu_hbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.HBoxContainer
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Rearranges selected nodes in a layout with minimum crossings between connections and uniform horizontal and vertical gap between nodes.
*/
//go:nosplit
func (self class) ArrangeNodes()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_arrange_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the specified [param node] as the one selected.
*/
//go:nosplit
func (self class) SetSelected(node object.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsGraphEdit() Expert { return self[0].AsGraphEdit() }


//go:nosplit
func (self Simple) AsGraphEdit() Simple { return self[0].AsGraphEdit() }


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
	case "_is_in_input_hotzone": return reflect.ValueOf(self._is_in_input_hotzone);
	case "_is_in_output_hotzone": return reflect.ValueOf(self._is_in_output_hotzone);
	case "_get_connection_line": return reflect.ValueOf(self._get_connection_line);
	case "_is_node_hover_valid": return reflect.ValueOf(self._is_node_hover_valid);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_is_in_input_hotzone": return reflect.ValueOf(self._is_in_input_hotzone);
	case "_is_in_output_hotzone": return reflect.ValueOf(self._is_in_output_hotzone);
	case "_get_connection_line": return reflect.ValueOf(self._get_connection_line);
	case "_is_node_hover_valid": return reflect.ValueOf(self._is_node_hover_valid);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("GraphEdit", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type PanningScheme = classdb.GraphEditPanningScheme

const (
/*[kbd]Mouse Wheel[/kbd] will zoom, [kbd]Ctrl + Mouse Wheel[/kbd] will move the view.*/
	ScrollZooms PanningScheme = 0
/*[kbd]Mouse Wheel[/kbd] will move the view, [kbd]Ctrl + Mouse Wheel[/kbd] will zoom.*/
	ScrollPans PanningScheme = 1
)
type GridPattern = classdb.GraphEditGridPattern

const (
/*Draw the grid using solid lines.*/
	GridPatternLines GridPattern = 0
/*Draw the grid using dots.*/
	GridPatternDots GridPattern = 1
)
