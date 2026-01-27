/*
img.convert(Image.FORMAT_RGBA8)
img.linear_to_srgb()
*/

package main

import "graphics.gd/classdb/Image"

func ExampleViewportTexture(img Image.Instance) {
	img.Convert(Image.FormatRgba8)
	img.LinearToSrgb()
}
