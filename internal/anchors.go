//go:build !generate

package gd

import (
	"runtime"
	"sync"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/threadcheck"
)

// Engine values created by code running on the main thread are tracked by the
// [pointers] tables and collected by the per-frame [pointers.Cycle]: they are
// frame-temporaries unless pinned or refreshed. That model is unsound for
// user goroutines and engine-owned threads, whose timelines are not aligned
// to frames — the main thread's Cycle would free a value the goroutine still
// holds, out from underneath it. Wrappers built off the main thread are
// therefore anchored instead: the engine value is pinned (invisible to
// Cycle) and its lifetime is handed to the Go garbage collector through a
// heap-allocated anchor carried inside the proxy value. When the proxy
// becomes unreachable, a [runtime.AddCleanup] cleanup releases the engine
// value — via [noescape.Free]'s cross-thread path, which queues the free in
// the dispatch ring FIFO behind any of the value's still-buffered uses.
//
// The anchor is only reachable through the proxy. Internal code that
// extracts the raw handle from a freshly-built proxy (for example, a
// generated binding converting a Go string into a StringName argument)
// drops the proxy before the buffered call is even pushed, so anchors are
// additionally kept alive in a two-generation epoch list that the frame
// drain cycles: an anchored value cannot be cleaned up before the second
// frame drain after its creation, by which point any use buffered by its
// creating expression has executed.
var anchors struct {
	mu   sync.Mutex
	keep [2][]any
}

func init() {
	// Entries born off the main thread are additionally protected by the
	// pointer nursery: the per-frame Cycle rescues them until they come of
	// age, because such a goroutine can be descheduled (or parked on the
	// cross-thread dispatch ring) across two cycles at ANY point — even
	// between a value's creation and its first use — which would otherwise
	// free the value out from underneath it. See pointers.OffMain.
	pointers.OffMain = func() bool { return !threadcheck.FrameTemporaries() }
}

// CycleAnchors rotates the anchor keep-alive generations. Called on the main
// thread once per frame, after the cross-thread dispatch ring has been
// drained (see startup/garbage_collector.go).
func CycleAnchors() {
	anchors.mu.Lock()
	defer anchors.mu.Unlock()
	recycled := anchors.keep[1]
	anchors.keep[1] = anchors.keep[0]
	anchors.keep[0] = recycled[:0]
}

// anchored pins the given freshly-constructed (tracked, off-main) value and
// returns a GC anchor for it along with the packed proxy state. The cleanup
// releases the engine value when the anchor becomes unreachable: owned
// values are freed (through the ring, ordered behind their buffered uses),
// borrowed ([pointers.Let]) values only release their table slot.
func anchored[T pointers.Generic[T, P], P pointers.Size](value T) (*T, complex128) {
	_, kind := pointers.Ask(value)
	if kind != pointers.Normal && kind != pointers.Letted {
		// Static and raw handles are not tracked by the pointer tables, the
		// per-frame Cycle cannot free them — nothing to anchor. Values that
		// are already pinned (generated callbacks pin their borrowed
		// arguments) have their lifetime managed by whoever pinned them;
		// attaching a cleanup would double-free them.
		return nil, pointers.Pack(value)
	}
	pinned := pointers.Pin(value)
	anchor := new(T)
	*anchor = pinned
	if kind == pointers.Letted {
		runtime.AddCleanup(anchor, func(raw T) { pointers.End(raw) }, pinned)
	} else {
		runtime.AddCleanup(anchor, func(raw T) { raw.Free() }, pinned)
	}
	anchors.mu.Lock()
	anchors.keep[0] = append(anchors.keep[0], anchor)
	anchors.mu.Unlock()
	return anchor, pointers.Pack(pinned)
}

// anchorTracked keeps a tracked engine value alive across a whole multi-call
// operation (converting an Array or Dictionary element by element). Off the
// main thread every element access can block on the cross-thread dispatch
// ring for a frame or more, while a tracked temporary only survives two of
// the main thread's pointer cycles — so a multi-element conversion outlives
// its source value and panics with "use of an invalid reference"
// (deterministically, once the ring is congested enough for each call to
// park a frame). The caller must hold the returned anchor with
// [runtime.KeepAlive] until the operation is done and use the returned
// (pinned) value in its place. On the main thread, or for values that are
// not tracked temporaries, the value is returned as-is with a nil anchor
// (KeepAlive(nil) is fine).
func anchorTracked[T pointers.Generic[T, P], P pointers.Size](value T) (*T, T) {
	if threadcheck.FrameTemporaries() {
		return nil, value
	}
	anchor, state := anchored(value)
	if anchor == nil {
		return nil, value
	}
	return anchor, pointers.Load[T](state)
}

