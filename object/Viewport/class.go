package Viewport

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A [Viewport] creates a different view into the screen, or a sub-view inside another viewport. Child 2D nodes will display on it, and child Camera3D 3D nodes will render on it too.
Optionally, a viewport can have its own 2D or 3D world, so it doesn't share what it draws with other viewports.
Viewports can also choose to be audio listeners, so they generate positional audio depending on a 2D or 3D camera child of it.
Also, viewports can be assigned to different screens in case the devices have multiple screens.
Finally, viewports can also behave as render targets, in which case they will not be visible unless the associated texture is used to draw.

*/
type Simple [1]classdb.Viewport
func (self Simple) SetWorld2d(world_2d [1]classdb.World2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWorld2d(world_2d)
}
func (self Simple) GetWorld2d() [1]classdb.World2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.World2D(Expert(self).GetWorld2d(gc))
}
func (self Simple) FindWorld2d() [1]classdb.World2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.World2D(Expert(self).FindWorld2d(gc))
}
func (self Simple) SetCanvasTransform(xform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCanvasTransform(xform)
}
func (self Simple) GetCanvasTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).GetCanvasTransform())
}
func (self Simple) SetGlobalCanvasTransform(xform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlobalCanvasTransform(xform)
}
func (self Simple) GetGlobalCanvasTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).GetGlobalCanvasTransform())
}
func (self Simple) GetFinalTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).GetFinalTransform())
}
func (self Simple) GetScreenTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).GetScreenTransform())
}
func (self Simple) GetVisibleRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetVisibleRect())
}
func (self Simple) SetTransparentBackground(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransparentBackground(enable)
}
func (self Simple) HasTransparentBackground() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasTransparentBackground())
}
func (self Simple) SetUseHdr2d(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseHdr2d(enable)
}
func (self Simple) IsUsingHdr2d() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingHdr2d())
}
func (self Simple) SetMsaa2d(msaa classdb.ViewportMSAA) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMsaa2d(msaa)
}
func (self Simple) GetMsaa2d() classdb.ViewportMSAA {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportMSAA(Expert(self).GetMsaa2d())
}
func (self Simple) SetMsaa3d(msaa classdb.ViewportMSAA) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMsaa3d(msaa)
}
func (self Simple) GetMsaa3d() classdb.ViewportMSAA {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportMSAA(Expert(self).GetMsaa3d())
}
func (self Simple) SetScreenSpaceAa(screen_space_aa classdb.ViewportScreenSpaceAA) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScreenSpaceAa(screen_space_aa)
}
func (self Simple) GetScreenSpaceAa() classdb.ViewportScreenSpaceAA {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportScreenSpaceAA(Expert(self).GetScreenSpaceAa())
}
func (self Simple) SetUseTaa(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseTaa(enable)
}
func (self Simple) IsUsingTaa() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingTaa())
}
func (self Simple) SetUseDebanding(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseDebanding(enable)
}
func (self Simple) IsUsingDebanding() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingDebanding())
}
func (self Simple) SetUseOcclusionCulling(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseOcclusionCulling(enable)
}
func (self Simple) IsUsingOcclusionCulling() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingOcclusionCulling())
}
func (self Simple) SetDebugDraw(debug_draw classdb.ViewportDebugDraw) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDebugDraw(debug_draw)
}
func (self Simple) GetDebugDraw() classdb.ViewportDebugDraw {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportDebugDraw(Expert(self).GetDebugDraw())
}
func (self Simple) GetRenderInfo(atype classdb.ViewportRenderInfoType, info classdb.ViewportRenderInfo) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetRenderInfo(atype, info)))
}
func (self Simple) GetTexture() [1]classdb.ViewportTexture {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.ViewportTexture(Expert(self).GetTexture(gc))
}
func (self Simple) SetPhysicsObjectPicking(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPhysicsObjectPicking(enable)
}
func (self Simple) GetPhysicsObjectPicking() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetPhysicsObjectPicking())
}
func (self Simple) SetPhysicsObjectPickingSort(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPhysicsObjectPickingSort(enable)
}
func (self Simple) GetPhysicsObjectPickingSort() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetPhysicsObjectPickingSort())
}
func (self Simple) SetPhysicsObjectPickingFirstOnly(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPhysicsObjectPickingFirstOnly(enable)
}
func (self Simple) GetPhysicsObjectPickingFirstOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetPhysicsObjectPickingFirstOnly())
}
func (self Simple) GetViewportRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetViewportRid())
}
func (self Simple) PushTextInput(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushTextInput(gc.String(text))
}
func (self Simple) PushInput(event [1]classdb.InputEvent, in_local_coords bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushInput(event, in_local_coords)
}
func (self Simple) PushUnhandledInput(event [1]classdb.InputEvent, in_local_coords bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushUnhandledInput(event, in_local_coords)
}
func (self Simple) GetMousePosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetMousePosition())
}
func (self Simple) WarpMouse(position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WarpMouse(position)
}
func (self Simple) UpdateMouseCursorState() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateMouseCursorState()
}
func (self Simple) GuiGetDragData() gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GuiGetDragData(gc))
}
func (self Simple) GuiIsDragging() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GuiIsDragging())
}
func (self Simple) GuiIsDragSuccessful() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GuiIsDragSuccessful())
}
func (self Simple) GuiReleaseFocus() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GuiReleaseFocus()
}
func (self Simple) GuiGetFocusOwner() [1]classdb.Control {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Control(Expert(self).GuiGetFocusOwner(gc))
}
func (self Simple) GuiGetHoveredControl() [1]classdb.Control {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Control(Expert(self).GuiGetHoveredControl(gc))
}
func (self Simple) SetDisableInput(disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisableInput(disable)
}
func (self Simple) IsInputDisabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsInputDisabled())
}
func (self Simple) SetPositionalShadowAtlasSize(size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPositionalShadowAtlasSize(gd.Int(size))
}
func (self Simple) GetPositionalShadowAtlasSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPositionalShadowAtlasSize()))
}
func (self Simple) SetPositionalShadowAtlas16Bits(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPositionalShadowAtlas16Bits(enable)
}
func (self Simple) GetPositionalShadowAtlas16Bits() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetPositionalShadowAtlas16Bits())
}
func (self Simple) SetSnapControlsToPixels(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSnapControlsToPixels(enabled)
}
func (self Simple) IsSnapControlsToPixelsEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSnapControlsToPixelsEnabled())
}
func (self Simple) SetSnap2dTransformsToPixel(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSnap2dTransformsToPixel(enabled)
}
func (self Simple) IsSnap2dTransformsToPixelEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSnap2dTransformsToPixelEnabled())
}
func (self Simple) SetSnap2dVerticesToPixel(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSnap2dVerticesToPixel(enabled)
}
func (self Simple) IsSnap2dVerticesToPixelEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSnap2dVerticesToPixelEnabled())
}
func (self Simple) SetPositionalShadowAtlasQuadrantSubdiv(quadrant int, subdiv classdb.ViewportPositionalShadowAtlasQuadrantSubdiv) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPositionalShadowAtlasQuadrantSubdiv(gd.Int(quadrant), subdiv)
}
func (self Simple) GetPositionalShadowAtlasQuadrantSubdiv(quadrant int) classdb.ViewportPositionalShadowAtlasQuadrantSubdiv {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportPositionalShadowAtlasQuadrantSubdiv(Expert(self).GetPositionalShadowAtlasQuadrantSubdiv(gd.Int(quadrant)))
}
func (self Simple) SetInputAsHandled() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputAsHandled()
}
func (self Simple) IsInputHandled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsInputHandled())
}
func (self Simple) SetHandleInputLocally(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHandleInputLocally(enable)
}
func (self Simple) IsHandlingInputLocally() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHandlingInputLocally())
}
func (self Simple) SetDefaultCanvasItemTextureFilter(mode classdb.ViewportDefaultCanvasItemTextureFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultCanvasItemTextureFilter(mode)
}
func (self Simple) GetDefaultCanvasItemTextureFilter() classdb.ViewportDefaultCanvasItemTextureFilter {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportDefaultCanvasItemTextureFilter(Expert(self).GetDefaultCanvasItemTextureFilter())
}
func (self Simple) SetEmbeddingSubwindows(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmbeddingSubwindows(enable)
}
func (self Simple) IsEmbeddingSubwindows() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEmbeddingSubwindows())
}
func (self Simple) GetEmbeddedSubwindows() gd.ArrayOf[[1]classdb.Window] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.Window](Expert(self).GetEmbeddedSubwindows(gc))
}
func (self Simple) SetCanvasCullMask(mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCanvasCullMask(gd.Int(mask))
}
func (self Simple) GetCanvasCullMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCanvasCullMask()))
}
func (self Simple) SetCanvasCullMaskBit(layer int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCanvasCullMaskBit(gd.Int(layer), enable)
}
func (self Simple) GetCanvasCullMaskBit(layer int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetCanvasCullMaskBit(gd.Int(layer)))
}
func (self Simple) SetDefaultCanvasItemTextureRepeat(mode classdb.ViewportDefaultCanvasItemTextureRepeat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultCanvasItemTextureRepeat(mode)
}
func (self Simple) GetDefaultCanvasItemTextureRepeat() classdb.ViewportDefaultCanvasItemTextureRepeat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportDefaultCanvasItemTextureRepeat(Expert(self).GetDefaultCanvasItemTextureRepeat())
}
func (self Simple) SetSdfOversize(oversize classdb.ViewportSDFOversize) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfOversize(oversize)
}
func (self Simple) GetSdfOversize() classdb.ViewportSDFOversize {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportSDFOversize(Expert(self).GetSdfOversize())
}
func (self Simple) SetSdfScale(scale classdb.ViewportSDFScale) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfScale(scale)
}
func (self Simple) GetSdfScale() classdb.ViewportSDFScale {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportSDFScale(Expert(self).GetSdfScale())
}
func (self Simple) SetMeshLodThreshold(pixels float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMeshLodThreshold(gd.Float(pixels))
}
func (self Simple) GetMeshLodThreshold() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMeshLodThreshold()))
}
func (self Simple) SetAsAudioListener2d(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAsAudioListener2d(enable)
}
func (self Simple) IsAudioListener2d() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAudioListener2d())
}
func (self Simple) GetCamera2d() [1]classdb.Camera2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Camera2D(Expert(self).GetCamera2d(gc))
}
func (self Simple) SetWorld3d(world_3d [1]classdb.World3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWorld3d(world_3d)
}
func (self Simple) GetWorld3d() [1]classdb.World3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.World3D(Expert(self).GetWorld3d(gc))
}
func (self Simple) FindWorld3d() [1]classdb.World3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.World3D(Expert(self).FindWorld3d(gc))
}
func (self Simple) SetUseOwnWorld3d(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseOwnWorld3d(enable)
}
func (self Simple) IsUsingOwnWorld3d() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingOwnWorld3d())
}
func (self Simple) GetCamera3d() [1]classdb.Camera3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Camera3D(Expert(self).GetCamera3d(gc))
}
func (self Simple) SetAsAudioListener3d(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAsAudioListener3d(enable)
}
func (self Simple) IsAudioListener3d() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAudioListener3d())
}
func (self Simple) SetDisable3d(disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisable3d(disable)
}
func (self Simple) Is3dDisabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).Is3dDisabled())
}
func (self Simple) SetUseXr(use bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseXr(use)
}
func (self Simple) IsUsingXr() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingXr())
}
func (self Simple) SetScaling3dMode(scaling_3d_mode classdb.ViewportScaling3DMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScaling3dMode(scaling_3d_mode)
}
func (self Simple) GetScaling3dMode() classdb.ViewportScaling3DMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportScaling3DMode(Expert(self).GetScaling3dMode())
}
func (self Simple) SetScaling3dScale(scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScaling3dScale(gd.Float(scale))
}
func (self Simple) GetScaling3dScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetScaling3dScale()))
}
func (self Simple) SetFsrSharpness(fsr_sharpness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFsrSharpness(gd.Float(fsr_sharpness))
}
func (self Simple) GetFsrSharpness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFsrSharpness()))
}
func (self Simple) SetTextureMipmapBias(texture_mipmap_bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureMipmapBias(gd.Float(texture_mipmap_bias))
}
func (self Simple) GetTextureMipmapBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTextureMipmapBias()))
}
func (self Simple) SetVrsMode(mode classdb.ViewportVRSMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVrsMode(mode)
}
func (self Simple) GetVrsMode() classdb.ViewportVRSMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportVRSMode(Expert(self).GetVrsMode())
}
func (self Simple) SetVrsUpdateMode(mode classdb.ViewportVRSUpdateMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVrsUpdateMode(mode)
}
func (self Simple) GetVrsUpdateMode() classdb.ViewportVRSUpdateMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ViewportVRSUpdateMode(Expert(self).GetVrsUpdateMode())
}
func (self Simple) SetVrsTexture(texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVrsTexture(texture)
}
func (self Simple) GetVrsTexture() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetVrsTexture(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Viewport
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetWorld2d(world_2d object.World2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(world_2d[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_world_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWorld2d(ctx gd.Lifetime) object.World2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_world_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.World2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the first valid [World2D] for this viewport, searching the [member world_2d] property of itself and any Viewport ancestor.
*/
//go:nosplit
func (self class) FindWorld2d(ctx gd.Lifetime) object.World2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_find_world_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.World2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCanvasTransform(xform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCanvasTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlobalCanvasTransform(xform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_global_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalCanvasTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_global_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transform from the viewport's coordinate system to the embedder's coordinate system.
*/
//go:nosplit
func (self class) GetFinalTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_final_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transform from the Viewport's coordinates to the screen coordinates of the containing window manager window.
*/
//go:nosplit
func (self class) GetScreenTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_screen_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the visible rectangle in global screen coordinates.
*/
//go:nosplit
func (self class) GetVisibleRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_visible_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransparentBackground(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_transparent_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasTransparentBackground() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_has_transparent_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseHdr2d(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_use_hdr_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingHdr2d() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_using_hdr_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMsaa2d(msaa classdb.ViewportMSAA)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_msaa_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsaa2d() classdb.ViewportMSAA {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportMSAA](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_msaa_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMsaa3d(msaa classdb.ViewportMSAA)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsaa3d() classdb.ViewportMSAA {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportMSAA](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScreenSpaceAa(screen_space_aa classdb.ViewportScreenSpaceAA)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen_space_aa)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScreenSpaceAa() classdb.ViewportScreenSpaceAA {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportScreenSpaceAA](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseTaa(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_use_taa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingTaa() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_using_taa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseDebanding(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingDebanding() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_using_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseOcclusionCulling(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_use_occlusion_culling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingOcclusionCulling() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_using_occlusion_culling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugDraw(debug_draw classdb.ViewportDebugDraw)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, debug_draw)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_debug_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDebugDraw() classdb.ViewportDebugDraw {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportDebugDraw](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_debug_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns rendering statistics of the given type. See [enum RenderInfoType] and [enum RenderInfo] for options.
*/
//go:nosplit
func (self class) GetRenderInfo(atype classdb.ViewportRenderInfoType, info classdb.ViewportRenderInfo) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, info)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_render_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the viewport's texture.
[b]Note:[/b] When trying to store the current texture (e.g. in a file), it might be completely black or outdated if used too early, especially when used in e.g. [method Node._ready]. To make sure the texture you get is correct, you can await [signal RenderingServer.frame_post_draw] signal.
[codeblock]
func _ready():
    await RenderingServer.frame_post_draw
    $Viewport.get_texture().get_image().save_png("user://Screenshot.png")
[/codeblock]
*/
//go:nosplit
func (self class) GetTexture(ctx gd.Lifetime) object.ViewportTexture {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.ViewportTexture
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPhysicsObjectPicking(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_physics_object_picking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsObjectPicking() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_physics_object_picking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPhysicsObjectPickingSort(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_physics_object_picking_sort, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsObjectPickingSort() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_physics_object_picking_sort, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPhysicsObjectPickingFirstOnly(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_physics_object_picking_first_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsObjectPickingFirstOnly() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_physics_object_picking_first_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the viewport's RID from the [RenderingServer].
*/
//go:nosplit
func (self class) GetViewportRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_viewport_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Helper method which calls the [code]set_text()[/code] method on the currently focused [Control], provided that it is defined (e.g. if the focused Control is [Button] or [LineEdit]).
*/
//go:nosplit
func (self class) PushTextInput(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_push_text_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Triggers the given [param event] in this [Viewport]. This can be used to pass an [InputEvent] between viewports, or to locally apply inputs that were sent over the network or saved to a file.
If [param in_local_coords] is [code]false[/code], the event's position is in the embedder's coordinates and will be converted to viewport coordinates. If [param in_local_coords] is [code]true[/code], the event's position is in viewport coordinates.
While this method serves a similar purpose as [method Input.parse_input_event], it does not remap the specified [param event] based on project settings like [member ProjectSettings.input_devices/pointing/emulate_touch_from_mouse].
Calling this method will propagate calls to child nodes for following methods in the given order:
- [method Node._input]
- [method Control._gui_input] for [Control] nodes
- [method Node._shortcut_input]
- [method Node._unhandled_key_input]
- [method Node._unhandled_input]
If an earlier method marks the input as handled via [method set_input_as_handled], any later method in this list will not be called.
If none of the methods handle the event and [member physics_object_picking] is [code]true[/code], the event is used for physics object picking.
*/
//go:nosplit
func (self class) PushInput(event object.InputEvent, in_local_coords bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(event[0].AsPointer())[0])
	callframe.Arg(frame, in_local_coords)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_push_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Triggers the given [param event] in this [Viewport]. This can be used to pass an [InputEvent] between viewports, or to locally apply inputs that were sent over the network or saved to a file.
If [param in_local_coords] is [code]false[/code], the event's position is in the embedder's coordinates and will be converted to viewport coordinates. If [param in_local_coords] is [code]true[/code], the event's position is in viewport coordinates.
Calling this method will propagate calls to child nodes for following methods in the given order:
- [method Node._shortcut_input]
- [method Node._unhandled_key_input]
- [method Node._unhandled_input]
If an earlier method marks the input as handled via [method set_input_as_handled], any later method in this list will not be called.
If none of the methods handle the event and [member physics_object_picking] is [code]true[/code], the event is used for physics object picking.
[b]Note:[/b] This method doesn't propagate input events to embedded [Window]s or [SubViewport]s.
*/
//go:nosplit
func (self class) PushUnhandledInput(event object.InputEvent, in_local_coords bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(event[0].AsPointer())[0])
	callframe.Arg(frame, in_local_coords)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_push_unhandled_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the mouse's position in this [Viewport] using the coordinate system of this [Viewport].
*/
//go:nosplit
func (self class) GetMousePosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_mouse_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves the mouse pointer to the specified position in this [Viewport] using the coordinate system of this [Viewport].
[b]Note:[/b] [method warp_mouse] is only supported on Windows, macOS and Linux. It has no effect on Android, iOS and Web.
*/
//go:nosplit
func (self class) WarpMouse(position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_warp_mouse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Force instantly updating the display based on the current mouse cursor position. This includes updating the mouse cursor shape and sending necessary [signal Control.mouse_entered], [signal CollisionObject2D.mouse_entered], [signal CollisionObject3D.mouse_entered] and [signal Window.mouse_entered] signals and their respective [code]mouse_exited[/code] counterparts.
*/
//go:nosplit
func (self class) UpdateMouseCursorState()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_update_mouse_cursor_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the drag data from the GUI, that was previously returned by [method Control._get_drag_data].
*/
//go:nosplit
func (self class) GuiGetDragData(ctx gd.Lifetime) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_gui_get_drag_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the viewport is currently performing a drag operation.
Alternative to [constant Node.NOTIFICATION_DRAG_BEGIN] and [constant Node.NOTIFICATION_DRAG_END] when you prefer polling the value.
*/
//go:nosplit
func (self class) GuiIsDragging() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_gui_is_dragging, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the drag operation is successful.
*/
//go:nosplit
func (self class) GuiIsDragSuccessful() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_gui_is_drag_successful, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the focus from the currently focused [Control] within this viewport. If no [Control] has the focus, does nothing.
*/
//go:nosplit
func (self class) GuiReleaseFocus()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_gui_release_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Control] having the focus within this viewport. If no [Control] has the focus, returns null.
*/
//go:nosplit
func (self class) GuiGetFocusOwner(ctx gd.Lifetime) object.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_gui_get_focus_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Control
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [Control] that the mouse is currently hovering over in this viewport. If no [Control] has the cursor, returns null.
Typically the leaf [Control] node or deepest level of the subtree which claims hover. This is very useful when used together with [method Node.is_ancestor_of] to find if the mouse is within a control tree.
*/
//go:nosplit
func (self class) GuiGetHoveredControl(ctx gd.Lifetime) object.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_gui_get_hovered_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Control
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisableInput(disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_disable_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsInputDisabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_input_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPositionalShadowAtlasSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_positional_shadow_atlas_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPositionalShadowAtlasSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_positional_shadow_atlas_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPositionalShadowAtlas16Bits(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_positional_shadow_atlas_16_bits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPositionalShadowAtlas16Bits() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_positional_shadow_atlas_16_bits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSnapControlsToPixels(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_snap_controls_to_pixels, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSnapControlsToPixelsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_snap_controls_to_pixels_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSnap2dTransformsToPixel(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_snap_2d_transforms_to_pixel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSnap2dTransformsToPixelEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_snap_2d_transforms_to_pixel_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSnap2dVerticesToPixel(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_snap_2d_vertices_to_pixel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSnap2dVerticesToPixelEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_snap_2d_vertices_to_pixel_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the number of subdivisions to use in the specified quadrant. A higher number of subdivisions allows you to have more shadows in the scene at once, but reduces the quality of the shadows. A good practice is to have quadrants with a varying number of subdivisions and to have as few subdivisions as possible.
*/
//go:nosplit
func (self class) SetPositionalShadowAtlasQuadrantSubdiv(quadrant gd.Int, subdiv classdb.ViewportPositionalShadowAtlasQuadrantSubdiv)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quadrant)
	callframe.Arg(frame, subdiv)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_positional_shadow_atlas_quadrant_subdiv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the positional shadow atlas quadrant subdivision of the specified quadrant.
*/
//go:nosplit
func (self class) GetPositionalShadowAtlasQuadrantSubdiv(quadrant gd.Int) classdb.ViewportPositionalShadowAtlasQuadrantSubdiv {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quadrant)
	var r_ret = callframe.Ret[classdb.ViewportPositionalShadowAtlasQuadrantSubdiv](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_positional_shadow_atlas_quadrant_subdiv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Stops the input from propagating further down the [SceneTree].
[b]Note:[/b] This does not affect the methods in [Input], only the way events are propagated.
*/
//go:nosplit
func (self class) SetInputAsHandled()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_input_as_handled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the current [InputEvent] has been handled. Input events are not handled until [method set_input_as_handled] has been called during the lifetime of an [InputEvent].
This is usually done as part of input handling methods like [method Node._input], [method Control._gui_input] or others, as well as in corresponding signal handlers.
If [member handle_input_locally] is set to [code]false[/code], this method will try finding the first parent viewport that is set to handle input locally, and return its value for [method is_input_handled] instead.
*/
//go:nosplit
func (self class) IsInputHandled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_input_handled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHandleInputLocally(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_handle_input_locally, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHandlingInputLocally() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_handling_input_locally, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultCanvasItemTextureFilter(mode classdb.ViewportDefaultCanvasItemTextureFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_default_canvas_item_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultCanvasItemTextureFilter() classdb.ViewportDefaultCanvasItemTextureFilter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportDefaultCanvasItemTextureFilter](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_default_canvas_item_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmbeddingSubwindows(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_embedding_subwindows, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEmbeddingSubwindows() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_embedding_subwindows, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a list of the visible embedded [Window]s inside the viewport.
[b]Note:[/b] [Window]s inside other viewports will not be listed.
*/
//go:nosplit
func (self class) GetEmbeddedSubwindows(ctx gd.Lifetime) gd.ArrayOf[object.Window] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_embedded_subwindows, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.Window](ret)
}
//go:nosplit
func (self class) SetCanvasCullMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_canvas_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCanvasCullMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_canvas_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set/clear individual bits on the rendering layer mask. This simplifies editing this [Viewport]'s layers.
*/
//go:nosplit
func (self class) SetCanvasCullMaskBit(layer gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_canvas_cull_mask_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an individual bit on the rendering layer mask.
*/
//go:nosplit
func (self class) GetCanvasCullMaskBit(layer gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_canvas_cull_mask_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultCanvasItemTextureRepeat(mode classdb.ViewportDefaultCanvasItemTextureRepeat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_default_canvas_item_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultCanvasItemTextureRepeat() classdb.ViewportDefaultCanvasItemTextureRepeat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportDefaultCanvasItemTextureRepeat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_default_canvas_item_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfOversize(oversize classdb.ViewportSDFOversize)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, oversize)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_sdf_oversize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfOversize() classdb.ViewportSDFOversize {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportSDFOversize](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_sdf_oversize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfScale(scale classdb.ViewportSDFScale)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_sdf_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfScale() classdb.ViewportSDFScale {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportSDFScale](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_sdf_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMeshLodThreshold(pixels gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_mesh_lod_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMeshLodThreshold() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_mesh_lod_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAsAudioListener2d(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_as_audio_listener_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAudioListener2d() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_audio_listener_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the currently active 2D camera. Returns null if there are no active cameras.
*/
//go:nosplit
func (self class) GetCamera2d(ctx gd.Lifetime) object.Camera2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_camera_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Camera2D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWorld3d(world_3d object.World3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(world_3d[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWorld3d(ctx gd.Lifetime) object.World3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.World3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the first valid [World3D] for this viewport, searching the [member world_3d] property of itself and any Viewport ancestor.
*/
//go:nosplit
func (self class) FindWorld3d(ctx gd.Lifetime) object.World3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_find_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.World3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseOwnWorld3d(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_use_own_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingOwnWorld3d() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_using_own_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the currently active 3D camera.
*/
//go:nosplit
func (self class) GetCamera3d(ctx gd.Lifetime) object.Camera3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_camera_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Camera3D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAsAudioListener3d(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_as_audio_listener_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAudioListener3d() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_audio_listener_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisable3d(disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_disable_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) Is3dDisabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_3d_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseXr(use bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_use_xr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingXr() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_is_using_xr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScaling3dMode(scaling_3d_mode classdb.ViewportScaling3DMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scaling_3d_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScaling3dMode() classdb.ViewportScaling3DMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportScaling3DMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScaling3dScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_scaling_3d_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScaling3dScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_scaling_3d_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFsrSharpness(fsr_sharpness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fsr_sharpness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFsrSharpness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureMipmapBias(texture_mipmap_bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture_mipmap_bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_texture_mipmap_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureMipmapBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_texture_mipmap_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVrsMode(mode classdb.ViewportVRSMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_vrs_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVrsMode() classdb.ViewportVRSMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportVRSMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_vrs_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVrsUpdateMode(mode classdb.ViewportVRSUpdateMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_vrs_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVrsUpdateMode() classdb.ViewportVRSUpdateMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportVRSUpdateMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_vrs_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVrsTexture(texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_set_vrs_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVrsTexture(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Viewport.Bind_get_vrs_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsViewport() Expert { return self[0].AsViewport() }


//go:nosplit
func (self Simple) AsViewport() Simple { return self[0].AsViewport() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Viewport", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type PositionalShadowAtlasQuadrantSubdiv = classdb.ViewportPositionalShadowAtlasQuadrantSubdiv

const (
/*This quadrant will not be used.*/
	ShadowAtlasQuadrantSubdivDisabled PositionalShadowAtlasQuadrantSubdiv = 0
/*This quadrant will only be used by one shadow map.*/
	ShadowAtlasQuadrantSubdiv1 PositionalShadowAtlasQuadrantSubdiv = 1
/*This quadrant will be split in 4 and used by up to 4 shadow maps.*/
	ShadowAtlasQuadrantSubdiv4 PositionalShadowAtlasQuadrantSubdiv = 2
/*This quadrant will be split 16 ways and used by up to 16 shadow maps.*/
	ShadowAtlasQuadrantSubdiv16 PositionalShadowAtlasQuadrantSubdiv = 3
/*This quadrant will be split 64 ways and used by up to 64 shadow maps.*/
	ShadowAtlasQuadrantSubdiv64 PositionalShadowAtlasQuadrantSubdiv = 4
/*This quadrant will be split 256 ways and used by up to 256 shadow maps. Unless the [member positional_shadow_atlas_size] is very high, the shadows in this quadrant will be very low resolution.*/
	ShadowAtlasQuadrantSubdiv256 PositionalShadowAtlasQuadrantSubdiv = 5
/*This quadrant will be split 1024 ways and used by up to 1024 shadow maps. Unless the [member positional_shadow_atlas_size] is very high, the shadows in this quadrant will be very low resolution.*/
	ShadowAtlasQuadrantSubdiv1024 PositionalShadowAtlasQuadrantSubdiv = 6
/*Represents the size of the [enum PositionalShadowAtlasQuadrantSubdiv] enum.*/
	ShadowAtlasQuadrantSubdivMax PositionalShadowAtlasQuadrantSubdiv = 7
)
type Scaling3DMode = classdb.ViewportScaling3DMode

const (
/*Use bilinear scaling for the viewport's 3D buffer. The amount of scaling can be set using [member scaling_3d_scale]. Values less than [code]1.0[/code] will result in undersampling while values greater than [code]1.0[/code] will result in supersampling. A value of [code]1.0[/code] disables scaling.*/
	Scaling3dModeBilinear Scaling3DMode = 0
/*Use AMD FidelityFX Super Resolution 1.0 upscaling for the viewport's 3D buffer. The amount of scaling can be set using [member scaling_3d_scale]. Values less than [code]1.0[/code] will be result in the viewport being upscaled using FSR. Values greater than [code]1.0[/code] are not supported and bilinear downsampling will be used instead. A value of [code]1.0[/code] disables scaling.*/
	Scaling3dModeFsr Scaling3DMode = 1
/*Use AMD FidelityFX Super Resolution 2.2 upscaling for the viewport's 3D buffer. The amount of scaling can be set using [member Viewport.scaling_3d_scale]. Values less than [code]1.0[/code] will be result in the viewport being upscaled using FSR2. Values greater than [code]1.0[/code] are not supported and bilinear downsampling will be used instead. A value of [code]1.0[/code] will use FSR2 at native resolution as a TAA solution.*/
	Scaling3dModeFsr2 Scaling3DMode = 2
/*Represents the size of the [enum Scaling3DMode] enum.*/
	Scaling3dModeMax Scaling3DMode = 3
)
type MSAA = classdb.ViewportMSAA

const (
/*Multisample antialiasing mode disabled. This is the default value, and is also the fastest setting.*/
	MsaaDisabled MSAA = 0
/*Use 2× Multisample Antialiasing. This has a moderate performance cost. It helps reduce aliasing noticeably, but 4× MSAA still looks substantially better.*/
	Msaa2x MSAA = 1
/*Use 4× Multisample Antialiasing. This has a significant performance cost, and is generally a good compromise between performance and quality.*/
	Msaa4x MSAA = 2
/*Use 8× Multisample Antialiasing. This has a very high performance cost. The difference between 4× and 8× MSAA may not always be visible in real gameplay conditions. Likely unsupported on low-end and older hardware.*/
	Msaa8x MSAA = 3
/*Represents the size of the [enum MSAA] enum.*/
	MsaaMax MSAA = 4
)
type ScreenSpaceAA = classdb.ViewportScreenSpaceAA

const (
/*Do not perform any antialiasing in the full screen post-process.*/
	ScreenSpaceAaDisabled ScreenSpaceAA = 0
/*Use fast approximate antialiasing. FXAA is a popular screen-space antialiasing method, which is fast but will make the image look blurry, especially at lower resolutions. It can still work relatively well at large resolutions such as 1440p and 4K.*/
	ScreenSpaceAaFxaa ScreenSpaceAA = 1
/*Represents the size of the [enum ScreenSpaceAA] enum.*/
	ScreenSpaceAaMax ScreenSpaceAA = 2
)
type RenderInfo = classdb.ViewportRenderInfo

const (
/*Amount of objects in frame.*/
	RenderInfoObjectsInFrame RenderInfo = 0
/*Amount of vertices in frame.*/
	RenderInfoPrimitivesInFrame RenderInfo = 1
/*Amount of draw calls in frame.*/
	RenderInfoDrawCallsInFrame RenderInfo = 2
/*Represents the size of the [enum RenderInfo] enum.*/
	RenderInfoMax RenderInfo = 3
)
type RenderInfoType = classdb.ViewportRenderInfoType

const (
/*Visible render pass (excluding shadows).*/
	RenderInfoTypeVisible RenderInfoType = 0
/*Shadow render pass. Objects will be rendered several times depending on the number of amounts of lights with shadows and the number of directional shadow splits.*/
	RenderInfoTypeShadow RenderInfoType = 1
/*Canvas item rendering. This includes all 2D rendering.*/
	RenderInfoTypeCanvas RenderInfoType = 2
/*Represents the size of the [enum RenderInfoType] enum.*/
	RenderInfoTypeMax RenderInfoType = 3
)
type DebugDraw = classdb.ViewportDebugDraw

const (
/*Objects are displayed normally.*/
	DebugDrawDisabled DebugDraw = 0
/*Objects are displayed without light information.*/
	DebugDrawUnshaded DebugDraw = 1
/*Objects are displayed without textures and only with lighting information.*/
	DebugDrawLighting DebugDraw = 2
/*Objects are displayed semi-transparent with additive blending so you can see where they are drawing over top of one another. A higher overdraw means you are wasting performance on drawing pixels that are being hidden behind others.*/
	DebugDrawOverdraw DebugDraw = 3
/*Objects are displayed as wireframe models.*/
	DebugDrawWireframe DebugDraw = 4
/*Objects are displayed without lighting information and their textures replaced by normal mapping.*/
	DebugDrawNormalBuffer DebugDraw = 5
/*Objects are displayed with only the albedo value from [VoxelGI]s.*/
	DebugDrawVoxelGiAlbedo DebugDraw = 6
/*Objects are displayed with only the lighting value from [VoxelGI]s.*/
	DebugDrawVoxelGiLighting DebugDraw = 7
/*Objects are displayed with only the emission color from [VoxelGI]s.*/
	DebugDrawVoxelGiEmission DebugDraw = 8
/*Draws the shadow atlas that stores shadows from [OmniLight3D]s and [SpotLight3D]s in the upper left quadrant of the [Viewport].*/
	DebugDrawShadowAtlas DebugDraw = 9
/*Draws the shadow atlas that stores shadows from [DirectionalLight3D]s in the upper left quadrant of the [Viewport].*/
	DebugDrawDirectionalShadowAtlas DebugDraw = 10
/*Draws the scene luminance buffer (if available) in the upper left quadrant of the [Viewport].*/
	DebugDrawSceneLuminance DebugDraw = 11
/*Draws the screen-space ambient occlusion texture instead of the scene so that you can clearly see how it is affecting objects. In order for this display mode to work, you must have [member Environment.ssao_enabled] set in your [WorldEnvironment].*/
	DebugDrawSsao DebugDraw = 12
/*Draws the screen-space indirect lighting texture instead of the scene so that you can clearly see how it is affecting objects. In order for this display mode to work, you must have [member Environment.ssil_enabled] set in your [WorldEnvironment].*/
	DebugDrawSsil DebugDraw = 13
/*Colors each PSSM split for the [DirectionalLight3D]s in the scene a different color so you can see where the splits are. In order, they will be colored red, green, blue, and yellow.*/
	DebugDrawPssmSplits DebugDraw = 14
/*Draws the decal atlas used by [Decal]s and light projector textures in the upper left quadrant of the [Viewport].*/
	DebugDrawDecalAtlas DebugDraw = 15
/*Draws the cascades used to render signed distance field global illumination (SDFGI).
Does nothing if the current environment's [member Environment.sdfgi_enabled] is [code]false[/code] or SDFGI is not supported on the platform.*/
	DebugDrawSdfgi DebugDraw = 16
/*Draws the probes used for signed distance field global illumination (SDFGI).
Does nothing if the current environment's [member Environment.sdfgi_enabled] is [code]false[/code] or SDFGI is not supported on the platform.*/
	DebugDrawSdfgiProbes DebugDraw = 17
/*Draws the buffer used for global illumination (GI).*/
	DebugDrawGiBuffer DebugDraw = 18
/*Draws all of the objects at their highest polycount, without low level of detail (LOD).*/
	DebugDrawDisableLod DebugDraw = 19
/*Draws the cluster used by [OmniLight3D] nodes to optimize light rendering.*/
	DebugDrawClusterOmniLights DebugDraw = 20
/*Draws the cluster used by [SpotLight3D] nodes to optimize light rendering.*/
	DebugDrawClusterSpotLights DebugDraw = 21
/*Draws the cluster used by [Decal] nodes to optimize decal rendering.*/
	DebugDrawClusterDecals DebugDraw = 22
/*Draws the cluster used by [ReflectionProbe] nodes to optimize decal rendering.*/
	DebugDrawClusterReflectionProbes DebugDraw = 23
/*Draws the buffer used for occlusion culling.*/
	DebugDrawOccluders DebugDraw = 24
/*Draws vector lines over the viewport to indicate the movement of pixels between frames.*/
	DebugDrawMotionVectors DebugDraw = 25
/*Draws the internal resolution buffer of the scene before post-processing is applied.*/
	DebugDrawInternalBuffer DebugDraw = 26
)
type DefaultCanvasItemTextureFilter = classdb.ViewportDefaultCanvasItemTextureFilter

const (
/*The texture filter reads from the nearest pixel only. This makes the texture look pixelated from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	DefaultCanvasItemTextureFilterNearest DefaultCanvasItemTextureFilter = 0
/*The texture filter blends between the nearest 4 pixels. This makes the texture look smooth from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	DefaultCanvasItemTextureFilterLinear DefaultCanvasItemTextureFilter = 1
/*The texture filter blends between the nearest 4 pixels and between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look smooth from up close, and smooth from a distance.
Use this for non-pixel art textures that may be viewed at a low scale (e.g. due to [Camera2D] zoom or sprite scaling), as mipmaps are important to smooth out pixels that are smaller than on-screen pixels.*/
	DefaultCanvasItemTextureFilterLinearWithMipmaps DefaultCanvasItemTextureFilter = 2
/*The texture filter reads from the nearest pixel and blends between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look pixelated from up close, and smooth from a distance.
Use this for non-pixel art textures that may be viewed at a low scale (e.g. due to [Camera2D] zoom or sprite scaling), as mipmaps are important to smooth out pixels that are smaller than on-screen pixels.*/
	DefaultCanvasItemTextureFilterNearestWithMipmaps DefaultCanvasItemTextureFilter = 3
/*Represents the size of the [enum DefaultCanvasItemTextureFilter] enum.*/
	DefaultCanvasItemTextureFilterMax DefaultCanvasItemTextureFilter = 4
)
type DefaultCanvasItemTextureRepeat = classdb.ViewportDefaultCanvasItemTextureRepeat

const (
/*Disables textures repeating. Instead, when reading UVs outside the 0-1 range, the value will be clamped to the edge of the texture, resulting in a stretched out look at the borders of the texture.*/
	DefaultCanvasItemTextureRepeatDisabled DefaultCanvasItemTextureRepeat = 0
/*Enables the texture to repeat when UV coordinates are outside the 0-1 range. If using one of the linear filtering modes, this can result in artifacts at the edges of a texture when the sampler filters across the edges of the texture.*/
	DefaultCanvasItemTextureRepeatEnabled DefaultCanvasItemTextureRepeat = 1
/*Flip the texture when repeating so that the edge lines up instead of abruptly changing.*/
	DefaultCanvasItemTextureRepeatMirror DefaultCanvasItemTextureRepeat = 2
/*Represents the size of the [enum DefaultCanvasItemTextureRepeat] enum.*/
	DefaultCanvasItemTextureRepeatMax DefaultCanvasItemTextureRepeat = 3
)
type SDFOversize = classdb.ViewportSDFOversize

const (
/*The signed distance field only covers the viewport's own rectangle.*/
	SdfOversize100Percent SDFOversize = 0
/*The signed distance field is expanded to cover 20% of the viewport's size around the borders.*/
	SdfOversize120Percent SDFOversize = 1
/*The signed distance field is expanded to cover 50% of the viewport's size around the borders.*/
	SdfOversize150Percent SDFOversize = 2
/*The signed distance field is expanded to cover 100% (double) of the viewport's size around the borders.*/
	SdfOversize200Percent SDFOversize = 3
/*Represents the size of the [enum SDFOversize] enum.*/
	SdfOversizeMax SDFOversize = 4
)
type SDFScale = classdb.ViewportSDFScale

const (
/*The signed distance field is rendered at full resolution.*/
	SdfScale100Percent SDFScale = 0
/*The signed distance field is rendered at half the resolution of this viewport.*/
	SdfScale50Percent SDFScale = 1
/*The signed distance field is rendered at a quarter the resolution of this viewport.*/
	SdfScale25Percent SDFScale = 2
/*Represents the size of the [enum SDFScale] enum.*/
	SdfScaleMax SDFScale = 3
)
type VRSMode = classdb.ViewportVRSMode

const (
/*Variable Rate Shading is disabled.*/
	VrsDisabled VRSMode = 0
/*Variable Rate Shading uses a texture. Note, for stereoscopic use a texture atlas with a texture for each view.*/
	VrsTexture VRSMode = 1
/*Variable Rate Shading's texture is supplied by the primary [XRInterface].*/
	VrsXr VRSMode = 2
/*Represents the size of the [enum VRSMode] enum.*/
	VrsMax VRSMode = 3
)
type VRSUpdateMode = classdb.ViewportVRSUpdateMode

const (
/*The input texture for variable rate shading will not be processed.*/
	VrsUpdateDisabled VRSUpdateMode = 0
/*The input texture for variable rate shading will be processed once.*/
	VrsUpdateOnce VRSUpdateMode = 1
/*The input texture for variable rate shading will be processed each frame.*/
	VrsUpdateAlways VRSUpdateMode = 2
/*Represents the size of the [enum VRSUpdateMode] enum.*/
	VrsUpdateMax VRSUpdateMode = 3
)
