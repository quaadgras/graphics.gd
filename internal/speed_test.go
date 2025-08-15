//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/classdb/GDScript"
	gd "graphics.gd/internal"
	"graphics.gd/variant/Object"
)

func BenchmarkMethodBindPointerCall(B *testing.B) {
	B.ReportAllocs()
	s := gd.NewString("Hello, World!")
	var sum int64
	for i := 0; i < B.N; i++ {
		sum += s.Length()
	}
	if sum != int64(B.N)*int64(len("Hello, World!")) {
		B.Fail()
	}
}

func BenchmarkScriptCall(B *testing.B) {
	B.ReportAllocs()
	var script = GDScript.New().AsScript()
	script.SetSourceCode(`extends Object
var n: int
func bench():
	var sum = 0
	for i in range(n):
		sum += "Hello, World!".length()
	return sum
`)
	script.Reload()
	obj := Object.New()
	obj.SetScript(script)
	obj[0].Set(gd.NewStringName("n"), gd.NewVariant(B.N))
	bench := gd.NewStringName("bench")
	array := gd.NewArray()
	var result gd.Variant
	B.Cleanup(func() {
		if result.Interface().(int64) != int64(B.N*len("Hello, World!")) {
			B.Fail()
		}
	})
	B.ResetTimer()
	result = obj[0].Callv(bench, array)
	obj[0].Free()
}

func BenchmarkCallable(B *testing.B) {
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
	obj[0].Set(gd.NewStringName("n"), gd.NewVariant(B.N))
	bench := gd.NewStringName("bench")
	array := gd.NewArray()
	array.PushFront(gd.NewVariant(gd.NewCallable(func() int {
		return 1
	})))
	var result gd.Variant
	B.Cleanup(func() {
		if result.Interface().(int64) != int64(B.N) {
			B.Fail()
		}
	})
	B.ResetTimer()
	result = obj[0].Callv(bench, array)
	obj[0].Free()
}
