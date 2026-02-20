package gdjson

// Virtual functions that require a ring buffer flush due to restrictions on what
// they can call.
var Flushables = map[string]bool{
	"CanvasItem._draw": true,
}
