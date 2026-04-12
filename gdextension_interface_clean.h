
/**
 * @name mem_alloc2
 * @since 4.6
 *
 * Allocates memory.
 *
 * @param p_bytes The amount of memory to allocate in bytes.
 * @param p_pad_align If true, the returned memory will have prepadding of at least 8 bytes.
 *
 * @return A pointer to the allocated memory, or NULL if unsuccessful.
 */
typedef void *(*GDExtensionInterfaceMemAlloc2)(size_t p_bytes, GDExtensionBool p_pad_align);

/**
 * @name mem_realloc2
 * @since 4.6
 *
 * Reallocates memory.
 *
 * @param p_ptr A pointer to the previously allocated memory.
 * @param p_bytes The number of bytes to resize the memory block to.
 * @param p_pad_align If true, the returned memory will have prepadding of at least 8 bytes.
 *
 * @return A pointer to the allocated memory, or NULL if unsuccessful.
 */
typedef void *(*GDExtensionInterfaceMemRealloc2)(void *p_ptr, size_t p_bytes, GDExtensionBool p_pad_align);

/**
 * @name mem_free2
 * @since 4.6
 *
 * Frees memory.
 *
 * @param p_ptr A pointer to the previously allocated memory.
 * @param p_pad_align If true, the given memory was allocated with prepadding.
 */
typedef void (*GDExtensionInterfaceMemFree2)(void *p_ptr, GDExtensionBool p_pad_align);

/**
 * @name print_error
 * @since 4.1
 *
 * Logs an error to Godot's built-in debugger and to the OS terminal.
 *
 * @param p_description The code triggering the error.
 * @param p_function The function name where the error occurred.
 * @param p_file The file where the error occurred.
 * @param p_line The line where the error occurred.
 * @param p_editor_notify Whether or not to notify the editor.
 */
typedef void (*GDExtensionInterfacePrintError)(const char *p_description, const char *p_function, const char *p_file, int32_t p_line, GDExtensionBool p_editor_notify);

/**
 * @name print_error_with_message
 * @since 4.1
 *
 * Logs an error with a message to Godot's built-in debugger and to the OS terminal.
 *
 * @param p_description The code triggering the error.
 * @param p_message The message to show along with the error.
 * @param p_function The function name where the error occurred.
 * @param p_file The file where the error occurred.
 * @param p_line The line where the error occurred.
 * @param p_editor_notify Whether or not to notify the editor.
 */
typedef void (*GDExtensionInterfacePrintErrorWithMessage)(const char *p_description, const char *p_message, const char *p_function, const char *p_file, int32_t p_line, GDExtensionBool p_editor_notify);

/**
 * @name print_warning
 * @since 4.1
 *
 * Logs a warning to Godot's built-in debugger and to the OS terminal.
 *
 * @param p_description The code triggering the warning.
 * @param p_function The function name where the warning occurred.
 * @param p_file The file where the warning occurred.
 * @param p_line The line where the warning occurred.
 * @param p_editor_notify Whether or not to notify the editor.
 */
typedef void (*GDExtensionInterfacePrintWarning)(const char *p_description, const char *p_function, const char *p_file, int32_t p_line, GDExtensionBool p_editor_notify);

/**
 * @name print_warning_with_message
 * @since 4.1
 *
 * Logs a warning with a message to Godot's built-in debugger and to the OS terminal.
 *
 * @param p_description The code triggering the warning.
 * @param p_message The message to show along with the warning.
 * @param p_function The function name where the warning occurred.
 * @param p_file The file where the warning occurred.
 * @param p_line The line where the warning occurred.
 * @param p_editor_notify Whether or not to notify the editor.
 */
typedef void (*GDExtensionInterfacePrintWarningWithMessage)(const char *p_description, const char *p_message, const char *p_function, const char *p_file, int32_t p_line, GDExtensionBool p_editor_notify);

/**
 * @name print_script_error
 * @since 4.1
 *
 * Logs a script error to Godot's built-in debugger and to the OS terminal.
 *
 * @param p_description The code triggering the error.
 * @param p_function The function name where the error occurred.
 * @param p_file The file where the error occurred.
 * @param p_line The line where the error occurred.
 * @param p_editor_notify Whether or not to notify the editor.
 */
typedef void (*GDExtensionInterfacePrintScriptError)(const char *p_description, const char *p_function, const char *p_file, int32_t p_line, GDExtensionBool p_editor_notify);

/**
 * @name print_script_error_with_message
 * @since 4.1
 *
 * Logs a script error with a message to Godot's built-in debugger and to the OS terminal.
 *
 * @param p_description The code triggering the error.
 * @param p_message The message to show along with the error.
 * @param p_function The function name where the error occurred.
 * @param p_file The file where the error occurred.
 * @param p_line The line where the error occurred.
 * @param p_editor_notify Whether or not to notify the editor.
 */
typedef void (*GDExtensionInterfacePrintScriptErrorWithMessage)(const char *p_description, const char *p_message, const char *p_function, const char *p_file, int32_t p_line, GDExtensionBool p_editor_notify);

/**
 * @name get_native_struct_size
 * @since 4.1
 *
 * Gets the size of a native struct (ex. ObjectID) in bytes.
 *
 * @param p_name A pointer to a StringName identifying the struct name.
 *
 * @return The size in bytes.
 */
typedef uint64_t (*GDExtensionInterfaceGetNativeStructSize)(GDExtensionConstStringNamePtr p_name);

/**
 * @name variant_new_copy
 * @since 4.1
 *
 * Copies one Variant into a another.
 *
 * @param r_dest A pointer to the destination Variant.
 * @param p_src A pointer to the source Variant.
 */
typedef void (*GDExtensionInterfaceVariantNewCopy)(GDExtensionUninitializedVariantPtr r_dest, GDExtensionConstVariantPtr p_src);

/**
 * @name variant_new_nil
 * @since 4.1
 *
 * Creates a new Variant containing nil.
 *
 * @param r_dest A pointer to the destination Variant.
 */
typedef void (*GDExtensionInterfaceVariantNewNil)(GDExtensionUninitializedVariantPtr r_dest);

/**
 * @name variant_destroy
 * @since 4.1
 *
 * Destroys a Variant.
 *
 * @param p_self A pointer to the Variant to destroy.
 */
typedef void (*GDExtensionInterfaceVariantDestroy)(GDExtensionVariantPtr p_self);

/**
 * @name variant_call
 * @since 4.1
 *
 * Calls a method on a Variant.
 *
 * @param p_self A pointer to the Variant.
 * @param p_method A pointer to a StringName identifying the method.
 * @param p_args A pointer to a C array of Variant.
 * @param p_argument_count The number of arguments.
 * @param r_return A pointer a Variant which will be assigned the return value.
 * @param r_error A pointer the structure which will hold error information.
 *
 * @see Variant::callp()
 */
typedef void (*GDExtensionInterfaceVariantCall)(GDExtensionVariantPtr p_self, GDExtensionConstStringNamePtr p_method, const GDExtensionConstVariantPtr *p_args, GDExtensionInt p_argument_count, GDExtensionUninitializedVariantPtr r_return, GDExtensionCallError *r_error);

/**
 * @name variant_call_static
 * @since 4.1
 *
 * Calls a static method on a Variant.
 *
 * @param p_type The variant type.
 * @param p_method A pointer to a StringName identifying the method.
 * @param p_args A pointer to a C array of Variant.
 * @param p_argument_count The number of arguments.
 * @param r_return A pointer a Variant which will be assigned the return value.
 * @param r_error A pointer the structure which will be updated with error information.
 *
 * @see Variant::call_static()
 */
typedef void (*GDExtensionInterfaceVariantCallStatic)(GDExtensionVariantType p_type, GDExtensionConstStringNamePtr p_method, const GDExtensionConstVariantPtr *p_args, GDExtensionInt p_argument_count, GDExtensionUninitializedVariantPtr r_return, GDExtensionCallError *r_error);

/**
 * @name variant_evaluate
 * @since 4.1
 *
 * Evaluate an operator on two Variants.
 *
 * @param p_op The operator to evaluate.
 * @param p_a The first Variant.
 * @param p_b The second Variant.
 * @param r_return A pointer a Variant which will be assigned the return value.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 *
 * @see Variant::evaluate()
 */
typedef void (*GDExtensionInterfaceVariantEvaluate)(GDExtensionVariantOperator p_op, GDExtensionConstVariantPtr p_a, GDExtensionConstVariantPtr p_b, GDExtensionUninitializedVariantPtr r_return, GDExtensionBool *r_valid);

/**
 * @name variant_set
 * @since 4.1
 *
 * Sets a key on a Variant to a value.
 *
 * @param p_self A pointer to the Variant.
 * @param p_key A pointer to a Variant representing the key.
 * @param p_value A pointer to a Variant representing the value.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 *
 * @see Variant::set()
 */
typedef void (*GDExtensionInterfaceVariantSet)(GDExtensionVariantPtr p_self, GDExtensionConstVariantPtr p_key, GDExtensionConstVariantPtr p_value, GDExtensionBool *r_valid);

/**
 * @name variant_set_named
 * @since 4.1
 *
 * Sets a named key on a Variant to a value.
 *
 * @param p_self A pointer to the Variant.
 * @param p_key A pointer to a StringName representing the key.
 * @param p_value A pointer to a Variant representing the value.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 *
 * @see Variant::set_named()
 */
typedef void (*GDExtensionInterfaceVariantSetNamed)(GDExtensionVariantPtr p_self, GDExtensionConstStringNamePtr p_key, GDExtensionConstVariantPtr p_value, GDExtensionBool *r_valid);

/**
 * @name variant_set_keyed
 * @since 4.1
 *
 * Sets a keyed property on a Variant to a value.
 *
 * @param p_self A pointer to the Variant.
 * @param p_key A pointer to a Variant representing the key.
 * @param p_value A pointer to a Variant representing the value.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 *
 * @see Variant::set_keyed()
 */
typedef void (*GDExtensionInterfaceVariantSetKeyed)(GDExtensionVariantPtr p_self, GDExtensionConstVariantPtr p_key, GDExtensionConstVariantPtr p_value, GDExtensionBool *r_valid);

/**
 * @name variant_set_indexed
 * @since 4.1
 *
 * Sets an index on a Variant to a value.
 *
 * @param p_self A pointer to the Variant.
 * @param p_index The index.
 * @param p_value A pointer to a Variant representing the value.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 * @param r_oob A pointer to a boolean which will be set to true if the index is out of bounds.
 */
typedef void (*GDExtensionInterfaceVariantSetIndexed)(GDExtensionVariantPtr p_self, GDExtensionInt p_index, GDExtensionConstVariantPtr p_value, GDExtensionBool *r_valid, GDExtensionBool *r_oob);

