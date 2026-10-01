package main

import (
	"fmt"
	"math"

	"graphics.gd/classdb"
	"graphics.gd/classdb/CameraAttributesPhysical"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/DirectionalLight3D"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/OmniLight3D"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/classdb/WorldEnvironment"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

// Options drives the physical light and camera settings of the scene.
type Options struct {
	Control.Extension[Options] `gd:"Options"`

	Sun              DirectionalLight3D.Instance `gd:"../DirectionalLight3D"`
	Lightbulb1       OmniLight3D.Instance        `gd:"../Lightbulb1"`
	Lightbulb2       OmniLight3D.Instance        `gd:"../Lightbulb2"`
	WorldEnvironment WorldEnvironment.Instance   `gd:"../WorldEnvironment"`

	TimeOfDayValue             Label.Instance `gd:"Light/TimeOfDay/Value"`
	SunIntensityValue          Label.Instance `gd:"Light/SunIntensity/Value"`
	Lightbulb1IntensityValue   Label.Instance `gd:"Light/Lightbulb1Intensity/Value"`
	Lightbulb1TemperatureValue Label.Instance `gd:"Light/Lightbulb1Temperature/Value"`
	Lightbulb2IntensityValue   Label.Instance `gd:"Light/Lightbulb2Intensity/Value"`
	Lightbulb2TemperatureValue Label.Instance `gd:"Light/Lightbulb2Temperature/Value"`
	FocusDistanceValue         Label.Instance `gd:"Camera/FocusDistance/Value"`
	FocalLengthValue           Label.Instance `gd:"Camera/FocalLength/Value"`
	ApertureValue              Label.Instance `gd:"Camera/Aperture/Value"`
	ShutterSpeedValue          Label.Instance `gd:"Camera/ShutterSpeed/Value"`
	SensitivityValue           Label.Instance `gd:"Camera/Sensitivity/Value"`
	AutoexposureSpeedValue     Label.Instance `gd:"Camera/AutoexposureSpeed/Value"`
}

// getColorFromTemperature returns color from a given temperature in kelvins (6500K is nearly white).
// Valid range is [1000; 15000]. As explained in the Filament documentation:
// https://google.github.io/filament/Filament.md.html#lighting/directlighting/lightsparameterization
//
// This is the same function as used internally by the engine when setting a
// Light3D's `light_temperature`, but converted to Go.
func getColorFromTemperature(temperature Float.X) Color.RGBA {
	t2 := temperature * temperature
	u := (0.860117757 + 1.54118254e-4*temperature + 1.28641212e-7*t2) /
		(1.0 + 8.42420235e-4*temperature + 7.08145163e-7*t2)
	v := (0.317398726 + 4.22806245e-5*temperature + 4.20481691e-8*t2) /
		(1.0 - 2.89741816e-5*temperature + 1.61456053e-7*t2)

	// Convert to xyY space.
	d := 1.0 / (2.0*u - 8.0*v + 4.0)
	x := 3.0 * u * d
	y := 2.0 * v * d

	// Convert to XYZ space.
	a := 1.0 / max(y, 1e-5)
	xyz := [3]Float.X{x * a, 1.0, (1.0 - x - y) * a}

	// Convert from XYZ to sRGB(linear).
	linear := [3]Float.X{
		3.2404542*xyz[0] - 1.5371385*xyz[1] - 0.4985314*xyz[2],
		-0.9692660*xyz[0] + 1.8760108*xyz[1] + 0.0415560*xyz[2],
		0.0556434*xyz[0] - 0.2040259*xyz[1] + 1.0572252*xyz[2],
	}
	largest := max(1e-5, linear[0], linear[1], linear[2])
	for i := range linear {
		linear[i] /= largest
	}
	// Normalize, clamp, and convert to sRGB.
	clamped := Color.Clamp(
		Color.RGBA{R: linear[0], G: linear[1], B: linear[2], A: 1},
		Color.RGBA{A: 0}, Color.RGBA{R: 1, G: 1, B: 1, A: 1},
	)
	return Color.ToSRGB(clamped)
}

func (o *Options) attributes() CameraAttributesPhysical.Instance {
	attributes, _ := Object.As[CameraAttributesPhysical.Instance](Viewport.Get(o.AsNode()).GetCamera3d().Attributes())
	return attributes
}

func (o *Options) OnTimeOfDayValueChanged(value Float.X) {
	const offset = Angle.Tau * 0.25
	rotation := o.Sun.AsNode3D().Rotation()
	rotation.X = offset + Angle.Radians(Float.Remap(value, 0, 1440, 0, Float.X(Angle.Tau)))

	// Improve and prevent light leaks by hiding the sun if it's below the horizon.
	const epsilon = 0.0001
	o.Sun.AsNode3D().SetVisible(rotation.X > Angle.Tau*0.5+epsilon && rotation.X < Angle.Tau-epsilon)
	o.Sun.AsNode3D().SetRotation(rotation)

	o.TimeOfDayValue.SetText(fmt.Sprintf("%02d:%02d", int(value/60), int(math.Mod(float64(value), 60))))
}

func (o *Options) OnSunIntensityValueChanged(value Float.X) {
	o.Sun.AsLight3D().SetLightIntensityLux(value)
	o.SunIntensityValue.SetText(fmt.Sprintf("%d lux", int(value)))
}

func (o *Options) OnLightbulb1IntensityValueChanged(value Float.X) {
	o.Lightbulb1.AsLight3D().SetLightIntensityLumens(value)
	o.Lightbulb1IntensityValue.SetText(fmt.Sprintf("%d lm", int(value)))
}

func (o *Options) OnLightbulb1TemperatureValueChanged(value Float.X) {
	o.Lightbulb1.AsLight3D().SetLightTemperature(value)
	o.Lightbulb1TemperatureValue.SetText(fmt.Sprintf("%d K", int(value)))
	o.Lightbulb1TemperatureValue.AsControl().AddThemeColorOverride("font_color", getColorFromTemperature(value))
}

func (o *Options) OnLightbulb2IntensityValueChanged(value Float.X) {
	o.Lightbulb2.AsLight3D().SetLightIntensityLumens(value)
	o.Lightbulb2IntensityValue.SetText(fmt.Sprintf("%d lm", int(value)))
}

func (o *Options) OnLightbulb2TemperatureValueChanged(value Float.X) {
	o.Lightbulb2.AsLight3D().SetLightTemperature(value)
	o.Lightbulb2TemperatureValue.SetText(fmt.Sprintf("%d K", int(value)))
	o.Lightbulb2TemperatureValue.AsControl().AddThemeColorOverride("font_color", getColorFromTemperature(value))
}

func (o *Options) OnFocusDistanceValueChanged(value Float.X) {
	o.attributes().SetFrustumFocusDistance(value)
	o.FocusDistanceValue.SetText(fmt.Sprintf("%.1f m", value))
}

func (o *Options) OnFocalLengthValueChanged(value Float.X) {
	o.attributes().SetFrustumFocalLength(value)
	o.FocalLengthValue.SetText(fmt.Sprintf("%d mm", int(value)))
}

func (o *Options) OnApertureValueChanged(value Float.X) {
	o.attributes().SetExposureAperture(value)
	o.ApertureValue.SetText(fmt.Sprintf("%.1f f-stop", value))
}

func (o *Options) OnShutterSpeedValueChanged(value Float.X) {
	o.attributes().SetExposureShutterSpeed(value)
	o.ShutterSpeedValue.SetText(fmt.Sprintf("1/%d", int(value)))
}

func (o *Options) OnSensitivityValueChanged(value Float.X) {
	o.attributes().AsCameraAttributes().SetExposureSensitivity(value)
	o.SensitivityValue.SetText(fmt.Sprintf("%d ISO", int(value)))
}

func (o *Options) OnAutoexposureSpeedValueChanged(value Float.X) {
	o.attributes().AsCameraAttributes().SetAutoExposureSpeed(value)
	o.AutoexposureSpeedValue.SetText(fmt.Sprintf("%.1f", value))
}

func (o *Options) OnSdfgiButtonToggled(buttonPressed bool) {
	o.WorldEnvironment.Environment().SetSdfgiEnabled(buttonPressed)
}

func main() {
	classdb.Register[Options]()
	startup.Scene()
}
