package Image

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Native image datatype. Contains image data which can be converted to an [ImageTexture] and provides commonly used [i]image processing[/i] methods. The maximum width and height for an [Image] are [constant MAX_WIDTH] and [constant MAX_HEIGHT].
An [Image] cannot be assigned to a texture property of an object directly (such as [member Sprite2D.texture]), and has to be converted manually to an [ImageTexture] first.
[b]Note:[/b] The maximum image size is 16384×16384 pixels due to graphics hardware limitations. Larger images may fail to import.

*/
type Simple [1]classdb.Image
func (self Simple) GetWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetWidth()))
}
func (self Simple) GetHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetHeight()))
}
func (self Simple) GetSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetSize())
}
func (self Simple) HasMipmaps() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasMipmaps())
}
func (self Simple) GetFormat() classdb.ImageFormat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ImageFormat(Expert(self).GetFormat())
}
func (self Simple) GetData() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).GetData(gc).Bytes())
}
func (self Simple) GetDataSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDataSize()))
}
func (self Simple) Convert(format classdb.ImageFormat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Convert(format)
}
func (self Simple) GetMipmapCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMipmapCount()))
}
func (self Simple) GetMipmapOffset(mipmap int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMipmapOffset(gd.Int(mipmap))))
}
func (self Simple) ResizeToPo2(square bool, interpolation classdb.ImageInterpolation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ResizeToPo2(square, interpolation)
}
func (self Simple) Resize(width int, height int, interpolation classdb.ImageInterpolation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Resize(gd.Int(width), gd.Int(height), interpolation)
}
func (self Simple) ShrinkX2() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ShrinkX2()
}
func (self Simple) Crop(width int, height int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Crop(gd.Int(width), gd.Int(height))
}
func (self Simple) FlipX() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FlipX()
}
func (self Simple) FlipY() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FlipY()
}
func (self Simple) GenerateMipmaps(renormalize bool) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).GenerateMipmaps(renormalize))
}
func (self Simple) ClearMipmaps() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearMipmaps()
}
func (self Simple) Create(width int, height int, use_mipmaps bool, format classdb.ImageFormat) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).Create(gc, gd.Int(width), gd.Int(height), use_mipmaps, format))
}
func (self Simple) CreateEmpty(width int, height int, use_mipmaps bool, format classdb.ImageFormat) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).CreateEmpty(gc, gd.Int(width), gd.Int(height), use_mipmaps, format))
}
func (self Simple) CreateFromData(width int, height int, use_mipmaps bool, format classdb.ImageFormat, data []byte) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).CreateFromData(gc, gd.Int(width), gd.Int(height), use_mipmaps, format, gc.PackedByteSlice(data)))
}
func (self Simple) SetData(width int, height int, use_mipmaps bool, format classdb.ImageFormat, data []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetData(gd.Int(width), gd.Int(height), use_mipmaps, format, gc.PackedByteSlice(data))
}
func (self Simple) IsEmpty() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEmpty())
}
func (self Simple) Load(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Load(gc.String(path)))
}
func (self Simple) LoadFromFile(path string) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).LoadFromFile(gc, gc.String(path)))
}
func (self Simple) SavePng(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SavePng(gc.String(path)))
}
func (self Simple) SavePngToBuffer() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).SavePngToBuffer(gc).Bytes())
}
func (self Simple) SaveJpg(path string, quality float64) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SaveJpg(gc.String(path), gd.Float(quality)))
}
func (self Simple) SaveJpgToBuffer(quality float64) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).SaveJpgToBuffer(gc, gd.Float(quality)).Bytes())
}
func (self Simple) SaveExr(path string, grayscale bool) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SaveExr(gc.String(path), grayscale))
}
func (self Simple) SaveExrToBuffer(grayscale bool) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).SaveExrToBuffer(gc, grayscale).Bytes())
}
func (self Simple) SaveWebp(path string, lossy bool, quality float64) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SaveWebp(gc.String(path), lossy, gd.Float(quality)))
}
func (self Simple) SaveWebpToBuffer(lossy bool, quality float64) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).SaveWebpToBuffer(gc, lossy, gd.Float(quality)).Bytes())
}
func (self Simple) DetectAlpha() classdb.ImageAlphaMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ImageAlphaMode(Expert(self).DetectAlpha())
}
func (self Simple) IsInvisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsInvisible())
}
func (self Simple) DetectUsedChannels(source classdb.ImageCompressSource) classdb.ImageUsedChannels {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ImageUsedChannels(Expert(self).DetectUsedChannels(source))
}
func (self Simple) Compress(mode classdb.ImageCompressMode, source classdb.ImageCompressSource, astc_format classdb.ImageASTCFormat) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Compress(mode, source, astc_format))
}
func (self Simple) CompressFromChannels(mode classdb.ImageCompressMode, channels classdb.ImageUsedChannels, astc_format classdb.ImageASTCFormat) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CompressFromChannels(mode, channels, astc_format))
}
func (self Simple) Decompress() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Decompress())
}
func (self Simple) IsCompressed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCompressed())
}
func (self Simple) Rotate90(direction gd.ClockDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Rotate90(direction)
}
func (self Simple) Rotate180() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Rotate180()
}
func (self Simple) FixAlphaEdges() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FixAlphaEdges()
}
func (self Simple) PremultiplyAlpha() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PremultiplyAlpha()
}
func (self Simple) SrgbToLinear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SrgbToLinear()
}
func (self Simple) NormalMapToXy() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).NormalMapToXy()
}
func (self Simple) RgbeToSrgb() [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).RgbeToSrgb(gc))
}
func (self Simple) BumpMapToNormalMap(bump_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BumpMapToNormalMap(gd.Float(bump_scale))
}
func (self Simple) ComputeImageMetrics(compared_image [1]classdb.Image, use_luma bool) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).ComputeImageMetrics(gc, compared_image, use_luma))
}
func (self Simple) BlitRect(src [1]classdb.Image, src_rect gd.Rect2i, dst gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BlitRect(src, src_rect, dst)
}
func (self Simple) BlitRectMask(src [1]classdb.Image, mask [1]classdb.Image, src_rect gd.Rect2i, dst gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BlitRectMask(src, mask, src_rect, dst)
}
func (self Simple) BlendRect(src [1]classdb.Image, src_rect gd.Rect2i, dst gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BlendRect(src, src_rect, dst)
}
func (self Simple) BlendRectMask(src [1]classdb.Image, mask [1]classdb.Image, src_rect gd.Rect2i, dst gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BlendRectMask(src, mask, src_rect, dst)
}
func (self Simple) Fill(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Fill(color)
}
func (self Simple) FillRect(rect gd.Rect2i, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FillRect(rect, color)
}
func (self Simple) GetUsedRect() gd.Rect2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2i(Expert(self).GetUsedRect())
}
func (self Simple) GetRegion(region gd.Rect2i) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).GetRegion(gc, region))
}
func (self Simple) CopyFrom(src [1]classdb.Image) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CopyFrom(src)
}
func (self Simple) GetPixelv(point gd.Vector2i) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetPixelv(point))
}
func (self Simple) GetPixel(x int, y int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetPixel(gd.Int(x), gd.Int(y)))
}
func (self Simple) SetPixelv(point gd.Vector2i, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPixelv(point, color)
}
func (self Simple) SetPixel(x int, y int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPixel(gd.Int(x), gd.Int(y), color)
}
func (self Simple) AdjustBcs(brightness float64, contrast float64, saturation float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AdjustBcs(gd.Float(brightness), gd.Float(contrast), gd.Float(saturation))
}
func (self Simple) LoadPngFromBuffer(buffer []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadPngFromBuffer(gc.PackedByteSlice(buffer)))
}
func (self Simple) LoadJpgFromBuffer(buffer []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadJpgFromBuffer(gc.PackedByteSlice(buffer)))
}
func (self Simple) LoadWebpFromBuffer(buffer []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadWebpFromBuffer(gc.PackedByteSlice(buffer)))
}
func (self Simple) LoadTgaFromBuffer(buffer []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadTgaFromBuffer(gc.PackedByteSlice(buffer)))
}
func (self Simple) LoadBmpFromBuffer(buffer []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadBmpFromBuffer(gc.PackedByteSlice(buffer)))
}
func (self Simple) LoadKtxFromBuffer(buffer []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadKtxFromBuffer(gc.PackedByteSlice(buffer)))
}
func (self Simple) LoadSvgFromBuffer(buffer []byte, scale float64) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadSvgFromBuffer(gc.PackedByteSlice(buffer), gd.Float(scale)))
}
func (self Simple) LoadSvgFromString(svg_str string, scale float64) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadSvgFromString(gc.String(svg_str), gd.Float(scale)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Image
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the image's width.
*/
//go:nosplit
func (self class) GetWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the image's height.
*/
//go:nosplit
func (self class) GetHeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the image's size (width and height).
*/
//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the image has generated mipmaps.
*/
//go:nosplit
func (self class) HasMipmaps() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_has_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the image's format. See [enum Format] constants.
*/
//go:nosplit
func (self class) GetFormat() classdb.ImageFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ImageFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a copy of the image's raw data.
*/
//go:nosplit
func (self class) GetData(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns size (in bytes) of the image's raw data.
*/
//go:nosplit
func (self class) GetDataSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_data_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Converts the image's format. See [enum Format] constants.
*/
//go:nosplit
func (self class) Convert(format classdb.ImageFormat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_convert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of mipmap levels or 0 if the image has no mipmaps. The largest main level image is not counted as a mipmap level by this method, so if you want to include it you can add 1 to this count.
*/
//go:nosplit
func (self class) GetMipmapCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_mipmap_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the offset where the image's mipmap with index [param mipmap] is stored in the [member data] dictionary.
*/
//go:nosplit
func (self class) GetMipmapOffset(mipmap gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mipmap)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_mipmap_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Resizes the image to the nearest power of 2 for the width and height. If [param square] is [code]true[/code] then set width and height to be the same. New pixels are calculated using the [param interpolation] mode defined via [enum Interpolation] constants.
*/
//go:nosplit
func (self class) ResizeToPo2(square bool, interpolation classdb.ImageInterpolation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, square)
	callframe.Arg(frame, interpolation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_resize_to_po2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Resizes the image to the given [param width] and [param height]. New pixels are calculated using the [param interpolation] mode defined via [enum Interpolation] constants.
*/
//go:nosplit
func (self class) Resize(width gd.Int, height gd.Int, interpolation classdb.ImageInterpolation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, interpolation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_resize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Shrinks the image by a factor of 2 on each axis (this divides the pixel count by 4).
*/
//go:nosplit
func (self class) ShrinkX2()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_shrink_x2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Crops the image to the given [param width] and [param height]. If the specified size is larger than the current size, the extra area is filled with black pixels.
*/
//go:nosplit
func (self class) Crop(width gd.Int, height gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_crop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Flips the image horizontally.
*/
//go:nosplit
func (self class) FlipX()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_flip_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Flips the image vertically.
*/
//go:nosplit
func (self class) FlipY()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_flip_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Generates mipmaps for the image. Mipmaps are precalculated lower-resolution copies of the image that are automatically used if the image needs to be scaled down when rendered. They help improve image quality and performance when rendering. This method returns an error if the image is compressed, in a custom format, or if the image's width/height is [code]0[/code]. Enabling [param renormalize] when generating mipmaps for normal map textures will make sure all resulting vector values are normalized.
It is possible to check if the image has mipmaps by calling [method has_mipmaps] or [method get_mipmap_count]. Calling [method generate_mipmaps] on an image that already has mipmaps will replace existing mipmaps in the image.
*/
//go:nosplit
func (self class) GenerateMipmaps(renormalize bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, renormalize)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the image's mipmaps.
*/
//go:nosplit
func (self class) ClearMipmaps()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_clear_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates an empty image of given size and format. See [enum Format] constants. If [param use_mipmaps] is [code]true[/code], then generate mipmaps for this image. See the [method generate_mipmaps].
*/
//go:nosplit
func (self class) Create(ctx gd.Lifetime, width gd.Int, height gd.Int, use_mipmaps bool, format classdb.ImageFormat) object.Image {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, use_mipmaps)
	callframe.Arg(frame, format)
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.Image.Bind_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates an empty image of given size and format. See [enum Format] constants. If [param use_mipmaps] is [code]true[/code], then generate mipmaps for this image. See the [method generate_mipmaps].
*/
//go:nosplit
func (self class) CreateEmpty(ctx gd.Lifetime, width gd.Int, height gd.Int, use_mipmaps bool, format classdb.ImageFormat) object.Image {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, use_mipmaps)
	callframe.Arg(frame, format)
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.Image.Bind_create_empty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new image of given size and format. See [enum Format] constants. Fills the image with the given raw data. If [param use_mipmaps] is [code]true[/code] then loads mipmaps for this image from [param data]. See [method generate_mipmaps].
*/
//go:nosplit
func (self class) CreateFromData(ctx gd.Lifetime, width gd.Int, height gd.Int, use_mipmaps bool, format classdb.ImageFormat, data gd.PackedByteArray) object.Image {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, use_mipmaps)
	callframe.Arg(frame, format)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.Image.Bind_create_from_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Overwrites data of an existing [Image]. Non-static equivalent of [method create_from_data].
*/
//go:nosplit
func (self class) SetData(width gd.Int, height gd.Int, use_mipmaps bool, format classdb.ImageFormat, data gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, use_mipmaps)
	callframe.Arg(frame, format)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the image has no data.
*/
//go:nosplit
func (self class) IsEmpty() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_is_empty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) Load(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_load, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new [Image] and loads data from the specified file.
*/
//go:nosplit
func (self class) LoadFromFile(ctx gd.Lifetime, path gd.String) object.Image {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.Image.Bind_load_from_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Saves the image as a PNG file to the file at [param path].
*/
//go:nosplit
func (self class) SavePng(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_save_png, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Saves the image as a PNG file to a byte array.
*/
//go:nosplit
func (self class) SavePngToBuffer(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_save_png_to_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Saves the image as a JPEG file to [param path] with the specified [param quality] between [code]0.01[/code] and [code]1.0[/code] (inclusive). Higher [param quality] values result in better-looking output at the cost of larger file sizes. Recommended [param quality] values are between [code]0.75[/code] and [code]0.90[/code]. Even at quality [code]1.00[/code], JPEG compression remains lossy.
[b]Note:[/b] JPEG does not save an alpha channel. If the [Image] contains an alpha channel, the image will still be saved, but the resulting JPEG file won't contain the alpha channel.
*/
//go:nosplit
func (self class) SaveJpg(path gd.String, quality gd.Float) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, quality)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_save_jpg, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Saves the image as a JPEG file to a byte array with the specified [param quality] between [code]0.01[/code] and [code]1.0[/code] (inclusive). Higher [param quality] values result in better-looking output at the cost of larger byte array sizes (and therefore memory usage). Recommended [param quality] values are between [code]0.75[/code] and [code]0.90[/code]. Even at quality [code]1.00[/code], JPEG compression remains lossy.
[b]Note:[/b] JPEG does not save an alpha channel. If the [Image] contains an alpha channel, the image will still be saved, but the resulting byte array won't contain the alpha channel.
*/
//go:nosplit
func (self class) SaveJpgToBuffer(ctx gd.Lifetime, quality gd.Float) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_save_jpg_to_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Saves the image as an EXR file to [param path]. If [param grayscale] is [code]true[/code] and the image has only one channel, it will be saved explicitly as monochrome rather than one red channel. This function will return [constant ERR_UNAVAILABLE] if Godot was compiled without the TinyEXR module.
[b]Note:[/b] The TinyEXR module is disabled in non-editor builds, which means [method save_exr] will return [constant ERR_UNAVAILABLE] when it is called from an exported project.
*/
//go:nosplit
func (self class) SaveExr(path gd.String, grayscale bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, grayscale)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_save_exr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Saves the image as an EXR file to a byte array. If [param grayscale] is [code]true[/code] and the image has only one channel, it will be saved explicitly as monochrome rather than one red channel. This function will return an empty byte array if Godot was compiled without the TinyEXR module.
[b]Note:[/b] The TinyEXR module is disabled in non-editor builds, which means [method save_exr] will return an empty byte array when it is called from an exported project.
*/
//go:nosplit
func (self class) SaveExrToBuffer(ctx gd.Lifetime, grayscale bool) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, grayscale)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_save_exr_to_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Saves the image as a WebP (Web Picture) file to the file at [param path]. By default it will save lossless. If [param lossy] is true, the image will be saved lossy, using the [param quality] setting between 0.0 and 1.0 (inclusive). Lossless WebP offers more efficient compression than PNG.
[b]Note:[/b] The WebP format is limited to a size of 16383×16383 pixels, while PNG can save larger images.
*/
//go:nosplit
func (self class) SaveWebp(path gd.String, lossy bool, quality gd.Float) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, lossy)
	callframe.Arg(frame, quality)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_save_webp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Saves the image as a WebP (Web Picture) file to a byte array. By default it will save lossless. If [param lossy] is true, the image will be saved lossy, using the [param quality] setting between 0.0 and 1.0 (inclusive). Lossless WebP offers more efficient compression than PNG.
[b]Note:[/b] The WebP format is limited to a size of 16383×16383 pixels, while PNG can save larger images.
*/
//go:nosplit
func (self class) SaveWebpToBuffer(ctx gd.Lifetime, lossy bool, quality gd.Float) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lossy)
	callframe.Arg(frame, quality)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_save_webp_to_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [constant ALPHA_BLEND] if the image has data for alpha values. Returns [constant ALPHA_BIT] if all the alpha values are stored in a single bit. Returns [constant ALPHA_NONE] if no data for alpha values is found.
*/
//go:nosplit
func (self class) DetectAlpha() classdb.ImageAlphaMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ImageAlphaMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_detect_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if all the image's pixels have an alpha value of 0. Returns [code]false[/code] if any pixel has an alpha value higher than 0.
*/
//go:nosplit
func (self class) IsInvisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_is_invisible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the color channels used by this image, as one of the [enum UsedChannels] constants. If the image is compressed, the original [param source] must be specified.
*/
//go:nosplit
func (self class) DetectUsedChannels(source classdb.ImageCompressSource) classdb.ImageUsedChannels {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret = callframe.Ret[classdb.ImageUsedChannels](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_detect_used_channels, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) Compress(mode classdb.ImageCompressMode, source classdb.ImageCompressSource, astc_format classdb.ImageASTCFormat) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	callframe.Arg(frame, source)
	callframe.Arg(frame, astc_format)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_compress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) CompressFromChannels(mode classdb.ImageCompressMode, channels classdb.ImageUsedChannels, astc_format classdb.ImageASTCFormat) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	callframe.Arg(frame, channels)
	callframe.Arg(frame, astc_format)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_compress_from_channels, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Decompresses the image if it is VRAM compressed in a supported format. Returns [constant OK] if the format is supported, otherwise [constant ERR_UNAVAILABLE].
[b]Note:[/b] The following formats can be decompressed: DXT, RGTC, BPTC. The formats ETC1 and ETC2 are not supported.
*/
//go:nosplit
func (self class) Decompress() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_decompress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the image is compressed.
*/
//go:nosplit
func (self class) IsCompressed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_is_compressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Rotates the image in the specified [param direction] by [code]90[/code] degrees. The width and height of the image must be greater than [code]1[/code]. If the width and height are not equal, the image will be resized.
*/
//go:nosplit
func (self class) Rotate90(direction gd.ClockDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_rotate_90, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Rotates the image by [code]180[/code] degrees. The width and height of the image must be greater than [code]1[/code].
*/
//go:nosplit
func (self class) Rotate180()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_rotate_180, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Blends low-alpha pixels with nearby pixels.
*/
//go:nosplit
func (self class) FixAlphaEdges()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_fix_alpha_edges, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Multiplies color values with alpha values. Resulting color values for a pixel are [code](color * alpha)/256[/code]. See also [member CanvasItemMaterial.blend_mode].
*/
//go:nosplit
func (self class) PremultiplyAlpha()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_premultiply_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Converts the raw data from the sRGB colorspace to a linear scale.
*/
//go:nosplit
func (self class) SrgbToLinear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_srgb_to_linear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Converts the image's data to represent coordinates on a 3D plane. This is used when the image represents a normal map. A normal map can add lots of detail to a 3D surface without increasing the polygon count.
*/
//go:nosplit
func (self class) NormalMapToXy()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_normal_map_to_xy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Converts a standard RGBE (Red Green Blue Exponent) image to an sRGB image.
*/
//go:nosplit
func (self class) RgbeToSrgb(ctx gd.Lifetime) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_rgbe_to_srgb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Converts a bump map to a normal map. A bump map provides a height offset per-pixel, while a normal map provides a normal direction per pixel.
*/
//go:nosplit
func (self class) BumpMapToNormalMap(bump_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bump_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_bump_map_to_normal_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Compute image metrics on the current image and the compared image.
The dictionary contains [code]max[/code], [code]mean[/code], [code]mean_squared[/code], [code]root_mean_squared[/code] and [code]peak_snr[/code].
*/
//go:nosplit
func (self class) ComputeImageMetrics(ctx gd.Lifetime, compared_image object.Image, use_luma bool) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(compared_image[0].AsPointer())[0])
	callframe.Arg(frame, use_luma)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_compute_image_metrics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Copies [param src_rect] from [param src] image to this image at coordinates [param dst], clipped accordingly to both image bounds. This image and [param src] image [b]must[/b] have the same format. [param src_rect] with non-positive size is treated as empty.
*/
//go:nosplit
func (self class) BlitRect(src object.Image, src_rect gd.Rect2i, dst gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src[0].AsPointer())[0])
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, dst)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_blit_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Blits [param src_rect] area from [param src] image to this image at the coordinates given by [param dst], clipped accordingly to both image bounds. [param src] pixel is copied onto [param dst] if the corresponding [param mask] pixel's alpha value is not 0. This image and [param src] image [b]must[/b] have the same format. [param src] image and [param mask] image [b]must[/b] have the same size (width and height) but they can have different formats. [param src_rect] with non-positive size is treated as empty.
*/
//go:nosplit
func (self class) BlitRectMask(src object.Image, mask object.Image, src_rect gd.Rect2i, dst gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(mask[0].AsPointer())[0])
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, dst)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_blit_rect_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Alpha-blends [param src_rect] from [param src] image to this image at coordinates [param dst], clipped accordingly to both image bounds. This image and [param src] image [b]must[/b] have the same format. [param src_rect] with non-positive size is treated as empty.
*/
//go:nosplit
func (self class) BlendRect(src object.Image, src_rect gd.Rect2i, dst gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src[0].AsPointer())[0])
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, dst)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_blend_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Alpha-blends [param src_rect] from [param src] image to this image using [param mask] image at coordinates [param dst], clipped accordingly to both image bounds. Alpha channels are required for both [param src] and [param mask]. [param dst] pixels and [param src] pixels will blend if the corresponding mask pixel's alpha value is not 0. This image and [param src] image [b]must[/b] have the same format. [param src] image and [param mask] image [b]must[/b] have the same size (width and height) but they can have different formats. [param src_rect] with non-positive size is treated as empty.
*/
//go:nosplit
func (self class) BlendRectMask(src object.Image, mask object.Image, src_rect gd.Rect2i, dst gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(mask[0].AsPointer())[0])
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, dst)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_blend_rect_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Fills the image with [param color].
*/
//go:nosplit
func (self class) Fill(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Fills [param rect] with [param color].
*/
//go:nosplit
func (self class) FillRect(rect gd.Rect2i, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_fill_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Rect2i] enclosing the visible portion of the image, considering each pixel with a non-zero alpha channel as visible.
*/
//go:nosplit
func (self class) GetUsedRect() gd.Rect2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_used_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a new [Image] that is a copy of this [Image]'s area specified with [param region].
*/
//go:nosplit
func (self class) GetRegion(ctx gd.Lifetime, region gd.Rect2i) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Copies [param src] image to this image.
*/
//go:nosplit
func (self class) CopyFrom(src object.Image)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_copy_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the color of the pixel at [param point].
This is the same as [method get_pixel], but with a [Vector2i] argument instead of two integer arguments.
*/
//go:nosplit
func (self class) GetPixelv(point gd.Vector2i) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_pixelv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_get_pixel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SetPixelv(point gd.Vector2i, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_set_pixelv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SetPixel(x gd.Int, y gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_set_pixel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adjusts this image's [param brightness], [param contrast], and [param saturation] by the given values. Does not work if the image is compressed (see [method is_compressed]).
*/
//go:nosplit
func (self class) AdjustBcs(brightness gd.Float, contrast gd.Float, saturation gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, brightness)
	callframe.Arg(frame, contrast)
	callframe.Arg(frame, saturation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_adjust_bcs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Loads an image from the binary contents of a PNG file.
*/
//go:nosplit
func (self class) LoadPngFromBuffer(buffer gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_load_png_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Loads an image from the binary contents of a JPEG file.
*/
//go:nosplit
func (self class) LoadJpgFromBuffer(buffer gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_load_jpg_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Loads an image from the binary contents of a WebP file.
*/
//go:nosplit
func (self class) LoadWebpFromBuffer(buffer gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_load_webp_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Loads an image from the binary contents of a TGA file.
[b]Note:[/b] This method is only available in engine builds with the TGA module enabled. By default, the TGA module is enabled, but it can be disabled at build-time using the [code]module_tga_enabled=no[/code] SCons option.
*/
//go:nosplit
func (self class) LoadTgaFromBuffer(buffer gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_load_tga_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) LoadBmpFromBuffer(buffer gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_load_bmp_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) LoadKtxFromBuffer(buffer gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_load_ktx_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) LoadSvgFromBuffer(buffer gd.PackedByteArray, scale gd.Float) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	callframe.Arg(frame, scale)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_load_svg_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Loads an image from the string contents of an SVG file ([b].svg[/b]).
[b]Note:[/b] This method is only available in engine builds with the SVG module enabled. By default, the SVG module is enabled, but it can be disabled at build-time using the [code]module_svg_enabled=no[/code] SCons option.
*/
//go:nosplit
func (self class) LoadSvgFromString(svg_str gd.String, scale gd.Float) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(svg_str))
	callframe.Arg(frame, scale)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Image.Bind_load_svg_from_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsImage() Expert { return self[0].AsImage() }


//go:nosplit
func (self Simple) AsImage() Simple { return self[0].AsImage() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Image", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Format = classdb.ImageFormat

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
type Interpolation = classdb.ImageInterpolation

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
type AlphaMode = classdb.ImageAlphaMode

const (
/*Image does not have alpha.*/
	AlphaNone AlphaMode = 0
/*Image stores alpha in a single bit.*/
	AlphaBit AlphaMode = 1
/*Image uses alpha.*/
	AlphaBlend AlphaMode = 2
)
type CompressMode = classdb.ImageCompressMode

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
type UsedChannels = classdb.ImageUsedChannels

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
type CompressSource = classdb.ImageCompressSource

const (
/*Source texture (before compression) is a regular texture. Default for all textures.*/
	CompressSourceGeneric CompressSource = 0
/*Source texture (before compression) is in sRGB space.*/
	CompressSourceSrgb CompressSource = 1
/*Source texture (before compression) is a normal texture (e.g. it can be compressed into two channels).*/
	CompressSourceNormal CompressSource = 2
)
type ASTCFormat = classdb.ImageASTCFormat

const (
/*Hint to indicate that the high quality 4×4 ASTC compression format should be used.*/
	AstcFormat4x4 ASTCFormat = 0
/*Hint to indicate that the low quality 8×8 ASTC compression format should be used.*/
	AstcFormat8x8 ASTCFormat = 1
)
