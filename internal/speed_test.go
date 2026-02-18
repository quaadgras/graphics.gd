//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/GDScript"
	gd "graphics.gd/internal"
	"graphics.gd/internal/threadcheck"
	"graphics.gd/variant/Object"
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

func BenchmarkMethodBindCall(B *testing.B) {
	benchOnMain(B, func(B *channelB) {
		B.ReportAllocs()
		for B.Loop() {
			Engine.GetFramesPerSecond()
		}
	})
}

func BenchmarkMethodBindCallWithArgument(B *testing.B) {
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

func BenchmarkScriptCall(B *testing.B) {
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
