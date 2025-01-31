package variant

import (
	"fmt"
	"reflect"
	"unsafe"

	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Plane"
	"graphics.gd/variant/Projection"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector3i"
	"graphics.gd/variant/Vector4"
	"graphics.gd/variant/Vector4i"
)

// Any is like the standard [any] type except it can store the following types without allocating them on the heap:
//   - [bool]
//   - [int][int8][int16][int32][int64][uint][uint8][uint16][uint32][uint64][uintptr]
//   - [Float.X][float32][float64][complex64][complex128]
//   - [string][String][[]byte][[]int32][[]int64][[]float32][[]float64][[]string]
//   - [[]Vector2.XY][[]Vector3.XYZ][[]Color.RGBA][[]Vector4.XYZW]
//   - [Rect2.PositionSize][Rect2i.PositionSize][Plane.NormalD]
//   - [Vector2.XY][Vector2i.XY][Vector3.XYZ][Vector3i.XYZ][Vector4.XYZW][Vector4i.XYZW][Color.RGBA]
type Any struct {
	local complex128
	value any
}

// Nil reference value.
var Nil Any

func load[T any](a Any) T {
	_, ok := a.value.(*T)
	if !ok {
		if val, ok := a.value.(T); ok {
			return val
		}
		panic("variant conversion: variant is " + reflect.TypeOf(a.value).String() + ", not " + reflect.TypeFor[T]().String())
	}
	return *(*T)(unsafe.Pointer(&a.local))
}
func loadPacked[T packable](a Any) []T {
	ptr, ok := a.value.(isPacked[T])
	if !ok {
		panic("variant conversion: variant is " + a.Type().String() + ", not " + reflect.TypeFor[T]().String())
	}
	lencap := *(*lencap)(unsafe.Pointer(&a.local))
	return unsafe.Slice(ptr, lencap.cap)[:lencap.len:lencap.cap]
}

