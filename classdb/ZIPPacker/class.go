// Package ZIPPacker provides methods for working with ZIPPacker object instances.
package ZIPPacker

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

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
type Instance [1]gdclass.ZIPPacker

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsZIPPacker() Instance
}

/*
Opens a zip file for writing at the given path using the specified write mode.
This must be called before everything else.
*/
func (self Instance) Open(path string) error { //gd:ZIPPacker.open
	return error(gd.ToError(class(self).Open(String.New(path), 0)))
}

/*
Starts writing to a file within the archive. Only one file can be written at the same time.
Must be called after [method open].
*/
func (self Instance) StartFile(path string) error { //gd:ZIPPacker.start_file
	return error(gd.ToError(class(self).StartFile(String.New(path))))
}

/*
Write the given [param data] to the file.
Needs to be called after [method start_file].
*/
func (self Instance) WriteFile(data []byte) error { //gd:ZIPPacker.write_file
	return error(gd.ToError(class(self).WriteFile(Packed.Bytes(Packed.New(data...)))))
}

/*
Stops writing to a file within the archive.
It will fail if there is no open file.
*/
func (self Instance) CloseFile() error { //gd:ZIPPacker.close_file
	return error(gd.ToError(class(self).CloseFile()))
}

/*
Closes the underlying resources used by this instance.
*/
func (self Instance) Close() error { //gd:ZIPPacker.close
	return error(gd.ToError(class(self).Close()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ZIPPacker

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ZIPPacker"))
	casted := Instance{*(*gdclass.ZIPPacker)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Opens a zip file for writing at the given path using the specified write mode.
This must be called before everything else.
*/
//go:nosplit
func (self class) Open(path String.Readable, append gdclass.ZIPPackerZipAppend) Error.Code { //gd:ZIPPacker.open
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, append)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPPacker.Bind_open, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Starts writing to a file within the archive. Only one file can be written at the same time.
Must be called after [method open].
*/
//go:nosplit
func (self class) StartFile(path String.Readable) Error.Code { //gd:ZIPPacker.start_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPPacker.Bind_start_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Write the given [param data] to the file.
Needs to be called after [method start_file].
*/
//go:nosplit
func (self class) WriteFile(data Packed.Bytes) Error.Code { //gd:ZIPPacker.write_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPPacker.Bind_write_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Stops writing to a file within the archive.
It will fail if there is no open file.
*/
//go:nosplit
func (self class) CloseFile() Error.Code { //gd:ZIPPacker.close_file
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPPacker.Bind_close_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Closes the underlying resources used by this instance.
*/
//go:nosplit
func (self class) Close() Error.Code { //gd:ZIPPacker.close
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPPacker.Bind_close, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsZIPPacker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsZIPPacker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("ZIPPacker", func(ptr gd.Object) any { return [1]gdclass.ZIPPacker{*(*gdclass.ZIPPacker)(unsafe.Pointer(&ptr))} })
}

type ZipAppend = gdclass.ZIPPackerZipAppend //gd:ZIPPacker.ZipAppend

const (
	/*Create a new zip archive at the given path.*/
	AppendCreate ZipAppend = 0
	/*Append a new zip archive to the end of the already existing file at the given path.*/
	AppendCreateafter ZipAppend = 1
	/*Add new files to the existing zip archive at the given path.*/
	AppendAddinzip ZipAppend = 2
)
