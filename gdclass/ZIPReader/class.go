package ZIPReader

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

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
type Instance [1]classdb.ZIPReader

/*
Opens the zip archive at the given [param path] and reads its file index.
*/
func (self Instance) Open(path string) gd.Error {
	return gd.Error(class(self).Open(gd.NewString(path)))
}

/*
Closes the underlying resources used by this instance.
*/
func (self Instance) Close() gd.Error {
	return gd.Error(class(self).Close())
}

/*
Returns the list of names of all files in the loaded archive.
Must be called after [method open].
*/
func (self Instance) GetFiles() []string {
	return []string(class(self).GetFiles().Strings())
}

/*
Loads the whole content of a file in the loaded zip archive into memory and returns it.
Must be called after [method open].
*/
func (self Instance) ReadFile(path string) []byte {
	return []byte(class(self).ReadFile(gd.NewString(path), true).Bytes())
}

/*
Returns [code]true[/code] if the file exists in the loaded zip archive.
Must be called after [method open].
*/
func (self Instance) FileExists(path string) bool {
	return bool(class(self).FileExists(gd.NewString(path), true))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ZIPReader

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ZIPReader"))
	return Instance{classdb.ZIPReader(object)}
}

/*
Opens the zip archive at the given [param path] and reads its file index.
*/
//go:nosplit
func (self class) Open(path gd.String) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPReader.Bind_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Closes the underlying resources used by this instance.
*/
//go:nosplit
func (self class) Close() int64 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPReader.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of names of all files in the loaded archive.
Must be called after [method open].
*/
//go:nosplit
func (self class) GetFiles() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPReader.Bind_get_files, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Loads the whole content of a file in the loaded zip archive into memory and returns it.
Must be called after [method open].
*/
//go:nosplit
func (self class) ReadFile(path gd.String, case_sensitive bool) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, case_sensitive)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPReader.Bind_read_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the file exists in the loaded zip archive.
Must be called after [method open].
*/
//go:nosplit
func (self class) FileExists(path gd.String, case_sensitive bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, case_sensitive)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPReader.Bind_file_exists, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsZIPReader() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsZIPReader() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
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
func init() { classdb.Register("ZIPReader", func(ptr gd.Object) any { return classdb.ZIPReader(ptr) }) }
