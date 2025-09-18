/*
var rd = RenderingDevice.new()
var compute_list = rd.compute_list_begin()

rd.compute_list_bind_compute_pipeline(compute_list, compute_shader_dilate_pipeline)
rd.compute_list_bind_uniform_set(compute_list, compute_base_uniform_set, 0)
rd.compute_list_bind_uniform_set(compute_list, dilate_uniform_set, 1)

for i in atlas_slices:
	rd.compute_list_set_push_constant(compute_list, push_constant, push_constant.size())
	rd.compute_list_dispatch(compute_list, group_size.x, group_size.y, group_size.z)
	# No barrier, let them run all together.

rd.compute_list_end()
*/

package main

import (
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Vector3i"
)

var (
	compute_shader_dilate_pipeline RID.ComputePipeline
	compute_base_uniform_set       RID.UniformSet
	dilate_uniform_set             RID.UniformSet
	atlas_slices                   []int
	push_constant                  []byte
	group_size                     Vector3i.XYZ
)

func RenderingDevice_ComputeListBegin() {
	var compute_list = rd.ComputeListBegin()
	rd.ComputeListBindComputePipeline(compute_list, compute_shader_dilate_pipeline)
	rd.ComputeListBindUniformSet(compute_list, compute_base_uniform_set, 0)
	rd.ComputeListBindUniformSet(compute_list, dilate_uniform_set, 1)
	for range atlas_slices {
		rd.ComputeListSetPushConstant(compute_list, push_constant, len(push_constant))
		rd.ComputeListDispatch(compute_list, int(group_size.X), int(group_size.Y), int(group_size.Z))
		// No barrier, let them run all together.
	}
	rd.ComputeListEnd()
}
