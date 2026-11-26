package ring

import (
	"structs"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

type Entry struct {
	_      structs.HostLayout
	Object uintptr // gdextension.Object
	_      pad     // (the engine reads these two as 8-byte words: see pad)
	Method uintptr // gdextension.MethodForClass
	_      pad
	Shape  uint64     // gdextension.Shape
	Args   [256]byte  // copied packed args
	Result [64]byte   // result slot (for future phases)
	Refs   [16]uint16 // intra-buffer references (for future phases)
	Owner  uintptr    // back-pointer for future result offloading
	PC     uintptr    // Go caller PC, captured at enqueue time
}

const Size = 256 // power of 2
const Mask = Size - 1

type Ring struct {
	_    structs.HostLayout
	head uint32
	tail uint32
	// floor is the oldest slot that can still be in use. With no flush
	// running it equals tail. While a flush runs, the batch it claimed sits
	// below tail and is still being read by the engine: floor stays at the
	// start of the outermost running batch, so [floor, head) covers every
	// running batch as well as the pending entries.
	floor uint32
	// depth counts the flushes of this ring that are running: a call in a
	// batch can call back into Go, which buffers and flushes more calls.
	depth   uint32
	Entries [Size]Entry
}

var Main Ring

var CrashIndex uint32 = 0xFFFFFFFF

func init() {
	// C drains Main directly (gd_ring_drain reads head/tail and advances tail
	// past executed entries): the prefix layout is load-bearing on both sides.
	if unsafe.Offsetof(Main.Entries) != 16 {
		panic("ring: Ring prefix must be (head, tail, floor, depth uint32) — gd.c's gd_ring_buffer depends on it")
	}
	if unsafe.Offsetof(threadsShared.seq) != 16 || unsafe.Offsetof(threadsShared.kind) != 1040 {
		panic("ring: mpscShared layout drifted — gd.c's gd_mpsc_shared depends on it")
	}
}

func (r *Ring) Pending() bool {
	return r.head != r.tail
}

func (r *Ring) Buffer(object, method uintptr, shape uint64, args unsafe.Pointer, pc uintptr) {
	if r.head-r.floor >= Size {
		if r.full(object, method, shape, args, pc) {
			return
		}
	}
	e := &r.Entries[r.head&Mask]
	e.Object = object
	e.Method = method
	e.Shape = shape
	n := gdextension.Shape(shape).SizeArguments()
	if n > 0 && args != nil {
		copyArgs(&e.Args, args, n)
	}
	e.PC = pc
	r.head++
}

// overflow[d] carries a call made at flush depth d that found no room in the
// ring. Calls at the same depth run one after another, so they share it.
var overflow []*Entry

// full makes room for one more entry by flushing. That always works when no
// flush is running. Inside a flush (a call in the batch called back into Go)
// the running batches cannot be flushed away, they are waiting on us, so when
// they occupy the whole ring the call runs right now from an entry of its
// own instead, and full reports true.
//
//go:noinline
func (r *Ring) full(object, method uintptr, shape uint64, args unsafe.Pointer, pc uintptr) bool {
	r.Flush()
	if r.head-r.floor < Size {
		return false
	}
	d := int(r.depth)
	for len(overflow) <= d {
		overflow = append(overflow, new(Entry))
	}
	e := overflow[d]
	e.Object = object
	e.Method = method
	e.Shape = shape
	if n := gdextension.Shape(shape).SizeArguments(); n > 0 && args != nil {
		copyArgs(&e.Args, args, n)
	}
	e.PC = pc
	r.depth++
	flush(unsafe.Pointer(e), 0, 1)
	r.depth--
	return true
}

// copyArgs copies the packed argument bytes into an entry. The common
// argument sizes take exact fixed-size copies (direct moves, no memmove
// call): a variable-length copy costs more than the arguments themselves
// for the small packs nearly every buffered call carries.
func copyArgs(dst *[256]byte, args unsafe.Pointer, n int) {
	p := unsafe.Pointer(dst)
	switch n {
	case 4:
		*(*[4]byte)(p) = *(*[4]byte)(args)
	case 8:
		*(*[8]byte)(p) = *(*[8]byte)(args)
	case 12:
		*(*[12]byte)(p) = *(*[12]byte)(args)
	case 16:
		*(*[16]byte)(p) = *(*[16]byte)(args)
	case 24:
		*(*[24]byte)(p) = *(*[24]byte)(args)
	case 32:
		*(*[32]byte)(p) = *(*[32]byte)(args)
	default:
		copy(dst[:n], unsafe.Slice((*byte)(args), n))
	}
}

// Flush runs the pending entries, including the ones that callbacks leave
// behind while it runs (a signal handler's calls are not drained when the
// handler returns), so that the ring is empty when it returns: the caller is
// about to make a call that must come after all of them.
//
// A batch is claimed (tail moves past it) before entering the engine, so
// that a flush from a callback only sees the entries buffered since. The
// claimed slots are still in use until the batch returns: floor keeps Buffer
// off them (see Ring.floor). gd.c's gd_ring_main_flush is the same protocol,
// for the C-side drain.
func (r *Ring) Flush() {
	for r.head != r.tail {
		tail := r.tail
		head := r.head
		r.tail = head
		r.depth++
		flush(unsafe.Pointer(&r.Entries[0]), tail, head)
		r.depth--
		if r.depth == 0 {
			// nothing is running any more: only pending entries are in use.
			r.floor = r.tail
		} else if r.head == r.tail {
			// A nested batch has finished and nothing was buffered behind
			// it: hand its slots back, the batches still running leave
			// little room.
			r.head, r.tail = tail, tail
		}
	}
}
