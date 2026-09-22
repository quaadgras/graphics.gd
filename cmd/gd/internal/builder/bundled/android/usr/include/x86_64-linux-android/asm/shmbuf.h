/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __ASM_X86_SHMBUF_H
#define __ASM_X86_SHMBUF_H
#if !defined(__x86_64__) || !defined(__ILP32__)
#include <asm-generic/shmbuf.h>
#else
#include <asm/ipcbuf.h>
#include <asm/posix_types.h>
struct shmid64_ds {
  struct ipc64_perm shm_perm;
  __kernel_size_t shm_segsz;
  __kernel_long_t shm_atime;
  __kernel_long_t shm_dtime;
  __kernel_long_t shm_ctime;
  __kernel_pid_t shm_cpid;
  __kernel_pid_t shm_lpid;
  __kernel_ulong_t shm_nattch;
  __kernel_ulong_t __unused4;
  __kernel_ulong_t __unused5;
};
struct shminfo64 {
  __kernel_ulong_t shmmax;
  __kernel_ulong_t shmmin;
  __kernel_ulong_t shmmni;
  __kernel_ulong_t shmseg;
  __kernel_ulong_t shmall;
  __kernel_ulong_t __unused1;
  __kernel_ulong_t __unused2;
  __kernel_ulong_t __unused3;
  __kernel_ulong_t __unused4;
};
#endif
#endif
