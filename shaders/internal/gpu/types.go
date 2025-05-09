package gpu

import (
	"reflect"

	"graphics.gd/classdb/Cubemap"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/classdb/Texture2DArray"
	"graphics.gd/classdb/Texture3D"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Color"
	FloatVariant "graphics.gd/variant/Float"
	"graphics.gd/variant/Projection"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector3i"
	"graphics.gd/variant/Vector4"
)

type internalExpression = Expression

type isEquivalentTo[T any] interface {
	equivalentTo(T)
}

type Sampler2D[T any] struct {
	internalExpression
	isEquivalentTo[Texture2D.Instance]
}

func (Sampler2D[T]) sampler2D() reflect.Type { return reflect.TypeFor[T]() }

type IsSampler2D interface {
	sampler2D() reflect.Type
}

type Sampler3D[T any] struct {
	internalExpression
	isEquivalentTo[Texture3D.Instance]
}

func (Sampler3D[T]) sampler3D() reflect.Type { return reflect.TypeFor[T]() }

type IsSampler3D interface {
	sampler3D() reflect.Type
}

type ArraySampler2D[T any] struct {
	internalExpression
	isEquivalentTo[Texture2DArray.Instance]
}

func (ArraySampler2D[T]) arraySampler2D() reflect.Type { return reflect.TypeFor[T]() }

type IsArraySampler2D interface {
	arraySampler2D() reflect.Type
}

type ArraySampler3D[T any] struct {
	internalExpression
}

func (ArraySampler3D[T]) arraySampler3D() reflect.Type { return reflect.TypeFor[T]() }

type IsArraySampler3D interface {
	arraySampler3D() reflect.Type
}

type CubeSampler[T any] struct {
	internalExpression
	isEquivalentTo[Cubemap.Instance]
}

func (CubeSampler[T]) cubeSampler() reflect.Type { return reflect.TypeFor[T]() }

type IsCubeSampler interface {
	cubeSampler() reflect.Type
}

func SamplerType(val any) reflect.Type {
	switch v := val.(type) {
	case IsSampler2D:
		return v.sampler2D()
	case IsSampler3D:
		return v.sampler3D()
	case IsArraySampler2D:
		return v.arraySampler2D()
	case IsArraySampler3D:
		return v.arraySampler3D()
	case IsCubeSampler:
		return v.cubeSampler()
	default:
		panic("unsupported sampler type " + reflect.TypeOf(v).String())
	}
}

type Bool = struct {
	internalExpression
	isEquivalentTo[bool]

	X bool
}

type AnyBool interface {
	~bool | ~Bool
}

func NewBool[T AnyBool](x T) Bool {
	rvalue := reflect.ValueOf(x)
	if rvalue.Kind() == reflect.Bool {
		return Bool{X: rvalue.Bool()}
	}
	return rvalue.Convert(reflect.TypeFor[Bool]()).Interface().(Bool)
}

func NewBoolExpression(e Expression) Bool { return Bool{internalExpression: e} }

type Int = struct {
	internalExpression
	isEquivalentTo[int32]

	X int
}

type AnyInt interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~Int
}

func NewInt[T AnyInt](x T) Int {
	rvalue := reflect.ValueOf(x)
	if rvalue.Kind() != reflect.Struct {
		return Int{X: int(rvalue.Int())}
	}
	return rvalue.Convert(reflect.TypeFor[Int]()).Interface().(Int)
}
func NewIntExpression(e Expression) Int { return Int{internalExpression: e} }

type Uint = struct {
	internalExpression
	isEquivalentTo[uint32]

	X uint
}

type AnyUint interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr | ~Uint
}

func NewUint[T AnyUint](x T) Uint {
	rvalue := reflect.ValueOf(x)
	if rvalue.Kind() != reflect.Struct {
		return Uint{X: uint(rvalue.Uint())}
	}
	return rvalue.Convert(reflect.TypeFor[Uint]()).Interface().(Uint)
}
func NewUintExpression(e Expression) Uint { return Uint{internalExpression: e} }

type Float = struct {
	internalExpression
	isEquivalentTo[FloatVariant.X]

	X float64
}

type AnyFloat interface {
	~float32 | ~float64 | ~Float
}

func NewFloat[T AnyFloat](x T) Float {
	rvalue := reflect.ValueOf(x)
	if rvalue.Kind() != reflect.Struct {
		return Float{X: rvalue.Float()}
	}
	return rvalue.Convert(reflect.TypeFor[Float]()).Interface().(Float)
}
func NewFloatExpression(e Expression) Float { return Float{internalExpression: e} }

type Vec2b = struct {
	internalExpression
	isEquivalentTo[[2]bool]

	X, Y Bool
}

