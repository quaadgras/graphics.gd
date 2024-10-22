package Theme

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A resource used for styling/skinning [Control] and [Window] nodes. While individual controls can be styled using their local theme overrides (see [method Control.add_theme_color_override]), theme resources allow you to store and apply the same settings across all controls sharing the same type (e.g. style all [Button]s the same). One theme resource can be used for the entire project, but you can also set a separate theme resource to a branch of control nodes. A theme resource assigned to a control applies to the control itself, as well as all of its direct and indirect children (as long as a chain of controls is uninterrupted).
Use [member ProjectSettings.gui/theme/custom] to set up a project-scope theme that will be available to every control in your project.
Use [member Control.theme] of any control node to set up a theme that will be available to that control and all of its direct and indirect children.

*/
type Go [1]classdb.Theme

/*
Creates or changes the value of the icon property defined by [param name] and [param theme_type]. Use [method clear_icon] to remove the property.
*/
func (self Go) SetIcon(name string, theme_type string, texture gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIcon(gc.StringName(name), gc.StringName(theme_type), texture)
}

/*
Returns the icon property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback icon value if the property doesn't exist (see [member ThemeDB.fallback_icon]). Use [method has_icon] to check for existence.
*/
func (self Go) GetIcon(name string, theme_type string) gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Texture2D(class(self).GetIcon(gc, gc.StringName(name), gc.StringName(theme_type)))
}

/*
Returns [code]true[/code] if the icon property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_icon] to define it.
*/
func (self Go) HasIcon(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasIcon(gc.StringName(name), gc.StringName(theme_type)))
}

/*
Renames the icon property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_icon] to check for existence, and [method clear_icon] to remove the existing property.
*/
func (self Go) RenameIcon(old_name string, name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RenameIcon(gc.StringName(old_name), gc.StringName(name), gc.StringName(theme_type))
}

/*
Removes the icon property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_icon] to check for existence.
*/
func (self Go) ClearIcon(name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearIcon(gc.StringName(name), gc.StringName(theme_type))
}

/*
Returns a list of names for icon properties defined with [param theme_type]. Use [method get_icon_type_list] to get a list of possible theme type names.
*/
func (self Go) GetIconList(theme_type string) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetIconList(gc, gc.String(theme_type)).Strings(gc))
}

/*
Returns a list of all unique theme type names for icon properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Go) GetIconTypeList() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetIconTypeList(gc).Strings(gc))
}

/*
Creates or changes the value of the [StyleBox] property defined by [param name] and [param theme_type]. Use [method clear_stylebox] to remove the property.
*/
func (self Go) SetStylebox(name string, theme_type string, texture gdclass.StyleBox) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStylebox(gc.StringName(name), gc.StringName(theme_type), texture)
}

/*
Returns the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback stylebox value if the property doesn't exist (see [member ThemeDB.fallback_stylebox]). Use [method has_stylebox] to check for existence.
*/
func (self Go) GetStylebox(name string, theme_type string) gdclass.StyleBox {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.StyleBox(class(self).GetStylebox(gc, gc.StringName(name), gc.StringName(theme_type)))
}

/*
Returns [code]true[/code] if the [StyleBox] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_stylebox] to define it.
*/
func (self Go) HasStylebox(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasStylebox(gc.StringName(name), gc.StringName(theme_type)))
}

/*
Renames the [StyleBox] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_stylebox] to check for existence, and [method clear_stylebox] to remove the existing property.
*/
func (self Go) RenameStylebox(old_name string, name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RenameStylebox(gc.StringName(old_name), gc.StringName(name), gc.StringName(theme_type))
}

/*
Removes the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_stylebox] to check for existence.
*/
func (self Go) ClearStylebox(name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearStylebox(gc.StringName(name), gc.StringName(theme_type))
}

/*
Returns a list of names for [StyleBox] properties defined with [param theme_type]. Use [method get_stylebox_type_list] to get a list of possible theme type names.
*/
func (self Go) GetStyleboxList(theme_type string) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetStyleboxList(gc, gc.String(theme_type)).Strings(gc))
}

/*
Returns a list of all unique theme type names for [StyleBox] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Go) GetStyleboxTypeList() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetStyleboxTypeList(gc).Strings(gc))
}

/*
Creates or changes the value of the [Font] property defined by [param name] and [param theme_type]. Use [method clear_font] to remove the property.
*/
func (self Go) SetFont(name string, theme_type string, font gdclass.Font) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFont(gc.StringName(name), gc.StringName(theme_type), font)
}

