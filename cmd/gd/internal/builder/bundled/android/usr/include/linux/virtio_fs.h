/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_VIRTIO_FS_H
#define _UAPI_LINUX_VIRTIO_FS_H
#include <linux/types.h>
#include <linux/virtio_ids.h>
#include <linux/virtio_config.h>
#include <linux/virtio_types.h>
struct virtio_fs_config {
  __u8 tag[36];
  __le32 num_request_queues;
} __attribute__((packed));
#define VIRTIO_FS_SHMCAP_ID_CACHE 0
#endif
