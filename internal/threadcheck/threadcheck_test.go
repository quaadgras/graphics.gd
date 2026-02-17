package threadcheck_test

import (
	"testing"

	"graphics.gd/internal/threadcheck"
)

var isMain = threadcheck.Main()

func TestMain(t *testing.T) {
	if !isMain {
		t.Fatal("expected Main() == true during init")
	}
	done := make(chan bool)
	go func() {
		done <- threadcheck.Main()
	}()
	if result := <-done; result {
		t.Fatal("expected Main() == false on a different goroutine")
	}
}
