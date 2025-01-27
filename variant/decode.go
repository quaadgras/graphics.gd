package variant

import (
	"fmt"
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

type decodable interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64 |
		Vector2.XY | Rect2.PositionSize | Vector3.XYZ | Transform2D.OriginXY | Plane.NormalD | Quaternion.IJKX |
		AABB.PositionSize | Basis.XYZ | Transform3D.BasisOrigin | Color.RGBA
}

func decode[T decodable](data []byte) (T, error) {
	if unsafe.Sizeof([1]T{}[0]) > uintptr(len(data)) {
		return [1]T{}[0], fmt.Errorf("decode: data too short for type %T", [1]T{}[0])
	}
	return *(*T)(unsafe.Pointer(&data[0])), nil
}

func decodePacked[T decodable](data []byte) ([]T, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("decodePacked: data too short for length")
	}
	length, err := decode[uint32](data)
	if err != nil {
		return nil, err
	}
	if uint32(len(data)) < 4+length*uint32(unsafe.Sizeof([1]T{}[0])) {
		return nil, fmt.Errorf("decodePacked: data too short for packed array of type %T", [1]T{}[0])
	}
	return *(*[]T)(unsafe.Pointer(&data[4])), nil
}

// DecodeString decodes a length-prefixed string from the bytes starting at offset. Fails if the byte count is
// insufficient. Returns an empty string if a valid string can't be decoded.
func decodeString(data []byte) (string, error) {
	length, err := decode[uint32](data)
	if err != nil {
		return "", err
	}
	if length > uint32(len(data)) {
		return "", fmt.Errorf("decodeString: data too short for string of length %d", length)
	}
	return string(data[4 : 4+length]), nil
}

const (
	typeNull = iota
	typeBool
	typeInteger
	typeFloat
	typeString
	typeVector2
	typeRect2
	typeVector3
	typeTransform2D
	typePlane
	typeQuaternion
	typeAABB
	typeBasis
	typeTransform3D
	typeColor
	typeNodePath
	typeRID
	typeObject
	typeDictionary
	typeArray
	typeRawArray
	typeInt32Array
	typeInt64Array
	typeFloat32Array
	typeFloat64Array
	typeStringArray
	typeVector2Array
	typeVector3Array
	typeColorArray
)

const (
	encodeFlag64 = 1 << 0
)

// UnmarshalAny a variant-encoded value from the byte slice.
func UnmarshalAny(data []byte) (any, error) { //gd:bytes_to_var bytes_to_var_with_objects
	header, err := decode[uint32](data)
	if err != nil {
		return nil, err
	}
	vtype := header & 0xFF
	flags := header >> 16
	data = data[4:]
	switch vtype {
	case typeNull:
		return nil, nil
	case typeBool:
		b, err := decode[uint32](data)
		if err != nil {
			return nil, err
		}
		return b != 0, nil
	case typeInteger:
		if flags&encodeFlag64 != 0 {
			return decode[int64](data)
		}
		return decode[int32](data)
	case typeFloat:
		if flags&encodeFlag64 != 0 {
			return decode[float64](data)
		}
		return decode[float32](data)
	case typeString:
		return decodeString(data)
	case typeVector2:
		return decode[Vector2.XY](data)
	case typeRect2:
		return decode[Rect2.PositionSize](data)
	case typeVector3:
		return decode[Vector3.XYZ](data)
	case typeTransform2D:
		return decode[Transform2D.OriginXY](data)
	case typePlane:
		return decode[Plane.NormalD](data)
	case typeQuaternion:
		return decode[Quaternion.IJKX](data)
	case typeAABB:
		return decode[AABB.PositionSize](data)
	case typeBasis:
		return decode[Basis.XYZ](data)
	case typeTransform3D:
		return decode[Transform3D.BasisOrigin](data)
	case typeColor:
		return decode[Color.RGBA](data)
	case typeNodePath:
		length, err := decode[uint32](data)
		if err != nil {
			return nil, err
		}
		data = data[4:]
		if length&0x80000000 != 0 {
			count := length & 0x7FFFFFFF
			flags, err := decode[uint32](data)
			if err != nil {
				return nil, err
			}
			absolute := flags&1 != 0
			var path NodePath.String
			if absolute {
				path = "/"
			}
			for i := uintptr(0); i < uintptr(count); i++ {
				s, err := decodeString(data)
				if err != nil {
					return nil, err
				}
				path += NodePath.String(s)
				data = data[(len(s)+3)&^3:] // must be multiple of 4 bytes
			}
			return path, nil
		}
		s, err := decodeString(data)
		if err != nil {
			return nil, err
		}
		return NodePath.String(s), nil
	case typeRID:
		return decode[uint64](data)
	case typeObject:
		return nil, nil
	case typeDictionary:
		length, err := decode[uint32](data)
		if err != nil {
			return nil, err
		}
		var dictionary = make(map[any]any)
		for i := uintptr(0); i < uintptr(length)&0x7FFFFFFF; i++ {
			keySize, err := UnmarshalSize(data)
			if err != nil {
				return nil, err
			}
			valSize, err := UnmarshalSize(data[keySize:])
			if err != nil {
				return nil, err
			}
			key, err := UnmarshalAny(data)
			if err != nil {
				return nil, err
			}
			value, err := UnmarshalAny(data[keySize:])
			if err != nil {
				return nil, err
			}
			dictionary[key] = value
			data = data[keySize+valSize:]
		}
		return dictionary, nil
	case typeArray:
		length, err := decode[uint32](data)
		if err != nil {
			return nil, err
		}
		var result = make([]any, length&0x7FFFFFFF)
		for i := range result {
			size, err := UnmarshalSize(data)
			if err != nil {
				return nil, err
			}
			result[i], err = UnmarshalAny(data)
			if err != nil {
				return nil, err
			}
			data = data[size:]
		}
		return result, nil
	case typeRawArray:
		length, err := decode[uint32](data)
		if err != nil {
			return nil, err
		}
		return data[4 : 4+uintptr(length)], nil
	case typeInt32Array:
		return decodePacked[int32](data)
	case typeInt64Array:
		return decodePacked[int64](data)
	case typeFloat32Array:
		return decodePacked[float32](data)
	case typeFloat64Array:
		return decodePacked[float64](data)
	case typeStringArray:
		length, err := decode[uint32](data)
		if err != nil {
			return nil, err
		}
		result := make([]string, length)
		for i := uint32(0); i < length; i++ {
			result[i], err = decodeString(data)
			if err != nil {
				return nil, err
			}
			data = data[(len(result[i])+3)&^3:] // must be multiple of 4 bytes
		}
		return result, nil
	case typeVector2Array:
		return decodePacked[Vector2.XY](data)
	case typeVector3Array:
		return decodePacked[Vector3.XYZ](data)
	case typeColorArray:
		return decodePacked[Color.RGBA](data)
	}
	return nil, fmt.Errorf("unsupported variant type %d", vtype)
}

