package GraphEdit

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
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
		IsInInputHotzone(in_node gd.Object, in_port int, mouse_position gd.Vector2) bool
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
		IsInOutputHotzone(in_node gd.Object, in_port int, mouse_position gd.Vector2) bool
		//Virtual method which can be overridden to customize how connections are drawn.
		GetConnectionLine(from_position gd.Vector2, to_position gd.Vector2) []gd.Vector2
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
		IsNodeHoverValid(from_node string, from_port int, to_node string, to_port int) bool
	}

*/
type Go [1]classdb.GraphEdit

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
func (Go) _is_in_input_hotzone(impl func(ptr unsafe.Pointer, in_node gd.Object, in_port int, mouse_position gd.Vector2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Go) _is_in_output_hotzone(impl func(ptr unsafe.Pointer, in_node gd.Object, in_port int, mouse_position gd.Vector2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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

/*
Virtual method which can be overridden to customize how connections are drawn.
*/
func (Go) _get_connection_line(impl func(ptr unsafe.Pointer, from_position gd.Vector2, to_position gd.Vector2) []gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var from_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		var to_position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_position, to_position)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedVector2Slice(ret)))
		gc.End()
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
func (Go) _is_node_hover_valid(impl func(ptr unsafe.Pointer, from_node string, from_port int, to_node string, to_port int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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

/*
Create a connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection already exists, no connection is created.
*/
func (self Go) ConnectNode(from_node string, from_port int, to_node string, to_port int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).ConnectNode(gc.StringName(from_node), gd.Int(from_port), gc.StringName(to_node), gd.Int(to_port)))
}

/*
Returns [code]true[/code] if the [param from_port] of the [param from_node] [GraphNode] is connected to the [param to_port] of the [param to_node] [GraphNode].
*/
func (self Go) IsNodeConnected(from_node string, from_port int, to_node string, to_port int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsNodeConnected(gc.StringName(from_node), gd.Int(from_port), gc.StringName(to_node), gd.Int(to_port)))
}

/*
Removes the connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection does not exist, no connection is removed.
*/
func (self Go) DisconnectNode(from_node string, from_port int, to_node string, to_port int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DisconnectNode(gc.StringName(from_node), gd.Int(from_port), gc.StringName(to_node), gd.Int(to_port))
}

/*
Sets the coloration of the connection between [param from_node]'s [param from_port] and [param to_node]'s [param to_port] with the color provided in the [theme_item activity] theme property. The color is linearly interpolated between the connection color and the activity color using [param amount] as weight.
*/
func (self Go) SetConnectionActivity(from_node string, from_port int, to_node string, to_port int, amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConnectionActivity(gc.StringName(from_node), gd.Int(from_port), gc.StringName(to_node), gd.Int(to_port), gd.Float(amount))
}

/*
Returns an [Array] containing the list of connections. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
func (self Go) GetConnectionList() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](class(self).GetConnectionList(gc))
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
func (self Go) GetClosestConnectionAtPoint(point gd.Vector2) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).GetClosestConnectionAtPoint(gc, point, gd.Float(4.0)))
}

/*
Returns an [Array] containing the list of connections that intersect with the given [Rect2]. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
func (self Go) GetConnectionsIntersectingWithRect(rect gd.Rect2) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](class(self).GetConnectionsIntersectingWithRect(gc, rect))
}

/*
Removes all connections between nodes.
*/
func (self Go) ClearConnections() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearConnections()
}

/*
Ends the creation of the current connection. In other words, if you are dragging a connection you can use this method to abort the process and remove the line that followed your cursor.
This is best used together with [signal connection_drag_started] and [signal connection_drag_ended] to add custom behavior like node addition through shortcuts.
[b]Note:[/b] This method suppresses any other connection request signals apart from [signal connection_drag_ended].
*/
func (self Go) ForceConnectionDragEnd() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ForceConnectionDragEnd()
}

/*
Allows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_right_disconnect_type].
*/
func (self Go) AddValidRightDisconnectType(atype int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddValidRightDisconnectType(gd.Int(atype))
}

/*
Disallows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_right_disconnect_type].
*/
func (self Go) RemoveValidRightDisconnectType(atype int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveValidRightDisconnectType(gd.Int(atype))
}

/*
Allows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_left_disconnect_type].
*/
func (self Go) AddValidLeftDisconnectType(atype int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddValidLeftDisconnectType(gd.Int(atype))
}

/*
Disallows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_left_disconnect_type].
*/
func (self Go) RemoveValidLeftDisconnectType(atype int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveValidLeftDisconnectType(gd.Int(atype))
}

/*
Allows the connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type] and [method remove_valid_connection_type].
*/
func (self Go) AddValidConnectionType(from_type int, to_type int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddValidConnectionType(gd.Int(from_type), gd.Int(to_type))
}

