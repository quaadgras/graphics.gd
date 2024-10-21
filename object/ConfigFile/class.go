package ConfigFile

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
This helper class can be used to store [Variant] values on the filesystem using INI-style formatting. The stored values are identified by a section and a key:
[codeblock lang=text]
[section]
some_key=42
string_example="Hello World3D!"
a_vector=Vector3(1, 0, 2)
[/codeblock]
The stored data can be saved to or parsed from a file, though ConfigFile objects can also be used directly without accessing the filesystem.
The following example shows how to create a simple [ConfigFile] and save it on disc:
[codeblocks]
[gdscript]
# Create new ConfigFile object.
var config = ConfigFile.new()

# Store some values.
config.set_value("Player1", "player_name", "Steve")
config.set_value("Player1", "best_score", 10)
config.set_value("Player2", "player_name", "V3geta")
config.set_value("Player2", "best_score", 9001)

# Save it to a file (overwrite if already exists).
config.save("user://scores.cfg")
[/gdscript]
[csharp]
// Create new ConfigFile object.
var config = new ConfigFile();

// Store some values.
config.SetValue("Player1", "player_name", "Steve");
config.SetValue("Player1", "best_score", 10);
config.SetValue("Player2", "player_name", "V3geta");
config.SetValue("Player2", "best_score", 9001);

// Save it to a file (overwrite if already exists).
config.Save("user://scores.cfg");
[/csharp]
[/codeblocks]
This example shows how the above file could be loaded:
[codeblocks]
[gdscript]
var score_data = {}
var config = ConfigFile.new()

# Load data from a file.
var err = config.load("user://scores.cfg")

# If the file didn't load, ignore it.
if err != OK:
    return

# Iterate over all sections.
for player in config.get_sections():
    # Fetch the data for each section.
    var player_name = config.get_value(player, "player_name")
    var player_score = config.get_value(player, "best_score")
    score_data[player_name] = player_score
[/gdscript]
[csharp]
var score_data = new Godot.Collections.Dictionary();
var config = new ConfigFile();

// Load data from a file.
Error err = config.Load("user://scores.cfg");

// If the file didn't load, ignore it.
if (err != Error.Ok)
{
    return;
}

// Iterate over all sections.
foreach (String player in config.GetSections())
{
    // Fetch the data for each section.
    var player_name = (String)config.GetValue(player, "player_name");
    var player_score = (int)config.GetValue(player, "best_score");
    score_data[player_name] = player_score;
}
[/csharp]
[/codeblocks]
Any operation that mutates the ConfigFile such as [method set_value], [method clear], or [method erase_section], only changes what is loaded in memory. If you want to write the change to a file, you have to save the changes with [method save], [method save_encrypted], or [method save_encrypted_pass].
Keep in mind that section and property names can't contain spaces. Anything after a space will be ignored on save and on load.
ConfigFiles can also contain manually written comment lines starting with a semicolon ([code];[/code]). Those lines will be ignored when parsing the file. Note that comments will be lost when saving the ConfigFile. This can still be useful for dedicated server configuration files, which are typically never overwritten without explicit user action.
[b]Note:[/b] The file extension given to a ConfigFile does not have any impact on its formatting or behavior. By convention, the [code].cfg[/code] extension is used here, but any other extension such as [code].ini[/code] is also valid. Since neither [code].cfg[/code] nor [code].ini[/code] are standardized, Godot's ConfigFile formatting may differ from files written by other programs.

