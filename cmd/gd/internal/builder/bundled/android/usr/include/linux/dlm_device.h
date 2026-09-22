/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_DLM_DEVICE_H
#define _LINUX_DLM_DEVICE_H
#include <linux/dlm.h>
#include <linux/types.h>
#define DLM_USER_LVB_LEN 32
#define DLM_DEVICE_VERSION_MAJOR 6
#define DLM_DEVICE_VERSION_MINOR 0
#define DLM_DEVICE_VERSION_PATCH 2
struct dlm_lock_params {
  __u8 mode;
  __u8 namelen;
  __u16 unused;
  __u32 flags;
  __u32 lkid;
  __u32 parent;
  __u64 xid;
  __u64 timeout;
  void  * castparam;
  void  * castaddr;
  void  * bastparam;
  void  * bastaddr;
  struct dlm_lksb  * lksb;
  char lvb[DLM_USER_LVB_LEN];
  char name[];
};
struct dlm_lspace_params {
  __u32 flags;
  __u32 minor;
  char name[];
};
struct dlm_purge_params {
  __u32 nodeid;
  __u32 pid;
};
struct dlm_write_request {
  __u32 version[3];
  __u8 cmd;
  __u8 is64bit;
  __u8 unused[2];
  union {
    struct dlm_lock_params lock;
    struct dlm_lspace_params lspace;
    struct dlm_purge_params purge;
  } i;
};
struct dlm_device_version {
  __u32 version[3];
};
struct dlm_lock_result {
  __u32 version[3];
  __u32 length;
  void  * user_astaddr;
  void  * user_astparam;
  struct dlm_lksb  * user_lksb;
  struct dlm_lksb lksb;
  __u8 bast_mode;
  __u8 unused[3];
  __u32 lvb_offset;
};
#define DLM_USER_LOCK 1
#define DLM_USER_UNLOCK 2
#define DLM_USER_QUERY 3
#define DLM_USER_CREATE_LOCKSPACE 4
#define DLM_USER_REMOVE_LOCKSPACE 5
#define DLM_USER_PURGE 6
#define DLM_USER_DEADLOCK 7
#define DLM_USER_LSFLG_AUTOFREE 1
#define DLM_USER_LSFLG_FORCEFREE 2
#endif
