/*
func get_exposure(aperture: float, shutter_speed: float, sensitivity: float):
	return log((aperture * aperture) / shutter_speed * (100.0 / sensitivity)) / log(2)
*/

package main

import "math"

func getExposure(aperture, shutterSpeed, sensitivity float64) float64 {
	return math.Log((aperture*aperture)/shutterSpeed*(100.0/sensitivity)) / math.Log(2)
}
