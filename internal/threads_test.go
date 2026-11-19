//go:build !generate

package gd_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/NavigationServer3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/ResourceLoader"
	"graphics.gd/classdb/SceneTree"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdreference"
	"graphics.gd/variant"
	"graphics.gd/variant/Array"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Vector2"
)

// These tests exercise the cross-thread dispatch ring (see
// internal/threadcheck/DESIGN.md): engine calls made from ordinary goroutines
// are queued and executed in order by the main thread. Before the ring, such
// calls crossed into the engine on whatever thread the goroutine happened to
// be running on and crashed (issue #260). Unlike the rest of the suite, they
// intentionally do NOT use runOnMain.

// TestGoroutineEngineCalls performs buffered void calls and blocking result
// calls from several goroutines at once.
func TestGoroutineEngineCalls(t *testing.T) {
	const goroutines = 8
	var wg sync.WaitGroup
	for g := range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			node := Node.New()
			name := fmt.Sprintf("goroutine_%d", g)
			node.SetName(name)                   // buffered void call
			if got := node.Name(); got != name { // blocking result call
				t.Errorf("expected name %q, got %q", name, got)
			}
		}()
	}
	wg.Wait()
}

// TestGoroutineValueRoundTrip round-trips value types through the queue,
// covering trivial (jumponly) methods, which must also be re-routed off the
// main thread so they observe the goroutine's earlier queued writes.
func TestGoroutineValueRoundTrip(t *testing.T) {
	const goroutines = 4
	var wg sync.WaitGroup
	for g := range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			node := Node2D.New()
			for i := range 8 {
				want := Vector2.New(g, i)
				node.SetPosition(want)
				if got := node.Position(); got != want {
					t.Errorf("expected position %v, got %v", want, got)
					return
				}
			}
		}()
	}
	wg.Wait()
}

// TestGoroutineDynamicCall covers the variadic variant-call path, which is
// dispatched to the main thread as a whole.
func TestGoroutineDynamicCall(t *testing.T) {
	finished := make(chan struct{})
	go func() {
		defer close(finished)
		object := Object.New()
		result := Object.Call(object, "get_class")
		if got := fmt.Sprint(result); got != "Object" {
			t.Errorf("expected class 'Object', got %q", got)
		}
	}()
	<-finished
}

// TestGoroutineSceneTree is the issue #260 scenario: manipulating nodes that
// are inside the scene tree from a goroutine. Called directly off-thread, the
// engine's thread guards reject this (and the cgo callback used to crash);
// queued through the dispatch ring, the calls run on the main thread where
// the guards are satisfied.
func TestGoroutineSceneTree(t *testing.T) {
	finished := make(chan struct{})
	go func() {
		defer close(finished)
		parent := Node.New()
		child := Node2D.New()
		parent.SetName("issue260")
		parent.AddChild(child.AsNode())
		SceneTree.Add(parent)
		parent.PropagateNotification(int(Node.NotificationPaused)) // the call from the issue
		child.SetPosition(Vector2.New(1, 2))
		if got := child.Position(); got != Vector2.New(1, 2) {
			t.Errorf("expected position (1, 2), got %v", got)
		}
		// SceneTree.Add falls back to the engine's deferred queue when the
		// tree is still setting up its first frames, so give the add a few
		// frames to land before judging it.
		inTree := parent.IsInsideTree()
		for range 120 {
			if inTree {
				break
			}
			waitFrames(1)
			inTree = parent.IsInsideTree()
		}
		if !inTree {
			t.Error("expected parent to be inside the scene tree")
		}
		parent.QueueFree()
	}()
	<-finished
}

// TestGoroutineMainThreadVisibility checks that once a goroutine's blocking
// call has returned, its earlier queued writes are visible to the main thread.
func TestGoroutineMainThreadVisibility(t *testing.T) {
	var node Node.Instance
	ready := make(chan struct{})
	go func() {
		defer close(ready)
		node = Node.New()
		node.SetName("crossthread")
		_ = node.Name() // blocks until the queue up to this point has executed
	}()
	<-ready
	runOnMain(t, func(t testing.TB) {
		if got := node.Name(); got != "crossthread" {
			t.Errorf("expected name %q, got %q", "crossthread", got)
		}
	})
}

// TestGoroutineThreadSafeSingletons calls thread-safe server singletons from
// goroutines: their generated bindings cross into the engine directly (the
// engine guards them internally) instead of queueing for the main thread.
func TestGoroutineThreadSafeSingletons(t *testing.T) {
	const goroutines = 4
	var wg sync.WaitGroup
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			canvas := RenderingServer.CanvasItemCreate()
			if canvas == 0 {
				t.Error("expected a valid canvas item RID")
			}
			RenderingServer.FreeRid(RID.Any(canvas))
			navmap := NavigationServer3D.MapCreate()
			if navmap == 0 {
				t.Error("expected a valid navigation map RID")
			}
			NavigationServer3D.FreeRid(RID.Any(navmap))
			if ResourceLoader.Exists("res://this_file_does_not_exist.tres", "") {
				t.Error("expected missing resource to not exist")
			}
		}()
	}
	wg.Wait()
}

