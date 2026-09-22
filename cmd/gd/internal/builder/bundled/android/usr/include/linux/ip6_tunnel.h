/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _IP6_TUNNEL_H
#define _IP6_TUNNEL_H
#include <linux/types.h>
#include <linux/if.h>
#include <linux/in6.h>
#define IPV6_TLV_TNL_ENCAP_LIMIT 4
#define IPV6_DEFAULT_TNL_ENCAP_LIMIT 4
#define IP6_TNL_F_IGN_ENCAP_LIMIT 0x1
#define IP6_TNL_F_USE_ORIG_TCLASS 0x2
#define IP6_TNL_F_USE_ORIG_FLOWLABEL 0x4
#define IP6_TNL_F_MIP6_DEV 0x8
#define IP6_TNL_F_RCV_DSCP_COPY 0x10
#define IP6_TNL_F_USE_ORIG_FWMARK 0x20
#define IP6_TNL_F_ALLOW_LOCAL_REMOTE 0x40
struct ip6_tnl_parm {
  char name[IFNAMSIZ];
  int link;
  __u8 proto;
  __u8 encap_limit;
  __u8 hop_limit;
  __be32 flowinfo;
  __u32 flags;
  struct in6_addr laddr;
  struct in6_addr raddr;
};
struct ip6_tnl_parm2 {
  char name[IFNAMSIZ];
  int link;
  __u8 proto;
  __u8 encap_limit;
  __u8 hop_limit;
  __be32 flowinfo;
  __u32 flags;
  struct in6_addr laddr;
  struct in6_addr raddr;
  __be16 i_flags;
  __be16 o_flags;
  __be32 i_key;
  __be32 o_key;
};
#endif
