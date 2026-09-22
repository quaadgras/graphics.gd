/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_NET_NAMESPACE_H_
#define _UAPI_LINUX_NET_NAMESPACE_H_
enum {
  NETNSA_NONE,
#define NETNSA_NSID_NOT_ASSIGNED - 1
  NETNSA_NSID,
  NETNSA_PID,
  NETNSA_FD,
  NETNSA_TARGET_NSID,
  NETNSA_CURRENT_NSID,
  __NETNSA_MAX,
};
#define NETNSA_MAX (__NETNSA_MAX - 1)
#endif
