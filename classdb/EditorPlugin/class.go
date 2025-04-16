// Package EditorPlugin provides methods for working with EditorPlugin object instances.
package EditorPlugin

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Node"
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
Plugins are used by the editor to extend functionality. The most common types of plugins are those which edit a given node or resource type, import plugins and export plugins. See also [EditorScript] to add functions to the editor.
[b]Note:[/b] Some names in this class contain "left" or "right" (e.g. [constant DOCK_SLOT_LEFT_UL]). These APIs assume left-to-right layout, and would be backwards when using right-to-left layout. These names are kept for compatibility reasons.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorPlugin)
*/
type Instance [1]gdclass.EditorPlugin
type Expanded [1]gdclass.EditorPlugin

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorPlugin() Instance
}
type Interface interface {
	//Called when there is a root node in the current edited scene, [method _handles] is implemented, and an [InputEvent] happens in the 2D viewport. If this method returns [code]true[/code], [param event] is intercepted by this [EditorPlugin], otherwise [param event] is forwarded to other Editor classes.
	//[codeblocks]
	//[gdscript]
	//# Prevents the InputEvent from reaching other Editor classes.
	//func _forward_canvas_gui_input(event):
	//    return true
	//[/gdscript]
	//[csharp]
	//// Prevents the InputEvent from reaching other Editor classes.
	//public override bool ForwardCanvasGuiInput(InputEvent @event)
	//{
	//    return true;
	//}
	//[/csharp]
	//[/codeblocks]
	//This method must return [code]false[/code] in order to forward the [InputEvent] to other Editor classes.
	//[codeblocks]
	//[gdscript]
	//# Consumes InputEventMouseMotion and forwards other InputEvent types.
	//func _forward_canvas_gui_input(event):
	//    if (event is InputEventMouseMotion):
	//        return true
	//    return false
	//[/gdscript]
	//[csharp]
	//// Consumes InputEventMouseMotion and forwards other InputEvent types.
	//public override bool _ForwardCanvasGuiInput(InputEvent @event)
	//{
	//    if (@event is InputEventMouseMotion)
	//    {
	//        return true;
	//    }
	//    return false;
	//}
	//[/csharp]
	//[/codeblocks]
	ForwardCanvasGuiInput(event [1]gdclass.InputEvent) bool
	//Called by the engine when the 2D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
	//[codeblocks]
	//[gdscript]
	//func _forward_canvas_draw_over_viewport(overlay):
	//    # Draw a circle at cursor position.
	//    overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)
	//
	//func _forward_canvas_gui_input(event):
	//    if event is InputEventMouseMotion:
	//        # Redraw viewport when cursor is moved.
	//        update_overlays()
	//        return true
	//    return false
	//[/gdscript]
	//[csharp]
	//public override void _ForwardCanvasDrawOverViewport(Control viewportControl)
	//{
	//    // Draw a circle at cursor position.
	//    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
	//}
	//
	//public override bool _ForwardCanvasGuiInput(InputEvent @event)
	//{
	//    if (@event is InputEventMouseMotion)
	//    {
	//        // Redraw viewport when cursor is moved.
	//        UpdateOverlays();
	//        return true;
	//    }
	//    return false;
	//}
	//[/csharp]
	//[/codeblocks]
	ForwardCanvasDrawOverViewport(viewport_control [1]gdclass.Control)
	//This method is the same as [method _forward_canvas_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
	//You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
	ForwardCanvasForceDrawOverViewport(viewport_control [1]gdclass.Control)
	//Called when there is a root node in the current edited scene, [method _handles] is implemented, and an [InputEvent] happens in the 3D viewport. The return value decides whether the [InputEvent] is consumed or forwarded to other [EditorPlugin]s. See [enum AfterGUIInput] for options.
	//[codeblocks]
	//[gdscript]
	//# Prevents the InputEvent from reaching other Editor classes.
	//func _forward_3d_gui_input(camera, event):
	//    return EditorPlugin.AFTER_GUI_INPUT_STOP
	//[/gdscript]
	//[csharp]
	//// Prevents the InputEvent from reaching other Editor classes.
	//public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)
	//{
	//    return EditorPlugin.AfterGuiInput.Stop;
	//}
	//[/csharp]
	//[/codeblocks]
	//This method must return [constant AFTER_GUI_INPUT_PASS] in order to forward the [InputEvent] to other Editor classes.
	//[codeblocks]
	//[gdscript]
	//# Consumes InputEventMouseMotion and forwards other InputEvent types.
	//func _forward_3d_gui_input(camera, event):
	//    return EditorPlugin.AFTER_GUI_INPUT_STOP if event is InputEventMouseMotion else EditorPlugin.AFTER_GUI_INPUT_PASS
	//[/gdscript]
	//[csharp]
	//// Consumes InputEventMouseMotion and forwards other InputEvent types.
	//public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)
	//{
	//    return @event is InputEventMouseMotion ? EditorPlugin.AfterGuiInput.Stop : EditorPlugin.AfterGuiInput.Pass;
	//}
	//[/csharp]
	//[/codeblocks]
	Forward3dGuiInput(viewport_camera [1]gdclass.Camera3D, event [1]gdclass.InputEvent) int
	//Called by the engine when the 3D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
	//[codeblocks]
	//[gdscript]
	//func _forward_3d_draw_over_viewport(overlay):
	//    # Draw a circle at cursor position.
	//    overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)
	//
	//func _forward_3d_gui_input(camera, event):
	//    if event is InputEventMouseMotion:
	//        # Redraw viewport when cursor is moved.
	//        update_overlays()
	//        return EditorPlugin.AFTER_GUI_INPUT_STOP
	//    return EditorPlugin.AFTER_GUI_INPUT_PASS
	//[/gdscript]
	//[csharp]
	//public override void _Forward3DDrawOverViewport(Control viewportControl)
	//{
	//    // Draw a circle at cursor position.
	//    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
	//}
	//
	//public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D viewportCamera, InputEvent @event)
	//{
	//    if (@event is InputEventMouseMotion)
	//    {
	//        // Redraw viewport when cursor is moved.
	//        UpdateOverlays();
	//        return EditorPlugin.AfterGuiInput.Stop;
	//    }
	//    return EditorPlugin.AfterGuiInput.Pass;
	//}
	//[/csharp]
	//[/codeblocks]
	Forward3dDrawOverViewport(viewport_control [1]gdclass.Control)
	//This method is the same as [method _forward_3d_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
	//You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
	Forward3dForceDrawOverViewport(viewport_control [1]gdclass.Control)
	//Override this method in your plugin to provide the name of the plugin when displayed in the Godot editor.
	//For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
	GetPluginName() string
	//Override this method in your plugin to return a [Texture2D] in order to give it an icon.
	//For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
	//Ideally, the plugin icon should be white with a transparent background and 16×16 pixels in size.
	//[codeblocks]
	//[gdscript]
	//func _get_plugin_icon():
	//    # You can use a custom icon:
	//    return preload("res://addons/my_plugin/my_plugin_icon.svg")
	//    # Or use a built-in icon:
	//    return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")
	//[/gdscript]
	//[csharp]
	//public override Texture2D _GetPluginIcon()
	//{
	//    // You can use a custom icon:
	//    return ResourceLoader.Load<Texture2D>("res://addons/my_plugin/my_plugin_icon.svg");
	//    // Or use a built-in icon:
	//    return EditorInterface.Singleton.GetEditorTheme().GetIcon("Node", "EditorIcons");
	//}
	//[/csharp]
	//[/codeblocks]
	GetPluginIcon() [1]gdclass.Texture2D
	//Returns [code]true[/code] if this is a main screen editor plugin (it goes in the workspace selector together with [b]2D[/b], [b]3D[/b], [b]Script[/b] and [b]AssetLib[/b]).
	//When the plugin's workspace is selected, other main screen plugins will be hidden, but your plugin will not appear automatically. It needs to be added as a child of [method EditorInterface.get_editor_main_screen] and made visible inside [method _make_visible].
	//Use [method _get_plugin_name] and [method _get_plugin_icon] to customize the plugin button's appearance.
	//[codeblock]
	//var plugin_control
	//
	//func _enter_tree():
	//    plugin_control = preload("my_plugin_control.tscn").instantiate()
	//    EditorInterface.get_editor_main_screen().add_child(plugin_control)
	//    plugin_control.hide()
	//
	//func _has_main_screen():
	//    return true
	//
	//func _make_visible(visible):
	//    plugin_control.visible = visible
	//
	//func _get_plugin_name():
	//    return "My Super Cool Plugin 3000"
	//
	//func _get_plugin_icon():
	//    return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")
	//[/codeblock]
	HasMainScreen() bool
	//This function will be called when the editor is requested to become visible. It is used for plugins that edit a specific object type.
	//Remember that you have to manage the visibility of all your editor controls manually.
	MakeVisible(visible bool)
	//This function is used for plugins that edit specific object types (nodes or resources). It requests the editor to edit the given object.
	//[param object] can be [code]null[/code] if the plugin was editing an object, but there is no longer any selected object handled by this plugin. It can be used to cleanup editing state.
	Edit(obj Object.Instance)
	//Implement this function if your plugin edits a specific type of object (Resource or Node). If you return [code]true[/code], then you will get the functions [method _edit] and [method _make_visible] called when the editor requests them. If you have declared the methods [method _forward_canvas_gui_input] and [method _forward_3d_gui_input] these will be called too.
	//[b]Note:[/b] Each plugin should handle only one type of objects at a time. If a plugin handles more types of objects and they are edited at the same time, it will result in errors.
	Handles(obj Object.Instance) bool
	//Override this method to provide a state data you want to be saved, like view position, grid settings, folding, etc. This is used when saving the scene (so state is kept when opening it again) and for switching tabs (so state can be restored when the tab returns). This data is automatically saved for each scene in an [code]editstate[/code] file in the editor metadata folder. If you want to store global (scene-independent) editor data for your plugin, you can use [method _get_window_layout] instead.
	//Use [method _set_state] to restore your saved state.
	//[b]Note:[/b] This method should not be used to save important settings that should persist with the project.
	//[b]Note:[/b] You must implement [method _get_plugin_name] for the state to be stored and restored correctly.
	//[codeblock]
	//func _get_state():
	//    var state = {"zoom": zoom, "preferred_color": my_color}
	//    return state
	//[/codeblock]
	GetState() map[any]any
	//Restore the state saved by [method _get_state]. This method is called when the current scene tab is changed in the editor.
	//[b]Note:[/b] Your plugin must implement [method _get_plugin_name], otherwise it will not be recognized and this method will not be called.
	//[codeblock]
	//func _set_state(data):
	//    zoom = data.get("zoom", 1.0)
	//    preferred_color = data.get("my_color", Color.WHITE)
	//[/codeblock]
	SetState(state map[any]any)
	//Clear all the state and reset the object being edited to zero. This ensures your plugin does not keep editing a currently existing node, or a node from the wrong scene.
	Clear()
	//Override this method to provide a custom message that lists unsaved changes. The editor will call this method when exiting or when closing a scene, and display the returned string in a confirmation dialog. Return empty string if the plugin has no unsaved changes.
	//When closing a scene, [param for_scene] is the path to the scene being closed. You can use it to handle built-in resources in that scene.
	//If the user confirms saving, [method _save_external_data] will be called, before closing the editor.
	//[codeblock]
	//func _get_unsaved_status(for_scene):
	//    if not unsaved:
	//        return ""
	//
	//    if for_scene.is_empty():
	//        return "Save changes in MyCustomPlugin before closing?"
	//    else:
	//        return "Scene %s has changes from MyCustomPlugin. Save before closing?" % for_scene.get_file()
	//
	//func _save_external_data():
	//    unsaved = false
	//[/codeblock]
	//If the plugin has no scene-specific changes, you can ignore the calls when closing scenes:
	//[codeblock]
	//func _get_unsaved_status(for_scene):
	//    if not for_scene.is_empty():
	//        return ""
	//[/codeblock]
	GetUnsavedStatus(for_scene string) string
	//This method is called after the editor saves the project or when it's closed. It asks the plugin to save edited external scenes/resources.
	SaveExternalData()
	//This method is called when the editor is about to save the project, switch to another tab, etc. It asks the plugin to apply any pending state changes to ensure consistency.
	//This is used, for example, in shader editors to let the plugin know that it must apply the shader code being written by the user to the object.
	ApplyChanges()
	//This is for editors that edit script-based objects. You can return a list of breakpoints in the format ([code]script:line[/code]), for example: [code]res://path_to_script.gd:25[/code].
	GetBreakpoints() []string
	//Restore the plugin GUI layout and data saved by [method _get_window_layout]. This method is called for every plugin on editor startup. Use the provided [param configuration] file to read your saved data.
	//[codeblock]
	//func _set_window_layout(configuration):
	//    $Window.position = configuration.get_value("MyPlugin", "window_position", Vector2())
	//    $Icon.modulate = configuration.get_value("MyPlugin", "icon_color", Color.WHITE)
	//[/codeblock]
	SetWindowLayout(configuration [1]gdclass.ConfigFile)
	//Override this method to provide the GUI layout of the plugin or any other data you want to be stored. This is used to save the project's editor layout when [method queue_save_layout] is called or the editor layout was changed (for example changing the position of a dock). The data is stored in the [code]editor_layout.cfg[/code] file in the editor metadata directory.
	//Use [method _set_window_layout] to restore your saved layout.
	//[codeblock]
	//func _get_window_layout(configuration):
	//    configuration.set_value("MyPlugin", "window_position", $Window.position)
	//    configuration.set_value("MyPlugin", "icon_color", $Icon.modulate)
	//[/codeblock]
	GetWindowLayout(configuration [1]gdclass.ConfigFile)
	//This method is called when the editor is about to run the project. The plugin can then perform required operations before the project runs.
	//This method must return a boolean. If this method returns [code]false[/code], the project will not run. The run is aborted immediately, so this also prevents all other plugins' [method _build] methods from running.
	Build() bool
	//Called by the engine when the user enables the [EditorPlugin] in the Plugin tab of the project settings window.
	EnablePlugin()
	//Called by the engine when the user disables the [EditorPlugin] in the Plugin tab of the project settings window.
	DisablePlugin()
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) ForwardCanvasGuiInput(event [1]gdclass.InputEvent) (_ bool)        { return }
func (self implementation) ForwardCanvasDrawOverViewport(viewport_control [1]gdclass.Control) { return }
func (self implementation) ForwardCanvasForceDrawOverViewport(viewport_control [1]gdclass.Control) {
	return
}
func (self implementation) Forward3dGuiInput(viewport_camera [1]gdclass.Camera3D, event [1]gdclass.InputEvent) (_ int) {
	return
}
func (self implementation) Forward3dDrawOverViewport(viewport_control [1]gdclass.Control) { return }
func (self implementation) Forward3dForceDrawOverViewport(viewport_control [1]gdclass.Control) {
	return
}
func (self implementation) GetPluginName() (_ string)                           { return }
func (self implementation) GetPluginIcon() (_ [1]gdclass.Texture2D)             { return }
func (self implementation) HasMainScreen() (_ bool)                             { return }
func (self implementation) MakeVisible(visible bool)                            { return }
func (self implementation) Edit(obj Object.Instance)                            { return }
func (self implementation) Handles(obj Object.Instance) (_ bool)                { return }
func (self implementation) GetState() (_ map[any]any)                           { return }
func (self implementation) SetState(state map[any]any)                          { return }
func (self implementation) Clear()                                              { return }
func (self implementation) GetUnsavedStatus(for_scene string) (_ string)        { return }
func (self implementation) SaveExternalData()                                   { return }
func (self implementation) ApplyChanges()                                       { return }
func (self implementation) GetBreakpoints() (_ []string)                        { return }
func (self implementation) SetWindowLayout(configuration [1]gdclass.ConfigFile) { return }
func (self implementation) GetWindowLayout(configuration [1]gdclass.ConfigFile) { return }
func (self implementation) Build() (_ bool)                                     { return }
func (self implementation) EnablePlugin()                                       { return }
func (self implementation) DisablePlugin()                                      { return }

