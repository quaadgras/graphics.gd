package Node

import (
	"reflect"
	"sync"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/String"
)

// Scene-struct binding: a plain struct type that embeds a class Instance can
// declare exported class-typed fields (and nested structs of them) naming the
// children of a node, mirroring the layout of a scene:
//
//	type SceneMain struct {
//		Node2D.Instance
//
//		Bullets Node2D.Instance
//		Player  struct {
//			Area2D.Instance
//
//			CollisionShape2D CollisionShape2D.Instance
//		}
//	}
//
// When such a struct is cast from an engine object (via Object.To, or by
// instantiating a PackedScene.Is[SceneMain]), each field becomes a deferred
// reference that resolves the correspondingly named child node (or the node
// named by the field's `gd:"..."` tag) on first use. Fields read as null when
// no such node exists, or when it is not of the field's class.
//
// Each (struct type, field) pair registers one [gdreference.Resolver] in the
// deferred function table; the per-cast wrappers carry the scene root's
// object ID as their resolution context.

type sceneField struct {
	offset   uintptr
	resolver gdreference.Resolver
}

var scenePlans sync.Map // reflect.Type -> []sceneField

func init() {
	gd.BindStruct = bindStruct
}

func bindStruct(value any, root [1]gdreference.Object) {
	rtype := reflect.TypeOf(value).Elem()
	if rtype.Kind() != reflect.Struct {
		return
	}
	plan := planFor(rtype)
	if len(plan) == 0 {
		return
	}
	raw := gdreference.GetObject(root[0])
	if raw == 0 {
		return
	}
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(raw, gdextension.CallReturns[gdextension.ObjectID](&id))
	base := reflect.ValueOf(value).UnsafePointer()
	for _, field := range plan {
		*(*gdreference.Object)(unsafe.Add(base, field.offset)) = gdreference.DeferObject(id, field.resolver)
	}
}

func planFor(rtype reflect.Type) []sceneField {
	if plan, ok := scenePlans.Load(rtype); ok {
		return plan.([]sceneField)
	}
	var plan []sceneField
	if _, isExtension := rtype.FieldByName("Extension"); !isExtension {
		// extension classes resolve their fields on Ready instead,
		// see the declarative children support in the classdb package.
		var walk func(rtype reflect.Type, base uintptr, prefix string)
		walk = func(rtype reflect.Type, base uintptr, prefix string) {
			for i := 0; i < rtype.NumField(); i++ {
				field := rtype.Field(i)
				if field.Anonymous {
					continue // bound to the parent node itself.
				}
				name := field.Name
				if tag, ok := field.Tag.Lookup("gd"); ok {
					if tag == "-" {
						continue
					}
					name = tag
				} else if !field.IsExported() {
					continue
				}
				path := prefix + name
				if isInstanceType(field.Type) {
					plan = append(plan, sceneField{
						offset:   base + field.Offset,
						resolver: resolverFor(path, field.Type),
					})
					continue
				}
				if field.Type.Kind() == reflect.Struct {
					if embed, ok := instanceEmbedOf(field.Type); ok {
						plan = append(plan, sceneField{
							offset:   base + field.Offset + embed.Offset,
							resolver: resolverFor(path, embed.Type),
						})
						walk(field.Type, base+field.Offset, path+"/")
					}
				}
			}
		}
		walk(rtype, 0, "")
	}
	existing, _ := scenePlans.LoadOrStore(rtype, plan)
	return existing.([]sceneField)
}

// isInstanceType reports whether rtype is a class Instance wrapper, in the
// layout-compatible-with-[gdreference.Object] sense that scene-struct binding
// relies upon to write deferred references directly into struct fields.
func isInstanceType(rtype reflect.Type) bool {
	return rtype.Size() == unsafe.Sizeof(gdreference.Object{}) &&
		rtype.Kind() != reflect.Struct &&
		reflect.PointerTo(rtype).Implements(reflect.TypeFor[gd.IsClassCastable]())
}

func instanceEmbedOf(rtype reflect.Type) (reflect.StructField, bool) {
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if field.Anonymous && isInstanceType(field.Type) {
			return field, true
		}
	}
	return reflect.StructField{}, false
}

func resolverFor(path string, rtype reflect.Type) gdreference.Resolver {
	return gdreference.NewResolver(func(context gdextension.ObjectID) gdextension.Object {
		root := gdextension.Host.Objects.Lookup(context)
		if root == 0 {
			return 0
		}
		parent := Instance([1]gdclass.Node{gdclass.NewNode(gdreference.RawObject(root))})
		node := class(parent).GetNodeOrNull(Path.ToNode(String.New(path)))
		raw := gdreference.GetObject(gdclass.GetNode(node[0])[0])
		if raw == 0 {
			return 0
		}
		if castable, ok := reflect.New(rtype).Interface().(gd.IsClassCastable); ok && !castable.SetObject([1]gdreference.Object{gdreference.RawObject(raw)}) {
			return 0 // the node is not of the field's class.
		}
		return raw
	})
}
