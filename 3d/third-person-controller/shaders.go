package main

import (
	"graphics.gd/shaders"
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/mat4"
	"graphics.gd/shaders/pipeline/CanvasItem"
	"graphics.gd/shaders/pipeline/Sky"
	"graphics.gd/shaders/pipeline/Spatial"
	"graphics.gd/shaders/rgb"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/swizzle"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

// Shaders ported from the .gdshader files in the original demo.

type sampler = texture.Sampler2D[vec4.XYZW]

// xyz returns the first three components of v.
func xyz(v vec4.XYZW) vec3.XYZ { return vec3.New(v.X, v.Y, v.Z) }

// color returns v as a color.
func color(v vec3.XYZ) vec3.RGB { return swizzle.RGB(v) }

// half maps x from the range -1 to 1 into the range 0 to 1.
func half(x float.X) float.X { return float.Div(float.Add(1.0, x), 2.0) }

// fresnel is 0 where the normal faces the view and 1 at grazing angles.
func fresnel(normal, view vec3.XYZ, amount float.X) float.X {
	return float.Pow(float.Sub(1.0, float.Clamp(vec3.Dot(normal, view), 0.0, 1.0)), amount)
}

// BeeWings glows in the pattern of the wings, pulsing over time.
type BeeWings struct {
	Spatial.Shader[BeeWings]

	WingPatternSampler sampler  `gd:",repeat_disable"`
	EmissionColor      vec3.RGB `gd:",source_color"`
	GlobalIntensity    float.X
}

func NewBeeWings() *BeeWings {
	return &BeeWings{EmissionColor: rgb.New(1.0, 1.0, 1.0), GlobalIntensity: float.New(1.0)}
}

func (*BeeWings) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.CullDisabled}
}

func (s *BeeWings) Material(frag Spatial.Fragment) Spatial.Material {
	pattern := s.WingPatternSampler.Sample(frag.UV).X
	pulse := float.Add(half(float.Sin(float.Add(frag.UV.Y, shaders.Time))), 0.5)
	return Spatial.Material{
		Albedo:   rgb.New(0.0, 0.0, 0.0),
		Emission: rgb.Mul(rgb.Mul(rgb.Mul(s.EmissionColor, float.SmoothStep(0.51, 0.49, pattern)), pulse), s.GlobalIntensity),
	}
}

// Smoke holds what the smoke ball shaders have in common: the ball is pushed
// out by a voronoi texture.
type Smoke struct {
	VoronoiSampler   sampler
	Progress         float.X
	Smoothness       float.X
	Ease             float.X
	BaseScale        float.X
	DeformationScale float.X
	TextureOffset    float.X
}

func newSmoke() Smoke {
	return Smoke{
		Progress:         float.New(0.0),
		Smoothness:       float.New(0.01),
		Ease:             float.New(2.0),
		BaseScale:        float.New(1.0),
		DeformationScale: float.New(1.0),
		TextureOffset:    float.New(0.0),
	}
}

// SmokeFragment passes the voronoi value of each vertex to the fragment stage.
type SmokeFragment struct {
	Spatial.Fragment

	Voronoi float.X
}

func (s *Smoke) deform(vertex Spatial.Vertex) SmokeFragment {
	deformation := float.Mul(s.DeformationScale, float.Sub(1.0, float.Mul(vec2.Distance(vertex.UV, vec2.New(0.5, 0.5)), 2.0)))
	voronoi := float.X(s.VoronoiSampler.Sample(vec2.Add(vertex.UV, vec2.New(0.0, s.TextureOffset))).X)
	push := float.Mul(float.Sub(1.0, float.Pow(voronoi, s.Ease)), deformation)
	return SmokeFragment{
		Fragment: Spatial.Fragment{
			VertexPosition: vec3.Mul(vec3.Add(vertex.Position, vec3.Mul(vertex.Normal, push)), s.BaseScale),
		},
		Voronoi: voronoi,
	}
}

// SmokeBall dissolves as progress goes to 1.
type SmokeBall struct {
	Spatial.Shader[SmokeBall]
	Smoke
}

func NewSmokeBall() *SmokeBall { return &SmokeBall{Smoke: newSmoke()} }

func (*SmokeBall) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.ShadowsDisabled, Spatial.CullDisabled}
}

func (s *SmokeBall) Fragment(vertex Spatial.Vertex) SmokeFragment { return s.deform(vertex) }

func (s *SmokeBall) Material(frag SmokeFragment) Spatial.Material {
	return Spatial.Material{
		Alpha:                 float.SmoothStep(float.Sub(s.Progress, s.Smoothness), float.Add(s.Progress, s.Smoothness), frag.Voronoi),
		AlphaScizzorThreshold: float.New(0.5),
	}
}

// SmokeBallMiddle fades out towards its silhouette.
type SmokeBallMiddle struct {
	Spatial.Shader[SmokeBallMiddle]
	Smoke

	FresnelOffset float.X
	GlobalAlpha   float.X
}

func NewSmokeBallMiddle() *SmokeBallMiddle {
	return &SmokeBallMiddle{Smoke: newSmoke(), FresnelOffset: float.New(1.0), GlobalAlpha: float.New(1.0)}
}

func (*SmokeBallMiddle) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.ShadowsDisabled, Spatial.CullBack}
}

func (s *SmokeBallMiddle) Fragment(vertex Spatial.Vertex) Spatial.Fragment {
	return s.deform(vertex).Fragment
}

func (s *SmokeBallMiddle) Material(frag Spatial.Fragment) Spatial.Material {
	return Spatial.Material{
		Alpha: float.Mul(float.Sub(1.0, fresnel(frag.Normal, frag.View(), float.Mul(2.0, s.FresnelOffset))), s.GlobalAlpha),
	}
}

