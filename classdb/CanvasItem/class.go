// Package CanvasItem provides methods for working with CanvasItem object instances.
package CanvasItem

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Transform2D"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Abstract base class for everything in 2D space. Canvas items are laid out in a tree; children inherit and extend their parent's transform. [CanvasItem] is extended by [Control] for GUI-related nodes, and by [Node2D] for 2D game objects.
Any [CanvasItem] can draw. For this, [method queue_redraw] is called by the engine, then [constant NOTIFICATION_DRAW] will be received on idle time to request a redraw. Because of this, canvas items don't need to be redrawn on every frame, improving the performance significantly. Several functions for drawing on the [CanvasItem] are provided (see [code]draw_*[/code] functions). However, they can only be used inside [method _draw], its corresponding [method Object._notification] or methods connected to the [signal draw] signal.
Canvas items are drawn in tree order on their canvas layer. By default, children are on top of their parents, so a root [CanvasItem] will be drawn behind everything. This behavior can be changed on a per-item basis.
A [CanvasItem] can be hidden, which will also hide its children. By adjusting various other properties of a [CanvasItem], you can also modulate its color (via [member modulate] or [member self_modulate]), change its Z-index, blend mode, and more.
Note that properties like transform, modulation, and visibility are only propagated to [i]direct[/i] [CanvasItem] child nodes. If there is a non-[CanvasItem] node in between, like [Node] or [AnimationPlayer], the [CanvasItem] nodes below will have an independent position and [member modulate] chain. See also [member top_level].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=CanvasItem)
*/
type Instance [1]gdclass.CanvasItem

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCanvasItem() Instance
}
type Interface interface {
	//Called when [CanvasItem] has been requested to redraw (after [method queue_redraw] is called, either manually or by the engine).
	//Corresponds to the [constant NOTIFICATION_DRAW] notification in [method Object._notification].
	Draw()
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) Draw() { return }

