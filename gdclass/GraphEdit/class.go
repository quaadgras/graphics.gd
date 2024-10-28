package GraphEdit

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
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
var _ = discreet.Root

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
		var in_node = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(in_node)
		var in_port = gd.UnsafeGet[gd.Int](p_args,1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, int(in_port), mouse_position)
		gd.UnsafeSet(p_back, ret)
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
		var in_node = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(in_node)
		var in_port = gd.UnsafeGet[gd.Int](p_args,1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, int(in_port), mouse_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method which can be overridden to customize how connections are drawn.
*/
func (Go) _get_connection_line(impl func(ptr unsafe.Pointer, from_position gd.Vector2, to_position gd.Vector2) []gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var from_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		var to_position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_position, to_position)
ptr, ok := discreet.End(gd.NewPackedVector2Slice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
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
		var from_node = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(from_node)
		var from_port = gd.UnsafeGet[gd.Int](p_args,1)
		var to_node = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,2))
		defer discreet.End(to_node)
		var to_port = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_node.String(), int(from_port), to_node.String(), int(to_port))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Create a connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection already exists, no connection is created.
*/
func (self Go) ConnectNode(from_node string, from_port int, to_node string, to_port int) gd.Error {
	return gd.Error(class(self).ConnectNode(gd.NewStringName(from_node), gd.Int(from_port), gd.NewStringName(to_node), gd.Int(to_port)))
}

/*
Returns [code]true[/code] if the [param from_port] of the [param from_node] [GraphNode] is connected to the [param to_port] of the [param to_node] [GraphNode].
*/
func (self Go) IsNodeConnected(from_node string, from_port int, to_node string, to_port int) bool {
	return bool(class(self).IsNodeConnected(gd.NewStringName(from_node), gd.Int(from_port), gd.NewStringName(to_node), gd.Int(to_port)))
}

/*
Removes the connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection does not exist, no connection is removed.
*/
func (self Go) DisconnectNode(from_node string, from_port int, to_node string, to_port int) {
	class(self).DisconnectNode(gd.NewStringName(from_node), gd.Int(from_port), gd.NewStringName(to_node), gd.Int(to_port))
}

/*
Sets the coloration of the connection between [param from_node]'s [param from_port] and [param to_node]'s [param to_port] with the color provided in the [theme_item activity] theme property. The color is linearly interpolated between the connection color and the activity color using [param amount] as weight.
*/
func (self Go) SetConnectionActivity(from_node string, from_port int, to_node string, to_port int, amount float64) {
	class(self).SetConnectionActivity(gd.NewStringName(from_node), gd.Int(from_port), gd.NewStringName(to_node), gd.Int(to_port), gd.Float(amount))
}

/*
Returns an [Array] containing the list of connections. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
func (self Go) GetConnectionList() gd.Array {
	return gd.Array(class(self).GetConnectionList())
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
	return gd.Dictionary(class(self).GetClosestConnectionAtPoint(point, gd.Float(4.0)))
}

/*
Returns an [Array] containing the list of connections that intersect with the given [Rect2]. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
func (self Go) GetConnectionsIntersectingWithRect(rect gd.Rect2) gd.Array {
	return gd.Array(class(self).GetConnectionsIntersectingWithRect(rect))
}

/*
Removes all connections between nodes.
*/
func (self Go) ClearConnections() {
	class(self).ClearConnections()
}

/*
Ends the creation of the current connection. In other words, if you are dragging a connection you can use this method to abort the process and remove the line that followed your cursor.
This is best used together with [signal connection_drag_started] and [signal connection_drag_ended] to add custom behavior like node addition through shortcuts.
[b]Note:[/b] This method suppresses any other connection request signals apart from [signal connection_drag_ended].
*/
func (self Go) ForceConnectionDragEnd() {
	class(self).ForceConnectionDragEnd()
}

/*
Allows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_right_disconnect_type].
*/
func (self Go) AddValidRightDisconnectType(atype int) {
	class(self).AddValidRightDisconnectType(gd.Int(atype))
}

/*
Disallows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_right_disconnect_type].
*/
func (self Go) RemoveValidRightDisconnectType(atype int) {
	class(self).RemoveValidRightDisconnectType(gd.Int(atype))
}

/*
Allows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_left_disconnect_type].
*/
func (self Go) AddValidLeftDisconnectType(atype int) {
	class(self).AddValidLeftDisconnectType(gd.Int(atype))
}

/*
Disallows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_left_disconnect_type].
*/
func (self Go) RemoveValidLeftDisconnectType(atype int) {
	class(self).RemoveValidLeftDisconnectType(gd.Int(atype))
}

