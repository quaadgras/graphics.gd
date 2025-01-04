package texture

import (
	"graphics.gd/shaders/internal/dsl"
	"graphics.gd/shaders/vec1"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

type expression = dsl.Expression

type Sampler2D[T float64 | int | uint] struct {
	expression
}

func (s Sampler2D[T]) TextureSize(lod vec1.T[int]) vec2.T[int] {
	return vec2.T[int]{Expression: dsl.Fn("textureSize", s, lod)}
}

func (s Sampler2D[T]) QueryLOD(p vec2.XY) vec2.XY {
	return vec2.XY{Expression: dsl.Fn("textureQueryLod", s, p)}
}

func (s Sampler2D[T]) QueryLevels() vec1.X {
	return vec1.X{Expression: dsl.Fn("textureQueryLevels", s)}
}

func (s Sampler2D[T]) Texture(uv vec2.XY) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("texture", s, uv)}
}

func (s Sampler2D[T]) TextureWithBias(uv vec2.XY, bias vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("texture", s, uv, bias)}
}

func (s Sampler2D[T]) TextureProj(uv vec3.XYZ) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureProj", s, uv)}
}

func (s Sampler2D[T]) TextureLOD(uv vec2.XY, lod vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureLod", s, uv, lod)}
}

func (s Sampler2D[T]) TextureProjLoad(uv vec3.XYZ) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureProj", s, uv)}
}

func (s Sampler2D[T]) TextureGrad(uv vec2.XY, dPdx, dPdy vec2.XY) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureGrad", s, uv, dPdx, dPdy)}
}

func (s Sampler2D[T]) TextureProjGrad(uv vec3.XYZ, dPdx, dPdy vec3.XYZ) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureProjGrad", s, uv, dPdx, dPdy)}
}

func (s Sampler2D[T]) TexelFetch(p vec2.XY, lod vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("texelFetch", s, p, lod)}
}

func (s Sampler2D[T]) TextureGather(p vec2.XY, comp vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureGather", s, p, comp)}
}

type Sampler3D[T float64 | int | uint] struct {
	expression
}

func (s Sampler3D[T]) Texture(uv vec3.XYZ) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("texture", s, uv)}
}

func (s Sampler3D[T]) TextureSize(lod vec1.T[int]) vec3.T[int] {
	return vec3.T[int]{Expression: dsl.Fn("textureSize", s, lod)}
}

func (s Sampler3D[T]) QueryLOD(p vec2.XY) vec2.XY {
	return vec2.XY{Expression: dsl.Fn("textureQueryLod", s, p)}
}

func (s Sampler3D[T]) QueryLevels() vec1.T[int] {
	return vec1.T[int]{Expression: dsl.Fn("textureQueryLevels", s)}
}

func (s Sampler3D[T]) TextureProj(uv vec4.XYZW) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureProj", s, uv)}
}

func (s Sampler3D[T]) TextureLOD(uv vec3.XYZ, lod vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureLod", s, uv, lod)}
}

func (s Sampler3D[T]) TextureProjLoad(uv vec4.XYZW) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureProj", s, uv)}
}

func (s Sampler3D[T]) TextureGrad(uv vec3.XYZ, dPdx, dPdy vec3.XYZ) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureGrad", s, uv, dPdx, dPdy)}
}

func (s Sampler3D[T]) TextureProjGrad(uv vec4.XYZW, dPdx, dPdy vec4.XYZW) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureProjGrad", s, uv, dPdx, dPdy)}
}

func (s Sampler3D[T]) TexelFetch(p vec3.XYZ, lod vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("texelFetch", s, p, lod)}
}

type ArraySampler2D[T float64 | int | uint] struct {
	expression
}

func (s ArraySampler2D[T]) TextureSize(lod vec1.T[int]) vec3.T[int] {
	return vec3.T[int]{Expression: dsl.Fn("textureSize", s, lod)}
}

func (s ArraySampler2D[T]) QueryLOD(p vec2.XY) vec3.XYZ {
	return vec3.XYZ{Expression: dsl.Fn("textureQueryLod", s, p)}
}

func (s ArraySampler2D[T]) QueryLevels() vec1.T[int] {
	return vec1.T[int]{Expression: dsl.Fn("textureQueryLevels", s)}
}

func (s ArraySampler2D[T]) Texture(uv vec3.XYZ) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("texture", s, uv)}
}

func (s ArraySampler2D[T]) TextureLOD(uv vec3.XYZ, lod vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureLod", s, uv, lod)}
}

func (s ArraySampler2D[T]) TextureGrad(uv vec3.XYZ, dPdx, dPdy vec3.XYZ) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureGrad", s, uv, dPdx, dPdy)}
}

func (s ArraySampler2D[T]) TexelFetch(p vec3.XYZ, lod vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("texelFetch", s, p, lod)}
}

func (s ArraySampler2D[T]) TextureGather(p vec2.XY, comp vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureGather", s, p, comp)}
}

type ArraySampler3D[T float64 | int | uint] struct {
	expression
}

type SamplerCube struct {
	expression
}

func (s SamplerCube) TextureSize(lod vec1.T[int]) vec2.T[int] {
	return vec2.T[int]{Expression: dsl.Fn("textureSize", s, lod)}
}

func (s SamplerCube) QueryLOD(p vec2.XY) vec2.XY {
	return vec2.XY{Expression: dsl.Fn("textureQueryLod", s, p)}
}

func (s SamplerCube) QueryLevels() vec1.T[int] {
	return vec1.T[int]{Expression: dsl.Fn("textureQueryLevels", s)}
}

func (s SamplerCube) Texture(uv vec3.XYZ) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("texture", s, uv)}
}

func (s SamplerCube) TextureLOD(uv vec3.XYZ, lod vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureLod", s, uv, lod)}
}

func (s SamplerCube) TextureGrad(uv vec3.XYZ, dPdx, dPdy vec3.XYZ) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureGrad", s, uv, dPdx, dPdy)}
}

func (s SamplerCube) TextureGather(p vec2.XY, comp vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureGather", s, p, comp)}
}

type SamplerCubeArray struct {
	expression
}

func (s SamplerCubeArray) TextureSize(lod vec1.T[int]) vec2.T[int] {
	return vec2.T[int]{Expression: dsl.Fn("textureSize", s, lod)}
}

func (s SamplerCubeArray) Texture(uv vec4.XYZW) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("texture", s, uv)}
}

func (s SamplerCubeArray) TextureLOD(uv vec4.XYZW, lod vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureLod", s, uv, lod)}
}

func (s SamplerCubeArray) TextureGrad(uv vec4.XYZW, dPdx, dPdy vec4.XYZW) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Fn("textureGrad", s, uv, dPdx, dPdy)}
}
