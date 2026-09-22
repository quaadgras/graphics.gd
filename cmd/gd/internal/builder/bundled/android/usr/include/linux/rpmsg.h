/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_RPMSG_H_
#define _UAPI_RPMSG_H_
#include <linux/ioctl.h>
#include <linux/types.h>
#define RPMSG_ADDR_ANY 0xFFFFFFFF
struct rpmsg_endpoint_info {
  char name[32];
  __u32 src;
  __u32 dst;
};
#define RPMSG_CREATE_EPT_IOCTL _IOW(0xb5, 0x1, struct rpmsg_endpoint_info)
#define RPMSG_DESTROY_EPT_IOCTL _IO(0xb5, 0x2)
#define RPMSG_CREATE_DEV_IOCTL _IOW(0xb5, 0x3, struct rpmsg_endpoint_info)
#define RPMSG_RELEASE_DEV_IOCTL _IOW(0xb5, 0x4, struct rpmsg_endpoint_info)
#define RPMSG_GET_OUTGOING_FLOWCONTROL _IOR(0xb5, 0x5, int)
#define RPMSG_SET_INCOMING_FLOWCONTROL _IOR(0xb5, 0x6, int)
#endif
