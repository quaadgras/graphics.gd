package objects

type PropertyHint int

const (
	/*The property has no hint for the editor.*/
	PropertyHintNone PropertyHint = 0
	/*Hints that an [int] or [float] property should be within a range specified via the hint string [code]"min,max"[/code] or [code]"min,max,step"[/code]. The hint string can optionally include [code]"or_greater"[/code] and/or [code]"or_less"[/code] to allow manual input going respectively above the max or below the min values.
	  [b]Example:[/b] [code]"-360,360,1,or_greater,or_less"[/code].
	  Additionally, other keywords can be included: [code]"exp"[/code] for exponential range editing, [code]"radians_as_degrees"[/code] for editing radian angles in degrees (the range values are also in degrees), [code]"degrees"[/code] to hint at an angle and [code]"hide_slider"[/code] to hide the slider.*/
	PropertyHintRange PropertyHint = 1
	/*Hints that an [int] or [String] property is an enumerated value to pick in a list specified via a hint string.
	  The hint string is a comma separated list of names such as [code]"Hello,Something,Else"[/code]. Whitespaces are [b]not[/b] removed from either end of a name. For integer properties, the first name in the list has value 0, the next 1, and so on. Explicit values can also be specified by appending [code]:integer[/code] to the name, e.g. [code]"Zero,One,Three:3,Four,Six:6"[/code].*/
	PropertyHintEnum PropertyHint = 2
	/*Hints that a [String] property can be an enumerated value to pick in a list specified via a hint string such as [code]"Hello,Something,Else"[/code].
	  Unlike [constant PROPERTY_HINT_ENUM], a property with this hint still accepts arbitrary values and can be empty. The list of values serves to suggest possible values.*/
	PropertyHintEnumSuggestion PropertyHint = 3
	/*Hints that a [float] property should be edited via an exponential easing function. The hint string can include [code]"attenuation"[/code] to flip the curve horizontally and/or [code]"positive_only"[/code] to exclude in/out easing and limit values to be greater than or equal to zero.*/
	PropertyHintExpEasing PropertyHint = 4
	/*Hints that a vector property should allow its components to be linked. For example, this allows [member Vector2.x] and [member Vector2.y] to be edited together.*/
	PropertyHintLink PropertyHint = 5
	/*Hints that an [int] property is a bitmask with named bit flags.
	  The hint string is a comma separated list of names such as [code]"Bit0,Bit1,Bit2,Bit3"[/code]. Whitespaces are [b]not[/b] removed from either end of a name. The first name in the list has value 1, the next 2, then 4, 8, 16 and so on. Explicit values can also be specified by appending [code]:integer[/code] to the name, e.g. [code]"A:4,B:8,C:16"[/code]. You can also combine several flags ([code]"A:4,B:8,AB:12,C:16"[/code]).
	  [b]Note:[/b] A flag value must be at least [code]1[/code] and at most [code]2 ** 32 - 1[/code].
	  [b]Note:[/b] Unlike [constant PROPERTY_HINT_ENUM], the previous explicit value is not taken into account. For the hint [code]"A:16,B,C"[/code], A is 16, B is 2, C is 4.*/
	PropertyHintFlags PropertyHint = 6
	/*Hints that an [int] property is a bitmask using the optionally named 2D render layers.*/
	PropertyHintLayers2dRender PropertyHint = 7
	/*Hints that an [int] property is a bitmask using the optionally named 2D physics layers.*/
	PropertyHintLayers2dPhysics PropertyHint = 8
	/*Hints that an [int] property is a bitmask using the optionally named 2D navigation layers.*/
	PropertyHintLayers2dNavigation PropertyHint = 9
	/*Hints that an [int] property is a bitmask using the optionally named 3D render layers.*/
	PropertyHintLayers3dRender PropertyHint = 10
	/*Hints that an [int] property is a bitmask using the optionally named 3D physics layers.*/
	PropertyHintLayers3dPhysics PropertyHint = 11
	/*Hints that an [int] property is a bitmask using the optionally named 3D navigation layers.*/
	PropertyHintLayers3dNavigation PropertyHint = 12
	/*Hints that an integer property is a bitmask using the optionally named avoidance layers.*/
	PropertyHintLayersAvoidance PropertyHint = 37
	/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code].*/
	PropertyHintFile PropertyHint = 13
	/*Hints that a [String] property is a path to a directory. Editing it will show a file dialog for picking the path.*/
	PropertyHintDir PropertyHint = 14
	/*Hints that a [String] property is an absolute path to a file outside the project folder. Editing it will show a file dialog for picking the path. The hint string can be a set of filters with wildcards, like [code]"*.png,*.jpg"[/code].*/
	PropertyHintGlobalFile PropertyHint = 15
	/*Hints that a [String] property is an absolute path to a directory outside the project folder. Editing it will show a file dialog for picking the path.*/
	PropertyHintGlobalDir PropertyHint = 16
	/*Hints that a property is an instance of a [Resource]-derived type, optionally specified via the hint string (e.g. [code]"Texture2D"[/code]). Editing it will show a popup menu of valid resource types to instantiate.*/
	PropertyHintResourceType PropertyHint = 17
	/*Hints that a [String] property is text with line breaks. Editing it will show a text input field where line breaks can be typed.*/
	PropertyHintMultilineText PropertyHint = 18
	/*Hints that a [String] property is an [Expression].*/
	PropertyHintExpression PropertyHint = 19
	/*Hints that a [String] property should show a placeholder text on its input field, if empty. The hint string is the placeholder text to use.*/
	PropertyHintPlaceholderText PropertyHint = 20
	/*Hints that a [Color] property should be edited without affecting its transparency ([member Color.a] is not editable).*/
	PropertyHintColorNoAlpha PropertyHint = 21
	/*Hints that the property's value is an object encoded as object ID, with its type specified in the hint string. Used by the debugger.*/
	PropertyHintObjectId PropertyHint = 22
	/*If a property is [String], hints that the property represents a particular type (class). This allows to select a type from the create dialog. The property will store the selected type as a string.
	  If a property is [Array], hints the editor how to show elements. The [code]hint_string[/code] must encode nested types using [code]":"[/code] and [code]"/"[/code].
	  [codeblocks]
	  [gdscript]
	  # Array of elem_type.
	  hint_string = "%d:" % [elem_type]
	  hint_string = "%d/%d:%s" % [elem_type, elem_hint, elem_hint_string]
	  # Two-dimensional array of elem_type (array of arrays of elem_type).
	  hint_string = "%d:%d:" % [TYPE_ARRAY, elem_type]
	  hint_string = "%d:%d/%d:%s" % [TYPE_ARRAY, elem_type, elem_hint, elem_hint_string]
	  # Three-dimensional array of elem_type (array of arrays of arrays of elem_type).
	  hint_string = "%d:%d:%d:" % [TYPE_ARRAY, TYPE_ARRAY, elem_type]
	  hint_string = "%d:%d:%d/%d:%s" % [TYPE_ARRAY, TYPE_ARRAY, elem_type, elem_hint, elem_hint_string]
	  [/gdscript]
	  [csharp]
	  // Array of elemType.
	  hintString = $"{elemType:D}:";
	  hintString = $"{elemType:}/{elemHint:D}:{elemHintString}";
	  // Two-dimensional array of elemType (array of arrays of elemType).
	  hintString = $"{Variant.Type.Array:D}:{elemType:D}:";
	  hintString = $"{Variant.Type.Array:D}:{elemType:D}/{elemHint:D}:{elemHintString}";
	  // Three-dimensional array of elemType (array of arrays of arrays of elemType).
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Array:D}:{elemType:D}:";
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Array:D}:{elemType:D}/{elemHint:D}:{elemHintString}";
	  [/csharp]
	  [/codeblocks]
	  Examples:
	  [codeblocks]
	  [gdscript]
	  hint_string = "%d:" % [TYPE_INT] # Array of integers.
	  hint_string = "%d/%d:1,10,1" % [TYPE_INT, PROPERTY_HINT_RANGE] # Array of integers (in range from 1 to 10).
	  hint_string = "%d/%d:Zero,One,Two" % [TYPE_INT, PROPERTY_HINT_ENUM] # Array of integers (an enum).
	  hint_string = "%d/%d:Zero,One,Three:3,Six:6" % [TYPE_INT, PROPERTY_HINT_ENUM] # Array of integers (an enum).
	  hint_string = "%d/%d:*.png" % [TYPE_STRING, PROPERTY_HINT_FILE] # Array of strings (file paths).
	  hint_string = "%d/%d:Texture2D" % [TYPE_OBJECT, PROPERTY_HINT_RESOURCE_TYPE] # Array of textures.

	  hint_string = "%d:%d:" % [TYPE_ARRAY, TYPE_FLOAT] # Two-dimensional array of floats.
	  hint_string = "%d:%d/%d:" % [TYPE_ARRAY, TYPE_STRING, PROPERTY_HINT_MULTILINE_TEXT] # Two-dimensional array of multiline strings.
	  hint_string = "%d:%d/%d:-1,1,0.1" % [TYPE_ARRAY, TYPE_FLOAT, PROPERTY_HINT_RANGE] # Two-dimensional array of floats (in range from -1 to 1).
	  hint_string = "%d:%d/%d:Texture2D" % [TYPE_ARRAY, TYPE_OBJECT, PROPERTY_HINT_RESOURCE_TYPE] # Two-dimensional array of textures.
	  [/gdscript]
	  [csharp]
	  hintString = $"{Variant.Type.Int:D}/{PropertyHint.Range:D}:1,10,1"; // Array of integers (in range from 1 to 10).
	  hintString = $"{Variant.Type.Int:D}/{PropertyHint.Enum:D}:Zero,One,Two"; // Array of integers (an enum).
	  hintString = $"{Variant.Type.Int:D}/{PropertyHint.Enum:D}:Zero,One,Three:3,Six:6"; // Array of integers (an enum).
	  hintString = $"{Variant.Type.String:D}/{PropertyHint.File:D}:*.png"; // Array of strings (file paths).
	  hintString = $"{Variant.Type.Object:D}/{PropertyHint.ResourceType:D}:Texture2D"; // Array of textures.

	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Float:D}:"; // Two-dimensional array of floats.
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.String:D}/{PropertyHint.MultilineText:D}:"; // Two-dimensional array of multiline strings.
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Float:D}/{PropertyHint.Range:D}:-1,1,0.1"; // Two-dimensional array of floats (in range from -1 to 1).
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Object:D}/{PropertyHint.ResourceType:D}:Texture2D"; // Two-dimensional array of textures.
	  [/csharp]
	  [/codeblocks]
	  [b]Note:[/b] The trailing colon is required for properly detecting built-in types.*/
	PropertyHintTypeString           PropertyHint = 23
	PropertyHintNodePathToEditedNode PropertyHint = 24
	/*Hints that an object is too big to be sent via the debugger.*/
	PropertyHintObjectTooBig PropertyHint = 25
	/*Hints that the hint string specifies valid node types for property of type [NodePath].*/
	PropertyHintNodePathValidTypes PropertyHint = 26
	/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path for the file to be saved at. The dialog has access to the project's directory. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code]. See also [member FileDialog.filters].*/
	PropertyHintSaveFile PropertyHint = 27
	/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path for the file to be saved at. The dialog has access to the entire filesystem. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code]. See also [member FileDialog.filters].*/
	PropertyHintGlobalSaveFile PropertyHint = 28
	PropertyHintIntIsObjectid  PropertyHint = 29
	/*Hints that an [int] property is a pointer. Used by GDExtension.*/
	PropertyHintIntIsPointer PropertyHint = 30
	/*Hints that a property is an [Array] with the stored type specified in the hint string.*/
	PropertyHintArrayType PropertyHint = 31
	/*Hints that a string property is a locale code. Editing it will show a locale dialog for picking language and country.*/
	PropertyHintLocaleId PropertyHint = 32
	/*Hints that a dictionary property is string translation map. Dictionary keys are locale codes and, values are translated strings.*/
	PropertyHintLocalizableString PropertyHint = 33
	/*Hints that a property is an instance of a [Node]-derived type, optionally specified via the hint string (e.g. [code]"Node2D"[/code]). Editing it will show a dialog for picking a node from the scene.*/
	PropertyHintNodeType PropertyHint = 34
	/*Hints that a quaternion property should disable the temporary euler editor.*/
	PropertyHintHideQuaternionEdit PropertyHint = 35
	/*Hints that a string property is a password, and every character is replaced with the secret character.*/
	PropertyHintPassword PropertyHint = 36
	/*Represents the size of the [enum PropertyHint] enum.*/
	PropertyHintMax PropertyHint = 38
)

