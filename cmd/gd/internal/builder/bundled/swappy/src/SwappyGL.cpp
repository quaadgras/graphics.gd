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

#include "SwappyGL.h"
#define LOG_TAG "SwappyGL"

#include <cinttypes>
#include <cmath>
#include <cstdlib>

#include "SwappyLog.h"
#include "Thread.h"
#include "Trace.h"
#include "system_utils.h"

namespace swappy {

using std::chrono::milliseconds;
using std::chrono::nanoseconds;

std::mutex SwappyGL::sInstanceMutex;
std::unique_ptr<SwappyGL> SwappyGL::sInstance;

bool SwappyGL::init(JNIEnv *env, jobject jactivity) {
    std::lock_guard<std::mutex> lock(sInstanceMutex);
    if (sInstance) {
        SWAPPY_LOGE("Attempted to initialize SwappyGL twice");
        return false;
    }
    sInstance = std::make_unique<SwappyGL>(env, jactivity, ConstructorTag{});
    if (!sInstance->mEnableSwappy) {
        SWAPPY_LOGE("Failed to initialize SwappyGL");
        return false;
    }

    return true;
}

void SwappyGL::onChoreographer(int64_t frameTimeNanos) {
    TRACE_CALL();

    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }

    swappy->mCommonBase.onChoreographer(frameTimeNanos);
}

bool SwappyGL::setWindow(ANativeWindow *window) {
    TRACE_CALL();

    SwappyGL *swappy = getInstance();
    if (!swappy) {
        SWAPPY_LOGE("Failed to get SwappyGL instance in setWindow");
        return false;
    }

    swappy->mCommonBase.setANativeWindow(window);
    return true;
}

bool SwappyGL::swap(EGLDisplay display, EGLSurface surface) {
    TRACE_CALL();

    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return EGL_FALSE;
    }

    if (swappy->enabled()) {
        return swappy->swapInternal(display, surface);
    } else {
        return swappy->getEgl()->swapBuffers(display, surface) == EGL_TRUE;
    }
}

bool SwappyGL::lastFrameIsComplete(EGLDisplay display) {
    bool pipelineMode = (mCommonBase.getCurrentPipelineMode() ==
                         SwappyCommon::PipelineMode::On);
    if (!getEgl()->lastFrameIsComplete(display, pipelineMode)) {
        gamesdk::ScopedTrace trace("lastFrameIncomplete");
        SWAPPY_LOGV("lastFrameIncomplete");
        return false;
    }
    return true;
}

bool SwappyGL::swapInternal(EGLDisplay display, EGLSurface surface) {
    const SwappyCommon::SwapHandlers handlers = {
        .lastFrameIsComplete = [&]() { return lastFrameIsComplete(display); },
        .getPrevFrameGpuTime =
            [&]() { return getEgl()->getFencePendingTime(); },
    };

    getEgl()->insertSyncFence(display);

    mCommonBase.onPreSwap(handlers);

    if (mCommonBase.needToSetPresentationTime()) {
        bool setPresentationTimeResult = setPresentationTime(display, surface);
        if (!setPresentationTimeResult) {
            return setPresentationTimeResult;
        }
    }

    bool swapBuffersResult =
        (getEgl()->swapBuffers(display, surface) == EGL_TRUE);

    mCommonBase.onPostSwap(handlers);

    return swapBuffersResult;
}

void SwappyGL::addTracer(const SwappyTracer *tracer) {
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    if (swappy->enabled() && tracer != nullptr)
        swappy->mCommonBase.addTracerCallbacks(*tracer);
}

void SwappyGL::removeTracer(const SwappyTracer *tracer) {
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    if (swappy->enabled() && tracer != nullptr)
        swappy->mCommonBase.removeTracerCallbacks(*tracer);
}

nanoseconds SwappyGL::getSwapDuration() {
    SwappyGL *swappy = getInstance();
    if (!swappy || !swappy->enabled()) {
        return -1ns;
    }
    return swappy->mCommonBase.getSwapDuration();
};

void SwappyGL::setAutoSwapInterval(bool enabled) {
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    if (swappy->enabled()) swappy->mCommonBase.setAutoSwapInterval(enabled);
}

void SwappyGL::setAutoPipelineMode(bool enabled) {
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    if (swappy->enabled()) swappy->mCommonBase.setAutoPipelineMode(enabled);
}

void SwappyGL::setMaxAutoSwapDuration(std::chrono::nanoseconds maxDuration) {
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    if (swappy->enabled())
        swappy->mCommonBase.setMaxAutoSwapDuration(maxDuration);
}

void SwappyGL::enableStats(bool enabled) {
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }

    if (swappy->mFrameStatistics) {
        swappy->mFrameStatistics->enableStats(enabled);
    }
}

