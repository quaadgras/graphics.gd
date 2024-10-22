package SceneTree

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/MainLoop"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
As one of the most important classes, the [SceneTree] manages the hierarchy of nodes in a scene, as well as scenes themselves. Nodes can be added, fetched and removed. The whole scene tree (and thus the current scene) can be paused. Scenes can be loaded, switched and reloaded.
You can also use the [SceneTree] to organize your nodes into [b]groups[/b]: every node can be added to as many groups as you want to create, e.g. an "enemy" group. You can then iterate these groups or even call methods and set properties on all the nodes belonging to any given group.
[SceneTree] is the default [MainLoop] implementation used by the engine, and is thus in charge of the game loop.

*/
type Go [1]classdb.SceneTree

/*
Returns [code]true[/code] if a node added to the given group [param name] exists in the tree.
*/
func (self Go) HasGroup(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasGroup(gc.StringName(name)))
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
func (self Go) CreateTimer(time_sec float64) gdclass.SceneTreeTimer {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.SceneTreeTimer(class(self).CreateTimer(gc, gd.Float(time_sec), true, false, false))
}

/*
Creates and returns a new [Tween] processed in this tree. The Tween will start automatically on the next process frame or physics frame (depending on its [enum Tween.TweenProcessMode]).
[b]Note:[/b] A [Tween] created using this method is not bound to any [Node]. It may keep working until there is nothing left to animate. If you want the [Tween] to be automatically killed when the [Node] is freed, use [method Node.create_tween] or [method Tween.bind_node].
*/
func (self Go) CreateTween() gdclass.Tween {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Tween(class(self).CreateTween(gc))
}

/*
Returns an [Array] of currently existing [Tween]s in the tree, including paused tweens.
*/
func (self Go) GetProcessedTweens() gd.ArrayOf[gdclass.Tween] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Tween](class(self).GetProcessedTweens(gc))
}

/*
Returns the number of nodes inside this tree.
*/
func (self Go) GetNodeCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetNodeCount()))
}

/*
Returns how many frames have been processed, since the application started. This is [i]not[/i] a measurement of elapsed time.
*/
func (self Go) GetFrame() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFrame()))
}

/*
Quits the application at the end of the current iteration, with the given [param exit_code].
By convention, an exit code of [code]0[/code] indicates success, whereas any other exit code indicates an error. For portability reasons, it should be between [code]0[/code] and [code]125[/code] (inclusive).
[b]Note:[/b] On iOS this method doesn't work. Instead, as recommended by the [url=https://developer.apple.com/library/archive/qa/qa1561/_index.html]iOS Human Interface Guidelines[/url], the user is expected to close apps via the Home button.
*/
func (self Go) Quit() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Quit(gd.Int(0))
}

/*
Queues the given [param obj] to be deleted, calling its [method Object.free] at the end of the current frame. This method is similar to [method Node.queue_free].
*/
func (self Go) QueueDelete(obj gd.Object) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).QueueDelete(obj)
}

/*
Calls [method Object.notification] with the given [param notification] to all nodes inside this tree added to the [param group]. Use [param call_flags] to customize this method's behavior (see [enum GroupCallFlags]).
*/
func (self Go) NotifyGroupFlags(call_flags int, group string, notification int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).NotifyGroupFlags(gd.Int(call_flags), gc.StringName(group), gd.Int(notification))
}

/*
Sets the given [param property] to [param value] on all nodes inside this tree added to the given [param group]. Nodes that do not have the [param property] are ignored. Use [param call_flags] to customize this method's behavior (see [enum GroupCallFlags]).
[b]Note:[/b] In C#, [param property] must be in snake_case when referring to built-in Godot properties. Prefer using the names exposed in the [code]PropertyName[/code] class to avoid allocating a new [StringName] on each call.
*/
func (self Go) SetGroupFlags(call_flags int, group string, property string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGroupFlags(gd.Int(call_flags), gc.StringName(group), gc.String(property), value)
}

