package gdjson

/*
OwnershipSemantics attempts to capture a proposal and pragmatic categorization
around the ownership semantics for non-RefCounted godot objects.
*/
type OwnershipSemantics int

const (
	// RefCountedManagement is the default ownership semantics for objects that
	// extend the RefCounted class.
	RefCountedManagement OwnershipSemantics = iota

	// OwnershipTransferred means that the receiver of the value is responsible
	// for freeing the memory associated with the value. After transferring the
	// ownership of a value this way, no direct references may be kept by the prior
	// owner, although they may keep copies of the InstanceID and use it to look
	// up the object in the future.
	OwnershipTransferred

	// ReversesTheOwnership applies to function arguments and means if (and only
	// if) the class is known to be the owner of the object, or is known to bind
	// the lifetime of this value, once the function returns, the caller will own
	// the object instead and is no longer under an obligation to keep it alive.
	// In such cases, the caller becomes responsible for freeing the object.
	ReversesTheOwnership

	// IsTemporaryReference only applies to function arguments and means that
	// the function borrows the argument (for the lifetime of the call) and does
	// not take ownership of it. The function may store the InstanceID of an
	// argument, in order to look it up in future.
	IsTemporaryReference

	// MustAssertInstanceID means that the receiver must always assert that
	// the instance ID is valid before using the object. The object may be
	// freed at any time.
	MustAssertInstanceID

	// LifetimeBoundToClass when applied to function arguments, means that the
	// receiver may keep a reference to the object and the caller must keep
	// it alive for as long as the class is alive. For return values, they
	// are only valid as long as the class they are returned from is alive,
	// the caller may not free the object.
	LifetimeBoundToClass
)

