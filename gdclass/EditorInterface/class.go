package EditorInterface

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[EditorInterface] gives you control over Godot editor's window. It allows customizing the window, saving and (re-)loading scenes, rendering mesh previews, inspecting and editing resources and objects, and provides access to [EditorSettings], [EditorFileSystem], [EditorResourcePreview], [ScriptEditor], the editor viewport, and information about scenes.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton directly by its name.
[codeblocks]
[gdscript]
var editor_settings = EditorInterface.get_editor_settings()
[/gdscript]
[csharp]
// In C# you can access it via the static Singleton property.
EditorSettings settings = EditorInterface.Singleton.GetEditorSettings();
[/csharp]
[/codeblocks]

*/
var self gdclass.EditorInterface
var once sync.Once
func singleton() {
	gc := gd.Static
	obj := gc.API.Object.GetSingleton(gc, gc.API.Singletons.EditorInterface)
	self = *(*gdclass.EditorInterface)(unsafe.Pointer(&obj))
}

/*
Restarts the editor. This closes the editor and then opens the same project. If [param save] is [code]true[/code], the project will be saved before restarting.
*/
func RestartEditor() {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).RestartEditor(true)
}

/*
Returns the editor's [EditorCommandPalette] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func GetCommandPalette() gdclass.EditorCommandPalette {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.EditorCommandPalette(class(self).GetCommandPalette(gc))
}

/*
Returns the editor's [EditorFileSystem] instance.
*/
func GetResourceFilesystem() gdclass.EditorFileSystem {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.EditorFileSystem(class(self).GetResourceFilesystem(gc))
}

/*
Returns the [EditorPaths] singleton.
*/
func GetEditorPaths() gdclass.EditorPaths {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.EditorPaths(class(self).GetEditorPaths(gc))
}

/*
Returns the editor's [EditorResourcePreview] instance.
*/
func GetResourcePreviewer() gdclass.EditorResourcePreview {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.EditorResourcePreview(class(self).GetResourcePreviewer(gc))
}

/*
Returns the editor's [EditorSelection] instance.
*/
func GetSelection() gdclass.EditorSelection {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.EditorSelection(class(self).GetSelection(gc))
}

/*
Returns the editor's [EditorSettings] instance.
*/
func GetEditorSettings() gdclass.EditorSettings {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.EditorSettings(class(self).GetEditorSettings(gc))
}

/*
Returns mesh previews rendered at the given size as an [Array] of [Texture2D]s.
*/
func MakeMeshPreviews(meshes gd.ArrayOf[gdclass.Mesh], preview_size int) gd.ArrayOf[gdclass.Texture2D] {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.ArrayOf[gdclass.Texture2D](class(self).MakeMeshPreviews(gc, meshes, gd.Int(preview_size)))
}

/*
Sets the enabled status of a plugin. The plugin name is the same as its directory name.
*/
func SetPluginEnabled(plugin string, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).SetPluginEnabled(gc.String(plugin), enabled)
}

/*
Returns [code]true[/code] if the specified [param plugin] is enabled. The plugin name is the same as its directory name.
*/
func IsPluginEnabled(plugin string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).IsPluginEnabled(gc.String(plugin)))
}

/*
Returns the editor's [Theme].
[b]Note:[/b] When creating custom editor UI, prefer accessing theme items directly from your GUI nodes using the [code]get_theme_*[/code] methods.
*/
func GetEditorTheme() gdclass.Theme {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.Theme(class(self).GetEditorTheme(gc))
}

/*
Returns the main container of Godot editor's window. For example, you can use it to retrieve the size of the container and place your controls accordingly.
[b]Warning:[/b] Removing and freeing this node will render the editor useless and may cause a crash.
*/
func GetBaseControl() gdclass.Control {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.Control(class(self).GetBaseControl(gc))
}

/*
Returns the editor control responsible for main screen plugins and tools. Use it with plugins that implement [method EditorPlugin._has_main_screen].
[b]Note:[/b] This node is a [VBoxContainer], which means that if you add a [Control] child to it, you need to set the child's [member Control.size_flags_vertical] to [constant Control.SIZE_EXPAND_FILL] to make it use the full available space.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func GetEditorMainScreen() gdclass.VBoxContainer {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.VBoxContainer(class(self).GetEditorMainScreen(gc))
}

/*
Returns the editor's [ScriptEditor] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func GetScriptEditor() gdclass.ScriptEditor {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.ScriptEditor(class(self).GetScriptEditor(gc))
}

/*
Returns the 2D editor [SubViewport]. It does not have a camera. Instead, the view transforms are done directly and can be accessed with [member Viewport.global_canvas_transform].
*/
func GetEditorViewport2d() gdclass.SubViewport {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.SubViewport(class(self).GetEditorViewport2d(gc))
}

