/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_FOU_H
#define _UAPI_LINUX_FOU_H
#define FOU_GENL_NAME "fou"
#define FOU_GENL_VERSION 1
enum {
  FOU_ENCAP_UNSPEC,
  FOU_ENCAP_DIRECT,
  FOU_ENCAP_GUE,
};
enum {
  FOU_ATTR_UNSPEC,
  FOU_ATTR_PORT,
  FOU_ATTR_AF,
  FOU_ATTR_IPPROTO,
  FOU_ATTR_TYPE,
  FOU_ATTR_REMCSUM_NOPARTIAL,
  FOU_ATTR_LOCAL_V4,
  FOU_ATTR_LOCAL_V6,
  FOU_ATTR_PEER_V4,
  FOU_ATTR_PEER_V6,
  FOU_ATTR_PEER_PORT,
  FOU_ATTR_IFINDEX,
  __FOU_ATTR_MAX
};
#define FOU_ATTR_MAX (__FOU_ATTR_MAX - 1)
enum {
  FOU_CMD_UNSPEC,
  FOU_CMD_ADD,
  FOU_CMD_DEL,
  FOU_CMD_GET,
  __FOU_CMD_MAX
};
#define FOU_CMD_MAX (__FOU_CMD_MAX - 1)
#endif
