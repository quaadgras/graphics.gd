// Package Node provides methods for working with Node object instances.
package Node

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
Nodes are Godot's building blocks. They can be assigned as the child of another node, resulting in a tree arrangement. A given node can contain any number of nodes as children with the requirement that all siblings (direct children of a node) should have unique names.
A tree of nodes is called a [i]scene[/i]. Scenes can be saved to the disk and then instantiated into other scenes. This allows for very high flexibility in the architecture and data model of Godot projects.
[b]Scene tree:[/b] The [SceneTree] contains the active tree of nodes. When a node is added to the scene tree, it receives the [constant NOTIFICATION_ENTER_TREE] notification and its [method _enter_tree] callback is triggered. Child nodes are always added [i]after[/i] their parent node, i.e. the [method _enter_tree] callback of a parent node will be triggered before its child's.
Once all nodes have been added in the scene tree, they receive the [constant NOTIFICATION_READY] notification and their respective [method _ready] callbacks are triggered. For groups of nodes, the [method _ready] callback is called in reverse order, starting with the children and moving up to the parent nodes.
This means that when adding a node to the scene tree, the following order will be used for the callbacks: [method _enter_tree] of the parent, [method _enter_tree] of the children, [method _ready] of the children and finally [method _ready] of the parent (recursively for the entire scene tree).
[b]Processing:[/b] Nodes can override the "process" state, so that they receive a callback on each frame requesting them to process (do something). Normal processing (callback [method _process], toggled with [method set_process]) happens as fast as possible and is dependent on the frame rate, so the processing time [i]delta[/i] (in seconds) is passed as an argument. Physics processing (callback [method _physics_process], toggled with [method set_physics_process]) happens a fixed number of times per second (60 by default) and is useful for code related to the physics engine.
Nodes can also process input events. When present, the [method _input] function will be called for each input that the program receives. In many cases, this can be overkill (unless used for simple projects), and the [method _unhandled_input] function might be preferred; it is called when the input event was not handled by anyone else (typically, GUI [Control] nodes), ensuring that the node only receives the events that were meant for it.
To keep track of the scene hierarchy (especially when instantiating scenes into other scenes), an "owner" can be set for the node with the [member owner] property. This keeps track of who instantiated what. This is mostly useful when writing editors and tools, though.
Finally, when a node is freed with [method Object.free] or [method queue_free], it will also free all its children.
[b]Groups:[/b] Nodes can be added to as many groups as you want to be easy to manage, you could create groups like "enemies" or "collectables" for example, depending on your game. See [method add_to_group], [method is_in_group] and [method remove_from_group]. You can then retrieve all nodes in these groups, iterate them and even call methods on groups via the methods on [SceneTree].
[b]Networking with nodes:[/b] After connecting to a server (or making one, see [ENetMultiplayerPeer]), it is possible to use the built-in RPC (remote procedure call) system to communicate over the network. By calling [method rpc] with a method name, it will be called locally and in all connected peers (peers = clients and the server that accepts connections). To identify which node receives the RPC call, Godot will use its [NodePath] (make sure node names are the same on all peers). Also, take a look at the high-level networking tutorial and corresponding demos.
[b]Note:[/b] The [code]script[/code] property is part of the [Object] class, not [Node]. It isn't exposed like most properties but does have a setter and getter (see [method Object.set_script] and [method Object.get_script]).

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=Node)
*/
type Instance [1]gdclass.Node

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNode() Instance
}
type Interface interface {
	//Called during the processing step of the main loop. Processing happens at every frame and as fast as possible, so the [param delta] time since the previous frame is not constant. [param delta] is in seconds.
	//It is only called if processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process].
	//Corresponds to the [constant NOTIFICATION_PROCESS] notification in [method Object._notification].
	//[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
	Process(delta Float.X)
	//Called during the physics processing step of the main loop. Physics processing means that the frame rate is synced to the physics, i.e. the [param delta] variable should be constant. [param delta] is in seconds.
	//It is only called if physics processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_physics_process].
	//Corresponds to the [constant NOTIFICATION_PHYSICS_PROCESS] notification in [method Object._notification].
	//[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
	PhysicsProcess(delta Float.X)
	//Called when the node enters the [SceneTree] (e.g. upon instantiating, scene changing, or after calling [method add_child] in a script). If the node has children, its [method _enter_tree] callback will be called first, and then that of the children.
	//Corresponds to the [constant NOTIFICATION_ENTER_TREE] notification in [method Object._notification].
	EnterTree()
	//Called when the node is about to leave the [SceneTree] (e.g. upon freeing, scene changing, or after calling [method remove_child] in a script). If the node has children, its [method _exit_tree] callback will be called last, after all its children have left the tree.
	//Corresponds to the [constant NOTIFICATION_EXIT_TREE] notification in [method Object._notification] and signal [signal tree_exiting]. To get notified when the node has already left the active tree, connect to the [signal tree_exited].
	ExitTree()
	//Called when the node is "ready", i.e. when both the node and its children have entered the scene tree. If the node has children, their [method _ready] callbacks get triggered first, and the parent node will receive the ready notification afterwards.
	//Corresponds to the [constant NOTIFICATION_READY] notification in [method Object._notification]. See also the [code]@onready[/code] annotation for variables.
	//Usually used for initialization. For even earlier initialization, [method Object._init] may be used. See also [method _enter_tree].
	//[b]Note:[/b] This method may be called only once for each node. After removing a node from the scene tree and adding it again, [method _ready] will [b]not[/b] be called a second time. This can be bypassed by requesting another call with [method request_ready], which may be called anywhere before adding the node again.
	Ready()
	//The elements in the array returned from this method are displayed as warnings in the Scene dock if the script that overrides it is a [code]tool[/code] script.
	//Returning an empty array produces no warnings.
	//Call [method update_configuration_warnings] when the warnings need to be updated for this node.
	//[codeblock]
	//@export var energy = 0:
	//    set(value):
	//        energy = value
	//        update_configuration_warnings()
	//
	//func _get_configuration_warnings():
	//    if energy < 0:
	//        return ["Energy must be 0 or greater."]
	//    else:
	//        return []
	//[/codeblock]
	GetConfigurationWarnings() []string
	//Called when there is an input event. The input event propagates up through the node tree until a node consumes it.
	//It is only called if input processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_input].
	//To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
	//For gameplay input, [method _unhandled_input] and [method _unhandled_key_input] are usually a better fit as they allow the GUI to intercept the events first.
	//[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
	Input(event [1]gdclass.InputEvent)
	//Called when an [InputEventKey], [InputEventShortcut], or [InputEventJoypadButton] hasn't been consumed by [method _input] or any GUI [Control] item. It is called before [method _unhandled_key_input] and [method _unhandled_input]. The input event propagates up through the node tree until a node consumes it.
	//It is only called if shortcut processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_shortcut_input].
	//To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
	//This method can be used to handle shortcuts. For generic GUI events, use [method _input] instead. Gameplay events should usually be handled with either [method _unhandled_input] or [method _unhandled_key_input].
	//[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not orphan).
	ShortcutInput(event [1]gdclass.InputEvent)
	//Called when an [InputEvent] hasn't been consumed by [method _input] or any GUI [Control] item. It is called after [method _shortcut_input] and after [method _unhandled_key_input]. The input event propagates up through the node tree until a node consumes it.
	//It is only called if unhandled input processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_unhandled_input].
	//To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
	//For gameplay input, this method is usually a better fit than [method _input], as GUI events need a higher priority. For keyboard shortcuts, consider using [method _shortcut_input] instead, as it is called before this method. Finally, to handle keyboard events, consider using [method _unhandled_key_input] for performance reasons.
	//[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
	UnhandledInput(event [1]gdclass.InputEvent)
	//Called when an [InputEventKey] hasn't been consumed by [method _input] or any GUI [Control] item. It is called after [method _shortcut_input] but before [method _unhandled_input]. The input event propagates up through the node tree until a node consumes it.
	//It is only called if unhandled key input processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_unhandled_key_input].
	//To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
	//This method can be used to handle Unicode character input with [kbd]Alt[/kbd], [kbd]Alt + Ctrl[/kbd], and [kbd]Alt + Shift[/kbd] modifiers, after shortcuts were handled.
	//For gameplay input, this and [method _unhandled_input] are usually a better fit than [method _input], as GUI events should be handled first. This method also performs better than [method _unhandled_input], since unrelated events such as [InputEventMouseMotion] are automatically filtered. For shortcuts, consider using [method _shortcut_input] instead.
	//[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
	UnhandledKeyInput(event [1]gdclass.InputEvent)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) Process(delta Float.X)                         { return }
func (self implementation) PhysicsProcess(delta Float.X)                  { return }
func (self implementation) EnterTree()                                    { return }
func (self implementation) ExitTree()                                     { return }
func (self implementation) Ready()                                        { return }
func (self implementation) GetConfigurationWarnings() (_ []string)        { return }
func (self implementation) Input(event [1]gdclass.InputEvent)             { return }
func (self implementation) ShortcutInput(event [1]gdclass.InputEvent)     { return }
func (self implementation) UnhandledInput(event [1]gdclass.InputEvent)    { return }
func (self implementation) UnhandledKeyInput(event [1]gdclass.InputEvent) { return }

/*
Called during the processing step of the main loop. Processing happens at every frame and as fast as possible, so the [param delta] time since the previous frame is not constant. [param delta] is in seconds.
It is only called if processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process].
Corresponds to the [constant NOTIFICATION_PROCESS] notification in [method Object._notification].
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
*/
func (Instance) _process(impl func(ptr unsafe.Pointer, delta Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var delta = gd.UnsafeGet[float64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(delta))
	}
}

/*
Called during the physics processing step of the main loop. Physics processing means that the frame rate is synced to the physics, i.e. the [param delta] variable should be constant. [param delta] is in seconds.
It is only called if physics processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_physics_process].
Corresponds to the [constant NOTIFICATION_PHYSICS_PROCESS] notification in [method Object._notification].
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
*/
func (Instance) _physics_process(impl func(ptr unsafe.Pointer, delta Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var delta = gd.UnsafeGet[float64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(delta))
	}
}

/*
Called when the node enters the [SceneTree] (e.g. upon instantiating, scene changing, or after calling [method add_child] in a script). If the node has children, its [method _enter_tree] callback will be called first, and then that of the children.
Corresponds to the [constant NOTIFICATION_ENTER_TREE] notification in [method Object._notification].
*/
func (Instance) _enter_tree(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the node is about to leave the [SceneTree] (e.g. upon freeing, scene changing, or after calling [method remove_child] in a script). If the node has children, its [method _exit_tree] callback will be called last, after all its children have left the tree.
Corresponds to the [constant NOTIFICATION_EXIT_TREE] notification in [method Object._notification] and signal [signal tree_exiting]. To get notified when the node has already left the active tree, connect to the [signal tree_exited].
*/
func (Instance) _exit_tree(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the node is "ready", i.e. when both the node and its children have entered the scene tree. If the node has children, their [method _ready] callbacks get triggered first, and the parent node will receive the ready notification afterwards.
Corresponds to the [constant NOTIFICATION_READY] notification in [method Object._notification]. See also the [code]@onready[/code] annotation for variables.
Usually used for initialization. For even earlier initialization, [method Object._init] may be used. See also [method _enter_tree].
[b]Note:[/b] This method may be called only once for each node. After removing a node from the scene tree and adding it again, [method _ready] will [b]not[/b] be called a second time. This can be bypassed by requesting another call with [method request_ready], which may be called anywhere before adding the node again.
*/
func (Instance) _ready(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
The elements in the array returned from this method are displayed as warnings in the Scene dock if the script that overrides it is a [code]tool[/code] script.
Returning an empty array produces no warnings.
Call [method update_configuration_warnings] when the warnings need to be updated for this node.
[codeblock]
@export var energy = 0:

	set(value):
	    energy = value
	    update_configuration_warnings()

func _get_configuration_warnings():

	if energy < 0:
	    return ["Energy must be 0 or greater."]
	else:
	    return []

[/codeblock]
*/
func (Instance) _get_configuration_warnings(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(Packed.MakeStrings(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Called when there is an input event. The input event propagates up through the node tree until a node consumes it.
It is only called if input processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_input].
To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
For gameplay input, [method _unhandled_input] and [method _unhandled_key_input] are usually a better fit as they allow the GUI to intercept the events first.
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
*/
func (Instance) _input(impl func(ptr unsafe.Pointer, event [1]gdclass.InputEvent)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, event)
	}
}

/*
Called when an [InputEventKey], [InputEventShortcut], or [InputEventJoypadButton] hasn't been consumed by [method _input] or any GUI [Control] item. It is called before [method _unhandled_key_input] and [method _unhandled_input]. The input event propagates up through the node tree until a node consumes it.
It is only called if shortcut processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_shortcut_input].
To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
This method can be used to handle shortcuts. For generic GUI events, use [method _input] instead. Gameplay events should usually be handled with either [method _unhandled_input] or [method _unhandled_key_input].
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not orphan).
*/
func (Instance) _shortcut_input(impl func(ptr unsafe.Pointer, event [1]gdclass.InputEvent)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, event)
	}
}

/*
Called when an [InputEvent] hasn't been consumed by [method _input] or any GUI [Control] item. It is called after [method _shortcut_input] and after [method _unhandled_key_input]. The input event propagates up through the node tree until a node consumes it.
It is only called if unhandled input processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_unhandled_input].
To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
For gameplay input, this method is usually a better fit than [method _input], as GUI events need a higher priority. For keyboard shortcuts, consider using [method _shortcut_input] instead, as it is called before this method. Finally, to handle keyboard events, consider using [method _unhandled_key_input] for performance reasons.
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
*/
func (Instance) _unhandled_input(impl func(ptr unsafe.Pointer, event [1]gdclass.InputEvent)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, event)
	}
}

/*
Called when an [InputEventKey] hasn't been consumed by [method _input] or any GUI [Control] item. It is called after [method _shortcut_input] but before [method _unhandled_input]. The input event propagates up through the node tree until a node consumes it.
It is only called if unhandled key input processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_unhandled_key_input].
To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
This method can be used to handle Unicode character input with [kbd]Alt[/kbd], [kbd]Alt + Ctrl[/kbd], and [kbd]Alt + Shift[/kbd] modifiers, after shortcuts were handled.
For gameplay input, this and [method _unhandled_input] are usually a better fit than [method _input], as GUI events should be handled first. This method also performs better than [method _unhandled_input], since unrelated events such as [InputEventMouseMotion] are automatically filtered. For shortcuts, consider using [method _shortcut_input] instead.
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
*/
func (Instance) _unhandled_key_input(impl func(ptr unsafe.Pointer, event [1]gdclass.InputEvent)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, event)
	}
}

/*
Prints all orphan nodes (nodes outside the [SceneTree]). Useful for debugging.
[b]Note:[/b] This method only works in debug builds. Does nothing in a project exported in release mode.
*/
func PrintOrphanNodes() { //gd:Node.print_orphan_nodes
	self := Instance{}
	class(self).PrintOrphanNodes()
}

