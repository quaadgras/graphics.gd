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

// ulpDiff returns the ULP distance between two float32 values.
func ulpDiff(a, b float32) uint32 {
	ai := math.Float32bits(a)
	bi := math.Float32bits(b)
	if ai > bi {
		return ai - bi
	}
	return bi - ai
}

func TestSin(t *testing.T) {
	cases := []struct {
		name string
		x    float32
	}{
		{"zero", 0},
		{"-zero", float32(math.Copysign(0, -1))},
		{"π/6", math.Pi / 6},
		{"π/4", math.Pi / 4},
		{"π/3", math.Pi / 3},
		{"π/2", math.Pi / 2},
		{"π", math.Pi},
		{"3π/2", 3 * math.Pi / 2},
		{"2π", 2 * math.Pi},
		{"-π/6", -math.Pi / 6},
		{"-π/2", -math.Pi / 2},
		{"-π", -math.Pi},
		{"small", 1e-7},
		{"-small", -1e-7},
		{"1000", 1000},
		{"-1000", -1000},
		{"0.1", 0.1},
		{"1.0", 1.0},
		{"2.0", 2.0},
		{"3.0", 3.0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Sin(Radians(c.x))
			want := float32(math.Sin(float64(c.x)))
			if d := ulpDiff(float32(got), want); d > 1 {
				t.Fatalf("Sin(%v) = %.10g; want %.10g (ULP diff %d)", c.x, got, want, d)
			}
		})
	}

	// Special values
	t.Run("NaN", func(t *testing.T) {
		got := Sin(Radians(float32(math.NaN())))
		if !math.IsNaN(float64(got)) {
			t.Fatalf("Sin(NaN) = %v; want NaN", got)
		}
	})
	t.Run("+Inf", func(t *testing.T) {
		got := Sin(Radians(float32(math.Inf(1))))
		if !math.IsNaN(float64(got)) {
			t.Fatalf("Sin(+Inf) = %v; want NaN", got)
		}
	})
	t.Run("-Inf", func(t *testing.T) {
		got := Sin(Radians(float32(math.Inf(-1))))
		if !math.IsNaN(float64(got)) {
			t.Fatalf("Sin(-Inf) = %v; want NaN", got)
		}
	})
}

func TestCos(t *testing.T) {
	cases := []struct {
		name string
		x    float32
	}{
		{"zero", 0},
		{"-zero", float32(math.Copysign(0, -1))},
		{"π/6", math.Pi / 6},
		{"π/4", math.Pi / 4},
		{"π/3", math.Pi / 3},
		{"π/2", math.Pi / 2},
		{"π", math.Pi},
		{"3π/2", 3 * math.Pi / 2},
		{"2π", 2 * math.Pi},
		{"-π/6", -math.Pi / 6},
		{"-π/2", -math.Pi / 2},
		{"-π", -math.Pi},
		{"small", 1e-7},
		{"-small", -1e-7},
		{"1000", 1000},
		{"-1000", -1000},
		{"0.1", 0.1},
		{"1.0", 1.0},
		{"2.0", 2.0},
		{"3.0", 3.0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Cos(Radians(c.x))
			want := float32(math.Cos(float64(c.x)))
			if d := ulpDiff(float32(got), want); d > 1 {
				t.Fatalf("Cos(%v) = %.10g; want %.10g (ULP diff %d)", c.x, got, want, d)
			}
		})
	}

	// Special values
	t.Run("NaN", func(t *testing.T) {
		got := Cos(Radians(float32(math.NaN())))
		if !math.IsNaN(float64(got)) {
			t.Fatalf("Cos(NaN) = %v; want NaN", got)
		}
	})
	t.Run("+Inf", func(t *testing.T) {
		got := Cos(Radians(float32(math.Inf(1))))
		if !math.IsNaN(float64(got)) {
			t.Fatalf("Cos(+Inf) = %v; want NaN", got)
		}
	})
	t.Run("-Inf", func(t *testing.T) {
		got := Cos(Radians(float32(math.Inf(-1))))
		if !math.IsNaN(float64(got)) {
			t.Fatalf("Cos(-Inf) = %v; want NaN", got)
		}
	})
}

func TestSinCos(t *testing.T) {
	angles := []float32{
		0, math.Pi / 6, math.Pi / 4, math.Pi / 3, math.Pi / 2,
		math.Pi, 3 * math.Pi / 2, 2 * math.Pi,
		-math.Pi / 4, -math.Pi / 2, -math.Pi,
		0.1, 1.0, 2.0, 3.0, 1000,
	}

	for _, x := range angles {
		s := Sin(Radians(x))
		c := Cos(Radians(x))
		v := Radians(x).AsVector2()
		if v.X != c || v.Y != s {
			t.Fatalf("AsVector2(%v) = {%v, %v}; want {%v, %v}", x, v.X, v.Y, c, s)
		}
	}
}

func BenchmarkSin(b *testing.B) {
	x := Radians(1.0)
	for b.Loop() {
		_ = Sin(x)
	}
}

func BenchmarkCos(b *testing.B) {
	x := Radians(1.0)
	for b.Loop() {
		_ = Cos(x)
	}
}

func BenchmarkSinCos(b *testing.B) {
	x := Radians(1.0)
	for b.Loop() {
		x.AsVector2()
	}
}

func BenchmarkSinStdlib(b *testing.B) {
	x := 1.0
	for b.Loop() {
		_ = math.Sin(x)
	}
}

func BenchmarkCosStdlib(b *testing.B) {
	x := 1.0
	for b.Loop() {
		_ = math.Cos(x)
	}
}