/*
Returns the [Font] property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font if the property doesn't exist and the default theme font is set up (see [member default_font]). Use [method has_font] to check for existence of the property and [method has_default_font] to check for existence of the default theme font.
Returns the engine fallback font value, if neither exist (see [member ThemeDB.fallback_font]).
*/
func (self Go) GetFont(name string, theme_type string) gdclass.Font {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Font(class(self).GetFont(gc, gc.StringName(name), gc.StringName(theme_type)))
}

/*
Returns [code]true[/code] if the [Font] property defined by [param name] and [param theme_type] exists, or if the default theme font is set up (see [method has_default_font]).
Returns [code]false[/code] if neither exist. Use [method set_font] to define the property.
*/
func (self Go) HasFont(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasFont(gc.StringName(name), gc.StringName(theme_type)))
}

/*
Renames the [Font] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font] to check for existence, and [method clear_font] to remove the existing property.
*/
func (self Go) RenameFont(old_name string, name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RenameFont(gc.StringName(old_name), gc.StringName(name), gc.StringName(theme_type))
}

/*
Removes the [Font] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font] to check for existence.
*/
func (self Go) ClearFont(name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearFont(gc.StringName(name), gc.StringName(theme_type))
}

/*
Returns a list of names for [Font] properties defined with [param theme_type]. Use [method get_font_type_list] to get a list of possible theme type names.
*/
func (self Go) GetFontList(theme_type string) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetFontList(gc, gc.String(theme_type)).Strings(gc))
}

/*
Returns a list of all unique theme type names for [Font] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Go) GetFontTypeList() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetFontTypeList(gc).Strings(gc))
}

/*
Creates or changes the value of the font size property defined by [param name] and [param theme_type]. Use [method clear_font_size] to remove the property.
*/
func (self Go) SetFontSize(name string, theme_type string, font_size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFontSize(gc.StringName(name), gc.StringName(theme_type), gd.Int(font_size))
}

/*
Returns the font size property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font size if the property doesn't exist and the default theme font size is set up (see [member default_font_size]). Use [method has_font_size] to check for existence of the property and [method has_default_font_size] to check for existence of the default theme font.
Returns the engine fallback font size value, if neither exist (see [member ThemeDB.fallback_font_size]).
*/
func (self Go) GetFontSize(name string, theme_type string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFontSize(gc.StringName(name), gc.StringName(theme_type))))
}

/*
Returns [code]true[/code] if the font size property defined by [param name] and [param theme_type] exists, or if the default theme font size is set up (see [method has_default_font_size]).
Returns [code]false[/code] if neither exist. Use [method set_font_size] to define the property.
*/
func (self Go) HasFontSize(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasFontSize(gc.StringName(name), gc.StringName(theme_type)))
}

/*
Renames the font size property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font_size] to check for existence, and [method clear_font_size] to remove the existing property.
*/
func (self Go) RenameFontSize(old_name string, name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RenameFontSize(gc.StringName(old_name), gc.StringName(name), gc.StringName(theme_type))
}

/*
Removes the font size property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font_size] to check for existence.
*/
func (self Go) ClearFontSize(name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearFontSize(gc.StringName(name), gc.StringName(theme_type))
}

/*
Returns a list of names for font size properties defined with [param theme_type]. Use [method get_font_size_type_list] to get a list of possible theme type names.
*/
func (self Go) GetFontSizeList(theme_type string) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetFontSizeList(gc, gc.String(theme_type)).Strings(gc))
}

/*
Returns a list of all unique theme type names for font size properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Go) GetFontSizeTypeList() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetFontSizeTypeList(gc).Strings(gc))
}

/*
Creates or changes the value of the [Color] property defined by [param name] and [param theme_type]. Use [method clear_color] to remove the property.
*/
func (self Go) SetColor(name string, theme_type string, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColor(gc.StringName(name), gc.StringName(theme_type), color)
}

/*
Returns the [Color] property defined by [param name] and [param theme_type], if it exists.
Returns the default color value if the property doesn't exist. Use [method has_color] to check for existence.
*/
func (self Go) GetColor(name string, theme_type string) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetColor(gc.StringName(name), gc.StringName(theme_type)))
}