/*
Adds a [param sibling] node to this node's parent, and moves the added sibling right below this node.
If [param force_readable_name] is [code]true[/code], improves the readability of the added [param sibling]. If not named, the [param sibling] is renamed to its type, and if it shares [member name] with a sibling, a number is suffixed more appropriately. This operation is very slow. As such, it is recommended leaving this to [code]false[/code], which assigns a dummy name featuring [code]@[/code] in both situations.
Use [method add_child] instead of this method if you don't need the child node to be added below a specific node in the list of children.
[b]Note:[/b] If this node is internal, the added sibling will be internal too (see [method add_child]'s [code]internal[/code] parameter).
*/
func (self Instance) AddSibling(sibling [1]gdclass.Node) { //gd:Node.add_sibling
	class(self).AddSibling(sibling, false)
}

/*
Adds a child [param node]. Nodes can have any number of children, but every child must have a unique name. Child nodes are automatically deleted when the parent node is deleted, so an entire scene can be removed by deleting its topmost node.
If [param force_readable_name] is [code]true[/code], improves the readability of the added [param node]. If not named, the [param node] is renamed to its type, and if it shares [member name] with a sibling, a number is suffixed more appropriately. This operation is very slow. As such, it is recommended leaving this to [code]false[/code], which assigns a dummy name featuring [code]@[/code] in both situations.
If [param internal] is different than [constant INTERNAL_MODE_DISABLED], the child will be added as internal node. These nodes are ignored by methods like [method get_children], unless their parameter [code]include_internal[/code] is [code]true[/code]. The intended usage is to hide the internal nodes from the user, so the user won't accidentally delete or modify them. Used by some GUI nodes, e.g. [ColorPicker]. See [enum InternalMode] for available modes.
[b]Note:[/b] If [param node] already has a parent, this method will fail. Use [method remove_child] first to remove [param node] from its current parent. For example:
[codeblocks]
[gdscript]
var child_node = get_child(0)
if child_node.get_parent():

	child_node.get_parent().remove_child(child_node)

add_child(child_node)
[/gdscript]
[csharp]
Node childNode = GetChild(0);
if (childNode.GetParent() != null)

	{
	    childNode.GetParent().RemoveChild(childNode);
	}

AddChild(childNode);
[/csharp]
[/codeblocks]
If you need the child node to be added below a specific node in the list of children, use [method add_sibling] instead of this method.
[b]Note:[/b] If you want a child to be persisted to a [PackedScene], you must set [member owner] in addition to calling [method add_child]. This is typically relevant for [url=$DOCS_URL/tutorials/plugins/running_code_in_the_editor.html]tool scripts[/url] and [url=$DOCS_URL/tutorials/plugins/editor/index.html]editor plugins[/url]. If [method add_child] is called without setting [member owner], the newly added [Node] will not be visible in the scene tree, though it will be visible in the 2D/3D view.
*/
func (self Instance) AddChild(node [1]gdclass.Node) { //gd:Node.add_child
	class(self).AddChild(node, false, 0)
}

/*
Removes a child [param node]. The [param node], along with its children, are [b]not[/b] deleted. To delete a node, see [method queue_free].
[b]Note:[/b] When this node is inside the tree, this method sets the [member owner] of the removed [param node] (or its descendants) to [code]null[/code], if their [member owner] is no longer an ancestor (see [method is_ancestor_of]).
*/
func (self Instance) RemoveChild(node [1]gdclass.Node) { //gd:Node.remove_child
	class(self).RemoveChild(node)
}

/*
Changes the parent of this [Node] to the [param new_parent]. The node needs to already have a parent. The node's [member owner] is preserved if its owner is still reachable from the new location (i.e., the node is still a descendant of the new parent after the operation).
If [param keep_global_transform] is [code]true[/code], the node's global transform will be preserved if supported. [Node2D], [Node3D] and [Control] support this argument (but [Control] keeps only position).
*/
func (self Instance) Reparent(new_parent [1]gdclass.Node) { //gd:Node.reparent
	class(self).Reparent(new_parent, true)
}

/*
Returns the number of children of this node.
If [param include_internal] is [code]false[/code], internal children are not counted (see [method add_child]'s [code]internal[/code] parameter).
*/
func (self Instance) GetChildCount() int { //gd:Node.get_child_count
	return int(int(class(self).GetChildCount(false)))
}

/*
Returns all children of this node inside an [Array].
If [param include_internal] is [code]false[/code], excludes internal children from the returned array (see [method add_child]'s [code]internal[/code] parameter).
*/
func (self Instance) GetChildren() [][1]gdclass.Node { //gd:Node.get_children
	return [][1]gdclass.Node(gd.ArrayAs[[][1]gdclass.Node](gd.InternalArray(class(self).GetChildren(false))))
}

/*
Fetches a child node by its index. Each child node has an index relative its siblings (see [method get_index]). The first child is at index 0. Negative values can also be used to start from the end of the list. This method can be used in combination with [method get_child_count] to iterate over this node's children. If no child exists at the given index, this method returns [code]null[/code] and an error is generated.
If [param include_internal] is [code]false[/code], internal children are ignored (see [method add_child]'s [code]internal[/code] parameter).
[codeblock]
# Assuming the following are children of this node, in order:
# First, Middle, Last.

var a = get_child(0).name  # a is "First"
var b = get_child(1).name  # b is "Middle"
var b = get_child(2).name  # b is "Last"
var c = get_child(-1).name # c is "Last"
[/codeblock]
[b]Note:[/b] To fetch a node by [NodePath], use [method get_node].
*/
func (self Instance) GetChild(idx int) [1]gdclass.Node { //gd:Node.get_child
	return [1]gdclass.Node(class(self).GetChild(int64(idx), false))
}

/*
Returns [code]true[/code] if the [param path] points to a valid node. See also [method get_node].
*/
func (self Instance) HasNode(path string) bool { //gd:Node.has_node
	return bool(class(self).HasNode(Path.ToNode(String.New(path))))
}

/*
Fetches a node. The [NodePath] can either be a relative path (from this node), or an absolute path (from the [member SceneTree.root]) to a node. If [param path] does not point to a valid node, generates an error and returns [code]null[/code]. Attempts to access methods on the return value will result in an [i]"Attempt to call <method> on a null instance."[/i] error.
[b]Note:[/b] Fetching by absolute path only works when the node is inside the scene tree (see [method is_inside_tree]).
[b]Example:[/b] Assume this method is called from the Character node, inside the following tree:
[codeblock lang=text]

	┖╴root
	   ┠╴Character (you are here!)
	   ┃  ┠╴Sword
	   ┃  ┖╴Backpack
	   ┃     ┖╴Dagger
	   ┠╴MyGame
	   ┖╴Swamp
	      ┠╴Alligator
	      ┠╴Mosquito
	      ┖╴Goblin

[/codeblock]
The following calls will return a valid node:
[codeblocks]
[gdscript]
get_node("Sword")
get_node("Backpack/Dagger")
get_node("../Swamp/Alligator")
get_node("/root/MyGame")
[/gdscript]
[csharp]
GetNode("Sword");
GetNode("Backpack/Dagger");
GetNode("../Swamp/Alligator");
GetNode("/root/MyGame");
[/csharp]
[/codeblocks]
*/
func (self Instance) GetNode(path string) [1]gdclass.Node { //gd:Node.get_node
	return [1]gdclass.Node(class(self).GetNode(Path.ToNode(String.New(path))))
}

/*
Fetches a node by [NodePath]. Similar to [method get_node], but does not generate an error if [param path] does not point to a valid node.
*/
func (self Instance) GetNodeOrNull(path string) [1]gdclass.Node { //gd:Node.get_node_or_null
	return [1]gdclass.Node(class(self).GetNodeOrNull(Path.ToNode(String.New(path))))
}

/*
Returns this node's parent node, or [code]null[/code] if the node doesn't have a parent.
*/
func (self Instance) GetParent() [1]gdclass.Node { //gd:Node.get_parent
	return [1]gdclass.Node(class(self).GetParent())
}

/*
Finds the first descendant of this node whose [member name] matches [param pattern], returning [code]null[/code] if no match is found. The matching is done against node names, [i]not[/i] their paths, through [method String.match]. As such, it is case-sensitive, [code]"*"[/code] matches zero or more characters, and [code]"?"[/code] matches any single character.
If [param recursive] is [code]false[/code], only this node's direct children are checked. Nodes are checked in tree order, so this node's first direct child is checked first, then its own direct children, etc., before moving to the second direct child, and so on. Internal children are also included in the search (see [code]internal[/code] parameter in [method add_child]).
If [param owned] is [code]true[/code], only descendants with a valid [member owner] node are checked.
[b]Note:[/b] This method can be very slow. Consider storing a reference to the found node in a variable. Alternatively, use [method get_node] with unique names (see [member unique_name_in_owner]).
[b]Note:[/b] To find all descendant nodes matching a pattern or a class type, see [method find_children].
*/
func (self Instance) FindChild(pattern string) [1]gdclass.Node { //gd:Node.find_child
	return [1]gdclass.Node(class(self).FindChild(String.New(pattern), true, true))
}

/*
Finds all descendants of this node whose names match [param pattern], returning an empty [Array] if no match is found. The matching is done against node names, [i]not[/i] their paths, through [method String.match]. As such, it is case-sensitive, [code]"*"[/code] matches zero or more characters, and [code]"?"[/code] matches any single character.
If [param type] is not empty, only ancestors inheriting from [param type] are included (see [method Object.is_class]).
If [param recursive] is [code]false[/code], only this node's direct children are checked. Nodes are checked in tree order, so this node's first direct child is checked first, then its own direct children, etc., before moving to the second direct child, and so on. Internal children are also included in the search (see [code]internal[/code] parameter in [method add_child]).
If [param owned] is [code]true[/code], only descendants with a valid [member owner] node are checked.
[b]Note:[/b] This method can be very slow. Consider storing references to the found nodes in a variable.
[b]Note:[/b] To find a single descendant node matching a pattern, see [method find_child].
*/
func (self Instance) FindChildren(pattern string) [][1]gdclass.Node { //gd:Node.find_children
	return [][1]gdclass.Node(gd.ArrayAs[[][1]gdclass.Node](gd.InternalArray(class(self).FindChildren(String.New(pattern), String.New(""), true, true))))
}

/*
Finds the first ancestor of this node whose [member name] matches [param pattern], returning [code]null[/code] if no match is found. The matching is done through [method String.match]. As such, it is case-sensitive, [code]"*"[/code] matches zero or more characters, and [code]"?"[/code] matches any single character. See also [method find_child] and [method find_children].
[b]Note:[/b] As this method walks upwards in the scene tree, it can be slow in large, deeply nested nodes. Consider storing a reference to the found node in a variable. Alternatively, use [method get_node] with unique names (see [member unique_name_in_owner]).
*/
func (self Instance) FindParent(pattern string) [1]gdclass.Node { //gd:Node.find_parent
	return [1]gdclass.Node(class(self).FindParent(String.New(pattern)))
}

/*
Returns [code]true[/code] if [param path] points to a valid node and its subnames point to a valid [Resource], e.g. [code]Area2D/CollisionShape2D:shape[/code]. Properties that are not [Resource] types (such as nodes or other [Variant] types) are not considered. See also [method get_node_and_resource].
*/
func (self Instance) HasNodeAndResource(path string) bool { //gd:Node.has_node_and_resource
	return bool(class(self).HasNodeAndResource(Path.ToNode(String.New(path))))
}

/*
Fetches a node and its most nested resource as specified by the [NodePath]'s subname. Returns an [Array] of size [code]3[/code] where:
- Element [code]0[/code] is the [Node], or [code]null[/code] if not found;
- Element [code]1[/code] is the subname's last nested [Resource], or [code]null[/code] if not found;
- Element [code]2[/code] is the remaining [NodePath], referring to an existing, non-[Resource] property (see [method Object.get_indexed]).
[b]Example:[/b] Assume that the child's [member Sprite2D.texture] has been assigned a [AtlasTexture]:
[codeblocks]
[gdscript]
var a = get_node_and_resource("Area2D/Sprite2D")
print(a[0].name) # Prints Sprite2D
print(a[1])      # Prints <null>
print(a[2])      # Prints ^""

var b = get_node_and_resource("Area2D/Sprite2D:texture:atlas")
print(b[0].name)        # Prints Sprite2D
print(b[1].get_class()) # Prints AtlasTexture
print(b[2])             # Prints ^""

var c = get_node_and_resource("Area2D/Sprite2D:texture:atlas:region")
print(c[0].name)        # Prints Sprite2D
print(c[1].get_class()) # Prints AtlasTexture
print(c[2])             # Prints ^":region"
[/gdscript]
[csharp]
var a = GetNodeAndResource(NodePath("Area2D/Sprite2D"));
GD.Print(a[0].Name); // Prints Sprite2D
GD.Print(a[1]);      // Prints <null>
GD.Print(a[2]);      // Prints ^"

var b = GetNodeAndResource(NodePath("Area2D/Sprite2D:texture:atlas"));
GD.Print(b[0].name);        // Prints Sprite2D
GD.Print(b[1].get_class()); // Prints AtlasTexture
GD.Print(b[2]);             // Prints ^""

var c = GetNodeAndResource(NodePath("Area2D/Sprite2D:texture:atlas:region"));
GD.Print(c[0].name);        // Prints Sprite2D
GD.Print(c[1].get_class()); // Prints AtlasTexture
GD.Print(c[2]);             // Prints ^":region"
[/csharp]
[/codeblocks]
*/
func (self Instance) GetNodeAndResource(path string) []any { //gd:Node.get_node_and_resource
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetNodeAndResource(Path.ToNode(String.New(path))))))
}

/*
Returns [code]true[/code] if this node is currently inside a [SceneTree]. See also [method get_tree].
*/
func (self Instance) IsInsideTree() bool { //gd:Node.is_inside_tree
	return bool(class(self).IsInsideTree())
}

/*
Returns [code]true[/code] if the node is part of the scene currently opened in the editor.
*/
func (self Instance) IsPartOfEditedScene() bool { //gd:Node.is_part_of_edited_scene
	return bool(class(self).IsPartOfEditedScene())
}

/*
Returns [code]true[/code] if the given [param node] is a direct or indirect child of this node.
*/
func (self Instance) IsAncestorOf(node [1]gdclass.Node) bool { //gd:Node.is_ancestor_of
	return bool(class(self).IsAncestorOf(node))
}

/*
Returns [code]true[/code] if the given [param node] occurs later in the scene hierarchy than this node. A node occurring later is usually processed last.
*/
func (self Instance) IsGreaterThan(node [1]gdclass.Node) bool { //gd:Node.is_greater_than
	return bool(class(self).IsGreaterThan(node))
}

/*
Returns the node's absolute path, relative to the [member SceneTree.root]. If the node is not inside the scene tree, this method fails and returns an empty [NodePath].
*/
func (self Instance) GetPath() string { //gd:Node.get_path
	return string(class(self).GetPath().String())
}