/*
Allows the connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type] and [method remove_valid_connection_type].
*/
func (self Go) AddValidConnectionType(from_type int, to_type int) {
	class(self).AddValidConnectionType(gd.Int(from_type), gd.Int(to_type))
}

/*
Disallows the connection between two different port types previously allowed by [method add_valid_connection_type]. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type].
*/
func (self Go) RemoveValidConnectionType(from_type int, to_type int) {
	class(self).RemoveValidConnectionType(gd.Int(from_type), gd.Int(to_type))
}

/*
Returns whether it's possible to make a connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method add_valid_connection_type] and [method remove_valid_connection_type].
*/
func (self Go) IsValidConnectionType(from_type int, to_type int) bool {
	return bool(class(self).IsValidConnectionType(gd.Int(from_type), gd.Int(to_type)))
}

/*
Returns the points which would make up a connection between [param from_node] and [param to_node].
*/
func (self Go) GetConnectionLine(from_node gd.Vector2, to_node gd.Vector2) []gd.Vector2 {
	return []gd.Vector2(class(self).GetConnectionLine(from_node, to_node).AsSlice())
}

/*
Attaches the [param element] [GraphElement] to the [param frame] [GraphFrame].
*/
func (self Go) AttachGraphElementToFrame(element string, frame_ string) {
	class(self).AttachGraphElementToFrame(gd.NewStringName(element), gd.NewStringName(frame_))
}

/*
Detaches the [param element] [GraphElement] from the [GraphFrame] it is currently attached to.
*/
func (self Go) DetachGraphElementFromFrame(element string) {
	class(self).DetachGraphElementFromFrame(gd.NewStringName(element))
}

/*
Returns the [GraphFrame] that contains the [GraphElement] with the given name.
*/
func (self Go) GetElementFrame(element string) gdclass.GraphFrame {
	return gdclass.GraphFrame(class(self).GetElementFrame(gd.NewStringName(element)))
}

/*
Returns an array of node names that are attached to the [GraphFrame] with the given name.
*/
func (self Go) GetAttachedNodesOfFrame(frame_ string) gd.Array {
	return gd.Array(class(self).GetAttachedNodesOfFrame(gd.NewStringName(frame_)))
}

/*
Gets the [HBoxContainer] that contains the zooming and grid snap controls in the top left of the graph. You can use this method to reposition the toolbar or to add your own custom controls to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetMenuHbox() gdclass.HBoxContainer {
	return gdclass.HBoxContainer(class(self).GetMenuHbox())
}

/*
Rearranges selected nodes in a layout with minimum crossings between connections and uniform horizontal and vertical gap between nodes.
*/
func (self Go) ArrangeNodes() {
	class(self).ArrangeNodes()
}

/*
Sets the specified [param node] as the one selected.
*/
func (self Go) SetSelected(node gdclass.Node) {
	class(self).SetSelected(node)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GraphEdit
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GraphEdit"))
	return Go{classdb.GraphEdit(object)}
}

func (self Go) ScrollOffset() gd.Vector2 {
		return gd.Vector2(class(self).GetScrollOffset())
}

func (self Go) SetScrollOffset(value gd.Vector2) {
	class(self).SetScrollOffset(value)
}

func (self Go) ShowGrid() bool {
		return bool(class(self).IsShowingGrid())
}

func (self Go) SetShowGrid(value bool) {
	class(self).SetShowGrid(value)
}

func (self Go) GridPattern() classdb.GraphEditGridPattern {
		return classdb.GraphEditGridPattern(class(self).GetGridPattern())
}

func (self Go) SetGridPattern(value classdb.GraphEditGridPattern) {
	class(self).SetGridPattern(value)
}

func (self Go) SnappingEnabled() bool {
		return bool(class(self).IsSnappingEnabled())
}

func (self Go) SetSnappingEnabled(value bool) {
	class(self).SetSnappingEnabled(value)
}

func (self Go) SnappingDistance() int {
		return int(int(class(self).GetSnappingDistance()))
}

func (self Go) SetSnappingDistance(value int) {
	class(self).SetSnappingDistance(gd.Int(value))
}

func (self Go) PanningScheme() classdb.GraphEditPanningScheme {
		return classdb.GraphEditPanningScheme(class(self).GetPanningScheme())
}

func (self Go) SetPanningScheme(value classdb.GraphEditPanningScheme) {
	class(self).SetPanningScheme(value)
}

func (self Go) RightDisconnects() bool {
		return bool(class(self).IsRightDisconnectsEnabled())
}

func (self Go) SetRightDisconnects(value bool) {
	class(self).SetRightDisconnects(value)
}

