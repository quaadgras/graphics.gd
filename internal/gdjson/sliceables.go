package gdjson

// Sliceable describes how a pointer+count parameter pair maps to a Go slice/array type.
type Sliceable struct {
	Elem  string // Elem is the Go element type (e.g. "Vector2.XY", "Float.X").
	Count string // Count is the name of the companion parameter that holds the element count.
}

// Sliceables maps pointer parameters that represent typed buffers to their
// element types and count companions. Keys are "Class.method.param_name".
var Sliceables = map[string]Sliceable{
	// Physics collision result buffers — contact point pairs.
	"PhysicsDirectSpaceState2DExtension._collide_shape.r_results": {Elem: "Vector2.XY", Count: "max_results"},
	"PhysicsDirectSpaceState3DExtension._collide_shape.r_results": {Elem: "Vector3.XYZ", Count: "max_results"},
	"PhysicsServer2DExtension._shape_collide.r_results":           {Elem: "Vector2.XY", Count: "result_max"},
	"PhysicsServer2DExtension._body_collide_shape.r_results":      {Elem: "Vector2.XY", Count: "result_max"},

	// Audio buffers.
	"AudioEffectInstance._process.src_buffer":                {Elem: "AudioFrame", Count: "frame_count"},
	"AudioEffectInstance._process.r_dst_buffer":              {Elem: "AudioFrame", Count: "frame_count"},
	"AudioStreamPlayback._mix.buffer":                        {Elem: "AudioFrame", Count: "frames"},
	"AudioStreamPlaybackResampled._mix_resampled.dst_buffer": {Elem: "AudioFrame", Count: "frame_count"},

	// Physics 2D — result arrays.
	"PhysicsDirectSpaceState2DExtension._intersect_point.r_results": {Elem: "ShapeResult", Count: "max_results"},
	"PhysicsDirectSpaceState2DExtension._intersect_shape.r_result":  {Elem: "ShapeResult", Count: "max_results"},

	// Physics 3D — result arrays.
	"PhysicsDirectSpaceState3DExtension._intersect_point.r_results":      {Elem: "ShapeResult", Count: "max_results"},
	"PhysicsDirectSpaceState3DExtension._intersect_shape.r_result_count": {Elem: "ShapeResult", Count: "max_results"},

	// Script profiling.
	"ScriptLanguageExtension._profiling_get_accumulated_data.info_array": {Elem: "ProfilingInfo", Count: "info_max"},
	"ScriptLanguageExtension._profiling_get_frame_data.info_array":       {Elem: "ProfilingInfo", Count: "info_max"},

	// Stream peer — byte buffers.
	"StreamPeerExtension._get_data.r_buffer":         {Elem: "byte", Count: "r_bytes"},
	"StreamPeerExtension._get_partial_data.r_buffer": {Elem: "byte", Count: "r_bytes"},
	"StreamPeerExtension._put_data.data":             {Elem: "byte", Count: "bytes"},
	"StreamPeerExtension._put_partial_data.data":     {Elem: "byte", Count: "bytes"},

	// Multiplayer/Packet peer — byte buffers.
	"MultiplayerPeerExtension._put_packet.buffer": {Elem: "byte", Count: "buffer_size"},
	"PacketPeerExtension._put_packet.buffer":      {Elem: "byte", Count: "buffer_size"},

	// Text server — font data.
	"TextServerExtension._font_set_data_ptr.data_ptr": {Elem: "byte", Count: "data_size"},
}
