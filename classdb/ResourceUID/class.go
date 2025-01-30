// Package ResourceUID provides methods for working with ResourceUID object instances.
package ResourceUID

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Resource UIDs (Unique IDentifiers) allow the engine to keep references between resources intact, even if files can renamed or moved. They can be accessed with [code]uid://[/code].
[ResourceUID] keeps track of all registered resource UIDs in a project, generates new UIDs, and converts between their string and integer representations.
*/
var self [1]gdclass.ResourceUID
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ResourceUID)
	self = *(*[1]gdclass.ResourceUID)(unsafe.Pointer(&obj))
}

/*
Converts the given UID to a [code]uid://[/code] string value.
*/
func IdToText(id int) string { //gd:ResourceUID.id_to_text
	once.Do(singleton)
	return string(class(self).IdToText(int64(id)).String())
}

/*
Extracts the UID value from the given [code]uid://[/code] string.
*/
func TextToId(text_id string) int { //gd:ResourceUID.text_to_id
	once.Do(singleton)
	return int(int(class(self).TextToId(String.New(text_id))))
}

/*
Generates a random resource UID which is guaranteed to be unique within the list of currently loaded UIDs.
In order for this UID to be registered, you must call [method add_id] or [method set_id].
*/
func CreateId() int { //gd:ResourceUID.create_id
	once.Do(singleton)
	return int(int(class(self).CreateId()))
}

/*
Returns whether the given UID value is known to the cache.
*/
func HasId(id int) bool { //gd:ResourceUID.has_id
	once.Do(singleton)
	return bool(class(self).HasId(int64(id)))
}

/*
Adds a new UID value which is mapped to the given resource path.
Fails with an error if the UID already exists, so be sure to check [method has_id] beforehand, or use [method set_id] instead.
*/
func AddId(id int, path string) { //gd:ResourceUID.add_id
	once.Do(singleton)
	class(self).AddId(int64(id), String.New(path))
}

/*
Updates the resource path of an existing UID.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand, or use [method add_id] instead.
*/
func SetId(id int, path string) { //gd:ResourceUID.set_id
	once.Do(singleton)
	class(self).SetId(int64(id), String.New(path))
}

/*
Returns the path that the given UID value refers to.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
func GetIdPath(id int) string { //gd:ResourceUID.get_id_path
	once.Do(singleton)
	return string(class(self).GetIdPath(int64(id)).String())
}

/*
Removes a loaded UID value from the cache.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
func RemoveId(id int) { //gd:ResourceUID.remove_id
	once.Do(singleton)
	class(self).RemoveId(int64(id))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.ResourceUID

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Converts the given UID to a [code]uid://[/code] string value.
*/
//go:nosplit
func (self class) IdToText(id int64) String.Readable { //gd:ResourceUID.id_to_text
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_id_to_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Extracts the UID value from the given [code]uid://[/code] string.
*/
//go:nosplit
func (self class) TextToId(text_id String.Readable) int64 { //gd:ResourceUID.text_to_id
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(text_id)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_text_to_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Generates a random resource UID which is guaranteed to be unique within the list of currently loaded UIDs.
In order for this UID to be registered, you must call [method add_id] or [method set_id].
*/
//go:nosplit
func (self class) CreateId() int64 { //gd:ResourceUID.create_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_create_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the given UID value is known to the cache.
*/
//go:nosplit
func (self class) HasId(id int64) bool { //gd:ResourceUID.has_id
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_has_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new UID value which is mapped to the given resource path.
Fails with an error if the UID already exists, so be sure to check [method has_id] beforehand, or use [method set_id] instead.
*/
//go:nosplit
func (self class) AddId(id int64, path String.Readable) { //gd:ResourceUID.add_id
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_add_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Updates the resource path of an existing UID.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand, or use [method add_id] instead.
*/
//go:nosplit
func (self class) SetId(id int64, path String.Readable) { //gd:ResourceUID.set_id
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_set_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the path that the given UID value refers to.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
//go:nosplit
func (self class) GetIdPath(id int64) String.Readable { //gd:ResourceUID.get_id_path
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_get_id_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Removes a loaded UID value from the cache.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
//go:nosplit
func (self class) RemoveId(id int64) { //gd:ResourceUID.remove_id
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceUID.Bind_remove_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("ResourceUID", func(ptr gd.Object) any { return [1]gdclass.ResourceUID{*(*gdclass.ResourceUID)(unsafe.Pointer(&ptr))} })
}
