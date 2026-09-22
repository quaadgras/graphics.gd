/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_VIRTIO_PMEM_H
#define _UAPI_LINUX_VIRTIO_PMEM_H
#include <linux/types.h>
#include <linux/virtio_ids.h>
#include <linux/virtio_config.h>
#define VIRTIO_PMEM_F_SHMEM_REGION 0
#define VIRTIO_PMEM_SHMEM_REGION_ID 0
struct virtio_pmem_config {
  __le64 start;
  __le64 size;
};
#define VIRTIO_PMEM_REQ_TYPE_FLUSH 0
struct virtio_pmem_resp {
  __le32 ret;
};
struct virtio_pmem_req {
  __le32 type;
};
#endif
