/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _DMA_BUF_UAPI_H_
#define _DMA_BUF_UAPI_H_
#include <linux/types.h>
struct dma_buf_sync {
  __u64 flags;
};
#define DMA_BUF_SYNC_READ (1 << 0)
#define DMA_BUF_SYNC_WRITE (2 << 0)
#define DMA_BUF_SYNC_RW (DMA_BUF_SYNC_READ | DMA_BUF_SYNC_WRITE)
#define DMA_BUF_SYNC_START (0 << 2)
#define DMA_BUF_SYNC_END (1 << 2)
#define DMA_BUF_SYNC_VALID_FLAGS_MASK (DMA_BUF_SYNC_RW | DMA_BUF_SYNC_END)
#define DMA_BUF_NAME_LEN 32
struct dma_buf_export_sync_file {
  __u32 flags;
  __s32 fd;
};
struct dma_buf_import_sync_file {
  __u32 flags;
  __s32 fd;
};
#define DMA_BUF_BASE 'b'
#define DMA_BUF_IOCTL_SYNC _IOW(DMA_BUF_BASE, 0, struct dma_buf_sync)
#define DMA_BUF_SET_NAME _IOW(DMA_BUF_BASE, 1, const char *)
#define DMA_BUF_SET_NAME_A _IOW(DMA_BUF_BASE, 1, __u32)
#define DMA_BUF_SET_NAME_B _IOW(DMA_BUF_BASE, 1, __u64)
#define DMA_BUF_IOCTL_EXPORT_SYNC_FILE _IOWR(DMA_BUF_BASE, 2, struct dma_buf_export_sync_file)
#define DMA_BUF_IOCTL_IMPORT_SYNC_FILE _IOW(DMA_BUF_BASE, 3, struct dma_buf_import_sync_file)
#endif