/**
 * @name variant_get
 * @since 4.1
 *
 * Gets the value of a key from a Variant.
 *
 * @param p_self A pointer to the Variant.
 * @param p_key A pointer to a Variant representing the key.
 * @param r_ret A pointer to a Variant which will be assigned the value.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 */
typedef void (*GDExtensionInterfaceVariantGet)(GDExtensionConstVariantPtr p_self, GDExtensionConstVariantPtr p_key, GDExtensionUninitializedVariantPtr r_ret, GDExtensionBool *r_valid);

/**
 * @name variant_get_named
 * @since 4.1
 *
 * Gets the value of a named key from a Variant.
 *
 * @param p_self A pointer to the Variant.
 * @param p_key A pointer to a StringName representing the key.
 * @param r_ret A pointer to a Variant which will be assigned the value.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 */
typedef void (*GDExtensionInterfaceVariantGetNamed)(GDExtensionConstVariantPtr p_self, GDExtensionConstStringNamePtr p_key, GDExtensionUninitializedVariantPtr r_ret, GDExtensionBool *r_valid);

/**
 * @name variant_get_keyed
 * @since 4.1
 *
 * Gets the value of a keyed property from a Variant.
 *
 * @param p_self A pointer to the Variant.
 * @param p_key A pointer to a Variant representing the key.
 * @param r_ret A pointer to a Variant which will be assigned the value.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 */
typedef void (*GDExtensionInterfaceVariantGetKeyed)(GDExtensionConstVariantPtr p_self, GDExtensionConstVariantPtr p_key, GDExtensionUninitializedVariantPtr r_ret, GDExtensionBool *r_valid);

/**
 * @name variant_get_indexed
 * @since 4.1
 *
 * Gets the value of an index from a Variant.
 *
 * @param p_self A pointer to the Variant.
 * @param p_index The index.
 * @param r_ret A pointer to a Variant which will be assigned the value.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 * @param r_oob A pointer to a boolean which will be set to true if the index is out of bounds.
 */
typedef void (*GDExtensionInterfaceVariantGetIndexed)(GDExtensionConstVariantPtr p_self, GDExtensionInt p_index, GDExtensionUninitializedVariantPtr r_ret, GDExtensionBool *r_valid, GDExtensionBool *r_oob);

/**
 * @name variant_iter_init
 * @since 4.1
 *
 * Initializes an iterator over a Variant.
 *
 * @param p_self A pointer to the Variant.
 * @param r_iter A pointer to a Variant which will be assigned the iterator.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 *
 * @return true if the operation is valid; otherwise false.
 *
 * @see Variant::iter_init()
 */
typedef GDExtensionBool (*GDExtensionInterfaceVariantIterInit)(GDExtensionConstVariantPtr p_self, GDExtensionUninitializedVariantPtr r_iter, GDExtensionBool *r_valid);

/**
 * @name variant_iter_next
 * @since 4.1
 *
 * Gets the next value for an iterator over a Variant.
 *
 * @param p_self A pointer to the Variant.
 * @param r_iter A pointer to a Variant which will be assigned the iterator.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 *
 * @return true if the operation is valid; otherwise false.
 *
 * @see Variant::iter_next()
 */
typedef GDExtensionBool (*GDExtensionInterfaceVariantIterNext)(GDExtensionConstVariantPtr p_self, GDExtensionVariantPtr r_iter, GDExtensionBool *r_valid);

/**
 * @name variant_iter_get
 * @since 4.1
 *
 * Gets the next value for an iterator over a Variant.
 *
 * @param p_self A pointer to the Variant.
 * @param r_iter A pointer to a Variant which will be assigned the iterator.
 * @param r_ret A pointer to a Variant which will be assigned false if the operation is invalid.
 * @param r_valid A pointer to a boolean which will be set to false if the operation is invalid.
 *
 * @see Variant::iter_get()
 */
typedef void (*GDExtensionInterfaceVariantIterGet)(GDExtensionConstVariantPtr p_self, GDExtensionVariantPtr r_iter, GDExtensionUninitializedVariantPtr r_ret, GDExtensionBool *r_valid);

/**
 * @name variant_hash
 * @since 4.1
 *
 * Gets the hash of a Variant.
 *
 * @param p_self A pointer to the Variant.
 *
 * @return The hash value.
 *
 * @see Variant::hash()
 */
typedef GDExtensionInt (*GDExtensionInterfaceVariantHash)(GDExtensionConstVariantPtr p_self);

/**
 * @name variant_recursive_hash
 * @since 4.1
 *
 * Gets the recursive hash of a Variant.
 *
 * @param p_self A pointer to the Variant.
 * @param p_recursion_count The number of recursive loops so far.
 *
 * @return The hash value.
 *
 * @see Variant::recursive_hash()
 */
typedef GDExtensionInt (*GDExtensionInterfaceVariantRecursiveHash)(GDExtensionConstVariantPtr p_self, GDExtensionInt p_recursion_count);

/**
 * @name variant_hash_compare
 * @since 4.1
 *
 * Compares two Variants by their hash.
 *
 * @param p_self A pointer to the Variant.
 * @param p_other A pointer to the other Variant to compare it to.
 *
 * @return The hash value.
 *
 * @see Variant::hash_compare()
 */
typedef GDExtensionBool (*GDExtensionInterfaceVariantHashCompare)(GDExtensionConstVariantPtr p_self, GDExtensionConstVariantPtr p_other);

/**
 * @name variant_booleanize
 * @since 4.1
 *
 * Converts a Variant to a boolean.
 *
 * @param p_self A pointer to the Variant.
 *
 * @return The boolean value of the Variant.
 */
typedef GDExtensionBool (*GDExtensionInterfaceVariantBooleanize)(GDExtensionConstVariantPtr p_self);

/**
 * @name variant_duplicate
 * @since 4.1
 *
 * Duplicates a Variant.
 *
 * @param p_self A pointer to the Variant.
 * @param r_ret A pointer to a Variant to store the duplicated value.
 * @param p_deep Whether or not to duplicate deeply (when supported by the Variant type).
 */
typedef void (*GDExtensionInterfaceVariantDuplicate)(GDExtensionConstVariantPtr p_self, GDExtensionVariantPtr r_ret, GDExtensionBool p_deep);

/**
 * @name variant_stringify
 * @since 4.1
 *
 * Converts a Variant to a string.
 *
 * @param p_self A pointer to the Variant.
 * @param r_ret A pointer to a String to store the resulting value.
 */
typedef void (*GDExtensionInterfaceVariantStringify)(GDExtensionConstVariantPtr p_self, GDExtensionStringPtr r_ret);

/**
 * @name variant_get_type
 * @since 4.1
 *
 * Gets the type of a Variant.
 *
 * @param p_self A pointer to the Variant.
 *
 * @return The variant type.
 */
typedef GDExtensionVariantType (*GDExtensionInterfaceVariantGetType)(GDExtensionConstVariantPtr p_self);

/**
 * @name variant_has_method
 * @since 4.1
 *
 * Checks if a Variant has the given method.
 *
 * @param p_self A pointer to the Variant.
 * @param p_method A pointer to a StringName with the method name.
 *
 * @return true if the variant has the given method; otherwise false.
 */
typedef GDExtensionBool (*GDExtensionInterfaceVariantHasMethod)(GDExtensionConstVariantPtr p_self, GDExtensionConstStringNamePtr p_method);

/**
 * @name variant_has_member
 * @since 4.1
 *
 * Checks if a type of Variant has the given member.
 *
 * @param p_type The Variant type.
 * @param p_member A pointer to a StringName with the member name.
 *
 * @return true if the variant has the given method; otherwise false.
 */
typedef GDExtensionBool (*GDExtensionInterfaceVariantHasMember)(GDExtensionVariantType p_type, GDExtensionConstStringNamePtr p_member);

/**
 * @name variant_has_key
 * @since 4.1
 *
 * Checks if a Variant has a key.
 *
 * @param p_self A pointer to the Variant.
 * @param p_key A pointer to a Variant representing the key.
 * @param r_valid A pointer to a boolean which will be set to false if the key doesn't exist.
 *
 * @return true if the key exists; otherwise false.
 */
typedef GDExtensionBool (*GDExtensionInterfaceVariantHasKey)(GDExtensionConstVariantPtr p_self, GDExtensionConstVariantPtr p_key, GDExtensionBool *r_valid);

/**
 * @name variant_get_object_instance_id
 * @since 4.4
 *
 * Gets the object instance ID from a variant of type GDEXTENSION_VARIANT_TYPE_OBJECT.
 *
 * If the variant isn't of type GDEXTENSION_VARIANT_TYPE_OBJECT, then zero will be returned.
 * The instance ID will be returned even if the object is no longer valid - use `object_get_instance_by_id()` to check if the object is still valid.
 *
 * @param p_self A pointer to the Variant.
 *
 * @return The instance ID for the contained object.
 */
typedef GDObjectInstanceID (*GDExtensionInterfaceVariantGetObjectInstanceId)(GDExtensionConstVariantPtr p_self);

/**
 * @name variant_get_type_name
 * @since 4.1
 *
 * Gets the name of a Variant type.
 *
 * @param p_type The Variant type.
 * @param r_name A pointer to a String to store the Variant type name.
 */
typedef void (*GDExtensionInterfaceVariantGetTypeName)(GDExtensionVariantType p_type, GDExtensionUninitializedStringPtr r_name);

/**
 * @name variant_can_convert
 * @since 4.1
 *
 * Checks if Variants can be converted from one type to another.
 *
 * @param p_from The Variant type to convert from.
 * @param p_to The Variant type to convert to.
 *
 * @return true if the conversion is possible; otherwise false.
 */
typedef GDExtensionBool (*GDExtensionInterfaceVariantCanConvert)(GDExtensionVariantType p_from, GDExtensionVariantType p_to);

/**
 * @name variant_can_convert_strict
 * @since 4.1
 *
 * Checks if Variant can be converted from one type to another using stricter rules.
 *
 * @param p_from The Variant type to convert from.
 * @param p_to The Variant type to convert to.
 *
 * @return true if the conversion is possible; otherwise false.
 */
typedef GDExtensionBool (*GDExtensionInterfaceVariantCanConvertStrict)(GDExtensionVariantType p_from, GDExtensionVariantType p_to);

/**
 * @name get_variant_from_type_constructor
 * @since 4.1
 *
 * Gets a pointer to a function that can create a Variant of the given type from a raw value.
 *
 * @param p_type The Variant type.
 *
 * @return A pointer to a function that can create a Variant of the given type from a raw value.
 */
typedef GDExtensionVariantFromTypeConstructorFunc (*GDExtensionInterfaceGetVariantFromTypeConstructor)(GDExtensionVariantType p_type);

