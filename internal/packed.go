package gd

import (
	"strings"
	"unsafe"

	"grow.graphics/gd/internal/callframe"
	"runtime.link/mmm"
)

type PackedByteArray mmm.Pointer[API, PackedByteArray, [2]uintptr]

func (p PackedByteArray) Index(idx Int) byte {
	return mmm.API(p).PackedByteArray.Index(p, idx)
}

func (p PackedByteArray) SetIndex(idx Int, value byte) {
	mmm.API(p).PackedByteArray.SetIndex(p, idx, value)
}

func (p PackedByteArray) UnsafePointer() unsafe.Pointer {
	return mmm.API(p).PackedByteArray.UnsafePointer(p)
}

func (p PackedByteArray) Len() int { return int(p.Size()) }
func (p PackedByteArray) Cap() int { return int(p.Size()) }

func (p PackedByteArray) Free() {
	var frame = callframe.New()
	mmm.API(p).typeset.destruct.PackedByteArray(callframe.Arg(frame, mmm.End(p)).Uintptr())
	frame.Free()
}

type PackedInt32Array mmm.Pointer[API, PackedInt32Array, [2]uintptr]

func (p PackedInt32Array) Index(idx Int) int32 {
	return mmm.API(p).PackedInt32Array.Index(p, idx)
}

func (p PackedInt32Array) SetIndex(idx Int, value int32) {
	mmm.API(p).PackedInt32Array.SetIndex(p, idx, value)
}

func (p PackedInt32Array) Free() {
	var frame = callframe.New()
	mmm.API(p).typeset.destruct.PackedInt32Array(callframe.Arg(frame, mmm.End(p)).Uintptr())
	frame.Free()
}

func (p PackedInt32Array) UnsafePointer() unsafe.Pointer {
	return mmm.API(p).PackedInt32Array.UnsafePointer(p)
}

func (p PackedInt32Array) Len() int { return int(p.Size()) }
func (p PackedInt32Array) Cap() int { return int(p.Size()) }

type PackedInt64Array mmm.Pointer[API, PackedInt64Array, [2]uintptr]

func (p PackedInt64Array) Index(idx Int) int64 {
	return mmm.API(p).PackedInt64Array.Index(p, idx)
}

func (p PackedInt64Array) SetIndex(idx Int, value int64) {
	mmm.API(p).PackedInt64Array.SetIndex(p, idx, value)
}

func (p PackedInt64Array) Free() {
	var frame = callframe.New()
	mmm.API(p).typeset.destruct.PackedInt64Array(callframe.Arg(frame, mmm.End(p)).Uintptr())
	frame.Free()
}

func (p PackedInt64Array) UnsafePointer() unsafe.Pointer {
	return mmm.API(p).PackedInt64Array.UnsafePointer(p)
}

func (p PackedInt64Array) Len() int { return int(p.Size()) }
func (p PackedInt64Array) Cap() int { return int(p.Size()) }

type PackedFloat32Array mmm.Pointer[API, PackedFloat32Array, [2]uintptr]

func (p PackedFloat32Array) Index(idx Int) float32 {
	return mmm.API(p).PackedFloat32Array.Index(p, idx)
}

func (p PackedFloat32Array) SetIndex(idx Int, value float32) {
	mmm.API(p).PackedFloat32Array.SetIndex(p, idx, value)
}

func (p PackedFloat32Array) Free() {
	var frame = callframe.New()
	mmm.API(p).typeset.destruct.PackedFloat32Array(callframe.Arg(frame, mmm.End(p)).Uintptr())
	frame.Free()
}

func (p PackedFloat32Array) Len() int { return int(p.Size()) }
func (p PackedFloat32Array) Cap() int { return int(p.Size()) }

func (p PackedFloat32Array) UnsafePointer() unsafe.Pointer {
	return mmm.API(p).PackedFloat32Array.UnsafePointer(p)
}

type PackedFloat64Array mmm.Pointer[API, PackedFloat64Array, [2]uintptr]

func (p PackedFloat64Array) Index(idx Int) float64 {
	return mmm.API(p).PackedFloat64Array.Index(p, idx)
}

func (p PackedFloat64Array) SetIndex(idx Int, value float64) {
	mmm.API(p).PackedFloat64Array.SetIndex(p, idx, value)
}

func (p PackedFloat64Array) Free() {
	var frame = callframe.New()
	mmm.API(p).typeset.destruct.PackedFloat64Array(callframe.Arg(frame, mmm.End(p)).Uintptr())
	frame.Free()
}

func (p PackedFloat64Array) UnsafePointer() unsafe.Pointer {
	return mmm.API(p).PackedFloat64Array.UnsafePointer(p)
}

func (p PackedFloat64Array) Len() int { return int(p.Size()) }
func (p PackedFloat64Array) Cap() int { return int(p.Size()) }

type PackedStringArray mmm.Pointer[API, PackedStringArray, [2]uintptr]

