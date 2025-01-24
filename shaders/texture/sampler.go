package texture

import (
	"graphics.gd/shaders/internal/gpu"
)

type expression = gpu.Expression

type AnyData interface {
	~gpu.RGBA | ~gpu.Vec4 | ~gpu.Vec4i | ~gpu.Vec4u
}

type Sampler2D[T AnyData] struct {
	expression
}

func (s Sampler2D[T]) TextureSize(lod gpu.Int) gpu.Vec2i { //glsl:textureSize(sampler2D,int)ivec2
	return gpu.NewVec2iExpression(gpu.Fn("textureSize", s, lod))
}
func (s Sampler2D[T]) Texture(uv gpu.Vec2) T { //glsl:texture(sampler2D,vec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texture", s, uv))
}
func (s Sampler2D[T]) TextureBias(uv gpu.Vec2, bias gpu.Float) T { //glsl:texture(sampler2D,vec2,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texture", s, uv, bias))
}
func (s Sampler2D[T]) TextureProj3(uv gpu.Vec3) T { //glsl:textureProj(sampler2D,vec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProj", s, uv))
}
func (s Sampler2D[T]) TextureProj3Bias(uv gpu.Vec3, bias gpu.Float) T { //glsl:textureProj(sampler2D,vec3,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProj", s, uv, bias))
}
func (s Sampler2D[T]) TextureProj4(uv gpu.Vec4) T { //glsl:textureProj(sampler2D,vec4)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProj", s, uv))
}
func (s Sampler2D[T]) TextureProj4Bias(uv gpu.Vec4, bias gpu.Float) T { //glsl:textureProj(sampler2D,vec4,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProj", s, uv, bias))
}
func (s Sampler2D[T]) TextureLOD(uv gpu.Vec2, lod gpu.Float) T { //glsl:textureLod(sampler2D,vec2,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureLod", s, uv, lod))
}
func (s Sampler2D[T]) TextureOffsetLOD(uv gpu.Vec2, offset gpu.Vec2i, lod gpu.Float) T { //glsl:textureLodOffset(sampler2D,vec2,float,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureLodOffset", s, uv, lod, offset))
}
func (s Sampler2D[T]) TextureGrad(uv, dPdx, dPdy gpu.Vec2) T { //glsl:textureGrad(sampler2D,vec2,vec2,vec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureGrad", s, uv, dPdx, dPdy))
}
func (s Sampler2D[T]) TextureGradOffset(uv, dPdx, dPdy gpu.Vec2, offset gpu.Vec2i) T { //glsl:textureGradOffset(sampler2D,vec2,vec2,vec2,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureGradOffset", s, uv, dPdx, dPdy, offset))
}
func (s Sampler2D[T]) TextureOffset(uv gpu.Vec2, offset gpu.Vec2i) T { //glsl:textureOffset(sampler2D,vec2,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureOffset", s, uv, offset))
}
func (s Sampler2D[T]) TextureOffsetBias(uv gpu.Vec2, offset gpu.Vec2i, bias gpu.Float) T { //glsl:textureOffset(sampler2D,vec2,ivec2,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureOffset", s, uv, offset, bias))
}
func (s Sampler2D[T]) TextureProj3Grad(uv gpu.Vec3, dPdx, dPdy gpu.Vec2) T { //glsl:textureProjGrad(sampler2D,vec3,vec2,vec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjGrad", s, uv, dPdx, dPdy))
}
func (s Sampler2D[T]) TextureProj4Grad(uv gpu.Vec4, dPdx, dPdy gpu.Vec2) T { //glsl:textureProjGrad(sampler2D,vec4,vec2,vec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjGrad", s, uv, dPdx, dPdy))
}
func (s Sampler2D[T]) TextureProj3GradOffset(uv gpu.Vec3, dPdx, dPdy gpu.Vec2, offset gpu.Vec2i) T { //glsl:textureProjGradOffset(sampler2D,vec3,vec2,vec2,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjGradOffset", s, uv, dPdx, dPdy, offset))
}
func (s Sampler2D[T]) TextureProj4GradOffset(uv gpu.Vec4, dPdx, dPdy gpu.Vec2, offset gpu.Vec2i) T { //glsl:textureProjGradOffset(sampler2D,vec4,vec2,vec2,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjGradOffset", s, uv, dPdx, dPdy, offset))
}
func (s Sampler2D[T]) TextureProj3LOD(uv gpu.Vec3, lod gpu.Float) T { //glsl:textureProjLod(sampler2D,vec3,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjLod", s, uv, lod))
}
func (s Sampler2D[T]) TextureProj4LOD(uv gpu.Vec4, lod gpu.Float) T { //glsl:textureProjLod(sampler2D,vec4,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjLod", s, uv, lod))
}
func (s Sampler2D[T]) TextureProj3LODOffset(uv gpu.Vec3, lod gpu.Float, offset gpu.Vec2i) T { //glsl:textureProjLodOffset(sampler2D,vec3,float,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjLodOffset", s, uv, lod, offset))
}
func (s Sampler2D[T]) TextureProj4LODOffset(uv gpu.Vec4, lod gpu.Float, offset gpu.Vec2i) T { //glsl:textureProjLodOffset(sampler2D,vec4,float,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjLodOffset", s, uv, lod, offset))
}
func (s Sampler2D[T]) TextureProj3Offset(uv gpu.Vec3, offset gpu.Vec2i) T { //glsl:textureProjOffset(sampler2D,vec3,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjOffset", s, uv, offset))
}
func (s Sampler2D[T]) TextureProj4Offset(uv gpu.Vec4, offset gpu.Vec2i) T { //glsl:textureProjOffset(sampler2D,vec4,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjOffset", s, uv, offset))
}
func (s Sampler2D[T]) TextureProj3OffsetBias(uv gpu.Vec3, offset gpu.Vec2i, bias gpu.Float) T { //glsl:textureProjOffset(sampler2D,vec3,ivec2,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjOffset", s, uv, offset, bias))
}
func (s Sampler2D[T]) TextureProj4OffsetBias(uv gpu.Vec4, offset gpu.Vec2i, bias gpu.Float) T { //glsl:textureProjOffset(sampler2D,vec4,ivec2,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjOffset", s, uv, offset, bias))
}

