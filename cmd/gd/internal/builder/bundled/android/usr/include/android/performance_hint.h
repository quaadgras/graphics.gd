/*
 * Copyright (C) 2021 The Android Open Source Project
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
 * @defgroup APerformanceHint Performance Hint Manager
 *
 * APerformanceHint allows apps to create performance hint sessions for groups
 * of threads, and provide hints to the system about the workload of those threads,
 * to help the system more accurately allocate resources for them. It is the NDK
 * counterpart to the Java PerformanceHintManager SDK API.
 *
 * This API is intended for periodic workloads, such as frame production. Clients are
 * expected to create an instance of APerformanceHintManager, create a session with
 * that, and then set a target duration for the session. Then, they can report the actual
 * work duration at the end of each cycle to inform the framework about how long those
 * workloads are taking. The framework will then compare the actual durations to the target
 * duration and attempt to help the client reach a steady state under the target.
 *
 * Unlike reportActualWorkDuration, the "notify..." hints are intended to be sent in
 * advance of large changes in the workload, to prevent them from going over the target
 * when there is a sudden, unforseen change. Their effects are intended to last for only
 * one cycle, after which reportActualWorkDuration will have a chance to catch up.
 * These hints should be used judiciously, only in cases where the workload is changing
 * substantially. To enforce that, they are tracked using a per-app rate limiter to avoid
 * excessive hinting and encourage clients to be mindful about when to send them.
 * @{
 */

/**
 * @file performance_hint.h
 * @brief API for creating and managing a hint session.
 */


#ifndef ANDROID_NATIVE_PERFORMANCE_HINT_H
#define ANDROID_NATIVE_PERFORMANCE_HINT_H

#include <sys/cdefs.h>
#include <jni.h>

/******************************************************************
 *
 * IMPORTANT NOTICE:
 *
 *   This file is part of Android's set of stable system headers
 *   exposed by the Android NDK (Native Development Kit).
 *
 *   Third-party source AND binary code relies on the definitions
 *   here to be FROZEN ON ALL UPCOMING PLATFORM RELEASES.
 *
 *   - DO NOT MODIFY ENUMS (EXCEPT IF YOU ADD NEW 32-BIT VALUES)
 *   - DO NOT MODIFY CONSTANTS OR FUNCTIONAL MACROS
 *   - DO NOT CHANGE THE SIGNATURE OF FUNCTIONS IN ANY WAY
 *   - DO NOT CHANGE THE LAYOUT OR SIZE OF STRUCTURES
 */

#include <stdbool.h>
#include <stdint.h>
#include <unistd.h>

__BEGIN_DECLS

struct APerformanceHintManager;
struct APerformanceHintSession;
struct AWorkDuration;
struct ANativeWindow;
struct ASurfaceControl;

/**
 * {@link AWorkDuration} is an opaque type that represents the breakdown of the
 * actual workload duration in each component internally.
 *
 * A new {@link AWorkDuration} can be obtained using
 * {@link AWorkDuration_create()}, when the client finishes using
 * {@link AWorkDuration}, {@link AWorkDuration_release()} must be
 * called to destroy and free up the resources associated with
 * {@link AWorkDuration}.
 *
 * This file provides a set of functions to allow clients to set the measured
 * work duration of each component on {@link AWorkDuration}.
 *
 * - AWorkDuration_setWorkPeriodStartTimestampNanos()
 * - AWorkDuration_setActualTotalDurationNanos()
 * - AWorkDuration_setActualCpuDurationNanos()
 * - AWorkDuration_setActualGpuDurationNanos()
 */
typedef struct AWorkDuration AWorkDuration;

/**
 * An opaque type representing a handle to a performance hint manager.
 *
 * To use:<ul>
 *    <li>Obtain the performance hint manager instance by calling
 *        {@link APerformanceHint_getManager} function.</li>
 *    <li>Create an {@link APerformanceHintSession} with
 *        {@link APerformanceHint_createSession}.</li>
 *    <li>Get the preferred update rate in nanoseconds with
 *        {@link APerformanceHint_getPreferredUpdateRateNanos}.</li>
 */
