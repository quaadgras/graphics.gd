package gd

import (
	"time"

	gd "grow.graphics/gd/internal"
	classdb "grow.graphics/gd/internal/classdb"
	"grow.graphics/rd"
	"grow.graphics/uc"
	"grow.graphics/xy"
	"runtime.link/ffi"
	"runtime.link/mmm"
)

// RenderingDevice returns a new [rd.Interface] implementation via Godot.
func RenderingDevice(godot Lifetime) rd.Interface {
	var device = renderingDevice{
		godot: godot.API,
		class: RenderingServer(godot).GetRenderingDevice(godot),
	}
	var itfc rd.Interface = device
	device.value = &itfc
	return itfc
}

type renderingDevice struct {
	godot *gd.API
	class classdb.RenderingDevice
	value *rd.Interface
}

func (gpu renderingDevice) scope(lifetime mmm.Lifetime) Lifetime {
	return gd.Lifetime{
		Lifetime: lifetime,
		API:      gpu.godot,
	}
}

func (gpu renderingDevice) bytes(godot Lifetime, data ffi.Bytes) gd.PackedByteArray {
	packed, ok := data.Interface.(gd.PackedByteArray)
	if ok {
		return packed
	}
	// TODO use Go memory directly? as Godot rendering device doesn't hold onto the array.
	var copy = godot.PackedByteArray()
	copy.Resize(Int(data.Len()))
	for i := 0; i < data.Len(); i++ {
		copy.SetIndex(Int(i), data.Index(i))
	}
	return copy
}

func (gpu renderingDevice) string(godot Lifetime, data ffi.Bytes) gd.String {
	packed, ok := data.Interface.(gd.String)
	if ok {
		return packed
	}
	// TODO use Go memory directly? as Godot rendering device doesn't hold onto the array.
	var copy = godot.PackedByteArray()
	copy.Resize(Int(data.Len()))
	for i := 0; i < data.Len(); i++ {
		copy.SetIndex(Int(i), data.Index(i))
	}
	return copy.GetStringFromUtf8(godot)
}

func (gpu renderingDevice) colors(godot Lifetime, data ffi.Slice[uc.Color]) gd.PackedColorArray {
	packed, ok := data.Interface.(gd.PackedColorArray)
	if ok {
		return packed
	}
	// TODO use Go memory directly? as Godot rendering device doesn't hold onto the array.
	var copy = godot.PackedColorArray()
	copy.Resize(Int(data.Len()))
	for i := 0; i < data.Len(); i++ {
		copy.SetIndex(Int(i), data.Index(i))
	}
	return copy
}

func (gpu renderingDevice) int32s(godot Lifetime, data ffi.Slice[int32]) gd.PackedInt32Array {
	packed, ok := data.Interface.(gd.PackedInt32Array)
	if ok {
		return packed
	}
	// TODO use Go memory directly? as Godot rendering device doesn't hold onto the array.
	var copy = godot.PackedInt32Array()
	copy.Resize(Int(data.Len()))
	for i := 0; i < data.Len(); i++ {
		copy.SetIndex(Int(i), data.Index(i))
	}
	return copy
}

func (gpu renderingDevice) int64s(godot Lifetime, data ffi.Slice[int64]) gd.PackedInt64Array {
	packed, ok := data.Interface.(gd.PackedInt64Array)
	if ok {
		return packed
	}
	// TODO use Go memory directly? as Godot rendering device doesn't hold onto the array.
	var copy = godot.PackedInt64Array()
	copy.Resize(Int(data.Len()))
	for i := 0; i < data.Len(); i++ {
		copy.SetIndex(Int(i), data.Index(i))
	}
	return copy
}

func (gpu renderingDevice) textures(data ffi.Managed[[]rd.Texture]) gd.ArrayOf[RID] {
	array, ok := data.Interface().(gd.ArrayOf[RID])
	if ok {
		return array
	}
	slice := data.Value()
	span := gpu.scope(data.Lifetime())
	array = NewArrayOf[RID](span)
	array.Resize(Int(len(slice)))
	for i, texture := range slice {
		array.SetIndex(Int(i), RID(mmm.Get(texture)))
	}
	data.Cache(array)
	return array
}

func (gpu renderingDevice) framebufferPasses(data ffi.Managed[[]rd.FramebufferPass]) gd.ArrayOf[classdb.RDFramebufferPass] {
	array, ok := data.Interface().(gd.ArrayOf[classdb.RDFramebufferPass])
	if ok {
		return array
	}
	slice := data.Value()
	span := gpu.scope(data.Lifetime())
	array = NewArrayOf[classdb.RDFramebufferPass](span)
	for _, pass := range slice {
		var converted = *Create(span, new(classdb.RDFramebufferPass))
		converted.SetColorAttachments(gpu.int32s(span, pass.ColorAttachments))
		converted.SetDepthAttachment(Int(pass.DepthAttachment))
		converted.SetInputAttachments(gpu.int32s(span, pass.InputAttachments))
		converted.SetPreserveAttachments(gpu.int32s(span, pass.PreserveAttachments))
		converted.SetResolveAttachments(gpu.int32s(span, pass.ResolveAttachments))
		array.Append(converted)
	}
	data.Cache(array)
	return array
}

func (gpu renderingDevice) attachments(data ffi.Managed[[]rd.AttachmentFormat]) gd.ArrayOf[classdb.RDAttachmentFormat] {
	array, ok := data.Interface().(gd.ArrayOf[classdb.RDAttachmentFormat])
	if ok {
		return array
	}
	slice := data.Value()
	span := gpu.scope(data.Lifetime())
	array = NewArrayOf[classdb.RDAttachmentFormat](span)
	for _, attachment := range slice {
		var converted = *Create(span, new(classdb.RDAttachmentFormat))
		converted.SetFormat(classdb.RenderingDeviceDataFormat(attachment.Format))
		converted.SetSamples(classdb.RenderingDeviceTextureSamples(attachment.Samples))
		converted.SetUsageFlags(int64(attachment.UsageFlags))
		array.Append(converted)
	}
	data.Cache(array)
	return array
}

func (gpu renderingDevice) Barrier(from, to rd.BarrierMask) {
	gpu.class.Barrier(classdb.RenderingDeviceBarrierMask(from), classdb.RenderingDeviceBarrierMask(to))
}

