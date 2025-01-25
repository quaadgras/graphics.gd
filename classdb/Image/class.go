// Package Image provides methods for working with Image object instances.
package Image

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Rect2i"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Native image datatype. Contains image data which can be converted to an [ImageTexture] and provides commonly used [i]image processing[/i] methods. The maximum width and height for an [Image] are [constant MAX_WIDTH] and [constant MAX_HEIGHT].
An [Image] cannot be assigned to a texture property of an object directly (such as [member Sprite2D.texture]), and has to be converted manually to an [ImageTexture] first.
[b]Note:[/b] The maximum image size is 16384×16384 pixels due to graphics hardware limitations. Larger images may fail to import.
*/
type Instance [1]gdclass.Image

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsImage() Instance
}

/*
Returns the image's width.
*/
func (self Instance) GetWidth() int {
	return int(int(class(self).GetWidth()))
}

/*
Returns the image's height.
*/
func (self Instance) GetHeight() int {
	return int(int(class(self).GetHeight()))
}

/*
Returns the image's size (width and height).
*/
func (self Instance) GetSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetSize())
}

/*
Returns [code]true[/code] if the image has generated mipmaps.
*/
func (self Instance) HasMipmaps() bool {
	return bool(class(self).HasMipmaps())
}

/*
Returns the image's format. See [enum Format] constants.
*/
func (self Instance) GetFormat() gdclass.ImageFormat {
	return gdclass.ImageFormat(class(self).GetFormat())
}

/*
Returns a copy of the image's raw data.
*/
func (self Instance) GetData() []byte {
	return []byte(class(self).GetData().Bytes())
}

/*
Returns size (in bytes) of the image's raw data.
*/
func (self Instance) GetDataSize() int {
	return int(int(class(self).GetDataSize()))
}

/*
Converts the image's format. See [enum Format] constants.
*/
func (self Instance) Convert(format gdclass.ImageFormat) {
	class(self).Convert(format)
}

/*
Returns the number of mipmap levels or 0 if the image has no mipmaps. The largest main level image is not counted as a mipmap level by this method, so if you want to include it you can add 1 to this count.
*/
func (self Instance) GetMipmapCount() int {
	return int(int(class(self).GetMipmapCount()))
}

/*
Returns the offset where the image's mipmap with index [param mipmap] is stored in the [member data] dictionary.
*/
func (self Instance) GetMipmapOffset(mipmap int) int {
	return int(int(class(self).GetMipmapOffset(gd.Int(mipmap))))
}

/*
Resizes the image to the nearest power of 2 for the width and height. If [param square] is [code]true[/code] then set width and height to be the same. New pixels are calculated using the [param interpolation] mode defined via [enum Interpolation] constants.
*/
func (self Instance) ResizeToPo2() {
	class(self).ResizeToPo2(false, 1)
}

/*
Resizes the image to the given [param width] and [param height]. New pixels are calculated using the [param interpolation] mode defined via [enum Interpolation] constants.
*/
func (self Instance) Resize(width int, height int) {
	class(self).Resize(gd.Int(width), gd.Int(height), 1)
}

/*
Shrinks the image by a factor of 2 on each axis (this divides the pixel count by 4).
*/
func (self Instance) ShrinkX2() {
	class(self).ShrinkX2()
}

/*
Crops the image to the given [param width] and [param height]. If the specified size is larger than the current size, the extra area is filled with black pixels.
*/
func (self Instance) Crop(width int, height int) {
	class(self).Crop(gd.Int(width), gd.Int(height))
}

/*
Flips the image horizontally.
*/
func (self Instance) FlipX() {
	class(self).FlipX()
}

/*
Flips the image vertically.
*/
func (self Instance) FlipY() {
	class(self).FlipY()
}

/*
Generates mipmaps for the image. Mipmaps are precalculated lower-resolution copies of the image that are automatically used if the image needs to be scaled down when rendered. They help improve image quality and performance when rendering. This method returns an error if the image is compressed, in a custom format, or if the image's width/height is [code]0[/code]. Enabling [param renormalize] when generating mipmaps for normal map textures will make sure all resulting vector values are normalized.
It is possible to check if the image has mipmaps by calling [method has_mipmaps] or [method get_mipmap_count]. Calling [method generate_mipmaps] on an image that already has mipmaps will replace existing mipmaps in the image.
*/
func (self Instance) GenerateMipmaps() error {
	return error(gd.ToError(class(self).GenerateMipmaps(false)))
}

/*
Removes the image's mipmaps.
*/
func (self Instance) ClearMipmaps() {
	class(self).ClearMipmaps()
}

/*
Creates an empty image of given size and format. See [enum Format] constants. If [param use_mipmaps] is [code]true[/code], then generate mipmaps for this image. See the [method generate_mipmaps].
*/
func Create(width int, height int, use_mipmaps bool, format gdclass.ImageFormat) [1]gdclass.Image {
	self := Instance{}
	return [1]gdclass.Image(class(self).Create(gd.Int(width), gd.Int(height), use_mipmaps, format))
}

/*
Creates an empty image of given size and format. See [enum Format] constants. If [param use_mipmaps] is [code]true[/code], then generate mipmaps for this image. See the [method generate_mipmaps].
*/
func CreateEmpty(width int, height int, use_mipmaps bool, format gdclass.ImageFormat) [1]gdclass.Image {
	self := Instance{}
	return [1]gdclass.Image(class(self).CreateEmpty(gd.Int(width), gd.Int(height), use_mipmaps, format))
}

/*
Creates a new image of given size and format. See [enum Format] constants. Fills the image with the given raw data. If [param use_mipmaps] is [code]true[/code] then loads mipmaps for this image from [param data]. See [method generate_mipmaps].
*/
func CreateFromData(width int, height int, use_mipmaps bool, format gdclass.ImageFormat, data []byte) [1]gdclass.Image {
	self := Instance{}
	return [1]gdclass.Image(class(self).CreateFromData(gd.Int(width), gd.Int(height), use_mipmaps, format, gd.NewPackedByteSlice(data)))
}

/*
Overwrites data of an existing [Image]. Non-static equivalent of [method create_from_data].
*/
func (self Instance) SetData(width int, height int, use_mipmaps bool, format gdclass.ImageFormat, data []byte) {
	class(self).SetData(gd.Int(width), gd.Int(height), use_mipmaps, format, gd.NewPackedByteSlice(data))
}

/*
Returns [code]true[/code] if the image has no data.
*/
func (self Instance) IsEmpty() bool {
	return bool(class(self).IsEmpty())
}

/*
Loads an image from file [param path]. See [url=$DOCS_URL/tutorials/assets_pipeline/importing_images.html#supported-image-formats]Supported image formats[/url] for a list of supported image formats and limitations.
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external images at run-time, such as images located at the [code]user://[/code] directory, and may not work in exported projects.
See also [ImageTexture] description for usage examples.
*/
func (self Instance) Load(path string) error {
	return error(gd.ToError(class(self).Load(gd.NewString(path))))
}

/*
Creates a new [Image] and loads data from the specified file.
*/
func LoadFromFile(path string) [1]gdclass.Image {
	self := Instance{}
	return [1]gdclass.Image(class(self).LoadFromFile(gd.NewString(path)))
}

/*
Saves the image as a PNG file to the file at [param path].
*/
func (self Instance) SavePng(path string) error {
	return error(gd.ToError(class(self).SavePng(gd.NewString(path))))
}

/*
Saves the image as a PNG file to a byte array.
*/
func (self Instance) SavePngToBuffer() []byte {
	return []byte(class(self).SavePngToBuffer().Bytes())
}

