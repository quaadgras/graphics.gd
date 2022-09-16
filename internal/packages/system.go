package packages

// System specific layers with shared components (Unix, Win32, Apple, etc.) in drivers.
var System = []string{
	"ConfigFile",
	"Directory",
	"File",
	"Mutex",
	"OS",
	"Semaphore",
	"Shortcut",
	"Thread",
	"Time",
	"Timer",
	"WorkerThreadPool",
}

// SystemInput layer and events.
var SystemInput = []string{
	"Input",
	"InputEvent",
	"InputEventAction",
	"InputEventFromWindow",
	"InputEventGesture",
	"InputEventJoypadButton",
	"InputEventJoypadMotion",
	"InputEventKey",
	"InputEventMagnifyGesture",
	"InputEventMIDI",
	"InputEventMouse",
	"InputEventMouseButton",
	"InputEventMouseMotion",
	"InputEventPanGesture",
	"InputEventScreenDrag",
	"InputEventScreenTouch",
	"InputEventShortcut",
	"InputEventWithModifiers",
	"InputMap",
}
