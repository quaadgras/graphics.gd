/*
[gdscript]
func _ready():
    await RenderingServer.frame_post_draw
    $Viewport.get_texture().get_image().save_png("user://Screenshot.png")
[/gdscript]
[csharp]
public async override void _Ready()
{
    await ToSignal(RenderingServer.Singleton, RenderingServer.SignalName.FramePostDraw);
    var viewport = GetNode<Viewport>("Viewport");
    viewport.GetTexture().GetImage().SavePng("user://Screenshot.png");
}
[/csharp]
*/

package main

import "graphics.gd/classdb/RenderingServer"

func Viewport_GetTexture() {
	RenderingServer.OnFramePostDraw(func() {
		viewport.GetTexture().AsTexture2D().GetImage().SavePng("user://Screenshot.png")
	})
}
