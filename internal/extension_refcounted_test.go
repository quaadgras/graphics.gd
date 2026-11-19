//go:build !generate

package gd_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"graphics.gd/classdb"
	"graphics.gd/classdb/StyleBox"
	"graphics.gd/classdb/Theme"
	"graphics.gd/variant/Object"
)

// TestingHeldBox is a RefCounted extension class made in Go and handed to
// the engine.
type TestingHeldBox struct {
	StyleBox.Extension[TestingHeldBox]

	id int
}

var registerHeldBox sync.Once

func collectGarbage() {
	runtime.GC()
	// let the cleanups run before the frame ends
	if runtime.GOOS == "js" {
		for range 100 {
			runtime.Gosched()
		}
	} else {
		time.Sleep(time.Millisecond)
	}
	runtime.GC()
}

// TestGoRefCountedFollowsEngineReferences hands Go-made RefCounted values to
// the engine and drops them in Go: they must stay alive while the engine
// holds them, and be freed once it lets go (they used to stay rooted forever,
// and were reported as leaked at exit).
func TestGoRefCountedFollowsEngineReferences(t *testing.T) {
	registerHeldBox.Do(func() { classdb.Register[TestingHeldBox]() })
	const n = 8
	var theme Theme.Instance
	var ids [n]Object.ID
	runOnMain(t, func(t testing.TB) {
		theme = Object.Leak(Theme.New()) // kept across frames, freed at the end
		for i := range n {
			box := &TestingHeldBox{id: i}
			ids[i] = Object.Instance(box.AsObject()).ID()
			theme.SetStylebox(fmt.Sprint("box", i), "Button", box.AsStyleBox())
		}
		collectGarbage()
	})
	// twice: handing a value back to Go (GetStylebox) must not unroot it
	// while the theme still holds it.
	for range 2 {
		runOnMain(t, func(t testing.TB) {
			collectGarbage()
			for i := range n {
				box, ok := Object.As[*TestingHeldBox](theme.GetStylebox(fmt.Sprint("box", i), "Button"))
				if !ok || box.id != i {
					t.Fatalf("box %d was collected while the theme held it", i)
				}
			}
		})
	}
	runOnMain(t, func(t testing.TB) {
		theme.Clear()
		collectGarbage()
	})
	for range 3 { // the frames in which the frees happen
		runOnMain(t, func(t testing.TB) { collectGarbage() })
	}
	runOnMain(t, func(t testing.TB) {
		Object.Free(theme)
		for i, id := range ids {
			if id.Instance() != Object.Nil {
				t.Errorf("box %d was not freed after the theme released it", i)
			}
		}
	})
}