func (gpu renderingDevice) BufferClear(buffer rd.Buffer, offset, size_bytes int, post_barrier rd.BarrierMask) error {
	return Error(gpu.class.BufferClear(RID(mmm.Get(buffer)), Int(offset), Int(size_bytes), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu renderingDevice) BufferGetData(lifetime mmm.Lifetime, buffer rd.Buffer, offset_bytes, size_bytes int) ffi.Bytes {
	return ffi.Bytes{Interface: gpu.class.BufferGetData(gpu.scope(lifetime), RID(mmm.Get(buffer)), Int(offset_bytes), Int(size_bytes))}
}

func (gpu renderingDevice) BufferUpdate(buffer rd.Buffer, offset, size_bytes int, data ffi.Bytes, post_barrier rd.BarrierMask) error {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return Error(gpu.class.BufferUpdate(RID(mmm.Get(buffer)), Int(offset), Int(size_bytes), gpu.bytes(tmp, data), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu renderingDevice) CaptureTimestamp(name string) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.CaptureTimestamp(tmp.String(name))
}

func (gpu renderingDevice) GetCapturedTimestampTimeCPU(index int) time.Duration {
	return time.Duration(gpu.class.GetCapturedTimestampCpuTime(Int(index))) * time.Microsecond
}

func (gpu renderingDevice) GetCapturedTimestampTimeGPU(index int) time.Duration {
	return time.Duration(gpu.class.GetCapturedTimestampGpuTime(Int(index))) * time.Microsecond
}

func (gpu renderingDevice) GetCapturedTimestampName(index int) string {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return gpu.class.GetCapturedTimestampName(tmp, Int(index)).String()
}

func (gpu renderingDevice) GetCapturedTimestampsCount() int {
	return int(gpu.class.GetCapturedTimestampsCount())
}

func (gpu renderingDevice) GetCapturedTimestampsFrame() int {
	return int(gpu.class.GetCapturedTimestampsFrame())
}

func (gpu renderingDevice) ComputeListBegin(allow_draw_overlap bool) rd.ComputeList {
	return rd.ComputeList(gpu.class.ComputeListBegin(Bool(allow_draw_overlap)))
}

func (gpu renderingDevice) ComputeListAddBarrier(compute_list rd.ComputeList) {
	gpu.class.ComputeListAddBarrier(Int(compute_list))
}

func (gpu renderingDevice) ComputeListBindComputePipeline(list rd.ComputeList, pipeline rd.ComputePipeline) {
	gpu.class.ComputeListBindComputePipeline(Int(list), RID(mmm.Get(pipeline)))
}

func (gpu renderingDevice) ComputeListBindUniformSet(list rd.ComputeList, set rd.UniformSet, set_index int) {
	gpu.class.ComputeListBindUniformSet(Int(list), RID(mmm.Get(set)), Int(set_index))
}

func (gpu renderingDevice) ComputeListDispatch(list rd.ComputeList, groups_x, groups_y, groups_z int) {
	gpu.class.ComputeListDispatch(Int(list), Int(groups_x), Int(groups_y), Int(groups_z))
}

func (gpu renderingDevice) ComputeListEnd(mask rd.BarrierMask) {
	gpu.class.ComputeListEnd(classdb.RenderingDeviceBarrierMask(mask))
}

func (gpu renderingDevice) ComputeListSetPushConstant(list rd.ComputeList, buffer ffi.Bytes, size_bytes int) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.ComputeListSetPushConstant(Int(list), gpu.bytes(tmp, buffer), Int(size_bytes))
}

func (gpu renderingDevice) ComputePipelineCreate(span mmm.Lifetime, shader rd.Shader, constants ffi.Managed[[]rd.PipelineSpecializationConstant]) rd.ComputePipeline {
	array, ok := constants.Interface().(gd.ArrayOf[classdb.RDPipelineSpecializationConstant])
	if !ok {
		span := gpu.scope(constants.Lifetime())
		array = NewArrayOf[classdb.RDPipelineSpecializationConstant](span)
		for _, constant := range constants.Value() {
			var converted = *Create(span, new(classdb.RDPipelineSpecializationConstant))
			converted.SetConstantId(Int(constant.ConstantID))
			converted.SetValue(span.Variant(constant.Value))
			array.Append(converted)
		}
		constants.Cache(array)
	}
	pipeline := uintptr(gpu.class.ComputePipelineCreate(RID(mmm.Get(shader)), array))
	return mmm.New[rd.ComputePipeline, rd.Interface, uintptr](span, gpu.value, pipeline)
}

func (gpu renderingDevice) ComputePipelineIsValid(pipeline rd.ComputePipeline) bool {
	return bool(gpu.class.ComputePipelineIsValid(RID(mmm.Get(pipeline))))
}

func (gpu renderingDevice) CreateLocalDevice(span mmm.Lifetime) rd.Interface {
	var device = renderingDevice{
		godot: gpu.godot,
		class: gpu.class.CreateLocalDevice(gpu.scope(span)),
	}
	var itfc rd.Interface = device
	device.value = &itfc
	return itfc
}

func (gpu renderingDevice) DrawCommandBeginLabel(name string, color uc.Color) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.DrawCommandBeginLabel(tmp.String(name), color)
}

func (gpu renderingDevice) DrawCommandEndLabel() {
	gpu.class.DrawCommandEndLabel()
}

func (gpu renderingDevice) DrawCommandInsertLabel(name string, color uc.Color) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.DrawCommandInsertLabel(tmp.String(name), color)
}

func (gpu renderingDevice) DrawListBegin(fb rd.Framebuffer,
	initial_color_action rd.InitialAction, final_color_action rd.FinalAction,
	initial_depth_action rd.InitialAction, final_depth_action rd.FinalAction,
	clear_color_values ffi.Slice[uc.Color], clear_depth float64, clear_stencil int,
	region xy.Rect2, storage_textures ffi.Managed[[]rd.Texture],
) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.DrawListBegin(RID(mmm.Get(fb)),
		classdb.RenderingDeviceInitialAction(initial_color_action), classdb.RenderingDeviceFinalAction(final_color_action),
		classdb.RenderingDeviceInitialAction(initial_depth_action), classdb.RenderingDeviceFinalAction(final_depth_action),
		gpu.colors(tmp, clear_color_values), Float(clear_depth), Int(clear_stencil),
		Rect2(region), gpu.textures(storage_textures),
	)
}

func (gpu renderingDevice) DrawListBeginForScreen(screen rd.Screen, clear_color uc.Color) rd.DrawList {
	return rd.DrawList(gpu.class.DrawListBeginForScreen(Int(screen), clear_color))
}

