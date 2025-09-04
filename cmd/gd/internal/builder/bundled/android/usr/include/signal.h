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

#ifndef _SIGNAL_H_
#define _SIGNAL_H_

#include <sys/cdefs.h>
#include <sys/types.h>

#include <asm/sigcontext.h>
#include <bits/pthread_types.h>
#include <bits/signal_types.h>
#include <bits/timespec.h>
#include <limits.h>

#include <sys/ucontext.h>
#define __BIONIC_HAVE_UCONTEXT_T

__BEGIN_DECLS

/* The kernel headers define SIG_DFL (0) and SIG_IGN (1) but not SIG_HOLD, since
 * SIG_HOLD is only used by the deprecated SysV signal API.
 */
#define SIG_HOLD __BIONIC_CAST(reinterpret_cast, sighandler_t, 2)

/* We take a few real-time signals for ourselves. May as well use the same names as glibc. */
#define SIGRTMIN (__libc_current_sigrtmin())
#define SIGRTMAX (__libc_current_sigrtmax())
int __libc_current_sigrtmin(void);
int __libc_current_sigrtmax(void);

extern const char* _Nonnull const sys_siglist[_NSIG];
extern const char* _Nonnull const sys_signame[_NSIG]; /* BSD compatibility. */

#define si_timerid si_tid /* glibc compatibility. */

int sigaction(int __signal, const struct sigaction* _Nullable __new_action, struct sigaction* _Nullable __old_action);

#if __BIONIC_AVAILABILITY_GUARD(28)
int sigaction64(int __signal, const struct sigaction64* _Nullable __new_action, struct sigaction64* _Nullable __old_action) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


int siginterrupt(int __signal, int __flag);

sighandler_t _Nonnull signal(int __signal, sighandler_t _Nullable __handler);
int sigaddset(sigset_t* _Nonnull __set, int __signal);

#if __BIONIC_AVAILABILITY_GUARD(28)
int sigaddset64(sigset64_t* _Nonnull __set, int __signal) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

int sigdelset(sigset_t* _Nonnull __set, int __signal);

#if __BIONIC_AVAILABILITY_GUARD(28)
int sigdelset64(sigset64_t* _Nonnull __set, int __signal) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

int sigemptyset(sigset_t* _Nonnull __set);

#if __BIONIC_AVAILABILITY_GUARD(28)
int sigemptyset64(sigset64_t* _Nonnull __set) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

int sigfillset(sigset_t* _Nonnull __set);

#if __BIONIC_AVAILABILITY_GUARD(28)
int sigfillset64(sigset64_t* _Nonnull __set) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

int sigismember(const sigset_t* _Nonnull __set, int __signal);

#if __BIONIC_AVAILABILITY_GUARD(28)
int sigismember64(const sigset64_t* _Nonnull __set, int __signal) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


int sigpending(sigset_t* _Nonnull __set);

#if __BIONIC_AVAILABILITY_GUARD(28)
int sigpending64(sigset64_t* _Nonnull __set) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

int sigprocmask(int __how, const sigset_t* _Nullable __new_set, sigset_t* _Nullable __old_set);

#if __BIONIC_AVAILABILITY_GUARD(28)
int sigprocmask64(int __how, const sigset64_t* _Nullable __new_set, sigset64_t* _Nullable __old_set) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

int sigsuspend(const sigset_t* _Nonnull __mask);

#if __BIONIC_AVAILABILITY_GUARD(28)
int sigsuspend64(const sigset64_t* _Nonnull __mask) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

int sigwait(const sigset_t* _Nonnull __set, int* _Nonnull __signal);

#if __BIONIC_AVAILABILITY_GUARD(28)
int sigwait64(const sigset64_t* _Nonnull __set, int* _Nonnull __signal) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */



#if __BIONIC_AVAILABILITY_GUARD(26)
int sighold(int __signal)
  __attribute__((__deprecated__("use sigprocmask() or pthread_sigmask() instead")))
  __INTRODUCED_IN(26);
