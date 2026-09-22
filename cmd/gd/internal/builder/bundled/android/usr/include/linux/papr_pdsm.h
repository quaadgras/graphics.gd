/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_ASM_POWERPC_PAPR_PDSM_H_
#define _UAPI_ASM_POWERPC_PAPR_PDSM_H_
#include <linux/types.h>
#include <linux/ndctl.h>
#define ND_PDSM_PAYLOAD_MAX_SIZE 184
#define ND_PDSM_HDR_SIZE (sizeof(struct nd_pkg_pdsm) - ND_PDSM_PAYLOAD_MAX_SIZE)
#define PAPR_PDSM_DIMM_HEALTHY 0
#define PAPR_PDSM_DIMM_UNHEALTHY 1
#define PAPR_PDSM_DIMM_CRITICAL 2
#define PAPR_PDSM_DIMM_FATAL 3
#define PDSM_DIMM_HEALTH_RUN_GAUGE_VALID 1
#define PDSM_DIMM_DSC_VALID 2
struct nd_papr_pdsm_health {
  union {
    struct {
      __u32 extension_flags;
      __u8 dimm_unarmed;
      __u8 dimm_bad_shutdown;
      __u8 dimm_bad_restore;
      __u8 dimm_scrubbed;
      __u8 dimm_locked;
      __u8 dimm_encrypted;
      __u16 dimm_health;
      __u16 dimm_fuel_gauge;
      __u64 dimm_dsc;
    };
    __u8 buf[ND_PDSM_PAYLOAD_MAX_SIZE];
  };
};
#define PDSM_SMART_INJECT_HEALTH_FATAL (1 << 0)
#define PDSM_SMART_INJECT_BAD_SHUTDOWN (1 << 1)
struct nd_papr_pdsm_smart_inject {
  union {
    struct {
      __u32 flags;
      __u8 fatal_enable;
      __u8 unsafe_shutdown_enable;
    };
    __u8 buf[ND_PDSM_PAYLOAD_MAX_SIZE];
  };
};
enum papr_pdsm {
  PAPR_PDSM_MIN = 0x0,
  PAPR_PDSM_HEALTH,
  PAPR_PDSM_SMART_INJECT,
  PAPR_PDSM_MAX,
};
union nd_pdsm_payload {
  struct nd_papr_pdsm_health health;
  struct nd_papr_pdsm_smart_inject smart_inject;
  __u8 buf[ND_PDSM_PAYLOAD_MAX_SIZE];
} __attribute__((__packed__));
struct nd_pkg_pdsm {
  __s32 cmd_status;
  __u16 reserved[2];
  union nd_pdsm_payload payload;
} __attribute__((__packed__));
#endif
