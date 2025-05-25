package Angle

import (
	"math"
	"testing"
)

const epsilon = 1e-5

func almostEqual(a, b float32) bool {
	return math.Abs(float64(a-b)) < epsilon
}

func TestDifference(t *testing.T) {
	cases := []struct {
		name     string
		from, to float32 // input angles in radians
		want     float32 // expected signed difference in [–π, +π]
	}{
		{"zero→zero", 0, 0, 0},
		{"0→π/2", 0, math.Pi / 2, math.Pi / 2},
		{"π/2→0", math.Pi / 2, 0, -math.Pi / 2},
		{"wrap past +π", 3 * math.Pi / 4, -3 * math.Pi / 4, math.Pi / 2},
		{"wrap past –π", -3 * math.Pi / 4, 3 * math.Pi / 4, -math.Pi / 2},
		{"exact opp 0→π", 0, math.Pi, -math.Pi},
		{"exact opp π→0", math.Pi, 0, math.Pi},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Difference(Radians(c.from), Radians(c.to))
			if !almostEqual(float32(got), c.want) {
				t.Fatalf(
					"Difference(%.3f, %.3f) = %.5f; want %.5f",
					c.from, c.to, float32(got), c.want,
				)
			}
		})
	}
}
