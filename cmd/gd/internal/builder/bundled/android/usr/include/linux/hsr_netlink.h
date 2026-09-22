/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __UAPI_HSR_NETLINK_H
#define __UAPI_HSR_NETLINK_H
enum {
  HSR_A_UNSPEC,
  HSR_A_NODE_ADDR,
  HSR_A_IFINDEX,
  HSR_A_IF1_AGE,
  HSR_A_IF2_AGE,
  HSR_A_NODE_ADDR_B,
  HSR_A_IF1_SEQ,
  HSR_A_IF2_SEQ,
  HSR_A_IF1_IFINDEX,
  HSR_A_IF2_IFINDEX,
  HSR_A_ADDR_B_IFINDEX,
  __HSR_A_MAX,
};
#define HSR_A_MAX (__HSR_A_MAX - 1)
enum {
  HSR_C_UNSPEC,
  HSR_C_RING_ERROR,
  HSR_C_NODE_DOWN,
  HSR_C_GET_NODE_STATUS,
  HSR_C_SET_NODE_STATUS,
  HSR_C_GET_NODE_LIST,
  HSR_C_SET_NODE_LIST,
  __HSR_C_MAX,
};
#define HSR_C_MAX (__HSR_C_MAX - 1)
#endif