// BushWind sways the leaves of bushes and trees in the wind.
type BushWind struct {
	Spatial.Shader[BushWind]

	WindNoiseSampler sampler
	AlbedoSampler    sampler `gd:",source_color"`
	OrmSampler       sampler
	NormalSampler    sampler
}

func (*BushWind) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.CullDisabled}
}

func (s *BushWind) Fragment(vertex Spatial.Vertex) Spatial.Fragment {
	pos, node := vertex.Position, vertex.NodePositionWorld
	uv := vec2.Add(vec2.Add(vec2.New(node.X, node.Z), float.Add(vertex.UV.Y, pos.Y)), vec2.New(0.0, float.Mul(shaders.Time, 0.2)))
	sway := float.Mul(float.Mul(s.WindNoiseSampler.Sample(uv).X, 0.05), vertex.UV.Y)
	return Spatial.Fragment{
		VertexPosition: vec3.New(float.Add(pos.X, sway), pos.Y, float.Add(pos.Z, sway)),
	}
}

func (s *BushWind) Material(frag Spatial.Fragment) Spatial.Material {
	albedo := s.AlbedoSampler.Sample(frag.UV)
	orm := s.OrmSampler.Sample(frag.UV)
	return Spatial.Material{
		Albedo:                color(vec3.Mul(xyz(albedo), orm.X)),
		Roughness:             float.X(orm.Y),
		NormalMap:             xyz(s.NormalSampler.Sample(frag.UV)),
		Alpha:                 float.X(albedo.W),
		AlphaScizzorThreshold: float.New(0.5),
	}
}

// FlowerPatch sways the flowers.
type FlowerPatch struct {
	Spatial.Shader[FlowerPatch]

	FlowerAlbedoSampler sampler `gd:",repeat_disable,source_color"`
	FlowerNormalSampler sampler `gd:",repeat_disable"`
}

func (*FlowerPatch) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.CullDisabled}
}

func (s *FlowerPatch) Fragment(vertex Spatial.Vertex) Spatial.Fragment {
	pos, node := vertex.Position, vertex.NodePositionWorld
	y := float.Sub(1.0, pos.Y)
	weight := float.Mul(float.Sub(1.0, y), 0.1)
	return Spatial.Fragment{
		VertexPosition: vec3.New(
			float.Add(pos.X, float.Mul(float.Sin(float.Add(float.Add(node.X, shaders.Time), y)), weight)),
			pos.Y,
			float.Add(pos.Z, float.Mul(float.Cos(float.Add(float.Add(node.Z, float.Mul(shaders.Time, 0.45)), y)), weight)),
		),
	}
}

func (s *FlowerPatch) Material(frag Spatial.Fragment) Spatial.Material {
	albedo := s.FlowerAlbedoSampler.Sample(frag.UV)
	return Spatial.Material{
		Albedo:                color(xyz(albedo)),
		NormalMap:             xyz(s.FlowerNormalSampler.Sample(frag.UV)),
		Alpha:                 float.X(albedo.W),
		AlphaScizzorThreshold: float.New(0.5),
	}
}

// Grass bends in the wind, lit as if it were flat ground.
type Grass struct {
	Spatial.Shader[Grass]

	WindNoiseSampler sampler
	Gradient         sampler `gd:",source_color"`
}

func (*Grass) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.CullDisabled, Spatial.WorldVertexCoords}
}

func (s *Grass) Fragment(vertex Spatial.Vertex) Spatial.Fragment {
	pos := vertex.Position
	uv := vec2.Add(vec2.Div(vec2.New(pos.X, pos.Z), 20.0), float.Mul(float.Mul(0.5, shaders.Time), 0.1))
	wind := float.Mul(float.Mul(float.Sub(s.WindNoiseSampler.Sample(uv).X, 0.5), 0.1), float.Sub(1.0, vertex.UV.Y))
	return Spatial.Fragment{
		Normal:         vec3.New(0.0, 1.0, 0.0),
		VertexPosition: vec3.New(float.Add(pos.X, wind), pos.Y, float.Add(pos.Z, wind)),
	}
}

func (s *Grass) Material(frag Spatial.Fragment) Spatial.Material {
	facing := float.Sub(float.Mul(float.FromBool(frag.FrontFacing()), 2.0), 1.0)
	return Spatial.Material{
		Albedo: color(xyz(s.Gradient.Sample(vec2.New(frag.UV.Y, frag.UV.X)))),
		Normal: vec3.Mul(frag.Normal, facing),
	}
}

// Waterfall scrolls voronoi noise down the waterfall, with foam at its edges.
type Waterfall struct {
	Spatial.Shader[Waterfall]

	VoronoiSampler   sampler
	WaterColorTop    vec3.RGB `gd:",source_color"`
	WaterColorBottom vec3.RGB `gd:",source_color"`
	DepthTexture     sampler  `gd:"DEPTH_TEXTURE,hint_depth_texture,filter_linear_mipmap"`
	ScreenTexture    sampler  `gd:"SCREEN_TEXTURE,hint_screen_texture,filter_linear_mipmap"`
}

func (s *Waterfall) Fragment(vertex Spatial.Vertex) Spatial.Fragment {
	voronoi := float.Sub(1.0, s.VoronoiSampler.Sample(vec2.Add(vertex.UV, vec2.New(float.Mul(shaders.Time, 0.1), shaders.Time))).X)
	return Spatial.Fragment{
		VertexPosition: vec3.Add(vertex.Position, vec3.Mul(vec3.Mul(vertex.Normal, voronoi), float.Sub(1.0, vertex.UV.Y))),
	}
}

