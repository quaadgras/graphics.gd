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

#ifndef _STDLIB_H
#define _STDLIB_H

#include <sys/cdefs.h>

#include <alloca.h>
#include <bits/wait.h>
#include <malloc.h>
#include <stddef.h>
#include <xlocale.h>

__BEGIN_DECLS

#define EXIT_FAILURE 1
#define EXIT_SUCCESS 0

__noreturn void abort(void) __attribute__((__nomerge__));
__noreturn void exit(int __status);
__noreturn void _Exit(int __status);

int atexit(void (* _Nonnull __fn)(void));

int at_quick_exit(void (* _Nonnull __fn)(void));
void quick_exit(int __status) __noreturn;

char* _Nullable getenv(const char* _Nonnull __name);
int putenv(char* _Nonnull __assignment);
int setenv(const char* _Nonnull __name, const char* _Nonnull __value, int __overwrite);
int unsetenv(const char* _Nonnull __name);
int clearenv(void);

char* _Nullable mkdtemp(char* _Nonnull __template);
char* _Nullable mktemp(char* _Nonnull __template) __attribute__((__deprecated__("mktemp is unsafe, use mkstemp or tmpfile instead")));


#if __BIONIC_AVAILABILITY_GUARD(23)
int mkostemp64(char* _Nonnull __template, int __flags) __INTRODUCED_IN(23);
int mkostemp(char* _Nonnull __template, int __flags) __INTRODUCED_IN(23);
int mkostemps64(char* _Nonnull __template, int __suffix_length, int __flags) __INTRODUCED_IN(23);
int mkostemps(char* _Nonnull __template, int __suffix_length, int __flags) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */

int mkstemp64(char* _Nonnull __template);
int mkstemp(char* _Nonnull __template);

#if __BIONIC_AVAILABILITY_GUARD(23)
int mkstemps64(char* _Nonnull __template, int __flags) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */

int mkstemps(char* _Nonnull __template, int __flags);

int posix_memalign(void* _Nullable * _Nullable __memptr, size_t __alignment, size_t __size);

/**
 * [aligned_alloc(3)](https://man7.org/linux/man-pages/man3/aligned_alloc.3.html)
 * allocates the given number of bytes with the given alignment.
 *
 * Returns a pointer to the allocated memory on success and returns a null
 * pointer and sets `errno` on failure.
 *
 * Available since API level 28.
 */

#if __BIONIC_AVAILABILITY_GUARD(28)
__nodiscard void* _Nullable aligned_alloc(size_t __alignment, size_t __size) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


__nodiscard char* _Nullable realpath(const char* _Nonnull __path, char* _Nullable __resolved);

/**
 * [system(3)](https://man7.org/linux/man-pages/man3/system.3.html) executes
 * the given command in a new shell process.
 *
 * On Android, the special case of `system(NULL)` always returns 1,
 * as specified by POSIX. Passing `NULL` to determine whether or
 * not a shell is available is not portable. Callers should just try
 * the command they actually want to run, since there are many reasons
 * why it might fail, both temporarily (for lack of resources, say)
 * or permanently (for lack of permission, say).
 *
 * Returns -1 and sets errno if process creation fails; returns a
 * [waitpid(2)](https://man7.org/linux/man-pages/man2/waitpid.2.html)
 * status otherwise.
 */
int system(const char* _Nonnull __command);

/**
 * [bsearch(3)](https://man7.org/linux/man-pages/man3/bsearch.3.html) searches
 * a sorted array.
 *
 * Returns a pointer to a matching item on success,
 * or NULL if no matching item is found.
 */
__nodiscard void* _Nullable bsearch(const void* _Nonnull __key, const void* _Nullable __base, size_t __nmemb, size_t __size, int (* _Nonnull __comparator)(const void* _Nonnull __lhs, const void* _Nonnull __rhs));

/**
 * [qsort(3)](https://man7.org/linux/man-pages/man3/qsort.3.html) sorts an array
 * of n elements each of the given size, using the given comparator.
 */
void qsort(void* _Nullable __array, size_t __n, size_t __size, int (* _Nonnull __comparator)(const void* _Nullable __lhs, const void* _Nullable __rhs));

/**
 * [qsort_r(3)](https://man7.org/linux/man-pages/man3/qsort_r.3.html) sorts an
 * array of n elements each of the given size, using the given comparator,
 * and passing the given context argument to the comparator.
 *
 * Available since API level 36.
 */

