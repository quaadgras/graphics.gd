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
	tmp := NewLifetime(gpu.godot)
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
	tmp := NewLifetime(t.from.godot)
	defer tmp.End()
	for i := range t.from.class.GetCapturedTimestampsCount() {
		if t.from.class.GetCapturedTimestampName(tmp, i).String() == t.name {
			return time.Duration(t.from.class.GetCapturedTimestampCpuTime(i)) * time.Microsecond, true
		}
	}
	return 0, false
}

func (t rdTimestamp) GPU() (time.Duration, bool) {
	tmp := NewLifetime(t.from.godot)
	defer tmp.End()
	for i := range t.from.class.GetCapturedTimestampsCount() {
		if t.from.class.GetCapturedTimestampName(tmp, i).String() == t.name {
			return time.Duration(t.from.class.GetCapturedTimestampGpuTime(i)) * time.Microsecond, true
		}
	}
	return 0, false
}

func (gpu *rdInterface) CompileBinary(data []byte) rd.Shader {
	tmp := NewLifetime(gpu.godot)
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
	tmp := NewLifetime(res.from.godot)
	defer tmp.End()
	res.from.class.SetResourceName(res.rid(), tmp.String(name))
}

type rdShader struct {
	rdNameable
}

func (gpu rdShader) Compile(data []byte) {
	tmp := NewLifetime(gpu.from.godot)
	defer tmp.End()
	gpu.name = gpu.from.class.ShaderCreateFromBytecode(tmp.PackedByteSlice(data), gpu.rid())
}

func (gpu rdShader) VertexInputAttributeMask() uint32 {
	return uint32(gpu.from.class.ShaderGetVertexInputAttributeMask(gpu.rid()))
}