func (gpu renderingDevice) DrawListBeginSplit(span mmm.Lifetime, fb rd.Framebuffer, splits int,
	initial_color_action rd.InitialAction, final_color_action rd.FinalAction,
	initial_depth_action rd.InitialAction, final_depth_action rd.FinalAction,
	clear_color_values ffi.Slice[uc.Color], clear_depth float64, clear_stencil int,
	region xy.Rect2, storage_textures ffi.Managed[[]rd.Texture],
) ffi.Slice[rd.DrawList] {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	lists := gpu.class.DrawListBeginSplit(gpu.scope(span), RID(mmm.Get(fb)), Int(splits),
		classdb.RenderingDeviceInitialAction(initial_color_action), classdb.RenderingDeviceFinalAction(final_color_action),
		classdb.RenderingDeviceInitialAction(initial_depth_action), classdb.RenderingDeviceFinalAction(final_depth_action),
		gpu.colors(tmp, clear_color_values), Float(clear_depth), Int(clear_stencil),
		Rect2(region), gpu.textures(storage_textures),
	)
	return ffi.Slice[rd.DrawList]{Interface: lists}
}

func (gpu renderingDevice) DrawListBindIndexArray(list rd.DrawList, array rd.IndexArray) {
	gpu.class.DrawListBindIndexArray(Int(list), RID(mmm.Get(array)))
}

func (gpu renderingDevice) DrawListBindRenderPipeline(list rd.DrawList, pipeline rd.RenderPipeline) {
	gpu.class.DrawListBindRenderPipeline(Int(list), RID(mmm.Get(pipeline)))
}

func (gpu renderingDevice) DrawListBindUniformSet(list rd.DrawList, set rd.UniformSet, set_index int) {
	gpu.class.DrawListBindUniformSet(Int(list), RID(mmm.Get(set)), Int(set_index))
}

func (gpu renderingDevice) DrawListBindVertexArray(list rd.DrawList, array rd.VertexArray) {
	gpu.class.DrawListBindVertexArray(Int(list), RID(mmm.Get(array)))
}

func (gpu renderingDevice) DrawListDisableScissor(list rd.DrawList) {
	gpu.class.DrawListDisableScissor(Int(list))
}

func (gpu renderingDevice) DrawListDraw(list rd.DrawList, use_indices bool, instances int, procedural_vertex_count int) {
	gpu.class.DrawListDraw(Int(list), Bool(use_indices), Int(instances), Int(procedural_vertex_count))
}

func (gpu renderingDevice) DrawListEnableScissor(list rd.DrawList, rect xy.Rect2) {
	gpu.class.DrawListEnableScissor(Int(list), Rect2(rect))
}

func (gpu renderingDevice) DrawListEnd(mask rd.BarrierMask) {
	gpu.class.DrawListEnd(classdb.RenderingDeviceBarrierMask(mask))
}

func (gpu renderingDevice) DrawListSetBlendConstants(list rd.DrawList, color uc.Color) {
	gpu.class.DrawListSetBlendConstants(Int(list), color)
}

func (gpu renderingDevice) DrawListSetPushConstant(list rd.DrawList, buffer ffi.Bytes, size_bytes int) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.DrawListSetPushConstant(Int(list), gpu.bytes(tmp, buffer), Int(size_bytes))
}

func (gpu renderingDevice) DrawListSwitchToNextPass() rd.DrawList {
	return rd.DrawList(gpu.class.DrawListSwitchToNextPass())
}

func (gpu renderingDevice) DrawListSwitchToNextPassSplit(span mmm.Lifetime, splits int) ffi.Slice[rd.DrawList] {
	lists := gpu.class.DrawListSwitchToNextPassSplit(gpu.scope(span), Int(splits))
	return ffi.Slice[rd.DrawList]{Interface: lists}
}

func (gpu renderingDevice) FramebufferCreate(span mmm.Lifetime, textures ffi.Managed[[]rd.Texture], validate_with_format, view_count int) rd.Framebuffer {
	framebuffer := gpu.class.FramebufferCreate(gpu.textures(textures), Int(validate_with_format), Int(view_count))
	return mmm.New[rd.Framebuffer, rd.Interface, uintptr](span, gpu.value, uintptr(framebuffer))
}

func (gpu renderingDevice) FramebufferCreateEmpty(span mmm.Lifetime, size xy.Vector2i, samples rd.TextureSamples, validate_with_format int) rd.Framebuffer {
	framebuffer := gpu.class.FramebufferCreateEmpty(Vector2i(size), classdb.RenderingDeviceTextureSamples(samples), Int(validate_with_format))
	return mmm.New[rd.Framebuffer, rd.Interface, uintptr](span, gpu.value, uintptr(framebuffer))
}

func (gpu renderingDevice) FramebufferCreateMultipass(span mmm.Lifetime, textures ffi.Managed[[]rd.Texture], passes ffi.Managed[[]rd.FramebufferPass], validate_with_format, view_count int) rd.Framebuffer {
	framebuffer := gpu.class.FramebufferCreateMultipass(gpu.textures(textures), gpu.framebufferPasses(passes), Int(validate_with_format), Int(view_count))
	return mmm.New[rd.Framebuffer, rd.Interface, uintptr](span, gpu.value, uintptr(framebuffer))
}

func (gpu renderingDevice) FramebufferGetFormat(fb rd.Framebuffer) rd.FramebufferFormat {
	return rd.FramebufferFormat(gpu.class.FramebufferGetFormat(RID(mmm.Get(fb))))
}

func (gpu renderingDevice) FramebufferIsValid(fb rd.Framebuffer) bool {
	return bool(gpu.class.FramebufferIsValid(RID(mmm.Get(fb))))
}

func (gpu renderingDevice) FramebufferFormatCreate(attachments ffi.Managed[[]rd.AttachmentFormat], view_count int) rd.FramebufferFormat {
	return rd.FramebufferFormat(gpu.class.FramebufferFormatCreate(gpu.attachments(attachments), Int(view_count)))
}

func (gpu renderingDevice) FramebufferFormatCreateEmpty(samples rd.TextureSamples) rd.FramebufferFormat {
	return rd.FramebufferFormat(gpu.class.FramebufferFormatCreateEmpty(classdb.RenderingDeviceTextureSamples(samples)))
}

func (gpu renderingDevice) FramebufferFormatCreateMultipass(attachments ffi.Managed[[]rd.AttachmentFormat], passes ffi.Managed[[]rd.FramebufferPass], view_count int) rd.FramebufferFormat {
	return rd.FramebufferFormat(gpu.class.FramebufferFormatCreateMultipass(gpu.attachments(attachments), gpu.framebufferPasses(passes), Int(view_count)))
}

func (gpu renderingDevice) FramebufferFormatGetTextureSamples(format1 rd.FramebufferFormat, render_pass int) rd.TextureSamples {
	return rd.TextureSamples(gpu.class.FramebufferFormatGetTextureSamples(Int(format1), Int(render_pass)))
}

func (gpu renderingDevice) FullBarrier() {
	gpu.class.FullBarrier()
}

func (gpu renderingDevice) GetDeviceName() string {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return gpu.class.GetDeviceName(tmp).String()
}