/*
Returns the relative [NodePath] from this node to the specified [param node]. Both nodes must be in the same [SceneTree] or scene hierarchy, otherwise this method fails and returns an empty [NodePath].
If [param use_unique_path] is [code]true[/code], returns the shortest path accounting for this node's unique name (see [member unique_name_in_owner]).
[b]Note:[/b] If you get a relative path which starts from a unique node, the path may be longer than a normal relative path, due to the addition of the unique node's name.
*/
func (self Instance) GetPathTo(node [1]gdclass.Node) string { //gd:Node.get_path_to
	return string(class(self).GetPathTo(node, false).String())
}

/*
Adds the node to the [param group]. Groups can be helpful to organize a subset of nodes, for example [code]"enemies"[/code] or [code]"collectables"[/code]. See notes in the description, and the group methods in [SceneTree].
If [param persistent] is [code]true[/code], the group will be stored when saved inside a [PackedScene]. All groups created and displayed in the Node dock are persistent.
[b]Note:[/b] To improve performance, the order of group names is [i]not[/i] guaranteed and may vary between project runs. Therefore, do not rely on the group order.
[b]Note:[/b] [SceneTree]'s group methods will [i]not[/i] work on this node if not inside the tree (see [method is_inside_tree]).
*/
func (self Instance) AddToGroup(group string) { //gd:Node.add_to_group
	class(self).AddToGroup(String.Name(String.New(group)), false)
}

/*
Removes the node from the given [param group]. Does nothing if the node is not in the [param group]. See also notes in the description, and the [SceneTree]'s group methods.
*/
func (self Instance) RemoveFromGroup(group string) { //gd:Node.remove_from_group
	class(self).RemoveFromGroup(String.Name(String.New(group)))
}

/*
Returns [code]true[/code] if this node has been added to the given [param group]. See [method add_to_group] and [method remove_from_group]. See also notes in the description, and the [SceneTree]'s group methods.
*/
func (self Instance) IsInGroup(group string) bool { //gd:Node.is_in_group
	return bool(class(self).IsInGroup(String.Name(String.New(group))))
}

/*
Moves [param child_node] to the given index. A node's index is the order among its siblings. If [param to_index] is negative, the index is counted from the end of the list. See also [method get_child] and [method get_index].
[b]Note:[/b] The processing order of several engine callbacks ([method _ready], [method _process], etc.) and notifications sent through [method propagate_notification] is affected by tree order. [CanvasItem] nodes are also rendered in tree order. See also [member process_priority].
*/
func (self Instance) MoveChild(child_node [1]gdclass.Node, to_index int) { //gd:Node.move_child
	class(self).MoveChild(child_node, int64(to_index))
}

/*
Returns an [Array] of group names that the node has been added to.
[b]Note:[/b] To improve performance, the order of group names is [i]not[/i] guaranteed and may vary between project runs. Therefore, do not rely on the group order.
[b]Note:[/b] This method may also return some group names starting with an underscore ([code]_[/code]). These are internally used by the engine. To avoid conflicts, do not use custom groups starting with underscores. To exclude internal groups, see the following code snippet:
[codeblocks]
[gdscript]
# Stores the node's non-internal groups only (as an array of StringNames).
var non_internal_groups = []
for group in get_groups():

	if not str(group).begins_with("_"):
	    non_internal_groups.push_back(group)

[/gdscript]
[csharp]
// Stores the node's non-internal groups only (as a List of StringNames).
List<string> nonInternalGroups = new List<string>();
foreach (string group in GetGroups())

	{
	    if (!group.BeginsWith("_"))
	        nonInternalGroups.Add(group);
	}

[/csharp]
[/codeblocks]
*/
func (self Instance) GetGroups() []string { //gd:Node.get_groups
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetGroups())))
}

/*
Returns this node's order among its siblings. The first node's index is [code]0[/code]. See also [method get_child].
If [param include_internal] is [code]false[/code], returns the index ignoring internal children. The first, non-internal child will have an index of [code]0[/code] (see [method add_child]'s [code]internal[/code] parameter).
*/
func (self Instance) GetIndex() int { //gd:Node.get_index
	return int(int(class(self).GetIndex(false)))
}

/*
Prints the node and its children to the console, recursively. The node does not have to be inside the tree. This method outputs [NodePath]s relative to this node, and is good for copy/pasting into [method get_node]. See also [method print_tree_pretty].
May print, for example:
[codeblock lang=text]
.
Menu
Menu/Label
Menu/Camera2D
SplashScreen
SplashScreen/Camera2D
[/codeblock]
*/
func (self Instance) PrintTree() { //gd:Node.print_tree
	class(self).PrintTree()
}

/*
Prints the node and its children to the console, recursively. The node does not have to be inside the tree. Similar to [method print_tree], but the graphical representation looks like what is displayed in the editor's Scene dock. It is useful for inspecting larger trees.
May print, for example:
[codeblock lang=text]

	┖╴TheGame
	   ┠╴Menu
	   ┃  ┠╴Label
	   ┃  ┖╴Camera2D
	   ┖╴SplashScreen
	      ┖╴Camera2D

[/codeblock]
*/
func (self Instance) PrintTreePretty() { //gd:Node.print_tree_pretty
	class(self).PrintTreePretty()
}

/*
Returns the tree as a [String]. Used mainly for debugging purposes. This version displays the path relative to the current node, and is good for copy/pasting into the [method get_node] function. It also can be used in game UI/UX.
May print, for example:
[codeblock lang=text]
TheGame
TheGame/Menu
TheGame/Menu/Label
TheGame/Menu/Camera2D
TheGame/SplashScreen
TheGame/SplashScreen/Camera2D
[/codeblock]
*/
func (self Instance) GetTreeString() string { //gd:Node.get_tree_string
	return string(class(self).GetTreeString().String())
}

/*
Similar to [method get_tree_string], this returns the tree as a [String]. This version displays a more graphical representation similar to what is displayed in the Scene Dock. It is useful for inspecting larger trees.
May print, for example:
[codeblock lang=text]

	┖╴TheGame
	   ┠╴Menu
	   ┃  ┠╴Label
	   ┃  ┖╴Camera2D
	   ┖╴SplashScreen
	      ┖╴Camera2D

[/codeblock]
*/
func (self Instance) GetTreeStringPretty() string { //gd:Node.get_tree_string_pretty
	return string(class(self).GetTreeStringPretty().String())
}

/*
Calls [method Object.notification] with [param what] on this node and all of its children, recursively.
*/
func (self Instance) PropagateNotification(what int) { //gd:Node.propagate_notification
	class(self).PropagateNotification(int64(what))
}

/*
Calls the given [param method] name, passing [param args] as arguments, on this node and all of its children, recursively.
If [param parent_first] is [code]true[/code], the method is called on this node first, then on all of its children. If [code]false[/code], the children's methods are called first.
*/
func (self Instance) PropagateCall(method string) { //gd:Node.propagate_call
	class(self).PropagateCall(String.Name(String.New(method)), Array.Nil, false)
}

/*
If set to [code]true[/code], enables physics (fixed framerate) processing. When a node is being processed, it will receive a [constant NOTIFICATION_PHYSICS_PROCESS] at a fixed (usually 60 FPS, see [member Engine.physics_ticks_per_second] to change) interval (and the [method _physics_process] callback will be called if it exists).
[b]Note:[/b] If [method _physics_process] is overridden, this will be automatically enabled before [method _ready] is called.
*/
func (self Instance) SetPhysicsProcess(enable bool) { //gd:Node.set_physics_process
	class(self).SetPhysicsProcess(enable)
}

/*
Returns the time elapsed (in seconds) since the last physics callback. This value is identical to [method _physics_process]'s [code]delta[/code] parameter, and is often consistent at run-time, unless [member Engine.physics_ticks_per_second] is changed. See also [constant NOTIFICATION_PHYSICS_PROCESS].
*/
func (self Instance) GetPhysicsProcessDeltaTime() Float.X { //gd:Node.get_physics_process_delta_time
	return Float.X(Float.X(class(self).GetPhysicsProcessDeltaTime()))
}

/*
Returns [code]true[/code] if physics processing is enabled (see [method set_physics_process]).
*/
func (self Instance) IsPhysicsProcessing() bool { //gd:Node.is_physics_processing
	return bool(class(self).IsPhysicsProcessing())
}

/*
Returns the time elapsed (in seconds) since the last process callback. This value is identical to [method _process]'s [code]delta[/code] parameter, and may vary from frame to frame. See also [constant NOTIFICATION_PROCESS].
*/
func (self Instance) GetProcessDeltaTime() Float.X { //gd:Node.get_process_delta_time
	return Float.X(Float.X(class(self).GetProcessDeltaTime()))
}

/*
If set to [code]true[/code], enables processing. When a node is being processed, it will receive a [constant NOTIFICATION_PROCESS] on every drawn frame (and the [method _process] callback will be called if it exists).
[b]Note:[/b] If [method _process] is overridden, this will be automatically enabled before [method _ready] is called.
[b]Note:[/b] This method only affects the [method _process] callback, i.e. it has no effect on other callbacks like [method _physics_process]. If you want to disable all processing for the node, set [member process_mode] to [constant PROCESS_MODE_DISABLED].
*/
func (self Instance) SetProcess(enable bool) { //gd:Node.set_process
	class(self).SetProcess(enable)
}

/*
Returns [code]true[/code] if processing is enabled (see [method set_process]).
*/
func (self Instance) IsProcessing() bool { //gd:Node.is_processing
	return bool(class(self).IsProcessing())
}

/*
If set to [code]true[/code], enables input processing.
[b]Note:[/b] If [method _input] is overridden, this will be automatically enabled before [method _ready] is called. Input processing is also already enabled for GUI controls, such as [Button] and [TextEdit].
*/
func (self Instance) SetProcessInput(enable bool) { //gd:Node.set_process_input
	class(self).SetProcessInput(enable)
}

/*
Returns [code]true[/code] if the node is processing input (see [method set_process_input]).
*/
func (self Instance) IsProcessingInput() bool { //gd:Node.is_processing_input
	return bool(class(self).IsProcessingInput())
}

/*
If set to [code]true[/code], enables shortcut processing for this node.
[b]Note:[/b] If [method _shortcut_input] is overridden, this will be automatically enabled before [method _ready] is called.
*/
func (self Instance) SetProcessShortcutInput(enable bool) { //gd:Node.set_process_shortcut_input
	class(self).SetProcessShortcutInput(enable)
}

/*
Returns [code]true[/code] if the node is processing shortcuts (see [method set_process_shortcut_input]).
*/
func (self Instance) IsProcessingShortcutInput() bool { //gd:Node.is_processing_shortcut_input
	return bool(class(self).IsProcessingShortcutInput())
}

/*
If set to [code]true[/code], enables unhandled input processing. It enables the node to receive all input that was not previously handled (usually by a [Control]).
[b]Note:[/b] If [method _unhandled_input] is overridden, this will be automatically enabled before [method _ready] is called. Unhandled input processing is also already enabled for GUI controls, such as [Button] and [TextEdit].
*/
func (self Instance) SetProcessUnhandledInput(enable bool) { //gd:Node.set_process_unhandled_input
	class(self).SetProcessUnhandledInput(enable)
}

/*
Returns [code]true[/code] if the node is processing unhandled input (see [method set_process_unhandled_input]).
*/
func (self Instance) IsProcessingUnhandledInput() bool { //gd:Node.is_processing_unhandled_input
	return bool(class(self).IsProcessingUnhandledInput())
}

/*
If set to [code]true[/code], enables unhandled key input processing.
[b]Note:[/b] If [method _unhandled_key_input] is overridden, this will be automatically enabled before [method _ready] is called.
*/
func (self Instance) SetProcessUnhandledKeyInput(enable bool) { //gd:Node.set_process_unhandled_key_input
	class(self).SetProcessUnhandledKeyInput(enable)
}

/*
Returns [code]true[/code] if the node is processing unhandled key input (see [method set_process_unhandled_key_input]).
*/
func (self Instance) IsProcessingUnhandledKeyInput() bool { //gd:Node.is_processing_unhandled_key_input
	return bool(class(self).IsProcessingUnhandledKeyInput())
}

/*
Returns [code]true[/code] if the node can receive processing notifications and input callbacks ([constant NOTIFICATION_PROCESS], [method _input], etc.) from the [SceneTree] and [Viewport]. The returned value depends on [member process_mode]:
- If set to [constant PROCESS_MODE_PAUSABLE], returns [code]true[/code] when the game is processing, i.e. [member SceneTree.paused] is [code]false[/code];
- If set to [constant PROCESS_MODE_WHEN_PAUSED], returns [code]true[/code] when the game is paused, i.e. [member SceneTree.paused] is [code]true[/code];
- If set to [constant PROCESS_MODE_ALWAYS], always returns [code]true[/code];
- If set to [constant PROCESS_MODE_DISABLED], always returns [code]false[/code];
- If set to [constant PROCESS_MODE_INHERIT], use the parent node's [member process_mode] to determine the result.
If the node is not inside the tree, returns [code]false[/code] no matter the value of [member process_mode].
*/
func (self Instance) CanProcess() bool { //gd:Node.can_process
	return bool(class(self).CanProcess())
}

/*
If set to [code]true[/code], the node appears folded in the Scene dock. As a result, all of its children are hidden. This method is intended to be used in editor plugins and tools, but it also works in release builds. See also [method is_displayed_folded].
*/
func (self Instance) SetDisplayFolded(fold bool) { //gd:Node.set_display_folded
	class(self).SetDisplayFolded(fold)
}

/*
Returns [code]true[/code] if the node is folded (collapsed) in the Scene dock. This method is intended to be used in editor plugins and tools. See also [method set_display_folded].
*/
func (self Instance) IsDisplayedFolded() bool { //gd:Node.is_displayed_folded
	return bool(class(self).IsDisplayedFolded())
}

/*
If set to [code]true[/code], enables internal processing for this node. Internal processing happens in isolation from the normal [method _process] calls and is used by some nodes internally to guarantee proper functioning even if the node is paused or processing is disabled for scripting ([method set_process]).
[b]Warning:[/b] Built-in nodes rely on internal processing for their internal logic. Disabling it is unsafe and may lead to unexpected behavior. Use this method if you know what you are doing.
*/
func (self Instance) SetProcessInternal(enable bool) { //gd:Node.set_process_internal
	class(self).SetProcessInternal(enable)
}

/*
Returns [code]true[/code] if internal processing is enabled (see [method set_process_internal]).
*/
func (self Instance) IsProcessingInternal() bool { //gd:Node.is_processing_internal
	return bool(class(self).IsProcessingInternal())
}

/*
If set to [code]true[/code], enables internal physics for this node. Internal physics processing happens in isolation from the normal [method _physics_process] calls and is used by some nodes internally to guarantee proper functioning even if the node is paused or physics processing is disabled for scripting ([method set_physics_process]).
[b]Warning:[/b] Built-in nodes rely on internal processing for their internal logic. Disabling it is unsafe and may lead to unexpected behavior. Use this method if you know what you are doing.
*/
func (self Instance) SetPhysicsProcessInternal(enable bool) { //gd:Node.set_physics_process_internal
	class(self).SetPhysicsProcessInternal(enable)
}