/*
Saves the image as a JPEG file to [param path] with the specified [param quality] between [code]0.01[/code] and [code]1.0[/code] (inclusive). Higher [param quality] values result in better-looking output at the cost of larger file sizes. Recommended [param quality] values are between [code]0.75[/code] and [code]0.90[/code]. Even at quality [code]1.00[/code], JPEG compression remains lossy.
[b]Note:[/b] JPEG does not save an alpha channel. If the [Image] contains an alpha channel, the image will still be saved, but the resulting JPEG file won't contain the alpha channel.
*/
func (self Instance) SaveJpg(path string) error {
	return error(gd.ToError(class(self).SaveJpg(gd.NewString(path), gd.Float(0.75))))
}

/*
Saves the image as a JPEG file to a byte array with the specified [param quality] between [code]0.01[/code] and [code]1.0[/code] (inclusive). Higher [param quality] values result in better-looking output at the cost of larger byte array sizes (and therefore memory usage). Recommended [param quality] values are between [code]0.75[/code] and [code]0.90[/code]. Even at quality [code]1.00[/code], JPEG compression remains lossy.
[b]Note:[/b] JPEG does not save an alpha channel. If the [Image] contains an alpha channel, the image will still be saved, but the resulting byte array won't contain the alpha channel.
*/
func (self Instance) SaveJpgToBuffer() []byte {
	return []byte(class(self).SaveJpgToBuffer(gd.Float(0.75)).Bytes())
}

/*
Saves the image as an EXR file to [param path]. If [param grayscale] is [code]true[/code] and the image has only one channel, it will be saved explicitly as monochrome rather than one red channel. This function will return [constant ERR_UNAVAILABLE] if Godot was compiled without the TinyEXR module.
[b]Note:[/b] The TinyEXR module is disabled in non-editor builds, which means [method save_exr] will return [constant ERR_UNAVAILABLE] when it is called from an exported project.
*/
func (self Instance) SaveExr(path string) error {
	return error(gd.ToError(class(self).SaveExr(gd.NewString(path), false)))
}

/*
Saves the image as an EXR file to a byte array. If [param grayscale] is [code]true[/code] and the image has only one channel, it will be saved explicitly as monochrome rather than one red channel. This function will return an empty byte array if Godot was compiled without the TinyEXR module.
[b]Note:[/b] The TinyEXR module is disabled in non-editor builds, which means [method save_exr] will return an empty byte array when it is called from an exported project.
*/
func (self Instance) SaveExrToBuffer() []byte {
	return []byte(class(self).SaveExrToBuffer(false).Bytes())
}

/*
Saves the image as a WebP (Web Picture) file to the file at [param path]. By default it will save lossless. If [param lossy] is true, the image will be saved lossy, using the [param quality] setting between 0.0 and 1.0 (inclusive). Lossless WebP offers more efficient compression than PNG.
[b]Note:[/b] The WebP format is limited to a size of 16383×16383 pixels, while PNG can save larger images.
*/
func (self Instance) SaveWebp(path string) error {
	return error(gd.ToError(class(self).SaveWebp(gd.NewString(path), false, gd.Float(0.75))))
}

/*
Saves the image as a WebP (Web Picture) file to a byte array. By default it will save lossless. If [param lossy] is true, the image will be saved lossy, using the [param quality] setting between 0.0 and 1.0 (inclusive). Lossless WebP offers more efficient compression than PNG.
[b]Note:[/b] The WebP format is limited to a size of 16383×16383 pixels, while PNG can save larger images.
*/
func (self Instance) SaveWebpToBuffer() []byte {
	return []byte(class(self).SaveWebpToBuffer(false, gd.Float(0.75)).Bytes())
}

/*
Returns [constant ALPHA_BLEND] if the image has data for alpha values. Returns [constant ALPHA_BIT] if all the alpha values are stored in a single bit. Returns [constant ALPHA_NONE] if no data for alpha values is found.
*/
func (self Instance) DetectAlpha() gdclass.ImageAlphaMode {
	return gdclass.ImageAlphaMode(class(self).DetectAlpha())
}

/*
Returns [code]true[/code] if all the image's pixels have an alpha value of 0. Returns [code]false[/code] if any pixel has an alpha value higher than 0.
*/
func (self Instance) IsInvisible() bool {
	return bool(class(self).IsInvisible())
}

/*
Returns the color channels used by this image, as one of the [enum UsedChannels] constants. If the image is compressed, the original [param source] must be specified.
*/
func (self Instance) DetectUsedChannels() gdclass.ImageUsedChannels {
	return gdclass.ImageUsedChannels(class(self).DetectUsedChannels(0))
}

/*
Compresses the image to use less memory. Can not directly access pixel data while the image is compressed. Returns error if the chosen compression mode is not available.
The [param source] parameter helps to pick the best compression method for DXT and ETC2 formats. It is ignored for ASTC compression.
For ASTC compression, the [param astc_format] parameter must be supplied.
*/
func (self Instance) Compress(mode gdclass.ImageCompressMode) error {
	return error(gd.ToError(class(self).Compress(mode, 0, 0)))
}

/*
Compresses the image to use less memory. Can not directly access pixel data while the image is compressed. Returns error if the chosen compression mode is not available.
This is an alternative to [method compress] that lets the user supply the channels used in order for the compressor to pick the best DXT and ETC2 formats. For other formats (non DXT or ETC2), this argument is ignored.
For ASTC compression, the [param astc_format] parameter must be supplied.
*/
func (self Instance) CompressFromChannels(mode gdclass.ImageCompressMode, channels gdclass.ImageUsedChannels) error {
	return error(gd.ToError(class(self).CompressFromChannels(mode, channels, 0)))
}

/*
Decompresses the image if it is VRAM compressed in a supported format. Returns [constant OK] if the format is supported, otherwise [constant ERR_UNAVAILABLE].
[b]Note:[/b] The following formats can be decompressed: DXT, RGTC, BPTC. The formats ETC1 and ETC2 are not supported.
*/
func (self Instance) Decompress() error {
	return error(gd.ToError(class(self).Decompress()))
}

/*
Returns [code]true[/code] if the image is compressed.
*/
func (self Instance) IsCompressed() bool {
	return bool(class(self).IsCompressed())
}

/*
Rotates the image in the specified [param direction] by [code]90[/code] degrees. The width and height of the image must be greater than [code]1[/code]. If the width and height are not equal, the image will be resized.
*/
func (self Instance) Rotate90(direction ClockDirection) {
	class(self).Rotate90(direction)
}

/*
Rotates the image by [code]180[/code] degrees. The width and height of the image must be greater than [code]1[/code].
*/
func (self Instance) Rotate180() {
	class(self).Rotate180()
}

/*
Blends low-alpha pixels with nearby pixels.
*/
func (self Instance) FixAlphaEdges() {
	class(self).FixAlphaEdges()
}

/*
Multiplies color values with alpha values. Resulting color values for a pixel are [code](color * alpha)/256[/code]. See also [member CanvasItemMaterial.blend_mode].
*/
func (self Instance) PremultiplyAlpha() {
	class(self).PremultiplyAlpha()
}

/*
Converts the raw data from the sRGB colorspace to a linear scale.
*/
func (self Instance) SrgbToLinear() {
	class(self).SrgbToLinear()
}

/*
Converts the image's data to represent coordinates on a 3D plane. This is used when the image represents a normal map. A normal map can add lots of detail to a 3D surface without increasing the polygon count.
*/
func (self Instance) NormalMapToXy() {
	class(self).NormalMapToXy()
}

/*
Converts a standard RGBE (Red Green Blue Exponent) image to an sRGB image.
*/
func (self Instance) RgbeToSrgb() [1]gdclass.Image {
	return [1]gdclass.Image(class(self).RgbeToSrgb())
}

/*
Converts a bump map to a normal map. A bump map provides a height offset per-pixel, while a normal map provides a normal direction per pixel.
*/
func (self Instance) BumpMapToNormalMap() {
	class(self).BumpMapToNormalMap(gd.Float(1.0))
}

