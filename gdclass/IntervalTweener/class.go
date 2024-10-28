package IntervalTweener

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Tweener"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[IntervalTweener] is used to make delays in a tweening sequence. See [method Tween.tween_interval] for more usage information.
[b]Note:[/b] [method Tween.tween_interval] is the only correct way to create [IntervalTweener]. Any [IntervalTweener] created manually will not function correctly.

*/
type Go [1]classdb.IntervalTweener
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.IntervalTweener
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("IntervalTweener"))
	return Go{classdb.IntervalTweener(object)}
}

func (self class) AsIntervalTweener() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsIntervalTweener() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTweener() Tweener.GD { return *((*Tweener.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTweener() Tweener.Go { return *((*Tweener.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTweener(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTweener(), name)
	}
}
func init() {classdb.Register("IntervalTweener", func(ptr gd.Object) any { return classdb.IntervalTweener(ptr) })}
