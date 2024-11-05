package main

import (
	"fmt"
	"strings"

	"grow.graphics/gd/internal/gdjson"
)

type ClassDB map[string]gdjson.Class

func (db ClassDB) nameOf(pkg, original string) string {
	class := db[original]
	if pkg == class.Package {
		return class.Name
	}
	if class.Package == "internal" {
		return "gd." + class.Name
	}
	return class.Package + "." + class.Name
}

func (db ClassDB) genArgs(pkg string, method gdjson.Method) string {
	var args []string
	for _, arg := range method.Arguments {
		if arg.Name == "type" {
			arg.Name = "typ"
		}
		if arg.Name == "func" {
			arg.Name = "fn"
		}
		if arg.Name == "default" {
			arg.Name = "def"
		}
		if arg.Name == "interface" {
			arg.Name = "iface"
		}
		if arg.Name == "range" {
			arg.Name = "rng"
		}
		if arg.Name == "var" {
			arg.Name = "variable"
		}
		if arg.Name == "map" {
			arg.Name = "mapping"
		}
		if arg.Name == "internal" {
			arg.Name = "internal_"
		}
		args = append(args, fmt.Sprintf("%v %v", arg.Name, db.convertType(pkg, arg.Meta, arg.Type)))
	}
	return strings.Join(args, ", ")
}

