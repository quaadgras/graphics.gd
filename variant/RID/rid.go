package RID

import (
	"graphics.gd/classdb/Resource"
	gd "graphics.gd/internal"
)

// New allocates a unique ID which can be used by the implementation to construct an RID.
// This is used mainly from native extensions to implement servers.
func New() Any { //gd:rid_allocate_id
	return gd.RidFromInt64(gd.RidAllocateId())
}

type Any = Resource.ID

// Int64 returns a resource ID from the given int64.
func Int64(id int64) Any { //gd:rid_from_int64
	return gd.RidFromInt64(gd.Int(id))
}

type Either[A, B ~uint64] Any

type (
	Mesh                             Any
	Material                         Any
	CameraAttributes                 Any
	Camera                           Any
	Compositor                       Any
	Environment                      Any
	Canvas                           Any
	CanvasItem                       Any
	MultiMesh                        Any
	Texture                          Any
	Particles                        Any
	Shape2D                          Any
	Shape3D                          Any
	Space2D                          Any
	Space3D                          Any
	Font                             Any
	Body2D                           Any
	Body3D                           Any
	SoftBody3D                       Any
	Area2D                           Any
	Area3D                           Any
	Joint2D                          Any
	Joint3D                          Any
	NativeMenu                       Any
	Framebuffer                      Any
	NavigationMap2D                  Any
	NavigationAgent2D                Any
	NavigationLink2D                 Any
	NavigationRegion2D               Any
	NavigationObstacle2D             Any
	NavigationSourceGeometryParser2D Any
	NavigationMap3D                  Any
	NavigationAgent3D                Any
	NavigationLink3D                 Any
	NavigationRegion3D               Any
	NavigationObstacle3D             Any
	NavigationSourceGeometryParser3D Any
	UniformBuffer                    Any
	Sampler                          Any
	VertexBuffer                     Any
	VertexArray                      Any
	IndexBuffer                      Any
	IndexArray                       Any
	Shader                           Any
	ShaderPlaceholder                Any
	StorageBuffer                    Any
	TextureBuffer                    Any
	UniformSet                       Any
	Buffer                           Any
	RenderPipeline                   Any
	ComputePipeline                  Any
	Texture2D                        Any
	Texture3D                        Any
	TextureProxy                     Any
	Skeleton                         Any
	Light                            Any
	Decal                            Any
	ReflectionProbe                  Any
	VoxelGI                          Any
	Lightmap                         Any
	ParticlesCollision               Any
	FogVolume                        Any
	VisibilityNotifier               Any
	Occluder                         Any
	Viewport                         Any
	Scenario                         Any
	Sky                              Any
	CompositorEffect                 Any
	ColorCorrection                  Any
	VisualInstance                   Any
	CanvasTexture                    Any
	CanvasLight                      Any
	CanvasLightOccluder              Any
	CanvasLightOccluderPolygon       Any
	TextBuffer                       Any
)