func NewVec2b[X, Y AnyBool](x X, y Y) Vec2b { return Vec2b{X: NewBool(x), Y: NewBool(y)} }
func NewVec2bExpression(e Expression) Vec2b { return Vec2b{internalExpression: e} }

type Vec3b = struct {
	internalExpression
	isEquivalentTo[[3]bool]

	X, Y, Z Bool
}

func NewVec3b[X, Y, Z AnyBool](x X, y Y, z Z) Vec3b {
	return Vec3b{X: NewBool(x), Y: NewBool(y), Z: NewBool(z)}
}
func NewVec3bExpression(e Expression) Vec3b { return Vec3b{internalExpression: e} }

type Vec4b = struct {
	internalExpression
	isEquivalentTo[[4]bool]

	X, Y, Z, W Bool
}

func NewVec4b[X, Y, Z, W AnyBool](x X, y Y, z Z, w W) Vec4b {
	return Vec4b{X: NewBool(x), Y: NewBool(y), Z: NewBool(z), W: NewBool(w)}
}
func NewVec4bExpression(e Expression) Vec4b { return Vec4b{internalExpression: e} }

type Vec2i = struct {
	internalExpression
	isEquivalentTo[Vector2i.XY]

	X, Y Int
}

func NewVec2i[X, Y AnyInt](x X, y Y) Vec2i  { return Vec2i{X: NewInt(x), Y: NewInt(y)} }
func NewVec2iExpression(e Expression) Vec2i { return Vec2i{internalExpression: e} }

type Vec3i = struct {
	internalExpression
	isEquivalentTo[Vector3i.XYZ]

	X, Y, Z Int
}

func NewVec3i[X, Y, Z AnyInt](x X, y Y, z Z) Vec3i {
	return Vec3i{X: NewInt(x), Y: NewInt(y), Z: NewInt(z)}
}
func NewVec3iExpression(e Expression) Vec3i { return Vec3i{internalExpression: e} }

type Vec4i = struct {
	internalExpression
	isEquivalentTo[[4]int32]

	X, Y, Z, W Int
}

func NewVec4i[X, Y, Z, W AnyInt](x X, y Y, z Z, w W) Vec4i {
	return Vec4i{X: NewInt(x), Y: NewInt(y), Z: NewInt(z), W: NewInt(w)}
}
func NewVec4iExpression(e Expression) Vec4i { return Vec4i{internalExpression: e} }

type Vec2u = struct {
	internalExpression
	isEquivalentTo[[2]uint32]

	X, Y Uint
}

func NewVec2u[X, Y AnyUint](x X, y Y) Vec2u { return Vec2u{X: NewUint(x), Y: NewUint(y)} }
func NewVec2uExpression(e Expression) Vec2u { return Vec2u{internalExpression: e} }

type Vec3u = struct {
	internalExpression
	isEquivalentTo[[3]uint32]

	X, Y, Z Uint
}

func NewVec3u[X, Y, Z AnyUint](x X, y Y, z Z) Vec3u {
	return Vec3u{X: NewUint(x), Y: NewUint(y), Z: NewUint(z)}
}
func NewVec3uExpression(e Expression) Vec3u { return Vec3u{internalExpression: e} }

type Vec4u = struct {
	internalExpression
	isEquivalentTo[[4]uint32]

	X, Y, Z, W Uint
}

func NewVec4u[X, Y, Z, W AnyUint](x X, y Y, z Z, w W) Vec4u {
	return Vec4u{X: NewUint(x), Y: NewUint(y), Z: NewUint(z), W: NewUint(w)}
}
func NewVec4uExpression(e Expression) Vec4u { return Vec4u{internalExpression: e} }

type Vec2 = struct {
	internalExpression
	isEquivalentTo[Vector2.XY]

	X, Y Float
}

func NewVec2[X, Y AnyFloat](x X, y Y) Vec2 { return Vec2{X: NewFloat(x), Y: NewFloat(y)} }
func NewVec2Expression(e Expression) Vec2  { return Vec2{internalExpression: e} }

type Vec3 = struct {
	internalExpression
	isEquivalentTo[Vector3.XYZ]

	X, Y, Z Float
}

func NewVec3[X, Y, Z AnyFloat](x X, y Y, z Z) Vec3 {
	return Vec3{X: NewFloat(x), Y: NewFloat(y), Z: NewFloat(z)}
}
func NewVec3Expression(e Expression) Vec3 { return Vec3{internalExpression: e} }

type Vec4 = struct {
	internalExpression
	isEquivalentTo[Vector4.XYZW]

	X, Y, Z, W Float
}

