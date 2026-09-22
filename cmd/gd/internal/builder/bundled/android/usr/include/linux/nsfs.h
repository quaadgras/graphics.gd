/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __LINUX_NSFS_H
#define __LINUX_NSFS_H
#include <linux/ioctl.h>
#include <linux/types.h>
#define NSIO 0xb7
#define NS_GET_USERNS _IO(NSIO, 0x1)
#define NS_GET_PARENT _IO(NSIO, 0x2)
#define NS_GET_NSTYPE _IO(NSIO, 0x3)
#define NS_GET_OWNER_UID _IO(NSIO, 0x4)
#define NS_GET_MNTNS_ID _IOR(NSIO, 0x5, __u64)
#define NS_GET_PID_FROM_PIDNS _IOR(NSIO, 0x6, int)
#define NS_GET_TGID_FROM_PIDNS _IOR(NSIO, 0x7, int)
#define NS_GET_PID_IN_PIDNS _IOR(NSIO, 0x8, int)
#define NS_GET_TGID_IN_PIDNS _IOR(NSIO, 0x9, int)
struct mnt_ns_info {
  __u32 size;
  __u32 nr_mounts;
  __u64 mnt_ns_id;
};
#define MNT_NS_INFO_SIZE_VER0 16
#define NS_MNT_GET_INFO _IOR(NSIO, 10, struct mnt_ns_info)
#define NS_MNT_GET_NEXT _IOR(NSIO, 11, struct mnt_ns_info)
#define NS_MNT_GET_PREV _IOR(NSIO, 12, struct mnt_ns_info)
#endif
