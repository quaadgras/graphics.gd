/*
var fv = FontVariation.new()
fv.base_font = load("res://RobotoFlex.ttf")
var variation_list = fv.get_supported_variation_list()
for tag in variation_list:
	var name = TextServerManager.get_primary_interface().tag_to_name(tag)
	var values = variation_list[tag]
	print("variation axis: %s (%d)\n\tmin, max, default: %s" % [name, tag, values])
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Font"
	"graphics.gd/classdb/FontVariation"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/TextServerManager"
)

func Font_GetSupportedVariationList() {
	var fv = FontVariation.New()
	fv.AsFontVariation().SetBaseFont(Resource.Load[Font.Instance]("res://RobotoFlex.ttf"))
	var variation_list = fv.AsFont().GetSupportedVariationList()
	for tag := range variation_list {
		var name = TextServerManager.GetPrimaryInterface().TagToName(tag)
		var values = variation_list[tag]
		fmt.Printf("variation axis: %s (%d)\n\tmin, max, default: %v\n", name, tag, values)
	}
}