func (s *Waterfall) Material(frag Spatial.Fragment) Spatial.Material {
	uv := vec2.Add(vec2.Mul(vec2.Mul(frag.UV, vec2.New(4.0, 1.0)), 0.5), vec2.New(float.Mul(shaders.Time, 0.1), shaders.Time))
	voronoi := float.X(s.VoronoiSampler.Sample(uv).X)
	edge := float.Sub(1.0, float.SmoothStep(0.84, 0.86, float.Add(voronoi, frag.UV.Y)))
	water := rgb.Mix(s.WaterColorBottom, s.WaterColorTop, frag.UV.Y)
	return Spatial.Material{
		Albedo:         rgb.Mix(water, rgb.New(1.0, 1.0, 1.0), edge),
		NormalMap:      vec3.Normalize(vec3.New(0.5, voronoi, 0.5)),
		NormalMapDepth: float.New(0.5),
		Roughness:      float.New(0.05),
	}
}

// JumpingPadScreen is the pixelated screen on top of the jumping pad.
type JumpingPadScreen struct {
	Spatial.Shader[JumpingPadScreen]

	ScreenColor vec3.RGB `gd:",source_color"`
	CrossSprite sampler  `gd:",source_color,filter_nearest"`
}

func NewJumpingPadScreen() *JumpingPadScreen {
	return &JumpingPadScreen{ScreenColor: rgb.New(1.0, 1.0, 1.0)}
}

func (s *JumpingPadScreen) Material(frag Spatial.Fragment) Spatial.Material {
	const res = 20.0
	pixel := vec2.Div(vec2.Round(vec2.Mul(frag.UV, res)), res)
	dist := vec2.Distance(vec2.New(0.5, 0.5), pixel)
	waves := half(float.Sin(float.Mul(float.Sub(dist, float.Mul(shaders.Time, 0.2)), 10.0)))
	grid2D := vec2.Sin(vec2.Mul(vec2.Fract(vec2.Add(vec2.Mul(frag.UV, res), 0.5)), 3.14))
	grid := float.Add(float.Mul(float.Mul(grid2D.X, grid2D.Y), 0.8), 0.2)
	scroll := vec2.New(float.Mod(float.Mul(float.Round(shaders.Time), 0.5), 1.0), 0.0)
	cross := float.Sub(1.0, float.Round(s.CrossSprite.Sample(vec2.Add(vec2.Mul(pixel, vec2.New(0.5, 1.0)), scroll)).X))
	return Spatial.Material{
		Albedo:   rgb.New(0.0, 0.0, 0.0),
		Emission: rgb.Mul(rgb.Mul(rgb.Mul(s.ScreenColor, float.Add(float.Mul(waves, 0.2), cross)), grid), 2.0),
	}
}

// BulletSpark is an animated spark, colored by its particle.
type BulletSpark struct {
	Spatial.Shader[BulletSpark]

	SparksSampler sampler
}

// SparkFragment passes the particle's values to the fragment stage.
type SparkFragment struct {
	Spatial.Fragment

	RandID        float.X
	Lifetime      float.X
	ParticleColor vec3.RGB
}

func (*BulletSpark) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.Unshaded}
}

func (s *BulletSpark) Fragment(vertex Spatial.Vertex) SparkFragment {
	// Random function from Patricio Gonzalez Vivo:
	// https://gist.github.com/patriciogonzalezvivo/670c22f3966e662d2f83
	random := float.Fract(float.Mul(float.Sin(float.FromInt(vertex.InstanceID)), 43758.5453123))
	return SparkFragment{
		RandID:        random,
		ParticleColor: swizzle.RGB(vertex.Color),
		Lifetime:      float.X(vertex.Color.A),
	}
}

func (s *BulletSpark) Material(frag SparkFragment) Spatial.Material {
	frame := float.Div(float.Round(float.Mod(float.Add(frag.RandID, float.Mul(shaders.Time, 8.0)), 1.0)), 2.0)
	mask := s.SparksSampler.Sample(vec2.Add(vec2.Mul(frag.UV, vec2.New(0.5, 1.0)), vec2.New(frame, 1.0))).X
	return Spatial.Material{
		Alpha:    float.Mul(mask, frag.Lifetime),
		Albedo:   frag.ParticleColor,
		Emission: rgb.Mul(frag.ParticleColor, 1.0),
	}
}

// BulletTrail is the wobbling trail behind a bullet.
type BulletTrail struct {
	Spatial.Shader[BulletTrail]

	NoiseSampler sampler
	Gradient1D   sampler `gd:"gradient_1D,source_color"`
}

func (s *BulletTrail) Fragment(vertex Spatial.Vertex) Spatial.Fragment {
	v := float.Sub(float.Sin(float.Mul(float.Add(float.Mul(shaders.Time, 4.0), vertex.UV.Y), 10.0)), 0.5)
	pos, normal := vertex.Position, vertex.Normal
	return Spatial.Fragment{
		VertexPosition: vec3.New(
			float.Add(pos.X, float.Mul(float.Mul(v, normal.X), 0.02)),
			pos.Y,
			float.Add(pos.Z, float.Mul(float.Mul(v, normal.Z), 0.02)),
		),
	}
}

func (s *BulletTrail) Material(frag Spatial.Fragment) Spatial.Material {
	n := float.X(s.NoiseSampler.Sample(vec2.Add(vec2.Mul(frag.UV, vec2.New(1.0, 0.2)), vec2.New(0.0, shaders.Time))).X)
	edge := float.SmoothStep(0.0, 0.8, frag.UV.Y)
	c := color(xyz(s.Gradient1D.Sample(vec2.New(n, 0.0))))
	return Spatial.Material{
		Alpha:     float.SmoothStep(0.0, 0.15, float.Mul(edge, float.Add(n, float.SmoothStep(0.6, 1.0, frag.UV.Y)))),
		Albedo:    c,
		Emission:  rgb.Mul(c, 1.4),
		Roughness: float.New(0.0),
		Specular:  float.New(0.0),
	}
}

