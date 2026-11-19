//go:build js

package gd_test

import (
	"testing"

	"graphics.gd/classdb/JavaScriptBridge"
)

// TestJavaScriptBridgeEval: the web build's JavaScriptBridge singleton
// evaluates JavaScript and returns its result.
func TestJavaScriptBridgeEval(t *testing.T) {
	if got := JavaScriptBridge.Eval("1 + 1", true); got != int64(2) && got != 2.0 && got != int(2) {
		t.Fatalf("Eval(1 + 1) = %#v (%T)", got, got)
	}
	if got, ok := JavaScriptBridge.Eval("navigator.userAgent", true).(string); !ok || got == "" {
		t.Fatalf("Eval(navigator.userAgent) = %#v", got)
	}
}
