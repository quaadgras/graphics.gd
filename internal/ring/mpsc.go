package ring

import (
	"runtime"
	"structs"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

// Threads is the cross-thread (MPSC) command ring described in DESIGN.md.
// Goroutines that are neither on the main thread nor on an engine-owned
// thread record their engine calls here instead of crossing into the engine
// directly (which is not thread-safe and crashes, see issue #260). The main
// thread drains the ring during Flush, executing the buffered calls in FIFO
// order on the thread the engine expects them on.
var Threads MPSC

// threadsEntries lives outside the MPSC struct: the entries are handed to C
// during the flush, so they must not share an allocation with Go pointers
// (the cgo pointer check rejects the argument otherwise).
var threadsEntries [Size]Entry

// mpscShared is the C-visible, pointer-free part of the MPSC state. It lives
// outside the MPSC struct for the same cgo-pointer-check reason as
// threadsEntries: it is handed to C once (see Adopt), so the C-side drain can
// consume published fire-and-forget entries at engine->Go callback boundaries
// without a crossing. gd.c's gd_mpsc_shared mirrors this layout exactly.
type mpscShared struct {
	_    structs.HostLayout
	head atomic.Uint32 // next index to claim (producers, any thread)

	// cursor is the drain position. Main-thread only — advanced by both the
	// Go drain and the C drain, which run on the same thread.
	cursor uint32

	// draining is nonzero while a Go-side drain is in flight: the C drain
	// must not consume entries the Go drain has read but not yet retired.
	// Main-thread only.
	draining uint32

	// cdrained is set by the C drain when it released slots: producers that
	// wrapped a full lap park on the Go-side cond, which C cannot signal, so
	// the next Go drain broadcasts on their behalf. Main-thread only.
	cdrained uint32

	// seq holds each slot's lifecycle position, see MPSC.
	seq [Size]atomic.Uint32

	// kind classifies each published entry for the C drain: kindCall entries
	// are plain engine calls C may execute and release; kindGo entries
	// (thunks and parked calls) require the Go drain — C stops at the first
	// one to preserve FIFO order. Written by the producer before publishing,
	// ordered by seq.
	kind [Size]uint8

	// executed counts fire-and-forget (kindCall) entries executed by either
	// drain, for loss diagnostics: compared against the producers' published
	// count (MPSC.buffered), a lower value means a queued call vanished
	// without running. Appended after kind so the C mirror's asserted
	// offsets are unchanged; gd.c increments it in its drain. Both drains
	// run on the main thread; atomic so [Counters] may read from any thread.
	executed atomic.Uint64
}

const (
	kindCall = 0 // fire-and-forget engine call: the C drain may execute it
	kindGo   = 1 // thunk or parked call: Go drain only
)

var threadsShared mpscShared

func init() {
	Threads.Init(&threadsShared, &threadsEntries)
}

// The follow-up window is how long the end-of-frame drain keeps polling for
// follow-up calls after it released blocked goroutines. A goroutine making
// sequential result calls is otherwise limited to one call per frame: the
// drain wakes it, finds the ring empty and returns, and its next call waits
// a whole frame. Within the window, the released goroutine's next call is
// picked up microseconds later instead. The window is a fixed deadline per
// frame, so it bounds the extra frame time even against a goroutine that
// keeps calling — and since the frame callback runs before the engine's
// frame-delay sleep, an idle main thread pays for the window out of time it
// would have slept anyway.
//
// The window is sized dynamically from the measured frame period (see
// FlushFrame): period/FollowUpRatio, clamped to [FollowUpMin, FollowUpMax].
// A capped or vsynced game affords a bigger window (paid out of the frame
// sleep) than a fast uncapped loop.
var (
	FollowUpRatio = 8 // fraction of the frame period spent on follow-ups
	FollowUpMin   = 50 * time.Microsecond
	FollowUpMax   = 2 * time.Millisecond
)

// ResultSize is the maximum size of a buffered call's return value, matching
// the result slot of [Entry].
const ResultSize = 64

// MPSC is a multiple-producer single-consumer command ring. Any goroutine
// may claim and publish entries; only the main thread may Flush.
//
// Each slot moves through a sequence-numbered lifecycle so that no party
// ever waits on another to make progress within a slot, and results are
// delivered in place (no per-call allocation or copy by the drainer). For
// index i in slot s = i&Mask, seq[s] reads:
//
//	i        free: the producer of index i may fill it (a producer that
//	         laps the ring onto a slot that has not been released yet
//	         parks here — that is the backpressure)
//	i+1      published: the drain may execute it
//	i+2      executed: the result is in the entry; only the goroutine
//	         parked on the call reads it back and releases the slot
//	i+Size   released: free for the producer of index i+Size
//
// The drain itself releases the slots of fire-and-forget entries, and skips
// past executed-but-unread slots without waiting: the main thread never
// waits on a goroutine. An unread slot only stalls the one producer that
// wraps around onto it a full lap later, until its reader wakes and
// releases it.
type MPSC struct {
	_ structs.HostLayout

	// shared is the C-visible state: head, seq, cursor and the entry-kind
	// table (see mpscShared). Everything below is Go-side only.
	shared *mpscShared

	// parked marks entries whose producing goroutine blocks on completion
	// (Call/Run): the drain leaves their slot in the executed state for the
	// goroutine to read back and release. Written by the producer before
	// publishing, read by the drain after the publish, so it is ordered by
	// seq. Go-side only: parked entries are kindGo, so the C drain skips them.
	parked [Size]bool

	// thunks holds the Go function of Run/Defer entries, which execute on
	// the main thread between engine crossings instead of being dispatched
	// to C. Written by the producer before publishing, cleared by the drain.
	thunks [Size]func()

	// panics holds the value a Run thunk panicked with, transferred to the
	// parked caller so the panic surfaces on the goroutine that made the
	// call, as it would have for a direct call. Written by the drain before
	// the executed transition, consumed by the caller before release.
	panics [Size]any

	// closed poisons the ring at engine shutdown (see Close): parked
	// goroutines wake up and cross-thread calls become no-ops, instead of
	// parking forever now that nothing drains.
	closed atomic.Bool

	// buffered counts fire-and-forget entries published via [Buffer], the
	// producer side of the loss diagnostic (see mpscShared.executed).
	buffered atomic.Uint64

	mu   sync.Mutex
	cond *sync.Cond

	// frame-period tracking for the dynamic follow-up window; main-thread only.
	lastFrame   time.Time
	framePeriod time.Duration

	// entries is kept in a separate pointer-free allocation (see
	// threadsEntries) so it can be passed to C during the flush.
	entries *[Size]Entry
}

// Init prepares the ring: slot s starts out free for index s.
func (r *MPSC) Init(shared *mpscShared, entries *[Size]Entry) {
	r.shared = shared
	r.entries = entries
	r.cond = sync.NewCond(&r.mu)
	for s := range r.shared.seq {
		r.shared.seq[s].Store(uint32(s))
	}
}

// dispatch executes entries [tail, head) of a ring in one engine crossing.
// It is a variable so unit tests can substitute an executor that does not
// require a linked engine.
var dispatch = flush

func (r *MPSC) Pending() bool {
	return r.shared.head.Load() != r.shared.cursor
}

// claim reserves the next index. Claims are tickets: the producer parks only
// if its slot has not been released by the previous lap yet. Reports false
// once the ring is closed: the caller must abandon the call.
func (r *MPSC) claim() (uint32, bool) {
	if r.closed.Load() {
		return 0, false
	}
	i := r.shared.head.Add(1) - 1
	if r.shared.seq[i&Mask].Load() != i {
		r.mu.Lock()
		for r.shared.seq[i&Mask].Load() != i {
			if r.closed.Load() {
				r.mu.Unlock()
				return 0, false
			}
			r.cond.Wait()
		}
		r.mu.Unlock()
	}
	return i, true
}

func (r *MPSC) fill(i uint32, object, method uintptr, shape uint64, args unsafe.Pointer, pc uintptr) {
	e := &r.entries[i&Mask]
	e.Object = object
	e.Method = method
	e.Shape = shape
	n := gdextension.Shape(shape).SizeArguments()
	if n > 0 && args != nil {
		copyArgs(&e.Args, args, n)
	}
	e.PC = pc
}

// publish makes slot i visible to the drain. The entry, parked flag, kind and
// thunk must be fully written beforehand.
func (r *MPSC) publish(i uint32) {
	r.shared.seq[i&Mask].Store(i + 1)
}

// await parks the calling goroutine until the drain has executed index i,
// reporting false if the ring closed before that happened (the entry will
// never execute; the slot must not be released or read).
func (r *MPSC) await(i uint32) bool {
	r.mu.Lock()
	for r.shared.seq[i&Mask].Load() != i+2 {
		if r.closed.Load() {
			r.mu.Unlock()
			return false
		}
		r.cond.Wait()
	}
	r.mu.Unlock()
	return true
}

// release recycles slot i for the next lap, waking a producer that may have
// wrapped around onto it. The wake-up is only needed when a producer has
// already claimed a full lap ahead.
func (r *MPSC) release(i uint32) {
	r.shared.seq[i&Mask].Store(i + Size)
	if r.shared.head.Load()-i >= Size {
		r.mu.Lock()
		r.cond.Broadcast()
		r.mu.Unlock()
	}
}

// Buffer records a void engine call (fire and forget). The argument bytes are
// copied into the ring, so the caller does not need to keep them alive.
// Ordering with respect to other calls from the same goroutine is preserved,
// including frees of the values referenced by the arguments: those are engine
// calls themselves and are queued behind this entry.
func (r *MPSC) Buffer(object, method uintptr, shape uint64, args unsafe.Pointer, pc uintptr) {
	i, ok := r.claim()
	if !ok {
		return
	}
	r.fill(i, object, method, shape, args, pc)
	r.parked[i&Mask] = false
	r.thunks[i&Mask] = nil
	r.shared.kind[i&Mask] = kindCall
	r.publish(i)
	r.buffered.Add(1)
}

// Counters reports how many fire-and-forget entries have been published and
// how many either drain has executed, for loss diagnostics: after a Barrier
// (or any completed blocking call) the two are equal unless a queued call
// vanished without running.
func (r *MPSC) Counters() (buffered, executed uint64) {
	return r.buffered.Load(), r.shared.executed.Load()
}

// Call records an engine call and blocks until the main thread has executed
// it, copying up to size bytes of the return value out of the ring into
// result. The slot is released by this goroutine once the result has been
// read back; the main thread never waits on it.
func (r *MPSC) Call(object, method uintptr, shape uint64, args unsafe.Pointer, pc uintptr, result unsafe.Pointer, size uintptr) {
	i, ok := r.claim()
	if !ok {
		return // engine shut down: the caller sees a zero result
	}
	r.fill(i, object, method, shape, args, pc)
	r.parked[i&Mask] = true
	r.thunks[i&Mask] = nil
	r.shared.kind[i&Mask] = kindGo
	r.publish(i)
	if !r.await(i) {
		return
	}
	if size > 0 && result != nil {
		copy(unsafe.Slice((*byte)(result), size), r.entries[i&Mask].Result[:size])
	}
	r.release(i)
}

// Run queues fn to run on the main thread, in FIFO order with the buffered
// engine calls, and blocks until it has executed. It covers calls that cannot
// be encoded as a ring entry (such as variadic variant calls).
func (r *MPSC) Run(fn func()) {
	i, ok := r.claim()
	if !ok {
		return // engine shut down: fn is not executed
	}
	r.parked[i&Mask] = true
	r.thunks[i&Mask] = fn
	r.shared.kind[i&Mask] = kindGo
	r.publish(i)
	if !r.await(i) {
		return
	}
	s := i & Mask
	failure := r.panics[s]
	r.panics[s] = nil
	r.release(i)
	if failure != nil {
		// surface fn's panic on the goroutine that made the call, as it
		// would have for a direct call.
		panic(failure)
	}
}

// Defer queues fn to run on the main thread, in FIFO order with the buffered
// engine calls, without waiting for it. It is used for destructors: queueing
// a value's free behind its buffered uses keeps the value alive until every
// queued call that references it has executed.
func (r *MPSC) Defer(fn func()) {
	i, ok := r.claim()
	if !ok {
		return // engine shut down: fn is not executed
	}
	r.parked[i&Mask] = false
	r.thunks[i&Mask] = fn
	r.shared.kind[i&Mask] = kindGo
	r.publish(i)
}

// Barrier is a flush for producers: only the main thread may drain the ring,
// so "flushing" from a side thread means parking until the next drain has
// executed everything published before this call. It queues a parked no-op
// behind the calling goroutine's buffered entries and waits for it, so a
// direct engine crossing made after Barrier returns observes the effects of
// all of this goroutine's buffered calls. Returns immediately if the ring is
// closed.
func (r *MPSC) Barrier() {
	r.Run(func() {})
}

// Flush drains the ring. Main thread only. Entries are executed in FIFO
// order; runs of engine calls are dispatched in a single crossing, with
// thunks executed in Go between them. A published entry can trigger engine
// callbacks into Go while it executes; if that Go code flushes again the
// nested flush is a no-op (the drain is already in progress and per-goroutine
// FIFO order must not be broken by processing later entries first).
func (r *MPSC) Flush() {
	r.FlushFor(0)
}

// FlushFrame drains the ring at a frame boundary, keeping a follow-up
// window open that is sized from the measured frame period: an exponential
// moving average of the time between FlushFrame calls, divided by
// FollowUpRatio and clamped to [FollowUpMin, FollowUpMax].
func (r *MPSC) FlushFrame() {
	now := time.Now()
	if !r.lastFrame.IsZero() {
		period := now.Sub(r.lastFrame)
		if r.framePeriod == 0 {
			r.framePeriod = period
		} else {
			r.framePeriod += (period - r.framePeriod) / 8
		}
	}
	r.lastFrame = now
	window := r.framePeriod / time.Duration(FollowUpRatio)
	if window < FollowUpMin {
		window = FollowUpMin
	}
	if window > FollowUpMax {
		window = FollowUpMax
	}
	r.FlushFor(window)
}

// FlushFor drains the ring like Flush and then, if goroutines blocked on
// results were released, keeps polling for their follow-up calls until the
// window closes. See the FollowUp variables for why: without the window, a
// goroutine making sequential result calls completes only one per drain.
func (r *MPSC) FlushFor(window time.Duration) {
	if r.shared.draining != 0 {
		return
	}
	r.shared.draining = 1
	defer func() { r.shared.draining = 0 }()
	// The C drain releases slots without signalling (it cannot touch the Go
	// cond): wake producers that wrapped a lap and parked on those slots.
	if r.shared.cdrained != 0 {
		r.shared.cdrained = 0
		r.mu.Lock()
		r.cond.Broadcast()
		r.mu.Unlock()
	}
	if !r.drain() || window <= 0 {
		return
	}
	deadline := time.Now().Add(window)
	for {
		for !r.Pending() {
			if time.Now().After(deadline) {
				return
			}
			runtime.Gosched()
		}
		r.drain()
		if time.Now().After(deadline) {
			return
		}
	}
}

// drain processes everything published in the ring, reporting whether any
// blocked goroutine was released. Main thread only, r.draining must be held.
// The drain never waits: slots of blocked callers are left in the executed
// state for their goroutine to read back and release, and the scan simply
// stops at the first slot that is not published yet (which includes a slot
// from the previous lap whose reader has not released it — FIFO order is
// preserved either way).
func (r *MPSC) drain() (released bool) {
	for {
		cursor := r.shared.cursor
		head := r.shared.head.Load()
		end := cursor
		for end != head && r.shared.seq[end&Mask].Load() == end+1 {
			end++
		}
		if end == cursor {
			return
		}
		for i := cursor; i != end; {
			j := i
			for j != end && r.thunks[j&Mask] == nil {
				j++
			}
			if j != i {
				dispatch(unsafe.Pointer(&r.entries[0]), i, j)
				for k := i; k != j; k++ {
					s := k & Mask
					if !r.parked[s] {
						r.shared.executed.Add(1)
					}
					if r.parked[s] {
						// executed: the parked goroutine reads the result
						// out of the entry and releases the slot itself.
						released = true
						r.shared.seq[s].Store(k + 2)
					} else {
						r.shared.seq[s].Store(k + Size)
					}
				}
			} else {
				s := i & Mask
				fn := r.thunks[s]
				r.thunks[s] = nil
				failure := protect(fn)
				j = i + 1
				if r.parked[s] {
					// transfer a panic to the parked caller (see Run); it
					// consumes panics[s] before releasing the slot.
					r.panics[s] = failure
					released = true
					r.shared.seq[s].Store(i + 2)
				} else {
					r.shared.seq[s].Store(i + Size)
					if failure != nil {
						// nobody is waiting on a deferred thunk: keep the
						// ring consistent, then let the panic propagate on
						// the main thread.
						r.shared.cursor = j
						r.mu.Lock()
						r.cond.Broadcast()
						r.mu.Unlock()
						panic(failure)
					}
				}
			}
			r.shared.cursor = j
			r.mu.Lock()
			r.cond.Broadcast()
			r.mu.Unlock()
			i = j
		}
	}
}

// protect runs fn, capturing the value it panicked with, if any.
func protect(fn func()) (failure any) {
	defer func() { failure = recover() }()
	fn()
	return nil
}

// Close drains any remaining buffered calls while the engine is still alive
// and then poisons the ring: goroutines parked on it wake up, and subsequent
// cross-thread calls become no-ops with zero results, rather than parking
// forever now that nothing will drain them. Main thread only, at engine
// shutdown.
func (r *MPSC) Close() {
	r.Flush()
	r.closed.Store(true)
	r.mu.Lock()
	r.cond.Broadcast()
	r.mu.Unlock()
}

// Open re-arms a closed ring (engine re-initialisation).
func (r *MPSC) Open() {
	r.closed.Store(false)
}
