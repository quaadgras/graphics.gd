/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __UAPI_LINUX_NSM_H
#define __UAPI_LINUX_NSM_H
#include <linux/ioctl.h>
#include <linux/types.h>
#define NSM_MAGIC 0x0A
#define NSM_REQUEST_MAX_SIZE 0x1000
#define NSM_RESPONSE_MAX_SIZE 0x3000
struct nsm_iovec {
  __u64 addr;
  __u64 len;
};
struct nsm_raw {
  struct nsm_iovec request;
  struct nsm_iovec response;
};
#define NSM_IOCTL_RAW _IOWR(NSM_MAGIC, 0x0, struct nsm_raw)
#endif
