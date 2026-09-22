/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_STM_H
#define _UAPI_LINUX_STM_H
#include <linux/types.h>
#define STP_MASTER_MAX 0xffff
#define STP_CHANNEL_MAX 0xffff
struct stp_policy_id {
  __u32 size;
  __u16 master;
  __u16 channel;
  __u16 width;
  __u16 __reserved_0;
  __u32 __reserved_1;
  char id[];
};
#define STP_POLICY_ID_SET _IOWR('%', 0, struct stp_policy_id)
#define STP_POLICY_ID_GET _IOR('%', 1, struct stp_policy_id)
#define STP_SET_OPTIONS _IOW('%', 2, __u64)
#endif
