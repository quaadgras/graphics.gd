/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_RPL_IPTUNNEL_H
#define _UAPI_LINUX_RPL_IPTUNNEL_H
enum {
  RPL_IPTUNNEL_UNSPEC,
  RPL_IPTUNNEL_SRH,
  __RPL_IPTUNNEL_MAX,
};
#define RPL_IPTUNNEL_MAX (__RPL_IPTUNNEL_MAX - 1)
#define RPL_IPTUNNEL_SRH_SIZE(srh) (((srh)->hdrlen + 1) << 3)
#endif
