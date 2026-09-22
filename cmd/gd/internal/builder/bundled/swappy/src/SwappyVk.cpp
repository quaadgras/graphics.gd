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

#include "SwappyVk.h"

#define LOG_TAG "SwappyVk"
#include "SwappyLog.h"

/* For tracking tracers internally within vulkan instance, we cannot store the
 * pointers as there is no requirement for the life of the object once the
 * SwappyVk_injectTracer(*t) call is returned. So we store the struct objects by
 * value. In turn, this implies that we have to compare the whole struct for
 * managing it. So we define a local "==" operator to this file in order to
 * handle the C struct API in the std::list.
 *
 * In addition, there is a risk that updates to the SwappyTracer struct could go
 * unnoticed silently. So we have a copy here and we check at compile time that
 * the sizes are the same, it is a weak check, but it is better than nothing.
 */

typedef struct SwappyTracerLocalStruct {
    SwappyPreWaitCallback preWait;
    SwappyPostWaitCallback postWait;
    SwappyPreSwapBuffersCallback preSwapBuffers;
    SwappyPostSwapBuffersCallback postSwapBuffers;
    SwappyStartFrameCallback startFrame;
    void* userData;
    SwappySwapIntervalChangedCallback swapIntervalChanged;
} SwappyTracerLocalStruct;

static bool operator==(const SwappyTracer& t1, const SwappyTracer& t2) {
    static_assert(sizeof(SwappyTracer) == sizeof(SwappyTracerLocalStruct),
                  "SwappyTracer struct appears to have changed, please "
                  "consider updating locally.");
    return (t1.preWait == t2.preWait) && (t1.postWait == t2.postWait) &&
           (t1.preSwapBuffers == t2.preSwapBuffers) &&
           (t1.postSwapBuffers == t2.postSwapBuffers) &&
           (t1.startFrame == t2.startFrame) && (t1.userData == t2.userData) &&
           (t1.swapIntervalChanged == t2.swapIntervalChanged);
}

namespace swappy {

class DefaultSwappyVkFunctionProvider {
   public:
    static bool Init() {
        if (!mLibVulkan) {
            // This is the first time we've been called
            mLibVulkan = dlopen("libvulkan.so", RTLD_NOW | RTLD_LOCAL);
            if (!mLibVulkan) {
                // If Vulkan doesn't exist, bail-out early:
                return false;
            }
        }
        return true;
    }
    static void* GetProcAddr(const char* name) {
        if (!mLibVulkan && !Init()) return nullptr;
        return dlsym(mLibVulkan, name);
    }
    static void Close() {
        if (mLibVulkan) {
            dlclose(mLibVulkan);
            mLibVulkan = nullptr;
        }
    }

