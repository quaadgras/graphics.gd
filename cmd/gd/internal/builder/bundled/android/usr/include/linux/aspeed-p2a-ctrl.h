/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_ASPEED_P2A_CTRL_H
#define _UAPI_LINUX_ASPEED_P2A_CTRL_H
#include <linux/ioctl.h>
#include <linux/types.h>
#define ASPEED_P2A_CTRL_READ_ONLY 0
#define ASPEED_P2A_CTRL_READWRITE 1
struct aspeed_p2a_ctrl_mapping {
  __u64 addr;
  __u32 length;
  __u32 flags;
};
#define __ASPEED_P2A_CTRL_IOCTL_MAGIC 0xb3
#define ASPEED_P2A_CTRL_IOCTL_SET_WINDOW _IOW(__ASPEED_P2A_CTRL_IOCTL_MAGIC, 0x00, struct aspeed_p2a_ctrl_mapping)
#define ASPEED_P2A_CTRL_IOCTL_GET_MEMORY_CONFIG _IOWR(__ASPEED_P2A_CTRL_IOCTL_MAGIC, 0x01, struct aspeed_p2a_ctrl_mapping)
#endif
