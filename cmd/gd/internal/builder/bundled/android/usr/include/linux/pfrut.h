/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __PFRUT_H__
#define __PFRUT_H__
#include <linux/ioctl.h>
#include <linux/types.h>
#define PFRUT_IOCTL_MAGIC 0xEE
#define PFRU_IOC_SET_REV _IOW(PFRUT_IOCTL_MAGIC, 0x01, unsigned int)
#define PFRU_IOC_STAGE _IOW(PFRUT_IOCTL_MAGIC, 0x02, unsigned int)
#define PFRU_IOC_ACTIVATE _IOW(PFRUT_IOCTL_MAGIC, 0x03, unsigned int)
#define PFRU_IOC_STAGE_ACTIVATE _IOW(PFRUT_IOCTL_MAGIC, 0x04, unsigned int)
#define PFRU_IOC_QUERY_CAP _IOR(PFRUT_IOCTL_MAGIC, 0x05, struct pfru_update_cap_info)
struct pfru_payload_hdr {
  __u32 sig;
  __u32 hdr_version;
  __u32 hdr_size;
  __u32 hw_ver;
  __u32 rt_ver;
  __u8 platform_id[16];
};
enum pfru_dsm_status {
  DSM_SUCCEED = 0,
  DSM_FUNC_NOT_SUPPORT = 1,
  DSM_INVAL_INPUT = 2,
  DSM_HARDWARE_ERR = 3,
  DSM_RETRY_SUGGESTED = 4,
  DSM_UNKNOWN = 5,
  DSM_FUNC_SPEC_ERR = 6,
};
struct pfru_update_cap_info {
  __u32 status;
  __u32 update_cap;
  __u8 code_type[16];
  __u32 fw_version;
  __u32 code_rt_version;
  __u8 drv_type[16];
  __u32 drv_rt_version;
  __u32 drv_svn;
  __u8 platform_id[16];
  __u8 oem_id[16];
  __u32 oem_info_len;
};
struct pfru_com_buf_info {
  __u32 status;
  __u32 ext_status;
  __u64 addr_lo;
  __u64 addr_hi;
  __u32 buf_size;
};
struct pfru_updated_result {
  __u32 status;
  __u32 ext_status;
  __u64 low_auth_time;
  __u64 high_auth_time;
  __u64 low_exec_time;
  __u64 high_exec_time;
};
struct pfrt_log_data_info {
  __u32 status;
  __u32 ext_status;
  __u64 chunk1_addr_lo;
  __u64 chunk1_addr_hi;
  __u64 chunk2_addr_lo;
  __u64 chunk2_addr_hi;
  __u32 max_data_size;
  __u32 chunk1_size;
  __u32 chunk2_size;
  __u32 rollover_cnt;
  __u32 reset_cnt;
};
struct pfrt_log_info {
  __u32 log_level;
  __u32 log_type;
  __u32 log_revid;
};
#define PFRT_LOG_IOC_SET_INFO _IOW(PFRUT_IOCTL_MAGIC, 0x06, struct pfrt_log_info)
#define PFRT_LOG_IOC_GET_INFO _IOR(PFRUT_IOCTL_MAGIC, 0x07, struct pfrt_log_info)
#define PFRT_LOG_IOC_GET_DATA_INFO _IOR(PFRUT_IOCTL_MAGIC, 0x08, struct pfrt_log_data_info)
#endif
