/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_MPLS_H
#define _UAPI_MPLS_H
#include <linux/types.h>
#include <asm/byteorder.h>
struct mpls_label {
  __be32 entry;
};
#define MPLS_LS_LABEL_MASK 0xFFFFF000
#define MPLS_LS_LABEL_SHIFT 12
#define MPLS_LS_TC_MASK 0x00000E00
#define MPLS_LS_TC_SHIFT 9
#define MPLS_LS_S_MASK 0x00000100
#define MPLS_LS_S_SHIFT 8
#define MPLS_LS_TTL_MASK 0x000000FF
#define MPLS_LS_TTL_SHIFT 0
#define MPLS_LABEL_IPV4NULL 0
#define MPLS_LABEL_RTALERT 1
#define MPLS_LABEL_IPV6NULL 2
#define MPLS_LABEL_IMPLNULL 3
#define MPLS_LABEL_ENTROPY 7
#define MPLS_LABEL_GAL 13
#define MPLS_LABEL_OAMALERT 14
#define MPLS_LABEL_EXTENSION 15
#define MPLS_LABEL_FIRST_UNRESERVED 16
enum {
  MPLS_STATS_UNSPEC,
  MPLS_STATS_LINK,
  __MPLS_STATS_MAX,
};
#define MPLS_STATS_MAX (__MPLS_STATS_MAX - 1)
struct mpls_link_stats {
  __u64 rx_packets;
  __u64 tx_packets;
  __u64 rx_bytes;
  __u64 tx_bytes;
  __u64 rx_errors;
  __u64 tx_errors;
  __u64 rx_dropped;
  __u64 tx_dropped;
  __u64 rx_noroute;
};
#endif
