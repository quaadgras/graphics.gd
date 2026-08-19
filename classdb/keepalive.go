package classdb

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"
	"unsafe"

	"graphics.gd/classdb/Node"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/threadsafe"
	"graphics.gd/variant/Object"
)

// debugKeepalive enables verbose tracing of compile_keepalive's recursion.
// Toggle with GDDEBUG=keepalive.
var debugKeepalive = strings.Contains(os.Getenv("GDDEBUG"), "keepalive")

func keepaliveLog(depth int, format string, args ...any) {
	if !debugKeepalive {
		return
	}
	fmt.Fprintf(os.Stderr, "[keepalive] %s%s\n", strings.Repeat("  ", depth), fmt.Sprintf(format, args...))
}

// roots passed to the engine.
var roots threadsafe.Map[reflect.Value, func(reflect.Value)]

var skips = make(map[reflect.Value]struct{}) // only accessed from [keep_reachable_instances_alive]

// debugKeepaliveCost accumulates per-root-type walk time and dumps a
// sorted table every few hundred frames. Toggle with GDDEBUG=keepalive-cost.
var debugKeepaliveCost = strings.Contains(os.Getenv("GDDEBUG"), "keepalive-cost")

var (
	costByType   = map[reflect.Type]time.Duration{}
	costRoots    = map[reflect.Type]int{}
	costFrames   int
	costInterval = 300
)

//go:linkname keep_reachable_instances_alive
func keep_reachable_instances_alive() {
	clear(skips)
	if debugKeepalive {
		var count int
		for ptr, keepalive := range roots.Iter() {
			count++
			fmt.Fprintf(os.Stderr, "[keepalive] root: %v keepalive=%v\n", ptr.Type(), keepalive != nil)
		}
		fmt.Fprintf(os.Stderr, "[keepalive] === frame: %d roots ===\n", count)
	}
	if debugKeepaliveCost {
		for ptr, keepalive := range roots.Iter() {
			if keepalive == nil {
				continue
			}
			start := time.Now()
			keepalive(ptr)
			costByType[ptr.Type()] += time.Since(start)
			costRoots[ptr.Type()]++
		}
		costFrames++
		if costFrames%costInterval == 0 {
			type row struct {
				t reflect.Type
				d time.Duration
			}
			rows := make([]row, 0, len(costByType))
			total := time.Duration(0)
			for t, d := range costByType {
				rows = append(rows, row{t, d})
				total += d
			}
			sort.Slice(rows, func(i, j int) bool { return rows[i].d > rows[j].d })
			fmt.Fprintf(os.Stderr, "[keepalive-cost] %v/frame over %d frames\n", total/time.Duration(costFrames), costFrames)
			for i, r := range rows {
				if i >= 12 {
					break
				}
				fmt.Fprintf(os.Stderr, "[keepalive-cost]   %8s %5d roots  %v\n", r.d/time.Duration(costFrames), costRoots[r.t]/costFrames, r.t)
			}
			clear(costByType)
			clear(costRoots)
			costFrames = 0
		}
		return
	}
	for ptr, keepalive := range roots.Iter() {
		if keepalive != nil {
			keepalive(ptr)
		}
	}
}

var compiled_keepalives = make(map[reflect.Type]func(reflect.Value))

// compiled_keepalives_mu guards compiled_keepalives: the per-frame keepalive
// walk on the main thread resolves interface values lazily through
// [compile_keepalive] while any goroutine adopting a class instance
// (adopt.go) compiles and caches new types — an unguarded map here was a
// "concurrent map read and map write" crash under load.
var compiled_keepalives_mu sync.Mutex

var keepaliveDepth int

func compile_keepalive(rtype reflect.Type) func(reflect.Value) {
	compiled_keepalives_mu.Lock()
	defer compiled_keepalives_mu.Unlock()
	return compile_keepalive_locked(rtype)
}