func pascal(name string) string {
	words := strings.Split(name, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func fixReserved(name string) string {
	switch name {
	case "bool":
		return "b"
	case "type":
		return "atype"
	case "range":
		return "arange"
	case "default":
		return "def"
	case "class":
		return "class_"
	case "func":
		return "fn"
	case "frame":
		return "frame_"
	case "interface":
		return "intf"
	case "internal":
		return "internal_"
	case "map":
		return "mapping"
	case "var":
		return "v"
	case "object":
		return "obj"
	case "string":
		return "s"
	case "RID":
		return "rid"
	default:
		return name
	}
}

func inCore(s string) bool {
	switch s {
	case "Object", "RefCounted":
		return true
	default:
		return false
	}
}

func isBuiltin(s string) bool {
	switch s {
	case "String", "StringName", "PackedStringArray", "PackedInt32Array", "PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedVector2Array", "PackedVector3Array", "PackedVector4Array", "PackedColorArray", "PackedByteArray",
		"Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i", "Transform2D", "Vector4", "Vector4i",
		"Plane", "Quaternion", "AABB", "Basis", "Transform3D", "Projection", "Color", "NodePath", "RID",
		"Callable", "Signal", "Dictionary", "Array":
		return true
	default:
		return false
	}
}

func importsVariant(class gdjson.Class, identifier, s string) string {
	switch s {
	case "Float", "float":
		return "grow.graphics/gd/variant/Float"
	case "Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i", "Transform2D", "Vector4", "Vector4i",
		"Plane", "Quaternion", "AABB", "Basis", "Transform3D", "Projection", "Color":
		return "grow.graphics/gd/variant/" + s
	case "NodePath":
		return "grow.graphics/gd/variant/Path"
	case "RID":
		if class.Name == "Resource" {
			return ""
		}
		return "grow.graphics/gd/objects/Resource"
	case "Array", "Dictionary", "Signal":
		return "grow.graphics/gd/variant/" + s
	case "PackedVector2Array":
		return "grow.graphics/gd/variant/Vector2"
	case "PackedVector3Array":
		return "grow.graphics/gd/variant/Vector3"
	case "PackedVector4Array":
		return "grow.graphics/gd/variant/Vector4"
	case "PackedColorArray":
		return "grow.graphics/gd/variant/Color"
	case "Callable":
		details := gdjson.Callables[identifier]
		if len(details) == 0 {
			return "grow.graphics/gd/variant/Callable"
		}
		for _, detail := range details {
			if detail == "void" {
				continue
			}
			detail, _, _ = strings.Cut(detail, " ")
			return importsVariant(class, "", detail)
		}
		return ""
	default:
		return ""
	}
}

func (classDB ClassDB) convertType(pkg, meta string, gdType string) string {
	maybeInternal := func(name string) string {
		return "gd." + name
	}
	if strings.HasPrefix(gdType, "typedarray::") {
		gdType = strings.TrimPrefix(gdType, "typedarray::")
		/*meta, rest, ok := strings.Cut(gdType, ":")
		if ok {
			gdType = rest
		}*/
		return maybeInternal("Array") // Of[" + classDB.convertType(pkg, meta, gdType) + "]
	}
	switch gdType {
	case "int", "Int":
		return maybeInternal("Int")
	case "float", "Float":
		return maybeInternal("Float")
	case "bool", "Bool":
		return "bool"
	case "String":
		return maybeInternal("String")
	case "StringName":
		return maybeInternal("StringName")
	case "enum::GDExtension.InitializationLevel":
		return maybeInternal("GDExtensionInitializationLevel")
	case "enum::GDExtensionManager.LoadStatus":
		return "classdb.GDExtensionManagerLoadStatus"
	case "PackedStringArray", "PackedInt32Array", "PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedVector2Array", "PackedVector3Array", "PackedVector4Array", "PackedColorArray", "PackedByteArray",
		"Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i", "Transform2D", "Vector4", "Vector4i",
		"Plane", "Quaternion", "AABB", "Basis", "Transform3D", "Projection", "Color", "NodePath", "RID",
		"Callable", "Signal", "Dictionary", "Array":
		return maybeInternal(gdType)
	case "Variant":
		return maybeInternal("Variant")
	case "enum::Variant.Type":
		return maybeInternal("VariantType")

	// strange C++ cases
	case "enum::Error":
		return "error"
	case "const uint8_t **":
		return "unsafe.Pointer"
	case "const void*", "const uint8_t*", "const uint8_t *":
		return "unsafe.Pointer"
	case "float*":
		return "*float64"
	case "int32_t*":
		return "*int32"
	case "void*", "uint8_t*":
		return "unsafe.Pointer"
	case "Object", "RefCounted":
		return "gd." + gdType
	default:
		gdType = strings.TrimPrefix(gdType, "const")

		if strings.HasSuffix(gdType, "*") {
			return "*classdb." + gdType[:len(gdType)-1]
		}

		if strings.HasPrefix(gdType, "enum::") || strings.HasPrefix(gdType, "bitfield::") {
			gdType = strings.TrimPrefix(gdType, "enum::")
			gdType = strings.TrimPrefix(gdType, "bitfield::")
			host, sub, hasHost := strings.Cut(gdType, ".")
			if isBuiltin(host) {
				return "gd." + host + sub
			}
			if hasHost {
				return "classdb." + host + sub
			} else {
				return gdType
			}
		}

		gdType = strings.Replace(gdType, ".", "", -1)
		gdType = strings.Replace(gdType, ".", "", -1)

		gdType = strings.Replace(gdType, ".", "", -1)

		if gdType == "Object" {
			return "gd.Object"
		}

		if class, ok := classDB[gdType]; ok {
			return "objects." + class.Name
		}

		if inCore(gdType) {
			return maybeInternal(gdType)
		}

		if gdType != "" {
			return "objects." + gdType
		}
		return gdType
	}
}

func (classDB ClassDB) convertTypeSimple(class gdjson.Class, lookup, meta string, gdType string) string {
	if strings.HasPrefix(gdType, "typedarray::") {
		gdType = strings.TrimPrefix(gdType, "typedarray::")
		return "gd.Array" // Of[+ classDB.convertType("", meta, gdType) + "]"
	}
	switch gdType {
	case "int", "Int":
		return "int"
	case "float", "Float":
		return "Float.X"
	case "bool", "Bool":
		return "bool"
	case "String", "StringName":
		return "string"
	case "PackedByteArray":
		return "[]byte"
	case "PackedStringArray":
		return "[]string"
	case "PackedInt32Array":
		return "[]int32"
	case "PackedInt64Array":
		return "[]int64"
	case "PackedFloat32Array":
		return "[]float32"
	case "PackedFloat64Array":
		return "[]float64"
	case "PackedVector2Array":
		return "[]Vector2.XY"
	case "PackedVector3Array":
		return "[]Vector3.XYZ"
	case "PackedVector4Array":
		return "[]Vector4.XYZW"
	case "PackedColorArray":
		return "[]Color.RGBA"
	case "enum::Error":
		return "error"
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
	case "NodePath":
		return "Path.String"
	case "RID":
		if class.Name == "Resource" {
			return "ID"
		}
		return "Resource.ID"
	case "ObjectID":
		return "objects.ID"
	case "Signal":
		return "Signal.Any"
	case "Dictionary":
		return "Dictionary.Any"
	case "Array":
		return "Array.Any"
	case "Variant":
		return "any"
	case "Callable":
		details, ok := gdjson.Callables[lookup]
		if !ok || len(details) == 0 {
			return "Callable.Any"
		}
		var ftype string = "func("
		for i, arg := range details[1:] {
			if i > 0 {
				ftype += ", "
			}
			atype, name, ok := strings.Cut(arg, " ")
			if ok {
				ftype += name + " " + classDB.convertTypeSimple(class, "", "", atype)
			} else {
				ftype += classDB.convertTypeSimple(class, "", "", arg)
			}
		}
		ftype += ")"
		if details[0] != "void" {
			ftype += " " + classDB.convertTypeSimple(class, "", "", details[0])
		}
		return ftype
	default:
		return classDB.convertType("", meta, gdType)
	}
}

func mapOperator(name string) string {
	switch name {
	case "==":
		return "Equals"
	case "!=":
		return "NotEqual"
	case "<":
		return "Less"
	case "<=":
		return "LessEqual"
	case ">":
		return "Greater"
	case ">=":
		return "GreaterEqual"
	case "+":
		return "Add"
	case "-":
		return "Subtract"
	case "*":
		return "Multiply"
	case "/":
		return "Divide"
	case "unary-":
		return "Negate"
	case "%":
		return "Module"
	case "**":
		return "Power"
	case "<<":
		return "ShiftLeft"
	case ">>":
		return "ShiftRight"
	case "&":
		return "BitAnd"
	case "|":
		return "BitOr"
	case "^":
		return "BitXor"
	case "~":
		return "BitNegate"
	case "and":
		return "And"
	case "or":
		return "Or"
	case "xor":
		return "Xor"
	case "not":
		return "Not"
	case "in":
		return "In"
	default:
		panic("unknown operator: " + name)
	}
}

func convertName(fnName string) string {
	if fnName == "seek" {
		return "SeekTo"
	}
	if fnName == "type_string" {
		return "TypeToString"
	}

	fnName = strings.ToLower(fnName)

	joins := []string{}
	split := strings.Split(fnName, "_")
	for _, word := range split {
		joins = append(joins, strings.Title(word))
	}
	/*if joins[0] == "Get" {
		backup := joins
		joins = joins[1:]

		if len(joins) == 0 {
			joins = backup
		} else {
			if _, err := strconv.Atoi(joins[0]); err == nil {
				joins = backup
			}
		}
	}*/
	return strings.Join(joins, "")
}

func (db ClassDB) isPointer(t string) (string, bool) {
	t = strings.TrimPrefix(t, "[1]")
	t = strings.TrimPrefix(t, "gd.")
	t = strings.TrimPrefix(t, "classdb.")
	t = strings.TrimPrefix(t, "objects.")
	if strings.HasPrefix(t, "ArrayOf") {
		return "[1]uintptr", true
	}
	switch t {
	case "String", "StringName", "NodePath",
		"Dictionary", "Array":
		return "[1]uintptr", true
	case "Callable", "Signal",
		"PackedByteArray", "PackedInt32Array",
		"PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedStringArray",
		"PackedVector2Array", "PackedVector3Array",
		"PackedVector4Array",
		"PackedColorArray":
		return "[2]uintptr", true
	case "Variant":
		return "[3]uintptr", true
	default:
		if entry, ok := db[t]; ok && !entry.IsEnum {
			return "[1]uintptr", true
		}
		return "", false
	}
}