/*
Disallows the connection between two different port types previously allowed by [method add_valid_connection_type]. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type].
*/
func (self Go) RemoveValidConnectionType(from_type int, to_type int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveValidConnectionType(gd.Int(from_type), gd.Int(to_type))
}

/*
Returns whether it's possible to make a connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method add_valid_connection_type] and [method remove_valid_connection_type].
*/
func (self Go) IsValidConnectionType(from_type int, to_type int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsValidConnectionType(gd.Int(from_type), gd.Int(to_type)))
}

/*
Returns the points which would make up a connection between [param from_node] and [param to_node].
*/
func (self Go) GetConnectionLine(from_node gd.Vector2, to_node gd.Vector2) []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return []gd.Vector2(class(self).GetConnectionLine(gc, from_node, to_node).AsSlice())
}

/*
Attaches the [param element] [GraphElement] to the [param frame] [GraphFrame].
*/
func (self Go) AttachGraphElementToFrame(element string, frame_ string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AttachGraphElementToFrame(gc.StringName(element), gc.StringName(frame_))
}

/*
Detaches the [param element] [GraphElement] from the [GraphFrame] it is currently attached to.
*/
func (self Go) DetachGraphElementFromFrame(element string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DetachGraphElementFromFrame(gc.StringName(element))
}

/*
Returns the [GraphFrame] that contains the [GraphElement] with the given name.
*/
func (self Go) GetElementFrame(element string) gdclass.GraphFrame {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.GraphFrame(class(self).GetElementFrame(gc, gc.StringName(element)))
}

/*
Returns an array of node names that are attached to the [GraphFrame] with the given name.
*/
func (self Go) GetAttachedNodesOfFrame(frame_ string) gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](class(self).GetAttachedNodesOfFrame(gc, gc.StringName(frame_)))
}

/*
Gets the [HBoxContainer] that contains the zooming and grid snap controls in the top left of the graph. You can use this method to reposition the toolbar or to add your own custom controls to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetMenuHbox() gdclass.HBoxContainer {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.HBoxContainer(class(self).GetMenuHbox(gc))
}

/*
Rearranges selected nodes in a layout with minimum crossings between connections and uniform horizontal and vertical gap between nodes.
*/
func (self Go) ArrangeNodes() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ArrangeNodes()
}

/*
Sets the specified [param node] as the one selected.
*/
func (self Go) SetSelected(node gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelected(node)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GraphEdit
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GraphEdit"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ScrollOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetScrollOffset())
}

func (self Go) SetScrollOffset(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScrollOffset(value)
}

func (self Go) ShowGrid() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShowingGrid())
}

func (self Go) SetShowGrid(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShowGrid(value)
}

func (self Go) GridPattern() classdb.GraphEditGridPattern {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GraphEditGridPattern(class(self).GetGridPattern())
}

func (self Go) SetGridPattern(value classdb.GraphEditGridPattern) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGridPattern(value)
}

func (self Go) SnappingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSnappingEnabled())
}

func (self Go) SetSnappingEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSnappingEnabled(value)
}

func (self Go) SnappingDistance() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSnappingDistance()))
}

func (self Go) SetSnappingDistance(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSnappingDistance(gd.Int(value))
}

func (self Go) PanningScheme() classdb.GraphEditPanningScheme {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GraphEditPanningScheme(class(self).GetPanningScheme())
}

func (self Go) SetPanningScheme(value classdb.GraphEditPanningScheme) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPanningScheme(value)
}

func (self Go) RightDisconnects() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsRightDisconnectsEnabled())
}

func (self Go) SetRightDisconnects(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRightDisconnects(value)
}

func (self Go) ConnectionLinesCurvature() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetConnectionLinesCurvature()))
}

func (self Go) SetConnectionLinesCurvature(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConnectionLinesCurvature(gd.Float(value))
}

func (self Go) ConnectionLinesThickness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetConnectionLinesThickness()))
}

func (self Go) SetConnectionLinesThickness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConnectionLinesThickness(gd.Float(value))
}

func (self Go) ConnectionLinesAntialiased() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsConnectionLinesAntialiased())
}

func (self Go) SetConnectionLinesAntialiased(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConnectionLinesAntialiased(value)
}

func (self Go) Zoom() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetZoom()))
}

func (self Go) SetZoom(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetZoom(gd.Float(value))
}

func (self Go) ZoomMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetZoomMin()))
}

func (self Go) SetZoomMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetZoomMin(gd.Float(value))
}

func (self Go) ZoomMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetZoomMax()))
}

func (self Go) SetZoomMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetZoomMax(gd.Float(value))
}

func (self Go) ZoomStep() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetZoomStep()))
}

func (self Go) SetZoomStep(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetZoomStep(gd.Float(value))
}

func (self Go) MinimapEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMinimapEnabled())
}

func (self Go) SetMinimapEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinimapEnabled(value)
}

func (self Go) MinimapSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetMinimapSize())
}

func (self Go) SetMinimapSize(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinimapSize(value)
}

