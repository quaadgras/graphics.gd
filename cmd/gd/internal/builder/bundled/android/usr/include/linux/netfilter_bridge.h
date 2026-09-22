/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI__LINUX_BRIDGE_NETFILTER_H
#define _UAPI__LINUX_BRIDGE_NETFILTER_H
#include <linux/in.h>
#include <linux/netfilter.h>
#include <linux/if_ether.h>
#include <linux/if_vlan.h>
#include <linux/if_pppox.h>
#include <limits.h>
#define NF_BR_PRE_ROUTING 0
#define NF_BR_LOCAL_IN 1
#define NF_BR_FORWARD 2
#define NF_BR_LOCAL_OUT 3
#define NF_BR_POST_ROUTING 4
#define NF_BR_BROUTING 5
#define NF_BR_NUMHOOKS 6
enum nf_br_hook_priorities {
  NF_BR_PRI_FIRST = INT_MIN,
  NF_BR_PRI_NAT_DST_BRIDGED = - 300,
  NF_BR_PRI_FILTER_BRIDGED = - 200,
  NF_BR_PRI_BRNF = 0,
  NF_BR_PRI_NAT_DST_OTHER = 100,
  NF_BR_PRI_FILTER_OTHER = 200,
  NF_BR_PRI_NAT_SRC = 300,
  NF_BR_PRI_LAST = INT_MAX,
};
#endif
