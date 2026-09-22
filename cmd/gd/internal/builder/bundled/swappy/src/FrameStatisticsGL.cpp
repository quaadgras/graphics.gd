/*
 * Copyright 2019 The Android Open Source Project
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

#include "FrameStatisticsGL.h"

#include <inttypes.h>

#include <cmath>
#include <string>

#include "EGL.h"
#include "SwappyCommon.h"
#include "Trace.h"

namespace swappy {

FrameStatisticsGL::FrameStatisticsGL(const EGL& egl,
                                     const SwappyCommon& swappyCommon)
    : mEgl(egl), mSwappyCommon(swappyCommon) {
    mPendingFrames.reserve(MAX_FRAME_LAG + 1);
}

FrameStatisticsGL::ThisFrame FrameStatisticsGL::getThisFrame(
    EGLDisplay dpy, EGLSurface surface) {
    const TimePoint frameStartTime = std::chrono::steady_clock::now();

    // first get the next frame id
    std::pair<bool, EGLuint64KHR> nextFrameId =
        mEgl.getNextFrameId(dpy, surface);
    if (nextFrameId.first) {
        mPendingFrames.push_back(
            {dpy, surface, nextFrameId.second, frameStartTime});
    }

    if (mPendingFrames.empty()) {
        return {};
    }

    EGLFrame frame = mPendingFrames.front();
    // make sure we don't lag behind the stats too much
    if (nextFrameId.first && nextFrameId.second - frame.id > MAX_FRAME_LAG) {
        while (mPendingFrames.size() > 1)
            mPendingFrames.erase(mPendingFrames.begin());
        mFrameStatsCommon.invalidateLastFrame();
        frame = mPendingFrames.front();
    }
    std::unique_ptr<EGL::FrameTimestamps> frameStats =
        mEgl.getFrameTimestamps(frame.dpy, frame.surface, frame.id);

    if (!frameStats) {
        return {frame.startFrameTime};
    }

    mPendingFrames.erase(mPendingFrames.begin());

    return {frame.startFrameTime, std::move(frameStats)};
}

// called once per swap
void FrameStatisticsGL::capture(EGLDisplay dpy, EGLSurface surface) {
    auto frame = getThisFrame(dpy, surface);

    if (!frame.stats) return;

    FrameTimings current = {
        static_cast<uint64_t>(frame.startTime.time_since_epoch().count()),
        static_cast<uint64_t>(frame.stats->requested),
        static_cast<uint64_t>(frame.stats->presented),
        static_cast<uint64_t>(frame.stats->compositionLatched -
                              frame.stats->renderingCompleted)};

    mFrameStatsCommon.updateFrameStats(
        current, mSwappyCommon.getRefreshPeriod().count());
}

void FrameStatisticsGL::enableStats(bool enabled) {
    mFrameStatsCommon.enableStats(enabled);
}

SwappyStats FrameStatisticsGL::getStats() {
    return mFrameStatsCommon.getStats();
}

void FrameStatisticsGL::clearStats() { mFrameStatsCommon.clearStats(); }

int32_t FrameStatisticsGL::lastLatencyRecorded() {
    return mFrameStatsCommon.lastLatencyRecorded();
}
}  // namespace swappy
