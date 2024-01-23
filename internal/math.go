//go:build !generate

package gd

import "math"

const cmpEpsilon = 0.00001

type ComponentWise[T any] interface {
	Float | Int | Vector2 | Vector2i | Vector3 | Vector3i | Vector4 | Vector4i

	abs() T
	ceil() T
	floor() T
	round() T
	sign() T
	snapped(T, T) T
}

type Lerpable[T any] interface {
	Float | Int | Vector2 | Vector3 | Vector4 | Color | Quaternion | Basis

	lerp(T, Float) T
}

func Abs[T ComponentWise[T]](val T) T   { return val.abs() }
func Ceil[T ComponentWise[T]](val T) T  { return val.ceil() }
func Floor[T ComponentWise[T]](val T) T { return val.floor() }
func Round[T ComponentWise[T]](val T) T { return val.round() }
func Sign[T ComponentWise[T]](val T) T  { return val.sign() }
func Snapped[T ComponentWise[T]](val T, by T) T {
	return val.snapped(by, by)
}

func Lerp[T Lerpable[T]](a, b T, t Float) T { return a.lerp(b, t) }

func absf[T ~float32 | ~float64](val T) T { return T(math.Abs(float64(val))) }

func absi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](val T) T {
	if val < 0 {
		return -val
	}
	return val
}

func sign[T ~float32 | ~float64 | ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func ceilf[T ~float32 | ~float64](val T) T       { return T(math.Ceil(float64(val))) }
func floorf[T ~float32 | ~float64](val T) T      { return T(math.Floor(float64(val))) }
func lerpf[T ~float32 | ~float64](a, b T, t T) T { return a + (b-a)*t }
func snapf[T ~float32 | ~float64](val T, by T) T {
	if by != 0 {
		val = T(math.Floor(float64(val/by)+0.5)) * by
	}
	return val
}
func snapi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](val T, by T) T {
	if by != 0 {
		val = (val / by) * by
	}
	return val
}

func (i Int) abs() Int           { return absi(i) }
func (f Float) abs() Float       { return absf(f) }
func (v Vector2) abs() Vector2   { return Vector2{absf(v[x]), absf(v[y])} }
func (v Vector2i) abs() Vector2i { return Vector2i{absi(v[x]), absi(v[y])} }
func (v Vector3) abs() Vector3   { return Vector3{absf(v[x]), absf(v[y]), absf(v[z])} }
func (v Vector3i) abs() Vector3i { return Vector3i{absi(v[x]), absi(v[y]), absi(v[z])} }
func (v Vector4) abs() Vector4   { return Vector4{absf(v[x]), absf(v[y]), absf(v[z]), absf(v[w])} }
func (v Vector4i) abs() Vector4i { return Vector4i{absi(v[x]), absi(v[y]), absi(v[z]), absi(v[w])} }

func (i Int) ceil() Int           { return i }
func (f Float) ceil() Float       { return Float(math.Ceil(float64(f))) }
func (v Vector2) ceil() Vector2   { return Vector2{ceilf(v[x]), ceilf(v[y])} }
func (v Vector2i) ceil() Vector2i { return v }
func (v Vector3) ceil() Vector3   { return Vector3{ceilf(v[x]), ceilf(v[y]), ceilf(v[z])} }
func (v Vector3i) ceil() Vector3i { return v }
func (v Vector4) ceil() Vector4   { return Vector4{ceilf(v[x]), ceilf(v[y]), ceilf(v[z]), ceilf(v[w])} }
func (v Vector4i) ceil() Vector4i { return v }

func (i Int) floor() Int           { return i }
func (f Float) floor() Float       { return Float(math.Floor(float64(f))) }
func (v Vector2) floor() Vector2   { return Vector2{floorf(v[x]), floorf(v[y])} }
func (v Vector2i) floor() Vector2i { return v }
func (v Vector3) floor() Vector3   { return Vector3{floorf(v[x]), floorf(v[y]), floorf(v[z])} }
func (v Vector3i) floor() Vector3i { return v }
func (v Vector4) floor() Vector4 {
	return Vector4{floorf(v[x]), floorf(v[y]), floorf(v[z]), floorf(v[w])}
}
func (v Vector4i) floor() Vector4i { return v }

func (i Int) round() Int     { return i }
func (f Float) round() Float { return Float(math.Round(float64(f))) }
func (v Vector2) round() Vector2 {
	return Vector2{float32(math.Round(float64(v[x]))), float32(math.Round(float64(v[y])))}
}
func (v Vector2i) round() Vector2i { return v }
func (v Vector3) round() Vector3 {
	return Vector3{float32(math.Round(float64(v[x]))), float32(math.Round(float64(v[y]))), float32(math.Round(float64(v[z])))}
}
func (v Vector3i) round() Vector3i { return v }
func (v Vector4) round() Vector4 {
	return Vector4{float32(math.Round(float64(v[x]))), float32(math.Round(float64(v[y]))), float32(math.Round(float64(v[z]))), float32(math.Round(float64(v[w])))}
}
func (v Vector4i) round() Vector4i { return v }

