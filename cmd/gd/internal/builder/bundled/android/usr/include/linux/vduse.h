/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_VDUSE_H_
#define _UAPI_VDUSE_H_
#include <linux/types.h>
#define VDUSE_BASE 0x81
#define VDUSE_API_VERSION 0
#define VDUSE_GET_API_VERSION _IOR(VDUSE_BASE, 0x00, __u64)
#define VDUSE_SET_API_VERSION _IOW(VDUSE_BASE, 0x01, __u64)
struct vduse_dev_config {
#define VDUSE_NAME_MAX 256
  char name[VDUSE_NAME_MAX];
  __u32 vendor_id;
  __u32 device_id;
  __u64 features;
  __u32 vq_num;
  __u32 vq_align;
  __u32 reserved[13];
  __u32 config_size;
  __u8 config[];
};
#define VDUSE_CREATE_DEV _IOW(VDUSE_BASE, 0x02, struct vduse_dev_config)
#define VDUSE_DESTROY_DEV _IOW(VDUSE_BASE, 0x03, char[VDUSE_NAME_MAX])
struct vduse_iotlb_entry {
  __u64 offset;
  __u64 start;
  __u64 last;
#define VDUSE_ACCESS_RO 0x1
#define VDUSE_ACCESS_WO 0x2
#define VDUSE_ACCESS_RW 0x3
  __u8 perm;
};
#define VDUSE_IOTLB_GET_FD _IOWR(VDUSE_BASE, 0x10, struct vduse_iotlb_entry)
#define VDUSE_DEV_GET_FEATURES _IOR(VDUSE_BASE, 0x11, __u64)
struct vduse_config_data {
  __u32 offset;
  __u32 length;
  __u8 buffer[];
};
#define VDUSE_DEV_SET_CONFIG _IOW(VDUSE_BASE, 0x12, struct vduse_config_data)
#define VDUSE_DEV_INJECT_CONFIG_IRQ _IO(VDUSE_BASE, 0x13)
struct vduse_vq_config {
  __u32 index;
  __u16 max_size;
  __u16 reserved[13];
};
#define VDUSE_VQ_SETUP _IOW(VDUSE_BASE, 0x14, struct vduse_vq_config)
struct vduse_vq_state_split {
  __u16 avail_index;
};
struct vduse_vq_state_packed {
  __u16 last_avail_counter;
  __u16 last_avail_idx;
  __u16 last_used_counter;
  __u16 last_used_idx;
};
struct vduse_vq_info {
  __u32 index;
  __u32 num;
  __u64 desc_addr;
  __u64 driver_addr;
  __u64 device_addr;
  union {
    struct vduse_vq_state_split split;
    struct vduse_vq_state_packed packed;
  };
  __u8 ready;
};
#define VDUSE_VQ_GET_INFO _IOWR(VDUSE_BASE, 0x15, struct vduse_vq_info)
struct vduse_vq_eventfd {
  __u32 index;
#define VDUSE_EVENTFD_DEASSIGN - 1
  int fd;
};
#define VDUSE_VQ_SETUP_KICKFD _IOW(VDUSE_BASE, 0x16, struct vduse_vq_eventfd)
#define VDUSE_VQ_INJECT_IRQ _IOW(VDUSE_BASE, 0x17, __u32)
struct vduse_iova_umem {
  __u64 uaddr;
  __u64 iova;
  __u64 size;
  __u64 reserved[3];
};
#define VDUSE_IOTLB_REG_UMEM _IOW(VDUSE_BASE, 0x18, struct vduse_iova_umem)
#define VDUSE_IOTLB_DEREG_UMEM _IOW(VDUSE_BASE, 0x19, struct vduse_iova_umem)
struct vduse_iova_info {
  __u64 start;
  __u64 last;
#define VDUSE_IOVA_CAP_UMEM (1 << 0)
  __u64 capability;
  __u64 reserved[3];
};
#define VDUSE_IOTLB_GET_INFO _IOWR(VDUSE_BASE, 0x1a, struct vduse_iova_info)
enum vduse_req_type {
  VDUSE_GET_VQ_STATE,
  VDUSE_SET_STATUS,
  VDUSE_UPDATE_IOTLB,
};
struct vduse_vq_state {
  __u32 index;
  union {
    struct vduse_vq_state_split split;
    struct vduse_vq_state_packed packed;
  };
};
struct vduse_dev_status {
  __u8 status;
};
struct vduse_iova_range {
  __u64 start;
  __u64 last;
};
struct vduse_dev_request {
  __u32 type;
  __u32 request_id;
  __u32 reserved[4];
  union {
    struct vduse_vq_state vq_state;
    struct vduse_dev_status s;
    struct vduse_iova_range iova;
    __u32 padding[32];
  };
};
struct vduse_dev_response {
  __u32 request_id;
#define VDUSE_REQ_RESULT_OK 0x00
#define VDUSE_REQ_RESULT_FAILED 0x01
  __u32 result;
  __u32 reserved[4];
  union {
    struct vduse_vq_state vq_state;
    __u32 padding[32];
  };
};
#endif