// compile_keepalive_for_class builds the walker for a root itself: a
// pointer to an extension class registered in `roots`. Nested pointers
// to adopted classes are skipped as self-rooted (see the Pointer case),
// so the root's own walker is assembled here from the struct walker,
// bypassing that shortcut.
func compile_keepalive_for_class(rtype reflect.Type) func(reflect.Value) {
	compiled_keepalives_mu.Lock()
	defer compiled_keepalives_mu.Unlock()
	keepalive := compile_keepalive_locked(rtype.Elem())
	if keepalive == nil {
		return nil
	}
	return func(val reflect.Value) {
		if val.Kind() != reflect.Pointer || val.IsNil() {
			return
		}
		keepalive(val.Elem())
	}
}

// compile_keepalive_locked is the recursive body of [compile_keepalive];
// compiled_keepalives_mu must be held. Closures it returns run without the
// lock (they re-enter through the public wrapper for lazy interface values).
func compile_keepalive_locked(rtype reflect.Type) (keepalive func(reflect.Value)) {
	if cached, ok := compiled_keepalives[rtype]; ok {
		return cached
	}
	keepaliveDepth++
	defer func() { keepaliveDepth-- }()
	keepaliveLog(keepaliveDepth, "compile_keepalive %v (kind=%v name=%q)", rtype, rtype.Kind(), rtype.Name())
	if rtype.Name() == "Instance" && rtype.Implements(reflect.TypeFor[Object.Any]()) && rtype.Kind() == reflect.Array && rtype.Len() == 1 { // FIXME
		keepaliveLog(keepaliveDepth, "  → MATCHED Instance, returning Object.Use closure")
		return func(ptr reflect.Value) {
			if ptr.CanAddr() {
				Object.Use((*Object.Instance)(ptr.Addr().UnsafePointer()))
			} else {
				Object.Use(ptr.Interface().(Object.Any))
			}
		}
	}
	compiled_keepalives[rtype] = nil // TBD - to support circular references
	defer func() {
		compiled_keepalives[rtype] = keepalive
	}()
	switch rtype.Kind() {
	case reflect.Struct:
		is_extension_class := rtype.Implements(reflect.TypeFor[gdclass.Interface]())
		var keepalives []keep_struct_field_alive
		for field := range rtype.Fields() {
			if is_extension_class && field.Type.Implements(reflect.TypeFor[Node.Any]()) && field.IsExported() {
				continue
			}
			if is_extension_class && field.Index[0] == 0 {
				continue
			}
			if keepalive := compile_keepalive_locked(field.Type); keepalive != nil {
				keepalives = append(keepalives, keep_struct_field_alive{
					rtype:  field.Type,
					index:  field.Index[0],
					offset: field.Offset,
					handle: keepalive,
					public: field.IsExported(),
				})
			}
		}
		if len(keepalives) == 0 {
			return nil
		}
		// No cycle guard here: a struct reached by value cannot be
		// reached twice except through a pointer, interface or map,
		// and those recursion points carry the `skips` guard.
		return func(val reflect.Value) {
			var can_addr = val.CanAddr()
			if is_extension_class {
				if can_addr {
					Object.Use((*Object.Instance)(val.Addr().UnsafePointer()))
				} else {
					Object.Use(Object.Instance(gdclass.GetObjectFromInterface(val.Interface().(gdclass.Interface))))
				}
			}
			if can_addr {
				ptr := val.Addr().UnsafePointer()
				for _, keepalive := range keepalives {
					keepalive.handle(reflect.NewAt(keepalive.rtype, unsafe.Add(ptr, keepalive.offset)).Elem())
				}
			} else {
				for _, keepalive := range keepalives {
					if keepalive.public {
						keepalive.handle(val.Field(keepalive.index))
					}
				}
			}
		}
	case reflect.Array:
		if keepalive := compile_keepalive_locked(rtype.Elem()); keepalive != nil && rtype.Len() > 0 {
			return func(val reflect.Value) {
				for i := 0; i < val.Len(); i++ {
					keepalive(val.Index(i))
				}
			}
		}
		return nil
	case reflect.Pointer:
		// A pointer to an adopted class instance never needs walking:
		// every live instance is inserted into `roots` at construction
		// (adopt.go, register_class.go) and keeps its own graph alive.
		// Game code holds whole registries of these (maps of props,
		// players, and so on), and following each one from every struct
		// that mentions it was the bulk of the per-frame walk. The root
		// itself compiles through [compile_keepalive_for_class], which
		// bypasses this shortcut.
		if rtype.Implements(reflect.TypeFor[gdclass.Interface]()) {
			return nil
		}
		if keepalive := compile_keepalive_locked(rtype.Elem()); keepalive != nil {
			return func(val reflect.Value) {
				if val.IsNil() {
					return
				}
				if _, ok := skips[val]; ok {
					return // shared or cyclic pointer: already walked this frame
				}
				skips[val] = struct{}{}
				keepalive(val.Elem())
			}
		}
		return nil
	case reflect.Slice:
		if keepalive := compile_keepalive_locked(rtype.Elem()); keepalive != nil {
			return func(val reflect.Value) {
				for i := 0; i < val.Len(); i++ {
					keepalive(val.Index(i))
				}
			}
		}
		return nil
	case reflect.Map:
		if keyKeepalive, valKeepalive := compile_keepalive_locked(rtype.Key()), compile_keepalive_locked(rtype.Elem()); keyKeepalive != nil || valKeepalive != nil {
			// Copying every entry out with MapIter.Key/Value allocates a
			// fresh box per entry per frame — the single largest source
			// of garbage in a running game. SetIterKey/SetIterValue reuse
			// one addressable scratch instead (which also lets the
			// Instance case take its no-alloc pointer path). Struct- and
			// map-typed entries keep per-entry copies: their keepalives
			// identity-check the value they receive against `skips`, and
			// a shared scratch would make every entry look like the first.
			keyType, keyReuse := rtype.Key(), scratch_reusable(rtype.Key())
			elemType, elemReuse := rtype.Elem(), scratch_reusable(rtype.Elem())
			return func(val reflect.Value) {
				if _, ok := skips[val]; ok {
					return
				}
				skips[val] = struct{}{}
				var keyScratch, valScratch reflect.Value
				if keyKeepalive != nil && keyReuse {
					keyScratch = reflect.New(keyType).Elem()
				}
				if valKeepalive != nil && elemReuse {
					valScratch = reflect.New(elemType).Elem()
				}
				var map_iter reflect.MapIter
				map_iter.Reset(val)
				for map_iter.Next() {
					if keyKeepalive != nil {
						if keyReuse {
							keyScratch.SetIterKey(&map_iter)
							keyKeepalive(keyScratch)
						} else {
							keyKeepalive(map_iter.Key())
						}
					}
					if valKeepalive != nil {
						if elemReuse {
							valScratch.SetIterValue(&map_iter)
							valKeepalive(valScratch)
						} else {
							valKeepalive(map_iter.Value())
						}
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
			// No guard needed here: a boxed pointer dispatches to the
			// pointer walker, which carries its own `skips` guard, and a
			// boxed value cannot be reached twice.
			if keepalive := compile_keepalive(val.Type()); keepalive != nil {
				keepalive(val)
			}
		}
	default:
		return nil
	}
}

// scratch_reusable reports whether one scratch value can be reused for
// every entry when walking a map of this type: true unless the type's
// keepalive (or one reachable inside a struct or array without
// indirection) identity-checks the value it receives against `skips` —
// a map field's identity is its address, which inside a shared scratch
// is the same for every entry, conflating distinct maps. Indirect kinds
// (pointer, slice, interface) are fine — their keepalives immediately
// resolve to memory outside the scratch. Engine Instance handles get
// their own dedicated keepalive and never consult `skips`.
func scratch_reusable(rtype reflect.Type) bool {
	if rtype.Name() == "Instance" && rtype.Implements(reflect.TypeFor[Object.Any]()) && rtype.Kind() == reflect.Array && rtype.Len() == 1 {
		return true
	}
	switch rtype.Kind() {
	case reflect.Struct, reflect.Map:
		return false
	case reflect.Array:
		return scratch_reusable(rtype.Elem())
	default:
		return true
	}
}

type keep_struct_field_alive struct {
	index int
	rtype reflect.Type

	offset uintptr
	handle func(reflect.Value)
	public bool
}