// Blast is the billboarded shock wave of an explosion.
type Blast struct {
	Spatial.Shader[Blast]

	NoiseSampler           sampler
	Noise2Sampler          sampler
	UVYOffset              float.X `gd:"uv_y_offset"`
	UVYScale               float.X `gd:"uv_y_scale"`
	Coverage               float.X `gd:",hint_range(0.0, 1.0, 0.01)"`
	Intensity              float.X
	BlastGradientSampler1D sampler `gd:"blast_gradient_sampler_1D,source_color"`
}

func NewBlast() *Blast {
	return &Blast{UVYOffset: float.New(0.0), UVYScale: float.New(0.0), Coverage: float.New(0.5), Intensity: float.New(1.0)}
}

func (s *Blast) Fragment(vertex Spatial.Vertex) Spatial.Fragment {
	// Face the camera, keeping the position of the model.
	inv := vertex.InvViewMatrix
	billboard := mat4.New(mat4.Column(inv, 0), mat4.Column(inv, 1), mat4.Column(inv, 2), mat4.Column(vertex.ModelMatrix, 3))
	return Spatial.Fragment{
		ModelViewMatrix: mat4.Mul(vertex.ViewMatrix, billboard),
		VertexPosition:  vec3.Mul(vertex.Position, 5.0),
	}
}

func (s *Blast) Material(frag Spatial.Fragment) Spatial.Material {
	noise := s.NoiseSampler.Sample(vec2.Add(vec2.Mul(frag.UV, vec2.New(2.0, 0.01)), vec2.New(0.0, s.UVYOffset))).X
	noise2 := s.Noise2Sampler.Sample(vec2.Sub(vec2.Mul(frag.UV, vec2.New(1.0, 0.1)), vec2.New(0.0, float.Mul(shaders.Time, 0.1)))).X
	n := float.Mul(noise, noise2)
	edgeEase := float.Sin(float.Mul(float.Add(frag.UV.Y, s.UVYScale), shaders.PI))
	mask := float.Mul(
		float.SmoothStep(float.Sub(s.Coverage, 0.01), float.Add(s.Coverage, 0.01), float.Mul(n, edgeEase)),
		float.Sin(float.Mul(frag.UV.Y, shaders.PI)),
	)
	c := color(xyz(s.BlastGradientSampler1D.Sample(vec2.New(n, 0.0))))
	return Spatial.Material{
		Alpha:    mask,
		Albedo:   c,
		Emission: rgb.Mul(c, s.Intensity),
	}
}

// ExplosionTrails are the glowing trails of the explosion's particles.
type ExplosionTrails struct {
	Spatial.Shader[ExplosionTrails]

	NoiseSampler sampler
	Gradient1D   sampler `gd:"gradient_1D,source_color"`
}

func (*ExplosionTrails) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.ParticleTrails, Spatial.CullDisabled}
}

func (s *ExplosionTrails) Material(frag Spatial.Fragment) Spatial.Material {
	n := s.NoiseSampler.Sample(vec2.Sub(vec2.Mul(frag.UV, vec2.New(0.5, 1.0)), vec2.New(0.0, float.Mul(shaders.Time, 4.0)))).X
	edge := float.Mul(float.Sin(float.Mul(frag.UV.X, shaders.PI)), float.Sin(float.Mul(frag.UV.Y, shaders.PI)))
	clamped := float.SmoothStep(0.6, 0.2, float.Mul(frag.UV.Y, n))
	c := color(xyz(s.Gradient1D.Sample(vec2.New(clamped, 0.0))))
	return Spatial.Material{
		Alpha:    float.Mul(float.SmoothStep(0.2, 0.4, edge), clamped),
		Emission: rgb.Mul(c, 1.2),
	}
}

// Shell is the expanding shell of an explosion, dissolving as progress goes to 1.
type Shell struct {
	Spatial.Shader[Shell]

	Intensity      float.X
	ShellColor     vec4.RGBA `gd:",source_color"`
	VoronoiSampler sampler
	Edge           float.X
	Progress       float.X `gd:",hint_range(0.0, 1.0, 0.01)"`
}

func NewShell() *Shell { return &Shell{Edge: float.New(0.05), Progress: float.New(0.5)} }

func (*Shell) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.CullDisabled}
}

func (s *Shell) Material(frag Spatial.Fragment) Spatial.Material {
	voronoi := s.VoronoiSampler.Sample(vec2.Add(vec2.Mul(frag.UV, vec2.New(1.0, 0.6)), float.Mul(shaders.Time, 0.06))).X
	mask := float.SmoothStep(float.Sub(s.Progress, s.Edge), float.Add(s.Progress, s.Edge), voronoi)
	return Spatial.Material{
		Alpha:    float.Mul(mask, s.ShellColor.A),
		Albedo:   swizzle.RGB(s.ShellColor),
		Emission: rgb.Mul(swizzle.RGB(s.ShellColor), s.Intensity),
	}
}

// GrenadeTarget marks where a grenade will land.
type GrenadeTarget struct {
	Spatial.Shader[GrenadeTarget]

	TargetMaskSampler      sampler  `gd:",source_color,repeat_disable,filter_linear_mipmap_anisotropic"`
	TargetFrameMaskSampler sampler  `gd:",source_color,repeat_disable"`
	FillColor              vec3.RGB `gd:",source_color"`
}

func NewGrenadeTarget() *GrenadeTarget { return &GrenadeTarget{FillColor: rgb.New(1.0, 1.0, 1.0)} }

func (*GrenadeTarget) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.AmbientLightDisabled}
}

// rotateUV rotates uv around the center of the texture.
func rotateUV(uv vec2.XY, rotation float.X) vec2.XY {
	cos, sin := float.Cos(rotation), float.Sin(rotation)
	u, v := float.Sub(uv.X, 0.5), float.Sub(uv.Y, 0.5)
	return vec2.New(
		float.Add(float.Add(float.Mul(cos, u), float.Mul(sin, v)), 0.5),
		float.Add(float.Sub(float.Mul(cos, v), float.Mul(sin, u)), 0.5),
	)
}

