/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_WMI_H
#define _UAPI_LINUX_WMI_H
#include <linux/ioctl.h>
#include <linux/types.h>
#define WMI_IOC 'W'
struct wmi_ioctl_buffer {
  __u64 length;
  __u8 data[];
};
struct calling_interface_buffer {
  __u16 cmd_class;
  __u16 cmd_select;
  volatile __u32 input[4];
  volatile __u32 output[4];
} __attribute__((__packed__));
struct dell_wmi_extensions {
  __u32 argattrib;
  __u32 blength;
  __u8 data[];
} __attribute__((__packed__));
struct dell_wmi_smbios_buffer {
  __u64 length;
  struct calling_interface_buffer std;
  struct dell_wmi_extensions ext;
} __attribute__((__packed__));
#define CLASS_TOKEN_READ 0
#define CLASS_TOKEN_WRITE 1
#define SELECT_TOKEN_STD 0
#define SELECT_TOKEN_BAT 1
#define SELECT_TOKEN_AC 2
#define CLASS_FLASH_INTERFACE 7
#define SELECT_FLASH_INTERFACE 3
#define CLASS_ADMIN_PROP 10
#define SELECT_ADMIN_PROP 3
#define CLASS_INFO 17
#define SELECT_RFKILL 11
#define SELECT_APP_REGISTRATION 3
#define SELECT_DOCK 22
#define CAPSULE_EN_TOKEN 0x0461
#define CAPSULE_DIS_TOKEN 0x0462
#define WSMT_EN_TOKEN 0x04EC
#define WSMT_DIS_TOKEN 0x04ED
#define DELL_WMI_SMBIOS_CMD _IOWR(WMI_IOC, 0, struct dell_wmi_smbios_buffer)
#endif
