package gdjson

var Consitution = map[string]map[string]string{
	"AnimatedTexture":               {"MAX_": "int"},
	"AnimationNodeBlendTree":        {"CONNECTION_": "-"},
	"AudioStreamInteractive":        {"CLIP_": "Clip"},
	"AudioStreamPlaybackPolyphonic": {"INVALID_ID": "Stream"},
	"AudioStreamPlaylist":           {"MAX_": "int"},
	"AudioStreamSynchronized":       {"MAX_": "int"},
	"CanvasItem":                    {"NOTIFICATION_": "Object.Notification"},
	"Container":                     {"NOTIFICATION_": "Object.Notification"},
	"Control":                       {"NOTIFICATION_": "Object.Notification"},
	"DisplayServer": {
		"SCREEN_":              "Screen",
		"MAIN_WINDOW_ID":       "Window",
		"INVALID_WINDOW_ID":    "Window",
		"INVALID_INDICATOR_ID": "StatusIndicator",
	},
	"ENetPacketPeer": {
		"PACKET_LOSS_SCALE":     "",
		"PACKET_THROTTLE_SCALE": "",
		"FLAG_":                 "Flags",
	},
	"EditorSceneFormatImporter": {
		"IMPORT_": "Flags",
	},
	"EditorSettings": {"NOTIFICATION_": "Object.Notification"},
	"GLTFState":      {"HANDLE_BINARY_": "BinaryHandler"},
	"GPUParticles3D": {"MAX_": "int"},
	"GridMap":        {"INVALID_CELL_ITEM": "CellItem"},
	"IP": {
		"RESOLVER_MAX":        "",
		"RESOLVER_INVALID_ID": "", // FIXME?
	},
	"Image":             {"MAX_": ""},
	"InputEvent":        {"DEVICE_ID_": "-"},
	"MainLoop":          {"NOTIFICATION_": "Object.Notification"},
	"Material":          {"RENDER_PRIORITY": "RenderPriority"},
	"MultiplayerPeer":   {"TARGET_PEER_": "TargetPeer"},
	"Node":              {"NOTIFICATION_": "Object.Notification"},
	"Node3D":            {"NOTIFICATION_": "Object.Notification"},
	"RDFramebufferPass": {"ATTACHMENT_UNUSED": ""},
	"RenderingDevice":   {"INVALID_ID": "", "INVALID_FORMAT_ID": ""},
	"RenderingServer": {
		"MAX_":                         "int",
		"ARRAY_CUSTOM_COUNT":           "int",
		"ARRAY_WEIGHTS_SIZE":           "int",
		"CANVAS_ITEM_Z_MIN":            "int",
		"CANVAS_ITEM_Z_MAX":            "int",
		"MATERIAL_RENDER_PRIORITY_MIN": "int",
		"MATERIAL_RENDER_PRIORITY_MAX": "int",
		"PARTICLES_EMIT_FLAG":          "ParticlesEmitFlag",
		"NO_INDEX_ARRAY":               "-",
	},
	"ResourceUID": {
		"INVALID_ID": "Resource.UID",
	},
	"Skeleton3D":         {"NOTIFICATION_": "Object.Notification"},
	"TileSetAtlasSource": {"TRANSFORM_": ""}, // FIXME?
	"VisualShader":       {"NODE_ID_": "NodeID"},
	"Window":             {"NOTIFICATION_": "Object.Notification"},
}