func (gpu renderingDevice) GetDevicePipelineCacheUUID() string {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return gpu.class.GetDevicePipelineCacheUuid(tmp).String()
}

func (gpu renderingDevice) GetDeviceVendorName() string {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return gpu.class.GetDeviceVendorName(tmp).String()
}

func (gpu renderingDevice) GetDriverResource(resource rd.DriverResource, rid rd.RID, index int) int {
	return int(gpu.class.GetDriverResource(classdb.RenderingDeviceDriverResource(resource), RID(rid), Int(index)))
}

func (gpu renderingDevice) GetFrameDelay() int {
	return int(gpu.class.GetFrameDelay())
}

func (gpu renderingDevice) GetMemoryUsage(mtype rd.MemoryType) int {
	return int(gpu.class.GetMemoryUsage(classdb.RenderingDeviceMemoryType(mtype)))
}

func (gpu renderingDevice) IndexArrayCreate(span mmm.Lifetime, buf rd.IndexBuffer, offset, count int) rd.IndexArray {
	array := gpu.class.IndexArrayCreate(RID(mmm.Get(buf)), Int(offset), Int(count))
	return mmm.New[rd.IndexArray, rd.Interface, uintptr](span, gpu.value, uintptr(array))
}

func (gpu renderingDevice) IndexBufferCreate(span mmm.Lifetime, size_indices int, format rd.IndexBufferFormat, data ffi.Bytes, use_restart_indicies bool) rd.IndexBuffer {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	buffer := gpu.class.IndexBufferCreate(Int(size_indices), classdb.RenderingDeviceIndexBufferFormat(format), gpu.bytes(tmp, data), Bool(use_restart_indicies))
	return mmm.New[rd.IndexBuffer, rd.Interface, uintptr](span, gpu.value, uintptr(buffer))
}

func (gpu renderingDevice) LimitGet(limit rd.Limit) int {
	return int(gpu.class.LimitGet(classdb.RenderingDeviceLimit(limit)))
}

func (gpu renderingDevice) RenderPipelineCreate(span mmm.Lifetime, shader rd.Shader, framebuffer_format rd.FramebufferFormat, vertex_format rd.VertexFormat,
	primitive rd.RenderPrimitive, rasterization_state ffi.Managed[rd.PipelineRasterizationState], multisample_state ffi.Managed[rd.PipelineMultisampleState],
	stencil_state ffi.Managed[rd.PipelineDepthStencilState], color_blend_state ffi.Managed[rd.PipelineColorBlendState], dynamic_state_flags rd.PipelineDynamicStateFlags,
	for_render_pass int, specialization_constants ffi.Managed[[]rd.PipelineSpecializationConstant],
) rd.RenderPipeline {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	sc, ok := specialization_constants.Interface().(gd.ArrayOf[classdb.RDPipelineSpecializationConstant])
	if !ok {
		span := gpu.scope(specialization_constants.Lifetime())
		sc = NewArrayOf[classdb.RDPipelineSpecializationConstant](span)
		for _, constant := range specialization_constants.Value() {
			var converted = *Create(span, new(classdb.RDPipelineSpecializationConstant))
			converted.SetConstantId(Int(constant.ConstantID))
			converted.SetValue(span.Variant(constant.Value))
			sc.Append(converted)
		}
		specialization_constants.Cache(sc)
	}
	rs, ok := rasterization_state.Interface().(classdb.RDPipelineRasterizationState)
	if !ok {
		value := rasterization_state.Value()
		rs = *Create(gpu.scope(rasterization_state.Lifetime()), new(classdb.RDPipelineRasterizationState))
		rs.SetCullMode(classdb.RenderingDevicePolygonCullMode(value.CullMode))
		rs.SetDepthBiasClamp(Float(value.DepthBiasClamp))
		rs.SetDepthBiasConstantFactor(Float(value.DepthBiasConstantFactor))
		rs.SetDepthBiasEnabled(Bool(value.DepthBiasEnabled))
		rs.SetDepthBiasSlopeFactor(Float(value.DepthBiasSlopeFactor))
		rs.SetDiscardPrimitives(Bool(value.DiscardPrimitives))
		rs.SetEnableDepthClamp(Bool(value.EnableDepthClamp))
		rs.SetFrontFace(classdb.RenderingDevicePolygonFrontFace(value.FrontFace))
		rs.SetLineWidth(Float(value.LineWidth))
		rs.SetPatchControlPoints(Int(value.PatchControlPoints))
		rs.SetWireframe(Bool(value.Wireframe))
		rasterization_state.Cache(rs)
	}
	ms, ok := multisample_state.Interface().(classdb.RDPipelineMultisampleState)
	if !ok {
		value := multisample_state.Value()
		ms = *Create(gpu.scope(multisample_state.Lifetime()), new(classdb.RDPipelineMultisampleState))
		ms.SetEnableAlphaToCoverage(Bool(value.EnableAlphaToCoverage))
		ms.SetEnableAlphaToOne(Bool(value.EnableAlphaToOne))
		ms.SetEnableSampleShading(Bool(value.EnableSampleShading))
		ms.SetMinSampleShading(Float(value.MinSampleShading))
		ms.SetSampleCount(classdb.RenderingDeviceTextureSamples(value.SampleCount))
		masks, ok := value.SampleMasks.Interface().(gd.ArrayOf[gd.Int])
		if !ok {
			masks = NewArrayOf[gd.Int](gpu.scope(multisample_state.Lifetime()))
			masks.Resize(Int(len(value.SampleMasks.Value())))
			for i, mask := range value.SampleMasks.Value() {
				masks.SetIndex(Int(i), Int(mask))
			}
			value.SampleMasks.Cache(masks)
		}
		ms.SetSampleMasks(masks)
		multisample_state.Cache(ms)
	}
	ss, ok := stencil_state.Interface().(classdb.RDPipelineDepthStencilState)
	if !ok {
		value := stencil_state.Value()
		ss = *Create(gpu.scope(stencil_state.Lifetime()), new(classdb.RDPipelineDepthStencilState))
		ss.SetBackOpCompare(classdb.RenderingDeviceCompareOperator(value.BackOperationCompare))
		ss.SetBackOpCompareMask(Int(value.BackOperationCompareMask))
		ss.SetBackOpDepthFail(classdb.RenderingDeviceStencilOperation(value.BackOperationDepthFail))
		ss.SetBackOpFail(classdb.RenderingDeviceStencilOperation(value.BackOperationFail))
		ss.SetBackOpPass(classdb.RenderingDeviceStencilOperation(value.BackOperationPass))
		ss.SetBackOpReference(Int(value.BackOperationReference))
		ss.SetBackOpWriteMask(Int(value.BackOperationWriteMask))
		ss.SetDepthCompareOperator(classdb.RenderingDeviceCompareOperator(value.DepthCompareOperator))
		ss.SetDepthRangeMax(Float(value.DepthRangeMax))
		ss.SetDepthRangeMin(Float(value.DepthRangeMin))
		ss.SetEnableDepthRange(Bool(value.EnableDepthRange))
		ss.SetEnableDepthTest(Bool(value.EnableDepthTest))
		ss.SetEnableDepthWrite(Bool(value.EnableDepthWrite))
		ss.SetEnableStencil(Bool(value.EnableStencil))
		ss.SetFrontOpCompare(classdb.RenderingDeviceCompareOperator(value.FrontOperationCompare))
		ss.SetFrontOpCompareMask(Int(value.FrontOperationCompareMask))
		ss.SetFrontOpDepthFail(classdb.RenderingDeviceStencilOperation(value.FrontOperationDepthFail))
		ss.SetFrontOpFail(classdb.RenderingDeviceStencilOperation(value.FrontOperationFail))
		ss.SetFrontOpPass(classdb.RenderingDeviceStencilOperation(value.FrontOperationPass))
		ss.SetFrontOpReference(Int(value.FrontOperationReference))
		ss.SetFrontOpWriteMask(Int(value.FrontOperationWriteMask))
		stencil_state.Cache(ss)
	}
	cb, ok := color_blend_state.Interface().(classdb.RDPipelineColorBlendState)
	if !ok {
		value := color_blend_state.Value()
		cb = *Create(gpu.scope(color_blend_state.Lifetime()), new(classdb.RDPipelineColorBlendState))
		attachments, ok := value.Attachments.Interface().(gd.ArrayOf[classdb.RDPipelineColorBlendStateAttachment])
		if !ok {
			attachments = NewArrayOf[classdb.RDPipelineColorBlendStateAttachment](gpu.scope(color_blend_state.Lifetime()))
			for _, attachment := range value.Attachments.Value() {
				var converted = *Create(gpu.scope(color_blend_state.Lifetime()), new(classdb.RDPipelineColorBlendStateAttachment))
				converted.SetAlphaBlendOp(classdb.RenderingDeviceBlendOperation(attachment.AlphaBlendOperation))
				converted.SetColorBlendOp(classdb.RenderingDeviceBlendOperation(attachment.ColorBlendOperation))
				converted.SetDstAlphaBlendFactor(classdb.RenderingDeviceBlendFactor(attachment.DestinationAlphaBlendFactor))
				converted.SetDstColorBlendFactor(classdb.RenderingDeviceBlendFactor(attachment.DestinationColorBlendFactor))
				converted.SetEnableBlend(Bool(attachment.EnableBlend))
				converted.SetSrcAlphaBlendFactor(classdb.RenderingDeviceBlendFactor(attachment.SourceAlphaBlendFactor))
				converted.SetSrcColorBlendFactor(classdb.RenderingDeviceBlendFactor(attachment.SourceColorBlendFactor))
				converted.SetWriteA(Bool(attachment.WriteA))
				converted.SetWriteB(Bool(attachment.WriteB))
				converted.SetWriteG(Bool(attachment.WriteG))
				converted.SetWriteR(Bool(attachment.WriteR))
				attachments.Append(converted)
			}
			value.Attachments.Cache(attachments)
		}
		cb.SetAttachments(attachments)
		cb.SetBlendConstant(value.BlendConstant)
		cb.SetEnableLogicOp(Bool(value.EnableLogicOperation))
		cb.SetLogicOp(classdb.RenderingDeviceLogicOperation(value.LogicOperation))
		color_blend_state.Cache(cb)
	}
	pipeline := gpu.class.RenderPipelineCreate(RID(mmm.Get(shader)), Int(framebuffer_format), Int(vertex_format),
		classdb.RenderingDeviceRenderPrimitive(primitive), rs, ms, ss, cb, classdb.RenderingDevicePipelineDynamicStateFlags(dynamic_state_flags), Int(for_render_pass), sc,
	)
	return mmm.New[rd.RenderPipeline, rd.Interface, uintptr](span, gpu.value, uintptr(pipeline))
}

