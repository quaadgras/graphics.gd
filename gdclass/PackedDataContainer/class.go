package PackedDataContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[PackedDataContainer] can be used to efficiently store data from untyped containers. The data is packed into raw bytes and can be saved to file. Only [Array] and [Dictionary] can be stored this way.
You can retrieve the data by iterating on the container, which will work as if iterating on the packed data itself. If the packed container is a [Dictionary], the data can be retrieved by key names ([String]/[StringName] only).
[codeblock]
var data = { "key": "value", "another_key": 123, "lock": Vector2() }
var packed = PackedDataContainer.new()
packed.pack(data)
ResourceSaver.save(packed, "packed_data.res")
[/codeblock]
[codeblock]
var container = load("packed_data.res")
for key in container:
    prints(key, container[key])

# Prints:
# key value
# lock (0, 0)
# another_key 123
[/codeblock]
Nested containers will be packed recursively. While iterating, they will be returned as [PackedDataContainerRef].

*/
type Go [1]classdb.PackedDataContainer

/*
Packs the given container into a binary representation. The [param value] must be either [Array] or [Dictionary], any other type will result in invalid data error.
[b]Note:[/b] Subsequent calls to this method will overwrite the existing data.
*/
func (self Go) Pack(value gd.Variant) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Pack(value))
}

/*
Returns the size of the packed container (see [method Array.size] and [method Dictionary.size]).
*/
func (self Go) Size() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).Size()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PackedDataContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PackedDataContainer"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Packs the given container into a binary representation. The [param value] must be either [Array] or [Dictionary], any other type will result in invalid data error.
[b]Note:[/b] Subsequent calls to this method will overwrite the existing data.
*/
//go:nosplit
func (self class) Pack(value gd.Variant) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(value))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PackedDataContainer.Bind_pack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the size of the packed container (see [method Array.size] and [method Dictionary.size]).
*/
//go:nosplit
func (self class) Size() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PackedDataContainer.Bind_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPackedDataContainer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPackedDataContainer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("PackedDataContainer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