func (self Go) MinimapOpacity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMinimapOpacity()))
}

func (self Go) SetMinimapOpacity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinimapOpacity(gd.Float(value))
}

func (self Go) ShowMenu() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShowingMenu())
}

func (self Go) SetShowMenu(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShowMenu(value)
}

func (self Go) ShowZoomLabel() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShowingZoomLabel())
}

func (self Go) SetShowZoomLabel(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShowZoomLabel(value)
}

func (self Go) ShowZoomButtons() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShowingZoomButtons())
}

func (self Go) SetShowZoomButtons(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShowZoomButtons(value)
}

func (self Go) ShowGridButtons() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShowingGridButtons())
}

func (self Go) SetShowGridButtons(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShowGridButtons(value)
}

func (self Go) ShowMinimapButton() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShowingMinimapButton())
}

func (self Go) SetShowMinimapButton(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShowMinimapButton(value)
}

func (self Go) ShowArrangeButton() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShowingArrangeButton())
}

func (self Go) SetShowArrangeButton(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShowArrangeButton(value)
}

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
func (self class) GetElementFrame(ctx gd.Lifetime, element gd.StringName) gdclass.GraphFrame {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(element))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_element_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.GraphFrame
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
func (self class) GetMenuHbox(ctx gd.Lifetime) gdclass.HBoxContainer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_get_menu_hbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.HBoxContainer
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
func (self class) SetSelected(node gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphEdit.Bind_set_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnConnectionRequest(cb func(from_node string, from_port int, to_node string, to_port int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("connection_request"), gc.Callable(cb), 0)
}


func (self Go) OnDisconnectionRequest(cb func(from_node string, from_port int, to_node string, to_port int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("disconnection_request"), gc.Callable(cb), 0)
}


func (self Go) OnConnectionToEmpty(cb func(from_node string, from_port int, release_position gd.Vector2)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("connection_to_empty"), gc.Callable(cb), 0)
}


func (self Go) OnConnectionFromEmpty(cb func(to_node string, to_port int, release_position gd.Vector2)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("connection_from_empty"), gc.Callable(cb), 0)
}


func (self Go) OnConnectionDragStarted(cb func(from_node string, from_port int, is_output bool)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("connection_drag_started"), gc.Callable(cb), 0)
}


func (self Go) OnConnectionDragEnded(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("connection_drag_ended"), gc.Callable(cb), 0)
}


func (self Go) OnCopyNodesRequest(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("copy_nodes_request"), gc.Callable(cb), 0)
}


func (self Go) OnPasteNodesRequest(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("paste_nodes_request"), gc.Callable(cb), 0)
}


func (self Go) OnDuplicateNodesRequest(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("duplicate_nodes_request"), gc.Callable(cb), 0)
}


func (self Go) OnDeleteNodesRequest(cb func(nodes gd.ArrayOf[gd.StringName])) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("delete_nodes_request"), gc.Callable(cb), 0)
}


func (self Go) OnNodeSelected(cb func(node gdclass.Node)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("node_selected"), gc.Callable(cb), 0)
}


func (self Go) OnNodeDeselected(cb func(node gdclass.Node)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("node_deselected"), gc.Callable(cb), 0)
}


func (self Go) OnFrameRectChanged(cb func(frame_ gdclass.GraphFrame, new_rect gd.Vector2)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("frame_rect_changed"), gc.Callable(cb), 0)
}


func (self Go) OnPopupRequest(cb func(at_position gd.Vector2)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("popup_request"), gc.Callable(cb), 0)
}


func (self Go) OnBeginNodeMove(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("begin_node_move"), gc.Callable(cb), 0)
}


func (self Go) OnEndNodeMove(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("end_node_move"), gc.Callable(cb), 0)
}


func (self Go) OnGraphElementsLinkedToFrameRequest(cb func(elements gd.Array, frame_ string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("graph_elements_linked_to_frame_request"), gc.Callable(cb), 0)
}


func (self Go) OnScrollOffsetChanged(cb func(offset gd.Vector2)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("scroll_offset_changed"), gc.Callable(cb), 0)
}


func (self class) AsGraphEdit() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGraphEdit() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_is_in_input_hotzone": return reflect.ValueOf(self._is_in_input_hotzone);
	case "_is_in_output_hotzone": return reflect.ValueOf(self._is_in_output_hotzone);
	case "_get_connection_line": return reflect.ValueOf(self._get_connection_line);
	case "_is_node_hover_valid": return reflect.ValueOf(self._is_node_hover_valid);
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_is_in_input_hotzone": return reflect.ValueOf(self._is_in_input_hotzone);
	case "_is_in_output_hotzone": return reflect.ValueOf(self._is_in_output_hotzone);
	case "_get_connection_line": return reflect.ValueOf(self._get_connection_line);
	case "_is_node_hover_valid": return reflect.ValueOf(self._is_node_hover_valid);
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {classdb.Register("GraphEdit", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