#if __BIONIC_AVAILABILITY_GUARD(36)
void qsort_r(void* _Nullable __array, size_t __n, size_t __size, int (* _Nonnull __comparator)(const void* _Nullable __lhs, const void* _Nullable __rhs, void* _Nullable __context), void* _Nullable __context) __INTRODUCED_IN(36);
#endif /* __BIONIC_AVAILABILITY_GUARD(36) */


uint32_t arc4random(void);
uint32_t arc4random_uniform(uint32_t __upper_bound);
void arc4random_buf(void* _Nonnull __buf, size_t __n);

#define RAND_MAX 0x7fffffff

int rand_r(unsigned int* _Nonnull __seed_ptr);

double drand48(void);
double erand48(unsigned short __xsubi[_Nonnull 3]);
long jrand48(unsigned short __xsubi[_Nonnull 3]);

#if __BIONIC_AVAILABILITY_GUARD(23)
void lcong48(unsigned short __param[_Nonnull 7]) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */

long lrand48(void);
long mrand48(void);
long nrand48(unsigned short __xsubi[_Nonnull 3]);
unsigned short* _Nonnull seed48(unsigned short __seed16v[_Nonnull 3]);
void srand48(long __seed);

char* _Nullable initstate(unsigned int __seed, char* _Nonnull __state, size_t __n);
char* _Nullable setstate(char* _Nonnull __state);

int getpt(void);
int posix_openpt(int __flags);
char* _Nullable ptsname(int __fd);
int ptsname_r(int __fd, char* _Nonnull __buf, size_t __n);
int unlockpt(int __fd);


#if __BIONIC_AVAILABILITY_GUARD(26)
int getsubopt(char* _Nonnull * _Nonnull __option, char* _Nonnull const* _Nonnull __tokens, char* _Nullable * _Nonnull __value_ptr) __INTRODUCED_IN(26);
#endif /* __BIONIC_AVAILABILITY_GUARD(26) */


typedef struct {
  int quot;
  int rem;
} div_t;

div_t div(int __numerator, int __denominator) __attribute_const__;

typedef struct {
  long int quot;
  long int rem;
} ldiv_t;

ldiv_t ldiv(long __numerator, long __denominator) __attribute_const__;

typedef struct {
  long long int quot;
  long long int rem;
} lldiv_t;

lldiv_t lldiv(long long __numerator, long long __denominator) __attribute_const__;

/**
 * [getloadavg(3)](https://man7.org/linux/man-pages/man3/getloadavg.3.html) queries the
 * number of runnable processes averaged over time. The Linux kernel supports averages
 * over the last 1, 5, and 15 minutes.
 *
 * Returns the number of samples written to `__averages` (at most 3), and returns -1 on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(29)
int getloadavg(double __averages[_Nonnull], int __n) __INTRODUCED_IN(29);
#endif /* __BIONIC_AVAILABILITY_GUARD(29) */


/* BSD compatibility. */
const char* _Nullable getprogname(void);
void setprogname(const char* _Nonnull __name);


#if __BIONIC_AVAILABILITY_GUARD(26)
int mblen(const char* _Nullable __s, size_t __n) __INTRODUCED_IN(26);
#endif /* __BIONIC_AVAILABILITY_GUARD(26) */

size_t mbstowcs(wchar_t* _Nullable __dst, const char* _Nullable __src, size_t __n);
int mbtowc(wchar_t* _Nullable __wc_ptr, const char*  _Nullable __s, size_t __n);
int wctomb(char* _Nullable __dst, wchar_t __wc);

size_t wcstombs(char* _Nullable __dst, const wchar_t* _Nullable __src, size_t __n);

size_t __ctype_get_mb_cur_max(void);
#define MB_CUR_MAX __ctype_get_mb_cur_max()

#if defined(__BIONIC_INCLUDE_FORTIFY_HEADERS)
#include <bits/fortify/stdlib.h>
#endif

int abs(int __x) __attribute_const__;
long labs(long __x) __attribute_const__;
long long llabs(long long __x) __attribute_const__;

int rand(void);
void srand(unsigned int __seed);
long random(void);
void srandom(unsigned int __seed);
int grantpt(int __fd);

/**
 * [atof(3)](https://man7.org/linux/man-pages/man3/atof.3.html) converts a
 * string to a double.
 *
 * Returns the double; use strtof() or strtod() if you need to detect errors.
 */
double atof(const char* _Nonnull __s) __attribute_pure__;

