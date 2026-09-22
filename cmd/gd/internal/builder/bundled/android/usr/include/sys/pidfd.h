/*
 * Copyright (C) 2021 The Android Open Source Project
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
 * @file sys/pidfd.h
 * @brief File descriptors representing processes.
 */

#include <sys/cdefs.h>
#include <sys/types.h>

#include <bits/signal_types.h>

__BEGIN_DECLS

/**
 * [pidfd_open(2)](https://man7.org/linux/man-pages/man2/pidfd_open.2.html)
 * opens a file descriptor that refers to a process. This file descriptor will
 * have the close-on-exec flag set by default.
 *
 * Returns a new file descriptor on success and returns -1 and sets `errno` on
 * failure.
 *
 * Available since API level 31.
 */

#if __BIONIC_AVAILABILITY_GUARD(31)
int pidfd_open(pid_t __pid, unsigned int __flags) __INTRODUCED_IN(31);

/**
 * [pidfd_getfd(2)](https://man7.org/linux/man-pages/man2/pidfd_getfd.2.html)
 * dups a file descriptor from another process. This file descriptor will have
 * the close-on-exec flag set by default.
 *
 * Returns a new file descriptor on success and returns -1 and sets `errno` on
 * failure.
 *
 * Available since API level 31.
 */
int pidfd_getfd(int __pidfd, int __targetfd, unsigned int __flags) __INTRODUCED_IN(31);

/**
 * [pidfd_send_signal(2)](https://man7.org/linux/man-pages/man2/pidfd_send_signal.2.html)
 * sends a signal to another process.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 *
 * Available since API level 31.
 */
int pidfd_send_signal(int __pidfd, int __sig, siginfo_t * _Nullable __info, unsigned int __flags) __INTRODUCED_IN(31);
#endif /* __BIONIC_AVAILABILITY_GUARD(31) */


__END_DECLS
