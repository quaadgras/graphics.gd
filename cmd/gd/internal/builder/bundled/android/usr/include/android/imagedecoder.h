/*
 * Copyright (C) 2019 The Android Open Source Project
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
 * @defgroup ImageDecoder Android Image Decoder
 *
 * Functions for converting encoded images into RGBA pixels.
 *
 * Similar to the Java counterpart android.graphics.ImageDecoder, it can be used
 * to decode images in the following formats:
 * - JPEG
 * - PNG
 * - GIF
 * - WebP
 * - BMP
 * - ICO
 * - WBMP
 * - HEIF
 * - Digital negatives (via the DNG SDK)
 * <p>It has similar options for scaling, cropping, and choosing the output format.
 * Unlike the Java API, which can create an android.graphics.Bitmap or
 * android.graphics.drawable.Drawable object, AImageDecoder decodes directly
 * into memory provided by the client. For more information, see the
 * <a href="https://developer.android.com/ndk/guides/image-decoder">Image decoder</a>
 * developer guide.
 * @{
 */

/**
 * @file imagedecoder.h
 * @brief API for decoding images.
 */

#ifndef ANDROID_IMAGE_DECODER_H
#define ANDROID_IMAGE_DECODER_H

#include "bitmap.h"
#include <android/rect.h>
#include <stdint.h>

#if !defined(__INTRODUCED_IN)
#define __INTRODUCED_IN(__api_level) /* nothing */
#endif

