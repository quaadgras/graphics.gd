package ZIPReader

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
type Simple [1]classdb.ZIPReader
func (self Simple) Open(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Open(gc.String(path)))
}
func (self Simple) Close() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Close())
}
func (self Simple) GetFiles() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetFiles(gc))
}
func (self Simple) ReadFile(path string, case_sensitive bool) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).ReadFile(gc, gc.String(path), case_sensitive).Bytes())
}
func (self Simple) FileExists(path string, case_sensitive bool) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).FileExists(gc.String(path), case_sensitive))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ZIPReader
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsZIPReader() Expert { return self[0].AsZIPReader() }


//go:nosplit
func (self Simple) AsZIPReader() Simple { return self[0].AsZIPReader() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ZIPReader", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
