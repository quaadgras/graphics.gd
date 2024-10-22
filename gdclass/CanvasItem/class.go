package CanvasItem

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Abstract base class for everything in 2D space. Canvas items are laid out in a tree; children inherit and extend their parent's transform. [CanvasItem] is extended by [Control] for GUI-related nodes, and by [Node2D] for 2D game objects.
Any [CanvasItem] can draw. For this, [method queue_redraw] is called by the engine, then [constant NOTIFICATION_DRAW] will be received on idle time to request a redraw. Because of this, canvas items don't need to be redrawn on every frame, improving the performance significantly. Several functions for drawing on the [CanvasItem] are provided (see [code]draw_*[/code] functions). However, they can only be used inside [method _draw], its corresponding [method Object._notification] or methods connected to the [signal draw] signal.
Canvas items are drawn in tree order on their canvas layer. By default, children are on top of their parents, so a root [CanvasItem] will be drawn behind everything. This behavior can be changed on a per-item basis.
A [CanvasItem] can be hidden, which will also hide its children. By adjusting various other properties of a [CanvasItem], you can also modulate its color (via [member modulate] or [member self_modulate]), change its Z-index, blend mode, and more.
Note that properties like transform, modulation, and visibility are only propagated to [i]direct[/i] [CanvasItem] child nodes. If there is a non-[CanvasItem] node in between, like [Node] or [AnimationPlayer], the [CanvasItem] nodes below will have an independent position and [member modulate] chain. See also [member top_level].
	// CanvasItem methods that can be overridden by a [Class] that extends it.
	type CanvasItem interface {
		//Called when [CanvasItem] has been requested to redraw (after [method queue_redraw] is called, either manually or by the engine).
		//Corresponds to the [constant NOTIFICATION_DRAW] notification in [method Object._notification].
		Draw() 
	}

*/
type Go [1]classdb.CanvasItem

/*
Called when [CanvasItem] has been requested to redraw (after [method queue_redraw] is called, either manually or by the engine).
Corresponds to the [constant NOTIFICATION_DRAW] notification in [method Object._notification].
*/
func (Go) _draw(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Returns the canvas item RID used by [RenderingServer] for this item.
*/
func (self Go) GetCanvasItem() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetCanvasItem())
}

/*
Returns [code]true[/code] if the node is present in the [SceneTree], its [member visible] property is [code]true[/code] and all its ancestors are also visible. If any ancestor is hidden, this node will not be visible in the scene tree, and is therefore not drawn (see [method _draw]).
Visibility is checked only in parent nodes that inherit from [CanvasItem], [CanvasLayer], and [Window]. If the parent is of any other type (such as [Node], [AnimationPlayer], or [Node3D]), it is assumed to be visible.
*/
func (self Go) IsVisibleInTree() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsVisibleInTree())
}

/*
Show the [CanvasItem] if it's currently hidden. This is equivalent to setting [member visible] to [code]true[/code]. For controls that inherit [Popup], the correct way to make them visible is to call one of the multiple [code]popup*()[/code] functions instead.
*/
func (self Go) Show() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Show()
}

/*
Hide the [CanvasItem] if it's currently visible. This is equivalent to setting [member visible] to [code]false[/code].
*/
func (self Go) Hide() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Hide()
}

/*
Queues the [CanvasItem] to redraw. During idle time, if [CanvasItem] is visible, [constant NOTIFICATION_DRAW] is sent and [method _draw] is called. This only occurs [b]once[/b] per frame, even if this method has been called multiple times.
*/
func (self Go) QueueRedraw() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).QueueRedraw()
}

/*
Moves this node to display on top of its siblings.
Internally, the node is moved to the bottom of parent's child list. The method has no effect on nodes without a parent.
*/
func (self Go) MoveToFront() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MoveToFront()
}

/*
Draws a line from a 2D point to another, with a given color and width. It can be optionally antialiased. See also [method draw_multiline] and [method draw_polyline].
If [param width] is negative, then a two-point primitive will be drawn instead of a four-point one. This means that when the CanvasItem is scaled, the line will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
func (self Go) DrawLine(from gd.Vector2, to gd.Vector2, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawLine(from, to, color, gd.Float(-1.0), false)
}

/*
Draws a dashed line from a 2D point to another, with a given color and width. See also [method draw_multiline] and [method draw_polyline].
If [param width] is negative, then a two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the line parts will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
If [param antialiased] is [code]true[/code], half transparent "feathers" will be attached to the boundary, making outlines smooth.
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
func (self Go) DrawDashedLine(from gd.Vector2, to gd.Vector2, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawDashedLine(from, to, color, gd.Float(-1.0), gd.Float(2.0), true, false)
}

/*
Draws interconnected line segments with a uniform [param color] and [param width] and optional antialiasing (supported only for positive [param width]). When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw disconnected lines, use [method draw_multiline] instead. See also [method draw_polygon].
If [param width] is negative, it will be ignored and the polyline will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the polyline will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
func (self Go) DrawPolyline(points []gd.Vector2, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawPolyline(gc.PackedVector2Slice(points), color, gd.Float(-1.0), false)
}

/*
Draws interconnected line segments with a uniform [param width], point-by-point coloring, and optional antialiasing (supported only for positive [param width]). Colors assigned to line points match by index between [param points] and [param colors], i.e. each line segment is filled with a gradient between the colors of the endpoints. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw disconnected lines, use [method draw_multiline_colors] instead. See also [method draw_polygon].
If [param width] is negative, it will be ignored and the polyline will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the polyline will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
func (self Go) DrawPolylineColors(points []gd.Vector2, colors []gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawPolylineColors(gc.PackedVector2Slice(points), gc.PackedColorSlice(colors), gd.Float(-1.0), false)
}

/*
Draws an unfilled arc between the given angles with a uniform [param color] and [param width] and optional antialiasing (supported only for positive [param width]). The larger the value of [param point_count], the smoother the curve. See also [method draw_circle].
If [param width] is negative, it will be ignored and the arc will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the arc will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
The arc is drawn from [param start_angle] towards the value of [param end_angle] so in clockwise direction if [code]start_angle < end_angle[/code] and counter-clockwise otherwise. Passing the same angles but in reversed order will produce the same arc. If absolute difference of [param start_angle] and [param end_angle] is greater than [constant @GDScript.TAU] radians, then a full circle arc is drawn (i.e. arc will not overlap itself).
*/
func (self Go) DrawArc(center gd.Vector2, radius float64, start_angle float64, end_angle float64, point_count int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawArc(center, gd.Float(radius), gd.Float(start_angle), gd.Float(end_angle), gd.Int(point_count), color, gd.Float(-1.0), false)
}

