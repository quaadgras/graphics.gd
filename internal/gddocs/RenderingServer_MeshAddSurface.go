/*
{
	# Required:
    "primitive": RenderingServer.PrimitiveType,
    "format": RenderingServer.ArrayFormat,
    "vertex_data": PackedByteArray,
    "vertex_count": int,
    "aabb": AABB,

	# Optional:
	"attribute_data": PackedByteArray,
	"skin_data": PackedByteArray,
	"index_data": PackedByteArray,
	"index_count": int, # Required if `index_data` is specified.
	"uv_scale": Vector4,
	"lods": [
		# Both values are required for each LOD level.
		{
			"edge_length": float,
			"index_data": PackedByteArray,
		},
	],
	"bone_aabbs": Array[AABB],
	"blend_shape_data": PackedByteArray,
	"material": Material,
}
*/

package main
