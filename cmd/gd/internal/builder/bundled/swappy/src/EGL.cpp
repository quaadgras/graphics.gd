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

#include "EGL.h"

#include <Trace.h>
#include <dlfcn.h>

#include <vector>

#define LOG_TAG "Swappy::EGL"

#include "SwappyLog.h"

using namespace std::chrono_literals;

namespace swappy {

std::unique_ptr<EGL> EGL::create(std::chrono::nanoseconds fenceTimeout) {
    auto eglLib = dlopen("libEGL.so", RTLD_LAZY | RTLD_LOCAL);
    if (eglLib == nullptr) {
        SWAPPY_LOGE("Can't load libEGL");
        return nullptr;
    }
    auto eglGetProcAddress = reinterpret_cast<eglGetProcAddress_type>(
        dlsym(eglLib, "eglGetProcAddress"));
    if (eglGetProcAddress == nullptr) {
        SWAPPY_LOGE("Failed to load eglGetProcAddress");
        return nullptr;
    }

    auto eglSwapBuffers =
        reinterpret_cast<eglSwapBuffers_type>(dlsym(eglLib, "eglSwapBuffers"));
    if (eglSwapBuffers == nullptr) {
        SWAPPY_LOGE("Failed to load eglSwapBuffers");
        return nullptr;
    }

    auto eglPresentationTimeANDROID =
        reinterpret_cast<eglPresentationTimeANDROID_type>(
            eglGetProcAddress("eglPresentationTimeANDROID"));
    if (eglPresentationTimeANDROID == nullptr) {
        SWAPPY_LOGE("Failed to load eglPresentationTimeANDROID");
        return nullptr;
    }

    auto eglCreateSyncKHR = reinterpret_cast<eglCreateSyncKHR_type>(
        eglGetProcAddress("eglCreateSyncKHR"));
    if (eglCreateSyncKHR == nullptr) {
        SWAPPY_LOGE("Failed to load eglCreateSyncKHR");
        return nullptr;
    }

    auto eglDestroySyncKHR = reinterpret_cast<eglDestroySyncKHR_type>(
        eglGetProcAddress("eglDestroySyncKHR"));
    if (eglDestroySyncKHR == nullptr) {
        SWAPPY_LOGE("Failed to load eglDestroySyncKHR");
        return nullptr;
    }

    auto eglGetSyncAttribKHR = reinterpret_cast<eglGetSyncAttribKHR_type>(
        eglGetProcAddress("eglGetSyncAttribKHR"));
    if (eglGetSyncAttribKHR == nullptr) {
        SWAPPY_LOGE("Failed to load eglGetSyncAttribKHR");
        return nullptr;
    }

    auto eglClientWaitSyncKHR = reinterpret_cast<eglClientWaitSyncKHR_type>(
        eglGetProcAddress("eglClientWaitSyncKHR"));
    if (eglClientWaitSyncKHR == nullptr) {
        SWAPPY_LOGE("Failed to load eglClientWaitSyncKHR");
        return nullptr;
    }

    auto eglGetError =
        reinterpret_cast<eglGetError_type>(eglGetProcAddress("eglGetError"));
    if (eglGetError == nullptr) {
        SWAPPY_LOGE("Failed to load eglGetError");
        return nullptr;
    }

    auto eglSurfaceAttrib = reinterpret_cast<eglSurfaceAttrib_type>(
        eglGetProcAddress("eglSurfaceAttrib"));
    if (eglSurfaceAttrib == nullptr) {
        SWAPPY_LOGE("Failed to load eglSurfaceAttrib");
        return nullptr;
    }

    // stats may not be supported on all versions
    auto eglGetNextFrameIdANDROID =
        reinterpret_cast<eglGetNextFrameIdANDROID_type>(
            eglGetProcAddress("eglGetNextFrameIdANDROID"));
    if (eglGetNextFrameIdANDROID == nullptr) {
        SWAPPY_LOGI("Failed to load eglGetNextFrameIdANDROID");
    }

    auto eglGetFrameTimestampsANDROID =
        reinterpret_cast<eglGetFrameTimestampsANDROID_type>(
            eglGetProcAddress("eglGetFrameTimestampsANDROID"));
    if (eglGetFrameTimestampsANDROID == nullptr) {
        SWAPPY_LOGI("Failed to load eglGetFrameTimestampsANDROID");
    }

    auto egl = std::make_unique<EGL>(fenceTimeout, eglGetProcAddress,
                                     ConstructorTag{});
    egl->eglLib = eglLib;
    egl->eglSwapBuffers = eglSwapBuffers;
    egl->eglGetProcAddress = eglGetProcAddress;
    egl->eglPresentationTimeANDROID = eglPresentationTimeANDROID;
    egl->eglCreateSyncKHR = eglCreateSyncKHR;
    egl->eglClientWaitSyncKHR = eglClientWaitSyncKHR;
    egl->eglDestroySyncKHR = eglDestroySyncKHR;
    egl->eglGetSyncAttribKHR = eglGetSyncAttribKHR;
    egl->eglGetError = eglGetError;
    egl->eglSurfaceAttrib = eglSurfaceAttrib;
    egl->eglGetNextFrameIdANDROID = eglGetNextFrameIdANDROID;
    egl->eglGetFrameTimestampsANDROID = eglGetFrameTimestampsANDROID;

    std::lock_guard<std::mutex> lock(egl->mWaiterThreadContext.lock);
    egl->mWaiterThreadContext.thread =
        Thread([egl = egl.get()]() { egl->waitForFenceThreadMain(); });

    return egl;
}

EGL::~EGL() {
    // Stop the fence waiter thread
    {
        std::lock_guard<std::mutex> lock(mWaiterThreadContext.lock);
        mWaiterThreadContext.running = false;
        mWaiterThreadContext.condition.notify_one();
    }

    mWaiterThreadContext.thread.join();

    while (mWaitPendingSyncs.size() > 0) {
        auto sync = mWaitPendingSyncs.front();
        mWaitPendingSyncs.pop_front();
        // There is no need to wait here as the API allows for queueing pending
        // sync for deleting.
        EGLBoolean result = eglDestroySyncKHR(sync.display, sync.fence);
        if (result == EGL_FALSE) {
            SWAPPY_LOGE("Failed to destroy sync fence");
        }
    }
    if (eglLib) {
        dlclose(eglLib);
    }
}

bool EGL::setPresentationTime(EGLDisplay display, EGLSurface surface,
                              std::chrono::steady_clock::time_point time) {
    eglPresentationTimeANDROID(display, surface,
                               time.time_since_epoch().count());
    return EGL_TRUE;
}

bool EGL::statsSupported() {
    return (eglGetNextFrameIdANDROID != nullptr &&
            eglGetFrameTimestampsANDROID != nullptr);
}

std::pair<bool, EGLuint64KHR> EGL::getNextFrameId(EGLDisplay dpy,
                                                  EGLSurface surface) const {
    if (eglGetNextFrameIdANDROID == nullptr) {
        SWAPPY_LOGE("stats are not supported on this platform");
        return {false, 0};
    }

    EGLuint64KHR frameId;
    EGLBoolean result = eglGetNextFrameIdANDROID(dpy, surface, &frameId);
    if (result == EGL_FALSE) {
        SWAPPY_LOGE("Failed to get next frame ID");
        return {false, 0};
    }

    return {true, frameId};
}

std::unique_ptr<EGL::FrameTimestamps> EGL::getFrameTimestamps(
    EGLDisplay dpy, EGLSurface surface, EGLuint64KHR frameId) const {
#if (not defined ANDROID_NDK_VERSION) || ANDROID_NDK_VERSION >= 15
    if (eglGetFrameTimestampsANDROID == nullptr) {
        SWAPPY_LOGE("stats are not supported on this platform");
        return nullptr;
    }
    const std::vector<EGLint> timestamps = {
        EGL_REQUESTED_PRESENT_TIME_ANDROID,
        EGL_RENDERING_COMPLETE_TIME_ANDROID,
        EGL_COMPOSITION_LATCH_TIME_ANDROID,
        EGL_DISPLAY_PRESENT_TIME_ANDROID,
    };

    std::vector<EGLnsecsANDROID> values(timestamps.size());

    EGLBoolean result =
        eglGetFrameTimestampsANDROID(dpy, surface, frameId, timestamps.size(),
                                     timestamps.data(), values.data());
    if (result == EGL_FALSE) {
        EGLint reason = eglGetError();
        if (reason == EGL_BAD_SURFACE) {
            eglSurfaceAttrib(dpy, surface, EGL_TIMESTAMPS_ANDROID, EGL_TRUE);
        } else {
            SWAPPY_LOGE_ONCE("Failed to get timestamps for frame %llu",
                             (unsigned long long)frameId);
        }
        return nullptr;
    }

    // try again if we got some pending stats
    for (auto i : values) {
        if (i == EGL_TIMESTAMP_PENDING_ANDROID) return nullptr;
    }

    std::unique_ptr<EGL::FrameTimestamps> frameTimestamps =
        std::make_unique<EGL::FrameTimestamps>();
    frameTimestamps->requested = values[0];
    frameTimestamps->renderingCompleted = values[1];
    frameTimestamps->compositionLatched = values[2];
    frameTimestamps->presented = values[3];

    return frameTimestamps;
#else
    return nullptr;
#endif
}

void EGL::insertSyncFence(EGLDisplay display) {
    EGLSyncKHR sync_fence =
        eglCreateSyncKHR(display, EGL_SYNC_FENCE_KHR, nullptr);

    if (sync_fence != EGL_NO_SYNC_KHR) {
        EGLSync sync = {display, sync_fence};
        // kick off the thread work to wait for the fence and measure its time
        std::lock_guard<std::mutex> lock(mWaiterThreadContext.lock);
        mWaitPendingSyncs.push_back(sync);
        mWaiterThreadContext.hasPendingWork = true;
        mWaiterThreadContext.condition.notify_all();
    } else {
        SWAPPY_LOGE("Failed to create sync fence");
    }
}

bool EGL::lastFrameIsComplete(EGLDisplay display, bool pipelineMode) {
    std::lock_guard<std::mutex> lock(mWaiterThreadContext.lock);
    if (pipelineMode) {
        // We are in pipeline mode so we need to check the fence of frame N-1
        return mWaitPendingSyncs.size() < 2;
    }
    // We are not in pipeline mode so we need to check the fence of the current
    // frame. i.e. there are not unsignaled frames
    return mWaitPendingSyncs.empty();
}

void EGL::waitForFenceThreadMain() {
    while (true) {
        bool waitingSyncsEmpty;
        {
            std::lock_guard<std::mutex> lock(mWaiterThreadContext.lock);

            mWaiterThreadContext.condition.wait(
                mWaiterThreadContext.lock,
                [&]() REQUIRES(mWaiterThreadContext.lock) {
                    return mWaiterThreadContext.hasPendingWork ||
                           !mWaiterThreadContext.running;
                });

            mWaiterThreadContext.hasPendingWork = false;

            if (!mWaiterThreadContext.running) {
                break;
            }

            waitingSyncsEmpty = mWaitPendingSyncs.empty();
        }

        // No other consumers can empty the syncs while this thread is running,
        // the destructor of EGL waits for this thread to finish before emptying
        // the pending syncs.
        while (!waitingSyncsEmpty) {
            EGLSync sync;
            {
                // Get the latest fence to wait on.
                std::lock_guard<std::mutex> lock(mWaiterThreadContext.lock);
                sync = mWaitPendingSyncs.front();
            }

            gamesdk::ScopedTrace tracer("Swappy: GPU frame time");
            const auto startTime = std::chrono::steady_clock::now();

            EGLBoolean result = eglClientWaitSyncKHR(sync.display, sync.fence,
                                                     0, mFenceTimeout.count());
            switch (result) {
                case EGL_FALSE:
                    SWAPPY_LOGE("Failed to wait sync");
                    break;
                case EGL_TIMEOUT_EXPIRED_KHR:
                    SWAPPY_LOGE("Timeout waiting for fence");
                    break;
            }

            mFencePendingTime = std::chrono::steady_clock::now() - startTime;

            {
                std::lock_guard<std::mutex> lock(mWaiterThreadContext.lock);
                mWaitPendingSyncs.pop_front();

                // Once the wait has timed out/succeeded, we can submit it for
                // deletion as the API allows for pending syncs to be queued for
                // deletion.
                result = eglDestroySyncKHR(sync.display, sync.fence);
                if (result == EGL_FALSE) {
                    SWAPPY_LOGE("Failed to destroy sync fence");
                }
                waitingSyncsEmpty = mWaitPendingSyncs.empty();
            }
        }
    }
}

}  // namespace swappy