typedef struct APerformanceHintManager APerformanceHintManager;

/**
 * An opaque type representing a handle to a performance hint session creation configuration.
 * It is consumed by {@link APerformanceHint_createSessionUsingConfig}.
 *
 * A session creation config encapsulates the required information for a session.
 * Additionally, the caller can set various settings for the session,
 * to be passed during creation, streamlining the session setup process.
 *
 * The caller may reuse this object and modify the settings in it
 * to create additional sessions.
 *
 */
typedef struct ASessionCreationConfig ASessionCreationConfig;

/**
 * An opaque type representing a handle to a performance hint session.
 * A session can only be acquired from a {@link APerformanceHintManager}
 * with {@link APerformanceHint_createSession}
 * or {@link APerformanceHint_createSessionUsingConfig}. It must be
 * freed with {@link APerformanceHint_closeSession} after use.
 *
 * A Session represents a group of threads with an inter-related workload such that hints for
 * their performance should be considered as a unit. The threads in a given session should be
 * long-lived and not created or destroyed dynamically.
 *
 * The work duration API can be used with periodic workloads to dynamically adjust thread
 * performance and keep the work on schedule while optimizing the available power budget.
 * When using the work duration API, the starting target duration should be specified
 * while creating the session, and can later be adjusted with
 * {@link APerformanceHint_updateTargetWorkDuration}. While using the work duration
 * API, the client is expected to call {@link APerformanceHint_reportActualWorkDuration} each
 * cycle to report the actual time taken to complete to the system.
 *
 * Note, methods of {@link APerformanceHintSession_*} are not thread safe so callers must
 * ensure thread safety.
 *
 * All timings should be from `std::chrono::steady_clock` or `clock_gettime(CLOCK_MONOTONIC, ...)`
 */
typedef struct APerformanceHintSession APerformanceHintSession;

typedef struct ANativeWindow ANativeWindow;
typedef struct ASurfaceControl ASurfaceControl;

/**
  * Acquire an instance of the performance hint manager.
  *
  * @return APerformanceHintManager instance on success, nullptr on failure.
  */
APerformanceHintManager* _Nullable APerformanceHint_getManager()
                         __INTRODUCED_IN(__ANDROID_API_T__);

/**
 * Creates a session for the given set of threads and sets their initial target work
 * duration.
 *
 * @param manager The performance hint manager instance.
 * @param threadIds The list of threads to be associated with this session. They must be part of
 *     this process' thread group.
 * @param size The size of the list of threadIds.
 * @param initialTargetWorkDurationNanos The target duration in nanoseconds for the new session.
 *     This must be positive if using the work duration API, or 0 otherwise.
 * @return APerformanceHintSession pointer on success, nullptr on failure.
 */
APerformanceHintSession* _Nullable APerformanceHint_createSession(
        APerformanceHintManager* _Nonnull manager,
        const int32_t* _Nonnull threadIds, size_t size,
        int64_t initialTargetWorkDurationNanos) __INTRODUCED_IN(__ANDROID_API_T__);

/**
 * Creates a session for the given set of threads that are graphics pipeline threads
 * and set their initial target work duration.
 *
 * @param manager The performance hint manager instance.
 * @param config The configuration struct containing required information
 *        to create a session.
 * @return APerformanceHintSession pointer on success, nullptr on failure.
 */
APerformanceHintSession* _Nullable APerformanceHint_createSessionUsingConfig(
        APerformanceHintManager* _Nonnull manager,
        ASessionCreationConfig* _Nonnull config)
        __INTRODUCED_IN(36);

/**
 * Get preferred update rate information for this device.
 *
 * @param manager The performance hint manager instance.
 * @return the preferred update rate supported by device software.
 */
