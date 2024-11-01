package Theme

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A resource used for styling/skinning [Control] and [Window] nodes. While individual controls can be styled using their local theme overrides (see [method Control.add_theme_color_override]), theme resources allow you to store and apply the same settings across all controls sharing the same type (e.g. style all [Button]s the same). One theme resource can be used for the entire project, but you can also set a separate theme resource to a branch of control nodes. A theme resource assigned to a control applies to the control itself, as well as all of its direct and indirect children (as long as a chain of controls is uninterrupted).
Use [member ProjectSettings.gui/theme/custom] to set up a project-scope theme that will be available to every control in your project.
Use [member Control.theme] of any control node to set up a theme that will be available to that control and all of its direct and indirect children.
*/
type Instance [1]classdb.Theme

/*
Creates or changes the value of the icon property defined by [param name] and [param theme_type]. Use [method clear_icon] to remove the property.
*/
func (self Instance) SetIcon(name string, theme_type string, texture objects.Texture2D) {
	class(self).SetIcon(gd.NewStringName(name), gd.NewStringName(theme_type), texture)
}

/*
Returns the icon property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback icon value if the property doesn't exist (see [member ThemeDB.fallback_icon]). Use [method has_icon] to check for existence.
*/
func (self Instance) GetIcon(name string, theme_type string) objects.Texture2D {
	return objects.Texture2D(class(self).GetIcon(gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Returns [code]true[/code] if the icon property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_icon] to define it.
*/
func (self Instance) HasIcon(name string, theme_type string) bool {
	return bool(class(self).HasIcon(gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Renames the icon property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_icon] to check for existence, and [method clear_icon] to remove the existing property.
*/
func (self Instance) RenameIcon(old_name string, name string, theme_type string) {
	class(self).RenameIcon(gd.NewStringName(old_name), gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Removes the icon property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_icon] to check for existence.
*/
func (self Instance) ClearIcon(name string, theme_type string) {
	class(self).ClearIcon(gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Returns a list of names for icon properties defined with [param theme_type]. Use [method get_icon_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetIconList(theme_type string) []string {
	return []string(class(self).GetIconList(gd.NewString(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for icon properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetIconTypeList() []string {
	return []string(class(self).GetIconTypeList().Strings())
}

/*
Creates or changes the value of the [StyleBox] property defined by [param name] and [param theme_type]. Use [method clear_stylebox] to remove the property.
*/
func (self Instance) SetStylebox(name string, theme_type string, texture objects.StyleBox) {
	class(self).SetStylebox(gd.NewStringName(name), gd.NewStringName(theme_type), texture)
}

/*
Returns the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback stylebox value if the property doesn't exist (see [member ThemeDB.fallback_stylebox]). Use [method has_stylebox] to check for existence.
*/
func (self Instance) GetStylebox(name string, theme_type string) objects.StyleBox {
	return objects.StyleBox(class(self).GetStylebox(gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Returns [code]true[/code] if the [StyleBox] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_stylebox] to define it.
*/
func (self Instance) HasStylebox(name string, theme_type string) bool {
	return bool(class(self).HasStylebox(gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Renames the [StyleBox] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_stylebox] to check for existence, and [method clear_stylebox] to remove the existing property.
*/
func (self Instance) RenameStylebox(old_name string, name string, theme_type string) {
	class(self).RenameStylebox(gd.NewStringName(old_name), gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Removes the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_stylebox] to check for existence.
*/
func (self Instance) ClearStylebox(name string, theme_type string) {
	class(self).ClearStylebox(gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Returns a list of names for [StyleBox] properties defined with [param theme_type]. Use [method get_stylebox_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetStyleboxList(theme_type string) []string {
	return []string(class(self).GetStyleboxList(gd.NewString(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for [StyleBox] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetStyleboxTypeList() []string {
	return []string(class(self).GetStyleboxTypeList().Strings())
}

/*
Creates or changes the value of the [Font] property defined by [param name] and [param theme_type]. Use [method clear_font] to remove the property.
*/
func (self Instance) SetFont(name string, theme_type string, font objects.Font) {
	class(self).SetFont(gd.NewStringName(name), gd.NewStringName(theme_type), font)
}

/*
Returns the [Font] property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font if the property doesn't exist and the default theme font is set up (see [member default_font]). Use [method has_font] to check for existence of the property and [method has_default_font] to check for existence of the default theme font.
Returns the engine fallback font value, if neither exist (see [member ThemeDB.fallback_font]).
*/
func (self Instance) GetFont(name string, theme_type string) objects.Font {
	return objects.Font(class(self).GetFont(gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Returns [code]true[/code] if the [Font] property defined by [param name] and [param theme_type] exists, or if the default theme font is set up (see [method has_default_font]).
Returns [code]false[/code] if neither exist. Use [method set_font] to define the property.
*/
func (self Instance) HasFont(name string, theme_type string) bool {
	return bool(class(self).HasFont(gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Renames the [Font] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font] to check for existence, and [method clear_font] to remove the existing property.
*/
func (self Instance) RenameFont(old_name string, name string, theme_type string) {
	class(self).RenameFont(gd.NewStringName(old_name), gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Removes the [Font] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font] to check for existence.
*/
func (self Instance) ClearFont(name string, theme_type string) {
	class(self).ClearFont(gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Returns a list of names for [Font] properties defined with [param theme_type]. Use [method get_font_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetFontList(theme_type string) []string {
	return []string(class(self).GetFontList(gd.NewString(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for [Font] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetFontTypeList() []string {
	return []string(class(self).GetFontTypeList().Strings())
}

/*
Creates or changes the value of the font size property defined by [param name] and [param theme_type]. Use [method clear_font_size] to remove the property.
*/
func (self Instance) SetFontSize(name string, theme_type string, font_size int) {
	class(self).SetFontSize(gd.NewStringName(name), gd.NewStringName(theme_type), gd.Int(font_size))
}

/*
Returns the font size property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font size if the property doesn't exist and the default theme font size is set up (see [member default_font_size]). Use [method has_font_size] to check for existence of the property and [method has_default_font_size] to check for existence of the default theme font.
Returns the engine fallback font size value, if neither exist (see [member ThemeDB.fallback_font_size]).
*/
func (self Instance) GetFontSize(name string, theme_type string) int {
	return int(int(class(self).GetFontSize(gd.NewStringName(name), gd.NewStringName(theme_type))))
}

/*
Returns [code]true[/code] if the font size property defined by [param name] and [param theme_type] exists, or if the default theme font size is set up (see [method has_default_font_size]).
Returns [code]false[/code] if neither exist. Use [method set_font_size] to define the property.
*/
func (self Instance) HasFontSize(name string, theme_type string) bool {
	return bool(class(self).HasFontSize(gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Renames the font size property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font_size] to check for existence, and [method clear_font_size] to remove the existing property.
*/
func (self Instance) RenameFontSize(old_name string, name string, theme_type string) {
	class(self).RenameFontSize(gd.NewStringName(old_name), gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Removes the font size property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font_size] to check for existence.
*/
func (self Instance) ClearFontSize(name string, theme_type string) {
	class(self).ClearFontSize(gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Returns a list of names for font size properties defined with [param theme_type]. Use [method get_font_size_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetFontSizeList(theme_type string) []string {
	return []string(class(self).GetFontSizeList(gd.NewString(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for font size properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetFontSizeTypeList() []string {
	return []string(class(self).GetFontSizeTypeList().Strings())
}

/*
Creates or changes the value of the [Color] property defined by [param name] and [param theme_type]. Use [method clear_color] to remove the property.
*/
func (self Instance) SetColor(name string, theme_type string, color gd.Color) {
	class(self).SetColor(gd.NewStringName(name), gd.NewStringName(theme_type), color)
}

/*
Returns the [Color] property defined by [param name] and [param theme_type], if it exists.
Returns the default color value if the property doesn't exist. Use [method has_color] to check for existence.
*/
func (self Instance) GetColor(name string, theme_type string) gd.Color {
	return gd.Color(class(self).GetColor(gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Returns [code]true[/code] if the [Color] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_color] to define it.
*/
func (self Instance) HasColor(name string, theme_type string) bool {
	return bool(class(self).HasColor(gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Renames the [Color] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_color] to check for existence, and [method clear_color] to remove the existing property.
*/
func (self Instance) RenameColor(old_name string, name string, theme_type string) {
	class(self).RenameColor(gd.NewStringName(old_name), gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Removes the [Color] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_color] to check for existence.
*/
func (self Instance) ClearColor(name string, theme_type string) {
	class(self).ClearColor(gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Returns a list of names for [Color] properties defined with [param theme_type]. Use [method get_color_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetColorList(theme_type string) []string {
	return []string(class(self).GetColorList(gd.NewString(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for [Color] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetColorTypeList() []string {
	return []string(class(self).GetColorTypeList().Strings())
}

/*
Creates or changes the value of the constant property defined by [param name] and [param theme_type]. Use [method clear_constant] to remove the property.
*/
func (self Instance) SetConstant(name string, theme_type string, constant int) {
	class(self).SetConstant(gd.NewStringName(name), gd.NewStringName(theme_type), gd.Int(constant))
}

/*
Returns the constant property defined by [param name] and [param theme_type], if it exists.
Returns [code]0[/code] if the property doesn't exist. Use [method has_constant] to check for existence.
*/
func (self Instance) GetConstant(name string, theme_type string) int {
	return int(int(class(self).GetConstant(gd.NewStringName(name), gd.NewStringName(theme_type))))
}

/*
Returns [code]true[/code] if the constant property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_constant] to define it.
*/
func (self Instance) HasConstant(name string, theme_type string) bool {
	return bool(class(self).HasConstant(gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Renames the constant property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_constant] to check for existence, and [method clear_constant] to remove the existing property.
*/
func (self Instance) RenameConstant(old_name string, name string, theme_type string) {
	class(self).RenameConstant(gd.NewStringName(old_name), gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Removes the constant property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_constant] to check for existence.
*/
func (self Instance) ClearConstant(name string, theme_type string) {
	class(self).ClearConstant(gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Returns a list of names for constant properties defined with [param theme_type]. Use [method get_constant_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetConstantList(theme_type string) []string {
	return []string(class(self).GetConstantList(gd.NewString(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for constant properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetConstantTypeList() []string {
	return []string(class(self).GetConstantTypeList().Strings())
}

/*
Returns [code]true[/code] if [member default_base_scale] has a valid value.
Returns [code]false[/code] if it doesn't. The value must be greater than [code]0.0[/code] to be considered valid.
*/
func (self Instance) HasDefaultBaseScale() bool {
	return bool(class(self).HasDefaultBaseScale())
}

/*
Returns [code]true[/code] if [member default_font] has a valid value.
Returns [code]false[/code] if it doesn't.
*/
func (self Instance) HasDefaultFont() bool {
	return bool(class(self).HasDefaultFont())
}

/*
Returns [code]true[/code] if [member default_font_size] has a valid value.
Returns [code]false[/code] if it doesn't. The value must be greater than [code]0[/code] to be considered valid.
*/
func (self Instance) HasDefaultFontSize() bool {
	return bool(class(self).HasDefaultFontSize())
}

/*
Creates or changes the value of the theme property of [param data_type] defined by [param name] and [param theme_type]. Use [method clear_theme_item] to remove the property.
Fails if the [param value] type is not accepted by [param data_type].
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) SetThemeItem(data_type classdb.ThemeDataType, name string, theme_type string, value gd.Variant) {
	class(self).SetThemeItem(data_type, gd.NewStringName(name), gd.NewStringName(theme_type), value)
}

/*
Returns the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback value if the property doesn't exist (see [ThemeDB]). Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) GetThemeItem(data_type classdb.ThemeDataType, name string, theme_type string) gd.Variant {
	return gd.Variant(class(self).GetThemeItem(data_type, gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Returns [code]true[/code] if the theme property of [param data_type] defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_theme_item] to define it.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) HasThemeItem(data_type classdb.ThemeDataType, name string, theme_type string) bool {
	return bool(class(self).HasThemeItem(data_type, gd.NewStringName(name), gd.NewStringName(theme_type)))
}

/*
Renames the theme property of [param data_type] defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_theme_item] to check for existence, and [method clear_theme_item] to remove the existing property.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) RenameThemeItem(data_type classdb.ThemeDataType, old_name string, name string, theme_type string) {
	class(self).RenameThemeItem(data_type, gd.NewStringName(old_name), gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Removes the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) ClearThemeItem(data_type classdb.ThemeDataType, name string, theme_type string) {
	class(self).ClearThemeItem(data_type, gd.NewStringName(name), gd.NewStringName(theme_type))
}

/*
Returns a list of names for properties of [param data_type] defined with [param theme_type]. Use [method get_theme_item_type_list] to get a list of possible theme type names.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) GetThemeItemList(data_type classdb.ThemeDataType, theme_type string) []string {
	return []string(class(self).GetThemeItemList(data_type, gd.NewString(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for [param data_type] properties. Use [method get_type_list] to get a list of all unique theme types.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) GetThemeItemTypeList(data_type classdb.ThemeDataType) []string {
	return []string(class(self).GetThemeItemTypeList(data_type).Strings())
}

/*
Marks [param theme_type] as a variation of [param base_type].
This adds [param theme_type] as a suggested option for [member Control.theme_type_variation] on a [Control] that is of the [param base_type] class.
Variations can also be nested, i.e. [param base_type] can be another variation. If a chain of variations ends with a [param base_type] matching the class of the [Control], the whole chain is going to be suggested as options.
[b]Note:[/b] Suggestions only show up if this theme resource is set as the project default theme. See [member ProjectSettings.gui/theme/custom].
*/
func (self Instance) SetTypeVariation(theme_type string, base_type string) {
	class(self).SetTypeVariation(gd.NewStringName(theme_type), gd.NewStringName(base_type))
}

/*
Returns [code]true[/code] if [param theme_type] is marked as a variation of [param base_type].
*/
func (self Instance) IsTypeVariation(theme_type string, base_type string) bool {
	return bool(class(self).IsTypeVariation(gd.NewStringName(theme_type), gd.NewStringName(base_type)))
}

/*
Unmarks [param theme_type] as being a variation of another theme type. See [method set_type_variation].
*/
func (self Instance) ClearTypeVariation(theme_type string) {
	class(self).ClearTypeVariation(gd.NewStringName(theme_type))
}

/*
Returns the name of the base theme type if [param theme_type] is a valid variation type. Returns an empty string otherwise.
*/
func (self Instance) GetTypeVariationBase(theme_type string) string {
	return string(class(self).GetTypeVariationBase(gd.NewStringName(theme_type)).String())
}

/*
Returns a list of all type variations for the given [param base_type].
*/
func (self Instance) GetTypeVariationList(base_type string) []string {
	return []string(class(self).GetTypeVariationList(gd.NewStringName(base_type)).Strings())
}

/*
Adds an empty theme type for every valid data type.
[b]Note:[/b] Empty types are not saved with the theme. This method only exists to perform in-memory changes to the resource. Use available [code]set_*[/code] methods to add theme items.
*/
func (self Instance) AddType(theme_type string) {
	class(self).AddType(gd.NewStringName(theme_type))
}

/*
Removes the theme type, gracefully discarding defined theme items. If the type is a variation, this information is also erased. If the type is a base for type variations, those variations lose their base.
*/
func (self Instance) RemoveType(theme_type string) {
	class(self).RemoveType(gd.NewStringName(theme_type))
}

/*
Returns a list of all unique theme type names. Use the appropriate [code]get_*_type_list[/code] method to get a list of unique theme types for a single data type.
*/
func (self Instance) GetTypeList() []string {
	return []string(class(self).GetTypeList().Strings())
}

/*
Adds missing and overrides existing definitions with values from the [param other] theme resource.
[b]Note:[/b] This modifies the current theme. If you want to merge two themes together without modifying either one, create a new empty theme and merge the other two into it one after another.
*/
func (self Instance) MergeWith(other objects.Theme) {
	class(self).MergeWith(other)
}

/*
Removes all the theme properties defined on the theme resource.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Theme

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Theme"))
	return Instance{classdb.Theme(object)}
}

func (self Instance) DefaultBaseScale() float64 {
	return float64(float64(class(self).GetDefaultBaseScale()))
}

func (self Instance) SetDefaultBaseScale(value float64) {
	class(self).SetDefaultBaseScale(gd.Float(value))
}

func (self Instance) DefaultFont() objects.Font {
	return objects.Font(class(self).GetDefaultFont())
}

func (self Instance) SetDefaultFont(value objects.Font) {
	class(self).SetDefaultFont(value)
}

func (self Instance) DefaultFontSize() int {
	return int(int(class(self).GetDefaultFontSize()))
}

func (self Instance) SetDefaultFontSize(value int) {
	class(self).SetDefaultFontSize(gd.Int(value))
}

/*
Creates or changes the value of the icon property defined by [param name] and [param theme_type]. Use [method clear_icon] to remove the property.
*/
//go:nosplit
func (self class) SetIcon(name gd.StringName, theme_type gd.StringName, texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the icon property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback icon value if the property doesn't exist (see [member ThemeDB.fallback_icon]). Use [method has_icon] to check for existence.
*/
//go:nosplit
func (self class) GetIcon(name gd.StringName, theme_type gd.StringName) objects.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the icon property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_icon] to define it.
*/
//go:nosplit
func (self class) HasIcon(name gd.StringName, theme_type gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the icon property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_icon] to check for existence, and [method clear_icon] to remove the existing property.
*/
//go:nosplit
func (self class) RenameIcon(old_name gd.StringName, name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(old_name))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the icon property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_icon] to check for existence.
*/
//go:nosplit
func (self class) ClearIcon(name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of names for icon properties defined with [param theme_type]. Use [method get_icon_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetIconList(theme_type gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_icon_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for icon properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetIconTypeList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_icon_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Creates or changes the value of the [StyleBox] property defined by [param name] and [param theme_type]. Use [method clear_stylebox] to remove the property.
*/
//go:nosplit
func (self class) SetStylebox(name gd.StringName, theme_type gd.StringName, texture objects.StyleBox) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback stylebox value if the property doesn't exist (see [member ThemeDB.fallback_stylebox]). Use [method has_stylebox] to check for existence.
*/
//go:nosplit
func (self class) GetStylebox(name gd.StringName, theme_type gd.StringName) objects.StyleBox {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.StyleBox{classdb.StyleBox(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [StyleBox] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_stylebox] to define it.
*/
//go:nosplit
func (self class) HasStylebox(name gd.StringName, theme_type gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the [StyleBox] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_stylebox] to check for existence, and [method clear_stylebox] to remove the existing property.
*/
//go:nosplit
func (self class) RenameStylebox(old_name gd.StringName, name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(old_name))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_stylebox] to check for existence.
*/
//go:nosplit
func (self class) ClearStylebox(name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of names for [StyleBox] properties defined with [param theme_type]. Use [method get_stylebox_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetStyleboxList(theme_type gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_stylebox_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for [StyleBox] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetStyleboxTypeList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_stylebox_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Creates or changes the value of the [Font] property defined by [param name] and [param theme_type]. Use [method clear_font] to remove the property.
*/
//go:nosplit
func (self class) SetFont(name gd.StringName, theme_type gd.StringName, font objects.Font) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Font] property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font if the property doesn't exist and the default theme font is set up (see [member default_font]). Use [method has_font] to check for existence of the property and [method has_default_font] to check for existence of the default theme font.
Returns the engine fallback font value, if neither exist (see [member ThemeDB.fallback_font]).
*/
//go:nosplit
func (self class) GetFont(name gd.StringName, theme_type gd.StringName) objects.Font {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Font{classdb.Font(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [Font] property defined by [param name] and [param theme_type] exists, or if the default theme font is set up (see [method has_default_font]).
Returns [code]false[/code] if neither exist. Use [method set_font] to define the property.
*/
//go:nosplit
func (self class) HasFont(name gd.StringName, theme_type gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the [Font] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font] to check for existence, and [method clear_font] to remove the existing property.
*/
//go:nosplit
func (self class) RenameFont(old_name gd.StringName, name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(old_name))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the [Font] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font] to check for existence.
*/
//go:nosplit
func (self class) ClearFont(name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of names for [Font] properties defined with [param theme_type]. Use [method get_font_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetFontList(theme_type gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for [Font] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetFontTypeList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Creates or changes the value of the font size property defined by [param name] and [param theme_type]. Use [method clear_font_size] to remove the property.
*/
//go:nosplit
func (self class) SetFontSize(name gd.StringName, theme_type gd.StringName, font_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the font size property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font size if the property doesn't exist and the default theme font size is set up (see [member default_font_size]). Use [method has_font_size] to check for existence of the property and [method has_default_font_size] to check for existence of the default theme font.
Returns the engine fallback font size value, if neither exist (see [member ThemeDB.fallback_font_size]).
*/
//go:nosplit
func (self class) GetFontSize(name gd.StringName, theme_type gd.StringName) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the font size property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font_size] to check for existence, and [method clear_font_size] to remove the existing property.
*/
//go:nosplit
func (self class) RenameFontSize(old_name gd.StringName, name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(old_name))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the font size property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font_size] to check for existence.
*/
//go:nosplit
func (self class) ClearFontSize(name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of names for font size properties defined with [param theme_type]. Use [method get_font_size_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetFontSizeList(theme_type gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font_size_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for font size properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetFontSizeTypeList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font_size_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Creates or changes the value of the [Color] property defined by [param name] and [param theme_type]. Use [method clear_color] to remove the property.
*/
//go:nosplit
func (self class) SetColor(name gd.StringName, theme_type gd.StringName, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Color] property defined by [param name] and [param theme_type], if it exists.
Returns the default color value if the property doesn't exist. Use [method has_color] to check for existence.
*/
//go:nosplit
func (self class) GetColor(name gd.StringName, theme_type gd.StringName) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the [Color] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_color] to check for existence, and [method clear_color] to remove the existing property.
*/
//go:nosplit
func (self class) RenameColor(old_name gd.StringName, name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(old_name))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the [Color] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_color] to check for existence.
*/
//go:nosplit
func (self class) ClearColor(name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of names for [Color] properties defined with [param theme_type]. Use [method get_color_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetColorList(theme_type gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_color_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for [Color] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetColorTypeList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_color_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Creates or changes the value of the constant property defined by [param name] and [param theme_type]. Use [method clear_constant] to remove the property.
*/
//go:nosplit
func (self class) SetConstant(name gd.StringName, theme_type gd.StringName, constant gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	callframe.Arg(frame, constant)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the constant property defined by [param name] and [param theme_type], if it exists.
Returns [code]0[/code] if the property doesn't exist. Use [method has_constant] to check for existence.
*/
//go:nosplit
func (self class) GetConstant(name gd.StringName, theme_type gd.StringName) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the constant property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_constant] to check for existence, and [method clear_constant] to remove the existing property.
*/
//go:nosplit
func (self class) RenameConstant(old_name gd.StringName, name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(old_name))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the constant property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_constant] to check for existence.
*/
//go:nosplit
func (self class) ClearConstant(name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of names for constant properties defined with [param theme_type]. Use [method get_constant_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetConstantList(theme_type gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_constant_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for constant properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetConstantTypeList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_constant_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultBaseScale(base_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, base_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultBaseScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultFont(font objects.Font) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_default_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultFont() objects.Font {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_default_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Font{classdb.Font(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [member default_font] has a valid value.
Returns [code]false[/code] if it doesn't.
*/
//go:nosplit
func (self class) HasDefaultFont() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_default_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultFontSize(font_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_default_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultFontSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_default_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_default_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SetThemeItem(data_type classdb.ThemeDataType, name gd.StringName, theme_type gd.StringName, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_theme_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback value if the property doesn't exist (see [ThemeDB]). Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) GetThemeItem(data_type classdb.ThemeDataType, name gd.StringName, theme_type gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_theme_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
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
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_theme_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) RenameThemeItem(data_type classdb.ThemeDataType, old_name gd.StringName, name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(old_name))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_theme_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) ClearThemeItem(data_type classdb.ThemeDataType, name gd.StringName, theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_theme_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of names for properties of [param data_type] defined with [param theme_type]. Use [method get_theme_item_type_list] to get a list of possible theme type names.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) GetThemeItemList(data_type classdb.ThemeDataType, theme_type gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_theme_item_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for [param data_type] properties. Use [method get_type_list] to get a list of all unique theme types.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) GetThemeItemTypeList(data_type classdb.ThemeDataType) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_theme_item_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
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
func (self class) SetTypeVariation(theme_type gd.StringName, base_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	callframe.Arg(frame, pointers.Get(base_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_type_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if [param theme_type] is marked as a variation of [param base_type].
*/
//go:nosplit
func (self class) IsTypeVariation(theme_type gd.StringName, base_type gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	callframe.Arg(frame, pointers.Get(base_type))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_is_type_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Unmarks [param theme_type] as being a variation of another theme type. See [method set_type_variation].
*/
//go:nosplit
func (self class) ClearTypeVariation(theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_type_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the name of the base theme type if [param theme_type] is a valid variation type. Returns an empty string otherwise.
*/
//go:nosplit
func (self class) GetTypeVariationBase(theme_type gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_type_variation_base, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of all type variations for the given [param base_type].
*/
//go:nosplit
func (self class) GetTypeVariationList(base_type gd.StringName) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(base_type))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_type_variation_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds an empty theme type for every valid data type.
[b]Note:[/b] Empty types are not saved with the theme. This method only exists to perform in-memory changes to the resource. Use available [code]set_*[/code] methods to add theme items.
*/
//go:nosplit
func (self class) AddType(theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_add_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the theme type, gracefully discarding defined theme items. If the type is a variation, this information is also erased. If the type is a base for type variations, those variations lose their base.
*/
//go:nosplit
func (self class) RemoveType(theme_type gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(theme_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_remove_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of all unique theme type names. Use the appropriate [code]get_*_type_list[/code] method to get a list of unique theme types for a single data type.
*/
//go:nosplit
func (self class) GetTypeList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_type_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds missing and overrides existing definitions with values from the [param other] theme resource.
[b]Note:[/b] This modifies the current theme. If you want to merge two themes together without modifying either one, create a new empty theme and merge the other two into it one after another.
*/
//go:nosplit
func (self class) MergeWith(other objects.Theme) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(other[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_merge_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all the theme properties defined on the theme resource.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsTheme() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTheme() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() { classdb.Register("Theme", func(ptr gd.Object) any { return classdb.Theme(ptr) }) }

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
