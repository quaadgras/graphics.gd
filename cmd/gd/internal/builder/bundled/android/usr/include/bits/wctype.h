/*
 * Copyright (C) 2016 The Android Open Source Project
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

#ifndef _BITS_WCTYPE_H_
#define _BITS_WCTYPE_H_

#include <sys/cdefs.h>

__BEGIN_DECLS

typedef __WINT_TYPE__ wint_t;

#define WEOF __BIONIC_CAST(static_cast, wint_t, -1)

int iswalnum(wint_t __wc);
int iswalpha(wint_t __wc);
int iswblank(wint_t __wc);
int iswcntrl(wint_t __wc);
int iswdigit(wint_t __wc);
int iswgraph(wint_t __wc);
int iswlower(wint_t __wc);
int iswprint(wint_t __wc);
int iswpunct(wint_t __wc);
int iswspace(wint_t __wc);
int iswupper(wint_t __wc);
int iswxdigit(wint_t __wc);

wint_t towlower(wint_t __wc);
wint_t towupper(wint_t __wc);

typedef long wctype_t;
wctype_t wctype(const char* _Nonnull __name);
int iswctype(wint_t __wc, wctype_t __type);

typedef const void* wctrans_t;

#if __BIONIC_AVAILABILITY_GUARD(26)
wint_t towctrans(wint_t __wc, wctrans_t _Nonnull __transform) __INTRODUCED_IN(26);
wctrans_t _Nullable wctrans(const char* _Nonnull __name) __INTRODUCED_IN(26);
#endif /* __BIONIC_AVAILABILITY_GUARD(26) */


__END_DECLS

#endif
