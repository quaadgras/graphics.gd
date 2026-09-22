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

#ifndef _WCHAR_H_
#define _WCHAR_H_

#include <sys/cdefs.h>
#include <stdio.h>

#include <stdarg.h>
#include <stddef.h>
#include <time.h>
#include <xlocale.h>

#include <bits/bionic_multibyte_result.h>
#include <bits/mbstate_t.h>
#include <bits/wchar_limits.h>
#include <bits/wctype.h>

__BEGIN_DECLS

wint_t btowc(int __ch);
int fwprintf(FILE* _Nonnull __fp, const wchar_t* _Nonnull __fmt, ...);
int fwscanf(FILE* _Nonnull __fp, const wchar_t* _Nonnull __fmt, ...);
wint_t fgetwc(FILE* _Nonnull __fp);
wchar_t* _Nullable fgetws(wchar_t* _Nonnull __buf, int __size, FILE* _Nonnull __fp);
wint_t fputwc(wchar_t __wc, FILE* _Nonnull __fp);
int fputws(const wchar_t* _Nonnull __s, FILE* _Nonnull __fp);
int fwide(FILE* _Nonnull __fp, int __mode);
wint_t getwc(FILE* _Nonnull __fp);
wint_t getwchar(void);
int mbsinit(const mbstate_t* _Nullable __ps);
size_t mbrlen(const char* _Nullable __s, size_t __n, mbstate_t* _Nullable __ps);
size_t mbrtowc(wchar_t* _Nullable __buf, const char* _Nullable __s, size_t __n, mbstate_t* _Nullable __ps);
size_t mbsrtowcs(wchar_t* _Nullable __dst, const char* _Nullable * _Nonnull __src, size_t __dst_n, mbstate_t* _Nullable __ps);
size_t mbsrtowcs_l(wchar_t* _Nullable __dst, const char* _Nullable * _Nonnull __src, size_t __dst_n, mbstate_t* _Nullable __ps, locale_t _Nonnull __l) __RENAME(mbsrtowcs);
size_t mbsnrtowcs(wchar_t* _Nullable __dst, const char* _Nullable * _Nullable  __src, size_t __src_n, size_t __dst_n, mbstate_t* _Nullable __ps);
wint_t putwc(wchar_t __wc, FILE* _Nonnull __fp);
wint_t putwchar(wchar_t __wc);
int swprintf(wchar_t* _Nonnull __buf, size_t __n, const wchar_t* _Nonnull __fmt, ...);
int swscanf(const wchar_t* _Nonnull __s, const wchar_t* _Nonnull __fmt, ...);
wint_t ungetwc(wint_t __wc, FILE* _Nonnull __fp);
int vfwprintf(FILE* _Nonnull __fp, const wchar_t* _Nonnull __fmt, va_list __args);
int vfwscanf(FILE* _Nonnull __fp, const wchar_t* _Nonnull __fmt, va_list __args);
int vswprintf(wchar_t* _Nonnull __buf, size_t __n, const wchar_t* _Nonnull __fmt, va_list __args);
int vswscanf(const wchar_t* _Nonnull __s, const wchar_t* _Nonnull __fmt, va_list __args);
int vwprintf(const wchar_t* _Nonnull __fmt, va_list __args);
int vwscanf(const wchar_t* _Nonnull __fmt, va_list __args);
wchar_t* _Nonnull wcpcpy(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src);
wchar_t* _Nonnull wcpncpy(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src, size_t __n);
size_t wcrtomb(char* _Nullable __buf, wchar_t __wc, mbstate_t* _Nullable __ps);
int wcscasecmp(const wchar_t* _Nonnull __lhs, const wchar_t* _Nonnull __rhs);

#if __BIONIC_AVAILABILITY_GUARD(23)
int wcscasecmp_l(const wchar_t* _Nonnull __lhs, const wchar_t* _Nonnull __rhs, locale_t _Nonnull __l) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */

wchar_t* _Nonnull wcscat(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src);
wchar_t* _Nullable wcschr(const wchar_t * _Nonnull __s, wchar_t __wc);
int wcscmp(const wchar_t* _Nonnull __lhs, const wchar_t* _Nonnull __rhs);
int wcscoll(const wchar_t* _Nonnull __lhs, const wchar_t* _Nonnull __rhs);
wchar_t* _Nonnull wcscpy(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src);
size_t wcscspn(const wchar_t* _Nonnull __s, const wchar_t* _Nonnull __accept);
size_t wcsftime(wchar_t* _Nonnull __buf, size_t __n, const wchar_t* _Nullable __fmt, const struct tm* _Nonnull __tm);

#if __BIONIC_AVAILABILITY_GUARD(28)
size_t wcsftime_l(wchar_t* _Nonnull __buf, size_t __n, const wchar_t* _Nullable __fmt, const struct tm* _Nonnull __tm, locale_t _Nonnull __l) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

size_t wcslen(const wchar_t* _Nonnull __s);
int wcsncasecmp(const wchar_t* _Nonnull __lhs, const wchar_t* _Nonnull __rhs, size_t __n);

