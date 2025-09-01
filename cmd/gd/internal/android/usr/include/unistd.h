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

#include <sys/cdefs.h>

#include <stddef.h>
#include <sys/types.h>
#include <sys/select.h>

#include <bits/fcntl.h>
#include <bits/getentropy.h>
#include <bits/getopt.h>
#include <bits/ioctl.h>
#include <bits/lockf.h>
#include <bits/posix_limits.h>
#include <bits/seek_constants.h>
#include <bits/sysconf.h>

__BEGIN_DECLS

#define STDIN_FILENO	0
#define STDOUT_FILENO	1
#define STDERR_FILENO	2

#define F_OK 0
#define X_OK 1
#define W_OK 2
#define R_OK 4

#define _PC_FILESIZEBITS 0
#define _PC_LINK_MAX 1
#define _PC_MAX_CANON 2
#define _PC_MAX_INPUT 3
#define _PC_NAME_MAX 4
#define _PC_PATH_MAX 5
#define _PC_PIPE_BUF 6
#define _PC_2_SYMLINKS 7
#define _PC_ALLOC_SIZE_MIN 8
#define _PC_REC_INCR_XFER_SIZE 9
#define _PC_REC_MAX_XFER_SIZE 10
#define _PC_REC_MIN_XFER_SIZE 11
#define _PC_REC_XFER_ALIGN 12
#define _PC_SYMLINK_MAX 13
#define _PC_CHOWN_RESTRICTED 14
#define _PC_NO_TRUNC 15
#define _PC_VDISABLE 16
#define _PC_ASYNC_IO 17
#define _PC_PRIO_IO 18
#define _PC_SYNC_IO 19

extern char* _Nullable * _Nullable environ;

__noreturn void _exit(int __status);

/**
 * [fork(2)](https://man7.org/linux/man-pages/man2/fork.2.html) creates a new
 * process. fork() runs any handlers set by pthread_atfork().
 *
 * Returns 0 in the child, the pid of the child in the parent,
 * and returns -1 and sets `errno` on failure.
 */
pid_t fork(void);

/**
 * _Fork() creates a new process. _Fork() differs from fork() in that it does
 * not run any handlers set by pthread_atfork(). In addition to any user-defined
 * ones, bionic uses pthread_atfork() handlers to ensure consistency of its own
 * state, so the child should only call
 * [POSIX async-safe](https://man7.org/linux/man-pages/man7/signal-safety.7.html)
 * functions.
 *
 * Returns 0 in the child, the pid of the child in the parent,
 * and returns -1 and sets `errno` on failure.
 *
 * Available since API level 35.
 */

#if __BIONIC_AVAILABILITY_GUARD(35)
pid_t _Fork(void) __INTRODUCED_IN(35);
#endif /* __BIONIC_AVAILABILITY_GUARD(35) */


/**
 * [vfork(2)](https://man7.org/linux/man-pages/man2/vfork.2.html) creates a new
 * process. vfork() differs from fork() in that it does not run any handlers
 * set by pthread_atfork(), and the parent is suspended until the child calls
 * exec() or exits.
 *
 * Returns 0 in the child, the pid of the child in the parent,
 * and returns -1 and sets `errno` on failure.
 */
pid_t vfork(void) __returns_twice;

/**
 * [getpid(2)](https://man7.org/linux/man-pages/man2/getpid.2.html) returns
 * the caller's process ID.
 *
 * Returns the caller's process ID.
 */
pid_t  getpid(void);

/**
 * [gettid(2)](https://man7.org/linux/man-pages/man2/gettid.2.html) returns
 * the caller's thread ID.
 *
 * Returns the caller's thread ID.
 */
pid_t  gettid(void);

pid_t  getpgid(pid_t __pid);
int    setpgid(pid_t __pid, pid_t __pgid);
pid_t  getppid(void);
pid_t  getpgrp(void);
int    setpgrp(void);
pid_t  getsid(pid_t __pid);
pid_t  setsid(void);

