// Package SceneTree provides methods for working with SceneTree object instances.
package SceneTree

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
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/MainLoop"
import "graphics.gd/variant/Float"

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

/*
As one of the most important classes, the [SceneTree] manages the hierarchy of nodes in a scene, as well as scenes themselves. Nodes can be added, fetched and removed. The whole scene tree (and thus the current scene) can be paused. Scenes can be loaded, switched and reloaded.
You can also use the [SceneTree] to organize your nodes into [b]groups[/b]: every node can be added to as many groups as you want to create, e.g. an "enemy" group. You can then iterate these groups or even call methods and set properties on all the nodes belonging to any given group.
[SceneTree] is the default [MainLoop] implementation used by the engine, and is thus in charge of the game loop.
*/
type Instance [1]gdclass.SceneTree

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSceneTree() Instance
}

/*
Returns [code]true[/code] if a node added to the given group [param name] exists in the tree.
*/
func (self Instance) HasGroup(name string) bool { //gd:SceneTree.has_group
	return bool(class(self).HasGroup(String.Name(String.New(name))))
}

/*
Returns a new [SceneTreeTimer]. After [param time_sec] in seconds have passed, the timer will emit [signal SceneTreeTimer.timeout] and will be automatically freed.
If [param process_always] is [code]false[/code], the timer will be paused when setting [member SceneTree.paused] to [code]true[/code].
If [param process_in_physics] is [code]true[/code], the timer will update at the end of the physics frame, instead of the process frame.
If [param ignore_time_scale] is [code]true[/code], the timer will ignore [member Engine.time_scale] and update with the real, elapsed time.
This method is commonly used to create a one-shot delay timer, as in the following example:
[codeblocks]
[gdscript]
func some_function():

	print("start")
	await get_tree().create_timer(1.0).timeout
	print("end")

[/gdscript]
[csharp]
public async Task SomeFunction()

	{
	    GD.Print("start");
	    await ToSignal(GetTree().CreateTimer(1.0f), SceneTreeTimer.SignalName.Timeout);
	    GD.Print("end");
	}

[/csharp]
[/codeblocks]
[b]Note:[/b] The timer is always updated [i]after[/i] all of the nodes in the tree. A node's [method Node._process] method would be called before the timer updates (or [method Node._physics_process] if [param process_in_physics] is set to [code]true[/code]).
*/
func (self Instance) CreateTimer(time_sec Float.X) [1]gdclass.SceneTreeTimer { //gd:SceneTree.create_timer
	return [1]gdclass.SceneTreeTimer(class(self).CreateTimer(gd.Float(time_sec), true, false, false))
}

/*
Creates and returns a new [Tween] processed in this tree. The Tween will start automatically on the next process frame or physics frame (depending on its [enum Tween.TweenProcessMode]).
[b]Note:[/b] A [Tween] created using this method is not bound to any [Node]. It may keep working until there is nothing left to animate. If you want the [Tween] to be automatically killed when the [Node] is freed, use [method Node.create_tween] or [method Tween.bind_node].
*/
func (self Instance) CreateTween() [1]gdclass.Tween { //gd:SceneTree.create_tween
	return [1]gdclass.Tween(class(self).CreateTween())
}

/*
Returns an [Array] of currently existing [Tween]s in the tree, including paused tweens.
*/
func (self Instance) GetProcessedTweens() [][1]gdclass.Tween { //gd:SceneTree.get_processed_tweens
	return [][1]gdclass.Tween(gd.ArrayAs[[][1]gdclass.Tween](gd.InternalArray(class(self).GetProcessedTweens())))
}

/*
Returns the number of nodes inside this tree.
*/
func (self Instance) GetNodeCount() int { //gd:SceneTree.get_node_count
	return int(int(class(self).GetNodeCount()))
}

/*
Returns how many frames have been processed, since the application started. This is [i]not[/i] a measurement of elapsed time.
*/
func (self Instance) GetFrame() int { //gd:SceneTree.get_frame
	return int(int(class(self).GetFrame()))
}

/*
Quits the application at the end of the current iteration, with the given [param exit_code].
By convention, an exit code of [code]0[/code] indicates success, whereas any other exit code indicates an error. For portability reasons, it should be between [code]0[/code] and [code]125[/code] (inclusive).
[b]Note:[/b] On iOS this method doesn't work. Instead, as recommended by the [url=https://developer.apple.com/library/archive/qa/qa1561/_index.html]iOS Human Interface Guidelines[/url], the user is expected to close apps via the Home button.
*/
func (self Instance) Quit() { //gd:SceneTree.quit
	class(self).Quit(gd.Int(0))
}

