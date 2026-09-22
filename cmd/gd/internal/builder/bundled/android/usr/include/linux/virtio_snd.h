/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef VIRTIO_SND_IF_H
#define VIRTIO_SND_IF_H
#include <linux/virtio_types.h>
enum {
  VIRTIO_SND_F_CTLS = 0
};
struct virtio_snd_config {
  __le32 jacks;
  __le32 streams;
  __le32 chmaps;
  __le32 controls;
};
enum {
  VIRTIO_SND_VQ_CONTROL = 0,
  VIRTIO_SND_VQ_EVENT,
  VIRTIO_SND_VQ_TX,
  VIRTIO_SND_VQ_RX,
  VIRTIO_SND_VQ_MAX
};
enum {
  VIRTIO_SND_D_OUTPUT = 0,
  VIRTIO_SND_D_INPUT
};
enum {
  VIRTIO_SND_R_JACK_INFO = 1,
  VIRTIO_SND_R_JACK_REMAP,
  VIRTIO_SND_R_PCM_INFO = 0x0100,
  VIRTIO_SND_R_PCM_SET_PARAMS,
  VIRTIO_SND_R_PCM_PREPARE,
  VIRTIO_SND_R_PCM_RELEASE,
  VIRTIO_SND_R_PCM_START,
  VIRTIO_SND_R_PCM_STOP,
  VIRTIO_SND_R_CHMAP_INFO = 0x0200,
  VIRTIO_SND_R_CTL_INFO = 0x0300,
  VIRTIO_SND_R_CTL_ENUM_ITEMS,
  VIRTIO_SND_R_CTL_READ,
  VIRTIO_SND_R_CTL_WRITE,
  VIRTIO_SND_R_CTL_TLV_READ,
  VIRTIO_SND_R_CTL_TLV_WRITE,
  VIRTIO_SND_R_CTL_TLV_COMMAND,
  VIRTIO_SND_EVT_JACK_CONNECTED = 0x1000,
  VIRTIO_SND_EVT_JACK_DISCONNECTED,
  VIRTIO_SND_EVT_PCM_PERIOD_ELAPSED = 0x1100,
  VIRTIO_SND_EVT_PCM_XRUN,
  VIRTIO_SND_EVT_CTL_NOTIFY = 0x1200,
  VIRTIO_SND_S_OK = 0x8000,
  VIRTIO_SND_S_BAD_MSG,
  VIRTIO_SND_S_NOT_SUPP,
  VIRTIO_SND_S_IO_ERR
};
struct virtio_snd_hdr {
  __le32 code;
};
struct virtio_snd_event {
  struct virtio_snd_hdr hdr;
  __le32 data;
};
struct virtio_snd_query_info {
  struct virtio_snd_hdr hdr;
  __le32 start_id;
  __le32 count;
  __le32 size;
};
struct virtio_snd_info {
  __le32 hda_fn_nid;
};
struct virtio_snd_jack_hdr {
  struct virtio_snd_hdr hdr;
  __le32 jack_id;
};
enum {
  VIRTIO_SND_JACK_F_REMAP = 0
};
struct virtio_snd_jack_info {
  struct virtio_snd_info hdr;
  __le32 features;
  __le32 hda_reg_defconf;
  __le32 hda_reg_caps;
  __u8 connected;
  __u8 padding[7];
};
struct virtio_snd_jack_remap {
  struct virtio_snd_jack_hdr hdr;
  __le32 association;
  __le32 sequence;
};
struct virtio_snd_pcm_hdr {
  struct virtio_snd_hdr hdr;
  __le32 stream_id;
};
enum {
  VIRTIO_SND_PCM_F_SHMEM_HOST = 0,
  VIRTIO_SND_PCM_F_SHMEM_GUEST,
  VIRTIO_SND_PCM_F_MSG_POLLING,
  VIRTIO_SND_PCM_F_EVT_SHMEM_PERIODS,
  VIRTIO_SND_PCM_F_EVT_XRUNS
};
enum {
  VIRTIO_SND_PCM_FMT_IMA_ADPCM = 0,
  VIRTIO_SND_PCM_FMT_MU_LAW,
  VIRTIO_SND_PCM_FMT_A_LAW,
  VIRTIO_SND_PCM_FMT_S8,
  VIRTIO_SND_PCM_FMT_U8,
  VIRTIO_SND_PCM_FMT_S16,
  VIRTIO_SND_PCM_FMT_U16,
  VIRTIO_SND_PCM_FMT_S18_3,
  VIRTIO_SND_PCM_FMT_U18_3,
  VIRTIO_SND_PCM_FMT_S20_3,
  VIRTIO_SND_PCM_FMT_U20_3,
  VIRTIO_SND_PCM_FMT_S24_3,
  VIRTIO_SND_PCM_FMT_U24_3,
  VIRTIO_SND_PCM_FMT_S20,
  VIRTIO_SND_PCM_FMT_U20,
  VIRTIO_SND_PCM_FMT_S24,
  VIRTIO_SND_PCM_FMT_U24,
  VIRTIO_SND_PCM_FMT_S32,
  VIRTIO_SND_PCM_FMT_U32,
  VIRTIO_SND_PCM_FMT_FLOAT,
  VIRTIO_SND_PCM_FMT_FLOAT64,
  VIRTIO_SND_PCM_FMT_DSD_U8,
  VIRTIO_SND_PCM_FMT_DSD_U16,
  VIRTIO_SND_PCM_FMT_DSD_U32,
  VIRTIO_SND_PCM_FMT_IEC958_SUBFRAME
};
enum {
  VIRTIO_SND_PCM_RATE_5512 = 0,
  VIRTIO_SND_PCM_RATE_8000,
  VIRTIO_SND_PCM_RATE_11025,
  VIRTIO_SND_PCM_RATE_16000,
  VIRTIO_SND_PCM_RATE_22050,
  VIRTIO_SND_PCM_RATE_32000,
  VIRTIO_SND_PCM_RATE_44100,
  VIRTIO_SND_PCM_RATE_48000,
  VIRTIO_SND_PCM_RATE_64000,
  VIRTIO_SND_PCM_RATE_88200,
  VIRTIO_SND_PCM_RATE_96000,
  VIRTIO_SND_PCM_RATE_176400,
  VIRTIO_SND_PCM_RATE_192000,
  VIRTIO_SND_PCM_RATE_384000
};
struct virtio_snd_pcm_info {
  struct virtio_snd_info hdr;
  __le32 features;
  __le64 formats;
  __le64 rates;
  __u8 direction;
  __u8 channels_min;
  __u8 channels_max;
  __u8 padding[5];
};
struct virtio_snd_pcm_set_params {
  struct virtio_snd_pcm_hdr hdr;
  __le32 buffer_bytes;
  __le32 period_bytes;
  __le32 features;
  __u8 channels;
  __u8 format;
  __u8 rate;
  __u8 padding;
};
struct virtio_snd_pcm_xfer {
  __le32 stream_id;
};
struct virtio_snd_pcm_status {
  __le32 status;
  __le32 latency_bytes;
};
struct virtio_snd_chmap_hdr {
  struct virtio_snd_hdr hdr;
  __le32 chmap_id;
};
enum {
  VIRTIO_SND_CHMAP_NONE = 0,
  VIRTIO_SND_CHMAP_NA,
  VIRTIO_SND_CHMAP_MONO,
  VIRTIO_SND_CHMAP_FL,
  VIRTIO_SND_CHMAP_FR,
  VIRTIO_SND_CHMAP_RL,
  VIRTIO_SND_CHMAP_RR,
  VIRTIO_SND_CHMAP_FC,
  VIRTIO_SND_CHMAP_LFE,
  VIRTIO_SND_CHMAP_SL,
  VIRTIO_SND_CHMAP_SR,
  VIRTIO_SND_CHMAP_RC,
  VIRTIO_SND_CHMAP_FLC,
  VIRTIO_SND_CHMAP_FRC,
  VIRTIO_SND_CHMAP_RLC,
  VIRTIO_SND_CHMAP_RRC,
  VIRTIO_SND_CHMAP_FLW,
  VIRTIO_SND_CHMAP_FRW,
  VIRTIO_SND_CHMAP_FLH,
  VIRTIO_SND_CHMAP_FCH,
  VIRTIO_SND_CHMAP_FRH,
  VIRTIO_SND_CHMAP_TC,
  VIRTIO_SND_CHMAP_TFL,
  VIRTIO_SND_CHMAP_TFR,
  VIRTIO_SND_CHMAP_TFC,
  VIRTIO_SND_CHMAP_TRL,
  VIRTIO_SND_CHMAP_TRR,
  VIRTIO_SND_CHMAP_TRC,
  VIRTIO_SND_CHMAP_TFLC,
  VIRTIO_SND_CHMAP_TFRC,
  VIRTIO_SND_CHMAP_TSL,
  VIRTIO_SND_CHMAP_TSR,
  VIRTIO_SND_CHMAP_LLFE,
  VIRTIO_SND_CHMAP_RLFE,
  VIRTIO_SND_CHMAP_BC,
  VIRTIO_SND_CHMAP_BLC,
  VIRTIO_SND_CHMAP_BRC
};
#define VIRTIO_SND_CHMAP_MAX_SIZE 18
struct virtio_snd_chmap_info {
  struct virtio_snd_info hdr;
  __u8 direction;
  __u8 channels;
  __u8 positions[VIRTIO_SND_CHMAP_MAX_SIZE];
};
struct virtio_snd_ctl_hdr {
  struct virtio_snd_hdr hdr;
  __le32 control_id;
};
enum {
  VIRTIO_SND_CTL_ROLE_UNDEFINED = 0,
  VIRTIO_SND_CTL_ROLE_VOLUME,
  VIRTIO_SND_CTL_ROLE_MUTE,
  VIRTIO_SND_CTL_ROLE_GAIN
};
enum {
  VIRTIO_SND_CTL_TYPE_BOOLEAN = 0,
  VIRTIO_SND_CTL_TYPE_INTEGER,
  VIRTIO_SND_CTL_TYPE_INTEGER64,
  VIRTIO_SND_CTL_TYPE_ENUMERATED,
  VIRTIO_SND_CTL_TYPE_BYTES,
  VIRTIO_SND_CTL_TYPE_IEC958
};
enum {
  VIRTIO_SND_CTL_ACCESS_READ = 0,
  VIRTIO_SND_CTL_ACCESS_WRITE,
  VIRTIO_SND_CTL_ACCESS_VOLATILE,
  VIRTIO_SND_CTL_ACCESS_INACTIVE,
  VIRTIO_SND_CTL_ACCESS_TLV_READ,
  VIRTIO_SND_CTL_ACCESS_TLV_WRITE,
  VIRTIO_SND_CTL_ACCESS_TLV_COMMAND
};
struct virtio_snd_ctl_info {
  struct virtio_snd_info hdr;
  __le32 role;
  __le32 type;
  __le32 access;
  __le32 count;
  __le32 index;
  __u8 name[44];
  union {
    struct {
      __le32 min;
      __le32 max;
      __le32 step;
    } integer;
    struct {
      __le64 min;
      __le64 max;
      __le64 step;
    } integer64;
    struct {
      __le32 items;
    } enumerated;
  } value;
};
struct virtio_snd_ctl_enum_item {
  __u8 item[64];
};
struct virtio_snd_ctl_iec958 {
  __u8 status[24];
  __u8 subcode[147];
  __u8 pad;
  __u8 dig_subframe[4];
};
struct virtio_snd_ctl_value {
  union {
    __le32 integer[128];
    __le64 integer64[64];
    __le32 enumerated[128];
    __u8 bytes[512];
    struct virtio_snd_ctl_iec958 iec958;
  } value;
};
enum {
  VIRTIO_SND_CTL_EVT_MASK_VALUE = 0,
  VIRTIO_SND_CTL_EVT_MASK_INFO,
  VIRTIO_SND_CTL_EVT_MASK_TLV
};
struct virtio_snd_ctl_event {
  struct virtio_snd_hdr hdr;
  __le16 control_id;
  __le16 mask;
};
#endif
