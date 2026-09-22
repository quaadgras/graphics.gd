/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_MEMPOLICY_H
#define _UAPI_LINUX_MEMPOLICY_H
#include <linux/errno.h>
enum {
  MPOL_DEFAULT,
  MPOL_PREFERRED,
  MPOL_BIND,
  MPOL_INTERLEAVE,
  MPOL_LOCAL,
  MPOL_PREFERRED_MANY,
  MPOL_WEIGHTED_INTERLEAVE,
  MPOL_MAX,
};
#define MPOL_F_STATIC_NODES (1 << 15)
#define MPOL_F_RELATIVE_NODES (1 << 14)
#define MPOL_F_NUMA_BALANCING (1 << 13)
#define MPOL_MODE_FLAGS (MPOL_F_STATIC_NODES | MPOL_F_RELATIVE_NODES | MPOL_F_NUMA_BALANCING)
#define MPOL_F_NODE (1 << 0)
#define MPOL_F_ADDR (1 << 1)
#define MPOL_F_MEMS_ALLOWED (1 << 2)
#define MPOL_MF_STRICT (1 << 0)
#define MPOL_MF_MOVE (1 << 1)
#define MPOL_MF_MOVE_ALL (1 << 2)
#define MPOL_MF_LAZY (1 << 3)
#define MPOL_MF_INTERNAL (1 << 4)
#define MPOL_MF_VALID (MPOL_MF_STRICT | MPOL_MF_MOVE | MPOL_MF_MOVE_ALL)
#define MPOL_F_SHARED (1 << 0)
#define MPOL_F_MOF (1 << 3)
#define MPOL_F_MORON (1 << 4)
#define RECLAIM_ZONE (1 << 0)
#define RECLAIM_WRITE (1 << 1)
#define RECLAIM_UNMAP (1 << 2)
#endif
