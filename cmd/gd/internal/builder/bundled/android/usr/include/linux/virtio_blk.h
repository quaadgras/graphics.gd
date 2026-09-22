/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_VIRTIO_BLK_H
#define _LINUX_VIRTIO_BLK_H
#include <linux/types.h>
#include <linux/virtio_ids.h>
#include <linux/virtio_config.h>
#include <linux/virtio_types.h>
#define VIRTIO_BLK_F_SIZE_MAX 1
#define VIRTIO_BLK_F_SEG_MAX 2
#define VIRTIO_BLK_F_GEOMETRY 4
#define VIRTIO_BLK_F_RO 5
#define VIRTIO_BLK_F_BLK_SIZE 6
#define VIRTIO_BLK_F_TOPOLOGY 10
#define VIRTIO_BLK_F_MQ 12
#define VIRTIO_BLK_F_DISCARD 13
#define VIRTIO_BLK_F_WRITE_ZEROES 14
#define VIRTIO_BLK_F_SECURE_ERASE 16
#define VIRTIO_BLK_F_ZONED 17
#ifndef VIRTIO_BLK_NO_LEGACY
#define VIRTIO_BLK_F_BARRIER 0
#define VIRTIO_BLK_F_SCSI 7
#define VIRTIO_BLK_F_FLUSH 9
#define VIRTIO_BLK_F_CONFIG_WCE 11
#define VIRTIO_BLK_F_WCE VIRTIO_BLK_F_FLUSH
#endif
#define VIRTIO_BLK_ID_BYTES 20
struct virtio_blk_config {
  __virtio64 capacity;
  __virtio32 size_max;
  __virtio32 seg_max;
  struct virtio_blk_geometry {
    __virtio16 cylinders;
    __u8 heads;
    __u8 sectors;
  } geometry;
  __virtio32 blk_size;
  __u8 physical_block_exp;
  __u8 alignment_offset;
  __virtio16 min_io_size;
  __virtio32 opt_io_size;
  __u8 wce;
  __u8 unused;
  __virtio16 num_queues;
  __virtio32 max_discard_sectors;
  __virtio32 max_discard_seg;
  __virtio32 discard_sector_alignment;
  __virtio32 max_write_zeroes_sectors;
  __virtio32 max_write_zeroes_seg;
  __u8 write_zeroes_may_unmap;
  __u8 unused1[3];
  __virtio32 max_secure_erase_sectors;
  __virtio32 max_secure_erase_seg;
  __virtio32 secure_erase_sector_alignment;
  struct virtio_blk_zoned_characteristics {
    __virtio32 zone_sectors;
    __virtio32 max_open_zones;
    __virtio32 max_active_zones;
    __virtio32 max_append_sectors;
    __virtio32 write_granularity;
    __u8 model;
    __u8 unused2[3];
  } zoned;
} __attribute__((packed));
#define VIRTIO_BLK_T_IN 0
#define VIRTIO_BLK_T_OUT 1
#ifndef VIRTIO_BLK_NO_LEGACY
#define VIRTIO_BLK_T_SCSI_CMD 2
#endif
#define VIRTIO_BLK_T_FLUSH 4
#define VIRTIO_BLK_T_GET_ID 8
#define VIRTIO_BLK_T_DISCARD 11
#define VIRTIO_BLK_T_WRITE_ZEROES 13
#define VIRTIO_BLK_T_SECURE_ERASE 14
#define VIRTIO_BLK_T_ZONE_APPEND 15
#define VIRTIO_BLK_T_ZONE_REPORT 16
#define VIRTIO_BLK_T_ZONE_OPEN 18
#define VIRTIO_BLK_T_ZONE_CLOSE 20
#define VIRTIO_BLK_T_ZONE_FINISH 22
#define VIRTIO_BLK_T_ZONE_RESET 24
#define VIRTIO_BLK_T_ZONE_RESET_ALL 26
#ifndef VIRTIO_BLK_NO_LEGACY
#define VIRTIO_BLK_T_BARRIER 0x80000000
#endif
struct virtio_blk_outhdr {
  __virtio32 type;
  __virtio32 ioprio;
  __virtio64 sector;
};
#define VIRTIO_BLK_Z_NONE 0
#define VIRTIO_BLK_Z_HM 1
#define VIRTIO_BLK_Z_HA 2
struct virtio_blk_zone_descriptor {
  __virtio64 z_cap;
  __virtio64 z_start;
  __virtio64 z_wp;
  __u8 z_type;
  __u8 z_state;
  __u8 reserved[38];
};
struct virtio_blk_zone_report {
  __virtio64 nr_zones;
  __u8 reserved[56];
  struct virtio_blk_zone_descriptor zones[];
};
#define VIRTIO_BLK_ZT_CONV 1
#define VIRTIO_BLK_ZT_SWR 2
#define VIRTIO_BLK_ZT_SWP 3
#define VIRTIO_BLK_ZS_NOT_WP 0
#define VIRTIO_BLK_ZS_EMPTY 1
#define VIRTIO_BLK_ZS_IOPEN 2
#define VIRTIO_BLK_ZS_EOPEN 3
#define VIRTIO_BLK_ZS_CLOSED 4
#define VIRTIO_BLK_ZS_RDONLY 13
#define VIRTIO_BLK_ZS_FULL 14
#define VIRTIO_BLK_ZS_OFFLINE 15
#define VIRTIO_BLK_WRITE_ZEROES_FLAG_UNMAP 0x00000001
struct virtio_blk_discard_write_zeroes {
  __le64 sector;
  __le32 num_sectors;
  __le32 flags;
};
#ifndef VIRTIO_BLK_NO_LEGACY
struct virtio_scsi_inhdr {
  __virtio32 errors;
  __virtio32 data_len;
  __virtio32 sense_len;
  __virtio32 residual;
};
#endif
#define VIRTIO_BLK_S_OK 0
#define VIRTIO_BLK_S_IOERR 1
#define VIRTIO_BLK_S_UNSUPP 2
#define VIRTIO_BLK_S_ZONE_INVALID_CMD 3
#define VIRTIO_BLK_S_ZONE_UNALIGNED_WP 4
#define VIRTIO_BLK_S_ZONE_OPEN_RESOURCE 5
#define VIRTIO_BLK_S_ZONE_ACTIVE_RESOURCE 6
#endif