// WrapString prepares proxy state for a [String] wrapper: tracked
// frame-temporary state where that model is sound (the main thread, see
// [threadcheck.FrameTemporaries]), anchored GC-managed state elsewhere.
// The remaining Wrap functions do the same for the other reference types;
// generated bindings and internal conversions must use these instead of
// packing tracked pointers directly.
func WrapString(s String) (StringProxy, complex128) {
	if threadcheck.FrameTemporaries() {
		return StringProxy{}, pointers.Pack(s)
	}
	anchor, state := anchored(s)
	return StringProxy{anchor: anchor}, state
}

// WrapStringName prepares proxy state for a [StringName] wrapper, see [WrapString].
func WrapStringName(s StringName) (StringNameProxy, complex128) {
	if threadcheck.FrameTemporaries() {
		return StringNameProxy{}, pointers.Pack(s)
	}
	anchor, state := anchored(s)
	return StringNameProxy{anchor: anchor}, state
}

// WrapNodePath prepares proxy state for a [NodePath] wrapper, see [WrapString].
func WrapNodePath(n NodePath) (NodePathProxy, complex128) {
	if threadcheck.FrameTemporaries() {
		return NodePathProxy{}, pointers.Pack(n)
	}
	anchor, state := anchored(n)
	return NodePathProxy{anchor: anchor}, state
}

// WrapArray prepares proxy state for an [Array] wrapper, see [WrapString].
func WrapArray[T any](a Array) (ArrayProxy[T], complex128) {
	if threadcheck.FrameTemporaries() {
		return ArrayProxy[T]{}, pointers.Pack(a)
	}
	anchor, state := anchored(a)
	return ArrayProxy[T]{anchor: anchor}, state
}

// WrapDictionary prepares proxy state for a [Dictionary] wrapper, see [WrapString].
func WrapDictionary[K comparable, V any](d Dictionary) (DictionaryProxy[K, V], complex128) {
	if threadcheck.FrameTemporaries() {
		return DictionaryProxy[K, V]{}, pointers.Pack(d)
	}
	anchor, state := anchored(d)
	return DictionaryProxy[K, V]{anchor: anchor}, state
}

// WrapPacked prepares proxy state for a packed array wrapper, see [WrapString].
func WrapPacked[P Packed[P, V], V gdextension.Packable](p P) (PackedProxy[P, V], complex128) {
	if threadcheck.FrameTemporaries() {
		return PackedProxy[P, V]{}, pointers.Pack[P, PackedPointers](p)
	}
	anchor, state := anchored[P, PackedPointers](p)
	return PackedProxy[P, V]{anchor: anchor}, state
}

// WrapPackedStrings prepares proxy state for a [PackedStringArray] wrapper, see [WrapString].
func WrapPackedStrings(p PackedStringArray) (PackedStringArrayProxy, complex128) {
	if threadcheck.FrameTemporaries() {
		return PackedStringArrayProxy{}, pointers.Pack(p)
	}
	anchor, state := anchored(p)
	return PackedStringArrayProxy{anchor: anchor}, state
}

// WrapVariant prepares proxy state for a [Variant] wrapper, see [WrapString].
func WrapVariant(v Variant) (VariantProxy, complex128) {
	if threadcheck.FrameTemporaries() {
		return VariantProxy{}, pointers.Pack(v)
	}
	anchor, state := anchored(v)
	return VariantProxy{anchor: anchor}, state
}

// WrapCallable prepares proxy state for a [Callable] wrapper, see [WrapString].
func WrapCallable(c Callable) (CallableProxy, complex128) {
	if threadcheck.FrameTemporaries() {
		return CallableProxy{}, pointers.Pack(c)
	}
	anchor, state := anchored(c)
	return CallableProxy{anchor: anchor}, state
}

// WrapSignal prepares proxy state for a [Signal] wrapper, see [WrapString].
func WrapSignal(s Signal) (SignalProxy, complex128) {
	if threadcheck.FrameTemporaries() {
		return SignalProxy{}, pointers.Pack(s)
	}
	anchor, state := anchored(s)
	return SignalProxy{anchor: anchor}, state
}
