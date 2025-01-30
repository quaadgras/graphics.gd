// Package XMLParser provides methods for working with XMLParser object instances.
package XMLParser

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
func (self Instance) Read() error { //gd:XMLParser.read
	return error(gd.ToError(class(self).Read()))
}

/*
Returns the type of the current node. Compare with [enum NodeType] constants.
*/
func (self Instance) GetNodeType() gdclass.XMLParserNodeType { //gd:XMLParser.get_node_type
	return gdclass.XMLParserNodeType(class(self).GetNodeType())
}

/*
Returns the name of a node. This method will raise an error if the currently parsed node is a text node.
[b]Note:[/b] The content of a [constant NODE_CDATA] node and the comment string of a [constant NODE_COMMENT] node are also considered names.
*/
func (self Instance) GetNodeName() string { //gd:XMLParser.get_node_name
	return string(class(self).GetNodeName().String())
}

/*
Returns the contents of a text node. This method will raise an error if the current parsed node is of any other type.
*/
func (self Instance) GetNodeData() string { //gd:XMLParser.get_node_data
	return string(class(self).GetNodeData().String())
}

/*
Returns the byte offset of the currently parsed node since the beginning of the file or buffer. This is usually equivalent to the number of characters before the read position.
*/
func (self Instance) GetNodeOffset() int { //gd:XMLParser.get_node_offset
	return int(int(class(self).GetNodeOffset()))
}

/*
Returns the number of attributes in the currently parsed element.
[b]Note:[/b] If this method is used while the currently parsed node is not [constant NODE_ELEMENT] or [constant NODE_ELEMENT_END], this count will not be updated and will still reflect the last element.
*/
func (self Instance) GetAttributeCount() int { //gd:XMLParser.get_attribute_count
	return int(int(class(self).GetAttributeCount()))
}

/*
Returns the name of an attribute of the currently parsed element, specified by the [param idx] index.
*/
func (self Instance) GetAttributeName(idx int) string { //gd:XMLParser.get_attribute_name
	return string(class(self).GetAttributeName(int64(idx)).String())
}

/*
Returns the value of an attribute of the currently parsed element, specified by the [param idx] index.
*/
func (self Instance) GetAttributeValue(idx int) string { //gd:XMLParser.get_attribute_value
	return string(class(self).GetAttributeValue(int64(idx)).String())
}

/*
Returns [code]true[/code] if the currently parsed element has an attribute with the [param name].
*/
func (self Instance) HasAttribute(name string) bool { //gd:XMLParser.has_attribute
	return bool(class(self).HasAttribute(String.New(name)))
}

/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will raise an error if the element has no such attribute.
*/
func (self Instance) GetNamedAttributeValue(name string) string { //gd:XMLParser.get_named_attribute_value
	return string(class(self).GetNamedAttributeValue(String.New(name)).String())
}

/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will return an empty string if the element has no such attribute.
*/
func (self Instance) GetNamedAttributeValueSafe(name string) string { //gd:XMLParser.get_named_attribute_value_safe
	return string(class(self).GetNamedAttributeValueSafe(String.New(name)).String())
}

/*
Returns [code]true[/code] if the currently parsed element is empty, e.g. [code]<element />[/code].
*/
func (self Instance) IsEmpty() bool { //gd:XMLParser.is_empty
	return bool(class(self).IsEmpty())
}

/*
Returns the current line in the parsed file, counting from 0.
*/
func (self Instance) GetCurrentLine() int { //gd:XMLParser.get_current_line
	return int(int(class(self).GetCurrentLine()))
}

/*
Skips the current section. If the currently parsed node contains more inner nodes, they will be ignored and the cursor will go to the closing of the current element.
*/
func (self Instance) SkipSection() { //gd:XMLParser.skip_section
	class(self).SkipSection()
}