int64_t APerformanceHint_getPreferredUpdateRateNanos(
        APerformanceHintManager* _Nonnull manager) __INTRODUCED_IN(__ANDROID_API_T__);

/**
 * Get maximum number of graphics pipieline threads per-app for this device.
 *
 * @param manager The performance hint manager instance.
 * @return the maximum number of graphics pipeline threads supported by device.
 */
 int APerformanceHint_getMaxGraphicsPipelineThreadsCount(
        APerformanceHintManager* _Nonnull manager) __INTRODUCED_IN(36);

/**
 * Updates this session's target duration for each cycle of work.
 *
 * @param session The performance hint session instance to update.
 * @param targetDurationNanos The new desired duration in nanoseconds. This must be positive.
 * @return 0 on success.
 *         EINVAL if targetDurationNanos is not positive.
 *         EPIPE if communication with the system service has failed.
 */
int APerformanceHint_updateTargetWorkDuration(
        APerformanceHintSession* _Nonnull session,
        int64_t targetDurationNanos) __INTRODUCED_IN(__ANDROID_API_T__);

/**
 * Reports the actual duration for the last cycle of work.
 *
 * The system will attempt to adjust the scheduling and performance of the
 * threads within the thread group to bring the actual duration close to the target duration.
 *
 * @param session The performance hint session instance to update.
 * @param actualDurationNanos The duration of time the thread group took to complete its last
 *     task in nanoseconds. This must be positive.
 * @return 0 on success.
 *         EINVAL if actualDurationNanos is not positive.
 *         EPIPE if communication with the system service has failed.
 */
int APerformanceHint_reportActualWorkDuration(
        APerformanceHintSession* _Nonnull session,
        int64_t actualDurationNanos) __INTRODUCED_IN(__ANDROID_API_T__);

/**
 * Release the performance hint manager pointer acquired via
 * {@link APerformanceHint_createSession}.
 *
 * This cannot be used to close a Java PerformanceHintManager.Session, as its
 * lifecycle is tied to the object in the SDK.
 *
 * @param session The performance hint session instance to release.
 */
void APerformanceHint_closeSession(
        APerformanceHintSession* _Nonnull session) __INTRODUCED_IN(__ANDROID_API_T__);

/**
 * Set a list of threads to the performance hint session. This operation will replace
 * the current list of threads with the given list of threads.
 *
 * @param session The performance hint session instance to update.
 * @param threadIds The list of threads to be associated with this session. They must be part of
 *     this app's thread group.
 * @param size The size of the list of threadIds.
 * @return 0 on success.
 *         EINVAL if the list of thread ids is empty or if any of the thread ids are not part of
               the thread group.
 *         EPIPE if communication with the system service has failed.
 *         EPERM if any thread id doesn't belong to the application.
 */
int APerformanceHint_setThreads(
        APerformanceHintSession* _Nonnull session,
        const pid_t* _Nonnull threadIds,
        size_t size) __INTRODUCED_IN(__ANDROID_API_U__);

/**
 * This tells the session that these threads can be
 * safely scheduled to prefer power efficiency over performance.
 *
 * @param session The performance hint session instance to update.
 * @param enabled The flag which sets whether this session will use power-efficient scheduling.
 * @return 0 on success.
 *         EPIPE if communication with the system service has failed.
 */
int APerformanceHint_setPreferPowerEfficiency(
        APerformanceHintSession* _Nonnull session,
        bool enabled) __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Reports the durations for the last cycle of work.
 *
 * The system will attempt to adjust the scheduling and performance of the
 * threads within the thread group to bring the actual duration close to the target duration.
 *
 * @param session The {@link APerformanceHintSession} instance to update.
 * @param workDuration The {@link AWorkDuration} structure of times the thread group took to
 *     complete its last task in nanoseconds breaking down into different components.
 *
 *     The work period start timestamp and actual total duration must be greater than zero.
 *
 *     The actual CPU and GPU durations must be greater than or equal to zero, and at least one
 *     of them must be greater than zero. When one of them is equal to zero, it means that type
 *     of work was not measured for this workload.
 *
 * @return 0 on success.
 *         EINVAL if any duration is an invalid number.
 *         EPIPE if communication with the system service has failed.
 */