func (s *GrenadeTarget) Material(frag Spatial.Fragment) Spatial.Material {
	dist := float.Sub(1.0, vec2.Distance(vec2.New(0.5, 0.5), frag.UV))
	circle := s.TargetMaskSampler.Sample(rotateUV(frag.UV, float.Mod(shaders.Time, shaders.PI))).X
	frame := s.TargetFrameMaskSampler.Sample(frag.UV).X
	ripple := float.Mul(float.Mul(half(float.Sin(float.Add(float.Mul(dist, 15.0), shaders.Time))), float.SmoothStep(0.5, 1.0, dist)), 0.1)
	return Spatial.Material{
		Albedo:   s.FillColor,
		Emission: rgb.Mul(s.FillColor, 1.0),
		Alpha:    float.Clamp(float.Add(float.Add(circle, frame), ripple), 0.0, 1.0),
	}
}

// GrenadeTrajectory is the arrowed path a grenade will follow.
type GrenadeTrajectory struct {
	Spatial.Shader[GrenadeTrajectory]

	ArrowSampler sampler  `gd:",source_color"`
	FillColor    vec3.RGB `gd:",source_color"`
	ArrowScale   float.X
}

func NewGrenadeTrajectory() *GrenadeTrajectory {
	return &GrenadeTrajectory{FillColor: rgb.New(1.0, 1.0, 1.0), ArrowScale: float.New(1.0)}
}

func (*GrenadeTrajectory) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.AmbientLightDisabled, Spatial.CullDisabled}
}

func (s *GrenadeTrajectory) Material(frag Spatial.Fragment) Spatial.Material {
	yEase := float.Sin(float.Mul(frag.UV.X, shaders.PI))
	arrowTime := float.Mod(shaders.Time, 1.0)
	bandTime := float.Mod(float.Mul(shaders.Time, 0.2), 1.0)
	arrows := float.Mul(yEase, s.ArrowSampler.Sample(vec2.Sub(vec2.Mul(frag.UV, vec2.New(s.ArrowScale, 1.0)), vec2.New(arrowTime, 0.0))).X)
	bandEase := float.Mul(yEase, float.Sin(float.Mul(frag.UV.Y, shaders.PI)))
	stripes := float.Add(float.Mul(half(float.Sin(float.Mul(float.Sub(float.Mul(frag.UV.X, 6.0), bandTime), shaders.TAU))), 0.1), 0.4)
	return Spatial.Material{
		Albedo:   s.FillColor,
		Emission: rgb.Add(s.FillColor, rgb.Mul(rgb.Mul(s.FillColor, arrows), 5.0)),
		Alpha:    float.Clamp(float.Add(float.Mul(bandEase, stripes), arrows), 0.0, 1.0),
	}
}

// Water is the lake, with caustics and foam where it meets the shore.
type Water struct {
	Spatial.Shader[Water]

	MainNoiseSampler   sampler  `gd:",hint_normal"`
	CausticSampler     sampler  `gd:",hint_normal,filter_linear"`
	DepthTexture       sampler  `gd:"DEPTH_TEXTURE,hint_depth_texture,filter_linear_mipmap"`
	ScreenTexture      sampler  `gd:"SCREEN_TEXTURE,hint_screen_texture,filter_linear_mipmap"`
	WaterColor         vec3.RGB `gd:",source_color"`
	WaterDepthColor    vec3.RGB `gd:",source_color"`
	WaterEdgeDistance  float.X
	WaterDepthDistance float.X
	WaterTextureScale  float.X
}

func NewWater() *Water {
	return &Water{
		WaterColor:         rgb.New(1.0, 1.0, 1.0),
		WaterDepthColor:    rgb.New(1.0, 1.0, 1.0),
		WaterEdgeDistance:  float.New(0.1),
		WaterDepthDistance: float.New(0.3),
		WaterTextureScale:  float.New(20.0),
	}
}

func (*Water) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.DepthDrawAlways}
}

func (s *Water) Fragment(vertex Spatial.Vertex) Spatial.Fragment {
	offset := vec2.Add(vec2.Mul(vertex.UV, 0.5), float.Mul(shaders.Time, 0.01))
	noise := func(offset vec2.XY) float.X { return float.X(s.MainNoiseSampler.Sample(offset).X) }
	h := noise(offset)
	pos := vertex.Position
	return Spatial.Fragment{
		VertexPosition: vec3.New(pos.X, float.Add(pos.Y, h), pos.Z),
		Normal: vec3.Normalize(vec3.New(
			float.Sub(h, noise(vec2.Add(offset, vec2.New(0.8, 0.0)))),
			8.0,
			float.Sub(h, noise(vec2.Add(offset, vec2.New(0.0, 0.8)))),
		)),
	}
}

