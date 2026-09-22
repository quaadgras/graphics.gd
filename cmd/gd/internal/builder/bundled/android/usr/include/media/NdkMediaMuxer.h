/*
 * Copyright (C) 2014 The Android Open Source Project
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
 * @file NdkMediaMuxer.h
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

#ifndef _NDK_MEDIA_MUXER_H
#define _NDK_MEDIA_MUXER_H

#include <sys/cdefs.h>
#include <sys/types.h>

#include "NdkMediaCodec.h"
#include "NdkMediaError.h"
#include "NdkMediaFormat.h"

__BEGIN_DECLS

struct AMediaMuxer;
typedef struct AMediaMuxer AMediaMuxer;

/**
 * Defines the output format. These constants are used with constructor.
 *
 * These enums match the ones used in {@link android.media.MediaMuxer.OutputFormat}
 */
typedef enum {
    /** MPEG4 media file format*/
    AMEDIAMUXER_OUTPUT_FORMAT_MPEG_4 = 0,
    /** WEBM media file format*/
    AMEDIAMUXER_OUTPUT_FORMAT_WEBM = 1,
    /** 3GPP media file format*/
    AMEDIAMUXER_OUTPUT_FORMAT_THREE_GPP = 2,
    /** HEIF media file format*/
    AMEDIAMUXER_OUTPUT_FORMAT_HEIF = 3,  // introduced in API 34
    /** Ogg media file format*/
    AMEDIAMUXER_OUTPUT_FORMAT_OGG = 4,  // introduced in API 34
} OutputFormat;

typedef enum {
    /* Last group of pictures(GOP) of video track can be incomplete, so it would be safe to
     * scrap that and rewrite.  If both audio and video tracks are present in a file, then
     * samples of audio track after last GOP of video would be scrapped too.
     * If only audio track is present, then no sample would be discarded.
     */
    AMEDIAMUXER_APPEND_IGNORE_LAST_VIDEO_GOP = 0,
    // Keep all existing samples as it is and append new samples after that only.
    AMEDIAMUXER_APPEND_TO_EXISTING_DATA = 1,
} AppendMode;

/**
 * Create new media muxer.
 *
 * Available since API level 21.
 */
AMediaMuxer* AMediaMuxer_new(int fd, OutputFormat format) __INTRODUCED_IN(21);

/**
 * Delete a previously created media muxer.
 *
 * Available since API level 21.
 */
media_status_t AMediaMuxer_delete(AMediaMuxer*) __INTRODUCED_IN(21);

/**
 * Set and store the geodata (latitude and longitude) in the output file.
 * This method should be called before AMediaMuxer_start. The geodata is stored
 * in udta box if the output format is AMEDIAMUXER_OUTPUT_FORMAT_MPEG_4, and is
 * ignored for other output formats.
 * The geodata is stored according to ISO-6709 standard.
 *
 * Both values are specified in degrees.
 * Latitude must be in the range [-90, 90].
 * Longitude must be in the range [-180, 180].
 *
 * Available since API level 21.
 */
media_status_t AMediaMuxer_setLocation(AMediaMuxer*,
        float latitude, float longitude) __INTRODUCED_IN(21);

/**
 * Sets the orientation hint for output video playback.
 * This method should be called before AMediaMuxer_start. Calling this
 * method will not rotate the video frame when muxer is generating the file,
 * but add a composition matrix containing the rotation angle in the output
 * video if the output format is AMEDIAMUXER_OUTPUT_FORMAT_MPEG_4, so that a
 * video player can choose the proper orientation for playback.
 * Note that some video players may choose to ignore the composition matrix
 * during playback.
 * The angle is specified in degrees, clockwise.
 * The supported angles are 0, 90, 180, and 270 degrees.
 *
 * Available since API level 21.
 */
media_status_t AMediaMuxer_setOrientationHint(AMediaMuxer*, int degrees) __INTRODUCED_IN(21);

/**
 * Adds a track with the specified format.
 * Returns the index of the new track or a negative value in case of failure,
 * which can be interpreted as a media_status_t.
 *
 * Available since API level 21.
 */
ssize_t AMediaMuxer_addTrack(AMediaMuxer*, const AMediaFormat* format) __INTRODUCED_IN(21);

/**
 * Start the muxer. Should be called after AMediaMuxer_addTrack and
 * before AMediaMuxer_writeSampleData.
 *
 * Available since API level 21.
 */
media_status_t AMediaMuxer_start(AMediaMuxer*) __INTRODUCED_IN(21);

/**
 * Stops the muxer.
 * Once the muxer stops, it can not be restarted.
 *
 * Available since API level 21.
 */
media_status_t AMediaMuxer_stop(AMediaMuxer*) __INTRODUCED_IN(21);

/**
 * Writes an encoded sample into the muxer.
 * The application needs to make sure that the samples are written into
 * the right tracks. Also, it needs to make sure the samples for each track
 * are written in chronological order (e.g. in the order they are provided
 * by the encoder.)
 *
 * Available since API level 21.
 */
media_status_t AMediaMuxer_writeSampleData(AMediaMuxer *muxer,
        size_t trackIdx, const uint8_t *data,
        const AMediaCodecBufferInfo *info) __INTRODUCED_IN(21);

/**
 * Creates a new media muxer for appending data to an existing MPEG4 file.
 * This is a synchronous API call and could take a while to return if the existing file is large.
 * Only works for MPEG4 files matching one of the following characteristics:
 * <ul>
 *    <li>a single audio track.</li>
 *    <li>a single video track.</li>
 *    <li>a single audio and a single video track.</li>
 * </ul>
 * @param fd Must be opened with read and write permission. Does not take ownership of
 * this fd i.e., caller is responsible for closing fd.
 * @param mode Specifies how data will be appended; the AppendMode enum describes
 *             the possible methods for appending..
 * @return Pointer to AMediaMuxer if the file(fd) has tracks already, otherwise, nullptr.
 * {@link AMediaMuxer_delete} should be used to free the returned pointer.
 *
 * Available since API level 31.
 */
AMediaMuxer* AMediaMuxer_append(int fd, AppendMode mode) __INTRODUCED_IN(31);

/**
 * Returns the number of tracks added in the file passed to {@link AMediaMuxer_new} or
 * the number of existing tracks in the file passed to {@link AMediaMuxer_append}.
 * Should be called in INITIALIZED or STARTED state, otherwise returns -1.
 *
 * Available since API level 31.
 */
ssize_t AMediaMuxer_getTrackCount(AMediaMuxer*) __INTRODUCED_IN(31);

/**
 * Returns AMediaFormat of the added track with index idx in the file passed to
 * {@link AMediaMuxer_new} or the AMediaFormat of the existing track with index idx
 * in the file passed to {@link AMediaMuxer_append}.
 * Should be called in INITIALIZED or STARTED state, otherwise returns nullptr.
 * {@link AMediaFormat_delete} should be used to free the returned pointer.
 *
 * Available since API level 31.
 */
AMediaFormat* AMediaMuxer_getTrackFormat(AMediaMuxer* muxer, size_t idx) __INTRODUCED_IN(31);

__END_DECLS

#endif // _NDK_MEDIA_MUXER_H

/** @} */
