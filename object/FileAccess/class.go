package FileAccess

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
This class can be used to permanently store data in the user device's file system and to read from it. This is useful for store game save data or player configuration files.
Here's a sample on how to write and read from a file:
[codeblocks]
[gdscript]
func save_to_file(content):
    var file = FileAccess.open("user://save_game.dat", FileAccess.WRITE)
    file.store_string(content)

func load_from_file():
    var file = FileAccess.open("user://save_game.dat", FileAccess.READ)
    var content = file.get_as_text()
    return content
[/gdscript]
[csharp]
public void SaveToFile(string content)
{
    using var file = FileAccess.Open("user://save_game.dat", FileAccess.ModeFlags.Write);
    file.StoreString(content);
}

public string LoadFromFile()
{
    using var file = FileAccess.Open("user://save_game.dat", FileAccess.ModeFlags.Read);
    string content = file.GetAsText();
    return content;
}
[/csharp]
[/codeblocks]
In the example above, the file will be saved in the user data folder as specified in the [url=$DOCS_URL/tutorials/io/data_paths.html]Data paths[/url] documentation.
[FileAccess] will close when it's freed, which happens when it goes out of scope or when it gets assigned with [code]null[/code]. [method close] can be used to close it before then explicitly. In C# the reference must be disposed manually, which can be done with the [code]using[/code] statement or by calling the [code]Dispose[/code] method directly.
[b]Note:[/b] To access project resources once exported, it is recommended to use [ResourceLoader] instead of [FileAccess], as some files are converted to engine-specific formats and their original source files might not be present in the exported PCK package.
[b]Note:[/b] Files are automatically closed only if the process exits "normally" (such as by clicking the window manager's close button or pressing [b]Alt + F4[/b]). If you stop the project execution by pressing [b]F8[/b] while the project is running, the file won't be closed as the game process will be killed. You can work around this by calling [method flush] at regular intervals.

*/
type Simple [1]classdb.FileAccess
func (self Simple) Open(path string, flags classdb.FileAccessModeFlags) [1]classdb.FileAccess {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.FileAccess(Expert(self).Open(gc, gc.String(path), flags))
}
func (self Simple) OpenEncrypted(path string, mode_flags classdb.FileAccessModeFlags, key []byte) [1]classdb.FileAccess {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.FileAccess(Expert(self).OpenEncrypted(gc, gc.String(path), mode_flags, gc.PackedByteSlice(key)))
}
func (self Simple) OpenEncryptedWithPass(path string, mode_flags classdb.FileAccessModeFlags, pass string) [1]classdb.FileAccess {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.FileAccess(Expert(self).OpenEncryptedWithPass(gc, gc.String(path), mode_flags, gc.String(pass)))
}
func (self Simple) OpenCompressed(path string, mode_flags classdb.FileAccessModeFlags, compression_mode classdb.FileAccessCompressionMode) [1]classdb.FileAccess {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.FileAccess(Expert(self).OpenCompressed(gc, gc.String(path), mode_flags, compression_mode))
}
func (self Simple) GetOpenError() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).GetOpenError(gc))
}
func (self Simple) GetFileAsBytes(path string) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).GetFileAsBytes(gc, gc.String(path)).Bytes())
}
func (self Simple) GetFileAsString(path string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetFileAsString(gc, gc.String(path)).String())
}
func (self Simple) Resize(length int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Resize(gd.Int(length)))
}
func (self Simple) Flush() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Flush()
}
func (self Simple) GetPath() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetPath(gc).String())
}
func (self Simple) GetPathAbsolute() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetPathAbsolute(gc).String())
}
func (self Simple) IsOpen() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOpen())
}
func (self Simple) SeekTo(position int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SeekTo(gd.Int(position))
}
func (self Simple) SeekEnd(position int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SeekEnd(gd.Int(position))
}
func (self Simple) GetPosition() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPosition()))
}
func (self Simple) GetLength() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLength()))
}
func (self Simple) EofReached() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).EofReached())
}
func (self Simple) Get8() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).Get8()))
}
func (self Simple) Get16() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).Get16()))
}
func (self Simple) Get32() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).Get32()))
}
func (self Simple) Get64() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).Get64()))
}
func (self Simple) GetFloat() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFloat()))
}
func (self Simple) GetDouble() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDouble()))
}
func (self Simple) GetReal() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetReal()))
}
func (self Simple) GetBuffer(length int) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).GetBuffer(gc, gd.Int(length)).Bytes())
}
func (self Simple) GetLine() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLine(gc).String())
}
func (self Simple) GetCsvLine(delim string) gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetCsvLine(gc, gc.String(delim)))
}
func (self Simple) GetAsText(skip_cr bool) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetAsText(gc, skip_cr).String())
}
func (self Simple) GetMd5(path string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetMd5(gc, gc.String(path)).String())
}
func (self Simple) GetSha256(path string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSha256(gc, gc.String(path)).String())
}
func (self Simple) IsBigEndian() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBigEndian())
}
func (self Simple) SetBigEndian(big_endian bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBigEndian(big_endian)
}
func (self Simple) GetError() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).GetError())
}
func (self Simple) GetVar(allow_objects bool) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetVar(gc, allow_objects))
}
func (self Simple) Store8(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Store8(gd.Int(value))
}
func (self Simple) Store16(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Store16(gd.Int(value))
}
func (self Simple) Store32(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Store32(gd.Int(value))
}
func (self Simple) Store64(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Store64(gd.Int(value))
}
func (self Simple) StoreFloat(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StoreFloat(gd.Float(value))
}
func (self Simple) StoreDouble(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StoreDouble(gd.Float(value))
}
func (self Simple) StoreReal(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StoreReal(gd.Float(value))
}
func (self Simple) StoreBuffer(buffer []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StoreBuffer(gc.PackedByteSlice(buffer))
}
func (self Simple) StoreLine(line string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StoreLine(gc.String(line))
}
func (self Simple) StoreCsvLine(values gd.PackedStringArray, delim string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StoreCsvLine(values, gc.String(delim))
}
func (self Simple) StoreString(s string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StoreString(gc.String(s))
}
func (self Simple) StoreVar(value gd.Variant, full_objects bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StoreVar(value, full_objects)
}
func (self Simple) StorePascalString(s string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StorePascalString(gc.String(s))
}
func (self Simple) GetPascalString() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetPascalString(gc).String())
}
func (self Simple) Close() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Close()
}
func (self Simple) FileExists(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).FileExists(gc, gc.String(path)))
}
func (self Simple) GetModifiedTime(file string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetModifiedTime(gc, gc.String(file))))
}
func (self Simple) GetUnixPermissions(file string) classdb.FileAccessUnixPermissionFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.FileAccessUnixPermissionFlags(Expert(self).GetUnixPermissions(gc, gc.String(file)))
}
func (self Simple) SetUnixPermissions(file string, permissions classdb.FileAccessUnixPermissionFlags) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SetUnixPermissions(gc, gc.String(file), permissions))
}
func (self Simple) GetHiddenAttribute(file string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetHiddenAttribute(gc, gc.String(file)))
}
func (self Simple) SetHiddenAttribute(file string, hidden bool) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SetHiddenAttribute(gc, gc.String(file), hidden))
}
func (self Simple) SetReadOnlyAttribute(file string, ro bool) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SetReadOnlyAttribute(gc, gc.String(file), ro))
}
func (self Simple) GetReadOnlyAttribute(file string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetReadOnlyAttribute(gc, gc.String(file)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.FileAccess
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates a new [FileAccess] object and opens the file for writing or reading, depending on the flags.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) Open(ctx gd.Lifetime, path gd.String, flags classdb.FileAccessModeFlags) object.FileAccess {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.FileAccess
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new [FileAccess] object and opens an encrypted file in write or read mode. You need to pass a binary key to encrypt/decrypt it.
[b]Note:[/b] The provided key must be 32 bytes long.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) OpenEncrypted(ctx gd.Lifetime, path gd.String, mode_flags classdb.FileAccessModeFlags, key gd.PackedByteArray) object.FileAccess {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mode_flags)
	callframe.Arg(frame, mmm.Get(key))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_open_encrypted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.FileAccess
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new [FileAccess] object and opens an encrypted file in write or read mode. You need to pass a password to encrypt/decrypt it.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) OpenEncryptedWithPass(ctx gd.Lifetime, path gd.String, mode_flags classdb.FileAccessModeFlags, pass gd.String) object.FileAccess {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mode_flags)
	callframe.Arg(frame, mmm.Get(pass))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_open_encrypted_with_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.FileAccess
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new [FileAccess] object and opens a compressed file for reading or writing.
[b]Note:[/b] [method open_compressed] can only read files that were saved by Godot, not third-party compression formats. See [url=https://github.com/godotengine/godot/issues/28999]GitHub issue #28999[/url] for a workaround.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) OpenCompressed(ctx gd.Lifetime, path gd.String, mode_flags classdb.FileAccessModeFlags, compression_mode classdb.FileAccessCompressionMode) object.FileAccess {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mode_flags)
	callframe.Arg(frame, compression_mode)
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_open_compressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.FileAccess
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the result of the last [method open] call in the current thread.
*/
//go:nosplit
func (self class) GetOpenError(ctx gd.Lifetime) int64 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_get_open_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the whole [param path] file contents as a [PackedByteArray] without any decoding.
Returns an empty [PackedByteArray] if an error occurred while opening the file. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) GetFileAsBytes(ctx gd.Lifetime, path gd.String) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_get_file_as_bytes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the whole [param path] file contents as a [String]. Text is interpreted as being UTF-8 encoded.
Returns an empty [String] if an error occurred while opening the file. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) GetFileAsString(ctx gd.Lifetime, path gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_get_file_as_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Resizes the file to a specified length. The file must be open in a mode that permits writing. If the file is extended, NUL characters are appended. If the file is truncated, all data from the end file to the original length of the file is lost.
*/
//go:nosplit
func (self class) Resize(length gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_resize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Writes the file's buffer to disk. Flushing is automatically performed when the file is closed. This means you don't need to call [method flush] manually before closing a file. Still, calling [method flush] can be used to ensure the data is safe even if the project crashes instead of being closed gracefully.
[b]Note:[/b] Only call [method flush] when you actually need it. Otherwise, it will decrease performance due to constant disk writes.
*/
//go:nosplit
func (self class) Flush()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_flush, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the path as a [String] for the current open file.
*/
//go:nosplit
func (self class) GetPath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the absolute path as a [String] for the current open file.
*/
//go:nosplit
func (self class) GetPathAbsolute(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_path_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the file is currently opened.
*/
//go:nosplit
func (self class) IsOpen() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_is_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Changes the file reading/writing cursor to the specified position (in bytes from the beginning of the file).
*/
//go:nosplit
func (self class) SeekTo(position gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Changes the file reading/writing cursor to the specified position (in bytes from the end of the file).
[b]Note:[/b] This is an offset, so you should use negative numbers or the cursor will be at the end of the file.
*/
//go:nosplit
func (self class) SeekEnd(position gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_seek_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the file cursor's position.
*/
//go:nosplit
func (self class) GetPosition() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the size of the file in bytes.
*/
//go:nosplit
func (self class) GetLength() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the file cursor has already read past the end of the file.
[b]Note:[/b] [code]eof_reached() == false[/code] cannot be used to check whether there is more data available. To loop while there is more data available, use:
[codeblocks]
[gdscript]
while file.get_position() < file.get_length():
    # Read data
[/gdscript]
[csharp]
while (file.GetPosition() < file.GetLength())
{
    // Read data
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) EofReached() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_eof_reached, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next 8 bits from the file as an integer. See [method store_8] for details on what values can be stored and retrieved this way.
*/
//go:nosplit
func (self class) Get8() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next 16 bits from the file as an integer. See [method store_16] for details on what values can be stored and retrieved this way.
*/
//go:nosplit
func (self class) Get16() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_16, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next 32 bits from the file as an integer. See [method store_32] for details on what values can be stored and retrieved this way.
*/
//go:nosplit
func (self class) Get32() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_32, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next 64 bits from the file as an integer. See [method store_64] for details on what values can be stored and retrieved this way.
*/
//go:nosplit
func (self class) Get64() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next 32 bits from the file as a floating-point number.
*/
//go:nosplit
func (self class) GetFloat() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_float, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next 64 bits from the file as a floating-point number.
*/
//go:nosplit
func (self class) GetDouble() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_double, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next bits from the file as a floating-point number.
*/
//go:nosplit
func (self class) GetReal() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_real, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns next [param length] bytes of the file as a [PackedByteArray].
*/
//go:nosplit
func (self class) GetBuffer(ctx gd.Lifetime, length gd.Int) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the next line of the file as a [String]. The returned string doesn't include newline ([code]\n[/code]) or carriage return ([code]\r[/code]) characters, but does include any other leading or trailing whitespace.
Text is interpreted as being UTF-8 encoded.
*/
//go:nosplit
func (self class) GetLine(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the next value of the file in CSV (Comma-Separated Values) format. You can pass a different delimiter [param delim] to use other than the default [code]","[/code] (comma). This delimiter must be one-character long, and cannot be a double quotation mark.
Text is interpreted as being UTF-8 encoded. Text values must be enclosed in double quotes if they include the delimiter character. Double quotes within a text value can be escaped by doubling their occurrence.
For example, the following CSV lines are valid and will be properly parsed as two strings each:
[codeblock lang=text]
Alice,"Hello, Bob!"
Bob,Alice! What a surprise!
Alice,"I thought you'd reply with ""Hello, world""."
[/codeblock]
Note how the second line can omit the enclosing quotes as it does not include the delimiter. However it [i]could[/i] very well use quotes, it was only written without for demonstration purposes. The third line must use [code]""[/code] for each quotation mark that needs to be interpreted as such instead of the end of a text value.
*/
//go:nosplit
func (self class) GetCsvLine(ctx gd.Lifetime, delim gd.String) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(delim))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_csv_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the whole file as a [String]. Text is interpreted as being UTF-8 encoded.
If [param skip_cr] is [code]true[/code], carriage return characters ([code]\r[/code], CR) will be ignored when parsing the UTF-8, so that only line feed characters ([code]\n[/code], LF) represent a new line (Unix convention).
*/
//go:nosplit
func (self class) GetAsText(ctx gd.Lifetime, skip_cr bool) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skip_cr)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_as_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an MD5 String representing the file at the given path or an empty [String] on failure.
*/
//go:nosplit
func (self class) GetMd5(ctx gd.Lifetime, path gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_get_md5, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an SHA-256 [String] representing the file at the given path or an empty [String] on failure.
*/
//go:nosplit
func (self class) GetSha256(ctx gd.Lifetime, path gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_get_sha256, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) IsBigEndian() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_is_big_endian, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBigEndian(big_endian bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, big_endian)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_set_big_endian, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the last error that happened when trying to perform operations. Compare with the [code]ERR_FILE_*[/code] constants from [enum Error].
*/
//go:nosplit
func (self class) GetError() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next [Variant] value from the file. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
//go:nosplit
func (self class) GetVar(ctx gd.Lifetime, allow_objects bool) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, allow_objects)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_var, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Stores an integer as 8 bits in the file.
[b]Note:[/b] The [param value] should lie in the interval [code][0, 255][/code]. Any other value will overflow and wrap around.
To store a signed integer, use [method store_64], or convert it manually (see [method store_16] for an example).
*/
//go:nosplit
func (self class) Store8(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stores an integer as 16 bits in the file.
[b]Note:[/b] The [param value] should lie in the interval [code][0, 2^16 - 1][/code]. Any other value will overflow and wrap around.
To store a signed integer, use [method store_64] or store a signed integer from the interval [code][-2^15, 2^15 - 1][/code] (i.e. keeping one bit for the signedness) and compute its sign manually when reading. For example:
[codeblocks]
[gdscript]
const MAX_15B = 1 << 15
const MAX_16B = 1 << 16

func unsigned16_to_signed(unsigned):
    return (unsigned + MAX_15B) % MAX_16B - MAX_15B

func _ready():
    var f = FileAccess.open("user://file.dat", FileAccess.WRITE_READ)
    f.store_16(-42) # This wraps around and stores 65494 (2^16 - 42).
    f.store_16(121) # In bounds, will store 121.
    f.seek(0) # Go back to start to read the stored value.
    var read1 = f.get_16() # 65494
    var read2 = f.get_16() # 121
    var converted1 = unsigned16_to_signed(read1) # -42
    var converted2 = unsigned16_to_signed(read2) # 121
[/gdscript]
[csharp]
public override void _Ready()
{
    using var f = FileAccess.Open("user://file.dat", FileAccess.ModeFlags.WriteRead);
    f.Store16(unchecked((ushort)-42)); // This wraps around and stores 65494 (2^16 - 42).
    f.Store16(121); // In bounds, will store 121.
    f.Seek(0); // Go back to start to read the stored value.
    ushort read1 = f.Get16(); // 65494
    ushort read2 = f.Get16(); // 121
    short converted1 = (short)read1; // -42
    short converted2 = (short)read2; // 121
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) Store16(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_16, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stores an integer as 32 bits in the file.
[b]Note:[/b] The [param value] should lie in the interval [code][0, 2^32 - 1][/code]. Any other value will overflow and wrap around.
To store a signed integer, use [method store_64], or convert it manually (see [method store_16] for an example).
*/
//go:nosplit
func (self class) Store32(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_32, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stores an integer as 64 bits in the file.
[b]Note:[/b] The [param value] must lie in the interval [code][-2^63, 2^63 - 1][/code] (i.e. be a valid [int] value).
*/
//go:nosplit
func (self class) Store64(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stores a floating-point number as 32 bits in the file.
*/
//go:nosplit
func (self class) StoreFloat(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_float, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stores a floating-point number as 64 bits in the file.
*/
//go:nosplit
func (self class) StoreDouble(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_double, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stores a floating-point number in the file.
*/
//go:nosplit
func (self class) StoreReal(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_real, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stores the given array of bytes in the file.
*/
//go:nosplit
func (self class) StoreBuffer(buffer gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Appends [param line] to the file followed by a line return character ([code]\n[/code]), encoding the text as UTF-8.
*/
//go:nosplit
func (self class) StoreLine(line gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(line))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Store the given [PackedStringArray] in the file as a line formatted in the CSV (Comma-Separated Values) format. You can pass a different delimiter [param delim] to use other than the default [code]","[/code] (comma). This delimiter must be one-character long.
Text will be encoded as UTF-8.
*/
//go:nosplit
func (self class) StoreCsvLine(values gd.PackedStringArray, delim gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(values))
	callframe.Arg(frame, mmm.Get(delim))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_csv_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Appends [param string] to the file without a line return, encoding the text as UTF-8.
[b]Note:[/b] This method is intended to be used to write text files. The string is stored as a UTF-8 encoded buffer without string length or terminating zero, which means that it can't be loaded back easily. If you want to store a retrievable string in a binary file, consider using [method store_pascal_string] instead. For retrieving strings from a text file, you can use [code]get_buffer(length).get_string_from_utf8()[/code] (if you know the length) or [method get_as_text].
*/
//go:nosplit
func (self class) StoreString(s gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stores any Variant value in the file. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
[b]Note:[/b] Not all properties are included. Only properties that are configured with the [constant PROPERTY_USAGE_STORAGE] flag set will be serialized. You can add a new usage flag to a property by overriding the [method Object._get_property_list] method in your class. You can also check how property usage is configured by calling [method Object._get_property_list]. See [enum PropertyUsageFlags] for the possible usage flags.
*/
//go:nosplit
func (self class) StoreVar(value gd.Variant, full_objects bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(value))
	callframe.Arg(frame, full_objects)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_var, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stores the given [String] as a line in the file in Pascal format (i.e. also store the length of the string).
Text will be encoded as UTF-8.
*/
//go:nosplit
func (self class) StorePascalString(s gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_store_pascal_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [String] saved in Pascal format from the file.
Text is interpreted as being UTF-8 encoded.
*/
//go:nosplit
func (self class) GetPascalString(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_get_pascal_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Closes the currently opened file and prevents subsequent read/write operations. Use [method flush] to persist the data to disk without closing the file.
[b]Note:[/b] [FileAccess] will automatically close when it's freed, which happens when it goes out of scope or when it gets assigned with [code]null[/code]. In C# the reference must be disposed after we are done using it, this can be done with the [code]using[/code] statement or calling the [code]Dispose[/code] method directly.
*/
//go:nosplit
func (self class) Close()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FileAccess.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the file exists in the given path.
[b]Note:[/b] Many resources types are imported (e.g. textures or sound files), and their source asset will not be included in the exported game, as only the imported version is used. See [method ResourceLoader.exists] for an alternative approach that takes resource remapping into account.
For a non-static, relative equivalent, use [method DirAccess.file_exists].
*/
//go:nosplit
func (self class) FileExists(ctx gd.Lifetime, path gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_file_exists, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the last time the [param file] was modified in Unix timestamp format, or [code]0[/code] on error. This Unix timestamp can be converted to another format using the [Time] singleton.
*/
//go:nosplit
func (self class) GetModifiedTime(ctx gd.Lifetime, file gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	var r_ret = callframe.Ret[gd.Int](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_get_modified_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns file UNIX permissions.
[b]Note:[/b] This method is implemented on iOS, Linux/BSD, and macOS.
*/
//go:nosplit
func (self class) GetUnixPermissions(ctx gd.Lifetime, file gd.String) classdb.FileAccessUnixPermissionFlags {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	var r_ret = callframe.Ret[classdb.FileAccessUnixPermissionFlags](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_get_unix_permissions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets file UNIX permissions.
[b]Note:[/b] This method is implemented on iOS, Linux/BSD, and macOS.
*/
//go:nosplit
func (self class) SetUnixPermissions(ctx gd.Lifetime, file gd.String, permissions classdb.FileAccessUnixPermissionFlags) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	callframe.Arg(frame, permissions)
	var r_ret = callframe.Ret[int64](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_set_unix_permissions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code], if file [code]hidden[/code] attribute is set.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
//go:nosplit
func (self class) GetHiddenAttribute(ctx gd.Lifetime, file gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	var r_ret = callframe.Ret[bool](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_get_hidden_attribute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets file [b]hidden[/b] attribute.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
//go:nosplit
func (self class) SetHiddenAttribute(ctx gd.Lifetime, file gd.String, hidden bool) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Ret[int64](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_set_hidden_attribute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets file [b]read only[/b] attribute.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
//go:nosplit
func (self class) SetReadOnlyAttribute(ctx gd.Lifetime, file gd.String, ro bool) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	callframe.Arg(frame, ro)
	var r_ret = callframe.Ret[int64](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_set_read_only_attribute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code], if file [code]read only[/code] attribute is set.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
//go:nosplit
func (self class) GetReadOnlyAttribute(ctx gd.Lifetime, file gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	var r_ret = callframe.Ret[bool](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.FileAccess.Bind_get_read_only_attribute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsFileAccess() Expert { return self[0].AsFileAccess() }


//go:nosplit
func (self Simple) AsFileAccess() Simple { return self[0].AsFileAccess() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("FileAccess", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ModeFlags = classdb.FileAccessModeFlags

const (
/*Opens the file for read operations. The cursor is positioned at the beginning of the file.*/
	Read ModeFlags = 1
/*Opens the file for write operations. The file is created if it does not exist, and truncated if it does.
[b]Note:[/b] When creating a file it must be in an already existing directory. To recursively create directories for a file path, see [method DirAccess.make_dir_recursive].*/
	Write ModeFlags = 2
/*Opens the file for read and write operations. Does not truncate the file. The cursor is positioned at the beginning of the file.*/
	ReadWrite ModeFlags = 3
/*Opens the file for read and write operations. The file is created if it does not exist, and truncated if it does. The cursor is positioned at the beginning of the file.
[b]Note:[/b] When creating a file it must be in an already existing directory. To recursively create directories for a file path, see [method DirAccess.make_dir_recursive].*/
	WriteRead ModeFlags = 7
)
type CompressionMode = classdb.FileAccessCompressionMode

const (
/*Uses the [url=https://fastlz.org/]FastLZ[/url] compression method.*/
	CompressionFastlz CompressionMode = 0
/*Uses the [url=https://en.wikipedia.org/wiki/DEFLATE]DEFLATE[/url] compression method.*/
	CompressionDeflate CompressionMode = 1
/*Uses the [url=https://facebook.github.io/zstd/]Zstandard[/url] compression method.*/
	CompressionZstd CompressionMode = 2
/*Uses the [url=https://www.gzip.org/]gzip[/url] compression method.*/
	CompressionGzip CompressionMode = 3
/*Uses the [url=https://github.com/google/brotli]brotli[/url] compression method (only decompression is supported).*/
	CompressionBrotli CompressionMode = 4
)
type UnixPermissionFlags = classdb.FileAccessUnixPermissionFlags

const (
/*Read for owner bit.*/
	UnixReadOwner UnixPermissionFlags = 256
/*Write for owner bit.*/
	UnixWriteOwner UnixPermissionFlags = 128
/*Execute for owner bit.*/
	UnixExecuteOwner UnixPermissionFlags = 64
/*Read for group bit.*/
	UnixReadGroup UnixPermissionFlags = 32
/*Write for group bit.*/
	UnixWriteGroup UnixPermissionFlags = 16
/*Execute for group bit.*/
	UnixExecuteGroup UnixPermissionFlags = 8
/*Read for other bit.*/
	UnixReadOther UnixPermissionFlags = 4
/*Write for other bit.*/
	UnixWriteOther UnixPermissionFlags = 2
/*Execute for other bit.*/
	UnixExecuteOther UnixPermissionFlags = 1
/*Set user id on execution bit.*/
	UnixSetUserId UnixPermissionFlags = 2048
/*Set group id on execution bit.*/
	UnixSetGroupId UnixPermissionFlags = 1024
/*Restricted deletion (sticky) bit.*/
	UnixRestrictedDelete UnixPermissionFlags = 512
)
