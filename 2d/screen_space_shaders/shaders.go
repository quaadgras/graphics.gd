package main

import (
	"graphics.gd/shaders"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/mat2"
	"graphics.gd/shaders/pipeline/CanvasItem"
	"graphics.gd/shaders/rgb"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/swizzle"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

// Shaders ported from the .gdshader files in the original demo. Each one
// samples what has already been drawn to the screen behind the TextureRect it
// is applied to and replaces the color channels, keeping the rect's alpha.

// Screen is the screen texture, with gaussian blurred copies in its mipmaps.
type Screen = texture.Sampler2D[vec4.RGBA]

// screen samples the screen behind the current pixel.
func screen[T float.Any](s Screen, uv vec2.XY, lod T) vec3.RGB {
	return swizzle.RGB(s.SampleLOD(uv, float.New(lod)))
}

// output sets the color channels of the pixel to c.
func output(frag CanvasItem.Fragment, c vec3.RGB) CanvasItem.Material {
	return CanvasItem.Material{
		Color: rgba.New(c.R, c.G, c.B, frag.Color.A),
	}
}

// brightness of c, as the average of its channels.
func brightness(c vec3.RGB) float.X {
	return float.X(rgb.Dot(c, rgb.New(0.33333, 0.33333, 0.33333)))
}

type Vignette struct {
	CanvasItem.Shader[Vignette]

	ScreenTexture Screen `gd:",hint_screen_texture,filter_linear_mipmap"`
	Vignette      texture.Sampler2D[vec4.RGBA]
}

func (s *Vignette) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	vignette := swizzle.RGB(s.Vignette.Sample(frag.UV))
	blurred := screen(s.ScreenTexture, frag.ScreenUV(), float.Mul(float.Sub(1.0, vignette.R), 4.0))
	return output(frag, rgb.Mul(blurred, vignette))
}

type Blur struct {
	CanvasItem.Shader[Blur]

	ScreenTexture Screen  `gd:",hint_screen_texture,filter_linear_mipmap"`
	Amount        float.X `gd:",hint_range(0.0, 5.0)"`
}

func (s *Blur) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	return output(frag, screen(s.ScreenTexture, frag.ScreenUV(), s.Amount))
}

type Pixelize struct {
	CanvasItem.Shader[Pixelize]

	ScreenTexture Screen `gd:",hint_screen_texture,filter_linear_mipmap"`
	SizeX         float.X
	SizeY         float.X
}

func NewPixelize() *Pixelize {
	return &Pixelize{SizeX: float.New(0.008), SizeY: float.New(0.008)}
}

func (s *Pixelize) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	uv := frag.ScreenUV()
	uv = vec2.Sub(uv, vec2.Mod(uv, vec2.New(s.SizeX, s.SizeY)))
	return output(frag, screen(s.ScreenTexture, uv, 0.0))
}

type Whirl struct {
	CanvasItem.Shader[Whirl]

	ScreenTexture Screen `gd:",hint_screen_texture,filter_linear_mipmap"`
	Rotation      float.X
}

func NewWhirl() *Whirl { return &Whirl{Rotation: float.New(3.0)} }

func (s *Whirl) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	rel := vec2.Sub(frag.ScreenUV(), 0.5)
	angle := float.Mul(vec2.Length(rel), s.Rotation)
	cos, sin := float.Cos(angle), float.Sin(angle)
	rot := mat2.New(vec2.New(cos, float.Neg(sin)), vec2.New(sin, cos))
	uv := vec2.ClampX(vec2.Add(mat2.Transform(rot, rel), 0.5), 0.0, 1.0)
	return output(frag, screen(s.ScreenTexture, uv, 0.0))
}

type Sepia struct {
	CanvasItem.Shader[Sepia]

	ScreenTexture Screen    `gd:",hint_screen_texture,filter_linear_mipmap"`
	Base          vec4.RGBA `gd:",source_color"`
}

func (s *Sepia) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	v := float.Sqrt(brightness(screen(s.ScreenTexture, frag.ScreenUV(), 0.0)))
	return output(frag, rgb.Mul(swizzle.RGB(s.Base), v))
}

type Negative struct {
	CanvasItem.Shader[Negative]

	ScreenTexture Screen `gd:",hint_screen_texture,filter_linear_mipmap"`
}

func (s *Negative) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	return output(frag, rgb.Sub(rgb.New(1.0, 1.0, 1.0), screen(s.ScreenTexture, frag.ScreenUV(), 0.0)))
}

type Contrasted struct {
	CanvasItem.Shader[Contrasted]

	ScreenTexture Screen `gd:",hint_screen_texture,filter_linear_mipmap"`
}