func (s *Water) Material(frag Spatial.Fragment) Spatial.Material {
	caustic := float.X(s.CausticSampler.Sample(vec2.Add(vec2.Mul(frag.UV, s.WaterTextureScale), float.Mul(shaders.Time, 0.05))).X)
	finalCaustic := float.Mul(float.Mul(caustic, float.Sub(1.0, fresnel(frag.Normal, frag.View(), float.New(1.0)))), 0.25)

	// Position of whatever is behind the water, in view space.
	screenUV := frag.ScreenUV()
	depth := s.DepthTexture.SampleLOD(screenUV, float.New(0.0)).X
	behind := mat4.Transform(frag.InvProjectionMatrix(), vec4.New(float.Sub(float.Mul(screenUV.X, 2.0), 1.0), float.Sub(float.Mul(screenUV.Y, 2.0), 1.0), depth, 1.0))
	behindZ := float.Div(behind.Z, behind.W)

	z := frag.VertexPosition.Z
	waterEdge := float.Add(behindZ, s.WaterEdgeDistance)
	edgeMask := float.Clamp(float.Sub(1.0, float.SmoothStep(waterEdge, float.Add(waterEdge, float.Mul(caustic, 0.5)), z)), 0.0, 1.0)
	depthMask := float.Clamp(float.Sub(1.0, float.SmoothStep(behindZ, float.Add(behindZ, s.WaterDepthDistance), z)), 0.0, 1.0)

	screen := color(xyz(s.ScreenTexture.SampleBias(screenUV, float.New(2.0))))
	water := rgb.Mix(s.WaterColor, rgb.Mul(screen, s.WaterDepthColor), float.Mul(depthMask, 0.75))
	return Spatial.Material{
		Roughness: float.New(0.95),
		Specular:  float.New(0.5),
		Albedo:    rgb.Add(water, float.Clamp(float.Add(edgeMask, finalCaustic), 0.0, 1.0)),
		Alpha:     float.SmoothStep(0.5, 0.25, vec2.Distance(frag.UV, vec2.New(0.5, 0.5))),
	}
}

// Triplanar holds the varyings for projecting textures onto a mesh along each
// of its axes, weighted by how much the surface faces that axis.
type Triplanar struct {
	Spatial.Fragment

	TriplanarPowerNormal vec3.XYZ
	TriplanarPos         vec3.XYZ
	ObjectNormal         vec3.XYZ
}

func triplanar(vertex Spatial.Vertex, scale, offset vec3.XYZ, sharpness float.X) Triplanar {
	power := vec3.Pow(vec3.Abs(vertex.Normal), vec3.New(sharpness, sharpness, sharpness))
	power = vec3.Div(power, vec3.Dot(power, vec3.New(1.0, 1.0, 1.0)))
	pos := vec3.Mul(vec3.Add(vec3.Mul(vertex.Position, scale), offset), vec3.New(1.0, -1.0, 1.0))
	return Triplanar{TriplanarPowerNormal: power, TriplanarPos: pos, ObjectNormal: vertex.Normal}
}

// triplanarTexture samples s along each axis at pos, blended by weights.
func triplanarTexture(s sampler, weights, pos vec3.XYZ) vec4.XYZW {
	sample := s.Sample(vec2.New(pos.X, pos.Y))
	sample = vec4.Add(vec4.Mul(sample, weights.Z), vec4.Mul(s.Sample(vec2.New(pos.X, pos.Z)), weights.Y))
	return vec4.Add(sample, vec4.Mul(s.Sample(vec2.Mul(vec2.New(pos.Z, pos.Y), vec2.New(-1.0, 1.0))), weights.X))
}

// motifMask blends along an edge, broken up by a motif.
func motifMask(edge, motif float.X, top float64) float.X {
	negative := float.SmoothStep(1.0, 0.5, edge)
	return float.SmoothStep(0.2, top, float.Mul(float.Sub(1.0, float.Mul(negative, motif)), edge))
}

// Terrain blends cliffs, grass and dirt paths, projected along each axis.
type Terrain struct {
	Spatial.Shader[Terrain]

	GrassPatchPatternSampler sampler `gd:",filter_linear_mipmap,repeat_enable"`
	WallMaskSampler          sampler `gd:",filter_linear_mipmap,source_color,repeat_enable"`
	CliffAlbedoSampler       sampler `gd:",filter_linear_mipmap,source_color,repeat_enable"`
	GrassAlbedoSampler       sampler `gd:",filter_linear_mipmap,source_color,repeat_enable"`
	DirtAlbedoSampler        sampler `gd:",filter_linear_mipmap,source_color,repeat_enable"`
	CliffOrmSampler          sampler `gd:",filter_linear_mipmap,hint_roughness_g,repeat_enable"`
	GrassOrmSampler          sampler `gd:",filter_linear_mipmap,hint_roughness_g,repeat_enable"`
	DirtOrmSampler           sampler `gd:",filter_linear_mipmap,hint_roughness_g,repeat_enable"`
	CliffNormalSampler       sampler `gd:",filter_linear_mipmap,hint_normal,repeat_enable"`
	GrassNormalSampler       sampler `gd:",filter_linear_mipmap,hint_normal,repeat_enable"`
	DirtNormalSampler        sampler `gd:",filter_linear_mipmap,hint_normal,repeat_enable"`
	WallEdgeOffset           float.X
	CliffScale               float.X
	TriplanarScale           vec3.XYZ
	TriplanarOffset          vec3.XYZ
	TriplanarSharpness       float.X
}

func NewTerrain() *Terrain {
	return &Terrain{
		WallEdgeOffset:     float.New(1.0),
		CliffScale:         float.New(1.0),
		TriplanarScale:     vec3.New(1.0, 1.0, 1.0),
		TriplanarOffset:    vec3.New(0.0, 0.0, 0.0),
		TriplanarSharpness: float.New(0.5),
	}
}

func (s *Terrain) Fragment(vertex Spatial.Vertex) Triplanar {
	n := vec3.Abs(vertex.Normal)
	tangent := vec3.Add(vec3.Add(vec3.Mul(vec3.New(0.0, 0.0, -1.0), n.X), vec3.Mul(vec3.New(1.0, 0.0, 0.0), n.Y)), vec3.Mul(vec3.New(1.0, 0.0, 0.0), n.Z))
	binormal := vec3.Add(vec3.Add(vec3.Mul(vec3.New(0.0, 1.0, 0.0), n.X), vec3.Mul(vec3.New(0.0, 0.0, -1.0), n.Y)), vec3.Mul(vec3.New(0.0, 1.0, 0.0), n.Z))
	out := triplanar(vertex, s.TriplanarScale, s.TriplanarOffset, s.TriplanarSharpness)
	out.Tangent = vec3.Normalize(tangent)
	out.Binormal = vec3.Normalize(binormal)
	return out
}

