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

// API entry points

#include <chrono>

#include "Settings.h"
#include "SwappyGL.h"
#include "swappy/swappyGL.h"

using namespace swappy;

extern "C" {

// Internal function to track Swappy version bundled in a binary.
void SWAPPY_VERSION_SYMBOL();

/**
 * @brief Initialize Swappy, getting the required Android parameters from the
 * display subsystem via JNI.
 * @param env The JNI environment where Swappy is used
 * @param jactivity The activity where Swappy is used
 * @return false if Swappy failed to initialize.
 * @see SwappyGL_destroy
 */
bool SwappyGL_init(JNIEnv *env, jobject jactivity) {
    // This call ensures that the header and the linked library are from the
    // same version (if not, a linker error will be triggered because of an
    // undefined symbolP).
    SWAPPY_VERSION_SYMBOL();
    return SwappyGL::init(env, jactivity);
}

void SwappyGL_destroy() { SwappyGL::destroyInstance(); }

void SwappyGL_onChoreographer(int64_t frameTimeNanos) {
    SwappyGL::onChoreographer(frameTimeNanos);
}

bool SwappyGL_setWindow(ANativeWindow *window) {
    return SwappyGL::setWindow(window);
}

bool SwappyGL_swap(EGLDisplay display, EGLSurface surface) {
    return SwappyGL::swap(display, surface);
}

void SwappyGL_setUseAffinity(bool tf) {
    Settings::getInstance()->setUseAffinity(tf);
}

void SwappyGL_setSwapIntervalNS(uint64_t swap_ns) {
    Settings::getInstance()->setSwapDuration(swap_ns);
}

uint64_t SwappyGL_getRefreshPeriodNanos() {
    return Settings::getInstance()->getDisplayTimings().refreshPeriod.count();
}

bool SwappyGL_getUseAffinity() {
    return Settings::getInstance()->getUseAffinity();
}

uint64_t SwappyGL_getSwapIntervalNS() {
    return SwappyGL::getSwapDuration().count();
}

void SwappyGL_injectTracer(const SwappyTracer *t) { SwappyGL::addTracer(t); }

void SwappyGL_setAutoSwapInterval(bool enabled) {
    SwappyGL::setAutoSwapInterval(enabled);
}

void SwappyGL_setMaxAutoSwapIntervalNS(uint64_t max_swap_ns) {
    SwappyGL::setMaxAutoSwapDuration(std::chrono::nanoseconds(max_swap_ns));
}

void SwappyGL_setAutoPipelineMode(bool enabled) {
    SwappyGL::setAutoPipelineMode(enabled);
}

void SwappyGL_enableStats(bool enabled) { SwappyGL::enableStats(enabled); }

void SwappyGL_recordFrameStart(EGLDisplay display, EGLSurface surface) {
    SwappyGL::recordFrameStart(display, surface);
}

void SwappyGL_getStats(SwappyStats *stats) { SwappyGL::getStats(stats); }

void SwappyGL_clearStats() { SwappyGL::clearStats(); }

bool SwappyGL_isEnabled() { return SwappyGL::isEnabled(); }

void SwappyGL_setFenceTimeoutNS(uint64_t t) {
    SwappyGL::setFenceTimeout(std::chrono::nanoseconds(t));
}

uint64_t SwappyGL_getFenceTimeoutNS() {
    return SwappyGL::getFenceTimeout().count();
}

void SwappyGL_setBufferStuffingFixWait(int32_t n_frames) {
    SwappyGL::setBufferStuffingFixWait(n_frames);
}

void SwappyGL_uninjectTracer(const SwappyTracer *t) {
    SwappyGL::removeTracer(t);
}

int SwappyGL_getSupportedRefreshPeriodsNS(uint64_t *out_refreshrates,
                                          int allocated_entries) {
    return SwappyGL::getSupportedRefreshPeriodsNS(out_refreshrates,
                                                  allocated_entries);
}

void SwappyGL_resetFramePacing() { SwappyGL::resetFramePacing(); }

void SwappyGL_enableFramePacing(bool enable) {
    SwappyGL::enableFramePacing(enable);
}

void SwappyGL_enableBlockingWait(bool enable) {
    SwappyGL::enableBlockingWait(enable);
}

}  // extern "C" {
