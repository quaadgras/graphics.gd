package gd

import (
	"io"
	"reflect"
	"time"
	"unsafe"

	gd "grow.graphics/gd/internal"
	classdb "grow.graphics/gd/internal/classdb"
	"grow.graphics/rd"
	"grow.graphics/uc"
	"grow.graphics/xy"
	"runtime.link/mmm"
)

// RenderingDevice returns a new [rd.Interface] implementation via Godot.
func RenderingDevice(godot Lifetime) rd.Interface {
	return &rdInterface{
		godot: godot,
		class: RenderingServer(godot).GetRenderingDevice(godot),
	}
}

type rdInterface struct {
	godot Lifetime
	class classdb.RenderingDevice

	compute rdCompute
	drawing rdDrawing
}

func (gpu *rdInterface) resource(id RID) *rdResource {
	return &rdResource{from: gpu, name: id} // TODO pool
}

func (gpu *rdInterface) Barrier(from, to rd.Barrier) {
	gpu.class.Barrier(classdb.RenderingDeviceBarrierMask(from), classdb.RenderingDeviceBarrierMask(to))
}

func (gpu *rdInterface) BarrierFull() { gpu.class.FullBarrier() }

func (gpu *rdInterface) CaptureTimestamp(name string) rd.Timestamp {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	gpu.class.CaptureTimestamp(tmp.String(name))
	return &rdTimestamp{
		from: gpu,
		name: name,
	}
}

type rdTimestamp struct {
	rd.Timestamp

	from *rdInterface
	name string
}

func (t rdTimestamp) CPU() (time.Duration, bool) {
	tmp := gd.NewContext(t.from.godot.API)
	defer tmp.End()
	for i := range t.from.class.GetCapturedTimestampsCount() {
		if t.from.class.GetCapturedTimestampName(tmp, i).String() == t.name {
			return time.Duration(t.from.class.GetCapturedTimestampCpuTime(i)) * time.Microsecond, true
		}
	}
	return 0, false
}

func (t rdTimestamp) GPU() (time.Duration, bool) {
	tmp := gd.NewContext(t.from.godot.API)
	defer tmp.End()
	for i := range t.from.class.GetCapturedTimestampsCount() {
		if t.from.class.GetCapturedTimestampName(tmp, i).String() == t.name {
			return time.Duration(t.from.class.GetCapturedTimestampGpuTime(i)) * time.Microsecond, true
		}
	}
	return 0, false
}

func (gpu *rdInterface) CompileBinary(data []byte) rd.Shader {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	return rdShader{rdNameable{gpu.resource(gpu.class.ShaderCreateFromBytecode(tmp.PackedByteSlice(data), 0))}}
}

type rdResource struct {
	rd.Resource

	from *rdInterface
	name RID
}

func (res *rdResource) rid() RID {
	if res.name == 0 {
		panic("use after free / nil resource dereference")
	}
	return res.name
}

func (res *rdResource) RID() uint64 { return uint64(res.name) }

func (res *rdResource) Free() {
	res.from.class.FreeRid(res.name)
	res.name = 0
}

type rdNameable struct {
	*rdResource
}

func (res rdNameable) SetResourceName(name string) {
	tmp := gd.NewContext(res.from.godot.API)
	defer tmp.End()
	res.from.class.SetResourceName(res.rid(), tmp.String(name))
}

type rdShader struct {
	rdNameable
}

func (gpu rdShader) Compile(data []byte) {
	tmp := gd.NewContext(gpu.from.godot.API)
	defer tmp.End()
	gpu.name = gpu.from.class.ShaderCreateFromBytecode(tmp.PackedByteSlice(data), gpu.rid())
}

func (gpu rdShader) VertexInputAttributeMask() uint32 {
	return uint32(gpu.from.class.ShaderGetVertexInputAttributeMask(gpu.rid()))
}

func (gpu rdShader) Variables(variables map[int]rd.Variable) rd.Variables {
	tmp := gd.NewContext(gpu.from.godot.API)
	defer tmp.End()
	array := NewArrayOf[classdb.RDUniform](tmp)
	for binding, variable := range variables {
		var uniform = *gd.Create(tmp, new(classdb.RDUniform))
		uniform.SetBinding(Int(binding))
		switch variable := variable.(type) {
		case rdSampler:
			uniform.SetUniformType(RenderingDeviceUniformTypeSampler)
			uniform.AddId(gd.RID(variable.RID()))
		case rd.SamplerWithTexture:
			uniform.SetUniformType(RenderingDeviceUniformTypeSampler)
			uniform.AddId(gd.RID(variable.Sampler.RID()))
			uniform.AddId(gd.RID(variable.Texture.RID()))
		case rdTexture:
			uniform.SetUniformType(RenderingDeviceUniformTypeTexture)
			uniform.AddId(gd.RID(variable.RID()))
		case rdTextureBuffer:
			uniform.SetUniformType(RenderingDeviceUniformTypeTextureBuffer)
			uniform.AddId(gd.RID(variable.RID()))
		case rd.SamplerWithTextureBuffer:
			uniform.SetUniformType(RenderingDeviceUniformTypeTextureBuffer)
			uniform.AddId(gd.RID(variable.Sampler.RID()))
		case rdUniformBuffer:
			uniform.SetUniformType(RenderingDeviceUniformTypeUniformBuffer)
			uniform.AddId(gd.RID(variable.RID()))
		case rdStorageBuffer:
			uniform.SetUniformType(RenderingDeviceUniformTypeStorageBuffer)
			uniform.AddId(gd.RID(variable.RID()))
		case rd.InputAttachment:
			uniform.SetUniformType(RenderingDeviceUniformTypeInputAttachment)
			uniform.AddId(gd.RID(variable.Texture.RID()))
		default:
			panic("unexpected rd.Variable type " + reflect.TypeOf(variable).String())
		}
		array.Append(uniform)
	}
	return rdVariables{gpu.from.resource(gpu.from.class.UniformSetCreate(array, gpu.rid(), 0))}
}

type rdVariables struct{ *rdResource }

func (vars rdVariables) AreValid() bool { return vars.from.class.UniformSetIsValid(vars.rid()) }

func (gpu *rdInterface) CompileSPIRV(name string, spirv rd.SPIRV) []byte {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	already, ok := spirv.(rdSPIRV)
	if !ok {
		converted := *gd.Create(tmp, new(classdb.RDShaderSPIRV))
		converted.SetStageBytecode(RenderingDeviceShaderStageCompute, tmp.PackedByteSlice(spirv.Compute()))
		converted.SetStageBytecode(RenderingDeviceShaderStageFragment, tmp.PackedByteSlice(spirv.Fragment()))
		converted.SetStageBytecode(RenderingDeviceShaderStageTesselationControl, tmp.PackedByteSlice(spirv.TesselationControl()))
		converted.SetStageBytecode(RenderingDeviceShaderStageTesselationEvaluation, tmp.PackedByteSlice(spirv.TesselationEvaluation()))
		converted.SetStageBytecode(RenderingDeviceShaderStageVertex, tmp.PackedByteSlice(spirv.Vertex()))
		return gpu.class.ShaderCompileBinaryFromSpirv(tmp, converted, tmp.String(name)).Bytes()
	}
	return gpu.class.ShaderCompileBinaryFromSpirv(tmp, already.RDShaderSPIRV, tmp.String(name)).Bytes()
}

func (gpu *rdInterface) CompileSource(cache bool, source rd.ShaderSource) rd.SPIRV {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	converted := *gd.Create(tmp, new(classdb.RDShaderSource))
	converted.SetLanguage(RenderingDeviceShaderLanguage(source.Language))
	converted.SetStageSource(RenderingDeviceShaderStageCompute, tmp.String(string(source.Compute)))
	converted.SetStageSource(RenderingDeviceShaderStageFragment, tmp.String(string(source.Fragment)))
	converted.SetStageSource(RenderingDeviceShaderStageTesselationControl, tmp.String(string(source.TesselationControl)))
	converted.SetStageSource(RenderingDeviceShaderStageTesselationEvaluation, tmp.String(string(source.TesselationEvaluation)))
	converted.SetStageSource(RenderingDeviceShaderStageVertex, tmp.String(string(source.Vertex)))
	spirv := gpu.class.ShaderCompileSpirvFromSource(tmp, converted, Bool(cache))
	return rdSPIRV{
		RDShaderSPIRV: spirv,
		from:          gpu,
	}
}

type rdSPIRV struct {
	RDShaderSPIRV

	from *rdInterface
}