/*
Queues the given [param obj] to be deleted, calling its [method Object.free] at the end of the current frame. This method is similar to [method Node.queue_free].
*/
func (self Instance) QueueDelete(obj Object.Instance) { //gd:SceneTree.queue_delete
	class(self).QueueDelete(obj)
}

/*
Calls [method Object.notification] with the given [param notification] to all nodes inside this tree added to the [param group]. Use [param call_flags] to customize this method's behavior (see [enum GroupCallFlags]).
*/
func (self Instance) NotifyGroupFlags(call_flags int, group string, notification int) { //gd:SceneTree.notify_group_flags
	class(self).NotifyGroupFlags(gd.Int(call_flags), String.Name(String.New(group)), gd.Int(notification))
}

/*
Sets the given [param property] to [param value] on all nodes inside this tree added to the given [param group]. Nodes that do not have the [param property] are ignored. Use [param call_flags] to customize this method's behavior (see [enum GroupCallFlags]).
[b]Note:[/b] In C#, [param property] must be in snake_case when referring to built-in Godot properties. Prefer using the names exposed in the [code]PropertyName[/code] class to avoid allocating a new [StringName] on each call.
*/
func (self Instance) SetGroupFlags(call_flags int, group string, property string, value any) { //gd:SceneTree.set_group_flags
	class(self).SetGroupFlags(gd.Int(call_flags), String.Name(String.New(group)), String.New(property), gd.NewVariant(value))
}

/*
Calls [method Object.notification] with the given [param notification] to all nodes inside this tree added to the [param group]. See also [url=$DOCS_URL/tutorials/best_practices/godot_notifications.html]Godot notifications[/url] and [method call_group] and [method set_group].
[b]Note:[/b] This method acts immediately on all selected nodes at once, which may cause stuttering in some performance-intensive situations.
*/
func (self Instance) NotifyGroup(group string, notification int) { //gd:SceneTree.notify_group
	class(self).NotifyGroup(String.Name(String.New(group)), gd.Int(notification))
}

/*
Sets the given [param property] to [param value] on all nodes inside this tree added to the given [param group]. Nodes that do not have the [param property] are ignored. See also [method call_group] and [method notify_group].
[b]Note:[/b] This method acts immediately on all selected nodes at once, which may cause stuttering in some performance-intensive situations.
[b]Note:[/b] In C#, [param property] must be in snake_case when referring to built-in Godot properties. Prefer using the names exposed in the [code]PropertyName[/code] class to avoid allocating a new [StringName] on each call.
*/
func (self Instance) SetGroup(group string, property string, value any) { //gd:SceneTree.set_group
	class(self).SetGroup(String.Name(String.New(group)), String.New(property), gd.NewVariant(value))
}

/*
Returns an [Array] containing all nodes inside this tree, that have been added to the given [param group], in scene hierarchy order.
*/
func (self Instance) GetNodesInGroup(group string) [][1]gdclass.Node { //gd:SceneTree.get_nodes_in_group
	return [][1]gdclass.Node(gd.ArrayAs[[][1]gdclass.Node](gd.InternalArray(class(self).GetNodesInGroup(String.Name(String.New(group))))))
}

/*
Returns the first [Node] found inside the tree, that has been added to the given [param group], in scene hierarchy order. Returns [code]null[/code] if no match is found. See also [method get_nodes_in_group].
*/
func (self Instance) GetFirstNodeInGroup(group string) [1]gdclass.Node { //gd:SceneTree.get_first_node_in_group
	return [1]gdclass.Node(class(self).GetFirstNodeInGroup(String.Name(String.New(group))))
}

/*
Returns the number of nodes assigned to the given group.
*/
func (self Instance) GetNodeCountInGroup(group string) int { //gd:SceneTree.get_node_count_in_group
	return int(int(class(self).GetNodeCountInGroup(String.Name(String.New(group)))))
}

