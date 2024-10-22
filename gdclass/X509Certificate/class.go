package X509Certificate

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
The X509Certificate class represents an X509 certificate. Certificates can be loaded and saved like any other [Resource].
They can be used as the server certificate in [method StreamPeerTLS.accept_stream] (along with the proper [CryptoKey]), and to specify the only certificate that should be accepted when connecting to a TLS server via [method StreamPeerTLS.connect_to_stream].

*/
type Go [1]classdb.X509Certificate

/*
Saves a certificate to the given [param path] (should be a "*.crt" file).
*/
func (self Go) Save(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Save(gc.String(path)))
}

/*
Loads a certificate from [param path] ("*.crt" file).
*/
func (self Go) Load(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Load(gc.String(path)))
}

/*
Returns a string representation of the certificate, or an empty string if the certificate is invalid.
*/
func (self Go) SaveToString() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).SaveToString(gc).String())
}

/*
Loads a certificate from the given [param string].
*/
func (self Go) LoadFromString(s string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).LoadFromString(gc.String(s)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.X509Certificate
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("X509Certificate"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Saves a certificate to the given [param path] (should be a "*.crt" file).
*/
//go:nosplit
func (self class) Save(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.X509Certificate.Bind_save, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Loads a certificate from [param path] ("*.crt" file).
*/
//go:nosplit
func (self class) Load(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.X509Certificate.Bind_load, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a string representation of the certificate, or an empty string if the certificate is invalid.
*/
//go:nosplit
func (self class) SaveToString(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.X509Certificate.Bind_save_to_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Loads a certificate from the given [param string].
*/
//go:nosplit
func (self class) LoadFromString(s gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.X509Certificate.Bind_load_from_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsX509Certificate() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsX509Certificate() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("X509Certificate", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