/*
Moves the buffer cursor to a certain offset (since the beginning) and reads the next node there. This method returns an error code.
*/
func (self Instance) SeekTo(position int) error { //gd:XMLParser.seek
	return error(gd.ToError(class(self).SeekTo(int64(position))))
}

/*
Opens an XML [param file] for parsing. This method returns an error code.
*/
func (self Instance) Open(file string) error { //gd:XMLParser.open
	return error(gd.ToError(class(self).Open(String.New(file))))
}

/*
Opens an XML raw [param buffer] for parsing. This method returns an error code.
*/
func (self Instance) OpenBuffer(buffer []byte) error { //gd:XMLParser.open_buffer
	return error(gd.ToError(class(self).OpenBuffer(Packed.Bytes(Packed.New(buffer...)))))
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
func (self class) Read() Error.Code { //gd:XMLParser.read
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_read, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the type of the current node. Compare with [enum NodeType] constants.
*/
//go:nosplit
func (self class) GetNodeType() gdclass.XMLParserNodeType { //gd:XMLParser.get_node_type
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
func (self class) GetNodeName() String.Readable { //gd:XMLParser.get_node_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_node_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the contents of a text node. This method will raise an error if the current parsed node is of any other type.
*/
//go:nosplit
func (self class) GetNodeData() String.Readable { //gd:XMLParser.get_node_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_node_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the byte offset of the currently parsed node since the beginning of the file or buffer. This is usually equivalent to the number of characters before the read position.
*/
//go:nosplit
func (self class) GetNodeOffset() int64 { //gd:XMLParser.get_node_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GetAttributeCount() int64 { //gd:XMLParser.get_attribute_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_attribute_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of an attribute of the currently parsed element, specified by the [param idx] index.
*/
//go:nosplit
func (self class) GetAttributeName(idx int64) String.Readable { //gd:XMLParser.get_attribute_name
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_attribute_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the value of an attribute of the currently parsed element, specified by the [param idx] index.
*/
//go:nosplit
func (self class) GetAttributeValue(idx int64) String.Readable { //gd:XMLParser.get_attribute_value
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_attribute_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the currently parsed element has an attribute with the [param name].
*/
//go:nosplit
func (self class) HasAttribute(name String.Readable) bool { //gd:XMLParser.has_attribute
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
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
func (self class) GetNamedAttributeValue(name String.Readable) String.Readable { //gd:XMLParser.get_named_attribute_value
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_named_attribute_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will return an empty string if the element has no such attribute.
*/
//go:nosplit
func (self class) GetNamedAttributeValueSafe(name String.Readable) String.Readable { //gd:XMLParser.get_named_attribute_value_safe
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_named_attribute_value_safe, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the currently parsed element is empty, e.g. [code]<element />[/code].
*/
//go:nosplit
func (self class) IsEmpty() bool { //gd:XMLParser.is_empty
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
func (self class) GetCurrentLine() int64 { //gd:XMLParser.get_current_line
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_get_current_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Skips the current section. If the currently parsed node contains more inner nodes, they will be ignored and the cursor will go to the closing of the current element.
*/
//go:nosplit
func (self class) SkipSection() { //gd:XMLParser.skip_section
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_skip_section, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves the buffer cursor to a certain offset (since the beginning) and reads the next node there. This method returns an error code.
*/
//go:nosplit
func (self class) SeekTo(position int64) Error.Code { //gd:XMLParser.seek
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Opens an XML [param file] for parsing. This method returns an error code.
*/
//go:nosplit
func (self class) Open(file String.Readable) Error.Code { //gd:XMLParser.open
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(file)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_open, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Opens an XML raw [param buffer] for parsing. This method returns an error code.
*/
//go:nosplit
func (self class) OpenBuffer(buffer Packed.Bytes) Error.Code { //gd:XMLParser.open_buffer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](buffer))))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XMLParser.Bind_open_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
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

type NodeType = gdclass.XMLParserNodeType //gd:XMLParser.NodeType

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
