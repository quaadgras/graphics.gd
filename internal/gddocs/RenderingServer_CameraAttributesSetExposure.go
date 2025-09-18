/*
func get_exposure_normalization(ev100: float):
	return 1.0 / (pow(2.0, ev100) * 1.2)
*/

package main

import "graphics.gd/variant/Float"

func RenderingServer_CameraAttributesSetExposure() {
	GetExposureNormalization := func(ev100 float32) float32 {
		return 1.0 / (Float.Pow(2, ev100) * 12 / 10)
	}
	_ = GetExposureNormalization
}
