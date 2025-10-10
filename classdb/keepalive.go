package classdb

import (
	"reflect"
	"unsafe"

	"graphics.gd/classdb/Node"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/threadsafe"
	"graphics.gd/variant/Object"
)

var roots threadsafe.Map[reflect.Value, func(reflect.Value)]
var skips = make(map[reflect.Value]struct{}) // only accessed from [keep_reachable_instances_alive]

//go:linkname keep_reachable_instances_alive
func keep_reachable_instances_alive() {
	clear(skips)
	for ptr, keepalive := range roots.Iter() {
		if keepalive != nil {
			keepalive(ptr)
		}
	}
}

var compiled_keepalives = make(map[reflect.Type]func(reflect.Value))

func compile_keepalive(rtype reflect.Type) (keepalive func(reflect.Value)) {
	if rtype.Name() == "Instance" && rtype.Implements(reflect.TypeFor[Object.Any]()) && rtype.Kind() == reflect.Array && rtype.Len() == 1 { // FIXME
		return func(ptr reflect.Value) {
			if ptr.CanAddr() {
				Object.Use((*Object.Instance)(ptr.Addr().UnsafePointer()))
			} else {
				Object.Use(ptr.Interface().(Object.Any))
			}
		}
	}
	if cached, ok := compiled_keepalives[rtype]; ok {
		return cached
	}
	compiled_keepalives[rtype] = nil // TBD - to support circular references
	defer func() {
		if keepalive != nil {
			compiled_keepalives[rtype] = keepalive
		} else {
			delete(compiled_keepalives, rtype)
		}
	}()
	switch rtype.Kind() {
	case reflect.Struct:
		is_extension_class := rtype.Implements(reflect.TypeFor[gdclass.Interface]())
		var keepalives []keep_struct_field_alive
		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)
			if is_extension_class && field.Type.Implements(reflect.TypeFor[Node.Any]()) && field.IsExported() {
				continue
			}
			if is_extension_class && i == 0 {
				continue
			}
			if keepalive := compile_keepalive(field.Type); keepalive != nil {
				keepalives = append(keepalives, keep_struct_field_alive{
					rtype:  field.Type,
					index:  field.Index[0],
					offset: field.Offset,
					handle: keepalive,
				})
			}
		}
		if len(keepalives) == 0 {
			return nil
		}
		return func(val reflect.Value) {
			if _, ok := skips[val]; ok {
				return
			}
			skips[val] = struct{}{}
			var can_addr = val.CanAddr()
			if is_extension_class {
				if can_addr {
					Object.Use((*Object.Instance)(val.Addr().UnsafePointer()))
				} else {
					Object.Use(Object.Instance(gdclass.GetObject(val.Interface().(gdclass.Interface))))
				}
			}
			if can_addr {
				ptr := val.Addr().UnsafePointer()
				for _, keepalive := range keepalives {
					keepalive.handle(reflect.NewAt(keepalive.rtype, unsafe.Add(ptr, keepalive.offset)).Elem())
				}
			} else {
				for _, keepalive := range keepalives {
					keepalive.handle(val.Field(keepalive.index))
				}
			}
		}
	case reflect.Array:
		if keepalive := compile_keepalive(rtype.Elem()); keepalive != nil && rtype.Len() > 0 {
			return func(val reflect.Value) {
				for i := 0; i < val.Len(); i++ {
					keepalive(val.Index(i))
				}
			}
		}
		return nil
	case reflect.Pointer:
		if keepalive := compile_keepalive(rtype.Elem()); keepalive != nil {
			return func(val reflect.Value) {
				if val.IsNil() {
					return
				}
				keepalive(val.Elem())
			}
		}
		return nil
	case reflect.Slice:
		if keepalive := compile_keepalive(rtype.Elem()); keepalive != nil {
			return func(val reflect.Value) {
				for i := 0; i < val.Len(); i++ {
					keepalive(val.Index(i))
				}
			}
		}
		return nil
	case reflect.Map:
		if keyKeepalive, valKeepalive := compile_keepalive(rtype.Key()), compile_keepalive(rtype.Elem()); keyKeepalive != nil || valKeepalive != nil {
			return func(val reflect.Value) {
				if _, ok := skips[val]; ok {
					return
				}
				skips[val] = struct{}{}
				var map_iter reflect.MapIter
				map_iter.Reset(val)
				for map_iter.Next() {
					if keyKeepalive != nil {
						keyKeepalive(map_iter.Key())
					}
					if valKeepalive != nil {
						valKeepalive(map_iter.Value())
					}
				}
			}
		}
		return nil
	case reflect.Interface:
		return func(val reflect.Value) {
			if val.IsNil() {
				return
			}
			val = val.Elem()
			if keepalive := compile_keepalive(val.Type()); keepalive != nil {
				keepalive(val)
			}
		}
	default:
		return nil
	}
}

type keep_struct_field_alive struct {
	index int
	rtype reflect.Type

	offset uintptr
	handle func(reflect.Value)
}