/**
 * @name get_variant_to_type_constructor
 * @since 4.1
 *
 * Gets a pointer to a function that can get the raw value from a Variant of the given type.
 *
 * @param p_type The Variant type.
 *
 * @return A pointer to a function that can get the raw value from a Variant of the given type.
 */
typedef GDExtensionTypeFromVariantConstructorFunc (*GDExtensionInterfaceGetVariantToTypeConstructor)(GDExtensionVariantType p_type);

/**
 * @name variant_get_ptr_internal_getter
 * @since 4.4
 *
 * Provides a function pointer for retrieving a pointer to a variant's internal value.
 *
 * Access to a variant's internal value can be used to modify it in-place, or to retrieve its value without the overhead of variant conversion functions.
 * It is recommended to cache the getter for all variant types in a function table to avoid retrieval overhead upon use.
 *
 * Each function assumes the variant's type has already been determined and matches the function.
 * Invoking the function with a variant of a mismatched type has undefined behavior, and may lead to a segmentation fault.
 *
 * @param p_type The Variant type.
 *
 * @return A pointer to a type-specific function that returns a pointer to the internal value of a variant. Check the implementation of this function (gdextension_variant_get_ptr_internal_getter) for pointee type info of each variant type.
 */
typedef GDExtensionVariantGetInternalPtrFunc (*GDExtensionInterfaceGetVariantGetInternalPtrFunc)(GDExtensionVariantType p_type);

/**
 * @name variant_get_ptr_operator_evaluator
 * @since 4.1
 *
 * Gets a pointer to a function that can evaluate the given Variant operator on the given Variant types.
 *
 * @param p_operator The variant operator.
 * @param p_type_a The type of the first Variant.
 * @param p_type_b The type of the second Variant.
 *
 * @return A pointer to a function that can evaluate the given Variant operator on the given Variant types.
 */
typedef GDExtensionPtrOperatorEvaluator (*GDExtensionInterfaceVariantGetPtrOperatorEvaluator)(GDExtensionVariantOperator p_operator, GDExtensionVariantType p_type_a, GDExtensionVariantType p_type_b);

/**
 * @name variant_get_ptr_builtin_method
 * @since 4.1
 *
 * Gets a pointer to a function that can call a builtin method on a type of Variant.
 *
 * @param p_type The Variant type.
 * @param p_method A pointer to a StringName with the method name.
 * @param p_hash A hash representing the method signature.
 *
 * @return A pointer to a function that can call a builtin method on a type of Variant.
 */
typedef GDExtensionPtrBuiltInMethod (*GDExtensionInterfaceVariantGetPtrBuiltinMethod)(GDExtensionVariantType p_type, GDExtensionConstStringNamePtr p_method, GDExtensionInt p_hash);

/**
 * @name variant_get_ptr_constructor
 * @since 4.1
 *
 * Gets a pointer to a function that can call one of the constructors for a type of Variant.
 *
 * @param p_type The Variant type.
 * @param p_constructor The index of the constructor.
 *
 * @return A pointer to a function that can call one of the constructors for a type of Variant.
 */
typedef GDExtensionPtrConstructor (*GDExtensionInterfaceVariantGetPtrConstructor)(GDExtensionVariantType p_type, int32_t p_constructor);

/**
 * @name variant_get_ptr_destructor
 * @since 4.1
 *
 * Gets a pointer to a function than can call the destructor for a type of Variant.
 *
 * @param p_type The Variant type.
 *
 * @return A pointer to a function than can call the destructor for a type of Variant.
 */
typedef GDExtensionPtrDestructor (*GDExtensionInterfaceVariantGetPtrDestructor)(GDExtensionVariantType p_type);

/**
 * @name variant_construct
 * @since 4.1
 *
 * Constructs a Variant of the given type, using the first constructor that matches the given arguments.
 *
 * @param p_type The Variant type.
 * @param r_base A pointer to a Variant to store the constructed value.
 * @param p_args A pointer to a C array of Variant pointers representing the arguments for the constructor.
 * @param p_argument_count The number of arguments to pass to the constructor.
 * @param r_error A pointer the structure which will be updated with error information.
 */
typedef void (*GDExtensionInterfaceVariantConstruct)(GDExtensionVariantType p_type, GDExtensionUninitializedVariantPtr r_base, const GDExtensionConstVariantPtr *p_args, int32_t p_argument_count, GDExtensionCallError *r_error);

/**
 * @name variant_get_ptr_setter
 * @since 4.1
 *
 * Gets a pointer to a function that can call a member's setter on the given Variant type.
 *
 * @param p_type The Variant type.
 * @param p_member A pointer to a StringName with the member name.
 *
 * @return A pointer to a function that can call a member's setter on the given Variant type.
 */
typedef GDExtensionPtrSetter (*GDExtensionInterfaceVariantGetPtrSetter)(GDExtensionVariantType p_type, GDExtensionConstStringNamePtr p_member);

/**
 * @name variant_get_ptr_getter
 * @since 4.1
 *
 * Gets a pointer to a function that can call a member's getter on the given Variant type.
 *
 * @param p_type The Variant type.
 * @param p_member A pointer to a StringName with the member name.
 *
 * @return A pointer to a function that can call a member's getter on the given Variant type.
 */
typedef GDExtensionPtrGetter (*GDExtensionInterfaceVariantGetPtrGetter)(GDExtensionVariantType p_type, GDExtensionConstStringNamePtr p_member);

/**
 * @name variant_get_ptr_indexed_setter
 * @since 4.1
 *
 * Gets a pointer to a function that can set an index on the given Variant type.
 *
 * @param p_type The Variant type.
 *
 * @return A pointer to a function that can set an index on the given Variant type.
 */
typedef GDExtensionPtrIndexedSetter (*GDExtensionInterfaceVariantGetPtrIndexedSetter)(GDExtensionVariantType p_type);

/**
 * @name variant_get_ptr_indexed_getter
 * @since 4.1
 *
 * Gets a pointer to a function that can get an index on the given Variant type.
 *
 * @param p_type The Variant type.
 *
 * @return A pointer to a function that can get an index on the given Variant type.
 */
typedef GDExtensionPtrIndexedGetter (*GDExtensionInterfaceVariantGetPtrIndexedGetter)(GDExtensionVariantType p_type);

/**
 * @name variant_get_ptr_keyed_setter
 * @since 4.1
 *
 * Gets a pointer to a function that can set a key on the given Variant type.
 *
 * @param p_type The Variant type.
 *
 * @return A pointer to a function that can set a key on the given Variant type.
 */
typedef GDExtensionPtrKeyedSetter (*GDExtensionInterfaceVariantGetPtrKeyedSetter)(GDExtensionVariantType p_type);

/**
 * @name variant_get_ptr_keyed_getter
 * @since 4.1
 *
 * Gets a pointer to a function that can get a key on the given Variant type.
 *
 * @param p_type The Variant type.
 *
 * @return A pointer to a function that can get a key on the given Variant type.
 */
typedef GDExtensionPtrKeyedGetter (*GDExtensionInterfaceVariantGetPtrKeyedGetter)(GDExtensionVariantType p_type);

/**
 * @name variant_get_ptr_keyed_checker
 * @since 4.1
 *
 * Gets a pointer to a function that can check a key on the given Variant type.
 *
 * @param p_type The Variant type.
 *
 * @return A pointer to a function that can check a key on the given Variant type.
 */
typedef GDExtensionPtrKeyedChecker (*GDExtensionInterfaceVariantGetPtrKeyedChecker)(GDExtensionVariantType p_type);

/**
 * @name variant_get_constant_value
 * @since 4.1
 *
 * Gets the value of a constant from the given Variant type.
 *
 * @param p_type The Variant type.
 * @param p_constant A pointer to a StringName with the constant name.
 * @param r_ret A pointer to a Variant to store the value.
 */
typedef void (*GDExtensionInterfaceVariantGetConstantValue)(GDExtensionVariantType p_type, GDExtensionConstStringNamePtr p_constant, GDExtensionUninitializedVariantPtr r_ret);

/**
 * @name variant_get_ptr_utility_function
 * @since 4.1
 *
 * Gets a pointer to a function that can call a Variant utility function.
 *
 * @param p_function A pointer to a StringName with the function name.
 * @param p_hash A hash representing the function signature.
 *
 * @return A pointer to a function that can call a Variant utility function.
 */
typedef GDExtensionPtrUtilityFunction (*GDExtensionInterfaceVariantGetPtrUtilityFunction)(GDExtensionConstStringNamePtr p_function, GDExtensionInt p_hash);

/**
 * @name string_new_with_latin1_chars
 * @since 4.1
 *
 * Creates a String from a Latin-1 encoded C string.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a Latin-1 encoded C string (null terminated).
 */
typedef void (*GDExtensionInterfaceStringNewWithLatin1Chars)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents);

/**
 * @name string_new_with_utf8_chars
 * @since 4.1
 *
 * Creates a String from a UTF-8 encoded C string.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a UTF-8 encoded C string (null terminated).
 */
typedef void (*GDExtensionInterfaceStringNewWithUtf8Chars)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents);

/**
 * @name string_new_with_utf16_chars
 * @since 4.1
 *
 * Creates a String from a UTF-16 encoded C string.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a UTF-16 encoded C string (null terminated).
 */
typedef void (*GDExtensionInterfaceStringNewWithUtf16Chars)(GDExtensionUninitializedStringPtr r_dest, const char16_t *p_contents);

/**
 * @name string_new_with_utf32_chars
 * @since 4.1
 *
 * Creates a String from a UTF-32 encoded C string.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a UTF-32 encoded C string (null terminated).
 */
typedef void (*GDExtensionInterfaceStringNewWithUtf32Chars)(GDExtensionUninitializedStringPtr r_dest, const char32_t *p_contents);

/**
 * @name string_new_with_wide_chars
 * @since 4.1
 *
 * Creates a String from a wide C string.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a wide C string (null terminated).
 */
typedef void (*GDExtensionInterfaceStringNewWithWideChars)(GDExtensionUninitializedStringPtr r_dest, const wchar_t *p_contents);

/**
 * @name string_new_with_latin1_chars_and_len
 * @since 4.1
 *
 * Creates a String from a Latin-1 encoded C string with the given length.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a Latin-1 encoded C string.
 * @param p_size The number of characters (= number of bytes).
 */
typedef void (*GDExtensionInterfaceStringNewWithLatin1CharsAndLen)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents, GDExtensionInt p_size);

/**
 * @name string_new_with_utf8_chars_and_len
 * @since 4.1
 * @deprecated Deprecated in Godot 4.3. Use `string_new_with_utf8_chars_and_len2` instead.
 *
 * Creates a String from a UTF-8 encoded C string with the given length.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a UTF-8 encoded C string.
 * @param p_size The number of bytes (not code units).
 */
typedef void (*GDExtensionInterfaceStringNewWithUtf8CharsAndLen)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents, GDExtensionInt p_size);

