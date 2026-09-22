/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_CACHEFILES_H
#define _LINUX_CACHEFILES_H
#include <linux/types.h>
#include <linux/ioctl.h>
#define CACHEFILES_MSG_MAX_SIZE 1024
enum cachefiles_opcode {
  CACHEFILES_OP_OPEN,
  CACHEFILES_OP_CLOSE,
  CACHEFILES_OP_READ,
};
struct cachefiles_msg {
  __u32 msg_id;
  __u32 opcode;
  __u32 len;
  __u32 object_id;
  __u8 data[];
};
struct cachefiles_open {
  __u32 volume_key_size;
  __u32 cookie_key_size;
  __u32 fd;
  __u32 flags;
  __u8 data[];
};
struct cachefiles_read {
  __u64 off;
  __u64 len;
};
#define CACHEFILES_IOC_READ_COMPLETE _IOW(0x98, 1, int)
#endif