/*
Called when [CanvasItem] has been requested to redraw (after [method queue_redraw] is called, either manually or by the engine).
Corresponds to the [constant NOTIFICATION_DRAW] notification in [method Object._notification].
*/
func (Instance) _draw(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns the canvas item RID used by [RenderingServer] for this item.
*/
func (self Instance) GetCanvasItem() Resource.ID {
	return Resource.ID(class(self).GetCanvasItem())
}

/*
Returns [code]true[/code] if the node is present in the [SceneTree], its [member visible] property is [code]true[/code] and all its ancestors are also visible. If any ancestor is hidden, this node will not be visible in the scene tree, and is therefore not drawn (see [method _draw]).
Visibility is checked only in parent nodes that inherit from [CanvasItem], [CanvasLayer], and [Window]. If the parent is of any other type (such as [Node], [AnimationPlayer], or [Node3D]), it is assumed to be visible.
*/
func (self Instance) IsVisibleInTree() bool {
	return bool(class(self).IsVisibleInTree())
}

/*
Show the [CanvasItem] if it's currently hidden. This is equivalent to setting [member visible] to [code]true[/code]. For controls that inherit [Popup], the correct way to make them visible is to call one of the multiple [code]popup*()[/code] functions instead.
*/
func (self Instance) Show() {
	class(self).Show()
}

/*
Hide the [CanvasItem] if it's currently visible. This is equivalent to setting [member visible] to [code]false[/code].
*/
func (self Instance) Hide() {
	class(self).Hide()
}

/*
Queues the [CanvasItem] to redraw. During idle time, if [CanvasItem] is visible, [constant NOTIFICATION_DRAW] is sent and [method _draw] is called. This only occurs [b]once[/b] per frame, even if this method has been called multiple times.
*/
func (self Instance) QueueRedraw() {
	class(self).QueueRedraw()
}

/*
Moves this node to display on top of its siblings.
Internally, the node is moved to the bottom of parent's child list. The method has no effect on nodes without a parent.
*/
func (self Instance) MoveToFront() {
	class(self).MoveToFront()
}

/*
Draws a line from a 2D point to another, with a given color and width. It can be optionally antialiased. See also [method draw_multiline] and [method draw_polyline].
If [param width] is negative, then a two-point primitive will be drawn instead of a four-point one. This means that when the CanvasItem is scaled, the line will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
func (self Instance) DrawLine(from Vector2.XY, to Vector2.XY, color Color.RGBA) {
	class(self).DrawLine(gd.Vector2(from), gd.Vector2(to), gd.Color(color), gd.Float(-1.0), false)
}

/*
Draws a dashed line from a 2D point to another, with a given color and width. See also [method draw_multiline] and [method draw_polyline].
If [param width] is negative, then a two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the line parts will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
If [param antialiased] is [code]true[/code], half transparent "feathers" will be attached to the boundary, making outlines smooth.
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
func (self Instance) DrawDashedLine(from Vector2.XY, to Vector2.XY, color Color.RGBA) {
	class(self).DrawDashedLine(gd.Vector2(from), gd.Vector2(to), gd.Color(color), gd.Float(-1.0), gd.Float(2.0), true, false)
}

/*
Draws interconnected line segments with a uniform [param color] and [param width] and optional antialiasing (supported only for positive [param width]). When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw disconnected lines, use [method draw_multiline] instead. See also [method draw_polygon].
If [param width] is negative, it will be ignored and the polyline will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the polyline will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
func (self Instance) DrawPolyline(points []Vector2.XY, color Color.RGBA) {
	class(self).DrawPolyline(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&points))), gd.Color(color), gd.Float(-1.0), false)
}

/*
Draws interconnected line segments with a uniform [param width], point-by-point coloring, and optional antialiasing (supported only for positive [param width]). Colors assigned to line points match by index between [param points] and [param colors], i.e. each line segment is filled with a gradient between the colors of the endpoints. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw disconnected lines, use [method draw_multiline_colors] instead. See also [method draw_polygon].
If [param width] is negative, it will be ignored and the polyline will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the polyline will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
func (self Instance) DrawPolylineColors(points []Vector2.XY, colors []Color.RGBA) {
	class(self).DrawPolylineColors(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&points))), gd.NewPackedColorSlice(*(*[]gd.Color)(unsafe.Pointer(&colors))), gd.Float(-1.0), false)
}

/*
Draws an unfilled arc between the given angles with a uniform [param color] and [param width] and optional antialiasing (supported only for positive [param width]). The larger the value of [param point_count], the smoother the curve. See also [method draw_circle].
If [param width] is negative, it will be ignored and the arc will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the arc will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
The arc is drawn from [param start_angle] towards the value of [param end_angle] so in clockwise direction if [code]start_angle < end_angle[/code] and counter-clockwise otherwise. Passing the same angles but in reversed order will produce the same arc. If absolute difference of [param start_angle] and [param end_angle] is greater than [constant @GDScript.TAU] radians, then a full circle arc is drawn (i.e. arc will not overlap itself).
*/
func (self Instance) DrawArc(center Vector2.XY, radius Float.X, start_angle Float.X, end_angle Float.X, point_count int, color Color.RGBA) {
	class(self).DrawArc(gd.Vector2(center), gd.Float(radius), gd.Float(start_angle), gd.Float(end_angle), gd.Int(point_count), gd.Color(color), gd.Float(-1.0), false)
}

/*
Draws multiple disconnected lines with a uniform [param width] and [param color]. Each line is defined by two consecutive points from [param points] array, i.e. i-th segment consists of [code]points[2 * i][/code], [code]points[2 * i + 1][/code] endpoints. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw interconnected lines, use [method draw_polyline] instead.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
func (self Instance) DrawMultiline(points []Vector2.XY, color Color.RGBA) {
	class(self).DrawMultiline(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&points))), gd.Color(color), gd.Float(-1.0), false)
}

/*
Draws multiple disconnected lines with a uniform [param width] and segment-by-segment coloring. Each segment is defined by two consecutive points from [param points] array and a corresponding color from [param colors] array, i.e. i-th segment consists of [code]points[2 * i][/code], [code]points[2 * i + 1][/code] endpoints and has [code]colors[i][/code] color. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw interconnected lines, use [method draw_polyline_colors] instead.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
func (self Instance) DrawMultilineColors(points []Vector2.XY, colors []Color.RGBA) {
	class(self).DrawMultilineColors(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&points))), gd.NewPackedColorSlice(*(*[]gd.Color)(unsafe.Pointer(&colors))), gd.Float(-1.0), false)
}

/*
Draws a rectangle. If [param filled] is [code]true[/code], the rectangle will be filled with the [param color] specified. If [param filled] is [code]false[/code], the rectangle will be drawn as a stroke with the [param color] and [param width] specified. See also [method draw_texture_rect].
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
If [param antialiased] is [code]true[/code], half transparent "feathers" will be attached to the boundary, making outlines smooth.
[b]Note:[/b] [param width] is only effective if [param filled] is [code]false[/code].
[b]Note:[/b] Unfilled rectangles drawn with a negative [param width] may not display perfectly. For example, corners may be missing or brighter due to overlapping lines (for a translucent [param color]).
*/
func (self Instance) DrawRect(rect Rect2.PositionSize, color Color.RGBA) {
	class(self).DrawRect(gd.Rect2(rect), gd.Color(color), true, gd.Float(-1.0), false)
}

/*
Draws a circle. See also [method draw_arc], [method draw_polyline], and [method draw_polygon].
If [param filled] is [code]true[/code], the circle will be filled with the [param color] specified. If [param filled] is [code]false[/code], the circle will be drawn as a stroke with the [param color] and [param width] specified.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
If [param antialiased] is [code]true[/code], half transparent "feathers" will be attached to the boundary, making outlines smooth.
[b]Note:[/b] [param width] is only effective if [param filled] is [code]false[/code].
*/
func (self Instance) DrawCircle(position Vector2.XY, radius Float.X, color Color.RGBA) {
	class(self).DrawCircle(gd.Vector2(position), gd.Float(radius), gd.Color(color), true, gd.Float(-1.0), false)
}

/*
Draws a texture at a given position.
*/
func (self Instance) DrawTexture(texture [1]gdclass.Texture2D, position Vector2.XY) {
	class(self).DrawTexture(texture, gd.Vector2(position), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draws a textured rectangle at a given position, optionally modulated by a color. If [param transpose] is [code]true[/code], the texture will have its X and Y coordinates swapped. See also [method draw_rect] and [method draw_texture_rect_region].
*/
func (self Instance) DrawTextureRect(texture [1]gdclass.Texture2D, rect Rect2.PositionSize, tile bool) {
	class(self).DrawTextureRect(texture, gd.Rect2(rect), tile, gd.Color(gd.Color{1, 1, 1, 1}), false)
}

/*
Draws a textured rectangle from a texture's region (specified by [param src_rect]) at a given position, optionally modulated by a color. If [param transpose] is [code]true[/code], the texture will have its X and Y coordinates swapped. See also [method draw_texture_rect].
*/
func (self Instance) DrawTextureRectRegion(texture [1]gdclass.Texture2D, rect Rect2.PositionSize, src_rect Rect2.PositionSize) {
	class(self).DrawTextureRectRegion(texture, gd.Rect2(rect), gd.Rect2(src_rect), gd.Color(gd.Color{1, 1, 1, 1}), false, true)
}

/*
Draws a textured rectangle region of the multi-channel signed distance field texture at a given position, optionally modulated by a color. See [member FontFile.multichannel_signed_distance_field] for more information and caveats about MSDF font rendering.
If [param outline] is positive, each alpha channel value of pixel in region is set to maximum value of true distance in the [param outline] radius.
Value of the [param pixel_range] should the same that was used during distance field texture generation.
*/
func (self Instance) DrawMsdfTextureRectRegion(texture [1]gdclass.Texture2D, rect Rect2.PositionSize, src_rect Rect2.PositionSize) {
	class(self).DrawMsdfTextureRectRegion(texture, gd.Rect2(rect), gd.Rect2(src_rect), gd.Color(gd.Color{1, 1, 1, 1}), gd.Float(0.0), gd.Float(4.0), gd.Float(1.0))
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
func (self Instance) DrawLcdTextureRectRegion(texture [1]gdclass.Texture2D, rect Rect2.PositionSize, src_rect Rect2.PositionSize) {
	class(self).DrawLcdTextureRectRegion(texture, gd.Rect2(rect), gd.Rect2(src_rect), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draws a styled rectangle.
*/
func (self Instance) DrawStyleBox(style_box [1]gdclass.StyleBox, rect Rect2.PositionSize) {
	class(self).DrawStyleBox(style_box, gd.Rect2(rect))
}

/*
Draws a custom primitive. 1 point for a point, 2 points for a line, 3 points for a triangle, and 4 points for a quad. If 0 points or more than 4 points are specified, nothing will be drawn and an error message will be printed. See also [method draw_line], [method draw_polyline], [method draw_polygon], and [method draw_rect].
*/
func (self Instance) DrawPrimitive(points []Vector2.XY, colors []Color.RGBA, uvs []Vector2.XY) {
	class(self).DrawPrimitive(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&points))), gd.NewPackedColorSlice(*(*[]gd.Color)(unsafe.Pointer(&colors))), gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&uvs))), [1][1]gdclass.Texture2D{}[0])
}

