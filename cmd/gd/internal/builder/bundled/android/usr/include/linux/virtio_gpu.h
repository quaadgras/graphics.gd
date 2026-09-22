/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef VIRTIO_GPU_HW_H
#define VIRTIO_GPU_HW_H
#include <linux/types.h>
#define VIRTIO_GPU_F_VIRGL 0
#define VIRTIO_GPU_F_EDID 1
#define VIRTIO_GPU_F_RESOURCE_UUID 2
#define VIRTIO_GPU_F_RESOURCE_BLOB 3
#define VIRTIO_GPU_F_CONTEXT_INIT 4
enum virtio_gpu_ctrl_type {
  VIRTIO_GPU_UNDEFINED = 0,
  VIRTIO_GPU_CMD_GET_DISPLAY_INFO = 0x0100,
  VIRTIO_GPU_CMD_RESOURCE_CREATE_2D,
  VIRTIO_GPU_CMD_RESOURCE_UNREF,
  VIRTIO_GPU_CMD_SET_SCANOUT,
  VIRTIO_GPU_CMD_RESOURCE_FLUSH,
  VIRTIO_GPU_CMD_TRANSFER_TO_HOST_2D,
  VIRTIO_GPU_CMD_RESOURCE_ATTACH_BACKING,
  VIRTIO_GPU_CMD_RESOURCE_DETACH_BACKING,
  VIRTIO_GPU_CMD_GET_CAPSET_INFO,
  VIRTIO_GPU_CMD_GET_CAPSET,
  VIRTIO_GPU_CMD_GET_EDID,
  VIRTIO_GPU_CMD_RESOURCE_ASSIGN_UUID,
  VIRTIO_GPU_CMD_RESOURCE_CREATE_BLOB,
  VIRTIO_GPU_CMD_SET_SCANOUT_BLOB,
  VIRTIO_GPU_CMD_CTX_CREATE = 0x0200,
  VIRTIO_GPU_CMD_CTX_DESTROY,
  VIRTIO_GPU_CMD_CTX_ATTACH_RESOURCE,
  VIRTIO_GPU_CMD_CTX_DETACH_RESOURCE,
  VIRTIO_GPU_CMD_RESOURCE_CREATE_3D,
  VIRTIO_GPU_CMD_TRANSFER_TO_HOST_3D,
  VIRTIO_GPU_CMD_TRANSFER_FROM_HOST_3D,
  VIRTIO_GPU_CMD_SUBMIT_3D,
  VIRTIO_GPU_CMD_RESOURCE_MAP_BLOB,
  VIRTIO_GPU_CMD_RESOURCE_UNMAP_BLOB,
  VIRTIO_GPU_CMD_UPDATE_CURSOR = 0x0300,
  VIRTIO_GPU_CMD_MOVE_CURSOR,
  VIRTIO_GPU_RESP_OK_NODATA = 0x1100,
  VIRTIO_GPU_RESP_OK_DISPLAY_INFO,
  VIRTIO_GPU_RESP_OK_CAPSET_INFO,
  VIRTIO_GPU_RESP_OK_CAPSET,
  VIRTIO_GPU_RESP_OK_EDID,
  VIRTIO_GPU_RESP_OK_RESOURCE_UUID,
  VIRTIO_GPU_RESP_OK_MAP_INFO,
  VIRTIO_GPU_RESP_ERR_UNSPEC = 0x1200,
  VIRTIO_GPU_RESP_ERR_OUT_OF_MEMORY,
  VIRTIO_GPU_RESP_ERR_INVALID_SCANOUT_ID,
  VIRTIO_GPU_RESP_ERR_INVALID_RESOURCE_ID,
  VIRTIO_GPU_RESP_ERR_INVALID_CONTEXT_ID,
  VIRTIO_GPU_RESP_ERR_INVALID_PARAMETER,
};
enum virtio_gpu_shm_id {
  VIRTIO_GPU_SHM_ID_UNDEFINED = 0,
  VIRTIO_GPU_SHM_ID_HOST_VISIBLE = 1
};
#define VIRTIO_GPU_FLAG_FENCE (1 << 0)
#define VIRTIO_GPU_FLAG_INFO_RING_IDX (1 << 1)
struct virtio_gpu_ctrl_hdr {
  __le32 type;
  __le32 flags;
  __le64 fence_id;
  __le32 ctx_id;
  __u8 ring_idx;
  __u8 padding[3];
};
struct virtio_gpu_cursor_pos {
  __le32 scanout_id;
  __le32 x;
  __le32 y;
  __le32 padding;
};
struct virtio_gpu_update_cursor {
  struct virtio_gpu_ctrl_hdr hdr;
  struct virtio_gpu_cursor_pos pos;
  __le32 resource_id;
  __le32 hot_x;
  __le32 hot_y;
  __le32 padding;
};
struct virtio_gpu_rect {
  __le32 x;
  __le32 y;
  __le32 width;
  __le32 height;
};
struct virtio_gpu_resource_unref {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 resource_id;
  __le32 padding;
};
struct virtio_gpu_resource_create_2d {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 resource_id;
  __le32 format;
  __le32 width;
  __le32 height;
};
struct virtio_gpu_set_scanout {
  struct virtio_gpu_ctrl_hdr hdr;
  struct virtio_gpu_rect r;
  __le32 scanout_id;
  __le32 resource_id;
};
struct virtio_gpu_resource_flush {
  struct virtio_gpu_ctrl_hdr hdr;
  struct virtio_gpu_rect r;
  __le32 resource_id;
  __le32 padding;
};
struct virtio_gpu_transfer_to_host_2d {
  struct virtio_gpu_ctrl_hdr hdr;
  struct virtio_gpu_rect r;
  __le64 offset;
  __le32 resource_id;
  __le32 padding;
};
struct virtio_gpu_mem_entry {
  __le64 addr;
  __le32 length;
  __le32 padding;
};
struct virtio_gpu_resource_attach_backing {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 resource_id;
  __le32 nr_entries;
};
struct virtio_gpu_resource_detach_backing {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 resource_id;
  __le32 padding;
};
#define VIRTIO_GPU_MAX_SCANOUTS 16
struct virtio_gpu_resp_display_info {
  struct virtio_gpu_ctrl_hdr hdr;
  struct virtio_gpu_display_one {
    struct virtio_gpu_rect r;
    __le32 enabled;
    __le32 flags;
  } pmodes[VIRTIO_GPU_MAX_SCANOUTS];
};
struct virtio_gpu_box {
  __le32 x, y, z;
  __le32 w, h, d;
};
struct virtio_gpu_transfer_host_3d {
  struct virtio_gpu_ctrl_hdr hdr;
  struct virtio_gpu_box box;
  __le64 offset;
  __le32 resource_id;
  __le32 level;
  __le32 stride;
  __le32 layer_stride;
};
#define VIRTIO_GPU_RESOURCE_FLAG_Y_0_TOP (1 << 0)
struct virtio_gpu_resource_create_3d {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 resource_id;
  __le32 target;
  __le32 format;
  __le32 bind;
  __le32 width;
  __le32 height;
  __le32 depth;
  __le32 array_size;
  __le32 last_level;
  __le32 nr_samples;
  __le32 flags;
  __le32 padding;
};
#define VIRTIO_GPU_CONTEXT_INIT_CAPSET_ID_MASK 0x000000ff
struct virtio_gpu_ctx_create {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 nlen;
  __le32 context_init;
  char debug_name[64];
};
struct virtio_gpu_ctx_destroy {
  struct virtio_gpu_ctrl_hdr hdr;
};
struct virtio_gpu_ctx_resource {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 resource_id;
  __le32 padding;
};
struct virtio_gpu_cmd_submit {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 size;
  __le32 padding;
};
#define VIRTIO_GPU_CAPSET_VIRGL 1
#define VIRTIO_GPU_CAPSET_VIRGL2 2
#define VIRTIO_GPU_CAPSET_VENUS 4
#define VIRTIO_GPU_CAPSET_DRM 6
struct virtio_gpu_get_capset_info {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 capset_index;
  __le32 padding;
};
struct virtio_gpu_resp_capset_info {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 capset_id;
  __le32 capset_max_version;
  __le32 capset_max_size;
  __le32 padding;
};
struct virtio_gpu_get_capset {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 capset_id;
  __le32 capset_version;
};
struct virtio_gpu_resp_capset {
  struct virtio_gpu_ctrl_hdr hdr;
  __u8 capset_data[];
};
struct virtio_gpu_cmd_get_edid {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 scanout;
  __le32 padding;
};
struct virtio_gpu_resp_edid {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 size;
  __le32 padding;
  __u8 edid[1024];
};
#define VIRTIO_GPU_EVENT_DISPLAY (1 << 0)
struct virtio_gpu_config {
  __le32 events_read;
  __le32 events_clear;
  __le32 num_scanouts;
  __le32 num_capsets;
};
enum virtio_gpu_formats {
  VIRTIO_GPU_FORMAT_B8G8R8A8_UNORM = 1,
  VIRTIO_GPU_FORMAT_B8G8R8X8_UNORM = 2,
  VIRTIO_GPU_FORMAT_A8R8G8B8_UNORM = 3,
  VIRTIO_GPU_FORMAT_X8R8G8B8_UNORM = 4,
  VIRTIO_GPU_FORMAT_R8G8B8A8_UNORM = 67,
  VIRTIO_GPU_FORMAT_X8B8G8R8_UNORM = 68,
  VIRTIO_GPU_FORMAT_A8B8G8R8_UNORM = 121,
  VIRTIO_GPU_FORMAT_R8G8B8X8_UNORM = 134,
};
struct virtio_gpu_resource_assign_uuid {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 resource_id;
  __le32 padding;
};
struct virtio_gpu_resp_resource_uuid {
  struct virtio_gpu_ctrl_hdr hdr;
  __u8 uuid[16];
};
struct virtio_gpu_resource_create_blob {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 resource_id;
#define VIRTIO_GPU_BLOB_MEM_GUEST 0x0001
#define VIRTIO_GPU_BLOB_MEM_HOST3D 0x0002
#define VIRTIO_GPU_BLOB_MEM_HOST3D_GUEST 0x0003
#define VIRTIO_GPU_BLOB_FLAG_USE_MAPPABLE 0x0001
#define VIRTIO_GPU_BLOB_FLAG_USE_SHAREABLE 0x0002
#define VIRTIO_GPU_BLOB_FLAG_USE_CROSS_DEVICE 0x0004
  __le32 blob_mem;
  __le32 blob_flags;
  __le32 nr_entries;
  __le64 blob_id;
  __le64 size;
};
struct virtio_gpu_set_scanout_blob {
  struct virtio_gpu_ctrl_hdr hdr;
  struct virtio_gpu_rect r;
  __le32 scanout_id;
  __le32 resource_id;
  __le32 width;
  __le32 height;
  __le32 format;
  __le32 padding;
  __le32 strides[4];
  __le32 offsets[4];
};
struct virtio_gpu_resource_map_blob {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 resource_id;
  __le32 padding;
  __le64 offset;
};
#define VIRTIO_GPU_MAP_CACHE_MASK 0x0f
#define VIRTIO_GPU_MAP_CACHE_NONE 0x00
#define VIRTIO_GPU_MAP_CACHE_CACHED 0x01
#define VIRTIO_GPU_MAP_CACHE_UNCACHED 0x02
#define VIRTIO_GPU_MAP_CACHE_WC 0x03
struct virtio_gpu_resp_map_info {
  struct virtio_gpu_ctrl_hdr hdr;
  __u32 map_info;
  __u32 padding;
};
struct virtio_gpu_resource_unmap_blob {
  struct virtio_gpu_ctrl_hdr hdr;
  __le32 resource_id;
  __le32 padding;
};
#endif