   private:
    static void* mLibVulkan;
};

void* DefaultSwappyVkFunctionProvider::mLibVulkan = nullptr;

bool SwappyVk::InitFunctions() {
    if (pFunctionProvider == nullptr) {
        static SwappyVkFunctionProvider c_provider;
        c_provider.init = &DefaultSwappyVkFunctionProvider::Init;
        c_provider.getProcAddr = &DefaultSwappyVkFunctionProvider::GetProcAddr;
        c_provider.close = &DefaultSwappyVkFunctionProvider::Close;
        pFunctionProvider = &c_provider;
    }
    if (pFunctionProvider->init()) {
        LoadVulkanFunctions(pFunctionProvider);
        return true;
    } else {
        return false;
    }
}
void SwappyVk::SetFunctionProvider(
    const SwappyVkFunctionProvider* functionProvider) {
    if (pFunctionProvider != nullptr) pFunctionProvider->close();
    pFunctionProvider = functionProvider;
}

/**
 * Generic/Singleton implementation of swappyVkDetermineDeviceExtensions.
 */
void SwappyVk::swappyVkDetermineDeviceExtensions(
    VkPhysicalDevice physicalDevice, uint32_t availableExtensionCount,
    VkExtensionProperties* pAvailableExtensions,
    uint32_t* pRequiredExtensionCount, char** pRequiredExtensions) {
#if (not defined ANDROID_NDK_VERSION) || ANDROID_NDK_VERSION >= 15
    // TODO: Refactor this to be more concise:
    if (!pRequiredExtensions) {
        for (uint32_t i = 0; i < availableExtensionCount; i++) {
            if (!strcmp(VK_GOOGLE_DISPLAY_TIMING_EXTENSION_NAME,
                        pAvailableExtensions[i].extensionName)) {
                (*pRequiredExtensionCount)++;
            }
        }
    } else {
        doesPhysicalDeviceHaveGoogleDisplayTiming[physicalDevice] = false;
        for (uint32_t i = 0, j = 0; i < availableExtensionCount; i++) {
            if (!strcmp(VK_GOOGLE_DISPLAY_TIMING_EXTENSION_NAME,
                        pAvailableExtensions[i].extensionName)) {
                if (j < *pRequiredExtensionCount) {
                    strcpy(pRequiredExtensions[j++],
                           VK_GOOGLE_DISPLAY_TIMING_EXTENSION_NAME);
                    doesPhysicalDeviceHaveGoogleDisplayTiming[physicalDevice] =
                        true;
                }
            }
        }
    }
#else
    doesPhysicalDeviceHaveGoogleDisplayTiming[physicalDevice] = false;
#endif
}

void SwappyVk::SetQueueFamilyIndex(VkDevice device, VkQueue queue,
                                   uint32_t queueFamilyIndex) {
    perQueueFamilyIndex[queue] = {device, queueFamilyIndex};
}

/**
 * Generic/Singleton implementation of swappyVkGetRefreshCycleDuration.
 */
bool SwappyVk::GetRefreshCycleDuration(JNIEnv* env, jobject jactivity,
                                       VkPhysicalDevice physicalDevice,
                                       VkDevice device,
                                       VkSwapchainKHR swapchain,
                                       uint64_t* pRefreshDuration) {
    auto& pImplementation = perSwapchainImplementation[swapchain];
    if (!pImplementation) {
        if (!InitFunctions()) {
            // If Vulkan doesn't exist, bail-out early
            return false;
        }

#if (not defined ANDROID_NDK_VERSION) || ANDROID_NDK_VERSION >= 15
        // First, based on whether VK_GOOGLE_display_timing is available
        // (determined and cached by swappyVkDetermineDeviceExtensions),
        // determine which derived class to use to implement the rest of the API
        if (doesPhysicalDeviceHaveGoogleDisplayTiming[physicalDevice]) {
            pImplementation = std::make_shared<SwappyVkGoogleDisplayTiming>(
                env, jactivity, physicalDevice, device, pFunctionProvider);
            SWAPPY_LOGV(
                "SwappyVk initialized for VkDevice %p using "
                "VK_GOOGLE_display_timing on Android",
                device);
        } else
#endif
        {
            pImplementation = std::make_shared<SwappyVkFallback>(
                env, jactivity, physicalDevice, device, pFunctionProvider);
            SWAPPY_LOGV(
                "SwappyVk initialized for VkDevice %p using Android fallback",
                device);
        }

        if (!pImplementation) {  // should never happen
            SWAPPY_LOGE(
                "SwappyVk could not find or create correct implementation for "
                "the current environment: "
                "%p, %p",
                physicalDevice, device);
            return false;
        }
    }

    // SwappyBase is constructed by this point, so we can add the tracers we
    // have so far.
    {
        std::lock_guard<std::mutex> lock(tracer_list_lock);
        for (const auto& tracer : tracer_list) {
            pImplementation->addTracer(&tracer);
        }
    }
    // Now, call that derived class to get the refresh duration to return
    return pImplementation->doGetRefreshCycleDuration(swapchain,
                                                      pRefreshDuration);
}

/**
 * Generic/Singleton implementation of swappyVkSetWindow.
 */
void SwappyVk::SetWindow(VkDevice device, VkSwapchainKHR swapchain,
                         ANativeWindow* window) {
    auto& pImplementation = perSwapchainImplementation[swapchain];
    if (!pImplementation) {
        return;
    }
    pImplementation->doSetWindow(window);
}

/**
 * Generic/Singleton implementation of swappyVkSetSwapInterval.
 */
void SwappyVk::SetSwapDuration(VkDevice device, VkSwapchainKHR swapchain,
                               uint64_t swapNs) {
    auto& pImplementation = perSwapchainImplementation[swapchain];
    if (!pImplementation) {
        return;
    }
    pImplementation->doSetSwapInterval(swapchain, swapNs);
}

/**
 * Generic/Singleton implementation of swappyVkQueuePresent.
 */
VkResult SwappyVk::QueuePresent(VkQueue queue,
                                const VkPresentInfoKHR* pPresentInfo) {
    if (perQueueFamilyIndex.find(queue) == perQueueFamilyIndex.end()) {
        SWAPPY_LOGE(
            "Unknown queue %p. Did you call SwappyVkSetQueueFamilyIndex ?",
            queue);
        return VK_INCOMPLETE;
    }

    // This command doesn't have a VkDevice.  It should have at least one
    // VkSwapchainKHR's.  For this command, all VkSwapchainKHR's will have the
    // same VkDevice and VkQueue.
    if ((pPresentInfo->swapchainCount == 0) || (!pPresentInfo->pSwapchains)) {
        // This shouldn't happen, but if it does, something is really wrong.
        return VK_ERROR_DEVICE_LOST;
    }
    auto& pImplementation =
        perSwapchainImplementation[*pPresentInfo->pSwapchains];
    if (pImplementation) {
        return pImplementation->doQueuePresent(
            queue, perQueueFamilyIndex[queue].queueFamilyIndex, pPresentInfo);
    } else {
        // This should only happen if the API was used wrong (e.g. they never
        // called swappyVkGetRefreshCycleDuration).
        // NOTE: Technically, a Vulkan library shouldn't protect a user from
        // themselves, but we'll be friendlier
        return VK_ERROR_DEVICE_LOST;
    }
}

void SwappyVk::DestroySwapchain(VkDevice /*device*/, VkSwapchainKHR swapchain) {
    auto swapchain_it = perSwapchainImplementation.find(swapchain);
    if (swapchain_it == perSwapchainImplementation.end()) return;
    perSwapchainImplementation.erase(swapchain);
}

void SwappyVk::DestroyDevice(VkDevice device) {
    {
        // Erase swapchains
        auto it = perSwapchainImplementation.begin();
        while (it != perSwapchainImplementation.end()) {
            if (it->second->getDevice() == device) {
                it = perSwapchainImplementation.erase(it);
            } else {
                ++it;
            }
        }
    }
    {
        // Erase the device
        auto it = perQueueFamilyIndex.begin();
        while (it != perQueueFamilyIndex.end()) {
            if (it->second.device == device) {
                it = perQueueFamilyIndex.erase(it);
            } else {
                ++it;
            }
        }
    }
}

void SwappyVk::SetAutoSwapInterval(bool enabled) {
    for (auto i : perSwapchainImplementation) {
        i.second->setAutoSwapInterval(enabled);
    }
}

void SwappyVk::SetAutoPipelineMode(bool enabled) {
    for (auto i : perSwapchainImplementation) {
        i.second->setAutoPipelineMode(enabled);
    }
}

void SwappyVk::SetMaxAutoSwapDuration(std::chrono::nanoseconds maxDuration) {
    for (auto i : perSwapchainImplementation) {
        i.second->setMaxAutoSwapDuration(maxDuration);
    }
}

void SwappyVk::SetFenceTimeout(std::chrono::nanoseconds t) {
    for (auto i : perSwapchainImplementation) {
        i.second->setFenceTimeout(t);
    }
}

std::chrono::nanoseconds SwappyVk::GetFenceTimeout() const {
    auto it = perSwapchainImplementation.begin();
    if (it != perSwapchainImplementation.end()) {
        return it->second->getFenceTimeout();
    }
    return std::chrono::nanoseconds(0);
}

std::chrono::nanoseconds SwappyVk::GetSwapInterval(VkSwapchainKHR swapchain) {
    auto it = perSwapchainImplementation.find(swapchain);
    if (it != perSwapchainImplementation.end())
        return it->second->getSwapInterval();
    return std::chrono::nanoseconds(0);
}

void SwappyVk::addTracer(const SwappyTracer* t) {
    if (t != nullptr) {
        std::lock_guard<std::mutex> lock(tracer_list_lock);
        tracer_list.push_back(*t);

        for (const auto& i : perSwapchainImplementation) {
            i.second->addTracer(t);
        }
    }
}

void SwappyVk::removeTracer(const SwappyTracer* t) {
    if (t != nullptr) {
        std::lock_guard<std::mutex> lock(tracer_list_lock);
        tracer_list.remove(*t);

        for (const auto& i : perSwapchainImplementation) {
            i.second->removeTracer(t);
        }
    }
}

int SwappyVk::GetSupportedRefreshPeriodsNS(uint64_t* out_refreshrates,
                                           int allocated_entries,
                                           VkSwapchainKHR swapchain) {
    return (*perSwapchainImplementation[swapchain])
        .getSupportedRefreshPeriodsNS(out_refreshrates, allocated_entries);
}

bool SwappyVk::IsEnabled(VkSwapchainKHR swapchain, bool* isEnabled) {
    auto& pImplementation = perSwapchainImplementation[swapchain];
    if (!pImplementation || !isEnabled) return false;
    *isEnabled = pImplementation->isEnabled();
    return true;
}

void SwappyVk::enableStats(VkSwapchainKHR swapchain, bool enabled) {
    auto it = perSwapchainImplementation.find(swapchain);
    if (it != perSwapchainImplementation.end())
        it->second->enableStats(enabled);
}

void SwappyVk::getStats(VkSwapchainKHR swapchain, SwappyStats* swappyStats) {
    auto it = perSwapchainImplementation.find(swapchain);
    if (it != perSwapchainImplementation.end())
        it->second->getStats(swappyStats);
}

void SwappyVk::recordFrameStart(VkQueue queue, VkSwapchainKHR swapchain,
                                uint32_t image) {
    auto it = perSwapchainImplementation.find(swapchain);
    if (it != perSwapchainImplementation.end())
        it->second->recordFrameStart(queue, image);
}

void SwappyVk::clearStats(VkSwapchainKHR swapchain) {
    auto it = perSwapchainImplementation.find(swapchain);
    if (it != perSwapchainImplementation.end()) it->second->clearStats();
}

void SwappyVk::resetFramePacing(VkSwapchainKHR swapchain) {
    auto it = perSwapchainImplementation.find(swapchain);
    if (it != perSwapchainImplementation.end()) it->second->resetFramePacing();
}

void SwappyVk::enableFramePacing(VkSwapchainKHR swapchain, bool enable) {
    auto it = perSwapchainImplementation.find(swapchain);
    if (it != perSwapchainImplementation.end())
        it->second->enableFramePacing(enable);
}

void SwappyVk::enableBlockingWait(VkSwapchainKHR swapchain, bool enable) {
    auto it = perSwapchainImplementation.find(swapchain);
    if (it != perSwapchainImplementation.end())
        it->second->enableBlockingWait(enable);
}

}  // namespace swappy