// TestGoroutineDirectWhileMainBlocked proves the dispatch is direct: the
// main thread is parked inside a callable, so the cross-thread ring cannot
// drain — a queued call could never complete, but a thread-safe singleton
// call still does.
func TestGoroutineDirectWhileMainBlocked(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		finished := make(chan struct{})
		go func() {
			defer close(finished)
			ResourceLoader.Exists("res://this_file_does_not_exist.tres", "")
		}()
		select {
		case <-finished:
		case <-time.After(10 * time.Second):
			t.Error("thread-safe singleton call did not complete while the main thread was blocked")
		}
	})
}

// TestGoroutineFollowUpThroughput checks that a goroutine making sequential
// result calls is not limited to one call per frame: the end-of-frame drain
// keeps a follow-up window open (sized from the frame period) that services
// its next calls within the same frame.
func TestGoroutineFollowUpThroughput(t *testing.T) {
	const calls = 50
	finished := make(chan struct{})
	var frames int
	go func() {
		defer close(finished)
		node := Node.New()
		node.SetName("throughput")
		before := Engine.GetProcessFrames()
		for range calls {
			if got := node.Name(); got != "throughput" { // blocking result call
				t.Errorf("expected name %q, got %q", "throughput", got)
				return
			}
		}
		frames = Engine.GetProcessFrames() - before
	}()
	<-finished
	if frames >= calls {
		t.Errorf("%d sequential calls took %d frames: follow-up window is not batching them", calls, frames)
	}
}

// TestGoroutineTemporaryGC creates and immediately drops objects on a
// goroutine while forcing garbage collection: the GC-driven frees must queue
// behind the dropped wrappers' still-buffered calls (regression: they used
// to run as deferred callables, which can execute before the frame drain,
// freeing an object whose queued call had not run yet).
func TestGoroutineTemporaryGC(t *testing.T) {
	finished := make(chan struct{})
	go func() {
		defer close(finished)
		for i := range 200 {
			node := Node2D.New()
			node.AsNode().SetName(fmt.Sprintf("temporary_%d", i))
			node.SetPosition(Vector2.New(i, i)) // buffered; wrapper dropped right after
			if i%10 == 0 {
				runtime.GC()
			}
		}
		runtime.GC()
	}()
	<-finished
}

// waitFrames blocks until at least the given number of engine frames have
// elapsed, giving the main thread's per-frame garbage collection every
// opportunity to (incorrectly) collect values the calling goroutine holds.
func waitFrames(frames int) {
	start := Engine.GetProcessFrames()
	for Engine.GetProcessFrames() < start+frames {
		time.Sleep(time.Millisecond)
	}
}

// TestGoroutineObjectAnchor covers the guarantee the generated bindings rely
// on to keep their receiver alive across an engine call: an object created off
// the main thread is freed by a runtime.AddCleanup on its wrapper, and that
// cleanup cannot run while the wrapper's one-word anchor is still reachable.
//
// Without it the bindings have a use-after-free window they cannot close. The
// receiver is dead, as far as the collector is concerned, from the moment
// gd.ObjectChecked hands its raw pointer to the call — and off the main thread
// the call is only *recorded* in the dispatch ring at that point. A collection
// inside that window queues the free first, so the main thread frees the object
// before running the call that uses it. On Android, where the whole suite runs
// off the main thread, that surfaced as an "invalid reference" panic under GC
// pressure.
//
// Only objects that gdreference.OwnObject built off the main thread carry a
// cleanup at all. On wasm threadcheck.Main is unconditionally true (one OS
// thread means thread identity cannot tell the frame-driving code from a free
// goroutine), so every object there is a pooled frame temporary whose lifetime
// the anchor does not govern, and there is nothing for this test to assert.
func TestGoroutineObjectAnchor(t *testing.T) {
	finished := make(chan struct{})
	var owned bool
	go func() {
		defer close(finished)
		node := Node.New()
		if _, kind := gdreference.AskObject(node.AsObject()[0]); kind != gdreference.TypeThread {
			return // frame temporary, not GC-anchored — skipped below
		}
		owned = true
		id := Object.Instance(node.AsObject()).ID()
		anchor := node.AsObject()[0].Anchor()
		if anchor == nil {
			t.Error("an object created off the main thread should carry a collectable anchor")
			return
		}
		node = Node.Nil // the wrapper is unreachable from here on, only the anchor is

		// Reachable anchor: the cleanup cannot run, so the object survives
		// every collection and every frame of the main thread's own cycle.
		for range 3 {
			runtime.GC()
			waitFrames(2)
		}
		if Object.ID(id).Instance() == Object.Nil {
			t.Error("object was freed while its anchor was still reachable")
		}
		runtime.KeepAlive(anchor)

		// Unreachable anchor: the cleanup runs and the free is queued behind
		// the calls already recorded, so the object goes away. Polled, because
		// it takes a collection and a drain, neither of which is synchronous.
		var freed bool
		for range 10 {
			runtime.GC()
			waitFrames(2)
			if Object.ID(id).Instance() == Object.Nil {
				freed = true
				break
			}
		}
		if !freed {
			t.Error("object was not freed after its anchor became unreachable")
		}
	}()
	<-finished
	if !owned {
		// Skipped from the test goroutine: t.Skip may not be called from any
		// other one.
		t.Skip("objects here are pooled frame temporaries, their lifetime is not anchored")
	}
}

