//go:build !generate

package gd

import (
	"context"
	"reflect"

	"runtime.link/mmm"
)

func (ctx Context) Variant(v any) Variant {
	var godot = ctx.API()
	var frame = godot.NewFrame()
	if v == nil {
		godot.Variants.Zero(frame.Back())
		var variant = FrameGet[[3]uintptr](frame)
		frame.Free()
		return mmm.Make[API, Variant](ctx, godot, variant)
	}
	switch val := v.(type) {
	case Variant:
		FrameSet[[3]uintptr](0, frame, val.Pointer())
		godot.Variants.Copy(frame.Back(), frame.Get(0))
	case bool:
		FrameSet[bool](0, frame, val)
		godot.variant.FromType[TypeBool](frame.Back(), frame.Get(0))
	case Int:
		FrameSet[Int](0, frame, val)
		godot.variant.FromType[TypeInt](frame.Back(), frame.Get(0))
	case Float:
		FrameSet[Float](0, frame, val)
		godot.variant.FromType[TypeFloat](frame.Back(), frame.Get(0))
	case String:
		FrameSet[uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypeString](frame.Back(), frame.Get(0))
	case Vector2:
		FrameSet[Vector2](0, frame, val)
		godot.variant.FromType[TypeVector2](frame.Back(), frame.Get(0))
	case Vector2i:
		FrameSet[Vector2i](0, frame, val)
		godot.variant.FromType[TypeVector2i](frame.Back(), frame.Get(0))
	case Rect2:
		FrameSet[Rect2](0, frame, val)
		godot.variant.FromType[TypeRect2](frame.Back(), frame.Get(0))
	case Rect2i:
		FrameSet[Rect2i](0, frame, val)
		godot.variant.FromType[TypeRect2i](frame.Back(), frame.Get(0))
	case Vector3:
		FrameSet[Vector3](0, frame, val)
		godot.variant.FromType[TypeVector3](frame.Back(), frame.Get(0))
	case Vector3i:
		FrameSet[Vector3i](0, frame, val)
		godot.variant.FromType[TypeVector3i](frame.Back(), frame.Get(0))
	case Transform2D:
		FrameSet[Transform2D](0, frame, val)
		godot.variant.FromType[TypeTransform2d](frame.Back(), frame.Get(0))
	case Vector4:
		FrameSet[Vector4](0, frame, val)
		godot.variant.FromType[TypeVector4](frame.Back(), frame.Get(0))
	case Vector4i:
		FrameSet[Vector4i](0, frame, val)
		godot.variant.FromType[TypeVector4i](frame.Back(), frame.Get(0))
	case Plane:
		FrameSet[Plane](0, frame, val)
		godot.variant.FromType[TypePlane](frame.Back(), frame.Get(0))
	case Quaternion:
		FrameSet[Quaternion](0, frame, val)
		godot.variant.FromType[TypeQuaternion](frame.Back(), frame.Get(0))
	case AABB:
		FrameSet[AABB](0, frame, val)
		godot.variant.FromType[TypeAabb](frame.Back(), frame.Get(0))
	case Basis:
		FrameSet[Basis](0, frame, val)
		godot.variant.FromType[TypeBasis](frame.Back(), frame.Get(0))
	case Transform3D:
		FrameSet[Transform3D](0, frame, val)
		godot.variant.FromType[TypeTransform3d](frame.Back(), frame.Get(0))
	case Projection:
		FrameSet[Projection](0, frame, val)
		godot.variant.FromType[TypeProjection](frame.Back(), frame.Get(0))
	case Color:
		FrameSet[Color](0, frame, val)
		godot.variant.FromType[TypeColor](frame.Back(), frame.Get(0))
	case StringName:
		FrameSet[uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypeStringName](frame.Back(), frame.Get(0))
	case NodePath:
		FrameSet[uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypeNodePath](frame.Back(), frame.Get(0))
	case RID:
		FrameSet[RID](0, frame, val)
		godot.variant.FromType[TypeRid](frame.Back(), frame.Get(0))
	case Object:
		FrameSet[uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypeObject](frame.Back(), frame.Get(0))
	case Callable:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypeCallable](frame.Back(), frame.Get(0))
	case Signal:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypeSignal](frame.Back(), frame.Get(0))
	case Dictionary:
		FrameSet[uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypeDictionary](frame.Back(), frame.Get(0))
	case Array:
		FrameSet[uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypeArray](frame.Back(), frame.Get(0))
	case PackedByteArray:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypePackedByteArray](frame.Back(), frame.Get(0))
	case PackedInt32Array:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypePackedInt32Array](frame.Back(), frame.Get(0))
	case PackedInt64Array:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypePackedInt64Array](frame.Back(), frame.Get(0))
	case PackedFloat32Array:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypePackedFloat32Array](frame.Back(), frame.Get(0))
	case PackedFloat64Array:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypePackedFloat64Array](frame.Back(), frame.Get(0))
	case PackedStringArray:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypePackedStringArray](frame.Back(), frame.Get(0))
	case PackedVector2Array:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypePackedVector2Array](frame.Back(), frame.Get(0))
	case PackedVector3Array:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypePackedVector3Array](frame.Back(), frame.Get(0))
	case PackedColorArray:
		FrameSet[[2]uintptr](0, frame, val.Pointer())
		godot.variant.FromType[TypePackedColorArray](frame.Back(), frame.Get(0))
	default:
		panic("gd.Variant: unsupported type " + reflect.TypeOf(v).String())
	}
	frame.Free()
	return mmm.Make[API, Variant](ctx, godot, FrameGet[[3]uintptr](frame))
}

