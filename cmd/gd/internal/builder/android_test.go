package builder

import (
	"strings"
	"testing"
)

// The dying process and the one android relaunches in its place write to the
// same log buffer, so the crash report arrives with the next run's progress
// lines spliced into it. The report has to survive that: this is the shape the
// CI emulator produces, where the stack used to be cut off after its first line.
const interleavedCrashLog = `=== RUN   TestAdvancedRoundTrip
--- PASS: TestAdvancedRoundTrip (0.00s)
=== RUN   TestResourceLibrary
panic: use of an invalid reference [recovered, repanicked]

--- PASS: TestBasis (0.05s)
goroutine 53 [running]:
=== RUN   TestCallables
testing.tRunner.func1.2({0x1160a5380, 0x116336e50})
	/src/testing/testing.go:1974 +0x1a0
graphics.gd/classdb/Texture2D.class.GetWidth(...)
	/graphics.gd/classdb/Texture2D/class.go:706 +0x48
--- PASS: TestCallables (0.01s)
graphics.gd/internal_test.TestResourceLibrary(0x8a26f3226c8)
	/graphics.gd/internal/scene_library_test.go:88 +0x30
created by testing.(*T).Run in goroutine 19
	/src/testing/testing.go:2101 +0x3a8
=== RUN   TestErrors
--- PASS: TestErrors (0.00s)
GDTEST_DONE 0
`

func TestAndroidResultsKeepsInterleavedCrash(t *testing.T) {
	report, crashed := androidResults(interleavedCrashLog)
	if !crashed {
		t.Fatal("expected the panic to be reported as a crash")
	}
	// Every frame of the stack has to survive, not just the panic line.
	for _, want := range []string{
		"panic: use of an invalid reference [recovered, repanicked]",
		"goroutine 53 [running]:",
		"graphics.gd/classdb/Texture2D.class.GetWidth(...)",
		"/graphics.gd/internal/scene_library_test.go:88 +0x30",
		"created by testing.(*T).Run in goroutine 19",
	} {
		if !strings.Contains(report, want) {
			t.Errorf("crash report is missing %q, got:\n%s", want, report)
		}
	}
	// The relaunched run's progress lines are not part of the crash report; they
	// are still reported once each, as results.
	if strings.Contains(report, "=== RUN") {
		t.Errorf("=== RUN lines should not be printed, got:\n%s", report)
	}
	if n := strings.Count(report, "--- PASS: TestBasis"); n != 1 {
		t.Errorf("--- PASS: TestBasis printed %d times, want 1:\n%s", n, report)
	}
}

// A clean run must not be reported as a crash, and repeated result lines from
// the relaunches collapse to one each.
func TestAndroidResultsCleanRun(t *testing.T) {
	const log = `=== RUN   TestOne
--- PASS: TestOne (0.01s)
    one_test.go:12: some detail
=== RUN   TestOne
--- PASS: TestOne (0.01s)
GDTEST_DONE 0
`
	report, crashed := androidResults(log)
	if crashed {
		t.Error("a clean run must not be reported as crashed")
	}
	if n := strings.Count(report, "--- PASS: TestOne"); n != 1 {
		t.Errorf("--- PASS: TestOne printed %d times, want 1:\n%s", n, report)
	}
	if !strings.Contains(report, "one_test.go:12: some detail") {
		t.Errorf("failure detail lines must be kept, got:\n%s", report)
	}
}
