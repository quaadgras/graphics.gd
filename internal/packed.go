package gd

import (
	"strings"
	"unsafe"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
)

func (p *PackedFloat32Array) Pointer() *PackedFloat32Array { return p }
func (p *PackedFloat64Array) Pointer() *PackedFloat64Array { return p }
func (p *PackedInt32Array) Pointer() *PackedInt32Array     { return p }
func (p *PackedInt64Array) Pointer() *PackedInt64Array     { return p }
func (p *PackedColorArray) Pointer() *PackedColorArray     { return p }
func (p *PackedStringArray) Pointer() *PackedStringArray   { return p }
func (p *PackedVector2Array) Pointer() *PackedVector2Array { return p }
func (p *PackedVector3Array) Pointer() *PackedVector3Array { return p }
func (p *PackedVector4Array) Pointer() *PackedVector4Array { return p }

func (p PackedInt32Array) AsSlice() []int32     { return Global.PackedInt32Array.CopyAsSlice(p) }
func (p PackedInt64Array) AsSlice() []int64     { return Global.PackedInt64Array.CopyAsSlice(p) }
func (p PackedFloat32Array) AsSlice() []float32 { return Global.PackedFloat32Array.CopyAsSlice(p) }
func (p PackedFloat64Array) AsSlice() []float64 { return Global.PackedFloat64Array.CopyAsSlice(p) }
func (p PackedVector2Array) AsSlice() []Vector2 { return Global.PackedVector2Array.CopyAsSlice(p) }
func (p PackedVector3Array) AsSlice() []Vector3 { return Global.PackedVector3Array.CopyAsSlice(p) }
func (p PackedVector4Array) AsSlice() []Vector4 { return Global.PackedVector4Array.CopyAsSlice(p) }
func (p PackedColorArray) AsSlice() []Color     { return Global.PackedColorArray.CopyAsSlice(p) }
func (p PackedStringArray) Strings() []string {
	var s = make([]string, p.Size())
	for i := Int(0); i < p.Size(); i++ {
		s[i] = p.Index(i).String()
	}
	return s
}

func (p PackedByteArray) Index(idx Int) Int {
	return Int(Global.PackedByteArray.Index(p, idx))
}
func (p PackedByteArray) ToByteArray() PackedByteArray { return p.Duplicate() }
func (p PackedByteArray) SetIndex(idx Int, value Int) {
	Global.PackedByteArray.SetIndex(p, idx, byte(value))
}

func (p PackedByteArray) UnsafePointer() unsafe.Pointer {
	return Global.PackedByteArray.UnsafePointer(p)
}

// Bytes returns a copy of the byte array as a byte slice.
func (p PackedByteArray) Bytes() []byte {
	var bytes = make([]byte, p.Size())
	copy(bytes, unsafe.Slice((*byte)(p.UnsafePointer()), len(bytes)))
	return bytes
}

func (p PackedByteArray) Len() int { return int(p.Size()) }
func (p PackedByteArray) Cap() int { return int(p.Size()) }

