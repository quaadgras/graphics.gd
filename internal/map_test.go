package gd_test

import (
	"testing"

	"grow.graphics/gd"
)

func TestMap(t *testing.T) {
	lt := gd.NewLifetime(godot)
	defer lt.End()

	var ages = gd.NewMap[gd.String, gd.Int](lt)
	ages.Set(lt.String("John"), gd.Int(32))
	ages.Set(lt.String("Jane"), gd.Int(28))

	if ages.Get(lt, lt.String("John")) != gd.Int(32) {
		t.Error("expected 32")
	}
	if ages.Get(lt, lt.String("Jane")) != gd.Int(28) {
		t.Error("expected 28")
	}
	if ages.Len() != 2 {
		t.Error("expected 2")
	}
}