// TestGoroutineReferenceLifetimes holds one of each engine-backed reference
// type on a goroutine across several frames (and Go collections) before
// using it. Wrappers created off the main thread are anchored to the Go
// garbage collector instead of participating in the main thread's
// frame-temporary collection, which is not aware of goroutine timelines and
// used to free these values underneath their holder (regression: temporary
// strings created on loader threads were collected mid-use by the main
// thread's pointers.Cycle).
func TestGoroutineReferenceLifetimes(t *testing.T) {
	finished := make(chan struct{})
	go func() {
		defer close(finished)
		node := Node.New()
		node.SetName("lifetimes")
		image := Image.CreateEmpty(2, 2, false, Image.FormatRgba8)
		var (
			name      = Node.Advanced(node).GetName()                             // String.Name
			path      = Node.Advanced(node).GetPathTo([1]gdclass.Node(node), false) // Path.ToNode
			renamed   = Node.Advanced(node).Renamed()                             // Signal.Any
			host      = OS.Advanced().GetName()                                   // String.Readable
			copyright = Engine.Advanced().GetCopyrightInfo()                      // Array.Contains[Dictionary.Any]
			numbers   = gd.ArrayFromSlice[Array.Contains[int64]]([]int64{10, 20, 30})
			data      = Image.Advanced(image).GetData()                           // Packed.Bytes
			class     = Object.Call(Object.New(), "get_class")                    // variant.Any
		)
		runtime.GC() // the values must survive the Go collector (they are still referenced) ...
		waitFrames(4)
		runtime.GC() // ... and several frames of the main thread's per-frame collection.

		if got := name.String(); got != "lifetimes" {
			t.Errorf("expected name %q, got %q", "lifetimes", got)
		}
		if got := path.String(); got != "." {
			t.Errorf("expected path %q, got %q", ".", got)
		}
		if got := renamed.Name().String(); got != "renamed" {
			t.Errorf("expected signal name %q, got %q", "renamed", got)
		}
		if host.String() == "" {
			t.Error("expected a non-empty host OS name")
		}
		if copyright.Len() == 0 {
			t.Error("expected non-empty copyright info")
		} else if got := fmt.Sprint(copyright.Index(0).Index(variant.New("name"))); got == "" {
			t.Error("expected a name in the first copyright entry")
		}
		if numbers.Len() != 3 || numbers.Index(2) != 30 {
			t.Errorf("expected array [10 20 30], got length %d", numbers.Len())
		}
		if data.Len() != 16 { // 2x2 RGBA8
			t.Errorf("expected 16 bytes of image data, got %d", data.Len())
		}
		if got := fmt.Sprint(class); got != "Object" {
			t.Errorf("expected class %q, got %q", "Object", got)
		}
	}()
	<-finished
}

// TestGoroutineArrayProxyCache passes a Go-native array to an engine call
// from a goroutine: the conversion caches an engine-backed proxy inside the
// user's array value, which must then remain valid on the goroutine after
// the frame-temporary collection has run.
func TestGoroutineArrayProxyCache(t *testing.T) {
	finished := make(chan struct{})
	go func() {
		defer close(finished)
		numbers := Array.New[int64]()
		for i := range int64(5) {
			numbers.Append(i * 11)
		}
		object := Object.New()
		Object.Call(object, "set_meta", variant.New("numbers"), variant.New(numbers))
		runtime.GC()
		waitFrames(4)
		if numbers.Len() != 5 {
			t.Errorf("expected rebound array to have 5 elements, got %d", numbers.Len())
		} else if got := numbers.Index(4); got != 44 {
			t.Errorf("expected element 4 to be 44, got %d", got)
		}
	}()
	<-finished
}

// TestGoroutineStress floods the ring well past its capacity from many
// goroutines at once, exercising backpressure against the frame-boundary
// drain.
func TestGoroutineStress(t *testing.T) {
	const goroutines = 16
	const iterations = 400 // goroutines*iterations >> ring size
	var wg sync.WaitGroup
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			node := Node2D.New()
			for i := range iterations {
				node.SetPosition(Vector2.New(i, i)) // fire and forget
			}
			want := Vector2.New(iterations-1, iterations-1)
			if got := node.Position(); got != want {
				t.Errorf("expected position %v, got %v", want, got)
			}
		}()
	}
	wg.Wait()
}
