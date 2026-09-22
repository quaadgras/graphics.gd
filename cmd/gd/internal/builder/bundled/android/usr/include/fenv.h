/*  $OpenBSD: fenv.h,v 1.2 2011/05/25 21:46:49 martynas Exp $ */
/*  $NetBSD: fenv.h,v 1.2.4.1 2011/02/08 16:18:55 bouyer Exp $  */

/*
 * Copyright (c) 2010 The NetBSD Foundation, Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE NETBSD FOUNDATION, INC. AND CONTRIBUTORS
 * ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED
 * TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE FOUNDATION OR CONTRIBUTORS
 * BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

#pragma once

/**
 * @file fenv.h
 * @brief Floating-point environment.
 */

#include <sys/cdefs.h>

#if defined(__aarch64__) || defined(__arm__)
#include <bits/fenv_arm.h>
#elif defined(__i386__)
#include <bits/fenv_x86.h>
#elif defined(__riscv)
#include <bits/fenv_riscv64.h>
#elif defined(__x86_64__)
#include <bits/fenv_x86_64.h>
#endif

__BEGIN_DECLS

/**
 * [feclearexcept(3)](https://man7.org/linux/man-pages/man3/feclearexcept.3.html)
 * clears the given `exceptions` in hardware.
 *
 * Returns 0 on success, and returns non-zero on failure.
 */
int feclearexcept(int __exceptions);

/**
 * [fegetexceptflag(3)](https://man7.org/linux/man-pages/man3/fegetexceptflag.3.html)
 * copies the state of the given `exceptions` from hardware into `*flag_ptr`.
 * See fesetexceptflag().
 *
 * Returns 0 on success, and returns non-zero on failure.
 */
int fegetexceptflag(fexcept_t* _Nonnull __flag_ptr, int __exceptions);

/**
 * [feraiseexcept(3)](https://man7.org/linux/man-pages/man3/feraiseexcept.3.html)
 * raises the given `exceptions` in hardware.
 *
 * Returns 0 on success, and returns non-zero on failure.
 */
int feraiseexcept(int __exceptions);

/**
 * [fesetexceptflag(3)](https://man7.org/linux/man-pages/man3/fesetexceptflag.3.html)
 * copies the state of the given `exceptions` from `*flag_ptr` into hardware.
 * See fesetexceptflag().
 *
 * Returns 0 on success, and returns non-zero on failure.
 */
int fesetexceptflag(const fexcept_t* _Nonnull __flag_ptr, int __exceptions);

/**
 * [fetestexcept(3)](https://man7.org/linux/man-pages/man3/fetestexcept.3.html)
 * tests whether the given `exceptions` are set in hardware.
 *
 * Returns the currently-set subset of `exceptions`.
 */
int fetestexcept(int __exceptions);

/**
 * [fegetround(3)](https://man7.org/linux/man-pages/man3/fegetround.3.html)
 * returns the current rounding mode.
 *
 * Returns the rounding mode on success, and returns a negative value on failure.
 */
int fegetround(void);

/**
 * [fesetround(3)](https://man7.org/linux/man-pages/man3/fesetround.3.html)
 * sets the current rounding mode.
 *
 * Returns 0 on success, and returns non-zero on failure.
 */
int fesetround(int __rounding_mode);

/**
 * [fegetenv(3)](https://man7.org/linux/man-pages/man3/fegetenv.3.html)
 * gets the current floating-point environment. See fesetenv().
 *
 * Returns 0 on success, and returns non-zero on failure.
 */
int fegetenv(fenv_t* _Nonnull __env);

/**
 * [feholdexcept(3)](https://man7.org/linux/man-pages/man3/feholdexcept.3.html)
 * gets the current floating-point environment, clears the status flags, and
 * ignores floating point exceptions. See fesetenv()/feupdateenv().
 *
 * Returns 0 on success, and returns non-zero on failure.
 */
int feholdexcept(fenv_t* _Nonnull __env);

/**
 * [fesetenv(3)](https://man7.org/linux/man-pages/man3/fesetenv.3.html)
 * sets the current floating-point environment. See fegetenv().
 *
 * Returns 0 on success, and returns non-zero on failure.
 */
int fesetenv(const fenv_t* _Nonnull __env);

/**
 * [feupdateenv(3)](https://man7.org/linux/man-pages/man3/feupdateenv.3.html)
 * sets the current floating-point environment to `*env` but with currently-raised
 * exceptions still raised. See fesetenv().
 *
 * Returns 0 on success, and returns non-zero on failure.
 */
int feupdateenv(const fenv_t* _Nonnull __env);

/**
 * [feenableexcept(3)](https://man7.org/linux/man-pages/man3/feenableexcept.3.html)
 * sets the given `exceptions` to trap, if the hardware supports it. This is not
 * generally useful on Android, because only x86/x86-64 can trap.
 *
 * Returns the previous set of enabled exceptions on success, and returns -1 on failure.
 */
int feenableexcept(int __exceptions);

/**
 * [fedisableexcept(3)](https://man7.org/linux/man-pages/man3/fedisableexcept.3.html)
 * sets the given `exceptions` to not trap, if the hardware supports it. This is not
 * generally useful on Android, because only x86/x86-64 can trap.
 *
 * Returns the previous set of enabled exceptions on success, and returns -1 on failure.
 */
int fedisableexcept(int __exceptions);

/**
 * [fegetexcept(3)](https://man7.org/linux/man-pages/man3/fegetexcept.3.html)
 * returns the exceptions that currently trap. This is not generally useful on
 * Android, because only x86/x86-64 can trap.
 *
 * Returns the exceptions that currently trap.
 */
int fegetexcept(void);

/** See FE_DFL_ENV. */
extern const fenv_t __fe_dfl_env;

/**
 * Constant representing the default floating-point environment
 * (that is, the one installed at program startup).
 *
 * It can be used as an argument to the functions that manage the floating-point
 * environment, namely fesetenv() and feupdateenv().
 */
#define FE_DFL_ENV (&__fe_dfl_env)

__END_DECLS
