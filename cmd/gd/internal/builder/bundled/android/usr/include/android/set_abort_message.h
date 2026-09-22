/*
 * Copyright (C) 2014 The Android Open Source Project
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
 * @file android/set_abort_message.h
 * @brief The android_set_abort_message() function.
 */

#include <sys/cdefs.h>

#include <stddef.h>
#include <stdint.h>
#include <string.h>

__BEGIN_DECLS

typedef struct crash_detail_t crash_detail_t;

/**
 * android_set_abort_message() sets the abort message passed to
 * [debuggerd](https://source.android.com/devices/tech/debug/native-crash)
 * for inclusion in any crash.
 *
 * This is meant for use by libraries that deliberately abort so that they can
 * provide an explanation. It is used within bionic to implement assert() and
 * all FORTIFY and fdsan failures.
 *
 * The message appears directly in logcat at the time of crash. It will
 * also be added to both the tombstone proto in the crash_detail field, and
 * in the tombstone text format.
 *
 * Tombstone proto definition:
 *   https://cs.android.com/android/platform/superproject/main/+/main:system/core/debuggerd/proto/tombstone.proto
 *
 * An app can get hold of these for any `REASON_CRASH_NATIVE` instance of
 * `android.app.ApplicationExitInfo`.
 *  https://developer.android.com/reference/android/app/ApplicationExitInfo#getTraceInputStream()
 *
 * The given message is copied at the time this function is called, and does
 * not need to be valid until the crash actually happens, but typically this
 * function is called immediately before aborting. See <android/crash_detail.h>
 * for API more suited to the use case where the caller doesn't _expect_ a
 * crash but would like to see the information _if_ a crash happens.
 */
void android_set_abort_message(const char* _Nullable __msg);

__END_DECLS