// Decodes a size of a Variant from the bytes starting at offset. Requires at least 4 bytes of data
// starting at the offset, otherwise fails.
func UnmarshalSize(data []byte) (uintptr, error) {
	header, err := decode[uint32](data)
	if err != nil {
		return 0, err
	}
	vtype := header & 0xFF
	flags := header >> 16
	data = data[4:]
	switch vtype {
	case typeNull:
		return 4, nil
	case typeBool:
		return 4 + 4, nil
	case typeInteger, typeFloat:
		if flags&encodeFlag64 != 0 {
			return 4 + 8, nil
		}
		return 4 + 4, nil
	case typeString:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		return 4 + uintptr(length), nil
	case typeVector2:
		return 4 + unsafe.Sizeof(Vector2.XY{}), nil
	case typeRect2:
		return 4 + unsafe.Sizeof(Rect2.PositionSize{}), nil
	case typeVector3:
		return 4 + unsafe.Sizeof(Vector3.XYZ{}), nil
	case typeTransform2D:
		return 4 + unsafe.Sizeof(Transform2D.OriginXY{}), nil
	case typePlane:
		return 4 + unsafe.Sizeof(Plane.NormalD{}), nil
	case typeQuaternion:
		return 4 + unsafe.Sizeof(Quaternion.IJKX{}), nil
	case typeAABB:
		return 4 + unsafe.Sizeof(AABB.PositionSize{}), nil
	case typeBasis:
		return 4 + unsafe.Sizeof(Basis.XYZ{}), nil
	case typeTransform3D:
		return 4 + unsafe.Sizeof(Transform3D.BasisOrigin{}), nil
	case typeColor:
		return 4 + unsafe.Sizeof(Color.RGBA{}), nil
	case typeNodePath:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		if length&0x80000000 != 0 {
			count := length & 0x7FFFFFFF
			var size uintptr = 4
			for i := uintptr(0); i < uintptr(count); i++ {
				length, err := decode[uint32](data[size:])
				if err != nil {
					return 0, err
				}
				size += 4 + uintptr(length)
				data = data[size:]
			}
			return size, nil
		}
		return 4 + uintptr(length), nil
	case typeRID:
		return 4 + 8, nil
	case typeObject:
		return 4, nil
	case typeDictionary:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		size := uintptr(4)
		for i := uint32(0); i < length; i++ {
			s, err := UnmarshalSize(data[size:])
			if err != nil {
				return 0, err
			}
			size += s
			data = data[size:]
		}
		return size, nil
	case typeArray:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		size := uintptr(4)
		for i := uint32(0); i < length; i++ {
			s, err := UnmarshalSize(data[size:])
			if err != nil {
				return 0, err
			}
			size += s
			data = data[size:]
		}
		return size, nil
	case typeInt32Array:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		return 4 + 4 + uintptr(length)*unsafe.Sizeof(int32(0)), nil
	case typeInt64Array:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		return 4 + 4 + uintptr(length)*unsafe.Sizeof(int64(0)), nil
	case typeFloat32Array:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		return 4 + 4 + uintptr(length)*unsafe.Sizeof(float32(0)), nil
	case typeFloat64Array:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		return 4 + 4 + uintptr(length)*unsafe.Sizeof(float64(0)), nil
	case typeStringArray:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		size := uintptr(4)
		for i := uint32(0); i < length; i++ {
			length, err := decode[uint32](data[size:])
			if err != nil {
				return 0, err
			}
			size += 4 + uintptr(length)
			data = data[size:]
		}
		return size, nil
	case typeVector2Array:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		return 4 + 4 + uintptr(length)*unsafe.Sizeof(Vector2.XY{}), nil
	case typeVector3Array:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		return 4 + 4 + uintptr(length)*unsafe.Sizeof(Vector3.XYZ{}), nil
	case typeColorArray:
		length, err := decode[uint32](data)
		if err != nil {
			return 0, err
		}
		return 4 + 4 + uintptr(length)*unsafe.Sizeof(Color.RGBA{}), nil
	}
	return 0, fmt.Errorf("unknown type %d", vtype)
}
