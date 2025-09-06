package gdtype

import (
	"strings"

	"graphics.gd/internal/gdjson"
)

func EnumNameOf(class, enum_name string) (string, string) {
	if rename := gdjson.Renumeration[class+"."+enum_name]; rename != "" {
		return rename, class + "." + enum_name
	}
	rename := enum_name
	if enum_name == "MouseMode" {
		rename = "MouseModeValue"
	}
	original := enum_name
	if class != "" {
		original = class + "." + original
	}
	rename = strings.Replace(rename, ".", "", -1)
	enum_name = strings.Replace(enum_name, ".", "", -1)
	return rename, original
}

func EngineTypeAsGoType(pkg, meta string, gdType string) string {
	maybeInternal := func(name string) string {
		return "gd." + name
	}
	if strings.HasPrefix(gdType, "typedarray::") {
		gdType = strings.TrimPrefix(gdType, "typedarray::")
		meta, rest, ok := strings.Cut(gdType, ":")
		if ok {
			gdType = rest
		}
		return "Array.Contains[" + EngineTypeAsGoType(pkg, meta, gdType) + "]"
	}
	switch gdType {
	case "int", "Int":
		return "int64"
	case "float", "Float":
		return "float64"
	case "bool", "Bool":
		return "bool"
	case "StringName":
		return "String.Name"
	case "enum::GDExtension.InitializationLevel":
		return maybeInternal("GDExtensionInitializationLevel")
	case "PackedInt32Array":
		return "Packed.Array[int32]"
	case "PackedInt64Array":
		return "Packed.Array[int64]"
	case "PackedFloat32Array":
		return "Packed.Array[float32]"
	case "PackedFloat64Array":
		return "Packed.Array[float64]"
	case "PackedVector2Array":
		return "Packed.Array[Vector2.XY]"
	case "PackedVector3Array":
		return "Packed.Array[Vector3.XYZ]"
	case "PackedVector4Array":
		return "Packed.Array[Vector4.XYZW]"
	case "PackedColorArray":
		return "Packed.Array[Color.RGBA]"
	case "PackedStringArray":
		return "Packed.Strings"
	case "PackedByteArray":
		return "Packed.Bytes"
	case "Vector2":
		return "Vector2.XY"
	case "Vector2i":
		return "Vector2i.XY"
	case "Rect2":
		return "Rect2.PositionSize"
	case "Rect2i":
		return "Rect2i.PositionSize"
	case "Vector3":
		return "Vector3.XYZ"
	case "Vector3i":
		return "Vector3i.XYZ"
	case "Transform2D":
		return "Transform2D.OriginXY"
	case "Vector4":
		return "Vector4.XYZW"
	case "Vector4i":
		return "Vector4i.XYZW"
	case "Plane":
		return "Plane.NormalD"
	case "Quaternion":
		return "Quaternion.IJKX"
	case "AABB":
		return "AABB.PositionSize"
	case "Basis":
		return "Basis.XYZ"
	case "Transform3D":
		return "Transform3D.BasisOrigin"
	case "Projection":
		return "Projection.XYZW"
	case "Color":
		return "Color.RGBA"
	case "RID":
		return "RID.Any"
	case "NodePath":
		return "Path.ToNode"
	case "Signal":
		return "Signal.Any"
	case "Array":
		return "Array.Any"
	case "Dictionary":
		return "Dictionary.Any"
	case "String":
		return "String.Readable"
	case "Callable":
		return "Callable.Function"
	case "Variant":
		return "variant.Any"
	case "enum::Variant.Type":
		return "variant.Type"
	// strange C++ cases
	case "enum::Error":
		return "Error.Code"
	case "const uint8_t **":
		return "gdextension.Pointer"
	case "const void*", "const uint8_t*", "const uint8_t *":
		return "gdextension.Pointer"
	case "float*":
		return "*float64"
	case "int32_t*":
		return "*int32"
	case "void*", "uint8_t*":
		return "gdextension.Pointer"
	case "Object":
		return "[1]gd.Object"
	case "RefCounted":
		return "gd." + gdType
	default:
		gdType = strings.TrimPrefix(gdType, "const")
		if strings.HasSuffix(gdType, "*") {
			gdType = strings.TrimPrefix(gdType, pkg)
			return "*" + gdType[:len(gdType)-1]
		}
		if strings.HasPrefix(gdType, "enum::") || strings.HasPrefix(gdType, "bitfield::") {
			gdType = strings.TrimPrefix(gdType, "enum::")
			gdType = strings.TrimPrefix(gdType, "bitfield::")
			if rename := gdjson.Renumeration[gdType]; rename != "" {
				gdType = rename
			}
			host, sub, hasHost := strings.Cut(gdType, ".")
			if sub == "MouseMode" {
				sub = "MouseModeValue"
			}
			if host == "RenderingDevice" {
				host = "Rendering"
			}
			if hasHost {
				if host == pkg {
					return sub
				}
				return host + "." + sub
			} else {
				return gdType
			}
		}
		gdType = strings.Replace(gdType, ".", "", -1)
		gdType = strings.Replace(gdType, ".", "", -1)
		gdType = strings.Replace(gdType, ".", "", -1)
		if gdType == "Object" {
			return "[1]gd.Object"
		}
		if class, ok := ClassDB[gdType]; ok {
			return "[1]gdclass." + class.Name
		}
		if Name(gdType).InCore() {
			return maybeInternal(gdType)
		}
		if gdType != "" {
			return "[1]gdclass." + gdType
		}
		return gdType
	}
}
