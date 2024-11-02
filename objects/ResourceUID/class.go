package ResourceUID

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Resource UIDs (Unique IDentifiers) allow the engine to keep references between resources intact, even if files can renamed or moved. They can be accessed with [code]uid://[/code].
[ResourceUID] keeps track of all registered resource UIDs in a project, generates new UIDs, and converts between their string and integer representations.
*/
var self objects.ResourceUID
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ResourceUID)
	self = *(*objects.ResourceUID)(unsafe.Pointer(&obj))
}

/*
Converts the given UID to a [code]uid://[/code] string value.
*/
func IdToText(id int) string {
	once.Do(singleton)
	return string(class(self).IdToText(gd.Int(id)).String())
}

/*
Extracts the UID value from the given [code]uid://[/code] string.
*/
func TextToId(text_id string) int {
	once.Do(singleton)
	return int(int(class(self).TextToId(gd.NewString(text_id))))
}

/*
Generates a random resource UID which is guaranteed to be unique within the list of currently loaded UIDs.
In order for this UID to be registered, you must call [method add_id] or [method set_id].
*/
func CreateId() int {
	once.Do(singleton)
	return int(int(class(self).CreateId()))
}

/*
Returns whether the given UID value is known to the cache.
*/
func HasId(id int) bool {
	once.Do(singleton)
	return bool(class(self).HasId(gd.Int(id)))
}

/*
Adds a new UID value which is mapped to the given resource path.
Fails with an error if the UID already exists, so be sure to check [method has_id] beforehand, or use [method set_id] instead.
*/
func AddId(id int, path string) {
	once.Do(singleton)
	class(self).AddId(gd.Int(id), gd.NewString(path))
}

/*
Updates the resource path of an existing UID.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand, or use [method add_id] instead.
*/
func SetId(id int, path string) {
	once.Do(singleton)
	class(self).SetId(gd.Int(id), gd.NewString(path))
}

/*
Returns the path that the given UID value refers to.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
func GetIdPath(id int) string {
	once.Do(singleton)
	return string(class(self).GetIdPath(gd.Int(id)).String())
}

/*
Removes a loaded UID value from the cache.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
func RemoveId(id int) {
	once.Do(singleton)
	class(self).RemoveId(gd.Int(id))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.ResourceUID

func (self class) AsObject() gd.Object { return self[0].AsObject() }

/*
Converts the given UID to a [code]uid://[/code] string value.
*/
//go:nosplit
func (self class) IdToText(id gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_id_to_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Extracts the UID value from the given [code]uid://[/code] string.
*/
//go:nosplit
func (self class) TextToId(text_id gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text_id))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_text_to_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Generates a random resource UID which is guaranteed to be unique within the list of currently loaded UIDs.
In order for this UID to be registered, you must call [method add_id] or [method set_id].
*/
//go:nosplit
func (self class) CreateId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_create_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the given UID value is known to the cache.
*/
//go:nosplit
func (self class) HasId(id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_has_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new UID value which is mapped to the given resource path.
Fails with an error if the UID already exists, so be sure to check [method has_id] beforehand, or use [method set_id] instead.
*/
//go:nosplit
func (self class) AddId(id gd.Int, path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_add_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Updates the resource path of an existing UID.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand, or use [method add_id] instead.
*/
//go:nosplit
func (self class) SetId(id gd.Int, path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_set_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the path that the given UID value refers to.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
//go:nosplit
func (self class) GetIdPath(id gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_get_id_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes a loaded UID value from the cache.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
//go:nosplit
func (self class) RemoveId(id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_remove_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("ResourceUID", func(ptr gd.Object) any { return classdb.ResourceUID(ptr) })
}
