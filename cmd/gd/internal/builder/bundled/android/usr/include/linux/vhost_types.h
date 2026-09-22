/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_VHOST_TYPES_H
#define _LINUX_VHOST_TYPES_H
#include <linux/types.h>
#include <linux/compiler.h>
#include <linux/virtio_config.h>
#include <linux/virtio_ring.h>
struct vhost_vring_state {
  unsigned int index;
  unsigned int num;
};
struct vhost_vring_file {
  unsigned int index;
  int fd;
};
struct vhost_vring_addr {
  unsigned int index;
  unsigned int flags;
#define VHOST_VRING_F_LOG 0
  __u64 desc_user_addr;
  __u64 used_user_addr;
  __u64 avail_user_addr;
  __u64 log_guest_addr;
};
struct vhost_worker_state {
  unsigned int worker_id;
};
struct vhost_vring_worker {
  unsigned int index;
  unsigned int worker_id;
};
struct vhost_iotlb_msg {
  __u64 iova;
  __u64 size;
  __u64 uaddr;
#define VHOST_ACCESS_RO 0x1
#define VHOST_ACCESS_WO 0x2
#define VHOST_ACCESS_RW 0x3
  __u8 perm;
#define VHOST_IOTLB_MISS 1
#define VHOST_IOTLB_UPDATE 2
#define VHOST_IOTLB_INVALIDATE 3
#define VHOST_IOTLB_ACCESS_FAIL 4
#define VHOST_IOTLB_BATCH_BEGIN 5
#define VHOST_IOTLB_BATCH_END 6
  __u8 type;
};
#define VHOST_IOTLB_MSG 0x1
#define VHOST_IOTLB_MSG_V2 0x2
struct vhost_msg {
  int type;
  union {
    struct vhost_iotlb_msg iotlb;
    __u8 padding[64];
  };
};
struct vhost_msg_v2 {
  __u32 type;
  __u32 asid;
  union {
    struct vhost_iotlb_msg iotlb;
    __u8 padding[64];
  };
};
struct vhost_memory_region {
  __u64 guest_phys_addr;
  __u64 memory_size;
  __u64 userspace_addr;
  __u64 flags_padding;
};
#define VHOST_PAGE_SIZE 0x1000
struct vhost_memory {
  __u32 nregions;
  __u32 padding;
  struct vhost_memory_region regions[];
};
#define VHOST_SCSI_ABI_VERSION 1
struct vhost_scsi_target {
  int abi_version;
  char vhost_wwpn[224];
  unsigned short vhost_tpgt;
  unsigned short reserved;
};
struct vhost_vdpa_config {
  __u32 off;
  __u32 len;
  __u8 buf[];
};
struct vhost_vdpa_iova_range {
  __u64 first;
  __u64 last;
};
#define VHOST_F_LOG_ALL 26
#define VHOST_NET_F_VIRTIO_NET_HDR 27
#define VHOST_BACKEND_F_IOTLB_MSG_V2 0x1
#define VHOST_BACKEND_F_IOTLB_BATCH 0x2
#define VHOST_BACKEND_F_IOTLB_ASID 0x3
#define VHOST_BACKEND_F_SUSPEND 0x4
#define VHOST_BACKEND_F_RESUME 0x5
#define VHOST_BACKEND_F_ENABLE_AFTER_DRIVER_OK 0x6
#define VHOST_BACKEND_F_DESC_ASID 0x7
#define VHOST_BACKEND_F_IOTLB_PERSIST 0x8
#endif
