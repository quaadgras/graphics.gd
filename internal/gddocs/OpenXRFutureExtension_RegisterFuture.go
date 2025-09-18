/*
var future_result = OpenXRFutureExtension.register_future(future)
await future_result.completed
if future_result.get_status() == OpenXRFutureResult.RESULT_FINISHED:
	# Handle your success
	pass
*/

package main

import (
	"graphics.gd/classdb/OpenXRFutureExtension"
	"graphics.gd/classdb/OpenXRFutureResult"
	"graphics.gd/variant/Signal"
)

var openXRFutureExtension OpenXRFutureExtension.Instance
var future int

func OpenXRFutureExtension_RegisterFuture() {
	var future_result = openXRFutureExtension.RegisterFuture(future)
	future_result.OnCompleted(func(result OpenXRFutureResult.Instance) {
		if result.GetStatus() == OpenXRFutureResult.ResultFinished {
			// Handle your success
		}
	}, Signal.OneShot)
}
