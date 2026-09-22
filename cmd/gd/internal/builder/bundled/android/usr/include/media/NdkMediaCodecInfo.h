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
 * @addtogroup Media
 * @{
 */

/**
 * @file NdkMediaCodecInfo.h
 */

/*
 * This file defines an NDK API.
 * Do not remove methods.
 * Do not change method signatures.
 * Do not change the value of constants.
 * Do not change the size of any of the classes defined in here.
 * Do not reference types that are not part of the NDK.
 * Do not #include files that aren't part of the NDK.
 */

#ifndef _NDK_MEDIA_CODEC_INFO_H
#define _NDK_MEDIA_CODEC_INFO_H

#include "NdkMediaError.h"
#include "NdkMediaFormat.h"

__BEGIN_DECLS

struct ACodecAudioCapabilities;
typedef struct ACodecAudioCapabilities ACodecAudioCapabilities;
struct ACodecPerformancePoint;
typedef struct ACodecPerformancePoint ACodecPerformancePoint;
struct ACodecVideoCapabilities;
typedef struct ACodecVideoCapabilities ACodecVideoCapabilities;
struct ACodecEncoderCapabilities;
typedef struct ACodecEncoderCapabilities ACodecEncoderCapabilities;
struct AMediaCodecInfo;
typedef struct AMediaCodecInfo AMediaCodecInfo;

/**
 * A uitlity structure describing the range of two integer values.
 */
typedef struct AIntRange {
    int32_t mLower;
    int32_t mUpper;
} AIntRange;

/**
 * A uitlity structure describing the range of two double values.
 */
typedef struct ADoubleRange {
    double mLower;
    double mUpper;
} ADoubleRange;

// AMediaCodecInfo

/**
 * Get the canonical name of a codec.
 *
 * @return      The char pointer to the canonical name.
 *              Encoded as ASCII and owned by the framework.
 *
 * @return NULL if @param info is invalid.
 */
const char* _Nullable AMediaCodecInfo_getCanonicalName(const AMediaCodecInfo* _Nonnull info)
        __INTRODUCED_IN(36);

typedef enum AMediaCodecKind : int32_t {
    /** invalid media codec info */
    AMediaCodecKind_INVALID = 0,

    /** decoder */
    AMediaCodecKind_DECODER = 1,

    /** encoder */
    AMediaCodecKind_ENCODER = 2,
} AMediaCodecKind;

/**
 * Query the kind of the codec.
 */
AMediaCodecKind AMediaCodecInfo_getKind(const AMediaCodecInfo* _Nonnull info) __INTRODUCED_IN(36);

/**
 * Query if the codec is provided by the Android platform or the device manufacturer.
 *
 * @return 1 if the codec is provided by the device manufacturer
 * @return 0 if the codec is provided by the Android platform
 * @return -1 if @param info is invalid.
 */
int32_t AMediaCodecInfo_isVendor(const AMediaCodecInfo* _Nonnull info) __INTRODUCED_IN(36);

/**
 * The type of codecs.
 */
typedef enum AMediaCodecType : int32_t {
    /**
     * Not a codec type. Used for indicating an invalid operation occurred.
     */
    AMediaCodecType_INVALID_CODEC_INFO = 0,

    /**
     * Software codec.
     *
     * Software-only codecs are more secure as they run in a tighter security sandbox.
     * On the other hand, software-only codecs do not provide any performance guarantees.
     */
    AMediaCodecType_SOFTWARE_ONLY = 1,

    /**
     * Hardware accelerated codec.
     *
     * Hardware codecs generally have higher performance or lower power consumption than
     * software codecs, but since they are specific to each device,
     * the actual performance details can vary.
     */
    AMediaCodecType_HARDWARE_ACCELERATED = 2,

    /**
     * Software codec but have device access.
     * Mainly referring to software codecs provided by vendors.
     */
    AMediaCodecType_SOFTWARE_WITH_DEVICE_ACCESS = 3,
} AMediaCodecType;

/**
 * Query if the codec is SOFTWARE_ONLY, HARDWARE_ACCELERATED or SOFTWARE_WITH_DEVICE_ACCESS.
 *
 * @return INVALID_CODEC_INFO if @param info is invalid.
 */
AMediaCodecType AMediaCodecInfo_getMediaCodecInfoType(
        const AMediaCodecInfo* _Nonnull info) __INTRODUCED_IN(36);

/**
 * Get the supported media type of the codec.
 *
 * @return  The char pointer to the media type(e.g. "video/hevc").
 *          It is ASCII encoded and owned by the framework with infinite lifetime.
 *
 * @return NULL if @param info is invalid.
 */
const char* _Nullable AMediaCodecInfo_getMediaType(const AMediaCodecInfo* _Nonnull info)
        __INTRODUCED_IN(36);

/**
 * Get the max number of the supported concurrent codec instances.
 *
 * This is a hint for an upper bound. Applications should not expect to successfully
 * operate more instances than the returned value, but the actual number of
 * concurrently operable instances may be less as it depends on the available
 * resources at time of use.
 *
 * @return -1 if @param info is invalid.
 */
int32_t AMediaCodecInfo_getMaxSupportedInstances(const AMediaCodecInfo* _Nonnull info)
        __INTRODUCED_IN(36);

/**
 * Query codec feature capabilities.
 *
 * These features are supported to be used by the codec.  These
 * include optional features that can be turned on, as well as
 * features that are always on.
 *
 * @param featureName   Get valid feature names from the defined constants in NdkMediaCodecInfo.h
 *                      with prefix AMediaCoecInfo_FEATURE_.
 *
 * @return 1 if the feature is supported;
 * @return 0 if the feature is unsupported;
 * @return -1 if @param info or @param featureName is invalid.
 */
int32_t AMediaCodecInfo_isFeatureSupported(const AMediaCodecInfo* _Nonnull info,
        const char* _Nonnull featureName) __INTRODUCED_IN(36);

/**
 * Query codec feature requirements.
 *
 * These features are required to be used by the codec, and as such,
 * they are always turned on.
 *
 * @param featureName   Get valid feature names from the defined constants in NdkMediaCodecInfo.h
 *                      with prefix AMediaCoecInfo_FEATURE_.
 *
 * @return 1 if the feature is required;
 * @return 0 if the feature is not required;
 * @return -1 if @param info or @param featureName is invalid.
 */
int32_t AMediaCodecInfo_isFeatureRequired(const AMediaCodecInfo* _Nonnull info,
        const char* _Nonnull featureName) __INTRODUCED_IN(36);

/**
 * Query whether codec supports a given @param format.
 *
 * @return 1 if the format is supported;
 * @return 0 if the format is unsupported;
 * @return -1 if @param info or @param format is invalid.
 */
int32_t AMediaCodecInfo_isFormatSupported(const AMediaCodecInfo* _Nonnull info,
        const AMediaFormat* _Nonnull format) __INTRODUCED_IN(36);

/**
 * Get the ACodecAudioCapabilities from the given AMediaCodecInfo.
 *
 * @param outAudioCaps        The pointer to the output ACodecAudioCapabilities.
 *                            It is owned by the framework and has an infinite lifetime.
 *
 * @return AMEDIA_OK if successfully got the ACodecAudioCapabilities.
 * @return AMEDIA_ERROR_UNSUPPORTED if the codec is not an audio codec.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if @param info or @param outAudioCaps is invalid.
 */
media_status_t AMediaCodecInfo_getAudioCapabilities(const AMediaCodecInfo* _Nonnull info,
        const ACodecAudioCapabilities* _Nullable * _Nonnull outAudioCaps) __INTRODUCED_IN(36);

/**
 * Get the ACodecVideoCapabilities from the given AMediaCodecInfo.
 *
 * @param outVideoCaps        The pointer to the output ACodecVideoCapabilities.
 *                            It is owned by the framework and has an infinite lifetime.
 *
 * @return AMEDIA_OK if successfully got the ACodecVideoCapabilities.
 * @return AMEDIA_ERROR_UNSUPPORTED if the codec is not a video codec.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if @param info or @param outVideoCaps is invalid.
 */
media_status_t AMediaCodecInfo_getVideoCapabilities(const AMediaCodecInfo* _Nonnull info,
        const ACodecVideoCapabilities* _Nullable * _Nonnull outVideoCaps) __INTRODUCED_IN(36);

/**
 * Get the ACodecEncoderCapabilities from the given AMediaCodecInfo.
 *
 * @param outEncoderCaps        The pointer to the output ACodecEncoderCapabilities.
 *                              It is owned by the framework and has an infinite lifetime.
 *
 * @return AMEDIA_OK if successfully got the ACodecEncoderCapabilities.
 * @return AMEDIA_ERROR_UNSUPPORTED if the codec is not an encoder.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if @param info or @param outEncoderCaps is invalid.
 */
media_status_t AMediaCodecInfo_getEncoderCapabilities(const AMediaCodecInfo* _Nonnull info,
        const ACodecEncoderCapabilities* _Nullable * _Nonnull outEncoderCaps) __INTRODUCED_IN(36);

// ACodecAudioCapabilities

/**
 * Get the range of supported bitrates in bits/second.
 *
 * @param outRange  The pointer to the range of supported bitrates.
 *                  Users are responsible for allocating a valid AIntRange structure and
 *                  managing the lifetime of it.
 *
 * @return AMEDIA_OK if got bitrates successfully.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param audioCaps and @param outRange is invalid.
 */
media_status_t ACodecAudioCapabilities_getBitrateRange(
        const ACodecAudioCapabilities* _Nonnull audioCaps,
        AIntRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Get the array of supported sample rates
 *
 * The array is sorted in ascending order.
 *
 * @param outArrayPtr   The pointer to the output sample rates array.
 *                      The array is owned by the framework and has an infinite lifetime.
 * @param outCount      The size of the output array.
 *
 * @return AMEDIA_OK if the codec supports only discrete values.
 * Otherwise, it returns AMEDIA_ERROR_UNSUPPORTED.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param audioCaps, @param outArrayPtr
 * and @param outCount is invalid.
 */
media_status_t ACodecAudioCapabilities_getSupportedSampleRates(
        const ACodecAudioCapabilities* _Nonnull audioCaps,
        const int* _Nullable * _Nonnull outArrayPtr,
        size_t* _Nonnull outCount) __INTRODUCED_IN(36);

/**
 * Get the array of supported sample rate ranges.
 *
 * The array is sorted in ascending order, and the ranges are distinct (non-intersecting).
 *
 * @param outArrayPtr   The pointer to the out sample rate ranges array.
 *                      The array is owned by the framework and has an infinite lifetime.
 * @param outCount      The size of the out array.
 *
 * @return AMEDIA_OK if got the sample rate ranges successfully.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param audioCaps, @param outArrayPtr
 * and @param outCount is invalid.
 */
media_status_t ACodecAudioCapabilities_getSupportedSampleRateRanges(
        const ACodecAudioCapabilities* _Nonnull audioCaps,
        const AIntRange* _Nullable * _Nonnull outArrayPtr, size_t* _Nonnull outCount)
        __INTRODUCED_IN(36);

/**
 * Return the maximum number of input channels supported.
 *
 * @return -1 if @param audioCaps is invalid.
 */
int32_t ACodecAudioCapabilities_getMaxInputChannelCount(
        const ACodecAudioCapabilities* _Nonnull audioCaps) __INTRODUCED_IN(36);

/**
 * Returns the minimum number of input channels supported.
 * This is often 1, but does vary for certain mime types.
 *
 * @return -1 if @param audioCaps is invalid.
 */
int32_t ACodecAudioCapabilities_getMinInputChannelCount(
        const ACodecAudioCapabilities* _Nonnull audioCaps) __INTRODUCED_IN(36);

/**
 * Get an array of ranges representing the number of input channels supported.
 * The codec supports any number of input channels within this range.
 * For many codecs, this will be a single range [1..N], for some N.
 *
 * The array is sorted in ascending order, and the ranges are distinct (non-intersecting).
 *
 * @param outArrayPtr   The pointer to the output array of input-channels ranges.
 *                      The array is owned by the framework and has an infinite lifetime.
 * @param outCount      The size of the output array.
 *
 * @return AMEDIA_OK if got the input channel array successfully.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param audioCaps, @param outArrayPtr
 * and @param outCount is invalid.
 */
media_status_t ACodecAudioCapabilities_getInputChannelCountRanges(
        const ACodecAudioCapabilities* _Nonnull audioCaps,
        const AIntRange* _Nullable * _Nonnull outArrayPtr, size_t* _Nonnull outCount)
        __INTRODUCED_IN(36);

/**
 * Query whether the sample rate is supported by the codec.
 *
 * @return 1 if the sample rate is supported.
 * @return 0 if the sample rate is unsupported
 * @return -1 if @param audioCaps is invalid.
 */
int32_t ACodecAudioCapabilities_isSampleRateSupported(
        const ACodecAudioCapabilities* _Nonnull audioCaps,
        int32_t sampleRate) __INTRODUCED_IN(36);

// ACodecPerformancePoint

/**
 * Create a performance point for a given frame size and frame rate.
 *
 * Video performance points are a set of standard performance points defined by number of
 * pixels, pixel rate and frame rate. Performance point represents an upper bound. This
 * means that it covers all performance points with fewer pixels, pixel rate and frame
 * rate.
 *
 * Users are responsible for calling
 * ACodecPerformancePoint_destroy(ACodecPerformancePoint *performancePoint) after use.
 *
 * @param width width of the frame in pixels
 * @param height height of the frame in pixels
 * @param frameRate frame rate in frames per second
 */
ACodecPerformancePoint* _Nonnull ACodecPerformancePoint_create(int32_t width, int32_t height,
        int32_t frameRate) __INTRODUCED_IN(36);

/**
 * Delete a created performance point.
 */
void ACodecPerformancePoint_destroy(
        ACodecPerformancePoint* _Nullable performancePoint) __INTRODUCED_IN(36);

/**
 * Checks whether the performance point covers a media format.
 *
 * @param format Stream format considered.
 *
 * @return 1 if the performance point covers the format.
 * @return 0 if the performance point does not cover the format.
 * @return -1 if @param performancePoint or @param format is invalid.
 */
int32_t ACodecPerformancePoint_coversFormat(const ACodecPerformancePoint* _Nonnull performancePoint,
        const AMediaFormat* _Nonnull format) __INTRODUCED_IN(36);

/**
 * Checks whether a performance point covers another performance point.
 *
 * Use this method to determine if a performance point advertised by a codec covers the
 * performance point required. This method can also be used for loose ordering as this
 * method is transitive.
 *
 * A Performance point represents an upper bound. This means that
 * it covers all performance points with fewer pixels, pixel rate and frame rate.
 *
 * @return 1 if @param one covers @param another.
 * @return 0 if @param one does not cover @param another.
 * @return -1 if @param one or @param another is invalid.
 */
int32_t ACodecPerformancePoint_covers(const ACodecPerformancePoint* _Nonnull one,
        const ACodecPerformancePoint* _Nonnull another) __INTRODUCED_IN(36);

/**
 * Checks whether two performance points are equal.
 *
 * @return 1 if @param one and @param another are equal.
 * @return 0 if @param one and @param another are not equal.
 * @return -1 if @param one or @param another is invalid.
 */
int32_t ACodecPerformancePoint_equals(const ACodecPerformancePoint* _Nonnull one,
        const ACodecPerformancePoint* _Nonnull another) __INTRODUCED_IN(36);

// ACodecVideoCapabilities

/**
 * Get the range of supported bitrates in bits/second.
 *
 * @param outRange  The pointer to the range of output bitrates.
 *                  Users are responsible for allocating a valid AIntRange structure and
 *                  managing the lifetime of it.
 *
 * @return AMEDIA_OK if got the supported bitrates successfully.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param videoCaps and @param outRange is invalid.
 */
media_status_t ACodecVideoCapabilities_getBitrateRange(
        const ACodecVideoCapabilities* _Nonnull videoCaps,
        AIntRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Get the range of supported video widths.
 *
 * @param outRange  The pointer to the range of output supported widths.
 *                  Users are responsible for allocating a valid AIntRange structure and
 *                  managing the lifetime of it.
 *
 * @return AMEDIA_OK if got the supported video widths successfully.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param videoCaps and @param outRange is invalid.
 */
media_status_t ACodecVideoCapabilities_getSupportedWidths(
        const ACodecVideoCapabilities* _Nonnull videoCaps,
        AIntRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Get the range of supported video heights.
 *
 * @param outRange  The pointer to the range of output supported heights.
 *                  Users are responsible for allocating a valid AIntRange structure and
 *                  managing the lifetime of it.
 *
 * @return AMEDIA_OK if got the supported video heights successfully.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param videoCaps and @param outRange is invalid.
 */
media_status_t ACodecVideoCapabilities_getSupportedHeights(
        const ACodecVideoCapabilities* _Nonnull videoCaps,
        AIntRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Get the alignment requirement for video width (in pixels).
 *
 * This is a power-of-2 value that video width must be a multiple of.
 *
 * @return -1 if @param videoCaps is invalid.
 */
int32_t ACodecVideoCapabilities_getWidthAlignment(
        const ACodecVideoCapabilities* _Nonnull videoCaps) __INTRODUCED_IN(36);

/**
 * Return the alignment requirement for video height (in pixels).
 *
 * This is a power-of-2 value that video height must be a multiple of.
 *
 * @return -1 if @param videoCaps is invalid.
 */
int32_t ACodecVideoCapabilities_getHeightAlignment(
        const ACodecVideoCapabilities* _Nonnull videoCaps) __INTRODUCED_IN(36);

/**
 * Get the range of supported frame rates.
 *
 * This is not a performance indicator. Rather, it expresses the limits specified in the coding
 * standard, based on the complexities of encoding material for later playback at a certain
 * frame rate, or the decoding of such material in non-realtime.
 *
 * @param outRange  The pointer to the range of output supported frame rates.
 *                  Users are responsible for allocating a valid AIntRange structure and
 *                  managing the lifetime of it.
 *
 * @return AMEDIA_OK if got the frame rate range successfully.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param videoCaps and @param outRange is invalid.
 */
media_status_t ACodecVideoCapabilities_getSupportedFrameRates(
        const ACodecVideoCapabilities* _Nonnull videoCaps,
        AIntRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Get the range of supported video widths for a video height.
 *
 * @param outRange      The pointer to the range of supported widths.
 *                      Users are responsible for allocating a valid AIntRange structure and
 *                      managing the lifetime of it.
 *
 * @return AMEDIA_OK if got the supported video width range successfully.
 * @return AMEDIA_ERROR_UNSUPPORTED if the height query is not supported.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param videoCaps and @param outRange is invalid.
 */
media_status_t ACodecVideoCapabilities_getSupportedWidthsFor(
        const ACodecVideoCapabilities* _Nonnull videoCaps, int32_t height,
        AIntRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Get the range of supported video heights for a video width.
 *
 * @param outRange      The pointer to the range of supported heights.
 *                      Users are responsible for allocating a valid AIntRange structure and
 *                      managing the lifetime of it.
 *
 * @return AMEDIA_OK if got the supported video height range successfully.
 * @return AMEDIA_ERROR_UNSUPPORTED if the width query is not supported.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param videoCaps and @param outRange is invalid.
 */
media_status_t ACodecVideoCapabilities_getSupportedHeightsFor(
        const ACodecVideoCapabilities* _Nonnull videoCaps, int32_t width,
        AIntRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Get the range of supported video frame rates for a video size.
 *
 * This is not a performance indicator.  Rather, it expresses the limits specified in the coding
 * standard, based on the complexities of encoding material of a given size for later playback at
 * a certain frame rate, or the decoding of such material in non-realtime.
 *
 * @param outRange      The pointer to the range of frame rates.
 *                      Users are responsible for allocating a valid ADoubleRange structure and
 *                      managing the lifetime of it.
 *
 * @return AMEDIA_OK if got the supported video frame rates successfully.
 * @return AMEDIA_ERROR_UNSUPPORTED if the size query is not supported.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param videoCaps and @param outRange is invalid.
 */
media_status_t ACodecVideoCapabilities_getSupportedFrameRatesFor(
        const ACodecVideoCapabilities* _Nonnull videoCaps, int32_t width, int32_t height,
        ADoubleRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Get the range of achievable video frame rates for a video size.
 *
 * This is based on manufacturer's performance measurements for this device and codec.
 * The measurements may not be available for all codecs or devices.
 *
 * @param outRange      The pointer to the range of frame rates.
  *                     Users are responsible for allocating a valid ADoubleRange structure and
 *                      managing the lifetime of it.
 *
 * @return AMEDIA_OK if got the achievable video frame rates successfully.
 * @return AMEDIA_ERROR_UNSUPPORTED if the codec did not publish any measurement data.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param videoCaps and @param outRange is invalid.
 */
media_status_t ACodecVideoCapabilities_getAchievableFrameRatesFor(
        const ACodecVideoCapabilities* _Nonnull videoCaps, int32_t width, int32_t height,
        ADoubleRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Get the supported performance points.
 *
 * This API returns the supported performance points in a sequence and stores them in a pointer at
 * outPerformancePoint. Initially, set the pointer pointed to by outPerformancePoint to NULL.
 * On successive calls, keep the last pointer value. When the sequence is over
 * AMEDIA_ERROR_UNSUPPORTED will be returned and the pointer will be set to NULL.
 *
 * @param outPerformancePoint   a pointer to (the ACodecPerformancePoint pointer) where the next
 *                              performance point will be stored.
 *                              The ACodecPerformancePoint object is owned by the framework
 *                              and has an infinite lifecycle.
 *
 * @return AMEDIA_OK if successfully got the performance point.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if @param videoCaps or @param outPerformancePoint
 * is invalid.
 * @return AMEDIA_ERROR_UNSUPPORTED if there are no more supported performance points.
 * *outPerformancePoint will also be set to NULL in this case.
 */
media_status_t ACodecVideoCapabilities_getNextSupportedPerformancePoint(
        const ACodecVideoCapabilities* _Nonnull videoCaps,
        const ACodecPerformancePoint* _Nullable * _Nonnull outPerformancePoint) __INTRODUCED_IN(36);

/**
 * Get whether a given video size and frameRate combination is supported.
 *
 * @return 1 if the size and rate are supported.
 * @return 0 if they are not supported.
 * @return -1 if @param videoCaps is invalid.
 */
int32_t ACodecVideoCapabilities_areSizeAndRateSupported(
        const ACodecVideoCapabilities* _Nonnull videoCaps,
        int32_t width, int32_t height, double frameRate) __INTRODUCED_IN(36);

/**
 * Get whether a given video size is supported.
 *
 * @return 1 if the size is supported.
 * @return 0 if the size is not supported.
 * @return -1 if @param videoCaps is invalid.
 */
int32_t ACodecVideoCapabilities_isSizeSupported(const ACodecVideoCapabilities* _Nonnull videoCaps,
        int32_t width, int32_t height) __INTRODUCED_IN(36);

// ACodecEncoderCapabilities

/**
 * Get the supported range of quality values.
 *
 * Quality is implementation-specific. As a general rule, a higher quality
 * setting results in a better image quality and a lower compression ratio.
 *
 * @param outRange      The pointer to the range of quality values.
 *                      Users are responsible for allocating a valid AIntRange structure and
 *                      managing the lifetime of it.
 *
 * @return AMEDIA_OK if successfully got the quality range.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param videoCaps and @param outRange is invalid.
 */
media_status_t ACodecEncoderCapabilities_getQualityRange(
        const ACodecEncoderCapabilities* _Nonnull encoderCaps,
        AIntRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Get the supported range of encoder complexity values.
 *
 * Some codecs may support multiple complexity levels, where higher complexity values use more
 * encoder tools (e.g. perform more intensive calculations) to improve the quality or the
 * compression ratio. Use a lower value to save power and/or time.
 *
 * @param outRange      The pointer to the range of encoder complexity values.
 *                      Users are responsible for allocating a valid AIntRange structure and
 *                      managing the lifetime of it.
 *
 * @return AMEDIA_OK if successfully got the complexity range.
 * @return AMEDIA_ERROR_INVALID_PARAMETER if any of @param videoCaps and @param outRange is invalid.
 */
media_status_t ACodecEncoderCapabilities_getComplexityRange(
        const ACodecEncoderCapabilities* _Nonnull encoderCaps,
        AIntRange* _Nonnull outRange) __INTRODUCED_IN(36);

/**
 * Encoder bitrate modes.
 */
typedef enum ABitrateMode : int32_t {
    ABITRATE_MODE_CQ = 0,
    ABITRATE_MODE_VBR = 1,
    ABITRATE_MODE_CBR = 2,
    ABITRATE_MODE_CBR_FD = 3
} ABitrateMode;

/**
 * Query whether a bitrate mode is supported.
 *
 * @return 1 if the bitrate mode is supported.
 * @return 0 if the bitrate mode is unsupported.
 * @return -1 if @param encoderCaps is invalid.
 */
int32_t ACodecEncoderCapabilities_isBitrateModeSupported(
        const ACodecEncoderCapabilities* _Nonnull encoderCaps, ABitrateMode mode)
        __INTRODUCED_IN(36);

// Feature Names

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video decoder only: codec supports seamless resolution changes.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_AdaptivePlayback __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video decoder only: codec supports secure decryption.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_SecurePlayback __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video or audio decoder only: codec supports tunneled playback.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_TunneledPlayback __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * The timestamp of each output buffer is derived from the timestamp of the input
 * buffer that produced the output. If false, the timestamp of each output buffer is
 * derived from the timestamp of the first input buffer.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_DynamicTimestamp __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * decoder only: codec supports partial (including multiple) access units
 * per input buffer.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_FrameParsing __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * codec supports multiple access units (for decoding, or to output for
 * encoders). If false, the codec only supports single access units. Producing multiple
 * access units for output is an optional feature.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_MultipleFrames __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video decoder only: codec supports queuing partial frames.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_PartialFrame __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video encoder only: codec supports intra refresh.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_IntraRefresh __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * decoder only: codec supports low latency decoding.
 * If supported, clients can enable the low latency mode for the decoder.
 * When the mode is enabled, the decoder doesn't hold input and output data more than
 * required by the codec standards.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_LowLatency __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video encoder only: codec supports quantization parameter bounds.
 * @see MediaFormat#KEY_VIDEO_QP_MAX
 * @see MediaFormat#KEY_VIDEO_QP_MIN
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_QpBounds __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video encoder only: codec supports exporting encoding statistics.
 * Encoders with this feature can provide the App clients with the encoding statistics
 * information about the frame.
 * The scope of encoding statistics is controlled by
 * {@link MediaFormat#KEY_VIDEO_ENCODING_STATISTICS_LEVEL}.
 *
 * @see MediaFormat#KEY_VIDEO_ENCODING_STATISTICS_LEVEL
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_EncodingStatistics __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video encoder only: codec supports HDR editing.
 *
 * HDR editing support means that the codec accepts 10-bit HDR
 * input surface, and it is capable of generating any HDR
 * metadata required from both YUV and RGB input when the
 * metadata is not present. This feature is only meaningful when
 * using an HDR capable profile (and 10-bit HDR input).
 *
 * This feature implies that the codec is capable of encoding at
 * least one HDR format, and that it supports RGBA_1010102 as
 * well as P010, and optionally RGBA_FP16 input formats, and
 * that the encoder can generate HDR metadata for all supported
 * HDR input formats.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_HdrEditing __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video encoder only: codec supports HLG editing.
 *
 * HLG editing support means that the codec accepts 10-bit HDR
 * input surface in both YUV and RGB pixel format. This feature
 * is only meaningful when using a 10-bit (HLG) profile and
 * 10-bit input.
 *
 * This feature implies that the codec is capable of encoding
 * 10-bit format, and that it supports RGBA_1010102 as
 * well as P010, and optionally RGBA_FP16 input formats.
 *
 * The difference between this feature and {@link
 * FEATURE_HdrEditing} is that HLG does not require the
 * generation of HDR metadata and does not use an explicit HDR
 * profile.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_HlgEditing __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video decoder only: codec supports dynamically
 * changing color aspects.
 *
 * If true, the codec can propagate color aspect changes during
 * decoding. This is only meaningful at session boundaries, e.g.
 * upon processing Picture Parameter Sets prior to a new IDR.
 * The color aspects may come from the bitstream, or may be
 * provided using {@link MediaCodec#setParameters} calls.
 *
 * If the codec supports both 8-bit and 10-bit profiles, this
 * feature means that the codec can dynamically switch between 8
 * and 10-bit profiles, but this is restricted to Surface mode
 * only.
 *
 * If the device supports HDR transfer functions, switching
 * between SDR and HDR transfer is also supported. Together with
 * the previous clause this means that switching between SDR and
 * HDR sessions are supported in Surface mode, as SDR is
 * typically encoded at 8-bit and HDR at 10-bit.
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_DynamicColorAspects __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video encoder only: codec supports region of interest encoding.
 *
 * RoI encoding support means the codec accepts information that specifies the relative
 * importance of different portions of each video frame. This allows the encoder to
 * separate a video frame into critical and non-critical regions, and use more bits
 * (better quality) to represent the critical regions and de-prioritize non-critical
 * regions. In other words, the encoder chooses a negative qp bias for the critical
 * portions and a zero or positive qp bias for the non-critical portions.
 *
 * At a basic level, if the encoder decides to encode each frame with a uniform
 * quantization value 'qpFrame' and a 'qpBias' is chosen/suggested for an LCU of the
 * frame, then the actual qp of the LCU will be 'qpFrame + qpBias', although this value
 * can be clamped basing on the min-max configured qp bounds for the current encoding
 * session.
 *
 * In a shot, if a group of LCUs pan out quickly they can be marked as non-critical
 * thereby enabling the encoder to reserve fewer bits during their encoding. Contrarily,
 * LCUs that remain in shot for a prolonged duration can be encoded at better quality in
 * one frame thereby setting-up an excellent long-term reference for all future frames.
 *
 * Note that by offsetting the quantization of each LCU, the overall bit allocation will
 * differ from the originally estimated bit allocation, and the encoder will adjust the
 * frame quantization for subsequent frames to meet the bitrate target. An effective
 * selection of critical regions can set-up a golden reference and this can compensate
 * for the bit burden that was introduced due to encoding RoI's at better quality.
 * On the other hand, an ineffective choice of critical regions might increase the
 * quality of certain parts of the image but this can hamper quality in subsequent frames.
 *
 * @see MediaCodec#PARAMETER_KEY_QP_OFFSET_MAP
 * @see MediaCodec#PARAMETER_KEY_QP_OFFSET_RECTS
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_Roi __INTRODUCED_IN(36);

/**
 * A feature constant for use with AMediaCodecInfo_isFeature* methods
 *
 * video decoder only: codec supports detaching the
 * output surface when in Surface mode.
 * If true, the codec can be configured in Surface mode
 * without an actual surface (in detached surface mode).
 * @see MediaCodec#CONFIGURE_FLAG_DETACHED_SURFACE
 */
extern const char* _Nonnull AMediaCodecInfo_FEATURE_DetachedSurface __INTRODUCED_IN(36);


__END_DECLS

#endif //_NDK_MEDIA_CODEC_INFO_H

/** @} */