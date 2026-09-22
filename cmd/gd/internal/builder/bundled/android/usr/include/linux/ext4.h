/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_EXT4_H
#define _UAPI_LINUX_EXT4_H
#include <linux/fiemap.h>
#include <linux/fs.h>
#include <linux/ioctl.h>
#include <linux/types.h>
#define EXT4_IOC_GETVERSION _IOR('f', 3, long)
#define EXT4_IOC_SETVERSION _IOW('f', 4, long)
#define EXT4_IOC_GETVERSION_OLD FS_IOC_GETVERSION
#define EXT4_IOC_SETVERSION_OLD FS_IOC_SETVERSION
#define EXT4_IOC_GETRSVSZ _IOR('f', 5, long)
#define EXT4_IOC_SETRSVSZ _IOW('f', 6, long)
#define EXT4_IOC_GROUP_EXTEND _IOW('f', 7, unsigned long)
#define EXT4_IOC_GROUP_ADD _IOW('f', 8, struct ext4_new_group_input)
#define EXT4_IOC_MIGRATE _IO('f', 9)
#define EXT4_IOC_ALLOC_DA_BLKS _IO('f', 12)
#define EXT4_IOC_MOVE_EXT _IOWR('f', 15, struct move_extent)
#define EXT4_IOC_RESIZE_FS _IOW('f', 16, __u64)
#define EXT4_IOC_SWAP_BOOT _IO('f', 17)
#define EXT4_IOC_PRECACHE_EXTENTS _IO('f', 18)
#define EXT4_IOC_CLEAR_ES_CACHE _IO('f', 40)
#define EXT4_IOC_GETSTATE _IOW('f', 41, __u32)
#define EXT4_IOC_GET_ES_CACHE _IOWR('f', 42, struct fiemap)
#define EXT4_IOC_CHECKPOINT _IOW('f', 43, __u32)
#define EXT4_IOC_GETFSUUID _IOR('f', 44, struct fsuuid)
#define EXT4_IOC_SETFSUUID _IOW('f', 44, struct fsuuid)
#define EXT4_IOC_SHUTDOWN _IOR('X', 125, __u32)
#define EXT4_IOC32_GETVERSION _IOR('f', 3, int)
#define EXT4_IOC32_SETVERSION _IOW('f', 4, int)
#define EXT4_IOC32_GETRSVSZ _IOR('f', 5, int)
#define EXT4_IOC32_SETRSVSZ _IOW('f', 6, int)
#define EXT4_IOC32_GROUP_EXTEND _IOW('f', 7, unsigned int)
#define EXT4_IOC32_GROUP_ADD _IOW('f', 8, struct compat_ext4_new_group_input)
#define EXT4_IOC32_GETVERSION_OLD FS_IOC32_GETVERSION
#define EXT4_IOC32_SETVERSION_OLD FS_IOC32_SETVERSION
#define EXT4_STATE_FLAG_EXT_PRECACHED 0x00000001
#define EXT4_STATE_FLAG_NEW 0x00000002
#define EXT4_STATE_FLAG_NEWENTRY 0x00000004
#define EXT4_STATE_FLAG_DA_ALLOC_CLOSE 0x00000008
#define EXT4_IOC_CHECKPOINT_FLAG_DISCARD 0x1
#define EXT4_IOC_CHECKPOINT_FLAG_ZEROOUT 0x2
#define EXT4_IOC_CHECKPOINT_FLAG_DRY_RUN 0x4
#define EXT4_IOC_CHECKPOINT_FLAG_VALID (EXT4_IOC_CHECKPOINT_FLAG_DISCARD | EXT4_IOC_CHECKPOINT_FLAG_ZEROOUT | EXT4_IOC_CHECKPOINT_FLAG_DRY_RUN)
struct fsuuid {
  __u32 fsu_len;
  __u32 fsu_flags;
  __u8 fsu_uuid[];
};
struct move_extent {
  __u32 reserved;
  __u32 donor_fd;
  __u64 orig_start;
  __u64 donor_start;
  __u64 len;
  __u64 moved_len;
};
#define EXT4_GOING_FLAGS_DEFAULT 0x0
#define EXT4_GOING_FLAGS_LOGFLUSH 0x1
#define EXT4_GOING_FLAGS_NOLOGFLUSH 0x2
struct ext4_new_group_input {
  __u32 group;
  __u64 block_bitmap;
  __u64 inode_bitmap;
  __u64 inode_table;
  __u32 blocks_count;
  __u16 reserved_blocks;
  __u16 unused;
};
#define EXT4_FIEMAP_EXTENT_HOLE 0x08000000
#endif
