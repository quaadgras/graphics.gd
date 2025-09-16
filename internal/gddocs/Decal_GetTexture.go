/*
[gdscript]
for i in Decal.TEXTURE_MAX:
	$NewDecal.set_texture(i, $OldDecal.get_texture(i))
[/gdscript]
[csharp]
for (int i = 0; i < (int)Decal.DecalTexture.Max; i++)
{
	GetNode<Decal>("NewDecal").SetTexture(i, GetNode<Decal>("OldDecal").GetTexture(i));
}
[/csharp]
*/

package main

import "graphics.gd/classdb/Decal"

func Decal_GetTexture() {
	for i := range Decal.TextureMax {
		Decal.Advanced(NewDecal).SetTexture(i, Decal.Advanced(OldDecal).GetTexture(i))
	}
}
