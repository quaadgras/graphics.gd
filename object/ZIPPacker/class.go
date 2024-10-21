package ZIPPacker

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
This class implements a writer that allows storing the multiple blobs in a zip archive.
[codeblock]
func write_zip_file():
    var writer := ZIPPacker.new()
    var err := writer.open("user://archive.zip")
    if err != OK:
        return err
    writer.start_file("hello.txt")
    writer.write_file("Hello World".to_utf8_buffer())
    writer.close_file()

    writer.close()
    return OK
[/codeblock]

*/
type Simple [1]classdb.ZIPPacker
func (self Simple) Open(path string, append classdb.ZIPPackerZipAppend) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Open(gc.String(path), append))
}
func (self Simple) StartFile(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).StartFile(gc.String(path)))
}
func (self Simple) WriteFile(data []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).WriteFile(gc.PackedByteSlice(data)))
}
func (self Simple) CloseFile() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CloseFile())
}
func (self Simple) Close() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Close())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ZIPPacker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Opens a zip file for writing at the given path using the specified write mode.
This must be called before everything else.
*/
//go:nosplit
func (self class) Open(path gd.String, append classdb.ZIPPackerZipAppend) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, append)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ZIPPacker.Bind_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Starts writing to a file within the archive. Only one file can be written at the same time.
Must be called after [method open].
*/
//go:nosplit
func (self class) StartFile(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ZIPPacker.Bind_start_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Write the given [param data] to the file.
Needs to be called after [method start_file].
*/
//go:nosplit
func (self class) WriteFile(data gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ZIPPacker.Bind_write_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Stops writing to a file within the archive.
It will fail if there is no open file.
*/
//go:nosplit
func (self class) CloseFile() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ZIPPacker.Bind_close_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ZIPPacker.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsZIPPacker() Expert { return self[0].AsZIPPacker() }


//go:nosplit
func (self Simple) AsZIPPacker() Simple { return self[0].AsZIPPacker() }


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
func init() {classdb.Register("ZIPPacker", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ZipAppend = classdb.ZIPPackerZipAppend

const (
/*Create a new zip archive at the given path.*/
	AppendCreate ZipAppend = 0
/*Append a new zip archive to the end of the already existing file at the given path.*/
	AppendCreateafter ZipAppend = 1
/*Add new files to the existing zip archive at the given path.*/
	AppendAddinzip ZipAppend = 2
)
