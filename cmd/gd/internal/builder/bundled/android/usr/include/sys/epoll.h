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
 * @file sys/epoll.h
 * @brief I/O event file descriptors.
 */

#include <sys/cdefs.h>
#include <sys/types.h>
#include <signal.h> /* For sigset_t. */

#include <linux/eventpoll.h>

__BEGIN_DECLS

/**
 * [epoll_create(2)](https://man7.org/linux/man-pages/man2/epoll_create.2.html)
 * creates a new [epoll](https://man7.org/linux/man-pages/man7/epoll.7.html)
 * file descriptor.
 *
 * Returns a new file descriptor on success and returns -1 and sets `errno` on
 * failure.
 */
int epoll_create(int __size);

/**
 * [epoll_create1(2)](https://man7.org/linux/man-pages/man2/epoll_create1.2.html)
 * creates a new [epoll](https://man7.org/linux/man-pages/man7/epoll.7.html)
 * file descriptor with the given flags.
 *
 * Returns a new file descriptor on success and returns -1 and sets `errno` on
 * failure.
 */
int epoll_create1(int __flags);

/**
 * [epoll_ctl(2)](https://man7.org/linux/man-pages/man2/epoll_ctl.2.html)
 * adds/modifies/removes file descriptors from the given epoll file descriptor.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int epoll_ctl(int __epoll_fd, int __op, int __fd, struct epoll_event* __BIONIC_COMPLICATED_NULLNESS __event);

/**
 * [epoll_wait(2)](https://man7.org/linux/man-pages/man2/epoll_wait.2.html)
 * waits for an event on the given epoll file descriptor.
 *
 * Returns the number of ready file descriptors on success, 0 on timeout,
 * or -1 and sets `errno` on failure.
 */
int epoll_wait(int __epoll_fd, struct epoll_event* _Nonnull __events, int __event_count, int __timeout_ms);

/**
 * Like epoll_wait() but atomically applying the given signal mask.
 */
int epoll_pwait(int __epoll_fd, struct epoll_event* _Nonnull __events, int __event_count, int __timeout_ms, const sigset_t* _Nullable __mask);

/**
 * Like epoll_pwait() but using a 64-bit signal mask even on 32-bit systems.
 *
 * Available since API level 28.
 */

#if __BIONIC_AVAILABILITY_GUARD(28)
int epoll_pwait64(int __epoll_fd, struct epoll_event* _Nonnull __events, int __event_count, int __timeout_ms, const sigset64_t* _Nullable __mask) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


/**
 * Like epoll_pwait() but with a `struct timespec` timeout, for nanosecond resolution.
 *
 * Available since API level 35.
 */

#if __BIONIC_AVAILABILITY_GUARD(35)
int epoll_pwait2(int __epoll_fd, struct epoll_event* _Nonnull __events, int __event_count, const struct timespec* _Nullable __timeout, const sigset_t* _Nullable __mask) __INTRODUCED_IN(35);

/**
 * Like epoll_pwait2() but using a 64-bit signal mask even on 32-bit systems.
 *
 * Available since API level 35.
 */
int epoll_pwait2_64(int __epoll_fd, struct epoll_event* _Nonnull __events, int __event_count, const struct timespec* _Nullable __timeout, const sigset64_t* _Nullable __mask) __INTRODUCED_IN(35);
#endif /* __BIONIC_AVAILABILITY_GUARD(35) */


__END_DECLS
