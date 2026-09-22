/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_ICMP_H
#define _UAPI_LINUX_ICMP_H
#include <linux/types.h>
#include <asm/byteorder.h>
#include <linux/if.h>
#include <linux/in6.h>
#define ICMP_ECHOREPLY 0
#define ICMP_DEST_UNREACH 3
#define ICMP_SOURCE_QUENCH 4
#define ICMP_REDIRECT 5
#define ICMP_ECHO 8
#define ICMP_TIME_EXCEEDED 11
#define ICMP_PARAMETERPROB 12
#define ICMP_TIMESTAMP 13
#define ICMP_TIMESTAMPREPLY 14
#define ICMP_INFO_REQUEST 15
#define ICMP_INFO_REPLY 16
#define ICMP_ADDRESS 17
#define ICMP_ADDRESSREPLY 18
#define NR_ICMP_TYPES 18
#define ICMP_NET_UNREACH 0
#define ICMP_HOST_UNREACH 1
#define ICMP_PROT_UNREACH 2
#define ICMP_PORT_UNREACH 3
#define ICMP_FRAG_NEEDED 4
#define ICMP_SR_FAILED 5
#define ICMP_NET_UNKNOWN 6
#define ICMP_HOST_UNKNOWN 7
#define ICMP_HOST_ISOLATED 8
#define ICMP_NET_ANO 9
#define ICMP_HOST_ANO 10
#define ICMP_NET_UNR_TOS 11
#define ICMP_HOST_UNR_TOS 12
#define ICMP_PKT_FILTERED 13
#define ICMP_PREC_VIOLATION 14
#define ICMP_PREC_CUTOFF 15
#define NR_ICMP_UNREACH 15
#define ICMP_REDIR_NET 0
#define ICMP_REDIR_HOST 1
#define ICMP_REDIR_NETTOS 2
#define ICMP_REDIR_HOSTTOS 3
#define ICMP_EXC_TTL 0
#define ICMP_EXC_FRAGTIME 1
#define ICMP_EXT_ECHO 42
#define ICMP_EXT_ECHOREPLY 43
#define ICMP_EXT_CODE_MAL_QUERY 1
#define ICMP_EXT_CODE_NO_IF 2
#define ICMP_EXT_CODE_NO_TABLE_ENT 3
#define ICMP_EXT_CODE_MULT_IFS 4
#define ICMP_EXT_ECHOREPLY_ACTIVE (1 << 2)
#define ICMP_EXT_ECHOREPLY_IPV4 (1 << 1)
#define ICMP_EXT_ECHOREPLY_IPV6 1
#define ICMP_EXT_ECHO_CTYPE_NAME 1
#define ICMP_EXT_ECHO_CTYPE_INDEX 2
#define ICMP_EXT_ECHO_CTYPE_ADDR 3
#define ICMP_AFI_IP 1
#define ICMP_AFI_IP6 2
struct icmphdr {
  __u8 type;
  __u8 code;
  __sum16 checksum;
  union {
    struct {
      __be16 id;
      __be16 sequence;
    } echo;
    __be32 gateway;
    struct {
#ifdef __BIONIC__
      __be16 __linux_unused;
#else
      __be16 __linux_unused;
#endif
      __be16 mtu;
    } frag;
    __u8 reserved[4];
  } un;
};
#define ICMP_FILTER 1
struct icmp_filter {
  __u32 data;
};
struct icmp_ext_hdr {
#ifdef __LITTLE_ENDIAN_BITFIELD
  __u8 reserved1 : 4, version : 4;
#elif defined(__BIG_ENDIAN_BITFIELD)
  __u8 version : 4, reserved1 : 4;
#else
#error "Please fix <asm/byteorder.h>"
#endif
  __u8 reserved2;
  __sum16 checksum;
};
struct icmp_extobj_hdr {
  __be16 length;
  __u8 class_num;
  __u8 class_type;
};
struct icmp_ext_echo_ctype3_hdr {
  __be16 afi;
  __u8 addrlen;
  __u8 reserved;
};
struct icmp_ext_echo_iio {
  struct icmp_extobj_hdr extobj_hdr;
  union {
    char name[IFNAMSIZ];
    __be32 ifindex;
    struct {
      struct icmp_ext_echo_ctype3_hdr ctype3_hdr;
      union {
        __be32 ipv4_addr;
        struct in6_addr ipv6_addr;
      } ip_addr;
    } addr;
  } ident;
};
#endif
