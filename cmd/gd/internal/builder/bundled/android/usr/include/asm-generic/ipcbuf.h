/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __ASM_GENERIC_IPCBUF_H
#define __ASM_GENERIC_IPCBUF_H
#include <linux/posix_types.h>
struct ipc64_perm {
  __kernel_key_t key;
  __kernel_uid32_t uid;
  __kernel_gid32_t gid;
  __kernel_uid32_t cuid;
  __kernel_gid32_t cgid;
  __kernel_mode_t mode;
  unsigned char __pad1[4 - sizeof(__kernel_mode_t)];
  unsigned short seq;
  unsigned short __pad2;
  __kernel_ulong_t __unused1;
  __kernel_ulong_t __unused2;
};
#endif
