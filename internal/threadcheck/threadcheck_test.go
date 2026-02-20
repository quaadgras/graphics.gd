package threadcheck_test

import (
	"runtime"
	"testing"

	"graphics.gd/internal/threadcheck"
)

var isMain = threadcheck.Main()

func TestMain(t *testing.T) {
	if !isMain {
		t.Fatal("expected Main() == true during init")
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	done := make(chan bool)
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		done <- threadcheck.Main()
	}()
	if result := <-done; result {
		t.Fatal("expected Main() == false on a different OS thread")
	}
}