/**
 * @name string_new_with_utf8_chars_and_len2
 * @since 4.3
 *
 * Creates a String from a UTF-8 encoded C string with the given length.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a UTF-8 encoded C string.
 * @param p_size The number of bytes (not code units).
 *
 * @return Error code signifying if the operation successful.
 */
typedef GDExtensionInt (*GDExtensionInterfaceStringNewWithUtf8CharsAndLen2)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents, GDExtensionInt p_size);

/**
 * @name string_new_with_utf16_chars_and_len
 * @since 4.1
 * @deprecated Deprecated in Godot 4.3. Use `string_new_with_utf16_chars_and_len2` instead.
 *
 * Creates a String from a UTF-16 encoded C string with the given length.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a UTF-16 encoded C string.
 * @param p_char_count The number of characters (not bytes).
 */
typedef void (*GDExtensionInterfaceStringNewWithUtf16CharsAndLen)(GDExtensionUninitializedStringPtr r_dest, const char16_t *p_contents, GDExtensionInt p_char_count);

/**
 * @name string_new_with_utf16_chars_and_len2
 * @since 4.3
 *
 * Creates a String from a UTF-16 encoded C string with the given length.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a UTF-16 encoded C string.
 * @param p_char_count The number of characters (not bytes).
 * @param p_default_little_endian If true, UTF-16 use little endian.
 *
 * @return Error code signifying if the operation successful.
 */
typedef GDExtensionInt (*GDExtensionInterfaceStringNewWithUtf16CharsAndLen2)(GDExtensionUninitializedStringPtr r_dest, const char16_t *p_contents, GDExtensionInt p_char_count, GDExtensionBool p_default_little_endian);

/**
 * @name string_new_with_utf32_chars_and_len
 * @since 4.1
 *
 * Creates a String from a UTF-32 encoded C string with the given length.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a UTF-32 encoded C string.
 * @param p_char_count The number of characters (not bytes).
 */
typedef void (*GDExtensionInterfaceStringNewWithUtf32CharsAndLen)(GDExtensionUninitializedStringPtr r_dest, const char32_t *p_contents, GDExtensionInt p_char_count);

/**
 * @name string_new_with_wide_chars_and_len
 * @since 4.1
 *
 * Creates a String from a wide C string with the given length.
 *
 * @param r_dest A pointer to a Variant to hold the newly created String.
 * @param p_contents A pointer to a wide C string.
 * @param p_char_count The number of characters (not bytes).
 */
typedef void (*GDExtensionInterfaceStringNewWithWideCharsAndLen)(GDExtensionUninitializedStringPtr r_dest, const wchar_t *p_contents, GDExtensionInt p_char_count);

/**
 * @name string_to_latin1_chars
 * @since 4.1
 *
 * Converts a String to a Latin-1 encoded C string.
 *
 * It doesn't write a null terminator.
 *
 * @param p_self A pointer to the String.
 * @param r_text A pointer to the buffer to hold the resulting data. If NULL is passed in, only the length will be computed.
 * @param p_max_write_length The maximum number of characters that can be written to r_text. It has no affect on the return value.
 *
 * @return The resulting encoded string length in characters (not bytes), not including a null terminator.
 */
typedef GDExtensionInt (*GDExtensionInterfaceStringToLatin1Chars)(GDExtensionConstStringPtr p_self, char *r_text, GDExtensionInt p_max_write_length);

/**
 * @name string_to_utf8_chars
 * @since 4.1
 *
 * Converts a String to a UTF-8 encoded C string.
 *
 * It doesn't write a null terminator.
 *
 * @param p_self A pointer to the String.
 * @param r_text A pointer to the buffer to hold the resulting data. If NULL is passed in, only the length will be computed.
 * @param p_max_write_length The maximum number of characters that can be written to r_text. It has no affect on the return value.
 *
 * @return The resulting encoded string length in characters (not bytes), not including a null terminator.
 */
typedef GDExtensionInt (*GDExtensionInterfaceStringToUtf8Chars)(GDExtensionConstStringPtr p_self, char *r_text, GDExtensionInt p_max_write_length);

/**
 * @name string_to_utf16_chars
 * @since 4.1
 *
 * Converts a String to a UTF-16 encoded C string.
 *
 * It doesn't write a null terminator.
 *
 * @param p_self A pointer to the String.
 * @param r_text A pointer to the buffer to hold the resulting data. If NULL is passed in, only the length will be computed.
 * @param p_max_write_length The maximum number of characters that can be written to r_text. It has no affect on the return value.
 *
 * @return The resulting encoded string length in characters (not bytes), not including a null terminator.
 */
typedef GDExtensionInt (*GDExtensionInterfaceStringToUtf16Chars)(GDExtensionConstStringPtr p_self, char16_t *r_text, GDExtensionInt p_max_write_length);

/**
 * @name string_to_utf32_chars
 * @since 4.1
 *
 * Converts a String to a UTF-32 encoded C string.
 *
 * It doesn't write a null terminator.
 *
 * @param p_self A pointer to the String.
 * @param r_text A pointer to the buffer to hold the resulting data. If NULL is passed in, only the length will be computed.
 * @param p_max_write_length The maximum number of characters that can be written to r_text. It has no affect on the return value.
 *
 * @return The resulting encoded string length in characters (not bytes), not including a null terminator.
 */
typedef GDExtensionInt (*GDExtensionInterfaceStringToUtf32Chars)(GDExtensionConstStringPtr p_self, char32_t *r_text, GDExtensionInt p_max_write_length);

/**
 * @name string_to_wide_chars
 * @since 4.1
 *
 * Converts a String to a wide C string.
 *
 * It doesn't write a null terminator.
 *
 * @param p_self A pointer to the String.
 * @param r_text A pointer to the buffer to hold the resulting data. If NULL is passed in, only the length will be computed.
 * @param p_max_write_length The maximum number of characters that can be written to r_text. It has no affect on the return value.
 *
 * @return The resulting encoded string length in characters (not bytes), not including a null terminator.
 */
typedef GDExtensionInt (*GDExtensionInterfaceStringToWideChars)(GDExtensionConstStringPtr p_self, wchar_t *r_text, GDExtensionInt p_max_write_length);

/**
 * @name string_operator_index
 * @since 4.1
 *
 * Gets a pointer to the character at the given index from a String.
 *
 * @param p_self A pointer to the String.
 * @param p_index The index.
 *
 * @return A pointer to the requested character.
 */
typedef char32_t *(*GDExtensionInterfaceStringOperatorIndex)(GDExtensionStringPtr p_self, GDExtensionInt p_index);

/**
 * @name string_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to the character at the given index from a String.
 *
 * @param p_self A pointer to the String.
 * @param p_index The index.
 *
 * @return A const pointer to the requested character.
 */
typedef const char32_t *(*GDExtensionInterfaceStringOperatorIndexConst)(GDExtensionConstStringPtr p_self, GDExtensionInt p_index);

/**
 * @name string_operator_plus_eq_string
 * @since 4.1
 *
 * Appends another String to a String.
 *
 * @param p_self A pointer to the String.
 * @param p_b A pointer to the other String to append.
 */
typedef void (*GDExtensionInterfaceStringOperatorPlusEqString)(GDExtensionStringPtr p_self, GDExtensionConstStringPtr p_b);

/**
 * @name string_operator_plus_eq_char
 * @since 4.1
 *
 * Appends a character to a String.
 *
 * @param p_self A pointer to the String.
 * @param p_b A pointer to the character to append.
 */
typedef void (*GDExtensionInterfaceStringOperatorPlusEqChar)(GDExtensionStringPtr p_self, char32_t p_b);

/**
 * @name string_operator_plus_eq_cstr
 * @since 4.1
 *
 * Appends a Latin-1 encoded C string to a String.
 *
 * @param p_self A pointer to the String.
 * @param p_b A pointer to a Latin-1 encoded C string (null terminated).
 */
typedef void (*GDExtensionInterfaceStringOperatorPlusEqCstr)(GDExtensionStringPtr p_self, const char *p_b);

/**
 * @name string_operator_plus_eq_wcstr
 * @since 4.1
 *
 * Appends a wide C string to a String.
 *
 * @param p_self A pointer to the String.
 * @param p_b A pointer to a wide C string (null terminated).
 */
typedef void (*GDExtensionInterfaceStringOperatorPlusEqWcstr)(GDExtensionStringPtr p_self, const wchar_t *p_b);

/**
 * @name string_operator_plus_eq_c32str
 * @since 4.1
 *
 * Appends a UTF-32 encoded C string to a String.
 *
 * @param p_self A pointer to the String.
 * @param p_b A pointer to a UTF-32 encoded C string (null terminated).
 */
typedef void (*GDExtensionInterfaceStringOperatorPlusEqC32str)(GDExtensionStringPtr p_self, const char32_t *p_b);

/**
 * @name string_resize
 * @since 4.2
 *
 * Resizes the underlying string data to the given number of characters.
 *
 * Space needs to be allocated for the null terminating character ('\0') which
 * also must be added manually, in order for all string functions to work correctly.
 *
 * Warning: This is an error-prone operation - only use it if there's no other
 * efficient way to accomplish your goal.
 *
 * @param p_self A pointer to the String.
 * @param p_resize The new length for the String.
 *
 * @return Error code signifying if the operation successful.
 */
typedef GDExtensionInt (*GDExtensionInterfaceStringResize)(GDExtensionStringPtr p_self, GDExtensionInt p_resize);

/**
 * @name string_name_new_with_latin1_chars
 * @since 4.2
 *
 * Creates a StringName from a Latin-1 encoded C string.
 *
 * If `p_is_static` is true, then:
 * - The StringName will reuse the `p_contents` buffer instead of copying it.
 * - You must guarantee that the buffer remains valid for the duration of the application (e.g. string literal).
 * - You must not call a destructor for this StringName. Incrementing the initial reference once should achieve this.
 *
 * `p_is_static` is purely an optimization and can easily introduce undefined behavior if used wrong. In case of doubt, set it to false.
 *
 * @param r_dest A pointer to uninitialized storage, into which the newly created StringName is constructed.
 * @param p_contents A pointer to a C string (null terminated and Latin-1 or ASCII encoded).
 * @param p_is_static Whether the StringName reuses the buffer directly (see above).
 */
typedef void (*GDExtensionInterfaceStringNameNewWithLatin1Chars)(GDExtensionUninitializedStringNamePtr r_dest, const char *p_contents, GDExtensionBool p_is_static);

/**
 * @name string_name_new_with_utf8_chars
 * @since 4.2
 *
 * Creates a StringName from a UTF-8 encoded C string.
 *
 * @param r_dest A pointer to uninitialized storage, into which the newly created StringName is constructed.
 * @param p_contents A pointer to a C string (null terminated and UTF-8 encoded).
 */
