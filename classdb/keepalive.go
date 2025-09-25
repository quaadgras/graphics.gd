package classdb

import (
	"fmt"
	"reflect"
	"unsafe"

	"graphics.gd/classdb/Node"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/threadsafe"
	"graphics.gd/variant/Object"
)

var roots threadsafe.Map[unsafe.Pointer, func(unsafe.Pointer)]
var skips = make(map[unsafe.Pointer]struct{})

//go:linkname keep_reachable_instances_alive
func keep_reachable_instances_alive() {
	clear(skips)
	for ptr, keepalive := range roots.Iter() {
		if keepalive != nil {
			keepalive(ptr)
		}
	}
}

var compiled_keepalives = make(map[reflect.Type]func(unsafe.Pointer))

func compile_keepalive(rtype reflect.Type) (keepalive func(unsafe.Pointer)) {
	if rtype.Name() == "Instance" && rtype.Implements(reflect.TypeFor[Object.Any]()) && rtype.Kind() == reflect.Array && rtype.Len() == 1 { // FIXME
		return func(ptr unsafe.Pointer) {
			Object.Use((*Object.Instance)(ptr))
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
			if is_extension_class && field.Type.Implements(reflect.TypeFor[Node.Any]()) {
				continue
			}
			if is_extension_class && i == 0 {
				continue
			}
			fmt.Println(field.Name, field.Type)
			if keepalive := compile_keepalive(field.Type); keepalive != nil {
				keepalives = append(keepalives, keep_struct_field_alive{
					offset: field.Offset,
					handle: keepalive,
				})
			}
		}
		if len(keepalives) == 0 {
			return nil
		}
		return func(ptr unsafe.Pointer) {
			if _, ok := skips[ptr]; ok {
				return
			}
			skips[ptr] = struct{}{}
			if is_extension_class {
				Object.Use((*Object.Instance)(ptr))
			}
			for _, keepalive := range keepalives {
				keepalive.handle(unsafe.Add(ptr, keepalive.offset))
			}
		}
	case reflect.Array:
		if keepalive := compile_keepalive(rtype.Elem()); keepalive != nil && rtype.Len() > 0 {
			return func(ptr unsafe.Pointer) {
				array := reflect.NewAt(rtype, ptr).Elem()
				for i := 0; i < array.Len(); i++ {
					keepalive(array.Index(i).Addr().UnsafePointer())
				}
			}
		}
		return nil
	case reflect.Pointer:
		if keepalive := compile_keepalive(rtype.Elem()); keepalive != nil {
			return func(ptr unsafe.Pointer) {
				p := reflect.NewAt(rtype, ptr)
				if p.IsNil() {
					return
				}
				keepalive(p.UnsafePointer())
			}
		}
		return nil
	case reflect.Slice:
		if keepalive := compile_keepalive(rtype.Elem()); keepalive != nil {
			return func(ptr unsafe.Pointer) {
				slice := reflect.NewAt(rtype, ptr).Elem()
				ptr = slice.UnsafePointer()
				if ptr == nil {
					return
				}
				if _, ok := skips[ptr]; ok {
					return
				}
				skips[ptr] = struct{}{}
				for i := 0; i < slice.Len(); i++ {
					keepalive(slice.Index(i).Addr().UnsafePointer())
				}
			}
		}
		return nil
	case reflect.Map:
		if keyKeepalive, valKeepalive := compile_keepalive(rtype.Key()), compile_keepalive(rtype.Elem()); keyKeepalive != nil || valKeepalive != nil {
			return func(ptr unsafe.Pointer) {
				m := reflect.NewAt(rtype, ptr).Elem()
				ptr = m.UnsafePointer()
				if ptr == nil {
					return
				}
				if _, ok := skips[ptr]; ok {
					return
				}
				skips[ptr] = struct{}{}
				for iter := m.MapRange(); iter.Next(); {
					if keyKeepalive != nil {
						keyKeepalive(iter.Key().Addr().UnsafePointer())
					}
					if valKeepalive != nil {
						valKeepalive(iter.Value().Addr().UnsafePointer())
					}
				}
			}
		}
		return nil
	case reflect.Interface:
		return func(ptr unsafe.Pointer) {
			i := reflect.NewAt(rtype, ptr).Elem()
			if i.IsNil() {
				return
			}
			elem := i.Elem()
			if elem.IsZero() {
				return
			}
			ptr = elem.UnsafePointer()
			if ptr == nil {
				return
			}
			switch elem.Kind() {
			case reflect.Map, reflect.Interface, reflect.Pointer:
			default:
				return
			}
			if _, ok := skips[ptr]; ok {
				return
			}
			skips[ptr] = struct{}{}
			if keepalive := compile_keepalive(elem.Type()); keepalive != nil {
				keepalive(elem.UnsafePointer())
			}
		}
	default:
		return nil
	}
}

type keep_struct_field_alive struct {
	offset uintptr
	handle func(unsafe.Pointer)
}
