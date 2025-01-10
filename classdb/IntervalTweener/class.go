// Package IntervalTweener provides methods for working with IntervalTweener object instances.
package IntervalTweener

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Tweener"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
[IntervalTweener] is used to make delays in a tweening sequence. See [method Tween.tween_interval] for more usage information.
[b]Note:[/b] [method Tween.tween_interval] is the only correct way to create [IntervalTweener]. Any [IntervalTweener] created manually will not function correctly.
*/
type Instance [1]gdclass.IntervalTweener

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsIntervalTweener() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.IntervalTweener

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("IntervalTweener"))
	casted := Instance{*(*gdclass.IntervalTweener)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsIntervalTweener() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsIntervalTweener() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTweener() Tweener.Advanced    { return *((*Tweener.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTweener() Tweener.Instance {
	return *((*Tweener.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Tweener.Advanced(self.AsTweener()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Tweener.Instance(self.AsTweener()), name)
	}
}
func init() {
	gdclass.Register("IntervalTweener", func(ptr gd.Object) any {
		return [1]gdclass.IntervalTweener{*(*gdclass.IntervalTweener)(unsafe.Pointer(&ptr))}
	})
}