func (gpu renderingDevice) RenderPipelineIsValid(pipeline rd.RenderPipeline) bool {
	return bool(gpu.class.RenderPipelineIsValid(RID(mmm.Get(pipeline))))
}

func (gpu renderingDevice) SamplerCreate(span mmm.Lifetime, state ffi.Managed[rd.SamplerState]) rd.Sampler {
	s, ok := state.Interface().(classdb.RDSamplerState)
	if !ok {
		value := state.Value()
		s = *Create(gpu.scope(state.Lifetime()), new(classdb.RDSamplerState))
		s.SetAnisotropyMax(Float(value.AnisotropyMax))
		s.SetBorderColor(classdb.RenderingDeviceSamplerBorderColor(value.BorderColor))
		s.SetCompareOp(classdb.RenderingDeviceCompareOperator(value.CompareOperator))
		s.SetEnableCompare(Bool(value.EnableCompare))
		s.SetLodBias(Float(value.LevelOfDetailBias))
		s.SetMagFilter(classdb.RenderingDeviceSamplerFilter(value.MagnificationFilter))
		s.SetMaxLod(Float(value.MaxLevelOfDetail))
		s.SetMinFilter(classdb.RenderingDeviceSamplerFilter(value.MinificationFilter))
		s.SetMinLod(Float(value.MinLevelOfDetail))
		s.SetMipFilter(classdb.RenderingDeviceSamplerFilter(value.MipmapFilter))
		s.SetRepeatU(classdb.RenderingDeviceSamplerRepeatMode(value.RepeatU))
		s.SetRepeatV(classdb.RenderingDeviceSamplerRepeatMode(value.RepeatV))
		s.SetRepeatW(classdb.RenderingDeviceSamplerRepeatMode(value.RepeatW))
		s.SetUnnormalizedUvw(Bool(value.UnnormalizedUVW))
		s.SetUseAnisotropy(Bool(value.UseAnisotropy))
		state.Cache(s)
	}
	sampler := gpu.class.SamplerCreate(s)
	return mmm.New[rd.Sampler, rd.Interface, uintptr](span, gpu.value, uintptr(sampler))
}

func (gpu renderingDevice) SamplerIsFormatSupportedForFilter(format rd.DataFormat, filter rd.SamplerFilter) bool {
	return bool(gpu.class.SamplerIsFormatSupportedForFilter(classdb.RenderingDeviceDataFormat(format), classdb.RenderingDeviceSamplerFilter(filter)))
}

func (gpu renderingDevice) ScreenGetFramebufferFormat() rd.FramebufferFormat {
	return rd.FramebufferFormat(gpu.class.ScreenGetFramebufferFormat())
}