/*
Returns [code]true[/code] if the [Color] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_color] to define it.
*/
func (self Go) HasColor(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasColor(gc.StringName(name), gc.StringName(theme_type)))
}

/*
Renames the [Color] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_color] to check for existence, and [method clear_color] to remove the existing property.
*/
func (self Go) RenameColor(old_name string, name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RenameColor(gc.StringName(old_name), gc.StringName(name), gc.StringName(theme_type))
}

/*
Removes the [Color] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_color] to check for existence.
*/
func (self Go) ClearColor(name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearColor(gc.StringName(name), gc.StringName(theme_type))
}

/*
Returns a list of names for [Color] properties defined with [param theme_type]. Use [method get_color_type_list] to get a list of possible theme type names.
*/
func (self Go) GetColorList(theme_type string) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetColorList(gc, gc.String(theme_type)).Strings(gc))
}

/*
Returns a list of all unique theme type names for [Color] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Go) GetColorTypeList() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetColorTypeList(gc).Strings(gc))
}

/*
Creates or changes the value of the constant property defined by [param name] and [param theme_type]. Use [method clear_constant] to remove the property.
*/
func (self Go) SetConstant(name string, theme_type string, constant int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConstant(gc.StringName(name), gc.StringName(theme_type), gd.Int(constant))
}

/*
Returns the constant property defined by [param name] and [param theme_type], if it exists.
Returns [code]0[/code] if the property doesn't exist. Use [method has_constant] to check for existence.
*/
func (self Go) GetConstant(name string, theme_type string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetConstant(gc.StringName(name), gc.StringName(theme_type))))
}

/*
Returns [code]true[/code] if the constant property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_constant] to define it.
*/
func (self Go) HasConstant(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasConstant(gc.StringName(name), gc.StringName(theme_type)))
}

/*
Renames the constant property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_constant] to check for existence, and [method clear_constant] to remove the existing property.
*/
func (self Go) RenameConstant(old_name string, name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RenameConstant(gc.StringName(old_name), gc.StringName(name), gc.StringName(theme_type))
}

/*
Removes the constant property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_constant] to check for existence.
*/
func (self Go) ClearConstant(name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearConstant(gc.StringName(name), gc.StringName(theme_type))
}

/*
Returns a list of names for constant properties defined with [param theme_type]. Use [method get_constant_type_list] to get a list of possible theme type names.
*/
func (self Go) GetConstantList(theme_type string) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetConstantList(gc, gc.String(theme_type)).Strings(gc))
}

/*
Returns a list of all unique theme type names for constant properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Go) GetConstantTypeList() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetConstantTypeList(gc).Strings(gc))
}

/*
Returns [code]true[/code] if [member default_base_scale] has a valid value.
Returns [code]false[/code] if it doesn't. The value must be greater than [code]0.0[/code] to be considered valid.
*/
func (self Go) HasDefaultBaseScale() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasDefaultBaseScale())
}

/*
Returns [code]true[/code] if [member default_font] has a valid value.
Returns [code]false[/code] if it doesn't.
*/
func (self Go) HasDefaultFont() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasDefaultFont())
}

/*
Returns [code]true[/code] if [member default_font_size] has a valid value.
Returns [code]false[/code] if it doesn't. The value must be greater than [code]0[/code] to be considered valid.
*/
func (self Go) HasDefaultFontSize() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasDefaultFontSize())
}

