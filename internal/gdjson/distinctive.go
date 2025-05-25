package gdjson

var Distinctions = map[string][][3]string{
	"Astar2D":                        {{"id", "int", "Point"}, {"get_id_path ", "PackedInt64Array", "[]Point"}, {"get_point_connections ", "PackedInt64Array", "[]Point"}, {"get_point_ids ", "PackedInt64Array", "[]Point"}},
	"Astar3D":                        {{"id", "int", "Point"}, {"get_id_path ", "PackedInt64Array", "[]Point"}, {"get_point_connections ", "PackedInt64Array", "[]Point"}, {"get_point_ids ", "PackedInt64Array", "[]Point"}},
	"AStarGrid2D":                    {{"id", "int", "Point"}, {"get_id_path ", "PackedInt64Array", "[]Point"}, {"get_point_connections ", "PackedInt64Array", "[]Point"}, {"get_point_ids ", "PackedInt64Array", "[]Point"}},
	"AudioServer":                    {{"bus_idx", "int", "Bus"}, {"effect_idx", "int", "Effect"}},
	"AudioStreamInteractive":         {{"clip", "int", "Clip"}},
	"AudioStreamPlaybackInteractive": {{"clip", "int", "AudioStreamInteractive.Clip"}, {"get_current_clip_index ", "int", "AudioStreamInteractive.Clip"}},
	"AudioStreamPlaybackPolyphonic":  {{"stream", "int", "Stream"}},
	"DisplayServer":                  {{"screen", "int", "Screen"}, {"status_indicator id", "int", "StatusIndicator"}, {"create_status_indicator ", "int", "StatusIndicator"}, {"window_id", "int", "Window"}},
	"ENetPacketPeer":                 {{"flags", "int", "Flags"}},
	"EditorSceneFormatImporter":      {{"import_ flags", "int", "Flags"}},
	"GLTFState":                      {{"set_handle_binary_image method", "int", "BinaryHandler"}},
	"GridMap":                        {{"item", "int", "CellItem"}, {"get_cell_item ", "int", "CellItem"}},
	"Input":                          {{"device", "int", "Device"}},
	"InputEvent":                     {{"set_device value", "int", "Device"}, {"get_device ", "int", "Device"}},
	"Material":                       {{"get_render_priority ", "", "RenderPriority"}, {"set_render_priority value", "int", "RenderPriority"}},
	"MultiplayerPeer":                {{"set_target_peer id", "int", "TargetPeer"}},
	"RenderingServer":                {{"emit_flags", "int", "ParticlesEmitFlag"}},
	"ResourceUID":                    {{"id", "int", "Resource.UID"}},
	"VisualShader":                   {{"node", "int", "NodeID"}, {"id ", "int", "NodeID"}},
}
