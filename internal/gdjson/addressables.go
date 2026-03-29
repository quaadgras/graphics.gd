package gdjson

// Addressables maps native pointer parameters and return values to their intended Go types.
// Keys are "Class.method.param_name" for arguments and "Class.method." for return values.
// Values are the Go type to use instead of the generic pointer type (e.g. "uintptr", "Pointer.Any").
// Empty string means unmapped/unknown.
var Addressables = map[string]string{
	// Audio buffers.
	"AudioEffectInstance._process.src_buffer":                "Engine.Pointer[AudioFrame]",
	"AudioEffectInstance._process.dst_buffer":                "Engine.Pointer[AudioFrame]",
	"AudioStreamPlayback._mix.buffer":                        "Engine.Pointer[AudioFrame]",
	"AudioStreamPlaybackResampled._mix_resampled.dst_buffer": "Engine.Pointer[AudioFrame]",
	"MovieWriter._write_frame.audio_frame_block":             "Engine.Pointer[int32]",

	// OpenXR typed structs — known C layouts in classdb/OpenXR.
	"OpenXRAPIExtension.transform_from_pose.pose":    "Engine.Pointer[OpenXR.Posef]",
	"OpenXRAPIExtension.set_custom_play_space.space": "uintptr",

	// OpenXR extension chain next pointers.
	"OpenXRExtensionWrapper._set_system_properties_and_get_next_pointer.next_pointer":                     "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_instance_create_info_and_get_next_pointer.next_pointer":                  "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_session_create_and_get_next_pointer.next_pointer":                        "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_swapchain_create_info_and_get_next_pointer.next_pointer":                 "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_hand_joint_locations_and_get_next_pointer.next_pointer":                  "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_projection_views_and_get_next_pointer.next_pointer":                      "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_frame_wait_info_and_get_next_pointer.next_pointer":                       "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_frame_end_info_and_get_next_pointer.next_pointer":                        "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_view_locate_info_and_get_next_pointer.next_pointer":                      "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_reference_space_create_info_and_get_next_pointer.next_pointer":           "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_view_configuration_and_get_next_pointer.next_pointer":                    "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_viewport_composition_layer_and_get_next_pointer.next_pointer":            "Engine.Pointer[OpenXR.Extension]",
	"OpenXRExtensionWrapper._set_android_surface_swapchain_create_info_and_get_next_pointer.next_pointer": "Engine.Pointer[OpenXR.Extension]",

	// OpenXR event and composition layer pointers.
	"OpenXRExtensionWrapper._on_event_polled.event":                                     "Engine.Pointer[OpenXR.EventDataBuffer]",
	"OpenXRExtensionWrapper._set_viewport_composition_layer_and_get_next_pointer.layer": "Engine.Pointer[OpenXR.CompositionLayer]",
	"OpenXRExtensionWrapper._on_viewport_composition_layer_destroyed.layer":             "Engine.Pointer[OpenXR.CompositionLayer]",

	// GDExtension function pointers.
	"GDExtensionManager.load_extension_from_function.init_func": "uintptr",

	// Multiplayer/Packet peer — byte buffers.
	"MultiplayerPeerExtension._get_packet.r_buffer":      "Engine.Pointer[Engine.Pointer[byte]]",
	"MultiplayerPeerExtension._get_packet.r_buffer_size": "Engine.Pointer[int32]",
	"MultiplayerPeerExtension._put_packet.p_buffer":      "Engine.Pointer[byte]",
	"PacketPeerExtension._get_packet.r_buffer":           "Engine.Pointer[Engine.Pointer[byte]]",
	"PacketPeerExtension._get_packet.r_buffer_size":      "Engine.Pointer[int32]",
	"PacketPeerExtension._put_packet.p_buffer":           "Engine.Pointer[byte]",

	// Physics 2D — results and out-params.
	"PhysicsDirectSpaceState2DExtension._intersect_ray.result":       "Engine.Pointer[RayResult]",
	"PhysicsDirectSpaceState2DExtension._intersect_point.results":    "Engine.Pointer[ShapeResult]",
	"PhysicsDirectSpaceState2DExtension._intersect_shape.result":     "Engine.Pointer[ShapeResult]",
	"PhysicsDirectSpaceState2DExtension._cast_motion.closest_safe":   "Engine.Pointer[float32]",
	"PhysicsDirectSpaceState2DExtension._cast_motion.closest_unsafe": "Engine.Pointer[float32]",
	"PhysicsDirectSpaceState2DExtension._collide_shape.results":      "Engine.Pointer[Vector2.XY]",
	"PhysicsDirectSpaceState2DExtension._collide_shape.result_count": "Engine.Pointer[int32]",
	"PhysicsDirectSpaceState2DExtension._rest_info.rest_info":        "Engine.Pointer[ShapeRestInfo]",

	// Physics 3D — results and out-params.
	"PhysicsDirectSpaceState3DExtension._intersect_ray.result":         "Engine.Pointer[RayResult]",
	"PhysicsDirectSpaceState3DExtension._intersect_point.results":      "Engine.Pointer[ShapeResult]",
	"PhysicsDirectSpaceState3DExtension._intersect_shape.result_count": "Engine.Pointer[ShapeResult]",
	"PhysicsDirectSpaceState3DExtension._cast_motion.closest_safe":     "Engine.Pointer[float32]",
	"PhysicsDirectSpaceState3DExtension._cast_motion.closest_unsafe":   "Engine.Pointer[float32]",
	"PhysicsDirectSpaceState3DExtension._cast_motion.info":             "Engine.Pointer[ShapeRestInfo]",
	"PhysicsDirectSpaceState3DExtension._collide_shape.results":        "Engine.Pointer[Vector3.XYZ]",
	"PhysicsDirectSpaceState3DExtension._collide_shape.result_count":   "Engine.Pointer[int32]",
	"PhysicsDirectSpaceState3DExtension._rest_info.rest_info":          "Engine.Pointer[ShapeRestInfo]",

	// Physics server — collision results and motion testing.
	"PhysicsServer2DExtension._shape_collide.results":           "Engine.Pointer[Vector2.XY]",
	"PhysicsServer2DExtension._shape_collide.result_count":      "Engine.Pointer[int32]",
	"PhysicsServer2DExtension._body_collide_shape.results":      "Engine.Pointer[Vector2.XY]",
	"PhysicsServer2DExtension._body_collide_shape.result_count": "Engine.Pointer[int32]",
	"PhysicsServer2DExtension._body_test_motion.result":         "Engine.Pointer[MotionResult]",
	"PhysicsServer3DExtension._body_test_motion.result":         "Engine.Pointer[MotionResult]",

	// Script extension — opaque handles and profiling.
	"ScriptExtension._placeholder_erased.placeholder":                    "uintptr",
	"ScriptExtension._instance_create.":                                  "uintptr",
	"ScriptExtension._placeholder_instance_create.":                      "uintptr",
	"ScriptLanguageExtension._debug_get_stack_level_instance.":           "uintptr",
	"ScriptLanguageExtension._profiling_get_accumulated_data.info_array": "Engine.Pointer[ProfilingInfo]",
	"ScriptLanguageExtension._profiling_get_frame_data.info_array":       "Engine.Pointer[ProfilingInfo]",

	// Stream peer — byte buffers and out-params.
	"StreamPeerExtension._get_data.r_buffer":           "Engine.Pointer[byte]",
	"StreamPeerExtension._get_data.r_received":         "Engine.Pointer[int32]",
	"StreamPeerExtension._get_partial_data.r_buffer":   "Engine.Pointer[byte]",
	"StreamPeerExtension._get_partial_data.r_received": "Engine.Pointer[int32]",
	"StreamPeerExtension._put_data.p_data":             "Engine.Pointer[byte]",
	"StreamPeerExtension._put_data.r_sent":             "Engine.Pointer[int32]",
	"StreamPeerExtension._put_partial_data.p_data":     "Engine.Pointer[byte]",
	"StreamPeerExtension._put_partial_data.r_sent":     "Engine.Pointer[int32]",

	// Text server — font data and glyph buffers.
	"TextServerExtension._font_set_data_ptr.data_ptr":       "Engine.Pointer[byte]",
	"TextServerExtension._shaped_text_get_glyphs.":          "Engine.Pointer[Glyph]",
	"TextServerExtension._shaped_text_sort_logical.":        "Engine.Pointer[Glyph]",
	"TextServerExtension._shaped_text_get_ellipsis_glyphs.": "Engine.Pointer[Glyph]",
	"TextServerExtension._shaped_text_get_carets.caret":     "Engine.Pointer[CaretInfo]",

	// WebRTC — byte buffers.
	"WebRTCDataChannelExtension._get_packet.r_buffer":      "Engine.Pointer[Engine.Pointer[byte]]",
	"WebRTCDataChannelExtension._get_packet.r_buffer_size": "Engine.Pointer[int32]",
	"WebRTCDataChannelExtension._put_packet.p_buffer":      "Engine.Pointer[byte]",
}
