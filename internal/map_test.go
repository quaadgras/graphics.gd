package gd_test

import (
	"testing"

	"grow.graphics/gd"
	"grow.graphics/gd/gdclass/Resource"
)

func TestMap(t *testing.T) {
	lt := gd.NewLifetime(godot)

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

	var resources = gd.NewMap[gd.String, Resource.Expert](lt)
	tmp := gd.NewLifetime(lt)
	res := *gd.New[Resource.Expert](tmp)
	res.SetName(tmp.String("res1"))
	resources.Set(lt.String("res1"), res)
	tmp.End()

	if !resources.Has(lt.String("res1")) {
		t.Error("expected res1")
	}
	res = resources.Get(lt, lt.String("res1"))
	if res.GetName(lt).String() != "res1" {
		t.Error("expected res1")
	}
	lt.End()
}