/*
Draws a solid polygon of any number of points, convex or concave. Unlike [method draw_colored_polygon], each point's color can be changed individually. See also [method draw_polyline] and [method draw_polyline_colors]. If you need more flexibility (such as being able to use bones), use [method RenderingServer.canvas_item_add_triangle_array] instead.
*/
func (self Instance) DrawPolygon(points []Vector2.XY, colors []Color.RGBA) {
	class(self).DrawPolygon(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&points))), gd.NewPackedColorSlice(*(*[]gd.Color)(unsafe.Pointer(&colors))), gd.NewPackedVector2Slice(nil), [1][1]gdclass.Texture2D{}[0])
}

/*
Draws a colored polygon of any number of points, convex or concave. Unlike [method draw_polygon], a single color must be specified for the whole polygon.
*/
func (self Instance) DrawColoredPolygon(points []Vector2.XY, color Color.RGBA) {
	class(self).DrawColoredPolygon(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&points))), gd.Color(color), gd.NewPackedVector2Slice(nil), [1][1]gdclass.Texture2D{}[0])
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
func (self Instance) DrawString(font [1]gdclass.Font, pos Vector2.XY, text string) {
	class(self).DrawString(font, gd.Vector2(pos), gd.NewString(text), 0, gd.Float(-1), gd.Int(16), gd.Color(gd.Color{1, 1, 1, 1}), 3, 0, 0)
}

/*
Breaks [param text] into lines and draws it using the specified [param font] at the [param pos] (top-left corner). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
func (self Instance) DrawMultilineString(font [1]gdclass.Font, pos Vector2.XY, text string) {
	class(self).DrawMultilineString(font, gd.Vector2(pos), gd.NewString(text), 0, gd.Float(-1), gd.Int(16), gd.Int(-1), gd.Color(gd.Color{1, 1, 1, 1}), 3, 3, 0, 0)
}

/*
Draws [param text] outline using the specified [param font] at the [param pos] (bottom-left corner using the baseline of the font). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
func (self Instance) DrawStringOutline(font [1]gdclass.Font, pos Vector2.XY, text string) {
	class(self).DrawStringOutline(font, gd.Vector2(pos), gd.NewString(text), 0, gd.Float(-1), gd.Int(16), gd.Int(1), gd.Color(gd.Color{1, 1, 1, 1}), 3, 0, 0)
}

/*
Breaks [param text] to the lines and draws text outline using the specified [param font] at the [param pos] (top-left corner). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
func (self Instance) DrawMultilineStringOutline(font [1]gdclass.Font, pos Vector2.XY, text string) {
	class(self).DrawMultilineStringOutline(font, gd.Vector2(pos), gd.NewString(text), 0, gd.Float(-1), gd.Int(16), gd.Int(-1), gd.Int(1), gd.Color(gd.Color{1, 1, 1, 1}), 3, 3, 0, 0)
}

/*
Draws a string first character using a custom font.
*/
func (self Instance) DrawChar(font [1]gdclass.Font, pos Vector2.XY, char string) {
	class(self).DrawChar(font, gd.Vector2(pos), gd.NewString(char), gd.Int(16), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draws a string first character outline using a custom font.
*/
func (self Instance) DrawCharOutline(font [1]gdclass.Font, pos Vector2.XY, char string) {
	class(self).DrawCharOutline(font, gd.Vector2(pos), gd.NewString(char), gd.Int(16), gd.Int(-1), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draws a [Mesh] in 2D, using the provided texture. See [MeshInstance2D] for related documentation.
*/
func (self Instance) DrawMesh(mesh [1]gdclass.Mesh, texture [1]gdclass.Texture2D) {
	class(self).DrawMesh(mesh, texture, gd.Transform2D(gd.NewTransform2D(1, 0, 0, 1, 0, 0)), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draws a [MultiMesh] in 2D with the provided texture. See [MultiMeshInstance2D] for related documentation.
*/
func (self Instance) DrawMultimesh(multimesh [1]gdclass.MultiMesh, texture [1]gdclass.Texture2D) {
	class(self).DrawMultimesh(multimesh, texture)
}

/*
Sets a custom transform for drawing via components. Anything drawn afterwards will be transformed by this.
[b]Note:[/b] [member FontFile.oversampling] does [i]not[/i] take [param scale] into account. This means that scaling up/down will cause bitmap fonts and rasterized (non-MSDF) dynamic fonts to appear blurry or pixelated. To ensure text remains crisp regardless of scale, you can enable MSDF font rendering by enabling [member ProjectSettings.gui/theme/default_font_multichannel_signed_distance_field] (applies to the default project font only), or enabling [b]Multichannel Signed Distance Field[/b] in the import options of a DynamicFont for custom fonts. On system fonts, [member SystemFont.multichannel_signed_distance_field] can be enabled in the inspector.
*/
func (self Instance) DrawSetTransform(position Vector2.XY) {
	class(self).DrawSetTransform(gd.Vector2(position), gd.Float(0.0), gd.Vector2(gd.Vector2{1, 1}))
}

/*
Sets a custom transform for drawing via matrix. Anything drawn afterwards will be transformed by this.
*/
func (self Instance) DrawSetTransformMatrix(xform Transform2D.OriginXY) {
	class(self).DrawSetTransformMatrix(gd.Transform2D(xform))
}

/*
Subsequent drawing commands will be ignored unless they fall within the specified animation slice. This is a faster way to implement animations that loop on background rather than redrawing constantly.
*/
func (self Instance) DrawAnimationSlice(animation_length Float.X, slice_begin Float.X, slice_end Float.X) {
	class(self).DrawAnimationSlice(gd.Float(animation_length), gd.Float(slice_begin), gd.Float(slice_end), gd.Float(0.0))
}

/*
After submitting all animations slices via [method draw_animation_slice], this function can be used to revert drawing to its default state (all subsequent drawing commands will be visible). If you don't care about this particular use case, usage of this function after submitting the slices is not required.
*/
func (self Instance) DrawEndAnimation() {
	class(self).DrawEndAnimation()
}

/*
Returns the transform matrix of this item.
*/
func (self Instance) GetTransform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetTransform())
}

/*
Returns the global transform matrix of this item, i.e. the combined transform up to the topmost [CanvasItem] node. The topmost item is a [CanvasItem] that either has no parent, has non-[CanvasItem] parent or it has [member top_level] enabled.
*/
func (self Instance) GetGlobalTransform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetGlobalTransform())
}

/*
Returns the transform from the local coordinate system of this [CanvasItem] to the [Viewport]s coordinate system.
*/
func (self Instance) GetGlobalTransformWithCanvas() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetGlobalTransformWithCanvas())
}

