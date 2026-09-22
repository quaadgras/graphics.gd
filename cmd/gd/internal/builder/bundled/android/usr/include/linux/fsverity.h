/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_FSVERITY_H
#define _UAPI_LINUX_FSVERITY_H
#include <linux/ioctl.h>
#include <linux/types.h>
#define FS_VERITY_HASH_ALG_SHA256 1
#define FS_VERITY_HASH_ALG_SHA512 2
struct fsverity_enable_arg {
  __u32 version;
  __u32 hash_algorithm;
  __u32 block_size;
  __u32 salt_size;
  __u64 salt_ptr;
  __u32 sig_size;
  __u32 __reserved1;
  __u64 sig_ptr;
  __u64 __reserved2[11];
};
struct fsverity_digest {
  __u16 digest_algorithm;
  __u16 digest_size;
  __u8 digest[];
};
struct fsverity_descriptor {
  __u8 version;
  __u8 hash_algorithm;
  __u8 log_blocksize;
  __u8 salt_size;
  __le32 __reserved_0x04;
  __le64 data_size;
  __u8 root_hash[64];
  __u8 salt[32];
  __u8 __reserved[144];
};
struct fsverity_formatted_digest {
  char magic[8];
  __le16 digest_algorithm;
  __le16 digest_size;
  __u8 digest[];
};
#define FS_VERITY_METADATA_TYPE_MERKLE_TREE 1
#define FS_VERITY_METADATA_TYPE_DESCRIPTOR 2
#define FS_VERITY_METADATA_TYPE_SIGNATURE 3
struct fsverity_read_metadata_arg {
  __u64 metadata_type;
  __u64 offset;
  __u64 length;
  __u64 buf_ptr;
  __u64 __reserved;
};
#define FS_IOC_ENABLE_VERITY _IOW('f', 133, struct fsverity_enable_arg)
#define FS_IOC_MEASURE_VERITY _IOWR('f', 134, struct fsverity_digest)
#define FS_IOC_READ_VERITY_METADATA _IOWR('f', 135, struct fsverity_read_metadata_arg)
#endif