#ifdef __cplusplus
extern "C" {
#endif

struct AAsset;

/**
 *  {@link AImageDecoder} functions result code.
 *
 *  Introduced in API 30.
 *
 *  Many functions will return this to indicate success
 *  ({@link ANDROID_IMAGE_DECODER_SUCCESS}) or the reason for the failure. On
 *  failure, any out-parameters should be considered uninitialized, except where
 *  specified. Use {@link AImageDecoder_resultToString} for a readable
 *  version of the result code.
 */
enum {
    /**
     * Decoding was successful and complete.
     */
    ANDROID_IMAGE_DECODER_SUCCESS = 0,
    /**
     * The input is incomplete.
     */
    ANDROID_IMAGE_DECODER_INCOMPLETE = -1,
    /**
     * The input contained an error after decoding some lines.
     */
    ANDROID_IMAGE_DECODER_ERROR = -2,
    /**
     * Could not convert. For example, attempting to decode an image with
     * alpha to an opaque format.
     */
    ANDROID_IMAGE_DECODER_INVALID_CONVERSION = -3,
    /**
     * The scale is invalid. It may have overflowed, or it may be incompatible
     * with the current alpha setting.
     */
    ANDROID_IMAGE_DECODER_INVALID_SCALE = -4,
    /**
     * Some other parameter is invalid.
     */
    ANDROID_IMAGE_DECODER_BAD_PARAMETER = -5,
    /**
     * Input was invalid before decoding any pixels.
     */
    ANDROID_IMAGE_DECODER_INVALID_INPUT = -6,
    /**
     * A seek was required and it failed.
     */
    ANDROID_IMAGE_DECODER_SEEK_ERROR = -7,
    /**
     * Some other error. For example, an internal allocation failed.
     */
    ANDROID_IMAGE_DECODER_INTERNAL_ERROR = -8,
    /**
     * AImageDecoder did not recognize the format.
     */
    ANDROID_IMAGE_DECODER_UNSUPPORTED_FORMAT = -9,
    /**
     * The animation has reached the end.
     */
    ANDROID_IMAGE_DECODER_FINISHED = -10,
    /**
     * This method cannot be called while the AImageDecoder is in its current
     * state. For example, various setters (like {@link AImageDecoder_setTargetSize})
     * can only be called while the AImageDecoder is set to decode the first
     * frame of an animation. This ensures that any blending and/or restoring
     * prior frames works correctly.
     */
    ANDROID_IMAGE_DECODER_INVALID_STATE = -11,
};

/**
 * Return a constant string value representing the error code.
 *
 * Introduced in API 31.
 *
 * Pass the return value from an {@link AImageDecoder} method (e.g.
 * {@link AImageDecoder_decodeImage}) for a text string representing the error
 * code.
 *
 * Errors:
 * - Returns null for a value out of range.
 */
const char* _Nullable AImageDecoder_resultToString(int)__INTRODUCED_IN(31);

struct AImageDecoder;

/**
 * Opaque handle for decoding images.
 *
 * Introduced in API 30
 *
 * Create using one of the following:
 * - {@link AImageDecoder_createFromAAsset}
 * - {@link AImageDecoder_createFromFd}
 * - {@link AImageDecoder_createFromBuffer}
 *
 * After creation, {@link AImageDecoder_getHeaderInfo} can be used to retrieve
 * information about the encoded image. Other functions, like
 * {@link AImageDecoder_setTargetSize}, can be used to specify how to decode, and
 * {@link AImageDecoder_decodeImage} will decode into client provided memory.
 *
 * {@link AImageDecoder} objects are NOT thread-safe, and should not be shared across
 * threads.
 */
typedef struct AImageDecoder AImageDecoder;

/**
 * Create a new {@link AImageDecoder} from an {@link AAsset}.
 *
 * Available since API level 30.
 *
 * @param asset {@link AAsset} containing encoded image data. Client is still
 *              responsible for calling {@link AAsset_close} on it, which may be
 *              done after deleting the returned {@link AImageDecoder}.
 * @param outDecoder On success (i.e. return value is
 *                   {@link ANDROID_IMAGE_DECODER_SUCCESS}), this will be set to
 *                   a newly created {@link AImageDecoder}. Caller is
 *                   responsible for calling {@link AImageDecoder_delete} on it.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_INCOMPLETE}: The asset was truncated before
 *   reading the image header.
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: One of the parameters is
 *   null.
 * - {@link ANDROID_IMAGE_DECODER_INVALID_INPUT}: There is an error in the
 *   header.
 * - {@link ANDROID_IMAGE_DECODER_SEEK_ERROR}: The asset failed to seek.
 * - {@link ANDROID_IMAGE_DECODER_INTERNAL_ERROR}: Some other error, like a
 *   failure to allocate memory.
 * - {@link ANDROID_IMAGE_DECODER_UNSUPPORTED_FORMAT}: The format is not
 *   supported.
 */
int AImageDecoder_createFromAAsset(struct AAsset* _Nonnull asset,
                                   AImageDecoder* _Nullable * _Nonnull outDecoder)
        __INTRODUCED_IN(30);

/**
 * Create a new {@link AImageDecoder} from a file descriptor.
 *
 * Available since API level 30.
 *
 * @param fd Seekable, readable, open file descriptor for encoded data.
 *           Client is still responsible for closing it, which may be done
 *           after deleting the returned {@link AImageDecoder}.
 * @param outDecoder On success (i.e. return value is
 *                   {@link ANDROID_IMAGE_DECODER_SUCCESS}), this will be set to
 *                   a newly created {@link AImageDecoder}. Caller is
 *                   responsible for calling {@link AImageDecoder_delete} on it.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_INCOMPLETE}: The file was truncated before
 *   reading the image header.
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The {@link AImageDecoder} is
 *   null, or |fd| does not represent a valid, seekable file descriptor.
 * - {@link ANDROID_IMAGE_DECODER_INVALID_INPUT}: There is an error in the
 *   header.
 * - {@link ANDROID_IMAGE_DECODER_SEEK_ERROR}: The descriptor failed to seek.
 * - {@link ANDROID_IMAGE_DECODER_INTERNAL_ERROR}: Some other error, like a
 *   failure to allocate memory.
 * - {@link ANDROID_IMAGE_DECODER_UNSUPPORTED_FORMAT}: The format is not
 *   supported.
 */
int AImageDecoder_createFromFd(int fd, AImageDecoder* _Nullable * _Nonnull outDecoder)
        __INTRODUCED_IN(30);

/**
 * Create a new AImageDecoder from a buffer.
 *
 * Available since API level 30.
 *
 * @param buffer Pointer to encoded data. Must be valid for the entire time
 *               the {@link AImageDecoder} is used.
 * @param length Byte length of buffer.
 * @param outDecoder On success (i.e. return value is
 *                   {@link ANDROID_IMAGE_DECODER_SUCCESS}), this will be set to
 *                   a newly created {@link AImageDecoder}. Caller is
 *                   responsible for calling {@link AImageDecoder_delete} on it.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_INCOMPLETE}: The encoded image was truncated before
 *   reading the image header.
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: One of the parameters is
 *   invalid.
 * - {@link ANDROID_IMAGE_DECODER_INVALID_INPUT}: There is an error in the
 *   header.
 * - {@link ANDROID_IMAGE_DECODER_INTERNAL_ERROR}: Some other error, like a
 *   failure to allocate memory.
 * - {@link ANDROID_IMAGE_DECODER_UNSUPPORTED_FORMAT}: The format is not
 *   supported.
 */
int AImageDecoder_createFromBuffer(const void* _Nonnull buffer, size_t length,
                                   AImageDecoder* _Nullable * _Nonnull outDecoder)
        __INTRODUCED_IN(30);

/**
 * Delete the AImageDecoder.
 * @param decoder {@link AImageDecoder} object created with one of AImageDecoder_createFrom...
 *        functions.
 * Available since API level 30.
 */
void AImageDecoder_delete(AImageDecoder* _Nullable decoder) __INTRODUCED_IN(30);

/**
 * Choose the desired output format.
 *
 * If the encoded image represents an animation, this must be called while on
 * the first frame (e.g. before calling {@link AImageDecoder_advanceFrame} or
 * after calling {@link AImageDecoder_rewind}).
 *
 * Available since API level 30.
 *
 * @param format {@link AndroidBitmapFormat} to use for the output.
 * @param decoder an {@link AImageDecoder} object.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure. On failure, the
 *         {@link AImageDecoder} uses the format it was already planning
 *         to use (either its default or a previously successful setting
 *         from this function).
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The
 *   {@link AImageDecoder} is null or |format| does not correspond to an
 *   {@link AndroidBitmapFormat}.
 * - {@link ANDROID_IMAGE_DECODER_INVALID_CONVERSION}: The
 *   {@link AndroidBitmapFormat} is incompatible with the image.
 * - {@link ANDROID_IMAGE_DECODER_INVALID_STATE}: The animation is not on
 *   the first frame.
 */
int AImageDecoder_setAndroidBitmapFormat(AImageDecoder* _Nonnull decoder,
        int32_t format) __INTRODUCED_IN(30);

/**
 * Specify whether the output's pixels should be unpremultiplied.
 *
 * By default, {@link AImageDecoder_decodeImage} will premultiply the pixels, if they have alpha.
 * Pass true to this method to leave them unpremultiplied. This has no effect on an
 * opaque image.
 *
 * If the encoded image represents an animation, this must be called while on
 * the first frame (e.g. before calling {@link AImageDecoder_advanceFrame} or
 * after calling {@link AImageDecoder_rewind}).
 *
 * Available since API level 30.
 *
 * @param decoder an {@link AImageDecoder} object.
 * @param unpremultipliedRequired Pass true to leave the pixels unpremultiplied.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_INVALID_CONVERSION}: Unpremultiplied is not
 *   possible due to an existing scale set by
 *   {@link AImageDecoder_setTargetSize}.
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The
 *   {@link AImageDecoder} is null.
 * - {@link ANDROID_IMAGE_DECODER_INVALID_STATE}: The animation is not on
 *   the first frame.
 */
int AImageDecoder_setUnpremultipliedRequired(AImageDecoder* _Nonnull decoder,
                                             bool unpremultipliedRequired) __INTRODUCED_IN(30);

/**
 * Choose the dataspace for the output.
 *
 * Ignored by {@link ANDROID_BITMAP_FORMAT_A_8}, which does not support
 * an {@link ADataSpace}.
 *
 * If the encoded image represents an animation, this must be called while on
 * the first frame (e.g. before calling {@link AImageDecoder_advanceFrame} or
 * after calling {@link AImageDecoder_rewind}).
 *
 * Available since API level 30.
 *
 * @param decoder an {@link AImageDecoder} object.
 * @param dataspace The {@link ADataSpace} to decode into. An ADataSpace
 *                  specifies how to interpret the colors. By default,
 *                  AImageDecoder will decode into the ADataSpace specified by
 *                  {@link AImageDecoderHeaderInfo_getDataSpace}. If this
 *                  parameter is set to a different ADataSpace, AImageDecoder
 *                  will transform the output into the specified ADataSpace.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The
 *   {@link AImageDecoder} is null or |dataspace| does not correspond to an
 *   {@link ADataSpace} value.
 * - {@link ANDROID_IMAGE_DECODER_INVALID_STATE}: The animation is not on
 *   the first frame.
 */
int AImageDecoder_setDataSpace(AImageDecoder* _Nonnull decoder, int32_t dataspace)
        __INTRODUCED_IN(30);

/**
 * Specify the output size for a decoded image.
 *
 * Future calls to {@link AImageDecoder_decodeImage} will sample or scale the
 * encoded image to reach the desired size. If a crop rect is set (via
 * {@link AImageDecoder_setCrop}), it must be contained within the dimensions
 * specified by width and height, and the output image will be the size of the
 * crop rect.
 *
 * If the encoded image represents an animation, this must be called while on
 * the first frame (e.g. before calling {@link AImageDecoder_advanceFrame} or
 * after calling {@link AImageDecoder_rewind}).
 *
 * It is strongly recommended to use setTargetSize only for downscaling, as it
 * is often more efficient to scale-up when rendering than up-front due to
 * reduced overall memory.
 *
 * Available since API level 30.
 *
 * @param decoder an {@link AImageDecoder} object.
 * @param width Width of the output (prior to cropping).
 *              This will affect future calls to
 *              {@link AImageDecoder_getMinimumStride}, which will now return
 *              a value based on this width.
 * @param height Height of the output (prior to cropping).
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The
 *   {@link AImageDecoder} is null.
 * - {@link ANDROID_IMAGE_DECODER_INVALID_SCALE}: |width| or |height| is <= 0,
 *   the size is too big, any existing crop is not contained by the new image
 *   dimensions, or the scale is incompatible with a previous call to
 *   {@link AImageDecoder_setUnpremultipliedRequired}(true).
 * - {@link ANDROID_IMAGE_DECODER_INVALID_STATE}: The animation is not on
 *   the first frame.
 */
int AImageDecoder_setTargetSize(AImageDecoder* _Nonnull decoder, int32_t width,
                                int32_t height) __INTRODUCED_IN(30);

/**
 * Compute the dimensions to use for a given sampleSize.
 *
 * Although AImageDecoder can scale to an arbitrary target size (see
 * {@link AImageDecoder_setTargetSize}), some sizes may be more efficient than
 * others. This computes the most efficient target size to use to reach a
 * particular sampleSize.
 *
 * Available since API level 30.
 *
 * @param decoder an {@link AImageDecoder} object.
 * @param sampleSize A subsampling rate of the original image. Must be greater
 *                   than or equal to 1. A sampleSize of 2 means to skip every
 *                   other pixel/line, resulting in a width and height that are
 *                   1/2 of the original dimensions, with 1/4 the number of
 *                   pixels.
 * @param width Out parameter for the width sampled by sampleSize, and rounded
 *              in the direction that the decoder can do most efficiently.
 * @param height Out parameter for the height sampled by sampleSize, and rounded
 *               in the direction that the decoder can do most efficiently.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The
 *   {@link AImageDecoder}, |width| or |height| is null or |sampleSize| is < 1.
 */
int AImageDecoder_computeSampledSize(const AImageDecoder* _Nonnull decoder, int sampleSize,
                                     int32_t* _Nonnull width, int32_t* _Nonnull height)
        __INTRODUCED_IN(30);

/**
 * Specify how to crop the output after scaling (if any).
 *
 * Future calls to {@link AImageDecoder_decodeImage} will crop their output to
 * the specified {@link ARect}. Clients will only need to allocate enough memory
 * for the cropped ARect.
 *
 * If the encoded image represents an animation, this must be called while on
 * the first frame (e.g. before calling {@link AImageDecoder_advanceFrame} or
 * after calling {@link AImageDecoder_rewind}).
 *
 * Available since API level 30.
 *
 * @param decoder an {@link AImageDecoder} object.
 * @param crop Rectangle describing a crop of the decode. It must be contained inside of
 *             the (possibly scaled, by {@link AImageDecoder_setTargetSize})
 *             image dimensions. This will affect future calls to
 *             {@link AImageDecoder_getMinimumStride}, which will now return a
 *             value based on the width of the crop. An empty ARect -
 *             specifically { 0, 0, 0, 0 } - may be used to remove the cropping
 *             behavior. Any other empty or unsorted ARects will result in
 *             returning {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The
 *   {@link AImageDecoder} is null, or the crop is not contained by the
 *   (possibly scaled) image dimensions.
 * - {@link ANDROID_IMAGE_DECODER_INVALID_STATE}: The animation is not on
 *   the first frame.
 */
int AImageDecoder_setCrop(AImageDecoder* _Nonnull decoder, ARect crop) __INTRODUCED_IN(30);

struct AImageDecoderHeaderInfo;
/**
 * Opaque handle for representing information about the encoded image.
 *
 * Introduced in API 30
 *
 * Retrieved using {@link AImageDecoder_getHeaderInfo} and passed to methods
 * like {@link AImageDecoderHeaderInfo_getWidth} and
 * {@link AImageDecoderHeaderInfo_getHeight}.
 */
typedef struct AImageDecoderHeaderInfo AImageDecoderHeaderInfo;

/**
 * Return an opaque handle for reading header info.
 *
 * This is owned by the {@link AImageDecoder} and will be destroyed when the
 * AImageDecoder is destroyed via {@link AImageDecoder_delete}.
 *
 * @param decoder an {@link AImageDecoder} object.
 *
 * Available since API level 30.
 */
const AImageDecoderHeaderInfo* _Nonnull  AImageDecoder_getHeaderInfo(
        const AImageDecoder* _Nonnull decoder) __INTRODUCED_IN(30);

/**
 * Report the native width of the encoded image. This is also the logical
 * pixel width of the output, unless {@link AImageDecoder_setTargetSize} is
 * used to choose a different size or {@link AImageDecoder_setCrop} is used to
 * set a crop rect.
 *
 * Available since API level 30.
 */
int32_t AImageDecoderHeaderInfo_getWidth(const AImageDecoderHeaderInfo* _Nonnull)
        __INTRODUCED_IN(30);

/**
 * Report the native height of the encoded image. This is also the logical
 * pixel height of the output, unless {@link AImageDecoder_setTargetSize} is
 * used to choose a different size or {@link AImageDecoder_setCrop} is used to
 * set a crop rect.
 *
 * Available since API level 30.
 */
int32_t AImageDecoderHeaderInfo_getHeight(const AImageDecoderHeaderInfo* _Nonnull)
        __INTRODUCED_IN(30);

/**
 * Report the mimeType of the encoded image.
 *
 * Available since API level 30.
 *
 * @return a string literal describing the mime type.
 */
const char* _Nonnull  AImageDecoderHeaderInfo_getMimeType(
        const AImageDecoderHeaderInfo* _Nonnull) __INTRODUCED_IN(30);

/**
 * Report the {@link AndroidBitmapFormat} the AImageDecoder will decode to
 * by default. {@link AImageDecoder} will try to choose one that is sensible
 * for the image and the system. Note that this does not indicate the
 * encoded format of the image.
 *
 * Available since API level 30.
 */
int32_t AImageDecoderHeaderInfo_getAndroidBitmapFormat(
        const AImageDecoderHeaderInfo* _Nonnull) __INTRODUCED_IN(30);

/**
 * Report how the {@link AImageDecoder} will handle alpha by default. If the image
 * contains no alpha (according to its header), this will return
 * {@link ANDROID_BITMAP_FLAGS_ALPHA_OPAQUE}. If the image may contain alpha,
 * this returns {@link ANDROID_BITMAP_FLAGS_ALPHA_PREMUL}, because
 * {@link AImageDecoder_decodeImage} will premultiply pixels by default.
 *
 * Available since API level 30.
 *
 * Starting in API level 31, an AImageDecoder may contain multiple frames of an
 * animation, but this method still only reports whether the first frame has
 * alpha.
 */
int AImageDecoderHeaderInfo_getAlphaFlags(
        const AImageDecoderHeaderInfo* _Nonnull) __INTRODUCED_IN(30);

/**
 * Report the dataspace the AImageDecoder will decode to by default.
 *
 * By default, {@link AImageDecoder_decodeImage} will not do any color
 * conversion.
 *
 * Available since API level 30.
 *
 * @return The {@link ADataSpace} representing the way the colors
 *         are encoded (or {@link ADATASPACE_UNKNOWN} if there is not a
 *         corresponding ADataSpace). This specifies how to interpret the colors
 *         in the decoded image, unless {@link AImageDecoder_setDataSpace} is
 *         called to decode to a different ADataSpace.
 *
 *         Note that ADataSpace only exposes a few values. This may return
 *         {@link ADATASPACE_UNKNOWN}, even for Named ColorSpaces, if they have
 *         no corresponding {@link ADataSpace}.
 */
int32_t AImageDecoderHeaderInfo_getDataSpace(
        const AImageDecoderHeaderInfo* _Nonnull) __INTRODUCED_IN(30);

/**
 * Return the minimum stride that can be used in
 * {@link AImageDecoder_decodeImage}.
 *
 * This stride provides no padding, meaning it will be exactly equal to the
 * width times the number of bytes per pixel for the {@link AndroidBitmapFormat}
 * being used.
 *
 * If the output is scaled (via {@link AImageDecoder_setTargetSize}) and/or
 * cropped (via {@link AImageDecoder_setCrop}), this takes those into account.
 *
 * @param decoder an {@link AImageDecoder} object.
 *
 * Available since API level 30.
 */
size_t AImageDecoder_getMinimumStride(AImageDecoder* _Nonnull decoder) __INTRODUCED_IN(30);

/**
 * Decode the image into pixels, using the settings of the {@link AImageDecoder}.
 *
 * Available since API level 30.
 *
 * Starting in API level 31, it can be used to decode all of the frames of an
 * animated image (i.e. GIF, WebP) using new APIs. Internally,
 * AImageDecoder keeps track of its "current frame" - that is, the frame that
 * will be decoded by a call to AImageDecoder_decodeImage. At creation time, the
 * current frame is always the first frame, and multiple calls to this method
 * will each decode the first frame. {@link AImageDecoder_advanceFrame} advances
 * the current frame to the following frame, so that future calls to this method
 * will decode that frame. Some frames may update only part of the image. They
 * may only update a sub-rectangle (see {@link
 * AImageDecoderFrameInfo_getFrameRect}), or they may have alpha (see
 * {@link AImageDecoderFrameInfo_hasAlphaWithinBounds}). In these cases, this
 * method assumes that the prior frame is still residing in the |pixels| buffer,
 * decodes only the new portion, and blends it with the buffer. Frames that change
 * the entire |pixels| buffer are "independent", and do not require the prior
 * frame to remain in the buffer. The first frame is always independent. A
 * sophisticated client can use information from the {@link AImageDecoderFrameInfo}
 * to determine whether other frames are independent, or what frames they rely on.
 *
 * If the current frame is marked {@link ANDROID_IMAGE_DECODER_DISPOSE_OP_PREVIOUS},
 * AImageDecoder_decodeImage will store the |pixels| buffer prior to decoding
 * (note: this only happens for the first in a string of consecutive
 * ANDROID_IMAGE_DECODER_DISPOSE_OP_PREVIOUS frames). After advancing to the
 * following frame, AImageDecoder_decodeImage will restore that buffer prior to
 * decoding that frame. This is the default behavior, but it can be disabled
 * by passing false to {@link AImageDecoder_setInternallyHandleDisposePrevious}.
 *
 * Ignoring timing information, display, etc, a client wishing to decode all
 * frames of an animated image may conceptually use code like the following:
 *
 * while (true) {
 *   int result = AImageDecoder_decodeImage(decoder, pixels, stride, size);
 *   if (result != ANDROID_IMAGE_DECODER_SUCCESS) break;
 *
 *   // Display or save the image in |pixels|, keeping the buffer intact for
 *   // AImageDecoder to decode the next frame correctly.
 *   Application_viewImage(pixels);
 *
 *   result = AImageDecoder_advanceFrame(decoder);
 *   if (result != ANDROID_IMAGE_DECODER_SUCCESS) break;
 * }
 *
 * @param decoder Opaque object representing the decoder.
 * @param pixels On success, will be filled with the result
 *               of the decode. Must be large enough to hold |size| bytes.
 * @param stride Width in bytes of a single row. Must be at least
 *               {@link AImageDecoder_getMinimumStride} and a multiple of the
 *               bytes per pixel of the {@link AndroidBitmapFormat}.
 * @param size Size of the pixel buffer in bytes. Must be at least
 *             stride * (height - 1) +
 *             {@link AImageDecoder_getMinimumStride}.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_INCOMPLETE}: The image was truncated. A
 *   partial image was decoded, and undecoded lines have been initialized to all
 *   zeroes.
 * - {@link ANDROID_IMAGE_DECODER_ERROR}: The image contained an error. A
 *   partial image was decoded, and undecoded lines have been initialized to all
 *   zeroes.
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The {@link AImageDecoder} or
 *   |pixels| is null, the stride is not large enough or not pixel aligned, or
 *   |size| is not large enough.
 * - {@link ANDROID_IMAGE_DECODER_SEEK_ERROR}: The asset or file descriptor
 *   failed to seek.
 * - {@link ANDROID_IMAGE_DECODER_INTERNAL_ERROR}: Some other error, like a
 *   failure to allocate memory.
 * - {@link ANDROID_IMAGE_DECODER_FINISHED}: The input contains no
 *   more frames. No decoding occurred. The client must call
 *   {@link AImageDecoder_rewind} before calling
 *   {@link AImageDecoder_decodeImage} again.
 */
int AImageDecoder_decodeImage(AImageDecoder* _Nonnull decoder,
                              void* _Nonnull pixels, size_t stride,
                              size_t size) __INTRODUCED_IN(30);

/**
 * Return true iff the image is animated - i.e. has multiple frames.
 *
 * Introduced in API 31.
 *
 * A single frame GIF is considered to *not* be animated. This may require
 * seeking past the first frame to verify whether there is a following frame.
 *
 * @param decoder an {@link AImageDecoder} object.
 *
 * Errors:
 * - returns false if |decoder| is null.
 */
bool AImageDecoder_isAnimated(AImageDecoder* _Nonnull decoder)
        __INTRODUCED_IN(31);

enum {
    /**
     * Reported by {@link AImageDecoder_getRepeatCount} if the
     * animation should repeat forever.
     *
     * Introduced in API 31
     */
    ANDROID_IMAGE_DECODER_INFINITE = INT32_MAX,
};

/**
 * Report how many times the animation should repeat.
 *
 * Introduced in API 31.
 *
 * This does not include the first play through. e.g. a repeat
 * count of 4 means that each frame is played 5 times.
 *
 * {@link ANDROID_IMAGE_DECODER_INFINITE} means to repeat forever.
 *
 * This may require seeking.
 *
 * For non-animated formats, this returns 0. It may return non-zero for
 * an image with only one frame (i.e. {@link AImageDecoder_isAnimated} returns
 * false) if the encoded image contains a repeat count.
 *
 * @param decoder an {@link AImageDecoder} object.
 * @return Number of times to repeat on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The AImageDecoder
 *   is null.
 */
int32_t AImageDecoder_getRepeatCount(AImageDecoder* _Nonnull decoder)
        __INTRODUCED_IN(31);

/**
 * Advance to the next frame in the animation.
 *
 * Introduced in API 31.
 *
 * The AImageDecoder keeps track internally which frame it is ready to decode
 * (the "current frame"). Initially it is set to decode the first frame, and
 * each call to {@link AImageDecoder_decodeImage} will continue to decode
 * the same frame until this method (or {@link AImageDecoder_rewind})
 * is called.
 *
 * Note that this can be used to skip a frame without decoding it. But
 * some frames depend on (i.e. blend with) prior frames, and
 * AImageDecoder_decodeImage assumes that the prior frame is in the
 * |pixels| buffer. In addition, AImageDecoder_decodeImage handles caching and
 * restoring frames (see {@link ANDROID_IMAGE_DECODER_DISPOSE_OP_PREVIOUS}), so
 * skipping frames in an image with such frames may not produce the correct
 * results.
 *
 * Only supported by {@link ANDROID_BITMAP_FORMAT_RGBA_8888} and
 * {@link ANDROID_BITMAP_FORMAT_RGBA_F16}.
 *
 * @param decoder an {@link AImageDecoder} object.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The AImageDecoder
 *   represents an image that is not animated (see
 *   {@link AImageDecoder_isAnimated}) or the AImageDecoder is null.
 * - {@link ANDROID_IMAGE_DECODER_INVALID_STATE): The requested
 *   {@link AndroidBitmapFormat} does not support animation.
 * - {@link ANDROID_IMAGE_DECODER_INCOMPLETE}: The input appears
 *   to be truncated. The client must call {@link AImageDecoder_rewind}
 *   before calling {@link AImageDecoder_decodeImage} again.
 * - {@link ANDROID_IMAGE_DECODER_ERROR}: The input contains an error.
 *   The client must call  {@link AImageDecoder_rewind} before
 *   calling {@link AImageDecoder_decodeImage} again.
 * - {@link ANDROID_IMAGE_DECODER_FINISHED}: The input contains no
 *   more frames. The client must call {@link AImageDecoder_rewind}
 *   before calling {@link AImageDecoder_decodeImage} again.
 */
int AImageDecoder_advanceFrame(AImageDecoder* _Nonnull decoder)
        __INTRODUCED_IN(31);

/**
 * Return to the beginning of the animation.
 *
 * Introduced in API 31.
 *
 * After this call, the AImageDecoder will be ready to decode the
 * first frame of the animation. This can be called after reaching
 * the end of the animation or an error or in the middle of the
 * animation.
 *
 * @param decoder an {@link AImageDecoder} object.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: The AImageDecoder
 *   represents an image that is not animated (see
 *   {@link AImageDecoder_isAnimated}) or the AImageDecoder is
 *   null.
 * - {@link ANDROID_IMAGE_DECODER_SEEK_ERROR}: The asset or file
 *   descriptor failed to seek.
 */
int AImageDecoder_rewind(AImageDecoder* _Nonnull decoder)
        __INTRODUCED_IN(31);

struct AImageDecoderFrameInfo;

/**
 * Opaque handle to animation information about a single frame.
 *
 * Introduced in API 31
 *
 * The duration (retrieved with {@link AImageDecoderFrameInfo_getDuration}) is
 * necessary for clients to display the animation at the proper speed. The other
 * information is helpful for a client that wants to determine what frames are
 * independent (or what frames they depend on), but is unnecessary for
 * a simple client that wants to sequentially display all frames.
 */
typedef struct AImageDecoderFrameInfo AImageDecoderFrameInfo;

/**
 * Create an uninitialized AImageDecoderFrameInfo.
 *
 * Introduced in API 31.
 *
 * This can be passed to {@link AImageDecoder_getFrameInfo} to fill
 * in information about the current frame. It may be reused.
 *
 * Must be deleted with {@link AImageDecoderFrameInfo_delete}.
 */
AImageDecoderFrameInfo* _Nullable AImageDecoderFrameInfo_create()
        __INTRODUCED_IN(31);

/**
 * Delete an AImageDecoderFrameInfo.
 *
 * Introduced in API 31.
 */
void AImageDecoderFrameInfo_delete(
        AImageDecoderFrameInfo* _Nullable info) __INTRODUCED_IN(31);

/**
 * Fill |info| with information about the current frame.
 *
 * Introduced in API 31.
 *
 * Initially, this will return information about the first frame.
 * {@link AImageDecoder_advanceFrame} and
 * {@link AImageDecoder_rewind} can be used to change which frame
 * is the current frame.
 *
 * If the image only has one frame, this will fill the {@link
 * AImageDecoderFrameInfo} with the encoded info and reasonable
 * defaults.
 *
 * If {@link AImageDecoder_advanceFrame} succeeded, this will succeed as well.
 *
 * @param decoder Opaque object representing the decoder.
 * @param info Opaque object to hold frame information. On success, will be
 *             filled with information regarding the current frame.
 * @return {@link ANDROID_IMAGE_DECODER_SUCCESS} on success or a value
 *         indicating the reason for the failure.
 *
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER}: One of the parameters is null.
 * - {@link ANDROID_IMAGE_DECODER_FINISHED}: The input contains no
 *   more frames. The client must call {@link AImageDecoder_rewind} to reset the
 *   current frame to a valid frame (0).
 */
int AImageDecoder_getFrameInfo(AImageDecoder* _Nonnull decoder,
        AImageDecoderFrameInfo* _Nonnull info) __INTRODUCED_IN(31);

/**
 * Report the number of nanoseconds to show the current frame.
 *
 * Introduced in API 31.
 *
 * Errors:
 * - returns {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER} if |info| is null.
 */
int64_t AImageDecoderFrameInfo_getDuration(
        const AImageDecoderFrameInfo* _Nonnull info) __INTRODUCED_IN(31);

/**
 * The rectangle of the image (within 0, 0,
 * {@link AImageDecoderHeaderInfo_getWidth}, {@link AImageDecoderHeaderInfo_getHeight})
 * updated by this frame.
 *
 * Introduced in API 31.
 *
 * Note that this is unaffected by calls to
 * {@link AImageDecoder_setTargetSize} or
 * {@link AImageDecoder_setCrop}.
 *
 * A frame may update only part of the image. This will always be
 * contained by the image’s dimensions.
 *
 * This, along with other information in AImageDecoderFrameInfo,
 * can be useful for determining whether a frame is independent, but
 * the decoder handles blending frames, so a simple
 * sequential client does not need this.
 *
 * Errors:
 * - returns an empty ARect if |info| is null.
 */
ARect AImageDecoderFrameInfo_getFrameRect(
        const AImageDecoderFrameInfo* _Nonnull info) __INTRODUCED_IN(31);

/**
 * Whether the new portion of this frame may contain alpha.
 *
 * Introduced in API 31.
 *
 * Unless this frame is independent (see {@link AImageDecoder_decodeImage}),
 * a single call to {@link AImageDecoder_decodeImage} will decode an updated
 * rectangle of pixels and then blend it with the existing pixels in the
 * |pixels| buffer according to {@link AImageDecoderFrameInfo_getBlendOp}. This
 * method returns whether the updated rectangle has alpha, prior to blending.
 * The return value is conservative; for example, if a color-index-based frame
 * has a color with alpha but does not use it, this will still return true.
 *
 * This, along with other information in AImageDecoderFrameInfo,
 * can be useful for determining whether a frame is independent, but
 * the decoder handles blending frames, so a simple
 * sequential client does not need this.
 *
 * Note that this may differ from whether the composed frame (that is, the
 * resulting image after blending) has alpha. If this frame does not fill the
 * entire image dimensions (see {@link AImageDecoderFrameInfo_getFrameRect})
 * or it blends with an opaque frame, for example, the composed frame’s alpha
 * may not match.
 *
 * Errors:
 * - returns false if |info| is null.
 */
bool AImageDecoderFrameInfo_hasAlphaWithinBounds(
        const AImageDecoderFrameInfo* _Nonnull info) __INTRODUCED_IN(31);

/**
 * How a frame is “disposed” before showing the next one.
 *
 * Introduced in API 31.
 *
 * This, along with other information in AImageDecoderFrameInfo,
 * can be useful for determining whether a frame is independent, but
 * the decoder handles disposing of frames, so a simple
 * sequential client does not need this.
 */
enum {
    /// No disposal. The following frame will be drawn directly
    /// on top of this one.
    ANDROID_IMAGE_DECODER_DISPOSE_OP_NONE = 1,
    /// The frame’s rectangle is cleared to transparent (by AImageDecoder)
    /// before decoding the next frame.
    ANDROID_IMAGE_DECODER_DISPOSE_OP_BACKGROUND = 2,
    /// The frame’s rectangle is reverted to the prior frame before decoding
    /// the next frame. This is handled by AImageDecoder, unless
    /// {@link AImageDecoder_setInternallyHandleDisposePrevious} is set to false.
    ANDROID_IMAGE_DECODER_DISPOSE_OP_PREVIOUS = 3,
};

/**
 * Return how this frame is “disposed” before showing the next one.
 *
 * Introduced in API 31.
 *
 * This, along with other information in AImageDecoderFrameInfo,
 * can be useful for determining whether a frame is independent, but
 * the decoder handles disposing of frames, so a simple
 * sequential client does not need this.
 *
 * @return one of:
 * - {@link ANDROID_IMAGE_DECODER_DISPOSE_OP_NONE}
 * - {@link ANDROID_IMAGE_DECODER_DISPOSE_OP_BACKGROUND}
 * - {@link ANDROID_IMAGE_DECODER_DISPOSE_OP_PREVIOUS}
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER} if |info| is null.
 */
int32_t AImageDecoderFrameInfo_getDisposeOp(
        const AImageDecoderFrameInfo* _Nonnull info) __INTRODUCED_IN(31);

/**
 * How a frame is blended with the previous frame.
 *
 * Introduced in API 31.
 *
 * This, along with other information in AImageDecoderFrameInfo,
 * can be useful for determining whether a frame is independent, but
 * the decoder handles blending frames, so a simple
 * sequential client does not need this.
 */
enum {
    /// This frame replaces existing content. This corresponds
    /// to webp’s “do not blend”.
    ANDROID_IMAGE_DECODER_BLEND_OP_SRC = 1,
    /// This frame blends with the previous frame.
    ANDROID_IMAGE_DECODER_BLEND_OP_SRC_OVER = 2,
};

/**
 * Return how this frame is blended with the previous frame.
 *
 * Introduced in API 31.
 *
 * This, along with other information in AImageDecoderFrameInfo,
 * can be useful for determining whether a frame is independent, but
 * the decoder handles blending frames, so a simple
 * sequential client does not need this.
 *
 * @return one of:
 * - {@link ANDROID_IMAGE_DECODER_BLEND_OP_SRC}
 * - {@link ANDROID_IMAGE_DECODER_BLEND_OP_SRC_OVER}
 * Errors:
 * - {@link ANDROID_IMAGE_DECODER_BAD_PARAMETER} if |info| is null.
 */
int32_t AImageDecoderFrameInfo_getBlendOp(
        const AImageDecoderFrameInfo* _Nonnull info)
        __INTRODUCED_IN(31);

/**
 * Whether to have AImageDecoder store the frame prior to a
 * frame marked {@link ANDROID_IMAGE_DECODER_DISPOSE_OP_PREVIOUS}.
 *
 * Introduced in API 31.
 *
 * The default is true. Many images will not have such a frame (it
 * is not supported by WebP, and only some GIFs use it). But
 * if frame i is ANDROID_IMAGE_DECODER_DISPOSE_OP_PREVIOUS, then i+1
 * may depend on i-1. When this setting is true, AImageDecoder will
 * defensively copy frame i-1 (i.e. the contents of |pixels| in
 * {@link AImageDecoder_decodeImage}) into an internal buffer so that
 * it can be used to decode i+1.
 *
 * AImageDecoder will only store a single frame, at the size specified
 * by {@link AImageDecoder_setTargetSize} (or the original dimensions
 * if that method has not been called), and will discard it when it is
 * no longer necessary.
 *
 * A client that desires to manually store such frames may set this to
 * false, so that AImageDecoder does not need to store this extra
 * frame. Instead, when decoding the same
 * ANDROID_IMAGE_DECODER_DISPOSE_OP_PREVIOUS frame i, AImageDecoder
 * will decode directly into |pixels|, assuming the client stored i-1.
 * When asked to decode frame i+1, AImageDecoder will now assume that
 * the client provided i-1 in |pixels|.
 *
 * @param decoder an {@link AImageDecoder} object.
 * @param handleInternally Whether AImageDecoder will internally
 *               handle ANDROID_IMAGE_DECODER_DISPOSE_OP_PREVIOUS
 *               frames.
 */
void AImageDecoder_setInternallyHandleDisposePrevious(
        AImageDecoder* _Nonnull decoder, bool handleInternally)
        __INTRODUCED_IN(31);


#ifdef __cplusplus
}
#endif

#endif // ANDROID_IMAGE_DECODER_H

/** @} */
