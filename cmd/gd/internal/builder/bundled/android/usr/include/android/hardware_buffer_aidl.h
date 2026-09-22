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
 * @file hardware_buffer_aidl.h
 * @brief HardwareBuffer NDK AIDL glue code
 */

/**
 * @addtogroup AHardwareBuffer
 *
 * Parcelable support for AHardwareBuffer. Can be used with libbinder_ndk
 *
 * @{
 */

#ifndef ANDROID_HARDWARE_BUFFER_AIDL_H
#define ANDROID_HARDWARE_BUFFER_AIDL_H

#include <android/binder_parcel.h>
#include <android/hardware_buffer.h>
#include <sys/cdefs.h>

#ifdef __cplusplus
#include <string>
#endif

__BEGIN_DECLS

/**
 * Read an AHardwareBuffer from a AParcel. The output buffer will have an
 * initial reference acquired and will need to be released with
 * AHardwareBuffer_release.
 *
 * Available since API level 34.
 *
 * \return STATUS_OK on success
 *         STATUS_BAD_VALUE if the parcel or outBuffer is null, or if there's an
 *                          issue deserializing (eg, corrupted parcel)
 *         STATUS_BAD_TYPE if the parcel's current data position is not that of
 *                         an AHardwareBuffer type
 *         STATUS_NO_MEMORY if an allocation fails
 */
binder_status_t AHardwareBuffer_readFromParcel(const AParcel* _Nonnull parcel,
        AHardwareBuffer* _Nullable* _Nonnull outBuffer) __INTRODUCED_IN(34);

/**
 * Write an AHardwareBuffer to an AParcel.
 *
 * Available since API level 34.
 *
 * \return STATUS_OK on success.
 *         STATUS_BAD_VALUE if either buffer or parcel is null, or if the AHardwareBuffer*
 *                          fails to serialize (eg, internally corrupted)
 *         STATUS_NO_MEMORY if the parcel runs out of space to store the buffer & is
 *                          unable to allocate more
 *         STATUS_FDS_NOT_ALLOWED if the parcel does not allow storing FDs
 */
binder_status_t AHardwareBuffer_writeToParcel(const AHardwareBuffer* _Nonnull buffer,
        AParcel* _Nonnull parcel) __INTRODUCED_IN(34);

__END_DECLS

// Only enable the AIDL glue helper if this is C++
#ifdef __cplusplus

namespace aidl::android::hardware {

/**
 * Wrapper class that enables interop with AIDL NDK generation
 * Takes ownership of the AHardwareBuffer* given to it in reset() and will automatically
 * destroy it in the destructor, similar to a smart pointer container
 */
class HardwareBuffer {
public:
    HardwareBuffer() noexcept {}
    HardwareBuffer(HardwareBuffer&& other) noexcept : mBuffer(other.release()) {}

    ~HardwareBuffer() {
        reset();
    }

    binder_status_t readFromParcel(const AParcel* _Nonnull parcel) {
        reset();
        if (__builtin_available(android __ANDROID_API_U__, *)) {
            return AHardwareBuffer_readFromParcel(parcel, &mBuffer);
        } else {
            return STATUS_FAILED_TRANSACTION;
        }
    }

    binder_status_t writeToParcel(AParcel* _Nonnull parcel) const {
        if (!mBuffer) {
            return STATUS_BAD_VALUE;
        }
        if (__builtin_available(android __ANDROID_API_U__, *)) {
            return AHardwareBuffer_writeToParcel(mBuffer, parcel);
        } else {
            return STATUS_FAILED_TRANSACTION;
        }
    }

    /**
     * Destroys any currently owned AHardwareBuffer* and takes ownership of the given
     * AHardwareBuffer*
     *
     * @param buffer The buffer to take ownership of
     */
    void reset(AHardwareBuffer* _Nullable buffer = nullptr) noexcept {
        if (mBuffer) {
            AHardwareBuffer_release(mBuffer);
            mBuffer = nullptr;
        }
        mBuffer = buffer;
    }

    inline AHardwareBuffer* _Nullable operator-> () const { return mBuffer;  }
    inline AHardwareBuffer* _Nullable get() const { return mBuffer; }
    inline explicit operator bool () const { return mBuffer != nullptr; }

    inline bool operator!=(const HardwareBuffer& rhs) const { return get() != rhs.get(); }
    inline bool operator<(const HardwareBuffer& rhs) const { return get() < rhs.get(); }
    inline bool operator<=(const HardwareBuffer& rhs) const { return get() <= rhs.get(); }
    inline bool operator==(const HardwareBuffer& rhs) const { return get() == rhs.get(); }
    inline bool operator>(const HardwareBuffer& rhs) const { return get() > rhs.get(); }
    inline bool operator>=(const HardwareBuffer& rhs) const { return get() >= rhs.get(); }

    HardwareBuffer& operator=(HardwareBuffer&& other) noexcept {
        reset(other.release());
        return *this;
    }

    /**
     * Stops managing any contained AHardwareBuffer*, returning it to the caller. Ownership
     * is released.
     * @return AHardwareBuffer* or null if this was empty
     */
    [[nodiscard]] AHardwareBuffer* _Nullable release() noexcept {
        AHardwareBuffer* _Nullable ret = mBuffer;
        mBuffer = nullptr;
        return ret;
    }

    inline std::string toString() const {
        if (!mBuffer) {
            return "<HardwareBuffer: Invalid>";
        }
        if (__builtin_available(android __ANDROID_API_S__, *)) {
            uint64_t id = 0;
            AHardwareBuffer_getId(mBuffer, &id);
            return "<HardwareBuffer " + std::to_string(id) + ">";
        } else {
            return "<HardwareBuffer (unknown)>";
        }
    }

private:
    HardwareBuffer(const HardwareBuffer& other) = delete;
    HardwareBuffer& operator=(const HardwareBuffer& other) = delete;

    AHardwareBuffer* _Nullable mBuffer = nullptr;
};

} // aidl::android::hardware

#endif // __cplusplus

#endif // ANDROID_HARDWARE_BUFFER_AIDL_H

/** @} */