#if __BIONIC_AVAILABILITY_GUARD(23)
int wcsncasecmp_l(const wchar_t* _Nonnull __lhs, const wchar_t* _Nonnull __rhs, size_t __n, locale_t _Nonnull __l) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */

wchar_t* _Nonnull wcsncat(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src, size_t __n);
int wcsncmp(const wchar_t* _Nonnull __lhs, const wchar_t* _Nonnull __rhs, size_t __n);
wchar_t* _Nonnull wcsncpy(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src, size_t __n);
size_t wcsnrtombs(char* _Nullable __dst, const wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __src, size_t __src_n, size_t __dst_n, mbstate_t* _Nullable __ps);
wchar_t* _Nullable wcspbrk(const wchar_t* _Nonnull __s, const wchar_t* _Nonnull __accept);
wchar_t* _Nullable wcsrchr(const wchar_t* _Nonnull __s, wchar_t __wc);
size_t wcsrtombs(char* _Nullable __dst, const wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __src, size_t __dst_n, mbstate_t* _Nullable __ps);
size_t wcsrtombs_l(char* _Nullable __dst, const wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __src, size_t __dst_n, mbstate_t* _Nullable __ps, locale_t _Nonnull __l) __RENAME(wcsrtombs);
size_t wcsspn(const wchar_t* _Nonnull __s, const wchar_t* _Nonnull __accept);
wchar_t* _Nullable wcsstr(const wchar_t* _Nonnull __haystack, const wchar_t* _Nonnull __needle);
double wcstod(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr);
double wcstod_l(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr, locale_t _Nonnull __l) __RENAME(wcstod);
float wcstof(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr);
float wcstof_l(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr, locale_t _Nonnull __l) __RENAME(wcstof);
wchar_t* _Nullable wcstok(wchar_t* _Nullable __s, const wchar_t* _Nonnull __delimiter, wchar_t* _Nonnull * _Nonnull __ptr);
long wcstol(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr, int __base);
long wcstol_l(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr, int __base, locale_t _Nonnull __l) __RENAME(wcstol);
long long wcstoll(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr, int __base);
long double wcstold(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr);
unsigned long wcstoul(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr, int __base);
unsigned long wcstoul_l(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr, int __base, locale_t _Nonnull __l) __RENAME(wcstoul);
unsigned long long wcstoull(const wchar_t* _Nonnull __s, wchar_t* __BIONIC_COMPLICATED_NULLNESS * _Nullable __end_ptr, int __base);
int wcswidth(const wchar_t* _Nonnull __s, size_t __n);
size_t wcsxfrm(wchar_t* __BIONIC_COMPLICATED_NULLNESS __dst, const wchar_t* _Nonnull __src, size_t __n);
int wctob(wint_t __wc);
int wcwidth(wchar_t __wc);
wchar_t* _Nullable wmemchr(const wchar_t* _Nonnull __src, wchar_t __wc, size_t __n);
int wmemcmp(const wchar_t* _Nullable __lhs, const wchar_t* _Nullable __rhs, size_t __n);
wchar_t* _Nonnull wmemcpy(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src, size_t __n);
#if defined(__USE_GNU)

#if __BIONIC_AVAILABILITY_GUARD(23)
wchar_t* _Nonnull wmempcpy(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src, size_t __n) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */

#endif
wchar_t* _Nonnull wmemmove(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src, size_t __n);
wchar_t* _Nonnull wmemset(wchar_t* _Nonnull __dst, wchar_t __wc, size_t __n);
int wprintf(const wchar_t* _Nonnull __fmt, ...);
int wscanf(const wchar_t* _Nonnull __fmt, ...);

long long wcstoll_l(const wchar_t* _Nonnull __s, wchar_t* _Nullable * _Nullable __end_ptr, int __base, locale_t _Nonnull __l);
unsigned long long wcstoull_l(const wchar_t* _Nonnull __s, wchar_t* _Nullable * _Nullable __end_ptr, int __base, locale_t _Nonnull __l);
long double wcstold_l(const wchar_t* _Nonnull __s, wchar_t* _Nullable * _Nullable __end_ptr, locale_t _Nonnull __l);

int wcscoll_l(const wchar_t* _Nonnull __lhs, const wchar_t* _Nonnull __rhs, locale_t _Nonnull __l) __attribute_pure__;
size_t wcsxfrm_l(wchar_t* __BIONIC_COMPLICATED_NULLNESS __dst, const wchar_t* _Nonnull __src, size_t __n, locale_t _Nonnull __l);
size_t wcslcat(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src, size_t __n);
size_t wcslcpy(wchar_t* _Nonnull __dst, const wchar_t* _Nonnull __src, size_t __n);


#if __BIONIC_AVAILABILITY_GUARD(23)
FILE* _Nullable open_wmemstream(wchar_t* _Nonnull * _Nonnull __ptr, size_t* _Nonnull  __size_ptr) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */

wchar_t* _Nullable wcsdup(const wchar_t* _Nonnull __s);
size_t wcsnlen(const wchar_t* _Nonnull __s, size_t __n);

__END_DECLS

#endif