/*
Calls [method Object.notification] with the given [param notification] to all nodes inside this tree added to the [param group]. See also [url=$DOCS_URL/tutorials/best_practices/godot_notifications.html]Godot notifications[/url] and [method call_group] and [method set_group].
[b]Note:[/b] This method acts immediately on all selected nodes at once, which may cause stuttering in some performance-intensive situations.
*/
func (self Go) NotifyGroup(group string, notification int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).NotifyGroup(gc.StringName(group), gd.Int(notification))
}

/*
Sets the given [param property] to [param value] on all nodes inside this tree added to the given [param group]. Nodes that do not have the [param property] are ignored. See also [method call_group] and [method notify_group].
[b]Note:[/b] This method acts immediately on all selected nodes at once, which may cause stuttering in some performance-intensive situations.
[b]Note:[/b] In C#, [param property] must be in snake_case when referring to built-in Godot properties. Prefer using the names exposed in the [code]PropertyName[/code] class to avoid allocating a new [StringName] on each call.
*/
func (self Go) SetGroup(group string, property string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGroup(gc.StringName(group), gc.String(property), value)
}

/*
Returns an [Array] containing all nodes inside this tree, that have been added to the given [param group], in scene hierarchy order.
*/
func (self Go) GetNodesInGroup(group string) gd.ArrayOf[gdclass.Node] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Node](class(self).GetNodesInGroup(gc, gc.StringName(group)))
}

/*
Returns the first [Node] found inside the tree, that has been added to the given [param group], in scene hierarchy order. Returns [code]null[/code] if no match is found. See also [method get_nodes_in_group].
*/
func (self Go) GetFirstNodeInGroup(group string) gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Node(class(self).GetFirstNodeInGroup(gc, gc.StringName(group)))
}

/*
Returns the number of nodes assigned to the given group.
*/
func (self Go) GetNodeCountInGroup(group string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetNodeCountInGroup(gc.StringName(group))))
}

/*
Changes the running scene to the one at the given [param path], after loading it into a [PackedScene] and creating a new instance.
Returns [constant OK] on success, [constant ERR_CANT_OPEN] if the [param path] cannot be loaded into a [PackedScene], or [constant ERR_CANT_CREATE] if that scene cannot be instantiated.
[b]Note:[/b] See [method change_scene_to_packed] for details on the order of operations.
*/
func (self Go) ChangeSceneToFile(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).ChangeSceneToFile(gc.String(path)))
}

/*
Changes the running scene to a new instance of the given [PackedScene] (which must be valid).
Returns [constant OK] on success, [constant ERR_CANT_CREATE] if the scene cannot be instantiated, or [constant ERR_INVALID_PARAMETER] if the scene is invalid.
[b]Note:[/b] Operations happen in the following order when [method change_scene_to_packed] is called:
1. The current scene node is immediately removed from the tree. From that point, [method Node.get_tree] called on the current (outgoing) scene will return [code]null[/code]. [member current_scene] will be [code]null[/code], too, because the new scene is not available yet.
2. At the end of the frame, the formerly current scene, already removed from the tree, will be deleted (freed from memory) and then the new scene will be instantiated and added to the tree. [method Node.get_tree] and [member current_scene] will be back to working as usual.
This ensures that both scenes aren't running at the same time, while still freeing the previous scene in a safe way similar to [method Node.queue_free].
*/
func (self Go) ChangeSceneToPacked(packed_scene gdclass.PackedScene) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).ChangeSceneToPacked(packed_scene))
}

/*
Reloads the currently active scene, replacing [member current_scene] with a new instance of its original [PackedScene].
Returns [constant OK] on success, [constant ERR_UNCONFIGURED] if no [member current_scene] is defined, [constant ERR_CANT_OPEN] if [member current_scene] cannot be loaded into a [PackedScene], or [constant ERR_CANT_CREATE] if the scene cannot be instantiated.
*/
func (self Go) ReloadCurrentScene() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).ReloadCurrentScene())
}

/*
If a current scene is loaded, calling this method will unload it.
*/
func (self Go) UnloadCurrentScene() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UnloadCurrentScene()
}

