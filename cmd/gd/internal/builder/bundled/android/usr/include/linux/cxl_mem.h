/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_CXL_MEM_H_
#define _UAPI_CXL_MEM_H_
#include <linux/types.h>
#define CXL_MEM_QUERY_COMMANDS _IOR(0xCE, 1, struct cxl_mem_query_commands)
#define CXL_MEM_SEND_COMMAND _IOWR(0xCE, 2, struct cxl_send_command)
#define CXL_CMDS ___C(INVALID, "Invalid Command"), ___C(IDENTIFY, "Identify Command"), ___C(RAW, "Raw device command"), ___C(GET_SUPPORTED_LOGS, "Get Supported Logs"), ___C(GET_FW_INFO, "Get FW Info"), ___C(GET_PARTITION_INFO, "Get Partition Information"), ___C(GET_LSA, "Get Label Storage Area"), ___C(GET_HEALTH_INFO, "Get Health Info"), ___C(GET_LOG, "Get Log"), ___C(SET_PARTITION_INFO, "Set Partition Information"), ___C(SET_LSA, "Set Label Storage Area"), ___C(GET_ALERT_CONFIG, "Get Alert Configuration"), ___C(SET_ALERT_CONFIG, "Set Alert Configuration"), ___C(GET_SHUTDOWN_STATE, "Get Shutdown State"), ___C(SET_SHUTDOWN_STATE, "Set Shutdown State"), ___DEPRECATED(GET_POISON, "Get Poison List"), ___DEPRECATED(INJECT_POISON, "Inject Poison"), ___DEPRECATED(CLEAR_POISON, "Clear Poison"), ___C(GET_SCAN_MEDIA_CAPS, "Get Scan Media Capabilities"), ___DEPRECATED(SCAN_MEDIA, "Scan Media"), ___DEPRECATED(GET_SCAN_MEDIA, "Get Scan Media Results"), ___C(GET_TIMESTAMP, "Get Timestamp"), ___C(GET_LOG_CAPS, "Get Log Capabilities"), ___C(CLEAR_LOG, "Clear Log"), ___C(GET_SUP_LOG_SUBLIST, "Get Supported Logs Sub-List"), ___C(MAX, "invalid / last command")
#define ___C(a,b) CXL_MEM_COMMAND_ID_ ##a
#define ___DEPRECATED(a,b) CXL_MEM_DEPRECATED_ID_ ##a
enum {
  CXL_CMDS
};
#undef ___C
#undef ___DEPRECATED
#define ___C(a,b) { b }
#define ___DEPRECATED(a,b) { "Deprecated " b }
static const struct {
  const char * name;
} cxl_command_names[] __attribute__((__unused__)) = {
  CXL_CMDS
};
#undef ___C
#undef ___DEPRECATED
#define ___C(a,b) (0)
#define ___DEPRECATED(a,b) (1)
#undef ___C
#undef ___DEPRECATED
struct cxl_command_info {
  __u32 id;
  __u32 flags;
#define CXL_MEM_COMMAND_FLAG_MASK GENMASK(1, 0)
#define CXL_MEM_COMMAND_FLAG_ENABLED BIT(0)
#define CXL_MEM_COMMAND_FLAG_EXCLUSIVE BIT(1)
  __u32 size_in;
  __u32 size_out;
};
struct cxl_mem_query_commands {
  __u32 n_commands;
  __u32 rsvd;
  struct cxl_command_info  commands[];
};
struct cxl_send_command {
  __u32 id;
  __u32 flags;
  union {
    struct {
      __u16 opcode;
      __u16 rsvd;
    } raw;
    __u32 rsvd;
  };
  __u32 retval;
  struct {
    __u32 size;
    __u32 rsvd;
    __u64 payload;
  } in;
  struct {
    __u32 size;
    __u32 rsvd;
    __u64 payload;
  } out;
};
#endif
