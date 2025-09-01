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
 * @file sys/select.h
 * @brief Wait for events on a set of file descriptors.
 * New code should prefer the different interface specified in <poll.h> instead,
 * because it scales better and easily avoids the limits on `fd_set` size.
 */

#include <sys/cdefs.h>
#include <sys/types.h>

#include <linux/time.h>
#include <signal.h>

__BEGIN_DECLS

typedef unsigned long fd_mask;

/**
 * The limit on the largest fd that can be used with type `fd_set`.
 * You can allocate your own memory,
 * but new code should prefer the different interface specified in <poll.h> instead,
 * because it scales better and easily avoids the limits on `fd_set` size.
 */
#define FD_SETSIZE 1024
#define NFDBITS (8 * sizeof(fd_mask))

/**
 * The type of a file descriptor set. Limited to 1024 fds.
 * The underlying system calls do not have this limit,
 * and callers can allocate their own sets with calloc().
 *
 * New code should prefer the different interface specified in <poll.h> instead,
 * because it scales better and easily avoids the limits on `fd_set` size.
 */
typedef struct {
  fd_mask fds_bits[FD_SETSIZE/NFDBITS];
} fd_set;

#define __FDELT(fd) ((fd) / NFDBITS)
#define __FDMASK(fd) (1UL << ((fd) % NFDBITS))
#define __FDS_BITS(type, set) (__BIONIC_CAST(static_cast, type, set)->fds_bits)

void __FD_CLR_chk(int, fd_set* _Nonnull , size_t);
void __FD_SET_chk(int, fd_set* _Nonnull, size_t);
int __FD_ISSET_chk(int, const fd_set* _Nonnull, size_t);

/**
 * FD_CLR() with no bounds checking for users that allocated their own set.
 * New code should prefer <poll.h> instead.
 */
#define __FD_CLR(fd, set) (__FDS_BITS(fd_set*, set)[__FDELT(fd)] &= ~__FDMASK(fd))

/**
 * FD_SET() with no bounds checking for users that allocated their own set.
 * New code should prefer <poll.h> instead.
 */
#define __FD_SET(fd, set) (__FDS_BITS(fd_set*, set)[__FDELT(fd)] |= __FDMASK(fd))

/**
 * FD_ISSET() with no bounds checking for users that allocated their own set.
 * New code should prefer <poll.h> instead.
 */
#define __FD_ISSET(fd, set) ((__FDS_BITS(const fd_set*, set)[__FDELT(fd)] & __FDMASK(fd)) != 0)

/**
 * Removes all 1024 fds from the given set.
 * Limited to fds under 1024.
 * New code should prefer <poll.h> instead for this reason,
 * rather than using memset() directly.
 */
#define FD_ZERO(set) __builtin_memset(set, 0, sizeof(*__BIONIC_CAST(static_cast, const fd_set*, set)))

/**
 * Removes `fd` from the given set.
 * Limited to fds under 1024.
 * New code should prefer <poll.h> instead for this reason,
 * rather than using __FD_CLR().
 */
#define FD_CLR(fd, set) __FD_CLR_chk(fd, set, __bos(set))

/**
 * Adds `fd` to the given set.
 * Limited to fds under 1024.
 * New code should prefer <poll.h> instead for this reason,
 * rather than using __FD_SET().
 */
#define FD_SET(fd, set) __FD_SET_chk(fd, set, __bos(set))

/**
 * Tests whether `fd` is in the given set.
 * Limited to fds under 1024.
 * New code should prefer <poll.h> instead for this reason,
 * rather than using __FD_ISSET().
 */
#define FD_ISSET(fd, set) __FD_ISSET_chk(fd, set, __bos(set))

/**
 * [select(2)](https://man7.org/linux/man-pages/man2/select.2.html) waits on a
 * set of file descriptors.
 *
 * New code should prefer poll() from <poll.h> instead,
 * because it scales better and easily avoids the limits on `fd_set` size.
 *
 * Returns the number of ready file descriptors on success, 0 for timeout,
 * and returns -1 and sets `errno` on failure.
 */
int select(int __max_fd_plus_one, fd_set* _Nullable __read_fds, fd_set* _Nullable __write_fds, fd_set* _Nullable __exception_fds, struct timeval* _Nullable __timeout);

/**
 * [pselect(2)](https://man7.org/linux/man-pages/man2/pselect.2.html) waits on a
 * set of file descriptors.
 *
 * New code should prefer ppoll() from <poll.h> instead,
 * because it scales better and easily avoids the limits on `fd_set` size.
 *
 * Returns the number of ready file descriptors on success, 0 for timeout,
 * and returns -1 and sets `errno` on failure.
 */
int pselect(int __max_fd_plus_one, fd_set* _Nullable __read_fds, fd_set* _Nullable __write_fds, fd_set* _Nullable __exception_fds, const struct timespec* _Nullable __timeout, const sigset_t* _Nullable __mask);

/**
 * [pselect64(2)](https://man7.org/linux/man-pages/man2/select.2.html) waits on a
 * set of file descriptors.
 *
 * New code should prefer ppoll64() from <poll.h> instead,
 * because it scales better and easily avoids the limits on `fd_set` size.
 *
 * Returns the number of ready file descriptors on success, 0 for timeout,
 * and returns -1 and sets `errno` on failure.
 *
 * Available since API level 28.
 */
#if __BIONIC_AVAILABILITY_GUARD(28)
int pselect64(int __max_fd_plus_one, fd_set* _Nullable __read_fds, fd_set* _Nullable __write_fds, fd_set* _Nullable __exception_fds, const struct timespec* _Nullable __timeout, const sigset64_t* _Nullable __mask) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

__END_DECLS
