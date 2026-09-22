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
 * @file sys/stat.h
 * @brief File status.
 */

#include <sys/cdefs.h>

#include <bits/timespec.h>
#include <linux/stat.h>
#include <sys/types.h>

__BEGIN_DECLS

#if defined(__aarch64__) || defined(__riscv)
#define __STAT64_BODY \
  dev_t st_dev; \
  ino_t st_ino; \
  mode_t st_mode; \
  nlink_t st_nlink; \
  uid_t st_uid; \
  gid_t st_gid; \
  dev_t st_rdev; \
  unsigned long __pad1; \
  off_t st_size; \
  int st_blksize; \
  int __pad2; \
  long st_blocks; \
  struct timespec st_atim; \
  struct timespec st_mtim; \
  struct timespec st_ctim; \
  unsigned int __unused4; \
  unsigned int __unused5; \

#elif defined(__x86_64__)
#define __STAT64_BODY \
  dev_t st_dev; \
  ino_t st_ino; \
  unsigned long st_nlink; \
  mode_t st_mode; \
  uid_t st_uid; \
  gid_t st_gid; \
  unsigned int __pad0; \
  dev_t st_rdev; \
  off_t st_size; \
  long st_blksize; \
  long st_blocks; \
  struct timespec st_atim; \
  struct timespec st_mtim; \
  struct timespec st_ctim; \
  long __pad3[3]; \

#else /* __arm__ || __i386__ */
#define __STAT64_BODY \
  unsigned long long st_dev; \
  unsigned char __pad0[4]; \
  unsigned long __st_ino; \
  unsigned int st_mode; \
  nlink_t st_nlink; \
  uid_t st_uid; \
  gid_t st_gid; \
  unsigned long long st_rdev; \
  unsigned char __pad3[4]; \
  long long st_size; \
  unsigned long st_blksize; \
  unsigned long long st_blocks; \
  struct timespec st_atim; \
  struct timespec st_mtim; \
  struct timespec st_ctim; \
  unsigned long long st_ino; \

#endif

/** The file information returned by fstat()/fstatat()/lstat()/stat(). */
struct stat { __STAT64_BODY };

/**
 * A synonym for `struct stat` on Android,
 * provided for source compatibility with other systems.
 */
struct stat64 { __STAT64_BODY };

#undef __STAT64_BODY

/* Compatibility with older versions of POSIX. */
#define st_atime st_atim.tv_sec
#define st_mtime st_mtim.tv_sec
#define st_ctime st_ctim.tv_sec
/* Compatibility with glibc. */
#define st_atimensec st_atim.tv_nsec
#define st_mtimensec st_mtim.tv_nsec
#define st_ctimensec st_ctim.tv_nsec
/* Compatibility with Linux headers and old NDKs. */
#define st_atime_nsec st_atim.tv_nsec
#define st_mtime_nsec st_mtim.tv_nsec
#define st_ctime_nsec st_ctim.tv_nsec

#if defined(__USE_BSD)
/* Permission macros provided by glibc for compatibility with BSDs. */
#define ACCESSPERMS (S_IRWXU | S_IRWXG | S_IRWXO) /* 0777 */
#define ALLPERMS    (S_ISUID | S_ISGID | S_ISVTX | S_IRWXU | S_IRWXG | S_IRWXO) /* 07777 */
#define DEFFILEMODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP | S_IROTH | S_IWOTH) /* 0666 */
#endif

#if defined(__USE_BSD) || defined(__USE_GNU)
#define S_IREAD S_IRUSR
#define S_IWRITE S_IWUSR
#define S_IEXEC S_IXUSR
#endif

/* POSIX mandates these, but Linux doesn't implement them as distinct file types. */
#define S_TYPEISMQ(__sb) 0
#define S_TYPEISSEM(__sb) 0
#define S_TYPEISSHM(__sb) 0
#define S_TYPEISTMO(__sb) 0

/**
 * [chmod(2)](https://man7.org/linux/man-pages/man2/chmod.2.html)
 * changes the mode of a file given a path.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int chmod(const char* _Nonnull __path, mode_t __mode);

/**
 * [fchmod(2)](https://man7.org/linux/man-pages/man2/fchmod.2.html)
 * changes the mode of a file given a file descriptor.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int fchmod(int __fd, mode_t __mode);

/**
 * [fchmodat(2)](https://man7.org/linux/man-pages/man2/fchmodat.2.html)
 * changes the mode of a file.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int fchmodat(int __dir_fd, const char* _Nonnull __path, mode_t __mode, int __flags);

/**
 * [chmod(2)](https://man7.org/linux/man-pages/man2/chmod.2.html)
 * changes the mode of a file given a path, without following symlinks.
 *
 * Equivalent to `fchmodat(AT_FDCWD, path, mode, AT_SYMLINK_NOFOLLOW)`.
 *
 * Available since API 36.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(36)
int lchmod(const char* _Nonnull __path, mode_t __mode) __INTRODUCED_IN(36);
#endif /* __BIONIC_AVAILABILITY_GUARD(36) */