/*
Returns [code]true[/code] if internal physics processing is enabled (see [method set_physics_process_internal]).
*/
func (self Instance) IsPhysicsProcessingInternal() bool { //gd:Node.is_physics_processing_internal
	return bool(class(self).IsPhysicsProcessingInternal())
}

/*
Returns [code]true[/code] if physics interpolation is enabled for this node (see [member physics_interpolation_mode]).
[b]Note:[/b] Interpolation will only be active if both the flag is set [b]and[/b] physics interpolation is enabled within the [SceneTree]. This can be tested using [method is_physics_interpolated_and_enabled].
*/
func (self Instance) IsPhysicsInterpolated() bool { //gd:Node.is_physics_interpolated
	return bool(class(self).IsPhysicsInterpolated())
}

/*
Returns [code]true[/code] if physics interpolation is enabled (see [member physics_interpolation_mode]) [b]and[/b] enabled in the [SceneTree].
This is a convenience version of [method is_physics_interpolated] that also checks whether physics interpolation is enabled globally.
See [member SceneTree.physics_interpolation] and [member ProjectSettings.physics/common/physics_interpolation].
*/
func (self Instance) IsPhysicsInterpolatedAndEnabled() bool { //gd:Node.is_physics_interpolated_and_enabled
	return bool(class(self).IsPhysicsInterpolatedAndEnabled())
}

/*
When physics interpolation is active, moving a node to a radically different transform (such as placement within a level) can result in a visible glitch as the object is rendered moving from the old to new position over the physics tick.
That glitch can be prevented by calling this method, which temporarily disables interpolation until the physics tick is complete.
The notification [constant NOTIFICATION_RESET_PHYSICS_INTERPOLATION] will be received by the node and all children recursively.
[b]Note:[/b] This function should be called [b]after[/b] moving the node, rather than before.
*/
func (self Instance) ResetPhysicsInterpolation() { //gd:Node.reset_physics_interpolation
	class(self).ResetPhysicsInterpolation()
}

/*
Returns the [Window] that contains this node. If the node is in the main window, this is equivalent to getting the root node ([code]get_tree().get_root()[/code]).
*/
func (self Instance) GetWindow() [1]gdclass.Window { //gd:Node.get_window
	return [1]gdclass.Window(class(self).GetWindow())
}

/*
Returns the [Window] that contains this node, or the last exclusive child in a chain of windows starting with the one that contains this node.
*/
func (self Instance) GetLastExclusiveWindow() [1]gdclass.Window { //gd:Node.get_last_exclusive_window
	return [1]gdclass.Window(class(self).GetLastExclusiveWindow())
}

/*
Returns the [SceneTree] that contains this node. If this node is not inside the tree, generates an error and returns [code]null[/code]. See also [method is_inside_tree].
*/
func (self Instance) GetTree() [1]gdclass.SceneTree { //gd:Node.get_tree
	return [1]gdclass.SceneTree(class(self).GetTree())
}

/*
Creates a new [Tween] and binds it to this node.
This is the equivalent of doing:
[codeblocks]
[gdscript]
get_tree().create_tween().bind_node(self)
[/gdscript]
[csharp]
GetTree().CreateTween().BindNode(this);
[/csharp]
[/codeblocks]
The Tween will start automatically on the next process frame or physics frame (depending on [enum Tween.TweenProcessMode]). See [method Tween.bind_node] for more info on Tweens bound to nodes.
[b]Note:[/b] The method can still be used when the node is not inside [SceneTree]. It can fail in an unlikely case of using a custom [MainLoop].
*/
func (self Instance) CreateTween() [1]gdclass.Tween { //gd:Node.create_tween
	return [1]gdclass.Tween(class(self).CreateTween())
}

/*
Duplicates the node, returning a new node with all of its properties, signals and groups copied from the original. The behavior can be tweaked through the [param flags] (see [enum DuplicateFlags]).
[b]Note:[/b] For nodes with a [Script] attached, if [method Object._init] has been defined with required parameters, the duplicated node will not have a [Script].
*/
func (self Instance) Duplicate() [1]gdclass.Node { //gd:Node.duplicate
	return [1]gdclass.Node(class(self).Duplicate(int64(15)))
}

/*
Replaces this node by the given [param node]. All children of this node are moved to [param node].
If [param keep_groups] is [code]true[/code], the [param node] is added to the same groups that the replaced node is in (see [method add_to_group]).
[b]Warning:[/b] The replaced node is removed from the tree, but it is [b]not[/b] deleted. To prevent memory leaks, store a reference to the node in a variable, or use [method Object.free].
*/
func (self Instance) ReplaceBy(node [1]gdclass.Node) { //gd:Node.replace_by
	class(self).ReplaceBy(node, false)
}

/*
If set to [code]true[/code], the node becomes a [InstancePlaceholder] when packed and instantiated from a [PackedScene]. See also [method get_scene_instance_load_placeholder].
*/
func (self Instance) SetSceneInstanceLoadPlaceholder(load_placeholder bool) { //gd:Node.set_scene_instance_load_placeholder
	class(self).SetSceneInstanceLoadPlaceholder(load_placeholder)
}

/*
Returns [code]true[/code] if this node is an instance load placeholder. See [InstancePlaceholder] and [method set_scene_instance_load_placeholder].
*/
func (self Instance) GetSceneInstanceLoadPlaceholder() bool { //gd:Node.get_scene_instance_load_placeholder
	return bool(class(self).GetSceneInstanceLoadPlaceholder())
}

/*
Set to [code]true[/code] to allow all nodes owned by [param node] to be available, and editable, in the Scene dock, even if their [member owner] is not the scene root. This method is intended to be used in editor plugins and tools, but it also works in release builds. See also [method is_editable_instance].
*/
func (self Instance) SetEditableInstance(node [1]gdclass.Node, is_editable bool) { //gd:Node.set_editable_instance
	class(self).SetEditableInstance(node, is_editable)
}

/*
Returns [code]true[/code] if [param node] has editable children enabled relative to this node. This method is intended to be used in editor plugins and tools. See also [method set_editable_instance].
*/
func (self Instance) IsEditableInstance(node [1]gdclass.Node) bool { //gd:Node.is_editable_instance
	return bool(class(self).IsEditableInstance(node))
}

/*
Returns the node's closest [Viewport] ancestor, if the node is inside the tree. Otherwise, returns [code]null[/code].
*/
func (self Instance) GetViewport() [1]gdclass.Viewport { //gd:Node.get_viewport
	return [1]gdclass.Viewport(class(self).GetViewport())
}

/*
Queues this node to be deleted at the end of the current frame. When deleted, all of its children are deleted as well, and all references to the node and its children become invalid.
Unlike with [method Object.free], the node is not deleted instantly, and it can still be accessed before deletion. It is also safe to call [method queue_free] multiple times. Use [method Object.is_queued_for_deletion] to check if the node will be deleted at the end of the frame.
[b]Note:[/b] The node will only be freed after all other deferred calls are finished. Using this method is not always the same as calling [method Object.free] through [method Object.call_deferred].
*/
func (self Instance) QueueFree() { //gd:Node.queue_free
	class(self).QueueFree()
}

/*
Requests [method _ready] to be called again the next time the node enters the tree. Does [b]not[/b] immediately call [method _ready].
[b]Note:[/b] This method only affects the current node. If the node's children also need to request ready, this method needs to be called for each one of them. When the node and its children enter the tree again, the order of [method _ready] callbacks will be the same as normal.
*/
func (self Instance) RequestReady() { //gd:Node.request_ready
	class(self).RequestReady()
}

/*
Returns [code]true[/code] if the node is ready, i.e. it's inside scene tree and all its children are initialized.
[method request_ready] resets it back to [code]false[/code].
*/
func (self Instance) IsNodeReady() bool { //gd:Node.is_node_ready
	return bool(class(self).IsNodeReady())
}

/*
Sets the node's multiplayer authority to the peer with the given peer [param id]. The multiplayer authority is the peer that has authority over the node on the network. Defaults to peer ID 1 (the server). Useful in conjunction with [method rpc_config] and the [MultiplayerAPI].
If [param recursive] is [code]true[/code], the given peer is recursively set as the authority for all children of this node.
[b]Warning:[/b] This does [b]not[/b] automatically replicate the new authority to other peers. It is the developer's responsibility to do so. You may replicate the new authority's information using [member MultiplayerSpawner.spawn_function], an RPC, or a [MultiplayerSynchronizer]. Furthermore, the parent's authority does [b]not[/b] propagate to newly added children.
*/
func (self Instance) SetMultiplayerAuthority(id int) { //gd:Node.set_multiplayer_authority
	class(self).SetMultiplayerAuthority(int64(id), true)
}

/*
Returns the peer ID of the multiplayer authority for this node. See [method set_multiplayer_authority].
*/
func (self Instance) GetMultiplayerAuthority() int { //gd:Node.get_multiplayer_authority
	return int(int(class(self).GetMultiplayerAuthority()))
}

/*
Returns [code]true[/code] if the local system is the multiplayer authority of this node.
*/
func (self Instance) IsMultiplayerAuthority() bool { //gd:Node.is_multiplayer_authority
	return bool(class(self).IsMultiplayerAuthority())
}

/*
Changes the RPC configuration for the given [param method]. [param config] should either be [code]null[/code] to disable the feature (as by default), or a [Dictionary] containing the following entries:
- [code]rpc_mode[/code]: see [enum MultiplayerAPI.RPCMode];
- [code]transfer_mode[/code]: see [enum MultiplayerPeer.TransferMode];
- [code]call_local[/code]: if [code]true[/code], the method will also be called locally;
- [code]channel[/code]: an [int] representing the channel to send the RPC on.
[b]Note:[/b] In GDScript, this method corresponds to the [annotation @GDScript.@rpc] annotation, with various parameters passed ([code]@rpc(any)[/code], [code]@rpc(authority)[/code]...). See also the [url=$DOCS_URL/tutorials/networking/high_level_multiplayer.html]high-level multiplayer[/url] tutorial.
*/
func (self Instance) RpcConfig(method string, config any) { //gd:Node.rpc_config
	class(self).RpcConfig(String.Name(String.New(method)), variant.New(config))
}

/*
Translates a [param message], using the translation catalogs configured in the Project Settings. Further [param context] can be specified to help with the translation. Note that most [Control] nodes automatically translate their strings, so this method is mostly useful for formatted strings or custom drawn text.
This method works the same as [method Object.tr], with the addition of respecting the [member auto_translate_mode] state.
If [method Object.can_translate_messages] is [code]false[/code], or no translation is available, this method returns the [param message] without changes. See [method Object.set_message_translation].
For detailed examples, see [url=$DOCS_URL/tutorials/i18n/internationalizing_games.html]Internationalizing games[/url].
*/
func (self Instance) Atr(message string) string { //gd:Node.atr
	return string(class(self).Atr(String.New(message), String.Name(String.New(""))).String())
}

/*
Translates a [param message] or [param plural_message], using the translation catalogs configured in the Project Settings. Further [param context] can be specified to help with the translation.
This method works the same as [method Object.tr_n], with the addition of respecting the [member auto_translate_mode] state.
If [method Object.can_translate_messages] is [code]false[/code], or no translation is available, this method returns [param message] or [param plural_message], without changes. See [method Object.set_message_translation].
The [param n] is the number, or amount, of the message's subject. It is used by the translation system to fetch the correct plural form for the current language.
For detailed examples, see [url=$DOCS_URL/tutorials/i18n/localization_using_gettext.html]Localization using gettext[/url].
[b]Note:[/b] Negative and [float] numbers may not properly apply to some countable subjects. It's recommended to handle these cases with [method atr].
*/
func (self Instance) AtrN(message string, plural_message string, n int) string { //gd:Node.atr_n
	return string(class(self).AtrN(String.New(message), String.Name(String.New(plural_message)), int64(n), String.Name(String.New(""))).String())
}

/*
Refreshes the warnings displayed for this node in the Scene dock. Use [method _get_configuration_warnings] to customize the warning messages to display.
*/
func (self Instance) UpdateConfigurationWarnings() { //gd:Node.update_configuration_warnings
	class(self).UpdateConfigurationWarnings()
}

/*
Similar to [method call_deferred_thread_group], but for setting properties.
*/
func (self Instance) SetDeferredThreadGroup(property string, value any) { //gd:Node.set_deferred_thread_group
	class(self).SetDeferredThreadGroup(String.Name(String.New(property)), variant.New(value))
}

/*
Similar to [method call_deferred_thread_group], but for notifications.
*/
func (self Instance) NotifyDeferredThreadGroup(what int) { //gd:Node.notify_deferred_thread_group
	class(self).NotifyDeferredThreadGroup(int64(what))
}

/*
Similar to [method call_thread_safe], but for setting properties.
*/
func (self Instance) SetThreadSafe(property string, value any) { //gd:Node.set_thread_safe
	class(self).SetThreadSafe(String.Name(String.New(property)), variant.New(value))
}