/*
Draws multiple disconnected lines with a uniform [param width] and [param color]. Each line is defined by two consecutive points from [param points] array, i.e. i-th segment consists of [code]points[2 * i][/code], [code]points[2 * i + 1][/code] endpoints. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw interconnected lines, use [method draw_polyline] instead.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
func (self Go) DrawMultiline(points []gd.Vector2, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawMultiline(gc.PackedVector2Slice(points), color, gd.Float(-1.0), false)
}

/*
Draws multiple disconnected lines with a uniform [param width] and segment-by-segment coloring. Each segment is defined by two consecutive points from [param points] array and a corresponding color from [param colors] array, i.e. i-th segment consists of [code]points[2 * i][/code], [code]points[2 * i + 1][/code] endpoints and has [code]colors[i][/code] color. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw interconnected lines, use [method draw_polyline_colors] instead.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
func (self Go) DrawMultilineColors(points []gd.Vector2, colors []gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawMultilineColors(gc.PackedVector2Slice(points), gc.PackedColorSlice(colors), gd.Float(-1.0), false)
}

/*
Draws a rectangle. If [param filled] is [code]true[/code], the rectangle will be filled with the [param color] specified. If [param filled] is [code]false[/code], the rectangle will be drawn as a stroke with the [param color] and [param width] specified. See also [method draw_texture_rect].
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
If [param antialiased] is [code]true[/code], half transparent "feathers" will be attached to the boundary, making outlines smooth.
[b]Note:[/b] [param width] is only effective if [param filled] is [code]false[/code].
[b]Note:[/b] Unfilled rectangles drawn with a negative [param width] may not display perfectly. For example, corners may be missing or brighter due to overlapping lines (for a translucent [param color]).
*/
func (self Go) DrawRect(rect gd.Rect2, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawRect(rect, color, true, gd.Float(-1.0), false)
}

/*
Draws a circle. See also [method draw_arc], [method draw_polyline], and [method draw_polygon].
If [param filled] is [code]true[/code], the circle will be filled with the [param color] specified. If [param filled] is [code]false[/code], the circle will be drawn as a stroke with the [param color] and [param width] specified.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
If [param antialiased] is [code]true[/code], half transparent "feathers" will be attached to the boundary, making outlines smooth.
[b]Note:[/b] [param width] is only effective if [param filled] is [code]false[/code].
*/
func (self Go) DrawCircle(position gd.Vector2, radius float64, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawCircle(position, gd.Float(radius), color, true, gd.Float(-1.0), false)
}

/*
Draws a texture at a given position.
*/
func (self Go) DrawTexture(texture gdclass.Texture2D, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawTexture(texture, position, gd.Color{1, 1, 1, 1})
}

/*
Draws a textured rectangle at a given position, optionally modulated by a color. If [param transpose] is [code]true[/code], the texture will have its X and Y coordinates swapped. See also [method draw_rect] and [method draw_texture_rect_region].
*/
func (self Go) DrawTextureRect(texture gdclass.Texture2D, rect gd.Rect2, tile bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawTextureRect(texture, rect, tile, gd.Color{1, 1, 1, 1}, false)
}

/*
Draws a textured rectangle from a texture's region (specified by [param src_rect]) at a given position, optionally modulated by a color. If [param transpose] is [code]true[/code], the texture will have its X and Y coordinates swapped. See also [method draw_texture_rect].
*/
func (self Go) DrawTextureRectRegion(texture gdclass.Texture2D, rect gd.Rect2, src_rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawTextureRectRegion(texture, rect, src_rect, gd.Color{1, 1, 1, 1}, false, true)
}

/*
Draws a textured rectangle region of the multi-channel signed distance field texture at a given position, optionally modulated by a color. See [member FontFile.multichannel_signed_distance_field] for more information and caveats about MSDF font rendering.
If [param outline] is positive, each alpha channel value of pixel in region is set to maximum value of true distance in the [param outline] radius.
Value of the [param pixel_range] should the same that was used during distance field texture generation.
*/
func (self Go) DrawMsdfTextureRectRegion(texture gdclass.Texture2D, rect gd.Rect2, src_rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawMsdfTextureRectRegion(texture, rect, src_rect, gd.Color{1, 1, 1, 1}, gd.Float(0.0), gd.Float(4.0), gd.Float(1.0))
}

/*
Draws a textured rectangle region of the font texture with LCD subpixel anti-aliasing at a given position, optionally modulated by a color.
Texture is drawn using the following blend operation, blend mode of the [CanvasItemMaterial] is ignored:
[codeblock]
dst.r = texture.r * modulate.r * modulate.a + dst.r * (1.0 - texture.r * modulate.a);
dst.g = texture.g * modulate.g * modulate.a + dst.g * (1.0 - texture.g * modulate.a);
dst.b = texture.b * modulate.b * modulate.a + dst.b * (1.0 - texture.b * modulate.a);
dst.a = modulate.a + dst.a * (1.0 - modulate.a);
[/codeblock]
*/
func (self Go) DrawLcdTextureRectRegion(texture gdclass.Texture2D, rect gd.Rect2, src_rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawLcdTextureRectRegion(texture, rect, src_rect, gd.Color{1, 1, 1, 1})
}

/*
Draws a styled rectangle.
*/
func (self Go) DrawStyleBox(style_box gdclass.StyleBox, rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawStyleBox(style_box, rect)
}

/*
Draws a custom primitive. 1 point for a point, 2 points for a line, 3 points for a triangle, and 4 points for a quad. If 0 points or more than 4 points are specified, nothing will be drawn and an error message will be printed. See also [method draw_line], [method draw_polyline], [method draw_polygon], and [method draw_rect].
*/
func (self Go) DrawPrimitive(points []gd.Vector2, colors []gd.Color, uvs []gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawPrimitive(gc.PackedVector2Slice(points), gc.PackedColorSlice(colors), gc.PackedVector2Slice(uvs), ([1]gdclass.Texture2D{}[0]))
}

/*
Draws a solid polygon of any number of points, convex or concave. Unlike [method draw_colored_polygon], each point's color can be changed individually. See also [method draw_polyline] and [method draw_polyline_colors]. If you need more flexibility (such as being able to use bones), use [method RenderingServer.canvas_item_add_triangle_array] instead.
*/
func (self Go) DrawPolygon(points []gd.Vector2, colors []gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawPolygon(gc.PackedVector2Slice(points), gc.PackedColorSlice(colors), gc.PackedVector2Slice(([1][]gd.Vector2{}[0])), ([1]gdclass.Texture2D{}[0]))
}

