package gdjson

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"sync"

	"graphics.gd/internal/bbparser"
)

var parser = bbparser.New()
var parser_once sync.Once
var parser_block int
var parser_class Class
var parser_links = make(map[string]string)
var parser_sample string

func DocsToGoDoc(docs string, classdb map[string]Class, className, codeblock string) string {
	docs = strings.ReplaceAll(docs, "*/", "")

	parser_class = classdb[className]
	parser_block = 0
	parser_sample = codeblock
	clear(parser_links)
	parser_once.Do(func() {
		code := func(tag bbparser.Tag, body string) string {
			parser_block++
			if parser_sample == "" {
				return body
			}
			name := fmt.Sprintf("%s.go", parser_sample)
			if parser_block > 1 {
				name = fmt.Sprintf("%s%d.go", parser_sample, parser_block)
			}
			trans, err := os.ReadFile("./internal/gddocs/" + name)
			s := string(trans)
			_, s, _ = strings.Cut(s, "package main\n")
			if strings.Contains(s, "func "+parser_sample+"() {") {
				_, s, _ = strings.Cut(s, "func "+parser_sample+"() {")
				var hasEquals bool
				s, _, hasEquals = strings.Cut(s, "	_ = ")
				if !hasEquals {
					s = strings.TrimSuffix(s, "}\n")
				}
				return s
			}
			if err == nil {
				return "\tpackage main\n" + strings.ReplaceAll(string(s), "\n", "\n\t")
			}
			return ""
		}

		parser.AddTag("codeblock", code)
		parser.AddTag("codeblocks", code)
		parser.AddTag("param", func(tag bbparser.Tag, body string) string {
			for key := range tag.Attributes {
				return fmt.Sprintf("'%s'", key)
			}
			return ""
		})
		parser.AddTag("method", func(tag bbparser.Tag, body string) string {
			for key := range tag.Attributes {
				if strings.Contains(key, ".") {
					class, method, _ := strings.Cut(key, ".")
					if classdb[class].IsSingleton {
						return fmt.Sprintf("[graphics.gd/classdb/%s.%s]", class, ConvertName(method))
					}
					for _, method := range classdb[class].Methods {
						if method.Name == key {
							if method.IsStatic {
								return fmt.Sprintf("[graphics.gd/classdb/%s.%s]", class, ConvertName(method.Name))
							}
							if method.IsVirtual {
								return fmt.Sprintf("[graphics.gd/classdb/%s.Interface.%s]", class, ConvertName(method.Name))
							}
							break
						}
					}
					return fmt.Sprintf("[graphics.gd/classdb/%s.Instance.%s]", class, ConvertName(method))
				}
				if parser_class.IsSingleton {
					return fmt.Sprintf("[%s]", ConvertName(key))
				}
				class_info := classdb[parser_class.Name]
				for _, method := range class_info.Methods {
					if method.Name == key {
						if method.IsStatic {
							return fmt.Sprintf("[%s]", ConvertName(key))
						}
						if method.IsVirtual {
							return fmt.Sprintf("[Interface.%s]", ConvertName(key))
						}
						break
					}
				}
				return fmt.Sprintf("[Instance.%s]", ConvertName(key))
			}
			return ""
		})
		parser.AddTag("constant", func(tag bbparser.Tag, body string) string {
			for key := range tag.Attributes {
				return fmt.Sprintf("[%s]", ConvertName(key))
			}
			return ""
		})
		parser.AddTag("i", func(t bbparser.Tag, s string) string { return s })
		parser.AddTag("b", func(tag bbparser.Tag, body string) string { return body })
		parser.AddTag("code", func(t bbparser.Tag, s string) string { return s })
		parser.AddTag("kbd", func(t bbparser.Tag, s string) string { return s })
		parser.AddTag("annotation", func(t bbparser.Tag, s string) string { return s })
		parser.AddTag("@GlobalScope", func(t bbparser.Tag, s string) string { return "standard" })

		parser.AddTag("lb", func(t bbparser.Tag, s string) string { return "[" })
		parser.AddTag("rb", func(t bbparser.Tag, s string) string { return "]" })
		parser.AddTag("member", func(tag bbparser.Tag, body string) string {
			for key := range tag.Attributes {
				if strings.Contains(key, ".") {
					class, member, _ := strings.Cut(key, ".")
					if strings.Contains(member, "/") {
						if classdb[class].IsSingleton {
							return fmt.Sprintf("[graphics.gd/classdb/%s] %q", class, member)
						}
						return fmt.Sprintf("[graphics.gd/classdb/%s.Instance] %q", class, member)
					}
					member = ConvertName(member)
					if classdb[class].IsSingleton {
						return fmt.Sprintf("[graphics.gd/classdb/%s.%s]", class, member)
					}
					return fmt.Sprintf("[graphics.gd/classdb/%s.Instance.%s]", class, member)
				}
				member := key
				if strings.Contains(member, "/") {
					return fmt.Sprintf("%q", member)
				}
				member = ConvertName(member)
				if parser_class.IsSingleton {
					return fmt.Sprintf("[%s]", member)
				}
				return fmt.Sprintf("[Instance.%s]", member)
			}
			return ""
		})
		parser.AddTag("signal", func(tag bbparser.Tag, body string) string {
			for key := range tag.Attributes {
				if parser_class.IsSingleton {
					return fmt.Sprintf("[On%s]", ConvertName(key))
				}
				return fmt.Sprintf("[Instance.On%s]", ConvertName(key))
			}
			return ""
		})

		parser.AddTag("enum", func(tag bbparser.Tag, body string) string {
			for key := range tag.Attributes {
				return fmt.Sprintf("[%s]", key)
			}
			return ""
		})
		parser.AddTag("theme_item", func(tag bbparser.Tag, body string) string {
			for key := range tag.Attributes {
				return fmt.Sprintf("theme's '%s'", key)
			}
			return ""
		})
		parser.AddTag("url", func(tag bbparser.Tag, body string) string {
			parser_links[body] = tag.Attributes["starting"]
			return fmt.Sprintf("[%s]", body)
		})

		parser.AddTag("float", func(tag bbparser.Tag, body string) string { return "[Float.X]" })
		parser.AddTag("int", func(tag bbparser.Tag, body string) string { return "int" })
		parser.AddTag("bool", func(tag bbparser.Tag, body string) string { return "bool" })
		parser.AddTag("String", func(tag bbparser.Tag, body string) string { return "string" })
		parser.AddTag("Callable", func(tag bbparser.Tag, body string) string { return "func" })
		parser.AddTag("Variant", func(tag bbparser.Tag, body string) string { return "any" })
		parser.AddTag("Vector2", func(tag bbparser.Tag, body string) string { return "[Vector2.XY]" })
		parser.AddTag("Vector2i", func(tag bbparser.Tag, body string) string { return "[Vector2i.XY]" })
		parser.AddTag("Vector3", func(tag bbparser.Tag, body string) string { return "[Vector3.XYZ]" })
		parser.AddTag("Vector3i", func(tag bbparser.Tag, body string) string { return "[Vector3i.XYZ]" })
		parser.AddTag("Vector4", func(tag bbparser.Tag, body string) string { return "[Vector4.XYZW]" })

		parser.AddTag("Quaternion", func(tag bbparser.Tag, body string) string { return "[Quaternion.IJKX]" })
		parser.AddTag("Color", func(tag bbparser.Tag, body string) string { return "[Color.RGBA]" })
		parser.AddTag("StringName", func(tag bbparser.Tag, body string) string { return "string" })
		parser.AddTag("PackedByteArray", func(tag bbparser.Tag, body string) string { return "[]byte" })
		parser.AddTag("PackedInt32Array", func(tag bbparser.Tag, body string) string { return "[]int32" })
		parser.AddTag("PackedInt64Array", func(tag bbparser.Tag, body string) string { return "[]int64" })
		parser.AddTag("PackedFloat32Array", func(tag bbparser.Tag, body string) string { return "[]float32" })
		parser.AddTag("PackedFloat64Array", func(tag bbparser.Tag, body string) string { return "[]float64" })
		parser.AddTag("PackedVector2Array", func(tag bbparser.Tag, body string) string { return "[][Vector2.XY]" })
		parser.AddTag("PackedVector3Array", func(tag bbparser.Tag, body string) string { return "[][Vector3.XYZ]" })
		parser.AddTag("PackedColorArray", func(tag bbparser.Tag, body string) string { return "[][Color.RGBA]" })
		parser.AddTag("PackedStringArray", func(tag bbparser.Tag, body string) string { return "[]string" })
		parser.AddTag("Transform2D", func(tag bbparser.Tag, body string) string { return "[Transform2D.OriginXY]" })
		parser.AddTag("Transform3D", func(tag bbparser.Tag, body string) string { return "[Transform3D.BasisOrigin]" })
		parser.AddTag("NodePath", func(tag bbparser.Tag, body string) string { return "node path" })
		parser.AddTag("Plane", func(tag bbparser.Tag, body string) string { return "[Plane.NormalD]" })
		parser.AddTag("Rect2", func(tag bbparser.Tag, body string) string { return "[Rect2.PositionSize]" })
		parser.AddTag("Rect2i", func(tag bbparser.Tag, body string) string { return "[Rect2i.PositionSize]" })
		parser.AddTag("AABB", func(tag bbparser.Tag, body string) string { return "[AABB.PositionSize]" })
		parser.AddTag("Basis", func(tag bbparser.Tag, body string) string { return "[Basis.XYZ]" })
		parser.AddTag("RID", func(tag bbparser.Tag, body string) string { return "[Resource.ID]" })
		parser.AddTag("Dictionary", func(tag bbparser.Tag, body string) string { return "data structure" })
		parser.AddTag("Array", func(tag bbparser.Tag, body string) string { return "slice" })
		parser.AddTag("Projection", func(tag bbparser.Tag, body string) string { return "[Projection.XYZW]" })
		parser.AddTag("TextServerFallback", func(tag bbparser.Tag, body string) string {
			return `[graphics.gd/classdb/TextServerFallback]`
		})
		parser.AddTag("@GDScript", func(tag bbparser.Tag, body string) string {
			return `[graphics.gd/classdb/GDScript]`
		})
		for _, class := range classdb {
			if class.Name == parser_class.Name {
				parser.AddTag(class.Name, func(tag bbparser.Tag, body string) string {
					return class.Name
				})
				continue
			}
			parser.AddTag(class.Name, func(tag bbparser.Tag, body string) string {
				return `[graphics.gd/classdb/` + class.Name + `]`
			})
		}
	})
	docs = strings.ReplaceAll(docs, "$DOCS_URL", "https://docs.godotengine.org")
	docs = strings.ReplaceAll(docs, "\n", "\n\n")

	var parsed = parser.Parse(docs)
	if len(parser_links) > 0 {
		parsed += "\n"
	}
	for _, name := range slices.Sorted(maps.Keys(parser_links)) {
		url := parser_links[name]
		parsed += fmt.Sprintf("\n[%s]: %s", name, url)
	}
	return parsed
}
