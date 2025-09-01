/*
 * Copyright (C) 2018 The Android Open Source Project
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

#include <limits.h>
#include <linux/signal.h>
#include <sys/types.h>

/**
 * The highest kernel-supported signal number, plus one.
 *
 * In theory this is useful for declaring an array with an entry for each signal.
 * In practice, that's less useful than it seems because of the real-time
 * signals and the reserved signals,
 * and the sig2str() and str2sig() functions cover the most common use case
 * of translating between signal numbers and signal names.
 *
 * Note also that although sigset_t and sigset64_t are the same type on LP64,
 * on ILP32 only sigset64_t is large enough to refer to the upper 32 signals.
 * NSIG does _not_ tell you anything about what can be used with sigset_t.
 *
 * See the
 * (32-bit ABI bugs)[https://android.googlesource.com/platform/bionic/+/main/docs/32-bit-abi.md#is-too-small-for-real_time-signals]
 * documentation.
 */
#define NSIG 65
/** A traditional alternative name for NSIG. */
#define _NSIG 65

/*
 * We rewrite the kernel's _NSIG to _KERNEL__NSIG
 * (because the kernel values are off by one from the userspace values),
 * but the kernel <asm/signal.h> headers define SIGRTMAX in terms of
 * _KERNEL__NSIG (or _NSIG, in the original kernel source),
 * so we need to provide a definition here.
 * (Ideally our uapi header rewriter would just hard-code _KERNEL__NSIG to 64.)
 */
#ifndef _KERNEL__NSIG
#define _KERNEL__NSIG 64
#endif

typedef int sig_atomic_t;

typedef __sighandler_t sig_t; /* BSD compatibility. */
typedef __sighandler_t sighandler_t; /* glibc compatibility. */

#if defined(__LP64__)
/**
 * The kernel LP64 sigset_t is large enough to support all signals;
 * this typedef is just for source compatibility with code that uses
 * real-time signals on ILP32.
 *
 * See the
 * (32-bit ABI bugs)[https://android.googlesource.com/platform/bionic/+/main/docs/32-bit-abi.md#is-too-small-for-real_time-signals]
 * documentation.
 */
typedef sigset_t sigset64_t;
#else
/**
 * The ILP32 sigset_t is only 32 bits, so we need a 64-bit sigset64_t
 * and associated functions to be able to support the real-time signals.
 *
 * See the
 * (32-bit ABI bugs)[https://android.googlesource.com/platform/bionic/+/main/docs/32-bit-abi.md#is-too-small-for-real_time-signals]
 * documentation.
 */
typedef struct { unsigned long __bits[64/(8*sizeof(long))]; } sigset64_t;
#endif

/* The kernel's struct sigaction doesn't match the POSIX one,
 * so we define struct sigaction ourselves. */

#if defined(__LP64__)

#define __SIGACTION_BODY \
  int sa_flags; \
  union { \
    sighandler_t sa_handler; \
    void (*sa_sigaction)(int, struct siginfo*, void*); \
  }; \
  sigset_t sa_mask; \
  void (*sa_restorer)(void); \

/**
 * Used with sigaction().
 *
 * On LP64, this supports all signals including real-time signals.
 * On ILP32, this only supports the first 32 signals.
 *
 * See the
 * (32-bit ABI bugs)[https://android.googlesource.com/platform/bionic/+/main/docs/32-bit-abi.md#is-too-small-for-real_time-signals]
 * documentation.
 */
struct sigaction { __SIGACTION_BODY };
/**
 * Used with sigaction64().
 *
 * On LP64, a synonym for struct sigaction for source compatibility with ILP32.
 * On ILP32, this is needed to support all signals including real-time signals
 * because struct sigaction only supports the first 32 signals.
 *
 * See the
 * (32-bit ABI bugs)[https://android.googlesource.com/platform/bionic/+/main/docs/32-bit-abi.md#is-too-small-for-real_time-signals]
 * documentation.
 */
struct sigaction64 { __SIGACTION_BODY };

#undef __SIGACTION_BODY

#else

/* The arm32 kernel headers pollute the namespace with these,
 * but our header scrubber doesn't know how to remove #defines. */
#undef sa_handler
#undef sa_sigaction

/**
 * Used with sigaction().
 *
 * On LP64, this supports all signals including real-time signals.
 * On ILP32, this only supports the first 32 signals.
 *
 * See the
 * (32-bit ABI bugs)[https://android.googlesource.com/platform/bionic/+/main/docs/32-bit-abi.md#is-too-small-for-real_time-signals]
 * documentation.
 */
struct sigaction {
  union {
    sighandler_t sa_handler;
    void (*sa_sigaction)(int, struct siginfo*, void*);
  };
  sigset_t sa_mask;
  int sa_flags;
  void (*sa_restorer)(void);
};

/**
 * Used with sigaction64().
 *
 * On LP64, a synonym for struct sigaction for source compatibility with ILP32.
 * On ILP32, this is needed to support all signals including real-time signals
 * because struct sigaction only supports the first 32 signals.
 *
 * See the
 * (32-bit ABI bugs)[https://android.googlesource.com/platform/bionic/+/main/docs/32-bit-abi.md#is-too-small-for-real_time-signals]
 * documentation.
 */
struct sigaction64 {
  union {
    sighandler_t sa_handler;
    void (*sa_sigaction)(int, struct siginfo*, void*);
  };
  int sa_flags;
  void (*sa_restorer)(void);
  sigset64_t sa_mask;
};

#endif
