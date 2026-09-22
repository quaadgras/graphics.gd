/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPILINUX_ATM_TCP_H
#define _UAPILINUX_ATM_TCP_H
#include <linux/atmapi.h>
#include <linux/atm.h>
#include <linux/atmioc.h>
#include <linux/types.h>
struct atmtcp_hdr {
  __u16 vpi;
  __u16 vci;
  __u32 length;
};
#define ATMTCP_HDR_MAGIC (~0)
#define ATMTCP_CTRL_OPEN 1
#define ATMTCP_CTRL_CLOSE 2
struct atmtcp_control {
  struct atmtcp_hdr hdr;
  int type;
  atm_kptr_t vcc;
  struct sockaddr_atmpvc addr;
  struct atm_qos qos;
  int result;
} __ATM_API_ALIGN;
#define SIOCSIFATMTCP _IO('a', ATMIOC_ITF)
#define ATMTCP_CREATE _IO('a', ATMIOC_ITF + 14)
#define ATMTCP_REMOVE _IO('a', ATMIOC_ITF + 15)
#endif
