/*
 * Copyright (C) 2024 The Android Open Source Project
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
 * @file android/crash_detail.h
 * @brief Attach extra information to android crashes.
 */

#include <sys/cdefs.h>

#include <stddef.h>

__BEGIN_DECLS

typedef struct crash_detail_t crash_detail_t;

/**
 * Register a new buffer to get logged into tombstones for crashes.
 *
 * It will be added to both the tombstone proto in the crash_detail field, and
 * in the tombstone text format.
 *
 * Tombstone proto definition:
 *   https://cs.android.com/android/platform/superproject/main/+/main:system/core/debuggerd/proto/tombstone.proto
 *
 * An app can get hold of these for any `REASON_CRASH_NATIVE` instance of
 * `android.app.ApplicationExitInfo`.
 *
 * https://developer.android.com/reference/android/app/ApplicationExitInfo#getTraceInputStream()

 * The lifetime of name and data has to be valid until the program crashes, or until
 * android_crash_detail_unregister is called.
 *
 * Example usage:
 *   const char* stageName = "garbage_collection";
 *   crash_detail_t* cd = android_crash_detail_register("stage", stageName, strlen(stageName));
 *   do_garbage_collection();
 *   android_crash_detail_unregister(cd);
 *
 * If this example crashes in do_garbage_collection, a line will show up in the textual representation of the tombstone:
 *   Extra crash detail: stage: 'garbage_collection'
 *
 * Introduced in API 35.
 *
 * \param name identifying name for this extra data.
 *             this should generally be a human-readable UTF-8 string, but we are treating
 *             it as arbitrary bytes because it could be corrupted by the crash.
 * \param name_size number of bytes of the buffer pointed to by name
 * \param data a buffer containing the extra detail bytes, if null the crash detail
 *             is disabled until android_crash_detail_replace_data replaces it with
 *             a non-null pointer.
 * \param data_size number of bytes of the buffer pointed to by data
 *
 * \return a handle to the extra crash detail.
 */

#if __BIONIC_AVAILABILITY_GUARD(35)
crash_detail_t* _Nullable android_crash_detail_register(
    const void* _Nonnull name, size_t name_size, const void* _Nullable data, size_t data_size) __INTRODUCED_IN(35);

/**
 * Unregister crash detail from being logged into tombstones.
 *
 * After this function returns, the lifetime of the objects crash_detail was
 * constructed from no longer needs to be valid.
 *
 * Introduced in API 35.
 *
 * \param crash_detail the crash_detail that should be removed.
 */
void android_crash_detail_unregister(crash_detail_t* _Nonnull crash_detail) __INTRODUCED_IN(35);

/**
 * Replace data of crash detail.
 *
 * This is more efficient than using android_crash_detail_unregister followed by
 * android_crash_detail_register. If you very frequently need to swap out the data,
 * you can hold onto the crash_detail.
 *
 * Introduced in API 35.
 *
 * \param data the new buffer containing the extra detail bytes, or null to disable until
 *             android_crash_detail_replace_data is called again with non-null data.
 * \param data_size the number of bytes of the buffer pointed to by data.
 */
void android_crash_detail_replace_data(crash_detail_t* _Nonnull crash_detail, const void* _Nullable data, size_t data_size) __INTRODUCED_IN(35);

/**
 * Replace name of crash detail.
 *
 * This is more efficient than using android_crash_detail_unregister followed by
 * android_crash_detail_register. If you very frequently need to swap out the name,
 * you can hold onto the crash_detail.
 *
 * Introduced in API 35.
 *
 * \param name identifying name for this extra data.
 * \param name_size number of bytes of the buffer pointed to by name
 */
void android_crash_detail_replace_name(crash_detail_t* _Nonnull crash_detail, const void* _Nonnull name, size_t name_size) __INTRODUCED_IN(35);
#endif /* __BIONIC_AVAILABILITY_GUARD(35) */


__END_DECLS
