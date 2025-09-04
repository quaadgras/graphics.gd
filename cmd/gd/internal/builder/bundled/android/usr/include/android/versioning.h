/*
 * Copyright (C) 2016 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#pragma once

/**
 * @def __ANDROID_UNAVAILABLE_SYMBOLS_ARE_WEAK__
 *
 * Controls whether calling APIs newer than the developer's minSdkVersion are a
 * build error (when not defined) or allowed as a weak reference with a
 * __builtin_available() guard (when defined).
 *
 * See https://developer.android.com/ndk/guides/using-newer-apis for more
 * details.
 */
#if defined(__ANDROID_UNAVAILABLE_SYMBOLS_ARE_WEAK__)
// In this mode, Clang will emit weak references to the APIs if the
// minSdkVersion is less than the __what argument. This allows the libraries to
// load even on systems too old to contain the API, but calls must be guarded
// with `__builtin_available(android api_level, *)` to avoid segfaults.
#define __BIONIC_AVAILABILITY(__what, ...) __attribute__((__availability__(android,__what __VA_OPT__(,) __VA_ARGS__)))

// When the caller is using weak API references, we should expose the decls for
// APIs which are not available in the caller's minSdkVersion, otherwise there's
// no way to take advantage of the weak references.
#define __BIONIC_AVAILABILITY_GUARD(api_level) 1
#else
// The 'strict' flag is required for NDK clients where the code was not written
// to handle the case where the API was available at build-time but not at
// run-time. Most 3p code ported to Android was not written to use
// `__builtin_available()` for run-time availability checking, and so would not
// compile in this mode (or worse, if the build doesn't use
// -Werror=unguarded-availability, it would build but crash at runtime).
#define __BIONIC_AVAILABILITY(__what, ...) __attribute__((__availability__(android,strict,__what __VA_OPT__(,) __VA_ARGS__)))

// When the caller is using strict API references, we hide APIs which are not
// available in the caller's minSdkVersion. This is a bionic-only deviation in
// behavior from the rest of the NDK headers, but it's necessary to maintain
// source compatibility with 3p libraries that either can't correctly detect API
// availability (either incorrectly detecting as always-available or as
// never-available, but neither is true), or define their own polyfills which
// conflict with our declarations.
//
// https://github.com/android/ndk/issues/2081
#define __BIONIC_AVAILABILITY_GUARD(api_level) (__ANDROID_MIN_SDK_VERSION__ >= (api_level))
#endif

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wc23-extensions"
// Passing no argument for the '...' parameter of a variadic macro is a C23 extension
#define __INTRODUCED_IN(api_level) __BIONIC_AVAILABILITY(introduced=api_level)
#pragma clang diagnostic pop

#define __DEPRECATED_IN(api_level, msg) __BIONIC_AVAILABILITY(deprecated=api_level, message=msg)
#define __REMOVED_IN(api_level, msg) __BIONIC_AVAILABILITY(obsoleted=api_level, message=msg)

// The same availability attribute can't be annotated multiple times. Therefore, the macros are
// defined for the configuration that it is valid for so that declarations like the below doesn't
// cause inconsistent availability values which is an error with -Wavailability:
//
// void foo() __INTRODUCED_IN_32(30) __INTRODUCED_IN_64(31);
//
#if !defined(__LP64__)
#define __INTRODUCED_IN_32(api_level) __BIONIC_AVAILABILITY(introduced=api_level)
#define __INTRODUCED_IN_64(api_level)
#else
#define __INTRODUCED_IN_32(api_level)
#define __INTRODUCED_IN_64(api_level) __BIONIC_AVAILABILITY(introduced=api_level)
#endif