/*
Compute image metrics on the current image and the compared image.
The dictionary contains [code]max[/code], [code]mean[/code], [code]mean_squared[/code], [code]root_mean_squared[/code] and [code]peak_snr[/code].
*/
func (self Instance) ComputeImageMetrics(compared_image [1]gdclass.Image, use_luma bool) map[any]any {
	return map[any]any(gd.DictionaryAs[any, any](class(self).ComputeImageMetrics(compared_image, use_luma)))
}

/*
Copies [param src_rect] from [param src] image to this image at coordinates [param dst], clipped accordingly to both image bounds. This image and [param src] image [b]must[/b] have the same format. [param src_rect] with non-positive size is treated as empty.
*/
func (self Instance) BlitRect(src [1]gdclass.Image, src_rect Rect2i.PositionSize, dst Vector2i.XY) {
	class(self).BlitRect(src, gd.Rect2i(src_rect), gd.Vector2i(dst))
}

/*
Blits [param src_rect] area from [param src] image to this image at the coordinates given by [param dst], clipped accordingly to both image bounds. [param src] pixel is copied onto [param dst] if the corresponding [param mask] pixel's alpha value is not 0. This image and [param src] image [b]must[/b] have the same format. [param src] image and [param mask] image [b]must[/b] have the same size (width and height) but they can have different formats. [param src_rect] with non-positive size is treated as empty.
*/
func (self Instance) BlitRectMask(src [1]gdclass.Image, mask [1]gdclass.Image, src_rect Rect2i.PositionSize, dst Vector2i.XY) {
	class(self).BlitRectMask(src, mask, gd.Rect2i(src_rect), gd.Vector2i(dst))
}

/*
Alpha-blends [param src_rect] from [param src] image to this image at coordinates [param dst], clipped accordingly to both image bounds. This image and [param src] image [b]must[/b] have the same format. [param src_rect] with non-positive size is treated as empty.
*/
func (self Instance) BlendRect(src [1]gdclass.Image, src_rect Rect2i.PositionSize, dst Vector2i.XY) {
	class(self).BlendRect(src, gd.Rect2i(src_rect), gd.Vector2i(dst))
}

/*
Alpha-blends [param src_rect] from [param src] image to this image using [param mask] image at coordinates [param dst], clipped accordingly to both image bounds. Alpha channels are required for both [param src] and [param mask]. [param dst] pixels and [param src] pixels will blend if the corresponding mask pixel's alpha value is not 0. This image and [param src] image [b]must[/b] have the same format. [param src] image and [param mask] image [b]must[/b] have the same size (width and height) but they can have different formats. [param src_rect] with non-positive size is treated as empty.
*/
func (self Instance) BlendRectMask(src [1]gdclass.Image, mask [1]gdclass.Image, src_rect Rect2i.PositionSize, dst Vector2i.XY) {
	class(self).BlendRectMask(src, mask, gd.Rect2i(src_rect), gd.Vector2i(dst))
}

/*
Fills the image with [param color].
*/
func (self Instance) Fill(color Color.RGBA) {
	class(self).Fill(gd.Color(color))
}

/*
Fills [param rect] with [param color].
*/
func (self Instance) FillRect(rect Rect2i.PositionSize, color Color.RGBA) {
	class(self).FillRect(gd.Rect2i(rect), gd.Color(color))
}

/*
Returns a [Rect2i] enclosing the visible portion of the image, considering each pixel with a non-zero alpha channel as visible.
*/
func (self Instance) GetUsedRect() Rect2i.PositionSize {
	return Rect2i.PositionSize(class(self).GetUsedRect())
}

/*
Returns a new [Image] that is a copy of this [Image]'s area specified with [param region].
*/
func (self Instance) GetRegion(region Rect2i.PositionSize) [1]gdclass.Image {
	return [1]gdclass.Image(class(self).GetRegion(gd.Rect2i(region)))
}

/*
Copies [param src] image to this image.
*/
func (self Instance) CopyFrom(src [1]gdclass.Image) {
	class(self).CopyFrom(src)
}

/*
Returns the color of the pixel at [param point].
This is the same as [method get_pixel], but with a [Vector2i] argument instead of two integer arguments.
*/
func (self Instance) GetPixelv(point Vector2i.XY) Color.RGBA {
	return Color.RGBA(class(self).GetPixelv(gd.Vector2i(point)))
}

/*
Returns the color of the pixel at [code](x, y)[/code].
This is the same as [method get_pixelv], but with two integer arguments instead of a [Vector2i] argument.
*/
func (self Instance) GetPixel(x int, y int) Color.RGBA {
	return Color.RGBA(class(self).GetPixel(gd.Int(x), gd.Int(y)))
}

/*
Sets the [Color] of the pixel at [param point] to [param color].
[b]Example:[/b]
[codeblocks]
[gdscript]
var img_width = 10
var img_height = 5
var img = Image.create(img_width, img_height, false, Image.FORMAT_RGBA8)

img.set_pixelv(Vector2i(1, 2), Color.RED) # Sets the color at (1, 2) to red.
[/gdscript]
[csharp]
int imgWidth = 10;
int imgHeight = 5;
var img = Image.Create(imgWidth, imgHeight, false, Image.Format.Rgba8);

img.SetPixelv(new Vector2I(1, 2), Colors.Red); // Sets the color at (1, 2) to red.
[/csharp]
[/codeblocks]
This is the same as [method set_pixel], but with a [Vector2i] argument instead of two integer arguments.
*/
func (self Instance) SetPixelv(point Vector2i.XY, color Color.RGBA) {
	class(self).SetPixelv(gd.Vector2i(point), gd.Color(color))
}

/*
Sets the [Color] of the pixel at [code](x, y)[/code] to [param color].
[b]Example:[/b]
[codeblocks]
[gdscript]
var img_width = 10
var img_height = 5
var img = Image.create(img_width, img_height, false, Image.FORMAT_RGBA8)

img.set_pixel(1, 2, Color.RED) # Sets the color at (1, 2) to red.
[/gdscript]
[csharp]
int imgWidth = 10;
int imgHeight = 5;
var img = Image.Create(imgWidth, imgHeight, false, Image.Format.Rgba8);

img.SetPixel(1, 2, Colors.Red); // Sets the color at (1, 2) to red.
[/csharp]
[/codeblocks]
This is the same as [method set_pixelv], but with a two integer arguments instead of a [Vector2i] argument.
*/
func (self Instance) SetPixel(x int, y int, color Color.RGBA) {
	class(self).SetPixel(gd.Int(x), gd.Int(y), gd.Color(color))
}

/*
Adjusts this image's [param brightness], [param contrast], and [param saturation] by the given values. Does not work if the image is compressed (see [method is_compressed]).
*/
func (self Instance) AdjustBcs(brightness Float.X, contrast Float.X, saturation Float.X) {
	class(self).AdjustBcs(gd.Float(brightness), gd.Float(contrast), gd.Float(saturation))
}

/*
Loads an image from the binary contents of a PNG file.
*/
func (self Instance) LoadPngFromBuffer(buffer []byte) error {
	return error(gd.ToError(class(self).LoadPngFromBuffer(gd.NewPackedByteSlice(buffer))))
}

/*
Loads an image from the binary contents of a JPEG file.
*/
func (self Instance) LoadJpgFromBuffer(buffer []byte) error {
	return error(gd.ToError(class(self).LoadJpgFromBuffer(gd.NewPackedByteSlice(buffer))))
}

/*
Loads an image from the binary contents of a WebP file.
*/
func (self Instance) LoadWebpFromBuffer(buffer []byte) error {
	return error(gd.ToError(class(self).LoadWebpFromBuffer(gd.NewPackedByteSlice(buffer))))
}

/*
Loads an image from the binary contents of a TGA file.
[b]Note:[/b] This method is only available in engine builds with the TGA module enabled. By default, the TGA module is enabled, but it can be disabled at build-time using the [code]module_tga_enabled=no[/code] SCons option.
*/
func (self Instance) LoadTgaFromBuffer(buffer []byte) error {
	return error(gd.ToError(class(self).LoadTgaFromBuffer(gd.NewPackedByteSlice(buffer))))
}

