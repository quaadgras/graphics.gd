package variant

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"unsafe"

	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/NodePath"
	"graphics.gd/variant/Plane"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
)

// Marshal a variant-encoded value.
func Marshal(value interface{}) ([]byte, error) { //gd:var_to_bytes var_to_bytes_with_objects
	var buf []byte
	switch value := value.(type) {
	case bool:
		buf = binary.LittleEndian.AppendUint32(buf, typeBool)
		if value {
			buf = binary.LittleEndian.AppendUint32(buf, 1)
		} else {
			buf = binary.LittleEndian.AppendUint32(buf, 0)
		}
	case int8:
		var i32 int32 = int32(value)
		buf = binary.LittleEndian.AppendUint32(buf, typeInteger)
		buf = binary.LittleEndian.AppendUint32(buf, *(*uint32)(unsafe.Pointer(&i32)))
	case int16:
		var i32 int32 = int32(value)
		buf = binary.LittleEndian.AppendUint32(buf, typeInteger)
		buf = binary.LittleEndian.AppendUint32(buf, *(*uint32)(unsafe.Pointer(&i32)))
	case int32:
		var i32 int32 = int32(value)
		buf = binary.LittleEndian.AppendUint32(buf, typeInteger)
		buf = binary.LittleEndian.AppendUint32(buf, *(*uint32)(unsafe.Pointer(&i32)))
	case int64:
		buf = binary.LittleEndian.AppendUint32(buf, typeInteger&encodeFlag64)
		buf = binary.LittleEndian.AppendUint64(buf, *(*uint64)(unsafe.Pointer(&value)))
	case int:
		if reflect.TypeOf(int(0)).Size() == 4 {
			var i32 int32 = int32(value)
			buf = binary.LittleEndian.AppendUint32(buf, typeInteger)
			buf = binary.LittleEndian.AppendUint32(buf, *(*uint32)(unsafe.Pointer(&i32)))
		} else {
			buf = binary.LittleEndian.AppendUint32(buf, typeInteger&encodeFlag64)
			buf = binary.LittleEndian.AppendUint64(buf, *(*uint64)(unsafe.Pointer(&value)))
		}
	case uint8:
		buf = binary.LittleEndian.AppendUint32(buf, typeInteger)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(value))
	case uint16:
		buf = binary.LittleEndian.AppendUint32(buf, typeInteger)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(value))
	case uint32:
		buf = binary.LittleEndian.AppendUint32(buf, typeInteger&encodeFlag64)
		buf = binary.LittleEndian.AppendUint64(buf, uint64(value))
	case uint64:
		buf = binary.LittleEndian.AppendUint32(buf, typeRID)
		buf = binary.LittleEndian.AppendUint64(buf, value)
	case uint:
		if reflect.TypeOf(uint(0)).Size() == 4 {
			buf = binary.LittleEndian.AppendUint32(buf, typeInteger&encodeFlag64)
			buf = binary.LittleEndian.AppendUint64(buf, uint64(value))
		} else {
			buf = binary.LittleEndian.AppendUint32(buf, typeRID)
			buf = binary.LittleEndian.AppendUint64(buf, uint64(value))
		}
	case float32:
		buf = binary.LittleEndian.AppendUint32(buf, typeFloat)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value))
	case float64:
		buf = binary.LittleEndian.AppendUint32(buf, typeFloat)
		buf = binary.LittleEndian.AppendUint64(buf, math.Float64bits(value))
	case string:
		buf = binary.LittleEndian.AppendUint32(buf, typeString)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		buf = append(buf, value...)
		for i := len(value); i%4 != 0; i++ {
			buf = append(buf, 0)
		}
	case []byte:
		buf = binary.LittleEndian.AppendUint32(buf, typeRawArray)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		buf = append(buf, value...)
		for i := len(value); i%4 != 0; i++ {
			buf = append(buf, 0)
		}
	case Vector2.XY:
		buf = binary.LittleEndian.AppendUint32(buf, typeVector2)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Y))
	case Rect2.PositionSize:
		buf = binary.LittleEndian.AppendUint32(buf, typeRect2)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Position.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Position.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Size.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Size.Y))
	case Vector3.XYZ:
		buf = binary.LittleEndian.AppendUint32(buf, typeVector3)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Z))
	case Transform2D.OriginXY:
		buf = binary.LittleEndian.AppendUint32(buf, typeTransform2D)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.X.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.X.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Y.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Y.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Origin.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Origin.Y))
	case Plane.NormalD:
		buf = binary.LittleEndian.AppendUint32(buf, typePlane)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Normal.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Normal.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Normal.Z))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.D))
	case Quaternion.IJKX:
		buf = binary.LittleEndian.AppendUint32(buf, typeQuaternion)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.I))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.J))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.K))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.X))
	case AABB.PositionSize:
		buf = binary.LittleEndian.AppendUint32(buf, typeAABB)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Position.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Position.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Position.Z))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Size.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Size.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Size.Z))
	case Basis.XYZ:
		buf = binary.LittleEndian.AppendUint32(buf, typeBasis)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.X.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.X.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.X.Z))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Y.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Y.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Y.Z))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Z.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Z.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Z.Z))
	case Transform3D.BasisOrigin:
		buf = binary.LittleEndian.AppendUint32(buf, typeTransform3D)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Basis.X.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Basis.X.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Basis.X.Z))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Basis.Y.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Basis.Y.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Basis.Y.Z))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Basis.Z.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Basis.Z.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Basis.Z.Z))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Origin.X))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Origin.Y))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.Origin.Z))
	case Color.RGBA:
		buf = binary.LittleEndian.AppendUint32(buf, typeColor)
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.R))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.G))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.B))
		buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(value.A))
	case NodePath.String:
		buf = binary.LittleEndian.AppendUint32(buf, typeNodePath)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		buf = append(buf, value...)
	case []int32:
		buf = binary.LittleEndian.AppendUint32(buf, typeInt32Array)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		for _, v := range value {
			var i32 int32 = int32(v)
			buf = binary.LittleEndian.AppendUint32(buf, *(*uint32)(unsafe.Pointer(&i32)))
		}
	case []int64:
		buf = binary.LittleEndian.AppendUint32(buf, typeInt64Array)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		for _, v := range value {
			buf = binary.LittleEndian.AppendUint64(buf, *(*uint64)(unsafe.Pointer(&v)))
		}
	case []float32:
		buf = binary.LittleEndian.AppendUint32(buf, typeFloat32Array)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		for _, v := range value {
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(v))
		}
	case []float64:
		buf = binary.LittleEndian.AppendUint32(buf, typeFloat64Array)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		for _, v := range value {
			buf = binary.LittleEndian.AppendUint64(buf, math.Float64bits(v))
		}
	case []string:
		buf = binary.LittleEndian.AppendUint32(buf, typeStringArray)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		for _, v := range value {
			buf = binary.LittleEndian.AppendUint32(buf, uint32(len(v)))
			buf = append(buf, v...)
			for i := len(v); i%4 != 0; i++ {
				buf = append(buf, 0)
			}
		}
	case []Vector2.XY:
		buf = binary.LittleEndian.AppendUint32(buf, typeVector2Array)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		for _, v := range value {
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(v.X))
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(v.Y))
		}
	case []Vector3.XYZ:
		buf = binary.LittleEndian.AppendUint32(buf, typeVector3Array)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		for _, v := range value {
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(v.X))
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(v.Y))
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(v.Z))
		}
	case []Color.RGBA:
		buf = binary.LittleEndian.AppendUint32(buf, typeColorArray)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(value)))
		for _, v := range value {
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(v.R))
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(v.G))
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(v.B))
			buf = binary.LittleEndian.AppendUint32(buf, math.Float32bits(v.A))
		}
	default:
		return nil, fmt.Errorf("unsupported type %T", value)
	}
	return buf, nil
}
