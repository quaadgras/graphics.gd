package nativeui

import (
	"sync"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/StyleBox"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
)

// glassShader draws Apple's Liquid Glass material (macOS/iOS 26+): what is
// behind the shape, blurred, bent inwards towards the rounded edges like a
// lens, lightly tinted, with a specular rim lit from the top left and a soft
// shadow outside.
const glassShaderCode = `shader_type canvas_item;
render_mode unshaded;

uniform sampler2D screen : hint_screen_texture, filter_linear_mipmap;
uniform vec4 rect;          // the glass shape, in the canvas item's coordinates
uniform float radius;
uniform vec4 tint;          // mixed over the background by its alpha
uniform float blur = 1.0;   // mip level of the screen copy
uniform float bend = 12.0;  // how far (px) the background is pulled at the rim
uniform float rim = 0.45;   // specular rim brightness
uniform vec4 shadow = vec4(0.0, 0.0, 0.0, 0.14);
uniform float shadow_size = 6.0;

varying vec2 local;

void vertex() {
	local = VERTEX;
}

float sdf(vec2 p, vec2 hs, float r) {
	vec2 q = abs(p) - hs + r;
	return length(max(q, 0.0)) + min(max(q.x, q.y), 0.0) - r;
}

void fragment() {
	vec2 hs = rect.zw * 0.5;
	vec2 p = local - (rect.xy + hs);
	float r = min(radius, min(hs.x, hs.y));
	float d = sdf(p, hs, r);
	if (d > 0.0) {
		float s = 1.0 - smoothstep(0.0, shadow_size, sdf(p - vec2(0.0, 1.5), hs, r));
		COLOR = vec4(shadow.rgb, shadow.a * s * s);
	} else {
		vec2 e = vec2(0.5, 0.0);
		vec2 n = normalize(vec2(sdf(p + e.xy, hs, r) - sdf(p - e.xy, hs, r), sdf(p + e.yx, hs, r) - sdf(p - e.yx, hs, r)) + 1e-6);
		// The lens: zero in the middle, strongest at the edge.
		float depth = clamp(-d / max(r, 1.0), 0.0, 1.0);
		float lens = (1.0 - depth) * (1.0 - depth);
		vec2 uv = SCREEN_UV - n * lens * bend * SCREEN_PIXEL_SIZE;
		vec3 bg = textureLod(screen, uv, blur).rgb;
		vec3 col = mix(bg, tint.rgb, tint.a);
		float light = 0.55 + 0.45 * dot(n, normalize(vec2(-0.5, -0.85)));
		float edgeLine = 1.0 - smoothstep(0.0, 1.6, -d);
		col += rim * edgeLine * light + 0.06 * lens;
		COLOR = vec4(col, clamp(-d * 1.5, 0.0, 1.0));
	}
}
`

// glassBox is a StyleBox drawn in Liquid Glass. A StyleBox can't give the
// control a shader, so it draws into a child canvas item of the control's
// own (behind the control's other drawing) that has the glass material.
type glassBox struct {
	StyleBox.Extension[glassBox] `gd:"NativeUIGlassBox"`

	radius     Float.X
	tint       Color.RGBA
	shadow     Color.RGBA
	shadowSize Float.X

	mutex    sync.Mutex
	children map[RID.CanvasItem]*glassLayers
}

// glassLayers are the glass layers under one canvas item. A control can
// draw the same StyleBox several times (a TabBar draws every tab into its
// own canvas item), and each needs its own material for its shape, so the
// layers are reused in order, starting over each frame.
type glassLayers struct {
	frame int
	used  int
	list  []glassChild
}

type glassChild struct {
	item     RID.CanvasItem
	material RID.Material
}

var (
	registerGlass sync.Once
	glassShader   struct {
		sync.Mutex
		id    RID.Shader
		users int // live glassBoxes; the shader is freed with the last
	}
)