int APerformanceHint_reportActualWorkDuration2(
        APerformanceHintSession* _Nonnull session,
        AWorkDuration* _Nonnull workDuration) __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Informs the framework of an upcoming increase in the workload of a graphics pipeline
 * bound to this session. The user can specify whether the increase is expected to be
 * on the CPU, GPU, or both.
 *
 * Sending hints for both CPU and GPU counts as two separate hints for the purposes of the
 * rate limiter.
 *
 * @param cpu Indicates if the workload increase is expected to affect the CPU.
 * @param gpu Indicates if the workload increase is expected to affect the GPU.
 * @param debugName A required string used to identify this specific hint during
 *        tracing. This debug string will only be held for the duration of the
 *        method, and can be safely discarded after.
 *
 * @return 0 on success.
 *         EINVAL if no hints were requested.
 *         EBUSY if the hint was rate limited.
 *         EPIPE if communication with the system service has failed.
 *         ENOTSUP if the hint is not supported.
 */
int APerformanceHint_notifyWorkloadIncrease(
        APerformanceHintSession* _Nonnull session,
        bool cpu, bool gpu, const char* _Nonnull debugName) __INTRODUCED_IN(36);

/**
 * Informs the framework of an upcoming reset in the workload of a graphics pipeline
 * bound to this session, or the imminent start of a new workload. The user can specify
 * whether the reset is expected to affect the CPU, GPU, or both.
 *
 * Sending hints for both CPU and GPU counts as two separate hints for the purposes of the
 * this load tracking.
 *
 * @param cpu Indicates if the workload reset is expected to affect the CPU.
 * @param gpu Indicates if the workload reset is expected to affect the GPU.
 * @param debugName A required string used to identify this specific hint during
 *        tracing. This debug string will only be held for the duration of the
 *        method, and can be safely discarded after.
 *
 * @return 0 on success.
 *         EINVAL if no hints were requested.
 *         EBUSY if the hint was rate limited.
 *         EPIPE if communication with the system service has failed.
 *         ENOTSUP if the hint is not supported.
 */
int APerformanceHint_notifyWorkloadReset(
        APerformanceHintSession* _Nonnull session,
        bool cpu, bool gpu, const char* _Nonnull debugName) __INTRODUCED_IN(36);

/**
 * Informs the framework of an upcoming one-off expensive frame for a graphics pipeline
 * bound to this session. This frame will be treated as not representative of the workload as a
 * whole, and it will be discarded the purposes of load tracking. The user can specify
 * whether the workload spike is expected to be on the CPU, GPU, or both.
 *
 * Sending hints for both CPU and GPU counts as two separate hints for the purposes of the
 * rate limiter.
 *
 * @param cpu Indicates if the workload spike is expected to affect the CPU.
 * @param gpu Indicates if the workload spike is expected to affect the GPU.
 * @param debugName A required string used to identify this specific hint during
 *        tracing. This debug string will only be held for the duration of the
 *        method, and can be safely discarded after.
 *
 * @return 0 on success.
 *         EINVAL if no hints were requested.
 *         EBUSY if the hint was rate limited.
 *         EPIPE if communication with the system service has failed.
 *         ENOTSUP if the hint is not supported.
 */
int APerformanceHint_notifyWorkloadSpike(
        APerformanceHintSession* _Nonnull session,
        bool cpu, bool gpu, const char* _Nonnull debugName) __INTRODUCED_IN(36);

