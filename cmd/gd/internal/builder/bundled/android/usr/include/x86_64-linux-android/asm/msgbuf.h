/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __ASM_X64_MSGBUF_H
#define __ASM_X64_MSGBUF_H
#if !defined(__x86_64__) || !defined(__ILP32__)
#include <asm-generic/msgbuf.h>
#else
#include <asm/ipcbuf.h>
struct msqid64_ds {
  struct ipc64_perm msg_perm;
  __kernel_long_t msg_stime;
  __kernel_long_t msg_rtime;
  __kernel_long_t msg_ctime;
  __kernel_ulong_t msg_cbytes;
  __kernel_ulong_t msg_qnum;
  __kernel_ulong_t msg_qbytes;
  __kernel_pid_t msg_lspid;
  __kernel_pid_t msg_lrpid;
  __kernel_ulong_t __unused4;
  __kernel_ulong_t __unused5;
};
#endif
#endif
