# FFI Command Buffer Design

## Overview

Replace direct per-call cgo crossings with a ring buffer (command buffer) that
records engine calls and flushes them in a single cgo crossing. Combined with
a fast main-thread check, this enables automatic cross-thread dispatch and
batched execution.

## Fast Main-Thread Detection

Implemented in `internal/threadcheck`. Uses the Go runtime's dedicated g
register (R14 on amd64, R28 on arm64) to compare the current goroutine pointer
against the one captured at init time. Since the main goroutine is always
`LockOSThread`'d, its g pointer is stable for the process lifetime.

- **amd64**: `MOVQ R14, ret+0(FP)` — one instruction
- **arm64**: `MOVD g, ret+0(FP)` — one instruction
- **wasm**: always returns true (single-threaded)
- **other**: falls back to `gdextension.Host.Threads.Main()` via cgo

Cost: ~1-3 cycles. No syscall, no cgo, no TLS indirection.

## Ring Buffer Architecture

### Entry Layout

Each entry is self-contained:

    object   uintptr          — Godot object pointer
    method   uintptr          — method bind pointer
    shape    uint64           — shape-encoded argument sizes
    args     [ARGS_SIZE]byte  — copied argument data
    result   [RESULT_SIZE]byte — result written here during flush
    refs     [16]uint16       — intra-buffer references (0 = literal arg,
                                 N = use entries[N-1].result)
    owner    *uintptr         — Go-side pointer to off-load result into

### Ring Structure

    head     uint32   — next write position
    tail     uint32   — next flush position
    entries  [SIZE]Entry

Main-thread ring is SPSC (single producer = locked main goroutine, single
consumer = same goroutine during flush). Cross-thread ring is MPSC (multiple
goroutines CAS head forward, main thread drains during flush).

## Pointer Integration

Pointers to Godot objects are opaque to Go — they are never dereferenced, only
passed to engine calls. This means a pointer can be in one of two states:

    bit 0 = 0  →  real Godot pointer (always aligned, bit 0 is free)
    bit 0 = 1  →  ring index (value >> 1 = entry index)

A ring-tagged pointer is a "virtual" pointer whose actual value hasn't been
materialized yet. It exists only as a future result in the ring buffer.

### Passing Ring-Tagged Pointers

When a ring-tagged pointer is passed as an argument to another buffered call,
the entry records a ref (refs[N] = ring index). During flush, the C-side loop
resolves refs by pointing directly at the referenced entry's result slot:

    void ring_flush(ring_buffer *ring, uint64_t through) {
        for (i = ring->tail; i <= through; i++) {
            ring_entry *e = &ring->entries[i & RING_MASK];
            void *points[16];
            prepare_callframe(1, &points[0], e->shape, e->args);
            for (j = 0; j < 16; j++) {
                if (e->refs[j] != 0)
                    points[j] = ring->entries[(e->refs[j]-1) & RING_MASK].result;
            }
            gdextension_object_method_bind_ptrcall(e->method, e->object, points, e->result);
            if (e->owner != NULL)
                *e->owner = *(uintptr_t *)e->result;
        }
        ring->tail = through + 1;
    }

### Reclamation (Write Head Wraps)

When the write head wraps around to a slot that has an unmaterialized result,
the flush processes that entry and writes the real pointer back into the
Go-side pointer via the owner back-pointer. The tag bit is cleared and all
future uses of that pointer are direct.

## Flush Triggers

Go code never observes raw Godot pointers directly. Reference type returns
(Object, String, Ref, etc.) are always ring-tagged. This means:

- **Reference type return** → buffer, return ring-tagged pointer. NO FLUSH.
- **Passing a ring-tagged pointer to another call** → buffer with ref. NO FLUSH.
- **Write head wraps to occupied slot** → off-load that entry. PARTIAL FLUSH.
- **Frame boundary** → flush everything. ONE cgo crossing.
- **Concrete value type return that Go code observes** → FLUSH (the only
  mid-frame flush in typical usage).

## Deferred Value Types (Future)

For the advanced API, even value type returns (int, float, Vector2, etc.) can
be deferred using pointer-like wrappers. Go code only forces a flush when it
actually inspects the value (e.g., uses it in an if statement or Go-side
arithmetic).

Engine-side arithmetic operations (e.g., Vector2.Add) can themselves be
buffered, keeping the entire computation in the ring:

    pos := node.GetPosition()                         // ring, deferred Vector2
    node.SetPosition(pos.WithY(pos.Y.Add(1.0)))       // ring, all deferred
    length := s.Length()                               // ring, deferred int
    array.Resize(length)                               // ring, refs length

Zero mid-frame flushes. The only "pipeline stall" is when Go code needs to
branch on or compute with a concrete value from the engine.

## Cross-Thread Dispatch

Using `threadcheck.Main()` (~1-3 cycles), calls from non-main goroutines are
routed to a separate MPSC ring. The main thread drains this ring during flush.

- **Off-thread void call** → push to MPSC ring (fire-and-forget)
- **Off-thread call needing result** → push to MPSC ring + block until
  entry status is DONE

### MPSC Ring Full: Backpressure

When the MPSC ring is full, off-thread goroutines must park until the main
thread drains space. Use an atomic fast path with a sync.Cond slow path:

    type MPSCRing struct {
        head    atomic.Uint32
        tail    atomic.Uint32
        cond    sync.Cond
        entries [SIZE]Entry
    }

    func (r *MPSCRing) Claim() *Entry {
        for {
            head := r.head.Load()
            tail := r.tail.Load()
            if head-tail >= SIZE {
                // slow path: ring full, park until main thread drains
                r.cond.L.Lock()
                for r.head.Load()-r.tail.Load() >= SIZE {
                    r.cond.Wait()
                }
                r.cond.L.Unlock()
                continue
            }
            if r.head.CompareAndSwap(head, head+1) {
                return &r.entries[head&MASK]
            }
        }
    }

The main thread, after draining entries during flush:

    func (r *MPSCRing) Flush() {
        // ... process entries in C ...
        r.tail.Store(newTail)
        r.cond.Broadcast()
    }

The fast path (ring not full, CAS succeeds) touches no mutex — just two
atomic loads and a CAS. Only when the ring is actually full does a goroutine
park on the sync.Cond.

Note: the SPSC main-thread ring never has this problem. The main goroutine is
both writer and flusher, so when the ring fills up it just flushes inline.

### Fundamental Constraint

If the main thread is inside the engine (running Godot's frame tick), it
cannot drain the MPSC ring. Off-thread goroutines are blocked until the next
Go callback or frame boundary where the main thread flushes. This is inherent
— Godot's API isn't thread-safe, so off-thread calls must wait for the main
thread regardless. The ring makes the waiting explicit and batched rather than
per-call.

## Example: Typical Frame

    Go callback:
      s := String.New("hello")        // entry 0, ring-tagged pointer
      node.SetName(s)                 // entry 1, refs[0] = entry 0
      node.SetVisible(true)           // entry 2, plain void
      pos := node.GetPosition()       // entry 3, ring-tagged (or deferred)
      node.SetPosition(pos)           // entry 4, refs entry 3
      // return to engine

    Flush (one cgo crossing):
      process entries 0-4 in C loop
      off-load any remaining owners
      advance tail to 5

    Result: 1 cgo crossing instead of 5.
