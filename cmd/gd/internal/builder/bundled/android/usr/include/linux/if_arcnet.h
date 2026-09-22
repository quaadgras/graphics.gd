/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_IF_ARCNET_H
#define _LINUX_IF_ARCNET_H
#include <linux/types.h>
#include <linux/if_ether.h>
#define ARC_P_IP 212
#define ARC_P_IPV6 196
#define ARC_P_ARP 213
#define ARC_P_RARP 214
#define ARC_P_IPX 250
#define ARC_P_NOVELL_EC 236
#define ARC_P_IP_RFC1051 240
#define ARC_P_ARP_RFC1051 241
#define ARC_P_ETHER 232
#define ARC_P_DATAPOINT_BOOT 0
#define ARC_P_DATAPOINT_MOUNT 1
#define ARC_P_POWERLAN_BEACON 8
#define ARC_P_POWERLAN_BEACON2 243
#define ARC_P_LANSOFT 251
#define ARC_P_ATALK 0xDD
#define ARCNET_ALEN 1
struct arc_rfc1201 {
  __u8 proto;
  __u8 split_flag;
  __be16 sequence;
  __u8 payload[];
};
#define RFC1201_HDR_SIZE 4
struct arc_rfc1051 {
  __u8 proto;
  __u8 payload[];
};
#define RFC1051_HDR_SIZE 1
struct arc_eth_encap {
  __u8 proto;
  struct ethhdr eth;
  __u8 payload[];
};
#define ETH_ENCAP_HDR_SIZE 14
struct arc_cap {
  __u8 proto;
  __u8 cookie[sizeof(int)];
  union {
    __u8 ack;
    __u8 raw[0];
  } mes;
};
struct arc_hardware {
  __u8 source;
  __u8 dest;
  __u8 offset[2];
};
#define ARC_HDR_SIZE 4
struct archdr {
  struct arc_hardware hard;
  union {
    struct arc_rfc1201 rfc1201;
    struct arc_rfc1051 rfc1051;
    struct arc_eth_encap eth_encap;
    struct arc_cap cap;
    __u8 raw[0];
  } soft;
};
#endif