func (p PackedStringArray) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	size := int(p.Size())
	ctx := NewContext(mmm.API(p))
	defer ctx.End()
	for i := 0; i < size; i++ {
		builder.WriteString(p.Index(ctx, Int(i)).String())
		if i < size-1 {
			builder.WriteString(" ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (p PackedStringArray) Index(ctx Lifetime, idx Int) String {
	return mmm.API(p).PackedStringArray.Index(ctx, p, idx)
}

func (p PackedStringArray) SetIndex(idx Int, value String) {
	mmm.API(p).PackedStringArray.SetIndex(p, idx, value)
}

func (p PackedStringArray) AsSlice(ctx Lifetime) []String {
	return mmm.API(p).PackedStringArray.CopyAsSlice(ctx, p)
}

func (p PackedStringArray) Free() {
	var frame = callframe.New()
	mmm.API(p).typeset.destruct.PackedStringArray(callframe.Arg(frame, mmm.End(p)).Uintptr())
	frame.Free()
}

type PackedVector2Array mmm.Pointer[API, PackedVector2Array, [2]uintptr]

func (p PackedVector2Array) Index(idx Int) Vector2 {
	return mmm.API(p).PackedVector2Array.Index(p, idx)
}

func (p PackedVector2Array) SetIndex(idx Int, value Vector2) {
	mmm.API(p).PackedVector2Array.SetIndex(p, idx, value)
}

func (p PackedVector2Array) Free() {
	var frame = callframe.New()
	mmm.API(p).typeset.destruct.PackedVector2Array(callframe.Arg(frame, mmm.End(p)).Uintptr())
	frame.Free()
}

func (p PackedVector2Array) Len() int { return int(p.Size()) }
func (p PackedVector2Array) Cap() int { return int(p.Size()) }
func (p PackedVector2Array) UnsafePointer() unsafe.Pointer {
	return mmm.API(p).PackedVector2Array.UnsafePointer(p)
}

type PackedVector3Array mmm.Pointer[API, PackedVector3Array, [2]uintptr]

func (p PackedVector3Array) Index(idx Int) Vector3 {
	return mmm.API(p).PackedVector3Array.Index(p, idx)
}

func (p PackedVector3Array) SetIndex(idx Int, value Vector3) {
	mmm.API(p).PackedVector3Array.SetIndex(p, idx, value)
}

func (p PackedVector3Array) Free() {
	var frame = callframe.New()
	mmm.API(p).typeset.destruct.PackedVector3Array(callframe.Arg(frame, mmm.End(p)).Uintptr())
	frame.Free()
}

func (p PackedVector3Array) Len() int { return int(p.Size()) }
func (p PackedVector3Array) Cap() int { return int(p.Size()) }
func (p PackedVector3Array) UnsafePointer() unsafe.Pointer {
	return mmm.API(p).PackedVector3Array.UnsafePointer(p)
}

type PackedColorArray mmm.Pointer[API, PackedColorArray, [2]uintptr]

func (p PackedColorArray) Index(idx Int) Color {
	return mmm.API(p).PackedColorArray.Index(p, idx)
}

func (p PackedColorArray) SetIndex(idx Int, value Color) {
	mmm.API(p).PackedColorArray.SetIndex(p, idx, value)
}

func (p PackedColorArray) Free() {
	var frame = callframe.New()
	mmm.API(p).typeset.destruct.PackedColorArray(callframe.Arg(frame, mmm.End(p)).Uintptr())
	frame.Free()
}

func (p PackedColorArray) Len() int { return int(p.Size()) }
func (p PackedColorArray) Cap() int { return int(p.Size()) }
func (p PackedColorArray) UnsafePointer() unsafe.Pointer {
	return mmm.API(p).PackedColorArray.UnsafePointer(p)
}

func (godot Lifetime) PackedByteArray() PackedByteArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	godot.API.typeset.creation.PackedByteArray[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[PackedByteArray](godot.Lifetime, godot.API, raw)
}

func (godot Lifetime) PackedColorArray() PackedColorArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	godot.API.typeset.creation.PackedColorArray[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[PackedColorArray](godot.Lifetime, godot.API, raw)
}

func (godot Lifetime) PackedFloat32Array() PackedFloat32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	godot.API.typeset.creation.PackedFloat32Array[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[PackedFloat32Array](godot.Lifetime, godot.API, raw)
}

func (godot Lifetime) PackedInt32Array() PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	godot.API.typeset.creation.PackedInt32Array[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[PackedInt32Array](godot.Lifetime, godot.API, raw)
}

func (godot Lifetime) PackedStringArray() PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	godot.API.typeset.creation.PackedStringArray[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[PackedStringArray](godot.Lifetime, godot.API, raw)
}

func (godot Lifetime) PackedVector2Array() PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	godot.API.typeset.creation.PackedVector2Array[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[PackedVector2Array](godot.Lifetime, godot.API, raw)
}

func (godot Lifetime) PackedVector3Array() PackedVector3Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	godot.API.typeset.creation.PackedVector3Array[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[PackedVector3Array](godot.Lifetime, godot.API, raw)
}

func (godot Lifetime) PackedInt64Array() PackedInt64Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	godot.API.typeset.creation.PackedInt64Array[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[PackedInt64Array](godot.Lifetime, godot.API, raw)
}

func (godot Lifetime) PackedFloat64Array() PackedFloat64Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	godot.API.typeset.creation.PackedFloat64Array[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[PackedFloat64Array](godot.Lifetime, godot.API, raw)
}
