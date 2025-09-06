/*
 * Copyright (C) 2008 The Android Open Source Project
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *  * Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in
 *    the documentation and/or other materials provided with the
 *    distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
 * OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED
 * AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT
 * OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */

#pragma once

#include <sys/cdefs.h>

#include <stdint.h>

/**
 * @addtogroup libdl Dynamic Linker
 * @{
 */

/**
 * \file
 * Standard dynamic library support.
 * See also the Android-specific functionality in `<android/dlext.h>`.
 */

__BEGIN_DECLS

/**
 * dladdr() returns information using this structure.
 */
typedef struct {
  /** Pathname of the shared object that contains the given address. */
  const char* _Nullable dli_fname;
  /** Address at which the shared object is loaded. */
  void* _Nullable dli_fbase;
  /** Name of the nearest symbol with an address lower than the given address. */
  const char* _Nullable dli_sname;
  /** Exact address of the symbol named in `dli_sname`. */
  void* _Nullable dli_saddr;
} Dl_info;

/**
 * [dlopen(3)](https://man7.org/linux/man-pages/man3/dlopen.3.html)
 * loads the given shared library.
 *
 * See also
 * [Android changes for NDK developers](https://android.googlesource.com/platform/bionic/+/main/android-changes-for-ndk-developers.md)
 * which should cover all dynamic linker behavioral changes relevant to app development.
 * It also explains how to debug issues with shared libraries on Android
 * using the `debug.ld.app.*` system properties.
 *
 * One Android-specific extension is particularly noteworthy.
 * The "my_zip_file.zip!/libs/libstuff.so" syntax to load a library
 * directly from an arbitrary zip file, including but not limited to your apk
 * (where this is how `extractNativeLibs=false` is implemented).
 * Related to that, until API level 36 PackageManager is fussy about what
 * files `extractNativeLibs` will actually extract.
 * To be compatible with all API levels,
 * always give files that need to be extracted a "lib" prefix and ".so" suffix,
 * or avoid using `extractNativeLibs`.
 *
 * See also android_dlopen_ext().
 *
 * Returns a pointer to an opaque handle for use with other <dlfcn.h> functions
 * on success, and returns NULL on failure, in which case dlerror() can be used
 * to retrieve the specific error.
 */
void* _Nullable dlopen(const char* _Nullable __filename, int __flag);

/**
 * [dlclose(3)](https://man7.org/linux/man-pages/man3/dlclose.3.html)
 * decrements the reference count for the given shared library (and
 * any libraries brought in by that library's DT_NEEDED entries).
 *
 * If a library's reference count hits zero, it may be unloaded.
 * Code that relies on this is not portable, and may not work on
 * future versions of Android.
 *
 * dlclose() is dangerous because function pointers may or may not
 * be rendered invalid, global data may or may not be rendered invalid,
 * and memory may or may not leak. Code with global constructors is
 * especially problematic. Instead of dlclose, prefer to leave the
 * library open or, if cleanup is necessary, dlopen() the library in
 * a child process which can later be killed by the parent or call
 * exit() itself.
 *
 * Note also that dlclose() interacts badly with thread local variables
 * with non-trivial destructors, with the
 * (exact behavior varying by API level)[https://android.googlesource.com/platform/bionic/+/main/android-changes-for-ndk-developers.md#dlclose-interacts-badly-with-thread-local-variables-with-non_trivial-destructors].
 *
 * Returns 0 on success, and returns -1 on failure, in which case
 * dlerror() can be used to retrieve the specific error.
 */
int dlclose(void* _Nonnull __handle);

/**
 * [dlerror(3)](https://man7.org/linux/man-pages/man3/dlerror.3.html)
 * returns a human-readable error message describing the most recent
 * failure from one of the <dlfcn.h> functions on the calling thread.
 *
 * This function also clears the error, so a second call (or a call
 * before any failure) will return NULL.
 *
 * Returns a pointer to an error on success, and returns NULL if no
 * error is pending.
 */
char* _Nullable dlerror(void);

/**
 * [dlsym(3)](https://man7.org/linux/man-pages/man3/dlsym.3.html)
 * returns a pointer to the symbol with the given name in the shared
 * library represented by the given handle. The handle may have been
 * returned from dlopen(), or can be RTLD_DEFAULT or RTLD_NEXT.
 *
 * Returns the address of the symbol on success, and returns NULL on failure,
 * in which case dlerror() can be used to retrieve the specific error.
 */
void* _Nullable dlsym(void* __BIONIC_COMPLICATED_NULLNESS __handle, const char* _Nullable __symbol);

/**
 * [dlvsym(3)](https://man7.org/linux/man-pages/man3/dlvsym.3.html)
 * returns a pointer to the symbol with the given name and version in the shared
 * library represented by the given handle. The handle may have been
 * returned from dlopen(), or can be RTLD_DEFAULT or RTLD_NEXT.
 *
 * Returns the address of the symbol on success, and returns NULL on failure,
 * in which case dlerror() can be used to retrieve the specific error.
 */

#if __BIONIC_AVAILABILITY_GUARD(24)
void* _Nullable dlvsym(void* __BIONIC_COMPLICATED_NULLNESS __handle, const char* _Nullable __symbol, const char* _Nullable __version) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */


/**
 * [dladdr(3)](https://man7.org/linux/man-pages/man3/dladdr.3.html)
 * returns information about the symbol at the given address.
 *
 * Returns non-zero on success, and returns 0 on failure. Note that unlike
 * the other <dlfcn.h> functions, in this case dlerror() will _not_ have
 * more information.
 */
int dladdr(const void* _Nonnull __addr, Dl_info* _Nonnull __info);

/**
 * A dlsym()/dlvsym() handle that returns the first symbol found in any
 * shared library using the default search order.
 */
#define RTLD_DEFAULT  __BIONIC_CAST(reinterpret_cast, void*, 0)

/**
 * A dlsym()/dlvsym() handle that returns the first symbol found in any
 * shared library that appears _after_ the object containing the caller.
 */
#define RTLD_NEXT     __BIONIC_CAST(reinterpret_cast, void*, -1L)

/**
 * A dlopen() flag to not make symbols from this library available to later
 * libraries. See also RTLD_GLOBAL.
 */
#define RTLD_LOCAL    0

/**
 * Not supported on Android. Android always uses RTLD_NOW for security reasons.
 * Resolving all undefined symbols before dlopen() returns means that RELRO
 * protections can be applied to the PLT before dlopen() returns.
 */
#define RTLD_LAZY     0x00001

/** A dlopen() flag to resolve all undefined symbols before dlopen() returns. */
#define RTLD_NOW      0x00002

/**
 * A dlopen() flag to not actually load the given library;
 * used to test whether the library is already loaded.
 */
#define RTLD_NOLOAD   0x00004

/**
 * A dlopen() flag to make symbols from this library available to later
 * libraries. See also RTLD_LOCAL.
 */
#define RTLD_GLOBAL   0x00100

/**
 * A dlopen() flag to ignore later dlclose() calls on this library.
 */
#define RTLD_NODELETE 0x01000

/* LP32 has historical ABI breakage. */
#if !defined(__LP64__)
#undef RTLD_DEFAULT
#define RTLD_DEFAULT  __BIONIC_CAST(reinterpret_cast, void*, 0xffffffff)
#undef RTLD_NEXT
#define RTLD_NEXT     __BIONIC_CAST(reinterpret_cast, void*, 0xfffffffe)
#undef RTLD_NOW
#define RTLD_NOW      0x00000
#undef RTLD_GLOBAL
#define RTLD_GLOBAL   0x00002
#endif

__END_DECLS

/** @} */
