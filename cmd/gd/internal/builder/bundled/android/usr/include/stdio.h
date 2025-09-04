/*	$OpenBSD: stdio.h,v 1.35 2006/01/13 18:10:09 miod Exp $	*/
/*	$NetBSD: stdio.h,v 1.18 1996/04/25 18:29:21 jtc Exp $	*/

/*-
 * Copyright (c) 1990 The Regents of the University of California.
 * All rights reserved.
 *
 * This code is derived from software contributed to Berkeley by
 * Chris Torek.
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
 *	@(#)stdio.h	5.17 (Berkeley) 6/3/91
 */

#ifndef	_STDIO_H_
#define	_STDIO_H_

#include <sys/cdefs.h>
#include <sys/types.h>

#include <stdarg.h>
#include <stddef.h>

#include <bits/seek_constants.h>

#if __ANDROID_API__ < 24
#include <bits/struct_file.h>
#endif

__BEGIN_DECLS

typedef off_t fpos_t;
typedef off64_t fpos64_t;

struct __sFILE;
typedef struct __sFILE FILE;

#if __ANDROID_API__ >= 23
extern FILE* _Nonnull stdin __INTRODUCED_IN(23);
extern FILE* _Nonnull stdout __INTRODUCED_IN(23);
extern FILE* _Nonnull stderr __INTRODUCED_IN(23);

/* C99 and earlier plus current C++ standards say these must be macros. */
#define stdin stdin
#define stdout stdout
#define stderr stderr
#else
/* Before M the actual symbols for stdin and friends had different names. */
extern FILE __sF[] __REMOVED_IN(23, "Use stdin/stdout/stderr");

#define stdin (&__sF[0])
#define stdout (&__sF[1])
#define stderr (&__sF[2])
#endif

/*
 * The following three definitions are for ANSI C, which took them
 * from System V, which brilliantly took internal interface macros and
 * made them official arguments to setvbuf(), without renaming them.
 * Hence, these ugly _IOxxx names are *supposed* to appear in user code.
 *
 * Although numbered as their counterparts above, the implementation
 * does not rely on this.
 */
#define	_IOFBF	0		/* setvbuf should set fully buffered */
#define	_IOLBF	1		/* setvbuf should set line buffered */
#define	_IONBF	2		/* setvbuf should set unbuffered */

#define	BUFSIZ	1024		/* size of buffer used by setbuf */
#define	EOF	(-1)

/*
 * FOPEN_MAX is a minimum maximum, and is the number of streams that
 * stdio can provide without attempting to allocate further resources
 * (which could fail).  Do not use this for anything.
 */
#define FOPEN_MAX 20
#define FILENAME_MAX 4096

#define L_tmpnam 4096
#define TMP_MAX 308915776

void clearerr(FILE* _Nonnull __fp);
int fclose(FILE* _Nonnull __fp);
__nodiscard int feof(FILE* _Nonnull __fp);
__nodiscard int ferror(FILE* _Nonnull __fp);
int fflush(FILE* _Nullable __fp);
__nodiscard int fgetc(FILE* _Nonnull __fp);
char* _Nullable fgets(char* _Nonnull __buf, int __size, FILE* _Nonnull __fp);
int fprintf(FILE* _Nonnull __fp , const char* _Nonnull __fmt, ...) __printflike(2, 3);
int fputc(int __ch, FILE* _Nonnull __fp);
int fputs(const char* _Nonnull __s, FILE* _Nonnull __fp);
size_t fread(void* _Nonnull __buf, size_t __size, size_t __count, FILE* _Nonnull __fp);
int fscanf(FILE* _Nonnull __fp, const char* _Nonnull __fmt, ...) __scanflike(2, 3);
size_t fwrite(const void* _Nonnull __buf, size_t __size, size_t __count, FILE* _Nonnull __fp);
__nodiscard int getc(FILE* _Nonnull __fp);
__nodiscard int getchar(void);
ssize_t getdelim(char* _Nullable * _Nonnull __line_ptr, size_t* _Nonnull __line_length_ptr, int __delimiter, FILE* _Nonnull __fp);
ssize_t getline(char* _Nullable * _Nonnull __line_ptr, size_t* _Nonnull __line_length_ptr, FILE* _Nonnull __fp);

