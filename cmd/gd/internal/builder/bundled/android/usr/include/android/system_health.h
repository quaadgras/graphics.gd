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
* @defgroup SystemHealth
* @{
*/

/**
 * @file system_health.h
 */

#ifndef _ANDROID_SYSTEM_HEALTH_H
#define _ANDROID_SYSTEM_HEALTH_H

#include <sys/cdefs.h>

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


#include <stdint.h>
#include <sys/types.h>

#if !defined(__INTRODUCED_IN)
#define __INTRODUCED_IN(__api_level) /* nothing */
#endif

#ifdef __cplusplus
extern "C" {
#endif

/**
 * Params used to customize the calculation of CPU headroom.
 *
 * Also see {@link ASystemHealth_getCpuHeadroom}.
 */
typedef struct ACpuHeadroomParams ACpuHeadroomParams;

/**
 * Params used to customize the calculation of GPU headroom.
 *
 * Also see {@link ASystemHealth_getGpuHeadroom}.
 */
typedef struct AGpuHeadroomParams AGpuHeadroomParams;

/**
 * Creates a new instance of ACpuHeadroomParams.
 *
 * When the client finishes using {@link ACpuHeadroomParams},
 * {@link ACpuHeadroomParams_destroy()} must be called to destroy
 * and free up the resources associated with {@link ACpuHeadroomParams}.
 *
 * Available since API level 36.
 *
 * @return A new instance of ACpuHeadroomParams.
 */
ACpuHeadroomParams *_Nonnull ACpuHeadroomParams_create()
__INTRODUCED_IN(36);

enum ACpuHeadroomCalculationType {
    /**
     * Use the minimum headroom value within the calculation window.
     * Introduced in API level 36.
     */
    ACPU_HEADROOM_CALCULATION_TYPE_MIN = 0,
    /**
     * Use the average headroom value within the calculation window.
     * Introduced in API level 36.
     */
    ACPU_HEADROOM_CALCULATION_TYPE_AVERAGE = 1,
};
typedef enum ACpuHeadroomCalculationType ACpuHeadroomCalculationType;

enum AGpuHeadroomCalculationType {
    /**
     * Use the minimum headroom value within the calculation window.
     * Introduced in API level 36.
     */
    AGPU_HEADROOM_CALCULATION_TYPE_MIN = 0,
    /**
     * Use the average headroom value within the calculation window.
     * Introduced in API level 36.
     */
    AGPU_HEADROOM_CALCULATION_TYPE_AVERAGE = 1,
};
typedef enum AGpuHeadroomCalculationType AGpuHeadroomCalculationType;

/**
 * Sets the headroom calculation window size in ACpuHeadroomParams.
 *
 * Available since API level 36.
 *
 * @param params The params to be set.
 * @param windowMillis The window size in milliseconds ranged from [50, 10000]. The smaller the
 *                     window size, the larger fluctuation in the headroom value should be expected.
 *                     The default value can be retrieved from the
 *                     {@link #ACpuHeadroomParams_getCalculationWindowMillis} if not set. The device
 *                     will try to use the closest feasible window size to this param.
 */
void ACpuHeadroomParams_setCalculationWindowMillis(ACpuHeadroomParams *_Nonnull params,
                                                   int windowMillis)
__INTRODUCED_IN(36);

/**
 * Gets the headroom calculation window size in ACpuHeadroomParams.
 *
 * Available since API level 36.
 *
 * @param params The params to be set.
 * @return This will return the default value chosen by the device if the params is not set.
 */
int ACpuHeadroomParams_getCalculationWindowMillis(ACpuHeadroomParams *_Nonnull params)
__INTRODUCED_IN(36);

/**
 * Sets the headroom calculation window size in AGpuHeadroomParams.
 *
 * Available since API level 36.
 *
 * @param params The params to be set.
 * @param windowMillis The window size in milliseconds ranged from [50, 10000]. The smaller the
 *                     window size, the larger fluctuation in the headroom value should be expected.
 *                     The default value can be retrieved from the
 *                     {@link #AGpuHeadroomParams_getCalculationWindowMillis} if not set. The device
 *                     will try to use the closest feasible window size to this param.
 */
void AGpuHeadroomParams_setCalculationWindowMillis(AGpuHeadroomParams *_Nonnull params,
                                                   int windowMillis)
__INTRODUCED_IN(36);

/**
 * Gets the headroom calculation window size in AGpuHeadroomParams.
 *
 * Available since API level 36.
 *
 * @param params The params to be set.
 * @return This will return the default value chosen by the device if the params is not set.
 */
int AGpuHeadroomParams_getCalculationWindowMillis(AGpuHeadroomParams *_Nonnull params)
__INTRODUCED_IN(36);

/**
 * Sets the headroom calculation type in ACpuHeadroomParams.
 *
 * Available since API level 36.
 *
 * @param params The params to be set.
 * @param calculationType The headroom calculation type.
 */
void ACpuHeadroomParams_setCalculationType(ACpuHeadroomParams *_Nonnull params,
                                           ACpuHeadroomCalculationType calculationType)
__INTRODUCED_IN(36);

/**
 * Gets the headroom calculation type in ACpuHeadroomParams.
 *
 * Available since API level 36.
 *
 * @param params The params to be set.
 * @return The headroom calculation type.
 */
ACpuHeadroomCalculationType
ACpuHeadroomParams_getCalculationType(ACpuHeadroomParams *_Nonnull params)
__INTRODUCED_IN(36);

/**
 * Sets the headroom calculation type in AGpuHeadroomParams.
 *
 * Available since API level 36.
 *
 * @param params The params to be set.
 * @param calculationType The headroom calculation type.
 */
void AGpuHeadroomParams_setCalculationType(AGpuHeadroomParams *_Nonnull params,
                                           AGpuHeadroomCalculationType calculationType)
__INTRODUCED_IN(36);

/**
 * Gets the headroom calculation type in AGpuHeadroomParams.
 *
 * Available since API level 36.
 *
 * @param params The params to be set.
 * @return The headroom calculation type.
 */
AGpuHeadroomCalculationType
AGpuHeadroomParams_getCalculationType(AGpuHeadroomParams *_Nonnull params)
__INTRODUCED_IN(36);

/**
 * Sets the thread TIDs to track in ACpuHeadroomParams.
 *
 * Available since API level 36.
 *
 * @param params The params to be set.
 * @param tids Non-null array of TIDs, maximum 5.
 * @param tidsSize The size of the tids array.
 */
void ACpuHeadroomParams_setTids(ACpuHeadroomParams *_Nonnull params, const int *_Nonnull tids,
                                int tidsSize)
__INTRODUCED_IN(36);

/**
 * Creates a new instance of AGpuHeadroomParams.
 *
 * When the client finishes using {@link AGpuHeadroomParams},
 * {@link AGpuHeadroomParams_destroy()} must be called to destroy
 * and free up the resources associated with {@link AGpuHeadroomParams}.
 *
 * Available since API level 36.
 *
 * @return A new instance of AGpuHeadroomParams.
 */
AGpuHeadroomParams *_Nonnull AGpuHeadroomParams_create()
__INTRODUCED_IN(36);

/**
 * Deletes the ACpuHeadroomParams instance.
 *
 * Available since API level 36.
 *
 * @param params The params to be deleted.
 */
void ACpuHeadroomParams_destroy(ACpuHeadroomParams *_Nonnull params)
__INTRODUCED_IN(36);

/**
 * Deletes the AGpuHeadroomParams instance.
 *
 * Available since API level 36.
 *
 * @param params The params to be deleted.
 */
void AGpuHeadroomParams_destroy(AGpuHeadroomParams *_Nonnull params)
__INTRODUCED_IN(36);

/**
 * Provides an estimate of available CPU capacity headroom of the device.
 *
 * The value can be used by the calling application to determine if the workload was CPU bound and
 * then take action accordingly to ensure that the workload can be completed smoothly. It can also
 * be used with the thermal status and headroom to determine if reducing the CPU bound workload can
 * help reduce the device temperature to avoid thermal throttling.
 *
 * Available since API level 36.
 *
 * @param params The params to customize the CPU headroom calculation, or nullptr to use the default
 * @param outHeadroom Non-null output pointer to a single float, which will be set to the CPU
 *                    headroom value. The value will be a single value or `Float.NaN` if it's
 *                    temporarily unavailable.
 *                    Each valid value ranges from [0, 100], where 0 indicates no more cpu resources
 *                    can be granted.
 * @return 0 on success
 *         EPIPE if failed to get the CPU headroom.
 *         EPERM if the TIDs do not belong to the same process.
 *         ENOTSUP if API or requested params is unsupported.
 */
int ASystemHealth_getCpuHeadroom(const ACpuHeadroomParams *_Nullable params,
                                 float *_Nonnull outHeadroom)
__INTRODUCED_IN(36);

/**
 * Provides an estimate of available GPU capacity headroom of the device.
 *
 * The value can be used by the calling application to determine if the workload was GPU bound and
 * then take action accordingly to ensure that the workload can be completed smoothly. It can also
 * be used with the thermal status and headroom to determine if reducing the GPU bound workload can
 * help reduce the device temperature to avoid thermal throttling.
 *
 * Available since API level 36
 *
 * @param params The params to customize the GPU headroom calculation, or nullptr to use the default
 * @param outHeadroom Non-null output pointer to a single float, which will be set to the GPU
 *                    headroom value. The value will be a single value or `Float.NaN` if it's
 *                    temporarily unavailable.
 *                    Each valid value ranges from [0, 100], where 0 indicates no more gpu resources
 *                    can be granted.
 * @return 0 on success
 *         EPIPE if failed to get the GPU headroom.
 *         ENOTSUP if API or requested params is unsupported.
 */
int ASystemHealth_getGpuHeadroom(const AGpuHeadroomParams *_Nullable params,
                                 float *_Nonnull outHeadroom)
__INTRODUCED_IN(36);

/**
 * Gets minimum polling interval for calling {@link ASystemHealth_getCpuHeadroom} in milliseconds.
 *
 * The getCpuHeadroom API may return cached result if called more frequently than the interval.
 *
 * Available since API level 36.
 *
 * @param outMinIntervalMillis Non-null output pointer to a int64_t, which
 *                will be set to the minimum polling interval in milliseconds.
 * @return 0 on success
 *         EPIPE if failed to get the minimum polling interval.
 *         ENOTSUP if API is unsupported.
 */
int ASystemHealth_getCpuHeadroomMinIntervalMillis(int64_t *_Nonnull outMinIntervalMillis)
__INTRODUCED_IN(36);

/**
 * Gets minimum polling interval for calling {@link ASystemHealth_getGpuHeadroom} in milliseconds.
 *
 * The getGpuHeadroom API may return cached result if called more frequent than the interval.
 *
 * Available since API level 36.
 *
 * @param outMinIntervalMillis Non-null output pointer to a int64_t, which
 *                will be set to the minimum polling interval in milliseconds.
 * @return 0 on success
 *         EPIPE if failed to get the minimum polling interval.
 *         ENOTSUP if API is unsupported.
 */
int ASystemHealth_getGpuHeadroomMinIntervalMillis(int64_t *_Nonnull outMinIntervalMillis)
__INTRODUCED_IN(36);

#ifdef __cplusplus
}
#endif

#endif // _ANDROID_SYSTEM_HEALTH_H

/** @} */
