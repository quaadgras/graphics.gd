// Package GraphEdit provides methods for working with GraphEdit object instances.
package GraphEdit

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Rect2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[GraphEdit] provides tools for creation, manipulation, and display of various graphs. Its main purpose in the engine is to power the visual programming systems, such as visual shaders, but it is also available for use in user projects.
[GraphEdit] by itself is only an empty container, representing an infinite grid where [GraphNode]s can be placed. Each [GraphNode] represents a node in the graph, a single unit of data in the connected scheme. [GraphEdit], in turn, helps to control various interactions with nodes and between nodes. When the user attempts to connect, disconnect, or delete a [GraphNode], a signal is emitted in the [GraphEdit], but no action is taken by default. It is the responsibility of the programmer utilizing this control to implement the necessary logic to determine how each request should be handled.
[b]Performance:[/b] It is greatly advised to enable low-processor usage mode (see [member OS.low_processor_usage_mode]) when using GraphEdits.
[b]Note:[/b] Keep in mind that [method Node.get_children] will also return the connection layer node named [code]_connection_layer[/code] due to technical limitations. This behavior may change in future releases.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=GraphEdit)
*/
type Instance [1]gdclass.GraphEdit

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGraphEdit() Instance
}
type Interface interface {
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
	IsInInputHotzone(in_node Object.Instance, in_port int, mouse_position Vector2.XY) bool
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
	IsInOutputHotzone(in_node Object.Instance, in_port int, mouse_position Vector2.XY) bool
	//Virtual method which can be overridden to customize how connections are drawn.
	GetConnectionLine(from_position Vector2.XY, to_position Vector2.XY) []Vector2.XY
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

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) IsInInputHotzone(in_node Object.Instance, in_port int, mouse_position Vector2.XY) (_ bool) {
	return
}
func (self implementation) IsInOutputHotzone(in_node Object.Instance, in_port int, mouse_position Vector2.XY) (_ bool) {
	return
}
func (self implementation) GetConnectionLine(from_position Vector2.XY, to_position Vector2.XY) (_ []Vector2.XY) {
	return
}
func (self implementation) IsNodeHoverValid(from_node string, from_port int, to_node string, to_port int) (_ bool) {
	return
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
func (Instance) _is_in_input_hotzone(impl func(ptr unsafe.Pointer, in_node Object.Instance, in_port int, mouse_position Vector2.XY) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var in_node = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(in_node[0])
		var in_port = gd.UnsafeGet[gd.Int](p_args, 1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args, 2)
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
func (Instance) _is_in_output_hotzone(impl func(ptr unsafe.Pointer, in_node Object.Instance, in_port int, mouse_position Vector2.XY) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var in_node = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(in_node[0])
		var in_port = gd.UnsafeGet[gd.Int](p_args, 1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, int(in_port), mouse_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method which can be overridden to customize how connections are drawn.
*/
func (Instance) _get_connection_line(impl func(ptr unsafe.Pointer, from_position Vector2.XY, to_position Vector2.XY) []Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from_position = gd.UnsafeGet[gd.Vector2](p_args, 0)
		var to_position = gd.UnsafeGet[gd.Vector2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_position, to_position)
		ptr, ok := pointers.End(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&ret))))
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
func (Instance) _is_node_hover_valid(impl func(ptr unsafe.Pointer, from_node string, from_port int, to_node string, to_port int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from_node = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(from_node)
		var from_port = gd.UnsafeGet[gd.Int](p_args, 1)
		var to_node = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(to_node)
		var to_port = gd.UnsafeGet[gd.Int](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_node.String(), int(from_port), to_node.String(), int(to_port))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Create a connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection already exists, no connection is created.
*/
func (self Instance) ConnectNode(from_node string, from_port int, to_node string, to_port int) error {
	return error(gd.ToError(class(self).ConnectNode(gd.NewStringName(from_node), gd.Int(from_port), gd.NewStringName(to_node), gd.Int(to_port))))
}

/*
Returns [code]true[/code] if the [param from_port] of the [param from_node] [GraphNode] is connected to the [param to_port] of the [param to_node] [GraphNode].
*/
func (self Instance) IsNodeConnected(from_node string, from_port int, to_node string, to_port int) bool {
	return bool(class(self).IsNodeConnected(gd.NewStringName(from_node), gd.Int(from_port), gd.NewStringName(to_node), gd.Int(to_port)))
}

/*
Removes the connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection does not exist, no connection is removed.
*/
func (self Instance) DisconnectNode(from_node string, from_port int, to_node string, to_port int) {
	class(self).DisconnectNode(gd.NewStringName(from_node), gd.Int(from_port), gd.NewStringName(to_node), gd.Int(to_port))
}

/*
Sets the coloration of the connection between [param from_node]'s [param from_port] and [param to_node]'s [param to_port] with the color provided in the [theme_item activity] theme property. The color is linearly interpolated between the connection color and the activity color using [param amount] as weight.
*/
func (self Instance) SetConnectionActivity(from_node string, from_port int, to_node string, to_port int, amount Float.X) {
	class(self).SetConnectionActivity(gd.NewStringName(from_node), gd.Int(from_port), gd.NewStringName(to_node), gd.Int(to_port), gd.Float(amount))
}

/*
Returns an [Array] containing the list of connections. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
func (self Instance) GetConnectionList() []map[any]any {
	return []map[any]any(gd.ArrayAs[[]map[any]any](class(self).GetConnectionList()))
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
func (self Instance) GetClosestConnectionAtPoint(point Vector2.XY) map[any]any {
	return map[any]any(gd.DictionaryAs[any, any](class(self).GetClosestConnectionAtPoint(gd.Vector2(point), gd.Float(4.0))))
}

/*
Returns an [Array] containing the list of connections that intersect with the given [Rect2]. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
func (self Instance) GetConnectionsIntersectingWithRect(rect Rect2.PositionSize) []map[any]any {
	return []map[any]any(gd.ArrayAs[[]map[any]any](class(self).GetConnectionsIntersectingWithRect(gd.Rect2(rect))))
}

/*
Removes all connections between nodes.
*/
func (self Instance) ClearConnections() {
	class(self).ClearConnections()
}

/*
Ends the creation of the current connection. In other words, if you are dragging a connection you can use this method to abort the process and remove the line that followed your cursor.
This is best used together with [signal connection_drag_started] and [signal connection_drag_ended] to add custom behavior like node addition through shortcuts.
[b]Note:[/b] This method suppresses any other connection request signals apart from [signal connection_drag_ended].
*/
func (self Instance) ForceConnectionDragEnd() {
	class(self).ForceConnectionDragEnd()
}

/*
Allows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_right_disconnect_type].
*/
func (self Instance) AddValidRightDisconnectType(atype int) {
	class(self).AddValidRightDisconnectType(gd.Int(atype))
}

/*
Disallows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_right_disconnect_type].
*/
func (self Instance) RemoveValidRightDisconnectType(atype int) {
	class(self).RemoveValidRightDisconnectType(gd.Int(atype))
}

/*
Allows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_left_disconnect_type].
*/
func (self Instance) AddValidLeftDisconnectType(atype int) {
	class(self).AddValidLeftDisconnectType(gd.Int(atype))
}

/*
Disallows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_left_disconnect_type].
*/
func (self Instance) RemoveValidLeftDisconnectType(atype int) {
	class(self).RemoveValidLeftDisconnectType(gd.Int(atype))
}

/*
Allows the connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type] and [method remove_valid_connection_type].
*/
func (self Instance) AddValidConnectionType(from_type int, to_type int) {
	class(self).AddValidConnectionType(gd.Int(from_type), gd.Int(to_type))
}

/*
Disallows the connection between two different port types previously allowed by [method add_valid_connection_type]. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type].
*/
func (self Instance) RemoveValidConnectionType(from_type int, to_type int) {
	class(self).RemoveValidConnectionType(gd.Int(from_type), gd.Int(to_type))
}

/*
Returns whether it's possible to make a connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method add_valid_connection_type] and [method remove_valid_connection_type].
*/
func (self Instance) IsValidConnectionType(from_type int, to_type int) bool {
	return bool(class(self).IsValidConnectionType(gd.Int(from_type), gd.Int(to_type)))
}

/*
Returns the points which would make up a connection between [param from_node] and [param to_node].
*/
func (self Instance) GetConnectionLine(from_node Vector2.XY, to_node Vector2.XY) []Vector2.XY {
	return []Vector2.XY(class(self).GetConnectionLine(gd.Vector2(from_node), gd.Vector2(to_node)).AsSlice())
}

/*
Attaches the [param element] [GraphElement] to the [param frame] [GraphFrame].
*/
func (self Instance) AttachGraphElementToFrame(element string, frame_ string) {
	class(self).AttachGraphElementToFrame(gd.NewStringName(element), gd.NewStringName(frame_))
}

/*
Detaches the [param element] [GraphElement] from the [GraphFrame] it is currently attached to.
*/
func (self Instance) DetachGraphElementFromFrame(element string) {
	class(self).DetachGraphElementFromFrame(gd.NewStringName(element))
}

/*
Returns the [GraphFrame] that contains the [GraphElement] with the given name.
*/
func (self Instance) GetElementFrame(element string) [1]gdclass.GraphFrame {
	return [1]gdclass.GraphFrame(class(self).GetElementFrame(gd.NewStringName(element)))
}

/*
Returns an array of node names that are attached to the [GraphFrame] with the given name.
*/
func (self Instance) GetAttachedNodesOfFrame(frame_ string) []string {
	return []string(gd.ArrayAs[[]string](class(self).GetAttachedNodesOfFrame(gd.NewStringName(frame_))))
}

/*
Gets the [HBoxContainer] that contains the zooming and grid snap controls in the top left of the graph. You can use this method to reposition the toolbar or to add your own custom controls to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetMenuHbox() [1]gdclass.HBoxContainer {
	return [1]gdclass.HBoxContainer(class(self).GetMenuHbox())
}

/*
Rearranges selected nodes in a layout with minimum crossings between connections and uniform horizontal and vertical gap between nodes.
*/
func (self Instance) ArrangeNodes() {
	class(self).ArrangeNodes()
}

/*
Sets the specified [param node] as the one selected.
*/
func (self Instance) SetSelected(node [1]gdclass.Node) {
	class(self).SetSelected(node)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GraphEdit

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GraphEdit"))
	casted := Instance{*(*gdclass.GraphEdit)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) ScrollOffset() Vector2.XY {
	return Vector2.XY(class(self).GetScrollOffset())
}

func (self Instance) SetScrollOffset(value Vector2.XY) {
	class(self).SetScrollOffset(gd.Vector2(value))
}

func (self Instance) ShowGrid() bool {
	return bool(class(self).IsShowingGrid())
}

func (self Instance) SetShowGrid(value bool) {
	class(self).SetShowGrid(value)
}

func (self Instance) GridPattern() gdclass.GraphEditGridPattern {
	return gdclass.GraphEditGridPattern(class(self).GetGridPattern())
}

func (self Instance) SetGridPattern(value gdclass.GraphEditGridPattern) {
	class(self).SetGridPattern(value)
}

func (self Instance) SnappingEnabled() bool {
	return bool(class(self).IsSnappingEnabled())
}

func (self Instance) SetSnappingEnabled(value bool) {
	class(self).SetSnappingEnabled(value)
}

func (self Instance) SnappingDistance() int {
	return int(int(class(self).GetSnappingDistance()))
}

func (self Instance) SetSnappingDistance(value int) {
	class(self).SetSnappingDistance(gd.Int(value))
}

func (self Instance) PanningScheme() gdclass.GraphEditPanningScheme {
	return gdclass.GraphEditPanningScheme(class(self).GetPanningScheme())
}

func (self Instance) SetPanningScheme(value gdclass.GraphEditPanningScheme) {
	class(self).SetPanningScheme(value)
}

func (self Instance) RightDisconnects() bool {
	return bool(class(self).IsRightDisconnectsEnabled())
}

func (self Instance) SetRightDisconnects(value bool) {
	class(self).SetRightDisconnects(value)
}

func (self Instance) ConnectionLinesCurvature() Float.X {
	return Float.X(Float.X(class(self).GetConnectionLinesCurvature()))
}

func (self Instance) SetConnectionLinesCurvature(value Float.X) {
	class(self).SetConnectionLinesCurvature(gd.Float(value))
}

func (self Instance) ConnectionLinesThickness() Float.X {
	return Float.X(Float.X(class(self).GetConnectionLinesThickness()))
}

func (self Instance) SetConnectionLinesThickness(value Float.X) {
	class(self).SetConnectionLinesThickness(gd.Float(value))
}

func (self Instance) ConnectionLinesAntialiased() bool {
	return bool(class(self).IsConnectionLinesAntialiased())
}

func (self Instance) SetConnectionLinesAntialiased(value bool) {
	class(self).SetConnectionLinesAntialiased(value)
}

func (self Instance) Zoom() Float.X {
	return Float.X(Float.X(class(self).GetZoom()))
}

func (self Instance) SetZoom(value Float.X) {
	class(self).SetZoom(gd.Float(value))
}

func (self Instance) ZoomMin() Float.X {
	return Float.X(Float.X(class(self).GetZoomMin()))
}

func (self Instance) SetZoomMin(value Float.X) {
	class(self).SetZoomMin(gd.Float(value))
}

func (self Instance) ZoomMax() Float.X {
	return Float.X(Float.X(class(self).GetZoomMax()))
}

func (self Instance) SetZoomMax(value Float.X) {
	class(self).SetZoomMax(gd.Float(value))
}

func (self Instance) ZoomStep() Float.X {
	return Float.X(Float.X(class(self).GetZoomStep()))
}

func (self Instance) SetZoomStep(value Float.X) {
	class(self).SetZoomStep(gd.Float(value))
}

func (self Instance) MinimapEnabled() bool {
	return bool(class(self).IsMinimapEnabled())
}

func (self Instance) SetMinimapEnabled(value bool) {
	class(self).SetMinimapEnabled(value)
}

func (self Instance) MinimapSize() Vector2.XY {
	return Vector2.XY(class(self).GetMinimapSize())
}

func (self Instance) SetMinimapSize(value Vector2.XY) {
	class(self).SetMinimapSize(gd.Vector2(value))
}

func (self Instance) MinimapOpacity() Float.X {
	return Float.X(Float.X(class(self).GetMinimapOpacity()))
}

func (self Instance) SetMinimapOpacity(value Float.X) {
	class(self).SetMinimapOpacity(gd.Float(value))
}

func (self Instance) ShowMenu() bool {
	return bool(class(self).IsShowingMenu())
}

func (self Instance) SetShowMenu(value bool) {
	class(self).SetShowMenu(value)
}

func (self Instance) ShowZoomLabel() bool {
	return bool(class(self).IsShowingZoomLabel())
}

func (self Instance) SetShowZoomLabel(value bool) {
	class(self).SetShowZoomLabel(value)
}

func (self Instance) ShowZoomButtons() bool {
	return bool(class(self).IsShowingZoomButtons())
}

func (self Instance) SetShowZoomButtons(value bool) {
	class(self).SetShowZoomButtons(value)
}

func (self Instance) ShowGridButtons() bool {
	return bool(class(self).IsShowingGridButtons())
}

func (self Instance) SetShowGridButtons(value bool) {
	class(self).SetShowGridButtons(value)
}

func (self Instance) ShowMinimapButton() bool {
	return bool(class(self).IsShowingMinimapButton())
}

func (self Instance) SetShowMinimapButton(value bool) {
	class(self).SetShowMinimapButton(value)
}

func (self Instance) ShowArrangeButton() bool {
	return bool(class(self).IsShowingArrangeButton())
}

func (self Instance) SetShowArrangeButton(value bool) {
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
func (class) _is_in_input_hotzone(impl func(ptr unsafe.Pointer, in_node [1]gd.Object, in_port gd.Int, mouse_position gd.Vector2) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var in_node = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(in_node[0])
		var in_port = gd.UnsafeGet[gd.Int](p_args, 1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args, 2)
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
func (class) _is_in_output_hotzone(impl func(ptr unsafe.Pointer, in_node [1]gd.Object, in_port gd.Int, mouse_position gd.Vector2) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var in_node = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(in_node[0])
		var in_port = gd.UnsafeGet[gd.Int](p_args, 1)
		var mouse_position = gd.UnsafeGet[gd.Vector2](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, in_port, mouse_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method which can be overridden to customize how connections are drawn.
*/
func (class) _get_connection_line(impl func(ptr unsafe.Pointer, from_position gd.Vector2, to_position gd.Vector2) gd.PackedVector2Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from_position = gd.UnsafeGet[gd.Vector2](p_args, 0)
		var to_position = gd.UnsafeGet[gd.Vector2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_position, to_position)
		ptr, ok := pointers.End(ret)
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
func (class) _is_node_hover_valid(impl func(ptr unsafe.Pointer, from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from_node = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		var from_port = gd.UnsafeGet[gd.Int](p_args, 1)
		var to_node = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		var to_port = gd.UnsafeGet[gd.Int](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_node, from_port, to_node, to_port)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Create a connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection already exists, no connection is created.
*/
//go:nosplit
func (self class) ConnectNode(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, pointers.Get(to_node))
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_connect_node, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	callframe.Arg(frame, pointers.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, pointers.Get(to_node))
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_node_connected, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection does not exist, no connection is removed.
*/
//go:nosplit
func (self class) DisconnectNode(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, pointers.Get(to_node))
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_disconnect_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the coloration of the connection between [param from_node]'s [param from_port] and [param to_node]'s [param to_port] with the color provided in the [theme_item activity] theme property. The color is linearly interpolated between the connection color and the activity color using [param amount] as weight.
*/
//go:nosplit
func (self class) SetConnectionActivity(from_node gd.StringName, from_port gd.Int, to_node gd.StringName, to_port gd.Int, amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from_node))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, pointers.Get(to_node))
	callframe.Arg(frame, to_port)
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_activity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an [Array] containing the list of connections. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
//go:nosplit
func (self class) GetConnectionList() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
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
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_closest_connection_at_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
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
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connections_intersecting_with_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes all connections between nodes.
*/
//go:nosplit
func (self class) ClearConnections() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_clear_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Ends the creation of the current connection. In other words, if you are dragging a connection you can use this method to abort the process and remove the line that followed your cursor.
This is best used together with [signal connection_drag_started] and [signal connection_drag_ended] to add custom behavior like node addition through shortcuts.
[b]Note:[/b] This method suppresses any other connection request signals apart from [signal connection_drag_ended].
*/
//go:nosplit
func (self class) ForceConnectionDragEnd() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_force_connection_drag_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetScrollOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Allows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_right_disconnect_type].
*/
//go:nosplit
func (self class) AddValidRightDisconnectType(atype gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_add_valid_right_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disallows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_right_disconnect_type].
*/
//go:nosplit
func (self class) RemoveValidRightDisconnectType(atype gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_remove_valid_right_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Allows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_left_disconnect_type].
*/
//go:nosplit
func (self class) AddValidLeftDisconnectType(atype gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_add_valid_left_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disallows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_left_disconnect_type].
*/
//go:nosplit
func (self class) RemoveValidLeftDisconnectType(atype gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_remove_valid_left_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Allows the connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type] and [method remove_valid_connection_type].
*/
//go:nosplit
func (self class) AddValidConnectionType(from_type gd.Int, to_type gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_add_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disallows the connection between two different port types previously allowed by [method add_valid_connection_type]. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type].
*/
//go:nosplit
func (self class) RemoveValidConnectionType(from_type gd.Int, to_type gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_remove_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Attaches the [param element] [GraphElement] to the [param frame] [GraphFrame].
*/
//go:nosplit
func (self class) AttachGraphElementToFrame(element gd.StringName, frame_ gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(element))
	callframe.Arg(frame, pointers.Get(frame_))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_attach_graph_element_to_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Detaches the [param element] [GraphElement] from the [GraphFrame] it is currently attached to.
*/
//go:nosplit
func (self class) DetachGraphElementFromFrame(element gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(element))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_detach_graph_element_from_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [GraphFrame] that contains the [GraphElement] with the given name.
*/
//go:nosplit
func (self class) GetElementFrame(element gd.StringName) [1]gdclass.GraphFrame {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(element))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_element_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.GraphFrame{gd.PointerLifetimeBoundTo[gdclass.GraphFrame](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns an array of node names that are attached to the [GraphFrame] with the given name.
*/
//go:nosplit
func (self class) GetAttachedNodesOfFrame(frame_ gd.StringName) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(frame_))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_attached_nodes_of_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPanningScheme(scheme gdclass.GraphEditPanningScheme) {
	var frame = callframe.New()
	callframe.Arg(frame, scheme)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_panning_scheme, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPanningScheme() gdclass.GraphEditPanningScheme {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GraphEditPanningScheme](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_panning_scheme, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZoom(zoom gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, zoom)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZoom() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZoomMin(zoom_min gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, zoom_min)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZoomMin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZoomMax(zoom_max gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, zoom_max)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZoomMax() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZoomStep(zoom_step gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, zoom_step)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZoomStep() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowGrid(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_grid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingGrid() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_grid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGridPattern(pattern gdclass.GraphEditGridPattern) {
	var frame = callframe.New()
	callframe.Arg(frame, pattern)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_grid_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGridPattern() gdclass.GraphEditGridPattern {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GraphEditGridPattern](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_grid_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnappingEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_snapping_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSnappingEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_snapping_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnappingDistance(pixels gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_snapping_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSnappingDistance() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_snapping_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConnectionLinesCurvature(curvature gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, curvature)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_lines_curvature, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetConnectionLinesCurvature() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_lines_curvature, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConnectionLinesThickness(pixels gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_lines_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetConnectionLinesThickness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_lines_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConnectionLinesAntialiased(pixels bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_lines_antialiased, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsConnectionLinesAntialiased() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_connection_lines_antialiased, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinimapSize(size gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_minimap_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinimapSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_minimap_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinimapOpacity(opacity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, opacity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_minimap_opacity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinimapOpacity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_minimap_opacity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinimapEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_minimap_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMinimapEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_minimap_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowMenu(hidden bool) {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingMenu() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowZoomLabel(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_zoom_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingZoomLabel() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_zoom_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowGridButtons(hidden bool) {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_grid_buttons, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingGridButtons() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_grid_buttons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowZoomButtons(hidden bool) {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_zoom_buttons, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingZoomButtons() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_zoom_buttons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowMinimapButton(hidden bool) {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_minimap_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingMinimapButton() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_minimap_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowArrangeButton(hidden bool) {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_arrange_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingArrangeButton() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_arrange_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRightDisconnects(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_right_disconnects, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRightDisconnectsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_right_disconnects_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets the [HBoxContainer] that contains the zooming and grid snap controls in the top left of the graph. You can use this method to reposition the toolbar or to add your own custom controls to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetMenuHbox() [1]gdclass.HBoxContainer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_menu_hbox, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.HBoxContainer{gd.PointerLifetimeBoundTo[gdclass.HBoxContainer](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Rearranges selected nodes in a layout with minimum crossings between connections and uniform horizontal and vertical gap between nodes.
*/
//go:nosplit
func (self class) ArrangeNodes() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_arrange_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the specified [param node] as the one selected.
*/
//go:nosplit
func (self class) SetSelected(node [1]gdclass.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_selected, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnConnectionRequest(cb func(from_node string, from_port int, to_node string, to_port int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDisconnectionRequest(cb func(from_node string, from_port int, to_node string, to_port int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("disconnection_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectionToEmpty(cb func(from_node string, from_port int, release_position Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_to_empty"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectionFromEmpty(cb func(to_node string, to_port int, release_position Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_from_empty"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectionDragStarted(cb func(from_node string, from_port int, is_output bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_drag_started"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectionDragEnded(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_drag_ended"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCopyNodesRequest(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("copy_nodes_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPasteNodesRequest(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("paste_nodes_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDuplicateNodesRequest(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("duplicate_nodes_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDeleteNodesRequest(cb func(nodes []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("delete_nodes_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNodeSelected(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNodeDeselected(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_deselected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFrameRectChanged(cb func(frame_ [1]gdclass.GraphFrame, new_rect Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("frame_rect_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPopupRequest(cb func(at_position Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("popup_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBeginNodeMove(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("begin_node_move"), gd.NewCallable(cb), 0)
}

func (self Instance) OnEndNodeMove(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("end_node_move"), gd.NewCallable(cb), 0)
}

func (self Instance) OnGraphElementsLinkedToFrameRequest(cb func(elements []any, frame_ string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("graph_elements_linked_to_frame_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnScrollOffsetChanged(cb func(offset Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("scroll_offset_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsGraphEdit() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGraphEdit() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_is_in_input_hotzone":
		return reflect.ValueOf(self._is_in_input_hotzone)
	case "_is_in_output_hotzone":
		return reflect.ValueOf(self._is_in_output_hotzone)
	case "_get_connection_line":
		return reflect.ValueOf(self._get_connection_line)
	case "_is_node_hover_valid":
		return reflect.ValueOf(self._is_node_hover_valid)
	default:
		return gd.VirtualByName(Control.Advanced(self.AsControl()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_is_in_input_hotzone":
		return reflect.ValueOf(self._is_in_input_hotzone)
	case "_is_in_output_hotzone":
		return reflect.ValueOf(self._is_in_output_hotzone)
	case "_get_connection_line":
		return reflect.ValueOf(self._get_connection_line)
	case "_is_node_hover_valid":
		return reflect.ValueOf(self._is_node_hover_valid)
	default:
		return gd.VirtualByName(Control.Instance(self.AsControl()), name)
	}
}
func init() {
	gdclass.Register("GraphEdit", func(ptr gd.Object) any { return [1]gdclass.GraphEdit{*(*gdclass.GraphEdit)(unsafe.Pointer(&ptr))} })
}

type PanningScheme = gdclass.GraphEditPanningScheme

const (
	/*[kbd]Mouse Wheel[/kbd] will zoom, [kbd]Ctrl + Mouse Wheel[/kbd] will move the view.*/
	ScrollZooms PanningScheme = 0
	/*[kbd]Mouse Wheel[/kbd] will move the view, [kbd]Ctrl + Mouse Wheel[/kbd] will zoom.*/
	ScrollPans PanningScheme = 1
)

type GridPattern = gdclass.GraphEditGridPattern

const (
	/*Draw the grid using solid lines.*/
	GridPatternLines GridPattern = 0
	/*Draw the grid using dots.*/
	GridPatternDots GridPattern = 1
)

type Error = gd.Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
