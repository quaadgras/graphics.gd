package gd_test

import (
	"testing"

	"graphics.gd/variant/Callable"
)

type fatalSentinel struct{}

// channelTB forwards test method calls back to the testing goroutine via a channel.
// Fatal/FailNow/SkipNow panic with fatalSentinel to unwind the main goroutine body,
// while the test goroutine receives and executes t.Fatal() (which calls runtime.Goexit
// on the correct goroutine).
type channelTB struct {
	*testing.T
	calls chan<- func(*testing.T)
}

func (c *channelTB) Error(args ...any) {
	c.calls <- func(t *testing.T) { t.Helper(); t.Error(args...) }
}
func (c *channelTB) Errorf(format string, args ...any) {
	c.calls <- func(t *testing.T) { t.Helper(); t.Errorf(format, args...) }
}
func (c *channelTB) Fatal(args ...any) {
	c.calls <- func(t *testing.T) { t.Helper(); t.Fatal(args...) }
	panic(fatalSentinel{})
}
func (c *channelTB) Fatalf(format string, args ...any) {
	c.calls <- func(t *testing.T) { t.Helper(); t.Fatalf(format, args...) }
	panic(fatalSentinel{})
}
func (c *channelTB) FailNow() {
	c.calls <- func(t *testing.T) { t.Helper(); t.FailNow() }
	panic(fatalSentinel{})
}
func (c *channelTB) Fail() {
	c.calls <- func(t *testing.T) { t.Fail() }
}
func (c *channelTB) Log(args ...any) {
	c.calls <- func(t *testing.T) { t.Helper(); t.Log(args...) }
}
func (c *channelTB) Logf(format string, args ...any) {
	c.calls <- func(t *testing.T) { t.Helper(); t.Logf(format, args...) }
}

// runOnMain dispatches fn to the main goroutine via Callable.Defer and blocks
// until it completes. Test method calls inside fn are forwarded back to the
// calling goroutine via a channel and executed on the real *testing.T.
func runOnMain(t *testing.T, fn func(testing.TB)) {
	t.Helper()
	calls := make(chan func(*testing.T))
	Callable.Defer(Callable.New(func() {
		defer close(calls)
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(fatalSentinel); !ok {
					panic(r)
				}
			}
		}()
		fn(&channelTB{T: t, calls: calls})
	}))
	for call := range calls {
		call(t)
	}
}

// channelB forwards benchmark method calls back to the benchmark goroutine via a channel.
type channelB struct {
	*testing.B
	calls chan<- func(*testing.B)
}

func (c *channelB) Error(args ...any) {
	c.calls <- func(b *testing.B) { b.Helper(); b.Error(args...) }
}
func (c *channelB) Errorf(format string, args ...any) {
	c.calls <- func(b *testing.B) { b.Helper(); b.Errorf(format, args...) }
}
func (c *channelB) Fatal(args ...any) {
	c.calls <- func(b *testing.B) { b.Helper(); b.Fatal(args...) }
	panic(fatalSentinel{})
}
func (c *channelB) Fatalf(format string, args ...any) {
	c.calls <- func(b *testing.B) { b.Helper(); b.Fatalf(format, args...) }
	panic(fatalSentinel{})
}
func (c *channelB) FailNow() {
	c.calls <- func(b *testing.B) { b.Helper(); b.FailNow() }
	panic(fatalSentinel{})
}
func (c *channelB) Fail() {
	c.calls <- func(b *testing.B) { b.Fail() }
}
func (c *channelB) Log(args ...any) {
	c.calls <- func(b *testing.B) { b.Helper(); b.Log(args...) }
}
func (c *channelB) Logf(format string, args ...any) {
	c.calls <- func(b *testing.B) { b.Helper(); b.Logf(format, args...) }
}

// benchOnMain dispatches fn to the main goroutine via Callable.Defer and blocks
// until it completes. The *testing.B is accessible through the channelB proxy,
// which forwards assertion calls back to the benchmark goroutine. The proxy
// embeds *testing.B so Loop(), N, ReportAllocs(), ResetTimer(), and Cleanup()
// are available directly.
func benchOnMain(b *testing.B, fn func(*channelB)) {
	b.Helper()
	calls := make(chan func(*testing.B))
	Callable.Defer(Callable.New(func() {
		defer close(calls)
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(fatalSentinel); !ok {
					panic(r)
				}
			}
		}()
		fn(&channelB{B: b, calls: calls})
	}))
	for call := range calls {
		call(b)
	}
}
