/*
var refresh_rate = DisplayServer.screen_get_refresh_rate()
if refresh_rate < 0:
    refresh_rate = 60.0
*/

package main

import "graphics.gd/classdb/DisplayServer"

func DisplayServer_ScreenGetRefreshRate() {
	var refresh_rate = DisplayServer.ScreenGetRefreshRate()
	if refresh_rate < 0 {
		refresh_rate = 60.0
	}
}
