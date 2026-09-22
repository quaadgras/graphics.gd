/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI__DLM_PLOCK_DOT_H__
#define _UAPI__DLM_PLOCK_DOT_H__
#include <linux/types.h>
#define DLM_PLOCK_MISC_NAME "dlm_plock"
#define DLM_PLOCK_VERSION_MAJOR 1
#define DLM_PLOCK_VERSION_MINOR 2
#define DLM_PLOCK_VERSION_PATCH 0
enum {
  DLM_PLOCK_OP_LOCK = 1,
  DLM_PLOCK_OP_UNLOCK,
  DLM_PLOCK_OP_GET,
  DLM_PLOCK_OP_CANCEL,
};
#define DLM_PLOCK_FL_CLOSE 1
struct dlm_plock_info {
  __u32 version[3];
  __u8 optype;
  __u8 ex;
  __u8 wait;
  __u8 flags;
  __u32 pid;
  __s32 nodeid;
  __s32 rv;
  __u32 fsid;
  __u64 number;
  __u64 start;
  __u64 end;
  __u64 owner;
};
#endif