func (variant Variant) Type() VariantType {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var vtype = godot.Variants.Type(frame.Get(0))
	frame.Free()
	return vtype
}

func (variant Variant) Interface(ctx context.Context) any {
	switch variant.Type() {
	case TypeNil:
		return nil
	case TypeBool:
		return variant.asBool()
	case TypeInt:
		return variant.asInt()
	case TypeFloat:
		return variant.asFloat()
	case TypeString:
		return variant.asString(ctx)
	case TypeVector2:
		return variant.asVector2()
	case TypeVector2i:
		return variant.asVector2i()
	case TypeRect2:
		return variant.asRect2()
	case TypeRect2i:
		return variant.asRect2i()
	case TypeVector3:
		return variant.asVector3()
	case TypeVector3i:
		return variant.asVector3i()
	case TypeTransform2d:
		return variant.asTransform2D()
	case TypeVector4:
		return variant.asVector4()
	case TypeVector4i:
		return variant.asVector4i()
	case TypePlane:
		return variant.asPlane()
	case TypeQuaternion:
		return variant.asQuaternion()
	case TypeAabb:
		return variant.asAABB()
	case TypeBasis:
		return variant.asBasis()
	case TypeTransform3d:
		return variant.asTransform3D()
	case TypeProjection:
		return variant.asProjection()
	case TypeColor:
		return variant.asColor()
	case TypeStringName:
		return variant.asStringName(ctx)
	case TypeNodePath:
		return variant.asNodePath(ctx)
	case TypeRid:
		return variant.asRID()
	case TypeObject:
		return variant.asObject(ctx)
	case TypeCallable:
		return variant.asCallable(ctx)
	case TypeSignal:
		return variant.asSignal(ctx)
	case TypeDictionary:
		return variant.asDictionary(ctx)
	case TypeArray:
		return variant.asArray(ctx)
	case TypePackedByteArray:
		return variant.asPackedByteArray(ctx)
	case TypePackedInt32Array:
		return variant.asPackedInt32Array(ctx)
	case TypePackedInt64Array:
		return variant.asPackedInt64Array(ctx)
	case TypePackedFloat32Array:
		return variant.asPackedFloat32Array(ctx)
	case TypePackedFloat64Array:
		return variant.asPackedFloat64Array(ctx)
	case TypePackedStringArray:
		return variant.asPackedStringArray(ctx)
	case TypePackedVector2Array:
		return variant.asPackedVector2Array(ctx)
	case TypePackedVector3Array:
		return variant.asPackedVector3Array(ctx)
	case TypePackedColorArray:
		return variant.asPackedColorArray(ctx)
	default:
		panic("gd.Variant.Interface: invalid variant type")
	}
}

func (variant Variant) asBool() bool {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeBool](frame.Back(), frame.Get(0))
	var ret = FrameGet[bool](frame)
	frame.Free()
	return ret
}

func (variant Variant) asInt() Int {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeInt](frame.Back(), frame.Get(0))
	var ret = FrameGet[Int](frame)
	frame.Free()
	return ret
}

func (variant Variant) asFloat() Float {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeFloat](frame.Back(), frame.Get(0))
	var ret = FrameGet[Float](frame)
	frame.Free()
	return ret
}

func (variant Variant) asString(ctx context.Context) String {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeString](frame.Back(), frame.Get(0))
	var ret = FrameGet[uintptr](frame)
	frame.Free()
	return mmm.Make[API, String](ctx, godot, ret)
}

func (variant Variant) asVector2() Vector2 {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeVector2](frame.Back(), frame.Get(0))
	var ret = FrameGet[Vector2](frame)
	frame.Free()
	return ret
}

func (variant Variant) asVector2i() Vector2i {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeVector2i](frame.Back(), frame.Get(0))
	var ret = FrameGet[Vector2i](frame)
	frame.Free()
	return ret
}

func (variant Variant) asRect2() Rect2 {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeRect2](frame.Back(), frame.Get(0))
	var ret = FrameGet[Rect2](frame)
	frame.Free()
	return ret
}

