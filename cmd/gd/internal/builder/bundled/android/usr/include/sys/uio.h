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

/**
 * @file sys/uio.h
 * @brief Multi-buffer ("vector") I/O operations using `struct iovec`.
 */

#include <sys/cdefs.h>
#include <sys/types.h>
#include <linux/uio.h>

__BEGIN_DECLS

/**
 * [readv(2)](https://man7.org/linux/man-pages/man2/readv.2.html) reads
 * from an fd into the `__count` buffers described by `__iov`.
 *
 * Returns the number of bytes read on success,
 * and returns -1 and sets `errno` on failure.
 */
ssize_t readv(int __fd, const struct iovec* _Nonnull __iov, int __count);

/**
 * [writev(2)](https://man7.org/linux/man-pages/man2/writev.2.html) writes
 * to an fd from the `__count` buffers described by `__iov`.
 *
 * Returns the number of bytes written on success,
 * and returns -1 and sets `errno` on failure.
 */
ssize_t writev(int __fd, const struct iovec* _Nonnull __iov, int __count);

#if defined(__USE_GNU)

/**
 * [preadv(2)](https://man7.org/linux/man-pages/man2/preadv.2.html) reads
 * from an fd into the `__count` buffers described by `__iov`, starting at
 * offset `__offset` into the file.
 *
 * Returns the number of bytes read on success,
 * and returns -1 and sets `errno` on failure.
 *
 * Available since API level 24.
 */

#if __BIONIC_AVAILABILITY_GUARD(24)
ssize_t preadv(int __fd, const struct iovec* _Nonnull __iov, int __count, off_t __offset) __RENAME_IF_FILE_OFFSET64(preadv64) __INTRODUCED_IN(24);

/**
 * [pwritev(2)](https://man7.org/linux/man-pages/man2/pwritev.2.html) writes
 * to an fd from the `__count` buffers described by `__iov`, starting at offset
 * `__offset` into the file.
 *
 * Returns the number of bytes written on success,
 * and returns -1 and sets `errno` on failure.
 *
 * Available since API level 24.
 */
ssize_t pwritev(int __fd, const struct iovec* _Nonnull __iov, int __count, off_t __offset) __RENAME_IF_FILE_OFFSET64(pwritev64) __INTRODUCED_IN(24);

/**
 * Like preadv() but with a 64-bit offset even in a 32-bit process.
 *
 * Available since API level 24.
 */
ssize_t preadv64(int __fd, const struct iovec* _Nonnull __iov, int __count, off64_t __offset) __INTRODUCED_IN(24);

/**
 * Like pwritev() but with a 64-bit offset even in a 32-bit process.
 *
 * Available since API level 24.
 */
ssize_t pwritev64(int __fd, const struct iovec* _Nonnull __iov, int __count, off64_t __offset) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */


/**
 * [preadv2(2)](https://man7.org/linux/man-pages/man2/preadv2.2.html) reads
 * from an fd into the `__count` buffers described by `__iov`, starting at
 * offset `__offset` into the file, with the given flags.
 *
 * Returns the number of bytes read on success,
 * and returns -1 and sets `errno` on failure.
 *
 * Available since API level 33.
 */

#if __BIONIC_AVAILABILITY_GUARD(33)
ssize_t preadv2(int __fd, const struct iovec* _Nonnull __iov, int __count, off_t __offset, int __flags) __RENAME_IF_FILE_OFFSET64(preadv64v2) __INTRODUCED_IN(33);

/**
 * [pwritev2(2)](https://man7.org/linux/man-pages/man2/pwritev2.2.html) writes
 * to an fd from the `__count` buffers described by `__iov`, starting at offset
 * `__offset` into the file, with the given flags.
 *
 * Returns the number of bytes written on success,
 * and returns -1 and sets `errno` on failure.
 *
 * Available since API level 33.
 */
ssize_t pwritev2(int __fd, const struct iovec* _Nonnull __iov, int __count, off_t __offset, int __flags) __RENAME_IF_FILE_OFFSET64(pwritev64v2) __INTRODUCED_IN(33);

/**
 * Like preadv2() but with a 64-bit offset even in a 32-bit process.
 *
 * Available since API level 33.
 */
ssize_t preadv64v2(int __fd, const struct iovec* _Nonnull __iov, int __count, off64_t __offset, int __flags) __INTRODUCED_IN(33);

/**
 * Like pwritev2() but with a 64-bit offset even in a 32-bit process.
 *
 * Available since API level 33.
 */
ssize_t pwritev64v2(int __fd, const struct iovec* _Nonnull __iov, int __count, off64_t __offset, int __flags) __INTRODUCED_IN(33);
#endif /* __BIONIC_AVAILABILITY_GUARD(33) */


/**
 * [process_vm_readv(2)](https://man7.org/linux/man-pages/man2/process_vm_readv.2.html)
 * reads from the address space of another process.
 *
 * Returns the number of bytes read on success,
 * and returns -1 and sets `errno` on failure.
 *
 * Available since API level 23.
 */

#if __BIONIC_AVAILABILITY_GUARD(23)
ssize_t process_vm_readv(pid_t __pid, const struct iovec* __BIONIC_COMPLICATED_NULLNESS __local_iov, unsigned long __local_iov_count, const struct iovec* __BIONIC_COMPLICATED_NULLNESS __remote_iov, unsigned long __remote_iov_count, unsigned long __flags) __INTRODUCED_IN(23);

/**
 * [process_vm_writev(2)](https://man7.org/linux/man-pages/man2/process_vm_writev.2.html)
 * writes to the address space of another process.
 *
 * Returns the number of bytes read on success,
 * and returns -1 and sets `errno` on failure.
 *
 * Available since API level 23.
 */
ssize_t process_vm_writev(pid_t __pid, const struct iovec* __BIONIC_COMPLICATED_NULLNESS __local_iov, unsigned long __local_iov_count, const struct iovec* __BIONIC_COMPLICATED_NULLNESS __remote_iov, unsigned long __remote_iov_count, unsigned long __flags) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


#endif

__END_DECLS