func (p PackedByteArray) Free() {
	if ptr, ok := pointers.End(p); ok {
		var frame = callframe.New()
		Global.typeset.destruct.PackedByteArray(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (p PackedInt32Array) Index(idx Int) Int {
	return Int(Global.PackedInt32Array.Index(p, idx))
}

func (p PackedInt32Array) SetIndex(idx Int, value Int) {
	Global.PackedInt32Array.SetIndex(p, idx, int32(value))
}

func (p PackedInt32Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		var frame = callframe.New()
		Global.typeset.destruct.PackedInt32Array(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (p PackedInt32Array) UnsafePointer() unsafe.Pointer {
	return Global.PackedInt32Array.UnsafePointer(p)
}

func (p PackedInt32Array) Len() int { return int(p.Size()) }
func (p PackedInt32Array) Cap() int { return int(p.Size()) }

func (p PackedInt64Array) Index(idx Int) Int {
	return Int(Global.PackedInt64Array.Index(p, idx))
}

func (p PackedInt64Array) SetIndex(idx Int, value Int) {
	Global.PackedInt64Array.SetIndex(p, idx, int64(value))
}

func (p PackedInt64Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		var frame = callframe.New()
		Global.typeset.destruct.PackedInt64Array(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (p PackedInt64Array) UnsafePointer() unsafe.Pointer {
	return Global.PackedInt64Array.UnsafePointer(p)
}

func (p PackedInt64Array) Len() int { return int(p.Size()) }
func (p PackedInt64Array) Cap() int { return int(p.Size()) }

func (p PackedFloat32Array) Index(idx Int) Float {
	return Float(Global.PackedFloat32Array.Index(p, idx))
}

func (p PackedFloat32Array) SetIndex(idx Int, value Float) {
	Global.PackedFloat32Array.SetIndex(p, idx, float32(value))
}

func (p PackedFloat32Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		var frame = callframe.New()
		Global.typeset.destruct.PackedFloat32Array(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (p PackedFloat32Array) Len() int { return int(p.Size()) }
func (p PackedFloat32Array) Cap() int { return int(p.Size()) }

func (p PackedFloat32Array) UnsafePointer() unsafe.Pointer {
	return Global.PackedFloat32Array.UnsafePointer(p)
}

func (p PackedFloat64Array) Index(idx Int) Float {
	return Float(Global.PackedFloat64Array.Index(p, idx))
}

func (p PackedFloat64Array) SetIndex(idx Int, value Float) {
	Global.PackedFloat64Array.SetIndex(p, idx, float64(value))
}

func (p PackedFloat64Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		var frame = callframe.New()
		Global.typeset.destruct.PackedFloat64Array(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (p PackedFloat64Array) UnsafePointer() unsafe.Pointer {
	return Global.PackedFloat64Array.UnsafePointer(p)
}

func (p PackedFloat64Array) Len() int { return int(p.Size()) }
func (p PackedFloat64Array) Cap() int { return int(p.Size()) }

func (p PackedStringArray) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	size := int(p.Size())
	for i := 0; i < size; i++ {
		builder.WriteString(p.Index(Int(i)).String())
		if i < size-1 {
			builder.WriteString(" ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (p PackedStringArray) Index(idx Int) String {
	return Global.PackedStringArray.Index(p, idx)
}

func (p PackedStringArray) SetIndex(idx Int, value String) {
	Global.PackedStringArray.SetIndex(p, idx, value)
}

func (p PackedStringArray) AsSlice() []String {
	return Global.PackedStringArray.CopyAsSlice(p)
}

func (p PackedStringArray) Free() {
	if ptr, ok := pointers.End(p); ok {
		var frame = callframe.New()
		Global.typeset.destruct.PackedStringArray(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (p PackedVector2Array) Index(idx Int) Vector2 {
	return Global.PackedVector2Array.Index(p, idx)
}

func (p PackedVector2Array) SetIndex(idx Int, value Vector2) {
	Global.PackedVector2Array.SetIndex(p, idx, value)
}

func (p PackedVector2Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		var frame = callframe.New()
		Global.typeset.destruct.PackedVector2Array(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (p PackedVector2Array) Len() int { return int(p.Size()) }
func (p PackedVector2Array) Cap() int { return int(p.Size()) }
func (p PackedVector2Array) UnsafePointer() unsafe.Pointer {
	return Global.PackedVector2Array.UnsafePointer(p)
}

func (p PackedVector3Array) Index(idx Int) Vector3 {
	return Global.PackedVector3Array.Index(p, idx)
}

func (p PackedVector3Array) SetIndex(idx Int, value Vector3) {
	Global.PackedVector3Array.SetIndex(p, idx, value)
}

func (p PackedVector3Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		var frame = callframe.New()
		Global.typeset.destruct.PackedVector3Array(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (p PackedVector3Array) Len() int { return int(p.Size()) }
func (p PackedVector3Array) Cap() int { return int(p.Size()) }
func (p PackedVector3Array) UnsafePointer() unsafe.Pointer {
	return Global.PackedVector3Array.UnsafePointer(p)
}

func (p PackedVector4Array) Index(idx Int) Vector4 {
	return Global.PackedVector4Array.Index(p, idx)
}

func (p PackedVector4Array) SetIndex(idx Int, value Vector4) {
	Global.PackedVector4Array.SetIndex(p, idx, value)
}

func (p PackedVector4Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		var frame = callframe.New()
		Global.typeset.destruct.PackedVector4Array(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (p PackedVector4Array) Len() int { return int(p.Size()) }
func (p PackedVector4Array) Cap() int { return int(p.Size()) }
func (p PackedVector4Array) UnsafePointer() unsafe.Pointer {
	return Global.PackedVector4Array.UnsafePointer(p)
}

func (p PackedColorArray) Index(idx Int) Color {
	return Global.PackedColorArray.Index(p, idx)
}

func (p PackedColorArray) SetIndex(idx Int, value Color) {
	Global.PackedColorArray.SetIndex(p, idx, value)
}

func (p PackedColorArray) Free() {
	if ptr, ok := pointers.End(p); ok {
		var frame = callframe.New()
		Global.typeset.destruct.PackedColorArray(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (p PackedColorArray) Len() int { return int(p.Size()) }
func (p PackedColorArray) Cap() int { return int(p.Size()) }
func (p PackedColorArray) UnsafePointer() unsafe.Pointer {
	return Global.PackedColorArray.UnsafePointer(p)
}

func NewPackedByteArray() PackedByteArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.PackedByteArray[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[PackedByteArray](raw)
}

// PackedByteSlice returns a [PackedByteArray] from a byte slice.
func NewPackedByteSlice(data []byte) PackedByteArray {
	var array = NewPackedByteArray()
	array.Resize(Int(len(data)))
	copy(unsafe.Slice((*byte)(array.UnsafePointer()), len(data)), data)
	return array
}

func NewPackedColorArray() PackedColorArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.PackedColorArray[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[PackedColorArray](raw)
}

func NewPackedColorSlice(data []Color) PackedColorArray {
	var array = NewPackedColorArray()
	array.Resize(Int(len(data)))
	copy(unsafe.Slice((*Color)(array.UnsafePointer()), len(data)), data)
	return array
}

func NewPackedFloat32Array() PackedFloat32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.PackedFloat32Array[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[PackedFloat32Array](raw)
}

func NewPackedFloat32Slice(data []float32) PackedFloat32Array {
	var array = NewPackedFloat32Array()
	array.Resize(Int(len(data)))
	copy(unsafe.Slice((*float32)(array.UnsafePointer()), len(data)), data)
	return array
}

func NewPackedInt32Array() PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.PackedInt32Array[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[PackedInt32Array](raw)
}

func NewPackedInt32Slice(data []int32) PackedInt32Array {
	var array = NewPackedInt32Array()
	array.Resize(Int(len(data)))
	copy(unsafe.Slice((*int32)(array.UnsafePointer()), len(data)), data)
	return array
}

func NewPackedStringArray() PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.PackedStringArray[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[PackedStringArray](raw)
}

func NewPackedStringSlice(data []string) PackedStringArray {
	var array = NewPackedStringArray()
	array.Resize(Int(len(data)))
	for i, str := range data {
		array.SetIndex(Int(i), NewString(str))
	}
	return array
}

func NewPackedVector2Array() PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.PackedVector2Array[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[PackedVector2Array](raw)
}

func NewPackedVector2Slice(data []Vector2) PackedVector2Array {
	var array = NewPackedVector2Array()
	array.Resize(Int(len(data)))
	copy(unsafe.Slice((*Vector2)(array.UnsafePointer()), len(data)), data)
	return array
}

func NewPackedVector3Array() PackedVector3Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.PackedVector3Array[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[PackedVector3Array](raw)
}

func NewPackedVector3Slice(data []Vector3) PackedVector3Array {
	var array = NewPackedVector3Array()
	array.Resize(Int(len(data)))
	copy(unsafe.Slice((*Vector3)(array.UnsafePointer()), len(data)), data)
	return array
}

func NewPackedVector4Array() PackedVector4Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.PackedVector4Array[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[PackedVector4Array](raw)
}

func NewPackedVector4Slice(data []Vector4) PackedVector4Array {
	var array = NewPackedVector4Array()
	array.Resize(Int(len(data)))
	copy(unsafe.Slice((*Vector4)(array.UnsafePointer()), len(data)), data)
	return array
}

func NewPackedInt64Array() PackedInt64Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.PackedInt64Array[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[PackedInt64Array](raw)
}

func NewPackedInt64Slice(data []int64) PackedInt64Array {
	var array = NewPackedInt64Array()
	array.Resize(Int(len(data)))
	copy(unsafe.Slice((*int64)(array.UnsafePointer()), len(data)), data)
	return array
}

func NewPackedFloat64Array() PackedFloat64Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	Global.typeset.creation.PackedFloat64Array[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[PackedFloat64Array](raw)
}

func NewPackedFloat64Slice(data []float64) PackedFloat64Array {
	var array = NewPackedFloat64Array()
	array.Resize(Int(len(data)))
	copy(unsafe.Slice((*float64)(array.UnsafePointer()), len(data)), data)
	return array
}