void perror(const char* _Nullable __msg);
int printf(const char* _Nonnull __fmt, ...) __printflike(1, 2);
int putc(int __ch, FILE* _Nonnull __fp);
int putchar(int __ch);
int puts(const char* _Nonnull __s);
int remove(const char* _Nonnull __path);
void rewind(FILE* _Nonnull __fp);
int scanf(const char* _Nonnull __fmt, ...) __scanflike(1, 2);
void setbuf(FILE* _Nonnull __fp, char* _Nullable __buf);
int setvbuf(FILE* _Nonnull __fp, char* _Nullable __buf, int __mode, size_t __size);
int sscanf(const char* _Nonnull __s, const char* _Nonnull __fmt, ...) __scanflike(2, 3);
int ungetc(int __ch, FILE* _Nonnull __fp);
int vfprintf(FILE* _Nonnull __fp, const char* _Nonnull __fmt, va_list __args) __printflike(2, 0);
int vprintf(const char* _Nonnull __fp, va_list __args) __printflike(1, 0);

int dprintf(int __fd, const char* _Nonnull __fmt, ...) __printflike(2, 3);
int vdprintf(int __fd, const char* _Nonnull __fmt, va_list __args) __printflike(2, 0);

#if (defined(__STDC_VERSION__) && __STDC_VERSION__ < 201112L) || \
    (defined(__cplusplus) && __cplusplus <= 201103L)
char* _Nullable gets(char* _Nonnull __buf) __attribute__((__deprecated__("gets is unsafe, use fgets instead")));
#endif
int sprintf(char* __BIONIC_COMPLICATED_NULLNESS __s, const char* _Nonnull __fmt, ...)
    __printflike(2, 3) __warnattr_strict("sprintf is often misused; please use snprintf");
int vsprintf(char* __BIONIC_COMPLICATED_NULLNESS __s, const char* _Nonnull __fmt, va_list __args)
    __printflike(2, 0) __warnattr_strict("vsprintf is often misused; please use vsnprintf");
char* _Nullable tmpnam(char* _Nullable __s)
    __warnattr("tmpnam is unsafe, use mkstemp or tmpfile instead");
#define P_tmpdir "/tmp/" /* deprecated */
char* _Nullable tempnam(const char* _Nullable __dir, const char* _Nullable __prefix)
    __warnattr("tempnam is unsafe, use mkstemp or tmpfile instead");

/**
 * [rename(2)](https://man7.org/linux/man-pages/man2/rename.2.html) changes
 * the name or location of a file.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int rename(const char* _Nonnull __old_path, const char* _Nonnull __new_path);

/**
 * [renameat(2)](https://man7.org/linux/man-pages/man2/renameat.2.html) changes
 * the name or location of a file, interpreting relative paths using an fd.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int renameat(int __old_dir_fd, const char* _Nonnull __old_path, int __new_dir_fd, const char* _Nonnull __new_path);

#if defined(__USE_GNU)

/**
 * Flag for [renameat2(2)](https://man7.org/linux/man-pages/man2/renameat2.2.html)
 * to fail if the new path already exists.
 */
#define RENAME_NOREPLACE (1<<0)

/**
 * Flag for [renameat2(2)](https://man7.org/linux/man-pages/man2/renameat2.2.html)
 * to atomically exchange the two paths.
 */
#define RENAME_EXCHANGE (1<<1)

/**
 * Flag for [renameat2(2)](https://man7.org/linux/man-pages/man2/renameat2.2.html)
 * to create a union/overlay filesystem object.
 */
#define RENAME_WHITEOUT (1<<2)

/**
 * [renameat2(2)](https://man7.org/linux/man-pages/man2/renameat2.2.html) changes
 * the name or location of a file, interpreting relative paths using an fd,
 * with optional `RENAME_` flags.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(30)
int renameat2(int __old_dir_fd, const char* _Nonnull __old_path, int __new_dir_fd, const char* _Nonnull __new_path, unsigned __flags) __INTRODUCED_IN(30);
#endif /* __BIONIC_AVAILABILITY_GUARD(30) */


#endif

int fseek(FILE* _Nonnull __fp, long __offset, int __whence);
__nodiscard long ftell(FILE* _Nonnull __fp);

/* See https://android.googlesource.com/platform/bionic/+/main/docs/32-bit-abi.md */
#if defined(__USE_FILE_OFFSET64)

#if __BIONIC_AVAILABILITY_GUARD(24)
int fgetpos(FILE* _Nonnull __fp, fpos_t* _Nonnull __pos) __RENAME(fgetpos64) __INTRODUCED_IN(24);
int fsetpos(FILE* _Nonnull __fp, const fpos_t* _Nonnull __pos) __RENAME(fsetpos64) __INTRODUCED_IN(24);
int fseeko(FILE* _Nonnull __fp, off_t __offset, int __whence) __RENAME(fseeko64) __INTRODUCED_IN(24);
__nodiscard off_t ftello(FILE* _Nonnull __fp) __RENAME(ftello64) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */

#  if defined(__USE_BSD)
/* If __read_fn and __write_fn are both nullptr, it will cause EINVAL */

#if __BIONIC_AVAILABILITY_GUARD(24)
__nodiscard FILE* _Nullable funopen(const void* _Nullable __cookie,
              int (* __BIONIC_COMPLICATED_NULLNESS __read_fn)(void* _Nonnull, char* _Nonnull, int),
              int (* __BIONIC_COMPLICATED_NULLNESS __write_fn)(void* _Nonnull, const char* _Nonnull, int),
              fpos_t (* _Nullable __seek_fn)(void* _Nonnull, fpos_t, int),
              int (* _Nullable __close_fn)(void* _Nonnull)) __RENAME(funopen64) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */

#  endif
#else
int fgetpos(FILE* _Nonnull __fp, fpos_t* _Nonnull __pos);
int fsetpos(FILE* _Nonnull __fp, const fpos_t* _Nonnull __pos);
int fseeko(FILE* _Nonnull __fp, off_t __offset, int __whence);
__nodiscard off_t ftello(FILE* _Nonnull __fp);
#  if defined(__USE_BSD)
/* If __read_fn and __write_fn are both nullptr, it will cause EINVAL */
__nodiscard FILE* _Nullable funopen(const void* _Nullable __cookie,
              int (* __BIONIC_COMPLICATED_NULLNESS __read_fn)(void* _Nonnull, char* _Nonnull, int),
              int (* __BIONIC_COMPLICATED_NULLNESS __write_fn)(void* _Nonnull, const char* _Nonnull, int),
              fpos_t (* _Nullable __seek_fn)(void* _Nonnull, fpos_t, int),
              int (* _Nullable __close_fn)(void* _Nonnull));
#  endif
#endif

#if __BIONIC_AVAILABILITY_GUARD(24)
int fgetpos64(FILE* _Nonnull __fp, fpos64_t* _Nonnull __pos) __INTRODUCED_IN(24);
int fsetpos64(FILE* _Nonnull __fp, const fpos64_t* _Nonnull __pos) __INTRODUCED_IN(24);
int fseeko64(FILE* _Nonnull __fp, off64_t __offset, int __whence) __INTRODUCED_IN(24);
__nodiscard off64_t ftello64(FILE* _Nonnull __fp) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */

#if defined(__USE_BSD)
/* If __read_fn and __write_fn are both nullptr, it will cause EINVAL */

#if __BIONIC_AVAILABILITY_GUARD(24)
__nodiscard FILE* _Nullable funopen64(const void* _Nullable __cookie,
                int (* __BIONIC_COMPLICATED_NULLNESS __read_fn)(void* _Nonnull, char* _Nonnull, int),
                int (* __BIONIC_COMPLICATED_NULLNESS __write_fn)(void* _Nonnull, const char* _Nonnull, int),
                fpos64_t (* _Nullable __seek_fn)(void* _Nonnull, fpos64_t, int),
                int (* _Nullable __close_fn)(void* _Nonnull)) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */

#endif

__nodiscard FILE* _Nullable fopen(const char* _Nonnull __path, const char* _Nonnull __mode);

#if __BIONIC_AVAILABILITY_GUARD(24)
__nodiscard FILE* _Nullable fopen64(const char* _Nonnull __path, const char* _Nonnull __mode) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */

FILE* _Nullable freopen(const char* _Nullable __path, const char* _Nonnull __mode, FILE* _Nonnull __fp);

#if __BIONIC_AVAILABILITY_GUARD(24)
FILE* _Nullable freopen64(const char* _Nullable __path, const char* _Nonnull __mode, FILE* _Nonnull __fp) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */

__nodiscard FILE* _Nullable tmpfile(void);

#if __BIONIC_AVAILABILITY_GUARD(24)
__nodiscard FILE* _Nullable tmpfile64(void) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */


