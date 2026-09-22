/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI__IF_TUN_H
#define _UAPI__IF_TUN_H
#include <linux/types.h>
#include <linux/if_ether.h>
#include <linux/filter.h>
#define TUN_READQ_SIZE 500
#define TUN_TUN_DEV IFF_TUN
#define TUN_TAP_DEV IFF_TAP
#define TUN_TYPE_MASK 0x000f
#define TUNSETNOCSUM _IOW('T', 200, int)
#define TUNSETDEBUG _IOW('T', 201, int)
#define TUNSETIFF _IOW('T', 202, int)
#define TUNSETPERSIST _IOW('T', 203, int)
#define TUNSETOWNER _IOW('T', 204, int)
#define TUNSETLINK _IOW('T', 205, int)
#define TUNSETGROUP _IOW('T', 206, int)
#define TUNGETFEATURES _IOR('T', 207, unsigned int)
#define TUNSETOFFLOAD _IOW('T', 208, unsigned int)
#define TUNSETTXFILTER _IOW('T', 209, unsigned int)
#define TUNGETIFF _IOR('T', 210, unsigned int)
#define TUNGETSNDBUF _IOR('T', 211, int)
#define TUNSETSNDBUF _IOW('T', 212, int)
#define TUNATTACHFILTER _IOW('T', 213, struct sock_fprog)
#define TUNDETACHFILTER _IOW('T', 214, struct sock_fprog)
#define TUNGETVNETHDRSZ _IOR('T', 215, int)
#define TUNSETVNETHDRSZ _IOW('T', 216, int)
#define TUNSETQUEUE _IOW('T', 217, int)
#define TUNSETIFINDEX _IOW('T', 218, unsigned int)
#define TUNGETFILTER _IOR('T', 219, struct sock_fprog)
#define TUNSETVNETLE _IOW('T', 220, int)
#define TUNGETVNETLE _IOR('T', 221, int)
#define TUNSETVNETBE _IOW('T', 222, int)
#define TUNGETVNETBE _IOR('T', 223, int)
#define TUNSETSTEERINGEBPF _IOR('T', 224, int)
#define TUNSETFILTEREBPF _IOR('T', 225, int)
#define TUNSETCARRIER _IOW('T', 226, int)
#define TUNGETDEVNETNS _IO('T', 227)
#define IFF_TUN 0x0001
#define IFF_TAP 0x0002
#define IFF_NAPI 0x0010
#define IFF_NAPI_FRAGS 0x0020
#define IFF_NO_CARRIER 0x0040
#define IFF_NO_PI 0x1000
#define IFF_ONE_QUEUE 0x2000
#define IFF_VNET_HDR 0x4000
#define IFF_TUN_EXCL 0x8000
#define IFF_MULTI_QUEUE 0x0100
#define IFF_ATTACH_QUEUE 0x0200
#define IFF_DETACH_QUEUE 0x0400
#define IFF_PERSIST 0x0800
#define IFF_NOFILTER 0x1000
#define TUN_TX_TIMESTAMP 1
#define TUN_F_CSUM 0x01
#define TUN_F_TSO4 0x02
#define TUN_F_TSO6 0x04
#define TUN_F_TSO_ECN 0x08
#define TUN_F_UFO 0x10
#define TUN_F_USO4 0x20
#define TUN_F_USO6 0x40
#define TUN_PKT_STRIP 0x0001
struct tun_pi {
  __u16 flags;
  __be16 proto;
};
#define TUN_FLT_ALLMULTI 0x0001
struct tun_filter {
  __u16 flags;
  __u16 count;
  __u8 addr[][ETH_ALEN];
};
#endif
