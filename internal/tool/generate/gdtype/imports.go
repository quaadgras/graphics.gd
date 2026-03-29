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
		if class.Name == "CSGShape3D" {
			imports["graphics.gd/classdb/Mesh"] = true
			imports["graphics.gd/variant/Transform3D"] = true
		}
		if class.Name == "OpenXRInterface" {
			imports["graphics.gd/classdb/OpenXRActionSet"] = true
		}
		if class.Name == "MeshLibrary" {
			imports["graphics.gd/classdb/Shape3D"] = true
		}
		if class.Name == "TextEdit" {
			imports["graphics.gd/variant/Rect2"] = true
		}
		if class.Name == "ResourceUID" {
			imports["graphics.gd/classdb/Resource"] = true
		}
		if class.Name == "AudioStreamPlaybackInteractive" {
			imports["graphics.gd/classdb/AudioStreamInteractive"] = true
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
		if after, ok := strings.CutPrefix(s, "typedarray::"); ok {
			s = after
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
		if strings.HasPrefix(s, "enum::") || strings.HasPrefix(s, "bitfield::") {
			s = strings.TrimPrefix(s, "enum::")
			s = strings.TrimPrefix(s, "bitfield::")
			if rename := gdjson.Renumeration[s]; rename != "" {
				s = rename
			}
			host, _, hasHost := strings.Cut(s, ".")
			if hasHost {
				if host == "RenderingDevice" {
					host = "Rendering"
				}
				if class.Name != host {
					if dependency, ok := ClassDB[host]; ok && !dependency.IsEnum {
						if !yield("graphics.gd/classdb/" + host) {
							return
						}
					}
				}
			}
			s = host
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
		// Check Addressables/Sliceables for any pointer-typed param (AudioFrame*, void*, float*, etc.)
		if identifier != "" {
			if s, ok := gdjson.Sliceables[identifier]; ok {
				if !yield("graphics.gd/internal/gdmemory") {
					return
				}
				switch s.Elem {
				case "byte", "int32", "int64", "float32", "float64",
					"Vector2.XY", "Vector3.XYZ", "Vector4.XYZW", "Color.RGBA":
					if !yield("graphics.gd/variant/Packed") {
						return
					}
				}
			}
			if mapped, ok := gdjson.Addressables[identifier]; ok {
				if strings.HasPrefix(mapped, "Engine.Pointer[") {
					if !yield("graphics.gd/classdb/Engine") {
						return
					}
					if !yield("graphics.gd/internal/gdmemory") {
						return
					}
					// Extract the inner type and check if it needs a classdb import.
					inner := strings.TrimPrefix(mapped, "Engine.Pointer[")
					inner = strings.TrimSuffix(inner, "]")
					if pkg, _, ok := strings.Cut(inner, "."); ok && pkg != class.Name {
						if _, exists := ClassDB[pkg]; exists || pkg == "OpenXR" {
							if !yield("graphics.gd/classdb/" + pkg) {
								return
							}
						}
					}
				}
			}
		}
	}
}