/*
Called when there is a root node in the current edited scene, [method _handles] is implemented, and an [InputEvent] happens in the 2D viewport. If this method returns [code]true[/code], [param event] is intercepted by this [EditorPlugin], otherwise [param event] is forwarded to other Editor classes.
[codeblocks]
[gdscript]
# Prevents the InputEvent from reaching other Editor classes.
func _forward_canvas_gui_input(event):

	return true

[/gdscript]
[csharp]
// Prevents the InputEvent from reaching other Editor classes.
public override bool ForwardCanvasGuiInput(InputEvent @event)

	{
	    return true;
	}

[/csharp]
[/codeblocks]
This method must return [code]false[/code] in order to forward the [InputEvent] to other Editor classes.
[codeblocks]
[gdscript]
# Consumes InputEventMouseMotion and forwards other InputEvent types.
func _forward_canvas_gui_input(event):

	if (event is InputEventMouseMotion):
	    return true
	return false

[/gdscript]
[csharp]
// Consumes InputEventMouseMotion and forwards other InputEvent types.
public override bool _ForwardCanvasGuiInput(InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        return true;
	    }
	    return false;
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _forward_canvas_gui_input(impl func(ptr unsafe.Pointer, event [1]gdclass.InputEvent) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called by the engine when the 2D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
[codeblocks]
[gdscript]
func _forward_canvas_draw_over_viewport(overlay):

	# Draw a circle at cursor position.
	overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)

func _forward_canvas_gui_input(event):

	if event is InputEventMouseMotion:
	    # Redraw viewport when cursor is moved.
	    update_overlays()
	    return true
	return false

[/gdscript]
[csharp]
public override void _ForwardCanvasDrawOverViewport(Control viewportControl)

	{
	    // Draw a circle at cursor position.
	    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
	}

public override bool _ForwardCanvasGuiInput(InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        // Redraw viewport when cursor is moved.
	        UpdateOverlays();
	        return true;
	    }
	    return false;
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _forward_canvas_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control [1]gdclass.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport_control = [1]gdclass.Control{pointers.New[gdclass.Control]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
This method is the same as [method _forward_canvas_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
*/
func (Instance) _forward_canvas_force_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control [1]gdclass.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport_control = [1]gdclass.Control{pointers.New[gdclass.Control]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
Called when there is a root node in the current edited scene, [method _handles] is implemented, and an [InputEvent] happens in the 3D viewport. The return value decides whether the [InputEvent] is consumed or forwarded to other [EditorPlugin]s. See [enum AfterGUIInput] for options.
[codeblocks]
[gdscript]
# Prevents the InputEvent from reaching other Editor classes.
func _forward_3d_gui_input(camera, event):

	return EditorPlugin.AFTER_GUI_INPUT_STOP

[/gdscript]
[csharp]
// Prevents the InputEvent from reaching other Editor classes.
public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)

	{
	    return EditorPlugin.AfterGuiInput.Stop;
	}

[/csharp]
[/codeblocks]
This method must return [constant AFTER_GUI_INPUT_PASS] in order to forward the [InputEvent] to other Editor classes.
[codeblocks]
[gdscript]
# Consumes InputEventMouseMotion and forwards other InputEvent types.
func _forward_3d_gui_input(camera, event):

	return EditorPlugin.AFTER_GUI_INPUT_STOP if event is InputEventMouseMotion else EditorPlugin.AFTER_GUI_INPUT_PASS

