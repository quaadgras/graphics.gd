package gd_test

// Regression tests for the lost-SetName CI flake (windows TestObjectIDs
// "expected name 'test', got ''"; TestGraphicsPackageGeneration losing the
// "Bullets" node name to the @Node2D@N auto-name — root causes were data
// races that let user-goroutine calls bypass or misuse the cross-thread
// dispatch ring). The file sorts first so these run at the top of the
// suite, where the observed losses clustered.

import (
	"fmt"
	"testing"

	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/internal/ring"
	"graphics.gd/variant/Object"
)

// TestSetNameReadback hammers the exact TestObjectIDs pattern: queue a
// fire-and-forget SetName from the test goroutine, then read the name back
// (a blocking call FIFO-ordered behind it).
func TestSetNameReadback(t *testing.T) {
	for i := 0; i < 300; i++ {
		node := Node.New()
		want := fmt.Sprintf("N%d", i)
		node.SetName(want)
		if got := node.Name(); got != want {
			buffered, executed := ring.Threads.Counters()
			t.Errorf("iteration %d: SetName lost: got %q (ring kindCall buffered=%d executed=%d, diff=%d)",
				i, got, buffered, executed, int64(buffered)-int64(executed))
		}
		Object.Free(node)
	}
}

// TestSetNameChildReadback hammers the makeSceneMain pattern: name a node,
// add it to a parent, and check the name survived (a lost SetName shows up
// as Godot's @Class@N auto-name assigned by add_child).
func TestSetNameChildReadback(t *testing.T) {
	for i := 0; i < 300; i++ {
		root := Node2D.New()
		child := Node2D.New()
		want := fmt.Sprintf("C%d", i)
		child.AsNode().SetName(want)
		root.AsNode().AddChild(child.AsNode())
		if got := child.AsNode().Name(); got != want {
			buffered, executed := ring.Threads.Counters()
			t.Errorf("iteration %d: SetName lost: child is %q (ring kindCall buffered=%d executed=%d, diff=%d)",
				i, got, buffered, executed, int64(buffered)-int64(executed))
		}
		Object.Free(root)
	}
}
