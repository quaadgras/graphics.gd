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

#if (not defined ANDROID_NDK_VERSION) || ANDROID_NDK_VERSION >= 15

#include "SwappyVkGoogleDisplayTiming.h"

#define LOG_TAG "SwappyVkGoogleDisplayTiming"
#include "SwappyLog.h"

using std::chrono::nanoseconds;

namespace swappy {

SwappyVkGoogleDisplayTiming::SwappyVkGoogleDisplayTiming(
    JNIEnv* env, jobject jactivity, VkPhysicalDevice physicalDevice,
    VkDevice device, const SwappyVkFunctionProvider* provider)
    : SwappyVkBase(env, jactivity, physicalDevice, device, provider) {
    mPendingFrames.reserve(MAX_FRAME_LAG + 1);
}

bool SwappyVkGoogleDisplayTiming::doGetRefreshCycleDuration(
    VkSwapchainKHR swapchain, uint64_t* pRefreshDuration) {
    if (!isEnabled()) {
        SWAPPY_LOGE("Swappy is disabled.");
        return false;
    }

    VkRefreshCycleDurationGOOGLE refreshCycleDuration;
    VkResult res = mpfnGetRefreshCycleDurationGOOGLE(mDevice, swapchain,
                                                     &refreshCycleDuration);
    if (res != VK_SUCCESS) {
        SWAPPY_LOGE("mpfnGetRefreshCycleDurationGOOGLE failed %d", res);
        return false;
    }

    *pRefreshDuration = mCommonBase.getRefreshPeriod().count();

    // refreshRate is only used for logging, which maybe disabled.
    [[maybe_unused]] double refreshRate = 1000000000.0 / *pRefreshDuration;
    SWAPPY_LOGI("Returning refresh duration of %" PRIu64 " nsec (approx %f Hz)",
                *pRefreshDuration, refreshRate);

    mSwapchain = swapchain;
    return true;
}

VkResult SwappyVkGoogleDisplayTiming::doQueuePresent(
    VkQueue queue, uint32_t queueFamilyIndex,
    const VkPresentInfoKHR* pPresentInfo) {
    if (!isEnabled()) {
        SWAPPY_LOGE("Swappy is disabled.");
        return VK_ERROR_INITIALIZATION_FAILED;
    }

    VkResult res = initializeVkSyncObjects(queue, queueFamilyIndex);
    if (res) {
        return res;
    }

    const SwappyCommon::SwapHandlers handlers = {
        .lastFrameIsComplete = std::bind(
            &SwappyVkGoogleDisplayTiming::lastFrameIsCompleted, this, queue),
        .getPrevFrameGpuTime = std::bind(
            &SwappyVkGoogleDisplayTiming::getLastFenceTime, this, queue),
    };

    VkSemaphore semaphore;
    res = injectFence(queue, pPresentInfo, &semaphore);
    if (res) {
        SWAPPY_LOGE("Failed to vkQueueSubmit %d", res);
        return res;
    }

    uint32_t waitSemaphoreCount;
    const VkSemaphore* pWaitSemaphores;
    if (semaphore != VK_NULL_HANDLE) {
        waitSemaphoreCount = 1;
        pWaitSemaphores = &semaphore;
    } else {
        waitSemaphoreCount = pPresentInfo->waitSemaphoreCount;
        pWaitSemaphores = pPresentInfo->pWaitSemaphores;
    }

    mCommonBase.onPreSwap(handlers);

    VkPresentTimeGOOGLE pPresentTimes[pPresentInfo->swapchainCount];
    VkPresentInfoKHR replacementPresentInfo;
    VkPresentTimesInfoGOOGLE presentTimesInfo;
    // Set up the new structures to pass:
    // if 0 is passed as desired present time, it is ignored by the loader.
    uint64_t desiredPresentTime =
        mCommonBase.needToSetPresentationTime()
            ? mCommonBase.getPresentationTime().time_since_epoch().count()
            : 0;
    for (uint32_t i = 0; i < pPresentInfo->swapchainCount; i++) {
        pPresentTimes[i].presentID = mPresentID;
        pPresentTimes[i].desiredPresentTime = desiredPresentTime;
    }

    presentTimesInfo = {VK_STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE,
                        pPresentInfo->pNext, pPresentInfo->swapchainCount,
                        pPresentTimes};

    replacementPresentInfo = {
        pPresentInfo->sType,          &presentTimesInfo,
        waitSemaphoreCount,           pWaitSemaphores,
        pPresentInfo->swapchainCount, pPresentInfo->pSwapchains,
        pPresentInfo->pImageIndices,  pPresentInfo->pResults};

    mPresentID++;

    res = mpfnQueuePresentKHR(queue, &replacementPresentInfo);
    mCommonBase.onPostSwap(handlers);

    return res;
}

void SwappyVkGoogleDisplayTiming::enableStats(bool enabled) {
    mFrameStatisticsCommon.enableStats(enabled);
}

void SwappyVkGoogleDisplayTiming::recordFrameStart(VkQueue queue,
                                                   uint32_t image) {
    uint64_t frameStartTime = static_cast<uint64_t>(
        std::chrono::steady_clock::now().time_since_epoch().count());
    mPendingFrames.push_back({mPresentID, frameStartTime});

    // No point in querying if the history is too short, as vulkan loader does
    // not return any history newer than 5 frames.
    // See MIN_NUM_FRAMES_AGO in
    // https://android.googlesource.com/platform/frameworks/native/+/refs/heads/master/vulkan/libvulkan/swapchain.cpp
    if (mPendingFrames.size() < MIN_FRAME_LAG) return;

    // The query for vulkan past presentation timings does not point to any
    // specific id. Instead, the loader just returns whatever timings are
    // available to the user. Query all the available timings, which is a max
    // of 10 in the vulkan loader currently.
    //
    // Check through each of the timing if we have a pending frame id, if we do
    // then populate the histogram with the available timings. There are a
    // couple of assumptions made here.
    //  * The maximum size of the vectors here is 10, so simplicity is
    //  prioritized.
    //  * The frames are in order.
    //  * If any of the presentTimings ids are not present in mPendingFrames,
    //  those are frames that must have been cleared and we do not care about
    //  the timings anymore.
    //  * [Performance] Under normal smooth circumstances, this should be 1
    //  frame handled at a time, if there is a situation where several frames
    //  are pending, it implies that the gpu & presentation engine are not
    //  keeping up. So spending a few CPU cycles here to go through a few extra
    //  frames is not going to impact overall performance.
    uint32_t pastTimingsCount = MAX_FRAME_LAG;
    VkResult result = mpfnGetPastPresentationTimingGOOGLE(
        mDevice, mSwapchain, &pastTimingsCount, &mPastTimes[0]);

    if (result == VK_INCOMPLETE) {
        SWAPPY_LOGI(
            "More past presentation times available. Consider increasing "
            "MAX_FRAME_LAG");
    }
    if (result != VK_SUCCESS && result != VK_INCOMPLETE) {
        SWAPPY_LOGE("Error collecting past presentation times with result %d",
                    result);
        return;
    }

    int i = 0;
    while (i < pastTimingsCount && mPendingFrames.size() > 1) {
        auto frame = mPendingFrames.front();

        if (frame.id == mPastTimes[i].presentID) {
            FrameTimings current = {
                frame.startFrameTime, mPastTimes[i].desiredPresentTime,
                mPastTimes[i].actualPresentTime, mPastTimes[i].presentMargin};

            mFrameStatisticsCommon.updateFrameStats(
                current, mCommonBase.getRefreshPeriod().count());
            i++;
        }
        // If the past timings returned do not match, then the pending frame is
        // too old. So remove it from the list.
        mPendingFrames.erase(mPendingFrames.begin());
    }

    // Clear the pending frames if we are lagging too much.
    if (mPendingFrames.size() > MAX_FRAME_LAG) {
        while (mPendingFrames.size() > MIN_FRAME_LAG) {
            mPendingFrames.erase(mPendingFrames.begin());
        }
        mFrameStatisticsCommon.invalidateLastFrame();
    }
}

void SwappyVkGoogleDisplayTiming::getStats(SwappyStats* swappyStats) {
    *swappyStats = mFrameStatisticsCommon.getStats();
}

void SwappyVkGoogleDisplayTiming::clearStats() {
    mFrameStatisticsCommon.clearStats();
}
}  // namespace swappy

#endif  // #if (not defined ANDROID_NDK_VERSION) || ANDROID_NDK_VERSION>=15
