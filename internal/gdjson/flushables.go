package gdjson

// Flushables lists methods that require a ring buffer flush because they have
// observable side effects (e.g. adding children triggers Ready notifications).
var Flushables = map[string]bool{
	"CanvasItem._draw": true,
	"Node.add_child":   true,
}