/**
 * [mkdir(2)](https://man7.org/linux/man-pages/man2/mkdir.2.html)
 * creates a directory.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int mkdir(const char* _Nonnull __path, mode_t __mode);

/**
 * [mkdirat(2)](https://man7.org/linux/man-pages/man2/mkdirat.2.html)
 * creates a directory.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int mkdirat(int __dir_fd, const char* _Nonnull __path, mode_t __mode);

/**
 * [fstat(2)](https://man7.org/linux/man-pages/man2/fstat.2.html)
 * gets file status given a file descriptor.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int fstat(int __fd, struct stat* _Nonnull __buf);

/** An alias for fstat(). */
int fstat64(int __fd, struct stat64* _Nonnull __buf);

/**
 * [fstatat(2)](https://man7.org/linux/man-pages/man2/fstatat.2.html)
 * gets file status.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int fstatat(int __dir_fd, const char* _Nullable __path, struct stat* _Nonnull __buf, int __flags);

/** An alias for fstatat(). */
int fstatat64(int __dir_fd, const char* _Nullable __path, struct stat64* _Nonnull __buf, int __flags);

/**
 * [lstat(2)](https://man7.org/linux/man-pages/man2/lstat.2.html)
 * gets file status given a path, without following symlinks.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int lstat(const char* _Nonnull __path, struct stat* _Nonnull __buf);

/** An alias for lstat(). */
int lstat64(const char* _Nonnull __path, struct stat64* _Nonnull __buf);

/**
 * [stat(2)](https://man7.org/linux/man-pages/man2/stat.2.html)
 * gets file status given a path.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int stat(const char* _Nonnull __path, struct stat* _Nonnull __buf);

/** An alias for stat(). */
int stat64(const char* _Nonnull __path, struct stat64* _Nonnull __buf);

/**
 * [mknod(2)](https://man7.org/linux/man-pages/man2/mknod.2.html)
 * creates a directory, special, or regular file.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int mknod(const char* _Nonnull __path, mode_t __mode, dev_t __dev);

/**
 * [mknodat(2)](https://man7.org/linux/man-pages/man2/mknodat.2.html)
 * creates a directory, special, or regular file.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int mknodat(int __dir_fd, const char* _Nonnull __path, mode_t __mode, dev_t __dev);

/**
 * [umask(2)](https://man7.org/linux/man-pages/man2/umask.2.html)
 * gets and sets the process-wide file mode creation mask.
 *
 * Returns the previous file mode creation mask.
 */
mode_t umask(mode_t __mask);

#if defined(__BIONIC_INCLUDE_FORTIFY_HEADERS)
#include <bits/fortify/stat.h>
#endif

/**
 * [mkfifo(2)](https://man7.org/linux/man-pages/man2/mkfifo.2.html)
 * creates a FIFO.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int mkfifo(const char* _Nonnull __path, mode_t __mode);

/**
 * [mkfifoat(2)](https://man7.org/linux/man-pages/man2/mkfifoat.2.html)
 * creates a FIFO.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(23)
int mkfifoat(int __dir_fd, const char* _Nonnull __path, mode_t __mode) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


/**
 * Used in the tv_nsec field of an argument to utimensat()/futimens()
 * to set that time to the current time.
 */
#define UTIME_NOW  ((1L << 30) - 1L)

/**
 * Used in the tv_nsec field of an argument to utimensat()/futimens()
 * to _not_ set that time.
 */
#define UTIME_OMIT ((1L << 30) - 2L)

/**
 * [utimensat(2)](https://man7.org/linux/man-pages/man2/utimensat.2.html) sets
 * file timestamps.
 *
 * Note: Linux supports `__path` being NULL (in which case `__dir_fd` need not
 * be a directory), allowing futimens() to be implemented with utimensat().
 * For normal use of utimensat(), though, `__path` should be non-null.
 *
 * `__times[0]` is the access time (atime), and `__times[1]` the last modification time (mtime).
 * If `__times` is NULL, both times are set to the current time.
 * See also UTIME_NOW and UTIME_OMIT.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int utimensat(int __dir_fd, const char* __BIONIC_COMPLICATED_NULLNESS __path, const struct timespec __times[_Nullable 2], int __flags);

/**
 * [futimens(3)](https://man7.org/linux/man-pages/man3/futimens.3.html) sets
 * the given file descriptor's timestamp.
 *
 * `__times[0]` is the access time (atime), and `__times[1]` the last modification time (mtime).
 * If `__times` is NULL, both times are set to the current time.
 * See also UTIME_NOW and UTIME_OMIT.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int futimens(int __fd, const struct timespec __times[_Nullable 2]);

#if defined(__USE_GNU)
/**
 * [statx(2)](https://man7.org/linux/man-pages/man2/statx.2.html) returns
 * extended file status information.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 *
 * Available since API level 30.
 */

#if __BIONIC_AVAILABILITY_GUARD(30)
int statx(int __dir_fd, const char* _Nullable __path, int __flags, unsigned __mask, struct statx* _Nonnull __buf) __INTRODUCED_IN(30);
#endif /* __BIONIC_AVAILABILITY_GUARD(30) */

#endif

__END_DECLS