func (gpu renderingDevice) ScreenGetHeight(screen rd.Screen) int64 {
	return int64(gpu.class.ScreenGetHeight(Int(screen)))
}

func (gpu renderingDevice) ScreenGetWidth(screen rd.Screen) int64 {
	return int64(gpu.class.ScreenGetWidth(Int(screen)))
}

func (gpu renderingDevice) SetResourceName(resource rd.RID, name string) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.SetResourceName(RID(resource), tmp.String(name))
}

func (gpu renderingDevice) ShaderCompileBinaryFromSPIRV(span mmm.Lifetime, shader ffi.Managed[rd.ShaderSPIRV], name string) ffi.Bytes {
	spirv, ok := shader.Interface().(classdb.RDShaderSPIRV)
	if !ok {
		value := shader.Value()
		span := gpu.scope(shader.Lifetime())
		spirv = *Create(gpu.scope(shader.Lifetime()), new(classdb.RDShaderSPIRV))
		spirv.SetStageBytecode(RenderingDeviceShaderStageCompute, gpu.bytes(span, value.Compute))
		spirv.SetStageBytecode(RenderingDeviceShaderStageFragment, gpu.bytes(span, value.Fragment))
		spirv.SetStageBytecode(RenderingDeviceShaderStageTesselationControl, gpu.bytes(span, value.TesselationControl))
		spirv.SetStageBytecode(RenderingDeviceShaderStageTesselationEvaluation, gpu.bytes(span, value.TesselationEvaluation))
		spirv.SetStageBytecode(RenderingDeviceShaderStageVertex, gpu.bytes(span, value.Vertex))
		shader.Cache(spirv)
	}
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return ffi.Bytes{Interface: gpu.class.ShaderCompileBinaryFromSpirv(gpu.scope(span), spirv, tmp.String(name))}
}

func (gpu renderingDevice) ShaderCompileSourceIntoSPIRV(span mmm.Lifetime, shader ffi.Managed[rd.ShaderSource], allow_cache bool) rd.ShaderSPIRV {
	src, ok := shader.Interface().(classdb.RDShaderSource)
	if !ok {
		value := shader.Value()
		span := gpu.scope(shader.Lifetime())
		src = *Create(gpu.scope(shader.Lifetime()), new(classdb.RDShaderSource))
		src.SetStageSource(RenderingDeviceShaderStageCompute, gpu.string(span, value.Compute))
		src.SetStageSource(RenderingDeviceShaderStageFragment, gpu.string(span, value.Fragment))
		src.SetStageSource(RenderingDeviceShaderStageTesselationControl, gpu.string(span, value.TesselationControl))
		src.SetStageSource(RenderingDeviceShaderStageTesselationEvaluation, gpu.string(span, value.TesselationEvaluation))
		src.SetStageSource(RenderingDeviceShaderStageVertex, gpu.string(span, value.Vertex))
		shader.Cache(src)
	}
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	var spirv = gpu.class.ShaderCompileSpirvFromSource(gpu.scope(span), src, allow_cache)
	return rd.ShaderSPIRV{
		Compute:               ffi.Bytes{Interface: spirv.GetStageBytecode(gpu.scope(span), RenderingDeviceShaderStageCompute)},
		Fragment:              ffi.Bytes{Interface: spirv.GetStageBytecode(gpu.scope(span), RenderingDeviceShaderStageFragment)},
		TesselationControl:    ffi.Bytes{Interface: spirv.GetStageBytecode(gpu.scope(span), RenderingDeviceShaderStageTesselationControl)},
		TesselationEvaluation: ffi.Bytes{Interface: spirv.GetStageBytecode(gpu.scope(span), RenderingDeviceShaderStageTesselationEvaluation)},
		Vertex:                ffi.Bytes{Interface: spirv.GetStageBytecode(gpu.scope(span), RenderingDeviceShaderStageVertex)},
	}
}

func (gpu renderingDevice) ShaderCreateFromBytecode(span mmm.Lifetime, code ffi.Bytes, id rd.ShaderPlaceholder) rd.Shader {
	shader := gpu.class.ShaderCreateFromBytecode(gpu.bytes(gpu.scope(span), code), RID(mmm.Get(id)))
	return mmm.New[rd.Shader, rd.Interface, uintptr](span, gpu.value, uintptr(shader))
}

func (gpu renderingDevice) ShaderCreateFromSPIRV(span mmm.Lifetime, shader ffi.Managed[rd.ShaderSPIRV], name string) rd.Shader {
	spirv, ok := shader.Interface().(classdb.RDShaderSPIRV)
	if !ok {
		value := shader.Value()
		span := gpu.scope(shader.Lifetime())
		spirv = *Create(gpu.scope(shader.Lifetime()), new(classdb.RDShaderSPIRV))
		spirv.SetStageBytecode(RenderingDeviceShaderStageCompute, gpu.bytes(span, value.Compute))
		spirv.SetStageBytecode(RenderingDeviceShaderStageFragment, gpu.bytes(span, value.Fragment))
		spirv.SetStageBytecode(RenderingDeviceShaderStageTesselationControl, gpu.bytes(span, value.TesselationControl))
		spirv.SetStageBytecode(RenderingDeviceShaderStageTesselationEvaluation, gpu.bytes(span, value.TesselationEvaluation))
		spirv.SetStageBytecode(RenderingDeviceShaderStageVertex, gpu.bytes(span, value.Vertex))
		shader.Cache(spirv)
	}
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	result := gpu.class.ShaderCreateFromSpirv(spirv, tmp.String(name))
	return mmm.New[rd.Shader, rd.Interface, uintptr](span, gpu.value, uintptr(result))
}

func (gpu renderingDevice) ShaderCreatePlaceholder(span mmm.Lifetime) rd.ShaderPlaceholder {
	placeholder := gpu.class.ShaderCreatePlaceholder()
	return mmm.New[rd.ShaderPlaceholder, rd.Interface, uintptr](span, gpu.value, uintptr(placeholder))
}

func (gpu renderingDevice) ShaderGetVertexInputAttributeMask(shader rd.Shader) int64 {
	return int64(gpu.class.ShaderGetVertexInputAttributeMask(RID(mmm.Get(shader))))
}

func (gpu renderingDevice) StorageBufferCreate(span mmm.Lifetime, size_bytes int, data ffi.Bytes, usage_flags rd.StorageBufferUsage) rd.StorageBuffer {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	buffer := gpu.class.StorageBufferCreate(Int(size_bytes), gpu.bytes(tmp, data), classdb.RenderingDeviceStorageBufferUsage(usage_flags))
	return mmm.New[rd.StorageBuffer, rd.Interface, uintptr](span, gpu.value, uintptr(buffer))
}

