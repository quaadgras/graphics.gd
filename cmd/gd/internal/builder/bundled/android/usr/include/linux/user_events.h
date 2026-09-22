/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_USER_EVENTS_H
#define _UAPI_LINUX_USER_EVENTS_H
#include <linux/types.h>
#include <linux/ioctl.h>
#define USER_EVENTS_SYSTEM "user_events"
#define USER_EVENTS_MULTI_SYSTEM "user_events_multi"
#define USER_EVENTS_PREFIX "u:"
#define DYN_LOC(offset,size) ((size) << 16 | (offset))
enum user_reg_flag {
  USER_EVENT_REG_PERSIST = 1U << 0,
  USER_EVENT_REG_MULTI_FORMAT = 1U << 1,
  USER_EVENT_REG_MAX = 1U << 2,
};
struct user_reg {
  __u32 size;
  __u8 enable_bit;
  __u8 enable_size;
  __u16 flags;
  __u64 enable_addr;
  __u64 name_args;
  __u32 write_index;
} __attribute__((__packed__));
struct user_unreg {
  __u32 size;
  __u8 disable_bit;
  __u8 __reserved;
  __u16 __reserved2;
  __u64 disable_addr;
} __attribute__((__packed__));
#define DIAG_IOC_MAGIC '*'
#define DIAG_IOCSREG _IOWR(DIAG_IOC_MAGIC, 0, struct user_reg *)
#define DIAG_IOCSDEL _IOW(DIAG_IOC_MAGIC, 1, char *)
#define DIAG_IOCSUNREG _IOW(DIAG_IOC_MAGIC, 2, struct user_unreg *)
#endif
