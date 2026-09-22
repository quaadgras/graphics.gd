/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_IPMI_BMC_H
#define _UAPI_LINUX_IPMI_BMC_H
#include <linux/ioctl.h>
#define __IPMI_BMC_IOCTL_MAGIC 0xB1
#define IPMI_BMC_IOCTL_SET_SMS_ATN _IO(__IPMI_BMC_IOCTL_MAGIC, 0x00)
#define IPMI_BMC_IOCTL_CLEAR_SMS_ATN _IO(__IPMI_BMC_IOCTL_MAGIC, 0x01)
#define IPMI_BMC_IOCTL_FORCE_ABORT _IO(__IPMI_BMC_IOCTL_MAGIC, 0x02)
#endif