int execv(const char* _Nonnull __path, char* _Nullable const* _Nullable __argv);
int execvp(const char* _Nonnull __file, char* _Nullable const* _Nullable __argv);
int execvpe(const char* _Nonnull __file, char* _Nullable const* _Nullable __argv, char* _Nullable const* _Nullable __envp);
int execve(const char* _Nonnull __file, char* _Nullable const* _Nullable __argv, char* _Nullable const* _Nullable __envp);
int execl(const char* _Nonnull __path, const char* _Nullable __arg0, ...) __attribute__((__sentinel__));
int execlp(const char* _Nonnull __file, const char* _Nullable __arg0, ...) __attribute__((__sentinel__));
int execle(const char* _Nonnull __path, const char* _Nullable __arg0, ... /*,  char* const* __envp */)
    __attribute__((__sentinel__(1)));

#if __BIONIC_AVAILABILITY_GUARD(28)
int fexecve(int __fd, char* _Nullable const* _Nullable __argv, char* _Nullable const* _Nullable __envp) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


int nice(int __incr);

/**
 * [setegid(2)](https://man7.org/linux/man-pages/man2/setegid.2.html) sets
 * the effective group ID.
 *
 * On Android, this function only affects the calling thread, not all threads
 * in the process.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int setegid(gid_t __gid);

/**
 * [seteuid(2)](https://man7.org/linux/man-pages/man2/seteuid.2.html) sets
 * the effective user ID.
 *
 * On Android, this function only affects the calling thread, not all threads
 * in the process.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int seteuid(uid_t __uid);

/**
 * [setgid(2)](https://man7.org/linux/man-pages/man2/setgid.2.html) sets
 * the group ID.
 *
 * On Android, this function only affects the calling thread, not all threads
 * in the process.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int setgid(gid_t __gid);

/**
 * [setregid(2)](https://man7.org/linux/man-pages/man2/setregid.2.html) sets
 * the real and effective group IDs (use -1 to leave an ID unchanged).
 *
 * On Android, this function only affects the calling thread, not all threads
 * in the process.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int setregid(gid_t __rgid, gid_t __egid);

/**
 * [setresgid(2)](https://man7.org/linux/man-pages/man2/setresgid.2.html) sets
 * the real, effective, and saved group IDs (use -1 to leave an ID unchanged).
 *
 * On Android, this function only affects the calling thread, not all threads
 * in the process.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int setresgid(gid_t __rgid, gid_t __egid, gid_t __sgid);

/**
 * [setresuid(2)](https://man7.org/linux/man-pages/man2/setresuid.2.html) sets
 * the real, effective, and saved user IDs (use -1 to leave an ID unchanged).
 *
 * On Android, this function only affects the calling thread, not all threads
 * in the process.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int setresuid(uid_t __ruid, uid_t __euid, uid_t __suid);

/**
 * [setreuid(2)](https://man7.org/linux/man-pages/man2/setreuid.2.html) sets
 * the real and effective group IDs (use -1 to leave an ID unchanged).
 *
 * On Android, this function only affects the calling thread, not all threads
 * in the process.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int setreuid(uid_t __ruid, uid_t __euid);

/**
 * [setuid(2)](https://man7.org/linux/man-pages/man2/setuid.2.html) sets
 * the user ID.
 *
 * On Android, this function only affects the calling thread, not all threads
 * in the process.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int setuid(uid_t __uid);

uid_t getuid(void);
uid_t geteuid(void);
gid_t getgid(void);
gid_t getegid(void);
int getgroups(int __size, gid_t* _Nullable __list);
int setgroups(size_t __size, const gid_t* _Nullable __list);
int getresuid(uid_t* _Nonnull __ruid, uid_t* _Nonnull __euid, uid_t* _Nonnull __suid);
int getresgid(gid_t* _Nonnull __rgid, gid_t* _Nonnull __egid, gid_t* _Nonnull __sgid);
char* _Nullable getlogin(void);

#if __BIONIC_AVAILABILITY_GUARD(28)
int getlogin_r(char* _Nonnull __buffer, size_t __buffer_size) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


long fpathconf(int __fd, int __name);
long pathconf(const char* _Nonnull __path, int __name);

int access(const char* _Nonnull __path, int __mode);
int faccessat(int __dirfd, const char* _Nonnull __path, int __mode, int __flags);
int link(const char* _Nonnull __old_path, const char* _Nonnull __new_path);
int linkat(int __old_dir_fd, const char* _Nonnull __old_path, int __new_dir_fd, const char* _Nonnull __new_path, int __flags);
int unlink(const char* _Nonnull __path);
int unlinkat(int __dirfd, const char* _Nonnull __path, int __flags);

/**
 * [chdir(2)](https://man7.org/linux/man-pages/man2/chdir.2.html) changes
 * the current working directory to the given path.
 *
 * This function affects all threads in the process, so is generally a bad idea
 * on Android where most code will be running in a multi-threaded context.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int chdir(const char* _Nonnull __path);

/**
 * [fchdir(2)](https://man7.org/linux/man-pages/man2/fchdir.2.html) changes
 * the current working directory to the given fd.
 *
 * This function affects all threads in the process, so is generally a bad idea
 * on Android where most code will be running in a multi-threaded context.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int fchdir(int __fd);

int rmdir(const char* _Nonnull __path);
int pipe(int __fds[_Nonnull 2]);
#if defined(__USE_GNU)
int pipe2(int __fds[_Nonnull 2], int __flags);
#endif
int chroot(const char* _Nonnull __path);
int symlink(const char* _Nonnull __old_path, const char* _Nonnull __new_path);
int symlinkat(const char* _Nonnull __old_path, int __new_dir_fd, const char* _Nonnull __new_path);
ssize_t readlink(const char* _Nonnull __path, char* _Nonnull __buf, size_t __buf_size);
ssize_t readlinkat(int __dir_fd, const char* _Nonnull __path, char* _Nonnull __buf, size_t __buf_size);
int chown(const char* _Nonnull __path, uid_t __owner, gid_t __group);
int fchown(int __fd, uid_t __owner, gid_t __group);
int fchownat(int __dir_fd, const char* _Nonnull __path, uid_t __owner, gid_t __group, int __flags);
int lchown(const char* _Nonnull __path, uid_t __owner, gid_t __group);
char* _Nullable getcwd(char* _Nullable __buf, size_t __size);

void sync(void);
#if defined(__USE_GNU)

#if __BIONIC_AVAILABILITY_GUARD(28)
int syncfs(int __fd) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */

#endif

int close(int __fd);

/**
 * [read(2)](https://man7.org/linux/man-pages/man2/read.2.html) reads
 * up to `__count` bytes from file descriptor `__fd` into `__buf`.
 *
 * Note: `__buf` is not normally nullable, but may be null in the
 * special case of a zero-length read(), which while not generally
 * useful may be meaningful to some device drivers.
 *
 * Returns the number of bytes read on success, and returns -1 and sets `errno` on failure.
 */
ssize_t read(int __fd, void* __BIONIC_COMPLICATED_NULLNESS __buf, size_t __count);

/**
 * [write(2)](https://man7.org/linux/man-pages/man2/write.2.html) writes
 * up to `__count` bytes to file descriptor `__fd` from `__buf`.
 *
 * Note: `__buf` is not normally nullable, but may be null in the
 * special case of a zero-length write(), which while not generally
 * useful may be meaningful to some device drivers.
 *
 * Returns the number of bytes written on success, and returns -1 and sets `errno` on failure.
 */
ssize_t write(int __fd, const void* __BIONIC_COMPLICATED_NULLNESS __buf, size_t __count);

int dup(int __old_fd);
int dup2(int __old_fd, int __new_fd);
int dup3(int __old_fd, int __new_fd, int __flags);
int fsync(int __fd);
int fdatasync(int __fd);

/* See https://android.googlesource.com/platform/bionic/+/main/docs/32-bit-abi.md */
#if defined(__USE_FILE_OFFSET64)
int truncate(const char* _Nonnull __path, off_t __length) __RENAME(truncate64);
off_t lseek(int __fd, off_t __offset, int __whence) __RENAME(lseek64);
ssize_t pread(int __fd, void* _Nonnull __buf, size_t __count, off_t __offset) __RENAME(pread64);
ssize_t pwrite(int __fd, const void* _Nonnull __buf, size_t __count, off_t __offset) __RENAME(pwrite64);
int ftruncate(int __fd, off_t __length) __RENAME(ftruncate64);
#else
int truncate(const char* _Nonnull __path, off_t __length);
off_t lseek(int __fd, off_t __offset, int __whence);
ssize_t pread(int __fd, void* _Nonnull __buf, size_t __count, off_t __offset);
ssize_t pwrite(int __fd, const void* _Nonnull __buf, size_t __count, off_t __offset);
int ftruncate(int __fd, off_t __length);
#endif

