// Package XMLParser provides methods for working with XMLParser object instances.
package XMLParser

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Provides a low-level interface for creating parsers for [url=https://en.wikipedia.org/wiki/XML]XML[/url] files. This class can serve as base to make custom XML parsers.
To parse XML, you must open a file with the [method open] method or a buffer with the [method open_buffer] method. Then, the [method read] method must be called to parse the next nodes. Most of the methods take into consideration the currently parsed node.
Here is an example of using [XMLParser] to parse an SVG file (which is based on XML), printing each element and its attributes as a dictionary:
[codeblocks]
[gdscript]
var parser = XMLParser.new()
parser.open("path/to/file.svg")
while parser.read() != ERR_FILE_EOF:

	if parser.get_node_type() == XMLParser.NODE_ELEMENT:
	    var node_name = parser.get_node_name()
	    var attributes_dict = {}
	    for idx in range(parser.get_attribute_count()):
	        attributes_dict[parser.get_attribute_name(idx)] = parser.get_attribute_value(idx)
	    print("The ", node_name, " element has the following attributes: ", attributes_dict)

[/gdscript]
[csharp]
var parser = new XmlParser();
parser.Open("path/to/file.svg");
while (parser.Read() != Error.FileEof)

	{
	    if (parser.GetNodeType() == XmlParser.NodeType.Element)
	    {
	        var nodeName = parser.GetNodeName();
	        var attributesDict = new Godot.Collections.Dictionary();
	        for (int idx = 0; idx < parser.GetAttributeCount(); idx++)
	        {
	            attributesDict[parser.GetAttributeName(idx)] = parser.GetAttributeValue(idx);
	        }
	        GD.Print($"The {nodeName} element has the following attributes: {attributesDict}");
	    }
	}

[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.XMLParser

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsXMLParser() Instance
}

/*
Parses the next node in the file. This method returns an error code.
*/
func (self Instance) Read() error {
	return error(class(self).Read())
}

/*
Returns the type of the current node. Compare with [enum NodeType] constants.
*/
func (self Instance) GetNodeType() gdclass.XMLParserNodeType {
	return gdclass.XMLParserNodeType(class(self).GetNodeType())
}

/*
Returns the name of a node. This method will raise an error if the currently parsed node is a text node.
[b]Note:[/b] The content of a [constant NODE_CDATA] node and the comment string of a [constant NODE_COMMENT] node are also considered names.
*/
func (self Instance) GetNodeName() string {
	return string(class(self).GetNodeName().String())
}

/*
Returns the contents of a text node. This method will raise an error if the current parsed node is of any other type.
*/
func (self Instance) GetNodeData() string {
	return string(class(self).GetNodeData().String())
}

/*
Returns the byte offset of the currently parsed node since the beginning of the file or buffer. This is usually equivalent to the number of characters before the read position.
*/
func (self Instance) GetNodeOffset() int {
	return int(int(class(self).GetNodeOffset()))
}

/*
Returns the number of attributes in the currently parsed element.
[b]Note:[/b] If this method is used while the currently parsed node is not [constant NODE_ELEMENT] or [constant NODE_ELEMENT_END], this count will not be updated and will still reflect the last element.
*/
func (self Instance) GetAttributeCount() int {
	return int(int(class(self).GetAttributeCount()))
}

/*
Returns the name of an attribute of the currently parsed element, specified by the [param idx] index.
*/
func (self Instance) GetAttributeName(idx int) string {
	return string(class(self).GetAttributeName(gd.Int(idx)).String())
}

/*
Returns the value of an attribute of the currently parsed element, specified by the [param idx] index.
*/
func (self Instance) GetAttributeValue(idx int) string {
	return string(class(self).GetAttributeValue(gd.Int(idx)).String())
}

/*
Returns [code]true[/code] if the currently parsed element has an attribute with the [param name].
*/
func (self Instance) HasAttribute(name string) bool {
	return bool(class(self).HasAttribute(gd.NewString(name)))
}

/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will raise an error if the element has no such attribute.
*/
func (self Instance) GetNamedAttributeValue(name string) string {
	return string(class(self).GetNamedAttributeValue(gd.NewString(name)).String())
}

/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will return an empty string if the element has no such attribute.
*/
func (self Instance) GetNamedAttributeValueSafe(name string) string {
	return string(class(self).GetNamedAttributeValueSafe(gd.NewString(name)).String())
}

/*
Returns [code]true[/code] if the currently parsed element is empty, e.g. [code]<element />[/code].
*/
func (self Instance) IsEmpty() bool {
	return bool(class(self).IsEmpty())
}

/*
Returns the current line in the parsed file, counting from 0.
*/
func (self Instance) GetCurrentLine() int {
	return int(int(class(self).GetCurrentLine()))
}

/*
Skips the current section. If the currently parsed node contains more inner nodes, they will be ignored and the cursor will go to the closing of the current element.
*/
func (self Instance) SkipSection() {
	class(self).SkipSection()
}

/*
Moves the buffer cursor to a certain offset (since the beginning) and reads the next node there. This method returns an error code.
*/
func (self Instance) SeekTo(position int) error {
	return error(class(self).SeekTo(gd.Int(position)))
}

/*
Opens an XML [param file] for parsing. This method returns an error code.
*/
func (self Instance) Open(file string) error {
	return error(class(self).Open(gd.NewString(file)))
}

/*
Opens an XML raw [param buffer] for parsing. This method returns an error code.
*/
func (self Instance) OpenBuffer(buffer []byte) error {
	return error(class(self).OpenBuffer(gd.NewPackedByteSlice(buffer)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XMLParser

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XMLParser"))
	casted := Instance{*(*gdclass.XMLParser)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Parses the next node in the file. This method returns an error code.
*/
//go:nosplit
func (self class) Read() error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_read, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the type of the current node. Compare with [enum NodeType] constants.
*/
//go:nosplit
func (self class) GetNodeType() gdclass.XMLParserNodeType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.XMLParserNodeType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_node_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of a node. This method will raise an error if the currently parsed node is a text node.
[b]Note:[/b] The content of a [constant NODE_CDATA] node and the comment string of a [constant NODE_COMMENT] node are also considered names.
*/
//go:nosplit
func (self class) GetNodeName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_node_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the contents of a text node. This method will raise an error if the current parsed node is of any other type.
*/
//go:nosplit
func (self class) GetNodeData() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_node_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the byte offset of the currently parsed node since the beginning of the file or buffer. This is usually equivalent to the number of characters before the read position.
*/
//go:nosplit
func (self class) GetNodeOffset() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_node_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of attributes in the currently parsed element.
[b]Note:[/b] If this method is used while the currently parsed node is not [constant NODE_ELEMENT] or [constant NODE_ELEMENT_END], this count will not be updated and will still reflect the last element.
*/
//go:nosplit
func (self class) GetAttributeCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_attribute_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of an attribute of the currently parsed element, specified by the [param idx] index.
*/
//go:nosplit
func (self class) GetAttributeName(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_attribute_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the value of an attribute of the currently parsed element, specified by the [param idx] index.
*/
//go:nosplit
func (self class) GetAttributeValue(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_attribute_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the currently parsed element has an attribute with the [param name].
*/
//go:nosplit
func (self class) HasAttribute(name gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_has_attribute, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will raise an error if the element has no such attribute.
*/
//go:nosplit
func (self class) GetNamedAttributeValue(name gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_named_attribute_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will return an empty string if the element has no such attribute.
*/
//go:nosplit
func (self class) GetNamedAttributeValueSafe(name gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_named_attribute_value_safe, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the currently parsed element is empty, e.g. [code]<element />[/code].
*/
//go:nosplit
func (self class) IsEmpty() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_is_empty, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current line in the parsed file, counting from 0.
*/
//go:nosplit
func (self class) GetCurrentLine() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_current_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Skips the current section. If the currently parsed node contains more inner nodes, they will be ignored and the cursor will go to the closing of the current element.
*/
//go:nosplit
func (self class) SkipSection() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_skip_section, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves the buffer cursor to a certain offset (since the beginning) and reads the next node there. This method returns an error code.
*/
//go:nosplit
func (self class) SeekTo(position gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Opens an XML [param file] for parsing. This method returns an error code.
*/
//go:nosplit
func (self class) Open(file gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(file))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_open, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Opens an XML raw [param buffer] for parsing. This method returns an error code.
*/
//go:nosplit
func (self class) OpenBuffer(buffer gd.PackedByteArray) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_open_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXMLParser() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXMLParser() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("XMLParser", func(ptr gd.Object) any { return [1]gdclass.XMLParser{*(*gdclass.XMLParser)(unsafe.Pointer(&ptr))} })
}

type NodeType = gdclass.XMLParserNodeType

const (
	/*There's no node (no file or buffer opened).*/
	NodeNone NodeType = 0
	/*An element node type, also known as a tag, e.g. [code]<title>[/code].*/
	NodeElement NodeType = 1
	/*An end of element node type, e.g. [code]</title>[/code].*/
	NodeElementEnd NodeType = 2
	/*A text node type, i.e. text that is not inside an element. This includes whitespace.*/
	NodeText NodeType = 3
	/*A comment node type, e.g. [code]<!--A comment-->[/code].*/
	NodeComment NodeType = 4
	/*A node type for CDATA (Character Data) sections, e.g. [code]<![CDATA[CDATA section]]>[/code].*/
	NodeCdata NodeType = 5
	/*An unknown node type.*/
	NodeUnknown NodeType = 6
)

type Error int

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