func (self Go) ConnectionLinesCurvature() float64 {
		return float64(float64(class(self).GetConnectionLinesCurvature()))
}

func (self Go) SetConnectionLinesCurvature(value float64) {
	class(self).SetConnectionLinesCurvature(gd.Float(value))
}

func (self Go) ConnectionLinesThickness() float64 {
		return float64(float64(class(self).GetConnectionLinesThickness()))
}

func (self Go) SetConnectionLinesThickness(value float64) {
	class(self).SetConnectionLinesThickness(gd.Float(value))
}

func (self Go) ConnectionLinesAntialiased() bool {
		return bool(class(self).IsConnectionLinesAntialiased())
}

func (self Go) SetConnectionLinesAntialiased(value bool) {
	class(self).SetConnectionLinesAntialiased(value)
}

func (self Go) Zoom() float64 {
		return float64(float64(class(self).GetZoom()))
}

func (self Go) SetZoom(value float64) {
	class(self).SetZoom(gd.Float(value))
}

func (self Go) ZoomMin() float64 {
		return float64(float64(class(self).GetZoomMin()))
}

func (self Go) SetZoomMin(value float64) {
	class(self).SetZoomMin(gd.Float(value))
}

func (self Go) ZoomMax() float64 {
		return float64(float64(class(self).GetZoomMax()))
}

func (self Go) SetZoomMax(value float64) {
	class(self).SetZoomMax(gd.Float(value))
}

func (self Go) ZoomStep() float64 {
		return float64(float64(class(self).GetZoomStep()))
}

func (self Go) SetZoomStep(value float64) {
	class(self).SetZoomStep(gd.Float(value))
}

func (self Go) MinimapEnabled() bool {
		return bool(class(self).IsMinimapEnabled())
}

func (self Go) SetMinimapEnabled(value bool) {
	class(self).SetMinimapEnabled(value)
}

func (self Go) MinimapSize() gd.Vector2 {
		return gd.Vector2(class(self).GetMinimapSize())
}

func (self Go) SetMinimapSize(value gd.Vector2) {
	class(self).SetMinimapSize(value)
}

func (self Go) MinimapOpacity() float64 {
		return float64(float64(class(self).GetMinimapOpacity()))
}

func (self Go) SetMinimapOpacity(value float64) {
	class(self).SetMinimapOpacity(gd.Float(value))
}

func (self Go) ShowMenu() bool {
		return bool(class(self).IsShowingMenu())
}

func (self Go) SetShowMenu(value bool) {
	class(self).SetShowMenu(value)
}

func (self Go) ShowZoomLabel() bool {
		return bool(class(self).IsShowingZoomLabel())
}

func (self Go) SetShowZoomLabel(value bool) {
	class(self).SetShowZoomLabel(value)
}

func (self Go) ShowZoomButtons() bool {
		return bool(class(self).IsShowingZoomButtons())
}

func (self Go) SetShowZoomButtons(value bool) {
	class(self).SetShowZoomButtons(value)
}

func (self Go) ShowGridButtons() bool {
		return bool(class(self).IsShowingGridButtons())
}

func (self Go) SetShowGridButtons(value bool) {
	class(self).SetShowGridButtons(value)
}

func (self Go) ShowMinimapButton() bool {
		return bool(class(self).IsShowingMinimapButton())
}

func (self Go) SetShowMinimapButton(value bool) {
	class(self).SetShowMinimapButton(value)
}

func (self Go) ShowArrangeButton() bool {
		return bool(class(self).IsShowingArrangeButton())
}

func (self Go) SetShowArrangeButton(value bool) {
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
		var in_node = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(in_node)
		var in_port = gd.UnsafeGet[gd.Int](p_args,1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, in_port, mouse_position)
		gd.UnsafeSet(p_back, ret)
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
		var in_node = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(in_node)
		var in_port = gd.UnsafeGet[gd.Int](p_args,1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, in_port, mouse_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method which can be overridden to customize how connections are drawn.
*/
func (class) _get_connection_line(impl func(ptr unsafe.Pointer, from_position gd.Vector2, to_position gd.Vector2) gd.PackedVector2Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var from_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		var to_position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_position, to_position)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
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
		var from_node = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		var from_port = gd.UnsafeGet[gd.Int](p_args,1)
		var to_node = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,2))
		var to_port = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_node, from_port, to_node, to_port)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Create a connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection already exists, no connection is created.