func newGlassBox(x box, scale float64) StyleBox.Instance {
	registerGlass.Do(func() { classdb.Register[glassBox]() })
	glassShader.Lock()
	if glassShader.users == 0 {
		glassShader.id = RenderingServer.ShaderCreate()
		RenderingServer.ShaderSetCode(glassShader.id, glassShaderCode)
	}
	glassShader.users++
	glassShader.Unlock()
	g := &glassBox{
		radius:     Float.X(float64(x.radius) * scale),
		tint:       x.bg,
		shadow:     x.shadow,
		shadowSize: Float.X(float64(x.shadowSize) * scale),
		children:   map[RID.CanvasItem]*glassLayers{},
	}
	sb := g.AsStyleBox()
	if x.setPad {
		sb.SetContentMarginLeft(Float.X(x.pad[0] * scale))
		sb.SetContentMarginTop(Float.X(x.pad[1] * scale))
		sb.SetContentMarginRight(Float.X(x.pad[2] * scale))
		sb.SetContentMarginBottom(Float.X(x.pad[3] * scale))
	}
	return sb
}

// notificationPredelete is Object.NOTIFICATION_PREDELETE.
const notificationPredelete Object.Notification = 1

// Notification frees the glass layers when the StyleBox goes (its Theme was
// freed, or the program is exiting). A StyleBox can't see when one of its
// controls is freed, so until then a freed control's layers stay allocated,
// orphaned and not drawn: a few RIDs per glass control the theme ever drew.
func (g *glassBox) Notification(what Object.Notification) {
	if what != notificationPredelete {
		return
	}
	g.mutex.Lock()
	defer g.mutex.Unlock()
	for _, layers := range g.children {
		for _, c := range layers.list {
			RenderingServer.FreeRid(RID.Any(c.item))
			RenderingServer.FreeRid(RID.Any(c.material))
		}
	}
	g.children = nil
	glassShader.Lock()
	if glassShader.users--; glassShader.users == 0 {
		RenderingServer.FreeRid(RID.Any(glassShader.id))
	}
	glassShader.Unlock()
}

func (g *glassBox) Draw(item RID.CanvasItem, rect Rect2.PositionSize) {
	g.mutex.Lock()
	if g.children == nil {
		g.children = map[RID.CanvasItem]*glassLayers{}
	}
	layers, ok := g.children[item]
	if !ok {
		layers = &glassLayers{frame: -1}
		g.children[item] = layers
	}
	if frame := Engine.GetFramesDrawn(); frame != layers.frame {
		layers.frame, layers.used = frame, 0
		for _, c := range layers.list {
			RenderingServer.CanvasItemClear(c.item)
		}
	}
	if layers.used == len(layers.list) {
		var c glassChild
		c.item = RenderingServer.CanvasItemCreate()
		c.material = RenderingServer.MaterialCreate()
		RenderingServer.MaterialSetShader(c.material, glassShader.id)
		RenderingServer.CanvasItemSetParent(c.item, item)
		RenderingServer.CanvasItemSetDrawBehindParent(c.item, true)
		RenderingServer.CanvasItemSetMaterial(c.item, c.material)
		layers.list = append(layers.list, c)
	}
	child := layers.list[layers.used]
	layers.used++
	g.mutex.Unlock()
	m := child.material
	RenderingServer.MaterialSetParam(m, "rect", Color.RGBA{R: rect.Position.X, G: rect.Position.Y, B: rect.Size.X, A: rect.Size.Y})
	RenderingServer.MaterialSetParam(m, "radius", g.radius)
	RenderingServer.MaterialSetParam(m, "tint", g.tint)
	RenderingServer.MaterialSetParam(m, "shadow", g.shadow)
	RenderingServer.MaterialSetParam(m, "shadow_size", max(g.shadowSize, 1))
	grow := g.shadowSize + 2
	area := Rect2.PositionSize{
		Position: Vector2.XY{X: rect.Position.X - grow, Y: rect.Position.Y - grow},
		Size:     Vector2.XY{X: rect.Size.X + 2*grow, Y: rect.Size.Y + 2*grow},
	}
	RenderingServer.CanvasItemAddRect(child.item, area, Color.RGBA{R: 1, G: 1, B: 1, A: 1}, false)
}
