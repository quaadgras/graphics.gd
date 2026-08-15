//go:build go1.26 && (amd64 || arm64)

// Package threadcheck provides a fast check for whether the current
// goroutine is running on the main OS thread.
//
// This works by reading the m (machine/OS-thread) pointer from the
// current g struct via the dedicated g register (R14 on amd64, R28
// on arm64). The m pointer is stable across different goroutines
// running on the same OS thread, unlike the g pointer which changes
// when Go assigns a different goroutine to a cgo callback.
//
// Guarded behind go1.26 because it depends on the offset of g.m
// (48 bytes on 64-bit) which is stable but internal to the runtime.
package threadcheck

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// currentm returns the m pointer (OS thread) from the current g struct.
// The g register holds the goroutine pointer, and g.m at offset 48
// points to the OS thread struct.
func currentm() uintptr

var mainM = currentm()

func Init() {
	mainM = currentm()
}

// Main reports whether the caller is running on the main OS thread.
func Main() bool {
	return currentm() == mainM
}

// FrameTemporaries reports whether reference wrappers created by the caller
// may be tracked as main-thread frame-temporaries (collected by the per-frame
// [pointers.Cycle]). Code running on the main thread is frame-aligned by
// construction; everything else must anchor its wrappers to the Go garbage
// collector instead (see gd.Wrap*).
func FrameTemporaries() bool {
	return Main()
}

// engineMs records the m pointers of OS threads owned by the engine: threads
// (other than main) on which the engine has called into Go, such as the
// dedicated resource-loading thread or WorkerThreadPool threads. Calls made
// back into the engine from these threads must stay on them (the engine may
// be blocked waiting for the work), so they are never routed through the
// cross-thread dispatch ring. Slots are append-only: an engine thread stays
// marked for the process lifetime.
var engineMs [64]atomic.Uintptr
var engineMu sync.Mutex

// hostCalls tracks, per OS thread (m), how deeply the thread is nested
// inside a Go-initiated engine call (see EnterCall). Open-addressed table
// keyed by the m pointer. Slots are claimed for the lifetime of the process
// (the number of OS threads is small and bounded). The counter is atomic
// and EnterCall locks the goroutine to its thread until the matching
// LeaveCall: without the lock a preemption between the two could migrate
// the goroutine to another m, incrementing one slot and decrementing
// another — the drifted counters then make inHostCall misreport, and
// Mark() misclassifies Go-owned threads as engine-owned, after which user
// goroutines bypass the cross-thread dispatch ring and race the engine
// directly (silently lost or corrupted calls).
var hostCalls [512]struct {
	m atomic.Uintptr
	n atomic.Int32
}

func hostCallSlot(m uintptr, alloc bool) int {
	for i, h := 0, (m>>4)&511; i < len(hostCalls); i, h = i+1, (h+1)&511 {
		v := hostCalls[h].m.Load()
		if v == m {
			return int(h)
		}
		if v == 0 {
			if !alloc {
				return -1
			}
			if hostCalls[h].m.CompareAndSwap(0, m) {
				return int(h)
			}
			// Another thread claimed this slot between the load and the
			// swap: keep probing.
		}
	}
	return -1
}

// EnterCall records that the current OS thread is entering the engine on
// behalf of Go code. Any engine→Go callbacks that fire before the matching
// LeaveCall are re-entrant on a Go-owned thread and must not Mark it as
// engine-owned. The goroutine stays locked to its thread until LeaveCall
// so the pair always hits the same slot (see hostCalls).
func EnterCall() {
	runtime.LockOSThread()
	if s := hostCallSlot(currentm(), true); s >= 0 {
		hostCalls[s].n.Add(1)
	}
}

// LeaveCall records that the current OS thread has returned from a
// Go-initiated engine call.
func LeaveCall() {
	if s := hostCallSlot(currentm(), false); s >= 0 {
		hostCalls[s].n.Add(-1)
	}
	runtime.UnlockOSThread()
}

func inHostCall(m uintptr) bool {
	s := hostCallSlot(m, false)
	return s >= 0 && hostCalls[s].n.Load() > 0
}

// Mark records the current OS thread as engine-owned. It is called on entry
// to every engine→Go callback: if the engine calls into Go on a thread of
// its own accord, that thread belongs to the engine. Callbacks that are
// re-entrant from a Go-initiated engine call (between EnterCall/LeaveCall)
// do not count: the engine is calling back on a thread Go owns.
// Cheap when already marked (or on main).
func Mark() {
	m := currentm()
	if m == mainM {
		return
	}
	if inHostCall(m) {
		return
	}
	for i := range engineMs {
		v := engineMs[i].Load()
		if v == m {
			return
		}
		if v == 0 {
			break
		}
	}
	engineMu.Lock()
	defer engineMu.Unlock()
	for i := range engineMs {
		v := engineMs[i].Load()
		if v == m {
			return
		}
		if v == 0 {
			engineMs[i].Store(m)
			return
		}
	}
	// Registry full: the thread is treated as a user thread and its calls
	// will go through the cross-thread dispatch ring, which is safe but slow.
}

// Engine reports whether the caller is running on an engine-owned OS thread
// (not counting the main thread, see Main).
func Engine() bool {
	m := currentm()
	for i := range engineMs {
		v := engineMs[i].Load()
		if v == 0 {
			return false
		}
		if v == m {
			return true
		}
	}
	return false
}
