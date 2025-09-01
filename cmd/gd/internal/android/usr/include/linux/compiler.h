#pragma once

/*
 * There is no `include/uapi/linux/compiler.h`, just `include/linux/compiler.h`.
 *
 * We don't need anything _in_ this file, but we do need this file.
 * The two #defines are for backwards compatibility.
 */

#define __force
#define __user
