/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_IPMI_SSIF_BMC_H
#define _UAPI_LINUX_IPMI_SSIF_BMC_H
#include <linux/types.h>
#define IPMI_SSIF_PAYLOAD_MAX 254
struct ipmi_ssif_msg {
  unsigned int len;
  __u8 payload[IPMI_SSIF_PAYLOAD_MAX];
};
#endif
