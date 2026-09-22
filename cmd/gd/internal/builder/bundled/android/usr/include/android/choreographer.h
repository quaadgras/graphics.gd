/*
 * Copyright (C) 2015 The Android Open Source Project
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
 * @addtogroup Choreographer
 *
 * Choreographer coordinates the timing of frame rendering. This is the C
 * version of the android.view.Choreographer object in Java. If you do not use
 * Choreographer to pace your render loop, you may render too quickly for the
 * display, increasing latency between frame submission and presentation.
 *
 * Input events are guaranteed to be processed before the frame callback is
 * called, and will not be run concurrently. Input and sensor events should not
 * be handled in the Choregrapher callback.
 *
 * The frame callback is also the appropriate place to run any per-frame state
 * update logic. For example, in a game, the frame callback should be
 * responsible for updating things like physics, AI, game state, and rendering
 * the frame. Input and sensors should be handled separately via callbacks
 * registered with AInputQueue and ASensorManager.
 *
 * As of API level 33, apps can follow proper frame pacing and even choose a future frame to render.
 * The API is used as follows:
 * 1. The app posts an {@link AChoreographer_vsyncCallback} to Choreographer to run on the next
 * frame.
 * 2. The callback is called when it is the time to start the frame with an {@link
 * AChoreographerFrameCallbackData} payload: information about multiple possible frame
 * timelines.
 * 3. Apps can choose a frame timeline from the {@link
 * AChoreographerFrameCallbackData} payload, depending on the frame deadline they can meet when
 * rendering the frame and their desired presentation time, and subsequently
 * {@link ASurfaceTransaction_setFrameTimeline notify SurfaceFlinger}
 * of the choice. Alternatively, for apps that do not choose a frame timeline, their frame would be
 * presented at the earliest possible timeline.
 *   - The preferred frame timeline is the default frame
 * timeline that the platform scheduled for the app, based on device configuration.
 * 4. SurfaceFlinger attempts to follow the chosen frame timeline, by not applying transactions or
 * latching buffers before the desired presentation time.
 *
 * On older devices, AChoreographer_postFrameCallback64 or
 * AChoreographer_postFrameCallback can be used to lesser effect. They cannot be
 * used to precisely plan your render timeline, but will rate limit to avoid
 * overloading the display pipeline and increasing frame latency.
 *
 * @{
 */

/**
 * @file choreographer.h
 */

#ifndef ANDROID_CHOREOGRAPHER_H
#define ANDROID_CHOREOGRAPHER_H

#include <stddef.h>
#include <stdint.h>
#include <sys/cdefs.h>

// This file may also be built on glibc or on Windows/MacOS libc's, so no-op
// and deprecated definitions are provided.
#if !defined(__INTRODUCED_IN)
#define __INTRODUCED_IN(__api_level) /* nothing */
#endif
#if !defined(__DEPRECATED_IN)
#define __DEPRECATED_IN(__api_level, ...) __attribute__((__deprecated__))
#endif

__BEGIN_DECLS

struct AChoreographer;
/**
 * Opaque type that provides access to an AChoreographer object.
 *
 * A pointer can be obtained using {@link AChoreographer_getInstance()}.
 */
typedef struct AChoreographer AChoreographer;


/**
 * The identifier of a frame timeline.
 */
typedef int64_t AVsyncId;

struct AChoreographerFrameCallbackData;
/**
 * Opaque type that provides access to an AChoreographerFrameCallbackData object, which contains
 * various methods to extract frame information.
 */
typedef struct AChoreographerFrameCallbackData AChoreographerFrameCallbackData;

/**
 * Prototype of the function that is called when a new frame is being rendered.
 * It's passed the time that the frame is being rendered as nanoseconds in the
 * CLOCK_MONOTONIC time base, as well as the data pointer provided by the
 * application that registered a callback. All callbacks that run as part of
 * rendering a frame will observe the same frame time, so it should be used
 * whenever events need to be synchronized (e.g. animations).
 */
typedef void (*AChoreographer_frameCallback)(long frameTimeNanos, void* data);

/**
 * Prototype of the function that is called when a new frame is being rendered.
 * It's passed the time that the frame is being rendered as nanoseconds in the
 * CLOCK_MONOTONIC time base, as well as the data pointer provided by the
 * application that registered a callback. All callbacks that run as part of
 * rendering a frame will observe the same frame time, so it should be used
 * whenever events need to be synchronized (e.g. animations).
 */
typedef void (*AChoreographer_frameCallback64)(int64_t frameTimeNanos, void* data);

/**
 * Prototype of the function that is called when a new frame is being rendered.
 * It is called with \c callbackData describing multiple frame timelines, as well as the \c data
 * pointer provided by the application that registered a callback. The \c callbackData does not
 * outlive the callback.
 */