/*
Returns the specified 3D editor [SubViewport], from [code]0[/code] to [code]3[/code]. The viewport can be used to access the active editor cameras with [method Viewport.get_camera_3d].
*/
func GetEditorViewport3d() gdclass.SubViewport {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.SubViewport(class(self).GetEditorViewport3d(gc, gd.Int(0)))
}

/*
Sets the editor's current main screen to the one specified in [param name]. [param name] must match the title of the tab in question exactly (e.g. [code]2D[/code], [code]3D[/code], [code skip-lint]Script[/code], or [code]AssetLib[/code] for default tabs).
*/
func SetMainScreenEditor(name string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).SetMainScreenEditor(gc.String(name))
}

/*
Returns [code]true[/code] if multiple window support is enabled in the editor. Multiple window support is enabled if [i]all[/i] of these statements are true:
- [member EditorSettings.interface/multi_window/enable] is [code]true[/code].
- [member EditorSettings.interface/editor/single_window_mode] is [code]false[/code].
- [member Viewport.gui_embed_subwindows] is [code]false[/code]. This is forced to [code]true[/code] on platforms that don't support multiple windows such as Web, or when the [code]--single-window[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url] is used.
*/
func IsMultiWindowEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).IsMultiWindowEnabled())
}

/*
Returns the actual scale of the editor UI ([code]1.0[/code] being 100% scale). This can be used to adjust position and dimensions of the UI added by plugins.
[b]Note:[/b] This value is set via the [code]interface/editor/display_scale[/code] and [code]interface/editor/custom_display_scale[/code] editor settings. Editor must be restarted for changes to be properly applied.
*/
func GetEditorScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return float64(float64(class(self).GetEditorScale()))
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
func PopupDialog(dialog gdclass.Window) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).PopupDialog(dialog, gd.NewRect2i(0, 0, 0, 0))
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
func PopupDialogCentered(dialog gdclass.Window) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).PopupDialogCentered(dialog, gd.Vector2i{0, 0})
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered_ratio]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
func PopupDialogCenteredRatio(dialog gdclass.Window) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).PopupDialogCenteredRatio(dialog, gd.Float(0.8))
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered_clamped]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
func PopupDialogCenteredClamped(dialog gdclass.Window) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).PopupDialogCenteredClamped(dialog, gd.Vector2i{0, 0}, gd.Float(0.75))
}

/*
Returns the name of the currently activated feature profile. If the default profile is currently active, an empty string is returned instead.
In order to get a reference to the [EditorFeatureProfile], you must load the feature profile using [method EditorFeatureProfile.load_from_file].
[b]Note:[/b] Feature profiles created via the user interface are loaded from the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
func GetCurrentFeatureProfile() string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetCurrentFeatureProfile(gc).String())
}

/*
Selects and activates the specified feature profile with the given [param profile_name]. Set [param profile_name] to an empty string to reset to the default feature profile.
A feature profile can be created programmatically using the [EditorFeatureProfile] class.
[b]Note:[/b] The feature profile that gets activated must be located in the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. If a profile could not be found, an error occurs. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
func SetCurrentFeatureProfile(profile_name string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).SetCurrentFeatureProfile(gc.String(profile_name))
}

/*
Pops up an editor dialog for selecting a [Node] from the edited scene. The [param callback] must take a single argument of type [NodePath]. It is called on the selected [NodePath] or the empty path [code]^""[/code] if the dialog is canceled. If [param valid_types] is provided, the dialog will only show Nodes that match one of the listed Node types.
[b]Example:[/b]
[codeblock]
func _ready():
    if Engine.is_editor_hint():
        EditorInterface.popup_node_selector(_on_node_selected, ["Button"])

func _on_node_selected(node_path):
    if node_path.is_empty():
        print("node selection canceled")
    else:
        print("selected ", node_path)
[/codeblock]
*/
func PopupNodeSelector(callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).PopupNodeSelector(callback, ([1]gd.ArrayOf[gd.StringName]{}[0]))
}

