// Package EditorInterface provides methods for working with EditorInterface object instances.
package EditorInterface

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/NodePath"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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
var self [1]gdclass.EditorInterface
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.EditorInterface)
	self = *(*[1]gdclass.EditorInterface)(unsafe.Pointer(&obj))
}

/*
Restarts the editor. This closes the editor and then opens the same project. If [param save] is [code]true[/code], the project will be saved before restarting.
*/
func RestartEditor() {
	once.Do(singleton)
	class(self).RestartEditor(true)
}

/*
Returns the editor's [EditorCommandPalette] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func GetCommandPalette() [1]gdclass.EditorCommandPalette {
	once.Do(singleton)
	return [1]gdclass.EditorCommandPalette(class(self).GetCommandPalette())
}

/*
Returns the editor's [EditorFileSystem] instance.
*/
func GetResourceFilesystem() [1]gdclass.EditorFileSystem {
	once.Do(singleton)
	return [1]gdclass.EditorFileSystem(class(self).GetResourceFilesystem())
}

/*
Returns the [EditorPaths] singleton.
*/
func GetEditorPaths() [1]gdclass.EditorPaths {
	once.Do(singleton)
	return [1]gdclass.EditorPaths(class(self).GetEditorPaths())
}

/*
Returns the editor's [EditorResourcePreview] instance.
*/
func GetResourcePreviewer() [1]gdclass.EditorResourcePreview {
	once.Do(singleton)
	return [1]gdclass.EditorResourcePreview(class(self).GetResourcePreviewer())
}

/*
Returns the editor's [EditorSelection] instance.
*/
func GetSelection() [1]gdclass.EditorSelection {
	once.Do(singleton)
	return [1]gdclass.EditorSelection(class(self).GetSelection())
}

/*
Returns the editor's [EditorSettings] instance.
*/
func GetEditorSettings() [1]gdclass.EditorSettings {
	once.Do(singleton)
	return [1]gdclass.EditorSettings(class(self).GetEditorSettings())
}

/*
Returns mesh previews rendered at the given size as an [Array] of [Texture2D]s.
*/
func MakeMeshPreviews(meshes gd.Array, preview_size int) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).MakeMeshPreviews(meshes, gd.Int(preview_size)))
}

/*
Sets the enabled status of a plugin. The plugin name is the same as its directory name.
*/
func SetPluginEnabled(plugin string, enabled bool) {
	once.Do(singleton)
	class(self).SetPluginEnabled(gd.NewString(plugin), enabled)
}

/*
Returns [code]true[/code] if the specified [param plugin] is enabled. The plugin name is the same as its directory name.
*/
func IsPluginEnabled(plugin string) bool {
	once.Do(singleton)
	return bool(class(self).IsPluginEnabled(gd.NewString(plugin)))
}

/*
Returns the editor's [Theme].
[b]Note:[/b] When creating custom editor UI, prefer accessing theme items directly from your GUI nodes using the [code]get_theme_*[/code] methods.
*/
func GetEditorTheme() [1]gdclass.Theme {
	once.Do(singleton)
	return [1]gdclass.Theme(class(self).GetEditorTheme())
}

/*
Returns the main container of Godot editor's window. For example, you can use it to retrieve the size of the container and place your controls accordingly.
[b]Warning:[/b] Removing and freeing this node will render the editor useless and may cause a crash.
*/
func GetBaseControl() [1]gdclass.Control {
	once.Do(singleton)
	return [1]gdclass.Control(class(self).GetBaseControl())
}

/*
Returns the editor control responsible for main screen plugins and tools. Use it with plugins that implement [method EditorPlugin._has_main_screen].
[b]Note:[/b] This node is a [VBoxContainer], which means that if you add a [Control] child to it, you need to set the child's [member Control.size_flags_vertical] to [constant Control.SIZE_EXPAND_FILL] to make it use the full available space.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func GetEditorMainScreen() [1]gdclass.VBoxContainer {
	once.Do(singleton)
	return [1]gdclass.VBoxContainer(class(self).GetEditorMainScreen())
}

/*
Returns the editor's [ScriptEditor] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func GetScriptEditor() [1]gdclass.ScriptEditor {
	once.Do(singleton)
	return [1]gdclass.ScriptEditor(class(self).GetScriptEditor())
}

/*
Returns the 2D editor [SubViewport]. It does not have a camera. Instead, the view transforms are done directly and can be accessed with [member Viewport.global_canvas_transform].
*/
func GetEditorViewport2d() [1]gdclass.SubViewport {
	once.Do(singleton)
	return [1]gdclass.SubViewport(class(self).GetEditorViewport2d())
}

