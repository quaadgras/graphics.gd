package PCKPacker

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
The [PCKPacker] is used to create packages that can be loaded into a running project using [method ProjectSettings.load_resource_pack].
[codeblocks]
[gdscript]
var packer = PCKPacker.new()
packer.pck_start("test.pck")
packer.add_file("res://text.txt", "text.txt")
packer.flush()
[/gdscript]
[csharp]
var packer = new PckPacker();
packer.PckStart("test.pck");
packer.AddFile("res://text.txt", "text.txt");
packer.Flush();
[/csharp]
[/codeblocks]
The above [PCKPacker] creates package [code]test.pck[/code], then adds a file named [code]text.txt[/code] at the root of the package.

*/
type Go [1]classdb.PCKPacker

/*
Creates a new PCK file with the name [param pck_name]. The [code].pck[/code] file extension isn't added automatically, so it should be part of [param pck_name] (even though it's not required).
*/
func (self Go) PckStart(pck_name string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).PckStart(gc.String(pck_name), gd.Int(32), gc.String("0000000000000000000000000000000000000000000000000000000000000000"), false))
}

/*
Adds the [param source_path] file to the current PCK package at the [param pck_path] internal path (should start with [code]res://[/code]).
*/
func (self Go) AddFile(pck_path string, source_path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).AddFile(gc.String(pck_path), gc.String(source_path), false))
}

/*
Writes the files specified using all [method add_file] calls since the last flush. If [param verbose] is [code]true[/code], a list of files added will be printed to the console for easier debugging.
*/
func (self Go) Flush() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Flush(false))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PCKPacker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PCKPacker"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Creates a new PCK file with the name [param pck_name]. The [code].pck[/code] file extension isn't added automatically, so it should be part of [param pck_name] (even though it's not required).
*/
//go:nosplit
func (self class) PckStart(pck_name gd.String, alignment gd.Int, key gd.String, encrypt_directory bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(pck_name))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, encrypt_directory)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PCKPacker.Bind_pck_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds the [param source_path] file to the current PCK package at the [param pck_path] internal path (should start with [code]res://[/code]).
*/
//go:nosplit
func (self class) AddFile(pck_path gd.String, source_path gd.String, encrypt bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(pck_path))
	callframe.Arg(frame, mmm.Get(source_path))
	callframe.Arg(frame, encrypt)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PCKPacker.Bind_add_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Writes the files specified using all [method add_file] calls since the last flush. If [param verbose] is [code]true[/code], a list of files added will be printed to the console for easier debugging.
*/
//go:nosplit
func (self class) Flush(verbose bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, verbose)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PCKPacker.Bind_flush, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPCKPacker() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPCKPacker() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("PCKPacker", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