func NewVec4[X, Y, Z, W AnyFloat](x X, y Y, z Z, w W) Vec4 {
	return Vec4{X: NewFloat(x), Y: NewFloat(y), Z: NewFloat(z), W: NewFloat(w)}
}
func NewVec4Expression(e Expression) Vec4 { return Vec4{internalExpression: e} }

type RGBA = struct {
	internalExpression
	isEquivalentTo[Color.RGBA]

	R, G, B, A Float
}

func NewRGBA[R, G, B, A AnyFloat](r R, g G, b B, a A) RGBA {
	return RGBA{R: NewFloat(r), G: NewFloat(g), B: NewFloat(b), A: NewFloat(a)}
}
func NewRGBAExpression(e Expression) RGBA { return RGBA{internalExpression: e} }

type RGB = struct {
	internalExpression
	isEquivalentTo[Color.RGB]

	R, G, B Float
}

func NewRGB[R, G, B AnyFloat](r R, g G, b B) RGB {
	return RGB{R: NewFloat(r), G: NewFloat(g), B: NewFloat(b)}
}
func NewRGBExpression(e Expression) RGB { return RGB{internalExpression: e} }

type Quad interface {
	~Vec4 | ~Vec4i | ~Vec4u | ~RGBA
}

func NewQuadExpression[Q Quad](expr Expression) Q {
	rtype := reflect.TypeFor[Q]()
	switch {
	case rtype.ConvertibleTo(reflect.TypeOf(Vec4{})):
		return reflect.ValueOf(NewVec4Expression(expr)).Convert(rtype).Interface().(Q)
	case rtype.ConvertibleTo(reflect.TypeOf(Vec4i{})):
		return reflect.ValueOf(NewVec4iExpression(expr)).Convert(rtype).Interface().(Q)
	case rtype.ConvertibleTo(reflect.TypeOf(Vec4u{})):
		return reflect.ValueOf(NewVec4uExpression(expr)).Convert(rtype).Interface().(Q)
	case rtype.ConvertibleTo(reflect.TypeOf(RGBA{})):
		return reflect.ValueOf(NewRGBAExpression(expr)).Convert(rtype).Interface().(Q)
	default:
		panic("unreachable")
	}
}

type Mat2 = struct {
	internalExpression
	isEquivalentTo[Transform2D.OriginXY]

	Columns [2][2]Float
}

func NewMat2[C00, C01, C10, C11 AnyFloat](c00 C00, c01 C01, c10 C10, c11 C11) Mat2 {
	return Mat2{Columns: [2][2]Float{
		{NewFloat(c00), NewFloat(c01)},
		{NewFloat(c10), NewFloat(c11)},
	}}
}
func NewMat2Expression(e Expression) Mat2 { return Mat2{internalExpression: e} }

type Mat3 = struct {
	internalExpression
	isEquivalentTo[Basis.XYZ]

	Columns [3][3]Float
}

func NewMat3[C00, C01, C02, C10, C11, C12, C20, C21, C22 AnyFloat](c00 C00, c01 C01, c02 C02, c10 C10, c11 C11, c12 C12, c20 C20, c21 C21, c22 C22) Mat3 {
	return Mat3{Columns: [3][3]Float{
		{NewFloat(c00), NewFloat(c01), NewFloat(c02)},
		{NewFloat(c10), NewFloat(c11), NewFloat(c12)},
		{NewFloat(c20), NewFloat(c21), NewFloat(c22)},
	}}
}

func NewMat3Expression(e Expression) Mat3 { return Mat3{internalExpression: e} }

type Mat4 = struct {
	internalExpression
	isEquivalentTo[Projection.XYZW]

	Columns [4][4]Float
}

func NewMat4[C00, C01, C02, C03, C10, C11, C12, C13, C20, C21, C22, C23, C30, C31, C32, C33 AnyFloat](c00 C00, c01 C01, c02 C02, c03 C03, c10 C10, c11 C11, c12 C12, c13 C13, c20 C20, c21 C21, c22 C22, c23 C23, c30 C30, c31 C31, c32 C32, c33 C33) Mat4 {
	return Mat4{Columns: [4][4]Float{
		{NewFloat(c00), NewFloat(c01), NewFloat(c02), NewFloat(c03)},
		{NewFloat(c10), NewFloat(c11), NewFloat(c12), NewFloat(c13)},
		{NewFloat(c20), NewFloat(c21), NewFloat(c22), NewFloat(c23)},
		{NewFloat(c30), NewFloat(c31), NewFloat(c32), NewFloat(c33)},
	}}
}

func NewMat4Expression(e Expression) Mat4 { return Mat4{internalExpression: e} }