/*
Draws a colored polygon of any number of points, convex or concave. Unlike [method draw_polygon], a single color must be specified for the whole polygon.
*/
func (self Go) DrawColoredPolygon(points []gd.Vector2, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawColoredPolygon(gc.PackedVector2Slice(points), color, gc.PackedVector2Slice(([1][]gd.Vector2{}[0])), ([1]gdclass.Texture2D{}[0]))
}

/*
Draws [param text] using the specified [param font] at the [param pos] (bottom-left corner using the baseline of the font). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
[b]Example using the default project font:[/b]
[codeblocks]
[gdscript]
# If using this method in a script that redraws constantly, move the
# `default_font` declaration to a member variable assigned in `_ready()`
# so the Control is only created once.
var default_font = ThemeDB.fallback_font
var default_font_size = ThemeDB.fallback_font_size
draw_string(default_font, Vector2(64, 64), "Hello world", HORIZONTAL_ALIGNMENT_LEFT, -1, default_font_size)
[/gdscript]
[csharp]
// If using this method in a script that redraws constantly, move the
// `default_font` declaration to a member variable assigned in `_Ready()`
// so the Control is only created once.
Font defaultFont = ThemeDB.FallbackFont;
int defaultFontSize = ThemeDB.FallbackFontSize;
DrawString(defaultFont, new Vector2(64, 64), "Hello world", HORIZONTAL_ALIGNMENT_LEFT, -1, defaultFontSize);
[/csharp]
[/codeblocks]
See also [method Font.draw_string].
*/
func (self Go) DrawString(font gdclass.Font, pos gd.Vector2, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawString(font, pos, gc.String(text), 0, gd.Float(-1), gd.Int(16), gd.Color{1, 1, 1, 1}, 3, 0, 0)
}

/*
Breaks [param text] into lines and draws it using the specified [param font] at the [param pos] (top-left corner). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
func (self Go) DrawMultilineString(font gdclass.Font, pos gd.Vector2, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawMultilineString(font, pos, gc.String(text), 0, gd.Float(-1), gd.Int(16), gd.Int(-1), gd.Color{1, 1, 1, 1}, 3, 3, 0, 0)
}

/*
Draws [param text] outline using the specified [param font] at the [param pos] (bottom-left corner using the baseline of the font). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
func (self Go) DrawStringOutline(font gdclass.Font, pos gd.Vector2, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawStringOutline(font, pos, gc.String(text), 0, gd.Float(-1), gd.Int(16), gd.Int(1), gd.Color{1, 1, 1, 1}, 3, 0, 0)
}

/*
Breaks [param text] to the lines and draws text outline using the specified [param font] at the [param pos] (top-left corner). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
func (self Go) DrawMultilineStringOutline(font gdclass.Font, pos gd.Vector2, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawMultilineStringOutline(font, pos, gc.String(text), 0, gd.Float(-1), gd.Int(16), gd.Int(-1), gd.Int(1), gd.Color{1, 1, 1, 1}, 3, 3, 0, 0)
}

/*
Draws a string first character using a custom font.
*/
func (self Go) DrawChar(font gdclass.Font, pos gd.Vector2, char string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawChar(font, pos, gc.String(char), gd.Int(16), gd.Color{1, 1, 1, 1})
}

/*
Draws a string first character outline using a custom font.
*/
func (self Go) DrawCharOutline(font gdclass.Font, pos gd.Vector2, char string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawCharOutline(font, pos, gc.String(char), gd.Int(16), gd.Int(-1), gd.Color{1, 1, 1, 1})
}

/*
Draws a [Mesh] in 2D, using the provided texture. See [MeshInstance2D] for related documentation.
*/
func (self Go) DrawMesh(mesh gdclass.Mesh, texture gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawMesh(mesh, texture, gd.NewTransform2D(1, 0, 0, 1, 0, 0), gd.Color{1, 1, 1, 1})
}

/*
Draws a [MultiMesh] in 2D with the provided texture. See [MultiMeshInstance2D] for related documentation.
*/
func (self Go) DrawMultimesh(multimesh gdclass.MultiMesh, texture gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawMultimesh(multimesh, texture)
}

/*
Sets a custom transform for drawing via components. Anything drawn afterwards will be transformed by this.
[b]Note:[/b] [member FontFile.oversampling] does [i]not[/i] take [param scale] into account. This means that scaling up/down will cause bitmap fonts and rasterized (non-MSDF) dynamic fonts to appear blurry or pixelated. To ensure text remains crisp regardless of scale, you can enable MSDF font rendering by enabling [member ProjectSettings.gui/theme/default_font_multichannel_signed_distance_field] (applies to the default project font only), or enabling [b]Multichannel Signed Distance Field[/b] in the import options of a DynamicFont for custom fonts. On system fonts, [member SystemFont.multichannel_signed_distance_field] can be enabled in the inspector.
*/
func (self Go) DrawSetTransform(position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawSetTransform(position, gd.Float(0.0), gd.Vector2{1, 1})
}

/*
Sets a custom transform for drawing via matrix. Anything drawn afterwards will be transformed by this.
*/
func (self Go) DrawSetTransformMatrix(xform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawSetTransformMatrix(xform)
}

/*
Subsequent drawing commands will be ignored unless they fall within the specified animation slice. This is a faster way to implement animations that loop on background rather than redrawing constantly.
*/
func (self Go) DrawAnimationSlice(animation_length float64, slice_begin float64, slice_end float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawAnimationSlice(gd.Float(animation_length), gd.Float(slice_begin), gd.Float(slice_end), gd.Float(0.0))
}

/*
After submitting all animations slices via [method draw_animation_slice], this function can be used to revert drawing to its default state (all subsequent drawing commands will be visible). If you don't care about this particular use case, usage of this function after submitting the slices is not required.
*/
func (self Go) DrawEndAnimation() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawEndAnimation()
}

/*
Returns the transform matrix of this item.
*/
func (self Go) GetTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(class(self).GetTransform())
}

/*
Returns the global transform matrix of this item, i.e. the combined transform up to the topmost [CanvasItem] node. The topmost item is a [CanvasItem] that either has no parent, has non-[CanvasItem] parent or it has [member top_level] enabled.
*/
func (self Go) GetGlobalTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(class(self).GetGlobalTransform())
}

/*
Returns the transform from the local coordinate system of this [CanvasItem] to the [Viewport]s coordinate system.
*/
func (self Go) GetGlobalTransformWithCanvas() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(class(self).GetGlobalTransformWithCanvas())
}

/*
Returns the transform from the coordinate system of the canvas, this item is in, to the [Viewport]s embedders coordinate system.
*/
func (self Go) GetViewportTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(class(self).GetViewportTransform())
}