int snprintf(char* __BIONIC_COMPLICATED_NULLNESS __buf, size_t __size, const char* _Nonnull __fmt, ...) __printflike(3, 4);
int vfscanf(FILE* _Nonnull __fp, const char* _Nonnull __fmt, va_list __args) __scanflike(2, 0);
int vscanf(const char* _Nonnull __fmt , va_list __args) __scanflike(1, 0);
int vsnprintf(char* __BIONIC_COMPLICATED_NULLNESS __buf, size_t __size, const char* _Nonnull __fmt, va_list __args) __printflike(3, 0);
int vsscanf(const char* _Nonnull __s, const char* _Nonnull __fmt, va_list __args) __scanflike(2, 0);

#define L_ctermid 1024 /* size for ctermid() */

#if __BIONIC_AVAILABILITY_GUARD(26)
char* _Nonnull ctermid(char* _Nullable __buf) __INTRODUCED_IN(26);
#endif /* __BIONIC_AVAILABILITY_GUARD(26) */


__nodiscard FILE* _Nullable fdopen(int __fd, const char* _Nonnull __mode);
__nodiscard int fileno(FILE* _Nonnull __fp);
int pclose(FILE* _Nonnull __fp);
__nodiscard FILE* _Nullable popen(const char* _Nonnull __command, const char* _Nonnull __mode);
void flockfile(FILE* _Nonnull  __fp);
int ftrylockfile(FILE* _Nonnull __fp);
void funlockfile(FILE* _Nonnull __fp);
__nodiscard int getc_unlocked(FILE* _Nonnull __fp);
__nodiscard int getchar_unlocked(void);
int putc_unlocked(int __ch, FILE* _Nonnull __fp);
int putchar_unlocked(int __ch);


#if __BIONIC_AVAILABILITY_GUARD(23)
__nodiscard FILE* _Nullable fmemopen(void* _Nullable __buf, size_t __size, const char* _Nonnull __mode) __INTRODUCED_IN(23);
__nodiscard FILE* _Nullable open_memstream(char* _Nonnull * _Nonnull __ptr, size_t* _Nonnull __size_ptr) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


#if defined(__USE_BSD) || defined(__BIONIC__) /* Historically bionic exposed these. */
int  asprintf(char* _Nullable * _Nonnull __s_ptr, const char* _Nonnull __fmt, ...) __printflike(2, 3);
char* _Nullable fgetln(FILE* _Nonnull __fp, size_t* _Nonnull __length_ptr);
int fpurge(FILE* _Nonnull __fp);
void setbuffer(FILE* _Nonnull __fp, char* _Nullable __buf, int __size);
int setlinebuf(FILE* _Nonnull __fp);
int vasprintf(char* _Nullable * _Nonnull __s_ptr, const char* _Nonnull __fmt, va_list __args) __printflike(2, 0);

#if __BIONIC_AVAILABILITY_GUARD(23)
void clearerr_unlocked(FILE* _Nonnull __fp) __INTRODUCED_IN(23);
__nodiscard int feof_unlocked(FILE* _Nonnull __fp) __INTRODUCED_IN(23);
__nodiscard int ferror_unlocked(FILE* _Nonnull __fp) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


#if __BIONIC_AVAILABILITY_GUARD(24)
__nodiscard int fileno_unlocked(FILE* _Nonnull __fp) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */

#define fropen(cookie, fn) funopen(cookie, fn, 0, 0, 0)
#define fwopen(cookie, fn) funopen(cookie, 0, fn, 0, 0)
#endif

#if defined(__USE_BSD)

#if __BIONIC_AVAILABILITY_GUARD(28)
int fflush_unlocked(FILE* _Nullable __fp) __INTRODUCED_IN(28);
__nodiscard int fgetc_unlocked(FILE* _Nonnull __fp) __INTRODUCED_IN(28);
int fputc_unlocked(int __ch, FILE* _Nonnull __fp) __INTRODUCED_IN(28);
size_t fread_unlocked(void* _Nonnull __buf, size_t __size, size_t __count, FILE* _Nonnull __fp) __INTRODUCED_IN(28);
size_t fwrite_unlocked(const void* _Nonnull __buf, size_t __size, size_t __count, FILE* _Nonnull __fp) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

#endif

#if defined(__USE_GNU)

#if __BIONIC_AVAILABILITY_GUARD(28)
int fputs_unlocked(const char* _Nonnull __s, FILE* _Nonnull __fp) __INTRODUCED_IN(28);
char* _Nullable fgets_unlocked(char* _Nonnull __buf, int __size, FILE* _Nonnull __fp) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

#endif

#if defined(__BIONIC_INCLUDE_FORTIFY_HEADERS)
#include <bits/fortify/stdio.h>
#endif

__END_DECLS

#endif