*/
//go:nosplit
func (self class) ConnectNode(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, discreet.Get(to_node))
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_connect_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [param from_port] of the [param from_node] [GraphNode] is connected to the [param to_port] of the [param to_node] [GraphNode].
*/
//go:nosplit
func (self class) IsNodeConnected(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, discreet.Get(to_node))
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_node_connected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection does not exist, no connection is removed.
*/
//go:nosplit
func (self class) DisconnectNode(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, discreet.Get(to_node))
	callframe.Arg(frame, to_port)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_disconnect_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the coloration of the connection between [param from_node]'s [param from_port] and [param to_node]'s [param to_port] with the color provided in the [theme_item activity] theme property. The color is linearly interpolated between the connection color and the activity color using [param amount] as weight.
*/
//go:nosplit
func (self class) SetConnectionActivity(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int, amount gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, discreet.Get(to_node))
	callframe.Arg(frame, to_port)
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_activity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an [Array] containing the list of connections. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
//go:nosplit
func (self class) GetConnectionList() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
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
func (self class) GetClosestConnectionAtPoint(point gd.Vector2, max_distance gd.Float) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, max_distance)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_closest_connection_at_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an [Array] containing the list of connections that intersect with the given [Rect2]. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
//go:nosplit
func (self class) GetConnectionsIntersectingWithRect(rect gd.Rect2) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connections_intersecting_with_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes all connections between nodes.
*/
//go:nosplit
func (self class) ClearConnections()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_clear_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Ends the creation of the current connection. In other words, if you are dragging a connection you can use this method to abort the process and remove the line that followed your cursor.
This is best used together with [signal connection_drag_started] and [signal connection_drag_ended] to add custom behavior like node addition through shortcuts.
[b]Note:[/b] This method suppresses any other connection request signals apart from [signal connection_drag_ended].
*/
//go:nosplit
func (self class) ForceConnectionDragEnd()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_force_connection_drag_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScrollOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScrollOffset(offset gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Allows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_right_disconnect_type].
*/
//go:nosplit
func (self class) AddValidRightDisconnectType(atype gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_add_valid_right_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disallows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_right_disconnect_type].
*/
//go:nosplit
func (self class) RemoveValidRightDisconnectType(atype gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_remove_valid_right_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Allows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_left_disconnect_type].
*/
//go:nosplit
func (self class) AddValidLeftDisconnectType(atype gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_add_valid_left_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disallows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_left_disconnect_type].
*/
//go:nosplit
func (self class) RemoveValidLeftDisconnectType(atype gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_remove_valid_left_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Allows the connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type] and [method remove_valid_connection_type].
*/
//go:nosplit
func (self class) AddValidConnectionType(from_type gd.Int, to_type gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_add_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disallows the connection between two different port types previously allowed by [method add_valid_connection_type]. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type].
*/
//go:nosplit
func (self class) RemoveValidConnectionType(from_type gd.Int, to_type gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_remove_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether it's possible to make a connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method add_valid_connection_type] and [method remove_valid_connection_type].
*/
//go:nosplit
func (self class) IsValidConnectionType(from_type gd.Int, to_type gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the points which would make up a connection between [param from_node] and [param to_node].
*/
//go:nosplit
func (self class) GetConnectionLine(from_node gd.Vector2, to_node gd.Vector2) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, to_node)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Attaches the [param element] [GraphElement] to the [param frame] [GraphFrame].
*/
//go:nosplit
func (self class) AttachGraphElementToFrame(element gd.StringName, frame_ gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(element))
	callframe.Arg(frame, discreet.Get(frame_))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_attach_graph_element_to_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Detaches the [param element] [GraphElement] from the [GraphFrame] it is currently attached to.
*/
//go:nosplit
func (self class) DetachGraphElementFromFrame(element gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(element))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_detach_graph_element_from_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [GraphFrame] that contains the [GraphElement] with the given name.
*/
//go:nosplit
func (self class) GetElementFrame(element gd.StringName) gdclass.GraphFrame {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(element))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_element_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.GraphFrame{classdb.GraphFrame(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns an array of node names that are attached to the [GraphFrame] with the given name.
*/
//go:nosplit
func (self class) GetAttachedNodesOfFrame(frame_ gd.StringName) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(frame_))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_attached_nodes_of_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPanningScheme(scheme classdb.GraphEditPanningScheme)  {
	var frame = callframe.New()
	callframe.Arg(frame, scheme)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_panning_scheme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPanningScheme() classdb.GraphEditPanningScheme {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GraphEditPanningScheme](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_panning_scheme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZoom(zoom gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, zoom)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZoom() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZoomMin(zoom_min gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, zoom_min)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZoomMin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZoomMax(zoom_max gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, zoom_max)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZoomMax() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZoomStep(zoom_step gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, zoom_step)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZoomStep() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowGrid(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_grid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingGrid() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_grid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGridPattern(pattern classdb.GraphEditGridPattern)  {
	var frame = callframe.New()
	callframe.Arg(frame, pattern)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_grid_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGridPattern() classdb.GraphEditGridPattern {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GraphEditGridPattern](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_grid_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSnappingEnabled(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_snapping_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSnappingEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_snapping_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSnappingDistance(pixels gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_snapping_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSnappingDistance() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_snapping_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConnectionLinesCurvature(curvature gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, curvature)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_lines_curvature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConnectionLinesCurvature() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_lines_curvature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConnectionLinesThickness(pixels gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_lines_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConnectionLinesThickness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_lines_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConnectionLinesAntialiased(pixels bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_lines_antialiased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsConnectionLinesAntialiased() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_connection_lines_antialiased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinimapSize(size gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_minimap_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinimapSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_minimap_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinimapOpacity(opacity gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, opacity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_minimap_opacity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinimapOpacity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_minimap_opacity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinimapEnabled(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_minimap_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMinimapEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_minimap_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowMenu(hidden bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingMenu() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowZoomLabel(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_zoom_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingZoomLabel() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_zoom_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowGridButtons(hidden bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_grid_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingGridButtons() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_grid_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowZoomButtons(hidden bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_zoom_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingZoomButtons() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_zoom_buttons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowMinimapButton(hidden bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_minimap_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingMinimapButton() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_minimap_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowArrangeButton(hidden bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_arrange_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingArrangeButton() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_arrange_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRightDisconnects(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_right_disconnects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRightDisconnectsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_right_disconnects_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the [HBoxContainer] that contains the zooming and grid snap controls in the top left of the graph. You can use this method to reposition the toolbar or to add your own custom controls to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetMenuHbox() gdclass.HBoxContainer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_menu_hbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.HBoxContainer{classdb.HBoxContainer(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Rearranges selected nodes in a layout with minimum crossings between connections and uniform horizontal and vertical gap between nodes.
*/
//go:nosplit
func (self class) ArrangeNodes()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_arrange_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the specified [param node] as the one selected.
*/
//go:nosplit
func (self class) SetSelected(node gdclass.Node)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(node[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnConnectionRequest(cb func(from_node string, from_port int, to_node string, to_port int)) {
	self[0].AsObject().Connect(gd.NewStringName("connection_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnDisconnectionRequest(cb func(from_node string, from_port int, to_node string, to_port int)) {
	self[0].AsObject().Connect(gd.NewStringName("disconnection_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnConnectionToEmpty(cb func(from_node string, from_port int, release_position gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("connection_to_empty"), gd.NewCallable(cb), 0)
}


func (self Go) OnConnectionFromEmpty(cb func(to_node string, to_port int, release_position gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("connection_from_empty"), gd.NewCallable(cb), 0)
}


func (self Go) OnConnectionDragStarted(cb func(from_node string, from_port int, is_output bool)) {
	self[0].AsObject().Connect(gd.NewStringName("connection_drag_started"), gd.NewCallable(cb), 0)
}


func (self Go) OnConnectionDragEnded(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("connection_drag_ended"), gd.NewCallable(cb), 0)
}


func (self Go) OnCopyNodesRequest(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("copy_nodes_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnPasteNodesRequest(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("paste_nodes_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnDuplicateNodesRequest(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("duplicate_nodes_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnDeleteNodesRequest(cb func(nodes gd.Array)) {
	self[0].AsObject().Connect(gd.NewStringName("delete_nodes_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnNodeSelected(cb func(node gdclass.Node)) {
	self[0].AsObject().Connect(gd.NewStringName("node_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnNodeDeselected(cb func(node gdclass.Node)) {
	self[0].AsObject().Connect(gd.NewStringName("node_deselected"), gd.NewCallable(cb), 0)
}


func (self Go) OnFrameRectChanged(cb func(frame_ gdclass.GraphFrame, new_rect gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("frame_rect_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnPopupRequest(cb func(at_position gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("popup_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnBeginNodeMove(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("begin_node_move"), gd.NewCallable(cb), 0)
}


func (self Go) OnEndNodeMove(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("end_node_move"), gd.NewCallable(cb), 0)
}


func (self Go) OnGraphElementsLinkedToFrameRequest(cb func(elements gd.Array, frame_ string)) {
	self[0].AsObject().Connect(gd.NewStringName("graph_elements_linked_to_frame_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnScrollOffsetChanged(cb func(offset gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("scroll_offset_changed"), gd.NewCallable(cb), 0)
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
func init() {classdb.Register("GraphEdit", func(ptr gd.Object) any { return classdb.GraphEdit(ptr) })}
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