func (spirv rdSPIRV) Compute() []byte {
	tmp := gd.NewContext(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return spirv.GetStageBytecode(tmp, RenderingDeviceShaderStageCompute).Bytes()
}

func (spirv rdSPIRV) Fragment() []byte {
	tmp := gd.NewContext(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return spirv.GetStageBytecode(tmp, RenderingDeviceShaderStageFragment).Bytes()
}

func (spirv rdSPIRV) TesselationControl() []byte {
	tmp := gd.NewContext(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return spirv.GetStageBytecode(tmp, RenderingDeviceShaderStageTesselationControl).Bytes()
}

func (spirv rdSPIRV) TesselationEvaluation() []byte {
	tmp := gd.NewContext(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return spirv.GetStageBytecode(tmp, RenderingDeviceShaderStageTesselationEvaluation).Bytes()
}

func (spirv rdSPIRV) Vertex() []byte {
	tmp := gd.NewContext(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return spirv.GetStageBytecode(tmp, RenderingDeviceShaderStageVertex).Bytes()
}

func (spirv rdSPIRV) Shader(name string) rd.Shader {
	tmp := gd.NewContext(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return rdShader{rdNameable{spirv.from.resource(spirv.from.class.ShaderCreateFromSpirv(spirv.RDShaderSPIRV, tmp.String(name)))}}
}

func (gpu *rdInterface) Compute(fn func(rd.Compute)) {
	list := gpu.class.ComputeListBegin(Bool(false)) // FIXME should this be configurable?
	gpu.compute.list = list
	gpu.compute.from = gpu
	fn(&gpu.compute)
	gpu.class.ComputeListEnd(RenderingDeviceBarrierMaskAllBarriers) // FIXME should this be configurable?
}

type rdCompute struct {
	from *rdInterface
	list Int
}

func (c *rdCompute) SetData(data []byte) {
	tmp := gd.NewContext(c.from.godot.API)
	defer tmp.End()
	c.from.class.ComputeListSetPushConstant(c.list, tmp.PackedByteSlice(data), Int(len(data))) // FIXME can we avoid the copy here?
}

func (c *rdCompute) SetProcessor(p rd.Processor) {
	c.from.class.ComputeListBindComputePipeline(c.list, p.(*rdProcessor).rid())
}

func (c *rdCompute) SetVariables(level rd.VariableLevel, variables rd.Variables) {
	c.from.class.ComputeListBindUniformSet(c.list, variables.(*rdVariables).rid(), Int(level))
}

func (c *rdCompute) Submit(x, y, z int) {
	c.from.class.ComputeListDispatch(c.list, Int(x), Int(y), Int(z))
}

type rdProcessor struct{ *rdResource }

func (gpu *rdInterface) DeviceName() string {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	return gpu.class.GetDeviceName(tmp).String()
}

func (gpu *rdInterface) DeviceVendor() string {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	return gpu.class.GetDeviceVendorName(tmp).String()
}

func (gpu *rdInterface) Drawing(frame rd.Frame, fn func(rd.Drawing)) {
	intial_color_action, final_color_action := frame.Color()
	initial_depth_action, final_depth_action := frame.Depth()

	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()

	storage := NewArrayOf[RID](tmp)
	storage.Resize(Int(len(frame.Storage)))
	for i, texture := range frame.Storage {
		storage.SetIndex(Int(i), texture.(*rdTexture).rid())
	}

	list := gpu.class.DrawListBegin(
		frame.Buffer.(*rdFramebuffer).rid(),
		classdb.RenderingDeviceInitialAction(intial_color_action),
		classdb.RenderingDeviceFinalAction(final_color_action),
		classdb.RenderingDeviceInitialAction(initial_depth_action),
		classdb.RenderingDeviceFinalAction(final_depth_action),
		gpu.godot.PackedColorSlice(frame.Clear.Colors),
		frame.Clear.Depth,
		Int(frame.Clear.Stencil),
		frame.Region,
		storage,
	)
	gpu.drawing.list = list
	gpu.drawing.from = gpu
	fn(&gpu.drawing)
	gpu.class.DrawListEnd(RenderingDeviceBarrierMaskAllBarriers) // FIXME should this be configurable?
}

type rdFramebuffer struct{ *rdResource }

func (fb *rdFramebuffer) IsValid() bool { return fb.from.class.FramebufferIsValid(fb.rid()) }

type rdTexture struct {
	rd.Variable
	rdNameable
}

func (tex *rdTexture) Clear(color uc.Color, base_mipmap, mipmap_count, base_layer, layer_count int, barrier rd.Barrier) error {
	return Error(tex.from.class.TextureClear(tex.rid(), color, Int(base_mipmap), Int(mipmap_count), Int(base_layer), Int(layer_count), RenderingDeviceBarrierMask(barrier)))
}

func (tex *rdTexture) Format() rd.TextureFormat {
	tmp := gd.NewContext(tex.from.godot.API)
	defer tmp.End()
	format := tex.from.class.TextureGetFormat(tmp, tex.rid())
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

func (tex *rdTexture) Handle() uintptr {
	return uintptr(tex.from.class.TextureGetNativeHandle(tex.rid()))
}

func (tex *rdTexture) IsShared() bool { return tex.from.class.TextureIsShared(tex.rid()) }

func (tex *rdTexture) IsValid() bool { return tex.from.class.TextureIsValid(tex.rid()) }

func (tex *rdTexture) Layer(layer int) rd.TextureData {
	return &rdTextureData{
		layer: layer,
		tex:   tex,
	}
}

type rdTextureData struct {
	tex   *rdTexture
	layer int

	array gd.PackedByteArray
	span  mmm.Lifetime

	r int
	w int
}

func (data *rdTextureData) Read(p []byte) (n int, err error) {
	if data.array == (gd.PackedByteArray{}) {
		tmp := gd.NewContext(data.tex.from.godot.API)
		data.array = data.tex.from.class.TextureGetData(tmp, data.tex.rid(), Int(data.layer))
		data.span = tmp.Lifetime
	}
	if data.r >= int(data.array.Size()) {
		return 0, io.EOF
	}
	buf := unsafe.Slice((*byte)(data.array.UnsafePointer()), data.array.Size())
	n = copy(p, buf[data.r:])
	data.r += n
	return
}

func (data *rdTextureData) ReadFrom(r io.Reader) (n int64, err error) {
	tmp := gd.NewContext(data.tex.from.godot.API)
	defer tmp.End()
	packed := tmp.PackedByteArray()
	for {
		var buf [4096]byte
		m, err := r.Read(buf[:])
		if m > 0 {
			packed.AppendArray(tmp.PackedByteSlice(buf[:m]))
		}
		if err == io.EOF {
			break
		}
	}
	err = Error(data.tex.from.class.TextureUpdate(data.tex.rid(), Int(data.layer), packed, RenderingDeviceBarrierMaskAllBarriers))
	return int64(packed.Size()), err
}

func (data *rdTextureData) Close() error { data.span.End(); return nil }

type rdDrawing struct {
	from *rdInterface
	list Int
}

func (d *rdDrawing) DebugBlock(name string, color uc.Color, block func()) {
	tmp := gd.NewContext(d.from.godot.API)
	defer tmp.End()
	d.from.class.DrawCommandBeginLabel(tmp.String(name), color)
	block()
	d.from.class.DrawCommandEndLabel()
}

func (d *rdDrawing) DebugLabel(name string, color uc.Color) {
	tmp := gd.NewContext(d.from.godot.API)
	defer tmp.End()
	d.from.class.DrawCommandInsertLabel(tmp.String(name), color)
}

func (d *rdDrawing) SetBlendConstant(color uc.Color) {
	d.from.class.DrawListSetBlendConstants(d.list, color)
}

func (d *rdDrawing) SetData(data []byte) {
	tmp := gd.NewContext(d.from.godot.API)
	defer tmp.End()
	d.from.class.DrawListSetPushConstant(d.list, tmp.PackedByteSlice(data), Int(len(data))) // FIXME can we avoid the copy here?
}

func (d *rdDrawing) SetIndexArray(array rd.IndexArray) {
	tmp := gd.NewContext(d.from.godot.API)
	defer tmp.End()
	d.from.class.DrawListBindIndexArray(d.list, array.(*rdIndexArray).rid())
}

func (d *rdDrawing) SetRenderer(renderer rd.Renderer) {
	d.from.class.DrawListBindRenderPipeline(d.list, renderer.(*rdRenderer).rid())
}

func (d *rdDrawing) SetScissor(region *Rect2) {
	if region == nil {
		d.from.class.DrawListDisableScissor(d.list)
	} else {
		d.from.class.DrawListEnableScissor(d.list, *region)
	}
}

func (d *rdDrawing) SetVariables(level rd.VariableLevel, variables rd.Variables) {
	d.from.class.DrawListBindUniformSet(d.list, variables.(*rdVariables).rid(), Int(level))
}

func (d *rdDrawing) SetVertexArray(array rd.VertexArray) {
	tmp := gd.NewContext(d.from.godot.API)
	defer tmp.End()
	d.from.class.DrawListBindVertexArray(d.list, array.(*rdVertexArray).rid())
}

func (d *rdDrawing) Submit(indices bool, instances, vertices int) {
	d.from.class.DrawListDraw(d.list, Bool(indices), Int(instances), Int(vertices))
}

func (d *rdDrawing) SwitchToNextPass() {
	d.from.class.DrawListSwitchToNextPass()
}

type rdIndexArray struct{ *rdResource }

type rdVertexArray struct{ *rdResource }

type rdRenderer struct{ *rdResource }

func (r rdRenderer) IsValid() bool {
	return r.from.class.RenderPipelineIsValid(r.rid())
}

func (gpu *rdInterface) DrawingOnScreen(screen rd.Screen, clear uc.Color, fn func(rd.Drawing)) {
	list := gpu.class.DrawListBeginForScreen(screen.(*rdScreen).name, clear)
	gpu.drawing.list = list
	gpu.drawing.from = gpu
	fn(&gpu.drawing)
	gpu.class.DrawListEnd(RenderingDeviceBarrierMaskAllBarriers) // FIXME should this be configurable?
}

type rdScreen struct {
	from *rdInterface
	name Int
}

func (sc *rdScreen) FramebufferFormat() rd.FramebufferFormat {
	return &rdFramebufferFormat{
		from: sc.from,
		name: sc.from.class.ScreenGetFramebufferFormat(),
		eyes: 1, // FIXME
	}
}

func (sc *rdScreen) Width() int {
	return int(sc.from.class.ScreenGetWidth(sc.name))
}

func (sc *rdScreen) Height() int {
	return int(sc.from.class.ScreenGetHeight(sc.name))
}

type rdFramebufferFormat struct {
	from   *rdInterface
	name   Int
	eyes   Int
	passes []rd.FramebufferPass
}

func (fmt *rdFramebufferFormat) Framebuffer(textures []rd.Texture) rd.Framebuffer {
	tmp := gd.NewContext(fmt.from.godot.API)
	defer tmp.End()
	array1 := NewArrayOf[RID](tmp)
	array1.Resize(Int(len(textures)))
	for i, texture := range textures {
		array1.SetIndex(Int(i), texture.(*rdTexture).rid())
	}
	if len(fmt.passes) == 0 {
		return &rdFramebuffer{fmt.from.resource(fmt.from.class.FramebufferCreate(array1, fmt.name, fmt.eyes))}
	}
	array2 := NewArrayOf[classdb.RDFramebufferPass](tmp)
	for _, pass := range fmt.passes {
		var converted = *gd.Create(tmp, new(classdb.RDFramebufferPass))
		converted.SetColorAttachments(tmp.PackedInt32Slice(pass.ColorAttachments))
		converted.SetDepthAttachment(Int(pass.DepthAttachment))
		converted.SetInputAttachments(tmp.PackedInt32Slice(pass.InputAttachments))
		converted.SetPreserveAttachments(tmp.PackedInt32Slice(pass.AttachmentsToPreserve))
		converted.SetResolveAttachments(tmp.PackedInt32Slice(pass.AttachmentsToResolve))
		array2.Append(converted)
	}
	return &rdFramebuffer{fmt.from.resource(fmt.from.class.FramebufferCreateMultipass(array1, array2, fmt.name, fmt.eyes))}
}

func (fmt *rdFramebufferFormat) TextureSamples(pass int) rd.TextureSamples {
	return rd.TextureSamples(fmt.from.class.FramebufferFormatGetTextureSamples(fmt.name, Int(pass)))
}

func (gpu *rdInterface) ExtensionTexture(ttype rd.TextureType, format rd.DataFormat, samples rd.TextureSamples, usage rd.TextureUsage, image uintptr, width, height, depth, layers int) rd.Texture {
	return &rdTexture{rdNameable: rdNameable{gpu.resource(gpu.class.TextureCreateFromExtension(
		RenderingDeviceTextureType(ttype),
		RenderingDeviceDataFormat(format),
		RenderingDeviceTextureSamples(samples),
		RenderingDeviceTextureUsageBits(usage), *(*Int)(unsafe.Pointer(&image)), Int(width), Int(height), Int(depth), Int(layers),
	))}}
}

func (gpu *rdInterface) FrameDelay() int { return int(gpu.class.GetFrameDelay()) }

func (gpu *rdInterface) FramebufferFormat(eyes int, attachments []rd.AttachmentFormat, passes []rd.FramebufferPass) rd.FramebufferFormat {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	array1 := NewArrayOf[classdb.RDAttachmentFormat](tmp)
	for _, attachment := range attachments {
		var converted = *gd.Create(tmp, new(classdb.RDAttachmentFormat))
		converted.SetFormat(RenderingDeviceDataFormat(attachment.Format))
		converted.SetSamples(RenderingDeviceTextureSamples(attachment.Samples))
		converted.SetUsageFlags(int64(attachment.Usage))
		array1.Append(converted)
	}
	array2 := NewArrayOf[classdb.RDFramebufferPass](tmp)
	for _, pass := range passes {
		var converted = *gd.Create(tmp, new(classdb.RDFramebufferPass))
		converted.SetColorAttachments(tmp.PackedInt32Slice(pass.ColorAttachments))
		converted.SetDepthAttachment(Int(pass.DepthAttachment))
		converted.SetInputAttachments(tmp.PackedInt32Slice(pass.InputAttachments))
		converted.SetPreserveAttachments(tmp.PackedInt32Slice(pass.AttachmentsToPreserve))
		converted.SetResolveAttachments(tmp.PackedInt32Slice(pass.AttachmentsToResolve))
		array2.Append(converted)
	}
	return &rdFramebufferFormat{
		from: gpu,
		name: gpu.class.FramebufferFormatCreateMultipass(array1, array2, Int(eyes)),
		eyes: Int(eyes),
	}
}

func (gpu *rdInterface) IndexBufferU16(data []uint16) rd.IndexBuffer {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	bytes := unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), uintptr(len(data))*unsafe.Sizeof(uint16(0)))
	return rdIndexBuffer{rdBuffer{rdNameable{gpu.resource(gpu.class.IndexBufferCreate(Int(len(bytes)), RenderingDeviceIndexBufferFormatUint16, tmp.PackedByteSlice(bytes), false))}}}
}

func (gpu *rdInterface) IndexBufferU32(data []uint32) rd.IndexBuffer {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	bytes := unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), uintptr(len(data))*unsafe.Sizeof(uint32(0)))
	return rdIndexBuffer{rdBuffer{rdNameable{gpu.resource(gpu.class.IndexBufferCreate(Int(len(bytes)), RenderingDeviceIndexBufferFormatUint32, tmp.PackedByteSlice(bytes), false))}}}
}

type rdIndexBuffer struct{ rdBuffer }

type rdBuffer struct{ rdNameable }

func (buf rdBuffer) Clear() error {
	return Error(buf.from.class.BufferClear(buf.rid(), 0, 0, RenderingDeviceBarrierMaskAllBarriers))
}

func (buf rdBuffer) ReadAt(b []byte, off int64) (n int, err error) {
	tmp := gd.NewContext(buf.from.godot.API)
	defer tmp.End()
	data := buf.from.class.BufferGetData(tmp, buf.rid(), Int(off), Int(len(b)))
	copy(b, unsafe.Slice((*byte)(data.UnsafePointer()), data.Size()))
	return int(data.Size()), nil
}

func (buf rdBuffer) WriteAt(b []byte, off int64) (n int, err error) {
	tmp := gd.NewContext(buf.from.godot.API)
	defer tmp.End()
	data := tmp.PackedByteSlice(b)
	err = Error(buf.from.class.BufferUpdate(buf.rid(), Int(off), Int(len(b)), data, RenderingDeviceBarrierMaskAllBarriers)) // FIXME can we avoid the copy here?
	return int(data.Size()), err
}

func (gpu rdInterface) Limit(limit rd.Limit) int {
	return int(gpu.class.LimitGet(RenderingDeviceLimit(limit)))
}

func (gpu rdInterface) MemoryUsage(mtype rd.MemoryType) int {
	return int(gpu.class.GetMemoryUsage(RenderingDeviceMemoryType(mtype)))
}

func (gpu rdInterface) PipelineCache() string {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	return gpu.class.GetDevicePipelineCacheUuid(tmp).String()
}

func (gpu rdInterface) Processor(shader rd.Shader, defines []any) rd.Processor {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	array := NewArrayOf[classdb.RDPipelineSpecializationConstant](tmp)
	for i, define := range defines {
		var converted = *gd.Create(tmp, new(classdb.RDPipelineSpecializationConstant))
		converted.SetConstantId(Int(i))
		rvalue := reflect.ValueOf(define)
		switch reflect.TypeOf(define).Kind() {
		case reflect.Bool:
			converted.SetValue(tmp.Variant(Bool(rvalue.Bool())))
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			converted.SetValue(tmp.Variant(Int(rvalue.Int())))
		case reflect.Uint8, reflect.Uint16, reflect.Uint32:
			converted.SetValue(tmp.Variant(Int(rvalue.Uint())))
		case reflect.Float32, reflect.Float64:
			converted.SetValue(tmp.Variant(Float(rvalue.Float())))
		default:
			panic("unexpected define type " + reflect.TypeOf(define).String())
		}
		array.Append(converted)
	}
	return rdProcessor{gpu.resource(gpu.class.ComputePipelineCreate(shader.(*rdShader).rid(), array))}
}

func (gpu rdInterface) Renderer(shader rd.Shader, options rd.RenderingOptions) rd.Renderer {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()

	rasterization := *Create(tmp, new(classdb.RDPipelineRasterizationState))
	rasterization.SetCullMode(RenderingDevicePolygonCullMode(options.Rasterization.CullMode))
	rasterization.SetDepthBiasClamp(options.Rasterization.DepthBiasClamp)
	rasterization.SetDepthBiasConstantFactor(options.Rasterization.DepthBiasConstantFactor)
	rasterization.SetDepthBiasEnabled(Bool(options.Rasterization.DepthBiasEnabled))
	rasterization.SetDepthBiasSlopeFactor(options.Rasterization.DepthBiasSlopeFactor)
	rasterization.SetDiscardPrimitives(Bool(options.Rasterization.DiscardPrimitives))
	rasterization.SetEnableDepthClamp(Bool(options.Rasterization.EnableDepthClamp))
	rasterization.SetFrontFace(RenderingDevicePolygonFrontFace(options.Rasterization.FrontFace))
	rasterization.SetLineWidth(options.Rasterization.LineWidth)
	rasterization.SetPatchControlPoints(options.Rasterization.PatchControlPoints)
	rasterization.SetWireframe(Bool(options.Rasterization.Wireframe))

	multisampling := *Create(tmp, new(classdb.RDPipelineMultisampleState))
	multisampling.SetEnableAlphaToCoverage(Bool(options.Multisampling.EnableAlphaToCoverage))
	multisampling.SetEnableAlphaToOne(Bool(options.Multisampling.EnableAlphaToOne))
	multisampling.SetEnableSampleShading(Bool(options.Multisampling.EnableSampleShading))
	multisampling.SetMinSampleShading(options.Multisampling.MinSampleShading)
	multisampling.SetSampleCount(RenderingDeviceTextureSamples(options.Multisampling.Samples))
	masks := NewArrayOf[Int](tmp)
	masks.Resize(Int(len(options.Multisampling.SampleMasks)))
	for i, mask := range options.Multisampling.SampleMasks {
		masks.SetIndex(Int(i), Int(mask))
	}
	multisampling.SetSampleMasks(masks)

	depthstencils := *Create(tmp, new(classdb.RDPipelineDepthStencilState))
	depthstencils.SetBackOpCompare(RenderingDeviceCompareOperator(options.DepthStencils.BackOperationComparison))
	depthstencils.SetBackOpCompareMask(int64(options.DepthStencils.BackOperationComparisonMask))
	depthstencils.SetBackOpDepthFail(RenderingDeviceStencilOperation(options.DepthStencils.BackOperationDepthFail))
	depthstencils.SetBackOpPass(RenderingDeviceStencilOperation(options.DepthStencils.BackOperationPass))
	depthstencils.SetBackOpFail(RenderingDeviceStencilOperation(options.DepthStencils.BackOperationFail))
	depthstencils.SetBackOpReference(int64(options.DepthStencils.BackOperationReference))
	depthstencils.SetBackOpWriteMask(int64(options.DepthStencils.BackOperationWriteMask))
	depthstencils.SetDepthCompareOperator(RenderingDeviceCompareOperator(options.DepthStencils.DepthComparison))
	depthstencils.SetDepthRangeMax(options.DepthStencils.DepthRangeMax)
	depthstencils.SetDepthRangeMin(options.DepthStencils.DepthRangeMin)
	depthstencils.SetEnableDepthRange(Bool(options.DepthStencils.EnableDepthRange))
	depthstencils.SetEnableDepthTest(Bool(options.DepthStencils.EnableDepthTest))
	depthstencils.SetEnableDepthWrite(Bool(options.DepthStencils.EnableDepthWrite))
	depthstencils.SetEnableStencil(Bool(options.DepthStencils.EnableStencil))
	depthstencils.SetFrontOpCompare(RenderingDeviceCompareOperator(options.DepthStencils.FrontOperationComparison))
	depthstencils.SetFrontOpCompareMask(int64(options.DepthStencils.FrontOperationComparisonMask))
	depthstencils.SetFrontOpDepthFail(RenderingDeviceStencilOperation(options.DepthStencils.FrontOperationDepthFail))
	depthstencils.SetFrontOpPass(RenderingDeviceStencilOperation(options.DepthStencils.FrontOperationPass))
	depthstencils.SetFrontOpFail(RenderingDeviceStencilOperation(options.DepthStencils.FrontOperationFail))
	depthstencils.SetFrontOpReference(int64(options.DepthStencils.FrontOperationReference))
	depthstencils.SetFrontOpWriteMask(int64(options.DepthStencils.FrontOperationWriteMask))

	colorblending := *Create(tmp, new(classdb.RDPipelineColorBlendState))
	colorblending.SetBlendConstant(options.ColorBlending.BlendConstant)
	colorblending.SetEnableLogicOp(Bool(options.ColorBlending.EnableLogicOperation))
	colorblending.SetLogicOp(RenderingDeviceLogicOperation(options.ColorBlending.LogicOperation))
	attachments := NewArrayOf[classdb.RDPipelineColorBlendStateAttachment](tmp)
	for _, attachment := range options.ColorBlending.Attachments {
		var converted = *gd.Create(tmp, new(classdb.RDPipelineColorBlendStateAttachment))
		converted.SetEnableBlend(Bool(attachment.EnableBlend))

		toBlendOperation := func(op Int) RenderingDeviceBlendOperation {
			switch op {
			case Int(rd.Source + rd.Destination):
				return RenderingDeviceBlendOpAdd
			case Int(rd.Source - rd.Destination):
				return RenderingDeviceBlendOpSubtract
			case Int(rd.Destination - rd.Source):
				return RenderingDeviceBlendOpReverseSubtract
			case Int(min(rd.Source, rd.Destination)):
				return RenderingDeviceBlendOpMinimum
			case Int(max(rd.Source, rd.Destination)):
				return RenderingDeviceBlendOpMaximum
			default:
				return RenderingDeviceBlendOperation(op)
			}
		}
		converted.SetAlphaBlendOp(toBlendOperation(Int(attachment.AlphaBlendOperation)))
		converted.SetColorBlendOp(toBlendOperation(Int(attachment.ColorBlendOperation)))

		toBlendFactor := func(op rd.BlendFactor) RenderingDeviceBlendFactor {
			switch op {
			case rd.One - rd.SourceColor:
				return RenderingDeviceBlendFactorOneMinusSrcColor
			case rd.One - rd.SourceAlpha:
				return RenderingDeviceBlendFactorOneMinusSrcAlpha
			case rd.One - rd.DestinationColor:
				return RenderingDeviceBlendFactorOneMinusDstColor
			case rd.One - rd.DestinationAlpha:
				return RenderingDeviceBlendFactorOneMinusDstAlpha
			case rd.One - rd.ConstantAlpha:
				return RenderingDeviceBlendFactorOneMinusConstantAlpha
			case rd.One - rd.ConstantColor:
				return RenderingDeviceBlendFactorOneMinusConstantColor
			case rd.One - rd.SecondSourceAlpha:
				return RenderingDeviceBlendFactorOneMinusSrc1Alpha
			case rd.One - rd.SecondSourceColor:
				return RenderingDeviceBlendFactorOneMinusSrc1Color
			default:
				return RenderingDeviceBlendFactor(op)
			}
		}
		converted.SetDstAlphaBlendFactor(toBlendFactor(attachment.DestinationAlphaBlendFactor))
		converted.SetDstColorBlendFactor(toBlendFactor(attachment.DestinationColorBlendFactor))
		converted.SetSrcAlphaBlendFactor(toBlendFactor(attachment.SourceAlphaBlendFactor))
		converted.SetSrcColorBlendFactor(toBlendFactor(attachment.SourceColorBlendFactor))

		converted.SetWriteA(Bool(attachment.WriteA))
		converted.SetWriteB(Bool(attachment.WriteB))
		converted.SetWriteG(Bool(attachment.WriteG))
		converted.SetWriteR(Bool(attachment.WriteR))

		attachments.Append(converted)
	}

	constants := NewArrayOf[classdb.RDPipelineSpecializationConstant](tmp)
	for i, define := range options.ShaderDefines {
		var converted = *gd.Create(tmp, new(classdb.RDPipelineSpecializationConstant))
		converted.SetConstantId(Int(i))
		rvalue := reflect.ValueOf(define)
		switch reflect.TypeOf(define).Kind() {
		case reflect.Bool:
			converted.SetValue(tmp.Variant(Bool(rvalue.Bool())))
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			converted.SetValue(tmp.Variant(Int(rvalue.Int())))
		case reflect.Uint8, reflect.Uint16, reflect.Uint32:
			converted.SetValue(tmp.Variant(Int(rvalue.Uint())))
		case reflect.Float32, reflect.Float64:
			converted.SetValue(tmp.Variant(Float(rvalue.Float())))
		default:
			panic("unexpected define type " + reflect.TypeOf(define).String())
		}
		constants.Append(converted)
	}

	rid := gpu.class.RenderPipelineCreate(
		shader.(*rdShader).rid(),
		options.FramebufferFormat.(*rdFramebufferFormat).name,
		Int(options.VertexFormat),
		RenderingDeviceRenderPrimitive(options.PrimitiveType),
		rasterization,
		multisampling,
		depthstencils,
		colorblending,
		RenderingDevicePipelineDynamicStateFlags(options.DynamicStates),
		Int(options.RenderingPass),
		constants,
	)
	return rdRenderer{gpu.resource(rid)}
}

func (gpu rdInterface) RenderingDevice() rd.Local {
	return rdLocal{&rdInterface{
		godot: gpu.godot,
		class: gpu.class.CreateLocalDevice(gpu.godot),
	}}
}

type rdLocal struct {
	*rdInterface
}

func (gpu rdLocal) Submit() { gpu.class.Submit() }
func (gpu rdLocal) Sync()   { gpu.class.Sync() }

func (gpu rdInterface) Sampler(state rd.SamplerState) rd.Sampler {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	converted := *gd.Create(tmp, new(classdb.RDSamplerState))
	converted.SetAnisotropyMax(state.AnisotropyMax)
	converted.SetBorderColor(RenderingDeviceSamplerBorderColor(state.BorderColor))
	converted.SetCompareOp(RenderingDeviceCompareOperator(state.Comparison))
	converted.SetEnableCompare(Bool(state.EnableComparison))
	converted.SetLodBias(state.LevelOfDetailBias)
	converted.SetMagFilter(RenderingDeviceSamplerFilter(state.MagnificationFilter))
	converted.SetMaxLod(state.MaxLevelOfDetail)
	converted.SetMinFilter(RenderingDeviceSamplerFilter(state.MinificationFilter))
	converted.SetMinLod(state.MinLevelOfDetail)
	converted.SetMipFilter(RenderingDeviceSamplerFilter(state.MipmapFilter))
	converted.SetRepeatU(RenderingDeviceSamplerRepeatMode(state.RepeatU))
	converted.SetRepeatV(RenderingDeviceSamplerRepeatMode(state.RepeatV))
	converted.SetRepeatW(RenderingDeviceSamplerRepeatMode(state.RepeatW))
	converted.SetUnnormalizedUvw(Bool(state.UnnormalizedUVW))
	converted.SetUseAnisotropy(Bool(state.UseAnisotropy))
	return rdSampler{rdNameable: rdNameable{gpu.resource(gpu.class.SamplerCreate(converted))}}
}

type rdSampler struct {
	rd.Variable
	rdNameable
}

func (s rdSampler) FormatSupportedForFilter(format rd.DataFormat, filter rd.Filter) bool {
	return s.from.class.SamplerIsFormatSupportedForFilter(RenderingDeviceDataFormat(format), RenderingDeviceSamplerFilter(filter))
}

func (gpu *rdInterface) Screen(n int) rd.Screen {
	return &rdScreen{
		from: gpu,
		name: Int(n),
	}
}

func (gpu *rdInterface) Shader() rd.Shader {
	return &rdShader{rdNameable{gpu.resource(gpu.class.ShaderCreatePlaceholder())}}
}

func (gpu *rdInterface) SharedTexture(view rd.TextureView, with rd.Texture) rd.Texture {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()

	converted := *gd.Create(tmp, new(classdb.RDTextureView))
	converted.SetFormatOverride(RenderingDeviceDataFormat(view.FormatOverride))
	converted.SetSwizzleA(RenderingDeviceTextureSwizzle(view.SwizzleAlpha))
	converted.SetSwizzleB(RenderingDeviceTextureSwizzle(view.SwizzleBlue))
	converted.SetSwizzleG(RenderingDeviceTextureSwizzle(view.SwizzleGreen))
	converted.SetSwizzleR(RenderingDeviceTextureSwizzle(view.SwizzleRed))

	return &rdTexture{rdNameable: rdNameable{gpu.resource(gpu.class.TextureCreateShared(converted, with.(*rdTexture).rid()))}}
}

func (gpu *rdInterface) StorageBuffer(usage rd.StorageBufferUsage, data []byte) rd.StorageBuffer {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	bytes := tmp.PackedByteSlice(data) // FIXME can we avoid the copy here?
	return rdStorageBuffer{rdBuffer: rdBuffer{rdNameable{gpu.resource(gpu.class.StorageBufferCreate(Int(len(data)), bytes, RenderingDeviceStorageBufferUsage(usage)))}}}
}

type rdStorageBuffer struct {
	rd.Variable
	rdBuffer
}

func (gpu *rdInterface) Texture(format rd.TextureFormat, view rd.TextureView, data [][]byte) rd.Texture {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()

	texture_format := *gd.Create(tmp, new(classdb.RDTextureFormat))
	texture_format.SetArrayLayers(Int(format.ArrayLayers))
	texture_format.SetDepth(Int(format.Depth))
	texture_format.SetFormat(RenderingDeviceDataFormat(format.Format))
	texture_format.SetHeight(Int(format.Height))
	texture_format.SetMipmaps(Int(format.Mipmaps))
	texture_format.SetSamples(RenderingDeviceTextureSamples(format.Samples))
	texture_format.SetTextureType(RenderingDeviceTextureType(format.TextureType))
	texture_format.SetUsageBits(RenderingDeviceTextureUsageBits(format.Usage))
	texture_format.SetWidth(Int(format.Width))
	for share := range format.ShareableFormats {
		texture_format.AddShareableFormat(RenderingDeviceDataFormat(share))
	}

	converted := *gd.Create(tmp, new(classdb.RDTextureView))
	converted.SetFormatOverride(RenderingDeviceDataFormat(view.FormatOverride))
	converted.SetSwizzleA(RenderingDeviceTextureSwizzle(view.SwizzleAlpha))
	converted.SetSwizzleB(RenderingDeviceTextureSwizzle(view.SwizzleBlue))
	converted.SetSwizzleG(RenderingDeviceTextureSwizzle(view.SwizzleGreen))
	converted.SetSwizzleR(RenderingDeviceTextureSwizzle(view.SwizzleRed))

	buffers := NewArrayOf[gd.PackedByteArray](tmp)
	for _, slice := range data {
		buffers.Append(tmp.PackedByteSlice(slice)) // FIXME can we avoid the copy here?
	}
	return &rdTexture{rdNameable: rdNameable{gpu.resource(gpu.class.TextureCreate(texture_format, converted, buffers))}}
}

func (gpu *rdInterface) TextureBuffer(format rd.DataFormat, data []byte) rd.TextureBuffer {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	bytes := tmp.PackedByteSlice(data) // FIXME can we avoid the copy here?
	return rdTextureBuffer{rdBuffer: rdBuffer{rdNameable{gpu.resource(gpu.class.TextureBufferCreate(Int(len(data)), RenderingDeviceDataFormat(format), bytes))}}}
}

type rdTextureBuffer struct {
	rd.Variable
	rdBuffer
}

func (gpu *rdInterface) TextureCopy(src, dst rd.Texture, from, into, size xy.Vector3, src_mipmap, dst_mipmp, src_layer, dst_layer int, barrier rd.Barrier) error {
	return Error(gpu.class.TextureCopy(
		src.(*rdTexture).rid(),
		dst.(*rdTexture).rid(),
		from,
		into,
		size,
		Int(src_mipmap),
		Int(dst_mipmp),
		Int(src_layer),
		Int(dst_layer),
		RenderingDeviceBarrierMask(barrier),
	))
}

func (gpu *rdInterface) TextureFormatIsSupportedForUsage(format rd.DataFormat, usage rd.TextureUsage) bool {
	return gpu.class.TextureIsFormatSupportedForUsage(RenderingDeviceDataFormat(format), RenderingDeviceTextureUsageBits(usage))
}

func (gpu *rdInterface) TextureResolveMultiSample(from, into rd.Texture, barrier rd.Barrier) error {
	return Error(gpu.class.TextureResolveMultisample(from.(*rdTexture).rid(), into.(*rdTexture).rid(), RenderingDeviceBarrierMask(barrier)))
}

func (gpu *rdInterface) UniformBuffer(data []byte) rd.UniformBuffer {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	bytes := tmp.PackedByteSlice(data) // FIXME can we avoid the copy here?
	return rdUniformBuffer{rdBuffer: rdBuffer{rdNameable{gpu.resource(gpu.class.UniformBufferCreate(Int(len(data)), bytes))}}}
}

type rdUniformBuffer struct {
	rd.Variable
	rdBuffer
}

// VertexArray creates a new vertex from the specified buffers, optionally, with offsets.
func (gpu *rdInterface) VertexArray(vertices int, format rd.VertexFormat, buffers []rd.Buffer, offsets []int64) rd.VertexArray {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	array1 := NewArrayOf[RID](tmp)
	array1.Resize(Int(len(buffers)))
	for i, buffer := range buffers {
		array1.SetIndex(Int(i), buffer.(*rdBuffer).rid())
	}
	return &rdVertexArray{gpu.resource(gpu.class.VertexArrayCreate(
		Int(vertices),
		Int(format),
		array1,
		tmp.PackedInt64Slice(offsets),
	))}
}

// VertexBuffer creates a new vertex buffer with the specified data.
func (gpu *rdInterface) VertexBuffer(data []byte) rd.VertexBuffer {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	bytes := tmp.PackedByteSlice(data) // FIXME can we avoid the copy here?
	return rdVertexBuffer{rdBuffer: rdBuffer{rdNameable{gpu.resource(gpu.class.VertexBufferCreate(Int(len(data)), bytes, false))}}}
}

type rdVertexBuffer struct {
	rd.Variable
	rdBuffer
}

// VertexFormat creates a new vertex format with the specified attributes.
func (gpu *rdInterface) VertexFormat(attributes []rd.VertexAttribute) rd.VertexFormat {
	tmp := gd.NewContext(gpu.godot.API)
	defer tmp.End()
	array := NewArrayOf[classdb.RDVertexAttribute](tmp)
	for _, attribute := range attributes {
		var converted = *gd.Create(tmp, new(classdb.RDVertexAttribute))
		converted.SetFormat(RenderingDeviceDataFormat(attribute.Format))
		converted.SetFrequency(RenderingDeviceVertexFrequency(attribute.Frequency))
		converted.SetLocation(Int(attribute.Location))
		converted.SetOffset(Int(attribute.Offset))
		converted.SetStride(Int(attribute.Stride))
		array.Append(converted)
	}
	return rd.VertexFormat(gpu.class.VertexFormatCreate(array))

}

/*func (gpu rdInterface) scope(lifetime mmm.Lifetime) Lifetime {
	return gd.Lifetime{
		Lifetime: lifetime,
		API:      gpu.godot.API,
	}
}

func (gpu rdInterface) bytes(godot Lifetime, data ffi.Bytes) gd.PackedByteArray {
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

func (gpu rdInterface) string(godot Lifetime, data ffi.Bytes) gd.String {
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

func (gpu rdInterface) colors(godot Lifetime, data ffi.Slice[uc.Color]) gd.PackedColorArray {
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

func (gpu rdInterface) int32s(godot Lifetime, data ffi.Slice[int32]) gd.PackedInt32Array {
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

func (gpu rdInterface) int64s(godot Lifetime, data ffi.Slice[int64]) gd.PackedInt64Array {
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

func (gpu rdInterface) textures(data ffi.Managed[[]rd.Texture]) gd.ArrayOf[RID] {
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

func (gpu rdInterface) framebufferPasses(data ffi.Managed[[]rd.FramebufferPass]) gd.ArrayOf[classdb.RDFramebufferPass] {
	array, ok := data.Interface().(gd.ArrayOf[classdb.RDFramebufferPass])
	if ok {
		return array
	}
	slice := data.Value()
	span := gpu.scope(data.Lifetime())
	array = NewArrayOf[classdb.RDFramebufferPass](span)
	for _, pass := range slice {
		var converted = *Create(span, new(classdb.RDFramebufferPass))
		converted.SetColorAttachments(gpu.int32s(span, ffi.Make(pass.ColorAttachments)))
		converted.SetDepthAttachment(Int(pass.DepthAttachment))
		converted.SetInputAttachments(gpu.int32s(span, ffi.Make(pass.InputAttachments)))
		converted.SetPreserveAttachments(gpu.int32s(span, ffi.Make(pass.PreserveAttachments)))
		converted.SetResolveAttachments(gpu.int32s(span, ffi.Make(pass.ResolveAttachments)))
		array.Append(converted)
	}
	data.Cache(array)
	return array
}

func (gpu rdInterface) attachments(data ffi.Managed[[]rd.AttachmentFormat]) gd.ArrayOf[classdb.RDAttachmentFormat] {
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
}*/

/*func (gpu rdInterface) BufferClear(buffer rd.Buffer, offset, size_bytes int, post_barrier rd.Barrier) error {
	return Error(gpu.class.BufferClear(RID(mmm.Get(buffer)), Int(offset), Int(size_bytes), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu rdInterface) BufferGetData(lifetime mmm.Lifetime, buffer rd.Buffer, offset_bytes, size_bytes int) ffi.Bytes {
	return ffi.Bytes{Interface: gpu.class.BufferGetData(gpu.scope(lifetime), RID(mmm.Get(buffer)), Int(offset_bytes), Int(size_bytes))}
}

func (gpu rdInterface) BufferUpdate(buffer rd.Buffer, offset, size_bytes int, data ffi.Bytes, post_barrier rd.Barrier) error {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return Error(gpu.class.BufferUpdate(RID(mmm.Get(buffer)), Int(offset), Int(size_bytes), gpu.bytes(tmp, data), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu rdInterface) ComputeListBegin(allow_draw_overlap bool) rd.ComputeList {
	return rd.ComputeList(gpu.class.ComputeListBegin(Bool(allow_draw_overlap)))
}

func (gpu rdInterface) ComputeListAddBarrier(compute_list rd.ComputeList) {
	gpu.class.ComputeListAddBarrier(Int(compute_list))
}

func (gpu rdInterface) ComputeListBindComputePipeline(list rd.ComputeList, pipeline rd.ComputePipeline) {
	gpu.class.ComputeListBindComputePipeline(Int(list), RID(mmm.Get(pipeline)))
}

func (gpu rdInterface) ComputeListBindUniformSet(list rd.ComputeList, set rd.UniformSet, set_index int) {
	gpu.class.ComputeListBindUniformSet(Int(list), RID(mmm.Get(set)), Int(set_index))
}

func (gpu rdInterface) ComputeListDispatch(list rd.ComputeList, groups_x, groups_y, groups_z int) {
	gpu.class.ComputeListDispatch(Int(list), Int(groups_x), Int(groups_y), Int(groups_z))
}

func (gpu rdInterface) ComputeListEnd(mask rd.Barrier) {
	gpu.class.ComputeListEnd(classdb.RenderingDeviceBarrierMask(mask))
}

func (gpu rdInterface) ComputeListSetPushConstant(list rd.ComputeList, buffer ffi.Bytes, size_bytes int) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.ComputeListSetPushConstant(Int(list), gpu.bytes(tmp, buffer), Int(size_bytes))
}

func (gpu rdInterface) ComputePipelineCreate(span mmm.Lifetime, shader rd.Shader, constants ffi.Managed[[]rd.ShaderDefine]) rd.ComputePipeline {
	array, ok := constants.Interface().(gd.ArrayOf[classdb.RDPipelineSpecializationConstant])
	if !ok {
		span := gpu.scope(constants.Lifetime())
		array = NewArrayOf[classdb.RDPipelineSpecializationConstant](span)
		for i, constant := range constants.Value() {
			var converted = *Create(span, new(classdb.RDPipelineSpecializationConstant))
			converted.SetConstantId(Int(i))
			converted.SetValue(span.Variant(constant.Interface()))
			array.Append(converted)
		}
		constants.Cache(array)
	}
	pipeline := uintptr(gpu.class.ComputePipelineCreate(RID(mmm.Get(shader)), array))
	return mmm.New[rd.ComputePipeline, rd.API, uintptr](span, gpu.value, pipeline)
}

func (gpu rdInterface) ComputePipelineIsValid(pipeline rd.ComputePipeline) bool {
	return bool(gpu.class.ComputePipelineIsValid(RID(mmm.Get(pipeline))))
}

func (gpu rdInterface) CreateLocalDevice(span mmm.Lifetime) *rd.Allocator {
	var device = rdInterface{
		godot: gpu.godot,
		class: gpu.class.CreateLocalDevice(gpu.scope(span)),
	}
	var itfc = rd.Allocator{Lifetime: span, API: device}
	device.value = &itfc.API
	return &itfc
}

func (gpu rdInterface) DrawCommandBeginLabel(name string, color uc.Color) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.DrawCommandBeginLabel(tmp.String(name), color)
}

func (gpu rdInterface) DrawCommandEndLabel() {
	gpu.class.DrawCommandEndLabel()
}

func (gpu rdInterface) DrawCommandInsertLabel(name string, color uc.Color) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.DrawCommandInsertLabel(tmp.String(name), color)
}

func (gpu rdInterface) DrawListBegin(frame rd.Frame) rd.DrawList {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	initial_color_action, final_color_action := frame.Color.Split()
	initial_depth_action, final_depth_action := frame.Depth.Split()
	return rd.DrawList(gpu.class.DrawListBegin(RID(mmm.Get(frame.Buffer)),
		classdb.RenderingDeviceInitialAction(initial_color_action), classdb.RenderingDeviceFinalAction(final_color_action),
		classdb.RenderingDeviceInitialAction(initial_depth_action), classdb.RenderingDeviceFinalAction(final_depth_action),
		gpu.colors(tmp, ffi.Make(frame.Clear.Colors)), Float(frame.Clear.Depth), Int(frame.Clear.Stencil),
		Rect2(frame.Region), gpu.textures(ffi.Manage(frame.Storage)),
	))
}

func (gpu rdInterface) DrawListBeginForScreen(screen rd.Screen, clear_color uc.Color) rd.DrawList {
	return rd.DrawList(gpu.class.DrawListBeginForScreen(Int(screen), clear_color))
}

func (gpu rdInterface) DrawListBeginSplit(span mmm.Lifetime, splits int, frame rd.Frame) ffi.Slice[rd.DrawList] {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	initial_color_action, final_color_action := frame.Color.Split()
	initial_depth_action, final_depth_action := frame.Depth.Split()
	lists := gpu.class.DrawListBeginSplit(gpu.scope(span), RID(mmm.Get(frame.Buffer)), Int(splits),
		classdb.RenderingDeviceInitialAction(initial_color_action), classdb.RenderingDeviceFinalAction(final_color_action),
		classdb.RenderingDeviceInitialAction(initial_depth_action), classdb.RenderingDeviceFinalAction(final_depth_action),
		gpu.colors(tmp, ffi.Make(frame.Clear.Colors)), Float(frame.Clear.Depth), Int(frame.Clear.Stencil),
		Rect2(frame.Region), gpu.textures(ffi.Manage(frame.Storage)),
	)
	return ffi.Slice[rd.DrawList]{Interface: lists}
}

func (gpu rdInterface) DrawListBindIndexArray(list rd.DrawList, array rd.IndexArray) {
	gpu.class.DrawListBindIndexArray(Int(list), RID(mmm.Get(array)))
}

func (gpu rdInterface) DrawListBindRenderPipeline(list rd.DrawList, pipeline rd.RenderPipeline) {
	gpu.class.DrawListBindRenderPipeline(Int(list), RID(mmm.Get(pipeline)))
}

func (gpu rdInterface) DrawListBindUniformSet(list rd.DrawList, set rd.UniformSet, set_index int) {
	gpu.class.DrawListBindUniformSet(Int(list), RID(mmm.Get(set)), Int(set_index))
}

func (gpu rdInterface) DrawListBindVertexArray(list rd.DrawList, array rd.VertexArray) {
	gpu.class.DrawListBindVertexArray(Int(list), RID(mmm.Get(array)))
}

func (gpu rdInterface) DrawListDisableScissor(list rd.DrawList) {
	gpu.class.DrawListDisableScissor(Int(list))
}

func (gpu rdInterface) DrawListDraw(list rd.DrawList, use_indices bool, instances int, procedural_vertex_count int) {
	gpu.class.DrawListDraw(Int(list), Bool(use_indices), Int(instances), Int(procedural_vertex_count))
}

func (gpu rdInterface) DrawListEnableScissor(list rd.DrawList, rect xy.Rect2) {
	gpu.class.DrawListEnableScissor(Int(list), Rect2(rect))
}

func (gpu rdInterface) DrawListEnd(mask rd.Barrier) {
	gpu.class.DrawListEnd(classdb.RenderingDeviceBarrierMask(mask))
}

func (gpu rdInterface) DrawListSetBlendConstants(list rd.DrawList, color uc.Color) {
	gpu.class.DrawListSetBlendConstants(Int(list), color)
}

func (gpu rdInterface) DrawListSetPushConstant(list rd.DrawList, buffer ffi.Bytes, size_bytes int) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.DrawListSetPushConstant(Int(list), gpu.bytes(tmp, buffer), Int(size_bytes))
}

func (gpu rdInterface) DrawListSwitchToNextPass() rd.DrawList {
	return rd.DrawList(gpu.class.DrawListSwitchToNextPass())
}

func (gpu rdInterface) DrawListSwitchToNextPassSplit(span mmm.Lifetime, splits int) ffi.Slice[rd.DrawList] {
	lists := gpu.class.DrawListSwitchToNextPassSplit(gpu.scope(span), Int(splits))
	return ffi.Slice[rd.DrawList]{Interface: lists}
}

func (gpu rdInterface) FramebufferCreate(span mmm.Lifetime, textures ffi.Managed[[]rd.Texture], validate_with_format int64, view_count int) rd.Framebuffer {
	framebuffer := gpu.class.FramebufferCreate(gpu.textures(textures), Int(validate_with_format), Int(view_count))
	return mmm.New[rd.Framebuffer, rd.API, uintptr](span, gpu.value, uintptr(framebuffer))
}

func (gpu rdInterface) FramebufferCreateEmpty(span mmm.Lifetime, size xy.Vector2i, samples rd.TextureSamples, validate_with_format int64) rd.Framebuffer {
	framebuffer := gpu.class.FramebufferCreateEmpty(Vector2i(size), classdb.RenderingDeviceTextureSamples(samples), Int(validate_with_format))
	return mmm.New[rd.Framebuffer, rd.API, uintptr](span, gpu.value, uintptr(framebuffer))
}

func (gpu rdInterface) FramebufferCreateMultipass(span mmm.Lifetime, textures ffi.Managed[[]rd.Texture], passes ffi.Managed[[]rd.FramebufferPass], validate_with_format int64, view_count int) rd.Framebuffer {
	framebuffer := gpu.class.FramebufferCreateMultipass(gpu.textures(textures), gpu.framebufferPasses(passes), Int(validate_with_format), Int(view_count))
	return mmm.New[rd.Framebuffer, rd.API, uintptr](span, gpu.value, uintptr(framebuffer))
}

func (gpu rdInterface) FramebufferGetFormat(fb rd.Framebuffer) int64 {
	return int64(gpu.class.FramebufferGetFormat(RID(mmm.Get(fb))))
}

func (gpu rdInterface) FramebufferIsValid(fb rd.Framebuffer) bool {
	return bool(gpu.class.FramebufferIsValid(RID(mmm.Get(fb))))
}

func (gpu rdInterface) FramebufferFormatCreate(attachments ffi.Managed[[]rd.AttachmentFormat], view_count int) int64 {
	return int64(gpu.class.FramebufferFormatCreate(gpu.attachments(attachments), Int(view_count)))
}

func (gpu rdInterface) FramebufferFormatCreateEmpty(samples rd.TextureSamples) int64 {
	return int64(gpu.class.FramebufferFormatCreateEmpty(classdb.RenderingDeviceTextureSamples(samples)))
}

func (gpu rdInterface) FramebufferFormatCreateMultipass(attachments ffi.Managed[[]rd.AttachmentFormat], passes ffi.Managed[[]rd.FramebufferPass], view_count int) int64 {
	return int64(gpu.class.FramebufferFormatCreateMultipass(gpu.attachments(attachments), gpu.framebufferPasses(passes), Int(view_count)))
}

func (gpu rdInterface) FramebufferFormatGetTextureSamples(format1 int64, render_pass int) rd.TextureSamples {
	return rd.TextureSamples(gpu.class.FramebufferFormatGetTextureSamples(Int(format1), Int(render_pass)))
}

func (gpu rdInterface) GetDeviceName() string {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return gpu.class.GetDeviceName(tmp).String()
}

func (gpu rdInterface) GetDevicePipelineCacheUUID() string {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return gpu.class.GetDevicePipelineCacheUuid(tmp).String()
}

func (gpu rdInterface) GetDeviceVendorName() string {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return gpu.class.GetDeviceVendorName(tmp).String()
}

func (gpu rdInterface) GetDriverResource(resource rd.DriverResource, rid rd.RID, index int) int {
	return int(gpu.class.GetDriverResource(classdb.RenderingDeviceDriverResource(resource), RID(rid), Int(index)))
}

func (gpu rdInterface) GetFrameDelay() int {
	return int(gpu.class.GetFrameDelay())
}

func (gpu rdInterface) GetMemoryUsage(mtype rd.MemoryType) int {
	return int(gpu.class.GetMemoryUsage(classdb.RenderingDeviceMemoryType(mtype)))
}

func (gpu rdInterface) IndexArrayCreate(span mmm.Lifetime, buf rd.IndexBuffer, offset, count int) rd.IndexArray {
	array := gpu.class.IndexArrayCreate(RID(mmm.Get(buf)), Int(offset), Int(count))
	return mmm.New[rd.IndexArray, rd.API, uintptr](span, gpu.value, uintptr(array))
}

func (gpu rdInterface) IndexBufferCreate(span mmm.Lifetime, size_indices int, format rd.IndexBufferFormat, data ffi.Bytes, use_restart_indicies bool) rd.IndexBuffer {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	buffer := gpu.class.IndexBufferCreate(Int(size_indices), classdb.RenderingDeviceIndexBufferFormat(format), gpu.bytes(tmp, data), Bool(use_restart_indicies))
	return mmm.New[rd.IndexBuffer, rd.API, uintptr](span, gpu.value, uintptr(buffer))
}

func (gpu rdInterface) LimitGet(limit rd.Limit) int {
	return int(gpu.class.LimitGet(classdb.RenderingDeviceLimit(limit)))
}

func (gpu rdInterface) RenderPipelineCreate(span mmm.Lifetime, shader rd.Shader, framebuffer_format int64, vertex_format rd.VertexFormat,
	primitive rd.PrimitiveType, rasterization_state ffi.Managed[rd.Rasterization], multisample_state ffi.Managed[rd.Multisampling],
	stencil_state ffi.Managed[rd.DepthStencils], color_blend_state ffi.Managed[rd.ColorBlending], dynamic_state_flags rd.DynamicStates,
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
	return mmm.New[rd.RenderPipeline, rd.API, uintptr](span, gpu.value, uintptr(pipeline))
}

func (gpu rdInterface) RenderPipelineIsValid(pipeline rd.RenderPipeline) bool {
	return bool(gpu.class.RenderPipelineIsValid(RID(mmm.Get(pipeline))))
}

func (gpu rdInterface) SamplerCreate(span mmm.Lifetime, state ffi.Managed[rd.SamplerState]) rd.Sampler {
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
	return mmm.New[rd.Sampler, rd.API, uintptr](span, gpu.value, uintptr(sampler))
}

func (gpu rdInterface) SamplerIsFormatSupportedForFilter(format rd.DataFormat, filter rd.SamplerFilter) bool {
	return bool(gpu.class.SamplerIsFormatSupportedForFilter(classdb.RenderingDeviceDataFormat(format), classdb.RenderingDeviceSamplerFilter(filter)))
}

func (gpu rdInterface) ScreenGetFramebufferFormat() int64 {
	return int64(gpu.class.ScreenGetFramebufferFormat())
}

func (gpu rdInterface) ScreenGetHeight(screen rd.Screen) int64 {
	return int64(gpu.class.ScreenGetHeight(Int(screen)))
}

func (gpu rdInterface) ScreenGetWidth(screen rd.Screen) int64 {
	return int64(gpu.class.ScreenGetWidth(Int(screen)))
}

func (gpu rdInterface) SetResourceName(resource rd.RID, name string) {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	gpu.class.SetResourceName(RID(resource), tmp.String(name))
}

func (gpu rdInterface) ShaderCompileBinaryFromSPIRV(span mmm.Lifetime, shader ffi.Managed[rd.ShaderSPIRV], name string) ffi.Bytes {
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

func (gpu rdInterface) ShaderCompileSourceIntoSPIRV(span mmm.Lifetime, shader ffi.Managed[rd.ShaderSource], allow_cache bool) rd.ShaderSPIRV {
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

func (gpu rdInterface) ShaderCreateFromBytecode(span mmm.Lifetime, code ffi.Bytes, id rd.ShaderPlaceholder) rd.Shader {
	shader := gpu.class.ShaderCreateFromBytecode(gpu.bytes(gpu.scope(span), code), RID(mmm.Get(id)))
	return mmm.New[rd.Shader, rd.API, uintptr](span, gpu.value, uintptr(shader))
}

func (gpu rdInterface) ShaderCreateFromSPIRV(span mmm.Lifetime, shader ffi.Managed[rd.ShaderSPIRV], name string) rd.Shader {
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
	return mmm.New[rd.Shader, rd.API, uintptr](span, gpu.value, uintptr(result))
}

func (gpu rdInterface) ShaderCreatePlaceholder(span mmm.Lifetime) rd.ShaderPlaceholder {
	placeholder := gpu.class.ShaderCreatePlaceholder()
	return mmm.New[rd.ShaderPlaceholder, rd.API, uintptr](span, gpu.value, uintptr(placeholder))
}

func (gpu rdInterface) ShaderGetVertexInputAttributeMask(shader rd.Shader) int64 {
	return int64(gpu.class.ShaderGetVertexInputAttributeMask(RID(mmm.Get(shader))))
}

func (gpu rdInterface) StorageBufferCreate(span mmm.Lifetime, size_bytes int, data ffi.Bytes, usage_flags rd.StorageBufferUsage) rd.StorageBuffer {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	buffer := gpu.class.StorageBufferCreate(Int(size_bytes), gpu.bytes(tmp, data), classdb.RenderingDeviceStorageBufferUsage(usage_flags))
	return mmm.New[rd.StorageBuffer, rd.API, uintptr](span, gpu.value, uintptr(buffer))
}

func (gpu rdInterface) Submit() {
	gpu.class.Submit()
}

func (gpu rdInterface) Sync() {
	gpu.class.Sync()
}

func (gpu rdInterface) TextureBufferCreate(span mmm.Lifetime, size_bytes int, data_format rd.DataFormat, data ffi.Bytes) rd.TextureBuffer {
	buffer := gpu.class.TextureBufferCreate(Int(size_bytes), classdb.RenderingDeviceDataFormat(data_format), gpu.bytes(gpu.scope(span), data))
	return mmm.New[rd.TextureBuffer, rd.API, uintptr](span, gpu.value, uintptr(buffer))
}

func (gpu rdInterface) TextureClear(texture rd.Texture, color uc.Color, base_mipmap, mipmap_count, base_layer, layer_count int, post_barrier rd.Barrier) error {
	return Error(gpu.class.TextureClear(RID(mmm.Get(texture)), color, Int(base_mipmap), Int(mipmap_count), Int(base_layer), Int(layer_count), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu rdInterface) TextureCopy(src, dst rd.Texture, src_pos, dst_pos, size xy.Vector3, src_mipmap, dst_mipmap, src_layer, dst_layer int, post_barrier rd.Barrier) error {
	return Error(gpu.class.TextureCopy(RID(mmm.Get(src)), RID(mmm.Get(dst)), Vector3(src_pos), Vector3(dst_pos), Vector3(size), Int(src_mipmap), Int(dst_mipmap), Int(src_layer), Int(dst_layer), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu rdInterface) TextureCreate(span mmm.Lifetime, format ffi.Managed[rd.TextureFormat], view ffi.Managed[rd.TextureView], data ffi.Managed[[]ffi.Bytes]) rd.Texture {
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
	return mmm.New[rd.Texture, rd.API, uintptr](span, gpu.value, uintptr(texture))
}

func (gpu rdInterface) TextureCreateFromExtension(span mmm.Lifetime,
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
	return mmm.New[rd.Texture, rd.API, uintptr](span, gpu.value, uintptr(texture))
}

func (gpu rdInterface) TextureCreateShared(span mmm.Lifetime, view ffi.Managed[rd.TextureView], with_texture rd.Texture) rd.Texture {
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
	return mmm.New[rd.Texture, rd.API, uintptr](span, gpu.value, uintptr(texture))
}

func (gpu rdInterface) TextureCreateSharedFromSlice(span mmm.Lifetime, view ffi.Managed[rd.TextureView], with_texture rd.Texture, layer, mipmap, mipmaps int, slice_type rd.TextureSliceType) rd.Texture {
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
	return mmm.New[rd.Texture, rd.API, uintptr](span, gpu.value, uintptr(texture))
}

func (gpu rdInterface) TextureGetData(span mmm.Lifetime, texture rd.Texture, layer int) ffi.Bytes {
	return ffi.Bytes{
		Interface: gpu.class.TextureGetData(gpu.scope(span), RID(mmm.Get(texture)), Int(layer)),
	}
}

func (gpu rdInterface) TextureGetFormat(texture rd.Texture) rd.TextureFormat {
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

func (gpu rdInterface) TextureGetNativeHandle(texture rd.Texture) uint64 {
	return uint64(gpu.class.TextureGetNativeHandle(RID(mmm.Get(texture))))
}

func (gpu rdInterface) TextureIsFormatSupportedForUsage(format rd.DataFormat, usage rd.TextureUsage) bool {
	return bool(gpu.class.TextureIsFormatSupportedForUsage(classdb.RenderingDeviceDataFormat(format), classdb.RenderingDeviceTextureUsageBits(usage)))
}

func (gpu rdInterface) TextureIsShared(tex rd.Texture) bool {
	return bool(gpu.class.TextureIsShared(RID(mmm.Get(tex))))
}

func (gpu rdInterface) TextureIsValid(tex rd.Texture) bool {
	return bool(gpu.class.TextureIsValid(RID(mmm.Get(tex))))
}

func (gpu rdInterface) TextureResolveMultisample(src, dst rd.Texture, post_barrier rd.Barrier) error {
	return Error(gpu.class.TextureResolveMultisample(RID(mmm.Get(src)), RID(mmm.Get(dst)), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu rdInterface) TextureUpdate(texture rd.Texture, layer int, data ffi.Bytes, post_barrier rd.Barrier) error {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	return Error(gpu.class.TextureUpdate(RID(mmm.Get(texture)), Int(layer), gpu.bytes(tmp, data), classdb.RenderingDeviceBarrierMask(post_barrier)))
}

func (gpu rdInterface) UniformBufferCreate(span mmm.Lifetime, size_bytes int, data ffi.Bytes) rd.UniformBuffer {
	tmp := gd.NewContext(gpu.godot)
	defer tmp.End()
	buffer := gpu.class.UniformBufferCreate(Int(size_bytes), gpu.bytes(tmp, data))
	return mmm.New[rd.UniformBuffer, rd.API, uintptr](span, gpu.value, uintptr(buffer))
}

func (gpu rdInterface) UniformSetCreate(span mmm.Lifetime, uniforms ffi.Managed[[]rd.Uniform], shader rd.Shader, shader_set int) rd.UniformSet {
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
	return mmm.New[rd.UniformSet, rd.API, uintptr](span, gpu.value, uintptr(set))
}

func (gpu rdInterface) UniformSetIsValid(set rd.UniformSet) bool {
	return bool(gpu.class.UniformSetIsValid(RID(mmm.Get(set))))
}

func (gpu rdInterface) VertexArrayCreate(span mmm.Lifetime, vertex_count int, vertex_format rd.VertexFormat, src_buffers ffi.Managed[[]rd.Buffer], offsets ffi.Slice[int64]) rd.VertexArray {
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
	return mmm.New[rd.VertexArray, rd.API, uintptr](span, gpu.value, uintptr(array))
}

func (gpu rdInterface) VertexFormatCreate(span mmm.Lifetime, vertex_descriptions ffi.Managed[[]rd.VertexAttribute]) rd.VertexFormat {
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

func (gpu rdInterface) FreeRID(rid uintptr) {
	gpu.class.FreeRid(RID(rid))
}*/