func (i Int) sign() Int     { return sign(i) }
func (f Float) sign() Float { return sign(f) }
func (v Vector2) sign() Vector2 {
	return Vector2{sign(v[x]), sign(v[y])}
}
func (v Vector2i) sign() Vector2i { return Vector2i{sign(v[x]), sign(v[y])} }
func (v Vector3) sign() Vector3 {
	return Vector3{sign(v[x]), sign(v[y]), sign(v[z])}
}
func (v Vector3i) sign() Vector3i { return Vector3i{sign(v[x]), sign(v[y]), sign(v[z])} }
func (v Vector4) sign() Vector4 {
	return Vector4{sign(v[x]), sign(v[y]), sign(v[z]), sign(v[w])}
}
func (v Vector4i) sign() Vector4i { return Vector4i{sign(v[x]), sign(v[y]), sign(v[z]), sign(v[w])} }

func (i Int) snapped(by, margin Int) Int       { return snapi(i, by) }
func (f Float) snapped(by, margin Float) Float { return snapf(f, by) }
func (v Vector2) snapped(by, margin Vector2) Vector2 {
	return Vector2{snapf(v[x], by[x]), snapf(v[y], by[y])}
}
func (v Vector2i) snapped(by, margin Vector2i) Vector2i {
	return Vector2i{snapi(v[x], by[x]), snapi(v[y], by[y])}
}
func (v Vector3) snapped(by, margin Vector3) Vector3 {
	return Vector3{snapf(v[x], by[x]), snapf(v[y], by[y]), snapf(v[z], by[z])}
}
func (v Vector3i) snapped(by, margin Vector3i) Vector3i {
	return Vector3i{snapi(v[x], by[x]), snapi(v[y], by[y]), snapi(v[z], by[z])}
}
func (v Vector4) snapped(by, margin Vector4) Vector4 {
	return Vector4{snapf(v[x], by[x]), snapf(v[y], by[y]), snapf(v[z], by[z]), snapf(v[w], by[w])}
}
func (v Vector4i) snapped(by, margin Vector4i) Vector4i {
	return Vector4i{snapi(v[x], by[x]), snapi(v[y], by[y]), snapi(v[z], by[z]), snapi(v[w], by[w])}
}

func (a Int) lerp(b Int, t Float) Int       { return a + Int(Float(b-a)*t) }
func (a Float) lerp(b Float, t Float) Float { return a + (b-a)*t }
func (a Vector2) lerp(b Vector2, t Float) Vector2 {
	return Vector2{a[x] + (b[x]-a[x])*float32(t), a[y] + (b[y]-a[y])*float32(t)}
}
func (a Vector3) lerp(b Vector3, t Float) Vector3 {
	return Vector3{a[x] + (b[x]-a[x])*float32(t), a[y] + (b[y]-a[y])*float32(t), a[z] + (b[z]-a[z])*float32(t)}
}
func (a Vector4) lerp(b Vector4, t Float) Vector4 {
	return Vector4{a[x] + (b[x]-a[x])*float32(t), a[y] + (b[y]-a[y])*float32(t), a[z] + (b[z]-a[z])*float32(t), a[w] + (b[w]-a[w])*float32(t)}
}
func (a Color) lerp(b Color, t Float) Color {
	return Color{a[x] + (b[x]-a[x])*float32(t), a[y] + (b[y]-a[y])*float32(t), a[z] + (b[z]-a[z])*float32(t), a[w] + (b[w]-a[w])*float32(t)}
}

func dot4[T ~[4]float32](val T) float32 {
	return val[x]*val[x] + val[y]*val[y] + val[z]*val[z] + val[w]*val[w]
}
func neg4[T ~[4]float32](val T) T {
	return T{-val[x], -val[y], -val[z], -val[w]}
}

func (a Quaternion) lerp(p_to Quaternion, p_weight Float) Quaternion {
	var to1 Quaternion
	var omega, cosom, sinom, scale0, scale1 float32

	// calc cosine
	cosom = dot4(p_to)

	// adjust signs (if necessary)
	if cosom < 0.0 {
		cosom = -cosom
		to1 = neg4(p_to)
	} else {
		to1 = p_to
	}

	// calculate coefficients

	if (1.0 - cosom) > cmpEpsilon {
		// standard case (slerp)
		omega = float32(math.Acos(float64(cosom)))
		sinom = float32(math.Sin(float64(omega)))
		scale0 = float32(math.Sin(float64((1.0-p_weight)*Float(omega)))) / sinom
		scale1 = float32(math.Sin(float64(p_weight*Float(omega)))) / sinom
	} else {
		// "from" and "to" quaternions are very close
		//  ... so we can do a linear interpolation
		scale0 = float32(1.0 - p_weight)
		scale1 = float32(p_weight)
	}
	// calculate final values
	return Quaternion{
		scale0*a[x] + scale1*to1[x],
		scale0*a[y] + scale1*to1[y],
		scale0*a[z] + scale1*to1[z],
		scale0*a[w] + scale1*to1[w]}
}

func (a Basis) lerp(p_to Basis, p_weight Float) Basis {
	//consider scale
	var from = a.GetQuaternion()
	var to = p_to.GetQuaternion()
	var b = from.lerp(to, p_weight).Basis()
	b.rows[0] = b.rows[0].ScaledBy(lerpf(a.rows[0].Length(), p_to.rows[0].Length(), p_weight))
	b.rows[1] = b.rows[1].ScaledBy(lerpf(a.rows[1].Length(), p_to.rows[1].Length(), p_weight))
	b.rows[2] = b.rows[2].ScaledBy(lerpf(a.rows[2].Length(), p_to.rows[2].Length(), p_weight))
	return b
}