/*
Changes the running scene to the one at the given [param path], after loading it into a [PackedScene] and creating a new instance.
Returns [constant OK] on success, [constant ERR_CANT_OPEN] if the [param path] cannot be loaded into a [PackedScene], or [constant ERR_CANT_CREATE] if that scene cannot be instantiated.
[b]Note:[/b] See [method change_scene_to_packed] for details on the order of operations.
*/
func (self Instance) ChangeSceneToFile(path string) error { //gd:SceneTree.change_scene_to_file
	return error(gd.ToError(class(self).ChangeSceneToFile(String.New(path))))
}

/*
Changes the running scene to a new instance of the given [PackedScene] (which must be valid).
Returns [constant OK] on success, [constant ERR_CANT_CREATE] if the scene cannot be instantiated, or [constant ERR_INVALID_PARAMETER] if the scene is invalid.
[b]Note:[/b] Operations happen in the following order when [method change_scene_to_packed] is called:
1. The current scene node is immediately removed from the tree. From that point, [method Node.get_tree] called on the current (outgoing) scene will return [code]null[/code]. [member current_scene] will be [code]null[/code], too, because the new scene is not available yet.
2. At the end of the frame, the formerly current scene, already removed from the tree, will be deleted (freed from memory) and then the new scene will be instantiated and added to the tree. [method Node.get_tree] and [member current_scene] will be back to working as usual.
This ensures that both scenes aren't running at the same time, while still freeing the previous scene in a safe way similar to [method Node.queue_free].
*/
func (self Instance) ChangeSceneToPacked(packed_scene [1]gdclass.PackedScene) error { //gd:SceneTree.change_scene_to_packed
	return error(gd.ToError(class(self).ChangeSceneToPacked(packed_scene)))
}

/*
Reloads the currently active scene, replacing [member current_scene] with a new instance of its original [PackedScene].
Returns [constant OK] on success, [constant ERR_UNCONFIGURED] if no [member current_scene] is defined, [constant ERR_CANT_OPEN] if [member current_scene] cannot be loaded into a [PackedScene], or [constant ERR_CANT_CREATE] if the scene cannot be instantiated.
*/
func (self Instance) ReloadCurrentScene() error { //gd:SceneTree.reload_current_scene
	return error(gd.ToError(class(self).ReloadCurrentScene()))
}

/*
If a current scene is loaded, calling this method will unload it.
*/
func (self Instance) UnloadCurrentScene() { //gd:SceneTree.unload_current_scene
	class(self).UnloadCurrentScene()
}

/*
Sets a custom [MultiplayerAPI] with the given [param root_path] (controlling also the relative subpaths), or override the default one if [param root_path] is empty.
[b]Note:[/b] No [MultiplayerAPI] must be configured for the subpath containing [param root_path], nested custom multiplayers are not allowed. I.e. if one is configured for [code]"/root/Foo"[/code] setting one for [code]"/root/Foo/Bar"[/code] will cause an error.
*/
func (self Instance) SetMultiplayer(multiplayer [1]gdclass.MultiplayerAPI) { //gd:SceneTree.set_multiplayer
	class(self).SetMultiplayer(multiplayer, Path.ToNode(String.New("")))
}