func (s *Terrain) Material(frag Triplanar) Spatial.Material {
	w, pos := frag.TriplanarPowerNormal, frag.TriplanarPos
	cliffPos := vec3.Mul(pos, s.CliffScale)

	cliffAlbedo := xyz(triplanarTexture(s.CliffAlbedoSampler, w, cliffPos))
	grassAlbedo := xyz(triplanarTexture(s.GrassAlbedoSampler, w, pos))
	dirtAlbedo := xyz(triplanarTexture(s.DirtAlbedoSampler, w, pos))

	cliffRoughness := float.X(triplanarTexture(s.CliffOrmSampler, w, cliffPos).Y)
	grassRoughness := float.X(triplanarTexture(s.GrassOrmSampler, w, pos).Y)
	dirtRoughness := float.X(triplanarTexture(s.DirtOrmSampler, w, pos).Y)

	cliffNormal := xyz(triplanarTexture(s.CliffNormalSampler, w, cliffPos))
	grassNormal := xyz(triplanarTexture(s.GrassNormalSampler, w, pos))
	dirtNormal := xyz(triplanarTexture(s.DirtNormalSampler, w, pos))

	wallMotif := float.X(triplanarTexture(s.WallMaskSampler, w, cliffPos).X)
	wallMask := motifMask(float.Mul(frag.ObjectNormal.Y, s.WallEdgeOffset), wallMotif, 0.8)
	pathMask := motifMask(float.Add(frag.Color.R, 0.5), float.Mul(grassRoughness, 2.0), 0.8)

	grassEdge := float.X(triplanarTexture(s.GrassPatchPatternSampler, w, vec3.Mul(pos, 0.05)).X)
	grassMixMask := motifMask(float.SmoothStep(0.5, 0.6, grassEdge), grassRoughness, 0.8)
	grassMixed := vec3.Mix(vec3.Mul(grassAlbedo, 0.85), vec3.Mul(grassAlbedo, 0.5), grassMixMask)

	groundAlbedo := vec3.Mix(grassMixed, dirtAlbedo, pathMask)
	groundRoughness := float.Mix(grassRoughness, dirtRoughness, pathMask)
	groundNormal := vec3.Mix(grassNormal, dirtNormal, pathMask)
	return Spatial.Material{
		Albedo:    color(vec3.Mix(cliffAlbedo, groundAlbedo, wallMask)),
		Roughness: float.Mix(cliffRoughness, groundRoughness, wallMask),
		NormalMap: vec3.Mix(cliffNormal, groundNormal, wallMask),
	}
}

// LargeTreeTrunk blends bark into the cut top of the trunk.
type LargeTreeTrunk struct {
	Spatial.Shader[LargeTreeTrunk]

	BarkAlbedoSampler     sampler `gd:",source_color"`
	TrunkTopAlbedoSampler sampler `gd:",source_color"`
	MossAlbedoSampler     sampler `gd:",source_color"`
	BarkOrmSampler        sampler
	TrunkTopOrmSampler    sampler
	MossOrmSampler        sampler
	BarkNormalSampler     sampler
	TrunkTopNormalSampler sampler
	MossNormalSampler     sampler
	BarkMaskSampler       sampler `gd:",source_color"`
	TriplanarScale        vec3.XYZ
	TriplanarOffset       vec3.XYZ
	TriplanarSharpness    float.X
}

func NewLargeTreeTrunk() *LargeTreeTrunk {
	return &LargeTreeTrunk{TriplanarScale: vec3.New(1.0, 1.0, 1.0), TriplanarSharpness: float.New(0.5)}
}

func (s *LargeTreeTrunk) Fragment(vertex Spatial.Vertex) Triplanar {
	out := triplanar(vertex, s.TriplanarScale, s.TriplanarOffset, s.TriplanarSharpness)
	out.TriplanarPos = vec3.Add(out.TriplanarPos, vertex.NodePositionWorld)
	return out
}

func (s *LargeTreeTrunk) Material(frag Triplanar) Spatial.Material {
	w, pos := frag.TriplanarPowerNormal, frag.TriplanarPos
	barkAlbedo := xyz(triplanarTexture(s.BarkAlbedoSampler, w, pos))
	barkRoughness := float.X(triplanarTexture(s.BarkOrmSampler, w, pos).X)
	barkNormal := xyz(triplanarTexture(s.BarkNormalSampler, w, pos))
	topAlbedo := xyz(triplanarTexture(s.TrunkTopAlbedoSampler, w, pos))
	// As in the original, the roughness of the top comes from its albedo texture.
	topRoughness := float.X(triplanarTexture(s.TrunkTopAlbedoSampler, w, pos).X)
	topNormal := xyz(triplanarTexture(s.TrunkTopNormalSampler, w, pos))
	barkMotif := float.X(triplanarTexture(s.BarkMaskSampler, w, pos).X)
	wallMask := motifMask(float.Mul(frag.ObjectNormal.Y, 1.0), barkMotif, 0.3)
	return Spatial.Material{
		Albedo:    color(vec3.Mix(barkAlbedo, topAlbedo, wallMask)),
		Roughness: float.Mix(barkRoughness, topRoughness, wallMask),
		NormalMap: vec3.Mix(barkNormal, topNormal, wallMask),
	}
}

// SkyShader is a gradient sky with a sun, stars and clouds.
type SkyShader struct {
	Sky.Shader[SkyShader]

	CloudsSampler sampler  `gd:",filter_linear_mipmap"`
	TopColor      vec3.RGB `gd:",source_color"`
	BottomColor   vec3.RGB `gd:",source_color"`
	SunScatter    vec3.RGB `gd:",source_color"`
	StarStength   float.X  `gd:"star_stength,hint_range(0.0, 5.0, 0.1)"`
}

