package IntervalTweener

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Tweener"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[IntervalTweener] is used to make delays in a tweening sequence. See [method Tween.tween_interval] for more usage information.
[b]Note:[/b] [method Tween.tween_interval] is the only correct way to create [IntervalTweener]. Any [IntervalTweener] created manually will not function correctly.

*/
type Simple [1]classdb.IntervalTweener
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.IntervalTweener
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsIntervalTweener() Expert { return self[0].AsIntervalTweener() }


//go:nosplit
func (self Simple) AsIntervalTweener() Simple { return self[0].AsIntervalTweener() }


//go:nosplit
func (self class) AsTweener() Tweener.Expert { return self[0].AsTweener() }


//go:nosplit
func (self Simple) AsTweener() Tweener.Simple { return self[0].AsTweener() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("IntervalTweener", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
