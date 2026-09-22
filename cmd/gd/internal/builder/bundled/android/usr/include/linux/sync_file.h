/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_SYNC_H
#define _UAPI_LINUX_SYNC_H
#include <linux/ioctl.h>
#include <linux/types.h>
struct sync_merge_data {
  char name[32];
  __s32 fd2;
  __s32 fence;
  __u32 flags;
  __u32 pad;
};
struct sync_fence_info {
  char obj_name[32];
  char driver_name[32];
  __s32 status;
  __u32 flags;
  __u64 timestamp_ns;
};
struct sync_file_info {
  char name[32];
  __s32 status;
  __u32 flags;
  __u32 num_fences;
  __u32 pad;
  __u64 sync_fence_info;
};
struct sync_set_deadline {
  __u64 deadline_ns;
  __u64 pad;
};
#define SYNC_IOC_MAGIC '>'
#define SYNC_IOC_MERGE _IOWR(SYNC_IOC_MAGIC, 3, struct sync_merge_data)
#define SYNC_IOC_FILE_INFO _IOWR(SYNC_IOC_MAGIC, 4, struct sync_file_info)
#define SYNC_IOC_SET_DEADLINE _IOW(SYNC_IOC_MAGIC, 5, struct sync_set_deadline)
#endif
