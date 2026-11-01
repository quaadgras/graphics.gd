package main

import (
	"graphics.gd/shaders"
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/int"
	"graphics.gd/shaders/over"
	"graphics.gd/shaders/pipeline/Sky"
	"graphics.gd/shaders/rgb"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/swizzle"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

// VolumetricClouds draws an analytic sky with raymarched clouds. Ported from
// sky_volumetric_clouds.gdshader, which is based on:
// A. Schneider. "The Real-Time Volumetric Cloudscapes Of Horizon: Zero Dawn".
// ACM SIGGRAPH. Los Angeles, CA: ACM SIGGRAPH, 2015. Web. 26 Aug. 2015.
type VolumetricClouds struct {
	Sky.Shader[VolumetricClouds]

	Worlnoise     texture.Sampler3D[vec4.RGBA] `gd:",filter_linear_mipmap,repeat_enable"`
	Perlworlnoise texture.Sampler3D[vec4.RGBA] `gd:",filter_linear_mipmap,repeat_enable"`
	CloudMap      texture.Sampler2D[vec4.RGBA] `gd:",filter_linear_mipmap,repeat_enable"`

	CloudDensity            float.X `gd:",hint_range(0.01, 0.99)"`
	CloudCoverage           float.X `gd:",hint_range(0.0, 0.99)"`
	CloudTimeScale          float.X `gd:",hint_range(0.0, 20.0)"`
	CloudTimeOffset         float.X `gd:",hint_range(0.0, 1000.0, 0.1)"`
	CloudDirection          vec2.XY
	CloudStepsRange         vec2.XY
	CloudsScreenLod         float.X   `gd:",hint_range(0.0, 1.0, 0.1)"`
	CloudModulate           vec4.RGBA `gd:",source_color"`
	CloudCameraPositionMult vec2.XY
	CloudCameraHeight       float.X

	Rayleigh        float.X   `gd:",hint_range(0, 64)"`
	RayleighColor   vec4.RGBA `gd:",source_color"`
	Mie             float.X   `gd:",hint_range(0, 1)"`
	MieEccentricity float.X   `gd:",hint_range(-1, 1)"`
	MieColor        vec4.RGBA `gd:",source_color"`

	Turbidity    float.X   `gd:",hint_range(0, 1000)"`
	SunDiskScale float.X   `gd:",hint_range(0, 360)"`
	GroundColor  vec4.RGBA `gd:",source_color"`
	Exposure     float.X   `gd:",hint_range(0, 128)"`

	SpaceEmit         texture.Sampler2D[vec4.RGBA] `gd:",hint_default_transparent,filter_linear"`
	SpaceEmitColor    vec4.RGBA                    `gd:",source_color"`
	SpaceEmitAngle    float.X                      `gd:",hint_range(0.0, 1.0, 0.001)"`
	SpaceEmitStrength float.X
}

func NewVolumetricClouds() *VolumetricClouds {
	return &VolumetricClouds{
		CloudDensity:            float.New(0.05),
		CloudCoverage:           float.New(0.2),
		CloudTimeScale:          float.New(1.0),
		CloudTimeOffset:         float.New(0.0),
		CloudDirection:          vec2.New(0.9, 1.0),
		CloudStepsRange:         vec2.New(192.0, 108.0),
		CloudsScreenLod:         float.New(1.0),
		CloudModulate:           rgba.New(1.0, 1.0, 1.0, 1.0),
		CloudCameraPositionMult: vec2.New(0.1, 0.2),
		Rayleigh:                float.New(2.0),
		RayleighColor:           rgba.New(0.26, 0.41, 0.58, 1.0),
		Mie:                     float.New(0.005),
		MieEccentricity:         float.New(0.8),
		MieColor:                rgba.New(0.63, 0.77, 0.92, 1.0),
		Turbidity:               float.New(10.0),
		SunDiskScale:            float.New(1.0),
		GroundColor:             rgba.New(1.0, 1.0, 1.0, 1.0),
		Exposure:                float.New(0.1),
		SpaceEmitColor:          rgba.New(1.0, 1.0, 1.0, 1.0),
		SpaceEmitStrength:       float.New(1.0),
	}
}

func (*VolumetricClouds) RenderMode() []Sky.RenderMode {
	return []Sky.RenderMode{Sky.UseHalfResPass, Sky.UseQuarterResPass}
}

// Approximately earth sizes.
const (
	groundRadius    = 6000000.0
	skyBottomRadius = 6001500.0 // Bottom of cloud layer.
	skyTopRadius    = 6004000.0 // Top of cloud layer.
)

// Sun constants.
const (
	solSize   = 0.00872663806
	sunEnergy = 1000.0
)

// Optical length at zenith for molecules.
const (
	rayleighZenithSize = 8.4e3
	mieZenithSize      = 1.25e3
)

const weatherScale = 0.00006

var randomVectors = [6][3]float64{
	{0.38051305, 0.92453449, -0.02111345},
	{-0.50625799, -0.03590792, -0.86163418},
	{-0.32509218, -0.94557439, 0.01428793},
	{0.09026238, -0.27376545, 0.95755165},
	{0.28128598, 0.42443639, -0.86065785},
	{-0.16852403, 0.14748697, 0.97460106},
}

var up = vec3.New(0.0, 1.0, 0.0)

func parabola[X, K float.Any](x X, k K) float.X {
	return float.Pow(float.Mul(float.Mul(4.0, x), float.Sub(1.0, x)), k)
}

// hash is from https://www.shadertoy.com/view/4sfGzS, credit to Inigo Quilez.
func hash(p vec3.XYZ) float.X {
	p = vec3.Fract(vec3.Add(vec3.Mul(p, 0.3183099), 0.1))
	p = vec3.Mul(p, 17.0)
	return float.Fract(float.Mul(float.Mul(float.Mul(p.X, p.Y), p.Z), float.Add(float.Add(p.X, p.Y), p.Z)))
}

// remap maps a value from one range to another (like GDScript's remap).
func remap[V, A, B, C, D float.Any](value V, oldMin A, oldMax B, newMin C, newMax D) float.X {
	return float.Add(newMin, float.Mul(float.Div(float.Sub(value, oldMin), float.Sub(oldMax, oldMin)), float.Sub(newMax, newMin)))
}

// henyeyGreenstein is the phase function.
func henyeyGreenstein[C, G float.Any](cosTheta C, g G) float.X {
	const k = 0.0795774715459
	g2 := float.Mul(g, g)
	return float.Div(float.Mul(k, float.Sub(1.0, g2)), float.Pow(float.Sub(float.Add(1.0, g2), float.Mul(float.Mul(2.0, g), cosTheta)), 1.5))
}

func toRGB(v vec3.XYZ) vec3.RGB { return swizzle.RGB(v) }

// atmosphere is a simple analytic sky. In a real project, you should use a
// texture to improve performance.
func (s *VolumetricClouds) atmosphere(sun Sky.Light, eyeDir vec3.XYZ, coords vec2.XY, space float.X) vec3.RGB {
	sunDir := vec3.Normalize(sun.Direction)
	zenithAngle := float.Clamp(vec3.Dot(up, sunDir), -1.0, 1.0)
	energy := float.Mul(float.Mul(float.Max(0.0, float.Sub(1.0, float.Exp(float.Neg(float.Sub(float.Mul(shaders.PI, 0.5), float.Acos(zenithAngle)))))), sunEnergy), sun.Energy)
	sunFade := float.Sub(1.0, float.Clamp(float.Sub(1.0, float.Exp(sun.Direction.Y)), 0.0, 1.0))

	// Rayleigh coefficients.
	rayleighCoefficient := float.Sub(s.Rayleigh, float.Mul(1.0, float.Sub(1.0, sunFade)))
	rayleighBeta := rgb.Mul(rgb.Mul(swizzle.RGB(s.RayleighColor), rayleighCoefficient), 0.0001)
	// Mie coefficients from Preetham.
	mieBeta := rgb.Mul(rgb.Mul(swizzle.RGB(s.MieColor), float.Mul(s.Turbidity, s.Mie)), 0.000434)

	// Optical length.
	zenith := float.Acos(float.Max(0.0, vec3.Dot(up, eyeDir)))
	opticalMass := float.Div(1.0, float.Add(float.Cos(zenith), float.Mul(0.15, float.Pow(float.Sub(93.885, float.Degrees(zenith)), -1.253))))
	rayleighScatter := float.Mul(rayleighZenithSize, opticalMass)
	mieScatter := float.Mul(mieZenithSize, opticalMass)

	// Light extinction based on thickness of atmosphere.
	extinction := rgb.Exp(rgb.Neg(rgb.Add(rgb.Mul(rayleighBeta, rayleighScatter), rgb.Mul(mieBeta, mieScatter))))

	// In scattering.
	cosTheta := vec3.Dot(eyeDir, sunDir)
	rayleighPhase := float.Mul(3.0/(16.0*3.141592653589793), float.Add(1.0, float.Pow(float.Add(float.Mul(cosTheta, 0.5), 0.5), 2.0)))
	betaRTheta := rgb.Mul(rayleighBeta, rayleighPhase)
	betaMTheta := rgb.Mul(mieBeta, henyeyGreenstein(cosTheta, s.MieEccentricity))

	scatter := rgb.Mul(rgb.Div(rgb.Add(betaRTheta, betaMTheta), rgb.Add(rayleighBeta, mieBeta)), energy)
	lin := rgb.Pow(rgb.Mul(scatter, rgb.Sub(rgb.New(1.0, 1.0, 1.0), extinction)), rgb.New(1.5, 1.5, 1.5))
	// Hack from https://github.com/mrdoob/three.js/blob/master/examples/jsm/objects/Sky.js.
	lin = rgb.Mul(lin, rgb.Mix(rgb.New(1.0, 1.0, 1.0), rgb.Pow(rgb.Mul(scatter, extinction), rgb.New(0.5, 0.5, 0.5)), float.Clamp(float.Pow(float.Sub(1.0, zenithAngle), 5.0), 0.0, 1.0)))

	// Hack in the ground color.
	lin = rgb.Mul(lin, rgb.Mix(swizzle.RGB(s.GroundColor), rgb.New(1.0, 1.0, 1.0), float.SmoothStep(-0.1, 0.1, vec3.Dot(up, eyeDir))))

	// Solar disk and out-scattering.
	sunAngularDiameterCos := float.Cos(float.Mul(solSize, s.SunDiskScale))
	sunAngularDiameterCos2 := float.Cos(float.Mul(float.Mul(solSize, s.SunDiskScale), 0.5))
	sundisk := float.SmoothStep(sunAngularDiameterCos, sunAngularDiameterCos2, cosTheta)
	l0 := rgb.Mul(rgb.Mul(rgb.Mul(extinction, float.Mul(energy, 1900.0)), sundisk), sun.Color)

	// Add night-time sky.
	spaceCoords := vec2.New(float.Fract(float.Add(coords.X, s.SpaceEmitAngle)), coords.Y)
	spaceTex := s.SpaceEmit.Sample(vec2.Mul(spaceCoords, vec2.New(1.0, 2.0)))
	spaceColor := rgb.Mul(rgb.Mul(rgb.Mul(swizzle.RGB(spaceTex), spaceTex.A), swizzle.RGB(s.SpaceEmitColor)), s.SpaceEmitColor.A)
	l0 = rgb.Add(l0, rgb.Mul(rgb.Mul(rgb.Mul(spaceColor, float.Mul(5000.0, s.SpaceEmitStrength)), extinction), space))

	color := rgb.Mul(rgb.Add(lin, l0), 0.04)
	exponent := float.Div(1.0, float.Add(1.2, float.Mul(1.2, sunFade)))
	color = rgb.Pow(color, rgb.New(exponent, exponent, exponent))
	return rgb.Mul(color, s.Exposure)
}

func heightFraction(position float.X) float.X {
	return float.Clamp(float.Div(float.Sub(position, skyBottomRadius), skyTopRadius-skyBottomRadius), 0.0, 1.0)
}

func mixGradients(cloudType float.X) vec4.XYZW {
	stratusGradient := vec4.New(0.02, 0.05, 0.09, 0.11)
	stratocumulusGradient := vec4.New(0.02, 0.2, 0.48, 0.625)
	cumulusGradient := vec4.New(0.01, 0.0625, 0.78, 1.0)
	stratus := float.Sub(1.0, float.Clamp(float.Mul(cloudType, 2.0), 0.0, 1.0))
	stratocumulus := float.Sub(1.0, float.Mul(float.Abs(float.Sub(cloudType, 0.5)), 2.0))
	cumulus := float.Mul(float.Clamp(float.Sub(cloudType, 0.5), 0.0, 1.0), 2.0)
	return vec4.Add(vec4.Add(vec4.Mul(stratusGradient, stratus), vec4.Mul(stratocumulusGradient, stratocumulus)), vec4.Mul(cumulusGradient, cumulus))
}

func densityHeightGradient(heightFrac, cloudType float.X) float.X {
	gradient := mixGradients(cloudType)
	return float.Sub(float.SmoothStep(gradient.X, gradient.Y, heightFrac), float.SmoothStep(gradient.Z, gradient.W, heightFrac))
}

func intersectSphere[R float.Any](pos, dir vec3.XYZ, r R) float.X {
	a := vec3.Dot(dir, dir)
	b := float.Mul(2.0, vec3.Dot(dir, pos))
	c := float.Sub(vec3.Dot(pos, pos), float.Mul(r, r))
	d := float.Sqrt(float.Sub(float.Mul(b, b), float.Mul(float.Mul(4.0, a), c)))
	p := float.Sub(float.Neg(b), d)
	p2 := float.Add(float.Neg(b), d)
	return float.Div(float.Max(p, p2), float.Mul(2.0, a))
}

// density returns the density of the clouds at a given point, heavily based on
// the method from Schneider.
func (s *VolumetricClouds) density(pip vec3.XYZ, weather vec4.RGBA, mip float.X) float.X {
	p := vec3.New(float.Add(pip.X, float.Mul(float.Mul(10.0, s.CloudTimeOffset), s.CloudTimeScale)), pip.Y, pip.Z)
	height := heightFraction(vec3.Length(p))
	n := s.Perlworlnoise.SampleLOD(vec3.Mul(p, 0.00008), float.Sub(mip, 2.0))
	fbm := float.Add(float.Add(float.Mul(n.G, 0.625), float.Mul(n.B, 0.25)), float.Mul(n.A, 0.125))
	g := densityHeightGradient(height, weather.R)
	baseCloud := remap(n.R, float.Neg(float.Sub(1.0, fbm)), 1.0, 0.0, 1.0)
	weatherCoverage := float.Mul(s.CloudCoverage, weather.B)
	baseCloud = remap(float.Mul(baseCloud, g), float.Sub(1.0, weatherCoverage), 1.0, 0.0, 1.0)
	baseCloud = float.Mul(baseCloud, weatherCoverage)
	shift := float.Mul(float.Mul(40.0, s.CloudTimeOffset), s.CloudTimeScale)
	p = vec3.New(float.Sub(p.X, shift), float.Sub(p.Y, shift), p.Z)
	hn := s.Worlnoise.SampleLOD(vec3.Mul(p, 0.001), mip)
	hfbm := float.Add(float.Add(float.Mul(hn.R, 0.625), float.Mul(hn.G, 0.25)), float.Mul(hn.B, 0.125))
	hfbm = float.Mix(hfbm, float.Sub(1.0, hfbm), float.Clamp(float.Mul(height, 4.0), 0.0, 1.0))
	baseCloud = remap(baseCloud, float.Mul(float.Mul(hfbm, 0.4), height), 1.0, 0.0, 1.0)
	return float.Pow(float.Clamp(baseCloud, 0.0, 1.0), float.Add(float.Mul(float.Sub(1.0, height), 0.8), 0.5))
}

func (s *VolumetricClouds) march(sun Sky.Light, pos, dir vec3.XYZ, depth int.X) vec4.RGBA {
	ss := vec3.Length(dir)
	dir = vec3.Normalize(dir)
	lss := float.New((skyTopRadius - skyBottomRadius) / 36.0)
	ldir := vec3.Normalize(sun.Direction)

	cosTheta := vec3.Dot(ldir, dir)
	// Stack multiple phase functions to emulate some backscattering.
	phase := float.Max(float.Max(henyeyGreenstein(cosTheta, 0.6), henyeyGreenstein(cosTheta, float.Sub(0.4, float.Mul(1.4, ldir.Y)))), henyeyGreenstein(cosTheta, -0.2))

	// Precalculate sun and ambient colors.
	// This should really come from a uniform or texture for performance reasons.
	atmosphereSun := rgb.Mul(rgb.Mul(rgb.Mul(s.atmosphere(sun, sun.Direction, vec2.New(0.0, 0.0), float.New(0.0)), sun.Energy), ss), 0.1)
	atmosphereAmbient := s.atmosphere(sun, vec3.Normalize(vec3.New(1.0, 1.0, 0.0)), vec2.New(0.0, 0.0), float.New(0.0))
	atmosphereGround := s.atmosphere(sun, vec3.Normalize(vec3.New(1.0, -1.0, 0.0)), vec2.New(0.0, 0.0), float.New(0.0))

	time := float.Mul(float.Mul(0.003, s.CloudTimeOffset), s.CloudTimeScale)
	weatherPos := vec2.New(float.Mul(time, s.CloudDirection.X), float.Mul(time, s.CloudDirection.Y))
	weatherAt := func(p vec3.XYZ, offset vec2.XY) vec4.RGBA {
		return s.CloudMap.Sample(vec2.Add(vec2.Add(vec2.Mul(vec2.New(p.X, p.Z), weatherScale), 0.5), offset))
	}

	T := float.New(1.0)
	alpha := float.New(0.0)
	L := rgb.New(0.0, 0.0, 0.0)
	p := vec3.Add(pos, float.Mul(hash(vec3.Mul(pos, 10.0)), ss))
	for range over.Int(depth) {
		p = vec3.Add(p, vec3.Mul(dir, ss))
		height := heightFraction(vec3.Length(p))
		t := s.density(p, weatherAt(p, weatherPos), float.New(0.0))
		dt := float.Exp(float.Mul(float.Mul(float.Neg(s.CloudDensity), t), ss))
		T = float.Mul(T, dt)
		if float.Gt(t, 0.0).Evaluate() {
			// Calculate lighting, but only when we're in the cloud.
			lp := p
			cd := float.New(0.0)
			for j, random := range randomVectors {
				offset := vec3.Mul(vec3.New(random[0], random[1], random[2]), float64(j))
				lp = vec3.Add(lp, vec3.Mul(vec3.Add(ldir, offset), lss))
				cd = float.Add(cd, s.density(lp, weatherAt(lp, weatherPos), float.New(float64(j))))
			}
			// Take a single distant sample.
			lp = vec3.Add(p, vec3.Mul(vec3.Mul(ldir, 18.0), lss))
			lheight := heightFraction(vec3.Length(lp))
			lt := float.Pow(s.density(lp, weatherAt(lp, vec2.New(0.0, 0.0)), float.New(5.0)), float.Add(float.Mul(float.Sub(1.0, lheight), 0.8), 0.5))
			cd = float.Add(cd, lt)

			// Captures the direct lighting from the sun.
			beers := float.Exp(float.Mul(float.Mul(float.Neg(s.CloudDensity), cd), lss))
			beers2 := float.Mul(float.Exp(float.Mul(float.Mul(float.Mul(float.Neg(s.CloudDensity), cd), lss), 0.25)), 0.7)
			beersTotal := float.Max(beers, beers2)

			ambient := rgb.Mul(rgb.Mul(rgb.Mul(rgb.Mix(atmosphereGround, rgb.New(1.0, 1.0, 1.0), float.SmoothStep(0.0, 1.0, height)), s.CloudDensity), rgb.Mix(atmosphereAmbient, rgb.New(1.0, 1.0, 1.0), 0.4)), sun.Direction.Y)
			alpha = float.Add(alpha, float.Mul(float.Sub(1.0, dt), float.Sub(1.0, alpha)))
			L = rgb.Add(L, rgb.Mul(rgb.Mul(rgb.Add(ambient, rgb.Mul(rgb.Mul(atmosphereSun, float.Mul(beersTotal, phase)), alpha)), T), t))
		}
	}
	light := rgb.ClampX(L, 0.0, 1.0)
	return rgba.New(light.R, light.G, light.B, float.Clamp(alpha, 0.0, 1.0))
}

func (s *VolumetricClouds) clouds(sun Sky.Light, dir vec3.XYZ, stepsIn vec2.XY, screenX float.X, camera vec3.XYZ) vec4.RGBA {
	if !float.Gt(dir.Y, 0.0).Evaluate() {
		return rgba.New(0.0, 0.0, 0.0, 0.0)
	}
	cameraPosition := vec3.New(
		float.Mul(camera.X, s.CloudCameraPositionMult.X),
		float.Add(float.Add(groundRadius, float.Mul(camera.Y, s.CloudCameraPositionMult.Y)), s.CloudCameraHeight),
		float.Mul(camera.Z, s.CloudCameraPositionMult.X),
	)
	start := vec3.Add(cameraPosition, vec3.Mul(dir, intersectSphere(cameraPosition, dir, skyBottomRadius)))
	end := vec3.Add(cameraPosition, vec3.Mul(dir, intersectSphere(cameraPosition, dir, skyTopRadius)))
	shellDistance := vec3.Length(vec3.Sub(end, start))
	// Take fewer steps towards horizon and screen sides.
	steps := float.Mix(stepsIn.Y, stepsIn.X, float.Clamp(float.Mul(vec3.Dot(dir, up), float.Mix(1.0, parabola(screenX, 0.75), s.CloudsScreenLod)), 0.0, 1.0))
	rayStep := vec3.Div(vec3.Mul(dir, shellDistance), steps)
	volume := s.march(sun, start, rayStep, int.FromFloat(steps))
	return rgba.Mul(volume, s.CloudModulate)
}

func (s *VolumetricClouds) cloudsPost(sun Sky.Light, direction vec3.XYZ, volume vec4.RGBA, background vec3.RGB) vec3.RGB {
	// Draw background at full resolution pass.
	if !float.Gt(direction.Y, 0.0).Evaluate() {
		return s.atmosphere(sun, direction, vec2.New(0.0, 0.0), float.New(0.0))
	}
	// Draw cloud shape.
	col := rgb.Add(rgb.Mul(background, float.Sub(1.0, volume.A)), swizzle.RGB(volume))
	// Blend distant clouds into the sky.
	return rgb.Mix(rgb.ClampX(col, 0.0, 1.0), rgb.ClampX(background, 0.0, 1.0), float.SmoothStep(0.6, 1.0, float.Sub(1.0, direction.Y)))
}

func (s *VolumetricClouds) Lighting(sky Sky.Snapshot) Sky.Lighting {
	sun := sky.Lights()[0]
	background := s.atmosphere(sun, sky.EyeDirection, sky.SkyCoords, float.New(1.0))
	switch {
	case bool.And(sky.AtHalfResPass, bool.Not(sky.AtCubemapPass)).Evaluate():
		// Draw to half-resolution buffer for main sky.
		col := s.clouds(sun, sky.EyeDirection, s.CloudStepsRange, sky.ScreenUV.X, sky.CameraPosition)
		return Sky.Lighting{Color: swizzle.RGB(col), Alpha: float.X(col.A)}
	case bool.And(sky.AtQuarterResPass, sky.AtCubemapPass).Evaluate():
		// Draw to quarter-resolution buffer for reflections.
		volume := s.clouds(sun, sky.EyeDirection, vec2.New(15.0, 10.0), sky.ScreenUV.X, sky.CameraPosition)
		return Sky.Lighting{Color: s.cloudsPost(sun, sky.EyeDirection, volume, background), Alpha: float.New(1.0)}
	case sky.AtCubemapPass.Evaluate():
		return Sky.Lighting{Color: swizzle.RGB(sky.QuarterResColor), Alpha: float.New(1.0)}
	default:
		return Sky.Lighting{Color: s.cloudsPost(sun, sky.EyeDirection, sky.HalfResColor, background), Alpha: float.New(1.0)}
	}
}
