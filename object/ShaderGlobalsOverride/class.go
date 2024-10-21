package ShaderGlobalsOverride

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
Similar to how a [WorldEnvironment] node can be used to override the environment while a specific scene is loaded, [ShaderGlobalsOverride] can be used to override global shader parameters temporarily. Once the node is removed, the project-wide values for the global shader parameters are restored. See the [RenderingServer] [code]global_shader_parameter_*[/code] methods for more information.
[b]Note:[/b] Only one [ShaderGlobalsOverride] can be used per scene. If there is more than one [ShaderGlobalsOverride] node in the scene tree, only the first node (in tree order) will be taken into account.
[b]Note:[/b] All [ShaderGlobalsOverride] nodes are made part of a [code]"shader_overrides_group"[/code] group when they are added to the scene tree. The currently active [ShaderGlobalsOverride] node also has a [code]"shader_overrides_group_active"[/code] group added to it. You can use this to check which [ShaderGlobalsOverride] node is currently active.

*/
type Simple [1]classdb.ShaderGlobalsOverride
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ShaderGlobalsOverride
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsShaderGlobalsOverride() Expert { return self[0].AsShaderGlobalsOverride() }


//go:nosplit
func (self Simple) AsShaderGlobalsOverride() Simple { return self[0].AsShaderGlobalsOverride() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ShaderGlobalsOverride", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