/*
Loads an image from the binary contents of a BMP file.
[b]Note:[/b] Godot's BMP module doesn't support 16-bit per pixel images. Only 1-bit, 4-bit, 8-bit, 24-bit, and 32-bit per pixel images are supported.
[b]Note:[/b] This method is only available in engine builds with the BMP module enabled. By default, the BMP module is enabled, but it can be disabled at build-time using the [code]module_bmp_enabled=no[/code] SCons option.
*/
func (self Instance) LoadBmpFromBuffer(buffer []byte) error {
	return error(gd.ToError(class(self).LoadBmpFromBuffer(gd.NewPackedByteSlice(buffer))))
}

/*
Loads an image from the binary contents of a [url=https://github.com/KhronosGroup/KTX-Software]KTX[/url] file. Unlike most image formats, KTX can store VRAM-compressed data and embed mipmaps.
[b]Note:[/b] Godot's libktx implementation only supports 2D images. Cubemaps, texture arrays, and de-padding are not supported.
[b]Note:[/b] This method is only available in engine builds with the KTX module enabled. By default, the KTX module is enabled, but it can be disabled at build-time using the [code]module_ktx_enabled=no[/code] SCons option.
*/
func (self Instance) LoadKtxFromBuffer(buffer []byte) error {
	return error(gd.ToError(class(self).LoadKtxFromBuffer(gd.NewPackedByteSlice(buffer))))
}

/*
Loads an image from the UTF-8 binary contents of an [b]uncompressed[/b] SVG file ([b].svg[/b]).
[b]Note:[/b] Beware when using compressed SVG files (like [b].svgz[/b]), they need to be [code]decompressed[/code] before loading.
[b]Note:[/b] This method is only available in engine builds with the SVG module enabled. By default, the SVG module is enabled, but it can be disabled at build-time using the [code]module_svg_enabled=no[/code] SCons option.
*/
func (self Instance) LoadSvgFromBuffer(buffer []byte) error {
	return error(gd.ToError(class(self).LoadSvgFromBuffer(gd.NewPackedByteSlice(buffer), gd.Float(1.0))))
}

