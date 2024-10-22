package Light2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Casts light in a 2D environment. A light is defined as a color, an energy value, a mode (see constants), and various other parameters (range and shadows-related).

*/
type Go [1]classdb.Light2D

/*
Sets the light's height, which is used in 2D normal mapping. See [member PointLight2D.height] and [member DirectionalLight2D.height].
*/
func (self Go) SetHeight(height float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHeight(gd.Float(height))
}

/*
Returns the light's height, which is used in 2D normal mapping. See [member PointLight2D.height] and [member DirectionalLight2D.height].
*/
func (self Go) GetHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetHeight()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Light2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Light2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Enabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEnabled())
}

func (self Go) SetEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnabled(value)
}

func (self Go) EditorOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEditorOnly())
}

func (self Go) SetEditorOnly(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditorOnly(value)
}

func (self Go) Color() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetColor())
}

func (self Go) SetColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColor(value)
}

func (self Go) Energy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEnergy()))
}

func (self Go) SetEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnergy(gd.Float(value))
}

func (self Go) BlendMode() classdb.Light2DBlendMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Light2DBlendMode(class(self).GetBlendMode())
}

func (self Go) SetBlendMode(value classdb.Light2DBlendMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBlendMode(value)
}

func (self Go) RangeZMin() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetZRangeMin()))
}

func (self Go) SetRangeZMin(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetZRangeMin(gd.Int(value))
}

func (self Go) RangeZMax() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetZRangeMax()))
}

func (self Go) SetRangeZMax(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetZRangeMax(gd.Int(value))
}

func (self Go) RangeLayerMin() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLayerRangeMin()))
}

func (self Go) SetRangeLayerMin(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLayerRangeMin(gd.Int(value))
}

func (self Go) RangeLayerMax() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLayerRangeMax()))
}

func (self Go) SetRangeLayerMax(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLayerRangeMax(gd.Int(value))
}

func (self Go) RangeItemCullMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetItemCullMask()))
}

func (self Go) SetRangeItemCullMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetItemCullMask(gd.Int(value))
}

func (self Go) ShadowEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShadowEnabled())
}

func (self Go) SetShadowEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShadowEnabled(value)
}

func (self Go) ShadowColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetShadowColor())
}

func (self Go) SetShadowColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShadowColor(value)
}

func (self Go) ShadowFilter() classdb.Light2DShadowFilter {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Light2DShadowFilter(class(self).GetShadowFilter())
}

func (self Go) SetShadowFilter(value classdb.Light2DShadowFilter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShadowFilter(value)
}

func (self Go) ShadowFilterSmooth() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetShadowSmooth()))
}

func (self Go) SetShadowFilterSmooth(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShadowSmooth(gd.Float(value))
}

func (self Go) ShadowItemCullMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetItemShadowCullMask()))
}

func (self Go) SetShadowItemCullMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetItemShadowCullMask(gd.Int(value))
}

//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEditorOnly(editor_only bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, editor_only)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_editor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditorOnly() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_is_editor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnergy(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZRangeMin(z gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, z)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_z_range_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZRangeMin() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_z_range_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZRangeMax(z gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, z)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_z_range_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZRangeMax() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_z_range_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLayerRangeMin(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_layer_range_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLayerRangeMin() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_layer_range_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLayerRangeMax(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_layer_range_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLayerRangeMax() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_layer_range_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetItemCullMask(item_cull_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item_cull_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_item_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetItemCullMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_item_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetItemShadowCullMask(item_shadow_cull_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item_shadow_cull_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_item_shadow_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetItemShadowCullMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_item_shadow_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadowEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_shadow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShadowEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_is_shadow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadowSmooth(smooth gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, smooth)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_shadow_smooth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadowSmooth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_shadow_smooth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadowFilter(filter classdb.Light2DShadowFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, filter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_shadow_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadowFilter() classdb.Light2DShadowFilter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Light2DShadowFilter](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_shadow_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadowColor(shadow_color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shadow_color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_shadow_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadowColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_shadow_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBlendMode(mode classdb.Light2DBlendMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBlendMode() classdb.Light2DBlendMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Light2DBlendMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the light's height, which is used in 2D normal mapping. See [member PointLight2D.height] and [member DirectionalLight2D.height].
*/
//go:nosplit
func (self class) SetHeight(height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the light's height, which is used in 2D normal mapping. See [member PointLight2D.height] and [member DirectionalLight2D.height].
*/
//go:nosplit
func (self class) GetHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light2D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsLight2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsLight2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {classdb.Register("Light2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type ShadowFilter = classdb.Light2DShadowFilter

const (
/*No filter applies to the shadow map. This provides hard shadow edges and is the fastest to render. See [member shadow_filter].*/
	ShadowFilterNone ShadowFilter = 0
/*Percentage closer filtering (5 samples) applies to the shadow map. This is slower compared to hard shadow rendering. See [member shadow_filter].*/
	ShadowFilterPcf5 ShadowFilter = 1
/*Percentage closer filtering (13 samples) applies to the shadow map. This is the slowest shadow filtering mode, and should be used sparingly. See [member shadow_filter].*/
	ShadowFilterPcf13 ShadowFilter = 2
)
type BlendMode = classdb.Light2DBlendMode

const (
/*Adds the value of pixels corresponding to the Light2D to the values of pixels under it. This is the common behavior of a light.*/
	BlendModeAdd BlendMode = 0
/*Subtracts the value of pixels corresponding to the Light2D to the values of pixels under it, resulting in inversed light effect.*/
	BlendModeSub BlendMode = 1
/*Mix the value of pixels corresponding to the Light2D to the values of pixels under it by linear interpolation.*/
	BlendModeMix BlendMode = 2
)