/*
Returns the viewport's boundaries as a [Rect2].
*/
func (self Go) GetViewportRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).GetViewportRect())
}

/*
Returns the transform from the coordinate system of the canvas, this item is in, to the [Viewport]s coordinate system.
*/
func (self Go) GetCanvasTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(class(self).GetCanvasTransform())
}

/*
Returns the transform of this [CanvasItem] in global screen coordinates (i.e. taking window position into account). Mostly useful for editor plugins.
Equals to [method get_global_transform] if the window is embedded (see [member Viewport.gui_embed_subwindows]).
*/
func (self Go) GetScreenTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(class(self).GetScreenTransform())
}

/*
Returns the mouse's position in this [CanvasItem] using the local coordinate system of this [CanvasItem].
*/
func (self Go) GetLocalMousePosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetLocalMousePosition())
}

/*
Returns the mouse's position in the [CanvasLayer] that this [CanvasItem] is in using the coordinate system of the [CanvasLayer].
[b]Note:[/b] For screen-space coordinates (e.g. when using a non-embedded [Popup]), you can use [method DisplayServer.mouse_get_position].
*/
func (self Go) GetGlobalMousePosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetGlobalMousePosition())
}

/*
Returns the [RID] of the [World2D] canvas where this item is in.
*/
func (self Go) GetCanvas() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetCanvas())
}

/*
Returns the [CanvasLayer] that contains this node, or [code]null[/code] if the node is not in any [CanvasLayer].
*/
func (self Go) GetCanvasLayerNode() gdclass.CanvasLayer {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.CanvasLayer(class(self).GetCanvasLayerNode(gc))
}

/*
Returns the [World2D] where this item is in.
*/
func (self Go) GetWorld2d() gdclass.World2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.World2D(class(self).GetWorld2d(gc))
}

/*
If [param enable] is [code]true[/code], this node will receive [constant NOTIFICATION_LOCAL_TRANSFORM_CHANGED] when its local transform changes.
*/
func (self Go) SetNotifyLocalTransform(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNotifyLocalTransform(enable)
}

/*
Returns [code]true[/code] if local transform notifications are communicated to children.
*/
func (self Go) IsLocalTransformNotificationEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLocalTransformNotificationEnabled())
}

/*
If [param enable] is [code]true[/code], this node will receive [constant NOTIFICATION_TRANSFORM_CHANGED] when its global transform changes.
*/
func (self Go) SetNotifyTransform(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNotifyTransform(enable)
}

/*
Returns [code]true[/code] if global transform notifications are communicated to children.
*/
func (self Go) IsTransformNotificationEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsTransformNotificationEnabled())
}

/*
Forces the transform to update. Transform changes in physics are not instant for performance reasons. Transforms are accumulated and then set. Use this if you need an up-to-date transform when doing physics operations.
*/
func (self Go) ForceUpdateTransform() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ForceUpdateTransform()
}

/*
Assigns [param screen_point] as this node's new local transform.
*/
func (self Go) MakeCanvasPositionLocal(screen_point gd.Vector2) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).MakeCanvasPositionLocal(screen_point))
}

/*
Transformations issued by [param event]'s inputs are applied in local space instead of global space.
*/
func (self Go) MakeInputLocal(event gdclass.InputEvent) gdclass.InputEvent {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.InputEvent(class(self).MakeInputLocal(gc, event))
}

/*
Set/clear individual bits on the rendering visibility layer. This simplifies editing this [CanvasItem]'s visibility layer.
*/
func (self Go) SetVisibilityLayerBit(layer int, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityLayerBit(gd.Int(layer), enabled)
}

/*
Returns an individual bit on the rendering visibility layer.
*/
func (self Go) GetVisibilityLayerBit(layer int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetVisibilityLayerBit(gd.Int(layer)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CanvasItem
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CanvasItem"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Visible() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVisible())
}

func (self Go) SetVisible(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisible(value)
}

func (self Go) Modulate() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetModulate())
}

func (self Go) SetModulate(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetModulate(value)
}

func (self Go) SelfModulate() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetSelfModulate())
}

func (self Go) SetSelfModulate(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelfModulate(value)
}

func (self Go) ShowBehindParent() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawBehindParentEnabled())
}

func (self Go) SetShowBehindParent(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawBehindParent(value)
}

func (self Go) TopLevel() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSetAsTopLevel())
}

func (self Go) SetTopLevel(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAsTopLevel(value)
}

func (self Go) ClipChildren() classdb.CanvasItemClipChildrenMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CanvasItemClipChildrenMode(class(self).GetClipChildrenMode())
}

func (self Go) SetClipChildren(value classdb.CanvasItemClipChildrenMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetClipChildrenMode(value)
}

func (self Go) LightMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLightMask()))
}

func (self Go) SetLightMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLightMask(gd.Int(value))
}

func (self Go) VisibilityLayer() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetVisibilityLayer()))
}

func (self Go) SetVisibilityLayer(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityLayer(gd.Int(value))
}

func (self Go) ZIndex() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetZIndex()))
}

func (self Go) SetZIndex(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetZIndex(gd.Int(value))
}

func (self Go) ZAsRelative() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsZRelative())
}

func (self Go) SetZAsRelative(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetZAsRelative(value)
}

func (self Go) YSortEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsYSortEnabled())
}

func (self Go) SetYSortEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetYSortEnabled(value)
}

func (self Go) TextureFilter() classdb.CanvasItemTextureFilter {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CanvasItemTextureFilter(class(self).GetTextureFilter())
}

func (self Go) SetTextureFilter(value classdb.CanvasItemTextureFilter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureFilter(value)
}

func (self Go) TextureRepeat() classdb.CanvasItemTextureRepeat {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CanvasItemTextureRepeat(class(self).GetTextureRepeat())
}

func (self Go) SetTextureRepeat(value classdb.CanvasItemTextureRepeat) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureRepeat(value)
}

func (self Go) Material() gdclass.Material {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Material(class(self).GetMaterial(gc))
}

func (self Go) SetMaterial(value gdclass.Material) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaterial(value)
}

func (self Go) UseParentMaterial() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUseParentMaterial())
}

func (self Go) SetUseParentMaterial(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseParentMaterial(value)
}