func (gpu rdShader) Variables(variables map[int]rd.Variable) rd.Variables {
	tmp := NewLifetime(gpu.from.godot)
	defer tmp.End()
	array := NewArrayOf[classdb.RDUniform](tmp)
	for binding, variable := range variables {
		var uniform = *Create(tmp, new(classdb.RDUniform))
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
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	already, ok := spirv.(rdSPIRV)
	if !ok {
		converted := *Create(tmp, new(classdb.RDShaderSPIRV))
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
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	converted := *Create(tmp, new(classdb.RDShaderSource))
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
	tmp := gd.NewLifetime(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return spirv.GetStageBytecode(tmp, RenderingDeviceShaderStageCompute).Bytes()
}

func (spirv rdSPIRV) Fragment() []byte {
	tmp := gd.NewLifetime(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return spirv.GetStageBytecode(tmp, RenderingDeviceShaderStageFragment).Bytes()
}

func (spirv rdSPIRV) TesselationControl() []byte {
	tmp := gd.NewLifetime(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return spirv.GetStageBytecode(tmp, RenderingDeviceShaderStageTesselationControl).Bytes()
}

func (spirv rdSPIRV) TesselationEvaluation() []byte {
	tmp := gd.NewLifetime(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return spirv.GetStageBytecode(tmp, RenderingDeviceShaderStageTesselationEvaluation).Bytes()
}

func (spirv rdSPIRV) Vertex() []byte {
	tmp := gd.NewLifetime(mmm.API(spirv.AsPointer()))
	defer tmp.End()
	return spirv.GetStageBytecode(tmp, RenderingDeviceShaderStageVertex).Bytes()
}

func (spirv rdSPIRV) Shader(name string) rd.Shader {
	tmp := gd.NewLifetime(mmm.API(spirv.AsPointer()))
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
	tmp := NewLifetime(c.from.godot)
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
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	return gpu.class.GetDeviceName(tmp).String()
}

func (gpu *rdInterface) DeviceVendor() string {
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	return gpu.class.GetDeviceVendorName(tmp).String()
}

func (gpu *rdInterface) Drawing(frame rd.Frame, fn func(rd.Drawing)) {
	intial_color_action, final_color_action := frame.Color()
	initial_depth_action, final_depth_action := frame.Depth()

	tmp := NewLifetime(gpu.godot)
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
	tmp := NewLifetime(tex.from.godot)
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
		tmp := NewLifetime(data.tex.from.godot)
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
	tmp := NewLifetime(data.tex.from.godot)
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
	tmp := NewLifetime(d.from.godot)
	defer tmp.End()
	d.from.class.DrawCommandBeginLabel(tmp.String(name), color)
	block()
	d.from.class.DrawCommandEndLabel()
}

func (d *rdDrawing) DebugLabel(name string, color uc.Color) {
	tmp := NewLifetime(d.from.godot)
	defer tmp.End()
	d.from.class.DrawCommandInsertLabel(tmp.String(name), color)
}

func (d *rdDrawing) SetBlendConstant(color uc.Color) {
	d.from.class.DrawListSetBlendConstants(d.list, color)
}

func (d *rdDrawing) SetData(data []byte) {
	tmp := NewLifetime(d.from.godot)
	defer tmp.End()
	d.from.class.DrawListSetPushConstant(d.list, tmp.PackedByteSlice(data), Int(len(data))) // FIXME can we avoid the copy here?
}

func (d *rdDrawing) SetIndexArray(array rd.IndexArray) {
	tmp := NewLifetime(d.from.godot)
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
	tmp := NewLifetime(d.from.godot)
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
	tmp := NewLifetime(fmt.from.godot)
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
		var converted = *Create(tmp, new(classdb.RDFramebufferPass))
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
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	array1 := NewArrayOf[classdb.RDAttachmentFormat](tmp)
	for _, attachment := range attachments {
		var converted = *Create(tmp, new(classdb.RDAttachmentFormat))
		converted.SetFormat(RenderingDeviceDataFormat(attachment.Format))
		converted.SetSamples(RenderingDeviceTextureSamples(attachment.Samples))
		converted.SetUsageFlags(int64(attachment.Usage))
		array1.Append(converted)
	}
	array2 := NewArrayOf[classdb.RDFramebufferPass](tmp)
	for _, pass := range passes {
		var converted = *Create(tmp, new(classdb.RDFramebufferPass))
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
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	bytes := unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), uintptr(len(data))*unsafe.Sizeof(uint16(0)))
	return rdIndexBuffer{rdBuffer{rdNameable{gpu.resource(gpu.class.IndexBufferCreate(Int(len(bytes)), RenderingDeviceIndexBufferFormatUint16, tmp.PackedByteSlice(bytes), false))}}}
}

func (gpu *rdInterface) IndexBufferU32(data []uint32) rd.IndexBuffer {
	tmp := NewLifetime(gpu.godot)
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
	tmp := NewLifetime(buf.from.godot)
	defer tmp.End()
	data := buf.from.class.BufferGetData(tmp, buf.rid(), Int(off), Int(len(b)))
	copy(b, unsafe.Slice((*byte)(data.UnsafePointer()), data.Size()))
	return int(data.Size()), nil
}

func (buf rdBuffer) WriteAt(b []byte, off int64) (n int, err error) {
	tmp := NewLifetime(buf.from.godot)
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
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	return gpu.class.GetDevicePipelineCacheUuid(tmp).String()
}

func (gpu rdInterface) Processor(shader rd.Shader, defines []any) rd.Processor {
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	array := NewArrayOf[classdb.RDPipelineSpecializationConstant](tmp)
	for i, define := range defines {
		var converted = *Create(tmp, new(classdb.RDPipelineSpecializationConstant))
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
	tmp := NewLifetime(gpu.godot)
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
		var converted = *Create(tmp, new(classdb.RDPipelineColorBlendStateAttachment))
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
		var converted = *Create(tmp, new(classdb.RDPipelineSpecializationConstant))
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
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	converted := *Create(tmp, new(classdb.RDSamplerState))
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
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()

	converted := *Create(tmp, new(classdb.RDTextureView))
	converted.SetFormatOverride(RenderingDeviceDataFormat(view.FormatOverride))
	converted.SetSwizzleA(RenderingDeviceTextureSwizzle(view.SwizzleAlpha))
	converted.SetSwizzleB(RenderingDeviceTextureSwizzle(view.SwizzleBlue))
	converted.SetSwizzleG(RenderingDeviceTextureSwizzle(view.SwizzleGreen))
	converted.SetSwizzleR(RenderingDeviceTextureSwizzle(view.SwizzleRed))

	return &rdTexture{rdNameable: rdNameable{gpu.resource(gpu.class.TextureCreateShared(converted, with.(*rdTexture).rid()))}}
}

func (gpu *rdInterface) StorageBuffer(usage rd.StorageBufferUsage, data []byte) rd.StorageBuffer {
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	bytes := tmp.PackedByteSlice(data) // FIXME can we avoid the copy here?
	return rdStorageBuffer{rdBuffer: rdBuffer{rdNameable{gpu.resource(gpu.class.StorageBufferCreate(Int(len(data)), bytes, RenderingDeviceStorageBufferUsage(usage)))}}}
}

type rdStorageBuffer struct {
	rd.Variable
	rdBuffer
}

func (gpu *rdInterface) Texture(format rd.TextureFormat, view rd.TextureView, data [][]byte) rd.Texture {
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()

	texture_format := *Create(tmp, new(classdb.RDTextureFormat))
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

	converted := *Create(tmp, new(classdb.RDTextureView))
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
	tmp := NewLifetime(gpu.godot)
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
	tmp := NewLifetime(gpu.godot)
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
	tmp := NewLifetime(gpu.godot)
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
	tmp := NewLifetime(gpu.godot)
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
	tmp := NewLifetime(gpu.godot)
	defer tmp.End()
	array := NewArrayOf[classdb.RDVertexAttribute](tmp)
	for _, attribute := range attributes {
		var converted = *Create(tmp, new(classdb.RDVertexAttribute))
		converted.SetFormat(RenderingDeviceDataFormat(attribute.Format))
		converted.SetFrequency(RenderingDeviceVertexFrequency(attribute.Frequency))
		converted.SetLocation(Int(attribute.Location))
		converted.SetOffset(Int(attribute.Offset))
		converted.SetStride(Int(attribute.Stride))
		array.Append(converted)
	}
	return rd.VertexFormat(gpu.class.VertexFormatCreate(array))

}
