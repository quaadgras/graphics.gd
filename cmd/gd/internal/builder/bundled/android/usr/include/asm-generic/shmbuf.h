/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __ASM_GENERIC_SHMBUF_H
#define __ASM_GENERIC_SHMBUF_H
#include <asm/bitsperlong.h>
#include <asm/ipcbuf.h>
#include <asm/posix_types.h>
struct shmid64_ds {
  struct ipc64_perm shm_perm;
  __kernel_size_t shm_segsz;
#if __BITS_PER_LONG == 64
  long shm_atime;
  long shm_dtime;
  long shm_ctime;
#else
  unsigned long shm_atime;
  unsigned long shm_atime_high;
  unsigned long shm_dtime;
  unsigned long shm_dtime_high;
  unsigned long shm_ctime;
  unsigned long shm_ctime_high;
#endif
  __kernel_pid_t shm_cpid;
  __kernel_pid_t shm_lpid;
  unsigned long shm_nattch;
  unsigned long __unused4;
  unsigned long __unused5;
};
struct shminfo64 {
  unsigned long shmmax;
  unsigned long shmmin;
  unsigned long shmmni;
  unsigned long shmseg;
  unsigned long shmall;
  unsigned long __unused1;
  unsigned long __unused2;
  unsigned long __unused3;
  unsigned long __unused4;
};
#endif
