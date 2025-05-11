package gdjson

var Renumeration = map[string]string{
	"Input.MouseMode":        "Input.MouseModeValue",
	"CameraServer.FeedImage": "CameraFeed.ImageType",
	"XRServer.TrackerType":   "XRTracker.Type",

	"Key":             "Input.Key",
	"KeyModifierMask": "Input.KeyModifierMask",
	"KeyLocation":     "Input.KeyLocation",
	"MouseButton":     "Input.MouseButton",
	"MouseButtonMask": "Input.MouseButtonMask",
	"JoyButton":       "Input.JoyButton",
	"JoyAxis":         "Input.JoyAxis",

	"Side":   "Rect2.Side",
	"Corner": "Rect2.Corner",

	"ClockDirection": "Angle.Direction",

	"Orientation":         "GUI.Orientation",
	"HorizontalAlignment": "GUI.HorizontalAlignment",
	"VerticalAlignment":   "GUI.VerticalAlignment",
	"InlineAlignment":     "GUI.InlineAlignment",

	"EulerOrder":  "Angle.Order",
	"MIDIMessage": "InputEventMIDI.Message",
	"Error":       "Error.Code",

	"PropertyHint":       "ClassDB.PropertyHint",
	"PropertyUsageFlags": "ClassDB.PropertyUsageFlags",
	"MethodFlags":        "ClassDB.MethodFlags",

	"Variant.Type":     "Variant.Type",
	"Variant.Operator": "Variant.Operator",
}