/**
 * Associates a session with any {@link ASurfaceControl} or {@link ANativeWindow}
 * instances managed by this session.
 *
 * This method is primarily intended for sessions that manage the timing of an entire
 * graphics pipeline end-to-end, such as those using the
 * {@link ASessionCreationConfig_setGraphicsPipeline} API. However, any session directly
 * or indirectly managing a graphics pipeline should still associate themselves with
 * directly relevant ASurfaceControl or ANativeWindow instances for better optimization.
 *
 * To see any benefit from this method, the client must make sure they are updating the framerate
 * of attached surfaces using methods such as {@link ANativeWindow_setFrameRate}, or by updating
 * any associated ASurfaceControls with transactions that have {ASurfaceTransaction_setFrameRate}.
 *
 * @param session The {@link APerformanceHintSession} instance to update.
 * @param nativeWindows A pointer to a list of ANativeWindows associated with this session.
 *        nullptr can be passed to indicate there are no associated ANativeWindows.
 * @param nativeWindowsSize The number of ANativeWindows in the list.
 * @param surfaceControls A pointer to a list of ASurfaceControls associated with this session.
 *        nullptr can be passed to indicate there are no associated ASurfaceControls.
 * @param surfaceControlsSize The number of ASurfaceControls in the list.
 *
 * @return 0 on success.
 *         EPIPE if communication has failed.
 *         ENOTSUP if unsupported.
 *         EINVAL if invalid or empty arguments passed.
 */

int APerformanceHint_setNativeSurfaces(APerformanceHintSession* _Nonnull session,
        ANativeWindow* _Nonnull* _Nullable nativeWindows, int nativeWindowsSize,
        ASurfaceControl* _Nonnull* _Nullable surfaceControls, int surfaceControlsSize)
        __INTRODUCED_IN(36);

/**
 * Creates a new AWorkDuration. When the client finishes using {@link AWorkDuration}, it should
 * call {@link AWorkDuration_release()} to destroy {@link AWorkDuration} and release all resources
 * associated with it.
 *
 * @return AWorkDuration pointer.
 */
AWorkDuration* _Nonnull AWorkDuration_create() __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Destroys a {@link AWorkDuration} and frees all resources associated with it.
 *
 * @param aWorkDuration The {@link AWorkDuration} created by calling {@link AWorkDuration_create()}
 */
void AWorkDuration_release(AWorkDuration* _Nonnull aWorkDuration)
     __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Sets the work period start timestamp in nanoseconds.
 *
 * @param aWorkDuration The {@link AWorkDuration} created by calling {@link AWorkDuration_create()}
 * @param workPeriodStartTimestampNanos The work period start timestamp in nanoseconds based on
 *        CLOCK_MONOTONIC about when the work starts. This timestamp must be greater than zero.
 */
void AWorkDuration_setWorkPeriodStartTimestampNanos(AWorkDuration* _Nonnull aWorkDuration,
        int64_t workPeriodStartTimestampNanos) __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Sets the actual total work duration in nanoseconds.
 *
 * @param aWorkDuration The {@link AWorkDuration} created by calling {@link AWorkDuration_create()}
 * @param actualTotalDurationNanos The actual total work duration in nanoseconds. This number must
 *        be greater than zero.
 */
void AWorkDuration_setActualTotalDurationNanos(AWorkDuration* _Nonnull aWorkDuration,
        int64_t actualTotalDurationNanos) __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Sets the actual CPU work duration in nanoseconds.
 *
 * @param aWorkDuration The {@link AWorkDuration} created by calling {@link AWorkDuration_create()}
 * @param actualCpuDurationNanos The actual CPU work duration in nanoseconds. This number must be
 *        greater than or equal to zero. If it is equal to zero, that means the CPU was not
 *        measured.
 */
