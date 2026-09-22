/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI__TIPC_SOCKETS_DIAG_H__
#define _UAPI__TIPC_SOCKETS_DIAG_H__
#include <linux/types.h>
#include <linux/sock_diag.h>
struct tipc_sock_diag_req {
  __u8 sdiag_family;
  __u8 sdiag_protocol;
  __u16 pad;
  __u32 tidiag_states;
};
#endif
