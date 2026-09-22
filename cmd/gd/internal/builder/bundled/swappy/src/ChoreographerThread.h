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

#include <mutex>

#include "SwappyDisplayManager.h"
#include "Thread.h"

namespace swappy {

class ChoreographerThread {
 public:
  enum class Type {
    // choreographer ticks are provided by application
    App,

    // register internally with choreographer
    Swappy,
  };

  static const char* CT_CLASS;
  static const JNINativeMethod CTNativeMethods[];
  static constexpr int CTNativeMethodsSize = 1;

  using RefreshRateChangedCallback = std::function<void()>;
  using ChoreographerCallback = std::function<void(
      std::optional<std::chrono::nanoseconds> sfToVsyncDelay)>;

  static std::unique_ptr<ChoreographerThread> createChoreographerThread(
      Type type, JavaVM* vm, jobject jactivity,
      ChoreographerCallback onChoreographer,
      RefreshRateChangedCallback onRefreshRateChanged, SdkVersion sdkVersion);

  virtual ~ChoreographerThread() = 0;

  virtual void postFrameCallbacks();

  bool isInitialized() { return mInitialized; }

 protected:
  ChoreographerThread(ChoreographerCallback onChoreographer);
  virtual void scheduleNextFrameCallback() REQUIRES(mWaitingMutex) = 0;
  virtual void onChoreographer(
      std::optional<std::chrono::nanoseconds> sfToVsyncDelay);

  std::mutex mWaitingMutex;
  int mCallbacksBeforeIdle GUARDED_BY(mWaitingMutex) = 0;
  ChoreographerCallback mCallback;
  bool mInitialized = false;

  static constexpr int MAX_CALLBACKS_BEFORE_IDLE = 10;
};

}  // namespace swappy