func NewSkyShader() *SkyShader {
	white := rgb.New(1.0, 1.0, 1.0)
	return &SkyShader{TopColor: white, BottomColor: white, SunScatter: white, StarStength: float.New(0.0)}
}

const sunRadius = 0.08

func hash(x vec3.XYZ) vec3.XYZ {
	x = vec3.New(
		vec3.Dot(x, vec3.New(127.1, 311.7, 74.7)),
		vec3.Dot(x, vec3.New(269.5, 183.3, 246.1)),
		vec3.Dot(x, vec3.New(113.5, 271.9, 124.6)),
	)
	return vec3.Fract(vec3.Mul(vec3.Sin(x), 43758.5453123))
}

// voronoi returns the distance to the nearest cell and that cell's id. The Go
// loops unroll into the shader, with each if written as a select.
//
// Voronoi method credit:
// The MIT License
// Copyright © 2013 Inigo Quilez
// https://www.shadertoy.com/view/ldl3Dl
func voronoi(x vec3.XYZ) (distance, id float.X) {
	p := vec3.Floor(x)
	f := vec3.Fract(x)
	id = float.New(0.0)
	res := vec2.New(100.0, 100.0)
	for k := -1.0; k <= 1; k++ {
		for j := -1.0; j <= 1; j++ {
			for i := -1.0; i <= 1; i++ {
				b := vec3.New(i, j, k)
				r := vec3.Add(vec3.Sub(b, f), hash(vec3.Add(p, b)))
				d := float.X(vec3.Dot(r, r))
				nearest := float.Lt(d, res.X)
				id = bool.Select(nearest, float.X(vec3.Dot(vec3.Add(p, b), vec3.New(1.0, 57.0, 113.0))), id)
				res = bool.Select(nearest, vec2.New(d, res.X),
					bool.Select(float.Lt(d, res.Y), vec2.New(res.X, d), res))
			}
		}
	}
	return float.Sqrt(res.X), float.Abs(id)
}

func (s *SkyShader) Lighting(sky Sky.Snapshot) Sky.Lighting {
	eye := sky.EyeDirection
	sun := sky.Lights()[0]
	gradient := rgb.Mix(s.BottomColor, s.TopColor, float.Clamp(eye.Y, 0.0, 1.0))
	sunPosition := vec3.Distance(eye, sun.Direction)
	sunEdge := float.Mul(float.SmoothStep(sunRadius, sunRadius*0.9, sunPosition), 0.2)
	sunMask := float.Mul(float.SmoothStep(sunRadius*0.7, sunRadius*0.65, sunPosition), 0.8)
	sunColor := float.Mul(float.Add(sunEdge, sunMask), sun.Energy)
	sunset := rgb.Mul(s.SunScatter, float.Sub(1.0, float.Abs(float.Mul(eye.Y, 1.0))))
	c := rgb.Add(rgb.Add(gradient, sunset), float.Mul(sunColor, float.Clamp(float.Mul(eye.Y, 4.0), 0.0, 1.0)))

	// Stars.
	distance, id := voronoi(vec3.Mul(eye, 25.0))
	twinkle := float.Add(0.025, float.Mul(half(float.Sin(float.Add(shaders.Time, id))), 0.05))
	stars := rgb.Add(c, float.Mul(float.SmoothStep(twinkle, 0.0, distance), s.StarStength))
	c = bool.Select(float.Gt(s.StarStength, 0.0), stars, c)

	// Clouds.
	cloudUV := vec2.Div(vec2.New(eye.X, eye.Z), eye.Y)
	cloudMask := float.Mul(s.CloudsSampler.Sample(vec2.Add(vec2.Mul(cloudUV, 0.05), float.Mul(shaders.Time, 0.001))).X, float.Step(sky.SkyCoords.Y, 0.5))
	return Sky.Lighting{
		Color: rgb.Mix(c, rgb.New(1.0, 1.0, 1.0), float.SmoothStep(0.75, 0.85, cloudMask)),
	}
}

// Blur blurs what is behind the demo page, along both diagonals.
type Blur struct {
	CanvasItem.Shader[Blur]

	ScreenTexture texture.Sampler2D[vec4.RGBA] `gd:",hint_screen_texture,repeat_disable,filter_nearest"`
	BlurSize      float.X
	ColorModulate vec3.RGB `gd:",source_color"`
}

func NewBlur() *Blur { return &Blur{BlurSize: float.New(1.0), ColorModulate: rgb.New(1.0, 1.0, 1.0)} }

var blurWeights = [8]float64{0.05467, 0.080657, 0.106483, 0.125794, 0.132981, 0.125794, 0.106483, 0.080657}

const blurTotalWeight = 1.627038

func (s *Blur) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	scale := vec2.Mul(frag.ScreenPixelSize(), s.BlurSize)
	directions := []vec2.XY{
		vec2.Mul(frag.TexturePixelSize(), vec2.New(s.BlurSize, s.BlurSize)),
		vec2.Mul(frag.TexturePixelSize(), vec2.New(s.BlurSize, float.Neg(s.BlurSize))),
	}
	sum := rgba.New(0.0, 0.0, 0.0, 0.0)
	for _, direction := range directions {
		for i := -4; i < 4; i++ {
			offset := vec2.Mul(vec2.Mul(direction, float64(i)), scale)
			sum = rgba.Add(sum, rgba.Mul(s.ScreenTexture.Sample(vec2.Add(frag.ScreenUV(), offset)), blurWeights[i+4]))
		}
	}
	sum = rgba.Div(sum, blurTotalWeight)
	return CanvasItem.Material{
		Color: rgba.New(float.Mul(sum.R, s.ColorModulate.R), float.Mul(sum.G, s.ColorModulate.G), float.Mul(sum.B, s.ColorModulate.B), sum.A),
	}
}