/*
Searches for the [MultiplayerAPI] configured for the given path, if one does not exist it searches the parent paths until one is found. If the path is empty, or none is found, the default one is returned. See [method set_multiplayer].
*/
func (self Instance) GetMultiplayer() [1]gdclass.MultiplayerAPI { //gd:SceneTree.get_multiplayer
	return [1]gdclass.MultiplayerAPI(class(self).GetMultiplayer(Path.ToNode(String.New(""))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SceneTree

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SceneTree"))
	casted := Instance{*(*gdclass.SceneTree)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) AutoAcceptQuit() bool {
	return bool(class(self).IsAutoAcceptQuit())
}

func (self Instance) SetAutoAcceptQuit(value bool) {
	class(self).SetAutoAcceptQuit(value)
}

func (self Instance) QuitOnGoBack() bool {
	return bool(class(self).IsQuitOnGoBack())
}

func (self Instance) SetQuitOnGoBack(value bool) {
	class(self).SetQuitOnGoBack(value)
}

func (self Instance) DebugCollisionsHint() bool {
	return bool(class(self).IsDebuggingCollisionsHint())
}

func (self Instance) SetDebugCollisionsHint(value bool) {
	class(self).SetDebugCollisionsHint(value)
}

func (self Instance) DebugPathsHint() bool {
	return bool(class(self).IsDebuggingPathsHint())
}

func (self Instance) SetDebugPathsHint(value bool) {
	class(self).SetDebugPathsHint(value)
}

func (self Instance) DebugNavigationHint() bool {
	return bool(class(self).IsDebuggingNavigationHint())
}

func (self Instance) SetDebugNavigationHint(value bool) {
	class(self).SetDebugNavigationHint(value)
}

func (self Instance) Paused() bool {
	return bool(class(self).IsPaused())
}

func (self Instance) SetPaused(value bool) {
	class(self).SetPause(value)
}

func (self Instance) EditedSceneRoot() [1]gdclass.Node {
	return [1]gdclass.Node(class(self).GetEditedSceneRoot())
}

func (self Instance) SetEditedSceneRoot(value [1]gdclass.Node) {
	class(self).SetEditedSceneRoot(value)
}

func (self Instance) CurrentScene() [1]gdclass.Node {
	return [1]gdclass.Node(class(self).GetCurrentScene())
}

func (self Instance) SetCurrentScene(value [1]gdclass.Node) {
	class(self).SetCurrentScene(value)
}

func (self Instance) Root() [1]gdclass.Window {
	return [1]gdclass.Window(class(self).GetRoot())
}

func (self Instance) MultiplayerPoll() bool {
	return bool(class(self).IsMultiplayerPollEnabled())
}

func (self Instance) SetMultiplayerPoll(value bool) {
	class(self).SetMultiplayerPollEnabled(value)
}

func (self Instance) PhysicsInterpolation() bool {
	return bool(class(self).IsPhysicsInterpolationEnabled())
}

func (self Instance) SetPhysicsInterpolation(value bool) {
	class(self).SetPhysicsInterpolationEnabled(value)
}

//go:nosplit
func (self class) GetRoot() [1]gdclass.Window { //gd:SceneTree.get_root
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_get_root, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Window{gd.PointerMustAssertInstanceID[gdclass.Window](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if a node added to the given group [param name] exists in the tree.
*/
//go:nosplit
func (self class) HasGroup(name String.Name) bool { //gd:SceneTree.has_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_has_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsAutoAcceptQuit() bool { //gd:SceneTree.is_auto_accept_quit
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_is_auto_accept_quit, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoAcceptQuit(enabled bool) { //gd:SceneTree.set_auto_accept_quit
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_auto_accept_quit, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsQuitOnGoBack() bool { //gd:SceneTree.is_quit_on_go_back
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_is_quit_on_go_back, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetQuitOnGoBack(enabled bool) { //gd:SceneTree.set_quit_on_go_back
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_quit_on_go_back, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetDebugCollisionsHint(enable bool) { //gd:SceneTree.set_debug_collisions_hint
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_debug_collisions_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDebuggingCollisionsHint() bool { //gd:SceneTree.is_debugging_collisions_hint
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_is_debugging_collisions_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDebugPathsHint(enable bool) { //gd:SceneTree.set_debug_paths_hint
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_debug_paths_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDebuggingPathsHint() bool { //gd:SceneTree.is_debugging_paths_hint
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_is_debugging_paths_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDebugNavigationHint(enable bool) { //gd:SceneTree.set_debug_navigation_hint
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_debug_navigation_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDebuggingNavigationHint() bool { //gd:SceneTree.is_debugging_navigation_hint
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_is_debugging_navigation_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEditedSceneRoot(scene [1]gdclass.Node) { //gd:SceneTree.set_edited_scene_root
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(scene[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_edited_scene_root, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEditedSceneRoot() [1]gdclass.Node { //gd:SceneTree.get_edited_scene_root
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_get_edited_scene_root, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPause(enable bool) { //gd:SceneTree.set_pause
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_pause, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPaused() bool { //gd:SceneTree.is_paused
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_is_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a new [SceneTreeTimer]. After [param time_sec] in seconds have passed, the timer will emit [signal SceneTreeTimer.timeout] and will be automatically freed.
If [param process_always] is [code]false[/code], the timer will be paused when setting [member SceneTree.paused] to [code]true[/code].
If [param process_in_physics] is [code]true[/code], the timer will update at the end of the physics frame, instead of the process frame.
If [param ignore_time_scale] is [code]true[/code], the timer will ignore [member Engine.time_scale] and update with the real, elapsed time.
This method is commonly used to create a one-shot delay timer, as in the following example:
[codeblocks]
[gdscript]
func some_function():
    print("start")
    await get_tree().create_timer(1.0).timeout
    print("end")
[/gdscript]
[csharp]
public async Task SomeFunction()
{
    GD.Print("start");
    await ToSignal(GetTree().CreateTimer(1.0f), SceneTreeTimer.SignalName.Timeout);
    GD.Print("end");
}
[/csharp]
[/codeblocks]
[b]Note:[/b] The timer is always updated [i]after[/i] all of the nodes in the tree. A node's [method Node._process] method would be called before the timer updates (or [method Node._physics_process] if [param process_in_physics] is set to [code]true[/code]).
*/
//go:nosplit
func (self class) CreateTimer(time_sec gd.Float, process_always bool, process_in_physics bool, ignore_time_scale bool) [1]gdclass.SceneTreeTimer { //gd:SceneTree.create_timer
	var frame = callframe.New()
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, process_always)
	callframe.Arg(frame, process_in_physics)
	callframe.Arg(frame, ignore_time_scale)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_create_timer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SceneTreeTimer{gd.PointerWithOwnershipTransferredToGo[gdclass.SceneTreeTimer](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates and returns a new [Tween] processed in this tree. The Tween will start automatically on the next process frame or physics frame (depending on its [enum Tween.TweenProcessMode]).
[b]Note:[/b] A [Tween] created using this method is not bound to any [Node]. It may keep working until there is nothing left to animate. If you want the [Tween] to be automatically killed when the [Node] is freed, use [method Node.create_tween] or [method Tween.bind_node].
*/
//go:nosplit
func (self class) CreateTween() [1]gdclass.Tween { //gd:SceneTree.create_tween
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_create_tween, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns an [Array] of currently existing [Tween]s in the tree, including paused tweens.
*/
//go:nosplit
func (self class) GetProcessedTweens() Array.Contains[[1]gdclass.Tween] { //gd:SceneTree.get_processed_tweens
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_get_processed_tweens, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Tween]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the number of nodes inside this tree.
*/
//go:nosplit
func (self class) GetNodeCount() gd.Int { //gd:SceneTree.get_node_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_get_node_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns how many frames have been processed, since the application started. This is [i]not[/i] a measurement of elapsed time.
*/
//go:nosplit
func (self class) GetFrame() gd.Int { //gd:SceneTree.get_frame
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_get_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Quits the application at the end of the current iteration, with the given [param exit_code].
By convention, an exit code of [code]0[/code] indicates success, whereas any other exit code indicates an error. For portability reasons, it should be between [code]0[/code] and [code]125[/code] (inclusive).
[b]Note:[/b] On iOS this method doesn't work. Instead, as recommended by the [url=https://developer.apple.com/library/archive/qa/qa1561/_index.html]iOS Human Interface Guidelines[/url], the user is expected to close apps via the Home button.
*/
//go:nosplit
func (self class) Quit(exit_code gd.Int) { //gd:SceneTree.quit
	var frame = callframe.New()
	callframe.Arg(frame, exit_code)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_quit, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetPhysicsInterpolationEnabled(enabled bool) { //gd:SceneTree.set_physics_interpolation_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_physics_interpolation_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPhysicsInterpolationEnabled() bool { //gd:SceneTree.is_physics_interpolation_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_is_physics_interpolation_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Queues the given [param obj] to be deleted, calling its [method Object.free] at the end of the current frame. This method is similar to [method Node.queue_free].
*/
//go:nosplit
func (self class) QueueDelete(obj [1]gd.Object) { //gd:SceneTree.queue_delete
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_queue_delete, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Calls [method Object.notification] with the given [param notification] to all nodes inside this tree added to the [param group]. Use [param call_flags] to customize this method's behavior (see [enum GroupCallFlags]).
*/
//go:nosplit
func (self class) NotifyGroupFlags(call_flags gd.Int, group String.Name, notification gd.Int) { //gd:SceneTree.notify_group_flags
	var frame = callframe.New()
	callframe.Arg(frame, call_flags)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(group)))
	callframe.Arg(frame, notification)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_notify_group_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the given [param property] to [param value] on all nodes inside this tree added to the given [param group]. Nodes that do not have the [param property] are ignored. Use [param call_flags] to customize this method's behavior (see [enum GroupCallFlags]).
[b]Note:[/b] In C#, [param property] must be in snake_case when referring to built-in Godot properties. Prefer using the names exposed in the [code]PropertyName[/code] class to avoid allocating a new [StringName] on each call.
*/
//go:nosplit
func (self class) SetGroupFlags(call_flags gd.Int, group String.Name, property String.Readable, value gd.Variant) { //gd:SceneTree.set_group_flags
	var frame = callframe.New()
	callframe.Arg(frame, call_flags)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(group)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(property)))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_group_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Calls [method Object.notification] with the given [param notification] to all nodes inside this tree added to the [param group]. See also [url=$DOCS_URL/tutorials/best_practices/godot_notifications.html]Godot notifications[/url] and [method call_group] and [method set_group].
[b]Note:[/b] This method acts immediately on all selected nodes at once, which may cause stuttering in some performance-intensive situations.
*/
//go:nosplit
func (self class) NotifyGroup(group String.Name, notification gd.Int) { //gd:SceneTree.notify_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(group)))
	callframe.Arg(frame, notification)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_notify_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the given [param property] to [param value] on all nodes inside this tree added to the given [param group]. Nodes that do not have the [param property] are ignored. See also [method call_group] and [method notify_group].
[b]Note:[/b] This method acts immediately on all selected nodes at once, which may cause stuttering in some performance-intensive situations.
[b]Note:[/b] In C#, [param property] must be in snake_case when referring to built-in Godot properties. Prefer using the names exposed in the [code]PropertyName[/code] class to avoid allocating a new [StringName] on each call.
*/
//go:nosplit
func (self class) SetGroup(group String.Name, property String.Readable, value gd.Variant) { //gd:SceneTree.set_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(group)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(property)))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an [Array] containing all nodes inside this tree, that have been added to the given [param group], in scene hierarchy order.
*/
//go:nosplit
func (self class) GetNodesInGroup(group String.Name) Array.Contains[[1]gdclass.Node] { //gd:SceneTree.get_nodes_in_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(group)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_get_nodes_in_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Node]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the first [Node] found inside the tree, that has been added to the given [param group], in scene hierarchy order. Returns [code]null[/code] if no match is found. See also [method get_nodes_in_group].
*/
//go:nosplit
func (self class) GetFirstNodeInGroup(group String.Name) [1]gdclass.Node { //gd:SceneTree.get_first_node_in_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(group)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_get_first_node_in_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the number of nodes assigned to the given group.
*/
//go:nosplit
func (self class) GetNodeCountInGroup(group String.Name) gd.Int { //gd:SceneTree.get_node_count_in_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(group)))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_get_node_count_in_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurrentScene(child_node [1]gdclass.Node) { //gd:SceneTree.set_current_scene
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(child_node[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_current_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurrentScene() [1]gdclass.Node { //gd:SceneTree.get_current_scene
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_get_current_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Changes the running scene to the one at the given [param path], after loading it into a [PackedScene] and creating a new instance.
Returns [constant OK] on success, [constant ERR_CANT_OPEN] if the [param path] cannot be loaded into a [PackedScene], or [constant ERR_CANT_CREATE] if that scene cannot be instantiated.
[b]Note:[/b] See [method change_scene_to_packed] for details on the order of operations.
*/
//go:nosplit
func (self class) ChangeSceneToFile(path String.Readable) gd.Error { //gd:SceneTree.change_scene_to_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_change_scene_to_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Changes the running scene to a new instance of the given [PackedScene] (which must be valid).
Returns [constant OK] on success, [constant ERR_CANT_CREATE] if the scene cannot be instantiated, or [constant ERR_INVALID_PARAMETER] if the scene is invalid.
[b]Note:[/b] Operations happen in the following order when [method change_scene_to_packed] is called:
1. The current scene node is immediately removed from the tree. From that point, [method Node.get_tree] called on the current (outgoing) scene will return [code]null[/code]. [member current_scene] will be [code]null[/code], too, because the new scene is not available yet.
2. At the end of the frame, the formerly current scene, already removed from the tree, will be deleted (freed from memory) and then the new scene will be instantiated and added to the tree. [method Node.get_tree] and [member current_scene] will be back to working as usual.
This ensures that both scenes aren't running at the same time, while still freeing the previous scene in a safe way similar to [method Node.queue_free].
*/
//go:nosplit
func (self class) ChangeSceneToPacked(packed_scene [1]gdclass.PackedScene) gd.Error { //gd:SceneTree.change_scene_to_packed
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(packed_scene[0])[0])
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_change_scene_to_packed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Reloads the currently active scene, replacing [member current_scene] with a new instance of its original [PackedScene].
Returns [constant OK] on success, [constant ERR_UNCONFIGURED] if no [member current_scene] is defined, [constant ERR_CANT_OPEN] if [member current_scene] cannot be loaded into a [PackedScene], or [constant ERR_CANT_CREATE] if the scene cannot be instantiated.
*/
//go:nosplit
func (self class) ReloadCurrentScene() gd.Error { //gd:SceneTree.reload_current_scene
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_reload_current_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If a current scene is loaded, calling this method will unload it.
*/
//go:nosplit
func (self class) UnloadCurrentScene() { //gd:SceneTree.unload_current_scene
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_unload_current_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets a custom [MultiplayerAPI] with the given [param root_path] (controlling also the relative subpaths), or override the default one if [param root_path] is empty.
[b]Note:[/b] No [MultiplayerAPI] must be configured for the subpath containing [param root_path], nested custom multiplayers are not allowed. I.e. if one is configured for [code]"/root/Foo"[/code] setting one for [code]"/root/Foo/Bar"[/code] will cause an error.
*/
//go:nosplit
func (self class) SetMultiplayer(multiplayer [1]gdclass.MultiplayerAPI, root_path Path.ToNode) { //gd:SceneTree.set_multiplayer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(multiplayer[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(root_path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_multiplayer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Searches for the [MultiplayerAPI] configured for the given path, if one does not exist it searches the parent paths until one is found. If the path is empty, or none is found, the default one is returned. See [method set_multiplayer].
*/
//go:nosplit
func (self class) GetMultiplayer(for_path Path.ToNode) [1]gdclass.MultiplayerAPI { //gd:SceneTree.get_multiplayer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(for_path)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_get_multiplayer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.MultiplayerAPI{gd.PointerWithOwnershipTransferredToGo[gdclass.MultiplayerAPI](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMultiplayerPollEnabled(enabled bool) { //gd:SceneTree.set_multiplayer_poll_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_set_multiplayer_poll_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMultiplayerPollEnabled() bool { //gd:SceneTree.is_multiplayer_poll_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneTree.Bind_is_multiplayer_poll_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnTreeChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tree_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTreeProcessModeChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tree_process_mode_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNodeAdded(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_added"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNodeRemoved(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_removed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNodeRenamed(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_renamed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNodeConfigurationWarningChanged(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_configuration_warning_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnProcessFrame(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("process_frame"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPhysicsFrame(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("physics_frame"), gd.NewCallable(cb), 0)
}

func (self class) AsSceneTree() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSceneTree() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMainLoop() MainLoop.Advanced {
	return *((*MainLoop.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMainLoop() MainLoop.Instance {
	return *((*MainLoop.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(MainLoop.Advanced(self.AsMainLoop()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(MainLoop.Instance(self.AsMainLoop()), name)
	}
}
func init() {
	gdclass.Register("SceneTree", func(ptr gd.Object) any { return [1]gdclass.SceneTree{*(*gdclass.SceneTree)(unsafe.Pointer(&ptr))} })
}

type GroupCallFlags = gdclass.SceneTreeGroupCallFlags //gd:SceneTree.GroupCallFlags

const (
	/*Call nodes within a group with no special behavior (default).*/
	GroupCallDefault GroupCallFlags = 0
	/*Call nodes within a group in reverse tree hierarchy order (all nested children are called before their respective parent nodes).*/
	GroupCallReverse GroupCallFlags = 1
	/*Call nodes within a group at the end of the current frame (can be either process or physics frame), similar to [method Object.call_deferred].*/
	GroupCallDeferred GroupCallFlags = 2
	/*Call nodes within a group only once, even if the call is executed many times in the same frame. Must be combined with [constant GROUP_CALL_DEFERRED] to work.
	  [b]Note:[/b] Different arguments are not taken into account. Therefore, when the same call is executed with different arguments, only the first call will be performed.*/
	GroupCallUnique GroupCallFlags = 4
)

type Error = gd.Error //gd:Error

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
