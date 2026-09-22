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

#pragma once

#include <swappy/swappy_common.h>

#include <atomic>
#include <chrono>
#include <mutex>

#include "Thread.h"

using namespace std::chrono_literals;

namespace swappy {

/* FrameTimings is defined so that EGL & vulkan can convert their timestamps
 * into this common struct. Within the common frame statistics, this is all the
 * information that is needed.
 */
typedef struct {
  uint64_t startFrameTime;
  uint64_t desiredPresentTime;
  uint64_t actualPresentTime;
  uint64_t presentMargin;
} FrameTimings;

class FrameStatistics {
 public:
  ~FrameStatistics() = default;

  void enableStats(bool enabled);
  void updateFrameStats(FrameTimings currentFrameTimings,
                        uint64_t refreshPeriod);
  SwappyStats getStats();
  void clearStats();
  void invalidateLastFrame();

  int32_t lastLatencyRecorded() { return mLastLatency; }

 private:
  static constexpr std::chrono::nanoseconds LOG_EVERY_N_NS = 1s;
  void logFrames() REQUIRES(mMutex);

  int32_t getFrameDelta(int64_t deltaTimeNS, uint64_t refreshPeriod);

  std::mutex mMutex;
  SwappyStats mStats GUARDED_BY(mMutex) = {};
  std::atomic<int32_t> mLastLatency = {0};
  FrameTimings mLast;

  // A flag to enable or disable frame stats histogram update.
  bool mFullStatsEnabled = false;
};

}  // namespace swappy
