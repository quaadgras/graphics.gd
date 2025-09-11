package Int

import (
	"testing"
)

func TestSnapped_PositiveCloseLower(t *testing.T) {
	if got := Snapped(7, 5); got != 5 {
		t.Errorf("Snapped(7, 5) = %d, expected %d", got, 5)
	}
}

func TestSnapped_PositiveCloseUpper(t *testing.T) {
	if got := Snapped(8, 5); got != 10 {
		t.Errorf("Snapped(8, 5) = %d, expected %d", got, 10)
	}
}

func TestSnapped_PositiveExactlyHalfway(t *testing.T) {
	if got := Snapped(15, 10); got != 20 {
		t.Errorf("Snapped(15, 10) = %d, expected %d", got, 20)
	}
}

func TestSnapped_NegativeCloserMoreNegative(t *testing.T) {
	if got := Snapped(-8, 5); got != -10 {
		t.Errorf("Snapped(-8, 5) = %d, expected %d", got, -10)
	}
}

func TestSnapped_NegativeCloserLessNegative(t *testing.T) {
	if got := Snapped(-3, 5); got != -5 {
		t.Errorf("Snapped(-3, 5) = %d, expected %d", got, -5)
	}
}

func TestSnapped_ZeroValue(t *testing.T) {
	if got := Snapped(0, 5); got != 0 {
		t.Errorf("Snapped(0, 5) = %d, expected %d", got, 0)
	}
}

func TestSnapped_AlreadyMultiple(t *testing.T) {
	if got := Snapped(10, 5); got != 10 {
		t.Errorf("Snapped(10, 5) = %d, expected %d", got, 10)
	}
}

func TestSnapped_StepSizeOne(t *testing.T) {
	if got := Snapped(7, 1); got != 7 {
		t.Errorf("Snapped(7, 1) = %d, expected %d", got, 7)
	}
}

func TestSnapped_SmallStepSize(t *testing.T) {
	if got := Snapped(8, 3); got != 9 {
		t.Errorf("Snapped(8, 3) = %d, expected %d", got, 9)
	}
}

func TestSnapped_LargeValues(t *testing.T) {
	if got := Snapped(127, 25); got != 125 {
		t.Errorf("Snapped(127, 25) = %d, expected %d", got, 125)
	}
}
