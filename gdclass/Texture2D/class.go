package Texture2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A texture works by registering an image in the video hardware, which then can be used in 3D models or 2D [Sprite2D] or GUI [Control].
Textures are often created by loading them from a file. See [method @GDScript.load].
[Texture2D] is a base for other resources. It cannot be used directly.
[b]Note:[/b] The maximum texture size is 16384×16384 pixels due to graphics hardware limitations. Larger textures may fail to import.
	// Texture2D methods that can be overridden by a [Class] that extends it.
	type Texture2D interface {
		//Called when the [Texture2D]'s width is queried.
		GetWidth() int
		//Called when the [Texture2D]'s height is queried.
		GetHeight() int
		//Called when a pixel's opaque state in the [Texture2D] is queried at the specified [code](x, y)[/code] position.
		IsPixelOpaque(x int, y int) bool
		//Called when the presence of an alpha channel in the [Texture2D] is queried.
		HasAlpha() bool
		//Called when the entire [Texture2D] is requested to be drawn over a [CanvasItem], with the top-left offset specified in [param pos]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
		//[b]Note:[/b] This is only used in 2D rendering, not 3D.
		Draw(to_canvas_item gd.RID, pos gd.Vector2, modulate gd.Color, transpose bool) 
		//Called when the [Texture2D] is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
		//[b]Note:[/b] This is only used in 2D rendering, not 3D.
		DrawRect(to_canvas_item gd.RID, rect gd.Rect2, tile bool, modulate gd.Color, transpose bool) 
		//Called when a part of the [Texture2D] specified by [param src_rect]'s coordinates is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
		//[b]Note:[/b] This is only used in 2D rendering, not 3D.
		DrawRectRegion(to_canvas_item gd.RID, rect gd.Rect2, src_rect gd.Rect2, modulate gd.Color, transpose bool, clip_uv bool) 
	}

*/
type Go [1]classdb.Texture2D

/*
Called when the [Texture2D]'s width is queried.
*/
func (Go) _get_width(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the [Texture2D]'s height is queried.
*/
func (Go) _get_height(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when a pixel's opaque state in the [Texture2D] is queried at the specified [code](x, y)[/code] position.
*/
func (Go) _is_pixel_opaque(impl func(ptr unsafe.Pointer, x int, y int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var x = gd.UnsafeGet[gd.Int](p_args,0)
		var y = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(x), int(y))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the presence of an alpha channel in the [Texture2D] is queried.
*/
func (Go) _has_alpha(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the entire [Texture2D] is requested to be drawn over a [CanvasItem], with the top-left offset specified in [param pos]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (Go) _draw(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, pos gd.Vector2, modulate gd.Color, transpose bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,1)
		var modulate = gd.UnsafeGet[gd.Color](p_args,2)
		var transpose = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, to_canvas_item, pos, modulate, transpose)
	}
}

/*
Called when the [Texture2D] is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (Go) _draw_rect(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2, tile bool, modulate gd.Color, transpose bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		var tile = gd.UnsafeGet[bool](p_args,2)
		var modulate = gd.UnsafeGet[gd.Color](p_args,3)
		var transpose = gd.UnsafeGet[bool](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, to_canvas_item, rect, tile, modulate, transpose)
	}
}

/*
Called when a part of the [Texture2D] specified by [param src_rect]'s coordinates is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (Go) _draw_rect_region(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2, src_rect gd.Rect2, modulate gd.Color, transpose bool, clip_uv bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		var src_rect = gd.UnsafeGet[gd.Rect2](p_args,2)
		var modulate = gd.UnsafeGet[gd.Color](p_args,3)
		var transpose = gd.UnsafeGet[bool](p_args,4)
		var clip_uv = gd.UnsafeGet[bool](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, to_canvas_item, rect, src_rect, modulate, transpose, clip_uv)
	}
}

/*
Returns the texture width in pixels.
*/
func (self Go) GetWidth() int {
	return int(int(class(self).GetWidth()))
}

/*
Returns the texture height in pixels.
*/
func (self Go) GetHeight() int {
	return int(int(class(self).GetHeight()))
}

/*
Returns the texture size in pixels.
*/
func (self Go) GetSize() gd.Vector2 {
	return gd.Vector2(class(self).GetSize())
}

/*
Returns [code]true[/code] if this [Texture2D] has an alpha channel.
*/
func (self Go) HasAlpha() bool {
	return bool(class(self).HasAlpha())
}

/*
Draws the texture using a [CanvasItem] with the [RenderingServer] API at the specified [param position].
*/
func (self Go) Draw(canvas_item gd.RID, position gd.Vector2) {
	class(self).Draw(canvas_item, position, gd.Color{1, 1, 1, 1}, false)
}

/*
Draws the texture using a [CanvasItem] with the [RenderingServer] API.
*/
func (self Go) DrawRect(canvas_item gd.RID, rect gd.Rect2, tile bool) {
	class(self).DrawRect(canvas_item, rect, tile, gd.Color{1, 1, 1, 1}, false)
}

/*
Draws a part of the texture using a [CanvasItem] with the [RenderingServer] API.
*/
func (self Go) DrawRectRegion(canvas_item gd.RID, rect gd.Rect2, src_rect gd.Rect2) {
	class(self).DrawRectRegion(canvas_item, rect, src_rect, gd.Color{1, 1, 1, 1}, false, true)
}

/*
Returns an [Image] that is a copy of data from this [Texture2D] (a new [Image] is created each time). [Image]s can be accessed and manipulated directly.
[b]Note:[/b] This will return [code]null[/code] if this [Texture2D] is invalid.
[b]Note:[/b] This will fetch the texture data from the GPU, which might cause performance problems when overused.
*/
func (self Go) GetImage() gdclass.Image {
	return gdclass.Image(class(self).GetImage())
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture2D]).
*/
func (self Go) CreatePlaceholder() gdclass.Resource {
	return gdclass.Resource(class(self).CreatePlaceholder())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Texture2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Texture2D"))
	return Go{classdb.Texture2D(object)}
}

