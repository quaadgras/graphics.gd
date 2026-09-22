/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_VIRTIO_PCIDEV_H
#define _UAPI_LINUX_VIRTIO_PCIDEV_H
#include <linux/types.h>
enum virtio_pcidev_ops {
  VIRTIO_PCIDEV_OP_RESERVED = 0,
  VIRTIO_PCIDEV_OP_CFG_READ,
  VIRTIO_PCIDEV_OP_CFG_WRITE,
  VIRTIO_PCIDEV_OP_MMIO_READ,
  VIRTIO_PCIDEV_OP_MMIO_WRITE,
  VIRTIO_PCIDEV_OP_MMIO_MEMSET,
  VIRTIO_PCIDEV_OP_INT,
  VIRTIO_PCIDEV_OP_MSI,
  VIRTIO_PCIDEV_OP_PME,
};
struct virtio_pcidev_msg {
  __u8 op;
  __u8 bar;
  __u16 reserved;
  __u32 size;
  __u64 addr;
  __u8 data[];
};
#endif
