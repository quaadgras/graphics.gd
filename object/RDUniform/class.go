package RDUniform

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object is used by [RenderingDevice].

*/
type Simple [1]classdb.RDUniform
func (self Simple) SetUniformType(p_member classdb.RenderingDeviceUniformType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUniformType(p_member)
}
func (self Simple) GetUniformType() classdb.RenderingDeviceUniformType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceUniformType(Expert(self).GetUniformType())
}
func (self Simple) SetBinding(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBinding(gd.Int(p_member))
}
func (self Simple) GetBinding() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBinding()))
}
func (self Simple) AddId(id gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddId(id)
}
func (self Simple) ClearIds() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearIds()
}
func (self Simple) GetIds() gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.RID](Expert(self).GetIds(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDUniform
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetUniformType(p_member classdb.RenderingDeviceUniformType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDUniform.Bind_set_uniform_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUniformType() classdb.RenderingDeviceUniformType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceUniformType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDUniform.Bind_get_uniform_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBinding(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDUniform.Bind_set_binding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBinding() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDUniform.Bind_get_binding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Binds the given id to the uniform. The data associated with the id is then used when the uniform is passed to a shader.
*/
//go:nosplit
func (self class) AddId(id gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDUniform.Bind_add_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unbinds all ids currently bound to the uniform.
*/
//go:nosplit
func (self class) ClearIds()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDUniform.Bind_clear_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of all ids currently bound to the uniform.
*/
//go:nosplit
func (self class) GetIds(ctx gd.Lifetime) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDUniform.Bind_get_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.RID](ret)
}

//go:nosplit
func (self class) AsRDUniform() Expert { return self[0].AsRDUniform() }


//go:nosplit
func (self Simple) AsRDUniform() Simple { return self[0].AsRDUniform() }


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
func init() {classdb.Register("RDUniform", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
