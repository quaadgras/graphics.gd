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

#define LOG_TAG "SwappyDisplayManager"

#include "SwappyDisplayManager.h"

#include <android/looper.h>
#include <jni.h>

#include <map>

#include "JNIUtil.h"
#include "Settings.h"
#include "SwappyLog.h"

namespace swappy {

// Forward declaration of the native methods of Java SwappyDisplayManager class
extern "C" {

JNIEXPORT void JNICALL
Java_com_google_androidgamesdk_SwappyDisplayManager_nSetSupportedRefreshPeriods(
    JNIEnv *env, jobject /* this */, jlong cookie, jlongArray refreshPeriods,
    jintArray modeIds);

JNIEXPORT void JNICALL
Java_com_google_androidgamesdk_SwappyDisplayManager_nOnRefreshPeriodChanged(
    JNIEnv *env, jobject /* this */, jlong cookie, jlong refreshPeriod,
    jlong appOffset, jlong sfOffset);
}

const char *SwappyDisplayManager::SDM_CLASS =
    "com/google/androidgamesdk/SwappyDisplayManager";

const JNINativeMethod SwappyDisplayManager::SDMNativeMethods[] = {
    {"nSetSupportedRefreshPeriods", "(J[J[I)V",
     (void
          *)&Java_com_google_androidgamesdk_SwappyDisplayManager_nSetSupportedRefreshPeriods},
    {"nOnRefreshPeriodChanged", "(JJJJ)V",
     (void
          *)&Java_com_google_androidgamesdk_SwappyDisplayManager_nOnRefreshPeriodChanged}};

bool SwappyDisplayManager::usesMinSdkOrLater(SdkVersion sdkVersion) {
    return sdkVersion.sdkInt >= MIN_SDK_VERSION;
}

bool SwappyDisplayManager::useSwappyDisplayManager(SdkVersion sdkVersion) {
    // SwappyDisplayManager uses APIs introduced in SDK 23 and we get spurious
    // window messages for SDK < 28, so restrict here.
    if (!usesMinSdkOrLater(sdkVersion)) {
        return false;
    }

    // SDK 31 and above doesn't need SwappyDisplayManager as it has native
    // support in NDK. SDK 30 has partial native support
    // (AChoreographer_registerRefreshRateCallback) but lacks synchronization
    // with DisplayManager to query app/sf offsets
    return !(sdkVersion.sdkInt >= 31 ||
             (sdkVersion.sdkInt == 30 && sdkVersion.previewSdkInt == 1));
}

SwappyDisplayManager::SwappyDisplayManager(JavaVM *vm, jobject mainActivity)
    : mJVM(vm) {
    if (!vm || !mainActivity) {
        return;
    }

    JNIEnv *env;
    mJVM->AttachCurrentThread(&env, nullptr);

    jclass swappyDisplayManagerClass = gamesdk::loadClass(
        env, mainActivity, SwappyDisplayManager::SDM_CLASS,
        (JNINativeMethod *)SwappyDisplayManager::SDMNativeMethods,
        SwappyDisplayManager::SDMNativeMethodsSize);

    if (!swappyDisplayManagerClass) return;

    jmethodID constructor = env->GetMethodID(
        swappyDisplayManagerClass, "<init>", "(JLandroid/app/Activity;)V");
    mSetPreferredDisplayModeId = env->GetMethodID(
        swappyDisplayManagerClass, "setPreferredDisplayModeId", "(I)V");
    mTerminate =
        env->GetMethodID(swappyDisplayManagerClass, "terminate", "()V");
    jobject swappyDisplayManager = env->NewObject(
        swappyDisplayManagerClass, constructor, (jlong)this, mainActivity);
    mJthis = env->NewGlobalRef(swappyDisplayManager);

    mInitialized = true;
}

SwappyDisplayManager::~SwappyDisplayManager() {
    JNIEnv *env;
    mJVM->AttachCurrentThread(&env, nullptr);

    env->CallVoidMethod(mJthis, mTerminate);
    env->DeleteGlobalRef(mJthis);
}

std::shared_ptr<SwappyDisplayManager::RefreshPeriodMap>
SwappyDisplayManager::getSupportedRefreshPeriods() {
    std::unique_lock<std::mutex> lock(mMutex);

    mCondition.wait(
        lock, [&]() { return mSupportedRefreshPeriods.get() != nullptr; });
    return mSupportedRefreshPeriods;
}

void SwappyDisplayManager::setPreferredDisplayModeId(int index) {
    JNIEnv *env;
    mJVM->AttachCurrentThread(&env, nullptr);

    env->CallVoidMethod(mJthis, mSetPreferredDisplayModeId, index);
}

// Helper class to wrap JNI entry points to SwappyDisplayManager
class SwappyDisplayManagerJNI {
   public:
    static void onSetSupportedRefreshPeriods(
        jlong, std::shared_ptr<SwappyDisplayManager::RefreshPeriodMap>);
    static void onRefreshPeriodChanged(jlong, long, long, long);
};

void SwappyDisplayManagerJNI::onSetSupportedRefreshPeriods(
    jlong cookie,
    std::shared_ptr<SwappyDisplayManager::RefreshPeriodMap> refreshPeriods) {
    auto *sDM = reinterpret_cast<SwappyDisplayManager *>(cookie);

    std::lock_guard<std::mutex> lock(sDM->mMutex);
    sDM->mSupportedRefreshPeriods = std::move(refreshPeriods);
    sDM->mCondition.notify_one();
}

void SwappyDisplayManagerJNI::onRefreshPeriodChanged(jlong /*cookie*/,
                                                     long refreshPeriod,
                                                     long appOffset,
                                                     long sfOffset) {
    SWAPPY_LOGV("onRefreshPeriodChanged: refresh rate: %.0fHz",
                1e9f / refreshPeriod);
    using std::chrono::nanoseconds;
    Settings::DisplayTimings displayTimings;
    displayTimings.refreshPeriod = nanoseconds(refreshPeriod);
    displayTimings.appOffset = nanoseconds(appOffset);
    displayTimings.sfOffset = nanoseconds(sfOffset);
    Settings::getInstance()->setDisplayTimings(displayTimings);
}

extern "C" {

JNIEXPORT void JNICALL
Java_com_google_androidgamesdk_SwappyDisplayManager_nSetSupportedRefreshPeriods(
    JNIEnv *env, jobject /* this */, jlong cookie, jlongArray refreshPeriods,
    jintArray modeIds) {
    int length = env->GetArrayLength(refreshPeriods);
    auto refreshPeriodsMap =
        std::make_shared<SwappyDisplayManager::RefreshPeriodMap>();

    jlong *refreshPeriodsArr = env->GetLongArrayElements(refreshPeriods, 0);
    jint *modeIdsArr = env->GetIntArrayElements(modeIds, 0);
    for (int i = 0; i < length; i++) {
        (*refreshPeriodsMap)[std::chrono::nanoseconds(refreshPeriodsArr[i])] =
            modeIdsArr[i];
    }
    env->ReleaseLongArrayElements(refreshPeriods, refreshPeriodsArr, 0);
    env->ReleaseIntArrayElements(modeIds, modeIdsArr, 0);

    SwappyDisplayManagerJNI::onSetSupportedRefreshPeriods(cookie,
                                                          refreshPeriodsMap);
}

JNIEXPORT void JNICALL
Java_com_google_androidgamesdk_SwappyDisplayManager_nOnRefreshPeriodChanged(
    JNIEnv *env, jobject /* this */, jlong cookie, jlong refreshPeriod,
    jlong appOffset, jlong sfOffset) {
    SwappyDisplayManagerJNI::onRefreshPeriodChanged(cookie, refreshPeriod,
                                                    appOffset, sfOffset);
}

}  // extern "C"

}  // namespace swappy
