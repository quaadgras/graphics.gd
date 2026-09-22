/*
 * Copyright 2022 The Android Open Source Project
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

/**
 * @file native_window_aidl.h
 * @brief NativeWindow NDK AIDL glue code
 */

/**
 * @addtogroup ANativeWindow
 *
 * Parcelable support for ANativeWindow. Can be used with libbinder_ndk
 *
 * @{
 */

#ifndef ANDROID_NATIVE_WINDOW_AIDL_H
#define ANDROID_NATIVE_WINDOW_AIDL_H

#include <android/binder_parcel.h>
#include <android/native_window.h>
#include <sys/cdefs.h>

// Only required by the AIDL glue helper
#ifdef __cplusplus
#include <sstream>
#include <string>
#endif // __cplusplus

__BEGIN_DECLS

/**
 * Read an ANativeWindow from a AParcel. The output buffer will have an
 * initial reference acquired and will need to be released with
 * ANativeWindow_release.
 *
 * Available since API level 34.
 *
 * \return STATUS_OK on success
 *         STATUS_BAD_VALUE if the parcel or outBuffer is null, or if there's an
 *                          issue deserializing (eg, corrupted parcel)
 *         STATUS_BAD_TYPE if the parcel's current data position is not that of
 *                         an ANativeWindow type
 *         STATUS_NO_MEMORY if an allocation fails
 */
binder_status_t ANativeWindow_readFromParcel(const AParcel* _Nonnull parcel,
        ANativeWindow* _Nullable* _Nonnull outWindow) __INTRODUCED_IN(__ANDROID_API_U__);

/**
 * Write an ANativeWindow to an AParcel.
 *
 * Available since API level 34.
 *
 * \return STATUS_OK on success.
 *         STATUS_BAD_VALUE if either buffer or parcel is null, or if the ANativeWindow*
 *                          fails to serialize (eg, internally corrupted)
 *         STATUS_NO_MEMORY if the parcel runs out of space to store the buffer & is
 *                          unable to allocate more
 *         STATUS_FDS_NOT_ALLOWED if the parcel does not allow storing FDs
 */
binder_status_t ANativeWindow_writeToParcel(ANativeWindow* _Nonnull window,
        AParcel* _Nonnull parcel) __INTRODUCED_IN(__ANDROID_API_U__);

__END_DECLS

// Only enable the AIDL glue helper if this is C++
#ifdef __cplusplus

namespace aidl::android::hardware {

/**
 * Wrapper class that enables interop with AIDL NDK generation
 * Takes ownership of the ANativeWindow* given to it in reset() and will automatically
 * destroy it in the destructor, similar to a smart pointer container
 */
class NativeWindow final {
public:
    NativeWindow() noexcept {}
    explicit NativeWindow(ANativeWindow* _Nullable window) {
        reset(window);
    }

    explicit NativeWindow(NativeWindow&& other) noexcept {
        mWindow = other.release(); // steal ownership from r-value
    }

    ~NativeWindow() {
        reset();
    }

    binder_status_t readFromParcel(const AParcel* _Nonnull parcel) {
        reset();
        if (__builtin_available(android __ANDROID_API_U__, *)) {
            return ANativeWindow_readFromParcel(parcel, &mWindow);
        } else {
            return STATUS_INVALID_OPERATION;
        }
    }

    binder_status_t writeToParcel(AParcel* _Nonnull parcel) const {
        if (!mWindow) {
            return STATUS_BAD_VALUE;
        }
        if (__builtin_available(android __ANDROID_API_U__, *)) {
            return ANativeWindow_writeToParcel(mWindow, parcel);
        } else {
            return STATUS_INVALID_OPERATION;
        }
    }

    /**
     * Destroys any currently owned ANativeWindow* and takes ownership of the given
     * ANativeWindow*
     *
     * @param buffer The buffer to take ownership of
     */
    void reset(ANativeWindow* _Nullable window = nullptr) noexcept {
        if (mWindow) {
            ANativeWindow_release(mWindow);
            mWindow = nullptr;
        }
        if (window != nullptr) {
            ANativeWindow_acquire(window);
        }
        mWindow = window;
    }

    inline ANativeWindow* _Nullable get() const { return mWindow; }

    NativeWindow& operator=(NativeWindow&& other) noexcept {
        mWindow = other.release(); // steal ownership from r-value
        return *this;
    }

    inline ANativeWindow* _Nullable operator->() const { return mWindow; }
    inline explicit operator bool() const { return mWindow != nullptr; }
    inline bool operator==(const NativeWindow& rhs) const { return mWindow == rhs.mWindow; }
    inline bool operator!=(const NativeWindow& rhs) const { return !(*this == rhs); }
    inline bool operator<(const NativeWindow& rhs) const { return mWindow < rhs.mWindow; }
    inline bool operator>(const NativeWindow& rhs) const { return rhs < *this; }
    inline bool operator>=(const NativeWindow& rhs) const { return !(*this < rhs); }
    inline bool operator<=(const NativeWindow& rhs) const { return !(*this > rhs); }

    std::string toString() const {
        std::ostringstream ss;
        ss << "NativeWindow: " << mWindow;
        return ss.str();
    }

    /**
     * Stops managing any contained ANativeWindow*, returning it to the caller. Ownership
     * is released.
     * @return ANativeWindow* or null if this was empty
     */
    [[nodiscard]] ANativeWindow* _Nullable release() noexcept {
        ANativeWindow* _Nullable ret = mWindow;
        mWindow = nullptr;
        return ret;
    }
private:
    ANativeWindow* _Nullable mWindow = nullptr;
    NativeWindow(const NativeWindow &other) = delete;
    NativeWindow& operator=(const NativeWindow &other) = delete;
};

} // aidl::android::hardware
  //
namespace aidl::android::view {
    using Surface = aidl::android::hardware::NativeWindow;
}

#endif // __cplusplus

#endif // ANDROID_NATIVE_WINDOW_AIDL_H

/** @} */