/*
Returns the transform from the coordinate system of the canvas, this item is in, to the [Viewport]s embedders coordinate system.
*/
func (self Instance) GetViewportTransform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetViewportTransform())
}

/*
Returns the viewport's boundaries as a [Rect2].
*/
func (self Instance) GetViewportRect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetViewportRect())
}

/*
Returns the transform from the coordinate system of the canvas, this item is in, to the [Viewport]s coordinate system.
*/
func (self Instance) GetCanvasTransform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetCanvasTransform())
}

/*
Returns the transform of this [CanvasItem] in global screen coordinates (i.e. taking window position into account). Mostly useful for editor plugins.
Equals to [method get_global_transform] if the window is embedded (see [member Viewport.gui_embed_subwindows]).
*/
func (self Instance) GetScreenTransform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetScreenTransform())
}

/*
Returns the mouse's position in this [CanvasItem] using the local coordinate system of this [CanvasItem].
*/
func (self Instance) GetLocalMousePosition() Vector2.XY {
	return Vector2.XY(class(self).GetLocalMousePosition())
}

/*
Returns the mouse's position in the [CanvasLayer] that this [CanvasItem] is in using the coordinate system of the [CanvasLayer].
[b]Note:[/b] For screen-space coordinates (e.g. when using a non-embedded [Popup]), you can use [method DisplayServer.mouse_get_position].
*/
func (self Instance) GetGlobalMousePosition() Vector2.XY {
	return Vector2.XY(class(self).GetGlobalMousePosition())
}

/*
Returns the [RID] of the [World2D] canvas where this item is in.
*/
func (self Instance) GetCanvas() Resource.ID {
	return Resource.ID(class(self).GetCanvas())
}

/*
Returns the [CanvasLayer] that contains this node, or [code]null[/code] if the node is not in any [CanvasLayer].
*/
func (self Instance) GetCanvasLayerNode() [1]gdclass.CanvasLayer {
	return [1]gdclass.CanvasLayer(class(self).GetCanvasLayerNode())
}

/*
Returns the [World2D] where this item is in.
*/
func (self Instance) GetWorld2d() [1]gdclass.World2D {
	return [1]gdclass.World2D(class(self).GetWorld2d())
}

/*
If [param enable] is [code]true[/code], this node will receive [constant NOTIFICATION_LOCAL_TRANSFORM_CHANGED] when its local transform changes.
*/
func (self Instance) SetNotifyLocalTransform(enable bool) {
	class(self).SetNotifyLocalTransform(enable)
}

/*
Returns [code]true[/code] if local transform notifications are communicated to children.
*/
func (self Instance) IsLocalTransformNotificationEnabled() bool {
	return bool(class(self).IsLocalTransformNotificationEnabled())
}

/*
If [param enable] is [code]true[/code], this node will receive [constant NOTIFICATION_TRANSFORM_CHANGED] when its global transform changes.
*/
func (self Instance) SetNotifyTransform(enable bool) {
	class(self).SetNotifyTransform(enable)
}

/*
Returns [code]true[/code] if global transform notifications are communicated to children.
*/
func (self Instance) IsTransformNotificationEnabled() bool {
	return bool(class(self).IsTransformNotificationEnabled())
}

/*
Forces the transform to update. Transform changes in physics are not instant for performance reasons. Transforms are accumulated and then set. Use this if you need an up-to-date transform when doing physics operations.
*/
func (self Instance) ForceUpdateTransform() {
	class(self).ForceUpdateTransform()
}

/*
Assigns [param screen_point] as this node's new local transform.
*/
func (self Instance) MakeCanvasPositionLocal(screen_point Vector2.XY) Vector2.XY {
	return Vector2.XY(class(self).MakeCanvasPositionLocal(gd.Vector2(screen_point)))
}

/*
Transformations issued by [param event]'s inputs are applied in local space instead of global space.
*/
func (self Instance) MakeInputLocal(event [1]gdclass.InputEvent) [1]gdclass.InputEvent {
	return [1]gdclass.InputEvent(class(self).MakeInputLocal(event))
}

/*
Set/clear individual bits on the rendering visibility layer. This simplifies editing this [CanvasItem]'s visibility layer.
*/
func (self Instance) SetVisibilityLayerBit(layer int, enabled bool) {
	class(self).SetVisibilityLayerBit(gd.Int(layer), enabled)
}

