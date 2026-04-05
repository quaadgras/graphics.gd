//go:build js

package startup

import (
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
)

//go:wasmimport gd object_unsafe_call
func wasm_gd_object_unsafe_call(obj uint32, method uint32, result uint32, shape_hi uint32, shape_lo uint32, args uint32)

func init() {
	gdextension.Host.Objects.Unsafe.Call = func(p0 gdextension.Object, p1 gdextension.MethodForClass, p2 gdextension.CallReturns[interface{}], shape gdextension.Shape, p4 gdextension.CallAccepts[interface{}]) {
		mem2 := gdmemory.MakeResult(shape)
		mem4 := gdmemory.CopyArguments(shape, p4)
		wasm_gd_object_unsafe_call(uint32(p0), uint32(p1), uint32(mem2), uint32(shape>>32), uint32(shape&0xFFFFFFFF), uint32(mem4))
		gdmemory.LoadResult(shape, p2, mem2)
	}
}
