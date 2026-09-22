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

#pragma once

/**
 * @file dirent.h
 * @brief Directory entry iteration.
 */

#include <sys/cdefs.h>

#include <stdint.h>
#include <sys/types.h>

__BEGIN_DECLS

/** d_type value when the type is not known. */
#define DT_UNKNOWN 0
/** d_type value for a FIFO. */
#define DT_FIFO 1
/** d_type value for a character device. */
#define DT_CHR 2
/** d_type value for a directory. */
#define DT_DIR 4
/** d_type value for a block device. */
#define DT_BLK 6
/** d_type value for a regular file. */
#define DT_REG 8
/** d_type value for a symbolic link. */
#define DT_LNK 10
/** d_type value for a socket. */
#define DT_SOCK 12
#define DT_WHT 14

#if defined(__LP64__)
#define __DIRENT64_INO_T ino_t
#else
#define __DIRENT64_INO_T uint64_t /* Historical accident. */
#endif

#define __DIRENT64_BODY \
    __DIRENT64_INO_T d_ino; \
    off64_t d_off; \
    unsigned short d_reclen; \
    unsigned char d_type; \
    char d_name[256]; \

/** The structure returned by readdir(). Identical to dirent64 on Android. */
struct dirent { __DIRENT64_BODY };
/** The structure returned by readdir64(). Identical to dirent on Android. */
struct dirent64 { __DIRENT64_BODY };

#undef __DIRENT64_BODY
#undef __DIRENT64_INO_T

/* glibc compatibility. */
#undef _DIRENT_HAVE_D_NAMLEN /* Linux doesn't have a d_namlen field. */
#define _DIRENT_HAVE_D_RECLEN
#define _DIRENT_HAVE_D_OFF
#define _DIRENT_HAVE_D_TYPE

#define d_fileno d_ino

/** The structure returned by opendir()/fdopendir(). */
typedef struct DIR DIR;

/**
 * [opendir(3)](https://man7.org/linux/man-pages/man3/opendir.3.html)
 * opens a directory stream for the directory at `__path`.
 *
 * Returns null and sets `errno` on failure.
 */
DIR* _Nullable opendir(const char* _Nonnull __path);

/**
 * [fdopendir(3)](https://man7.org/linux/man-pages/man3/fdopendir.3.html)
 * opens a directory stream for the directory at `__dir_fd`.
 *
 * Returns null and sets `errno` on failure.
 */
DIR* _Nullable fdopendir(int __dir_fd);

/**
 * [readdir(3)](https://man7.org/linux/man-pages/man3/readdir.3.html)
 * returns the next directory entry in the given directory.
 *
 * Returns a pointer to a directory entry on success,
 * or returns null and leaves `errno` unchanged at the end of the directory,
 * or returns null and sets `errno` on failure.
 */
struct dirent* _Nullable readdir(DIR* _Nonnull __dir);

/**
 * [readdir64(3)](https://man7.org/linux/man-pages/man3/readdir.3.html)
 * returns the next directory entry in the given directory.
 *
 * Returns a pointer to a directory entry on success,
 * or returns null and leaves `errno` unchanged at the end of the directory,
 * or returns null and sets `errno` on failure.
 */
struct dirent64* _Nullable readdir64(DIR* _Nonnull __dir);

int readdir_r(DIR* _Nonnull __dir, struct dirent* _Nonnull __entry, struct dirent* _Nullable * _Nonnull __buffer) __attribute__((__deprecated__("readdir_r is deprecated; use readdir instead")));
int readdir64_r(DIR* _Nonnull __dir, struct dirent64* _Nonnull __entry, struct dirent64* _Nullable * _Nonnull __buffer) __attribute__((__deprecated__("readdir64_r is deprecated; use readdir64 instead")));

/**
 * [closedir(3)](https://man7.org/linux/man-pages/man3/closedir.3.html)
 * closes a directory stream.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int closedir(DIR* _Nonnull __dir);

/**
 * [rewinddir(3)](https://man7.org/linux/man-pages/man3/rewinddir.3.html)
 * rewinds a directory stream to the first entry.
 */
void rewinddir(DIR* _Nonnull __dir);

/**
 * [seekdir(3)](https://man7.org/linux/man-pages/man3/seekdir.3.html)
 * seeks a directory stream to the given entry, which must be a value returned
 * by telldir().
 *
 * Available since API level 23.
 */

#if __BIONIC_AVAILABILITY_GUARD(23)
void seekdir(DIR* _Nonnull __dir, long __location) __INTRODUCED_IN(23);

/**
 * [telldir(3)](https://man7.org/linux/man-pages/man3/telldir.3.html)
 * returns a value representing the current position in the directory
 * for use with seekdir().
 *
 * Returns the current position on success and returns -1 and sets `errno` on failure.
 *
 * Available since API level 23.
 */