/*
Loads an image from the string contents of an SVG file ([b].svg[/b]).
[b]Note:[/b] This method is only available in engine builds with the SVG module enabled. By default, the SVG module is enabled, but it can be disabled at build-time using the [code]module_svg_enabled=no[/code] SCons option.
*/
func (self Instance) LoadSvgFromString(svg_str string) error {
	return error(gd.ToError(class(self).LoadSvgFromString(gd.NewString(svg_str), gd.Float(1.0))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Image

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Image"))
	casted := Instance{*(*gdclass.Image)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the image's width.
*/
//go:nosplit
func (self class) GetWidth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the image's height.
*/
//go:nosplit
func (self class) GetHeight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the image's size (width and height).
*/
//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the image has generated mipmaps.
*/
//go:nosplit
func (self class) HasMipmaps() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_has_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the image's format. See [enum Format] constants.
*/
//go:nosplit
func (self class) GetFormat() gdclass.ImageFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ImageFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a copy of the image's raw data.
*/
//go:nosplit
func (self class) GetData() gd.PackedByteArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns size (in bytes) of the image's raw data.
*/
//go:nosplit
func (self class) GetDataSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_data_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Converts the image's format. See [enum Format] constants.
*/
//go:nosplit
func (self class) Convert(format gdclass.ImageFormat) {
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_convert, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of mipmap levels or 0 if the image has no mipmaps. The largest main level image is not counted as a mipmap level by this method, so if you want to include it you can add 1 to this count.
*/
//go:nosplit
func (self class) GetMipmapCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_mipmap_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the offset where the image's mipmap with index [param mipmap] is stored in the [member data] dictionary.
*/
//go:nosplit
func (self class) GetMipmapOffset(mipmap gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, mipmap)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_mipmap_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Resizes the image to the nearest power of 2 for the width and height. If [param square] is [code]true[/code] then set width and height to be the same. New pixels are calculated using the [param interpolation] mode defined via [enum Interpolation] constants.
*/
//go:nosplit
func (self class) ResizeToPo2(square bool, interpolation gdclass.ImageInterpolation) {
	var frame = callframe.New()
	callframe.Arg(frame, square)
	callframe.Arg(frame, interpolation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_resize_to_po2, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Resizes the image to the given [param width] and [param height]. New pixels are calculated using the [param interpolation] mode defined via [enum Interpolation] constants.
*/
//go:nosplit
func (self class) Resize(width gd.Int, height gd.Int, interpolation gdclass.ImageInterpolation) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, interpolation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_resize, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Shrinks the image by a factor of 2 on each axis (this divides the pixel count by 4).
*/
//go:nosplit
func (self class) ShrinkX2() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_shrink_x2, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Crops the image to the given [param width] and [param height]. If the specified size is larger than the current size, the extra area is filled with black pixels.
*/
//go:nosplit
func (self class) Crop(width gd.Int, height gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_crop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Flips the image horizontally.
*/
//go:nosplit
func (self class) FlipX() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_flip_x, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Flips the image vertically.
*/
//go:nosplit
func (self class) FlipY() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_flip_y, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Generates mipmaps for the image. Mipmaps are precalculated lower-resolution copies of the image that are automatically used if the image needs to be scaled down when rendered. They help improve image quality and performance when rendering. This method returns an error if the image is compressed, in a custom format, or if the image's width/height is [code]0[/code]. Enabling [param renormalize] when generating mipmaps for normal map textures will make sure all resulting vector values are normalized.
It is possible to check if the image has mipmaps by calling [method has_mipmaps] or [method get_mipmap_count]. Calling [method generate_mipmaps] on an image that already has mipmaps will replace existing mipmaps in the image.
*/
//go:nosplit
func (self class) GenerateMipmaps(renormalize bool) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, renormalize)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the image's mipmaps.
*/
//go:nosplit
func (self class) ClearMipmaps() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_clear_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates an empty image of given size and format. See [enum Format] constants. If [param use_mipmaps] is [code]true[/code], then generate mipmaps for this image. See the [method generate_mipmaps].
*/
//go:nosplit
func (self class) Create(width gd.Int, height gd.Int, use_mipmaps bool, format gdclass.ImageFormat) [1]gdclass.Image {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, use_mipmaps)
	callframe.Arg(frame, format)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates an empty image of given size and format. See [enum Format] constants. If [param use_mipmaps] is [code]true[/code], then generate mipmaps for this image. See the [method generate_mipmaps].
*/
//go:nosplit
func (self class) CreateEmpty(width gd.Int, height gd.Int, use_mipmaps bool, format gdclass.ImageFormat) [1]gdclass.Image {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, use_mipmaps)
	callframe.Arg(frame, format)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_create_empty, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new image of given size and format. See [enum Format] constants. Fills the image with the given raw data. If [param use_mipmaps] is [code]true[/code] then loads mipmaps for this image from [param data]. See [method generate_mipmaps].
*/
//go:nosplit
func (self class) CreateFromData(width gd.Int, height gd.Int, use_mipmaps bool, format gdclass.ImageFormat, data gd.PackedByteArray) [1]gdclass.Image {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, use_mipmaps)
	callframe.Arg(frame, format)
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_create_from_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Overwrites data of an existing [Image]. Non-static equivalent of [method create_from_data].
*/
//go:nosplit
func (self class) SetData(width gd.Int, height gd.Int, use_mipmaps bool, format gdclass.ImageFormat, data gd.PackedByteArray) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, use_mipmaps)
	callframe.Arg(frame, format)
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the image has no data.
*/
//go:nosplit
func (self class) IsEmpty() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_is_empty, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads an image from file [param path]. See [url=$DOCS_URL/tutorials/assets_pipeline/importing_images.html#supported-image-formats]Supported image formats[/url] for a list of supported image formats and limitations.
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external images at run-time, such as images located at the [code]user://[/code] directory, and may not work in exported projects.
See also [ImageTexture] description for usage examples.
*/
//go:nosplit
func (self class) Load(path gd.String) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_load, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new [Image] and loads data from the specified file.
*/
//go:nosplit
func (self class) LoadFromFile(path gd.String) [1]gdclass.Image {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_load_from_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Saves the image as a PNG file to the file at [param path].
*/
//go:nosplit
func (self class) SavePng(path gd.String) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_save_png, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the image as a PNG file to a byte array.
*/
//go:nosplit
func (self class) SavePngToBuffer() gd.PackedByteArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_save_png_to_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Saves the image as a JPEG file to [param path] with the specified [param quality] between [code]0.01[/code] and [code]1.0[/code] (inclusive). Higher [param quality] values result in better-looking output at the cost of larger file sizes. Recommended [param quality] values are between [code]0.75[/code] and [code]0.90[/code]. Even at quality [code]1.00[/code], JPEG compression remains lossy.
[b]Note:[/b] JPEG does not save an alpha channel. If the [Image] contains an alpha channel, the image will still be saved, but the resulting JPEG file won't contain the alpha channel.
*/
//go:nosplit
func (self class) SaveJpg(path gd.String, quality gd.Float) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, quality)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_save_jpg, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the image as a JPEG file to a byte array with the specified [param quality] between [code]0.01[/code] and [code]1.0[/code] (inclusive). Higher [param quality] values result in better-looking output at the cost of larger byte array sizes (and therefore memory usage). Recommended [param quality] values are between [code]0.75[/code] and [code]0.90[/code]. Even at quality [code]1.00[/code], JPEG compression remains lossy.
[b]Note:[/b] JPEG does not save an alpha channel. If the [Image] contains an alpha channel, the image will still be saved, but the resulting byte array won't contain the alpha channel.
*/
//go:nosplit
func (self class) SaveJpgToBuffer(quality gd.Float) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_save_jpg_to_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Saves the image as an EXR file to [param path]. If [param grayscale] is [code]true[/code] and the image has only one channel, it will be saved explicitly as monochrome rather than one red channel. This function will return [constant ERR_UNAVAILABLE] if Godot was compiled without the TinyEXR module.
[b]Note:[/b] The TinyEXR module is disabled in non-editor builds, which means [method save_exr] will return [constant ERR_UNAVAILABLE] when it is called from an exported project.
*/
//go:nosplit
func (self class) SaveExr(path gd.String, grayscale bool) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, grayscale)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_save_exr, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the image as an EXR file to a byte array. If [param grayscale] is [code]true[/code] and the image has only one channel, it will be saved explicitly as monochrome rather than one red channel. This function will return an empty byte array if Godot was compiled without the TinyEXR module.
[b]Note:[/b] The TinyEXR module is disabled in non-editor builds, which means [method save_exr] will return an empty byte array when it is called from an exported project.
*/
//go:nosplit
func (self class) SaveExrToBuffer(grayscale bool) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, grayscale)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_save_exr_to_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Saves the image as a WebP (Web Picture) file to the file at [param path]. By default it will save lossless. If [param lossy] is true, the image will be saved lossy, using the [param quality] setting between 0.0 and 1.0 (inclusive). Lossless WebP offers more efficient compression than PNG.
[b]Note:[/b] The WebP format is limited to a size of 16383×16383 pixels, while PNG can save larger images.
*/
//go:nosplit
func (self class) SaveWebp(path gd.String, lossy bool, quality gd.Float) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, lossy)
	callframe.Arg(frame, quality)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_save_webp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the image as a WebP (Web Picture) file to a byte array. By default it will save lossless. If [param lossy] is true, the image will be saved lossy, using the [param quality] setting between 0.0 and 1.0 (inclusive). Lossless WebP offers more efficient compression than PNG.
[b]Note:[/b] The WebP format is limited to a size of 16383×16383 pixels, while PNG can save larger images.
*/
//go:nosplit
func (self class) SaveWebpToBuffer(lossy bool, quality gd.Float) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, lossy)
	callframe.Arg(frame, quality)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_save_webp_to_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [constant ALPHA_BLEND] if the image has data for alpha values. Returns [constant ALPHA_BIT] if all the alpha values are stored in a single bit. Returns [constant ALPHA_NONE] if no data for alpha values is found.
*/
//go:nosplit
func (self class) DetectAlpha() gdclass.ImageAlphaMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ImageAlphaMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_detect_alpha, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if all the image's pixels have an alpha value of 0. Returns [code]false[/code] if any pixel has an alpha value higher than 0.
*/
//go:nosplit
func (self class) IsInvisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_is_invisible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the color channels used by this image, as one of the [enum UsedChannels] constants. If the image is compressed, the original [param source] must be specified.
*/
//go:nosplit
func (self class) DetectUsedChannels(source gdclass.ImageCompressSource) gdclass.ImageUsedChannels {
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret = callframe.Ret[gdclass.ImageUsedChannels](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_detect_used_channels, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Compresses the image to use less memory. Can not directly access pixel data while the image is compressed. Returns error if the chosen compression mode is not available.
The [param source] parameter helps to pick the best compression method for DXT and ETC2 formats. It is ignored for ASTC compression.
For ASTC compression, the [param astc_format] parameter must be supplied.
*/
//go:nosplit
func (self class) Compress(mode gdclass.ImageCompressMode, source gdclass.ImageCompressSource, astc_format gdclass.ImageASTCFormat) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	callframe.Arg(frame, source)
	callframe.Arg(frame, astc_format)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_compress, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Compresses the image to use less memory. Can not directly access pixel data while the image is compressed. Returns error if the chosen compression mode is not available.
This is an alternative to [method compress] that lets the user supply the channels used in order for the compressor to pick the best DXT and ETC2 formats. For other formats (non DXT or ETC2), this argument is ignored.
For ASTC compression, the [param astc_format] parameter must be supplied.
*/
//go:nosplit
func (self class) CompressFromChannels(mode gdclass.ImageCompressMode, channels gdclass.ImageUsedChannels, astc_format gdclass.ImageASTCFormat) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	callframe.Arg(frame, channels)
	callframe.Arg(frame, astc_format)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_compress_from_channels, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Decompresses the image if it is VRAM compressed in a supported format. Returns [constant OK] if the format is supported, otherwise [constant ERR_UNAVAILABLE].
[b]Note:[/b] The following formats can be decompressed: DXT, RGTC, BPTC. The formats ETC1 and ETC2 are not supported.
*/
//go:nosplit
func (self class) Decompress() gd.Error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_decompress, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the image is compressed.
*/
//go:nosplit
func (self class) IsCompressed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_is_compressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Rotates the image in the specified [param direction] by [code]90[/code] degrees. The width and height of the image must be greater than [code]1[/code]. If the width and height are not equal, the image will be resized.
*/
//go:nosplit
func (self class) Rotate90(direction ClockDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_rotate_90, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Rotates the image by [code]180[/code] degrees. The width and height of the image must be greater than [code]1[/code].
*/
//go:nosplit
func (self class) Rotate180() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_rotate_180, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Blends low-alpha pixels with nearby pixels.
*/
//go:nosplit
func (self class) FixAlphaEdges() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_fix_alpha_edges, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Multiplies color values with alpha values. Resulting color values for a pixel are [code](color * alpha)/256[/code]. See also [member CanvasItemMaterial.blend_mode].
*/
//go:nosplit
func (self class) PremultiplyAlpha() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_premultiply_alpha, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Converts the raw data from the sRGB colorspace to a linear scale.
*/
//go:nosplit
func (self class) SrgbToLinear() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_srgb_to_linear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Converts the image's data to represent coordinates on a 3D plane. This is used when the image represents a normal map. A normal map can add lots of detail to a 3D surface without increasing the polygon count.
*/
//go:nosplit
func (self class) NormalMapToXy() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_normal_map_to_xy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Converts a standard RGBE (Red Green Blue Exponent) image to an sRGB image.
*/
//go:nosplit
func (self class) RgbeToSrgb() [1]gdclass.Image {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_rgbe_to_srgb, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Converts a bump map to a normal map. A bump map provides a height offset per-pixel, while a normal map provides a normal direction per pixel.
*/
//go:nosplit
func (self class) BumpMapToNormalMap(bump_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bump_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_bump_map_to_normal_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Compute image metrics on the current image and the compared image.
The dictionary contains [code]max[/code], [code]mean[/code], [code]mean_squared[/code], [code]root_mean_squared[/code] and [code]peak_snr[/code].
*/
//go:nosplit
func (self class) ComputeImageMetrics(compared_image [1]gdclass.Image, use_luma bool) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(compared_image[0])[0])
	callframe.Arg(frame, use_luma)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_compute_image_metrics, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Copies [param src_rect] from [param src] image to this image at coordinates [param dst], clipped accordingly to both image bounds. This image and [param src] image [b]must[/b] have the same format. [param src_rect] with non-positive size is treated as empty.
*/
//go:nosplit
func (self class) BlitRect(src [1]gdclass.Image, src_rect gd.Rect2i, dst gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src[0])[0])
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, dst)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_blit_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Blits [param src_rect] area from [param src] image to this image at the coordinates given by [param dst], clipped accordingly to both image bounds. [param src] pixel is copied onto [param dst] if the corresponding [param mask] pixel's alpha value is not 0. This image and [param src] image [b]must[/b] have the same format. [param src] image and [param mask] image [b]must[/b] have the same size (width and height) but they can have different formats. [param src_rect] with non-positive size is treated as empty.
*/
//go:nosplit
func (self class) BlitRectMask(src [1]gdclass.Image, mask [1]gdclass.Image, src_rect gd.Rect2i, dst gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src[0])[0])
	callframe.Arg(frame, pointers.Get(mask[0])[0])
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, dst)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_blit_rect_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Alpha-blends [param src_rect] from [param src] image to this image at coordinates [param dst], clipped accordingly to both image bounds. This image and [param src] image [b]must[/b] have the same format. [param src_rect] with non-positive size is treated as empty.
*/
//go:nosplit
func (self class) BlendRect(src [1]gdclass.Image, src_rect gd.Rect2i, dst gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src[0])[0])
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, dst)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_blend_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Alpha-blends [param src_rect] from [param src] image to this image using [param mask] image at coordinates [param dst], clipped accordingly to both image bounds. Alpha channels are required for both [param src] and [param mask]. [param dst] pixels and [param src] pixels will blend if the corresponding mask pixel's alpha value is not 0. This image and [param src] image [b]must[/b] have the same format. [param src] image and [param mask] image [b]must[/b] have the same size (width and height) but they can have different formats. [param src_rect] with non-positive size is treated as empty.
*/
//go:nosplit
func (self class) BlendRectMask(src [1]gdclass.Image, mask [1]gdclass.Image, src_rect gd.Rect2i, dst gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src[0])[0])
	callframe.Arg(frame, pointers.Get(mask[0])[0])
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, dst)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_blend_rect_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Fills the image with [param color].
*/
//go:nosplit
func (self class) Fill(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_fill, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Fills [param rect] with [param color].
*/
//go:nosplit
func (self class) FillRect(rect gd.Rect2i, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_fill_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a [Rect2i] enclosing the visible portion of the image, considering each pixel with a non-zero alpha channel as visible.
*/
//go:nosplit
func (self class) GetUsedRect() gd.Rect2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_used_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a new [Image] that is a copy of this [Image]'s area specified with [param region].
*/
//go:nosplit
func (self class) GetRegion(region gd.Rect2i) [1]gdclass.Image {
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Copies [param src] image to this image.
*/
//go:nosplit
func (self class) CopyFrom(src [1]gdclass.Image) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_copy_from, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the color of the pixel at [param point].
This is the same as [method get_pixel], but with a [Vector2i] argument instead of two integer arguments.
*/
//go:nosplit
func (self class) GetPixelv(point gd.Vector2i) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_pixelv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the color of the pixel at [code](x, y)[/code].
This is the same as [method get_pixelv], but with two integer arguments instead of a [Vector2i] argument.
*/
//go:nosplit
func (self class) GetPixel(x gd.Int, y gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_get_pixel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Color] of the pixel at [param point] to [param color].
[b]Example:[/b]
[codeblocks]
[gdscript]
var img_width = 10
var img_height = 5
var img = Image.create(img_width, img_height, false, Image.FORMAT_RGBA8)

img.set_pixelv(Vector2i(1, 2), Color.RED) # Sets the color at (1, 2) to red.
[/gdscript]
[csharp]
int imgWidth = 10;
int imgHeight = 5;
var img = Image.Create(imgWidth, imgHeight, false, Image.Format.Rgba8);

img.SetPixelv(new Vector2I(1, 2), Colors.Red); // Sets the color at (1, 2) to red.
[/csharp]
[/codeblocks]
This is the same as [method set_pixel], but with a [Vector2i] argument instead of two integer arguments.
*/
//go:nosplit
func (self class) SetPixelv(point gd.Vector2i, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_set_pixelv, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [Color] of the pixel at [code](x, y)[/code] to [param color].
[b]Example:[/b]
[codeblocks]
[gdscript]
var img_width = 10
var img_height = 5
var img = Image.create(img_width, img_height, false, Image.FORMAT_RGBA8)

img.set_pixel(1, 2, Color.RED) # Sets the color at (1, 2) to red.
[/gdscript]
[csharp]
int imgWidth = 10;
int imgHeight = 5;
var img = Image.Create(imgWidth, imgHeight, false, Image.Format.Rgba8);

img.SetPixel(1, 2, Colors.Red); // Sets the color at (1, 2) to red.
[/csharp]
[/codeblocks]
This is the same as [method set_pixelv], but with a two integer arguments instead of a [Vector2i] argument.
*/
//go:nosplit
func (self class) SetPixel(x gd.Int, y gd.Int, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_set_pixel, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adjusts this image's [param brightness], [param contrast], and [param saturation] by the given values. Does not work if the image is compressed (see [method is_compressed]).
*/
//go:nosplit
func (self class) AdjustBcs(brightness gd.Float, contrast gd.Float, saturation gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, brightness)
	callframe.Arg(frame, contrast)
	callframe.Arg(frame, saturation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_adjust_bcs, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Loads an image from the binary contents of a PNG file.
*/
//go:nosplit
func (self class) LoadPngFromBuffer(buffer gd.PackedByteArray) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_load_png_from_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads an image from the binary contents of a JPEG file.
*/
//go:nosplit
func (self class) LoadJpgFromBuffer(buffer gd.PackedByteArray) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_load_jpg_from_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads an image from the binary contents of a WebP file.
*/
//go:nosplit
func (self class) LoadWebpFromBuffer(buffer gd.PackedByteArray) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_load_webp_from_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads an image from the binary contents of a TGA file.
[b]Note:[/b] This method is only available in engine builds with the TGA module enabled. By default, the TGA module is enabled, but it can be disabled at build-time using the [code]module_tga_enabled=no[/code] SCons option.
*/
//go:nosplit
func (self class) LoadTgaFromBuffer(buffer gd.PackedByteArray) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_load_tga_from_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads an image from the binary contents of a BMP file.
[b]Note:[/b] Godot's BMP module doesn't support 16-bit per pixel images. Only 1-bit, 4-bit, 8-bit, 24-bit, and 32-bit per pixel images are supported.
[b]Note:[/b] This method is only available in engine builds with the BMP module enabled. By default, the BMP module is enabled, but it can be disabled at build-time using the [code]module_bmp_enabled=no[/code] SCons option.
*/
//go:nosplit
func (self class) LoadBmpFromBuffer(buffer gd.PackedByteArray) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_load_bmp_from_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads an image from the binary contents of a [url=https://github.com/KhronosGroup/KTX-Software]KTX[/url] file. Unlike most image formats, KTX can store VRAM-compressed data and embed mipmaps.
[b]Note:[/b] Godot's libktx implementation only supports 2D images. Cubemaps, texture arrays, and de-padding are not supported.
[b]Note:[/b] This method is only available in engine builds with the KTX module enabled. By default, the KTX module is enabled, but it can be disabled at build-time using the [code]module_ktx_enabled=no[/code] SCons option.
*/
//go:nosplit
func (self class) LoadKtxFromBuffer(buffer gd.PackedByteArray) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_load_ktx_from_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads an image from the UTF-8 binary contents of an [b]uncompressed[/b] SVG file ([b].svg[/b]).
[b]Note:[/b] Beware when using compressed SVG files (like [b].svgz[/b]), they need to be [code]decompressed[/code] before loading.
[b]Note:[/b] This method is only available in engine builds with the SVG module enabled. By default, the SVG module is enabled, but it can be disabled at build-time using the [code]module_svg_enabled=no[/code] SCons option.
*/
//go:nosplit
func (self class) LoadSvgFromBuffer(buffer gd.PackedByteArray, scale gd.Float) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	callframe.Arg(frame, scale)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_load_svg_from_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads an image from the string contents of an SVG file ([b].svg[/b]).
[b]Note:[/b] This method is only available in engine builds with the SVG module enabled. By default, the SVG module is enabled, but it can be disabled at build-time using the [code]module_svg_enabled=no[/code] SCons option.
*/
//go:nosplit
func (self class) LoadSvgFromString(svg_str gd.String, scale gd.Float) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(svg_str))
	callframe.Arg(frame, scale)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Image.Bind_load_svg_from_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsImage() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsImage() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Image", func(ptr gd.Object) any { return [1]gdclass.Image{*(*gdclass.Image)(unsafe.Pointer(&ptr))} })
}

type Format = gdclass.ImageFormat

const (
	/*Texture format with a single 8-bit depth representing luminance.*/
	FormatL8 Format = 0
	/*OpenGL texture format with two values, luminance and alpha each stored with 8 bits.*/
	FormatLa8 Format = 1
	/*OpenGL texture format [code]RED[/code] with a single component and a bitdepth of 8.*/
	FormatR8 Format = 2
	/*OpenGL texture format [code]RG[/code] with two components and a bitdepth of 8 for each.*/
	FormatRg8 Format = 3
	/*OpenGL texture format [code]RGB[/code] with three components, each with a bitdepth of 8.
	  [b]Note:[/b] When creating an [ImageTexture], an sRGB to linear color space conversion is performed.*/
	FormatRgb8 Format = 4
	/*OpenGL texture format [code]RGBA[/code] with four components, each with a bitdepth of 8.
	  [b]Note:[/b] When creating an [ImageTexture], an sRGB to linear color space conversion is performed.*/
	FormatRgba8 Format = 5
	/*OpenGL texture format [code]RGBA[/code] with four components, each with a bitdepth of 4.*/
	FormatRgba4444 Format = 6
	/*OpenGL texture format [code]RGB[/code] with three components. Red and blue have a bitdepth of 5, and green has a bitdepth of 6.*/
	FormatRgb565 Format = 7
	/*OpenGL texture format [code]GL_R32F[/code] where there's one component, a 32-bit floating-point value.*/
	FormatRf Format = 8
	/*OpenGL texture format [code]GL_RG32F[/code] where there are two components, each a 32-bit floating-point values.*/
	FormatRgf Format = 9
	/*OpenGL texture format [code]GL_RGB32F[/code] where there are three components, each a 32-bit floating-point values.*/
	FormatRgbf Format = 10
	/*OpenGL texture format [code]GL_RGBA32F[/code] where there are four components, each a 32-bit floating-point values.*/
	FormatRgbaf Format = 11
	/*OpenGL texture format [code]GL_R16F[/code] where there's one component, a 16-bit "half-precision" floating-point value.*/
	FormatRh Format = 12
	/*OpenGL texture format [code]GL_RG16F[/code] where there are two components, each a 16-bit "half-precision" floating-point value.*/
	FormatRgh Format = 13
	/*OpenGL texture format [code]GL_RGB16F[/code] where there are three components, each a 16-bit "half-precision" floating-point value.*/
	FormatRgbh Format = 14
	/*OpenGL texture format [code]GL_RGBA16F[/code] where there are four components, each a 16-bit "half-precision" floating-point value.*/
	FormatRgbah Format = 15
	/*A special OpenGL texture format where the three color components have 9 bits of precision and all three share a single 5-bit exponent.*/
	FormatRgbe9995 Format = 16
	/*The [url=https://en.wikipedia.org/wiki/S3_Texture_Compression]S3TC[/url] texture format that uses Block Compression 1, and is the smallest variation of S3TC, only providing 1 bit of alpha and color data being premultiplied with alpha.
	  [b]Note:[/b] When creating an [ImageTexture], an sRGB to linear color space conversion is performed.*/
	FormatDxt1 Format = 17
	/*The [url=https://en.wikipedia.org/wiki/S3_Texture_Compression]S3TC[/url] texture format that uses Block Compression 2, and color data is interpreted as not having been premultiplied by alpha. Well suited for images with sharp alpha transitions between translucent and opaque areas.
	  [b]Note:[/b] When creating an [ImageTexture], an sRGB to linear color space conversion is performed.*/
	FormatDxt3 Format = 18
	/*The [url=https://en.wikipedia.org/wiki/S3_Texture_Compression]S3TC[/url] texture format also known as Block Compression 3 or BC3 that contains 64 bits of alpha channel data followed by 64 bits of DXT1-encoded color data. Color data is not premultiplied by alpha, same as DXT3. DXT5 generally produces superior results for transparent gradients compared to DXT3.
	  [b]Note:[/b] When creating an [ImageTexture], an sRGB to linear color space conversion is performed.*/
	FormatDxt5 Format = 19
	/*Texture format that uses [url=https://www.khronos.org/opengl/wiki/Red_Green_Texture_Compression]Red Green Texture Compression[/url], normalizing the red channel data using the same compression algorithm that DXT5 uses for the alpha channel.*/
	FormatRgtcR Format = 20
	/*Texture format that uses [url=https://www.khronos.org/opengl/wiki/Red_Green_Texture_Compression]Red Green Texture Compression[/url], normalizing the red and green channel data using the same compression algorithm that DXT5 uses for the alpha channel.*/
	FormatRgtcRg Format = 21
	/*Texture format that uses [url=https://www.khronos.org/opengl/wiki/BPTC_Texture_Compression]BPTC[/url] compression with unsigned normalized RGBA components.
	  [b]Note:[/b] When creating an [ImageTexture], an sRGB to linear color space conversion is performed.*/
	FormatBptcRgba Format = 22
	/*Texture format that uses [url=https://www.khronos.org/opengl/wiki/BPTC_Texture_Compression]BPTC[/url] compression with signed floating-point RGB components.*/
	FormatBptcRgbf Format = 23
	/*Texture format that uses [url=https://www.khronos.org/opengl/wiki/BPTC_Texture_Compression]BPTC[/url] compression with unsigned floating-point RGB components.*/
	FormatBptcRgbfu Format = 24
	/*[url=https://en.wikipedia.org/wiki/Ericsson_Texture_Compression#ETC1]Ericsson Texture Compression format 1[/url], also referred to as "ETC1", and is part of the OpenGL ES graphics standard. This format cannot store an alpha channel.*/
	FormatEtc Format = 25
	/*[url=https://en.wikipedia.org/wiki/Ericsson_Texture_Compression#ETC2_and_EAC]Ericsson Texture Compression format 2[/url] ([code]R11_EAC[/code] variant), which provides one channel of unsigned data.*/
	FormatEtc2R11 Format = 26
	/*[url=https://en.wikipedia.org/wiki/Ericsson_Texture_Compression#ETC2_and_EAC]Ericsson Texture Compression format 2[/url] ([code]SIGNED_R11_EAC[/code] variant), which provides one channel of signed data.*/
	FormatEtc2R11s Format = 27
	/*[url=https://en.wikipedia.org/wiki/Ericsson_Texture_Compression#ETC2_and_EAC]Ericsson Texture Compression format 2[/url] ([code]RG11_EAC[/code] variant), which provides two channels of unsigned data.*/
	FormatEtc2Rg11 Format = 28
	/*[url=https://en.wikipedia.org/wiki/Ericsson_Texture_Compression#ETC2_and_EAC]Ericsson Texture Compression format 2[/url] ([code]SIGNED_RG11_EAC[/code] variant), which provides two channels of signed data.*/
	FormatEtc2Rg11s Format = 29
	/*[url=https://en.wikipedia.org/wiki/Ericsson_Texture_Compression#ETC2_and_EAC]Ericsson Texture Compression format 2[/url] ([code]RGB8[/code] variant), which is a follow-up of ETC1 and compresses RGB888 data.
	  [b]Note:[/b] When creating an [ImageTexture], an sRGB to linear color space conversion is performed.*/
	FormatEtc2Rgb8 Format = 30
	/*[url=https://en.wikipedia.org/wiki/Ericsson_Texture_Compression#ETC2_and_EAC]Ericsson Texture Compression format 2[/url] ([code]RGBA8[/code]variant), which compresses RGBA8888 data with full alpha support.
	  [b]Note:[/b] When creating an [ImageTexture], an sRGB to linear color space conversion is performed.*/
	FormatEtc2Rgba8 Format = 31
	/*[url=https://en.wikipedia.org/wiki/Ericsson_Texture_Compression#ETC2_and_EAC]Ericsson Texture Compression format 2[/url] ([code]RGB8_PUNCHTHROUGH_ALPHA1[/code] variant), which compresses RGBA data to make alpha either fully transparent or fully opaque.
	  [b]Note:[/b] When creating an [ImageTexture], an sRGB to linear color space conversion is performed.*/
	FormatEtc2Rgb8a1 Format = 32
	/*[url=https://en.wikipedia.org/wiki/Ericsson_Texture_Compression#ETC2_and_EAC]Ericsson Texture Compression format 2[/url] ([code]RGBA8[/code] variant), which compresses RA data and interprets it as two channels (red and green). See also [constant FORMAT_ETC2_RGBA8].*/
	FormatEtc2RaAsRg Format = 33
	/*The [url=https://en.wikipedia.org/wiki/S3_Texture_Compression]S3TC[/url] texture format also known as Block Compression 3 or BC3, which compresses RA data and interprets it as two channels (red and green). See also [constant FORMAT_DXT5].*/
	FormatDxt5RaAsRg Format = 34
	/*[url=https://en.wikipedia.org/wiki/Adaptive_scalable_texture_compression]Adaptive Scalable Texture Compression[/url]. This implements the 4×4 (high quality) mode.*/
	FormatAstc4x4 Format = 35
	/*Same format as [constant FORMAT_ASTC_4x4], but with the hint to let the GPU know it is used for HDR.*/
	FormatAstc4x4Hdr Format = 36
	/*[url=https://en.wikipedia.org/wiki/Adaptive_scalable_texture_compression]Adaptive Scalable Texture Compression[/url]. This implements the 8×8 (low quality) mode.*/
	FormatAstc8x8 Format = 37
	/*Same format as [constant FORMAT_ASTC_8x8], but with the hint to let the GPU know it is used for HDR.*/
	FormatAstc8x8Hdr Format = 38
	/*Represents the size of the [enum Format] enum.*/
	FormatMax Format = 39
)

type Interpolation = gdclass.ImageInterpolation

const (
	/*Performs nearest-neighbor interpolation. If the image is resized, it will be pixelated.*/
	InterpolateNearest Interpolation = 0
	/*Performs bilinear interpolation. If the image is resized, it will be blurry. This mode is faster than [constant INTERPOLATE_CUBIC], but it results in lower quality.*/
	InterpolateBilinear Interpolation = 1
	/*Performs cubic interpolation. If the image is resized, it will be blurry. This mode often gives better results compared to [constant INTERPOLATE_BILINEAR], at the cost of being slower.*/
	InterpolateCubic Interpolation = 2
	/*Performs bilinear separately on the two most-suited mipmap levels, then linearly interpolates between them.
	  It's slower than [constant INTERPOLATE_BILINEAR], but produces higher-quality results with far fewer aliasing artifacts.
	  If the image does not have mipmaps, they will be generated and used internally, but no mipmaps will be generated on the resulting image.
	  [b]Note:[/b] If you intend to scale multiple copies of the original image, it's better to call [method generate_mipmaps]] on it in advance, to avoid wasting processing power in generating them again and again.
	  On the other hand, if the image already has mipmaps, they will be used, and a new set will be generated for the resulting image.*/
	InterpolateTrilinear Interpolation = 3
	/*Performs Lanczos interpolation. This is the slowest image resizing mode, but it typically gives the best results, especially when downscaling images.*/
	InterpolateLanczos Interpolation = 4
)

type AlphaMode = gdclass.ImageAlphaMode

const (
	/*Image does not have alpha.*/
	AlphaNone AlphaMode = 0
	/*Image stores alpha in a single bit.*/
	AlphaBit AlphaMode = 1
	/*Image uses alpha.*/
	AlphaBlend AlphaMode = 2
)

type CompressMode = gdclass.ImageCompressMode

const (
	/*Use S3TC compression.*/
	CompressS3tc CompressMode = 0
	/*Use ETC compression.*/
	CompressEtc CompressMode = 1
	/*Use ETC2 compression.*/
	CompressEtc2 CompressMode = 2
	/*Use BPTC compression.*/
	CompressBptc CompressMode = 3
	/*Use ASTC compression.*/
	CompressAstc CompressMode = 4
	/*Represents the size of the [enum CompressMode] enum.*/
	CompressMax CompressMode = 5
)

type UsedChannels = gdclass.ImageUsedChannels

const (
	/*The image only uses one channel for luminance (grayscale).*/
	UsedChannelsL UsedChannels = 0
	/*The image uses two channels for luminance and alpha, respectively.*/
	UsedChannelsLa UsedChannels = 1
	/*The image only uses the red channel.*/
	UsedChannelsR UsedChannels = 2
	/*The image uses two channels for red and green.*/
	UsedChannelsRg UsedChannels = 3
	/*The image uses three channels for red, green, and blue.*/
	UsedChannelsRgb UsedChannels = 4
	/*The image uses four channels for red, green, blue, and alpha.*/
	UsedChannelsRgba UsedChannels = 5
)

type CompressSource = gdclass.ImageCompressSource

const (
	/*Source texture (before compression) is a regular texture. Default for all textures.*/
	CompressSourceGeneric CompressSource = 0
	/*Source texture (before compression) is in sRGB space.*/
	CompressSourceSrgb CompressSource = 1
	/*Source texture (before compression) is a normal texture (e.g. it can be compressed into two channels).*/
	CompressSourceNormal CompressSource = 2
)

type ASTCFormat = gdclass.ImageASTCFormat

const (
	/*Hint to indicate that the high quality 4×4 ASTC compression format should be used.*/
	AstcFormat4x4 ASTCFormat = 0
	/*Hint to indicate that the low quality 8×8 ASTC compression format should be used.*/
	AstcFormat8x8 ASTCFormat = 1
)

type ClockDirection int

const (
	/*Clockwise rotation. Used by some methods (e.g. [method Image.rotate_90]).*/
	Clockwise ClockDirection = 0
	/*Counter-clockwise rotation. Used by some methods (e.g. [method Image.rotate_90]).*/
	Counterclockwise ClockDirection = 1
)

type Error = gd.Error

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