/*
Creates or changes the value of the theme property of [param data_type] defined by [param name] and [param theme_type]. Use [method clear_theme_item] to remove the property.
Fails if the [param value] type is not accepted by [param data_type].
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Go) SetThemeItem(data_type classdb.ThemeDataType, name string, theme_type string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetThemeItem(data_type, gc.StringName(name), gc.StringName(theme_type), value)
}

/*
Returns the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback value if the property doesn't exist (see [ThemeDB]). Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Go) GetThemeItem(data_type classdb.ThemeDataType, name string, theme_type string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetThemeItem(gc, data_type, gc.StringName(name), gc.StringName(theme_type)))
}

/*
Returns [code]true[/code] if the theme property of [param data_type] defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_theme_item] to define it.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Go) HasThemeItem(data_type classdb.ThemeDataType, name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasThemeItem(data_type, gc.StringName(name), gc.StringName(theme_type)))
}

/*
Renames the theme property of [param data_type] defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_theme_item] to check for existence, and [method clear_theme_item] to remove the existing property.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Go) RenameThemeItem(data_type classdb.ThemeDataType, old_name string, name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RenameThemeItem(data_type, gc.StringName(old_name), gc.StringName(name), gc.StringName(theme_type))
}

/*
Removes the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Go) ClearThemeItem(data_type classdb.ThemeDataType, name string, theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearThemeItem(data_type, gc.StringName(name), gc.StringName(theme_type))
}

/*
Returns a list of names for properties of [param data_type] defined with [param theme_type]. Use [method get_theme_item_type_list] to get a list of possible theme type names.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Go) GetThemeItemList(data_type classdb.ThemeDataType, theme_type string) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetThemeItemList(gc, data_type, gc.String(theme_type)).Strings(gc))
}

/*
Returns a list of all unique theme type names for [param data_type] properties. Use [method get_type_list] to get a list of all unique theme types.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Go) GetThemeItemTypeList(data_type classdb.ThemeDataType) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetThemeItemTypeList(gc, data_type).Strings(gc))
}

/*
Marks [param theme_type] as a variation of [param base_type].
This adds [param theme_type] as a suggested option for [member Control.theme_type_variation] on a [Control] that is of the [param base_type] class.
Variations can also be nested, i.e. [param base_type] can be another variation. If a chain of variations ends with a [param base_type] matching the class of the [Control], the whole chain is going to be suggested as options.
[b]Note:[/b] Suggestions only show up if this theme resource is set as the project default theme. See [member ProjectSettings.gui/theme/custom].
*/
func (self Go) SetTypeVariation(theme_type string, base_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTypeVariation(gc.StringName(theme_type), gc.StringName(base_type))
}

/*
Returns [code]true[/code] if [param theme_type] is marked as a variation of [param base_type].
*/
func (self Go) IsTypeVariation(theme_type string, base_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsTypeVariation(gc.StringName(theme_type), gc.StringName(base_type)))
}

/*
Unmarks [param theme_type] as being a variation of another theme type. See [method set_type_variation].
*/
func (self Go) ClearTypeVariation(theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearTypeVariation(gc.StringName(theme_type))
}

/*
Returns the name of the base theme type if [param theme_type] is a valid variation type. Returns an empty string otherwise.
*/
func (self Go) GetTypeVariationBase(theme_type string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetTypeVariationBase(gc, gc.StringName(theme_type)).String())
}

/*
Returns a list of all type variations for the given [param base_type].
*/
func (self Go) GetTypeVariationList(base_type string) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetTypeVariationList(gc, gc.StringName(base_type)).Strings(gc))
}

/*
Adds an empty theme type for every valid data type.
[b]Note:[/b] Empty types are not saved with the theme. This method only exists to perform in-memory changes to the resource. Use available [code]set_*[/code] methods to add theme items.
*/
func (self Go) AddType(theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddType(gc.StringName(theme_type))
}

/*
Removes the theme type, gracefully discarding defined theme items. If the type is a variation, this information is also erased. If the type is a base for type variations, those variations lose their base.
*/
func (self Go) RemoveType(theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveType(gc.StringName(theme_type))
}

/*
Returns a list of all unique theme type names. Use the appropriate [code]get_*_type_list[/code] method to get a list of unique theme types for a single data type.
*/
func (self Go) GetTypeList() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetTypeList(gc).Strings(gc))
}

/*
Adds missing and overrides existing definitions with values from the [param other] theme resource.
[b]Note:[/b] This modifies the current theme. If you want to merge two themes together without modifying either one, create a new empty theme and merge the other two into it one after another.
*/
func (self Go) MergeWith(other gdclass.Theme) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MergeWith(other)
}

