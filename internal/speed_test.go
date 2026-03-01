//go:build !generate

package gd_test

import (
	"sync"
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Input"
	gd "graphics.gd/internal"
	"graphics.gd/internal/threadcheck"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

func BenchmarkBuiltinPointerCall(B *testing.B) {
	benchOnMain(B, func(B *channelB) {
		B.ReportAllocs()
		s := gd.NewString("Hello, World!")
		var sum int64
		for B.Loop() {
			sum += s.Length()
		}
		if sum != int64(B.N)*int64(len("Hello, World!")) {
			B.Fail()
		}
	})
}

func BenchmarkMethodBindCallWithReturnValue(B *testing.B) {
	benchOnMain(B, func(B *channelB) {
		B.ReportAllocs()
		for B.Loop() {
			Engine.GetFramesPerSecond()
		}
	})
}

func BenchmarkMethodBindCallThatReturnsVoid(B *testing.B) {
	benchOnMain(B, func(B *channelB) {
		if !threadcheck.Main() {
			B.Fatal("not main!")
		}
		B.ReportAllocs()
		for B.Loop() {
			Engine.SetMaxFps(60)
		}
	})
}

func BenchmarkGDScriptCall(B *testing.B) {
	benchOnMain(B, func(B *channelB) {
		B.ReportAllocs()
		var script = GDScript.New().AsScript()
		script.SetSourceCode(`extends Object
var n: int
func bench():
	var sum = 0
	for i in range(n):
		Engine.set_max_fps(60)
	return sum
`)
		script.Reload()
		obj := Object.New()
		obj.SetScript(script)
		gd.ObjectSet(obj[0], gd.NewStringName("n"), gd.NewVariant(B.N))
		bench := gd.NewStringName("bench")

		B.ResetTimer()
		gd.ObjectCall(obj[0], bench)
		gd.ObjectFree(obj[0])
	})
}

// BenchVirtualControl is used by BenchmarkVirtualCall to measure the
// Godot→Go checked_call path. Calling GetMinimumSize() on this from Go
// triggers: Go→Godot ptrcall → C++ GDVIRTUAL_CALL(_get_minimum_size) →
// cgo_class_call_virtual_with_data_func → go_on_extension_instance_checked_call → Go.
type BenchVirtualControl struct {
	Control.Extension[BenchVirtualControl] `gd:"BenchVirtualControl"`
}

func (b *BenchVirtualControl) GetMinimumSize() Vector2.XY { return Vector2.XY{} }

var registerBenchVirtualOnce sync.Once

func BenchmarkVirtualCallback(B *testing.B) {
	benchOnMain(B, func(B *channelB) {
		B.ReportAllocs()
		registerBenchVirtualOnce.Do(func() {
			classdb.Register[BenchVirtualControl]()
		})
		ctrl := new(BenchVirtualControl)
		instance := ctrl.AsControl()
		for B.Loop() {
			instance.GetMinimumSize()
		}
	})
}

func BenchmarkCallable(B *testing.B) {
	benchOnMain(B, func(B *channelB) {
		B.ReportAllocs()
		var script = GDScript.New().AsScript()
		script.SetSourceCode(`extends Object
var n: int
func bench(c):
	var sum = 0
	for i in range(n):
		sum += c.call()
	return sum
`)
		script.Reload()
		obj := Object.New()
		obj.SetScript(script)
		gd.ObjectSet(obj[0], gd.NewStringName("n"), gd.NewVariant(B.N))
		bench := gd.NewStringName("bench")
		var array []gd.Variant
		array = append(array, gd.NewVariant(gd.NewCallable(func() int {
			return 1
		})))
		var result gd.Variant
		B.Cleanup(func() {
			if result.Interface().(int64) != int64(B.N) {
				B.Fail()
			}
		})
		B.ResetTimer()
		result, _ = gd.ObjectCall(obj[0], bench, array...)
		gd.ObjectFree(obj[0])
	})
}

func BenchmarkMethodBindCallWithString(t *testing.B) {
	benchOnMain(t, func(B *channelB) {
		B.ReportAllocs()
		for B.Loop() {
			Input.IsActionPressed("ui_left", false)
		}
	})
}
