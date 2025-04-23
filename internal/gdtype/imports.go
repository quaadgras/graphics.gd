package gdtype

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"

	"graphics.gd/internal/gdjson"
)

func ImportsForClass(class gdjson.Class) iter.Seq[string] {
	return func(yield func(string) bool) {
		var imports = map[string]bool{
			"graphics.gd/variant/Object":     true,
			"graphics.gd/variant/Float":      true,
			"graphics.gd/variant/RefCounted": true,
			"graphics.gd/variant/Array":      true,
			"graphics.gd/variant/Callable":   true,
			"graphics.gd/variant/Dictionary": true,
			"graphics.gd/variant/RID":        true,
			"graphics.gd/variant/String":     true,
			"graphics.gd/variant/Path":       true,
			"graphics.gd/variant/Packed":     true,
			"graphics.gd/variant/Error":      true,
		}
		if class.Name == "TextEdit" {
			imports["graphics.gd/variant/Rect2"] = true
		}
		if class.Inherits != "" {
			super := ClassDB[class.Inherits]
			for super.Name != "" && super.Name != "Object" && super.Name != "RefCounted" && !ClassDB[super.Name].IsSingleton {
				path := fmt.Sprintf("graphics.gd/classdb/%s", super.Name)
				imports[path] = true
				super = ClassDB[super.Inherits]
			}
		}
		if class.Name == "IP" {
			imports["net/netip"] = true
		}
		for _, method := range class.Methods {
			if _, ok := gdjson.Relocations[class.Name+"."+method.Name]; ok {
				continue
			}
			for _, arg := range method.Arguments {
				for pkg := range importsForEngineType(class, class.Name+"."+method.Name+"."+arg.Name, arg.Type) {
					if !imports[pkg] {
						imports[pkg] = true
					}
				}
			}
			for pkg := range importsForEngineType(class, "", method.ReturnValue.Type) {
				imports[pkg] = true
			}
		}
		for _, peer_method := range gdjson.RelocationsReverse[class.Name] {
			peer, method_name, _ := strings.Cut(peer_method, ".")
			imports["graphics.gd/classdb/"+peer] = true
			peerClass := ClassDB[peer]
			method := gdjson.Method{}
			for _, peerMethod := range peerClass.Methods {
				if peerMethod.Name == method_name {
					method = peerMethod
					break
				}
			}
			for _, arg := range method.Arguments {
				for pkg := range importsForEngineType(class, class.Name+"."+method.Name+"."+arg.Name, arg.Type) {
					if !imports[pkg] {
						imports[pkg] = true
					}
				}
			}
			for pkg := range importsForEngineType(class, "", method.ReturnValue.Type) {
				imports[pkg] = true
			}
		}
		for _, signal := range class.Signals {
			for _, arg := range signal.Arguments {
				for pkg := range importsForEngineType(class, class.Name+"."+signal.Name+"."+arg.Name, arg.Type) {
					imports[pkg] = true
				}
			}
		}
		for _, pkg := range slices.Sorted(maps.Keys(imports)) {
			if !yield(pkg) {
				return
			}
		}
	}
}

func importsForEngineType(class gdjson.Class, identifier, s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		if strings.HasPrefix(s, "typedarray::") {
			s = strings.TrimPrefix(s, "typedarray::")
			for pkg := range importsForEngineType(class, identifier, s) {
				if !yield(pkg) {
					return
				}
			}
			return
		}
		if _, ok := ClassDB[s]; ok && s != "Object" && s != class.Name {
			if !yield("graphics.gd/classdb/" + s) {
				return
			}
		}
		switch s {
		case "Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i", "Transform2D", "Vector4", "Vector4i",
			"Plane", "Quaternion", "AABB", "Basis", "Transform3D", "Projection", "Color":
			yield("graphics.gd/variant/" + s)
		case "PackedVector2Array":
			yield("graphics.gd/variant/Vector2")
		case "PackedVector3Array":
			yield("graphics.gd/variant/Vector3")
		case "PackedVector4Array":
			yield("graphics.gd/variant/Vector4")
		case "PackedColorArray":
			yield("graphics.gd/variant/Color")
		case "Callable":
			details := gdjson.Callables[identifier]
			if len(details) == 0 {
				return
			}
			for _, detail := range details {
				if detail == "void" {
					continue
				}
				detail, _, _ = strings.Cut(detail, " ")
				for pkg := range importsForEngineType(class, "", detail) {
					if !yield(pkg) {
						return
					}
				}
			}
		}
	}
}
