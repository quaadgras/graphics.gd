/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_VIRTIO_9P_H
#define _LINUX_VIRTIO_9P_H
#include <linux/virtio_types.h>
#include <linux/virtio_ids.h>
#include <linux/virtio_config.h>
#define VIRTIO_9P_MOUNT_TAG 0
struct virtio_9p_config {
  __virtio16 tag_len;
  __u8 tag[];
} __attribute__((packed));
#endif