func (s Sampler2D[T]) TexelFetch(p gpu.Vec2i, lod gpu.Int) T { //glsl:texelFetch(sampler2D,ivec2,int)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texelFetch", s, p, lod))
}

func (s Sampler2D[T]) TexelFetchOffset(p gpu.Vec2i, lod gpu.Int, offset gpu.Vec2i) T { //glsl:texelFetchOffset(sampler2D,ivec2,int,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texelFetchOffset", s, p, lod, offset))
}

type Sampler3D[T AnyData] struct {
	expression
}

func (s Sampler3D[T]) TextureSize(lod gpu.Int) gpu.Vec3i { //glsl:textureSize(sampler3D,int)ivec3
	return gpu.NewVec3iExpression(gpu.Fn("textureSize", s, lod))
}
func (s Sampler3D[T]) Texture(uv gpu.Vec3) T { //glsl:texture(sampler3D,vec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texture", s, uv))
}
func (s Sampler3D[T]) TextureBias(uv gpu.Vec3, bias gpu.Float) T { //glsl:texture(sampler3D,vec3,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texture", s, uv, bias))
}
func (s Sampler3D[T]) TextureGrad(uv gpu.Vec3, dPdx, dPdy gpu.Vec3) T { //glsl:textureGrad(sampler3D,vec3,vec3,vec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureGrad", s, uv, dPdx, dPdy))
}
func (s Sampler3D[T]) TextureGradOffset(uv gpu.Vec3, dPdx, dPdy gpu.Vec3, offset gpu.Vec3i) T { //glsl:textureGradOffset(sampler3D,vec3,vec3,vec3,ivec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureGradOffset", s, uv, dPdx, dPdy, offset))
}
func (s Sampler3D[T]) TextureLOD(uv gpu.Vec3, lod gpu.Float) T { //glsl:textureLod(sampler3D,vec3,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureLod", s, uv, lod))
}
func (s Sampler3D[T]) TextureOffsetLOD(uv gpu.Vec3, offset gpu.Vec3i, lod gpu.Float) T { //glsl:textureLodOffset(sampler3D,vec3,float,ivec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureLodOffset", s, uv, lod, offset))
}
func (s Sampler3D[T]) TextureOffset(uv gpu.Vec3, offset gpu.Vec3i) T { //glsl:textureOffset(sampler3D,vec3,ivec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureOffset", s, uv, offset))
}
func (s Sampler3D[T]) TextureOffsetBias(uv gpu.Vec3, offset gpu.Vec3i, bias gpu.Float) T { //glsl:textureOffset(sampler3D,vec3,ivec3,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureOffset", s, uv, offset, bias))
}
func (s Sampler3D[T]) TextureProj(uv gpu.Vec4) T { //glsl:textureProj(sampler3D,vec4)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProj", s, uv))
}
func (s Sampler3D[T]) TextureProjBias(uv gpu.Vec4, bias gpu.Float) T { //glsl:textureProj(sampler3D,vec4,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProj", s, uv, bias))
}
func (s Sampler3D[T]) TextureProjGrad(uv gpu.Vec4, dPdx, dPdy gpu.Vec3) T { //glsl:textureProjGrad(sampler3D,vec4,vec3,vec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjGrad", s, uv, dPdx, dPdy))
}
func (s Sampler3D[T]) TextureProjGradOffset(uv gpu.Vec4, dPdx, dPdy gpu.Vec3, offset gpu.Vec3i) T { //glsl:textureProjGradOffset(sampler3D,vec4,vec3,vec3,ivec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjGradOffset", s, uv, dPdx, dPdy, offset))
}
func (s Sampler3D[T]) TextureProjLOD(uv gpu.Vec4, lod gpu.Float) T { //glsl:textureProjLod(sampler3D,vec4,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjLod", s, uv, lod))
}
func (s Sampler3D[T]) TextureProjLODOffset(uv gpu.Vec4, lod gpu.Float, offset gpu.Vec3i) T { //glsl:textureProjLodOffset(sampler3D,vec4,float,ivec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjLodOffset", s, uv, lod, offset))
}
func (s Sampler3D[T]) TextureProjOffset(uv gpu.Vec4, offset gpu.Vec3i) T { //glsl:textureProjOffset(sampler3D,vec4,ivec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjOffset", s, uv, offset))
}
func (s Sampler3D[T]) TextureProjOffsetBias(uv gpu.Vec4, offset gpu.Vec3i, bias gpu.Float) T { //glsl:textureProjOffset(sampler3D,vec4,ivec3,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureProjOffset", s, uv, offset, bias))
}