void AWorkDuration_setActualCpuDurationNanos(AWorkDuration* _Nonnull aWorkDuration,
        int64_t actualCpuDurationNanos) __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Sets the actual GPU work duration in nanoseconds.
 *
 * @param aWorkDuration The {@link AWorkDuration} created by calling {@link AWorkDuration_create()}.
 * @param actualGpuDurationNanos The actual GPU work duration in nanoseconds, the number must be
 *        greater than or equal to zero. If it is equal to zero, that means the GPU was not
 *        measured.
 */
void AWorkDuration_setActualGpuDurationNanos(AWorkDuration* _Nonnull aWorkDuration,
        int64_t actualGpuDurationNanos) __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Return the APerformanceHintSession wrapped by a Java PerformanceHintManager.Session object.
 *
 * The Java session maintains ownership over the wrapped native session, so it cannot be
 * closed using {@link APerformanceHint_closeSession}.
 *
 * @param env The Java environment where the PerformanceHintManager.Session lives.
 * @param sessionObj The Java Session to unwrap.
 *
 * @return A pointer to the APerformanceHintManager that backs the Java Session.
 */
APerformanceHintSession* _Nonnull APerformanceHint_borrowSessionFromJava(
        JNIEnv* _Nonnull env, jobject _Nonnull sessionObj) __INTRODUCED_IN(36);

/*
 * Creates a new ASessionCreationConfig.
 *
 * When the client finishes using {@link ASessionCreationConfig}, it should
 * call {@link ASessionCreationConfig_release()} to destroy
 * {@link ASessionCreationConfig} and release all resources
 * associated with it.
 *
 * @return ASessionCreationConfig pointer.
 */
ASessionCreationConfig* _Nonnull ASessionCreationConfig_create()
                __INTRODUCED_IN(36);


/**
 * Destroys a {@link ASessionCreationConfig} and frees all
 * resources associated with it.
 *
 * @param config The {@link ASessionCreationConfig}
 *        created by calling {@link ASessionCreationConfig_create()}.
 */
void ASessionCreationConfig_release(
                ASessionCreationConfig* _Nonnull config) __INTRODUCED_IN(36);

/**
 * Sets the tids to be associated with the session to be created.
 *
 * @param config The {@link ASessionCreationConfig}
 *        created by calling {@link ASessionCreationConfig_create()}
 * @param tids The list of tids to be associated with this session. They must be part of
 *        this process' thread group.
 * @param size The size of the list of tids.
 *
 * @return 0 on success.
 *         EINVAL if invalid array pointer or the value of size
 */
int ASessionCreationConfig_setTids(
        ASessionCreationConfig* _Nonnull config,
        const pid_t* _Nonnull tids, size_t size)  __INTRODUCED_IN(36);

/**
 * Sets the initial target work duration in nanoseconds for the session to be created.
 *
 * @param config The {@link ASessionCreationConfig}
 *        created by calling {@link ASessionCreationConfig_create()}.
 * @param targetWorkDurationNanos The parameter to specify a target duration
 *        in nanoseconds for the new session; this value must be positive to use
 *        the work duration API.
 *
 * @return 0 on success.
 *         ENOTSUP if unsupported
 *         EINVAL if invalid value
 */
int ASessionCreationConfig_setTargetWorkDurationNanos(
        ASessionCreationConfig* _Nonnull config,
        int64_t targetWorkDurationNanos)  __INTRODUCED_IN(36);

/**
 * Sets whether power efficiency mode will be enabled for the session.
 * This tells the session that these threads can be
 * safely scheduled to prefer power efficiency over performance.
 *
 * @param config The {@link ASessionCreationConfig}
 *        created by calling {@link ASessionCreationConfig_create()}.
 * @param enabled Whether power efficiency mode will be enabled.
 *
 * @return 0 on success.
 *         ENOTSUP if unsupported
 *         EINVAL if invalid pointer to creation config
 */
int ASessionCreationConfig_setPreferPowerEfficiency(
        ASessionCreationConfig* _Nonnull config, bool enabled)  __INTRODUCED_IN(36);

