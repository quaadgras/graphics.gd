//go:build cgo && !(gd && wasm)

package ring

// #include <stdint.h>
// extern void gd_ring_flush(void *entries, uint32_t tail, uint32_t head, uint32_t *crash_index);
// extern void *gd_ring_flush_g0_addr(void);
// extern void gd_ring_adopt(void *ring, uint32_t *crash_index, void *threads_shared, void *threads_entries);
import "C"
import (
	"structs"
	"unsafe"

	"graphics.gd/internal/fastcb"
	"graphics.gd/internal/threadcheck"
)

// flushPack mirrors gd.c's struct gd_ring_flush_pack. A Go global, because the
// resident flush path (fastcb.CallC) must not hand C a pointer into a
// goroutine stack: callbacks triggered by flushed entries can move it.
// Main-thread only, and gd_ring_flush is re-entrancy-safe by tail/head
// snapshot, so a single pack suffices: a re-entrant flush from a nested
// callback overwrites the pack only after the outer C call has consumed it
// (gd_ring_flush_g0 reads the fields before the first ptrcall).
var flushPack struct {
	_          structs.HostLayout
	entries    unsafe.Pointer
	tail, head uint32
	crashIndex *uint32
}

// flushG0Addr holds gd_ring_flush_g0's address (a constant), resolved at init.
var flushG0Addr = C.gd_ring_flush_g0_addr()

func init() { flushPack.crashIndex = &CrashIndex }

// adopted tracks whether Main has been handed to C. Main-thread only.
var adopted bool

// Adopt hands Main to the C side, which from then on drains it at engine->Go
// callback boundaries (see gd_ring_drain in gd.c): entries a callback buffered
// execute in C as the callback returns, without a Go->C crossing to pay for
// the flush. Must be called on the main thread (C records it as the ring's
// owner thread); idempotent.
func Adopt() {
	if adopted || !threadcheck.Main() {
		return
	}
	adopted = true
	C.gd_ring_adopt(unsafe.Pointer(&Main), (*C.uint32_t)(unsafe.Pointer(&CrashIndex)),
		unsafe.Pointer(&threadsShared), unsafe.Pointer(&threadsEntries))
}

func flush(entries unsafe.Pointer, tail, head uint32) {
	if threadcheck.Main() && fastcb.Resident() {
		flushPack.entries = entries
		flushPack.tail = tail
		flushPack.head = head
		fastcb.CallC(flushG0Addr, unsafe.Pointer(&flushPack))
		return
	}
	C.gd_ring_flush(entries, C.uint32_t(tail), C.uint32_t(head), (*C.uint32_t)(unsafe.Pointer(&CrashIndex)))
}