long telldir(DIR* _Nonnull __dir) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


/**
 * [dirfd(3)](https://man7.org/linux/man-pages/man3/dirfd.3.html)
 * returns the file descriptor backing the given directory stream.
 *
 * Returns a file descriptor on success and returns -1 and sets `errno` on failure.
 */
int dirfd(DIR* _Nonnull __dir);

/**
 * [alphasort(3)](https://man7.org/linux/man-pages/man3/alphasort.3.html) is a
 * comparator for use with scandir() that uses strcoll().
 */
int alphasort(const struct dirent* _Nonnull * _Nonnull __lhs, const struct dirent* _Nonnull * _Nonnull __rhs);

/**
 * [alphasort64(3)](https://man7.org/linux/man-pages/man3/alphasort.3.html) is a
 * comparator for use with scandir64() that uses strcmp().
 */
int alphasort64(const struct dirent64* _Nonnull * _Nonnull __lhs, const struct dirent64* _Nonnull * _Nonnull __rhs);

/**
 * [scandir(3)](https://man7.org/linux/man-pages/man3/scandir.3.html)
 * scans all the directory `__path`, filtering entries with `__filter` and
 * sorting them with qsort() using the given `__comparator`, and storing them
 * into `__name_list`. Passing NULL as the filter accepts all entries.
 * Passing NULL as the comparator skips sorting.
 *
 * Returns the number of entries returned in the list on success,
 * and returns -1 and sets `errno` on failure.
 */
int scandir(const char* _Nonnull __path, struct dirent* _Nonnull * _Nonnull * _Nonnull __name_list, int (* _Nullable __filter)(const struct dirent* _Nonnull), int (* _Nullable __comparator)(const struct dirent* _Nonnull * _Nonnull, const struct dirent* _Nonnull * _Nonnull));

/**
 * [scandir64(3)](https://man7.org/linux/man-pages/man3/scandir.3.html)
 * scans all the directory `__path`, filtering entries with `__filter` and
 * sorting them with qsort() using the given `__comparator`, and storing them
 * into `__name_list`. Passing NULL as the filter accepts all entries.
 * Passing NULL as the comparator skips sorting.
 *
 * Returns the number of entries returned in the list on success,
 * and returns -1 and sets `errno` on failure.
 */
int scandir64(const char* _Nonnull __path, struct dirent64* _Nonnull * _Nonnull * _Nonnull __name_list, int (* _Nullable __filter)(const struct dirent64* _Nonnull), int (* _Nullable __comparator)(const struct dirent64* _Nonnull * _Nonnull, const struct dirent64* _Nonnull * _Nonnull));

#if defined(__USE_GNU)

/**
 * [scandirat64(3)](https://man7.org/linux/man-pages/man3/scandirat.3.html)
 * scans all the directory referenced by the pair of `__dir_fd` and `__path`,
 * filtering entries with `__filter` and sorting them with qsort() using the
 * given `__comparator`, and storing them into `__name_list`. Passing NULL as
 * the filter accepts all entries.
 * Passing NULL as the comparator skips sorting.
 *
 * Returns the number of entries returned in the list on success,
 * and returns -1 and sets `errno` on failure.
 *
 * Available since API level 24.
 */

#if __BIONIC_AVAILABILITY_GUARD(24)
int scandirat64(int __dir_fd, const char* _Nonnull __path, struct dirent64* _Nonnull * _Nonnull * _Nonnull __name_list, int (* _Nullable __filter)(const struct dirent64* _Nonnull), int (* _Nullable __comparator)(const struct dirent64* _Nonnull * _Nonnull, const struct dirent64* _Nonnull * _Nonnull)) __INTRODUCED_IN(24);

/**
 * [scandirat(3)](https://man7.org/linux/man-pages/man3/scandirat.3.html)
 * scans all the directory referenced by the pair of `__dir_fd` and `__path`,
 * filtering entries with `__filter` and sorting them with qsort() using the
 * given `__comparator`, and storing them into `__name_list`. Passing NULL as
 * the filter accepts all entries.
 * Passing NULL as the comparator skips sorting.
 *
 * Returns the number of entries returned in the list on success,
 * and returns -1 and sets `errno` on failure.
 *
 * Available since API level 24.
 */
int scandirat(int __dir_fd, const char* _Nonnull __path, struct dirent* _Nonnull * _Nonnull * _Nonnull __name_list, int (* _Nullable __filter)(const struct dirent* _Nonnull), int (* _Nullable __comparator)(const struct dirent* _Nonnull * _Nonnull, const struct dirent* _Nonnull * _Nonnull)) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */


#endif

__END_DECLS
