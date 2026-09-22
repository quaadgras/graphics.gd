/*
 * Copyright 2022 The Android Open Source Project
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

#include "FrameStatistics.h"

#include "SwappyCommon.h"

#define LOG_TAG "FrameStatistics"
#include "SwappyLog.h"

namespace swappy {

// NB This is only needed for C++14
constexpr std::chrono::nanoseconds FrameStatistics::LOG_EVERY_N_NS;

int32_t FrameStatistics::getFrameDelta(int64_t deltaTimeNS,
                                       uint64_t refreshPeriod) {
    int32_t numFrames = deltaTimeNS / refreshPeriod;
    numFrames = std::max(
        0, std::min(numFrames, static_cast<int32_t>(MAX_FRAME_BUCKETS) - 1));
    return numFrames;
}

void FrameStatistics::clearStats() {
    std::lock_guard<std::mutex> lock(mMutex);
    mStats.totalFrames = 0;

    for (int i = 0; i < MAX_FRAME_BUCKETS; i++) {
        mStats.idleFrames[i] = 0;
        mStats.lateFrames[i] = 0;
        mStats.offsetFromPreviousFrame[i] = 0;
        mStats.latencyFrames[i] = 0;
    }
}

void FrameStatistics::invalidateLastFrame() { mLast = {0, 0, 0, 0}; }

void FrameStatistics::updateFrameStats(FrameTimings current,
                                       uint64_t refreshPeriod) {
    std::lock_guard<std::mutex> lock(mMutex);
    // Latency is always collected
    int latency = getFrameDelta(
        current.actualPresentTime - current.startFrameTime, refreshPeriod);

    // Use incoming frame timings to build the histogram.
    if (mFullStatsEnabled) {
        int idle = getFrameDelta(current.presentMargin, refreshPeriod);
        int late = getFrameDelta(
            current.actualPresentTime - current.desiredPresentTime,
            refreshPeriod);

        mStats.totalFrames++;
        mStats.idleFrames[idle]++;
        mStats.lateFrames[late]++;
        mStats.latencyFrames[latency]++;

        // Update the previous frame only if last frame has valid data
        if (mLast.actualPresentTime) {
            int offset = getFrameDelta(
                current.actualPresentTime - mLast.actualPresentTime,
                refreshPeriod);

            mStats.offsetFromPreviousFrame[offset]++;
        }

        logFrames();
    }

    mLastLatency = latency;
    mLast = current;
}
void FrameStatistics::logFrames() {
    static auto previousLogTime = std::chrono::steady_clock::now();

    if (std::chrono::steady_clock::now() - previousLogTime < LOG_EVERY_N_NS) {
        return;
    }

    std::string message;
    SWAPPY_LOGI("== Frame statistics ==");
    SWAPPY_LOGI("total frames: %" PRIu64, mStats.totalFrames);
    message += "Buckets:                    ";
    for (int i = 0; i < MAX_FRAME_BUCKETS; i++)
        message += "\t[" + swappy::to_string(i) + "]";
    SWAPPY_LOGI("%s", message.c_str());

    message = "";
    message += "idle frames:                ";
    for (int i = 0; i < MAX_FRAME_BUCKETS; i++)
        message += "\t " + swappy::to_string(mStats.idleFrames[i]);
    SWAPPY_LOGI("%s", message.c_str());

    message = "";
    message += "late frames:                ";
    for (int i = 0; i < MAX_FRAME_BUCKETS; i++)
        message += "\t " + swappy::to_string(mStats.lateFrames[i]);
    SWAPPY_LOGI("%s", message.c_str());

    message = "";
    message += "offset from previous frame: ";
    for (int i = 0; i < MAX_FRAME_BUCKETS; i++)
        message += "\t " + swappy::to_string(mStats.offsetFromPreviousFrame[i]);
    SWAPPY_LOGI("%s", message.c_str());

    message = "";
    message += "frame latency:              ";
    for (int i = 0; i < MAX_FRAME_BUCKETS; i++)
        message += "\t " + swappy::to_string(mStats.latencyFrames[i]);
    SWAPPY_LOGI("%s", message.c_str());

    previousLogTime = std::chrono::steady_clock::now();
}

void FrameStatistics::enableStats(bool enabled) { mFullStatsEnabled = enabled; }

SwappyStats FrameStatistics::getStats() {
    std::lock_guard<std::mutex> lock(mMutex);
    return mStats;
}

}  // namespace swappy
