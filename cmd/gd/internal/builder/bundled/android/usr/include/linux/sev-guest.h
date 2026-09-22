/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __UAPI_LINUX_SEV_GUEST_H_
#define __UAPI_LINUX_SEV_GUEST_H_
#include <linux/types.h>
#define SNP_REPORT_USER_DATA_SIZE 64
struct snp_report_req {
  __u8 user_data[SNP_REPORT_USER_DATA_SIZE];
  __u32 vmpl;
  __u8 rsvd[28];
};
struct snp_report_resp {
  __u8 data[4000];
};
struct snp_derived_key_req {
  __u32 root_key_select;
  __u32 rsvd;
  __u64 guest_field_select;
  __u32 vmpl;
  __u32 guest_svn;
  __u64 tcb_version;
};
struct snp_derived_key_resp {
  __u8 data[64];
};
struct snp_guest_request_ioctl {
  __u8 msg_version;
  __u64 req_data;
  __u64 resp_data;
  union {
    __u64 exitinfo2;
    struct {
      __u32 fw_error;
      __u32 vmm_error;
    };
  };
};
struct snp_ext_report_req {
  struct snp_report_req data;
  __u64 certs_address;
  __u32 certs_len;
};
#define SNP_GUEST_REQ_IOC_TYPE 'S'
#define SNP_GET_REPORT _IOWR(SNP_GUEST_REQ_IOC_TYPE, 0x0, struct snp_guest_request_ioctl)
#define SNP_GET_DERIVED_KEY _IOWR(SNP_GUEST_REQ_IOC_TYPE, 0x1, struct snp_guest_request_ioctl)
#define SNP_GET_EXT_REPORT _IOWR(SNP_GUEST_REQ_IOC_TYPE, 0x2, struct snp_guest_request_ioctl)
#define SNP_GUEST_FW_ERR_MASK GENMASK_ULL(31, 0)
#define SNP_GUEST_VMM_ERR_SHIFT 32
#define SNP_GUEST_VMM_ERR(x) (((u64) x) << SNP_GUEST_VMM_ERR_SHIFT)
#define SNP_GUEST_FW_ERR(x) ((x) & SNP_GUEST_FW_ERR_MASK)
#define SNP_GUEST_ERR(vmm_err,fw_err) (SNP_GUEST_VMM_ERR(vmm_err) | SNP_GUEST_FW_ERR(fw_err))
#define SNP_GUEST_VMM_ERR_INVALID_LEN 1
#define SNP_GUEST_VMM_ERR_BUSY 2
#endif
