/*
 * Copyright (C) 2024 The Android Open Source Project
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
 * @file surface_control_input_receiver.h
 */

#pragma once

#include <stdint.h>
#include <android/input.h>
#include <android/surface_control.h>
#include <android/input_transfer_token_jni.h>

__BEGIN_DECLS

/**
 * The AInputReceiver_onMotionEvent callback is invoked when the registered input channel receives
 * a motion event.
 *
 * \param context Optional context provided by the client that is passed when creating the
 * AInputReceiverCallbacks.
 *
 * \param motionEvent The motion event. This must be released with AInputEvent_release.
 *
 * Available since API level 35.
 */
typedef bool (*AInputReceiver_onMotionEvent)(void *_Null_unspecified context,
                                             AInputEvent *_Nonnull motionEvent)
                                            __INTRODUCED_IN(__ANDROID_API_V__);
/**
 * The AInputReceiver_onKeyEvent callback is invoked when the registered input channel receives
 * a key event.
 *
 * \param context Optional context provided by the client that is passed when creating the
 * AInputReceiverCallbacks.
 *
 * \param keyEvent The key event. This must be released with AInputEvent_release.
 *
 * Available since API level 35.
 */
typedef bool (*AInputReceiver_onKeyEvent)(void *_Null_unspecified context,
                                          AInputEvent *_Nonnull keyEvent)
                                          __INTRODUCED_IN(__ANDROID_API_V__);

typedef struct AInputReceiverCallbacks AInputReceiverCallbacks;

/**
 * The InputReceiver that holds the reference to the registered input channel. This must be released
 * using AInputReceiver_release
 */
typedef struct AInputReceiver AInputReceiver;

/**
 * Registers an input receiver for an ASurfaceControl that will receive batched input event. For
 * those events that are batched, the invocation will happen once per AChoreographer frame, and
 * other input events will be delivered immediately.
 *
 * This is different from AInputReceiver_createUnbatchedInputReceiver in that the input events are
 * received batched. The caller must invoke AInputReceiver_release to clean up the resources when
 * no longer needing to use the input receiver.
 *
 * \param aChoreographer         The AChoreographer used for batching. This should match the
 *                               rendering AChoreographer.
 * \param hostInputTransferToken The host token to link the embedded. This is used to handle
 *                               transferring touch gesture from host to embedded and for ANRs
 *                               to ensure the host receives the ANR if any issues with
 *                               touch on the embedded. This can be retrieved for the host window
 *                               by calling AttachedSurfaceControl#getInputTransferToken()
 * \param aSurfaceControl        The ASurfaceControl to register the InputChannel for
 * \param aInputReceiverCallbacks The SurfaceControlInputReceiver that will receive the input events
 *
 * Returns the reference to AInputReceiver to clean up resources when done.
 *
 * Available since API level 35.
 */
AInputReceiver* _Nonnull
AInputReceiver_createBatchedInputReceiver(AChoreographer* _Nonnull aChoreographer,
                                        const AInputTransferToken* _Nonnull hostInputTransferToken,
                                        const ASurfaceControl* _Nonnull aSurfaceControl,
                                        AInputReceiverCallbacks* _Nonnull aInputReceiverCallbacks)
                                        __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Registers an input receiver for an ASurfaceControl that will receive every input event.
 * This is different from AInputReceiver_createBatchedInputReceiver in that the input events are
 * received unbatched. The caller must invoke AInputReceiver_release to clean up the resources when
 * no longer needing to use the input receiver.
 *
 * \param aLooper                The looper to use when invoking callbacks.
 * \param hostInputTransferToken The host token to link the embedded. This is used to handle
 *                               transferring touch gesture from host to embedded and for ANRs
 *                               to ensure the host receives the ANR if any issues with
 *                               touch on the embedded. This can be retrieved for the host window
 *                               by calling AttachedSurfaceControl#getInputTransferToken()
 * \param aSurfaceControl        The ASurfaceControl to register the InputChannel for
 * \param aInputReceiverCallbacks The SurfaceControlInputReceiver that will receive the input events
 *
 * Returns the reference to AInputReceiver to clean up resources when done.
 *
 * Available since API level 35.
 */
AInputReceiver* _Nonnull
AInputReceiver_createUnbatchedInputReceiver(ALooper* _Nonnull aLooper,
                                         const AInputTransferToken* _Nonnull hostInputTransferToken,
                                         const ASurfaceControl* _Nonnull aSurfaceControl,
                                         AInputReceiverCallbacks* _Nonnull aInputReceiverCallbacks)
                                         __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Returns the AInputTransferToken that can be used to transfer touch gesture to or from other
 * windows. This InputTransferToken is associated with the SurfaceControl that registered an input
 * receiver and can be used with the host token for things like transfer touch gesture via
 * WindowManager#transferTouchGesture().
 *
 * This must be released with AInputTransferToken_release.
 *
 * \param aInputReceiver The inputReceiver object to retrieve the AInputTransferToken for.
 *
 * Available since API level 35.
 */
const AInputTransferToken *_Nonnull
AInputReceiver_getInputTransferToken(AInputReceiver *_Nonnull aInputReceiver)
        __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Unregisters the input channel and deletes the AInputReceiver. This must be called on the same
 * looper thread it was created with.
 *
 * \param aInputReceiver The inputReceiver object to release.
 *
 * Available since API level 35.
 */
void
AInputReceiver_release(AInputReceiver *_Nullable aInputReceiver) __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Creates a AInputReceiverCallbacks object that is used when registering for an AInputReceiver.
 * This must be released using AInputReceiverCallbacks_release
 *
 * \param context Optional context provided by the client that will be passed into the callbacks.
 *
 * Available since API level 35.
 */
AInputReceiverCallbacks* _Nonnull AInputReceiverCallbacks_create(void* _Nullable context)
                                                        __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Releases the AInputReceiverCallbacks. This must be called on the same
 * looper thread the AInputReceiver was created with. The receiver will not invoke any callbacks
 * once it's been released.
 *
 * Available since API level 35
 */
void AInputReceiverCallbacks_release(AInputReceiverCallbacks* _Nullable callbacks)
                                     __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Sets a AInputReceiver_onMotionEvent callback for an AInputReceiverCallbacks
 *
 * \param callbacks The callback object to set the motion event on.
 * \param onMotionEvent The motion event that will be invoked
 *
 * Available since API level 35.
 */
void AInputReceiverCallbacks_setMotionEventCallback(AInputReceiverCallbacks* _Nonnull callbacks,
                                                AInputReceiver_onMotionEvent _Nonnull onMotionEvent)
                                                __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Sets a AInputReceiver_onKeyEvent callback for an AInputReceiverCallbacks
 *
 * \param callbacks The callback object to set the motion event on.
 * \param onMotionEvent The key event that will be invoked
 *
 * Available since API level 35.
 */
void AInputReceiverCallbacks_setKeyEventCallback(AInputReceiverCallbacks* _Nonnull callbacks,
                                                 AInputReceiver_onKeyEvent _Nonnull onKeyEvent)
                                                 __INTRODUCED_IN(__ANDROID_API_V__);

__END_DECLS