typedef void (*AChoreographer_vsyncCallback)(
        const AChoreographerFrameCallbackData* callbackData, void* data);

/**
 * Prototype of the function that is called when the display refresh rate
 * changes. It's passed the new vsync period in nanoseconds, as well as the data
 * pointer provided by the application that registered a callback.
 */
typedef void (*AChoreographer_refreshRateCallback)(int64_t vsyncPeriodNanos, void* data);

/**
 * Get the AChoreographer instance for the current thread. This must be called
 * on an ALooper thread.
 *
 * Available since API level 24.
 */
AChoreographer* AChoreographer_getInstance() __INTRODUCED_IN(24);

/**
 * Post a callback to be run when the application should begin rendering the
 * next frame. The data pointer provided will be passed to the callback function
 * when it's called.
 *
 * The callback will only be run for the next frame, not all subsequent frames,
 * so to render continuously the callback should itself call
 * AChoreographer_postFrameCallback.
 *
 * \bug The callback receives the frame time in nanoseconds as a long. On 32-bit
 * systems, long is 32-bit, so the frame time will roll over roughly every two
 * seconds. If your minSdkVersion is 29 or higher, switch to
 * AChoreographer_postFrameCallback64, which uses a 64-bit frame time for all
 * platforms. For older OS versions, you must combine the argument with the
 * upper bits of clock_gettime(CLOCK_MONOTONIC, ...) on 32-bit systems.
 *
 * \deprecated Use AChoreographer_postFrameCallback64, which does not have the
 * bug described above.
 */
void AChoreographer_postFrameCallback(AChoreographer* choreographer,
                                      AChoreographer_frameCallback callback, void* data)
        __INTRODUCED_IN(24) __DEPRECATED_IN(29, "Use AChoreographer_postFrameCallback64 instead");

/**
 * Post a callback to be run when the application should begin rendering the
 * next frame following the specified delay. The data pointer provided will be
 * passed to the callback function when it's called.
 *
 * The callback will only be run for the next frame after the delay, not all
 * subsequent frames, so to render continuously the callback should itself call
 * AChoreographer_postFrameCallbackDelayed.
 *
 * \bug The callback receives the frame time in nanoseconds as a long. On 32-bit
 * systems, long is 32-bit, so the frame time will roll over roughly every two
 * seconds. If your minSdkVersion is 29 or higher, switch to
 * AChoreographer_postFrameCallbackDelayed64, which uses a 64-bit frame time for
 * all platforms. For older OS versions, you must combine the argument with the
 * upper bits of clock_gettime(CLOCK_MONOTONIC, ...) on 32-bit systems.
 *
 * \deprecated Use AChoreographer_postFrameCallbackDelayed64, which does not
 * have the bug described above.
 */
void AChoreographer_postFrameCallbackDelayed(AChoreographer* choreographer,
                                             AChoreographer_frameCallback callback, void* data,
                                             long delayMillis) __INTRODUCED_IN(24)
        __DEPRECATED_IN(29, "Use AChoreographer_postFrameCallbackDelayed64 instead");

/**
 * Post a callback to be run when the application should begin rendering the
 * next frame. The data pointer provided will be passed to the callback function
 * when it's called.
 *
 * The callback will only be run on the next frame, not all subsequent frames,
 * so to render continuously the callback should itself call
 * AChoreographer_postFrameCallback64.
 *
 * Available since API level 29.
 */
void AChoreographer_postFrameCallback64(AChoreographer* choreographer,
                                        AChoreographer_frameCallback64 callback, void* data)
        __INTRODUCED_IN(29);

/**
 * Post a callback to be run when the application should begin rendering the
 * next frame following the specified delay. The data pointer provided will be
 * passed to the callback function when it's called.
 *
 * The callback will only be run for the next frame after the delay, not all
 * subsequent frames, so to render continuously the callback should itself call
 * AChoreographer_postFrameCallbackDelayed64.
 *
 * Available since API level 29.
 */
void AChoreographer_postFrameCallbackDelayed64(AChoreographer* choreographer,
                                               AChoreographer_frameCallback64 callback, void* data,
                                               uint32_t delayMillis) __INTRODUCED_IN(29);

/**
 * Posts a callback to be run when the application should begin rendering the
 * next frame. The data pointer provided will be passed to the callback function
 * when it's called.
 *
 * The callback will only be run for the next frame, not all subsequent frames,
 * so to render continuously the callback should itself call
 * AChoreographer_postVsyncCallback.
 *
 * Available since API level 33.
 */
void AChoreographer_postVsyncCallback(AChoreographer* choreographer,
                                        AChoreographer_vsyncCallback callback, void* data)
        __INTRODUCED_IN(33);

