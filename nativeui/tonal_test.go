package nativeui

import (
	"math"
	"testing"
)

// TestTonalRoundTrip: a colour's own tone gives the colour back, and tones
// are CIE L* (white is 100, black 0).
func TestTonalRoundTrip(t *testing.T) {
	c := hex("#387AFF")
	l, _, _ := toLab(c)
	got := tonalOf(c).tone(l)
	if math.Abs(float64(got.R-c.R)) > 0.01 || math.Abs(float64(got.G-c.G)) > 0.01 || math.Abs(float64(got.B-c.B)) > 0.01 {
		t.Fatalf("round trip %v -> %v", c, got)
	}
	if w := tonalOf(c).withChroma(0).tone(100); w.R < 0.99 || w.B < 0.99 {
		t.Fatalf("tone 100 = %v", w)
	}
}