int truncate64(const char* _Nonnull __path, off64_t __length);
off64_t lseek64(int __fd, off64_t __offset, int __whence);
ssize_t pread64(int __fd, void* _Nonnull __buf, size_t __count, off64_t __offset);
ssize_t pwrite64(int __fd, const void* _Nonnull __buf, size_t __count, off64_t __offset);
int ftruncate64(int __fd, off64_t __length);

int pause(void);
unsigned int alarm(unsigned int __seconds);
unsigned int sleep(unsigned int __seconds);
int usleep(useconds_t __microseconds);

int gethostname(char* _Nonnull _buf, size_t __buf_size);

#if __BIONIC_AVAILABILITY_GUARD(23)
int sethostname(const char* _Nonnull __name, size_t __n) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


int brk(void* _Nonnull __addr);
void* _Nullable sbrk(ptrdiff_t __increment);

int isatty(int __fd);
char* _Nullable ttyname(int __fd);
int ttyname_r(int __fd, char* _Nonnull __buf, size_t __buf_size);

int acct(const char* _Nullable __path);

/**
 * [getpagesize(2)](https://man7.org/linux/man-pages/man2/getpagesize.2.html)
 * returns the system's page size. This is slightly faster than going via
 * sysconf(), and avoids the linear search in getauxval().
 *
 * Returns the system's page size in bytes.
 */
int getpagesize(void) __attribute_const__;

long syscall(long __number, ...);

int daemon(int __no_chdir, int __no_close);

#if defined(__arm__)
/**
 * New code should use __builtin___clear_cache() instead, which works on
 * all architectures.
 */
int cacheflush(long __addr, long __nbytes, long __cache);
#endif

pid_t tcgetpgrp(int __fd);
int tcsetpgrp(int __fd, pid_t __pid);

/* Used to retry syscalls that can return EINTR. */
#define TEMP_FAILURE_RETRY(exp) ({         \
    __typeof__(exp) _rc;                   \
    do {                                   \
        _rc = (exp);                       \
    } while (_rc == -1 && errno == EINTR); \
    _rc; })


#if __BIONIC_AVAILABILITY_GUARD(26)
int getdomainname(char* _Nonnull __buf, size_t __buf_size) __INTRODUCED_IN(26);
int setdomainname(const char* _Nonnull __name, size_t __n) __INTRODUCED_IN(26);
#endif /* __BIONIC_AVAILABILITY_GUARD(26) */


/**
 * [copy_file_range(2)](https://man7.org/linux/man-pages/man2/copy_file_range.2.html) copies
 * a range of data from one file descriptor to another.
 *
 * Available since API level 34.
 *
 * Returns the number of bytes copied on success, and returns -1 and sets
 * `errno` on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(34)
ssize_t copy_file_range(int __fd_in, off64_t* _Nullable __off_in, int __fd_out, off64_t* _Nullable __off_out, size_t __length, unsigned int __flags) __INTRODUCED_IN(34);
#endif /* __BIONIC_AVAILABILITY_GUARD(34) */


#if __ANDROID_API__ >= 28
void swab(const void* _Nonnull __src, void* _Nonnull __dst, ssize_t __byte_count) __INTRODUCED_IN(28);
#endif

/**
 * [close_range(2)](https://man7.org/linux/man-pages/man2/close_range.2.html)
 * performs an action (which depends on value of flags) on an inclusive range
 * of file descriptors.
 *
 * Available since API level 34.
 *
 * Note: there is no emulation on too old kernels, hence this will fail with
 * -1/ENOSYS on pre-5.9 kernels, -1/EINVAL for unsupported flags.  In particular
 * CLOSE_RANGE_CLOEXEC requires 5.11, though support was backported to Android
 * Common Kernel 5.10-T.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(34)
int close_range(unsigned int __min_fd, unsigned int __max_fd, int __flags) __INTRODUCED_IN(34);
#endif /* __BIONIC_AVAILABILITY_GUARD(34) */


#if defined(__BIONIC_INCLUDE_FORTIFY_HEADERS)
#define _UNISTD_H_
#include <bits/fortify/unistd.h>
#undef _UNISTD_H_
#endif

__END_DECLS

#include <android/legacy_unistd_inlines.h>