func (a Any) Bool() bool {
	if proxy, ok := a.value.(API); ok {
		return proxy.Bool(a.local)
	}
	return load[bool](a)
}
func (a Any) Int() int {
	if proxy, ok := a.value.(API); ok {
		return int(proxy.Int(a.local))
	}
	rtype := reflect.TypeOf(a.value)
	if rtype.Kind() != reflect.Ptr {
		panic("variant conversion: variant is " + a.Type().String() + ", not " + reflect.TypeFor[int]().String())
	}
	switch rtype.Elem().Kind() {
	case reflect.Int8:
		return int(*(*int8)(unsafe.Pointer(&a.local)))
	case reflect.Int16:
		return int(*(*int16)(unsafe.Pointer(&a.local)))
	case reflect.Int32:
		return int(*(*int32)(unsafe.Pointer(&a.local)))
	case reflect.Int64:
		return int(*(*int64)(unsafe.Pointer(&a.local)))
	case reflect.Int:
		return int(*(*int)(unsafe.Pointer(&a.local)))
	default:
		panic("variant conversion: variant is " + a.Type().String() + ", not " + reflect.TypeFor[int]().String())
	}
}
func (a Any) Int8() int8 {
	if proxy, ok := a.value.(API); ok {
		return int8(proxy.Int(a.local))
	}
	return load[int8](a)
}
func (a Any) Int16() int16 {
	if proxy, ok := a.value.(API); ok {
		return int16(proxy.Int(a.local))
	}
	return load[int16](a)
}
func (a Any) Int32() int32 {
	if proxy, ok := a.value.(API); ok {
		return int32(proxy.Int(a.local))
	}
	return load[int32](a)
}
func (a Any) Int64() int64 {
	if proxy, ok := a.value.(API); ok {
		return int64(proxy.Int(a.local))
	}
	return load[int64](a)
}
func (a Any) Uint() uint {
	if proxy, ok := a.value.(API); ok {
		return uint(proxy.RID(a.local))
	}
	return uint(load[uint](a))
}
func (a Any) Uint8() uint8 {
	if proxy, ok := a.value.(API); ok {
		return uint8(proxy.Int(a.local))
	}
	return load[uint8](a)
}
func (a Any) Uint16() uint16 {
	if proxy, ok := a.value.(API); ok {
		return uint16(proxy.Int(a.local))
	}
	return load[uint16](a)
}
func (a Any) Uint32() uint32 {
	if proxy, ok := a.value.(API); ok {
		return uint32(proxy.Int(a.local))
	}
	return load[uint32](a)
}
func (a Any) Uint64() uint64 {
	if proxy, ok := a.value.(API); ok {
		return uint64(proxy.RID(a.local))
	}
	return load[uint64](a)
}
func (a Any) RID() RID.Any {
	return RID.Any(a.Uint64())
}
func (a Any) Float32() float32 {
	if proxy, ok := a.value.(API); ok {
		return float32(proxy.Float(a.local))
	}
	return load[float32](a)
}
func (a Any) Float64() float64 {
	if proxy, ok := a.value.(API); ok {
		return float64(proxy.Float(a.local))
	}
	return load[float64](a)
}
func (a Any) Uintptr() uintptr {
	if proxy, ok := a.value.(API); ok {
		return uintptr(proxy.RID(a.local))
	}
	return uintptr(load[uintptr](a))
}
func (a Any) Complex64() complex64 {
	if proxy, ok := a.value.(API); ok {
		vec2 := proxy.Vector2(a.local)
		return complex(float32(vec2.X), float32(vec2.Y))
	}
	return load[complex64](a)
}
func (a Any) Complex128() complex128 {
	if proxy, ok := a.value.(API); ok {
		vec2 := proxy.Vector2(a.local)
		return complex(float64(vec2.X), float64(vec2.Y))
	}
	return load[complex128](a)
}
func (a Any) String() string {
	if proxy, ok := a.value.(API); ok {
		return proxy.String(a.local)
	}
	ptr, ok := a.value.(isString)
	if !ok {
		return a.toString()
	}
	l := *(*int)(unsafe.Pointer(&a.local))
	return unsafe.String(ptr, l)
}
func (a Any) Vector2() Vector2.XY {
	if proxy, ok := a.value.(API); ok {
		return proxy.Vector2(a.local)
	}
	return load[Vector2.XY](a)
}
func (a Any) Vector2i() Vector2i.XY {
	if proxy, ok := a.value.(API); ok {
		return proxy.Vector2i(a.local)
	}
	return load[Vector2i.XY](a)
}
func (a Any) Rect2() Rect2.PositionSize {
	if proxy, ok := a.value.(API); ok {
		return proxy.Rect2(a.local)
	}
	return load[Rect2.PositionSize](a)
}
func (a Any) Rect2i() Rect2i.PositionSize {
	if proxy, ok := a.value.(API); ok {
		return proxy.Rect2i(a.local)
	}
	return load[Rect2i.PositionSize](a)
}
func (a Any) Vector3() Vector3.XYZ {
	if proxy, ok := a.value.(API); ok {
		return proxy.Vector3(a.local)
	}
	return load[Vector3.XYZ](a)
}
func (a Any) Vector3i() Vector3i.XYZ {
	if proxy, ok := a.value.(API); ok {
		return proxy.Vector3i(a.local)
	}
	return load[Vector3i.XYZ](a)
}
func (a Any) Transform2D() Transform2D.OriginXY {
	if proxy, ok := a.value.(API); ok {
		return proxy.Transform2D(a.local)
	}
	return load[Transform2D.OriginXY](a)
}
func (a Any) Vector4() Vector4.XYZW {
	if proxy, ok := a.value.(API); ok {
		return proxy.Vector4(a.local)
	}
	return load[Vector4.XYZW](a)
}
func (a Any) Vector4i() Vector4i.XYZW {
	if proxy, ok := a.value.(API); ok {
		return proxy.Vector4i(a.local)
	}
	return load[Vector4i.XYZW](a)
}
func (a Any) Plane() Plane.NormalD {
	if proxy, ok := a.value.(API); ok {
		return proxy.Plane(a.local)
	}
	return load[Plane.NormalD](a)
}
func (a Any) Quaternion() Quaternion.IJKX {
	if proxy, ok := a.value.(API); ok {
		return proxy.Quaternion(a.local)
	}
	return a.value.(Quaternion.IJKX)
}
func (a Any) AABB() AABB.PositionSize {
	if proxy, ok := a.value.(API); ok {
		return proxy.AABB(a.local)
	}
	return a.value.(AABB.PositionSize)
}
func (a Any) Basis() Basis.XYZ {
	if proxy, ok := a.value.(API); ok {
		return proxy.Basis(a.local)
	}
	return a.value.(Basis.XYZ)
}
func (a Any) Transform3D() Transform3D.BasisOrigin {
	if proxy, ok := a.value.(API); ok {
		return proxy.Transform3D(a.local)
	}
	return a.value.(Transform3D.BasisOrigin)
}
func (a Any) Projection() Projection.XYZW {
	if proxy, ok := a.value.(API); ok {
		return proxy.Projection(a.local)
	}
	return a.value.(Projection.XYZW)
}
func (a Any) Color() Color.RGBA {
	if proxy, ok := a.value.(API); ok {
		return proxy.Color(a.local)
	}
	return load[Color.RGBA](a)
}
func (a Any) Bytes() []byte {
	if proxy, ok := a.value.(API); ok {
		return proxy.Bytes(a.local)
	}
	return loadPacked[byte](a)
}

func (a Any) toString() string {
	return fmt.Sprint(a.Interface())
}

// Interface returns the value of the variant as an interface{}.
func (a Any) Interface() interface{} {
	if a == Nil {
		return nil
	}
	rtype := reflect.TypeOf(a.value)
	rvalue := reflect.ValueOf(a.value)
	switch rtype.Kind() {
	case reflect.Pointer:
		if rvalue.Pointer() != 0 {
			switch value := a.value.(type) {
			case isString:
				l := *(*int)(unsafe.Pointer(&a.local))
				return unsafe.String(value, l)
			case isPacked[byte]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[int32]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[int64]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[float32]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[float64]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[string]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[Vector2.XY]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[Vector3.XYZ]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[Color.RGBA]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[Vector4.XYZW]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			}
			lencap := *(*lencap)(unsafe.Pointer(&a.local))
			if lencap.len > 0 {
				return reflect.SliceAt(rtype.Elem(), rvalue.UnsafePointer(), lencap.cap).Slice3(0, lencap.len, lencap.cap).Interface()
			}
			return a.value
		}
		var heap = a.local
		return reflect.NewAt(rtype.Elem(), unsafe.Pointer(&heap)).Elem().Interface()
	default:
		return a.value
	}
}