/**
 * [atoi(3)](https://man7.org/linux/man-pages/man3/atoi.3.html) converts a
 * string to an int.
 *
 * Returns the int or 0 on error; use strtol() if you need to detect errors.
 */
int atoi(const char* _Nonnull __s) __attribute_pure__;

/**
 * [atol(3)](https://man7.org/linux/man-pages/man3/atol.3.html) converts a
 * string to a long.
 *
 * Returns the long or 0 on error; use strtol() if you need to detect errors.
 */
long atol(const char* _Nonnull __s) __attribute_pure__;

/**
 * [atoll(3)](https://man7.org/linux/man-pages/man3/atoll.3.html) converts a
 * string to a long long.
 *
 * Returns the long long or 0 on error; use strtol() if you need to detect errors.
 */
long long atoll(const char* _Nonnull __s) __attribute_pure__;

/**
 * [strtol(3)](https://man7.org/linux/man-pages/man3/strtol.3.html) converts a
 * string to a long.
 *
 * Returns the long.
 * `__end_ptr` is set to the last character in `__s` that was converted.
 * errno is set to ERANGE if the result overflowed or underflowed.
 */
long strtol(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, int __base);

/** Equivalent to strtol() on Android. */
long strtol_l(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, int, locale_t _Nonnull __l) __RENAME(strtol);

/**
 * [strtoll(3)](https://man7.org/linux/man-pages/man3/strtoll.3.html) converts a
 * string to a long long.
 *
 * Returns the long long.
 * `__end_ptr` is set to the last character in `__s` that was converted.
 * errno is set to ERANGE if the result overflowed or underflowed.
 */
long long strtoll(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, int __base);

/** Equivalent to strtoll() on Android. */
long long strtoll_l(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, int __base, locale_t _Nonnull __l);

/**
 * [strtoul(3)](https://man7.org/linux/man-pages/man3/strtoul.3.html) converts a
 * string to an unsigned long.
 *
 * Returns the unsigned long.
 * `__end_ptr` is set to the last character in `__s` that was converted.
 * errno is set to ERANGE if the result overflowed or underflowed.
 */
unsigned long strtoul(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, int __base);

/** Equivalent to strtoul() on Android. */
unsigned long strtoul_l(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, int __base, locale_t _Nonnull __l) __RENAME(strtoul);

/**
 * [strtoull(3)](https://man7.org/linux/man-pages/man3/strtoull.3.html) converts a
 * string to an unsigned long long.
 *
 * Returns the unsigned long long.
 * `__end_ptr` is set to the last character in `__s` that was converted.
 * errno is set to ERANGE if the result overflowed or underflowed.
 */
unsigned long long strtoull(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, int __base);

/** Equivalent to strtoull() on Android. */
unsigned long long strtoull_l(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, int __base, locale_t _Nonnull __l);

/**
 * [strtof(3)](https://man7.org/linux/man-pages/man3/strtof.3.html) converts a
 * string to a float.
 *
 * Returns the float.
 * `__end_ptr` is set to the last character in `__s` that was converted.
 * errno is set to ERANGE if the result overflowed or underflowed.
 */
float strtof(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr);

/**
 * [strtod(3)](https://man7.org/linux/man-pages/man3/strtod.3.html) converts a
 * string to a double.
 *
 * Returns the double.
 * `__end_ptr` is set to the last character in `__s` that was converted.
 * errno is set to ERANGE if the result overflowed or underflowed.
 */
double strtod(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr);

/**
 * [strtold(3)](https://man7.org/linux/man-pages/man3/strtold.3.html) converts a
 * string to a long double.
 *
 * Returns the long double.
 * `__end_ptr` is set to the last character in `__s` that was converted.
 * errno is set to ERANGE if the result overflowed or underflowed.
 */
long double strtold(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr);

/** Equivalent to strtold() on Android. */
long double strtold_l(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, locale_t _Nonnull __l);

#if __ANDROID_API__ >= 26
/** Equivalent to strtod() on Android. */
double strtod_l(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, locale_t _Nonnull __l) __INTRODUCED_IN(26);
/** Equivalent to strtof() on Android. */
float strtof_l(const char* _Nonnull __s, char* _Nullable * _Nullable __end_ptr, locale_t _Nonnull __l) __INTRODUCED_IN(26);
#else
// Implemented as static inlines before 26.
#endif

__END_DECLS

#include <android/legacy_stdlib_inlines.h>

#endif /* _STDLIB_H */
