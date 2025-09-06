#pragma once

/*
 * There is no `include/uapi/linux/compiler_types.h`, just
 * `include/linux/compiler_types.h`.
 *
 * We don't need anything _in_ this file, but we do need this file.
 * The #include is for backwards compatibility.
 */

#include <linux/compiler.h>
