// Package ConfigFile provides methods for working with ConfigFile object instances.
package ConfigFile

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
type Instance [1]gdclass.ConfigFile

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsConfigFile() Instance
}

/*
Assigns a value to the specified key of the specified section. If either the section or the key do not exist, they are created. Passing a [code]null[/code] value deletes the specified key if it exists, and deletes the section if it ends up empty once the key has been removed.
*/
func (self Instance) SetValue(section string, key string, value any) { //gd:ConfigFile.set_value
	class(self).SetValue(String.New(section), String.New(key), variant.New(value))
}

/*
Returns the current value for the specified section and key. If either the section or the key do not exist, the method returns the fallback [param default] value. If [param default] is not specified or set to [code]null[/code], an error is also raised.
*/
func (self Instance) GetValue(section string, key string) any { //gd:ConfigFile.get_value
	return any(class(self).GetValue(String.New(section), String.New(key), variant.New([1]any{}[0])).Interface())
}

/*
Returns [code]true[/code] if the specified section exists.
*/
func (self Instance) HasSection(section string) bool { //gd:ConfigFile.has_section
	return bool(class(self).HasSection(String.New(section)))
}

/*
Returns [code]true[/code] if the specified section-key pair exists.
*/
func (self Instance) HasSectionKey(section string, key string) bool { //gd:ConfigFile.has_section_key
	return bool(class(self).HasSectionKey(String.New(section), String.New(key)))
}

/*
Returns an array of all defined section identifiers.
*/
func (self Instance) GetSections() []string { //gd:ConfigFile.get_sections
	return []string(class(self).GetSections().Strings())
}

/*
Returns an array of all defined key identifiers in the specified section. Raises an error and returns an empty array if the section does not exist.
*/
func (self Instance) GetSectionKeys(section string) []string { //gd:ConfigFile.get_section_keys
	return []string(class(self).GetSectionKeys(String.New(section)).Strings())
}

/*
Deletes the specified section along with all the key-value pairs inside. Raises an error if the section does not exist.
*/
func (self Instance) EraseSection(section string) { //gd:ConfigFile.erase_section
	class(self).EraseSection(String.New(section))
}

/*
Deletes the specified key in a section. Raises an error if either the section or the key do not exist.
*/
func (self Instance) EraseSectionKey(section string, key string) { //gd:ConfigFile.erase_section_key
	class(self).EraseSectionKey(String.New(section), String.New(key))
}

/*
Loads the config file specified as a parameter. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) Load(path string) error { //gd:ConfigFile.load
	return error(gd.ToError(class(self).Load(String.New(path))))
}

/*
Parses the passed string as the contents of a config file. The string is parsed and loaded in the ConfigFile object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) Parse(data string) error { //gd:ConfigFile.parse
	return error(gd.ToError(class(self).Parse(String.New(data))))
}

/*
Saves the contents of the [ConfigFile] object to the file specified as a parameter. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) Save(path string) error { //gd:ConfigFile.save
	return error(gd.ToError(class(self).Save(String.New(path))))
}

/*
Obtain the text version of this config file (the same text that would be written to a file).
*/
func (self Instance) EncodeToText() string { //gd:ConfigFile.encode_to_text
	return string(class(self).EncodeToText().String())
}

