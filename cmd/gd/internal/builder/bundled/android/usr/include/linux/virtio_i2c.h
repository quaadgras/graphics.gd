/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_VIRTIO_I2C_H
#define _UAPI_LINUX_VIRTIO_I2C_H
#include <linux/const.h>
#include <linux/types.h>
#define VIRTIO_I2C_F_ZERO_LENGTH_REQUEST 0
#define VIRTIO_I2C_FLAGS_FAIL_NEXT _BITUL(0)
#define VIRTIO_I2C_FLAGS_M_RD _BITUL(1)
struct virtio_i2c_out_hdr {
  __le16 addr;
  __le16 padding;
  __le32 flags;
};
struct virtio_i2c_in_hdr {
  __u8 status;
};
#define VIRTIO_I2C_MSG_OK 0
#define VIRTIO_I2C_MSG_ERR 1
#endif