/*
Sets a custom [MultiplayerAPI] with the given [param root_path] (controlling also the relative subpaths), or override the default one if [param root_path] is empty.
[b]Note:[/b] No [MultiplayerAPI] must be configured for the subpath containing [param root_path], nested custom multiplayers are not allowed. I.e. if one is configured for [code]"/root/Foo"[/code] setting one for [code]"/root/Foo/Bar"[/code] will cause an error.
*/
func (self Go) SetMultiplayer(multiplayer gdclass.MultiplayerAPI) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMultiplayer(multiplayer, gc.String("").NodePath(gc))
}

/*
Searches for the [MultiplayerAPI] configured for the given path, if one does not exist it searches the parent paths until one is found. If the path is empty, or none is found, the default one is returned. See [method set_multiplayer].
*/
func (self Go) GetMultiplayer() gdclass.MultiplayerAPI {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.MultiplayerAPI(class(self).GetMultiplayer(gc, gc.String("").NodePath(gc)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SceneTree
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SceneTree"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) AutoAcceptQuit() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAutoAcceptQuit())
}

func (self Go) SetAutoAcceptQuit(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoAcceptQuit(value)
}

func (self Go) QuitOnGoBack() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsQuitOnGoBack())
}

func (self Go) SetQuitOnGoBack(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetQuitOnGoBack(value)
}

func (self Go) DebugCollisionsHint() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDebuggingCollisionsHint())
}

func (self Go) SetDebugCollisionsHint(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDebugCollisionsHint(value)
}

func (self Go) DebugPathsHint() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDebuggingPathsHint())
}

func (self Go) SetDebugPathsHint(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDebugPathsHint(value)
}

func (self Go) DebugNavigationHint() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDebuggingNavigationHint())
}

func (self Go) SetDebugNavigationHint(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDebugNavigationHint(value)
}

func (self Go) Paused() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPaused())
}

func (self Go) SetPaused(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPause(value)
}

func (self Go) EditedSceneRoot() gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Node(class(self).GetEditedSceneRoot(gc))
}

func (self Go) SetEditedSceneRoot(value gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditedSceneRoot(value)
}

func (self Go) CurrentScene() gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Node(class(self).GetCurrentScene(gc))
}

func (self Go) SetCurrentScene(value gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCurrentScene(value)
}

func (self Go) Root() gdclass.Window {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Window(class(self).GetRoot(gc))
}

func (self Go) MultiplayerPoll() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMultiplayerPollEnabled())
}

func (self Go) SetMultiplayerPoll(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMultiplayerPollEnabled(value)
}

func (self Go) PhysicsInterpolation() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPhysicsInterpolationEnabled())
}

func (self Go) SetPhysicsInterpolation(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPhysicsInterpolationEnabled(value)
}

