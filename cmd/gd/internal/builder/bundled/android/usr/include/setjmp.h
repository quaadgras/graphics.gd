/*	$OpenBSD: setjmp.h,v 1.5 2005/12/13 00:35:22 millert Exp $	*/
/*	$NetBSD: setjmp.h,v 1.11 1994/12/20 10:35:44 cgd Exp $	*/

/*-
 * Copyright (c) 1990, 1993
 *	The Regents of the University of California.  All rights reserved.
 * (c) UNIX System Laboratories, Inc.
 * All or some portions of this file are derived from material licensed
 * to the University of California by American Telephone and Telegraph
 * Co. or Unix System Laboratories, Inc. and are reproduced herein with
 * the permission of UNIX System Laboratories, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)setjmp.h	8.2 (Berkeley) 1/21/94
 */

#pragma once

/**
 * @file setjmp.h
 * @brief Non-local jumps.
 */

#include <sys/cdefs.h>

#if defined(__aarch64__)
/**
 * The size in words of an arm64 jmp_buf. Room for callee-saved registers,
 * including floating point, stack pointer and program counter, various
 * internal implementation details, and leaving some free space.
 *
 * Coincidentally matches OpenBSD, though they also save/restore the
 * floating point status register too.
 */
#define _JBLEN 32
#elif defined(__arm__)
/** The size in words of an arm32 jmp_buf. Inherited from OpenBSD. */
#define _JBLEN 64
#elif defined(__i386__)
/** The size in words of an x86 jmp_buf. Inherited from OpenBSD. */
#define _JBLEN 10
#elif defined(__riscv)
/**
 * The size in words of a riscv64 jmp_buf. Room for callee-saved registers,
 * including floating point, stack pointer and program counter, various
 * internal implementation details, and leaving lots of free space.
 *
 * Deliberately very large given the uncertainty around the final form of
 * hardware shadow stack, and the fact that x86-64 glibc needed to steal
 * space from their enormous sigset_t (which we don't have) to be able to
 * implement the CET shadow stack.
 */
#define _JBLEN 64
#elif defined(__x86_64__)
/** The size in words of an x86-64 jmp_buf. Inherited from OpenBSD. */
#define _JBLEN 11
#endif

/** The type of the buffer used by sigsetjmp()/siglongjmp(). */
typedef long sigjmp_buf[_JBLEN + 1];

/** The type of the buffer used by setjmp()/longjmp(). */
typedef long jmp_buf[_JBLEN];

#undef _JBLEN

__BEGIN_DECLS

/**
 * Equivalent to sigsetjmp() with the second argument 0, so that the signal
 * mask is not saved.
 */
int _setjmp(jmp_buf __env) __returns_twice;

/** Equivalent to siglongjmp(). */
__noreturn void _longjmp(jmp_buf __env, int __value);

/**
 * Equivalent to sigsetjmp() with the second argument 1, so that the signal
 * mask is saved.
 */
int setjmp(jmp_buf __env) __returns_twice;

/** C11 says setjmp() must be a macro, but Android already had a function. */
#define setjmp(__env) setjmp(__env)

/** Equivalent to siglongjmp(). */
__noreturn void longjmp(jmp_buf __env, int __value);

/**
 * [sigsetjmp(3)](https://man7.org/linux/man-pages/man3/sigsetjmp.3.html)
 * sets the target of a future siglongjmp() call, saving or not saving the
 * current signal mask based on the second argument.
 *
 * Returns 0 when first called, and returns the value passed to siglongjmp()
 * when returning here as a result of a siglongjmp() call.
 */
int sigsetjmp(sigjmp_buf __env, int __save_signal_mask) __returns_twice;

/**
 * [siglongjmp(3)](https://man7.org/linux/man-pages/man3/siglongjmp.3.html)
 * transfers control back to the site of the sigsetjmp() call that initialized
 * the given jump buffer, returning the given value.
 *
 * Does not return.
 */
__noreturn void siglongjmp(sigjmp_buf __env, int __value);

__END_DECLS
