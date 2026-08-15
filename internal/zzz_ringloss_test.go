package gd_test

// Regression net for lost cross-thread calls (a fire-and-forget SetName
// that silently never executed was a real CI flake): after everything else
// in the suite has run (zzz_ sorts last), a ring barrier flushes the
// cross-thread queue and the published/executed counters must agree — any
// queued call that vanished without running shows up here as a nonzero
// difference, regardless of which test queued it.

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"graphics.gd/internal/ring"
)

// Hang watchdog: if the suite is somehow still running after ten minutes,
// something has deadlocked — dump every goroutine stack to a file so the
// hang is diagnosable even where outside SIGQUIT/SIGABRT are swallowed
// (the static musl editor does this). Not on js: the browser has no
// filesystem to write to.
func init() {
	if runtime.GOOS == "js" {
		return
	}
	go func() {
		time.Sleep(10 * time.Minute)
		buf := make([]byte, 1<<24)
		n := runtime.Stack(buf, true)
		os.WriteFile(filepath.Join(os.TempDir(), "gd-test-hang-goroutines.txt"), buf[:n], 0o644)
	}()
}

func TestRingNoLostCalls(t *testing.T) {
	if runtime.GOOS == "js" {
		// Single-threaded: every goroutine dispatches directly, so the
		// cross-thread ring is never used (the counters are always 0/0).
		// The barrier also wakes this final test inside the engine's frame
		// export, making the suite exit from within it — which tears down
		// the wasm instance mid-callback and fails the web run under -v.
		t.Skip("the cross-thread ring is unused on js")
	}
	// A goroutine leaked by an earlier test may still be queueing calls, so
	// an imbalance only counts if it survives several barriers.
	var buffered, executed uint64
	for attempt := 0; attempt < 5; attempt++ {
		ring.Threads.Barrier()
		buffered, executed = ring.Threads.Counters()
		if buffered == executed {
			t.Logf("ring balance ok: %d calls buffered and executed", buffered)
			return
		}
	}
	t.Errorf("cross-thread ring lost %d fire-and-forget call(s): buffered=%d executed=%d",
		int64(buffered)-int64(executed), buffered, executed)
}