func (variant Variant) asRect2i() Rect2i {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeRect2i](frame.Back(), frame.Get(0))
	var ret = FrameGet[Rect2i](frame)
	frame.Free()
	return ret
}

func (variant Variant) asVector3() Vector3 {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeVector3](frame.Back(), frame.Get(0))
	var ret = FrameGet[Vector3](frame)
	frame.Free()
	return ret
}

func (variant Variant) asVector3i() Vector3i {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeVector3i](frame.Back(), frame.Get(0))
	var ret = FrameGet[Vector3i](frame)
	frame.Free()
	return ret
}

func (variant Variant) asTransform2D() Transform2D {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeTransform2d](frame.Back(), frame.Get(0))
	var ret = FrameGet[Transform2D](frame)
	frame.Free()
	return ret
}

func (variant Variant) asVector4() Vector4 {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeVector4](frame.Back(), frame.Get(0))
	var ret = FrameGet[Vector4](frame)
	frame.Free()
	return ret
}

func (variant Variant) asVector4i() Vector4i {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypeVector4i](frame.Back(), frame.Get(0))
	var ret = FrameGet[Vector4i](frame)
	frame.Free()
	return ret
}

func (variant Variant) asPlane() Plane {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	godot.variant.IntoType[TypePlane](frame.Back(), frame.Get(0))
	var ret = FrameGet[Plane](frame)
	frame.Free()
	return ret
}

func (variant Variant) asQuaternion() Quaternion {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[Quaternion](frame)
	frame.Free()
	return ret
}

func (variant Variant) asAABB() AABB {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[AABB](frame)
	frame.Free()
	return ret
}

func (variant Variant) asBasis() Basis {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[Basis](frame)
	frame.Free()
	return ret
}

func (variant Variant) asTransform3D() Transform3D {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[Transform3D](frame)
	frame.Free()
	return ret
}

func (variant Variant) asProjection() Projection {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[Projection](frame)
	frame.Free()
	return ret
}

func (variant Variant) asColor() Color {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[Color](frame)
	frame.Free()
	return ret
}

func (variant Variant) asStringName(ctx context.Context) StringName {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[uintptr](frame)
	frame.Free()
	return mmm.Make[API, StringName](ctx, godot, ret)
}

func (variant Variant) asNodePath(ctx context.Context) NodePath {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[uintptr](frame)
	frame.Free()
	return mmm.Make[API, NodePath](ctx, godot, ret)
}

func (variant Variant) asRID() RID {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[RID](frame)
	frame.Free()
	return ret
}

func (variant Variant) asObject(ctx context.Context) Object {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[uintptr](frame)
	frame.Free()
	var object Object
	object.SetPointer(mmm.Make[API, Pointer](ctx, godot, ret))
	return object
}

func (variant Variant) asCallable(ctx context.Context) Callable {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, Callable](ctx, godot, ret)
}

func (variant Variant) asSignal(ctx context.Context) Signal {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, Signal](ctx, godot, ret)
}

func (variant Variant) asDictionary(ctx context.Context) Dictionary {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[uintptr](frame)
	frame.Free()
	return mmm.Make[API, Dictionary](ctx, godot, ret)
}

func (variant Variant) asArray(ctx context.Context) Array {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[uintptr](frame)
	frame.Free()
	return mmm.Make[API, Array](ctx, godot, ret)
}

func (variant Variant) asPackedByteArray(ctx context.Context) PackedByteArray {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, PackedByteArray](ctx, godot, ret)
}

func (variant Variant) asPackedInt32Array(ctx context.Context) PackedInt32Array {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, PackedInt32Array](ctx, godot, ret)
}

func (variant Variant) asPackedInt64Array(ctx context.Context) PackedInt64Array {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, PackedInt64Array](ctx, godot, ret)
}

func (variant Variant) asPackedFloat32Array(ctx context.Context) PackedFloat32Array {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, PackedFloat32Array](ctx, godot, ret)
}

func (variant Variant) asPackedFloat64Array(ctx context.Context) PackedFloat64Array {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, PackedFloat64Array](ctx, godot, ret)
}

func (variant Variant) asPackedStringArray(ctx context.Context) PackedStringArray {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, PackedStringArray](ctx, godot, ret)
}

func (variant Variant) asPackedVector2Array(ctx context.Context) PackedVector2Array {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, PackedVector2Array](ctx, godot, ret)
}

func (variant Variant) asPackedVector3Array(ctx context.Context) PackedVector3Array {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, PackedVector3Array](ctx, godot, ret)
}

func (variant Variant) asPackedColorArray(ctx context.Context) PackedColorArray {
	var godot = variant.API
	var frame = godot.NewFrame()
	FrameSet[[3]uintptr](0, frame, variant.Pointer())
	var ret = FrameGet[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[API, PackedColorArray](ctx, godot, ret)
}
