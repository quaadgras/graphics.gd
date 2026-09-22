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

#include <array>
#include <atomic>
#include <map>
#include <vector>

#include "EGL.h"
#include "FrameStatistics.h"
#include "SwappyCommon.h"
#include "Thread.h"

using TimePoint = std::chrono::steady_clock::time_point;
using namespace std::chrono_literals;

namespace swappy {

class FrameStatisticsGL {
 public:
  FrameStatisticsGL(const EGL& egl, const SwappyCommon& swappyCommon);
  ~FrameStatisticsGL() = default;

  void enableStats(bool enabled);
  void capture(EGLDisplay dpy, EGLSurface surface);
  SwappyStats getStats();
  void clearStats();

  int32_t lastLatencyRecorded();

 protected:
  static constexpr int MAX_FRAME_LAG = 10;
  struct ThisFrame {
    TimePoint startTime;
    std::unique_ptr<EGL::FrameTimestamps> stats;
  };
  ThisFrame getThisFrame(EGLDisplay dpy, EGLSurface surface);

  const EGL& mEgl;
  const SwappyCommon& mSwappyCommon;

  struct EGLFrame {
    EGLDisplay dpy;
    EGLSurface surface;
    EGLuint64KHR id;
    TimePoint startFrameTime;
  };
  std::vector<EGLFrame> mPendingFrames;
  FrameStatistics mFrameStatsCommon;
};

}  // namespace swappy