/*
Loads the encrypted config file specified as a parameter, using the provided [param key] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) LoadEncrypted(path string, key []byte) error { //gd:ConfigFile.load_encrypted
	return error(gd.ToError(class(self).LoadEncrypted(String.New(path), Packed.Bytes(Packed.New(key...)))))
}

/*
Loads the encrypted config file specified as a parameter, using the provided [param password] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) LoadEncryptedPass(path string, password string) error { //gd:ConfigFile.load_encrypted_pass
	return error(gd.ToError(class(self).LoadEncryptedPass(String.New(path), String.New(password))))
}

/*
Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [param key] to encrypt it. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) SaveEncrypted(path string, key []byte) error { //gd:ConfigFile.save_encrypted
	return error(gd.ToError(class(self).SaveEncrypted(String.New(path), Packed.Bytes(Packed.New(key...)))))
}

/*
Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [param password] to encrypt it. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) SaveEncryptedPass(path string, password string) error { //gd:ConfigFile.save_encrypted_pass
	return error(gd.ToError(class(self).SaveEncryptedPass(String.New(path), String.New(password))))
}

/*
Removes the entire contents of the config.
*/
func (self Instance) Clear() { //gd:ConfigFile.clear
	class(self).Clear()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ConfigFile

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ConfigFile"))
	casted := Instance{*(*gdclass.ConfigFile)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Assigns a value to the specified key of the specified section. If either the section or the key do not exist, they are created. Passing a [code]null[/code] value deletes the specified key if it exists, and deletes the section if it ends up empty once the key has been removed.
*/
//go:nosplit
func (self class) SetValue(section String.Readable, key String.Readable, value variant.Any) { //gd:ConfigFile.set_value
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(section)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(key)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_set_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current value for the specified section and key. If either the section or the key do not exist, the method returns the fallback [param default] value. If [param default] is not specified or set to [code]null[/code], an error is also raised.
*/
//go:nosplit
func (self class) GetValue(section String.Readable, key String.Readable, def variant.Any) variant.Any { //gd:ConfigFile.get_value
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(section)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(key)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(def)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_get_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the specified section exists.
*/
//go:nosplit
func (self class) HasSection(section String.Readable) bool { //gd:ConfigFile.has_section
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(section)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_has_section, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the specified section-key pair exists.
*/
//go:nosplit
func (self class) HasSectionKey(section String.Readable, key String.Readable) bool { //gd:ConfigFile.has_section_key
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(section)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(key)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_has_section_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of all defined section identifiers.
*/
//go:nosplit
func (self class) GetSections() Packed.Strings { //gd:ConfigFile.get_sections
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_get_sections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns an array of all defined key identifiers in the specified section. Raises an error and returns an empty array if the section does not exist.
*/
//go:nosplit
func (self class) GetSectionKeys(section String.Readable) Packed.Strings { //gd:ConfigFile.get_section_keys
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(section)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_get_section_keys, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Deletes the specified section along with all the key-value pairs inside. Raises an error if the section does not exist.
*/
//go:nosplit
func (self class) EraseSection(section String.Readable) { //gd:ConfigFile.erase_section
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(section)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_erase_section, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Deletes the specified key in a section. Raises an error if either the section or the key do not exist.
*/
//go:nosplit
func (self class) EraseSectionKey(section String.Readable, key String.Readable) { //gd:ConfigFile.erase_section_key
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(section)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(key)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_erase_section_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Loads the config file specified as a parameter. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) Load(path String.Readable) Error.Code { //gd:ConfigFile.load
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_load, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Parses the passed string as the contents of a config file. The string is parsed and loaded in the ConfigFile object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) Parse(data String.Readable) Error.Code { //gd:ConfigFile.parse
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(data)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_parse, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Saves the contents of the [ConfigFile] object to the file specified as a parameter. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) Save(path String.Readable) Error.Code { //gd:ConfigFile.save
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_save, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Obtain the text version of this config file (the same text that would be written to a file).
*/
//go:nosplit
func (self class) EncodeToText() String.Readable { //gd:ConfigFile.encode_to_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_encode_to_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Loads the encrypted config file specified as a parameter, using the provided [param key] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) LoadEncrypted(path String.Readable, key Packed.Bytes) Error.Code { //gd:ConfigFile.load_encrypted
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](key))))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_load_encrypted, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Loads the encrypted config file specified as a parameter, using the provided [param password] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) LoadEncryptedPass(path String.Readable, password String.Readable) Error.Code { //gd:ConfigFile.load_encrypted_pass
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(password)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_load_encrypted_pass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [param key] to encrypt it. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) SaveEncrypted(path String.Readable, key Packed.Bytes) Error.Code { //gd:ConfigFile.save_encrypted
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](key))))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_save_encrypted, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [param password] to encrypt it. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) SaveEncryptedPass(path String.Readable, password String.Readable) Error.Code { //gd:ConfigFile.save_encrypted_pass
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(password)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_save_encrypted_pass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes the entire contents of the config.
*/
//go:nosplit
func (self class) Clear() { //gd:ConfigFile.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsConfigFile() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsConfigFile() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("ConfigFile", func(ptr gd.Object) any { return [1]gdclass.ConfigFile{*(*gdclass.ConfigFile)(unsafe.Pointer(&ptr))} })
}
