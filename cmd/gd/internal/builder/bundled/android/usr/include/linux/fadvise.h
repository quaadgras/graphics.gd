/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef FADVISE_H_INCLUDED
#define FADVISE_H_INCLUDED
#define POSIX_FADV_NORMAL 0
#define POSIX_FADV_RANDOM 1
#define POSIX_FADV_SEQUENTIAL 2
#define POSIX_FADV_WILLNEED 3
#ifdef __s390x__
#define POSIX_FADV_DONTNEED 6
#define POSIX_FADV_NOREUSE 7
#else
#define POSIX_FADV_DONTNEED 4
#define POSIX_FADV_NOREUSE 5
#endif
#endif