typedef void (*GDExtensionInterfaceStringNameNewWithUtf8Chars)(GDExtensionUninitializedStringNamePtr r_dest, const char *p_contents);

/**
 * @name string_name_new_with_utf8_chars_and_len
 * @since 4.2
 *
 * Creates a StringName from a UTF-8 encoded string with a given number of characters.
 *
 * @param r_dest A pointer to uninitialized storage, into which the newly created StringName is constructed.
 * @param p_contents A pointer to a C string (null terminated and UTF-8 encoded).
 * @param p_size The number of bytes (not UTF-8 code points).
 */
typedef void (*GDExtensionInterfaceStringNameNewWithUtf8CharsAndLen)(GDExtensionUninitializedStringNamePtr r_dest, const char *p_contents, GDExtensionInt p_size);

/**
 * @name xml_parser_open_buffer
 * @since 4.1
 *
 * Opens a raw XML buffer on an XMLParser instance.
 *
 * @param p_instance A pointer to an XMLParser object.
 * @param p_buffer A pointer to the buffer.
 * @param p_size The size of the buffer.
 *
 * @return A Godot error code (ex. OK, ERR_INVALID_DATA, etc).
 *
 * @see XMLParser::open_buffer()
 */
typedef GDExtensionInt (*GDExtensionInterfaceXmlParserOpenBuffer)(GDExtensionObjectPtr p_instance, const uint8_t *p_buffer, size_t p_size);

/**
 * @name file_access_store_buffer
 * @since 4.1
 *
 * Stores the given buffer using an instance of FileAccess.
 *
 * @param p_instance A pointer to a FileAccess object.
 * @param p_src A pointer to the buffer.
 * @param p_length The size of the buffer.
 *
 * @see FileAccess::store_buffer()
 */
typedef void (*GDExtensionInterfaceFileAccessStoreBuffer)(GDExtensionObjectPtr p_instance, const uint8_t *p_src, uint64_t p_length);

/**
 * @name file_access_get_buffer
 * @since 4.1
 *
 * Reads the next p_length bytes into the given buffer using an instance of FileAccess.
 *
 * @param p_instance A pointer to a FileAccess object.
 * @param p_dst A pointer to the buffer to store the data.
 * @param p_length The requested number of bytes to read.
 *
 * @return The actual number of bytes read (may be less than requested).
 */
typedef uint64_t (*GDExtensionInterfaceFileAccessGetBuffer)(GDExtensionConstObjectPtr p_instance, uint8_t *p_dst, uint64_t p_length);

/**
 * @name image_ptrw
 * @since 4.3
 *
 * Returns writable pointer to internal Image buffer.
 *
 * @param p_instance A pointer to a Image object.
 *
 * @return Pointer to internal Image buffer.
 *
 * @see Image::ptrw()
 */
typedef uint8_t *(*GDExtensionInterfaceImagePtrw)(GDExtensionObjectPtr p_instance);

/**
 * @name image_ptr
 * @since 4.3
 *
 * Returns read only pointer to internal Image buffer.
 *
 * @param p_instance A pointer to a Image object.
 *
 * @return Pointer to internal Image buffer.
 *
 * @see Image::ptr()
 */
typedef const uint8_t *(*GDExtensionInterfaceImagePtr)(GDExtensionObjectPtr p_instance);

/**
 * @name worker_thread_pool_add_native_group_task
 * @since 4.1
 *
 * Adds a group task to an instance of WorkerThreadPool.
 *
 * @param p_instance A pointer to a WorkerThreadPool object.
 * @param p_func A pointer to a function to run in the thread pool.
 * @param p_userdata A pointer to arbitrary data which will be passed to p_func.
 * @param p_elements The number of element needed in the group.
 * @param p_tasks The number of tasks needed in the group.
 * @param p_high_priority Whether or not this is a high priority task.
 * @param p_description A pointer to a String with the task description.
 *
 * @return The task group ID.
 *
 * @see WorkerThreadPool::add_group_task()
 */
typedef int64_t (*GDExtensionInterfaceWorkerThreadPoolAddNativeGroupTask)(GDExtensionObjectPtr p_instance, GDExtensionWorkerThreadPoolGroupTask p_func, void *p_userdata, int32_t p_elements, int32_t p_tasks, GDExtensionBool p_high_priority, GDExtensionConstStringPtr p_description);

/**
 * @name worker_thread_pool_add_native_task
 * @since 4.1
 *
 * Adds a task to an instance of WorkerThreadPool.
 *
 * @param p_instance A pointer to a WorkerThreadPool object.
 * @param p_func A pointer to a function to run in the thread pool.
 * @param p_userdata A pointer to arbitrary data which will be passed to p_func.
 * @param p_high_priority Whether or not this is a high priority task.
 * @param p_description A pointer to a String with the task description.
 *
 * @return The task ID.
 */
typedef int64_t (*GDExtensionInterfaceWorkerThreadPoolAddNativeTask)(GDExtensionObjectPtr p_instance, GDExtensionWorkerThreadPoolTask p_func, void *p_userdata, GDExtensionBool p_high_priority, GDExtensionConstStringPtr p_description);

/**
 * @name packed_byte_array_operator_index
 * @since 4.1
 *
 * Gets a pointer to a byte in a PackedByteArray.
 *
 * @param p_self A pointer to a PackedByteArray object.
 * @param p_index The index of the byte to get.
 *
 * @return A pointer to the requested byte.
 */
