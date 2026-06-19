/*
[gdscript]
var rd = RenderingDevice.new()
assert(rd.has_feature(RenderingDevice.SUPPORTS_RAYTRACING_PIPELINE))

# Create a BLAS for a mesh.
var geometry = RDAccelerationStructureGeometry.new()
geometry.flags = RenderingDevice.ACCELERATION_STRUCTURE_GEOMETRY_OPAQUE_BIT
geometry.vertex_buffer = vertex_buffer
geometry.vertex_stride = 12
geometry.vertex_format = RenderingDevice.DATA_FORMAT_R32G32B32_SFLOAT
geometry.vertex_count = 3
geometry.index_buffer = index_buffer
geometry.index_count = 3
geometries.push_back(geometry)

blas = rd.blas_create([geometry], 0)

# Create TLAS.
tlas = rd.tlas_create(1, 0)

# Build acceleration structures.
rd.blas_build(blas)

var instance = RDAccelerationStructureInstance.new()
instance.blas = blas

instance.hit_sbt_range = rd.hit_sbt_range_alloc(hit_sbt, 1)
rd.hit_sbt_range_update(hit_sbt, instance.hit_sbt_range, 0, [0])

rd.tlas_build(tlas, [instance])

var raylist = rd.raytracing_list_begin()

# Bind pipeline and uniforms.
rd.raytracing_list_bind_raytracing_pipeline(raylist, raytracing_pipeline)
rd.raytracing_list_bind_uniform_set(raylist, uniform_set, 0)

# Trace rays.
var width = get_viewport().size.x
var height = get_viewport().size.y
rd.raytracing_list_trace_rays(raylist, 0, hit_sbt, width, height, 1)

rd.raytracing_list_end()
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/RDAccelerationStructureGeometry"
	"graphics.gd/classdb/RDAccelerationStructureInstance"
	"graphics.gd/classdb/Rendering"
	"graphics.gd/classdb/RenderingDevice"
	"graphics.gd/variant/RID"
)

func ExampleRaytracingList(
	rd RenderingDevice.Instance,
	vertexBuffer RID.VertexBuffer, indexBuffer RID.IndexBuffer,
	hitSbt RID.ShaderBindingTable, raytracingPipeline RID.RaytracingPipeline,
	uniformSet RID.UniformSet, width, height int,
) {
	if !rd.HasFeature(Rendering.SupportsRaytracingPipeline) {
		return
	}

	// Create a BLAS for a mesh.
	var geometry = RDAccelerationStructureGeometry.New()
	geometry.SetFlags(Rendering.AccelerationStructureGeometryOpaqueBit)
	geometry.SetVertexBuffer(RID.Any(vertexBuffer))
	geometry.SetVertexStride(12)
	geometry.SetVertexFormat(Rendering.DataFormatR32g32b32Sfloat)
	geometry.SetVertexCount(3)
	geometry.SetIndexBuffer(RID.Any(indexBuffer))
	geometry.SetIndexCount(3)

	var blas = rd.BlasCreate([]RDAccelerationStructureGeometry.Instance{geometry}, 0)

	// Create TLAS.
	var tlas = rd.TlasCreate(1, 0)

	// Build acceleration structures.
	rd.BlasBuild(blas)

	var instance = RDAccelerationStructureInstance.New()
	instance.SetBlas(RID.Any(blas))

	var hitSbtRange = rd.HitSbtRangeAlloc(hitSbt, 1)
	instance.SetHitSbtRange(hitSbtRange)
	rd.HitSbtRangeUpdate(hitSbt, hitSbtRange, 0, []int32{0})

	rd.TlasBuild(tlas, []RDAccelerationStructureInstance.Instance{instance})

	var raylist = rd.RaytracingListBegin()

	// Bind pipeline and uniforms.
	rd.RaytracingListBindRaytracingPipeline(raylist, raytracingPipeline)
	rd.RaytracingListBindUniformSet(raylist, uniformSet, 0)

	// Trace rays.
	rd.RaytracingListTraceRays(raylist, 0, hitSbt, width, height, 1)

	rd.RaytracingListEnd()
}
