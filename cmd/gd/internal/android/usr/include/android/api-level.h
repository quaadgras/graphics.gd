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

/**
 * @defgroup apilevels API Levels
 *
 * Defines functions and constants for working with Android API levels.
 * @{
 */

/**
 * @file android/api-level.h
 * @brief Functions and constants for dealing with multiple API levels.
 *
 * See
 * https://android.googlesource.com/platform/bionic/+/main/docs/defines.md.
 */

#include <sys/cdefs.h>

__BEGIN_DECLS

/**
 * Magic version number for an Android OS build which has not yet turned
 * into an official release, for comparison against `__ANDROID_API__`. See
 * https://android.googlesource.com/platform/bionic/+/main/docs/defines.md.
 */
#define __ANDROID_API_FUTURE__ 10000

/* This #ifndef should never be true except when doxygen is generating docs. */
#ifndef __ANDROID_API__
/**
 * `__ANDROID_API__` is the [API
 * level](https://developer.android.com/guide/topics/manifest/uses-sdk-element#ApiLevels)
 * this code is being built for. The resulting binaries are only guaranteed to
 * be compatible with devices which have an API level greater than or equal to
 * `__ANDROID_API__`.
 *
 * For NDK and APEX builds, this macro will always be defined. It is set
 * automatically by Clang using the version suffix that is a part of the target
 * name. For example, `__ANDROID_API__` will be 24 when Clang is given the
 * argument `-target aarch64-linux-android24`.
 *
 * For non-APEX OS code, this defaults to  __ANDROID_API_FUTURE__.
 *
 * The value of `__ANDROID_API__` can be compared to the named constants in
 * `<android/api-level.h>`.
 *
 * The interpretation of `__ANDROID_API__` is similar to the AndroidManifest.xml
 * `minSdkVersion`. In most cases `__ANDROID_API__` will be identical to
 * `minSdkVersion`, but as it is a build time constant it is possible for
 * library code to use a different value than the app it will be included in.
 * When libraries and applications build for different API levels, the
 * `minSdkVersion` of the application must be at least as high as the highest
 * API level used by any of its libraries which are loaded unconditionally.
 *
 * Note that in some cases the resulting binaries may load successfully on
 * devices with an older API level. That behavior should not be relied upon,
 * even if you are careful to avoid using new APIs, as the toolchain may make
 * use of new features by default. For example, additional FORTIFY features may
 * implicitly make use of new APIs, SysV hashes may be omitted in favor of GNU
 * hashes to improve library load times, or relocation packing may be enabled to
 * reduce binary size.
 *
 * See android_get_device_api_level(),
 * android_get_application_target_sdk_version() and
 * https://android.googlesource.com/platform/bionic/+/main/docs/defines.md.
 */
#define __ANDROID_API__ __ANDROID_API_FUTURE__
#endif

/** Names the Gingerbread API level (9), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_G__ 9

/** Names the Ice-Cream Sandwich API level (14), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_I__ 14

/** Names the Jellybean API level (16), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_J__ 16

/** Names the Jellybean MR1 API level (17), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_J_MR1__ 17

/** Names the Jellybean MR2 API level (18), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_J_MR2__ 18

/** Names the KitKat API level (19), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_K__ 19

/** Names the Lollipop API level (21), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_L__ 21

/** Names the Lollipop MR1 API level (22), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_L_MR1__ 22

/** Names the Marshmallow API level (23), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_M__ 23

/** Names the Nougat API level (24), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_N__ 24

/** Names the Nougat MR1 API level (25), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_N_MR1__ 25

/** Names the Oreo API level (26), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_O__ 26

/** Names the Oreo MR1 API level (27), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_O_MR1__ 27

/** Names the Pie API level (28), for comparison against `__ANDROID_API__`. */
#define __ANDROID_API_P__ 28

/**
 * Names the Android 10 (aka "Q" or "Quince Tart") API level (29), for
 * comparison against `__ANDROID_API__`.
 */
#define __ANDROID_API_Q__ 29

/**
 * Names the Android 11 (aka "R" or "Red Velvet Cake") API level (30), for
 * comparison against `__ANDROID_API__`.
 */
#define __ANDROID_API_R__ 30

/**
 * Names the Android 12 (aka "S" or "Snowcone") API level (31), for
 * comparison against `__ANDROID_API__`.
 */
#define __ANDROID_API_S__ 31

/**
 * Names the Android 13 (aka "T" or "Tiramisu") API level (33), for
 * comparison against `__ANDROID_API__`.
 */
#define __ANDROID_API_T__ 33

/**
 * Names the Android 14 (aka "U" or "UpsideDownCake") API level (34),
 * for comparison against `__ANDROID_API__`.
 */
#define __ANDROID_API_U__ 34

/**
 * Names the Android 15 (aka "V" or "VanillaIceCream") API level (35),
 * for comparison against `__ANDROID_API__`.
 */
#define __ANDROID_API_V__ 35

/* This file is included in <features.h>, and might be used from .S files. */
#if !defined(__ASSEMBLY__)

/**
 * Returns the `targetSdkVersion` of the caller, or `__ANDROID_API_FUTURE__` if
 * there is no known target SDK version (for code not running in the context of
 * an app).
 *
 * The returned values correspond to the named constants in `<android/api-level.h>`,
 * and is equivalent to the AndroidManifest.xml `targetSdkVersion`.
 *
 * See also android_get_device_api_level().
 *
 * Available since API level 24.
 */

#if __BIONIC_AVAILABILITY_GUARD(24)
int android_get_application_target_sdk_version() __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */


#if __ANDROID_API__ < 29

/* android_get_device_api_level is a static inline before API level 29. */
#define __BIONIC_GET_DEVICE_API_LEVEL_INLINE static __inline
#include <bits/get_device_api_level_inlines.h>
#undef __BIONIC_GET_DEVICE_API_LEVEL_INLINE

#else

/**
 * Returns the API level of the device we're actually running on, or -1 on failure.
 * The returned values correspond to the named constants in `<android/api-level.h>`,
 * and is equivalent to the Java `Build.VERSION.SDK_INT` API.
 *
 * See also android_get_application_target_sdk_version().
 */
int android_get_device_api_level() __INTRODUCED_IN(29);

#endif

#endif /* defined(__ASSEMBLY__) */

__END_DECLS

/** @} */
