/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_TDX_GUEST_H_
#define _UAPI_LINUX_TDX_GUEST_H_
#include <linux/ioctl.h>
#include <linux/types.h>
#define TDX_REPORTDATA_LEN 64
#define TDX_REPORT_LEN 1024
struct tdx_report_req {
  __u8 reportdata[TDX_REPORTDATA_LEN];
  __u8 tdreport[TDX_REPORT_LEN];
};
#define TDX_CMD_GET_REPORT0 _IOWR('T', 1, struct tdx_report_req)
#endif