var ClassMethodOwnership = map[string]map[string]map[string]OwnershipSemantics{
	"AcceptDialog": {
		"get_ok_button": {
			"return value": LifetimeBoundToClass,
		},
		"get_label": {
			"return value": LifetimeBoundToClass,
		},
		"remove_button": {
			"button": ReversesTheOwnership,
		},
		"register_text_enter": {
			"line_edit": IsTemporaryReference,
		},
		"add_button": {
			"return value": LifetimeBoundToClass,
		},
		"add_cancel_button": {
			"return value": LifetimeBoundToClass,
		},
	},
	"BoxContainer": {
		"add_spacer": {
			"return value": LifetimeBoundToClass,
		},
	},
	"ButtonGroup": {
		"get_pressed_button": {
			"return value": MustAssertInstanceID,
		},
	},
	"Area2D": {
		"overlaps_body": {
			"body": IsTemporaryReference,
		},
		"overlaps_area": {
			"area": IsTemporaryReference,
		},
	},
	"Area3D": {
		"overlaps_body": {
			"body": IsTemporaryReference,
		},
		"overlaps_area": {
			"area": IsTemporaryReference,
		},
	},
	"CPUParticles2D": {
		"convert_from_particles": {
			"particles": IsTemporaryReference,
		},
	},
	"CPUParticles3D": {
		"convert_from_particles": {
			"particles": IsTemporaryReference,
		},
	},
	"Camera2D": {
		"set_custom_viewport": {
			"viewport": MustAssertInstanceID,
		},
		"get_custom_viewport": {
			"return value": MustAssertInstanceID,
		},
	},
	"CanvasLayer": {
		"set_custom_viewport": {
			"viewport": MustAssertInstanceID,
		},
		"get_custom_viewport": {
			"return value": MustAssertInstanceID,
		},
	},
	"ColorPickerButton": {
		"get_picker": {
			"return value": LifetimeBoundToClass,
		},
		"get_popup": {
			"return value": LifetimeBoundToClass,
		},
	},
	"ConfirmationDialog": {
		"get_cancel_button": {
			"return value": LifetimeBoundToClass,
		},
	},
	"Container": {
		"fit_child_in_rect": {
			"child": IsTemporaryReference,
		},
	},
	"Control": {
		"force_drag": {
			"preview": OwnershipTransferred,
		},
		"set_drag_preview": {
			"control": OwnershipTransferred,
		},
		"set_shortcut_context": {
			"node": MustAssertInstanceID,
		},
		"find_prev_valid_focus": {
			"return value": MustAssertInstanceID,
		},
		"find_next_valid_focus": {
			"return value": MustAssertInstanceID,
		},
		"find_valid_focus_neighbor": {
			"return value": MustAssertInstanceID,
		},
		"get_parent_control": {
			"return value": MustAssertInstanceID,
		},
		"get_shortcut_context": {
			"return value": MustAssertInstanceID,
		},
	},
	"EditorDebuggerSession": {
		"add_session_tab": {
			"control": OwnershipTransferred,
		},
		"remove_session_tab": {
			"control": ReversesTheOwnership,
		},
	},
	"EditorFileDialog": {
		"add_side_menu": {
			"menu": OwnershipTransferred,
		},
		"get_vbox": {
			"return value": LifetimeBoundToClass,
		},
		"get_line_edit": {
			"return value": LifetimeBoundToClass,
		},
	},
	"EditorFileSystem": {
		"get_filesystem": {
			"return value": LifetimeBoundToClass,
		},
		"get_filesystem_path": {
			"return value": LifetimeBoundToClass,
		},
	},
	"EditorFileSystemDirectory": {
		"get_subdir": {
			"return value": MustAssertInstanceID,
		},
		"get_parent": {
			"return value": MustAssertInstanceID,
		},
	},
	"EditorInspectorPlugin": {
		"add_custom_control": {
			"control": OwnershipTransferred,
		},
		"add_property_editor": {
			"editor": OwnershipTransferred,
		},
		"add_property_editor_for_multiple_properties": {
			"editor": OwnershipTransferred,
		},
	},
	"EditorInterface": {
		"get_command_palette": {
			"return value": LifetimeBoundToClass,
		},
		"get_resource_filesystem": {
			"return value": LifetimeBoundToClass,
		},
		"get_editor_paths": {
			"return value": LifetimeBoundToClass,
		},
		"get_resource_previewer": {
			"return value": LifetimeBoundToClass,
		},
		"get_selection": {
			"return value": LifetimeBoundToClass,
		},
		"get_base_control": {
			"return value": LifetimeBoundToClass,
		},
		"get_editor_main_screen": {
			"return value": LifetimeBoundToClass,
		},
		"get_script_editor": {
			"return value": LifetimeBoundToClass,
		},
		"get_editor_viewport_2d": {
			"return value": LifetimeBoundToClass,
		},
		"get_editor_viewport_3d": {
			"return value": LifetimeBoundToClass,
		},
		"get_file_system_dock": {
			"return value": LifetimeBoundToClass,
		},
		"get_inspector": {
			"return value": LifetimeBoundToClass,
		},
		"get_edited_scene_root": {
			"return value": MustAssertInstanceID,
		},
		"popup_dialog": {
			"dialog": IsTemporaryReference,
		},
		"popup_dialog_centered": {
			"dialog": IsTemporaryReference,
		},
		"popup_dialog_centered_ratio": {
			"dialog": IsTemporaryReference,
		},
		"popup_dialog_centered_clamped": {
			"dialog": IsTemporaryReference,
		},
		"edit_node": {
			"node": MustAssertInstanceID,
		},
	},
	"EditorNode3DGizmo": {
		"set_node_3d": {
			"node": LifetimeBoundToClass,
		},
		"get_node_3d": {
			"return value": LifetimeBoundToClass,
		},
	},
	"EditorPlugin": {
		"add_control_to_container": {
			"control": OwnershipTransferred,
		},
		"add_control_to_bottom_panel": {
			"return value": MustAssertInstanceID,
			"control":      OwnershipTransferred,
		},
		"add_control_to_dock": {
			"control": OwnershipTransferred,
		},
		"remove_control_from_docks": {
			"control": ReversesTheOwnership,
		},
		"remove_control_from_bottom_panel": {
			"control": ReversesTheOwnership,
		},
		"remove_control_from_container": {
			"control": ReversesTheOwnership,
		},
		"add_tool_submenu_item": {
			"submenu": OwnershipTransferred,
		},
		"get_export_as_menu": {
			"return value": LifetimeBoundToClass,
		},
		"make_bottom_panel_item_visible": {
			"item": IsTemporaryReference,
		},
		"get_undo_redo": {
			"return value": LifetimeBoundToClass,
		},
		"get_editor_interface": {
			"return value": LifetimeBoundToClass,
		},
		"get_script_create_dialog": {
			"return value": LifetimeBoundToClass,
		},
	},
	"EditorProperty": {
		"add_focusable": {
			"control": LifetimeBoundToClass,
		},
		"set_bottom_editor": {
			"editor": LifetimeBoundToClass,
		},
	},
	"EditorResourceTooltipPlugin": {
		"request_thumbnail": {
			"control": MustAssertInstanceID,
		},
	},
	"EditorScript": {
		"add_root_node": {
			"node": OwnershipTransferred,
		},
		"get_scene": {
			"return value": MustAssertInstanceID,
		},
		"get_editor_interface": {
			"return value": LifetimeBoundToClass,
		},
	},
	"EditorScriptPicker": {
		"set_script_owner": {
			"owner_node": LifetimeBoundToClass,
		},
		"get_script_owner": {
			"return value": MustAssertInstanceID,
		},
	},
	"EditorSelection": {
		"add_node": {
			"node": OwnershipTransferred,
		},
		"remove_node": {
			"node": ReversesTheOwnership,
		},
	},
	"EditorUndoRedoManager": {
		"get_history_undo_redo": {
			"return value": LifetimeBoundToClass,
		},
	},
	"Engine": {
		"get_main_loop": {
			"return value": LifetimeBoundToClass,
		},
		"register_script_language": {
			"language": OwnershipTransferred,
		},
		"unregister_script_language": {
			"language": ReversesTheOwnership,
		},
		"get_script_language": {
			"return value": MustAssertInstanceID,
		},
	},
	"FileDialog": {
		"get_vbox": {
			"return value": LifetimeBoundToClass,
		},
		"get_line_edit": {
			"return value": LifetimeBoundToClass,
		},
	},
	"GLTFCamera": {
		"from_node": {
			"camera_node": IsTemporaryReference,
		},
		"to_node": {
			"return value": OwnershipTransferred,
		},
	},
	"GLTFDocument": {
		"append_from_scene": {
			"node": IsTemporaryReference,
		},
		"generate_scene": {
			"return value": OwnershipTransferred,
		},
	},
	"GLTFLight": {
		"from_node": {
			"light_node": IsTemporaryReference,
		},
		"to_node": {
			"return value": OwnershipTransferred,
		},
	},
	"GLTFPhysicsBody": {
		"from_node": {
			"body_node": IsTemporaryReference,
		},
		"to_node": {
			"return value": OwnershipTransferred,
		},
	},
	"GLTFPhysicsShape": {
		"from_node": {
			"shape_node": IsTemporaryReference,
		},
		"to_node": {
			"return value": OwnershipTransferred,
		},
	},
	"GLTFSkeleton": {
		"get_godot_skeleton": {
			"return value": OwnershipTransferred,
		},
		"get_bone_attachment": {
			"return value": OwnershipTransferred,
		},
	},
	"GLTFState": {
		"get_animation_player": {
			"return value": MustAssertInstanceID,
		},
		"get_scene_node": {
			"return value": MustAssertInstanceID,
		},
		"get_node_index": {
			"scene_node": IsTemporaryReference,
		},
	},
	"GPUParticles2D": {
		"convert_from_particles": {
			"particles": IsTemporaryReference,
		},
	},
	"GPUParticles3D": {
		"convert_from_particles": {
			"particles": IsTemporaryReference,
		},
	},
	"GraphEdit": {
		"get_menu_hbox": {
			"return value": LifetimeBoundToClass,
		},
		"set_selected": {
			"node": IsTemporaryReference,
		},
	},
	"GraphNode": {
		"get_titlebar_hbox": {
			"return value": LifetimeBoundToClass,
		},
	},
	"InstancePlaceholder": {
		"create_instance": {
			"return value": OwnershipTransferred,
		},
	},
	"ItemList": {
		"get_v_scroll_bar": {
			"return value": LifetimeBoundToClass,
		},
	},
	"LineEdit": {
		"get_menu": {
			"return value": LifetimeBoundToClass,
		},
	},
	"MenuBar": {
		"get_menu_popup": {
			"return value": LifetimeBoundToClass,
		},
	},
	"MenuButton": {
		"get_popup": {
			"return value": LifetimeBoundToClass,
		},
	},
	"MovieWriter": {
		"add_writer": {
			"writer": OwnershipTransferred,
		},
	},
	"MultiplayerSpawner": {
		"spawn": {
			"return value": OwnershipTransferred,
		},
	},
	"NavigationMeshGenerator": {
		"bake": {
			"root_node": IsTemporaryReference,
		},
		"parse_source_geometry_data": {
			"root_node": IsTemporaryReference,
		},
	},
	"NavigationServer2D": {
		"parse_source_geometry_data": {
			"root_node": IsTemporaryReference,
		},
	},
	"NavigationServer3D": {
		"region_bake_navigation_mesh": {
			"root_node": IsTemporaryReference,
		},
		"parse_source_geometry_data": {
			"root_node": IsTemporaryReference,
		},
	},
	"Node": {
		"add_sibling": {
			"sibling": OwnershipTransferred,
		},
		"add_child": {
			"node": OwnershipTransferred,
		},
		"remove_child": {
			"node": ReversesTheOwnership,
		},
		"reparent": {
			"new_parent": IsTemporaryReference,
		},
		"get_child": {
			"return value": MustAssertInstanceID,
		},
		"get_node": {
			"return value": MustAssertInstanceID,
		},
		"get_node_or_null": {
			"return value": MustAssertInstanceID,
		},
		"get_parent": {
			"return value": MustAssertInstanceID,
		},
		"find_child": {
			"return value": MustAssertInstanceID,
		},
		"find_parent": {
			"return value": MustAssertInstanceID,
		},
		"is_ancestor_of": {
			"node": IsTemporaryReference,
		},
		"is_greater_than": {
			"node": IsTemporaryReference,
		},
		"get_path_to": {
			"node": IsTemporaryReference,
		},
		"move_child": {
			"child_node": IsTemporaryReference,
		},
		"set_owner": {
			"owner": MustAssertInstanceID,
		},
		"get_owner": {
			"return value": MustAssertInstanceID,
		},
		"get_window": {
			"return value": MustAssertInstanceID,
		},
		"get_last_exclusive_window": {
			"return value": MustAssertInstanceID,
		},
		"get_tree": {
			"return value": MustAssertInstanceID,
		},
		"duplicate": {
			"return value": OwnershipTransferred,
		},
		"replace_by": {
			"node": OwnershipTransferred,
		},
		"set_editable_instance": {
			"node": IsTemporaryReference,
		},
		"is_editable_instance": {
			"node": IsTemporaryReference,
		},
		"get_viewport": {
			"return value": MustAssertInstanceID,
		},
	},
	"Node2D": {
		"get_relative_transform_to_parent": {
			"parent": IsTemporaryReference,
		},
	},
	"Node3D": {
		"get_parent_node_3d": {
			"return value": MustAssertInstanceID,
		},
	},
	"OptionButton": {
		"get_popup": {
			"return value": LifetimeBoundToClass,
		},
	},
	"PackedScene": {
		"pack": {
			"path": IsTemporaryReference,
		},
		"instantiate": {
			"return value": OwnershipTransferred,
		},
	},
	"PhysicalBone2D": {
		"get_joint": {
			"return value": MustAssertInstanceID,
		},
	},
	"PhysicsBody2D": {
		"add_collision_exception_with": {
			"body": IsTemporaryReference,
		},
		"remove_collision_exception_with": {
			"body": IsTemporaryReference,
		},
	},
	"PhysicsBody3D": {
		"add_collision_exception_with": {
			"body": IsTemporaryReference,
		},
		"remove_collision_exception_with": {
			"body": IsTemporaryReference,
		},
	},
	"PhysicsDirectBodyState2D": {
		"get_space_state": {
			"return value": MustAssertInstanceID,
		},
	},
	"PhysicsDirectBodyState3D": {
		"get_space_state": {
			"return value": MustAssertInstanceID,
		},
	},
	"PhysicsServer2D": {
		"space_get_direct_state": {
			"return value": MustAssertInstanceID,
		},
		"body_get_direct_state": {
			"return value": MustAssertInstanceID,
		},
	},
	"PhysicsServer3D": {
		"space_get_direct_state": {
			"return value": MustAssertInstanceID,
		},
		"body_get_direct_state": {
			"return value": MustAssertInstanceID,
		},
	},
	"Range": {
		"share": {
			"with": IsTemporaryReference,
		},
	},
	"RayCast2D": {
		"add_exception": {
			"node": IsTemporaryReference,
		},
		"remove_exception": {
			"node": IsTemporaryReference,
		},
	},
	"RayCast3D": {
		"add_exception": {
			"node": IsTemporaryReference,
		},
		"remove_exception": {
			"node": IsTemporaryReference,
		},
	},
	"RenderingDevice": {
		"create_local_device": {
			"return value": OwnershipTransferred,
		},
	},
	"RenderingServer": {
		"get_rendering_device": {
			"return value": IsTemporaryReference,
		},
		"create_local_rendering_device": {
			"return value": OwnershipTransferred,
		},
	},
	"Resource": {
		"get_local_scene": {
			"return value": MustAssertInstanceID,
		},
	},
	"RichTextLabel": {
		"get_v_scroll_bar": {
			"return value": LifetimeBoundToClass,
		},
		"get_menu": {
			"return value": LifetimeBoundToClass,
		},
	},
	"SceneTree": {
		"get_root": {
			"return value": MustAssertInstanceID,
		},
		"set_edited_scene_root": {
			"scene": OwnershipTransferred,
		},
		"get_edited_scene_root": {
			"return value": MustAssertInstanceID,
		},
		"get_first_node_in_group": {
			"return value": MustAssertInstanceID,
		},
		"set_current_scene": {
			"child_node": OwnershipTransferred,
		},
		"get_current_scene": {
			"return value": MustAssertInstanceID,
		},
	},
	"ScriptEditor": {
		"get_current_editor": {
			"return value": MustAssertInstanceID,
		},
	},
	"ScriptEditorBase": {
		"get_base_editor": {
			"return value": MustAssertInstanceID,
		},
	},
	"ScrollContainer": {
		"get_v_scroll_bar": {
			"return value": LifetimeBoundToClass,
		},
		"get_h_scroll_bar": {
			"return value": LifetimeBoundToClass,
		},
		"ensure_control_visible": {
			"control": IsTemporaryReference,
		},
	},
	"ShapeCast2D": {
		"add_exception": {
			"node": IsTemporaryReference,
		},
		"remove_exception": {
			"node": IsTemporaryReference,
		},
	},
	"ShapeCast3D": {
		"add_exception": {
			"node": IsTemporaryReference,
		},
		"remove_exception": {
			"node": IsTemporaryReference,
		},
	},
	"Skeleton2D": {
		"get_bone": {
			"return value": MustAssertInstanceID,
		},
	},
	"SkeletonIK3D": {
		"get_parent_skeleton": {
			"return value": MustAssertInstanceID,
		},
	},
	"SkeletonModificationStack2D": {
		"get_skeleton": {
			"return value": MustAssertInstanceID,
		},
	},
	"SoftBody3D": {
		"add_collision_exception_with": {
			"body": IsTemporaryReference,
		},
		"remove_collision_exception_with": {
			"body": IsTemporaryReference,
		},
	},
	"SpinBox": {
		"get_line_edit": {
			"return value": LifetimeBoundToClass,
		},
	},
	"StyleBox": {
		"get_current_item_drawn": {
			"return value": MustAssertInstanceID,
		},
	},
	"SyntaxHighlighter": {
		"get_text_edit": {
			"return value": MustAssertInstanceID,
		},
	},
	"TabContainer": {
		"get_current_tab_control": {
			"return value": MustAssertInstanceID,
		},
		"get_tab_bar": {
			"return value": LifetimeBoundToClass,
		},
		"get_tab_control": {
			"return value": MustAssertInstanceID,
		},
		"get_tab_idx_from_control": {
			"control": IsTemporaryReference,
		},
		"set_popup": {
			"popup": OwnershipTransferred,
		},
		"get_popup": {
			"return value": MustAssertInstanceID,
		},
	},
	"TextEdit": {
		"get_v_scroll_bar": {
			"return value": LifetimeBoundToClass,
		},
		"get_h_scroll_bar": {
			"return value": LifetimeBoundToClass,
		},
		"get_menu": {
			"return value": LifetimeBoundToClass,
		},
	},
	"TileMap": {
		"get_cell_tile_data": {
			"return value": MustAssertInstanceID,
		},
	},
	"TileSetAtlasSource": {
		"get_tile_data": {
			"return value": MustAssertInstanceID,
		},
	},
	"Tree": {
		"create_item": {
			"parent":       IsTemporaryReference,
			"return value": MustAssertInstanceID,
		},
		"get_root": {
			"return value": MustAssertInstanceID,
		},
		"get_next_selected": {
			"from":         IsTemporaryReference,
			"return value": MustAssertInstanceID,
		},
		"get_selected": {
			"return value": MustAssertInstanceID,
		},
		"set_selected": {
			"item": IsTemporaryReference,
		},
		"get_edited": {
			"return value": MustAssertInstanceID,
		},
		"get_item_area_rect": {
			"item": IsTemporaryReference,
		},
		"get_item_at_position": {
			"return value": MustAssertInstanceID,
		},
		"scroll_to_item": {
			"item": IsTemporaryReference,
		},
	},
	"TreeItem": {
		"create_child": {
			"return value": MustAssertInstanceID,
		},
		"add_child": {
			"child": OwnershipTransferred,
		},
		"remove_child": {
			"child": ReversesTheOwnership,
		},
		"get_tree": {
			"return value": MustAssertInstanceID,
		},
		"get_next": {
			"return value": MustAssertInstanceID,
		},
		"get_prev": {
			"return value": MustAssertInstanceID,
		},
		"get_parent": {
			"return value": MustAssertInstanceID,
		},
		"get_first_child": {
			"return value": MustAssertInstanceID,
		},
		"get_next_in_tree": {
			"return value": MustAssertInstanceID,
		},
		"get_prev_in_tree": {
			"return value": MustAssertInstanceID,
		},
		"get_next_visible": {
			"return value": MustAssertInstanceID,
		},
		"get_prev_visible": {
			"return value": MustAssertInstanceID,
		},
		"get_child": {
			"return value": MustAssertInstanceID,
		},
		"move_before": {
			"item": IsTemporaryReference,
		},
		"move_after": {
			"item": IsTemporaryReference,
		},
	},
	"Tween": {
		"bind_node": {
			"node": MustAssertInstanceID,
		},
	},
	"VehicleWheel3D": {
		"get_contact_body": {
			"return value": MustAssertInstanceID,
		},
	},
	"Viewport": {
		"get_camera_2d": {
			"return value": MustAssertInstanceID,
		},
		"gui_get_focus_owner": {
			"return value": MustAssertInstanceID,
		},
		"get_camera_3d": {
			"return value": MustAssertInstanceID,
		},
	},
	"VoxelGI": {
		"bake": {
			"from_node": IsTemporaryReference,
		},
	},
	"Window": {
		"popup_exclusive": {
			"from_node": IsTemporaryReference,
		},
		"popup_exclusive_on_parent": {
			"from_node": IsTemporaryReference,
		},
		"popup_exclusive_centered": {
			"from_node": IsTemporaryReference,
		},
		"popup_exclusive_centered_ratio": {
			"from_node": IsTemporaryReference,
		},
		"popup_exclusive_centered_clamped": {
			"from_node": IsTemporaryReference,
		},
	},
	"World2D": {
		"get_direct_space_state": {
			"return value": MustAssertInstanceID,
		},
	},
	"World3D": {
		"get_direct_space_state": {
			"return value": MustAssertInstanceID,
		},
	},
}