/*
Returns the specified 3D editor [SubViewport], from [code]0[/code] to [code]3[/code]. The viewport can be used to access the active editor cameras with [method Viewport.get_camera_3d].
*/
func GetEditorViewport3d() [1]gdclass.SubViewport {
	once.Do(singleton)
	return [1]gdclass.SubViewport(class(self).GetEditorViewport3d(gd.Int(0)))
}

/*
Sets the editor's current main screen to the one specified in [param name]. [param name] must match the title of the tab in question exactly (e.g. [code]2D[/code], [code]3D[/code], [code skip-lint]Script[/code], or [code]AssetLib[/code] for default tabs).
*/
func SetMainScreenEditor(name string) {
	once.Do(singleton)
	class(self).SetMainScreenEditor(gd.NewString(name))
}

/*
Returns [code]true[/code] if multiple window support is enabled in the editor. Multiple window support is enabled if [i]all[/i] of these statements are true:
- [member EditorSettings.interface/multi_window/enable] is [code]true[/code].
- [member EditorSettings.interface/editor/single_window_mode] is [code]false[/code].
- [member Viewport.gui_embed_subwindows] is [code]false[/code]. This is forced to [code]true[/code] on platforms that don't support multiple windows such as Web, or when the [code]--single-window[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url] is used.
*/
func IsMultiWindowEnabled() bool {
	once.Do(singleton)
	return bool(class(self).IsMultiWindowEnabled())
}