int sigignore(int __signal)
  __attribute__((__deprecated__("use sigaction() instead"))) __INTRODUCED_IN(26);
int sigpause(int __signal)
  __attribute__((__deprecated__("use sigsuspend() instead"))) __INTRODUCED_IN(26);
int sigrelse(int __signal)
  __attribute__((__deprecated__("use sigprocmask() or pthread_sigmask() instead")))
  __INTRODUCED_IN(26);
sighandler_t _Nonnull sigset(int __signal, sighandler_t _Nullable __handler)
  __attribute__((__deprecated__("use sigaction() instead"))) __INTRODUCED_IN(26);
#endif /* __BIONIC_AVAILABILITY_GUARD(26) */


int raise(int __signal);
int kill(pid_t __pid, int __signal);
int killpg(int __pgrp, int __signal);
int tgkill(int __tgid, int __tid, int __signal);

int sigaltstack(const stack_t* _Nullable __new_signal_stack, stack_t*  _Nullable __old_signal_stack);

void psiginfo(const siginfo_t* _Nonnull __info, const char* _Nullable __msg);
void psignal(int __signal, const char* _Nullable __msg);

int pthread_kill(pthread_t __pthread, int __signal);
#if defined(__USE_GNU)

#if __BIONIC_AVAILABILITY_GUARD(29)
int pthread_sigqueue(pthread_t __pthread, int __signal, const union sigval __value) __INTRODUCED_IN(29);
#endif /* __BIONIC_AVAILABILITY_GUARD(29) */

#endif

int pthread_sigmask(int __how, const sigset_t* _Nullable __new_set, sigset_t* _Nullable __old_set);

#if __BIONIC_AVAILABILITY_GUARD(28)
int pthread_sigmask64(int __how, const sigset64_t* _Nullable __new_set, sigset64_t* _Nullable __old_set) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */



#if __BIONIC_AVAILABILITY_GUARD(23)
int sigqueue(pid_t __pid, int __signal, const union sigval __value) __INTRODUCED_IN(23);
int sigtimedwait(const sigset_t* _Nonnull __set, siginfo_t* _Nullable __info, const struct timespec* _Nullable __timeout) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


#if __BIONIC_AVAILABILITY_GUARD(28)
int sigtimedwait64(const sigset64_t* _Nonnull __set, siginfo_t* _Nullable __info, const struct timespec* _Nullable __timeout) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


#if __BIONIC_AVAILABILITY_GUARD(23)
int sigwaitinfo(const sigset_t* _Nonnull __set, siginfo_t* _Nullable __info) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


#if __BIONIC_AVAILABILITY_GUARD(28)
int sigwaitinfo64(const sigset64_t* _Nonnull __set, siginfo_t* _Nullable __info) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


/**
 * Buffer size suitable for any call to sig2str().
 */
#define SIG2STR_MAX 32

/**
 * [sig2str(3)](https://man7.org/linux/man-pages/man3/sig2str.3.html)
 * converts the integer corresponding to SIGSEGV (say) into a string
 * like "SEGV" (not including the "SIG" used in the constants).
 * SIG2STR_MAX is a safe size to use for the buffer.
 *
 * Returns 0 on success, and returns -1 _without_ setting errno otherwise.
 *
 * Available since API level 36.
 */

#if __BIONIC_AVAILABILITY_GUARD(36)
int sig2str(int __signal, char* _Nonnull __buf) __INTRODUCED_IN(36);

/**
 * [str2sig(3)](https://man7.org/linux/man-pages/man3/str2sig.3.html)
 * converts a string like "SEGV" (not including the "SIG" used in the constants)
 * into the integer corresponding to SIGSEGV.
 *
 * Returns 0 on success, and returns -1 _without_ setting errno otherwise.
 *
 * Available since API level 36.
 */
int str2sig(const char* _Nonnull __name, int* _Nonnull __signal) __INTRODUCED_IN(36);
#endif /* __BIONIC_AVAILABILITY_GUARD(36) */


__END_DECLS

#endif
