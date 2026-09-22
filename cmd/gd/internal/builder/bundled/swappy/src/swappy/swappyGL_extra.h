/*
 * Copyright 2018 The Android Open Source Project
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
 * @defgroup swappyGL_extra Swappy for OpenGL extras
 * Extra utility functions to use Swappy with OpenGL.
 * @{
 */

#pragma once

#include <EGL/egl.h>
#include <EGL/eglext.h>
#include <jni.h>
#include <stdint.h>

#include "swappy_common.h"

#ifdef __cplusplus
extern "C" {
#endif

/**
 * @brief If an app wishes to use the Android choreographer to provide ticks to
 * Swappy, it can call this function.
 *
 * @warning This function *must* be called before the first `Swappy_swap()`
 * call. Afterwards, call this function every choreographer tick.
 */
void SwappyGL_onChoreographer(int64_t frameTimeNanos);

/** @brief Pass callbacks to be called each frame to trace execution. */
void SwappyGL_injectTracer(const SwappyTracer *t);

/**
 * @brief Toggle auto-swap interval detection on/off
 *
 * By default, Swappy will adjust the swap interval based on actual frame
 * rendering time. If an app wants to override the swap interval calculated by
 * Swappy, it can call `SwappyGL_setSwapIntervalNS`. This will temporarily
 * override Swappy's frame timings but, unless
 * `SwappyGL_setAutoSwapInterval(false)` is called, the timings will continue to
 * be be updated dynamically, so the swap interval may change.
 */
void SwappyGL_setAutoSwapInterval(bool enabled);

/**
 * @brief Sets the maximal duration for auto-swap interval in milliseconds.
 *
 * If Swappy is operating in auto-swap interval and the frame duration is longer
 * than `max_swap_ns`, Swappy will not do any pacing and just submit the frame
 * as soon as possible.
 */
void SwappyGL_setMaxAutoSwapIntervalNS(uint64_t max_swap_ns);

/**
 * @brief Toggle auto-pipeline mode on/off
 *
 * By default, if auto-swap interval is on, auto-pipelining is on and Swappy
 * will try to reduce latency by scheduling cpu and gpu work in the same
 * pipeline stage, if it fits.
 */
void SwappyGL_setAutoPipelineMode(bool enabled);

/**
 * @brief Toggle statistics collection on/off
 *
 * By default, stats collection is off and there is no overhead related to
 * stats. An app can turn on stats collection by calling
 * `SwappyGL_enableStats(true)`. Then, the app is expected to call
 * ::SwappyGL_recordFrameStart for each frame before starting to do any CPU
 * related work. Stats will be logged to logcat with a 'FrameStatistics' tag. An
 * app can get the stats by calling ::SwappyGL_getStats.
 */
void SwappyGL_enableStats(bool enabled);

/**
 * @brief Should be called if stats have been enabled with SwappyGL_enableStats.
 *
 * When stats collection is enabled with SwappyGL_enableStats, the app is
 * expected to call this function for each frame before starting to do any CPU
 * related work.
 *
 * @see SwappyGL_enableStats.
 */
void SwappyGL_recordFrameStart(EGLDisplay display, EGLSurface surface);

/**
 * @brief Returns the stats collected, if statistics collection was toggled on.
 *
 * @param swappyStats Pointer to a SwappyStats that will be populated with
 * collected stats.
 * @see SwappyStats
 * @see SwappyGL_enableStats
 */
void SwappyGL_getStats(SwappyStats *swappyStats);

/**
 * @brief Clears the frame statistics collected so far.
 *
 * All the frame statistics collected are reset to 0, frame statistics are
 * collected normally after this call.
 */
void SwappyGL_clearStats(void);

/** @brief Remove callbacks that were previously added using
 * SwappyGL_injectTracer. */
void SwappyGL_uninjectTracer(const SwappyTracer *t);

/**
 * @brief Reset the swappy pacing mechanism
 *
 * In cases where the frame timing history is irrelevant (for example during
 * scene/level transitions or after loading screens), calling this would
 * remove all the history for frame pacing. Calling this entry point
 * would reset the frame rate to the initial state at the end of the current
 * frame. Then swappy would just pace as normal with fresh state from next
 * frame. There are no error conditions associated with this call.
 */
void SwappyGL_resetFramePacing(void);

/**
 * @brief Enable/Disable the swappy pacing mechanism
 *
 * By default frame pacing is enabled when swappy is used, however it can be
 * disabled at runtime by calling `SwappyGL_enableFramePacing(false)`. When the
 * frame pacing is disabled, ::SwappyGL_enableBlockingWait will control whether
 * ::SwappyGL_swap will return immediately, or will block until the previous
 * frame's GPU work is completed. All enabling/disabling effects take place from
 * next frame. Once disabled,
 * `SwappyGL_enableFramePacing(true)` will enable frame pacing again with a
 * fresh frame time state - equivalent of calling ::SwappyGL_resetFramePacing().
 *
 * @param enable  If true, enables frame pacing otherwise disables it
 */
void SwappyGL_enableFramePacing(bool enable);

/**
 * @brief Enable/Disable blocking wait when frame-pacing is disabled
 *
 * By default ::SwappyGL_swap will do a blocking wait until previous frame's GPU
 * work is completed. However when frame pacing is disabled, calling
 * `SwappyGL_enableBlockingWait(false)` will ensure that ::SwappyGL_swap returns
 * without waiting for previous frame's GPU work. This behaviour impacts the
 * GPU time returned in the ::SwappyPostWaitCallback callback. If the last
 * frame's GPU work is done by the time ::SwappyGL_swap for this frame is
 * called, then the previous frame's GPU time will be returned otherwise `-1`
 * will be delivered in the callback for the frame. This setting has no impact
 * when frame pacing is enabled.
 *
 * @param enable  If true, ::SwappyGL_swap will block until the previous frame's
 *                GPU work is complete.
 */
void SwappyGL_enableBlockingWait(bool enable);

#ifdef __cplusplus
};
#endif

/** @} */
