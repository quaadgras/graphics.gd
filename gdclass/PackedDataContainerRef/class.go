package PackedDataContainerRef

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
When packing nested containers using [PackedDataContainer], they are recursively packed into [PackedDataContainerRef] (only applies to [Array] and [Dictionary]). Their data can be retrieved the same way as from [PackedDataContainer].
[codeblock]
var packed = PackedDataContainer.new()
packed.pack([1, 2, 3, ["abc", "def"], 4, 5, 6])

for element in packed:
    if element is PackedDataContainerRef:
        for subelement in element:
            print("::", subelement)
    else:
        print(element)

# Prints:
# 1
# 2
# 3
# ::abc
# ::def
# 4
# 5
# 6
[/codeblock]

*/
type Go [1]classdb.PackedDataContainerRef

/*
Returns the size of the packed container (see [method Array.size] and [method Dictionary.size]).
*/
func (self Go) Size() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).Size()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PackedDataContainerRef
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PackedDataContainerRef"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the size of the packed container (see [method Array.size] and [method Dictionary.size]).
*/
//go:nosplit
func (self class) Size() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PackedDataContainerRef.Bind_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPackedDataContainerRef() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPackedDataContainerRef() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("PackedDataContainerRef", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