//go:nosplit
func (self class) GetRoot(ctx gd.Lifetime) gdclass.Window {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_get_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Window
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a node added to the given group [param name] exists in the tree.
*/
//go:nosplit
func (self class) HasGroup(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_has_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) IsAutoAcceptQuit() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_is_auto_accept_quit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoAcceptQuit(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_auto_accept_quit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsQuitOnGoBack() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_is_quit_on_go_back, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetQuitOnGoBack(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_quit_on_go_back, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDebugCollisionsHint(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_debug_collisions_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDebuggingCollisionsHint() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_is_debugging_collisions_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugPathsHint(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_debug_paths_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDebuggingPathsHint() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_is_debugging_paths_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugNavigationHint(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_debug_navigation_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDebuggingNavigationHint() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_is_debugging_navigation_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEditedSceneRoot(scene gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(scene[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_edited_scene_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEditedSceneRoot(ctx gd.Lifetime) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_get_edited_scene_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPause(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_pause, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPaused() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_is_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) CreateTimer(ctx gd.Lifetime, time_sec gd.Float, process_always bool, process_in_physics bool, ignore_time_scale bool) gdclass.SceneTreeTimer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, time_sec)
	callframe.Arg(frame, process_always)
	callframe.Arg(frame, process_in_physics)
	callframe.Arg(frame, ignore_time_scale)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_create_timer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.SceneTreeTimer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates and returns a new [Tween] processed in this tree. The Tween will start automatically on the next process frame or physics frame (depending on its [enum Tween.TweenProcessMode]).
[b]Note:[/b] A [Tween] created using this method is not bound to any [Node]. It may keep working until there is nothing left to animate. If you want the [Tween] to be automatically killed when the [Node] is freed, use [method Node.create_tween] or [method Tween.bind_node].
*/
//go:nosplit
func (self class) CreateTween(ctx gd.Lifetime) gdclass.Tween {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_create_tween, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Tween
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an [Array] of currently existing [Tween]s in the tree, including paused tweens.
*/
//go:nosplit
func (self class) GetProcessedTweens(ctx gd.Lifetime) gd.ArrayOf[gdclass.Tween] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_get_processed_tweens, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Tween](ret)
}
/*
Returns the number of nodes inside this tree.
*/
//go:nosplit
func (self class) GetNodeCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_get_node_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns how many frames have been processed, since the application started. This is [i]not[/i] a measurement of elapsed time.
*/
//go:nosplit
func (self class) GetFrame() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_get_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) Quit(exit_code gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exit_code)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_quit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetPhysicsInterpolationEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_physics_interpolation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPhysicsInterpolationEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_is_physics_interpolation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Queues the given [param obj] to be deleted, calling its [method Object.free] at the end of the current frame. This method is similar to [method Node.queue_free].
*/
//go:nosplit
func (self class) QueueDelete(obj gd.Object)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_queue_delete, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Calls [method Object.notification] with the given [param notification] to all nodes inside this tree added to the [param group]. Use [param call_flags] to customize this method's behavior (see [enum GroupCallFlags]).
*/
//go:nosplit
func (self class) NotifyGroupFlags(call_flags gd.Int, group gd.StringName, notification gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, call_flags)
	callframe.Arg(frame, mmm.Get(group))
	callframe.Arg(frame, notification)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_notify_group_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given [param property] to [param value] on all nodes inside this tree added to the given [param group]. Nodes that do not have the [param property] are ignored. Use [param call_flags] to customize this method's behavior (see [enum GroupCallFlags]).
[b]Note:[/b] In C#, [param property] must be in snake_case when referring to built-in Godot properties. Prefer using the names exposed in the [code]PropertyName[/code] class to avoid allocating a new [StringName] on each call.
*/
//go:nosplit
func (self class) SetGroupFlags(call_flags gd.Int, group gd.StringName, property gd.String, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, call_flags)
	callframe.Arg(frame, mmm.Get(group))
	callframe.Arg(frame, mmm.Get(property))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_group_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Calls [method Object.notification] with the given [param notification] to all nodes inside this tree added to the [param group]. See also [url=$DOCS_URL/tutorials/best_practices/godot_notifications.html]Godot notifications[/url] and [method call_group] and [method set_group].
[b]Note:[/b] This method acts immediately on all selected nodes at once, which may cause stuttering in some performance-intensive situations.
*/
//go:nosplit
func (self class) NotifyGroup(group gd.StringName, notification gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(group))
	callframe.Arg(frame, notification)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_notify_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given [param property] to [param value] on all nodes inside this tree added to the given [param group]. Nodes that do not have the [param property] are ignored. See also [method call_group] and [method notify_group].
[b]Note:[/b] This method acts immediately on all selected nodes at once, which may cause stuttering in some performance-intensive situations.
[b]Note:[/b] In C#, [param property] must be in snake_case when referring to built-in Godot properties. Prefer using the names exposed in the [code]PropertyName[/code] class to avoid allocating a new [StringName] on each call.
*/
//go:nosplit
func (self class) SetGroup(group gd.StringName, property gd.String, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(group))
	callframe.Arg(frame, mmm.Get(property))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an [Array] containing all nodes inside this tree, that have been added to the given [param group], in scene hierarchy order.
*/
//go:nosplit
func (self class) GetNodesInGroup(ctx gd.Lifetime, group gd.StringName) gd.ArrayOf[gdclass.Node] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(group))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_get_nodes_in_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Node](ret)
}
/*
Returns the first [Node] found inside the tree, that has been added to the given [param group], in scene hierarchy order. Returns [code]null[/code] if no match is found. See also [method get_nodes_in_group].
*/
//go:nosplit
func (self class) GetFirstNodeInGroup(ctx gd.Lifetime, group gd.StringName) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(group))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_get_first_node_in_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the number of nodes assigned to the given group.
*/
//go:nosplit
func (self class) GetNodeCountInGroup(group gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(group))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_get_node_count_in_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCurrentScene(child_node gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(child_node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_current_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurrentScene(ctx gd.Lifetime) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_get_current_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Changes the running scene to the one at the given [param path], after loading it into a [PackedScene] and creating a new instance.
Returns [constant OK] on success, [constant ERR_CANT_OPEN] if the [param path] cannot be loaded into a [PackedScene], or [constant ERR_CANT_CREATE] if that scene cannot be instantiated.
[b]Note:[/b] See [method change_scene_to_packed] for details on the order of operations.
*/
//go:nosplit
func (self class) ChangeSceneToFile(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_change_scene_to_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) ChangeSceneToPacked(packed_scene gdclass.PackedScene) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(packed_scene[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_change_scene_to_packed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Reloads the currently active scene, replacing [member current_scene] with a new instance of its original [PackedScene].
Returns [constant OK] on success, [constant ERR_UNCONFIGURED] if no [member current_scene] is defined, [constant ERR_CANT_OPEN] if [member current_scene] cannot be loaded into a [PackedScene], or [constant ERR_CANT_CREATE] if the scene cannot be instantiated.
*/
//go:nosplit
func (self class) ReloadCurrentScene() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_reload_current_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If a current scene is loaded, calling this method will unload it.
*/
//go:nosplit
func (self class) UnloadCurrentScene()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_unload_current_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a custom [MultiplayerAPI] with the given [param root_path] (controlling also the relative subpaths), or override the default one if [param root_path] is empty.
[b]Note:[/b] No [MultiplayerAPI] must be configured for the subpath containing [param root_path], nested custom multiplayers are not allowed. I.e. if one is configured for [code]"/root/Foo"[/code] setting one for [code]"/root/Foo/Bar"[/code] will cause an error.
*/
//go:nosplit
func (self class) SetMultiplayer(multiplayer gdclass.MultiplayerAPI, root_path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(multiplayer[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(root_path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_multiplayer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Searches for the [MultiplayerAPI] configured for the given path, if one does not exist it searches the parent paths until one is found. If the path is empty, or none is found, the default one is returned. See [method set_multiplayer].
*/
//go:nosplit
func (self class) GetMultiplayer(ctx gd.Lifetime, for_path gd.NodePath) gdclass.MultiplayerAPI {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(for_path))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_get_multiplayer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.MultiplayerAPI
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMultiplayerPollEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_set_multiplayer_poll_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMultiplayerPollEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneTree.Bind_is_multiplayer_poll_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnTreeChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tree_changed"), gc.Callable(cb), 0)
}


func (self Go) OnTreeProcessModeChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("tree_process_mode_changed"), gc.Callable(cb), 0)
}


func (self Go) OnNodeAdded(cb func(node gdclass.Node)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("node_added"), gc.Callable(cb), 0)
}


func (self Go) OnNodeRemoved(cb func(node gdclass.Node)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("node_removed"), gc.Callable(cb), 0)
}


func (self Go) OnNodeRenamed(cb func(node gdclass.Node)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("node_renamed"), gc.Callable(cb), 0)
}


func (self Go) OnNodeConfigurationWarningChanged(cb func(node gdclass.Node)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("node_configuration_warning_changed"), gc.Callable(cb), 0)
}


func (self Go) OnProcessFrame(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("process_frame"), gc.Callable(cb), 0)
}


func (self Go) OnPhysicsFrame(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("physics_frame"), gc.Callable(cb), 0)
}


func (self class) AsSceneTree() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSceneTree() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMainLoop() MainLoop.GD { return *((*MainLoop.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMainLoop() MainLoop.Go { return *((*MainLoop.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMainLoop(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMainLoop(), name)
	}
}
func init() {classdb.Register("SceneTree", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type GroupCallFlags = classdb.SceneTreeGroupCallFlags

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