typedef uint8_t *(*GDExtensionInterfacePackedByteArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_byte_array_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a byte in a PackedByteArray.
 *
 * @param p_self A const pointer to a PackedByteArray object.
 * @param p_index The index of the byte to get.
 *
 * @return A const pointer to the requested byte.
 */
typedef const uint8_t *(*GDExtensionInterfacePackedByteArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_float32_array_operator_index
 * @since 4.1
 *
 * Gets a pointer to a 32-bit float in a PackedFloat32Array.
 *
 * @param p_self A pointer to a PackedFloat32Array object.
 * @param p_index The index of the float to get.
 *
 * @return A pointer to the requested 32-bit float.
 */
typedef float *(*GDExtensionInterfacePackedFloat32ArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_float32_array_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a 32-bit float in a PackedFloat32Array.
 *
 * @param p_self A const pointer to a PackedFloat32Array object.
 * @param p_index The index of the float to get.
 *
 * @return A const pointer to the requested 32-bit float.
 */
typedef const float *(*GDExtensionInterfacePackedFloat32ArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_float64_array_operator_index
 * @since 4.1
 *
 * Gets a pointer to a 64-bit float in a PackedFloat64Array.
 *
 * @param p_self A pointer to a PackedFloat64Array object.
 * @param p_index The index of the float to get.
 *
 * @return A pointer to the requested 64-bit float.
 */
typedef double *(*GDExtensionInterfacePackedFloat64ArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_float64_array_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a 64-bit float in a PackedFloat64Array.
 *
 * @param p_self A const pointer to a PackedFloat64Array object.
 * @param p_index The index of the float to get.
 *
 * @return A const pointer to the requested 64-bit float.
 */
typedef const double *(*GDExtensionInterfacePackedFloat64ArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_int32_array_operator_index
 * @since 4.1
 *
 * Gets a pointer to a 32-bit integer in a PackedInt32Array.
 *
 * @param p_self A pointer to a PackedInt32Array object.
 * @param p_index The index of the integer to get.
 *
 * @return A pointer to the requested 32-bit integer.
 */
typedef int32_t *(*GDExtensionInterfacePackedInt32ArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_int32_array_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a 32-bit integer in a PackedInt32Array.
 *
 * @param p_self A const pointer to a PackedInt32Array object.
 * @param p_index The index of the integer to get.
 *
 * @return A const pointer to the requested 32-bit integer.
 */
typedef const int32_t *(*GDExtensionInterfacePackedInt32ArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_int64_array_operator_index
 * @since 4.1
 *
 * Gets a pointer to a 64-bit integer in a PackedInt64Array.
 *
 * @param p_self A pointer to a PackedInt64Array object.
 * @param p_index The index of the integer to get.
 *
 * @return A pointer to the requested 64-bit integer.
 */
typedef int64_t *(*GDExtensionInterfacePackedInt64ArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_int64_array_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a 64-bit integer in a PackedInt64Array.
 *
 * @param p_self A const pointer to a PackedInt64Array object.
 * @param p_index The index of the integer to get.
 *
 * @return A const pointer to the requested 64-bit integer.
 */
typedef const int64_t *(*GDExtensionInterfacePackedInt64ArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_string_array_operator_index
 * @since 4.1
 *
 * Gets a pointer to a string in a PackedStringArray.
 *
 * @param p_self A pointer to a PackedStringArray object.
 * @param p_index The index of the String to get.
 *
 * @return A pointer to the requested String.
 */
typedef GDExtensionStringPtr (*GDExtensionInterfacePackedStringArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_string_array_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a string in a PackedStringArray.
 *
 * @param p_self A const pointer to a PackedStringArray object.
 * @param p_index The index of the String to get.
 *
 * @return A const pointer to the requested String.
 */
typedef GDExtensionStringPtr (*GDExtensionInterfacePackedStringArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_vector2_array_operator_index
 * @since 4.1
 *
 * Gets a pointer to a Vector2 in a PackedVector2Array.
 *
 * @param p_self A pointer to a PackedVector2Array object.
 * @param p_index The index of the Vector2 to get.
 *
 * @return A pointer to the requested Vector2.
 */
typedef GDExtensionTypePtr (*GDExtensionInterfacePackedVector2ArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_vector2_array_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a Vector2 in a PackedVector2Array.
 *
 * @param p_self A const pointer to a PackedVector2Array object.
 * @param p_index The index of the Vector2 to get.
 *
 * @return A const pointer to the requested Vector2.
 */
typedef GDExtensionTypePtr (*GDExtensionInterfacePackedVector2ArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_vector3_array_operator_index
 * @since 4.1
 *
 * Gets a pointer to a Vector3 in a PackedVector3Array.
 *
 * @param p_self A pointer to a PackedVector3Array object.
 * @param p_index The index of the Vector3 to get.
 *
 * @return A pointer to the requested Vector3.
 */
typedef GDExtensionTypePtr (*GDExtensionInterfacePackedVector3ArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_vector3_array_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a Vector3 in a PackedVector3Array.
 *
 * @param p_self A const pointer to a PackedVector3Array object.
 * @param p_index The index of the Vector3 to get.
 *
 * @return A const pointer to the requested Vector3.
 */
typedef GDExtensionTypePtr (*GDExtensionInterfacePackedVector3ArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_vector4_array_operator_index
 * @since 4.3
 *
 * Gets a pointer to a Vector4 in a PackedVector4Array.
 *
 * @param p_self A pointer to a PackedVector4Array object.
 * @param p_index The index of the Vector4 to get.
 *
 * @return A pointer to the requested Vector4.
 */
typedef GDExtensionTypePtr (*GDExtensionInterfacePackedVector4ArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_vector4_array_operator_index_const
 * @since 4.3
 *
 * Gets a const pointer to a Vector4 in a PackedVector4Array.
 *
 * @param p_self A const pointer to a PackedVector4Array object.
 * @param p_index The index of the Vector4 to get.
 *
 * @return A const pointer to the requested Vector4.
 */
typedef GDExtensionTypePtr (*GDExtensionInterfacePackedVector4ArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_color_array_operator_index
 * @since 4.1
 *
 * Gets a pointer to a color in a PackedColorArray.
 *
 * @param p_self A pointer to a PackedColorArray object.
 * @param p_index The index of the Color to get.
 *
 * @return A pointer to the requested Color.
 */
typedef GDExtensionTypePtr (*GDExtensionInterfacePackedColorArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name packed_color_array_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a color in a PackedColorArray.
 *
 * @param p_self A const pointer to a PackedColorArray object.
 * @param p_index The index of the Color to get.
 *
 * @return A const pointer to the requested Color.
 */
typedef GDExtensionTypePtr (*GDExtensionInterfacePackedColorArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name array_operator_index
 * @since 4.1
 *
 * Gets a pointer to a Variant in an Array.
 *
 * @param p_self A pointer to an Array object.
 * @param p_index The index of the Variant to get.
 *
 * @return A pointer to the requested Variant.
 */
typedef GDExtensionVariantPtr (*GDExtensionInterfaceArrayOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionInt p_index);

/**
 * @name array_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a Variant in an Array.
 *
 * @param p_self A const pointer to an Array object.
 * @param p_index The index of the Variant to get.
 *
 * @return A const pointer to the requested Variant.
 */
typedef GDExtensionVariantPtr (*GDExtensionInterfaceArrayOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionInt p_index);

/**
 * @name array_ref
 * @since 4.1
 * @deprecated Deprecated in Godot 4.5. Removed from interface. Use copy constructor instead.
 *
 * Sets an Array to be a reference to another Array object.
 *
 * @param p_self A pointer to the Array object to update.
 * @param p_from A pointer to the Array object to reference.
 */
typedef void (*GDExtensionInterfaceArrayRef)(GDExtensionTypePtr p_self, GDExtensionConstTypePtr p_from);

/**
 * @name array_set_typed
 * @since 4.1
 *
 * Makes an Array into a typed Array.
 *
 * @param p_self A pointer to the Array.
 * @param p_type The type of Variant the Array will store.
 * @param p_class_name A pointer to a StringName with the name of the object (if p_type is GDEXTENSION_VARIANT_TYPE_OBJECT).
 * @param p_script A pointer to a Script object (if p_type is GDEXTENSION_VARIANT_TYPE_OBJECT and the base class is extended by a script).
 */
typedef void (*GDExtensionInterfaceArraySetTyped)(GDExtensionTypePtr p_self, GDExtensionVariantType p_type, GDExtensionConstStringNamePtr p_class_name, GDExtensionConstVariantPtr p_script);

/**
 * @name dictionary_operator_index
 * @since 4.1
 *
 * Gets a pointer to a Variant in a Dictionary with the given key.
 *
 * @param p_self A pointer to a Dictionary object.
 * @param p_key A pointer to a Variant representing the key.
 *
 * @return A pointer to a Variant representing the value at the given key.
 */
typedef GDExtensionVariantPtr (*GDExtensionInterfaceDictionaryOperatorIndex)(GDExtensionTypePtr p_self, GDExtensionConstVariantPtr p_key);

/**
 * @name dictionary_operator_index_const
 * @since 4.1
 *
 * Gets a const pointer to a Variant in a Dictionary with the given key.
 *
 * @param p_self A const pointer to a Dictionary object.
 * @param p_key A pointer to a Variant representing the key.
 *
 * @return A const pointer to a Variant representing the value at the given key.
 */
typedef GDExtensionVariantPtr (*GDExtensionInterfaceDictionaryOperatorIndexConst)(GDExtensionConstTypePtr p_self, GDExtensionConstVariantPtr p_key);

/**
 * @name dictionary_set_typed
 * @since 4.4
 *
 * Makes a Dictionary into a typed Dictionary.
 *
 * @param p_self A pointer to the Dictionary.
 * @param p_key_type The type of Variant the Dictionary key will store.
 * @param p_key_class_name A pointer to a StringName with the name of the object (if p_key_type is GDEXTENSION_VARIANT_TYPE_OBJECT).
 * @param p_key_script A pointer to a Script object (if p_key_type is GDEXTENSION_VARIANT_TYPE_OBJECT and the base class is extended by a script).
 * @param p_value_type The type of Variant the Dictionary value will store.
 * @param p_value_class_name A pointer to a StringName with the name of the object (if p_value_type is GDEXTENSION_VARIANT_TYPE_OBJECT).
 * @param p_value_script A pointer to a Script object (if p_value_type is GDEXTENSION_VARIANT_TYPE_OBJECT and the base class is extended by a script).
 */
typedef void (*GDExtensionInterfaceDictionarySetTyped)(GDExtensionTypePtr p_self, GDExtensionVariantType p_key_type, GDExtensionConstStringNamePtr p_key_class_name, GDExtensionConstVariantPtr p_key_script, GDExtensionVariantType p_value_type, GDExtensionConstStringNamePtr p_value_class_name, GDExtensionConstVariantPtr p_value_script);

/**
 * @name object_method_bind_call
 * @since 4.1
 *
 * Calls a method on an Object.
 *
 * @param p_method_bind A pointer to the MethodBind representing the method on the Object's class.
 * @param p_instance A pointer to the Object.
 * @param p_args A pointer to a C array of Variants representing the arguments.
 * @param p_arg_count The number of arguments.
 * @param r_ret A pointer to Variant which will receive the return value.
 * @param r_error A pointer to a GDExtensionCallError struct that will receive error information.
 */
typedef void (*GDExtensionInterfaceObjectMethodBindCall)(GDExtensionMethodBindPtr p_method_bind, GDExtensionObjectPtr p_instance, const GDExtensionConstVariantPtr *p_args, GDExtensionInt p_arg_count, GDExtensionUninitializedVariantPtr r_ret, GDExtensionCallError *r_error);

/**
 * @name object_method_bind_ptrcall
 * @since 4.1
 *
 * Calls a method on an Object (using a "ptrcall").
 *
 * @param p_method_bind A pointer to the MethodBind representing the method on the Object's class.
 * @param p_instance A pointer to the Object.
 * @param p_args A pointer to a C array representing the arguments.
 * @param r_ret A pointer to the Object that will receive the return value.
 */
typedef void (*GDExtensionInterfaceObjectMethodBindPtrcall)(GDExtensionMethodBindPtr p_method_bind, GDExtensionObjectPtr p_instance, const GDExtensionConstTypePtr *p_args, GDExtensionTypePtr r_ret);

/**
 * @name object_destroy
 * @since 4.1
 *
 * Destroys an Object.
 *
 * @param p_o A pointer to the Object.
 */
typedef void (*GDExtensionInterfaceObjectDestroy)(GDExtensionObjectPtr p_o);

/**
 * @name global_get_singleton
 * @since 4.1
 *
 * Gets a global singleton by name.
 *
 * @param p_name A pointer to a StringName with the singleton name.
 *
 * @return A pointer to the singleton Object.
 */
typedef GDExtensionObjectPtr (*GDExtensionInterfaceGlobalGetSingleton)(GDExtensionConstStringNamePtr p_name);

/**
 * @name object_get_instance_binding
 * @since 4.1
 *
 * Gets a pointer representing an Object's instance binding.
 *
 * @param p_o A pointer to the Object.
 * @param p_token A token the library received by the GDExtension's entry point function.
 * @param p_callbacks A pointer to a GDExtensionInstanceBindingCallbacks struct.
 *
 * @return A pointer to the instance binding.
 */
typedef void *(*GDExtensionInterfaceObjectGetInstanceBinding)(GDExtensionObjectPtr p_o, void *p_token, const GDExtensionInstanceBindingCallbacks *p_callbacks);

/**
 * @name object_set_instance_binding
 * @since 4.1
 *
 * Sets an Object's instance binding.
 *
 * @param p_o A pointer to the Object.
 * @param p_token A token the library received by the GDExtension's entry point function.
 * @param p_binding A pointer to the instance binding.
 * @param p_callbacks A pointer to a GDExtensionInstanceBindingCallbacks struct.
 */
typedef void (*GDExtensionInterfaceObjectSetInstanceBinding)(GDExtensionObjectPtr p_o, void *p_token, void *p_binding, const GDExtensionInstanceBindingCallbacks *p_callbacks);

/**
 * @name object_free_instance_binding
 * @since 4.2
 *
 * Free an Object's instance binding.
 *
 * @param p_o A pointer to the Object.
 * @param p_token A token the library received by the GDExtension's entry point function.
 */
typedef void (*GDExtensionInterfaceObjectFreeInstanceBinding)(GDExtensionObjectPtr p_o, void *p_token);

/**
 * @name object_set_instance
 * @since 4.1
 *
 * Sets an extension class instance on a Object.
 *
 * `p_classname` should be a registered extension class and should extend the `p_o` Object's class.
 *
 * @param p_o A pointer to the Object.
 * @param p_classname A pointer to a StringName with the registered extension class's name.
 * @param p_instance A pointer to the extension class instance.
 */
typedef void (*GDExtensionInterfaceObjectSetInstance)(GDExtensionObjectPtr p_o, GDExtensionConstStringNamePtr p_classname, GDExtensionClassInstancePtr p_instance);

/**
 * @name object_get_class_name
 * @since 4.1
 *
 * Gets the class name of an Object.
 *
 * If the GDExtension wraps the Godot object in an abstraction specific to its class, this is the
 * function that should be used to determine which wrapper to use.
 *
 * @param p_object A pointer to the Object.
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param r_class_name A pointer to a String to receive the class name.
 *
 * @return true if successful in getting the class name; otherwise false.
 */
typedef GDExtensionBool (*GDExtensionInterfaceObjectGetClassName)(GDExtensionConstObjectPtr p_object, GDExtensionClassLibraryPtr p_library, GDExtensionUninitializedStringNamePtr r_class_name);

/**
 * @name object_cast_to
 * @since 4.1
 *
 * Casts an Object to a different type.
 *
 * @param p_object A pointer to the Object.
 * @param p_class_tag A pointer uniquely identifying a built-in class in the ClassDB.
 *
 * @return Returns a pointer to the Object, or NULL if it can't be cast to the requested type.
 */
typedef GDExtensionObjectPtr (*GDExtensionInterfaceObjectCastTo)(GDExtensionConstObjectPtr p_object, void *p_class_tag);

/**
 * @name object_get_instance_from_id
 * @since 4.1
 *
 * Gets an Object by its instance ID.
 *
 * @param p_instance_id The instance ID.
 *
 * @return A pointer to the Object.
 */
typedef GDExtensionObjectPtr (*GDExtensionInterfaceObjectGetInstanceFromId)(GDObjectInstanceID p_instance_id);

/**
 * @name object_get_instance_id
 * @since 4.1
 *
 * Gets the instance ID from an Object.
 *
 * @param p_object A pointer to the Object.
 *
 * @return The instance ID.
 */
typedef GDObjectInstanceID (*GDExtensionInterfaceObjectGetInstanceId)(GDExtensionConstObjectPtr p_object);

/**
 * @name object_has_script_method
 * @since 4.3
 *
 * Checks if this object has a script with the given method.
 *
 * @param p_object A pointer to the Object.
 * @param p_method A pointer to a StringName identifying the method.
 *
 * @return true if the object has a script and that script has a method with the given name. Returns false if the object has no script.
 */
typedef GDExtensionBool (*GDExtensionInterfaceObjectHasScriptMethod)(GDExtensionConstObjectPtr p_object, GDExtensionConstStringNamePtr p_method);

/**
 * @name object_call_script_method
 * @since 4.3
 *
 * Call the given script method on this object.
 *
 * @param p_object A pointer to the Object.
 * @param p_method A pointer to a StringName identifying the method.
 * @param p_args A pointer to a C array of Variant.
 * @param p_argument_count The number of arguments.
 * @param r_return A pointer a Variant which will be assigned the return value.
 * @param r_error A pointer the structure which will hold error information.
 */
typedef void (*GDExtensionInterfaceObjectCallScriptMethod)(GDExtensionObjectPtr p_object, GDExtensionConstStringNamePtr p_method, const GDExtensionConstVariantPtr *p_args, GDExtensionInt p_argument_count, GDExtensionUninitializedVariantPtr r_return, GDExtensionCallError *r_error);

/**
 * @name ref_get_object
 * @since 4.1
 *
 * Gets the Object from a reference.
 *
 * @param p_ref A pointer to the reference.
 *
 * @return A pointer to the Object from the reference or NULL.
 */
typedef GDExtensionObjectPtr (*GDExtensionInterfaceRefGetObject)(GDExtensionConstRefPtr p_ref);

/**
 * @name ref_set_object
 * @since 4.1
 *
 * Sets the Object referred to by a reference.
 *
 * @param p_ref A pointer to the reference.
 * @param p_object A pointer to the Object to refer to.
 */
typedef void (*GDExtensionInterfaceRefSetObject)(GDExtensionRefPtr p_ref, GDExtensionObjectPtr p_object);

/**
 * @name script_instance_create
 * @since 4.1
 * @deprecated Deprecated in Godot 4.2. Use `script_instance_create3` instead.
 *
 * Creates a script instance that contains the given info and instance data.
 *
 * @param p_info A pointer to a GDExtensionScriptInstanceInfo struct.
 * @param p_instance_data A pointer to a data representing the script instance in the GDExtension. This will be passed to all the function pointers on p_info.
 *
 * @return A pointer to a ScriptInstanceExtension object.
 */
typedef GDExtensionScriptInstancePtr (*GDExtensionInterfaceScriptInstanceCreate)(const GDExtensionScriptInstanceInfo *p_info, GDExtensionScriptInstanceDataPtr p_instance_data);

/**
 * @name script_instance_create2
 * @since 4.2
 * @deprecated Deprecated in Godot 4.3. Use `script_instance_create3` instead.
 *
 * Creates a script instance that contains the given info and instance data.
 *
 * @param p_info A pointer to a GDExtensionScriptInstanceInfo2 struct.
 * @param p_instance_data A pointer to a data representing the script instance in the GDExtension. This will be passed to all the function pointers on p_info.
 *
 * @return A pointer to a ScriptInstanceExtension object.
 */
typedef GDExtensionScriptInstancePtr (*GDExtensionInterfaceScriptInstanceCreate2)(const GDExtensionScriptInstanceInfo2 *p_info, GDExtensionScriptInstanceDataPtr p_instance_data);

/**
 * @name script_instance_create3
 * @since 4.3
 *
 * Creates a script instance that contains the given info and instance data.
 *
 * @param p_info A pointer to a GDExtensionScriptInstanceInfo3 struct.
 * @param p_instance_data A pointer to a data representing the script instance in the GDExtension. This will be passed to all the function pointers on p_info.
 *
 * @return A pointer to a ScriptInstanceExtension object.
 */
typedef GDExtensionScriptInstancePtr (*GDExtensionInterfaceScriptInstanceCreate3)(const GDExtensionScriptInstanceInfo3 *p_info, GDExtensionScriptInstanceDataPtr p_instance_data);

/**
 * @name placeholder_script_instance_create
 * @since 4.2
 *
 * Creates a placeholder script instance for a given script and instance.
 *
 * This interface is optional as a custom placeholder could also be created with script_instance_create().
 *
 * @param p_language A pointer to a ScriptLanguage.
 * @param p_script A pointer to a Script.
 * @param p_owner A pointer to an Object.
 *
 * @return A pointer to a PlaceHolderScriptInstance object.
 */
typedef GDExtensionScriptInstancePtr (*GDExtensionInterfacePlaceHolderScriptInstanceCreate)(GDExtensionObjectPtr p_language, GDExtensionObjectPtr p_script, GDExtensionObjectPtr p_owner);

/**
 * @name placeholder_script_instance_update
 * @since 4.2
 *
 * Updates a placeholder script instance with the given properties and values.
 *
 * The passed in placeholder must be an instance of PlaceHolderScriptInstance
 * such as the one returned by placeholder_script_instance_create().
 *
 * @param p_placeholder A pointer to a PlaceHolderScriptInstance.
 * @param p_properties A pointer to an Array of Dictionary representing PropertyInfo.
 * @param p_values A pointer to a Dictionary mapping StringName to Variant values.
 */
typedef void (*GDExtensionInterfacePlaceHolderScriptInstanceUpdate)(GDExtensionScriptInstancePtr p_placeholder, GDExtensionConstTypePtr p_properties, GDExtensionConstTypePtr p_values);

/**
 * @name object_get_script_instance
 * @since 4.2
 *
 * Get the script instance data attached to this object.
 *
 * @param p_object A pointer to the Object.
 * @param p_language A pointer to the language expected for this script instance.
 *
 * @return A GDExtensionScriptInstanceDataPtr that was attached to this object as part of script_instance_create.
 */
typedef GDExtensionScriptInstanceDataPtr (*GDExtensionInterfaceObjectGetScriptInstance)(GDExtensionConstObjectPtr p_object, GDExtensionObjectPtr p_language);

/**
 * @name object_set_script_instance
 * @since 4.5
 *
 * Set the script instance data attached to this object.
 *
 * @param p_object A pointer to the Object.
 * @param p_script_instance A pointer to the script instance data to attach to this object.
 */
typedef void (*GDExtensionInterfaceObjectSetScriptInstance)(GDExtensionObjectPtr p_object, GDExtensionScriptInstanceDataPtr p_script_instance);

/**
 * @name callable_custom_create
 * @since 4.2
 * @deprecated Deprecated in Godot 4.3. Use `callable_custom_create2` instead.
 *
 * Creates a custom Callable object from a function pointer.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param r_callable A pointer that will receive the new Callable.
 * @param p_callable_custom_info The info required to construct a Callable.
 */
typedef void (*GDExtensionInterfaceCallableCustomCreate)(GDExtensionUninitializedTypePtr r_callable, GDExtensionCallableCustomInfo *p_callable_custom_info);

/**
 * @name callable_custom_create2
 * @since 4.3
 *
 * Creates a custom Callable object from a function pointer.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param r_callable A pointer that will receive the new Callable.
 * @param p_callable_custom_info The info required to construct a Callable.
 */
typedef void (*GDExtensionInterfaceCallableCustomCreate2)(GDExtensionUninitializedTypePtr r_callable, GDExtensionCallableCustomInfo2 *p_callable_custom_info);

/**
 * @name callable_custom_get_userdata
 * @since 4.2
 *
 * Retrieves the userdata pointer from a custom Callable.
 *
 * If the Callable is not a custom Callable or the token does not match the one provided to callable_custom_create() via GDExtensionCallableCustomInfo then NULL will be returned.
 *
 * @param p_callable A pointer to a Callable.
 * @param p_token A pointer to an address that uniquely identifies the GDExtension.
 *
 * @return The userdata pointer given when creating this custom Callable.
 */
typedef void *(*GDExtensionInterfaceCallableCustomGetUserData)(GDExtensionConstTypePtr p_callable, void *p_token);

/**
 * @name classdb_construct_object
 * @since 4.1
 * @deprecated Deprecated in Godot 4.4. Use `classdb_construct_object2` instead.
 *
 * Constructs an Object of the requested class.
 *
 * The passed class must be a built-in godot class, or an already-registered extension class. In both cases, object_set_instance() should be called to fully initialize the object.
 *
 * @param p_classname A pointer to a StringName with the class name.
 *
 * @return A pointer to the newly created Object.
 */
typedef GDExtensionObjectPtr (*GDExtensionInterfaceClassdbConstructObject)(GDExtensionConstStringNamePtr p_classname);

/**
 * @name classdb_construct_object2
 * @since 4.4
 *
 * Constructs an Object of the requested class.
 *
 * The passed class must be a built-in godot class, or an already-registered extension class. In both cases, object_set_instance() should be called to fully initialize the object.
 *
 * "NOTIFICATION_POSTINITIALIZE" must be sent after construction.
 *
 * @param p_classname A pointer to a StringName with the class name.
 *
 * @return A pointer to the newly created Object.
 */
typedef GDExtensionObjectPtr (*GDExtensionInterfaceClassdbConstructObject2)(GDExtensionConstStringNamePtr p_classname);

/**
 * @name classdb_get_method_bind
 * @since 4.1
 *
 * Gets a pointer to the MethodBind in ClassDB for the given class, method and hash.
 *
 * @param p_classname A pointer to a StringName with the class name.
 * @param p_methodname A pointer to a StringName with the method name.
 * @param p_hash A hash representing the function signature.
 *
 * @return A pointer to the MethodBind from ClassDB.
 */
typedef GDExtensionMethodBindPtr (*GDExtensionInterfaceClassdbGetMethodBind)(GDExtensionConstStringNamePtr p_classname, GDExtensionConstStringNamePtr p_methodname, GDExtensionInt p_hash);

/**
 * @name classdb_get_class_tag
 * @since 4.1
 *
 * Gets a pointer uniquely identifying the given built-in class in the ClassDB.
 *
 * @param p_classname A pointer to a StringName with the class name.
 *
 * @return A pointer uniquely identifying the built-in class in the ClassDB.
 */
typedef void *(*GDExtensionInterfaceClassdbGetClassTag)(GDExtensionConstStringNamePtr p_classname);

/**
 * @name classdb_register_extension_class
 * @since 4.1
 * @deprecated Deprecated in Godot 4.2. Use `classdb_register_extension_class5` instead.
 *
 * Registers an extension class in the ClassDB.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_parent_class_name A pointer to a StringName with the parent class name.
 * @param p_extension_funcs A pointer to a GDExtensionClassCreationInfo struct.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClass)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, GDExtensionConstStringNamePtr p_parent_class_name, const GDExtensionClassCreationInfo *p_extension_funcs);

/**
 * @name classdb_register_extension_class2
 * @since 4.2
 * @deprecated Deprecated in Godot 4.3. Use `classdb_register_extension_class5` instead.
 *
 * Registers an extension class in the ClassDB.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_parent_class_name A pointer to a StringName with the parent class name.
 * @param p_extension_funcs A pointer to a GDExtensionClassCreationInfo2 struct.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClass2)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, GDExtensionConstStringNamePtr p_parent_class_name, const GDExtensionClassCreationInfo2 *p_extension_funcs);

/**
 * @name classdb_register_extension_class3
 * @since 4.3
 * @deprecated Deprecated in Godot 4.4. Use `classdb_register_extension_class5` instead.
 *
 * Registers an extension class in the ClassDB.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_parent_class_name A pointer to a StringName with the parent class name.
 * @param p_extension_funcs A pointer to a GDExtensionClassCreationInfo3 struct.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClass3)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, GDExtensionConstStringNamePtr p_parent_class_name, const GDExtensionClassCreationInfo3 *p_extension_funcs);

/**
 * @name classdb_register_extension_class4
 * @since 4.4
 * @deprecated Deprecated in Godot 4.5. Use `classdb_register_extension_class5` instead.
 *
 * Registers an extension class in the ClassDB.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_parent_class_name A pointer to a StringName with the parent class name.
 * @param p_extension_funcs A pointer to a GDExtensionClassCreationInfo4 struct.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClass4)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, GDExtensionConstStringNamePtr p_parent_class_name, const GDExtensionClassCreationInfo4 *p_extension_funcs);

/**
 * @name classdb_register_extension_class5
 * @since 4.5
 *
 * Registers an extension class in the ClassDB.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_parent_class_name A pointer to a StringName with the parent class name.
 * @param p_extension_funcs A pointer to a GDExtensionClassCreationInfo5 struct.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClass5)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, GDExtensionConstStringNamePtr p_parent_class_name, const GDExtensionClassCreationInfo5 *p_extension_funcs);

/**
 * @name classdb_register_extension_class_method
 * @since 4.1
 *
 * Registers a method on an extension class in the ClassDB.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_method_info A pointer to a GDExtensionClassMethodInfo struct.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClassMethod)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, const GDExtensionClassMethodInfo *p_method_info);

/**
 * @name classdb_register_extension_class_virtual_method
 * @since 4.3
 *
 * Registers a virtual method on an extension class in ClassDB, that can be implemented by scripts or other extensions.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_method_info A pointer to a GDExtensionClassMethodInfo struct.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClassVirtualMethod)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, const GDExtensionClassVirtualMethodInfo *p_method_info);

/**
 * @name classdb_register_extension_class_integer_constant
 * @since 4.1
 *
 * Registers an integer constant on an extension class in the ClassDB.
 *
 * Note about registering bitfield values (if p_is_bitfield is true): even though p_constant_value is signed, language bindings are
 * advised to treat bitfields as uint64_t, since this is generally clearer and can prevent mistakes like using -1 for setting all bits.
 * Language APIs should thus provide an abstraction that registers bitfields (uint64_t) separately from regular constants (int64_t).
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_enum_name A pointer to a StringName with the enum name.
 * @param p_constant_name A pointer to a StringName with the constant name.
 * @param p_constant_value The constant value.
 * @param p_is_bitfield Whether or not this constant is part of a bitfield.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClassIntegerConstant)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, GDExtensionConstStringNamePtr p_enum_name, GDExtensionConstStringNamePtr p_constant_name, GDExtensionInt p_constant_value, GDExtensionBool p_is_bitfield);

/**
 * @name classdb_register_extension_class_property
 * @since 4.1
 *
 * Registers a property on an extension class in the ClassDB.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_info A pointer to a GDExtensionPropertyInfo struct.
 * @param p_setter A pointer to a StringName with the name of the setter method.
 * @param p_getter A pointer to a StringName with the name of the getter method.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClassProperty)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, const GDExtensionPropertyInfo *p_info, GDExtensionConstStringNamePtr p_setter, GDExtensionConstStringNamePtr p_getter);

/**
 * @name classdb_register_extension_class_property_indexed
 * @since 4.2
 *
 * Registers an indexed property on an extension class in the ClassDB.
 *
 * Provided struct can be safely freed once the function returns.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_info A pointer to a GDExtensionPropertyInfo struct.
 * @param p_setter A pointer to a StringName with the name of the setter method.
 * @param p_getter A pointer to a StringName with the name of the getter method.
 * @param p_index The index to pass as the first argument to the getter and setter methods.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClassPropertyIndexed)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, const GDExtensionPropertyInfo *p_info, GDExtensionConstStringNamePtr p_setter, GDExtensionConstStringNamePtr p_getter, GDExtensionInt p_index);

/**
 * @name classdb_register_extension_class_property_group
 * @since 4.1
 *
 * Registers a property group on an extension class in the ClassDB.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_group_name A pointer to a String with the group name.
 * @param p_prefix A pointer to a String with the prefix used by properties in this group.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClassPropertyGroup)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, GDExtensionConstStringPtr p_group_name, GDExtensionConstStringPtr p_prefix);

/**
 * @name classdb_register_extension_class_property_subgroup
 * @since 4.1
 *
 * Registers a property subgroup on an extension class in the ClassDB.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_subgroup_name A pointer to a String with the subgroup name.
 * @param p_prefix A pointer to a String with the prefix used by properties in this subgroup.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClassPropertySubgroup)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, GDExtensionConstStringPtr p_subgroup_name, GDExtensionConstStringPtr p_prefix);

/**
 * @name classdb_register_extension_class_signal
 * @since 4.1
 *
 * Registers a signal on an extension class in the ClassDB.
 *
 * Provided structs can be safely freed once the function returns.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 * @param p_signal_name A pointer to a StringName with the signal name.
 * @param p_argument_info A pointer to a GDExtensionPropertyInfo struct.
 * @param p_argument_count The number of arguments the signal receives.
 */
typedef void (*GDExtensionInterfaceClassdbRegisterExtensionClassSignal)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name, GDExtensionConstStringNamePtr p_signal_name, const GDExtensionPropertyInfo *p_argument_info, GDExtensionInt p_argument_count);

/**
 * @name classdb_unregister_extension_class
 * @since 4.1
 *
 * Unregisters an extension class in the ClassDB.
 *
 * Unregistering a parent class before a class that inherits it will result in failure. Inheritors must be unregistered first.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_class_name A pointer to a StringName with the class name.
 */
typedef void (*GDExtensionInterfaceClassdbUnregisterExtensionClass)(GDExtensionClassLibraryPtr p_library, GDExtensionConstStringNamePtr p_class_name);

/**
 * @name get_library_path
 * @since 4.1
 *
 * Gets the path to the current GDExtension library.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param r_path A pointer to a String which will receive the path.
 */
typedef void (*GDExtensionInterfaceGetLibraryPath)(GDExtensionClassLibraryPtr p_library, GDExtensionUninitializedStringPtr r_path);

/**
 * @name editor_add_plugin
 * @since 4.1
 *
 * Adds an editor plugin.
 *
 * It's safe to call during initialization.
 *
 * @param p_class_name A pointer to a StringName with the name of a class (descending from EditorPlugin) which is already registered with ClassDB.
 */
typedef void (*GDExtensionInterfaceEditorAddPlugin)(GDExtensionConstStringNamePtr p_class_name);

/**
 * @name editor_remove_plugin
 * @since 4.1
 *
 * Removes an editor plugin.
 *
 * @param p_class_name A pointer to a StringName with the name of a class that was previously added as an editor plugin.
 */
typedef void (*GDExtensionInterfaceEditorRemovePlugin)(GDExtensionConstStringNamePtr p_class_name);

/**
 * @name editor_help_load_xml_from_utf8_chars
 * @since 4.3
 *
 * Loads new XML-formatted documentation data in the editor.
 *
 * The provided pointer can be immediately freed once the function returns.
 *
 * @param p_data A pointer to a UTF-8 encoded C string (null terminated).
 */
typedef void (*GDExtensionsInterfaceEditorHelpLoadXmlFromUtf8Chars)(const char *p_data);

/**
 * @name editor_help_load_xml_from_utf8_chars_and_len
 * @since 4.3
 *
 * Loads new XML-formatted documentation data in the editor.
 *
 * The provided pointer can be immediately freed once the function returns.
 *
 * @param p_data A pointer to a UTF-8 encoded C string.
 * @param p_size The number of bytes (not code units).
 */
typedef void (*GDExtensionsInterfaceEditorHelpLoadXmlFromUtf8CharsAndLen)(const char *p_data, GDExtensionInt p_size);

/**
 * @name editor_register_get_classes_used_callback
 * @since 4.5
 *
 * Registers a callback that Godot can call to get the list of all classes (from ClassDB) that may be used by the calling GDExtension.
 *
 * This is used by the editor to generate a build profile (in "Tools" > "Engine Compilation Configuration Editor..." > "Detect from project"),
 * in order to recompile Godot with only the classes used.
 * In the provided callback, the GDExtension should provide the list of classes that _may_ be used statically, thus the time of invocation shouldn't matter.
 * If a GDExtension doesn't register a callback, Godot will assume that it could be using any classes.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_callback The callback to retrieve the list of classes used.
 */
typedef void (*GDExtensionInterfaceEditorRegisterGetClassesUsedCallback)(GDExtensionClassLibraryPtr p_library, GDExtensionEditorGetClassesUsedCallback p_callback);

/**
 * @name register_main_loop_callbacks
 * @since 4.5
 *
 * Registers callbacks to be called at different phases of the main loop.
 *
 * @param p_library A pointer the library received by the GDExtension's entry point function.
 * @param p_callbacks A pointer to the structure that contains the callbacks.
 */
typedef void (*GDExtensionInterfaceRegisterMainLoopCallbacks)(GDExtensionClassLibraryPtr p_library, const GDExtensionMainLoopCallbacks *p_callbacks);

#ifdef __cplusplus
}
#endif