/*
Returns an individual bit on the rendering visibility layer.
*/
func (self Instance) GetVisibilityLayerBit(layer int) bool {
	return bool(class(self).GetVisibilityLayerBit(gd.Int(layer)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CanvasItem

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CanvasItem"))
	casted := Instance{*(*gdclass.CanvasItem)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Visible() bool {
	return bool(class(self).IsVisible())
}

func (self Instance) SetVisible(value bool) {
	class(self).SetVisible(value)
}

func (self Instance) Modulate() Color.RGBA {
	return Color.RGBA(class(self).GetModulate())
}

func (self Instance) SetModulate(value Color.RGBA) {
	class(self).SetModulate(gd.Color(value))
}

func (self Instance) SelfModulate() Color.RGBA {
	return Color.RGBA(class(self).GetSelfModulate())
}

func (self Instance) SetSelfModulate(value Color.RGBA) {
	class(self).SetSelfModulate(gd.Color(value))
}

func (self Instance) ShowBehindParent() bool {
	return bool(class(self).IsDrawBehindParentEnabled())
}

func (self Instance) SetShowBehindParent(value bool) {
	class(self).SetDrawBehindParent(value)
}

func (self Instance) TopLevel() bool {
	return bool(class(self).IsSetAsTopLevel())
}

func (self Instance) SetTopLevel(value bool) {
	class(self).SetAsTopLevel(value)
}

func (self Instance) ClipChildren() gdclass.CanvasItemClipChildrenMode {
	return gdclass.CanvasItemClipChildrenMode(class(self).GetClipChildrenMode())
}

func (self Instance) SetClipChildren(value gdclass.CanvasItemClipChildrenMode) {
	class(self).SetClipChildrenMode(value)
}

func (self Instance) LightMask() int {
	return int(int(class(self).GetLightMask()))
}

func (self Instance) SetLightMask(value int) {
	class(self).SetLightMask(gd.Int(value))
}

func (self Instance) VisibilityLayer() int {
	return int(int(class(self).GetVisibilityLayer()))
}

func (self Instance) SetVisibilityLayer(value int) {
	class(self).SetVisibilityLayer(gd.Int(value))
}

func (self Instance) ZIndex() int {
	return int(int(class(self).GetZIndex()))
}

func (self Instance) SetZIndex(value int) {
	class(self).SetZIndex(gd.Int(value))
}

func (self Instance) ZAsRelative() bool {
	return bool(class(self).IsZRelative())
}

func (self Instance) SetZAsRelative(value bool) {
	class(self).SetZAsRelative(value)
}

func (self Instance) YSortEnabled() bool {
	return bool(class(self).IsYSortEnabled())
}

func (self Instance) SetYSortEnabled(value bool) {
	class(self).SetYSortEnabled(value)
}

func (self Instance) TextureFilter() gdclass.CanvasItemTextureFilter {
	return gdclass.CanvasItemTextureFilter(class(self).GetTextureFilter())
}

func (self Instance) SetTextureFilter(value gdclass.CanvasItemTextureFilter) {
	class(self).SetTextureFilter(value)
}

func (self Instance) TextureRepeat() gdclass.CanvasItemTextureRepeat {
	return gdclass.CanvasItemTextureRepeat(class(self).GetTextureRepeat())
}

func (self Instance) SetTextureRepeat(value gdclass.CanvasItemTextureRepeat) {
	class(self).SetTextureRepeat(value)
}

func (self Instance) Material() [1]gdclass.Material {
	return [1]gdclass.Material(class(self).GetMaterial())
}

func (self Instance) SetMaterial(value [1]gdclass.Material) {
	class(self).SetMaterial(value)
}

func (self Instance) UseParentMaterial() bool {
	return bool(class(self).GetUseParentMaterial())
}

func (self Instance) SetUseParentMaterial(value bool) {
	class(self).SetUseParentMaterial(value)
}

/*
Called when [CanvasItem] has been requested to redraw (after [method queue_redraw] is called, either manually or by the engine).
Corresponds to the [constant NOTIFICATION_DRAW] notification in [method Object._notification].
*/
func (class) _draw(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns the canvas item RID used by [RenderingServer] for this item.
*/
//go:nosplit
func (self class) GetCanvasItem() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_canvas_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisible(visible bool) {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_is_visible_in_tree, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Show the [CanvasItem] if it's currently hidden. This is equivalent to setting [member visible] to [code]true[/code]. For controls that inherit [Popup], the correct way to make them visible is to call one of the multiple [code]popup*()[/code] functions instead.
*/
//go:nosplit
func (self class) Show() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_show, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Hide the [CanvasItem] if it's currently visible. This is equivalent to setting [member visible] to [code]false[/code].
*/
//go:nosplit
func (self class) Hide() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_hide, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Queues the [CanvasItem] to redraw. During idle time, if [CanvasItem] is visible, [constant NOTIFICATION_DRAW] is sent and [method _draw] is called. This only occurs [b]once[/b] per frame, even if this method has been called multiple times.
*/
//go:nosplit
func (self class) QueueRedraw() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_queue_redraw, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves this node to display on top of its siblings.
Internally, the node is moved to the bottom of parent's child list. The method has no effect on nodes without a parent.
*/
//go:nosplit
func (self class) MoveToFront() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_move_to_front, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAsTopLevel(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_as_top_level, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSetAsTopLevel() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_is_set_as_top_level, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLightMask(light_mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, light_mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_light_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLightMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_light_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetModulate(modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetModulate() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelfModulate(self_modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, self_modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_self_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSelfModulate() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_self_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZIndex(z_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, z_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_z_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_z_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZAsRelative(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_z_as_relative, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsZRelative() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_is_z_relative, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetYSortEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_y_sort_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsYSortEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_is_y_sort_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawBehindParent(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_draw_behind_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawBehindParentEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_is_draw_behind_parent_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draws a line from a 2D point to another, with a given color and width. It can be optionally antialiased. See also [method draw_multiline] and [method draw_polyline].
If [param width] is negative, then a two-point primitive will be drawn instead of a four-point one. This means that when the CanvasItem is scaled, the line will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
//go:nosplit
func (self class) DrawLine(from gd.Vector2, to gd.Vector2, color gd.Color, width gd.Float, antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a dashed line from a 2D point to another, with a given color and width. See also [method draw_multiline] and [method draw_polyline].
If [param width] is negative, then a two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the line parts will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
If [param antialiased] is [code]true[/code], half transparent "feathers" will be attached to the boundary, making outlines smooth.
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
//go:nosplit
func (self class) DrawDashedLine(from gd.Vector2, to gd.Vector2, color gd.Color, width gd.Float, dash gd.Float, aligned bool, antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, dash)
	callframe.Arg(frame, aligned)
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_dashed_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws interconnected line segments with a uniform [param color] and [param width] and optional antialiasing (supported only for positive [param width]). When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw disconnected lines, use [method draw_multiline] instead. See also [method draw_polygon].
If [param width] is negative, it will be ignored and the polyline will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the polyline will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
//go:nosplit
func (self class) DrawPolyline(points gd.PackedVector2Array, color gd.Color, width gd.Float, antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_polyline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws interconnected line segments with a uniform [param width], point-by-point coloring, and optional antialiasing (supported only for positive [param width]). Colors assigned to line points match by index between [param points] and [param colors], i.e. each line segment is filled with a gradient between the colors of the endpoints. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw disconnected lines, use [method draw_multiline_colors] instead. See also [method draw_polygon].
If [param width] is negative, it will be ignored and the polyline will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the polyline will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
*/
//go:nosplit
func (self class) DrawPolylineColors(points gd.PackedVector2Array, colors gd.PackedColorArray, width gd.Float, antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	callframe.Arg(frame, pointers.Get(colors))
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_polyline_colors, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws an unfilled arc between the given angles with a uniform [param color] and [param width] and optional antialiasing (supported only for positive [param width]). The larger the value of [param point_count], the smoother the curve. See also [method draw_circle].
If [param width] is negative, it will be ignored and the arc will be drawn using [constant RenderingServer.PRIMITIVE_LINE_STRIP]. This means that when the CanvasItem is scaled, the arc will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
The arc is drawn from [param start_angle] towards the value of [param end_angle] so in clockwise direction if [code]start_angle < end_angle[/code] and counter-clockwise otherwise. Passing the same angles but in reversed order will produce the same arc. If absolute difference of [param start_angle] and [param end_angle] is greater than [constant @GDScript.TAU] radians, then a full circle arc is drawn (i.e. arc will not overlap itself).
*/
//go:nosplit
func (self class) DrawArc(center gd.Vector2, radius gd.Float, start_angle gd.Float, end_angle gd.Float, point_count gd.Int, color gd.Color, width gd.Float, antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, center)
	callframe.Arg(frame, radius)
	callframe.Arg(frame, start_angle)
	callframe.Arg(frame, end_angle)
	callframe.Arg(frame, point_count)
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_arc, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws multiple disconnected lines with a uniform [param width] and [param color]. Each line is defined by two consecutive points from [param points] array, i.e. i-th segment consists of [code]points[2 * i][/code], [code]points[2 * i + 1][/code] endpoints. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw interconnected lines, use [method draw_polyline] instead.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
//go:nosplit
func (self class) DrawMultiline(points gd.PackedVector2Array, color gd.Color, width gd.Float, antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_multiline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws multiple disconnected lines with a uniform [param width] and segment-by-segment coloring. Each segment is defined by two consecutive points from [param points] array and a corresponding color from [param colors] array, i.e. i-th segment consists of [code]points[2 * i][/code], [code]points[2 * i + 1][/code] endpoints and has [code]colors[i][/code] color. When drawing large amounts of lines, this is faster than using individual [method draw_line] calls. To draw interconnected lines, use [method draw_polyline_colors] instead.
If [param width] is negative, then two-point primitives will be drawn instead of a four-point ones. This means that when the CanvasItem is scaled, the lines will remain thin. If this behavior is not desired, then pass a positive [param width] like [code]1.0[/code].
[b]Note:[/b] [param antialiased] is only effective if [param width] is greater than [code]0.0[/code].
*/
//go:nosplit
func (self class) DrawMultilineColors(points gd.PackedVector2Array, colors gd.PackedColorArray, width gd.Float, antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	callframe.Arg(frame, pointers.Get(colors))
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_multiline_colors, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) DrawRect(rect gd.Rect2, color gd.Color, filled bool, width gd.Float, antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	callframe.Arg(frame, color)
	callframe.Arg(frame, filled)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) DrawCircle(position gd.Vector2, radius gd.Float, color gd.Color, filled bool, width gd.Float, antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, radius)
	callframe.Arg(frame, color)
	callframe.Arg(frame, filled)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_circle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a texture at a given position.
*/
//go:nosplit
func (self class) DrawTexture(texture [1]gdclass.Texture2D, position gd.Vector2, modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, position)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a textured rectangle at a given position, optionally modulated by a color. If [param transpose] is [code]true[/code], the texture will have its X and Y coordinates swapped. See also [method draw_rect] and [method draw_texture_rect_region].
*/
//go:nosplit
func (self class) DrawTextureRect(texture [1]gdclass.Texture2D, rect gd.Rect2, tile bool, modulate gd.Color, transpose bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, rect)
	callframe.Arg(frame, tile)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_texture_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a textured rectangle from a texture's region (specified by [param src_rect]) at a given position, optionally modulated by a color. If [param transpose] is [code]true[/code], the texture will have its X and Y coordinates swapped. See also [method draw_texture_rect].
*/
//go:nosplit
func (self class) DrawTextureRectRegion(texture [1]gdclass.Texture2D, rect gd.Rect2, src_rect gd.Rect2, modulate gd.Color, transpose bool, clip_uv bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, rect)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	callframe.Arg(frame, clip_uv)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_texture_rect_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a textured rectangle region of the multi-channel signed distance field texture at a given position, optionally modulated by a color. See [member FontFile.multichannel_signed_distance_field] for more information and caveats about MSDF font rendering.
If [param outline] is positive, each alpha channel value of pixel in region is set to maximum value of true distance in the [param outline] radius.
Value of the [param pixel_range] should the same that was used during distance field texture generation.
*/
//go:nosplit
func (self class) DrawMsdfTextureRectRegion(texture [1]gdclass.Texture2D, rect gd.Rect2, src_rect gd.Rect2, modulate gd.Color, outline gd.Float, pixel_range gd.Float, scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, rect)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, outline)
	callframe.Arg(frame, pixel_range)
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_msdf_texture_rect_region, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) DrawLcdTextureRectRegion(texture [1]gdclass.Texture2D, rect gd.Rect2, src_rect gd.Rect2, modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, rect)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_lcd_texture_rect_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a styled rectangle.
*/
//go:nosplit
func (self class) DrawStyleBox(style_box [1]gdclass.StyleBox, rect gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(style_box[0])[0])
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_style_box, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a custom primitive. 1 point for a point, 2 points for a line, 3 points for a triangle, and 4 points for a quad. If 0 points or more than 4 points are specified, nothing will be drawn and an error message will be printed. See also [method draw_line], [method draw_polyline], [method draw_polygon], and [method draw_rect].
*/
//go:nosplit
func (self class) DrawPrimitive(points gd.PackedVector2Array, colors gd.PackedColorArray, uvs gd.PackedVector2Array, texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	callframe.Arg(frame, pointers.Get(colors))
	callframe.Arg(frame, pointers.Get(uvs))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_primitive, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a solid polygon of any number of points, convex or concave. Unlike [method draw_colored_polygon], each point's color can be changed individually. See also [method draw_polyline] and [method draw_polyline_colors]. If you need more flexibility (such as being able to use bones), use [method RenderingServer.canvas_item_add_triangle_array] instead.
*/
//go:nosplit
func (self class) DrawPolygon(points gd.PackedVector2Array, colors gd.PackedColorArray, uvs gd.PackedVector2Array, texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	callframe.Arg(frame, pointers.Get(colors))
	callframe.Arg(frame, pointers.Get(uvs))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a colored polygon of any number of points, convex or concave. Unlike [method draw_polygon], a single color must be specified for the whole polygon.
*/
//go:nosplit
func (self class) DrawColoredPolygon(points gd.PackedVector2Array, color gd.Color, uvs gd.PackedVector2Array, texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	callframe.Arg(frame, color)
	callframe.Arg(frame, pointers.Get(uvs))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_colored_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) DrawString(font [1]gdclass.Font, pos gd.Vector2, text gd.String, alignment HorizontalAlignment, width gd.Float, font_size gd.Int, modulate gd.Color, justification_flags gdclass.TextServerJustificationFlag, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Breaks [param text] into lines and draws it using the specified [param font] at the [param pos] (top-left corner). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
//go:nosplit
func (self class) DrawMultilineString(font [1]gdclass.Font, pos gd.Vector2, text gd.String, alignment HorizontalAlignment, width gd.Float, font_size gd.Int, max_lines gd.Int, modulate gd.Color, brk_flags gdclass.TextServerLineBreakFlag, justification_flags gdclass.TextServerJustificationFlag, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, max_lines)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, brk_flags)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_multiline_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws [param text] outline using the specified [param font] at the [param pos] (bottom-left corner using the baseline of the font). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
//go:nosplit
func (self class) DrawStringOutline(font [1]gdclass.Font, pos gd.Vector2, text gd.String, alignment HorizontalAlignment, width gd.Float, font_size gd.Int, size gd.Int, modulate gd.Color, justification_flags gdclass.TextServerJustificationFlag, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_string_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Breaks [param text] to the lines and draws text outline using the specified [param font] at the [param pos] (top-left corner). The text will have its color multiplied by [param modulate]. If [param width] is greater than or equal to 0, the text will be clipped if it exceeds the specified width.
*/
//go:nosplit
func (self class) DrawMultilineStringOutline(font [1]gdclass.Font, pos gd.Vector2, text gd.String, alignment HorizontalAlignment, width gd.Float, font_size gd.Int, max_lines gd.Int, size gd.Int, modulate gd.Color, brk_flags gdclass.TextServerLineBreakFlag, justification_flags gdclass.TextServerJustificationFlag, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, pointers.Get(text))
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
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_multiline_string_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a string first character using a custom font.
*/
//go:nosplit
func (self class) DrawChar(font [1]gdclass.Font, pos gd.Vector2, char gd.String, font_size gd.Int, modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, pointers.Get(char))
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a string first character outline using a custom font.
*/
//go:nosplit
func (self class) DrawCharOutline(font [1]gdclass.Font, pos gd.Vector2, char gd.String, font_size gd.Int, size gd.Int, modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, pos)
	callframe.Arg(frame, pointers.Get(char))
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_char_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a [Mesh] in 2D, using the provided texture. See [MeshInstance2D] for related documentation.
*/
//go:nosplit
func (self class) DrawMesh(mesh [1]gdclass.Mesh, texture [1]gdclass.Texture2D, transform gd.Transform2D, modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, transform)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a [MultiMesh] in 2D with the provided texture. See [MultiMeshInstance2D] for related documentation.
*/
//go:nosplit
func (self class) DrawMultimesh(multimesh [1]gdclass.MultiMesh, texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(multimesh[0])[0])
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_multimesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets a custom transform for drawing via components. Anything drawn afterwards will be transformed by this.
[b]Note:[/b] [member FontFile.oversampling] does [i]not[/i] take [param scale] into account. This means that scaling up/down will cause bitmap fonts and rasterized (non-MSDF) dynamic fonts to appear blurry or pixelated. To ensure text remains crisp regardless of scale, you can enable MSDF font rendering by enabling [member ProjectSettings.gui/theme/default_font_multichannel_signed_distance_field] (applies to the default project font only), or enabling [b]Multichannel Signed Distance Field[/b] in the import options of a DynamicFont for custom fonts. On system fonts, [member SystemFont.multichannel_signed_distance_field] can be enabled in the inspector.
*/
//go:nosplit
func (self class) DrawSetTransform(position gd.Vector2, rotation gd.Float, scale gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, rotation)
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets a custom transform for drawing via matrix. Anything drawn afterwards will be transformed by this.
*/
//go:nosplit
func (self class) DrawSetTransformMatrix(xform gd.Transform2D) {
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_set_transform_matrix, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Subsequent drawing commands will be ignored unless they fall within the specified animation slice. This is a faster way to implement animations that loop on background rather than redrawing constantly.
*/
//go:nosplit
func (self class) DrawAnimationSlice(animation_length gd.Float, slice_begin gd.Float, slice_end gd.Float, offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, animation_length)
	callframe.Arg(frame, slice_begin)
	callframe.Arg(frame, slice_end)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_animation_slice, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
After submitting all animations slices via [method draw_animation_slice], this function can be used to revert drawing to its default state (all subsequent drawing commands will be visible). If you don't care about this particular use case, usage of this function after submitting the slices is not required.
*/
//go:nosplit
func (self class) DrawEndAnimation() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_draw_end_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the transform matrix of this item.
*/
//go:nosplit
func (self class) GetTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the global transform matrix of this item, i.e. the combined transform up to the topmost [CanvasItem] node. The topmost item is a [CanvasItem] that either has no parent, has non-[CanvasItem] parent or it has [member top_level] enabled.
*/
//go:nosplit
func (self class) GetGlobalTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_global_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the transform from the local coordinate system of this [CanvasItem] to the [Viewport]s coordinate system.
*/
//go:nosplit
func (self class) GetGlobalTransformWithCanvas() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_global_transform_with_canvas, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the transform from the coordinate system of the canvas, this item is in, to the [Viewport]s embedders coordinate system.
*/
//go:nosplit
func (self class) GetViewportTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_viewport_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the viewport's boundaries as a [Rect2].
*/
//go:nosplit
func (self class) GetViewportRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_viewport_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the transform from the coordinate system of the canvas, this item is in, to the [Viewport]s coordinate system.
*/
//go:nosplit
func (self class) GetCanvasTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_screen_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the mouse's position in this [CanvasItem] using the local coordinate system of this [CanvasItem].
*/
//go:nosplit
func (self class) GetLocalMousePosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_local_mouse_position, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_global_mouse_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [RID] of the [World2D] canvas where this item is in.
*/
//go:nosplit
func (self class) GetCanvas() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_canvas, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [CanvasLayer] that contains this node, or [code]null[/code] if the node is not in any [CanvasLayer].
*/
//go:nosplit
func (self class) GetCanvasLayerNode() [1]gdclass.CanvasLayer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_canvas_layer_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CanvasLayer{gd.PointerLifetimeBoundTo[gdclass.CanvasLayer](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [World2D] where this item is in.
*/
//go:nosplit
func (self class) GetWorld2d() [1]gdclass.World2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_world_2d, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.World2D{gd.PointerWithOwnershipTransferredToGo[gdclass.World2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaterial(material [1]gdclass.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaterial() [1]gdclass.Material {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseParentMaterial(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_use_parent_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseParentMaterial() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_use_parent_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enable] is [code]true[/code], this node will receive [constant NOTIFICATION_LOCAL_TRANSFORM_CHANGED] when its local transform changes.
*/
//go:nosplit
func (self class) SetNotifyLocalTransform(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_notify_local_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if local transform notifications are communicated to children.
*/
//go:nosplit
func (self class) IsLocalTransformNotificationEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_is_local_transform_notification_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enable] is [code]true[/code], this node will receive [constant NOTIFICATION_TRANSFORM_CHANGED] when its global transform changes.
*/
//go:nosplit
func (self class) SetNotifyTransform(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_notify_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if global transform notifications are communicated to children.
*/
//go:nosplit
func (self class) IsTransformNotificationEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_is_transform_notification_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Forces the transform to update. Transform changes in physics are not instant for performance reasons. Transforms are accumulated and then set. Use this if you need an up-to-date transform when doing physics operations.
*/
//go:nosplit
func (self class) ForceUpdateTransform() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_force_update_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Assigns [param screen_point] as this node's new local transform.
*/
//go:nosplit
func (self class) MakeCanvasPositionLocal(screen_point gd.Vector2) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_make_canvas_position_local, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Transformations issued by [param event]'s inputs are applied in local space instead of global space.
*/
//go:nosplit
func (self class) MakeInputLocal(event [1]gdclass.InputEvent) [1]gdclass.InputEvent {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_make_input_local, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.InputEvent{gd.PointerWithOwnershipTransferredToGo[gdclass.InputEvent](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityLayer(layer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_visibility_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityLayer() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_visibility_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set/clear individual bits on the rendering visibility layer. This simplifies editing this [CanvasItem]'s visibility layer.
*/
//go:nosplit
func (self class) SetVisibilityLayerBit(layer gd.Int, enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_visibility_layer_bit, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an individual bit on the rendering visibility layer.
*/
//go:nosplit
func (self class) GetVisibilityLayerBit(layer gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_visibility_layer_bit, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureFilter(mode gdclass.CanvasItemTextureFilter) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureFilter() gdclass.CanvasItemTextureFilter {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CanvasItemTextureFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureRepeat(mode gdclass.CanvasItemTextureRepeat) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureRepeat() gdclass.CanvasItemTextureRepeat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CanvasItemTextureRepeat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClipChildrenMode(mode gdclass.CanvasItemClipChildrenMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_set_clip_children_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetClipChildrenMode() gdclass.CanvasItemClipChildrenMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CanvasItemClipChildrenMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItem.Bind_get_clip_children_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnDraw(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("draw"), gd.NewCallable(cb), 0)
}

func (self Instance) OnVisibilityChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("visibility_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnHidden(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("hidden"), gd.NewCallable(cb), 0)
}

func (self Instance) OnItemRectChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("item_rect_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsCanvasItem() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCanvasItem() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced     { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance  { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_draw":
		return reflect.ValueOf(self._draw)
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_draw":
		return reflect.ValueOf(self._draw)
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("CanvasItem", func(ptr gd.Object) any { return [1]gdclass.CanvasItem{*(*gdclass.CanvasItem)(unsafe.Pointer(&ptr))} })
}

type TextureFilter = gdclass.CanvasItemTextureFilter

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

type TextureRepeat = gdclass.CanvasItemTextureRepeat

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

type ClipChildrenMode = gdclass.CanvasItemClipChildrenMode

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

type HorizontalAlignment int

const (
	/*Horizontal left alignment, usually for text-derived classes.*/
	HorizontalAlignmentLeft HorizontalAlignment = 0
	/*Horizontal center alignment, usually for text-derived classes.*/
	HorizontalAlignmentCenter HorizontalAlignment = 1
	/*Horizontal right alignment, usually for text-derived classes.*/
	HorizontalAlignmentRight HorizontalAlignment = 2
	/*Expand row to fit width, usually for text-derived classes.*/
	HorizontalAlignmentFill HorizontalAlignment = 3
)