/*
Pops up an editor dialog for selecting properties from [param object]. The [param callback] must take a single argument of type [NodePath]. It is called on the selected property path (see [method NodePath.get_as_property_path]) or the empty path [code]^""[/code] if the dialog is canceled. If [param type_filter] is provided, the dialog will only show properties that match one of the listed [enum Variant.Type] values.
[b]Example:[/b]
[codeblock]
func _ready():
    if Engine.is_editor_hint():
        EditorInterface.popup_property_selector(this, _on_property_selected, [TYPE_INT])

func _on_property_selected(property_path):
    if property_path.is_empty():
        print("property selection canceled")
    else:
        print("selected ", property_path)
[/codeblock]
*/
func PopupPropertySelector(obj gd.Object, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).PopupPropertySelector(obj, callback, gc.PackedInt32Slice(([1][]int32{}[0])))
}

/*
Returns the editor's [FileSystemDock] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func GetFileSystemDock() gdclass.FileSystemDock {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.FileSystemDock(class(self).GetFileSystemDock(gc))
}

/*
Selects the file, with the path provided by [param file], in the FileSystem dock.
*/
func SelectFile(file string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).SelectFile(gc.String(file))
}

/*
Returns an array containing the paths of the currently selected files (and directories) in the [FileSystemDock].
*/
func GetSelectedPaths() []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).GetSelectedPaths(gc).Strings(gc))
}

/*
Returns the current path being viewed in the [FileSystemDock].
*/
func GetCurrentPath() string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetCurrentPath(gc).String())
}

/*
Returns the current directory being viewed in the [FileSystemDock]. If a file is selected, its base directory will be returned using [method String.get_base_dir] instead.
*/
func GetCurrentDirectory() string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetCurrentDirectory(gc).String())
}

/*
Returns the editor's [EditorInspector] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func GetInspector() gdclass.EditorInspector {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.EditorInspector(class(self).GetInspector(gc))
}

/*
Shows the given property on the given [param object] in the editor's Inspector dock. If [param inspector_only] is [code]true[/code], plugins will not attempt to edit [param object].
*/
func InspectObject(obj gd.Object) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).InspectObject(obj, gc.String(""), false)
}

/*
Edits the given [Resource]. If the resource is a [Script] you can also edit it with [method edit_script] to specify the line and column position.
*/
func EditResource(resource gdclass.Resource) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).EditResource(resource)
}

/*
Edits the given [Node]. The node will be also selected if it's inside the scene tree.
*/
func EditNode(node gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).EditNode(node)
}

/*
Edits the given [Script]. The line and column on which to open the script can also be specified. The script will be open with the user-configured editor for the script's language which may be an external editor.
*/
func EditScript(script gdclass.Script) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).EditScript(script, gd.Int(-1), gd.Int(0), true)
}

/*
Opens the scene at the given path.
*/
func OpenSceneFromPath(scene_filepath string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).OpenSceneFromPath(gc.String(scene_filepath))
}

/*
Reloads the scene at the given path.
*/
func ReloadSceneFromPath(scene_filepath string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).ReloadSceneFromPath(gc.String(scene_filepath))
}

/*
Returns an [Array] with the file paths of the currently opened scenes.
*/
func GetOpenScenes() []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).GetOpenScenes(gc).Strings(gc))
}

/*
Returns the edited (current) scene's root [Node].
*/
func GetEditedSceneRoot() gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.Node(class(self).GetEditedSceneRoot(gc))
}

/*
Saves the currently active scene. Returns either [constant OK] or [constant ERR_CANT_CREATE].
*/
func SaveScene() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Error(class(self).SaveScene())
}

/*
Saves the currently active scene as a file at [param path].
*/
func SaveSceneAs(path string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).SaveSceneAs(gc.String(path), true)
}

/*
Saves all opened scenes in the editor.
*/
func SaveAllScenes() {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).SaveAllScenes()
}

/*
Marks the current scene tab as unsaved.
*/
func MarkSceneAsUnsaved() {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).MarkSceneAsUnsaved()
}

/*
Plays the main scene.
*/
func PlayMainScene() {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).PlayMainScene()
}

/*
Plays the currently active scene.
*/
func PlayCurrentScene() {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).PlayCurrentScene()
}

/*
Plays the scene specified by its filepath.
*/
func PlayCustomScene(scene_filepath string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).PlayCustomScene(gc.String(scene_filepath))
}

/*
Stops the scene that is currently playing.
*/
func StopPlayingScene() {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).StopPlayingScene()
}

/*
Returns [code]true[/code] if a scene is currently being played, [code]false[/code] otherwise. Paused scenes are considered as being played.
*/
func IsPlayingScene() bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).IsPlayingScene())
}