/*
Similar to [method call_thread_safe], but for notifications.
*/
func (self Instance) NotifyThreadSafe(what int) { //gd:Node.notify_thread_safe
	class(self).NotifyThreadSafe(int64(what))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Node

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Node"))
	casted := Instance{*(*gdclass.Node)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Name() string {
	return string(class(self).GetName().String())
}

func (self Instance) SetName(value string) {
	class(self).SetName(String.New(value))
}

func (self Instance) UniqueNameInOwner() bool {
	return bool(class(self).IsUniqueNameInOwner())
}

func (self Instance) SetUniqueNameInOwner(value bool) {
	class(self).SetUniqueNameInOwner(value)
}

func (self Instance) SceneFilePath() string {
	return string(class(self).GetSceneFilePath().String())
}

func (self Instance) SetSceneFilePath(value string) {
	class(self).SetSceneFilePath(String.New(value))
}

func (self Instance) Owner() [1]gdclass.Node {
	return [1]gdclass.Node(class(self).GetOwner())
}

func (self Instance) SetOwner(value [1]gdclass.Node) {
	class(self).SetOwner(value)
}

func (self Instance) Multiplayer() [1]gdclass.MultiplayerAPI {
	return [1]gdclass.MultiplayerAPI(class(self).GetMultiplayer())
}

func (self Instance) ProcessMode() gdclass.NodeProcessMode {
	return gdclass.NodeProcessMode(class(self).GetProcessMode())
}

func (self Instance) SetProcessMode(value gdclass.NodeProcessMode) {
	class(self).SetProcessMode(value)
}

func (self Instance) ProcessPriority() int {
	return int(int(class(self).GetProcessPriority()))
}

func (self Instance) SetProcessPriority(value int) {
	class(self).SetProcessPriority(int64(value))
}

func (self Instance) ProcessPhysicsPriority() int {
	return int(int(class(self).GetPhysicsProcessPriority()))
}

func (self Instance) SetProcessPhysicsPriority(value int) {
	class(self).SetPhysicsProcessPriority(int64(value))
}

func (self Instance) ProcessThreadGroup() gdclass.NodeProcessThreadGroup {
	return gdclass.NodeProcessThreadGroup(class(self).GetProcessThreadGroup())
}

func (self Instance) SetProcessThreadGroup(value gdclass.NodeProcessThreadGroup) {
	class(self).SetProcessThreadGroup(value)
}

func (self Instance) ProcessThreadGroupOrder() int {
	return int(int(class(self).GetProcessThreadGroupOrder()))
}

func (self Instance) SetProcessThreadGroupOrder(value int) {
	class(self).SetProcessThreadGroupOrder(int64(value))
}

func (self Instance) ProcessThreadMessages() gdclass.NodeProcessThreadMessages {
	return gdclass.NodeProcessThreadMessages(class(self).GetProcessThreadMessages())
}

func (self Instance) SetProcessThreadMessages(value gdclass.NodeProcessThreadMessages) {
	class(self).SetProcessThreadMessages(value)
}

func (self Instance) PhysicsInterpolationMode() gdclass.NodePhysicsInterpolationMode {
	return gdclass.NodePhysicsInterpolationMode(class(self).GetPhysicsInterpolationMode())
}

func (self Instance) SetPhysicsInterpolationMode(value gdclass.NodePhysicsInterpolationMode) {
	class(self).SetPhysicsInterpolationMode(value)
}

func (self Instance) AutoTranslateMode() gdclass.NodeAutoTranslateMode {
	return gdclass.NodeAutoTranslateMode(class(self).GetAutoTranslateMode())
}

func (self Instance) SetAutoTranslateMode(value gdclass.NodeAutoTranslateMode) {
	class(self).SetAutoTranslateMode(value)
}

func (self Instance) EditorDescription() string {
	return string(class(self).GetEditorDescription().String())
}

func (self Instance) SetEditorDescription(value string) {
	class(self).SetEditorDescription(String.New(value))
}

/*
Called during the processing step of the main loop. Processing happens at every frame and as fast as possible, so the [param delta] time since the previous frame is not constant. [param delta] is in seconds.
It is only called if processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process].
Corresponds to the [constant NOTIFICATION_PROCESS] notification in [method Object._notification].
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
*/
func (class) _process(impl func(ptr unsafe.Pointer, delta float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var delta = gd.UnsafeGet[float64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, delta)
	}
}

/*
Called during the physics processing step of the main loop. Physics processing means that the frame rate is synced to the physics, i.e. the [param delta] variable should be constant. [param delta] is in seconds.
It is only called if physics processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_physics_process].
Corresponds to the [constant NOTIFICATION_PHYSICS_PROCESS] notification in [method Object._notification].
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
*/
func (class) _physics_process(impl func(ptr unsafe.Pointer, delta float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var delta = gd.UnsafeGet[float64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, delta)
	}
}

/*
Called when the node enters the [SceneTree] (e.g. upon instantiating, scene changing, or after calling [method add_child] in a script). If the node has children, its [method _enter_tree] callback will be called first, and then that of the children.
Corresponds to the [constant NOTIFICATION_ENTER_TREE] notification in [method Object._notification].
*/
func (class) _enter_tree(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the node is about to leave the [SceneTree] (e.g. upon freeing, scene changing, or after calling [method remove_child] in a script). If the node has children, its [method _exit_tree] callback will be called last, after all its children have left the tree.
Corresponds to the [constant NOTIFICATION_EXIT_TREE] notification in [method Object._notification] and signal [signal tree_exiting]. To get notified when the node has already left the active tree, connect to the [signal tree_exited].
*/
func (class) _exit_tree(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the node is "ready", i.e. when both the node and its children have entered the scene tree. If the node has children, their [method _ready] callbacks get triggered first, and the parent node will receive the ready notification afterwards.
Corresponds to the [constant NOTIFICATION_READY] notification in [method Object._notification]. See also the [code]@onready[/code] annotation for variables.
Usually used for initialization. For even earlier initialization, [method Object._init] may be used. See also [method _enter_tree].
[b]Note:[/b] This method may be called only once for each node. After removing a node from the scene tree and adding it again, [method _ready] will [b]not[/b] be called a second time. This can be bypassed by requesting another call with [method request_ready], which may be called anywhere before adding the node again.
*/
func (class) _ready(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
The elements in the array returned from this method are displayed as warnings in the Scene dock if the script that overrides it is a [code]tool[/code] script.
Returning an empty array produces no warnings.
Call [method update_configuration_warnings] when the warnings need to be updated for this node.
[codeblock]
@export var energy = 0:

	set(value):
	    energy = value
	    update_configuration_warnings()

func _get_configuration_warnings():

	if energy < 0:
	    return ["Energy must be 0 or greater."]
	else:
	    return []

[/codeblock]
*/
func (class) _get_configuration_warnings(impl func(ptr unsafe.Pointer) Packed.Strings) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Called when there is an input event. The input event propagates up through the node tree until a node consumes it.
It is only called if input processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_input].
To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
For gameplay input, [method _unhandled_input] and [method _unhandled_key_input] are usually a better fit as they allow the GUI to intercept the events first.
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
*/
func (class) _input(impl func(ptr unsafe.Pointer, event [1]gdclass.InputEvent)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, event)
	}
}

/*
Called when an [InputEventKey], [InputEventShortcut], or [InputEventJoypadButton] hasn't been consumed by [method _input] or any GUI [Control] item. It is called before [method _unhandled_key_input] and [method _unhandled_input]. The input event propagates up through the node tree until a node consumes it.
It is only called if shortcut processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_shortcut_input].
To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
This method can be used to handle shortcuts. For generic GUI events, use [method _input] instead. Gameplay events should usually be handled with either [method _unhandled_input] or [method _unhandled_key_input].
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not orphan).
*/
func (class) _shortcut_input(impl func(ptr unsafe.Pointer, event [1]gdclass.InputEvent)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, event)
	}
}

/*
Called when an [InputEvent] hasn't been consumed by [method _input] or any GUI [Control] item. It is called after [method _shortcut_input] and after [method _unhandled_key_input]. The input event propagates up through the node tree until a node consumes it.
It is only called if unhandled input processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_unhandled_input].
To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
For gameplay input, this method is usually a better fit than [method _input], as GUI events need a higher priority. For keyboard shortcuts, consider using [method _shortcut_input] instead, as it is called before this method. Finally, to handle keyboard events, consider using [method _unhandled_key_input] for performance reasons.
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
*/
func (class) _unhandled_input(impl func(ptr unsafe.Pointer, event [1]gdclass.InputEvent)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, event)
	}
}

/*
Called when an [InputEventKey] hasn't been consumed by [method _input] or any GUI [Control] item. It is called after [method _shortcut_input] but before [method _unhandled_input]. The input event propagates up through the node tree until a node consumes it.
It is only called if unhandled key input processing is enabled, which is done automatically if this method is overridden, and can be toggled with [method set_process_unhandled_key_input].
To consume the input event and stop it propagating further to other nodes, [method Viewport.set_input_as_handled] can be called.
This method can be used to handle Unicode character input with [kbd]Alt[/kbd], [kbd]Alt + Ctrl[/kbd], and [kbd]Alt + Shift[/kbd] modifiers, after shortcuts were handled.
For gameplay input, this and [method _unhandled_input] are usually a better fit than [method _input], as GUI events should be handled first. This method also performs better than [method _unhandled_input], since unrelated events such as [InputEventMouseMotion] are automatically filtered. For shortcuts, consider using [method _shortcut_input] instead.
[b]Note:[/b] This method is only called if the node is present in the scene tree (i.e. if it's not an orphan).
*/
func (class) _unhandled_key_input(impl func(ptr unsafe.Pointer, event [1]gdclass.InputEvent)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, event)
	}
}

