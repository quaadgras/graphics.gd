/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI__RFKILL_H
#define _UAPI__RFKILL_H
#include <linux/types.h>
#define RFKILL_STATE_SOFT_BLOCKED 0
#define RFKILL_STATE_UNBLOCKED 1
#define RFKILL_STATE_HARD_BLOCKED 2
enum rfkill_type {
  RFKILL_TYPE_ALL = 0,
  RFKILL_TYPE_WLAN,
  RFKILL_TYPE_BLUETOOTH,
  RFKILL_TYPE_UWB,
  RFKILL_TYPE_WIMAX,
  RFKILL_TYPE_WWAN,
  RFKILL_TYPE_GPS,
  RFKILL_TYPE_FM,
  RFKILL_TYPE_NFC,
  NUM_RFKILL_TYPES,
};
enum rfkill_operation {
  RFKILL_OP_ADD = 0,
  RFKILL_OP_DEL,
  RFKILL_OP_CHANGE,
  RFKILL_OP_CHANGE_ALL,
};
enum rfkill_hard_block_reasons {
  RFKILL_HARD_BLOCK_SIGNAL = 1 << 0,
  RFKILL_HARD_BLOCK_NOT_OWNER = 1 << 1,
};
struct rfkill_event {
  __u32 idx;
  __u8 type;
  __u8 op;
  __u8 soft;
  __u8 hard;
} __attribute__((packed));
struct rfkill_event_ext {
  __u32 idx;
  __u8 type;
  __u8 op;
  __u8 soft;
  __u8 hard;
  __u8 hard_block_reasons;
} __attribute__((packed));
#define RFKILL_EVENT_SIZE_V1 sizeof(struct rfkill_event)
#define RFKILL_IOC_MAGIC 'R'
#define RFKILL_IOC_NOINPUT 1
#define RFKILL_IOCTL_NOINPUT _IO(RFKILL_IOC_MAGIC, RFKILL_IOC_NOINPUT)
#define RFKILL_IOC_MAX_SIZE 2
#define RFKILL_IOCTL_MAX_SIZE _IOW(RFKILL_IOC_MAGIC, RFKILL_IOC_MAX_SIZE, __u32)
#endif
