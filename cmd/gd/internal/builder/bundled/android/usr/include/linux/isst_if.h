/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __ISST_IF_H
#define __ISST_IF_H
#include <linux/types.h>
struct isst_if_platform_info {
  __u16 api_version;
  __u16 driver_version;
  __u16 max_cmds_per_ioctl;
  __u8 mbox_supported;
  __u8 mmio_supported;
};
struct isst_if_cpu_map {
  __u32 logical_cpu;
  __u32 physical_cpu;
};
struct isst_if_cpu_maps {
  __u32 cmd_count;
  struct isst_if_cpu_map cpu_map[1];
};
struct isst_if_io_reg {
  __u32 read_write;
  __u32 logical_cpu;
  __u32 reg;
  __u32 value;
};
struct isst_if_io_regs {
  __u32 req_count;
  struct isst_if_io_reg io_reg[1];
};
struct isst_if_mbox_cmd {
  __u32 logical_cpu;
  __u32 parameter;
  __u32 req_data;
  __u32 resp_data;
  __u16 command;
  __u16 sub_command;
  __u32 reserved;
};
struct isst_if_mbox_cmds {
  __u32 cmd_count;
  struct isst_if_mbox_cmd mbox_cmd[1];
};
struct isst_if_msr_cmd {
  __u32 read_write;
  __u32 logical_cpu;
  __u64 msr;
  __u64 data;
};
struct isst_if_msr_cmds {
  __u32 cmd_count;
  struct isst_if_msr_cmd msr_cmd[1];
};
struct isst_core_power {
  __u8 get_set;
  __u8 socket_id;
  __u8 power_domain_id;
  __u8 enable;
  __u8 supported;
  __u8 priority_type;
};
struct isst_clos_param {
  __u8 get_set;
  __u8 socket_id;
  __u8 power_domain_id;
  __u8 clos;
  __u16 min_freq_mhz;
  __u16 max_freq_mhz;
  __u8 prop_prio;
};
struct isst_if_clos_assoc {
  __u8 socket_id;
  __u8 power_domain_id;
  __u16 logical_cpu;
  __u16 clos;
};
struct isst_if_clos_assoc_cmds {
  __u16 cmd_count;
  __u16 get_set;
  __u16 punit_cpu_map;
  struct isst_if_clos_assoc assoc_info[1];
};
struct isst_tpmi_instance_count {
  __u8 socket_id;
  __u8 count;
  __u16 valid_mask;
};
struct isst_perf_level_info {
  __u8 socket_id;
  __u8 power_domain_id;
  __u8 max_level;
  __u8 feature_rev;
  __u8 level_mask;
  __u8 current_level;
  __u8 feature_state;
  __u8 locked;
  __u8 enabled;
  __u8 sst_tf_support;
  __u8 sst_bf_support;
};
struct isst_perf_level_control {
  __u8 socket_id;
  __u8 power_domain_id;
  __u8 level;
};
struct isst_perf_feature_control {
  __u8 socket_id;
  __u8 power_domain_id;
  __u8 feature;
};
#define TRL_MAX_BUCKETS 8
#define TRL_MAX_LEVELS 6
struct isst_perf_level_data_info {
  __u8 socket_id;
  __u8 power_domain_id;
  __u16 level;
  __u16 tdp_ratio;
  __u16 base_freq_mhz;
  __u16 base_freq_avx2_mhz;
  __u16 base_freq_avx512_mhz;
  __u16 base_freq_amx_mhz;
  __u16 thermal_design_power_w;
  __u16 tjunction_max_c;
  __u16 max_memory_freq_mhz;
  __u16 cooling_type;
  __u16 p0_freq_mhz;
  __u16 p1_freq_mhz;
  __u16 pn_freq_mhz;
  __u16 pm_freq_mhz;
  __u16 p0_fabric_freq_mhz;
  __u16 p1_fabric_freq_mhz;
  __u16 pn_fabric_freq_mhz;
  __u16 pm_fabric_freq_mhz;
  __u16 max_buckets;
  __u16 max_trl_levels;
  __u16 bucket_core_counts[TRL_MAX_BUCKETS];
  __u16 trl_freq_mhz[TRL_MAX_LEVELS][TRL_MAX_BUCKETS];
};
struct isst_perf_level_cpu_mask {
  __u8 socket_id;
  __u8 power_domain_id;
  __u8 level;
  __u8 punit_cpu_map;
  __u64 mask;
  __u16 cpu_buffer_size;
  __s8 cpu_buffer[1];
};
struct isst_base_freq_info {
  __u8 socket_id;
  __u8 power_domain_id;
  __u16 level;
  __u16 high_base_freq_mhz;
  __u16 low_base_freq_mhz;
  __u16 tjunction_max_c;
  __u16 thermal_design_power_w;
};
struct isst_turbo_freq_info {
  __u8 socket_id;
  __u8 power_domain_id;
  __u16 level;
  __u16 max_clip_freqs;
  __u16 max_buckets;
  __u16 max_trl_levels;
  __u16 lp_clip_freq_mhz[TRL_MAX_LEVELS];
  __u16 bucket_core_counts[TRL_MAX_BUCKETS];
  __u16 trl_freq_mhz[TRL_MAX_LEVELS][TRL_MAX_BUCKETS];
};
#define ISST_IF_MAGIC 0xFE
#define ISST_IF_GET_PLATFORM_INFO _IOR(ISST_IF_MAGIC, 0, struct isst_if_platform_info *)
#define ISST_IF_GET_PHY_ID _IOWR(ISST_IF_MAGIC, 1, struct isst_if_cpu_map *)
#define ISST_IF_IO_CMD _IOW(ISST_IF_MAGIC, 2, struct isst_if_io_regs *)
#define ISST_IF_MBOX_COMMAND _IOWR(ISST_IF_MAGIC, 3, struct isst_if_mbox_cmds *)
#define ISST_IF_MSR_COMMAND _IOWR(ISST_IF_MAGIC, 4, struct isst_if_msr_cmds *)
#define ISST_IF_COUNT_TPMI_INSTANCES _IOR(ISST_IF_MAGIC, 5, struct isst_tpmi_instance_count *)
#define ISST_IF_CORE_POWER_STATE _IOWR(ISST_IF_MAGIC, 6, struct isst_core_power *)
#define ISST_IF_CLOS_PARAM _IOWR(ISST_IF_MAGIC, 7, struct isst_clos_param *)
#define ISST_IF_CLOS_ASSOC _IOWR(ISST_IF_MAGIC, 8, struct isst_if_clos_assoc_cmds *)
#define ISST_IF_PERF_LEVELS _IOWR(ISST_IF_MAGIC, 9, struct isst_perf_level_info *)
#define ISST_IF_PERF_SET_LEVEL _IOW(ISST_IF_MAGIC, 10, struct isst_perf_level_control *)
#define ISST_IF_PERF_SET_FEATURE _IOW(ISST_IF_MAGIC, 11, struct isst_perf_feature_control *)
#define ISST_IF_GET_PERF_LEVEL_INFO _IOR(ISST_IF_MAGIC, 12, struct isst_perf_level_data_info *)
#define ISST_IF_GET_PERF_LEVEL_CPU_MASK _IOR(ISST_IF_MAGIC, 13, struct isst_perf_level_cpu_mask *)
#define ISST_IF_GET_BASE_FREQ_INFO _IOR(ISST_IF_MAGIC, 14, struct isst_base_freq_info *)
#define ISST_IF_GET_BASE_FREQ_CPU_MASK _IOR(ISST_IF_MAGIC, 15, struct isst_perf_level_cpu_mask *)
#define ISST_IF_GET_TURBO_FREQ_INFO _IOR(ISST_IF_MAGIC, 16, struct isst_turbo_freq_info *)
#endif
