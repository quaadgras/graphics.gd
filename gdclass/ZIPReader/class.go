package ZIPReader

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
This class implements a reader that can extract the content of individual files inside a zip archive.
[codeblock]
func read_zip_file():
    var reader := ZIPReader.new()
    var err := reader.open("user://archive.zip")
    if err != OK:
        return PackedByteArray()
    var res := reader.read_file("hello.txt")
    reader.close()
    return res
[/codeblock]

*/
type Go [1]classdb.ZIPReader

/*
Opens the zip archive at the given [param path] and reads its file index.
*/
func (self Go) Open(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Open(gc.String(path)))
}

/*
Closes the underlying resources used by this instance.
*/
func (self Go) Close() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Close())
}

/*
Returns the list of names of all files in the loaded archive.
Must be called after [method open].
*/
func (self Go) GetFiles() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetFiles(gc).Strings(gc))
}

/*
Loads the whole content of a file in the loaded zip archive into memory and returns it.
Must be called after [method open].
*/
func (self Go) ReadFile(path string) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).ReadFile(gc, gc.String(path), true).Bytes())
}

/*
Returns [code]true[/code] if the file exists in the loaded zip archive.
Must be called after [method open].
*/
func (self Go) FileExists(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FileExists(gc.String(path), true))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ZIPReader
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ZIPReader"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Opens the zip archive at the given [param path] and reads its file index.
*/
//go:nosplit
func (self class) Open(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ZIPReader.Bind_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Closes the underlying resources used by this instance.
*/
//go:nosplit
func (self class) Close() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ZIPReader.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of names of all files in the loaded archive.
Must be called after [method open].
*/
//go:nosplit
func (self class) GetFiles(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ZIPReader.Bind_get_files, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Loads the whole content of a file in the loaded zip archive into memory and returns it.
Must be called after [method open].
*/
//go:nosplit
func (self class) ReadFile(ctx gd.Lifetime, path gd.String, case_sensitive bool) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, case_sensitive)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ZIPReader.Bind_read_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the file exists in the loaded zip archive.
Must be called after [method open].
*/
//go:nosplit
func (self class) FileExists(path gd.String, case_sensitive bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, case_sensitive)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ZIPReader.Bind_file_exists, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsZIPReader() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsZIPReader() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("ZIPReader", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
