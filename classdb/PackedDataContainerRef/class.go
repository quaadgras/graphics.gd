// Package PackedDataContainerRef provides methods for working with PackedDataContainerRef object instances.
package PackedDataContainerRef

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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
type Instance [1]gdclass.PackedDataContainerRef

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPackedDataContainerRef() Instance
}

/*
Returns the size of the packed container (see [method Array.size] and [method Dictionary.size]).
*/
func (self Instance) Size() int {
	return int(int(class(self).Size()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PackedDataContainerRef

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PackedDataContainerRef"))
	casted := Instance{*(*gdclass.PackedDataContainerRef)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the size of the packed container (see [method Array.size] and [method Dictionary.size]).
*/
//go:nosplit
func (self class) Size() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PackedDataContainerRef.Bind_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPackedDataContainerRef() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPackedDataContainerRef() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("PackedDataContainerRef", func(ptr gd.Object) any {
		return [1]gdclass.PackedDataContainerRef{*(*gdclass.PackedDataContainerRef)(unsafe.Pointer(&ptr))}
	})
}