/*
Called when the [Texture2D]'s width is queried.
*/
func (class) _get_width(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture2D]'s height is queried.
*/
func (class) _get_height(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when a pixel's opaque state in the [Texture2D] is queried at the specified [code](x, y)[/code] position.
*/
func (class) _is_pixel_opaque(impl func(ptr unsafe.Pointer, x gd.Int, y gd.Int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var x = gd.UnsafeGet[gd.Int](p_args,0)
		var y = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, x, y)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the presence of an alpha channel in the [Texture2D] is queried.
*/
func (class) _has_alpha(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the entire [Texture2D] is requested to be drawn over a [CanvasItem], with the top-left offset specified in [param pos]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (class) _draw(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, pos gd.Vector2, modulate gd.Color, transpose bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,1)
		var modulate = gd.UnsafeGet[gd.Color](p_args,2)
		var transpose = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, to_canvas_item, pos, modulate, transpose)
	}
}

/*
Called when the [Texture2D] is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (class) _draw_rect(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2, tile bool, modulate gd.Color, transpose bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		var tile = gd.UnsafeGet[bool](p_args,2)
		var modulate = gd.UnsafeGet[gd.Color](p_args,3)
		var transpose = gd.UnsafeGet[bool](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, to_canvas_item, rect, tile, modulate, transpose)
	}
}

/*
Called when a part of the [Texture2D] specified by [param src_rect]'s coordinates is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (class) _draw_rect_region(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2, src_rect gd.Rect2, modulate gd.Color, transpose bool, clip_uv bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		var src_rect = gd.UnsafeGet[gd.Rect2](p_args,2)
		var modulate = gd.UnsafeGet[gd.Color](p_args,3)
		var transpose = gd.UnsafeGet[bool](p_args,4)
		var clip_uv = gd.UnsafeGet[bool](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, to_canvas_item, rect, src_rect, modulate, transpose, clip_uv)
	}
}

/*
Returns the texture width in pixels.
*/
//go:nosplit
func (self class) GetWidth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the texture height in pixels.
*/
//go:nosplit
func (self class) GetHeight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the texture size in pixels.
*/
//go:nosplit
func (self class) GetSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this [Texture2D] has an alpha channel.
*/
//go:nosplit
func (self class) HasAlpha() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_has_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draws the texture using a [CanvasItem] with the [RenderingServer] API at the specified [param position].
*/
//go:nosplit
func (self class) Draw(canvas_item gd.RID, position gd.Vector2, modulate gd.Color, transpose bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, position)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws the texture using a [CanvasItem] with the [RenderingServer] API.
*/
//go:nosplit
func (self class) DrawRect(canvas_item gd.RID, rect gd.Rect2, tile bool, modulate gd.Color, transpose bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, tile)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_draw_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a part of the texture using a [CanvasItem] with the [RenderingServer] API.
*/
//go:nosplit
func (self class) DrawRectRegion(canvas_item gd.RID, rect gd.Rect2, src_rect gd.Rect2, modulate gd.Color, transpose bool, clip_uv bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	callframe.Arg(frame, clip_uv)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_draw_rect_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an [Image] that is a copy of data from this [Texture2D] (a new [Image] is created each time). [Image]s can be accessed and manipulated directly.
[b]Note:[/b] This will return [code]null[/code] if this [Texture2D] is invalid.
[b]Note:[/b] This will fetch the texture data from the GPU, which might cause performance problems when overused.
*/
//go:nosplit
func (self class) GetImage() gdclass.Image {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_get_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Image{classdb.Image(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Creates a placeholder version of this resource ([PlaceholderTexture2D]).
*/
//go:nosplit
func (self class) CreatePlaceholder() gdclass.Resource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Resource{classdb.Resource(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsTexture2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_width": return reflect.ValueOf(self._get_width);
	case "_get_height": return reflect.ValueOf(self._get_height);
	case "_is_pixel_opaque": return reflect.ValueOf(self._is_pixel_opaque);
	case "_has_alpha": return reflect.ValueOf(self._has_alpha);
	case "_draw": return reflect.ValueOf(self._draw);
	case "_draw_rect": return reflect.ValueOf(self._draw_rect);
	case "_draw_rect_region": return reflect.ValueOf(self._draw_rect_region);
	default: return gd.VirtualByName(self.AsTexture(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_width": return reflect.ValueOf(self._get_width);
	case "_get_height": return reflect.ValueOf(self._get_height);
	case "_is_pixel_opaque": return reflect.ValueOf(self._is_pixel_opaque);
	case "_has_alpha": return reflect.ValueOf(self._has_alpha);
	case "_draw": return reflect.ValueOf(self._draw);
	case "_draw_rect": return reflect.ValueOf(self._draw_rect);
	case "_draw_rect_region": return reflect.ValueOf(self._draw_rect_region);
	default: return gd.VirtualByName(self.AsTexture(), name)
	}
}
func init() {classdb.Register("Texture2D", func(ptr gd.Object) any { return classdb.Texture2D(ptr) })}
