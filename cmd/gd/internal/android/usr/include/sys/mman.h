/*
 * Copyright (C) 2008 The Android Open Source Project
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *  * Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in
 *    the documentation and/or other materials provided with the
 *    distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
 * OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED
 * AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT
 * OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */

#pragma once

#include <sys/cdefs.h>
#include <sys/types.h>
#include <linux/memfd.h>
#include <linux/mman.h>
#include <linux/uio.h>

__BEGIN_DECLS

/** Alternative spelling of the `MAP_ANONYMOUS` flag for mmap(). */
#define MAP_ANON MAP_ANONYMOUS

/** Return value for mmap(). */
#define MAP_FAILED __BIONIC_CAST(reinterpret_cast, void*, -1)

/**
 * [mmap(2)](https://man7.org/linux/man-pages/man2/mmap.2.html)
 * creates a memory mapping for the given range.
 *
 * Returns the address of the mapping on success,
 * and returns `MAP_FAILED` and sets `errno` on failure.
 */
#if defined(__USE_FILE_OFFSET64)
void* _Nonnull mmap(void* _Nullable __addr, size_t __size, int __prot, int __flags, int __fd, off_t __offset) __RENAME(mmap64);
#else
void* _Nonnull mmap(void* _Nullable __addr, size_t __size, int __prot, int __flags, int __fd, off_t __offset);
#endif

/**
 * mmap64() is a variant of mmap() that takes a 64-bit offset even on LP32.
 *
 * See https://android.googlesource.com/platform/bionic/+/main/docs/32-bit-abi.md
 */
void* _Nonnull mmap64(void* _Nullable __addr, size_t __size, int __prot, int __flags, int __fd, off64_t __offset);

/**
 * [munmap(2)](https://man7.org/linux/man-pages/man2/munmap.2.html)
 * deletes a memory mapping for the given range.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int munmap(void* _Nonnull __addr, size_t __size);

/**
 * [msync(2)](https://man7.org/linux/man-pages/man2/msync.2.html)
 * flushes changes to a memory-mapped file to disk.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int msync(void* _Nonnull __addr, size_t __size, int __flags);

/**
 * [mprotect(2)](https://man7.org/linux/man-pages/man2/mprotect.2.html)
 * sets the protection on a memory region.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int mprotect(void* _Nonnull __addr, size_t __size, int __prot);

/** Flag for mremap(). */
#define MREMAP_MAYMOVE  1

/** Flag for mremap(). */
#define MREMAP_FIXED    2

/**
 * [mremap(2)](https://man7.org/linux/man-pages/man2/mremap.2.html)
 * expands or shrinks an existing memory mapping.
 *
 * Returns the address of the mapping on success,
 * and returns `MAP_FAILED` and sets `errno` on failure.
 */
void* _Nonnull mremap(void* _Nonnull __old_addr, size_t __old_size, size_t __new_size, int __flags, ...);

/**
 * [mlockall(2)](https://man7.org/linux/man-pages/man2/mlockall.2.html)
 * locks pages (preventing swapping).
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int mlockall(int __flags);

/**
 * [munlockall(2)](https://man7.org/linux/man-pages/man2/munlockall.2.html)
 * unlocks pages (allowing swapping).
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int munlockall(void);

/**
 * [mlock(2)](https://man7.org/linux/man-pages/man2/mlock.2.html)
 * locks pages (preventing swapping).
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int mlock(const void* _Nonnull __addr, size_t __size);

/**
 * [mlock2(2)](https://man7.org/linux/man-pages/man2/mlock2.2.html)
 * locks pages (preventing swapping), with optional flags.
 *
 * Available since API level 30.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(30)
int mlock2(const void* _Nonnull __addr, size_t __size, int __flags) __INTRODUCED_IN(30);
#endif /* __BIONIC_AVAILABILITY_GUARD(30) */


/**
 * [munlock(2)](https://man7.org/linux/man-pages/man2/munlock.2.html)
 * unlocks pages (allowing swapping).
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int munlock(const void* _Nonnull __addr, size_t __size);

/**
 * [mincore(2)](https://man7.org/linux/man-pages/man2/mincore.2.html)
 * tests whether pages are resident in memory.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int mincore(void* _Nonnull __addr, size_t __size, unsigned char* _Nonnull __vector);

/**
 * [madvise(2)](https://man7.org/linux/man-pages/man2/madvise.2.html)
 * gives the kernel advice about future usage patterns.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int madvise(void* _Nonnull __addr, size_t __size, int __advice);

/**
 * [process_madvise(2)](https://man7.org/linux/man-pages/man2/process_madvise.2.html)
 * works just like madvise(2) but applies to the process specified by the given
 * PID file descriptor.
 *
 * Available since API level 31. Its sibling process_mrelease() does not have a
 * libc wrapper and should be called using syscall() instead. Given the lack of
 * widespread applicability of this system call and the absence of wrappers in
 * other libcs, it was probably a mistake to have added this wrapper to bionic.
 *
 * Returns the number of bytes advised on success, and returns -1 and sets `errno` on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(31)
ssize_t process_madvise(int __pid_fd, const struct iovec* _Nonnull __iov, size_t __count, int __advice, unsigned __flags) __INTRODUCED_IN(31);
#endif /* __BIONIC_AVAILABILITY_GUARD(31) */


#if defined(__USE_GNU)

/**
 * [memfd_create(2)](https://man7.org/linux/man-pages/man2/memfd_create.2.html)
 * creates an anonymous file.
 *
 * Available since API level 30.
 *
 * Returns an fd on success, and returns -1 and sets `errno` on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(30)
int memfd_create(const char* _Nonnull __name, unsigned __flags) __INTRODUCED_IN(30);
#endif /* __BIONIC_AVAILABILITY_GUARD(30) */


#endif

#if __ANDROID_API__ >= 23

/*
 * Some third-party code uses the existence of POSIX_MADV_NORMAL to detect the
 * availability of posix_madvise. This is not correct, since having up-to-date
 * UAPI headers says nothing about the C library, but for the time being we
 * don't want to harm adoption of the unified headers.
 *
 * https://github.com/android-ndk/ndk/issues/395
 */

/** Flag for posix_madvise(). */
#define POSIX_MADV_NORMAL     MADV_NORMAL
/** Flag for posix_madvise(). */
#define POSIX_MADV_RANDOM     MADV_RANDOM
/** Flag for posix_madvise(). */
#define POSIX_MADV_SEQUENTIAL MADV_SEQUENTIAL
/** Flag for posix_madvise(). */
#define POSIX_MADV_WILLNEED   MADV_WILLNEED
/** Flag for posix_madvise(). */
#define POSIX_MADV_DONTNEED   MADV_DONTNEED

#endif

/**
 * [posix_madvise(3)](https://man7.org/linux/man-pages/man3/posix_madvise.3.html)
 * gives the kernel advice about future usage patterns.
 *
 * Available since API level 23.
 * See also madvise() which is available at all API levels.
 *
 * Returns 0 on success, and returns a positive error number on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(23)
int posix_madvise(void* _Nonnull __addr, size_t __size, int __advice) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


/**
 * [mseal(2)](https://man7.org/linux/man-pages/man2/mseal.2.html)
 * seals the given range to prevent modifications such as mprotect() calls.
 *
 * Available since API level 36.
 * Requires a Linux 6.10 or newer kernel.
 * Always fails for 32-bit processes.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(36)
int mseal(void* _Nonnull __addr, size_t __size, unsigned long __flags) __INTRODUCED_IN(36);
#endif /* __BIONIC_AVAILABILITY_GUARD(36) */


__END_DECLS