/*
Returns the name of the scene that is being played. If no scene is currently being played, returns an empty string.
*/
func GetPlayingScene() string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetPlayingScene(gc).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.EditorInterface
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func DistractionFreeMode() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDistractionFreeModeEnabled())
}

func SetDistractionFreeMode(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDistractionFreeMode(value)
}

func MovieMakerEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMovieMakerEnabled())
}

func SetMovieMakerEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMovieMakerEnabled(value)
}

/*
Restarts the editor. This closes the editor and then opens the same project. If [param save] is [code]true[/code], the project will be saved before restarting.
*/
//go:nosplit
func (self class) RestartEditor(save bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, save)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_restart_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the editor's [EditorCommandPalette] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetCommandPalette(ctx gd.Lifetime) gdclass.EditorCommandPalette {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_command_palette, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorCommandPalette
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the editor's [EditorFileSystem] instance.
*/
//go:nosplit
func (self class) GetResourceFilesystem(ctx gd.Lifetime) gdclass.EditorFileSystem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_resource_filesystem, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorFileSystem
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [EditorPaths] singleton.
*/
//go:nosplit
func (self class) GetEditorPaths(ctx gd.Lifetime) gdclass.EditorPaths {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_editor_paths, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorPaths
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the editor's [EditorResourcePreview] instance.
*/
//go:nosplit
func (self class) GetResourcePreviewer(ctx gd.Lifetime) gdclass.EditorResourcePreview {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_resource_previewer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorResourcePreview
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the editor's [EditorSelection] instance.
*/
//go:nosplit
func (self class) GetSelection(ctx gd.Lifetime) gdclass.EditorSelection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorSelection
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the editor's [EditorSettings] instance.
*/
//go:nosplit
func (self class) GetEditorSettings(ctx gd.Lifetime) gdclass.EditorSettings {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_editor_settings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorSettings
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns mesh previews rendered at the given size as an [Array] of [Texture2D]s.
*/
//go:nosplit
func (self class) MakeMeshPreviews(ctx gd.Lifetime, meshes gd.ArrayOf[gdclass.Mesh], preview_size gd.Int) gd.ArrayOf[gdclass.Texture2D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(meshes))
	callframe.Arg(frame, preview_size)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_make_mesh_previews, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Texture2D](ret)
}
/*
Sets the enabled status of a plugin. The plugin name is the same as its directory name.
*/
//go:nosplit
func (self class) SetPluginEnabled(plugin gd.String, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(plugin))
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_set_plugin_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the specified [param plugin] is enabled. The plugin name is the same as its directory name.
*/
//go:nosplit
func (self class) IsPluginEnabled(plugin gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(plugin))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_is_plugin_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the editor's [Theme].
[b]Note:[/b] When creating custom editor UI, prefer accessing theme items directly from your GUI nodes using the [code]get_theme_*[/code] methods.
*/
//go:nosplit
func (self class) GetEditorTheme(ctx gd.Lifetime) gdclass.Theme {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_editor_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Theme
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the main container of Godot editor's window. For example, you can use it to retrieve the size of the container and place your controls accordingly.
[b]Warning:[/b] Removing and freeing this node will render the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetBaseControl(ctx gd.Lifetime) gdclass.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_base_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Control
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the editor control responsible for main screen plugins and tools. Use it with plugins that implement [method EditorPlugin._has_main_screen].
[b]Note:[/b] This node is a [VBoxContainer], which means that if you add a [Control] child to it, you need to set the child's [member Control.size_flags_vertical] to [constant Control.SIZE_EXPAND_FILL] to make it use the full available space.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetEditorMainScreen(ctx gd.Lifetime) gdclass.VBoxContainer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_editor_main_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.VBoxContainer
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the editor's [ScriptEditor] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetScriptEditor(ctx gd.Lifetime) gdclass.ScriptEditor {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_script_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.ScriptEditor
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the 2D editor [SubViewport]. It does not have a camera. Instead, the view transforms are done directly and can be accessed with [member Viewport.global_canvas_transform].
*/
//go:nosplit
func (self class) GetEditorViewport2d(ctx gd.Lifetime) gdclass.SubViewport {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_editor_viewport_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.SubViewport
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the specified 3D editor [SubViewport], from [code]0[/code] to [code]3[/code]. The viewport can be used to access the active editor cameras with [method Viewport.get_camera_3d].
*/
//go:nosplit
func (self class) GetEditorViewport3d(ctx gd.Lifetime, idx gd.Int) gdclass.SubViewport {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_editor_viewport_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.SubViewport
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the editor's current main screen to the one specified in [param name]. [param name] must match the title of the tab in question exactly (e.g. [code]2D[/code], [code]3D[/code], [code skip-lint]Script[/code], or [code]AssetLib[/code] for default tabs).
*/
//go:nosplit
func (self class) SetMainScreenEditor(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_set_main_screen_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDistractionFreeMode(enter bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_set_distraction_free_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDistractionFreeModeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_is_distraction_free_mode_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if multiple window support is enabled in the editor. Multiple window support is enabled if [i]all[/i] of these statements are true:
- [member EditorSettings.interface/multi_window/enable] is [code]true[/code].
- [member EditorSettings.interface/editor/single_window_mode] is [code]false[/code].
- [member Viewport.gui_embed_subwindows] is [code]false[/code]. This is forced to [code]true[/code] on platforms that don't support multiple windows such as Web, or when the [code]--single-window[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url] is used.
*/
//go:nosplit
func (self class) IsMultiWindowEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_is_multi_window_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the actual scale of the editor UI ([code]1.0[/code] being 100% scale). This can be used to adjust position and dimensions of the UI added by plugins.
[b]Note:[/b] This value is set via the [code]interface/editor/display_scale[/code] and [code]interface/editor/custom_display_scale[/code] editor settings. Editor must be restarted for changes to be properly applied.
*/
//go:nosplit
func (self class) GetEditorScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_editor_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
//go:nosplit
func (self class) PopupDialog(dialog gdclass.Window, rect gd.Rect2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dialog[0].AsPointer())[0])
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_popup_dialog, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
//go:nosplit
func (self class) PopupDialogCentered(dialog gdclass.Window, minsize gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dialog[0].AsPointer())[0])
	callframe.Arg(frame, minsize)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_popup_dialog_centered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered_ratio]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
//go:nosplit
func (self class) PopupDialogCenteredRatio(dialog gdclass.Window, ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dialog[0].AsPointer())[0])
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_popup_dialog_centered_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered_clamped]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
//go:nosplit
func (self class) PopupDialogCenteredClamped(dialog gdclass.Window, minsize gd.Vector2i, fallback_ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dialog[0].AsPointer())[0])
	callframe.Arg(frame, minsize)
	callframe.Arg(frame, fallback_ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_popup_dialog_centered_clamped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the currently activated feature profile. If the default profile is currently active, an empty string is returned instead.
In order to get a reference to the [EditorFeatureProfile], you must load the feature profile using [method EditorFeatureProfile.load_from_file].
[b]Note:[/b] Feature profiles created via the user interface are loaded from the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
//go:nosplit
func (self class) GetCurrentFeatureProfile(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_current_feature_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Selects and activates the specified feature profile with the given [param profile_name]. Set [param profile_name] to an empty string to reset to the default feature profile.
A feature profile can be created programmatically using the [EditorFeatureProfile] class.
[b]Note:[/b] The feature profile that gets activated must be located in the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. If a profile could not be found, an error occurs. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
//go:nosplit
func (self class) SetCurrentFeatureProfile(profile_name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(profile_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_set_current_feature_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Pops up an editor dialog for selecting a [Node] from the edited scene. The [param callback] must take a single argument of type [NodePath]. It is called on the selected [NodePath] or the empty path [code]^""[/code] if the dialog is canceled. If [param valid_types] is provided, the dialog will only show Nodes that match one of the listed Node types.
[b]Example:[/b]
[codeblock]
func _ready():
    if Engine.is_editor_hint():
        EditorInterface.popup_node_selector(_on_node_selected, ["Button"])

func _on_node_selected(node_path):
    if node_path.is_empty():
        print("node selection canceled")
    else:
        print("selected ", node_path)
[/codeblock]
*/
//go:nosplit
func (self class) PopupNodeSelector(callback gd.Callable, valid_types gd.ArrayOf[gd.StringName])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(valid_types))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_popup_node_selector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Pops up an editor dialog for selecting properties from [param object]. The [param callback] must take a single argument of type [NodePath]. It is called on the selected property path (see [method NodePath.get_as_property_path]) or the empty path [code]^""[/code] if the dialog is canceled. If [param type_filter] is provided, the dialog will only show properties that match one of the listed [enum Variant.Type] values.
[b]Example:[/b]
[codeblock]
func _ready():
    if Engine.is_editor_hint():
        EditorInterface.popup_property_selector(this, _on_property_selected, [TYPE_INT])

func _on_property_selected(property_path):
    if property_path.is_empty():
        print("property selection canceled")
    else:
        print("selected ", property_path)
[/codeblock]
*/
//go:nosplit
func (self class) PopupPropertySelector(obj gd.Object, callback gd.Callable, type_filter gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(type_filter))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_popup_property_selector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the editor's [FileSystemDock] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetFileSystemDock(ctx gd.Lifetime) gdclass.FileSystemDock {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_file_system_dock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.FileSystemDock
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Selects the file, with the path provided by [param file], in the FileSystem dock.
*/
//go:nosplit
func (self class) SelectFile(file gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_select_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array containing the paths of the currently selected files (and directories) in the [FileSystemDock].
*/
//go:nosplit
func (self class) GetSelectedPaths(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_selected_paths, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the current path being viewed in the [FileSystemDock].
*/
//go:nosplit
func (self class) GetCurrentPath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_current_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the current directory being viewed in the [FileSystemDock]. If a file is selected, its base directory will be returned using [method String.get_base_dir] instead.
*/
//go:nosplit
func (self class) GetCurrentDirectory(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_current_directory, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the editor's [EditorInspector] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetInspector(ctx gd.Lifetime) gdclass.EditorInspector {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_inspector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorInspector
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Shows the given property on the given [param object] in the editor's Inspector dock. If [param inspector_only] is [code]true[/code], plugins will not attempt to edit [param object].
*/
//go:nosplit
func (self class) InspectObject(obj gd.Object, for_property gd.String, inspector_only bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(for_property))
	callframe.Arg(frame, inspector_only)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_inspect_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Edits the given [Resource]. If the resource is a [Script] you can also edit it with [method edit_script] to specify the line and column position.
*/
//go:nosplit
func (self class) EditResource(resource gdclass.Resource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(resource[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_edit_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Edits the given [Node]. The node will be also selected if it's inside the scene tree.
*/
//go:nosplit
func (self class) EditNode(node gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_edit_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Edits the given [Script]. The line and column on which to open the script can also be specified. The script will be open with the user-configured editor for the script's language which may be an external editor.
*/
//go:nosplit
func (self class) EditScript(script gdclass.Script, line gd.Int, column gd.Int, grab_focus bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(script[0].AsPointer())[0])
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	callframe.Arg(frame, grab_focus)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_edit_script, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Opens the scene at the given path.
*/
//go:nosplit
func (self class) OpenSceneFromPath(scene_filepath gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(scene_filepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_open_scene_from_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Reloads the scene at the given path.
*/
//go:nosplit
func (self class) ReloadSceneFromPath(scene_filepath gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(scene_filepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_reload_scene_from_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an [Array] with the file paths of the currently opened scenes.
*/
//go:nosplit
func (self class) GetOpenScenes(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_open_scenes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the edited (current) scene's root [Node].
*/
//go:nosplit
func (self class) GetEditedSceneRoot(ctx gd.Lifetime) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_edited_scene_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Saves the currently active scene. Returns either [constant OK] or [constant ERR_CANT_CREATE].
*/
//go:nosplit
func (self class) SaveScene() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_save_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Saves the currently active scene as a file at [param path].
*/
//go:nosplit
func (self class) SaveSceneAs(path gd.String, with_preview bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, with_preview)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_save_scene_as, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Saves all opened scenes in the editor.
*/
//go:nosplit
func (self class) SaveAllScenes()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_save_all_scenes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Marks the current scene tab as unsaved.
*/
//go:nosplit
func (self class) MarkSceneAsUnsaved()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_mark_scene_as_unsaved, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Plays the main scene.
*/
//go:nosplit
func (self class) PlayMainScene()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_play_main_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Plays the currently active scene.
*/
//go:nosplit
func (self class) PlayCurrentScene()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_play_current_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Plays the scene specified by its filepath.
*/
//go:nosplit
func (self class) PlayCustomScene(scene_filepath gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(scene_filepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_play_custom_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops the scene that is currently playing.
*/
//go:nosplit
func (self class) StopPlayingScene()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_stop_playing_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if a scene is currently being played, [code]false[/code] otherwise. Paused scenes are considered as being played.
*/
//go:nosplit
func (self class) IsPlayingScene() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_is_playing_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the scene that is being played. If no scene is currently being played, returns an empty string.
*/
//go:nosplit
func (self class) GetPlayingScene(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_get_playing_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMovieMakerEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_set_movie_maker_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMovieMakerEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInterface.Bind_is_movie_maker_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("EditorInterface", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