/*
Called when [CanvasItem] has been requested to redraw (after [method queue_redraw] is called, either manually or by the engine).
Corresponds to the [constant NOTIFICATION_DRAW] notification in [method Object._notification].
*/
func (class) _draw(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Returns the canvas item RID used by [RenderingServer] for this item.
*/
//go:nosplit
func (self class) GetCanvasItem() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_canvas_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the node is present in the [SceneTree], its [member visible] property is [code]true[/code] and all its ancestors are also visible. If any ancestor is hidden, this node will not be visible in the scene tree, and is therefore not drawn (see [method _draw]).
Visibility is checked only in parent nodes that inherit from [CanvasItem], [CanvasLayer], and [Window]. If the parent is of any other type (such as [Node], [AnimationPlayer], or [Node3D]), it is assumed to be visible.
*/
//go:nosplit
func (self class) IsVisibleInTree() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_is_visible_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Show the [CanvasItem] if it's currently hidden. This is equivalent to setting [member visible] to [code]true[/code]. For controls that inherit [Popup], the correct way to make them visible is to call one of the multiple [code]popup*()[/code] functions instead.
*/
//go:nosplit
func (self class) Show()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_show, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Hide the [CanvasItem] if it's currently visible. This is equivalent to setting [member visible] to [code]false[/code].
*/
//go:nosplit
func (self class) Hide()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_hide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Queues the [CanvasItem] to redraw. During idle time, if [CanvasItem] is visible, [constant NOTIFICATION_DRAW] is sent and [method _draw] is called. This only occurs [b]once[/b] per frame, even if this method has been called multiple times.
*/
//go:nosplit
func (self class) QueueRedraw()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_queue_redraw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves this node to display on top of its siblings.
Internally, the node is moved to the bottom of parent's child list. The method has no effect on nodes without a parent.
*/
//go:nosplit
func (self class) MoveToFront()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_move_to_front, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAsTopLevel(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_as_top_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSetAsTopLevel() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_is_set_as_top_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLightMask(light_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_light_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLightMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_light_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetModulate(modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetModulate() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelfModulate(self_modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, self_modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_self_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSelfModulate() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_self_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZIndex(z_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, z_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_z_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_z_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZAsRelative(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_z_as_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsZRelative() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_is_z_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetYSortEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_y_sort_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsYSortEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_is_y_sort_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawBehindParent(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_draw_behind_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawBehindParentEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_is_draw_behind_parent_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draws a line from a 2D point to another, with a given color and width. It can be optionally antialiased. See also [method draw_multiline] and [method draw_polyline].
If [param width] is negative, then a two-point primitive will be drawn instead of a four-point one. This means that when the CanvasItem is scaled, the line will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
//go:nosplit
func (self class) DrawLine(from gd.Vector2, to gd.Vector2, color gd.Color, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a dashed line from a 2D point to another, with a given color and width. See also [method draw_multiline] and [method draw_polyline].
If [param width] is negative, then a two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the line parts will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
If [param antialiased] is [code]true[/code], half transparent "feathers" will be attached to the boundary, making outlines smooth.
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
//go:nosplit
func (self class) DrawDashedLine(from gd.Vector2, to gd.Vector2, color gd.Color, width gd.Float, dash gd.Float, aligned bool, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, dash)
	callframe.Arg(frame, aligned)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_dashed_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws interconnected line segments with a uniform [param color] and [param width] and optional antialiasing (supported only for positive [param width]). When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw disconnected lines, use [method draw_multiline] instead. See also [method draw_polygon].
If [param width] is negative, it will be ignored and the polyline will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the polyline will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
//go:nosplit
func (self class) DrawPolyline(points gd.PackedVector2Array, color gd.Color, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_polyline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws interconnected line segments with a uniform [param width], point-by-point coloring, and optional antialiasing (supported only for positive [param width]). Colors assigned to line points match by index between [param points] and [param colors], i.e. each line segment is filled with a gradient between the colors of the endpoints. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw disconnected lines, use [method draw_multiline_colors] instead. See also [method draw_polygon].
If [param width] is negative, it will be ignored and the polyline will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the polyline will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
//go:nosplit
func (self class) DrawPolylineColors(points gd.PackedVector2Array, colors gd.PackedColorArray, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, mmm.Get(colors))
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_polyline_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws an unfilled arc between the given angles with a uniform [param color] and [param width] and optional antialiasing (supported only for positive [param width]). The larger the value of [param point_count], the smoother the curve. See also [method draw_circle].
If [param width] is negative, it will be ignored and the arc will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the arc will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
The arc is drawn from [param start_angle] towards the value of [param end_angle] so in clockwise direction if [code]start_angle < end_angle[/code] and counter-clockwise otherwise. Passing the same angles but in reversed order will produce the same arc. If absolute difference of [param start_angle] and [param end_angle] is greater than [constant @GDScript.TAU] radians, then a full circle arc is drawn (i.e. arc will not overlap itself).
*/
//go:nosplit
func (self class) DrawArc(center gd.Vector2, radius gd.Float, start_angle gd.Float, end_angle gd.Float, point_count gd.Int, color gd.Color, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, center)
	callframe.Arg(frame, radius)
	callframe.Arg(frame, start_angle)
	callframe.Arg(frame, end_angle)
	callframe.Arg(frame, point_count)
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_arc, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws multiple disconnected lines with a uniform [param width] and [param color]. Each line is defined by two consecutive points from [param points] array, i.e. i-th segment consists of [code]points[2 * i][/code], [code]points[2 * i + 1][/code] endpoints. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw interconnected lines, use [method draw_polyline] instead.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
//go:nosplit
func (self class) DrawMultiline(points gd.PackedVector2Array, color gd.Color, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_multiline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws multiple disconnected lines with a uniform [param width] and segment-by-segment coloring. Each segment is defined by two consecutive points from [param points] array and a corresponding color from [param colors] array, i.e. i-th segment consists of [code]points[2 * i][/code], [code]points[2 * i + 1][/code] endpoints and has [code]colors[i][/code] color. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw interconnected lines, use [method draw_polyline_colors] instead.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
//go:nosplit
func (self class) DrawMultilineColors(points gd.PackedVector2Array, colors gd.PackedColorArray, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, mmm.Get(colors))
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_multiline_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a rectangle. If [param filled] is [code]true[/code], the rectangle will be filled with the [param color] specified. If [param filled] is [code]false[/code], the rectangle will be drawn as a stroke with the [param color] and [param width] specified. See also [method draw_texture_rect].
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
If [param antialiased] is [code]true[/code], half transparent "feathers" will be attached to the boundary, making outlines smooth.
[b]Note:[/b] [param width] is only effective if [param filled] is [code]false[/code].
[b]Note:[/b] Unfilled rectangles drawn with a negative [param width] may not display perfectly. For example, corners may be missing or brighter due to overlapping lines (for a translucent [param color]).
*/
//go:nosplit
func (self class) DrawRect(rect gd.Rect2, color gd.Color, filled bool, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	callframe.Arg(frame, color)
	callframe.Arg(frame, filled)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a circle. See also [method draw_arc], [method draw_polyline], and [method draw_polygon].
If [param filled] is [code]true[/code], the circle will be filled with the [param color] specified. If [param filled] is [code]false[/code], the circle will be drawn as a stroke with the [param color] and [param width] specified.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
If [param antialiased] is [code]true[/code], half transparent "feathers" will be attached to the boundary, making outlines smooth.
[b]Note:[/b] [param width] is only effective if [param filled] is [code]false[/code].
*/
//go:nosplit
func (self class) DrawCircle(position gd.Vector2, radius gd.Float, color gd.Color, filled bool, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, radius)
	callframe.Arg(frame, color)
	callframe.Arg(frame, filled)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_circle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a texture at a given position.
*/
//go:nosplit
func (self class) DrawTexture(texture gdclass.Texture2D, position gd.Vector2, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, position)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a textured rectangle at a given position, optionally modulated by a color. If [param transpose] is [code]true[/code], the texture will have its X and Y coordinates swapped. See also [method draw_rect] and [method draw_texture_rect_region].
*/
//go:nosplit
func (self class) DrawTextureRect(texture gdclass.Texture2D, rect gd.Rect2, tile bool, modulate gd.Color, transpose bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, rect)
	callframe.Arg(frame, tile)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_texture_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a textured rectangle from a texture's region (specified by [param src_rect]) at a given position, optionally modulated by a color. If [param transpose] is [code]true[/code], the texture will have its X and Y coordinates swapped. See also [method draw_texture_rect].
*/
//go:nosplit
func (self class) DrawTextureRectRegion(texture gdclass.Texture2D, rect gd.Rect2, src_rect gd.Rect2, modulate gd.Color, transpose bool, clip_uv bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, rect)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	callframe.Arg(frame, clip_uv)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_texture_rect_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a textured rectangle region of the multi-channel signed distance field texture at a given position, optionally modulated by a color. See [member FontFile.multichannel_signed_distance_field] for more information and caveats about MSDF font rendering.
If [param outline] is positive, each alpha channel value of pixel in region is set to maximum value of true distance in the [param outline] radius.
Value of the [param pixel_range] should the same that was used during distance field texture generation.
*/
//go:nosplit
func (self class) DrawMsdfTextureRectRegion(texture gdclass.Texture2D, rect gd.Rect2, src_rect gd.Rect2, modulate gd.Color, outline gd.Float, pixel_range gd.Float, scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, rect)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, outline)
	callframe.Arg(frame, pixel_range)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_msdf_texture_rect_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a textured rectangle region of the font texture with LCD subpixel anti-aliasing at a given position, optionally modulated by a color.
Texture is drawn using the following blend operation, blend mode of the [CanvasItemMaterial] is ignored:
[codeblock]
dst.r = texture.r * modulate.r * modulate.a + dst.r * (1.0 - texture.r * modulate.a);
dst.g = texture.g * modulate.g * modulate.a + dst.g * (1.0 - texture.g * modulate.a);
dst.b = texture.b * modulate.b * modulate.a + dst.b * (1.0 - texture.b * modulate.a);
dst.a = modulate.a + dst.a * (1.0 - modulate.a);
[/codeblock]
*/
//go:nosplit
func (self class) DrawLcdTextureRectRegion(texture gdclass.Texture2D, rect gd.Rect2, src_rect gd.Rect2, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, rect)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_lcd_texture_rect_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a styled rectangle.
*/
//go:nosplit
func (self class) DrawStyleBox(style_box gdclass.StyleBox, rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(style_box[0].AsPointer())[0])
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_style_box, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a custom primitive. 1 point for a point, 2 points for a line, 3 points for a triangle, and 4 points for a quad. If 0 points or more than 4 points are specified, nothing will be drawn and an error message will be printed. See also [method draw_line], [method draw_polyline], [method draw_polygon], and [method draw_rect].
*/
//go:nosplit
func (self class) DrawPrimitive(points gd.PackedVector2Array, colors gd.PackedColorArray, uvs gd.PackedVector2Array, texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, mmm.Get(colors))
	callframe.Arg(frame, mmm.Get(uvs))
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_primitive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a solid polygon of any number of points, convex or concave. Unlike [method draw_colored_polygon], each point's color can be changed individually. See also [method draw_polyline] and [method draw_polyline_colors]. If you need more flexibility (such as being able to use bones), use [method RenderingServer.canvas_item_add_triangle_array] instead.
*/
//go:nosplit
func (self class) DrawPolygon(points gd.PackedVector2Array, colors gd.PackedColorArray, uvs gd.PackedVector2Array, texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, mmm.Get(colors))
	callframe.Arg(frame, mmm.Get(uvs))
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a colored polygon of any number of points, convex or concave. Unlike [method draw_polygon], a single color must be specified for the whole polygon.
*/
//go:nosplit
func (self class) DrawColoredPolygon(points gd.PackedVector2Array, color gd.Color, uvs gd.PackedVector2Array, texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, color)
	callframe.Arg(frame, mmm.Get(uvs))
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_colored_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws [param text] using the specified [param font] at the [param pos] (bottom-left corner using the baseline of the font). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
[b]Example using the default project font:[/b]
[codeblocks]
[gdscript]
# If using this method in a script that redraws constantly, move the
# `default_font` declaration to a member variable assigned in `_ready()`
# so the Control is only created once.
var default_font = ThemeDB.fallback_font
var default_font_size = ThemeDB.fallback_font_size
draw_string(default_font, Vector2(64, 64), "Hello world", HORIZONTAL_ALIGNMENT_LEFT, -1, default_font_size)
[/gdscript]
[csharp]
// If using this method in a script that redraws constantly, move the
// `default_font` declaration to a member variable assigned in `_Ready()`
// so the Control is only created once.
Font defaultFont = ThemeDB.FallbackFont;
int defaultFontSize = ThemeDB.FallbackFontSize;
DrawString(defaultFont, new Vector2(64, 64), "Hello world", HORIZONTAL_ALIGNMENT_LEFT, -1, defaultFontSize);
[/csharp]
[/codeblocks]
See also [method Font.draw_string].
*/
//go:nosplit
func (self class) DrawString(font gdclass.Font, pos gd.Vector2, text gd.String, alignment gd.HorizontalAlignment, width gd.Float, font_size gd.Int, modulate gd.Color, justification_flags classdb.TextServerJustificationFlag, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Breaks [param text] into lines and draws it using the specified [param font] at the [param pos] (top-left corner). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
//go:nosplit
func (self class) DrawMultilineString(font gdclass.Font, pos gd.Vector2, text gd.String, alignment gd.HorizontalAlignment, width gd.Float, font_size gd.Int, max_lines gd.Int, modulate gd.Color, brk_flags classdb.TextServerLineBreakFlag, justification_flags classdb.TextServerJustificationFlag, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, max_lines)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, brk_flags)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_multiline_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws [param text] outline using the specified [param font] at the [param pos] (bottom-left corner using the baseline of the font). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
//go:nosplit
func (self class) DrawStringOutline(font gdclass.Font, pos gd.Vector2, text gd.String, alignment gd.HorizontalAlignment, width gd.Float, font_size gd.Int, size gd.Int, modulate gd.Color, justification_flags classdb.TextServerJustificationFlag, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_string_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Breaks [param text] to the lines and draws text outline using the specified [param font] at the [param pos] (top-left corner). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
//go:nosplit
func (self class) DrawMultilineStringOutline(font gdclass.Font, pos gd.Vector2, text gd.String, alignment gd.HorizontalAlignment, width gd.Float, font_size gd.Int, max_lines gd.Int, size gd.Int, modulate gd.Color, brk_flags classdb.TextServerLineBreakFlag, justification_flags classdb.TextServerJustificationFlag, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, max_lines)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, brk_flags)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_multiline_string_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a string first character using a custom font.
*/
//go:nosplit
func (self class) DrawChar(font gdclass.Font, pos gd.Vector2, char gd.String, font_size gd.Int, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, mmm.Get(char))
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a string first character outline using a custom font.
*/
//go:nosplit
func (self class) DrawCharOutline(font gdclass.Font, pos gd.Vector2, char gd.String, font_size gd.Int, size gd.Int, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, mmm.Get(char))
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_char_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a [Mesh] in 2D, using the provided texture. See [MeshInstance2D] for related documentation.
*/
//go:nosplit
func (self class) DrawMesh(mesh gdclass.Mesh, texture gdclass.Texture2D, transform gd.Transform2D, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, transform)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a [MultiMesh] in 2D with the provided texture. See [MultiMeshInstance2D] for related documentation.
*/
//go:nosplit
func (self class) DrawMultimesh(multimesh gdclass.MultiMesh, texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(multimesh[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_multimesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a custom transform for drawing via components. Anything drawn afterwards will be transformed by this.
[b]Note:[/b] [member FontFile.oversampling] does [i]not[/i] take [param scale] into account. This means that scaling up/down will cause bitmap fonts and rasterized (non-MSDF) dynamic fonts to appear blurry or pixelated. To ensure text remains crisp regardless of scale, you can enable MSDF font rendering by enabling [member ProjectSettings.gui/theme/default_font_multichannel_signed_distance_field] (applies to the default project font only), or enabling [b]Multichannel Signed Distance Field[/b] in the import options of a DynamicFont for custom fonts. On system fonts, [member SystemFont.multichannel_signed_distance_field] can be enabled in the inspector.
*/
//go:nosplit
func (self class) DrawSetTransform(position gd.Vector2, rotation gd.Float, scale gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, rotation)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a custom transform for drawing via matrix. Anything drawn afterwards will be transformed by this.
*/
//go:nosplit
func (self class) DrawSetTransformMatrix(xform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_set_transform_matrix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Subsequent drawing commands will be ignored unless they fall within the specified animation slice. This is a faster way to implement animations that loop on background rather than redrawing constantly.
*/
//go:nosplit
func (self class) DrawAnimationSlice(animation_length gd.Float, slice_begin gd.Float, slice_end gd.Float, offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, animation_length)
	callframe.Arg(frame, slice_begin)
	callframe.Arg(frame, slice_end)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_animation_slice, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
After submitting all animations slices via [method draw_animation_slice], this function can be used to revert drawing to its default state (all subsequent drawing commands will be visible). If you don't care about this particular use case, usage of this function after submitting the slices is not required.
*/
//go:nosplit
func (self class) DrawEndAnimation()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_draw_end_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the transform matrix of this item.
*/
//go:nosplit
func (self class) GetTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the global transform matrix of this item, i.e. the combined transform up to the topmost [CanvasItem] node. The topmost item is a [CanvasItem] that either has no parent, has non-[CanvasItem] parent or it has [member top_level] enabled.
*/
//go:nosplit
func (self class) GetGlobalTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_global_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transform from the local coordinate system of this [CanvasItem] to the [Viewport]s coordinate system.
*/
//go:nosplit
func (self class) GetGlobalTransformWithCanvas() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_global_transform_with_canvas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transform from the coordinate system of the canvas, this item is in, to the [Viewport]s embedders coordinate system.
*/
//go:nosplit
func (self class) GetViewportTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_viewport_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the viewport's boundaries as a [Rect2].
*/
//go:nosplit
func (self class) GetViewportRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_viewport_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transform from the coordinate system of the canvas, this item is in, to the [Viewport]s coordinate system.
*/
//go:nosplit
func (self class) GetCanvasTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transform of this [CanvasItem] in global screen coordinates (i.e. taking window position into account). Mostly useful for editor plugins.
Equals to [method get_global_transform] if the window is embedded (see [member Viewport.gui_embed_subwindows]).
*/
//go:nosplit
func (self class) GetScreenTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_screen_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the mouse's position in this [CanvasItem] using the local coordinate system of this [CanvasItem].
*/
//go:nosplit
func (self class) GetLocalMousePosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_local_mouse_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the mouse's position in the [CanvasLayer] that this [CanvasItem] is in using the coordinate system of the [CanvasLayer].
[b]Note:[/b] For screen-space coordinates (e.g. when using a non-embedded [Popup]), you can use [method DisplayServer.mouse_get_position].
*/
//go:nosplit
func (self class) GetGlobalMousePosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_global_mouse_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [RID] of the [World2D] canvas where this item is in.
*/
//go:nosplit
func (self class) GetCanvas() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_canvas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [CanvasLayer] that contains this node, or [code]null[/code] if the node is not in any [CanvasLayer].
*/
//go:nosplit
func (self class) GetCanvasLayerNode(ctx gd.Lifetime) gdclass.CanvasLayer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_canvas_layer_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.CanvasLayer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [World2D] where this item is in.
*/
//go:nosplit
func (self class) GetWorld2d(ctx gd.Lifetime) gdclass.World2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_world_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.World2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaterial(material gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime) gdclass.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseParentMaterial(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_use_parent_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseParentMaterial() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_use_parent_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param enable] is [code]true[/code], this node will receive [constant NOTIFICATION_LOCAL_TRANSFORM_CHANGED] when its local transform changes.
*/
//go:nosplit
func (self class) SetNotifyLocalTransform(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_notify_local_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if local transform notifications are communicated to children.
*/
//go:nosplit
func (self class) IsLocalTransformNotificationEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_is_local_transform_notification_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param enable] is [code]true[/code], this node will receive [constant NOTIFICATION_TRANSFORM_CHANGED] when its global transform changes.
*/
//go:nosplit
func (self class) SetNotifyTransform(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_notify_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if global transform notifications are communicated to children.
*/
//go:nosplit
func (self class) IsTransformNotificationEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_is_transform_notification_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Forces the transform to update. Transform changes in physics are not instant for performance reasons. Transforms are accumulated and then set. Use this if you need an up-to-date transform when doing physics operations.
*/
//go:nosplit
func (self class) ForceUpdateTransform()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_force_update_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Assigns [param screen_point] as this node's new local transform.
*/
//go:nosplit
func (self class) MakeCanvasPositionLocal(screen_point gd.Vector2) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_make_canvas_position_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Transformations issued by [param event]'s inputs are applied in local space instead of global space.
*/
//go:nosplit
func (self class) MakeInputLocal(ctx gd.Lifetime, event gdclass.InputEvent) gdclass.InputEvent {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(event[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_make_input_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.InputEvent
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityLayer(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_visibility_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityLayer() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_visibility_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set/clear individual bits on the rendering visibility layer. This simplifies editing this [CanvasItem]'s visibility layer.
*/
//go:nosplit
func (self class) SetVisibilityLayerBit(layer gd.Int, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_visibility_layer_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an individual bit on the rendering visibility layer.
*/
//go:nosplit
func (self class) GetVisibilityLayerBit(layer gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_visibility_layer_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureFilter(mode classdb.CanvasItemTextureFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureFilter() classdb.CanvasItemTextureFilter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CanvasItemTextureFilter](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureRepeat(mode classdb.CanvasItemTextureRepeat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureRepeat() classdb.CanvasItemTextureRepeat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CanvasItemTextureRepeat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetClipChildrenMode(mode classdb.CanvasItemClipChildrenMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_set_clip_children_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetClipChildrenMode() classdb.CanvasItemClipChildrenMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CanvasItemClipChildrenMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItem.Bind_get_clip_children_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnDraw(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("draw"), gc.Callable(cb), 0)
}


func (self Go) OnVisibilityChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("visibility_changed"), gc.Callable(cb), 0)
}


func (self Go) OnHidden(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("hidden"), gc.Callable(cb), 0)
}


func (self Go) OnItemRectChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("item_rect_changed"), gc.Callable(cb), 0)
}


func (self class) AsCanvasItem() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_draw": return reflect.ValueOf(self._draw);
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_draw": return reflect.ValueOf(self._draw);
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("CanvasItem", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type TextureFilter = classdb.CanvasItemTextureFilter

const (
/*The [CanvasItem] will inherit the filter from its parent.*/
	TextureFilterParentNode TextureFilter = 0
/*The texture filter reads from the nearest pixel only. This makes the texture look pixelated from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	TextureFilterNearest TextureFilter = 1
/*The texture filter blends between the nearest 4 pixels. This makes the texture look smooth from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	TextureFilterLinear TextureFilter = 2
/*The texture filter reads from the nearest pixel and blends between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look pixelated from up close, and smooth from a distance.
Use this for non-pixel art textures that may be viewed at a low scale (e.g. due to [Camera2D] zoom or sprite scaling), as mipmaps are important to smooth out pixels that are smaller than on-screen pixels.*/
	TextureFilterNearestWithMipmaps TextureFilter = 3
/*The texture filter blends between the nearest 4 pixels and between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look smooth from up close, and smooth from a distance.
Use this for non-pixel art textures that may be viewed at a low scale (e.g. due to [Camera2D] zoom or sprite scaling), as mipmaps are important to smooth out pixels that are smaller than on-screen pixels.*/
	TextureFilterLinearWithMipmaps TextureFilter = 4
/*The texture filter reads from the nearest pixel and blends between 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]) based on the angle between the surface and the camera view. This makes the texture look pixelated from up close, and smooth from a distance. Anisotropic filtering improves texture quality on surfaces that are almost in line with the camera, but is slightly slower. The anisotropic filtering level can be changed by adjusting [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].
[b]Note:[/b] This texture filter is rarely useful in 2D projects. [constant TEXTURE_FILTER_NEAREST_WITH_MIPMAPS] is usually more appropriate in this case.*/
	TextureFilterNearestWithMipmapsAnisotropic TextureFilter = 5
/*The texture filter blends between the nearest 4 pixels and blends between 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]) based on the angle between the surface and the camera view. This makes the texture look smooth from up close, and smooth from a distance. Anisotropic filtering improves texture quality on surfaces that are almost in line with the camera, but is slightly slower. The anisotropic filtering level can be changed by adjusting [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].
[b]Note:[/b] This texture filter is rarely useful in 2D projects. [constant TEXTURE_FILTER_LINEAR_WITH_MIPMAPS] is usually more appropriate in this case.*/
	TextureFilterLinearWithMipmapsAnisotropic TextureFilter = 6
/*Represents the size of the [enum TextureFilter] enum.*/
	TextureFilterMax TextureFilter = 7
)
type TextureRepeat = classdb.CanvasItemTextureRepeat

const (
/*The [CanvasItem] will inherit the filter from its parent.*/
	TextureRepeatParentNode TextureRepeat = 0
/*Texture will not repeat.*/
	TextureRepeatDisabled TextureRepeat = 1
/*Texture will repeat normally.*/
	TextureRepeatEnabled TextureRepeat = 2
/*Texture will repeat in a 2×2 tiled mode, where elements at even positions are mirrored.*/
	TextureRepeatMirror TextureRepeat = 3
/*Represents the size of the [enum TextureRepeat] enum.*/
	TextureRepeatMax TextureRepeat = 4
)
type ClipChildrenMode = classdb.CanvasItemClipChildrenMode

const (
/*Child draws over parent and is not clipped.*/
	ClipChildrenDisabled ClipChildrenMode = 0
/*Parent is used for the purposes of clipping only. Child is clipped to the parent's visible area, parent is not drawn.*/
	ClipChildrenOnly ClipChildrenMode = 1
/*Parent is used for clipping child, but parent is also drawn underneath child as normal before clipping child to its visible area.*/
	ClipChildrenAndDraw ClipChildrenMode = 2
/*Represents the size of the [enum ClipChildrenMode] enum.*/
	ClipChildrenMax ClipChildrenMode = 3
)
