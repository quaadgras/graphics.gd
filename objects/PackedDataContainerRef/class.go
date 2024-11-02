package PackedDataContainerRef

import "unsafe"
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
type Instance [1]classdb.PackedDataContainerRef

/*
Returns the size of the packed container (see [method Array.size] and [method Dictionary.size]).
*/
func (self Instance) Size() int {
	return int(int(class(self).Size()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PackedDataContainerRef

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PackedDataContainerRef"))
	return Instance{classdb.PackedDataContainerRef(object)}
}

/*
Returns the size of the packed container (see [method Array.size] and [method Dictionary.size]).
*/
//go:nosplit
func (self class) Size() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PackedDataContainerRef.Bind_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPackedDataContainerRef() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPackedDataContainerRef() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("PackedDataContainerRef", func(ptr gd.Object) any { return classdb.PackedDataContainerRef(ptr) })
}