func (gpu renderingDevice) Submit() {
	gpu.class.Submit()
}

func (gpu renderingDevice) Sync() {
	gpu.class.Sync()
}

func (gpu renderingDevice) TextureBufferCreate(span mmm.Lifetime, size_bytes int, data_format rd.DataFormat, data ffi.Bytes) rd.TextureBuffer {
	buffer := gpu.class.TextureBufferCreate(Int(size_bytes), classdb.RenderingDeviceDataFormat(data_format), gpu.bytes(gpu.scope(span), data))
	return mmm.New[rd.TextureBuffer, rd.Interface, uintptr](span, gpu.value, uintptr(buffer))
}

func (gpu renderingDevice) TextureClear(texture rd.Texture, color uc.Color, base_mipmap, mipmap_count, base_layer, layer_count int, post_barrier rd.BarrierMask) error {
	return Error(gpu.class.TextureClear(RID(mmm.Get(texture)), color, Int(base_mipmap), Int(mipmap_count), Int(base_layer), Int(layer_count), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu renderingDevice) TextureCopy(src, dst rd.Texture, src_pos, dst_pos, size xy.Vector3, src_mipmap, dst_mipmap, src_layer, dst_layer int, post_barrier rd.BarrierMask) error {
	return Error(gpu.class.TextureCopy(RID(mmm.Get(src)), RID(mmm.Get(dst)), Vector3(src_pos), Vector3(dst_pos), Vector3(size), Int(src_mipmap), Int(dst_mipmap), Int(src_layer), Int(dst_layer), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu renderingDevice) TextureCreate(span mmm.Lifetime, format ffi.Managed[rd.TextureFormat], view ffi.Managed[rd.TextureView], data ffi.Managed[[]ffi.Bytes]) rd.Texture {
	fmt, ok := format.Interface().(classdb.RDTextureFormat)
	if !ok {
		value := format.Value()
		fmt = *Create(gpu.scope(format.Lifetime()), new(classdb.RDTextureFormat))
		fmt.SetArrayLayers(Int(value.ArrayLayers))
		fmt.SetDepth(Int(value.Depth))
		fmt.SetFormat(classdb.RenderingDeviceDataFormat(value.Format))
		fmt.SetHeight(Int(value.Height))
		fmt.SetMipmaps(Int(value.Mipmaps))
		fmt.SetSamples(classdb.RenderingDeviceTextureSamples(value.Samples))
		fmt.SetTextureType(classdb.RenderingDeviceTextureType(value.TextureType))
		fmt.SetUsageBits(classdb.RenderingDeviceTextureUsageBits(value.Usage))
		fmt.SetWidth(Int(value.Width))
		format.Cache(fmt)
	}
	vue, ok := view.Interface().(classdb.RDTextureView)
	if !ok {
		value := view.Value()
		vue = *Create(gpu.scope(view.Lifetime()), new(classdb.RDTextureView))
		vue.SetFormatOverride(classdb.RenderingDeviceDataFormat(value.FormatOverride))
		vue.SetSwizzleA(classdb.RenderingDeviceTextureSwizzle(value.SwizzleAlpha))
		vue.SetSwizzleB(classdb.RenderingDeviceTextureSwizzle(value.SwizzleBlue))
		vue.SetSwizzleG(classdb.RenderingDeviceTextureSwizzle(value.SwizzleGreen))
		vue.SetSwizzleR(classdb.RenderingDeviceTextureSwizzle(value.SwizzleRed))
		view.Cache(vue)
	}
	buf, ok := data.Interface().(gd.ArrayOf[gd.PackedByteArray])
	if !ok {
		buf = NewArrayOf[gd.PackedByteArray](gpu.scope(data.Lifetime()))
		for _, bytes := range data.Value() {
			buf.Append(gpu.bytes(gpu.scope(data.Lifetime()), bytes))
		}
		data.Cache(buf)
	}
	texture := gpu.class.TextureCreate(fmt, vue, buf)
	return mmm.New[rd.Texture, rd.Interface, uintptr](span, gpu.value, uintptr(texture))
}

func (gpu renderingDevice) TextureCreateFromExtension(span mmm.Lifetime,
	texture_type rd.TextureType, format rd.DataFormat, samples rd.TextureSamples, usage_flags rd.TextureUsage,
	image, width, height, depth, layers int,
) rd.Texture {
	texture := gpu.class.TextureCreateFromExtension(
		classdb.RenderingDeviceTextureType(texture_type),
		classdb.RenderingDeviceDataFormat(format),
		classdb.RenderingDeviceTextureSamples(samples),
		classdb.RenderingDeviceTextureUsageBits(usage_flags),
		Int(image), Int(width), Int(height), Int(depth), Int(layers),
	)
	return mmm.New[rd.Texture, rd.Interface, uintptr](span, gpu.value, uintptr(texture))
}

func (gpu renderingDevice) TextureCreateShared(span mmm.Lifetime, view ffi.Managed[rd.TextureView], with_texture rd.Texture) rd.Texture {
	vue, ok := view.Interface().(classdb.RDTextureView)
	if !ok {
		value := view.Value()
		vue = *Create(gpu.scope(view.Lifetime()), new(classdb.RDTextureView))
		vue.SetFormatOverride(classdb.RenderingDeviceDataFormat(value.FormatOverride))
		vue.SetSwizzleA(classdb.RenderingDeviceTextureSwizzle(value.SwizzleAlpha))
		vue.SetSwizzleB(classdb.RenderingDeviceTextureSwizzle(value.SwizzleBlue))
		vue.SetSwizzleG(classdb.RenderingDeviceTextureSwizzle(value.SwizzleGreen))
		vue.SetSwizzleR(classdb.RenderingDeviceTextureSwizzle(value.SwizzleRed))
		view.Cache(vue)
	}
	texture := gpu.class.TextureCreateShared(vue, RID(mmm.Get(with_texture)))
	return mmm.New[rd.Texture, rd.Interface, uintptr](span, gpu.value, uintptr(texture))
}

func (gpu renderingDevice) TextureCreateSharedFromSlice(span mmm.Lifetime, view ffi.Managed[rd.TextureView], with_texture rd.Texture, layer, mipmap, mipmaps int, slice_type rd.TextureSliceType) rd.Texture {
	vue, ok := view.Interface().(classdb.RDTextureView)
	if !ok {
		value := view.Value()
		vue = *Create(gpu.scope(view.Lifetime()), new(classdb.RDTextureView))
		vue.SetFormatOverride(classdb.RenderingDeviceDataFormat(value.FormatOverride))
		vue.SetSwizzleA(classdb.RenderingDeviceTextureSwizzle(value.SwizzleAlpha))
		vue.SetSwizzleB(classdb.RenderingDeviceTextureSwizzle(value.SwizzleBlue))
		vue.SetSwizzleG(classdb.RenderingDeviceTextureSwizzle(value.SwizzleGreen))
		vue.SetSwizzleR(classdb.RenderingDeviceTextureSwizzle(value.SwizzleRed))
		view.Cache(vue)
	}
	texture := gpu.class.TextureCreateSharedFromSlice(vue, RID(mmm.Get(with_texture)), Int(layer), Int(mipmap), Int(mipmaps), classdb.RenderingDeviceTextureSliceType(slice_type))
	return mmm.New[rd.Texture, rd.Interface, uintptr](span, gpu.value, uintptr(texture))
}

func (gpu renderingDevice) TextureGetData(span mmm.Lifetime, texture rd.Texture, layer int) ffi.Bytes {
	return ffi.Bytes{
		Interface: gpu.class.TextureGetData(gpu.scope(span), RID(mmm.Get(texture)), Int(layer)),
	}
}

func (gpu renderingDevice) TextureGetFormat(texture rd.Texture) rd.TextureFormat {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	format := gpu.class.TextureGetFormat(tmp, RID(mmm.Get(texture)))
	return rd.TextureFormat{
		ArrayLayers: int(format.GetArrayLayers()),
		Depth:       int(format.GetDepth()),
		Format:      rd.DataFormat(format.GetFormat()),
		Height:      int(format.GetHeight()),
		Mipmaps:     int(format.GetMipmaps()),
		Samples:     rd.TextureSamples(format.GetSamples()),
		TextureType: rd.TextureType(format.GetTextureType()),
		Usage:       rd.TextureUsage(format.GetUsageBits()),
		Width:       int(format.GetWidth()),
	}
}

func (gpu renderingDevice) TextureGetNativeHandle(texture rd.Texture) uint64 {
	return uint64(gpu.class.TextureGetNativeHandle(RID(mmm.Get(texture))))
}

func (gpu renderingDevice) TextureIsFormatSupportedForUsage(format rd.DataFormat, usage rd.TextureUsage) bool {
	return bool(gpu.class.TextureIsFormatSupportedForUsage(classdb.RenderingDeviceDataFormat(format), classdb.RenderingDeviceTextureUsageBits(usage)))
}

func (gpu renderingDevice) TextureIsShared(tex rd.Texture) bool {
	return bool(gpu.class.TextureIsShared(RID(mmm.Get(tex))))
}

func (gpu renderingDevice) TextureIsValid(tex rd.Texture) bool {
	return bool(gpu.class.TextureIsValid(RID(mmm.Get(tex))))
}

func (gpu renderingDevice) TextureResolveMultisample(src, dst rd.Texture, post_barrier rd.BarrierMask) error {
	return Error(gpu.class.TextureResolveMultisample(RID(mmm.Get(src)), RID(mmm.Get(dst)), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu renderingDevice) TextureUpdate(texture rd.Texture, layer int, data ffi.Bytes, post_barrier rd.BarrierMask) error {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return Error(gpu.class.TextureUpdate(RID(mmm.Get(texture)), Int(layer), gpu.bytes(tmp, data), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu renderingDevice) UniformBufferCreate(span mmm.Lifetime, size_bytes int, data ffi.Bytes) rd.UniformBuffer {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	buffer := gpu.class.UniformBufferCreate(Int(size_bytes), gpu.bytes(tmp, data))
	return mmm.New[rd.UniformBuffer, rd.Interface, uintptr](span, gpu.value, uintptr(buffer))
}

func (gpu renderingDevice) UniformSetCreate(span mmm.Lifetime, uniforms ffi.Managed[[]rd.Uniform], shader rd.Shader, shader_set int) rd.UniformSet {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	var uniforms2 = NewArrayOf[classdb.RDUniform](gpu.scope(span))
	for _, uniform := range uniforms.Value() {
		var converted = *Create(gpu.scope(uniforms.Lifetime()), new(classdb.RDUniform))
		converted.SetBinding(Int(uniform.Binding))
		converted.SetUniformType(classdb.RenderingDeviceUniformType(uniform.UniformType))
		uniforms2.Append(converted)
	}
	uniforms.Cache(uniforms2)
	set := gpu.class.UniformSetCreate(uniforms2, RID(mmm.Get(shader)), Int(shader_set))
	return mmm.New[rd.UniformSet, rd.Interface, uintptr](span, gpu.value, uintptr(set))
}

func (gpu renderingDevice) UniformSetIsValid(set rd.UniformSet) bool {
	return bool(gpu.class.UniformSetIsValid(RID(mmm.Get(set))))
}

func (gpu renderingDevice) VertexArrayCreate(span mmm.Lifetime, vertex_count int, vertex_format rd.VertexFormat, src_buffers ffi.Managed[[]rd.Buffer], offsets ffi.Slice[int64]) rd.VertexArray {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	buffers, ok := src_buffers.Interface().(gd.ArrayOf[gd.RID])
	if !ok {
		var buffers2 = NewArrayOf[gd.RID](gpu.scope(src_buffers.Lifetime()))
		for _, buffer := range src_buffers.Value() {
			buffers2.Append(RID(mmm.Get(buffer)))
		}
		src_buffers.Cache(buffers2)
	}
	array := gpu.class.VertexArrayCreate(Int(vertex_count), Int(vertex_format), buffers, gpu.int64s(tmp, offsets))
	return mmm.New[rd.VertexArray, rd.Interface, uintptr](span, gpu.value, uintptr(array))
}

func (gpu renderingDevice) VertexFormatCreate(span mmm.Lifetime, vertex_descriptions ffi.Managed[[]rd.VertexAttribute]) rd.VertexFormat {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	descriptions, ok := vertex_descriptions.Interface().(gd.ArrayOf[classdb.RDVertexAttribute])
	if !ok {
		descriptions = NewArrayOf[classdb.RDVertexAttribute](gpu.scope(span))
		for _, desc := range vertex_descriptions.Value() {
			var converted = *Create(gpu.scope(vertex_descriptions.Lifetime()), new(classdb.RDVertexAttribute))
			converted.SetFormat(classdb.RenderingDeviceDataFormat(desc.Format))
			converted.SetFrequency(classdb.RenderingDeviceVertexFrequency(desc.Frequency))
			converted.SetLocation(Int(desc.Location))
			converted.SetOffset(Int(desc.Offset))
			converted.SetStride(Int(desc.Stride))
			descriptions.Append(converted)
		}
		vertex_descriptions.Cache(descriptions)
	}
	return rd.VertexFormat(gpu.class.VertexFormatCreate(descriptions))

}

func (gpu renderingDevice) FreeRID(rid uintptr) {
	gpu.class.FreeRid(RID(rid))
}