void SwappyGL::recordFrameStart(EGLDisplay display, EGLSurface surface) {
    TRACE_CALL();
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }

    if (swappy->mFrameStatistics) {
        swappy->mFrameStatistics->capture(display, surface);
    }
}

void SwappyGL::getStats(SwappyStats *stats) {
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    if (swappy->mFrameStatistics) {
        *stats = swappy->mFrameStatistics->getStats();
    }
}

void SwappyGL::clearStats() {
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    if (swappy->mFrameStatistics) {
        swappy->mFrameStatistics->clearStats();
    }
}

SwappyGL *SwappyGL::getInstance() {
    std::lock_guard<std::mutex> lock(sInstanceMutex);
    return sInstance.get();
}

bool SwappyGL::isEnabled() {
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        // This is a case of error.
        // We do not log anything here, so that we do not spam
        // the user when this function is called each frame.
        return false;
    }
    return swappy->enabled();
}

void SwappyGL::destroyInstance() {
    std::lock_guard<std::mutex> lock(sInstanceMutex);
    sInstance.reset();
}

void SwappyGL::setFenceTimeout(std::chrono::nanoseconds t) {
    SwappyGL *swappy = getInstance();
    if (!swappy || !swappy->enabled()) {
        return;
    }
    swappy->mCommonBase.setFenceTimeout(t);
}

std::chrono::nanoseconds SwappyGL::getFenceTimeout() {
    SwappyGL *swappy = getInstance();
    if (!swappy || !swappy->enabled()) {
        return std::chrono::nanoseconds(0);
    }
    return swappy->mCommonBase.getFenceTimeout();
}

EGL *SwappyGL::getEgl() {
    static thread_local EGL *egl = nullptr;
    if (!egl) {
        std::lock_guard<std::mutex> lock(mEglMutex);
        egl = mEgl.get();
    }
    return egl;
}

SwappyGL::SwappyGL(JNIEnv *env, jobject jactivity, ConstructorTag)
    : mFrameStatistics(nullptr), mCommonBase(env, jactivity) {
    {
        std::lock_guard<std::mutex> lock(mEglMutex);
        mEgl = EGL::create(mCommonBase.getFenceTimeout());
        if (!mEgl) {
            SWAPPY_LOGE("Failed to load EGL functions");
            mEnableSwappy = false;
            return;
        }
    }

    if (!mCommonBase.isValid()) {
        SWAPPY_LOGE("SwappyCommon could not initialize correctly.");
        mEnableSwappy = false;
        return;
    }

    mEnableSwappy =
        !gamesdk::GetSystemPropAsBool(SWAPPY_SYSTEM_PROP_KEY_DISABLE, false);
    if (!enabled()) {
        SWAPPY_LOGI("Swappy is disabled");
        return;
    }

    if (mEgl->statsSupported()) {
        mFrameStatistics =
            std::make_unique<FrameStatisticsGL>(*mEgl, mCommonBase);
        mCommonBase.setLastLatencyRecordedCallback(
            [this]() { return this->mFrameStatistics->lastLatencyRecorded(); });
    } else {
        SWAPPY_LOGI("stats are not suppored on this platform");
    }
    SWAPPY_LOGI("SwappyGL initialized successfully");
}

bool SwappyGL::setPresentationTime(EGLDisplay display, EGLSurface surface) {
    TRACE_CALL();

    auto displayTimings = Settings::getInstance()->getDisplayTimings();

    // if we are too close to the vsync, there is no need to set presentation
    // time
    if ((mCommonBase.getPresentationTime() - std::chrono::steady_clock::now()) <
        (mCommonBase.getRefreshPeriod() - displayTimings.sfOffset)) {
        return EGL_TRUE;
    }
    return getEgl()->setPresentationTime(display, surface,
                                         mCommonBase.getPresentationTime());
}

void SwappyGL::setBufferStuffingFixWait(int32_t n_frames) {
    TRACE_CALL();
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    swappy->mCommonBase.setBufferStuffingFixWait(n_frames);
}

int SwappyGL::getSupportedRefreshPeriodsNS(uint64_t *out_refreshrates,
                                           int allocated_entries) {
    TRACE_CALL();
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return -1;
    }
    return swappy->mCommonBase.getSupportedRefreshPeriodsNS(out_refreshrates,
                                                            allocated_entries);
}

void SwappyGL::resetFramePacing() {
    TRACE_CALL();
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    swappy->mCommonBase.resetFramePacing();
}

void SwappyGL::enableFramePacing(bool enable) {
    TRACE_INT("enableFramePacing", (int)enable);
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    swappy->mCommonBase.enableFramePacing(enable);
}

void SwappyGL::enableBlockingWait(bool enable) {
    TRACE_INT("enableBlockingWait", (int)enable);
    SwappyGL *swappy = getInstance();
    if (!swappy) {
        return;
    }
    swappy->mCommonBase.enableBlockingWait(enable);
}

}  // namespace swappy