/*
Removes all the theme properties defined on the theme resource.
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Theme
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Theme"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) DefaultBaseScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDefaultBaseScale()))
}

func (self Go) SetDefaultBaseScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDefaultBaseScale(gd.Float(value))
}

func (self Go) DefaultFont() gdclass.Font {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Font(class(self).GetDefaultFont(gc))
}

func (self Go) SetDefaultFont(value gdclass.Font) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDefaultFont(value)
}

func (self Go) DefaultFontSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetDefaultFontSize()))
}

func (self Go) SetDefaultFontSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDefaultFontSize(gd.Int(value))
}

/*
Creates or changes the value of the icon property defined by [param name] and [param theme_type]. Use [method clear_icon] to remove the property.
*/
//go:nosplit
func (self class) SetIcon(name gd.StringName, theme_type gd.StringName, texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the icon property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback icon value if the property doesn't exist (see [member ThemeDB.fallback_icon]). Use [method has_icon] to check for existence.
*/
//go:nosplit
func (self class) GetIcon(ctx gd.Lifetime, name gd.StringName, theme_type gd.StringName) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the icon property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_icon] to define it.
*/
//go:nosplit
func (self class) HasIcon(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_has_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Renames the icon property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_icon] to check for existence, and [method clear_icon] to remove the existing property.
*/
//go:nosplit
func (self class) RenameIcon(old_name gd.StringName, name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(old_name))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_rename_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the icon property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_icon] to check for existence.
*/
//go:nosplit
func (self class) ClearIcon(name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_clear_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of names for icon properties defined with [param theme_type]. Use [method get_icon_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetIconList(ctx gd.Lifetime, theme_type gd.String) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_icon_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of all unique theme type names for icon properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetIconTypeList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_icon_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates or changes the value of the [StyleBox] property defined by [param name] and [param theme_type]. Use [method clear_stylebox] to remove the property.
*/
//go:nosplit
func (self class) SetStylebox(name gd.StringName, theme_type gd.StringName, texture gdclass.StyleBox)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback stylebox value if the property doesn't exist (see [member ThemeDB.fallback_stylebox]). Use [method has_stylebox] to check for existence.
*/
//go:nosplit
func (self class) GetStylebox(ctx gd.Lifetime, name gd.StringName, theme_type gd.StringName) gdclass.StyleBox {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.StyleBox
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [StyleBox] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_stylebox] to define it.
*/
//go:nosplit
func (self class) HasStylebox(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_has_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Renames the [StyleBox] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_stylebox] to check for existence, and [method clear_stylebox] to remove the existing property.
*/
//go:nosplit
func (self class) RenameStylebox(old_name gd.StringName, name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(old_name))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_rename_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_stylebox] to check for existence.
*/
//go:nosplit
func (self class) ClearStylebox(name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_clear_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of names for [StyleBox] properties defined with [param theme_type]. Use [method get_stylebox_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetStyleboxList(ctx gd.Lifetime, theme_type gd.String) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_stylebox_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of all unique theme type names for [StyleBox] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetStyleboxTypeList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_stylebox_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates or changes the value of the [Font] property defined by [param name] and [param theme_type]. Use [method clear_font] to remove the property.
*/
//go:nosplit
func (self class) SetFont(name gd.StringName, theme_type gd.StringName, font gdclass.Font)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Font] property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font if the property doesn't exist and the default theme font is set up (see [member default_font]). Use [method has_font] to check for existence of the property and [method has_default_font] to check for existence of the default theme font.
Returns the engine fallback font value, if neither exist (see [member ThemeDB.fallback_font]).
*/
//go:nosplit
func (self class) GetFont(ctx gd.Lifetime, name gd.StringName, theme_type gd.StringName) gdclass.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [Font] property defined by [param name] and [param theme_type] exists, or if the default theme font is set up (see [method has_default_font]).
Returns [code]false[/code] if neither exist. Use [method set_font] to define the property.
*/
//go:nosplit
func (self class) HasFont(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_has_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Renames the [Font] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font] to check for existence, and [method clear_font] to remove the existing property.
*/
//go:nosplit
func (self class) RenameFont(old_name gd.StringName, name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(old_name))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_rename_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the [Font] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font] to check for existence.
*/
//go:nosplit
func (self class) ClearFont(name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_clear_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of names for [Font] properties defined with [param theme_type]. Use [method get_font_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetFontList(ctx gd.Lifetime, theme_type gd.String) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_font_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of all unique theme type names for [Font] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetFontTypeList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_font_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates or changes the value of the font size property defined by [param name] and [param theme_type]. Use [method clear_font_size] to remove the property.
*/
//go:nosplit
func (self class) SetFontSize(name gd.StringName, theme_type gd.StringName, font_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the font size property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font size if the property doesn't exist and the default theme font size is set up (see [member default_font_size]). Use [method has_font_size] to check for existence of the property and [method has_default_font_size] to check for existence of the default theme font.
Returns the engine fallback font size value, if neither exist (see [member ThemeDB.fallback_font_size]).
*/
//go:nosplit
func (self class) GetFontSize(name gd.StringName, theme_type gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the font size property defined by [param name] and [param theme_type] exists, or if the default theme font size is set up (see [method has_default_font_size]).
Returns [code]false[/code] if neither exist. Use [method set_font_size] to define the property.
*/
//go:nosplit
func (self class) HasFontSize(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_has_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Renames the font size property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font_size] to check for existence, and [method clear_font_size] to remove the existing property.
*/
//go:nosplit
func (self class) RenameFontSize(old_name gd.StringName, name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(old_name))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_rename_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the font size property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font_size] to check for existence.
*/
//go:nosplit
func (self class) ClearFontSize(name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_clear_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of names for font size properties defined with [param theme_type]. Use [method get_font_size_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetFontSizeList(ctx gd.Lifetime, theme_type gd.String) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_font_size_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of all unique theme type names for font size properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetFontSizeTypeList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_font_size_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates or changes the value of the [Color] property defined by [param name] and [param theme_type]. Use [method clear_color] to remove the property.
*/
//go:nosplit
func (self class) SetColor(name gd.StringName, theme_type gd.StringName, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Color] property defined by [param name] and [param theme_type], if it exists.
Returns the default color value if the property doesn't exist. Use [method has_color] to check for existence.
*/
//go:nosplit
func (self class) GetColor(name gd.StringName, theme_type gd.StringName) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [Color] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_color] to define it.
*/
//go:nosplit
func (self class) HasColor(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_has_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Renames the [Color] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_color] to check for existence, and [method clear_color] to remove the existing property.
*/
//go:nosplit
func (self class) RenameColor(old_name gd.StringName, name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(old_name))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_rename_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the [Color] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_color] to check for existence.
*/
//go:nosplit
func (self class) ClearColor(name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_clear_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of names for [Color] properties defined with [param theme_type]. Use [method get_color_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetColorList(ctx gd.Lifetime, theme_type gd.String) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_color_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of all unique theme type names for [Color] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetColorTypeList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_color_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates or changes the value of the constant property defined by [param name] and [param theme_type]. Use [method clear_constant] to remove the property.
*/
//go:nosplit
func (self class) SetConstant(name gd.StringName, theme_type gd.StringName, constant gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	callframe.Arg(frame, constant)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the constant property defined by [param name] and [param theme_type], if it exists.
Returns [code]0[/code] if the property doesn't exist. Use [method has_constant] to check for existence.
*/
//go:nosplit
func (self class) GetConstant(name gd.StringName, theme_type gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the constant property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_constant] to define it.
*/
//go:nosplit
func (self class) HasConstant(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_has_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Renames the constant property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_constant] to check for existence, and [method clear_constant] to remove the existing property.
*/
//go:nosplit
func (self class) RenameConstant(old_name gd.StringName, name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(old_name))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_rename_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the constant property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_constant] to check for existence.
*/
//go:nosplit
func (self class) ClearConstant(name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_clear_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of names for constant properties defined with [param theme_type]. Use [method get_constant_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetConstantList(ctx gd.Lifetime, theme_type gd.String) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_constant_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of all unique theme type names for constant properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetConstantTypeList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_constant_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultBaseScale(base_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, base_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultBaseScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [member default_base_scale] has a valid value.
Returns [code]false[/code] if it doesn't. The value must be greater than [code]0.0[/code] to be considered valid.
*/
//go:nosplit
func (self class) HasDefaultBaseScale() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_has_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultFont(font gdclass.Font)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_default_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultFont(ctx gd.Lifetime) gdclass.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_default_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [member default_font] has a valid value.
Returns [code]false[/code] if it doesn't.
*/
//go:nosplit
func (self class) HasDefaultFont() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_has_default_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultFontSize(font_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_default_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultFontSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_default_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [member default_font_size] has a valid value.
Returns [code]false[/code] if it doesn't. The value must be greater than [code]0[/code] to be considered valid.
*/
//go:nosplit
func (self class) HasDefaultFontSize() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_has_default_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates or changes the value of the theme property of [param data_type] defined by [param name] and [param theme_type]. Use [method clear_theme_item] to remove the property.
Fails if the [param value] type is not accepted by [param data_type].
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) SetThemeItem(data_type classdb.ThemeDataType, name gd.StringName, theme_type gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_theme_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback value if the property doesn't exist (see [ThemeDB]). Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) GetThemeItem(ctx gd.Lifetime, data_type classdb.ThemeDataType, name gd.StringName, theme_type gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_theme_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the theme property of [param data_type] defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_theme_item] to define it.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) HasThemeItem(data_type classdb.ThemeDataType, name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_has_theme_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Renames the theme property of [param data_type] defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_theme_item] to check for existence, and [method clear_theme_item] to remove the existing property.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) RenameThemeItem(data_type classdb.ThemeDataType, old_name gd.StringName, name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, mmm.Get(old_name))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_rename_theme_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) ClearThemeItem(data_type classdb.ThemeDataType, name gd.StringName, theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_clear_theme_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of names for properties of [param data_type] defined with [param theme_type]. Use [method get_theme_item_type_list] to get a list of possible theme type names.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) GetThemeItemList(ctx gd.Lifetime, data_type classdb.ThemeDataType, theme_type gd.String) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_theme_item_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of all unique theme type names for [param data_type] properties. Use [method get_type_list] to get a list of all unique theme types.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) GetThemeItemTypeList(ctx gd.Lifetime, data_type classdb.ThemeDataType) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_theme_item_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Marks [param theme_type] as a variation of [param base_type].
This adds [param theme_type] as a suggested option for [member Control.theme_type_variation] on a [Control] that is of the [param base_type] class.
Variations can also be nested, i.e. [param base_type] can be another variation. If a chain of variations ends with a [param base_type] matching the class of the [Control], the whole chain is going to be suggested as options.
[b]Note:[/b] Suggestions only show up if this theme resource is set as the project default theme. See [member ProjectSettings.gui/theme/custom].
*/
//go:nosplit
func (self class) SetTypeVariation(theme_type gd.StringName, base_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	callframe.Arg(frame, mmm.Get(base_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_set_type_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if [param theme_type] is marked as a variation of [param base_type].
*/
//go:nosplit
func (self class) IsTypeVariation(theme_type gd.StringName, base_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	callframe.Arg(frame, mmm.Get(base_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_is_type_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Unmarks [param theme_type] as being a variation of another theme type. See [method set_type_variation].
*/
//go:nosplit
func (self class) ClearTypeVariation(theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_clear_type_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the base theme type if [param theme_type] is a valid variation type. Returns an empty string otherwise.
*/
//go:nosplit
func (self class) GetTypeVariationBase(ctx gd.Lifetime, theme_type gd.StringName) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_type_variation_base, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of all type variations for the given [param base_type].
*/
//go:nosplit
func (self class) GetTypeVariationList(ctx gd.Lifetime, base_type gd.StringName) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(base_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_type_variation_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds an empty theme type for every valid data type.
[b]Note:[/b] Empty types are not saved with the theme. This method only exists to perform in-memory changes to the resource. Use available [code]set_*[/code] methods to add theme items.
*/
//go:nosplit
func (self class) AddType(theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_add_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the theme type, gracefully discarding defined theme items. If the type is a variation, this information is also erased. If the type is a base for type variations, those variations lose their base.
*/
//go:nosplit
func (self class) RemoveType(theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_remove_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of all unique theme type names. Use the appropriate [code]get_*_type_list[/code] method to get a list of unique theme types for a single data type.
*/
//go:nosplit
func (self class) GetTypeList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_get_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds missing and overrides existing definitions with values from the [param other] theme resource.
[b]Note:[/b] This modifies the current theme. If you want to merge two themes together without modifying either one, create a new empty theme and merge the other two into it one after another.
*/
//go:nosplit
func (self class) MergeWith(other gdclass.Theme)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(other[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_merge_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all the theme properties defined on the theme resource.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Theme.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsTheme() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTheme() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("Theme", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type DataType = classdb.ThemeDataType

const (
/*Theme's [Color] item type.*/
	DataTypeColor DataType = 0
/*Theme's constant item type.*/
	DataTypeConstant DataType = 1
/*Theme's [Font] item type.*/
	DataTypeFont DataType = 2
/*Theme's font size item type.*/
	DataTypeFontSize DataType = 3
/*Theme's icon [Texture2D] item type.*/
	DataTypeIcon DataType = 4
/*Theme's [StyleBox] item type.*/
	DataTypeStylebox DataType = 5
/*Maximum value for the DataType enum.*/
	DataTypeMax DataType = 6
)