/**
 * Registers a callback to be run when the display refresh rate changes. The
 * data pointer provided will be passed to the callback function when it's
 * called. The same callback may be registered multiple times, provided that a
 * different data pointer is provided each time.
 *
 * If an application registers a callback for this choreographer instance when
 * no new callbacks were previously registered, that callback is guaranteed to
 * be dispatched. However, if the callback and associated data pointer are
 * unregistered prior to running the callback, then the callback may be silently
 * dropped.
 *
 * This api is thread-safe. Any thread is allowed to register a new refresh
 * rate callback for the choreographer instance.
 *
 * Note that in API level 30, this api is not guaranteed to be atomic with
 * DisplayManager. That is, calling Display#getRefreshRate very soon after
 * a refresh rate callback is invoked may return a stale refresh rate. If any
 * Display properties would be required by this callback, then it is recommended
 * to listen directly to DisplayManager.DisplayListener#onDisplayChanged events
 * instead.
 *
 * As of API level 31, this api is guaranteed to have a consistent view with DisplayManager;
 * Display#getRefreshRate is guaranteed to not return a stale refresh rate when invoked from this
 * callback.
 *
 * Available since API level 30.
 */
void AChoreographer_registerRefreshRateCallback(AChoreographer* choreographer,
                                                AChoreographer_refreshRateCallback, void* data)
        __INTRODUCED_IN(30);

/**
 * Unregisters a callback to be run when the display refresh rate changes, along
 * with the data pointer previously provided when registering the callback. The
 * callback is only unregistered when the data pointer matches one that was
 * previously registered.
 *
 * This api is thread-safe. Any thread is allowed to unregister an existing
 * refresh rate callback for the choreographer instance. When a refresh rate
 * callback and associated data pointer are unregistered, then there is a
 * guarantee that when the unregistration completes that that callback will not
 * be run with the data pointer passed.
 *
 * Available since API level 30.
 */
void AChoreographer_unregisterRefreshRateCallback(AChoreographer* choreographer,
                                                  AChoreographer_refreshRateCallback, void* data)
        __INTRODUCED_IN(30);

/**
 * The time in nanoseconds at which the frame started being rendered.
 *
 * Note that this time should \b not be used to advance animation clocks.
 * Instead, see AChoreographerFrameCallbackData_getFrameTimelineExpectedPresentationTimeNanos().
 *
 * Available since API level 33.
 */
int64_t AChoreographerFrameCallbackData_getFrameTimeNanos(
        const AChoreographerFrameCallbackData* data) __INTRODUCED_IN(33);

/**
 * The number of possible frame timelines.
 *
 * Available since API level 33.
 */
size_t AChoreographerFrameCallbackData_getFrameTimelinesLength(
        const AChoreographerFrameCallbackData* data) __INTRODUCED_IN(33);

/**
 * Gets the index of the platform-preferred frame timeline.
 * The preferred frame timeline is the default
 * by which the platform scheduled the app, based on the device configuration.
 *
 * Available since API level 33.
 */
size_t AChoreographerFrameCallbackData_getPreferredFrameTimelineIndex(
        const AChoreographerFrameCallbackData* data) __INTRODUCED_IN(33);

/**
 * Gets the token used by the platform to identify the frame timeline at the given \c index.
 *
 * Available since API level 33.
 *
 * \param index index of a frame timeline, in \f( [0, FrameTimelinesLength) \f). See
 * AChoreographerFrameCallbackData_getFrameTimelinesLength()
 *
 */
AVsyncId AChoreographerFrameCallbackData_getFrameTimelineVsyncId(
        const AChoreographerFrameCallbackData* data, size_t index) __INTRODUCED_IN(33);

/**
 * Gets the time in nanoseconds at which the frame described at the given \c index is expected to
 * be presented. This time should be used to advance any animation clocks.
 *
 * Available since API level 33.
 *
 * \param index index of a frame timeline, in \f( [0, FrameTimelinesLength) \f). See
 * AChoreographerFrameCallbackData_getFrameTimelinesLength()
 */
int64_t AChoreographerFrameCallbackData_getFrameTimelineExpectedPresentationTimeNanos(
        const AChoreographerFrameCallbackData* data, size_t index) __INTRODUCED_IN(33);

/**
 * Gets the time in nanoseconds at which the frame described at the given \c index needs to be
 * ready by in order to be presented on time.
 *
 * Available since API level 33.
 *
 * \param index index of a frame timeline, in \f( [0, FrameTimelinesLength) \f). See
 * AChoreographerFrameCallbackData_getFrameTimelinesLength()
 */
int64_t AChoreographerFrameCallbackData_getFrameTimelineDeadlineNanos(
        const AChoreographerFrameCallbackData* data, size_t index) __INTRODUCED_IN(33);

__END_DECLS

#endif // ANDROID_CHOREOGRAPHER_H

/** @} */
