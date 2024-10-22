package HashingContext

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
The HashingContext class provides an interface for computing cryptographic hashes over multiple iterations. Useful for computing hashes of big files (so you don't have to load them all in memory), network streams, and data streams in general (so you don't have to hold buffers).
The [enum HashType] enum shows the supported hashing algorithms.
[codeblocks]
[gdscript]
const CHUNK_SIZE = 1024

func hash_file(path):
    # Check that file exists.
    if not FileAccess.file_exists(path):
        return
    # Start an SHA-256 context.
    var ctx = HashingContext.new()
    ctx.start(HashingContext.HASH_SHA256)
    # Open the file to hash.
    var file = FileAccess.open(path, FileAccess.READ)
    # Update the context after reading each chunk.
    while file.get_position() < file.get_length():
        var remaining = file.get_length() - file.get_position()
        ctx.update(file.get_buffer(min(remaining, CHUNK_SIZE)))
    # Get the computed hash.
    var res = ctx.finish()
    # Print the result as hex string and array.
    printt(res.hex_encode(), Array(res))
[/gdscript]
[csharp]
public const int ChunkSize = 1024;

public void HashFile(string path)
{
    // Check that file exists.
    if (!FileAccess.FileExists(path))
    {
        return;
    }
    // Start an SHA-256 context.
    var ctx = new HashingContext();
    ctx.Start(HashingContext.HashType.Sha256);
    // Open the file to hash.
    using var file = FileAccess.Open(path, FileAccess.ModeFlags.Read);
    // Update the context after reading each chunk.
    while (file.GetPosition() < file.GetLength())
    {
        int remaining = (int)(file.GetLength() - file.GetPosition());
        ctx.Update(file.GetBuffer(Mathf.Min(remaining, ChunkSize)));
    }
    // Get the computed hash.
    byte[] res = ctx.Finish();
    // Print the result as hex string and array.
    GD.PrintT(res.HexEncode(), (Variant)res);
}
[/csharp]
[/codeblocks]

*/
type Go [1]classdb.HashingContext

/*
Starts a new hash computation of the given [param type] (e.g. [constant HASH_SHA256] to start computation of an SHA-256).
*/
func (self Go) Start(atype classdb.HashingContextHashType) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Start(atype))
}

/*
Updates the computation with the given [param chunk] of data.
*/
func (self Go) Update(chunk []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Update(gc.PackedByteSlice(chunk)))
}

/*
Closes the current context, and return the computed hash.
*/
func (self Go) Finish() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).Finish(gc).Bytes())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.HashingContext
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("HashingContext"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Starts a new hash computation of the given [param type] (e.g. [constant HASH_SHA256] to start computation of an SHA-256).
*/
//go:nosplit
func (self class) Start(atype classdb.HashingContextHashType) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HashingContext.Bind_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the computation with the given [param chunk] of data.
*/
//go:nosplit
func (self class) Update(chunk gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(chunk))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HashingContext.Bind_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Closes the current context, and return the computed hash.
*/
//go:nosplit
func (self class) Finish(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HashingContext.Bind_finish, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsHashingContext() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsHashingContext() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("HashingContext", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type HashType = classdb.HashingContextHashType

const (
/*Hashing algorithm: MD5.*/
	HashMd5 HashType = 0
/*Hashing algorithm: SHA-1.*/
	HashSha1 HashType = 1
/*Hashing algorithm: SHA-256.*/
	HashSha256 HashType = 2
)
