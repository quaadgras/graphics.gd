/*
 * Copyright 2018 The Android Open Source Project
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

#include <jni.h>

#include <chrono>
#include <mutex>

#include "EGL.h"
#include "FrameStatisticsGL.h"
#include "SwappyCommon.h"
#include "swappy/swappyGL.h"
#include "swappy/swappyGL_extra.h"

namespace swappy {

using EGLDisplay = void *;
using EGLSurface = void *;

using namespace std::chrono_literals;

class SwappyGL {
 private:
  // Allows construction with std::unique_ptr from a static method, but
  // disallows construction outside of the class since no one else can
  // construct a ConstructorTag
  struct ConstructorTag {};

 public:
  SwappyGL(JNIEnv *env, jobject jactivity, ConstructorTag);
  static bool init(JNIEnv *env, jobject jactivity);

  static bool setWindow(ANativeWindow *window);

  static void onChoreographer(int64_t frameTimeNanos);

  static bool swap(EGLDisplay display, EGLSurface surface);

  // Pass callbacks for tracing within the swap function
  static void addTracer(const SwappyTracer *tracer);

  static void removeTracer(const SwappyTracer *tracer);

  static std::chrono::nanoseconds getSwapDuration();

  static void setAutoSwapInterval(bool enabled);

  static void setAutoPipelineMode(bool enabled);

  static void setMaxAutoSwapDuration(std::chrono::nanoseconds maxDuration);

  static void enableStats(bool enabled);
  static void recordFrameStart(EGLDisplay display, EGLSurface surface);
  static void getStats(SwappyStats *stats);
  static void clearStats();

  static bool isEnabled();
  static void destroyInstance();

  static void setFenceTimeout(std::chrono::nanoseconds t);
  static std::chrono::nanoseconds getFenceTimeout();

  static void setBufferStuffingFixWait(int32_t n_frames);

  static int getSupportedRefreshPeriodsNS(uint64_t *out_refreshrates,
                                          int allocated_entries);

  static void resetFramePacing();

  static void enableFramePacing(bool enable);
  static void enableBlockingWait(bool enable);

 private:
  static SwappyGL *getInstance();

  bool enabled() const { return mEnableSwappy; }

  EGL *getEgl();

  bool swapInternal(EGLDisplay display, EGLSurface surface);

  bool lastFrameIsComplete(EGLDisplay display);

  // Computes the desired presentation time based on the swap interval and
  // sets it using eglPresentationTimeANDROID
  bool setPresentationTime(EGLDisplay display, EGLSurface surface);

  bool mEnableSwappy = true;

  static std::mutex sInstanceMutex;
  static std::unique_ptr<SwappyGL> sInstance GUARDED_BY(sInstanceMutex);

  std::mutex mEglMutex;
  std::unique_ptr<EGL> mEgl;

  std::unique_ptr<FrameStatisticsGL> mFrameStatistics;

  SwappyCommon mCommonBase;
};

}  // namespace swappy