/**
 * Sessions setting this hint are expected to time the critical path of
 * graphics pipeline from end to end, with the total work duration
 * representing the time from the start of frame production until the
 * buffer is fully finished drawing.
 *
 * It should include any threads on the critical path of that pipeline,
 * up to a limit accessible from {@link getMaxGraphicsPipelineThreadsCount()}.
 *
 * @param config The {@link ASessionCreationConfig}
 *        created by calling {@link ASessionCreationConfig_create()}.
 * @param enabled Whether this session manages a graphics pipeline's critical path.
 *
 * @return 0 on success.
 *         ENOTSUP if unsupported
 *         EINVAL if invalid pointer to creation config or maximum threads for graphics
                  pipeline is reached.
 */
int ASessionCreationConfig_setGraphicsPipeline(
        ASessionCreationConfig* _Nonnull config, bool enabled)  __INTRODUCED_IN(36);

/**
 * Associates a session with any {@link ASurfaceControl} or {@link ANativeWindow}
 * instances managed by this session. See {@link APerformanceHint_setNativeSurfaces}
 * for more details.
 *
 * @param config The {@link ASessionCreationConfig}
 *        created by calling {@link ASessionCreationConfig_create()}.
 * @param nativeWindows A pointer to a list of ANativeWindows associated with this session.
 *        nullptr can be passed to indicate there are no associated ANativeWindows.
 * @param nativeWindowsSize The number of ANativeWindows in the list.
 * @param surfaceControls A pointer to a list of ASurfaceControls associated with this session.
 *        nullptr can be passed to indicate there are no associated ASurfaceControls.
 * @param surfaceControlsSize The number of ASurfaceControls in the list.
 *
 * @return 0 on success.
 *         ENOTSUP if unsupported.
 *         EINVAL if invalid or empty arguments passed.
 */
int ASessionCreationConfig_setNativeSurfaces(
        ASessionCreationConfig* _Nonnull config,
        ANativeWindow* _Nonnull* _Nullable nativeWindows, int nativeWindowsSize,
        ASurfaceControl* _Nonnull* _Nullable surfaceControls, int surfaceControlsSize)
        __INTRODUCED_IN(36);

/**
 * Enable automatic timing mode for sessions using the GRAPHICS_PIPELINE API with an attached
 * surface. In this mode, sessions do not need to report actual durations and only need
 * to keep their thread list up-to-date, set a native surface, call
 * {@link ASessionCreationConfig_setGraphicsPipeline()} to signal that the session is in
 * "graphics pipeline" mode, and then set whether automatic timing is desired for the
 * CPU, GPU, or both, using this method.
 *
 * It is still be beneficial to set an accurate target time, as this may help determine
 * timing information for some workloads where there is less information available from
 * the framework, such as games. Additionally, reported CPU durations will be ignored
 * while automatic CPU timing is enabled, and similarly GPU durations will be ignored
 * when automatic GPU timing is enabled. When both are enabled, the entire
 * reportActualWorkDuration call will be ignored, and the session will be managed
 * completely automatically.
 *
 * This mode will not work unless the client makes sure they are updating the framerate
 * of attached surfaces with methods such as {@link ANativeWindow_setFrameRate}, or updating
 * any associated ASurfaceControls with transactions that have {ASurfaceTransaction_setFrameRate}.
 *
 * @param config The {@link ASessionCreationConfig}
 *        created by calling {@link ASessionCreationConfig_create()}.
 * @param cpu Whether to enable automatic timing for the CPU for this session.
 * @param gpu Whether to enable automatic timing for the GPU for this session.
 *
 * @return 0 on success.
 *         ENOTSUP if unsupported.
 */
int ASessionCreationConfig_setUseAutoTiming(
        ASessionCreationConfig* _Nonnull config,
        bool cpu, bool gpu)
        __INTRODUCED_IN(36);

__END_DECLS

#endif // ANDROID_NATIVE_PERFORMANCE_HINT_H

/** @} */
