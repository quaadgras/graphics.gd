package Resource

import (
	"reflect"
	"sync"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/String"
)

// ID that uniquely identifies a resource.
type ID = gd.RID

// UID uniquely identifies a resource that has been imported into the graphics directory.
// Remains valid even if the resource is relocated or renamed.
type UID int

var (
	rid_allocate_id       gdextension.FunctionID
	rid_allocate_id_setup = sync.OnceFunc(func() {
		rid_allocate_id = gdextension.Host.Builtin.Functions.Name(pointers.Get(gd.NewStringName("rid_allocate_id")), 701202648)
	})
)

// New allocates a unique ID which can be used by the implementation to construct an RID.
// This is used mainly from native extensions to implement servers.
func AllocateID() ID { //gd:rid_allocate_id
	rid_allocate_id_setup()
	var id int64
	gdextension.Host.Builtin.Functions.Call(rid_allocate_id, gdextension.CallReturns[any](&id), gdextension.SizeInt, nil)
	return Int64(id)
}

// Int64 returns a resource ID from the given int64.
func Int64(id int64) ID { //gd:rid_from_int64
	return *(*gd.RID)(unsafe.Pointer(&id))
}

var preloaded_resources []gd.RefCounted

var startup []func()

func init() {
	preload := func() {
		for _, f := range startup {
			f()
		}
		gd.RegisterCleanup(func() {
			for _, resource := range preloaded_resources {
				if resource.Unreference() {
					gdextension.Host.Objects.Unsafe.Free(gdextension.Object(pointers.Get(resource.AsObject()[0])[0]))
				}
			}
		})
	}
	Callable.Defer(Callable.New(preload))
}

// Load behaves like the builtin "load" function in GDScript. It can also be used to preload a resource if called before
// startup.
func Load[T Any, P string | Path.ToResource](path_to_resource P) T {
	path := Path.ToResource(String.New(path_to_resource))
	if !gd.Linked {
		var placeholder T
		*(*gd.Object)(unsafe.Pointer(&placeholder)) = pointers.Add[gd.Object]([3]uint64{})
		preloaded_resources = append(preloaded_resources, *(*gd.RefCounted)(unsafe.Pointer(&placeholder)))
		startup = append(startup, func() {
			resource := Instance(load(String.Readable(path), String.New(""), 1))
			result, ok := Object.As[T](resource)
			if !ok {
				panic("Resource \"" + path.String() + "\" is " + resource.AsObject()[0].GetClass().String() + " not " + reflect.TypeFor[T]().String())
			}
			raw, ok := pointers.End(result.AsObject()[0])
			if ok {
				pointers.Set(*(*gd.Object)(unsafe.Pointer(&placeholder)), raw)
			}
		})
		return placeholder
	}
	resource := Instance(load(String.Readable(path), String.New(""), 1))
	if resource == (Instance{}) {
		return [1]T{}[0]
	}
	result, ok := Object.As[T](resource)
	if !ok {
		panic("Resource \"" + path.String() + "\" is " + resource.AsObject()[0].GetClass().String() + " not " + reflect.TypeFor[T]().String())
	}
	return result
}

var loader_sname gdextension.StringName
var loader_methods struct {
	load gdextension.MethodForClass `hash:"3358495409"`
}

func init() {
	gd.Links = append(gd.Links, func() {
		loader_sname = gdextension.Host.Strings.Intern.UTF8("ResourceLoader")
		gd.LinkMethods(loader_sname, &loader_methods, false)
	})
	gd.RegisterCleanup(func() {
		pointers.Raw[gd.StringName](loader_sname).Free()
	})
}

var self [1]gdclass.ResourceLoader
var once sync.Once

func singleton() {
	self[0] = pointers.Raw[gdclass.ResourceLoader]([3]uint64{uint64(gdextension.Host.Objects.Global(loader_sname))})
}

func load(path String.Readable, type_hint String.Readable, cache_mode int) [1]gdclass.Resource { //gd:ResourceLoader.load
	once.Do(singleton)
	var r_ret = gdextension.Call[gdextension.Object](gdextension.Object(gd.ObjectChecked(self[0].AsObject())), loader_methods.load, gdextension.SizeObject|(gdextension.SizeString<<4)|(gdextension.SizeString<<8)|(gdextension.SizeInt<<12), unsafe.Pointer(&struct {
		path       gdextension.String
		type_hint  gdextension.String
		cache_mode int64
	}{pointers.Get(gd.InternalString(path)), pointers.Get(gd.InternalString(type_hint)), int64(cache_mode)}))
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret)}
	return ret
}
