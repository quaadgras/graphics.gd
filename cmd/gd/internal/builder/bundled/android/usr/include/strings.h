/*	$NetBSD: strings.h,v 1.10 2005/02/03 04:39:32 perry Exp $	*/

/*-
 * Copyright (c) 1998 The NetBSD Foundation, Inc.
 * All rights reserved.
 *
 * This code is derived from software contributed to The NetBSD Foundation
 * by Klaus Klein.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *        This product includes software developed by the NetBSD
 *        Foundation, Inc. and its contributors.
 * 4. Neither the name of The NetBSD Foundation nor the names of its
 *    contributors may be used to endorse or promote products derived
 *    from this software without specific prior written permission.
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
 * @file strings.h
 * @brief Extra string functions.
 */

#include <sys/cdefs.h>

#include <sys/types.h>
#include <xlocale.h>

#include <bits/strcasecmp.h>

#if !defined(__BIONIC_STRINGS_INLINE)
#define __BIONIC_STRINGS_INLINE static __inline
#endif

#undef ffs
#undef ffsl
#undef ffsll

__BEGIN_DECLS

/** Deprecated. Use memmove() instead. */
#define bcopy(b1, b2, len) __bionic_bcopy((b1), (b2), (len))
static __inline __always_inline void __bionic_bcopy(const void* _Nonnull b1, void* _Nonnull b2, size_t len) {
  __builtin_memmove(b2, b1, len);
}

/** Deprecated. Use memset() instead. */
#define bzero(b, len) __bionic_bzero((b), (len))
static __inline __always_inline void __bionic_bzero(void* _Nonnull b, size_t len) {
  __builtin_memset(b, 0, len);
}

/**
 * [ffs(3)](https://man7.org/linux/man-pages/man3/ffs.3.html) finds the
 * first set bit in `__n`.
 *
 * Returns 0 if no bit is set, or the index of the lowest set bit (counting
 * from 1) otherwise.
 */
__BIONIC_STRINGS_INLINE int ffs(int __n) {
  return __builtin_ffs(__n);
}

/**
 * [ffsl(3)](https://man7.org/linux/man-pages/man3/ffsl.3.html) finds the
 * first set bit in `__n`.
 *
 * Returns 0 if no bit is set, or the index of the lowest set bit (counting
 * from 1) otherwise.
 */
__BIONIC_STRINGS_INLINE int ffsl(long __n) {
  return __builtin_ffsl(__n);
}

/**
 * [ffsll(3)](https://man7.org/linux/man-pages/man3/ffsll.3.html) finds the
 * first set bit in `__n`.
 *
 * Returns 0 if no bit is set, or the index of the lowest set bit (counting
 * from 1) otherwise.
 */
__BIONIC_STRINGS_INLINE int ffsll(long long __n) {
  return __builtin_ffsll(__n);
}

#if defined(__BIONIC_INCLUDE_FORTIFY_HEADERS)
#include <bits/fortify/strings.h>
#endif

__END_DECLS