type PropertyUsageFlags int

const (
	/*The property is not stored, and does not display in the editor. This is the default for non-exported properties.*/
	PropertyUsageNone PropertyUsageFlags = 0
	/*The property is serialized and saved in the scene file (default for exported properties).*/
	PropertyUsageStorage PropertyUsageFlags = 2
	/*The property is shown in the [EditorInspector] (default for exported properties).*/
	PropertyUsageEditor PropertyUsageFlags = 4
	/*The property is excluded from the class reference.*/
	PropertyUsageInternal PropertyUsageFlags = 8
	/*The property can be checked in the [EditorInspector].*/
	PropertyUsageCheckable PropertyUsageFlags = 16
	/*The property is checked in the [EditorInspector].*/
	PropertyUsageChecked PropertyUsageFlags = 32
	/*Used to group properties together in the editor. See [EditorInspector].*/
	PropertyUsageGroup PropertyUsageFlags = 64
	/*Used to categorize properties together in the editor.*/
	PropertyUsageCategory PropertyUsageFlags = 128
	/*Used to group properties together in the editor in a subgroup (under a group). See [EditorInspector].*/
	PropertyUsageSubgroup PropertyUsageFlags = 256
	/*The property is a bitfield, i.e. it contains multiple flags represented as bits.*/
	PropertyUsageClassIsBitfield PropertyUsageFlags = 512
	/*The property does not save its state in [PackedScene].*/
	PropertyUsageNoInstanceState PropertyUsageFlags = 1024
	/*Editing the property prompts the user for restarting the editor.*/
	PropertyUsageRestartIfChanged PropertyUsageFlags = 2048
	/*The property is a script variable which should be serialized and saved in the scene file.*/
	PropertyUsageScriptVariable PropertyUsageFlags = 4096
	/*The property value of type [Object] will be stored even if its value is [code]null[/code].*/
	PropertyUsageStoreIfNull PropertyUsageFlags = 8192
	/*If this property is modified, all inspector fields will be refreshed.*/
	PropertyUsageUpdateAllIfModified PropertyUsageFlags = 16384
	PropertyUsageScriptDefaultValue  PropertyUsageFlags = 32768
	/*The property is an enum, i.e. it only takes named integer constants from its associated enumeration.*/
	PropertyUsageClassIsEnum PropertyUsageFlags = 65536
	/*If property has [code]nil[/code] as default value, its type will be [Variant].*/
	PropertyUsageNilIsVariant PropertyUsageFlags = 131072
	/*The property is an array.*/
	PropertyUsageArray PropertyUsageFlags = 262144
	/*When duplicating a resource with [method Resource.duplicate], and this flag is set on a property of that resource, the property should always be duplicated, regardless of the [code]subresources[/code] bool parameter.*/
	PropertyUsageAlwaysDuplicate PropertyUsageFlags = 524288
	/*When duplicating a resource with [method Resource.duplicate], and this flag is set on a property of that resource, the property should never be duplicated, regardless of the [code]subresources[/code] bool parameter.*/
	PropertyUsageNeverDuplicate PropertyUsageFlags = 1048576
	/*The property is only shown in the editor if modern renderers are supported (the Compatibility rendering method is excluded).*/
	PropertyUsageHighEndGfx PropertyUsageFlags = 2097152
	/*The [NodePath] property will always be relative to the scene's root. Mostly useful for local resources.*/
	PropertyUsageNodePathFromSceneRoot PropertyUsageFlags = 4194304
	/*Use when a resource is created on the fly, i.e. the getter will always return a different instance. [ResourceSaver] needs this information to properly save such resources.*/
	PropertyUsageResourceNotPersistent PropertyUsageFlags = 8388608
	/*Inserting an animation key frame of this property will automatically increment the value, allowing to easily keyframe multiple values in a row.*/
	PropertyUsageKeyingIncrements    PropertyUsageFlags = 16777216
	PropertyUsageDeferredSetResource PropertyUsageFlags = 33554432
	/*When this property is a [Resource] and base object is a [Node], a resource instance will be automatically created whenever the node is created in the editor.*/
	PropertyUsageEditorInstantiateObject PropertyUsageFlags = 67108864
	/*The property is considered a basic setting and will appear even when advanced mode is disabled. Used for project settings.*/
	PropertyUsageEditorBasicSetting PropertyUsageFlags = 134217728
	/*The property is read-only in the [EditorInspector].*/
	PropertyUsageReadOnly PropertyUsageFlags = 268435456
	/*An export preset property with this flag contains confidential information and is stored separately from the rest of the export preset configuration.*/
	PropertyUsageSecret PropertyUsageFlags = 536870912
	/*Default usage (storage and editor).*/
	PropertyUsageDefault PropertyUsageFlags = 6
	/*Default usage but without showing the property in the editor (storage).*/
	PropertyUsageNoEditor PropertyUsageFlags = 2
)

type MethodFlags int

const (
	/*Flag for a normal method.*/
	MethodFlagNormal MethodFlags = 1
	/*Flag for an editor method.*/
	MethodFlagEditor MethodFlags = 2
	/*Flag for a constant method.*/
	MethodFlagConst MethodFlags = 4
	/*Flag for a virtual method.*/
	MethodFlagVirtual MethodFlags = 8
	/*Flag for a method with a variable number of arguments.*/
	MethodFlagVararg MethodFlags = 16
	/*Flag for a static method.*/
	MethodFlagStatic MethodFlags = 32
	/*Used internally. Allows to not dump core virtual methods (such as [method Object._notification]) to the JSON API.*/
	MethodFlagObjectCore MethodFlags = 64
	/*Default method flags (normal).*/
	MethodFlagsDefault MethodFlags = 1
)
