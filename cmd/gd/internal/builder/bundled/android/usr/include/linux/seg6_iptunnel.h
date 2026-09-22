/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_SEG6_IPTUNNEL_H
#define _UAPI_LINUX_SEG6_IPTUNNEL_H
#include <linux/seg6.h>
enum {
  SEG6_IPTUNNEL_UNSPEC,
  SEG6_IPTUNNEL_SRH,
  __SEG6_IPTUNNEL_MAX,
};
#define SEG6_IPTUNNEL_MAX (__SEG6_IPTUNNEL_MAX - 1)
struct seg6_iptunnel_encap {
  int mode;
  struct ipv6_sr_hdr srh[];
};
#define SEG6_IPTUN_ENCAP_SIZE(x) ((sizeof(* x)) + (((x)->srh->hdrlen + 1) << 3))
enum {
  SEG6_IPTUN_MODE_INLINE,
  SEG6_IPTUN_MODE_ENCAP,
  SEG6_IPTUN_MODE_L2ENCAP,
  SEG6_IPTUN_MODE_ENCAP_RED,
  SEG6_IPTUN_MODE_L2ENCAP_RED,
};
#endif