*/
type Simple [1]classdb.ConfigFile
func (self Simple) SetValue(section string, key string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetValue(gc.String(section), gc.String(key), value)
}
func (self Simple) GetValue(section string, key string, def gd.Variant) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetValue(gc, gc.String(section), gc.String(key), def))
}
func (self Simple) HasSection(section string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasSection(gc.String(section)))
}
func (self Simple) HasSectionKey(section string, key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasSectionKey(gc.String(section), gc.String(key)))
}
func (self Simple) GetSections() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetSections(gc))
}
func (self Simple) GetSectionKeys(section string) gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetSectionKeys(gc, gc.String(section)))
}
func (self Simple) EraseSection(section string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EraseSection(gc.String(section))
}
func (self Simple) EraseSectionKey(section string, key string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EraseSectionKey(gc.String(section), gc.String(key))
}
func (self Simple) Load(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Load(gc.String(path)))
}
func (self Simple) Parse(data string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Parse(gc.String(data)))
}
func (self Simple) Save(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Save(gc.String(path)))
}
func (self Simple) EncodeToText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).EncodeToText(gc).String())
}
func (self Simple) LoadEncrypted(path string, key []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadEncrypted(gc.String(path), gc.PackedByteSlice(key)))
}
func (self Simple) LoadEncryptedPass(path string, password string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadEncryptedPass(gc.String(path), gc.String(password)))
}
func (self Simple) SaveEncrypted(path string, key []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SaveEncrypted(gc.String(path), gc.PackedByteSlice(key)))
}
func (self Simple) SaveEncryptedPass(path string, password string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SaveEncryptedPass(gc.String(path), gc.String(password)))
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ConfigFile
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Assigns a value to the specified key of the specified section. If either the section or the key do not exist, they are created. Passing a [code]null[/code] value deletes the specified key if it exists, and deletes the section if it ends up empty once the key has been removed.
*/
//go:nosplit
func (self class) SetValue(section gd.String, key gd.String, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(section))
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_set_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current value for the specified section and key. If either the section or the key do not exist, the method returns the fallback [param default] value. If [param default] is not specified or set to [code]null[/code], an error is also raised.
*/
//go:nosplit
func (self class) GetValue(ctx gd.Lifetime, section gd.String, key gd.String, def gd.Variant) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(section))
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, mmm.Get(def))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_get_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified section exists.
*/
//go:nosplit
func (self class) HasSection(section gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(section))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_has_section, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified section-key pair exists.
*/
//go:nosplit
func (self class) HasSectionKey(section gd.String, key gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(section))
	callframe.Arg(frame, mmm.Get(key))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_has_section_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of all defined section identifiers.
*/
//go:nosplit
func (self class) GetSections(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_get_sections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of all defined key identifiers in the specified section. Raises an error and returns an empty array if the section does not exist.
*/
//go:nosplit
func (self class) GetSectionKeys(ctx gd.Lifetime, section gd.String) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(section))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_get_section_keys, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Deletes the specified section along with all the key-value pairs inside. Raises an error if the section does not exist.
*/
//go:nosplit
func (self class) EraseSection(section gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(section))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_erase_section, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deletes the specified key in a section. Raises an error if either the section or the key do not exist.
*/
//go:nosplit
func (self class) EraseSectionKey(section gd.String, key gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(section))
	callframe.Arg(frame, mmm.Get(key))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_erase_section_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Loads the config file specified as a parameter. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) Load(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_load, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Parses the passed string as the contents of a config file. The string is parsed and loaded in the ConfigFile object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) Parse(data gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_parse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Saves the contents of the [ConfigFile] object to the file specified as a parameter. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) Save(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_save, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Obtain the text version of this config file (the same text that would be written to a file).
*/
//go:nosplit
func (self class) EncodeToText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_encode_to_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Loads the encrypted config file specified as a parameter, using the provided [param key] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) LoadEncrypted(path gd.String, key gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(key))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_load_encrypted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Loads the encrypted config file specified as a parameter, using the provided [param password] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) LoadEncryptedPass(path gd.String, password gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(password))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_load_encrypted_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [param key] to encrypt it. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) SaveEncrypted(path gd.String, key gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(key))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_save_encrypted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [param password] to encrypt it. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) SaveEncryptedPass(path gd.String, password gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(password))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_save_encrypted_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the entire contents of the config.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConfigFile.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsConfigFile() Expert { return self[0].AsConfigFile() }


//go:nosplit
func (self Simple) AsConfigFile() Simple { return self[0].AsConfigFile() }


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
func init() {classdb.Register("ConfigFile", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