/*
Returns the actual scale of the editor UI ([code]1.0[/code] being 100% scale). This can be used to adjust position and dimensions of the UI added by plugins.
[b]Note:[/b] This value is set via the [code]interface/editor/display_scale[/code] and [code]interface/editor/custom_display_scale[/code] editor settings. Editor must be restarted for changes to be properly applied.
*/
func GetEditorScale() Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetEditorScale()))
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
func PopupDialog(dialog [1]gdclass.Window) {
	once.Do(singleton)
	class(self).PopupDialog(dialog, gd.Rect2i(gd.NewRect2i(0, 0, 0, 0)))
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
func PopupDialogCentered(dialog [1]gdclass.Window) {
	once.Do(singleton)
	class(self).PopupDialogCentered(dialog, gd.Vector2i(gd.Vector2i{0, 0}))
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered_ratio]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
func PopupDialogCenteredRatio(dialog [1]gdclass.Window) {
	once.Do(singleton)
	class(self).PopupDialogCenteredRatio(dialog, gd.Float(0.8))
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered_clamped]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
func PopupDialogCenteredClamped(dialog [1]gdclass.Window) {
	once.Do(singleton)
	class(self).PopupDialogCenteredClamped(dialog, gd.Vector2i(gd.Vector2i{0, 0}), gd.Float(0.75))
}

/*
Returns the name of the currently activated feature profile. If the default profile is currently active, an empty string is returned instead.
In order to get a reference to the [EditorFeatureProfile], you must load the feature profile using [method EditorFeatureProfile.load_from_file].
[b]Note:[/b] Feature profiles created via the user interface are loaded from the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
func GetCurrentFeatureProfile() string {
	once.Do(singleton)
	return string(class(self).GetCurrentFeatureProfile().String())
}

/*
Selects and activates the specified feature profile with the given [param profile_name]. Set [param profile_name] to an empty string to reset to the default feature profile.
A feature profile can be created programmatically using the [EditorFeatureProfile] class.
[b]Note:[/b] The feature profile that gets activated must be located in the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. If a profile could not be found, an error occurs. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
func SetCurrentFeatureProfile(profile_name string) {
	once.Do(singleton)
	class(self).SetCurrentFeatureProfile(gd.NewString(profile_name))
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
func PopupNodeSelector(callback func(selected NodePath.String)) {
	once.Do(singleton)
	class(self).PopupNodeSelector(gd.NewCallable(callback), [1]gd.Array{}[0])
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
func PopupPropertySelector(obj Object.Instance, callback func(selected NodePath.String)) {
	once.Do(singleton)
	class(self).PopupPropertySelector(obj, gd.NewCallable(callback), gd.NewPackedInt32Slice([1][]int32{}[0]))
}

/*
Returns the editor's [FileSystemDock] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func GetFileSystemDock() [1]gdclass.FileSystemDock {
	once.Do(singleton)
	return [1]gdclass.FileSystemDock(class(self).GetFileSystemDock())
}

/*
Selects the file, with the path provided by [param file], in the FileSystem dock.
*/
func SelectFile(file string) {
	once.Do(singleton)
	class(self).SelectFile(gd.NewString(file))
}

/*
Returns an array containing the paths of the currently selected files (and directories) in the [FileSystemDock].
*/
func GetSelectedPaths() []string {
	once.Do(singleton)
	return []string(class(self).GetSelectedPaths().Strings())
}

/*
Returns the current path being viewed in the [FileSystemDock].
*/
func GetCurrentPath() string {
	once.Do(singleton)
	return string(class(self).GetCurrentPath().String())
}

/*
Returns the current directory being viewed in the [FileSystemDock]. If a file is selected, its base directory will be returned using [method String.get_base_dir] instead.
*/
func GetCurrentDirectory() string {
	once.Do(singleton)
	return string(class(self).GetCurrentDirectory().String())
}

/*
Returns the editor's [EditorInspector] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func GetInspector() [1]gdclass.EditorInspector {
	once.Do(singleton)
	return [1]gdclass.EditorInspector(class(self).GetInspector())
}

/*
Shows the given property on the given [param object] in the editor's Inspector dock. If [param inspector_only] is [code]true[/code], plugins will not attempt to edit [param object].
*/
func InspectObject(obj Object.Instance) {
	once.Do(singleton)
	class(self).InspectObject(obj, gd.NewString(""), false)
}

/*
Edits the given [Resource]. If the resource is a [Script] you can also edit it with [method edit_script] to specify the line and column position.
*/
func EditResource(resource [1]gdclass.Resource) {
	once.Do(singleton)
	class(self).EditResource(resource)
}

/*
Edits the given [Node]. The node will be also selected if it's inside the scene tree.
*/
func EditNode(node [1]gdclass.Node) {
	once.Do(singleton)
	class(self).EditNode(node)
}

/*
Edits the given [Script]. The line and column on which to open the script can also be specified. The script will be open with the user-configured editor for the script's language which may be an external editor.
*/
func EditScript(script [1]gdclass.Script) {
	once.Do(singleton)
	class(self).EditScript(script, gd.Int(-1), gd.Int(0), true)
}

/*
Opens the scene at the given path.
*/
func OpenSceneFromPath(scene_filepath string) {
	once.Do(singleton)
	class(self).OpenSceneFromPath(gd.NewString(scene_filepath))
}

/*
Reloads the scene at the given path.
*/
func ReloadSceneFromPath(scene_filepath string) {
	once.Do(singleton)
	class(self).ReloadSceneFromPath(gd.NewString(scene_filepath))
}

/*
Returns an [Array] with the file paths of the currently opened scenes.
*/
func GetOpenScenes() []string {
	once.Do(singleton)
	return []string(class(self).GetOpenScenes().Strings())
}

/*
Returns the edited (current) scene's root [Node].
*/
func GetEditedSceneRoot() [1]gdclass.Node {
	once.Do(singleton)
	return [1]gdclass.Node(class(self).GetEditedSceneRoot())
}

/*
Saves the currently active scene. Returns either [constant OK] or [constant ERR_CANT_CREATE].
*/
func SaveScene() error {
	once.Do(singleton)
	return error(class(self).SaveScene())
}

/*
Saves the currently active scene as a file at [param path].
*/
func SaveSceneAs(path string) {
	once.Do(singleton)
	class(self).SaveSceneAs(gd.NewString(path), true)
}

/*
Saves all opened scenes in the editor.
*/
func SaveAllScenes() {
	once.Do(singleton)
	class(self).SaveAllScenes()
}

/*
Marks the current scene tab as unsaved.
*/
func MarkSceneAsUnsaved() {
	once.Do(singleton)
	class(self).MarkSceneAsUnsaved()
}

/*
Plays the main scene.
*/
func PlayMainScene() {
	once.Do(singleton)
	class(self).PlayMainScene()
}

/*
Plays the currently active scene.
*/
func PlayCurrentScene() {
	once.Do(singleton)
	class(self).PlayCurrentScene()
}

/*
Plays the scene specified by its filepath.
*/
func PlayCustomScene(scene_filepath string) {
	once.Do(singleton)
	class(self).PlayCustomScene(gd.NewString(scene_filepath))
}

/*
Stops the scene that is currently playing.
*/
func StopPlayingScene() {
	once.Do(singleton)
	class(self).StopPlayingScene()
}

/*
Returns [code]true[/code] if a scene is currently being played, [code]false[/code] otherwise. Paused scenes are considered as being played.
*/
func IsPlayingScene() bool {
	once.Do(singleton)
	return bool(class(self).IsPlayingScene())
}

/*
Returns the name of the scene that is being played. If no scene is currently being played, returns an empty string.
*/
func GetPlayingScene() string {
	once.Do(singleton)
	return string(class(self).GetPlayingScene().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.EditorInterface

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

func DistractionFreeMode() bool {
	return bool(class(self).IsDistractionFreeModeEnabled())
}

func SetDistractionFreeMode(value bool) {
	class(self).SetDistractionFreeMode(value)
}

func MovieMakerEnabled() bool {
	return bool(class(self).IsMovieMakerEnabled())
}

func SetMovieMakerEnabled(value bool) {
	class(self).SetMovieMakerEnabled(value)
}

/*
Restarts the editor. This closes the editor and then opens the same project. If [param save] is [code]true[/code], the project will be saved before restarting.
*/
//go:nosplit
func (self class) RestartEditor(save bool) {
	var frame = callframe.New()
	callframe.Arg(frame, save)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_restart_editor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the editor's [EditorCommandPalette] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetCommandPalette() [1]gdclass.EditorCommandPalette {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_command_palette, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorCommandPalette{gd.PointerLifetimeBoundTo[gdclass.EditorCommandPalette](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the editor's [EditorFileSystem] instance.
*/
//go:nosplit
func (self class) GetResourceFilesystem() [1]gdclass.EditorFileSystem {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_resource_filesystem, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorFileSystem{gd.PointerLifetimeBoundTo[gdclass.EditorFileSystem](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [EditorPaths] singleton.
*/
//go:nosplit
func (self class) GetEditorPaths() [1]gdclass.EditorPaths {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_editor_paths, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorPaths{gd.PointerLifetimeBoundTo[gdclass.EditorPaths](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the editor's [EditorResourcePreview] instance.
*/
//go:nosplit
func (self class) GetResourcePreviewer() [1]gdclass.EditorResourcePreview {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_resource_previewer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorResourcePreview{gd.PointerLifetimeBoundTo[gdclass.EditorResourcePreview](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the editor's [EditorSelection] instance.
*/
//go:nosplit
func (self class) GetSelection() [1]gdclass.EditorSelection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorSelection{gd.PointerLifetimeBoundTo[gdclass.EditorSelection](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the editor's [EditorSettings] instance.
*/
//go:nosplit
func (self class) GetEditorSettings() [1]gdclass.EditorSettings {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_editor_settings, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorSettings{gd.PointerWithOwnershipTransferredToGo[gdclass.EditorSettings](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns mesh previews rendered at the given size as an [Array] of [Texture2D]s.
*/
//go:nosplit
func (self class) MakeMeshPreviews(meshes gd.Array, preview_size gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(meshes))
	callframe.Arg(frame, preview_size)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_make_mesh_previews, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the enabled status of a plugin. The plugin name is the same as its directory name.
*/
//go:nosplit
func (self class) SetPluginEnabled(plugin gd.String, enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin))
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_set_plugin_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified [param plugin] is enabled. The plugin name is the same as its directory name.
*/
//go:nosplit
func (self class) IsPluginEnabled(plugin gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_is_plugin_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the editor's [Theme].
[b]Note:[/b] When creating custom editor UI, prefer accessing theme items directly from your GUI nodes using the [code]get_theme_*[/code] methods.
*/
//go:nosplit
func (self class) GetEditorTheme() [1]gdclass.Theme {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_editor_theme, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Theme{gd.PointerWithOwnershipTransferredToGo[gdclass.Theme](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the main container of Godot editor's window. For example, you can use it to retrieve the size of the container and place your controls accordingly.
[b]Warning:[/b] Removing and freeing this node will render the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetBaseControl() [1]gdclass.Control {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_base_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Control{gd.PointerLifetimeBoundTo[gdclass.Control](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the editor control responsible for main screen plugins and tools. Use it with plugins that implement [method EditorPlugin._has_main_screen].
[b]Note:[/b] This node is a [VBoxContainer], which means that if you add a [Control] child to it, you need to set the child's [member Control.size_flags_vertical] to [constant Control.SIZE_EXPAND_FILL] to make it use the full available space.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetEditorMainScreen() [1]gdclass.VBoxContainer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_editor_main_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.VBoxContainer{gd.PointerLifetimeBoundTo[gdclass.VBoxContainer](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the editor's [ScriptEditor] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetScriptEditor() [1]gdclass.ScriptEditor {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_script_editor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ScriptEditor{gd.PointerLifetimeBoundTo[gdclass.ScriptEditor](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the 2D editor [SubViewport]. It does not have a camera. Instead, the view transforms are done directly and can be accessed with [member Viewport.global_canvas_transform].
*/
//go:nosplit
func (self class) GetEditorViewport2d() [1]gdclass.SubViewport {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_editor_viewport_2d, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SubViewport{gd.PointerLifetimeBoundTo[gdclass.SubViewport](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the specified 3D editor [SubViewport], from [code]0[/code] to [code]3[/code]. The viewport can be used to access the active editor cameras with [method Viewport.get_camera_3d].
*/
//go:nosplit
func (self class) GetEditorViewport3d(idx gd.Int) [1]gdclass.SubViewport {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_editor_viewport_3d, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SubViewport{gd.PointerLifetimeBoundTo[gdclass.SubViewport](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the editor's current main screen to the one specified in [param name]. [param name] must match the title of the tab in question exactly (e.g. [code]2D[/code], [code]3D[/code], [code skip-lint]Script[/code], or [code]AssetLib[/code] for default tabs).
*/
//go:nosplit
func (self class) SetMainScreenEditor(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_set_main_screen_editor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetDistractionFreeMode(enter bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enter)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_set_distraction_free_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDistractionFreeModeEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_is_distraction_free_mode_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_is_multi_window_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_editor_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
//go:nosplit
func (self class) PopupDialog(dialog [1]gdclass.Window, rect gd.Rect2i) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dialog[0])[0])
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_popup_dialog, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
//go:nosplit
func (self class) PopupDialogCentered(dialog [1]gdclass.Window, minsize gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dialog[0])[0])
	callframe.Arg(frame, minsize)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_popup_dialog_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered_ratio]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
//go:nosplit
func (self class) PopupDialogCenteredRatio(dialog [1]gdclass.Window, ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dialog[0])[0])
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_popup_dialog_centered_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Pops up the [param dialog] in the editor UI with [method Window.popup_exclusive_centered_clamped]. The dialog must have no current parent, otherwise the method fails.
See also [method Window.set_unparent_when_invisible].
*/
//go:nosplit
func (self class) PopupDialogCenteredClamped(dialog [1]gdclass.Window, minsize gd.Vector2i, fallback_ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dialog[0])[0])
	callframe.Arg(frame, minsize)
	callframe.Arg(frame, fallback_ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_popup_dialog_centered_clamped, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the name of the currently activated feature profile. If the default profile is currently active, an empty string is returned instead.
In order to get a reference to the [EditorFeatureProfile], you must load the feature profile using [method EditorFeatureProfile.load_from_file].
[b]Note:[/b] Feature profiles created via the user interface are loaded from the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
//go:nosplit
func (self class) GetCurrentFeatureProfile() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_current_feature_profile, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Selects and activates the specified feature profile with the given [param profile_name]. Set [param profile_name] to an empty string to reset to the default feature profile.
A feature profile can be created programmatically using the [EditorFeatureProfile] class.
[b]Note:[/b] The feature profile that gets activated must be located in the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. If a profile could not be found, an error occurs. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
//go:nosplit
func (self class) SetCurrentFeatureProfile(profile_name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(profile_name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_set_current_feature_profile, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) PopupNodeSelector(callback gd.Callable, valid_types gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(valid_types))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_popup_node_selector, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) PopupPropertySelector(obj [1]gd.Object, callback gd.Callable, type_filter gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(type_filter))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_popup_property_selector, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the editor's [FileSystemDock] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetFileSystemDock() [1]gdclass.FileSystemDock {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_file_system_dock, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.FileSystemDock{gd.PointerLifetimeBoundTo[gdclass.FileSystemDock](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Selects the file, with the path provided by [param file], in the FileSystem dock.
*/
//go:nosplit
func (self class) SelectFile(file gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(file))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_select_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array containing the paths of the currently selected files (and directories) in the [FileSystemDock].
*/
//go:nosplit
func (self class) GetSelectedPaths() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_selected_paths, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current path being viewed in the [FileSystemDock].
*/
//go:nosplit
func (self class) GetCurrentPath() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_current_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current directory being viewed in the [FileSystemDock]. If a file is selected, its base directory will be returned using [method String.get_base_dir] instead.
*/
//go:nosplit
func (self class) GetCurrentDirectory() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_current_directory, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the editor's [EditorInspector] instance.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetInspector() [1]gdclass.EditorInspector {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_inspector, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorInspector{gd.PointerLifetimeBoundTo[gdclass.EditorInspector](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Shows the given property on the given [param object] in the editor's Inspector dock. If [param inspector_only] is [code]true[/code], plugins will not attempt to edit [param object].
*/
//go:nosplit
func (self class) InspectObject(obj [1]gd.Object, for_property gd.String, inspector_only bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	callframe.Arg(frame, pointers.Get(for_property))
	callframe.Arg(frame, inspector_only)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_inspect_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Edits the given [Resource]. If the resource is a [Script] you can also edit it with [method edit_script] to specify the line and column position.
*/
//go:nosplit
func (self class) EditResource(resource [1]gdclass.Resource) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(resource[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_edit_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Edits the given [Node]. The node will be also selected if it's inside the scene tree.
*/
//go:nosplit
func (self class) EditNode(node [1]gdclass.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_edit_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Edits the given [Script]. The line and column on which to open the script can also be specified. The script will be open with the user-configured editor for the script's language which may be an external editor.
*/
//go:nosplit
func (self class) EditScript(script [1]gdclass.Script, line gd.Int, column gd.Int, grab_focus bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(script[0])[0])
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	callframe.Arg(frame, grab_focus)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_edit_script, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Opens the scene at the given path.
*/
//go:nosplit
func (self class) OpenSceneFromPath(scene_filepath gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_filepath))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_open_scene_from_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Reloads the scene at the given path.
*/
//go:nosplit
func (self class) ReloadSceneFromPath(scene_filepath gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_filepath))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_reload_scene_from_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an [Array] with the file paths of the currently opened scenes.
*/
//go:nosplit
func (self class) GetOpenScenes() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_open_scenes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the edited (current) scene's root [Node].
*/
//go:nosplit
func (self class) GetEditedSceneRoot() [1]gdclass.Node {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_edited_scene_root, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Saves the currently active scene. Returns either [constant OK] or [constant ERR_CANT_CREATE].
*/
//go:nosplit
func (self class) SaveScene() gd.Error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_save_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the currently active scene as a file at [param path].
*/
//go:nosplit
func (self class) SaveSceneAs(path gd.String, with_preview bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, with_preview)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_save_scene_as, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Saves all opened scenes in the editor.
*/
//go:nosplit
func (self class) SaveAllScenes() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_save_all_scenes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Marks the current scene tab as unsaved.
*/
//go:nosplit
func (self class) MarkSceneAsUnsaved() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_mark_scene_as_unsaved, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Plays the main scene.
*/
//go:nosplit
func (self class) PlayMainScene() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_play_main_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Plays the currently active scene.
*/
//go:nosplit
func (self class) PlayCurrentScene() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_play_current_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Plays the scene specified by its filepath.
*/
//go:nosplit
func (self class) PlayCustomScene(scene_filepath gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_filepath))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_play_custom_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops the scene that is currently playing.
*/
//go:nosplit
func (self class) StopPlayingScene() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_stop_playing_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if a scene is currently being played, [code]false[/code] otherwise. Paused scenes are considered as being played.
*/
//go:nosplit
func (self class) IsPlayingScene() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_is_playing_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the scene that is being played. If no scene is currently being played, returns an empty string.
*/
//go:nosplit
func (self class) GetPlayingScene() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_get_playing_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMovieMakerEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_set_movie_maker_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMovieMakerEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInterface.Bind_is_movie_maker_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("EditorInterface", func(ptr gd.Object) any {
		return [1]gdclass.EditorInterface{*(*gdclass.EditorInterface)(unsafe.Pointer(&ptr))}
	})
}

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