/*
Prints all orphan nodes (nodes outside the [SceneTree]). Useful for debugging.
[b]Note:[/b] This method only works in debug builds. Does nothing in a project exported in release mode.
*/
//go:nosplit
func (self class) PrintOrphanNodes() { //gd:Node.print_orphan_nodes
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_print_orphan_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [param sibling] node to this node's parent, and moves the added sibling right below this node.
If [param force_readable_name] is [code]true[/code], improves the readability of the added [param sibling]. If not named, the [param sibling] is renamed to its type, and if it shares [member name] with a sibling, a number is suffixed more appropriately. This operation is very slow. As such, it is recommended leaving this to [code]false[/code], which assigns a dummy name featuring [code]@[/code] in both situations.
Use [method add_child] instead of this method if you don't need the child node to be added below a specific node in the list of children.
[b]Note:[/b] If this node is internal, the added sibling will be internal too (see [method add_child]'s [code]internal[/code] parameter).
*/
//go:nosplit
func (self class) AddSibling(sibling [1]gdclass.Node, force_readable_name bool) { //gd:Node.add_sibling
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(sibling[0].AsObject()[0]))
	callframe.Arg(frame, force_readable_name)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_add_sibling, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetName(name String.Readable) { //gd:Node.set_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetName() String.Name { //gd:Node.get_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Adds a child [param node]. Nodes can have any number of children, but every child must have a unique name. Child nodes are automatically deleted when the parent node is deleted, so an entire scene can be removed by deleting its topmost node.
If [param force_readable_name] is [code]true[/code], improves the readability of the added [param node]. If not named, the [param node] is renamed to its type, and if it shares [member name] with a sibling, a number is suffixed more appropriately. This operation is very slow. As such, it is recommended leaving this to [code]false[/code], which assigns a dummy name featuring [code]@[/code] in both situations.
If [param internal] is different than [constant INTERNAL_MODE_DISABLED], the child will be added as internal node. These nodes are ignored by methods like [method get_children], unless their parameter [code]include_internal[/code] is [code]true[/code]. The intended usage is to hide the internal nodes from the user, so the user won't accidentally delete or modify them. Used by some GUI nodes, e.g. [ColorPicker]. See [enum InternalMode] for available modes.
[b]Note:[/b] If [param node] already has a parent, this method will fail. Use [method remove_child] first to remove [param node] from its current parent. For example:
[codeblocks]
[gdscript]
var child_node = get_child(0)
if child_node.get_parent():
    child_node.get_parent().remove_child(child_node)
add_child(child_node)
[/gdscript]
[csharp]
Node childNode = GetChild(0);
if (childNode.GetParent() != null)
{
    childNode.GetParent().RemoveChild(childNode);
}
AddChild(childNode);
[/csharp]
[/codeblocks]
If you need the child node to be added below a specific node in the list of children, use [method add_sibling] instead of this method.
[b]Note:[/b] If you want a child to be persisted to a [PackedScene], you must set [member owner] in addition to calling [method add_child]. This is typically relevant for [url=$DOCS_URL/tutorials/plugins/running_code_in_the_editor.html]tool scripts[/url] and [url=$DOCS_URL/tutorials/plugins/editor/index.html]editor plugins[/url]. If [method add_child] is called without setting [member owner], the newly added [Node] will not be visible in the scene tree, though it will be visible in the 2D/3D view.
*/
//go:nosplit
func (self class) AddChild(node [1]gdclass.Node, force_readable_name bool, internal_ gdclass.NodeInternalMode) { //gd:Node.add_child
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(node[0].AsObject()[0]))
	callframe.Arg(frame, force_readable_name)
	callframe.Arg(frame, internal_)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_add_child, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a child [param node]. The [param node], along with its children, are [b]not[/b] deleted. To delete a node, see [method queue_free].
[b]Note:[/b] When this node is inside the tree, this method sets the [member owner] of the removed [param node] (or its descendants) to [code]null[/code], if their [member owner] is no longer an ancestor (see [method is_ancestor_of]).
*/
//go:nosplit
func (self class) RemoveChild(node [1]gdclass.Node) { //gd:Node.remove_child
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_remove_child, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Changes the parent of this [Node] to the [param new_parent]. The node needs to already have a parent. The node's [member owner] is preserved if its owner is still reachable from the new location (i.e., the node is still a descendant of the new parent after the operation).
If [param keep_global_transform] is [code]true[/code], the node's global transform will be preserved if supported. [Node2D], [Node3D] and [Control] support this argument (but [Control] keeps only position).
*/
//go:nosplit
func (self class) Reparent(new_parent [1]gdclass.Node, keep_global_transform bool) { //gd:Node.reparent
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(new_parent[0])[0])
	callframe.Arg(frame, keep_global_transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_reparent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of children of this node.
If [param include_internal] is [code]false[/code], internal children are not counted (see [method add_child]'s [code]internal[/code] parameter).
*/
//go:nosplit
func (self class) GetChildCount(include_internal bool) int64 { //gd:Node.get_child_count
	var frame = callframe.New()
	callframe.Arg(frame, include_internal)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_child_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns all children of this node inside an [Array].
If [param include_internal] is [code]false[/code], excludes internal children from the returned array (see [method add_child]'s [code]internal[/code] parameter).
*/
//go:nosplit
func (self class) GetChildren(include_internal bool) Array.Contains[[1]gdclass.Node] { //gd:Node.get_children
	var frame = callframe.New()
	callframe.Arg(frame, include_internal)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_children, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Node]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Fetches a child node by its index. Each child node has an index relative its siblings (see [method get_index]). The first child is at index 0. Negative values can also be used to start from the end of the list. This method can be used in combination with [method get_child_count] to iterate over this node's children. If no child exists at the given index, this method returns [code]null[/code] and an error is generated.
If [param include_internal] is [code]false[/code], internal children are ignored (see [method add_child]'s [code]internal[/code] parameter).
[codeblock]
# Assuming the following are children of this node, in order:
# First, Middle, Last.

var a = get_child(0).name  # a is "First"
var b = get_child(1).name  # b is "Middle"
var b = get_child(2).name  # b is "Last"
var c = get_child(-1).name # c is "Last"
[/codeblock]
[b]Note:[/b] To fetch a node by [NodePath], use [method get_node].
*/
//go:nosplit
func (self class) GetChild(idx int64, include_internal bool) [1]gdclass.Node { //gd:Node.get_child
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, include_internal)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_child, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [param path] points to a valid node. See also [method get_node].
*/
//go:nosplit
func (self class) HasNode(path Path.ToNode) bool { //gd:Node.has_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_has_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Fetches a node. The [NodePath] can either be a relative path (from this node), or an absolute path (from the [member SceneTree.root]) to a node. If [param path] does not point to a valid node, generates an error and returns [code]null[/code]. Attempts to access methods on the return value will result in an [i]"Attempt to call <method> on a null instance."[/i] error.
[b]Note:[/b] Fetching by absolute path only works when the node is inside the scene tree (see [method is_inside_tree]).
[b]Example:[/b] Assume this method is called from the Character node, inside the following tree:
[codeblock lang=text]
 ┖╴root
    ┠╴Character (you are here!)
    ┃  ┠╴Sword
    ┃  ┖╴Backpack
    ┃     ┖╴Dagger
    ┠╴MyGame
    ┖╴Swamp
       ┠╴Alligator
       ┠╴Mosquito
       ┖╴Goblin
[/codeblock]
The following calls will return a valid node:
[codeblocks]
[gdscript]
get_node("Sword")
get_node("Backpack/Dagger")
get_node("../Swamp/Alligator")
get_node("/root/MyGame")
[/gdscript]
[csharp]
GetNode("Sword");
GetNode("Backpack/Dagger");
GetNode("../Swamp/Alligator");
GetNode("/root/MyGame");
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetNode(path Path.ToNode) [1]gdclass.Node { //gd:Node.get_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Fetches a node by [NodePath]. Similar to [method get_node], but does not generate an error if [param path] does not point to a valid node.
*/
//go:nosplit
func (self class) GetNodeOrNull(path Path.ToNode) [1]gdclass.Node { //gd:Node.get_node_or_null
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_node_or_null, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns this node's parent node, or [code]null[/code] if the node doesn't have a parent.
*/
//go:nosplit
func (self class) GetParent() [1]gdclass.Node { //gd:Node.get_parent
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Finds the first descendant of this node whose [member name] matches [param pattern], returning [code]null[/code] if no match is found. The matching is done against node names, [i]not[/i] their paths, through [method String.match]. As such, it is case-sensitive, [code]"*"[/code] matches zero or more characters, and [code]"?"[/code] matches any single character.
If [param recursive] is [code]false[/code], only this node's direct children are checked. Nodes are checked in tree order, so this node's first direct child is checked first, then its own direct children, etc., before moving to the second direct child, and so on. Internal children are also included in the search (see [code]internal[/code] parameter in [method add_child]).
If [param owned] is [code]true[/code], only descendants with a valid [member owner] node are checked.
[b]Note:[/b] This method can be very slow. Consider storing a reference to the found node in a variable. Alternatively, use [method get_node] with unique names (see [member unique_name_in_owner]).
[b]Note:[/b] To find all descendant nodes matching a pattern or a class type, see [method find_children].
*/
//go:nosplit
func (self class) FindChild(pattern String.Readable, recursive bool, owned bool) [1]gdclass.Node { //gd:Node.find_child
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(pattern)))
	callframe.Arg(frame, recursive)
	callframe.Arg(frame, owned)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_find_child, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Finds all descendants of this node whose names match [param pattern], returning an empty [Array] if no match is found. The matching is done against node names, [i]not[/i] their paths, through [method String.match]. As such, it is case-sensitive, [code]"*"[/code] matches zero or more characters, and [code]"?"[/code] matches any single character.
If [param type] is not empty, only ancestors inheriting from [param type] are included (see [method Object.is_class]).
If [param recursive] is [code]false[/code], only this node's direct children are checked. Nodes are checked in tree order, so this node's first direct child is checked first, then its own direct children, etc., before moving to the second direct child, and so on. Internal children are also included in the search (see [code]internal[/code] parameter in [method add_child]).
If [param owned] is [code]true[/code], only descendants with a valid [member owner] node are checked.
[b]Note:[/b] This method can be very slow. Consider storing references to the found nodes in a variable.
[b]Note:[/b] To find a single descendant node matching a pattern, see [method find_child].
*/
//go:nosplit
func (self class) FindChildren(pattern String.Readable, atype String.Readable, recursive bool, owned bool) Array.Contains[[1]gdclass.Node] { //gd:Node.find_children
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(pattern)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(atype)))
	callframe.Arg(frame, recursive)
	callframe.Arg(frame, owned)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_find_children, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Node]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Finds the first ancestor of this node whose [member name] matches [param pattern], returning [code]null[/code] if no match is found. The matching is done through [method String.match]. As such, it is case-sensitive, [code]"*"[/code] matches zero or more characters, and [code]"?"[/code] matches any single character. See also [method find_child] and [method find_children].
[b]Note:[/b] As this method walks upwards in the scene tree, it can be slow in large, deeply nested nodes. Consider storing a reference to the found node in a variable. Alternatively, use [method get_node] with unique names (see [member unique_name_in_owner]).
*/
//go:nosplit
func (self class) FindParent(pattern String.Readable) [1]gdclass.Node { //gd:Node.find_parent
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(pattern)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_find_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [param path] points to a valid node and its subnames point to a valid [Resource], e.g. [code]Area2D/CollisionShape2D:shape[/code]. Properties that are not [Resource] types (such as nodes or other [Variant] types) are not considered. See also [method get_node_and_resource].
*/
//go:nosplit
func (self class) HasNodeAndResource(path Path.ToNode) bool { //gd:Node.has_node_and_resource
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_has_node_and_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Fetches a node and its most nested resource as specified by the [NodePath]'s subname. Returns an [Array] of size [code]3[/code] where:
- Element [code]0[/code] is the [Node], or [code]null[/code] if not found;
- Element [code]1[/code] is the subname's last nested [Resource], or [code]null[/code] if not found;
- Element [code]2[/code] is the remaining [NodePath], referring to an existing, non-[Resource] property (see [method Object.get_indexed]).
[b]Example:[/b] Assume that the child's [member Sprite2D.texture] has been assigned a [AtlasTexture]:
[codeblocks]
[gdscript]
var a = get_node_and_resource("Area2D/Sprite2D")
print(a[0].name) # Prints Sprite2D
print(a[1])      # Prints <null>
print(a[2])      # Prints ^""

var b = get_node_and_resource("Area2D/Sprite2D:texture:atlas")
print(b[0].name)        # Prints Sprite2D
print(b[1].get_class()) # Prints AtlasTexture
print(b[2])             # Prints ^""

var c = get_node_and_resource("Area2D/Sprite2D:texture:atlas:region")
print(c[0].name)        # Prints Sprite2D
print(c[1].get_class()) # Prints AtlasTexture
print(c[2])             # Prints ^":region"
[/gdscript]
[csharp]
var a = GetNodeAndResource(NodePath("Area2D/Sprite2D"));
GD.Print(a[0].Name); // Prints Sprite2D
GD.Print(a[1]);      // Prints <null>
GD.Print(a[2]);      // Prints ^"

var b = GetNodeAndResource(NodePath("Area2D/Sprite2D:texture:atlas"));
GD.Print(b[0].name);        // Prints Sprite2D
GD.Print(b[1].get_class()); // Prints AtlasTexture
GD.Print(b[2]);             // Prints ^""

var c = GetNodeAndResource(NodePath("Area2D/Sprite2D:texture:atlas:region"));
GD.Print(c[0].name);        // Prints Sprite2D
GD.Print(c[1].get_class()); // Prints AtlasTexture
GD.Print(c[2]);             // Prints ^":region"
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetNodeAndResource(path Path.ToNode) Array.Any { //gd:Node.get_node_and_resource
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_node_and_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this node is currently inside a [SceneTree]. See also [method get_tree].
*/
//go:nosplit
func (self class) IsInsideTree() bool { //gd:Node.is_inside_tree
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_inside_tree, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the node is part of the scene currently opened in the editor.
*/
//go:nosplit
func (self class) IsPartOfEditedScene() bool { //gd:Node.is_part_of_edited_scene
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_part_of_edited_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [param node] is a direct or indirect child of this node.
*/
//go:nosplit
func (self class) IsAncestorOf(node [1]gdclass.Node) bool { //gd:Node.is_ancestor_of
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_ancestor_of, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [param node] occurs later in the scene hierarchy than this node. A node occurring later is usually processed last.
*/
//go:nosplit
func (self class) IsGreaterThan(node [1]gdclass.Node) bool { //gd:Node.is_greater_than
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_greater_than, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the node's absolute path, relative to the [member SceneTree.root]. If the node is not inside the scene tree, this method fails and returns an empty [NodePath].
*/
//go:nosplit
func (self class) GetPath() Path.ToNode { //gd:Node.get_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the relative [NodePath] from this node to the specified [param node]. Both nodes must be in the same [SceneTree] or scene hierarchy, otherwise this method fails and returns an empty [NodePath].
If [param use_unique_path] is [code]true[/code], returns the shortest path accounting for this node's unique name (see [member unique_name_in_owner]).
[b]Note:[/b] If you get a relative path which starts from a unique node, the path may be longer than a normal relative path, due to the addition of the unique node's name.
*/
//go:nosplit
func (self class) GetPathTo(node [1]gdclass.Node, use_unique_path bool) Path.ToNode { //gd:Node.get_path_to
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, use_unique_path)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_path_to, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Adds the node to the [param group]. Groups can be helpful to organize a subset of nodes, for example [code]"enemies"[/code] or [code]"collectables"[/code]. See notes in the description, and the group methods in [SceneTree].
If [param persistent] is [code]true[/code], the group will be stored when saved inside a [PackedScene]. All groups created and displayed in the Node dock are persistent.
[b]Note:[/b] To improve performance, the order of group names is [i]not[/i] guaranteed and may vary between project runs. Therefore, do not rely on the group order.
[b]Note:[/b] [SceneTree]'s group methods will [i]not[/i] work on this node if not inside the tree (see [method is_inside_tree]).
*/
//go:nosplit
func (self class) AddToGroup(group String.Name, persistent bool) { //gd:Node.add_to_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(group)))
	callframe.Arg(frame, persistent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_add_to_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the node from the given [param group]. Does nothing if the node is not in the [param group]. See also notes in the description, and the [SceneTree]'s group methods.
*/
//go:nosplit
func (self class) RemoveFromGroup(group String.Name) { //gd:Node.remove_from_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(group)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_remove_from_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if this node has been added to the given [param group]. See [method add_to_group] and [method remove_from_group]. See also notes in the description, and the [SceneTree]'s group methods.
*/
//go:nosplit
func (self class) IsInGroup(group String.Name) bool { //gd:Node.is_in_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(group)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_in_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves [param child_node] to the given index. A node's index is the order among its siblings. If [param to_index] is negative, the index is counted from the end of the list. See also [method get_child] and [method get_index].
[b]Note:[/b] The processing order of several engine callbacks ([method _ready], [method _process], etc.) and notifications sent through [method propagate_notification] is affected by tree order. [CanvasItem] nodes are also rendered in tree order. See also [member process_priority].
*/
//go:nosplit
func (self class) MoveChild(child_node [1]gdclass.Node, to_index int64) { //gd:Node.move_child
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(child_node[0])[0])
	callframe.Arg(frame, to_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_move_child, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an [Array] of group names that the node has been added to.
[b]Note:[/b] To improve performance, the order of group names is [i]not[/i] guaranteed and may vary between project runs. Therefore, do not rely on the group order.
[b]Note:[/b] This method may also return some group names starting with an underscore ([code]_[/code]). These are internally used by the engine. To avoid conflicts, do not use custom groups starting with underscores. To exclude internal groups, see the following code snippet:
[codeblocks]
[gdscript]
# Stores the node's non-internal groups only (as an array of StringNames).
var non_internal_groups = []
for group in get_groups():
    if not str(group).begins_with("_"):
        non_internal_groups.push_back(group)
[/gdscript]
[csharp]
// Stores the node's non-internal groups only (as a List of StringNames).
List<string> nonInternalGroups = new List<string>();
foreach (string group in GetGroups())
{
    if (!group.BeginsWith("_"))
        nonInternalGroups.Add(group);
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetGroups() Array.Contains[String.Name] { //gd:Node.get_groups
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_groups, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Name]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOwner(owner [1]gdclass.Node) { //gd:Node.set_owner
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(owner[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOwner() [1]gdclass.Node { //gd:Node.get_owner
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns this node's order among its siblings. The first node's index is [code]0[/code]. See also [method get_child].
If [param include_internal] is [code]false[/code], returns the index ignoring internal children. The first, non-internal child will have an index of [code]0[/code] (see [method add_child]'s [code]internal[/code] parameter).
*/
//go:nosplit
func (self class) GetIndex(include_internal bool) int64 { //gd:Node.get_index
	var frame = callframe.New()
	callframe.Arg(frame, include_internal)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Prints the node and its children to the console, recursively. The node does not have to be inside the tree. This method outputs [NodePath]s relative to this node, and is good for copy/pasting into [method get_node]. See also [method print_tree_pretty].
May print, for example:
[codeblock lang=text]
.
Menu
Menu/Label
Menu/Camera2D
SplashScreen
SplashScreen/Camera2D
[/codeblock]
*/
//go:nosplit
func (self class) PrintTree() { //gd:Node.print_tree
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_print_tree, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Prints the node and its children to the console, recursively. The node does not have to be inside the tree. Similar to [method print_tree], but the graphical representation looks like what is displayed in the editor's Scene dock. It is useful for inspecting larger trees.
May print, for example:
[codeblock lang=text]
 ┖╴TheGame
    ┠╴Menu
    ┃  ┠╴Label
    ┃  ┖╴Camera2D
    ┖╴SplashScreen
       ┖╴Camera2D
[/codeblock]
*/
//go:nosplit
func (self class) PrintTreePretty() { //gd:Node.print_tree_pretty
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_print_tree_pretty, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the tree as a [String]. Used mainly for debugging purposes. This version displays the path relative to the current node, and is good for copy/pasting into the [method get_node] function. It also can be used in game UI/UX.
May print, for example:
[codeblock lang=text]
TheGame
TheGame/Menu
TheGame/Menu/Label
TheGame/Menu/Camera2D
TheGame/SplashScreen
TheGame/SplashScreen/Camera2D
[/codeblock]
*/
//go:nosplit
func (self class) GetTreeString() String.Readable { //gd:Node.get_tree_string
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_tree_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Similar to [method get_tree_string], this returns the tree as a [String]. This version displays a more graphical representation similar to what is displayed in the Scene Dock. It is useful for inspecting larger trees.
May print, for example:
[codeblock lang=text]
 ┖╴TheGame
    ┠╴Menu
    ┃  ┠╴Label
    ┃  ┖╴Camera2D
    ┖╴SplashScreen
       ┖╴Camera2D
[/codeblock]
*/
//go:nosplit
func (self class) GetTreeStringPretty() String.Readable { //gd:Node.get_tree_string_pretty
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_tree_string_pretty, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSceneFilePath(scene_file_path String.Readable) { //gd:Node.set_scene_file_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(scene_file_path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_scene_file_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSceneFilePath() String.Readable { //gd:Node.get_scene_file_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_scene_file_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Calls [method Object.notification] with [param what] on this node and all of its children, recursively.
*/
//go:nosplit
func (self class) PropagateNotification(what int64) { //gd:Node.propagate_notification
	var frame = callframe.New()
	callframe.Arg(frame, what)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_propagate_notification, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Calls the given [param method] name, passing [param args] as arguments, on this node and all of its children, recursively.
If [param parent_first] is [code]true[/code], the method is called on this node first, then on all of its children. If [code]false[/code], the children's methods are called first.
*/
//go:nosplit
func (self class) PropagateCall(method String.Name, args Array.Any, parent_first bool) { //gd:Node.propagate_call
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(method)))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(args)))
	callframe.Arg(frame, parent_first)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_propagate_call, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If set to [code]true[/code], enables physics (fixed framerate) processing. When a node is being processed, it will receive a [constant NOTIFICATION_PHYSICS_PROCESS] at a fixed (usually 60 FPS, see [member Engine.physics_ticks_per_second] to change) interval (and the [method _physics_process] callback will be called if it exists).
[b]Note:[/b] If [method _physics_process] is overridden, this will be automatically enabled before [method _ready] is called.
*/
//go:nosplit
func (self class) SetPhysicsProcess(enable bool) { //gd:Node.set_physics_process
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_physics_process, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the time elapsed (in seconds) since the last physics callback. This value is identical to [method _physics_process]'s [code]delta[/code] parameter, and is often consistent at run-time, unless [member Engine.physics_ticks_per_second] is changed. See also [constant NOTIFICATION_PHYSICS_PROCESS].
*/
//go:nosplit
func (self class) GetPhysicsProcessDeltaTime() float64 { //gd:Node.get_physics_process_delta_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_physics_process_delta_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if physics processing is enabled (see [method set_physics_process]).
*/
//go:nosplit
func (self class) IsPhysicsProcessing() bool { //gd:Node.is_physics_processing
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_physics_processing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the time elapsed (in seconds) since the last process callback. This value is identical to [method _process]'s [code]delta[/code] parameter, and may vary from frame to frame. See also [constant NOTIFICATION_PROCESS].
*/
//go:nosplit
func (self class) GetProcessDeltaTime() float64 { //gd:Node.get_process_delta_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_process_delta_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], enables processing. When a node is being processed, it will receive a [constant NOTIFICATION_PROCESS] on every drawn frame (and the [method _process] callback will be called if it exists).
[b]Note:[/b] If [method _process] is overridden, this will be automatically enabled before [method _ready] is called.
[b]Note:[/b] This method only affects the [method _process] callback, i.e. it has no effect on other callbacks like [method _physics_process]. If you want to disable all processing for the node, set [member process_mode] to [constant PROCESS_MODE_DISABLED].
*/
//go:nosplit
func (self class) SetProcess(enable bool) { //gd:Node.set_process
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetProcessPriority(priority int64) { //gd:Node.set_process_priority
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProcessPriority() int64 { //gd:Node.get_process_priority
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_process_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsProcessPriority(priority int64) { //gd:Node.set_physics_process_priority
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_physics_process_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsProcessPriority() int64 { //gd:Node.get_physics_process_priority
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_physics_process_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if processing is enabled (see [method set_process]).
*/
//go:nosplit
func (self class) IsProcessing() bool { //gd:Node.is_processing
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_processing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], enables input processing.
[b]Note:[/b] If [method _input] is overridden, this will be automatically enabled before [method _ready] is called. Input processing is also already enabled for GUI controls, such as [Button] and [TextEdit].
*/
//go:nosplit
func (self class) SetProcessInput(enable bool) { //gd:Node.set_process_input
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process_input, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the node is processing input (see [method set_process_input]).
*/
//go:nosplit
func (self class) IsProcessingInput() bool { //gd:Node.is_processing_input
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_processing_input, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], enables shortcut processing for this node.
[b]Note:[/b] If [method _shortcut_input] is overridden, this will be automatically enabled before [method _ready] is called.
*/
//go:nosplit
func (self class) SetProcessShortcutInput(enable bool) { //gd:Node.set_process_shortcut_input
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process_shortcut_input, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the node is processing shortcuts (see [method set_process_shortcut_input]).
*/
//go:nosplit
func (self class) IsProcessingShortcutInput() bool { //gd:Node.is_processing_shortcut_input
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_processing_shortcut_input, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], enables unhandled input processing. It enables the node to receive all input that was not previously handled (usually by a [Control]).
[b]Note:[/b] If [method _unhandled_input] is overridden, this will be automatically enabled before [method _ready] is called. Unhandled input processing is also already enabled for GUI controls, such as [Button] and [TextEdit].
*/
//go:nosplit
func (self class) SetProcessUnhandledInput(enable bool) { //gd:Node.set_process_unhandled_input
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process_unhandled_input, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the node is processing unhandled input (see [method set_process_unhandled_input]).
*/
//go:nosplit
func (self class) IsProcessingUnhandledInput() bool { //gd:Node.is_processing_unhandled_input
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_processing_unhandled_input, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], enables unhandled key input processing.
[b]Note:[/b] If [method _unhandled_key_input] is overridden, this will be automatically enabled before [method _ready] is called.
*/
//go:nosplit
func (self class) SetProcessUnhandledKeyInput(enable bool) { //gd:Node.set_process_unhandled_key_input
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process_unhandled_key_input, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the node is processing unhandled key input (see [method set_process_unhandled_key_input]).
*/
//go:nosplit
func (self class) IsProcessingUnhandledKeyInput() bool { //gd:Node.is_processing_unhandled_key_input
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_processing_unhandled_key_input, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProcessMode(mode gdclass.NodeProcessMode) { //gd:Node.set_process_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProcessMode() gdclass.NodeProcessMode { //gd:Node.get_process_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NodeProcessMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_process_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the node can receive processing notifications and input callbacks ([constant NOTIFICATION_PROCESS], [method _input], etc.) from the [SceneTree] and [Viewport]. The returned value depends on [member process_mode]:
- If set to [constant PROCESS_MODE_PAUSABLE], returns [code]true[/code] when the game is processing, i.e. [member SceneTree.paused] is [code]false[/code];
- If set to [constant PROCESS_MODE_WHEN_PAUSED], returns [code]true[/code] when the game is paused, i.e. [member SceneTree.paused] is [code]true[/code];
- If set to [constant PROCESS_MODE_ALWAYS], always returns [code]true[/code];
- If set to [constant PROCESS_MODE_DISABLED], always returns [code]false[/code];
- If set to [constant PROCESS_MODE_INHERIT], use the parent node's [member process_mode] to determine the result.
If the node is not inside the tree, returns [code]false[/code] no matter the value of [member process_mode].
*/
//go:nosplit
func (self class) CanProcess() bool { //gd:Node.can_process
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_can_process, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProcessThreadGroup(mode gdclass.NodeProcessThreadGroup) { //gd:Node.set_process_thread_group
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process_thread_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProcessThreadGroup() gdclass.NodeProcessThreadGroup { //gd:Node.get_process_thread_group
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NodeProcessThreadGroup](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_process_thread_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProcessThreadMessages(flags gdclass.NodeProcessThreadMessages) { //gd:Node.set_process_thread_messages
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process_thread_messages, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProcessThreadMessages() gdclass.NodeProcessThreadMessages { //gd:Node.get_process_thread_messages
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NodeProcessThreadMessages](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_process_thread_messages, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProcessThreadGroupOrder(order int64) { //gd:Node.set_process_thread_group_order
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process_thread_group_order, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProcessThreadGroupOrder() int64 { //gd:Node.get_process_thread_group_order
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_process_thread_group_order, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], the node appears folded in the Scene dock. As a result, all of its children are hidden. This method is intended to be used in editor plugins and tools, but it also works in release builds. See also [method is_displayed_folded].
*/
//go:nosplit
func (self class) SetDisplayFolded(fold bool) { //gd:Node.set_display_folded
	var frame = callframe.New()
	callframe.Arg(frame, fold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_display_folded, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the node is folded (collapsed) in the Scene dock. This method is intended to be used in editor plugins and tools. See also [method set_display_folded].
*/
//go:nosplit
func (self class) IsDisplayedFolded() bool { //gd:Node.is_displayed_folded
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_displayed_folded, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], enables internal processing for this node. Internal processing happens in isolation from the normal [method _process] calls and is used by some nodes internally to guarantee proper functioning even if the node is paused or processing is disabled for scripting ([method set_process]).
[b]Warning:[/b] Built-in nodes rely on internal processing for their internal logic. Disabling it is unsafe and may lead to unexpected behavior. Use this method if you know what you are doing.
*/
//go:nosplit
func (self class) SetProcessInternal(enable bool) { //gd:Node.set_process_internal
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_process_internal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if internal processing is enabled (see [method set_process_internal]).
*/
//go:nosplit
func (self class) IsProcessingInternal() bool { //gd:Node.is_processing_internal
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_processing_internal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], enables internal physics for this node. Internal physics processing happens in isolation from the normal [method _physics_process] calls and is used by some nodes internally to guarantee proper functioning even if the node is paused or physics processing is disabled for scripting ([method set_physics_process]).
[b]Warning:[/b] Built-in nodes rely on internal processing for their internal logic. Disabling it is unsafe and may lead to unexpected behavior. Use this method if you know what you are doing.
*/
//go:nosplit
func (self class) SetPhysicsProcessInternal(enable bool) { //gd:Node.set_physics_process_internal
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_physics_process_internal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if internal physics processing is enabled (see [method set_physics_process_internal]).
*/
//go:nosplit
func (self class) IsPhysicsProcessingInternal() bool { //gd:Node.is_physics_processing_internal
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_physics_processing_internal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsInterpolationMode(mode gdclass.NodePhysicsInterpolationMode) { //gd:Node.set_physics_interpolation_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_physics_interpolation_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsInterpolationMode() gdclass.NodePhysicsInterpolationMode { //gd:Node.get_physics_interpolation_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NodePhysicsInterpolationMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_physics_interpolation_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if physics interpolation is enabled for this node (see [member physics_interpolation_mode]).
[b]Note:[/b] Interpolation will only be active if both the flag is set [b]and[/b] physics interpolation is enabled within the [SceneTree]. This can be tested using [method is_physics_interpolated_and_enabled].
*/
//go:nosplit
func (self class) IsPhysicsInterpolated() bool { //gd:Node.is_physics_interpolated
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_physics_interpolated, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if physics interpolation is enabled (see [member physics_interpolation_mode]) [b]and[/b] enabled in the [SceneTree].
This is a convenience version of [method is_physics_interpolated] that also checks whether physics interpolation is enabled globally.
See [member SceneTree.physics_interpolation] and [member ProjectSettings.physics/common/physics_interpolation].
*/
//go:nosplit
func (self class) IsPhysicsInterpolatedAndEnabled() bool { //gd:Node.is_physics_interpolated_and_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_physics_interpolated_and_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
When physics interpolation is active, moving a node to a radically different transform (such as placement within a level) can result in a visible glitch as the object is rendered moving from the old to new position over the physics tick.
That glitch can be prevented by calling this method, which temporarily disables interpolation until the physics tick is complete.
The notification [constant NOTIFICATION_RESET_PHYSICS_INTERPOLATION] will be received by the node and all children recursively.
[b]Note:[/b] This function should be called [b]after[/b] moving the node, rather than before.
*/
//go:nosplit
func (self class) ResetPhysicsInterpolation() { //gd:Node.reset_physics_interpolation
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_reset_physics_interpolation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAutoTranslateMode(mode gdclass.NodeAutoTranslateMode) { //gd:Node.set_auto_translate_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_auto_translate_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoTranslateMode() gdclass.NodeAutoTranslateMode { //gd:Node.get_auto_translate_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NodeAutoTranslateMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_auto_translate_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Window] that contains this node. If the node is in the main window, this is equivalent to getting the root node ([code]get_tree().get_root()[/code]).
*/
//go:nosplit
func (self class) GetWindow() [1]gdclass.Window { //gd:Node.get_window
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_window, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Window{gd.PointerMustAssertInstanceID[gdclass.Window](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [Window] that contains this node, or the last exclusive child in a chain of windows starting with the one that contains this node.
*/
//go:nosplit
func (self class) GetLastExclusiveWindow() [1]gdclass.Window { //gd:Node.get_last_exclusive_window
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_last_exclusive_window, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Window{gd.PointerMustAssertInstanceID[gdclass.Window](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [SceneTree] that contains this node. If this node is not inside the tree, generates an error and returns [code]null[/code]. See also [method is_inside_tree].
*/
//go:nosplit
func (self class) GetTree() [1]gdclass.SceneTree { //gd:Node.get_tree
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_tree, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SceneTree{gd.PointerMustAssertInstanceID[gdclass.SceneTree](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new [Tween] and binds it to this node.
This is the equivalent of doing:
[codeblocks]
[gdscript]
get_tree().create_tween().bind_node(self)
[/gdscript]
[csharp]
GetTree().CreateTween().BindNode(this);
[/csharp]
[/codeblocks]
The Tween will start automatically on the next process frame or physics frame (depending on [enum Tween.TweenProcessMode]). See [method Tween.bind_node] for more info on Tweens bound to nodes.
[b]Note:[/b] The method can still be used when the node is not inside [SceneTree]. It can fail in an unlikely case of using a custom [MainLoop].
*/
//go:nosplit
func (self class) CreateTween() [1]gdclass.Tween { //gd:Node.create_tween
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_create_tween, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Duplicates the node, returning a new node with all of its properties, signals and groups copied from the original. The behavior can be tweaked through the [param flags] (see [enum DuplicateFlags]).
[b]Note:[/b] For nodes with a [Script] attached, if [method Object._init] has been defined with required parameters, the duplicated node will not have a [Script].
*/
//go:nosplit
func (self class) Duplicate(flags int64) [1]gdclass.Node { //gd:Node.duplicate
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_duplicate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerWithOwnershipTransferredToGo[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Replaces this node by the given [param node]. All children of this node are moved to [param node].
If [param keep_groups] is [code]true[/code], the [param node] is added to the same groups that the replaced node is in (see [method add_to_group]).
[b]Warning:[/b] The replaced node is removed from the tree, but it is [b]not[/b] deleted. To prevent memory leaks, store a reference to the node in a variable, or use [method Object.free].
*/
//go:nosplit
func (self class) ReplaceBy(node [1]gdclass.Node, keep_groups bool) { //gd:Node.replace_by
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(node[0].AsObject()[0]))
	callframe.Arg(frame, keep_groups)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_replace_by, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If set to [code]true[/code], the node becomes a [InstancePlaceholder] when packed and instantiated from a [PackedScene]. See also [method get_scene_instance_load_placeholder].
*/
//go:nosplit
func (self class) SetSceneInstanceLoadPlaceholder(load_placeholder bool) { //gd:Node.set_scene_instance_load_placeholder
	var frame = callframe.New()
	callframe.Arg(frame, load_placeholder)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_scene_instance_load_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if this node is an instance load placeholder. See [InstancePlaceholder] and [method set_scene_instance_load_placeholder].
*/
//go:nosplit
func (self class) GetSceneInstanceLoadPlaceholder() bool { //gd:Node.get_scene_instance_load_placeholder
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_scene_instance_load_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set to [code]true[/code] to allow all nodes owned by [param node] to be available, and editable, in the Scene dock, even if their [member owner] is not the scene root. This method is intended to be used in editor plugins and tools, but it also works in release builds. See also [method is_editable_instance].
*/
//go:nosplit
func (self class) SetEditableInstance(node [1]gdclass.Node, is_editable bool) { //gd:Node.set_editable_instance
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, is_editable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_editable_instance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if [param node] has editable children enabled relative to this node. This method is intended to be used in editor plugins and tools. See also [method set_editable_instance].
*/
//go:nosplit
func (self class) IsEditableInstance(node [1]gdclass.Node) bool { //gd:Node.is_editable_instance
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_editable_instance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the node's closest [Viewport] ancestor, if the node is inside the tree. Otherwise, returns [code]null[/code].
*/
//go:nosplit
func (self class) GetViewport() [1]gdclass.Viewport { //gd:Node.get_viewport
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_viewport, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Viewport{gd.PointerMustAssertInstanceID[gdclass.Viewport](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Queues this node to be deleted at the end of the current frame. When deleted, all of its children are deleted as well, and all references to the node and its children become invalid.
Unlike with [method Object.free], the node is not deleted instantly, and it can still be accessed before deletion. It is also safe to call [method queue_free] multiple times. Use [method Object.is_queued_for_deletion] to check if the node will be deleted at the end of the frame.
[b]Note:[/b] The node will only be freed after all other deferred calls are finished. Using this method is not always the same as calling [method Object.free] through [method Object.call_deferred].
*/
//go:nosplit
func (self class) QueueFree() { //gd:Node.queue_free
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_queue_free, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Requests [method _ready] to be called again the next time the node enters the tree. Does [b]not[/b] immediately call [method _ready].
[b]Note:[/b] This method only affects the current node. If the node's children also need to request ready, this method needs to be called for each one of them. When the node and its children enter the tree again, the order of [method _ready] callbacks will be the same as normal.
*/
//go:nosplit
func (self class) RequestReady() { //gd:Node.request_ready
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_request_ready, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the node is ready, i.e. it's inside scene tree and all its children are initialized.
[method request_ready] resets it back to [code]false[/code].
*/
//go:nosplit
func (self class) IsNodeReady() bool { //gd:Node.is_node_ready
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_node_ready, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the node's multiplayer authority to the peer with the given peer [param id]. The multiplayer authority is the peer that has authority over the node on the network. Defaults to peer ID 1 (the server). Useful in conjunction with [method rpc_config] and the [MultiplayerAPI].
If [param recursive] is [code]true[/code], the given peer is recursively set as the authority for all children of this node.
[b]Warning:[/b] This does [b]not[/b] automatically replicate the new authority to other peers. It is the developer's responsibility to do so. You may replicate the new authority's information using [member MultiplayerSpawner.spawn_function], an RPC, or a [MultiplayerSynchronizer]. Furthermore, the parent's authority does [b]not[/b] propagate to newly added children.
*/
//go:nosplit
func (self class) SetMultiplayerAuthority(id int64, recursive bool) { //gd:Node.set_multiplayer_authority
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, recursive)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_multiplayer_authority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the peer ID of the multiplayer authority for this node. See [method set_multiplayer_authority].
*/
//go:nosplit
func (self class) GetMultiplayerAuthority() int64 { //gd:Node.get_multiplayer_authority
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_multiplayer_authority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the local system is the multiplayer authority of this node.
*/
//go:nosplit
func (self class) IsMultiplayerAuthority() bool { //gd:Node.is_multiplayer_authority
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_multiplayer_authority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMultiplayer() [1]gdclass.MultiplayerAPI { //gd:Node.get_multiplayer
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_multiplayer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.MultiplayerAPI{gd.PointerWithOwnershipTransferredToGo[gdclass.MultiplayerAPI](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Changes the RPC configuration for the given [param method]. [param config] should either be [code]null[/code] to disable the feature (as by default), or a [Dictionary] containing the following entries:
- [code]rpc_mode[/code]: see [enum MultiplayerAPI.RPCMode];
- [code]transfer_mode[/code]: see [enum MultiplayerPeer.TransferMode];
- [code]call_local[/code]: if [code]true[/code], the method will also be called locally;
- [code]channel[/code]: an [int] representing the channel to send the RPC on.
[b]Note:[/b] In GDScript, this method corresponds to the [annotation @GDScript.@rpc] annotation, with various parameters passed ([code]@rpc(any)[/code], [code]@rpc(authority)[/code]...). See also the [url=$DOCS_URL/tutorials/networking/high_level_multiplayer.html]high-level multiplayer[/url] tutorial.
*/
//go:nosplit
func (self class) RpcConfig(method String.Name, config variant.Any) { //gd:Node.rpc_config
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(method)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(config)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_rpc_config, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetEditorDescription(editor_description String.Readable) { //gd:Node.set_editor_description
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(editor_description)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_editor_description, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEditorDescription() String.Readable { //gd:Node.get_editor_description
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_get_editor_description, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUniqueNameInOwner(enable bool) { //gd:Node.set_unique_name_in_owner
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_unique_name_in_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUniqueNameInOwner() bool { //gd:Node.is_unique_name_in_owner
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_is_unique_name_in_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Translates a [param message], using the translation catalogs configured in the Project Settings. Further [param context] can be specified to help with the translation. Note that most [Control] nodes automatically translate their strings, so this method is mostly useful for formatted strings or custom drawn text.
This method works the same as [method Object.tr], with the addition of respecting the [member auto_translate_mode] state.
If [method Object.can_translate_messages] is [code]false[/code], or no translation is available, this method returns the [param message] without changes. See [method Object.set_message_translation].
For detailed examples, see [url=$DOCS_URL/tutorials/i18n/internationalizing_games.html]Internationalizing games[/url].
*/
//go:nosplit
func (self class) Atr(message String.Readable, context String.Name) String.Readable { //gd:Node.atr
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(message)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(context)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_atr, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Translates a [param message] or [param plural_message], using the translation catalogs configured in the Project Settings. Further [param context] can be specified to help with the translation.
This method works the same as [method Object.tr_n], with the addition of respecting the [member auto_translate_mode] state.
If [method Object.can_translate_messages] is [code]false[/code], or no translation is available, this method returns [param message] or [param plural_message], without changes. See [method Object.set_message_translation].
The [param n] is the number, or amount, of the message's subject. It is used by the translation system to fetch the correct plural form for the current language.
For detailed examples, see [url=$DOCS_URL/tutorials/i18n/localization_using_gettext.html]Localization using gettext[/url].
[b]Note:[/b] Negative and [float] numbers may not properly apply to some countable subjects. It's recommended to handle these cases with [method atr].
*/
//go:nosplit
func (self class) AtrN(message String.Readable, plural_message String.Name, n int64, context String.Name) String.Readable { //gd:Node.atr_n
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(message)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(plural_message)))
	callframe.Arg(frame, n)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(context)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_atr_n, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Refreshes the warnings displayed for this node in the Scene dock. Use [method _get_configuration_warnings] to customize the warning messages to display.
*/
//go:nosplit
func (self class) UpdateConfigurationWarnings() { //gd:Node.update_configuration_warnings
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_update_configuration_warnings, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Similar to [method call_deferred_thread_group], but for setting properties.
*/
//go:nosplit
func (self class) SetDeferredThreadGroup(property String.Name, value variant.Any) { //gd:Node.set_deferred_thread_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(property)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_deferred_thread_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Similar to [method call_deferred_thread_group], but for notifications.
*/
//go:nosplit
func (self class) NotifyDeferredThreadGroup(what int64) { //gd:Node.notify_deferred_thread_group
	var frame = callframe.New()
	callframe.Arg(frame, what)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_notify_deferred_thread_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Similar to [method call_thread_safe], but for setting properties.
*/
//go:nosplit
func (self class) SetThreadSafe(property String.Name, value variant.Any) { //gd:Node.set_thread_safe
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(property)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_set_thread_safe, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Similar to [method call_thread_safe], but for notifications.
*/
//go:nosplit
func (self class) NotifyThreadSafe(what int64) { //gd:Node.notify_thread_safe
	var frame = callframe.New()
	callframe.Arg(frame, what)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node.Bind_notify_thread_safe, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnReady(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("ready"), gd.NewCallable(cb), 0)
}

func (self Instance) OnRenamed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("renamed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTreeEntered(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tree_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTreeExiting(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tree_exiting"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTreeExited(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tree_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnChildEnteredTree(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("child_entered_tree"), gd.NewCallable(cb), 0)
}

func (self Instance) OnChildExitingTree(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("child_exiting_tree"), gd.NewCallable(cb), 0)
}

func (self Instance) OnChildOrderChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("child_order_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnReplacingBy(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("replacing_by"), gd.NewCallable(cb), 0)
}

func (self Instance) OnEditorDescriptionChanged(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("editor_description_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsNode() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_process":
		return reflect.ValueOf(self._process)
	case "_physics_process":
		return reflect.ValueOf(self._physics_process)
	case "_enter_tree":
		return reflect.ValueOf(self._enter_tree)
	case "_exit_tree":
		return reflect.ValueOf(self._exit_tree)
	case "_ready":
		return reflect.ValueOf(self._ready)
	case "_get_configuration_warnings":
		return reflect.ValueOf(self._get_configuration_warnings)
	case "_input":
		return reflect.ValueOf(self._input)
	case "_shortcut_input":
		return reflect.ValueOf(self._shortcut_input)
	case "_unhandled_input":
		return reflect.ValueOf(self._unhandled_input)
	case "_unhandled_key_input":
		return reflect.ValueOf(self._unhandled_key_input)
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_process":
		return reflect.ValueOf(self._process)
	case "_physics_process":
		return reflect.ValueOf(self._physics_process)
	case "_enter_tree":
		return reflect.ValueOf(self._enter_tree)
	case "_exit_tree":
		return reflect.ValueOf(self._exit_tree)
	case "_ready":
		return reflect.ValueOf(self._ready)
	case "_get_configuration_warnings":
		return reflect.ValueOf(self._get_configuration_warnings)
	case "_input":
		return reflect.ValueOf(self._input)
	case "_shortcut_input":
		return reflect.ValueOf(self._shortcut_input)
	case "_unhandled_input":
		return reflect.ValueOf(self._unhandled_input)
	case "_unhandled_key_input":
		return reflect.ValueOf(self._unhandled_key_input)
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("Node", func(ptr gd.Object) any { return [1]gdclass.Node{*(*gdclass.Node)(unsafe.Pointer(&ptr))} })
}

type ProcessMode = gdclass.NodeProcessMode //gd:Node.ProcessMode

const (
	/*Inherits [member process_mode] from the node's parent. This is the default for any newly created node.*/
	ProcessModeInherit ProcessMode = 0
	/*Stops processing when [member SceneTree.paused] is [code]true[/code]. This is the inverse of [constant PROCESS_MODE_WHEN_PAUSED], and the default for the root node.*/
	ProcessModePausable ProcessMode = 1
	/*Process [b]only[/b] when [member SceneTree.paused] is [code]true[/code]. This is the inverse of [constant PROCESS_MODE_PAUSABLE].*/
	ProcessModeWhenPaused ProcessMode = 2
	/*Always process. Keeps processing, ignoring [member SceneTree.paused]. This is the inverse of [constant PROCESS_MODE_DISABLED].*/
	ProcessModeAlways ProcessMode = 3
	/*Never process. Completely disables processing, ignoring [member SceneTree.paused]. This is the inverse of [constant PROCESS_MODE_ALWAYS].*/
	ProcessModeDisabled ProcessMode = 4
)

type ProcessThreadGroup = gdclass.NodeProcessThreadGroup //gd:Node.ProcessThreadGroup

const (
	/*Process this node based on the thread group mode of the first parent (or grandparent) node that has a thread group mode that is not inherit. See [member process_thread_group] for more information.*/
	ProcessThreadGroupInherit ProcessThreadGroup = 0
	/*Process this node (and child nodes set to inherit) on the main thread. See [member process_thread_group] for more information.*/
	ProcessThreadGroupMainThread ProcessThreadGroup = 1
	/*Process this node (and child nodes set to inherit) on a sub-thread. See [member process_thread_group] for more information.*/
	ProcessThreadGroupSubThread ProcessThreadGroup = 2
)

type ProcessThreadMessages = gdclass.NodeProcessThreadMessages //gd:Node.ProcessThreadMessages

const (
	/*Allows this node to process threaded messages created with [method call_deferred_thread_group] right before [method _process] is called.*/
	FlagProcessThreadMessages ProcessThreadMessages = 1
	/*Allows this node to process threaded messages created with [method call_deferred_thread_group] right before [method _physics_process] is called.*/
	FlagProcessThreadMessagesPhysics ProcessThreadMessages = 2
	/*Allows this node to process threaded messages created with [method call_deferred_thread_group] right before either [method _process] or [method _physics_process] are called.*/
	FlagProcessThreadMessagesAll ProcessThreadMessages = 3
)

type PhysicsInterpolationMode = gdclass.NodePhysicsInterpolationMode //gd:Node.PhysicsInterpolationMode

const (
	/*Inherits [member physics_interpolation_mode] from the node's parent. This is the default for any newly created node.*/
	PhysicsInterpolationModeInherit PhysicsInterpolationMode = 0
	/*Enables physics interpolation for this node and for children set to [constant PHYSICS_INTERPOLATION_MODE_INHERIT]. This is the default for the root node.*/
	PhysicsInterpolationModeOn PhysicsInterpolationMode = 1
	/*Disables physics interpolation for this node and for children set to [constant PHYSICS_INTERPOLATION_MODE_INHERIT].*/
	PhysicsInterpolationModeOff PhysicsInterpolationMode = 2
)

type DuplicateFlags = gdclass.NodeDuplicateFlags //gd:Node.DuplicateFlags

const (
	/*Duplicate the node's signal connections.*/
	DuplicateSignals DuplicateFlags = 1
	/*Duplicate the node's groups.*/
	DuplicateGroups DuplicateFlags = 2
	/*Duplicate the node's script (also overriding the duplicated children's scripts, if combined with [constant DUPLICATE_USE_INSTANTIATION]).*/
	DuplicateScripts DuplicateFlags = 4
	/*Duplicate using [method PackedScene.instantiate]. If the node comes from a scene saved on disk, re-uses [method PackedScene.instantiate] as the base for the duplicated node and its children.*/
	DuplicateUseInstantiation DuplicateFlags = 8
)

type InternalMode = gdclass.NodeInternalMode //gd:Node.InternalMode

const (
	/*The node will not be internal.*/
	InternalModeDisabled InternalMode = 0
	/*The node will be placed at the beginning of the parent's children, before any non-internal sibling.*/
	InternalModeFront InternalMode = 1
	/*The node will be placed at the end of the parent's children, after any non-internal sibling.*/
	InternalModeBack InternalMode = 2
)

type AutoTranslateMode = gdclass.NodeAutoTranslateMode //gd:Node.AutoTranslateMode

const (
	/*Inherits [member auto_translate_mode] from the node's parent. This is the default for any newly created node.*/
	AutoTranslateModeInherit AutoTranslateMode = 0
	/*Always automatically translate. This is the inverse of [constant AUTO_TRANSLATE_MODE_DISABLED], and the default for the root node.*/
	AutoTranslateModeAlways AutoTranslateMode = 1
	/*Never automatically translate. This is the inverse of [constant AUTO_TRANSLATE_MODE_ALWAYS].
	  String parsing for POT generation will be skipped for this node and children that are set to [constant AUTO_TRANSLATE_MODE_INHERIT].*/
	AutoTranslateModeDisabled AutoTranslateMode = 2
)