[/gdscript]
[csharp]
// Consumes InputEventMouseMotion and forwards other InputEvent types.
public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)

	{
	    return @event is InputEventMouseMotion ? EditorPlugin.AfterGuiInput.Stop : EditorPlugin.AfterGuiInput.Pass;
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _forward_3d_gui_input(impl func(ptr unsafe.Pointer, viewport_camera [1]gdclass.Camera3D, event [1]gdclass.InputEvent) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport_camera = [1]gdclass.Camera3D{pointers.New[gdclass.Camera3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport_camera[0])
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, viewport_camera, event)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Called by the engine when the 3D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
[codeblocks]
[gdscript]
func _forward_3d_draw_over_viewport(overlay):

	# Draw a circle at cursor position.
	overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)

func _forward_3d_gui_input(camera, event):

	if event is InputEventMouseMotion:
	    # Redraw viewport when cursor is moved.
	    update_overlays()
	    return EditorPlugin.AFTER_GUI_INPUT_STOP
	return EditorPlugin.AFTER_GUI_INPUT_PASS

[/gdscript]
[csharp]
public override void _Forward3DDrawOverViewport(Control viewportControl)

	{
	    // Draw a circle at cursor position.
	    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
	}

public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D viewportCamera, InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        // Redraw viewport when cursor is moved.
	        UpdateOverlays();
	        return EditorPlugin.AfterGuiInput.Stop;
	    }
	    return EditorPlugin.AfterGuiInput.Pass;
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _forward_3d_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control [1]gdclass.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport_control = [1]gdclass.Control{pointers.New[gdclass.Control]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
This method is the same as [method _forward_3d_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
*/
func (Instance) _forward_3d_force_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control [1]gdclass.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport_control = [1]gdclass.Control{pointers.New[gdclass.Control]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
Override this method in your plugin to provide the name of the plugin when displayed in the Godot editor.
For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
*/
func (Instance) _get_plugin_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method in your plugin to return a [Texture2D] in order to give it an icon.
For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
Ideally, the plugin icon should be white with a transparent background and 16×16 pixels in size.
[codeblocks]
[gdscript]
func _get_plugin_icon():

	# You can use a custom icon:
	return preload("res://addons/my_plugin/my_plugin_icon.svg")
	# Or use a built-in icon:
	return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")

[/gdscript]
[csharp]
public override Texture2D _GetPluginIcon()

	{
	    // You can use a custom icon:
	    return ResourceLoader.Load<Texture2D>("res://addons/my_plugin/my_plugin_icon.svg");
	    // Or use a built-in icon:
	    return EditorInterface.Singleton.GetEditorTheme().GetIcon("Node", "EditorIcons");
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _get_plugin_icon(impl func(ptr unsafe.Pointer) [1]gdclass.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this is a main screen editor plugin (it goes in the workspace selector together with [b]2D[/b], [b]3D[/b], [b]Script[/b] and [b]AssetLib[/b]).
When the plugin's workspace is selected, other main screen plugins will be hidden, but your plugin will not appear automatically. It needs to be added as a child of [method EditorInterface.get_editor_main_screen] and made visible inside [method _make_visible].
Use [method _get_plugin_name] and [method _get_plugin_icon] to customize the plugin button's appearance.
[codeblock]
var plugin_control

func _enter_tree():

	plugin_control = preload("my_plugin_control.tscn").instantiate()
	EditorInterface.get_editor_main_screen().add_child(plugin_control)
	plugin_control.hide()

func _has_main_screen():

	return true

func _make_visible(visible):

	plugin_control.visible = visible

func _get_plugin_name():

	return "My Super Cool Plugin 3000"

func _get_plugin_icon():

	return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")

[/codeblock]
*/
func (Instance) _has_main_screen(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
This function will be called when the editor is requested to become visible. It is used for plugins that edit a specific object type.
Remember that you have to manage the visibility of all your editor controls manually.
*/
func (Instance) _make_visible(impl func(ptr unsafe.Pointer, visible bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var visible = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, visible)
	}
}

/*
This function is used for plugins that edit specific object types (nodes or resources). It requests the editor to edit the given object.
[param object] can be [code]null[/code] if the plugin was editing an object, but there is no longer any selected object handled by this plugin. It can be used to cleanup editing state.
*/
func (Instance) _edit(impl func(ptr unsafe.Pointer, obj Object.Instance)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj)
	}
}

/*
Implement this function if your plugin edits a specific type of object (Resource or Node). If you return [code]true[/code], then you will get the functions [method _edit] and [method _make_visible] called when the editor requests them. If you have declared the methods [method _forward_canvas_gui_input] and [method _forward_3d_gui_input] these will be called too.
[b]Note:[/b] Each plugin should handle only one type of objects at a time. If a plugin handles more types of objects and they are edited at the same time, it will result in errors.
*/
func (Instance) _handles(impl func(ptr unsafe.Pointer, obj Object.Instance) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to provide a state data you want to be saved, like view position, grid settings, folding, etc. This is used when saving the scene (so state is kept when opening it again) and for switching tabs (so state can be restored when the tab returns). This data is automatically saved for each scene in an [code]editstate[/code] file in the editor metadata folder. If you want to store global (scene-independent) editor data for your plugin, you can use [method _get_window_layout] instead.
Use [method _set_state] to restore your saved state.
[b]Note:[/b] This method should not be used to save important settings that should persist with the project.
[b]Note:[/b] You must implement [method _get_plugin_name] for the state to be stored and restored correctly.
[codeblock]
func _get_state():

	var state = {"zoom": zoom, "preferred_color": my_color}
	return state

[/codeblock]
*/
func (Instance) _get_state(impl func(ptr unsafe.Pointer) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Restore the state saved by [method _get_state]. This method is called when the current scene tab is changed in the editor.
[b]Note:[/b] Your plugin must implement [method _get_plugin_name], otherwise it will not be recognized and this method will not be called.
[codeblock]
func _set_state(data):

	zoom = data.get("zoom", 1.0)
	preferred_color = data.get("my_color", Color.WHITE)

[/codeblock]
*/
func (Instance) _set_state(impl func(ptr unsafe.Pointer, state map[any]any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var state = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalDictionary(state))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gd.DictionaryAs[map[any]any](state))
	}
}

/*
Clear all the state and reset the object being edited to zero. This ensures your plugin does not keep editing a currently existing node, or a node from the wrong scene.
*/
func (Instance) _clear(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Override this method to provide a custom message that lists unsaved changes. The editor will call this method when exiting or when closing a scene, and display the returned string in a confirmation dialog. Return empty string if the plugin has no unsaved changes.
When closing a scene, [param for_scene] is the path to the scene being closed. You can use it to handle built-in resources in that scene.
If the user confirms saving, [method _save_external_data] will be called, before closing the editor.
[codeblock]
func _get_unsaved_status(for_scene):

	if not unsaved:
	    return ""

	if for_scene.is_empty():
	    return "Save changes in MyCustomPlugin before closing?"
	else:
	    return "Scene %s has changes from MyCustomPlugin. Save before closing?" % for_scene.get_file()

func _save_external_data():

	unsaved = false

[/codeblock]
If the plugin has no scene-specific changes, you can ignore the calls when closing scenes:
[codeblock]
func _get_unsaved_status(for_scene):

	if not for_scene.is_empty():
	    return ""

[/codeblock]
*/
func (Instance) _get_unsaved_status(impl func(ptr unsafe.Pointer, for_scene string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var for_scene = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(for_scene))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_scene.String())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
This method is called after the editor saves the project or when it's closed. It asks the plugin to save edited external scenes/resources.
*/
func (Instance) _save_external_data(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
This method is called when the editor is about to save the project, switch to another tab, etc. It asks the plugin to apply any pending state changes to ensure consistency.
This is used, for example, in shader editors to let the plugin know that it must apply the shader code being written by the user to the object.
*/
func (Instance) _apply_changes(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
This is for editors that edit script-based objects. You can return a list of breakpoints in the format ([code]script:line[/code]), for example: [code]res://path_to_script.gd:25[/code].
*/
func (Instance) _get_breakpoints(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
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
Restore the plugin GUI layout and data saved by [method _get_window_layout]. This method is called for every plugin on editor startup. Use the provided [param configuration] file to read your saved data.
[codeblock]
func _set_window_layout(configuration):

	$Window.position = configuration.get_value("MyPlugin", "window_position", Vector2())
	$Icon.modulate = configuration.get_value("MyPlugin", "icon_color", Color.WHITE)

[/codeblock]
*/
func (Instance) _set_window_layout(impl func(ptr unsafe.Pointer, configuration [1]gdclass.ConfigFile)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var configuration = [1]gdclass.ConfigFile{pointers.New[gdclass.ConfigFile]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(configuration[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, configuration)
	}
}

/*
Override this method to provide the GUI layout of the plugin or any other data you want to be stored. This is used to save the project's editor layout when [method queue_save_layout] is called or the editor layout was changed (for example changing the position of a dock). The data is stored in the [code]editor_layout.cfg[/code] file in the editor metadata directory.
Use [method _set_window_layout] to restore your saved layout.
[codeblock]
func _get_window_layout(configuration):

	configuration.set_value("MyPlugin", "window_position", $Window.position)
	configuration.set_value("MyPlugin", "icon_color", $Icon.modulate)

[/codeblock]
*/
func (Instance) _get_window_layout(impl func(ptr unsafe.Pointer, configuration [1]gdclass.ConfigFile)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var configuration = [1]gdclass.ConfigFile{pointers.New[gdclass.ConfigFile]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(configuration[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, configuration)
	}
}

/*
This method is called when the editor is about to run the project. The plugin can then perform required operations before the project runs.
This method must return a boolean. If this method returns [code]false[/code], the project will not run. The run is aborted immediately, so this also prevents all other plugins' [method _build] methods from running.
*/
func (Instance) _build(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called by the engine when the user enables the [EditorPlugin] in the Plugin tab of the project settings window.
*/
func (Instance) _enable_plugin(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called by the engine when the user disables the [EditorPlugin] in the Plugin tab of the project settings window.
*/
func (Instance) _disable_plugin(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Adds a custom control to a container (see [enum CustomControlContainer]). There are many locations where custom controls can be added in the editor UI.
Please remember that you have to manage the visibility of your custom controls yourself (and likely hide it after adding it).
When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_container] and free it with [method Node.queue_free].
*/
func (self Instance) AddControlToContainer(container gdclass.EditorPluginCustomControlContainer, control [1]gdclass.Control) { //gd:EditorPlugin.add_control_to_container
	Advanced(self).AddControlToContainer(container, control)
}

/*
Adds a control to the bottom panel (together with Output, Debug, Animation, etc.). Returns a reference to the button added. It's up to you to hide/show the button when needed. When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_bottom_panel] and free it with [method Node.queue_free].
Optionally, you can specify a shortcut parameter. When pressed, this shortcut will toggle the bottom panel's visibility. See the default editor bottom panel shortcuts in the Editor Settings for inspiration. Per convention, they all use [kbd]Alt[/kbd] modifier.
*/
func (self Instance) AddControlToBottomPanel(control [1]gdclass.Control, title string) [1]gdclass.Button { //gd:EditorPlugin.add_control_to_bottom_panel
	return [1]gdclass.Button(Advanced(self).AddControlToBottomPanel(control, String.New(title), [1][1]gdclass.Shortcut{}[0]))
}

/*
Adds a control to the bottom panel (together with Output, Debug, Animation, etc.). Returns a reference to the button added. It's up to you to hide/show the button when needed. When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_bottom_panel] and free it with [method Node.queue_free].
Optionally, you can specify a shortcut parameter. When pressed, this shortcut will toggle the bottom panel's visibility. See the default editor bottom panel shortcuts in the Editor Settings for inspiration. Per convention, they all use [kbd]Alt[/kbd] modifier.
*/
func (self Expanded) AddControlToBottomPanel(control [1]gdclass.Control, title string, shortcut [1]gdclass.Shortcut) [1]gdclass.Button { //gd:EditorPlugin.add_control_to_bottom_panel
	return [1]gdclass.Button(Advanced(self).AddControlToBottomPanel(control, String.New(title), shortcut))
}

/*
Adds the control to a specific dock slot (see [enum DockSlot] for options).
If the dock is repositioned and as long as the plugin is active, the editor will save the dock position on further sessions.
When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_docks] and free it with [method Node.queue_free].
Optionally, you can specify a shortcut parameter. When pressed, this shortcut will toggle the dock's visibility once it's moved to the bottom panel (this shortcut does not affect the dock otherwise). See the default editor bottom panel shortcuts in the Editor Settings for inspiration. Per convention, they all use [kbd]Alt[/kbd] modifier.
*/
func (self Instance) AddControlToDock(slot gdclass.EditorPluginDockSlot, control [1]gdclass.Control) { //gd:EditorPlugin.add_control_to_dock
	Advanced(self).AddControlToDock(slot, control, [1][1]gdclass.Shortcut{}[0])
}

/*
Adds the control to a specific dock slot (see [enum DockSlot] for options).
If the dock is repositioned and as long as the plugin is active, the editor will save the dock position on further sessions.
When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_docks] and free it with [method Node.queue_free].
Optionally, you can specify a shortcut parameter. When pressed, this shortcut will toggle the dock's visibility once it's moved to the bottom panel (this shortcut does not affect the dock otherwise). See the default editor bottom panel shortcuts in the Editor Settings for inspiration. Per convention, they all use [kbd]Alt[/kbd] modifier.
*/
func (self Expanded) AddControlToDock(slot gdclass.EditorPluginDockSlot, control [1]gdclass.Control, shortcut [1]gdclass.Shortcut) { //gd:EditorPlugin.add_control_to_dock
	Advanced(self).AddControlToDock(slot, control, shortcut)
}

/*
Removes the control from the dock. You have to manually [method Node.queue_free] the control.
*/
func (self Instance) RemoveControlFromDocks(control [1]gdclass.Control) { //gd:EditorPlugin.remove_control_from_docks
	Advanced(self).RemoveControlFromDocks(control)
}

/*
Removes the control from the bottom panel. You have to manually [method Node.queue_free] the control.
*/
func (self Instance) RemoveControlFromBottomPanel(control [1]gdclass.Control) { //gd:EditorPlugin.remove_control_from_bottom_panel
	Advanced(self).RemoveControlFromBottomPanel(control)
}

/*
Removes the control from the specified container. You have to manually [method Node.queue_free] the control.
*/
func (self Instance) RemoveControlFromContainer(container gdclass.EditorPluginCustomControlContainer, control [1]gdclass.Control) { //gd:EditorPlugin.remove_control_from_container
	Advanced(self).RemoveControlFromContainer(container, control)
}

/*
Sets the tab icon for the given control in a dock slot. Setting to [code]null[/code] removes the icon.
*/
func (self Instance) SetDockTabIcon(control [1]gdclass.Control, icon [1]gdclass.Texture2D) { //gd:EditorPlugin.set_dock_tab_icon
	Advanced(self).SetDockTabIcon(control, icon)
}

/*
Adds a custom menu item to [b]Project > Tools[/b] named [param name]. When clicked, the provided [param callable] will be called.
*/
func (self Instance) AddToolMenuItem(name string, callable func()) { //gd:EditorPlugin.add_tool_menu_item
	Advanced(self).AddToolMenuItem(String.New(name), Callable.New(callable))
}

/*
Adds a custom [PopupMenu] submenu under [b]Project > Tools >[/b] [param name]. Use [method remove_tool_menu_item] on plugin clean up to remove the menu.
*/
func (self Instance) AddToolSubmenuItem(name string, submenu [1]gdclass.PopupMenu) { //gd:EditorPlugin.add_tool_submenu_item
	Advanced(self).AddToolSubmenuItem(String.New(name), submenu)
}

/*
Removes a menu [param name] from [b]Project > Tools[/b].
*/
func (self Instance) RemoveToolMenuItem(name string) { //gd:EditorPlugin.remove_tool_menu_item
	Advanced(self).RemoveToolMenuItem(String.New(name))
}

/*
Returns the [PopupMenu] under [b]Scene > Export As...[/b].
*/
func (self Instance) GetExportAsMenu() [1]gdclass.PopupMenu { //gd:EditorPlugin.get_export_as_menu
	return [1]gdclass.PopupMenu(Advanced(self).GetExportAsMenu())
}

/*
Adds a custom type, which will appear in the list of nodes or resources.
When a given node or resource is selected, the base type will be instantiated (e.g. "Node3D", "Control", "Resource"), then the script will be loaded and set to this object.
[b]Note:[/b] The base type is the base engine class which this type's class hierarchy inherits, not any custom type parent classes.
You can use the virtual method [method _handles] to check if your custom object is being edited by checking the script or using the [code]is[/code] keyword.
During run-time, this will be a simple object with a script so this function does not need to be called then.
[b]Note:[/b] Custom types added this way are not true classes. They are just a helper to create a node with specific script.
*/
func (self Instance) AddCustomType(atype string, base string, script [1]gdclass.Script, icon [1]gdclass.Texture2D) { //gd:EditorPlugin.add_custom_type
	Advanced(self).AddCustomType(String.New(atype), String.New(base), script, icon)
}

/*
Removes a custom type added by [method add_custom_type].
*/
func (self Instance) RemoveCustomType(atype string) { //gd:EditorPlugin.remove_custom_type
	Advanced(self).RemoveCustomType(String.New(atype))
}

/*
Adds a script at [param path] to the Autoload list as [param name].
*/
func (self Instance) AddAutoloadSingleton(name string, path string) { //gd:EditorPlugin.add_autoload_singleton
	Advanced(self).AddAutoloadSingleton(String.New(name), String.New(path))
}

/*
Removes an Autoload [param name] from the list.
*/
func (self Instance) RemoveAutoloadSingleton(name string) { //gd:EditorPlugin.remove_autoload_singleton
	Advanced(self).RemoveAutoloadSingleton(String.New(name))
}

/*
Updates the overlays of the 2D and 3D editor viewport. Causes methods [method _forward_canvas_draw_over_viewport], [method _forward_canvas_force_draw_over_viewport], [method _forward_3d_draw_over_viewport] and [method _forward_3d_force_draw_over_viewport] to be called.
*/
func (self Instance) UpdateOverlays() int { //gd:EditorPlugin.update_overlays
	return int(int(Advanced(self).UpdateOverlays()))
}

/*
Makes a specific item in the bottom panel visible.
*/
func (self Instance) MakeBottomPanelItemVisible(item [1]gdclass.Control) { //gd:EditorPlugin.make_bottom_panel_item_visible
	Advanced(self).MakeBottomPanelItemVisible(item)
}

/*
Minimizes the bottom panel.
*/
func (self Instance) HideBottomPanel() { //gd:EditorPlugin.hide_bottom_panel
	Advanced(self).HideBottomPanel()
}

/*
Gets the undo/redo object. Most actions in the editor can be undoable, so use this object to make sure this happens when it's worth it.
*/
func (self Instance) GetUndoRedo() [1]gdclass.EditorUndoRedoManager { //gd:EditorPlugin.get_undo_redo
	return [1]gdclass.EditorUndoRedoManager(Advanced(self).GetUndoRedo())
}

/*
Hooks a callback into the undo/redo action creation when a property is modified in the inspector. This allows, for example, to save other properties that may be lost when a given property is modified.
The callback should have 4 arguments: [Object] [code]undo_redo[/code], [Object] [code]modified_object[/code], [String] [code]property[/code] and [Variant] [code]new_value[/code]. They are, respectively, the [UndoRedo] object used by the inspector, the currently modified object, the name of the modified property and the new value the property is about to take.
*/
func (self Instance) AddUndoRedoInspectorHookCallback(callable func(undo_redo Object.Instance, modified_object Object.Instance, property string, new_value any)) { //gd:EditorPlugin.add_undo_redo_inspector_hook_callback
	Advanced(self).AddUndoRedoInspectorHookCallback(Callable.New(callable))
}

/*
Removes a callback previously added by [method add_undo_redo_inspector_hook_callback].
*/
func (self Instance) RemoveUndoRedoInspectorHookCallback(callable Callable.Function) { //gd:EditorPlugin.remove_undo_redo_inspector_hook_callback
	Advanced(self).RemoveUndoRedoInspectorHookCallback(Callable.New(callable))
}

/*
Queue save the project's editor layout.
*/
func (self Instance) QueueSaveLayout() { //gd:EditorPlugin.queue_save_layout
	Advanced(self).QueueSaveLayout()
}

/*
Registers a custom translation parser plugin for extracting translatable strings from custom files.
*/
func (self Instance) AddTranslationParserPlugin(parser [1]gdclass.EditorTranslationParserPlugin) { //gd:EditorPlugin.add_translation_parser_plugin
	Advanced(self).AddTranslationParserPlugin(parser)
}

/*
Removes a custom translation parser plugin registered by [method add_translation_parser_plugin].
*/
func (self Instance) RemoveTranslationParserPlugin(parser [1]gdclass.EditorTranslationParserPlugin) { //gd:EditorPlugin.remove_translation_parser_plugin
	Advanced(self).RemoveTranslationParserPlugin(parser)
}

/*
Registers a new [EditorImportPlugin]. Import plugins are used to import custom and unsupported assets as a custom [Resource] type.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
[b]Note:[/b] If you want to import custom 3D asset formats use [method add_scene_format_importer_plugin] instead.
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
func (self Instance) AddImportPlugin(importer [1]gdclass.EditorImportPlugin) { //gd:EditorPlugin.add_import_plugin
	Advanced(self).AddImportPlugin(importer, false)
}

/*
Registers a new [EditorImportPlugin]. Import plugins are used to import custom and unsupported assets as a custom [Resource] type.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
[b]Note:[/b] If you want to import custom 3D asset formats use [method add_scene_format_importer_plugin] instead.
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
func (self Expanded) AddImportPlugin(importer [1]gdclass.EditorImportPlugin, first_priority bool) { //gd:EditorPlugin.add_import_plugin
	Advanced(self).AddImportPlugin(importer, first_priority)
}

/*
Removes an import plugin registered by [method add_import_plugin].
*/
func (self Instance) RemoveImportPlugin(importer [1]gdclass.EditorImportPlugin) { //gd:EditorPlugin.remove_import_plugin
	Advanced(self).RemoveImportPlugin(importer)
}

/*
Registers a new [EditorSceneFormatImporter]. Scene importers are used to import custom 3D asset formats as scenes.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
*/
func (self Instance) AddSceneFormatImporterPlugin(scene_format_importer [1]gdclass.EditorSceneFormatImporter) { //gd:EditorPlugin.add_scene_format_importer_plugin
	Advanced(self).AddSceneFormatImporterPlugin(scene_format_importer, false)
}

/*
Registers a new [EditorSceneFormatImporter]. Scene importers are used to import custom 3D asset formats as scenes.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
*/
func (self Expanded) AddSceneFormatImporterPlugin(scene_format_importer [1]gdclass.EditorSceneFormatImporter, first_priority bool) { //gd:EditorPlugin.add_scene_format_importer_plugin
	Advanced(self).AddSceneFormatImporterPlugin(scene_format_importer, first_priority)
}

/*
Removes a scene format importer registered by [method add_scene_format_importer_plugin].
*/
func (self Instance) RemoveSceneFormatImporterPlugin(scene_format_importer [1]gdclass.EditorSceneFormatImporter) { //gd:EditorPlugin.remove_scene_format_importer_plugin
	Advanced(self).RemoveSceneFormatImporterPlugin(scene_format_importer)
}

/*
Add a [EditorScenePostImportPlugin]. These plugins allow customizing the import process of 3D assets by adding new options to the import dialogs.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
*/
func (self Instance) AddScenePostImportPlugin(scene_import_plugin [1]gdclass.EditorScenePostImportPlugin) { //gd:EditorPlugin.add_scene_post_import_plugin
	Advanced(self).AddScenePostImportPlugin(scene_import_plugin, false)
}

/*
Add a [EditorScenePostImportPlugin]. These plugins allow customizing the import process of 3D assets by adding new options to the import dialogs.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
*/
func (self Expanded) AddScenePostImportPlugin(scene_import_plugin [1]gdclass.EditorScenePostImportPlugin, first_priority bool) { //gd:EditorPlugin.add_scene_post_import_plugin
	Advanced(self).AddScenePostImportPlugin(scene_import_plugin, first_priority)
}

/*
Remove the [EditorScenePostImportPlugin], added with [method add_scene_post_import_plugin].
*/
func (self Instance) RemoveScenePostImportPlugin(scene_import_plugin [1]gdclass.EditorScenePostImportPlugin) { //gd:EditorPlugin.remove_scene_post_import_plugin
	Advanced(self).RemoveScenePostImportPlugin(scene_import_plugin)
}

/*
Registers a new [EditorExportPlugin]. Export plugins are used to perform tasks when the project is being exported.
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
func (self Instance) AddExportPlugin(plugin [1]gdclass.EditorExportPlugin) { //gd:EditorPlugin.add_export_plugin
	Advanced(self).AddExportPlugin(plugin)
}

/*
Removes an export plugin registered by [method add_export_plugin].
*/
func (self Instance) RemoveExportPlugin(plugin [1]gdclass.EditorExportPlugin) { //gd:EditorPlugin.remove_export_plugin
	Advanced(self).RemoveExportPlugin(plugin)
}

/*
Registers a new [EditorExportPlatform]. Export platforms provides functionality of exporting to the specific platform.
*/
func (self Instance) AddExportPlatform(platform [1]gdclass.EditorExportPlatform) { //gd:EditorPlugin.add_export_platform
	Advanced(self).AddExportPlatform(platform)
}

/*
Removes an export platform registered by [method add_export_platform].
*/
func (self Instance) RemoveExportPlatform(platform [1]gdclass.EditorExportPlatform) { //gd:EditorPlugin.remove_export_platform
	Advanced(self).RemoveExportPlatform(platform)
}

/*
Registers a new [EditorNode3DGizmoPlugin]. Gizmo plugins are used to add custom gizmos to the 3D preview viewport for a [Node3D].
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
func (self Instance) AddNode3dGizmoPlugin(plugin [1]gdclass.EditorNode3DGizmoPlugin) { //gd:EditorPlugin.add_node_3d_gizmo_plugin
	Advanced(self).AddNode3dGizmoPlugin(plugin)
}

/*
Removes a gizmo plugin registered by [method add_node_3d_gizmo_plugin].
*/
func (self Instance) RemoveNode3dGizmoPlugin(plugin [1]gdclass.EditorNode3DGizmoPlugin) { //gd:EditorPlugin.remove_node_3d_gizmo_plugin
	Advanced(self).RemoveNode3dGizmoPlugin(plugin)
}

/*
Registers a new [EditorInspectorPlugin]. Inspector plugins are used to extend [EditorInspector] and provide custom configuration tools for your object's properties.
[b]Note:[/b] Always use [method remove_inspector_plugin] to remove the registered [EditorInspectorPlugin] when your [EditorPlugin] is disabled to prevent leaks and an unexpected behavior.
[codeblocks]
[gdscript]
const MyInspectorPlugin = preload("res://addons/your_addon/path/to/your/script.gd")
var inspector_plugin = MyInspectorPlugin.new()

func _enter_tree():

	add_inspector_plugin(inspector_plugin)

func _exit_tree():

	remove_inspector_plugin(inspector_plugin)

[/gdscript]
[/codeblocks]
*/
func (self Instance) AddInspectorPlugin(plugin [1]gdclass.EditorInspectorPlugin) { //gd:EditorPlugin.add_inspector_plugin
	Advanced(self).AddInspectorPlugin(plugin)
}

/*
Removes an inspector plugin registered by [method add_inspector_plugin].
*/
func (self Instance) RemoveInspectorPlugin(plugin [1]gdclass.EditorInspectorPlugin) { //gd:EditorPlugin.remove_inspector_plugin
	Advanced(self).RemoveInspectorPlugin(plugin)
}

/*
Registers a new [EditorResourceConversionPlugin]. Resource conversion plugins are used to add custom resource converters to the editor inspector.
See [EditorResourceConversionPlugin] for an example of how to create a resource conversion plugin.
*/
func (self Instance) AddResourceConversionPlugin(plugin [1]gdclass.EditorResourceConversionPlugin) { //gd:EditorPlugin.add_resource_conversion_plugin
	Advanced(self).AddResourceConversionPlugin(plugin)
}

/*
Removes a resource conversion plugin registered by [method add_resource_conversion_plugin].
*/
func (self Instance) RemoveResourceConversionPlugin(plugin [1]gdclass.EditorResourceConversionPlugin) { //gd:EditorPlugin.remove_resource_conversion_plugin
	Advanced(self).RemoveResourceConversionPlugin(plugin)
}

/*
Use this method if you always want to receive inputs from 3D view screen inside [method _forward_3d_gui_input]. It might be especially usable if your plugin will want to use raycast in the scene.
*/
func (self Instance) SetInputEventForwardingAlwaysEnabled() { //gd:EditorPlugin.set_input_event_forwarding_always_enabled
	Advanced(self).SetInputEventForwardingAlwaysEnabled()
}

/*
Enables calling of [method _forward_canvas_force_draw_over_viewport] for the 2D editor and [method _forward_3d_force_draw_over_viewport] for the 3D editor when their viewports are updated. You need to call this method only once and it will work permanently for this plugin.
*/
func (self Instance) SetForceDrawOverForwardingEnabled() { //gd:EditorPlugin.set_force_draw_over_forwarding_enabled
	Advanced(self).SetForceDrawOverForwardingEnabled()
}

/*
Adds a plugin to the context menu. [param slot] is the context menu where the plugin will be added.
See [enum EditorContextMenuPlugin.ContextMenuSlot] for available context menus. A plugin instance can belong only to a single context menu slot.
*/
func (self Instance) AddContextMenuPlugin(slot gdclass.EditorContextMenuPluginContextMenuSlot, plugin [1]gdclass.EditorContextMenuPlugin) { //gd:EditorPlugin.add_context_menu_plugin
	Advanced(self).AddContextMenuPlugin(slot, plugin)
}

/*
Removes the specified context menu plugin.
*/
func (self Instance) RemoveContextMenuPlugin(plugin [1]gdclass.EditorContextMenuPlugin) { //gd:EditorPlugin.remove_context_menu_plugin
	Advanced(self).RemoveContextMenuPlugin(plugin)
}

/*
Returns the [EditorInterface] singleton instance.
*/
func (self Instance) GetEditorInterface() [1]gdclass.EditorInterface { //gd:EditorPlugin.get_editor_interface
	return [1]gdclass.EditorInterface(Advanced(self).GetEditorInterface())
}

/*
Gets the Editor's dialog used for making scripts.
[b]Note:[/b] Users can configure it before use.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func (self Instance) GetScriptCreateDialog() [1]gdclass.ScriptCreateDialog { //gd:EditorPlugin.get_script_create_dialog
	return [1]gdclass.ScriptCreateDialog(Advanced(self).GetScriptCreateDialog())
}

/*
Adds a [Script] as debugger plugin to the Debugger. The script must extend [EditorDebuggerPlugin].
*/
func (self Instance) AddDebuggerPlugin(script [1]gdclass.EditorDebuggerPlugin) { //gd:EditorPlugin.add_debugger_plugin
	Advanced(self).AddDebuggerPlugin(script)
}

/*
Removes the debugger plugin with given script from the Debugger.
*/
func (self Instance) RemoveDebuggerPlugin(script [1]gdclass.EditorDebuggerPlugin) { //gd:EditorPlugin.remove_debugger_plugin
	Advanced(self).RemoveDebuggerPlugin(script)
}

/*
Provide the version of the plugin declared in the [code]plugin.cfg[/code] config file.
*/
func (self Instance) GetPluginVersion() string { //gd:EditorPlugin.get_plugin_version
	return string(Advanced(self).GetPluginVersion().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorPlugin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorPlugin"))
	casted := Instance{*(*gdclass.EditorPlugin)(unsafe.Pointer(&object))}
	return casted
}

/*
Called when there is a root node in the current edited scene, [method _handles] is implemented, and an [InputEvent] happens in the 2D viewport. If this method returns [code]true[/code], [param event] is intercepted by this [EditorPlugin], otherwise [param event] is forwarded to other Editor classes.
[codeblocks]
[gdscript]
# Prevents the InputEvent from reaching other Editor classes.
func _forward_canvas_gui_input(event):

	return true

[/gdscript]
[csharp]
// Prevents the InputEvent from reaching other Editor classes.
public override bool ForwardCanvasGuiInput(InputEvent @event)

	{
	    return true;
	}

[/csharp]
[/codeblocks]
This method must return [code]false[/code] in order to forward the [InputEvent] to other Editor classes.
[codeblocks]
[gdscript]
# Consumes InputEventMouseMotion and forwards other InputEvent types.
func _forward_canvas_gui_input(event):

	if (event is InputEventMouseMotion):
	    return true
	return false

[/gdscript]
[csharp]
// Consumes InputEventMouseMotion and forwards other InputEvent types.
public override bool _ForwardCanvasGuiInput(InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        return true;
	    }
	    return false;
	}

[/csharp]
[/codeblocks]
*/
func (class) _forward_canvas_gui_input(impl func(ptr unsafe.Pointer, event [1]gdclass.InputEvent) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called by the engine when the 2D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
[codeblocks]
[gdscript]
func _forward_canvas_draw_over_viewport(overlay):

	# Draw a circle at cursor position.
	overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)

func _forward_canvas_gui_input(event):

	if event is InputEventMouseMotion:
	    # Redraw viewport when cursor is moved.
	    update_overlays()
	    return true
	return false

[/gdscript]
[csharp]
public override void _ForwardCanvasDrawOverViewport(Control viewportControl)

	{
	    // Draw a circle at cursor position.
	    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
	}

public override bool _ForwardCanvasGuiInput(InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        // Redraw viewport when cursor is moved.
	        UpdateOverlays();
	        return true;
	    }
	    return false;
	}

[/csharp]
[/codeblocks]
*/
func (class) _forward_canvas_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control [1]gdclass.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport_control = [1]gdclass.Control{pointers.New[gdclass.Control]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
This method is the same as [method _forward_canvas_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
*/
func (class) _forward_canvas_force_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control [1]gdclass.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport_control = [1]gdclass.Control{pointers.New[gdclass.Control]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
Called when there is a root node in the current edited scene, [method _handles] is implemented, and an [InputEvent] happens in the 3D viewport. The return value decides whether the [InputEvent] is consumed or forwarded to other [EditorPlugin]s. See [enum AfterGUIInput] for options.
[codeblocks]
[gdscript]
# Prevents the InputEvent from reaching other Editor classes.
func _forward_3d_gui_input(camera, event):

	return EditorPlugin.AFTER_GUI_INPUT_STOP

[/gdscript]
[csharp]
// Prevents the InputEvent from reaching other Editor classes.
public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)

	{
	    return EditorPlugin.AfterGuiInput.Stop;
	}

[/csharp]
[/codeblocks]
This method must return [constant AFTER_GUI_INPUT_PASS] in order to forward the [InputEvent] to other Editor classes.
[codeblocks]
[gdscript]
# Consumes InputEventMouseMotion and forwards other InputEvent types.
func _forward_3d_gui_input(camera, event):

	return EditorPlugin.AFTER_GUI_INPUT_STOP if event is InputEventMouseMotion else EditorPlugin.AFTER_GUI_INPUT_PASS

[/gdscript]
[csharp]
// Consumes InputEventMouseMotion and forwards other InputEvent types.
public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)

	{
	    return @event is InputEventMouseMotion ? EditorPlugin.AfterGuiInput.Stop : EditorPlugin.AfterGuiInput.Pass;
	}

[/csharp]
[/codeblocks]
*/
func (class) _forward_3d_gui_input(impl func(ptr unsafe.Pointer, viewport_camera [1]gdclass.Camera3D, event [1]gdclass.InputEvent) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport_camera = [1]gdclass.Camera3D{pointers.New[gdclass.Camera3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport_camera[0])
		var event = [1]gdclass.InputEvent{pointers.New[gdclass.InputEvent]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, viewport_camera, event)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called by the engine when the 3D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
[codeblocks]
[gdscript]
func _forward_3d_draw_over_viewport(overlay):

	# Draw a circle at cursor position.
	overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)

func _forward_3d_gui_input(camera, event):

	if event is InputEventMouseMotion:
	    # Redraw viewport when cursor is moved.
	    update_overlays()
	    return EditorPlugin.AFTER_GUI_INPUT_STOP
	return EditorPlugin.AFTER_GUI_INPUT_PASS

[/gdscript]
[csharp]
public override void _Forward3DDrawOverViewport(Control viewportControl)

	{
	    // Draw a circle at cursor position.
	    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
	}

public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D viewportCamera, InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        // Redraw viewport when cursor is moved.
	        UpdateOverlays();
	        return EditorPlugin.AfterGuiInput.Stop;
	    }
	    return EditorPlugin.AfterGuiInput.Pass;
	}

[/csharp]
[/codeblocks]
*/
func (class) _forward_3d_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control [1]gdclass.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport_control = [1]gdclass.Control{pointers.New[gdclass.Control]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
This method is the same as [method _forward_3d_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
*/
func (class) _forward_3d_force_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control [1]gdclass.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var viewport_control = [1]gdclass.Control{pointers.New[gdclass.Control]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
Override this method in your plugin to provide the name of the plugin when displayed in the Godot editor.
For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
*/
func (class) _get_plugin_name(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method in your plugin to return a [Texture2D] in order to give it an icon.
For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
Ideally, the plugin icon should be white with a transparent background and 16×16 pixels in size.
[codeblocks]
[gdscript]
func _get_plugin_icon():

	# You can use a custom icon:
	return preload("res://addons/my_plugin/my_plugin_icon.svg")
	# Or use a built-in icon:
	return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")

[/gdscript]
[csharp]
public override Texture2D _GetPluginIcon()

	{
	    // You can use a custom icon:
	    return ResourceLoader.Load<Texture2D>("res://addons/my_plugin/my_plugin_icon.svg");
	    // Or use a built-in icon:
	    return EditorInterface.Singleton.GetEditorTheme().GetIcon("Node", "EditorIcons");
	}

[/csharp]
[/codeblocks]
*/
func (class) _get_plugin_icon(impl func(ptr unsafe.Pointer) [1]gdclass.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this is a main screen editor plugin (it goes in the workspace selector together with [b]2D[/b], [b]3D[/b], [b]Script[/b] and [b]AssetLib[/b]).
When the plugin's workspace is selected, other main screen plugins will be hidden, but your plugin will not appear automatically. It needs to be added as a child of [method EditorInterface.get_editor_main_screen] and made visible inside [method _make_visible].
Use [method _get_plugin_name] and [method _get_plugin_icon] to customize the plugin button's appearance.
[codeblock]
var plugin_control

func _enter_tree():

	plugin_control = preload("my_plugin_control.tscn").instantiate()
	EditorInterface.get_editor_main_screen().add_child(plugin_control)
	plugin_control.hide()

func _has_main_screen():

	return true

func _make_visible(visible):

	plugin_control.visible = visible

func _get_plugin_name():

	return "My Super Cool Plugin 3000"

func _get_plugin_icon():

	return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")

[/codeblock]
*/
func (class) _has_main_screen(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
This function will be called when the editor is requested to become visible. It is used for plugins that edit a specific object type.
Remember that you have to manage the visibility of all your editor controls manually.
*/
func (class) _make_visible(impl func(ptr unsafe.Pointer, visible bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var visible = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, visible)
	}
}

/*
This function is used for plugins that edit specific object types (nodes or resources). It requests the editor to edit the given object.
[param object] can be [code]null[/code] if the plugin was editing an object, but there is no longer any selected object handled by this plugin. It can be used to cleanup editing state.
*/
func (class) _edit(impl func(ptr unsafe.Pointer, obj [1]gd.Object)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj)
	}
}

/*
Implement this function if your plugin edits a specific type of object (Resource or Node). If you return [code]true[/code], then you will get the functions [method _edit] and [method _make_visible] called when the editor requests them. If you have declared the methods [method _forward_canvas_gui_input] and [method _forward_3d_gui_input] these will be called too.
[b]Note:[/b] Each plugin should handle only one type of objects at a time. If a plugin handles more types of objects and they are edited at the same time, it will result in errors.
*/
func (class) _handles(impl func(ptr unsafe.Pointer, obj [1]gd.Object) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to provide a state data you want to be saved, like view position, grid settings, folding, etc. This is used when saving the scene (so state is kept when opening it again) and for switching tabs (so state can be restored when the tab returns). This data is automatically saved for each scene in an [code]editstate[/code] file in the editor metadata folder. If you want to store global (scene-independent) editor data for your plugin, you can use [method _get_window_layout] instead.
Use [method _set_state] to restore your saved state.
[b]Note:[/b] This method should not be used to save important settings that should persist with the project.
[b]Note:[/b] You must implement [method _get_plugin_name] for the state to be stored and restored correctly.
[codeblock]
func _get_state():

	var state = {"zoom": zoom, "preferred_color": my_color}
	return state

[/codeblock]
*/
func (class) _get_state(impl func(ptr unsafe.Pointer) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Restore the state saved by [method _get_state]. This method is called when the current scene tab is changed in the editor.
[b]Note:[/b] Your plugin must implement [method _get_plugin_name], otherwise it will not be recognized and this method will not be called.
[codeblock]
func _set_state(data):

	zoom = data.get("zoom", 1.0)
	preferred_color = data.get("my_color", Color.WHITE)

[/codeblock]
*/
func (class) _set_state(impl func(ptr unsafe.Pointer, state Dictionary.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var state = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalDictionary(state))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state)
	}
}

/*
Clear all the state and reset the object being edited to zero. This ensures your plugin does not keep editing a currently existing node, or a node from the wrong scene.
*/
func (class) _clear(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Override this method to provide a custom message that lists unsaved changes. The editor will call this method when exiting or when closing a scene, and display the returned string in a confirmation dialog. Return empty string if the plugin has no unsaved changes.
When closing a scene, [param for_scene] is the path to the scene being closed. You can use it to handle built-in resources in that scene.
If the user confirms saving, [method _save_external_data] will be called, before closing the editor.
[codeblock]
func _get_unsaved_status(for_scene):

	if not unsaved:
	    return ""

	if for_scene.is_empty():
	    return "Save changes in MyCustomPlugin before closing?"
	else:
	    return "Scene %s has changes from MyCustomPlugin. Save before closing?" % for_scene.get_file()

func _save_external_data():

	unsaved = false

[/codeblock]
If the plugin has no scene-specific changes, you can ignore the calls when closing scenes:
[codeblock]
func _get_unsaved_status(for_scene):

	if not for_scene.is_empty():
	    return ""

[/codeblock]
*/
func (class) _get_unsaved_status(impl func(ptr unsafe.Pointer, for_scene String.Readable) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var for_scene = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(for_scene))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_scene)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
This method is called after the editor saves the project or when it's closed. It asks the plugin to save edited external scenes/resources.
*/
func (class) _save_external_data(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
This method is called when the editor is about to save the project, switch to another tab, etc. It asks the plugin to apply any pending state changes to ensure consistency.
This is used, for example, in shader editors to let the plugin know that it must apply the shader code being written by the user to the object.
*/
func (class) _apply_changes(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
This is for editors that edit script-based objects. You can return a list of breakpoints in the format ([code]script:line[/code]), for example: [code]res://path_to_script.gd:25[/code].
*/
func (class) _get_breakpoints(impl func(ptr unsafe.Pointer) Packed.Strings) (cb gd.ExtensionClassCallVirtualFunc) {
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
Restore the plugin GUI layout and data saved by [method _get_window_layout]. This method is called for every plugin on editor startup. Use the provided [param configuration] file to read your saved data.
[codeblock]
func _set_window_layout(configuration):

	$Window.position = configuration.get_value("MyPlugin", "window_position", Vector2())
	$Icon.modulate = configuration.get_value("MyPlugin", "icon_color", Color.WHITE)

[/codeblock]
*/
func (class) _set_window_layout(impl func(ptr unsafe.Pointer, configuration [1]gdclass.ConfigFile)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var configuration = [1]gdclass.ConfigFile{pointers.New[gdclass.ConfigFile]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(configuration[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, configuration)
	}
}

/*
Override this method to provide the GUI layout of the plugin or any other data you want to be stored. This is used to save the project's editor layout when [method queue_save_layout] is called or the editor layout was changed (for example changing the position of a dock). The data is stored in the [code]editor_layout.cfg[/code] file in the editor metadata directory.
Use [method _set_window_layout] to restore your saved layout.
[codeblock]
func _get_window_layout(configuration):

	configuration.set_value("MyPlugin", "window_position", $Window.position)
	configuration.set_value("MyPlugin", "icon_color", $Icon.modulate)

[/codeblock]
*/
func (class) _get_window_layout(impl func(ptr unsafe.Pointer, configuration [1]gdclass.ConfigFile)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var configuration = [1]gdclass.ConfigFile{pointers.New[gdclass.ConfigFile]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(configuration[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, configuration)
	}
}

/*
This method is called when the editor is about to run the project. The plugin can then perform required operations before the project runs.
This method must return a boolean. If this method returns [code]false[/code], the project will not run. The run is aborted immediately, so this also prevents all other plugins' [method _build] methods from running.
*/
func (class) _build(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called by the engine when the user enables the [EditorPlugin] in the Plugin tab of the project settings window.
*/
func (class) _enable_plugin(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called by the engine when the user disables the [EditorPlugin] in the Plugin tab of the project settings window.
*/
func (class) _disable_plugin(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Adds a custom control to a container (see [enum CustomControlContainer]). There are many locations where custom controls can be added in the editor UI.
Please remember that you have to manage the visibility of your custom controls yourself (and likely hide it after adding it).
When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_container] and free it with [method Node.queue_free].
*/
//go:nosplit
func (self class) AddControlToContainer(container gdclass.EditorPluginCustomControlContainer, control [1]gdclass.Control) { //gd:EditorPlugin.add_control_to_container
	var frame = callframe.New()
	callframe.Arg(frame, container)
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(control[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_control_to_container, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a control to the bottom panel (together with Output, Debug, Animation, etc.). Returns a reference to the button added. It's up to you to hide/show the button when needed. When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_bottom_panel] and free it with [method Node.queue_free].
Optionally, you can specify a shortcut parameter. When pressed, this shortcut will toggle the bottom panel's visibility. See the default editor bottom panel shortcuts in the Editor Settings for inspiration. Per convention, they all use [kbd]Alt[/kbd] modifier.
*/
//go:nosplit
func (self class) AddControlToBottomPanel(control [1]gdclass.Control, title String.Readable, shortcut [1]gdclass.Shortcut) [1]gdclass.Button { //gd:EditorPlugin.add_control_to_bottom_panel
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(control[0].AsObject()[0]))
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_control_to_bottom_panel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Button{gd.PointerMustAssertInstanceID[gdclass.Button](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adds the control to a specific dock slot (see [enum DockSlot] for options).
If the dock is repositioned and as long as the plugin is active, the editor will save the dock position on further sessions.
When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_docks] and free it with [method Node.queue_free].
Optionally, you can specify a shortcut parameter. When pressed, this shortcut will toggle the dock's visibility once it's moved to the bottom panel (this shortcut does not affect the dock otherwise). See the default editor bottom panel shortcuts in the Editor Settings for inspiration. Per convention, they all use [kbd]Alt[/kbd] modifier.
*/
//go:nosplit
func (self class) AddControlToDock(slot gdclass.EditorPluginDockSlot, control [1]gdclass.Control, shortcut [1]gdclass.Shortcut) { //gd:EditorPlugin.add_control_to_dock
	var frame = callframe.New()
	callframe.Arg(frame, slot)
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(control[0].AsObject()[0]))
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_control_to_dock, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the control from the dock. You have to manually [method Node.queue_free] the control.
*/
//go:nosplit
func (self class) RemoveControlFromDocks(control [1]gdclass.Control) { //gd:EditorPlugin.remove_control_from_docks
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_control_from_docks, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the control from the bottom panel. You have to manually [method Node.queue_free] the control.
*/
//go:nosplit
func (self class) RemoveControlFromBottomPanel(control [1]gdclass.Control) { //gd:EditorPlugin.remove_control_from_bottom_panel
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_control_from_bottom_panel, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the control from the specified container. You have to manually [method Node.queue_free] the control.
*/
//go:nosplit
func (self class) RemoveControlFromContainer(container gdclass.EditorPluginCustomControlContainer, control [1]gdclass.Control) { //gd:EditorPlugin.remove_control_from_container
	var frame = callframe.New()
	callframe.Arg(frame, container)
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_control_from_container, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the tab icon for the given control in a dock slot. Setting to [code]null[/code] removes the icon.
*/
//go:nosplit
func (self class) SetDockTabIcon(control [1]gdclass.Control, icon [1]gdclass.Texture2D) { //gd:EditorPlugin.set_dock_tab_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_set_dock_tab_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a custom menu item to [b]Project > Tools[/b] named [param name]. When clicked, the provided [param callable] will be called.
*/
//go:nosplit
func (self class) AddToolMenuItem(name String.Readable, callable Callable.Function) { //gd:EditorPlugin.add_tool_menu_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callable)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_tool_menu_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a custom [PopupMenu] submenu under [b]Project > Tools >[/b] [param name]. Use [method remove_tool_menu_item] on plugin clean up to remove the menu.
*/
//go:nosplit
func (self class) AddToolSubmenuItem(name String.Readable, submenu [1]gdclass.PopupMenu) { //gd:EditorPlugin.add_tool_submenu_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(submenu[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_tool_submenu_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a menu [param name] from [b]Project > Tools[/b].
*/
//go:nosplit
func (self class) RemoveToolMenuItem(name String.Readable) { //gd:EditorPlugin.remove_tool_menu_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_tool_menu_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [PopupMenu] under [b]Scene > Export As...[/b].
*/
//go:nosplit
func (self class) GetExportAsMenu() [1]gdclass.PopupMenu { //gd:EditorPlugin.get_export_as_menu
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_get_export_as_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PopupMenu{gd.PointerLifetimeBoundTo[gdclass.PopupMenu](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adds a custom type, which will appear in the list of nodes or resources.
When a given node or resource is selected, the base type will be instantiated (e.g. "Node3D", "Control", "Resource"), then the script will be loaded and set to this object.
[b]Note:[/b] The base type is the base engine class which this type's class hierarchy inherits, not any custom type parent classes.
You can use the virtual method [method _handles] to check if your custom object is being edited by checking the script or using the [code]is[/code] keyword.
During run-time, this will be a simple object with a script so this function does not need to be called then.
[b]Note:[/b] Custom types added this way are not true classes. They are just a helper to create a node with specific script.
*/
//go:nosplit
func (self class) AddCustomType(atype String.Readable, base String.Readable, script [1]gdclass.Script, icon [1]gdclass.Texture2D) { //gd:EditorPlugin.add_custom_type
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(atype)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(base)))
	callframe.Arg(frame, pointers.Get(script[0])[0])
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_custom_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a custom type added by [method add_custom_type].
*/
//go:nosplit
func (self class) RemoveCustomType(atype String.Readable) { //gd:EditorPlugin.remove_custom_type
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(atype)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_custom_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a script at [param path] to the Autoload list as [param name].
*/
//go:nosplit
func (self class) AddAutoloadSingleton(name String.Readable, path String.Readable) { //gd:EditorPlugin.add_autoload_singleton
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_autoload_singleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes an Autoload [param name] from the list.
*/
//go:nosplit
func (self class) RemoveAutoloadSingleton(name String.Readable) { //gd:EditorPlugin.remove_autoload_singleton
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_autoload_singleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Updates the overlays of the 2D and 3D editor viewport. Causes methods [method _forward_canvas_draw_over_viewport], [method _forward_canvas_force_draw_over_viewport], [method _forward_3d_draw_over_viewport] and [method _forward_3d_force_draw_over_viewport] to be called.
*/
//go:nosplit
func (self class) UpdateOverlays() int64 { //gd:EditorPlugin.update_overlays
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_update_overlays, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Makes a specific item in the bottom panel visible.
*/
//go:nosplit
func (self class) MakeBottomPanelItemVisible(item [1]gdclass.Control) { //gd:EditorPlugin.make_bottom_panel_item_visible
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(item[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_make_bottom_panel_item_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Minimizes the bottom panel.
*/
//go:nosplit
func (self class) HideBottomPanel() { //gd:EditorPlugin.hide_bottom_panel
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_hide_bottom_panel, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets the undo/redo object. Most actions in the editor can be undoable, so use this object to make sure this happens when it's worth it.
*/
//go:nosplit
func (self class) GetUndoRedo() [1]gdclass.EditorUndoRedoManager { //gd:EditorPlugin.get_undo_redo
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_get_undo_redo, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorUndoRedoManager{gd.PointerLifetimeBoundTo[gdclass.EditorUndoRedoManager](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Hooks a callback into the undo/redo action creation when a property is modified in the inspector. This allows, for example, to save other properties that may be lost when a given property is modified.
The callback should have 4 arguments: [Object] [code]undo_redo[/code], [Object] [code]modified_object[/code], [String] [code]property[/code] and [Variant] [code]new_value[/code]. They are, respectively, the [UndoRedo] object used by the inspector, the currently modified object, the name of the modified property and the new value the property is about to take.
*/
//go:nosplit
func (self class) AddUndoRedoInspectorHookCallback(callable Callable.Function) { //gd:EditorPlugin.add_undo_redo_inspector_hook_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callable)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_undo_redo_inspector_hook_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a callback previously added by [method add_undo_redo_inspector_hook_callback].
*/
//go:nosplit
func (self class) RemoveUndoRedoInspectorHookCallback(callable Callable.Function) { //gd:EditorPlugin.remove_undo_redo_inspector_hook_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callable)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_undo_redo_inspector_hook_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Queue save the project's editor layout.
*/
//go:nosplit
func (self class) QueueSaveLayout() { //gd:EditorPlugin.queue_save_layout
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_queue_save_layout, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers a custom translation parser plugin for extracting translatable strings from custom files.
*/
//go:nosplit
func (self class) AddTranslationParserPlugin(parser [1]gdclass.EditorTranslationParserPlugin) { //gd:EditorPlugin.add_translation_parser_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parser[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_translation_parser_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a custom translation parser plugin registered by [method add_translation_parser_plugin].
*/
//go:nosplit
func (self class) RemoveTranslationParserPlugin(parser [1]gdclass.EditorTranslationParserPlugin) { //gd:EditorPlugin.remove_translation_parser_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parser[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_translation_parser_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers a new [EditorImportPlugin]. Import plugins are used to import custom and unsupported assets as a custom [Resource] type.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
[b]Note:[/b] If you want to import custom 3D asset formats use [method add_scene_format_importer_plugin] instead.
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
//go:nosplit
func (self class) AddImportPlugin(importer [1]gdclass.EditorImportPlugin, first_priority bool) { //gd:EditorPlugin.add_import_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(importer[0])[0])
	callframe.Arg(frame, first_priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_import_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes an import plugin registered by [method add_import_plugin].
*/
//go:nosplit
func (self class) RemoveImportPlugin(importer [1]gdclass.EditorImportPlugin) { //gd:EditorPlugin.remove_import_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(importer[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_import_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers a new [EditorSceneFormatImporter]. Scene importers are used to import custom 3D asset formats as scenes.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
*/
//go:nosplit
func (self class) AddSceneFormatImporterPlugin(scene_format_importer [1]gdclass.EditorSceneFormatImporter, first_priority bool) { //gd:EditorPlugin.add_scene_format_importer_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_format_importer[0])[0])
	callframe.Arg(frame, first_priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_scene_format_importer_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a scene format importer registered by [method add_scene_format_importer_plugin].
*/
//go:nosplit
func (self class) RemoveSceneFormatImporterPlugin(scene_format_importer [1]gdclass.EditorSceneFormatImporter) { //gd:EditorPlugin.remove_scene_format_importer_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_format_importer[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_scene_format_importer_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Add a [EditorScenePostImportPlugin]. These plugins allow customizing the import process of 3D assets by adding new options to the import dialogs.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
*/
//go:nosplit
func (self class) AddScenePostImportPlugin(scene_import_plugin [1]gdclass.EditorScenePostImportPlugin, first_priority bool) { //gd:EditorPlugin.add_scene_post_import_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_import_plugin[0])[0])
	callframe.Arg(frame, first_priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_scene_post_import_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Remove the [EditorScenePostImportPlugin], added with [method add_scene_post_import_plugin].
*/
//go:nosplit
func (self class) RemoveScenePostImportPlugin(scene_import_plugin [1]gdclass.EditorScenePostImportPlugin) { //gd:EditorPlugin.remove_scene_post_import_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_import_plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_scene_post_import_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers a new [EditorExportPlugin]. Export plugins are used to perform tasks when the project is being exported.
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
//go:nosplit
func (self class) AddExportPlugin(plugin [1]gdclass.EditorExportPlugin) { //gd:EditorPlugin.add_export_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_export_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes an export plugin registered by [method add_export_plugin].
*/
//go:nosplit
func (self class) RemoveExportPlugin(plugin [1]gdclass.EditorExportPlugin) { //gd:EditorPlugin.remove_export_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_export_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers a new [EditorExportPlatform]. Export platforms provides functionality of exporting to the specific platform.
*/
//go:nosplit
func (self class) AddExportPlatform(platform [1]gdclass.EditorExportPlatform) { //gd:EditorPlugin.add_export_platform
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(platform[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_export_platform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes an export platform registered by [method add_export_platform].
*/
//go:nosplit
func (self class) RemoveExportPlatform(platform [1]gdclass.EditorExportPlatform) { //gd:EditorPlugin.remove_export_platform
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(platform[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_export_platform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers a new [EditorNode3DGizmoPlugin]. Gizmo plugins are used to add custom gizmos to the 3D preview viewport for a [Node3D].
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
//go:nosplit
func (self class) AddNode3dGizmoPlugin(plugin [1]gdclass.EditorNode3DGizmoPlugin) { //gd:EditorPlugin.add_node_3d_gizmo_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_node_3d_gizmo_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a gizmo plugin registered by [method add_node_3d_gizmo_plugin].
*/
//go:nosplit
func (self class) RemoveNode3dGizmoPlugin(plugin [1]gdclass.EditorNode3DGizmoPlugin) { //gd:EditorPlugin.remove_node_3d_gizmo_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_node_3d_gizmo_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers a new [EditorInspectorPlugin]. Inspector plugins are used to extend [EditorInspector] and provide custom configuration tools for your object's properties.
[b]Note:[/b] Always use [method remove_inspector_plugin] to remove the registered [EditorInspectorPlugin] when your [EditorPlugin] is disabled to prevent leaks and an unexpected behavior.
[codeblocks]
[gdscript]
const MyInspectorPlugin = preload("res://addons/your_addon/path/to/your/script.gd")
var inspector_plugin = MyInspectorPlugin.new()

func _enter_tree():
    add_inspector_plugin(inspector_plugin)

func _exit_tree():
    remove_inspector_plugin(inspector_plugin)
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) AddInspectorPlugin(plugin [1]gdclass.EditorInspectorPlugin) { //gd:EditorPlugin.add_inspector_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_inspector_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes an inspector plugin registered by [method add_inspector_plugin].
*/
//go:nosplit
func (self class) RemoveInspectorPlugin(plugin [1]gdclass.EditorInspectorPlugin) { //gd:EditorPlugin.remove_inspector_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_inspector_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers a new [EditorResourceConversionPlugin]. Resource conversion plugins are used to add custom resource converters to the editor inspector.
See [EditorResourceConversionPlugin] for an example of how to create a resource conversion plugin.
*/
//go:nosplit
func (self class) AddResourceConversionPlugin(plugin [1]gdclass.EditorResourceConversionPlugin) { //gd:EditorPlugin.add_resource_conversion_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_resource_conversion_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a resource conversion plugin registered by [method add_resource_conversion_plugin].
*/
//go:nosplit
func (self class) RemoveResourceConversionPlugin(plugin [1]gdclass.EditorResourceConversionPlugin) { //gd:EditorPlugin.remove_resource_conversion_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_resource_conversion_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Use this method if you always want to receive inputs from 3D view screen inside [method _forward_3d_gui_input]. It might be especially usable if your plugin will want to use raycast in the scene.
*/
//go:nosplit
func (self class) SetInputEventForwardingAlwaysEnabled() { //gd:EditorPlugin.set_input_event_forwarding_always_enabled
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_set_input_event_forwarding_always_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Enables calling of [method _forward_canvas_force_draw_over_viewport] for the 2D editor and [method _forward_3d_force_draw_over_viewport] for the 3D editor when their viewports are updated. You need to call this method only once and it will work permanently for this plugin.
*/
//go:nosplit
func (self class) SetForceDrawOverForwardingEnabled() { //gd:EditorPlugin.set_force_draw_over_forwarding_enabled
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_set_force_draw_over_forwarding_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a plugin to the context menu. [param slot] is the context menu where the plugin will be added.
See [enum EditorContextMenuPlugin.ContextMenuSlot] for available context menus. A plugin instance can belong only to a single context menu slot.
*/
//go:nosplit
func (self class) AddContextMenuPlugin(slot gdclass.EditorContextMenuPluginContextMenuSlot, plugin [1]gdclass.EditorContextMenuPlugin) { //gd:EditorPlugin.add_context_menu_plugin
	var frame = callframe.New()
	callframe.Arg(frame, slot)
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_context_menu_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the specified context menu plugin.
*/
//go:nosplit
func (self class) RemoveContextMenuPlugin(plugin [1]gdclass.EditorContextMenuPlugin) { //gd:EditorPlugin.remove_context_menu_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_context_menu_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [EditorInterface] singleton instance.
*/
//go:nosplit
func (self class) GetEditorInterface() [1]gdclass.EditorInterface { //gd:EditorPlugin.get_editor_interface
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_get_editor_interface, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorInterface{gd.PointerLifetimeBoundTo[gdclass.EditorInterface](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Gets the Editor's dialog used for making scripts.
[b]Note:[/b] Users can configure it before use.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetScriptCreateDialog() [1]gdclass.ScriptCreateDialog { //gd:EditorPlugin.get_script_create_dialog
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_get_script_create_dialog, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ScriptCreateDialog{gd.PointerLifetimeBoundTo[gdclass.ScriptCreateDialog](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adds a [Script] as debugger plugin to the Debugger. The script must extend [EditorDebuggerPlugin].
*/
//go:nosplit
func (self class) AddDebuggerPlugin(script [1]gdclass.EditorDebuggerPlugin) { //gd:EditorPlugin.add_debugger_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(script[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_debugger_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the debugger plugin with given script from the Debugger.
*/
//go:nosplit
func (self class) RemoveDebuggerPlugin(script [1]gdclass.EditorDebuggerPlugin) { //gd:EditorPlugin.remove_debugger_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(script[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_debugger_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Provide the version of the plugin declared in the [code]plugin.cfg[/code] config file.
*/
//go:nosplit
func (self class) GetPluginVersion() String.Readable { //gd:EditorPlugin.get_plugin_version
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_get_plugin_version, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}
func (self Instance) OnSceneChanged(cb func(scene_root [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("scene_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSceneClosed(cb func(filepath string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("scene_closed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMainScreenChanged(cb func(screen_name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("main_screen_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourceSaved(cb func(resource [1]gdclass.Resource)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resource_saved"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSceneSaved(cb func(filepath string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("scene_saved"), gd.NewCallable(cb), 0)
}

func (self Instance) OnProjectSettingsChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("project_settings_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorPlugin() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorPlugin() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced       { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance    { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_forward_canvas_gui_input":
		return reflect.ValueOf(self._forward_canvas_gui_input)
	case "_forward_canvas_draw_over_viewport":
		return reflect.ValueOf(self._forward_canvas_draw_over_viewport)
	case "_forward_canvas_force_draw_over_viewport":
		return reflect.ValueOf(self._forward_canvas_force_draw_over_viewport)
	case "_forward_3d_gui_input":
		return reflect.ValueOf(self._forward_3d_gui_input)
	case "_forward_3d_draw_over_viewport":
		return reflect.ValueOf(self._forward_3d_draw_over_viewport)
	case "_forward_3d_force_draw_over_viewport":
		return reflect.ValueOf(self._forward_3d_force_draw_over_viewport)
	case "_get_plugin_name":
		return reflect.ValueOf(self._get_plugin_name)
	case "_get_plugin_icon":
		return reflect.ValueOf(self._get_plugin_icon)
	case "_has_main_screen":
		return reflect.ValueOf(self._has_main_screen)
	case "_make_visible":
		return reflect.ValueOf(self._make_visible)
	case "_edit":
		return reflect.ValueOf(self._edit)
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_get_state":
		return reflect.ValueOf(self._get_state)
	case "_set_state":
		return reflect.ValueOf(self._set_state)
	case "_clear":
		return reflect.ValueOf(self._clear)
	case "_get_unsaved_status":
		return reflect.ValueOf(self._get_unsaved_status)
	case "_save_external_data":
		return reflect.ValueOf(self._save_external_data)
	case "_apply_changes":
		return reflect.ValueOf(self._apply_changes)
	case "_get_breakpoints":
		return reflect.ValueOf(self._get_breakpoints)
	case "_set_window_layout":
		return reflect.ValueOf(self._set_window_layout)
	case "_get_window_layout":
		return reflect.ValueOf(self._get_window_layout)
	case "_build":
		return reflect.ValueOf(self._build)
	case "_enable_plugin":
		return reflect.ValueOf(self._enable_plugin)
	case "_disable_plugin":
		return reflect.ValueOf(self._disable_plugin)
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_forward_canvas_gui_input":
		return reflect.ValueOf(self._forward_canvas_gui_input)
	case "_forward_canvas_draw_over_viewport":
		return reflect.ValueOf(self._forward_canvas_draw_over_viewport)
	case "_forward_canvas_force_draw_over_viewport":
		return reflect.ValueOf(self._forward_canvas_force_draw_over_viewport)
	case "_forward_3d_gui_input":
		return reflect.ValueOf(self._forward_3d_gui_input)
	case "_forward_3d_draw_over_viewport":
		return reflect.ValueOf(self._forward_3d_draw_over_viewport)
	case "_forward_3d_force_draw_over_viewport":
		return reflect.ValueOf(self._forward_3d_force_draw_over_viewport)
	case "_get_plugin_name":
		return reflect.ValueOf(self._get_plugin_name)
	case "_get_plugin_icon":
		return reflect.ValueOf(self._get_plugin_icon)
	case "_has_main_screen":
		return reflect.ValueOf(self._has_main_screen)
	case "_make_visible":
		return reflect.ValueOf(self._make_visible)
	case "_edit":
		return reflect.ValueOf(self._edit)
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_get_state":
		return reflect.ValueOf(self._get_state)
	case "_set_state":
		return reflect.ValueOf(self._set_state)
	case "_clear":
		return reflect.ValueOf(self._clear)
	case "_get_unsaved_status":
		return reflect.ValueOf(self._get_unsaved_status)
	case "_save_external_data":
		return reflect.ValueOf(self._save_external_data)
	case "_apply_changes":
		return reflect.ValueOf(self._apply_changes)
	case "_get_breakpoints":
		return reflect.ValueOf(self._get_breakpoints)
	case "_set_window_layout":
		return reflect.ValueOf(self._set_window_layout)
	case "_get_window_layout":
		return reflect.ValueOf(self._get_window_layout)
	case "_build":
		return reflect.ValueOf(self._build)
	case "_enable_plugin":
		return reflect.ValueOf(self._enable_plugin)
	case "_disable_plugin":
		return reflect.ValueOf(self._disable_plugin)
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("EditorPlugin", func(ptr gd.Object) any {
		return [1]gdclass.EditorPlugin{*(*gdclass.EditorPlugin)(unsafe.Pointer(&ptr))}
	})
}

type CustomControlContainer = gdclass.EditorPluginCustomControlContainer //gd:EditorPlugin.CustomControlContainer

const (
	/*Main editor toolbar, next to play buttons.*/
	ContainerToolbar CustomControlContainer = 0
	/*The toolbar that appears when 3D editor is active.*/
	ContainerSpatialEditorMenu CustomControlContainer = 1
	/*Left sidebar of the 3D editor.*/
	ContainerSpatialEditorSideLeft CustomControlContainer = 2
	/*Right sidebar of the 3D editor.*/
	ContainerSpatialEditorSideRight CustomControlContainer = 3
	/*Bottom panel of the 3D editor.*/
	ContainerSpatialEditorBottom CustomControlContainer = 4
	/*The toolbar that appears when 2D editor is active.*/
	ContainerCanvasEditorMenu CustomControlContainer = 5
	/*Left sidebar of the 2D editor.*/
	ContainerCanvasEditorSideLeft CustomControlContainer = 6
	/*Right sidebar of the 2D editor.*/
	ContainerCanvasEditorSideRight CustomControlContainer = 7
	/*Bottom panel of the 2D editor.*/
	ContainerCanvasEditorBottom CustomControlContainer = 8
	/*Bottom section of the inspector.*/
	ContainerInspectorBottom CustomControlContainer = 9
	/*Tab of Project Settings dialog, to the left of other tabs.*/
	ContainerProjectSettingTabLeft CustomControlContainer = 10
	/*Tab of Project Settings dialog, to the right of other tabs.*/
	ContainerProjectSettingTabRight CustomControlContainer = 11
)

type DockSlot = gdclass.EditorPluginDockSlot //gd:EditorPlugin.DockSlot

const (
	/*Dock slot, left side, upper-left (empty in default layout).*/
	DockSlotLeftUl DockSlot = 0
	/*Dock slot, left side, bottom-left (empty in default layout).*/
	DockSlotLeftBl DockSlot = 1
	/*Dock slot, left side, upper-right (in default layout includes Scene and Import docks).*/
	DockSlotLeftUr DockSlot = 2
	/*Dock slot, left side, bottom-right (in default layout includes FileSystem dock).*/
	DockSlotLeftBr DockSlot = 3
	/*Dock slot, right side, upper-left (in default layout includes Inspector, Node, and History docks).*/
	DockSlotRightUl DockSlot = 4
	/*Dock slot, right side, bottom-left (empty in default layout).*/
	DockSlotRightBl DockSlot = 5
	/*Dock slot, right side, upper-right (empty in default layout).*/
	DockSlotRightUr DockSlot = 6
	/*Dock slot, right side, bottom-right (empty in default layout).*/
	DockSlotRightBr DockSlot = 7
	/*Represents the size of the [enum DockSlot] enum.*/
	DockSlotMax DockSlot = 8
)

type AfterGUIInput = gdclass.EditorPluginAfterGUIInput //gd:EditorPlugin.AfterGUIInput

const (
	/*Forwards the [InputEvent] to other EditorPlugins.*/
	AfterGuiInputPass AfterGUIInput = 0
	/*Prevents the [InputEvent] from reaching other Editor classes.*/
	AfterGuiInputStop AfterGUIInput = 1
	/*Pass the [InputEvent] to other editor plugins except the main [Node3D] one. This can be used to prevent node selection changes and work with sub-gizmos instead.*/
	AfterGuiInputCustom AfterGUIInput = 2
)
