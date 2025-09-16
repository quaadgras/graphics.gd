/*
[gdscript]
extends EditorResourceConversionPlugin

func _handles(resource: Resource):
	return resource is ImageTexture

func _converts_to():
	return "PortableCompressedTexture2D"

func _convert(itex: Resource):
	var ptex = PortableCompressedTexture2D.new()
	ptex.create_from_image(itex.get_image(), PortableCompressedTexture2D.COMPRESSION_MODE_LOSSLESS)
	return ptex
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/EditorResourceConversionPlugin"
	"graphics.gd/classdb/ImageTexture"
	"graphics.gd/classdb/PortableCompressedTexture2D"
	"graphics.gd/classdb/Resource"
	"graphics.gd/variant/Object"
)

type ExampleEditorResourceConversionPlugin struct {
	EditorResourceConversionPlugin.Extension[ExampleEditorResourceConversionPlugin]
}

func (e *ExampleEditorResourceConversionPlugin) Handles(resource Resource.Instance) bool {
	return Object.Is[ImageTexture.Instance](resource)
}

func (e *ExampleEditorResourceConversionPlugin) ConvertsTo() string {
	return "PortableCompressedTexture2D"
}

func (e *ExampleEditorResourceConversionPlugin) Convert(itex Resource.Instance) Resource.Instance {
	var ptex = PortableCompressedTexture2D.New()
	ptex.CreateFromImage(Object.To[ImageTexture.Instance](itex).AsTexture2D().GetImage(), PortableCompressedTexture2D.CompressionModeLossless)
	return ptex.AsResource()
}
