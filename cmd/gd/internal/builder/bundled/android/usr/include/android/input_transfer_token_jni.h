/*
 * Copyright 2024 The Android Open Source Project
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
/**
 * @addtogroup NativeActivity Native Activity
 * @{
 */
/**
 * @file input_transfer_token_jni.h
 */

#pragma once

#include <sys/cdefs.h>
#include <jni.h>

__BEGIN_DECLS
struct AInputTransferToken;

/**
 * AInputTransferToken can be used to request focus on or to transfer touch gesture to and from
 * an embedded SurfaceControl
 */
typedef struct AInputTransferToken AInputTransferToken;

/**
 * Return the AInputTransferToken wrapped by a Java InputTransferToken object. This must be released
 * using AInputTransferToken_release
 *
 * inputTransferTokenObj must be a non-null instance of android.window.InputTransferToken.
 *
 * Available since API level 35.
 */
AInputTransferToken* _Nonnull AInputTransferToken_fromJava(JNIEnv* _Nonnull env,
        jobject _Nonnull inputTransferTokenObj) __INTRODUCED_IN(__ANDROID_API_V__);
/**
 * Return the Java InputTransferToken object that wraps AInputTransferToken
 *
 * aInputTransferToken must be non null and the returned value is an object of instance
 * android.window.InputTransferToken.
 *
 * Available since API level 35.
 */
jobject _Nonnull AInputTransferToken_toJava(JNIEnv* _Nonnull env,
        const AInputTransferToken* _Nonnull aInputTransferToken) __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Removes a reference that was previously acquired in native.
 *
 * Available since API level 35.
 */
void AInputTransferToken_release(AInputTransferToken* _Nullable aInputTransferToken)
        __INTRODUCED_IN(__ANDROID_API_V__);

__END_DECLS
/** @} */
