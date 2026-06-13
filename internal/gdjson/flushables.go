package gdjson

// Flushables lists methods that require a ring buffer flush because they have
// observable side effects (e.g. adding children triggers Ready notifications).
var Flushables = map[string]bool{
	"CanvasItem._draw": true,
	"Node.add_child":   true,
	// Runs last in the glTF import pipeline, receiving the finished scene
	// root. User code typically mutates that tree here (adding children,
	// setting properties), buffering void calls into ring.Main. Godot frees
	// the GLTFState/root and finalises the import as soon as this returns, so
	// the buffer must be drained before then or the deferred flush ptrcalls a
	// freed object.
	"GLTFDocumentExtension._import_post": true,
}