func (s Sampler3D[T]) TexelFetch(p gpu.Vec3i, lod gpu.Int) T { //glsl:texelFetch(sampler3D,ivec3,int)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texelFetch", s, p, lod))
}

func (s Sampler3D[T]) TexelFetchOffset(p gpu.Vec3i, lod gpu.Int, offset gpu.Vec3i) T { //glsl:texelFetchOffset(sampler3D,ivec3,int,ivec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texelFetchOffset", s, p, lod, offset))
}

type ArraySampler2D[T AnyData] struct {
	expression
}

func (s ArraySampler2D[T]) TextureSize(lod gpu.Int) gpu.Vec3i { //glsl:textureSize(sampler2DArray,int)ivec3
	return gpu.NewVec3iExpression(gpu.Fn("textureSize", s, lod))
}
func (s ArraySampler2D[T]) Texture(uv gpu.Vec3) T { //glsl:texture(sampler2DArray,vec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texture", s, uv))
}
func (s ArraySampler2D[T]) TextureWithBias(uv gpu.Vec3, bias gpu.Float) T { //glsl:texture(sampler2DArray,vec3,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texture", s, uv, bias))
}
func (s ArraySampler2D[T]) TextureGrad(uv gpu.Vec3, dPdx, dPdy gpu.Vec2) T { //glsl:textureGrad(sampler2DArray,vec3,vec2,vec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureGrad", s, uv, dPdx, dPdy))
}
func (s ArraySampler2D[T]) TextureGradOffset(uv gpu.Vec3, dPdx, dPdy gpu.Vec2, offset gpu.Vec2i) T { //glsl:textureGradOffset(sampler2DArray,vec3,vec2,vec2,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureGradOffset", s, uv, dPdx, dPdy, offset))
}
func (s ArraySampler2D[T]) TextureLOD(uv gpu.Vec3, lod gpu.Float) T { //glsl:textureLod(sampler2DArray,vec3,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureLod", s, uv, lod))
}
func (s ArraySampler2D[T]) TextureOffsetLOD(uv gpu.Vec3, offset gpu.Vec2i, lod gpu.Float) T { //glsl:textureLodOffset(sampler2DArray,vec3,float,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureLodOffset", s, uv, lod, offset))
}

func (s ArraySampler2D[T]) TexelFetch(p gpu.Vec3i, lod gpu.Int) T { //glsl:texelFetch(sampler2DArray,ivec3,int)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texelFetch", s, p, lod))
}

func (s ArraySampler2D[T]) TexelFetchOffset(p gpu.Vec3i, lod gpu.Int, offset gpu.Vec2i) T { //glsl:texelFetchOffset(sampler2DArray,ivec3,int,ivec2)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texelFetchOffset", s, p, lod, offset))
}

type ArraySampler3D[T float64 | int | uint] struct {
	expression
}

type CubeSampler[T AnyData] struct {
	expression
}

func (s CubeSampler[T]) Texture(uv gpu.Vec3) T { //glsl:texture(samplerCube,vec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texture", s, uv))
}
func (s CubeSampler[T]) TextureWithBias(uv gpu.Vec3, bias gpu.Float) T { //glsl:texture(samplerCube,vec3,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("texture", s, uv, bias))
}
func (s CubeSampler[T]) TextureSize(lod gpu.Int) gpu.Vec2i { //glsl:textureSize(samplerCube,int)ivec2
	return gpu.NewVec2iExpression(gpu.Fn("textureSize", s, lod))
}
func (s CubeSampler[T]) TextureGrad(uv, dPdx, dPdy gpu.Vec3) T { //glsl:textureGrad(samplerCube,vec3,vec3,vec3)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureGrad", s, uv, dPdx, dPdy))
}
func (s CubeSampler[T]) TextureLOD(uv gpu.Vec3, lod gpu.Float) T { //glsl:textureLod(samplerCube,vec3,float)rgba
	return gpu.NewQuadExpression[T](gpu.Fn("textureLod", s, uv, lod))
}