func (s *Contrasted) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	c := screen(s.ScreenTexture, frag.ScreenUV(), 0.0)
	return output(frag, rgb.Mod(rgb.Add(c, 0.5), 1.0))
}

type Normalized struct {
	CanvasItem.Shader[Normalized]

	ScreenTexture Screen `gd:",hint_screen_texture,filter_linear_mipmap"`
}

func (s *Normalized) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	return output(frag, rgb.Normalize(screen(s.ScreenTexture, frag.ScreenUV(), 0.0)))
}

// BCS adjusts the brightness, contrast and saturation of the screen.
type BCS struct {
	CanvasItem.Shader[BCS]

	ScreenTexture Screen `gd:",hint_screen_texture,filter_linear_mipmap"`
	Brightness    float.X
	Contrast      float.X
	Saturation    float.X
}

func NewBCS() *BCS {
	return &BCS{Brightness: float.New(0.8), Contrast: float.New(1.5), Saturation: float.New(1.8)}
}

func (s *BCS) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	c := screen(s.ScreenTexture, frag.ScreenUV(), 0.0)
	c = rgb.Mix(rgb.New(0.0, 0.0, 0.0), c, s.Brightness)
	c = rgb.Mix(rgb.New(0.5, 0.5, 0.5), c, s.Contrast)
	grey := float.Mul(rgb.Dot(rgb.New(1.0, 1.0, 1.0), c), 0.33333)
	c = rgb.Mix(rgb.New(grey, grey, grey), c, s.Saturation)
	return output(frag, c)
}

type Mirage struct {
	CanvasItem.Shader[Mirage]

	ScreenTexture Screen `gd:",hint_screen_texture,filter_linear_mipmap"`
	Frequency     float.X
	Depth         float.X
}

func NewMirage() *Mirage { return &Mirage{Frequency: float.New(60.0), Depth: float.New(0.005)} }

func (s *Mirage) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	uv := frag.ScreenUV()
	wave := float.Mul(float.Sin(float.Add(float.Mul(uv.Y, s.Frequency), shaders.Time)), s.Depth)
	uv = vec2.New(float.Clamp(float.Add(uv.X, wave), 0.0, 1.0), uv.Y)
	return output(frag, screen(s.ScreenTexture, uv, 0.0))
}

type OldFilm struct {
	CanvasItem.Shader[OldFilm]

	ScreenTexture Screen    `gd:",hint_screen_texture,filter_linear_mipmap"`
	Base          vec4.RGBA `gd:",source_color"`
	Grain         texture.Sampler2D[vec4.RGBA]
	GrainStrength float.X
	Vignette      texture.Sampler2D[vec4.RGBA]
	FPS           float.X `gd:"fps"`
	Stretch       float.X
	Flashing      float.X
}

func NewOldFilm() *OldFilm {
	return &OldFilm{
		GrainStrength: float.New(0.3),
		FPS:           float.New(12.0),
		Stretch:       float.New(0.5),
		Flashing:      float.New(0.01),
	}
}

// grain samples the film grain at a pseudo-random offset for the given time.
func (s *OldFilm) grain(time float.X, uv vec2.XY) float.X {
	wobble := func(a, b float.X) float.X {
		return float.Sin(float.Mul(float.Mul(a, time), float.Sin(float.Mul(time, b))))
	}
	offset := vec2.New(wobble(float.New(41.0), float.New(123.0)), wobble(float.New(27.0), float.New(312.0)))
	return float.X(s.Grain.Sample(vec2.Mul(vec2.Add(uv, vec2.Mod(offset, 1.0)), s.Stretch)).R)
}

func (s *OldFilm) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	v := float.Sqrt(brightness(screen(s.ScreenTexture, frag.ScreenUV(), 0.0)))

	// The grain only changes FPS times per second.
	f := float.Div(1.0, s.FPS)
	frame := float.Sub(shaders.Time, float.Mod(shaders.Time, f))
	g := s.grain(frame, frag.UV)
	g = float.Max(g, float.Mul(s.grain(float.Add(frame, f), frag.UV), 0.5))
	g = float.Max(g, float.Mul(s.grain(float.Add(frame, float.Mul(f, 2.0)), frag.UV), 0.25))

	c := rgb.Sub(rgb.Mul(swizzle.RGB(s.Base), v), float.Mul(g, s.GrainStrength))
	c = rgb.Mul(c, s.Vignette.Sample(frag.UV).R)
	ft := float.Mul(shaders.Time, 0.002)
	flash := float.Sin(float.Mul(float.Mul(75.0, ft), float.Sin(float.Mul(ft, 123.0))))
	return output(frag, rgb.Add(c, float.Mul(flash, s.Flashing)))
}
