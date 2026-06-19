/*
[gdscript]
func _process(_delta):
	# output_max_linear_value may change often, so do this every frame.
	var max_linear_value = get_window().get_output_max_linear_value()
	# Replace this with your color:
	var original_color = Color.PURPLE
	# Normalize to max_linear_value to produce the brightest color possible,
	# regardless of SDR or HDR output:
	var bright_color = normalize_color(original_color, max_linear_value)


func normalize_color(srgb_color, max_linear_value = 1.0):
	# Color must be linear-encoded to use math operations.
	var linear_color = srgb_color.srgb_to_linear()
	var max_rgb_value = maxf(linear_color.r, maxf(linear_color.g, linear_color.b))
	var brightness_scale = max_linear_value / max_rgb_value
	linear_color *= brightness_scale
	# Undo changes to the alpha channel, which should not be modified.
	linear_color.a = srgb_color.a
	# Convert back to nonlinear sRGB encoding, which is required for Color in
	# Godot unless stated otherwise.
	return linear_color.linear_to_srgb()
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Window"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

func exampleWindowOutputMaxLinear(window Window.Instance) {
	// output_max_linear_value may change often, so do this every frame.
	var maxLinearValue = window.GetOutputMaxLinearValue()
	// Replace this with your color:
	var originalColor = Color.W3C.Purple
	// Normalize to maxLinearValue to produce the brightest color possible,
	// regardless of SDR or HDR output:
	var brightColor = normalizeColor(originalColor, maxLinearValue)
	_ = brightColor
}

func normalizeColor(srgbColor Color.RGBA, maxLinearValue Float.X) Color.RGBA {
	// Color must be linear-encoded to use math operations.
	var linearColor = Color.ToLinear(srgbColor)
	var maxRgbValue = max(linearColor.R, max(linearColor.G, linearColor.B))
	var brightnessScale = maxLinearValue / maxRgbValue
	linearColor.R *= brightnessScale
	linearColor.G *= brightnessScale
	linearColor.B *= brightnessScale
	// Undo changes to the alpha channel, which should not be modified.
	linearColor.A = srgbColor.A
	// Convert back to nonlinear sRGB encoding.
	return Color.ToSRGB(linearColor)
}
