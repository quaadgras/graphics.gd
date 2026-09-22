/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef KFD_IOCTL_H_INCLUDED
#define KFD_IOCTL_H_INCLUDED
#include <drm/drm.h>
#include <linux/ioctl.h>
#define KFD_IOCTL_MAJOR_VERSION 1
#define KFD_IOCTL_MINOR_VERSION 17
struct kfd_ioctl_get_version_args {
  __u32 major_version;
  __u32 minor_version;
};
#define KFD_IOC_QUEUE_TYPE_COMPUTE 0x0
#define KFD_IOC_QUEUE_TYPE_SDMA 0x1
#define KFD_IOC_QUEUE_TYPE_COMPUTE_AQL 0x2
#define KFD_IOC_QUEUE_TYPE_SDMA_XGMI 0x3
#define KFD_IOC_QUEUE_TYPE_SDMA_BY_ENG_ID 0x4
#define KFD_MAX_QUEUE_PERCENTAGE 100
#define KFD_MAX_QUEUE_PRIORITY 15
struct kfd_ioctl_create_queue_args {
  __u64 ring_base_address;
  __u64 write_pointer_address;
  __u64 read_pointer_address;
  __u64 doorbell_offset;
  __u32 ring_size;
  __u32 gpu_id;
  __u32 queue_type;
  __u32 queue_percentage;
  __u32 queue_priority;
  __u32 queue_id;
  __u64 eop_buffer_address;
  __u64 eop_buffer_size;
  __u64 ctx_save_restore_address;
  __u32 ctx_save_restore_size;
  __u32 ctl_stack_size;
  __u32 sdma_engine_id;
  __u32 pad;
};
struct kfd_ioctl_destroy_queue_args {
  __u32 queue_id;
  __u32 pad;
};
struct kfd_ioctl_update_queue_args {
  __u64 ring_base_address;
  __u32 queue_id;
  __u32 ring_size;
  __u32 queue_percentage;
  __u32 queue_priority;
};
struct kfd_ioctl_set_cu_mask_args {
  __u32 queue_id;
  __u32 num_cu_mask;
  __u64 cu_mask_ptr;
};
struct kfd_ioctl_get_queue_wave_state_args {
  __u64 ctl_stack_address;
  __u32 ctl_stack_used_size;
  __u32 save_area_used_size;
  __u32 queue_id;
  __u32 pad;
};
struct kfd_ioctl_get_available_memory_args {
  __u64 available;
  __u32 gpu_id;
  __u32 pad;
};
struct kfd_dbg_device_info_entry {
  __u64 exception_status;
  __u64 lds_base;
  __u64 lds_limit;
  __u64 scratch_base;
  __u64 scratch_limit;
  __u64 gpuvm_base;
  __u64 gpuvm_limit;
  __u32 gpu_id;
  __u32 location_id;
  __u32 vendor_id;
  __u32 device_id;
  __u32 revision_id;
  __u32 subsystem_vendor_id;
  __u32 subsystem_device_id;
  __u32 fw_version;
  __u32 gfx_target_version;
  __u32 simd_count;
  __u32 max_waves_per_simd;
  __u32 array_count;
  __u32 simd_arrays_per_engine;
  __u32 num_xcc;
  __u32 capability;
  __u32 debug_prop;
};
#define KFD_IOC_CACHE_POLICY_COHERENT 0
#define KFD_IOC_CACHE_POLICY_NONCOHERENT 1
struct kfd_ioctl_set_memory_policy_args {
  __u64 alternate_aperture_base;
  __u64 alternate_aperture_size;
  __u32 gpu_id;
  __u32 default_policy;
  __u32 alternate_policy;
  __u32 pad;
};
struct kfd_ioctl_get_clock_counters_args {
  __u64 gpu_clock_counter;
  __u64 cpu_clock_counter;
  __u64 system_clock_counter;
  __u64 system_clock_freq;
  __u32 gpu_id;
  __u32 pad;
};
struct kfd_process_device_apertures {
  __u64 lds_base;
  __u64 lds_limit;
  __u64 scratch_base;
  __u64 scratch_limit;
  __u64 gpuvm_base;
  __u64 gpuvm_limit;
  __u32 gpu_id;
  __u32 pad;
};
#define NUM_OF_SUPPORTED_GPUS 7
struct kfd_ioctl_get_process_apertures_args {
  struct kfd_process_device_apertures process_apertures[NUM_OF_SUPPORTED_GPUS];
  __u32 num_of_nodes;
  __u32 pad;
};
struct kfd_ioctl_get_process_apertures_new_args {
  __u64 kfd_process_device_apertures_ptr;
  __u32 num_of_nodes;
  __u32 pad;
};
#define MAX_ALLOWED_NUM_POINTS 100
#define MAX_ALLOWED_AW_BUFF_SIZE 4096
#define MAX_ALLOWED_WAC_BUFF_SIZE 128
struct kfd_ioctl_dbg_register_args {
  __u32 gpu_id;
  __u32 pad;
};
struct kfd_ioctl_dbg_unregister_args {
  __u32 gpu_id;
  __u32 pad;
};
struct kfd_ioctl_dbg_address_watch_args {
  __u64 content_ptr;
  __u32 gpu_id;
  __u32 buf_size_in_bytes;
};
struct kfd_ioctl_dbg_wave_control_args {
  __u64 content_ptr;
  __u32 gpu_id;
  __u32 buf_size_in_bytes;
};
#define KFD_INVALID_FD 0xffffffff
#define KFD_IOC_EVENT_SIGNAL 0
#define KFD_IOC_EVENT_NODECHANGE 1
#define KFD_IOC_EVENT_DEVICESTATECHANGE 2
#define KFD_IOC_EVENT_HW_EXCEPTION 3
#define KFD_IOC_EVENT_SYSTEM_EVENT 4
#define KFD_IOC_EVENT_DEBUG_EVENT 5
#define KFD_IOC_EVENT_PROFILE_EVENT 6
#define KFD_IOC_EVENT_QUEUE_EVENT 7
#define KFD_IOC_EVENT_MEMORY 8
#define KFD_IOC_WAIT_RESULT_COMPLETE 0
#define KFD_IOC_WAIT_RESULT_TIMEOUT 1
#define KFD_IOC_WAIT_RESULT_FAIL 2
#define KFD_SIGNAL_EVENT_LIMIT 4096
#define KFD_HW_EXCEPTION_WHOLE_GPU_RESET 0
#define KFD_HW_EXCEPTION_PER_ENGINE_RESET 1
#define KFD_HW_EXCEPTION_GPU_HANG 0
#define KFD_HW_EXCEPTION_ECC 1
#define KFD_MEM_ERR_NO_RAS 0
#define KFD_MEM_ERR_SRAM_ECC 1
#define KFD_MEM_ERR_POISON_CONSUMED 2
#define KFD_MEM_ERR_GPU_HANG 3
struct kfd_ioctl_create_event_args {
  __u64 event_page_offset;
  __u32 event_trigger_data;
  __u32 event_type;
  __u32 auto_reset;
  __u32 node_id;
  __u32 event_id;
  __u32 event_slot_index;
};
struct kfd_ioctl_destroy_event_args {
  __u32 event_id;
  __u32 pad;
};
struct kfd_ioctl_set_event_args {
  __u32 event_id;
  __u32 pad;
};
struct kfd_ioctl_reset_event_args {
  __u32 event_id;
  __u32 pad;
};
struct kfd_memory_exception_failure {
  __u32 NotPresent;
  __u32 ReadOnly;
  __u32 NoExecute;
  __u32 imprecise;
};
struct kfd_hsa_memory_exception_data {
  struct kfd_memory_exception_failure failure;
  __u64 va;
  __u32 gpu_id;
  __u32 ErrorType;
};
struct kfd_hsa_hw_exception_data {
  __u32 reset_type;
  __u32 reset_cause;
  __u32 memory_lost;
  __u32 gpu_id;
};
struct kfd_hsa_signal_event_data {
  __u64 last_event_age;
};
struct kfd_event_data {
  union {
    struct kfd_hsa_memory_exception_data memory_exception_data;
    struct kfd_hsa_hw_exception_data hw_exception_data;
    struct kfd_hsa_signal_event_data signal_event_data;
  };
  __u64 kfd_event_data_ext;
  __u32 event_id;
  __u32 pad;
};
struct kfd_ioctl_wait_events_args {
  __u64 events_ptr;
  __u32 num_events;
  __u32 wait_for_all;
  __u32 timeout;
  __u32 wait_result;
};
struct kfd_ioctl_set_scratch_backing_va_args {
  __u64 va_addr;
  __u32 gpu_id;
  __u32 pad;
};
struct kfd_ioctl_get_tile_config_args {
  __u64 tile_config_ptr;
  __u64 macro_tile_config_ptr;
  __u32 num_tile_configs;
  __u32 num_macro_tile_configs;
  __u32 gpu_id;
  __u32 gb_addr_config;
  __u32 num_banks;
  __u32 num_ranks;
};
struct kfd_ioctl_set_trap_handler_args {
  __u64 tba_addr;
  __u64 tma_addr;
  __u32 gpu_id;
  __u32 pad;
};
struct kfd_ioctl_acquire_vm_args {
  __u32 drm_fd;
  __u32 gpu_id;
};
#define KFD_IOC_ALLOC_MEM_FLAGS_VRAM (1 << 0)
#define KFD_IOC_ALLOC_MEM_FLAGS_GTT (1 << 1)
#define KFD_IOC_ALLOC_MEM_FLAGS_USERPTR (1 << 2)
#define KFD_IOC_ALLOC_MEM_FLAGS_DOORBELL (1 << 3)
#define KFD_IOC_ALLOC_MEM_FLAGS_MMIO_REMAP (1 << 4)
#define KFD_IOC_ALLOC_MEM_FLAGS_WRITABLE (1 << 31)
#define KFD_IOC_ALLOC_MEM_FLAGS_EXECUTABLE (1 << 30)
#define KFD_IOC_ALLOC_MEM_FLAGS_PUBLIC (1 << 29)
#define KFD_IOC_ALLOC_MEM_FLAGS_NO_SUBSTITUTE (1 << 28)
#define KFD_IOC_ALLOC_MEM_FLAGS_AQL_QUEUE_MEM (1 << 27)
#define KFD_IOC_ALLOC_MEM_FLAGS_COHERENT (1 << 26)
#define KFD_IOC_ALLOC_MEM_FLAGS_UNCACHED (1 << 25)
#define KFD_IOC_ALLOC_MEM_FLAGS_EXT_COHERENT (1 << 24)
#define KFD_IOC_ALLOC_MEM_FLAGS_CONTIGUOUS (1 << 23)
struct kfd_ioctl_alloc_memory_of_gpu_args {
  __u64 va_addr;
  __u64 size;
  __u64 handle;
  __u64 mmap_offset;
  __u32 gpu_id;
  __u32 flags;
};
struct kfd_ioctl_free_memory_of_gpu_args {
  __u64 handle;
};
struct kfd_ioctl_map_memory_to_gpu_args {
  __u64 handle;
  __u64 device_ids_array_ptr;
  __u32 n_devices;
  __u32 n_success;
};
struct kfd_ioctl_unmap_memory_from_gpu_args {
  __u64 handle;
  __u64 device_ids_array_ptr;
  __u32 n_devices;
  __u32 n_success;
};
struct kfd_ioctl_alloc_queue_gws_args {
  __u32 queue_id;
  __u32 num_gws;
  __u32 first_gws;
  __u32 pad;
};
struct kfd_ioctl_get_dmabuf_info_args {
  __u64 size;
  __u64 metadata_ptr;
  __u32 metadata_size;
  __u32 gpu_id;
  __u32 flags;
  __u32 dmabuf_fd;
};
struct kfd_ioctl_import_dmabuf_args {
  __u64 va_addr;
  __u64 handle;
  __u32 gpu_id;
  __u32 dmabuf_fd;
};
struct kfd_ioctl_export_dmabuf_args {
  __u64 handle;
  __u32 flags;
  __u32 dmabuf_fd;
};
enum kfd_smi_event {
  KFD_SMI_EVENT_NONE = 0,
  KFD_SMI_EVENT_VMFAULT = 1,
  KFD_SMI_EVENT_THERMAL_THROTTLE = 2,
  KFD_SMI_EVENT_GPU_PRE_RESET = 3,
  KFD_SMI_EVENT_GPU_POST_RESET = 4,
  KFD_SMI_EVENT_MIGRATE_START = 5,
  KFD_SMI_EVENT_MIGRATE_END = 6,
  KFD_SMI_EVENT_PAGE_FAULT_START = 7,
  KFD_SMI_EVENT_PAGE_FAULT_END = 8,
  KFD_SMI_EVENT_QUEUE_EVICTION = 9,
  KFD_SMI_EVENT_QUEUE_RESTORE = 10,
  KFD_SMI_EVENT_UNMAP_FROM_GPU = 11,
  KFD_SMI_EVENT_ALL_PROCESS = 64
};
enum KFD_MIGRATE_TRIGGERS {
  KFD_MIGRATE_TRIGGER_PREFETCH,
  KFD_MIGRATE_TRIGGER_PAGEFAULT_GPU,
  KFD_MIGRATE_TRIGGER_PAGEFAULT_CPU,
  KFD_MIGRATE_TRIGGER_TTM_EVICTION
};
enum KFD_QUEUE_EVICTION_TRIGGERS {
  KFD_QUEUE_EVICTION_TRIGGER_SVM,
  KFD_QUEUE_EVICTION_TRIGGER_USERPTR,
  KFD_QUEUE_EVICTION_TRIGGER_TTM,
  KFD_QUEUE_EVICTION_TRIGGER_SUSPEND,
  KFD_QUEUE_EVICTION_CRIU_CHECKPOINT,
  KFD_QUEUE_EVICTION_CRIU_RESTORE
};
enum KFD_SVM_UNMAP_TRIGGERS {
  KFD_SVM_UNMAP_TRIGGER_MMU_NOTIFY,
  KFD_SVM_UNMAP_TRIGGER_MMU_NOTIFY_MIGRATE,
  KFD_SVM_UNMAP_TRIGGER_UNMAP_FROM_CPU
};
#define KFD_SMI_EVENT_MASK_FROM_INDEX(i) (1ULL << ((i) - 1))
#define KFD_SMI_EVENT_MSG_SIZE 96
struct kfd_ioctl_smi_events_args {
  __u32 gpuid;
  __u32 anon_fd;
};
#define KFD_EVENT_FMT_UPDATE_GPU_RESET(reset_seq_num,reset_cause) "%x %s\n", (reset_seq_num), (reset_cause)
#define KFD_EVENT_FMT_THERMAL_THROTTLING(bitmask,counter) "%llx:%llx\n", (bitmask), (counter)
#define KFD_EVENT_FMT_VMFAULT(pid,task_name) "%x:%s\n", (pid), (task_name)
#define KFD_EVENT_FMT_PAGEFAULT_START(ns,pid,addr,node,rw) "%lld -%d @%lx(%x) %c\n", (ns), (pid), (addr), (node), (rw)
#define KFD_EVENT_FMT_PAGEFAULT_END(ns,pid,addr,node,migrate_update) "%lld -%d @%lx(%x) %c\n", (ns), (pid), (addr), (node), (migrate_update)
#define KFD_EVENT_FMT_MIGRATE_START(ns,pid,start,size,from,to,prefetch_loc,preferred_loc,migrate_trigger) "%lld -%d @%lx(%lx) %x->%x %x:%x %d\n", (ns), (pid), (start), (size), (from), (to), (prefetch_loc), (preferred_loc), (migrate_trigger)
#define KFD_EVENT_FMT_MIGRATE_END(ns,pid,start,size,from,to,migrate_trigger) "%lld -%d @%lx(%lx) %x->%x %d\n", (ns), (pid), (start), (size), (from), (to), (migrate_trigger)
#define KFD_EVENT_FMT_QUEUE_EVICTION(ns,pid,node,evict_trigger) "%lld -%d %x %d\n", (ns), (pid), (node), (evict_trigger)
#define KFD_EVENT_FMT_QUEUE_RESTORE(ns,pid,node,rescheduled) "%lld -%d %x %c\n", (ns), (pid), (node), (rescheduled)
#define KFD_EVENT_FMT_UNMAP_FROM_GPU(ns,pid,addr,size,node,unmap_trigger) "%lld -%d @%lx(%lx) %x %d\n", (ns), (pid), (addr), (size), (node), (unmap_trigger)
enum kfd_criu_op {
  KFD_CRIU_OP_PROCESS_INFO,
  KFD_CRIU_OP_CHECKPOINT,
  KFD_CRIU_OP_UNPAUSE,
  KFD_CRIU_OP_RESTORE,
  KFD_CRIU_OP_RESUME,
};
struct kfd_ioctl_criu_args {
  __u64 devices;
  __u64 bos;
  __u64 priv_data;
  __u64 priv_data_size;
  __u32 num_devices;
  __u32 num_bos;
  __u32 num_objects;
  __u32 pid;
  __u32 op;
};
struct kfd_criu_device_bucket {
  __u32 user_gpu_id;
  __u32 actual_gpu_id;
  __u32 drm_fd;
  __u32 pad;
};
struct kfd_criu_bo_bucket {
  __u64 addr;
  __u64 size;
  __u64 offset;
  __u64 restored_offset;
  __u32 gpu_id;
  __u32 alloc_flags;
  __u32 dmabuf_fd;
  __u32 pad;
};
enum kfd_mmio_remap {
  KFD_MMIO_REMAP_HDP_MEM_FLUSH_CNTL = 0,
  KFD_MMIO_REMAP_HDP_REG_FLUSH_CNTL = 4,
};
#define KFD_IOCTL_SVM_FLAG_HOST_ACCESS 0x00000001
#define KFD_IOCTL_SVM_FLAG_COHERENT 0x00000002
#define KFD_IOCTL_SVM_FLAG_HIVE_LOCAL 0x00000004
#define KFD_IOCTL_SVM_FLAG_GPU_RO 0x00000008
#define KFD_IOCTL_SVM_FLAG_GPU_EXEC 0x00000010
#define KFD_IOCTL_SVM_FLAG_GPU_READ_MOSTLY 0x00000020
#define KFD_IOCTL_SVM_FLAG_GPU_ALWAYS_MAPPED 0x00000040
#define KFD_IOCTL_SVM_FLAG_EXT_COHERENT 0x00000080
enum kfd_ioctl_svm_op {
  KFD_IOCTL_SVM_OP_SET_ATTR,
  KFD_IOCTL_SVM_OP_GET_ATTR
};
enum kfd_ioctl_svm_location {
  KFD_IOCTL_SVM_LOCATION_SYSMEM = 0,
  KFD_IOCTL_SVM_LOCATION_UNDEFINED = 0xffffffff
};
enum kfd_ioctl_svm_attr_type {
  KFD_IOCTL_SVM_ATTR_PREFERRED_LOC,
  KFD_IOCTL_SVM_ATTR_PREFETCH_LOC,
  KFD_IOCTL_SVM_ATTR_ACCESS,
  KFD_IOCTL_SVM_ATTR_ACCESS_IN_PLACE,
  KFD_IOCTL_SVM_ATTR_NO_ACCESS,
  KFD_IOCTL_SVM_ATTR_SET_FLAGS,
  KFD_IOCTL_SVM_ATTR_CLR_FLAGS,
  KFD_IOCTL_SVM_ATTR_GRANULARITY
};
struct kfd_ioctl_svm_attribute {
  __u32 type;
  __u32 value;
};
struct kfd_ioctl_svm_args {
  __u64 start_addr;
  __u64 size;
  __u32 op;
  __u32 nattr;
  struct kfd_ioctl_svm_attribute attrs[];
};
struct kfd_ioctl_set_xnack_mode_args {
  __s32 xnack_enabled;
};
enum kfd_dbg_trap_override_mode {
  KFD_DBG_TRAP_OVERRIDE_OR = 0,
  KFD_DBG_TRAP_OVERRIDE_REPLACE = 1
};
enum kfd_dbg_trap_mask {
  KFD_DBG_TRAP_MASK_FP_INVALID = 1,
  KFD_DBG_TRAP_MASK_FP_INPUT_DENORMAL = 2,
  KFD_DBG_TRAP_MASK_FP_DIVIDE_BY_ZERO = 4,
  KFD_DBG_TRAP_MASK_FP_OVERFLOW = 8,
  KFD_DBG_TRAP_MASK_FP_UNDERFLOW = 16,
  KFD_DBG_TRAP_MASK_FP_INEXACT = 32,
  KFD_DBG_TRAP_MASK_INT_DIVIDE_BY_ZERO = 64,
  KFD_DBG_TRAP_MASK_DBG_ADDRESS_WATCH = 128,
  KFD_DBG_TRAP_MASK_DBG_MEMORY_VIOLATION = 256,
  KFD_DBG_TRAP_MASK_TRAP_ON_WAVE_START = (1 << 30),
  KFD_DBG_TRAP_MASK_TRAP_ON_WAVE_END = (1 << 31)
};
enum kfd_dbg_trap_wave_launch_mode {
  KFD_DBG_TRAP_WAVE_LAUNCH_MODE_NORMAL = 0,
  KFD_DBG_TRAP_WAVE_LAUNCH_MODE_HALT = 1,
  KFD_DBG_TRAP_WAVE_LAUNCH_MODE_DEBUG = 3
};
enum kfd_dbg_trap_address_watch_mode {
  KFD_DBG_TRAP_ADDRESS_WATCH_MODE_READ = 0,
  KFD_DBG_TRAP_ADDRESS_WATCH_MODE_NONREAD = 1,
  KFD_DBG_TRAP_ADDRESS_WATCH_MODE_ATOMIC = 2,
  KFD_DBG_TRAP_ADDRESS_WATCH_MODE_ALL = 3
};
enum kfd_dbg_trap_flags {
  KFD_DBG_TRAP_FLAG_SINGLE_MEM_OP = 1,
  KFD_DBG_TRAP_FLAG_SINGLE_ALU_OP = 2,
};
enum kfd_dbg_trap_exception_code {
  EC_NONE = 0,
  EC_QUEUE_WAVE_ABORT = 1,
  EC_QUEUE_WAVE_TRAP = 2,
  EC_QUEUE_WAVE_MATH_ERROR = 3,
  EC_QUEUE_WAVE_ILLEGAL_INSTRUCTION = 4,
  EC_QUEUE_WAVE_MEMORY_VIOLATION = 5,
  EC_QUEUE_WAVE_APERTURE_VIOLATION = 6,
  EC_QUEUE_PACKET_DISPATCH_DIM_INVALID = 16,
  EC_QUEUE_PACKET_DISPATCH_GROUP_SEGMENT_SIZE_INVALID = 17,
  EC_QUEUE_PACKET_DISPATCH_CODE_INVALID = 18,
  EC_QUEUE_PACKET_RESERVED = 19,
  EC_QUEUE_PACKET_UNSUPPORTED = 20,
  EC_QUEUE_PACKET_DISPATCH_WORK_GROUP_SIZE_INVALID = 21,
  EC_QUEUE_PACKET_DISPATCH_REGISTER_INVALID = 22,
  EC_QUEUE_PACKET_VENDOR_UNSUPPORTED = 23,
  EC_QUEUE_PREEMPTION_ERROR = 30,
  EC_QUEUE_NEW = 31,
  EC_DEVICE_QUEUE_DELETE = 32,
  EC_DEVICE_MEMORY_VIOLATION = 33,
  EC_DEVICE_RAS_ERROR = 34,
  EC_DEVICE_FATAL_HALT = 35,
  EC_DEVICE_NEW = 36,
  EC_PROCESS_RUNTIME = 48,
  EC_PROCESS_DEVICE_REMOVE = 49,
  EC_MAX
};
#define KFD_EC_MASK(ecode) (1ULL << (ecode - 1))
#define KFD_EC_MASK_QUEUE (KFD_EC_MASK(EC_QUEUE_WAVE_ABORT) | KFD_EC_MASK(EC_QUEUE_WAVE_TRAP) | KFD_EC_MASK(EC_QUEUE_WAVE_MATH_ERROR) | KFD_EC_MASK(EC_QUEUE_WAVE_ILLEGAL_INSTRUCTION) | KFD_EC_MASK(EC_QUEUE_WAVE_MEMORY_VIOLATION) | KFD_EC_MASK(EC_QUEUE_WAVE_APERTURE_VIOLATION) | KFD_EC_MASK(EC_QUEUE_PACKET_DISPATCH_DIM_INVALID) | KFD_EC_MASK(EC_QUEUE_PACKET_DISPATCH_GROUP_SEGMENT_SIZE_INVALID) | KFD_EC_MASK(EC_QUEUE_PACKET_DISPATCH_CODE_INVALID) | KFD_EC_MASK(EC_QUEUE_PACKET_RESERVED) | KFD_EC_MASK(EC_QUEUE_PACKET_UNSUPPORTED) | KFD_EC_MASK(EC_QUEUE_PACKET_DISPATCH_WORK_GROUP_SIZE_INVALID) | KFD_EC_MASK(EC_QUEUE_PACKET_DISPATCH_REGISTER_INVALID) | KFD_EC_MASK(EC_QUEUE_PACKET_VENDOR_UNSUPPORTED) | KFD_EC_MASK(EC_QUEUE_PREEMPTION_ERROR) | KFD_EC_MASK(EC_QUEUE_NEW))
#define KFD_EC_MASK_DEVICE (KFD_EC_MASK(EC_DEVICE_QUEUE_DELETE) | KFD_EC_MASK(EC_DEVICE_RAS_ERROR) | KFD_EC_MASK(EC_DEVICE_FATAL_HALT) | KFD_EC_MASK(EC_DEVICE_MEMORY_VIOLATION) | KFD_EC_MASK(EC_DEVICE_NEW))
#define KFD_EC_MASK_PROCESS (KFD_EC_MASK(EC_PROCESS_RUNTIME) | KFD_EC_MASK(EC_PROCESS_DEVICE_REMOVE))
#define KFD_EC_MASK_PACKET (KFD_EC_MASK(EC_QUEUE_PACKET_DISPATCH_DIM_INVALID) | KFD_EC_MASK(EC_QUEUE_PACKET_DISPATCH_GROUP_SEGMENT_SIZE_INVALID) | KFD_EC_MASK(EC_QUEUE_PACKET_DISPATCH_CODE_INVALID) | KFD_EC_MASK(EC_QUEUE_PACKET_RESERVED) | KFD_EC_MASK(EC_QUEUE_PACKET_UNSUPPORTED) | KFD_EC_MASK(EC_QUEUE_PACKET_DISPATCH_WORK_GROUP_SIZE_INVALID) | KFD_EC_MASK(EC_QUEUE_PACKET_DISPATCH_REGISTER_INVALID) | KFD_EC_MASK(EC_QUEUE_PACKET_VENDOR_UNSUPPORTED))
#define KFD_DBG_EC_IS_VALID(ecode) (ecode > EC_NONE && ecode < EC_MAX)
#define KFD_DBG_EC_TYPE_IS_QUEUE(ecode) (KFD_DBG_EC_IS_VALID(ecode) && ! ! (KFD_EC_MASK(ecode) & KFD_EC_MASK_QUEUE))
#define KFD_DBG_EC_TYPE_IS_DEVICE(ecode) (KFD_DBG_EC_IS_VALID(ecode) && ! ! (KFD_EC_MASK(ecode) & KFD_EC_MASK_DEVICE))
#define KFD_DBG_EC_TYPE_IS_PROCESS(ecode) (KFD_DBG_EC_IS_VALID(ecode) && ! ! (KFD_EC_MASK(ecode) & KFD_EC_MASK_PROCESS))
#define KFD_DBG_EC_TYPE_IS_PACKET(ecode) (KFD_DBG_EC_IS_VALID(ecode) && ! ! (KFD_EC_MASK(ecode) & KFD_EC_MASK_PACKET))
enum kfd_dbg_runtime_state {
  DEBUG_RUNTIME_STATE_DISABLED = 0,
  DEBUG_RUNTIME_STATE_ENABLED = 1,
  DEBUG_RUNTIME_STATE_ENABLED_BUSY = 2,
  DEBUG_RUNTIME_STATE_ENABLED_ERROR = 3
};
struct kfd_runtime_info {
  __u64 r_debug;
  __u32 runtime_state;
  __u32 ttmp_setup;
};
#define KFD_RUNTIME_ENABLE_MODE_ENABLE_MASK 1
#define KFD_RUNTIME_ENABLE_MODE_TTMP_SAVE_MASK 2
struct kfd_ioctl_runtime_enable_args {
  __u64 r_debug;
  __u32 mode_mask;
  __u32 capabilities_mask;
};
struct kfd_queue_snapshot_entry {
  __u64 exception_status;
  __u64 ring_base_address;
  __u64 write_pointer_address;
  __u64 read_pointer_address;
  __u64 ctx_save_restore_address;
  __u32 queue_id;
  __u32 gpu_id;
  __u32 ring_size;
  __u32 queue_type;
  __u32 ctx_save_restore_area_size;
  __u32 reserved;
};
#define KFD_DBG_QUEUE_ERROR_BIT 30
#define KFD_DBG_QUEUE_INVALID_BIT 31
#define KFD_DBG_QUEUE_ERROR_MASK (1 << KFD_DBG_QUEUE_ERROR_BIT)
#define KFD_DBG_QUEUE_INVALID_MASK (1 << KFD_DBG_QUEUE_INVALID_BIT)
struct kfd_context_save_area_header {
  struct {
    __u32 control_stack_offset;
    __u32 control_stack_size;
    __u32 wave_state_offset;
    __u32 wave_state_size;
  } wave_state;
  __u32 debug_offset;
  __u32 debug_size;
  __u64 err_payload_addr;
  __u32 err_event_id;
  __u32 reserved1;
};
enum kfd_dbg_trap_operations {
  KFD_IOC_DBG_TRAP_ENABLE = 0,
  KFD_IOC_DBG_TRAP_DISABLE = 1,
  KFD_IOC_DBG_TRAP_SEND_RUNTIME_EVENT = 2,
  KFD_IOC_DBG_TRAP_SET_EXCEPTIONS_ENABLED = 3,
  KFD_IOC_DBG_TRAP_SET_WAVE_LAUNCH_OVERRIDE = 4,
  KFD_IOC_DBG_TRAP_SET_WAVE_LAUNCH_MODE = 5,
  KFD_IOC_DBG_TRAP_SUSPEND_QUEUES = 6,
  KFD_IOC_DBG_TRAP_RESUME_QUEUES = 7,
  KFD_IOC_DBG_TRAP_SET_NODE_ADDRESS_WATCH = 8,
  KFD_IOC_DBG_TRAP_CLEAR_NODE_ADDRESS_WATCH = 9,
  KFD_IOC_DBG_TRAP_SET_FLAGS = 10,
  KFD_IOC_DBG_TRAP_QUERY_DEBUG_EVENT = 11,
  KFD_IOC_DBG_TRAP_QUERY_EXCEPTION_INFO = 12,
  KFD_IOC_DBG_TRAP_GET_QUEUE_SNAPSHOT = 13,
  KFD_IOC_DBG_TRAP_GET_DEVICE_SNAPSHOT = 14
};
struct kfd_ioctl_dbg_trap_enable_args {
  __u64 exception_mask;
  __u64 rinfo_ptr;
  __u32 rinfo_size;
  __u32 dbg_fd;
};
struct kfd_ioctl_dbg_trap_send_runtime_event_args {
  __u64 exception_mask;
  __u32 gpu_id;
  __u32 queue_id;
};
struct kfd_ioctl_dbg_trap_set_exceptions_enabled_args {
  __u64 exception_mask;
};
struct kfd_ioctl_dbg_trap_set_wave_launch_override_args {
  __u32 override_mode;
  __u32 enable_mask;
  __u32 support_request_mask;
  __u32 pad;
};
struct kfd_ioctl_dbg_trap_set_wave_launch_mode_args {
  __u32 launch_mode;
  __u32 pad;
};
struct kfd_ioctl_dbg_trap_suspend_queues_args {
  __u64 exception_mask;
  __u64 queue_array_ptr;
  __u32 num_queues;
  __u32 grace_period;
};
struct kfd_ioctl_dbg_trap_resume_queues_args {
  __u64 queue_array_ptr;
  __u32 num_queues;
  __u32 pad;
};
struct kfd_ioctl_dbg_trap_set_node_address_watch_args {
  __u64 address;
  __u32 mode;
  __u32 mask;
  __u32 gpu_id;
  __u32 id;
};
struct kfd_ioctl_dbg_trap_clear_node_address_watch_args {
  __u32 gpu_id;
  __u32 id;
};
struct kfd_ioctl_dbg_trap_set_flags_args {
  __u32 flags;
  __u32 pad;
};
struct kfd_ioctl_dbg_trap_query_debug_event_args {
  __u64 exception_mask;
  __u32 gpu_id;
  __u32 queue_id;
};
struct kfd_ioctl_dbg_trap_query_exception_info_args {
  __u64 info_ptr;
  __u32 info_size;
  __u32 source_id;
  __u32 exception_code;
  __u32 clear_exception;
};
struct kfd_ioctl_dbg_trap_queue_snapshot_args {
  __u64 exception_mask;
  __u64 snapshot_buf_ptr;
  __u32 num_queues;
  __u32 entry_size;
};
struct kfd_ioctl_dbg_trap_device_snapshot_args {
  __u64 exception_mask;
  __u64 snapshot_buf_ptr;
  __u32 num_devices;
  __u32 entry_size;
};
struct kfd_ioctl_dbg_trap_args {
  __u32 pid;
  __u32 op;
  union {
    struct kfd_ioctl_dbg_trap_enable_args enable;
    struct kfd_ioctl_dbg_trap_send_runtime_event_args send_runtime_event;
    struct kfd_ioctl_dbg_trap_set_exceptions_enabled_args set_exceptions_enabled;
    struct kfd_ioctl_dbg_trap_set_wave_launch_override_args launch_override;
    struct kfd_ioctl_dbg_trap_set_wave_launch_mode_args launch_mode;
    struct kfd_ioctl_dbg_trap_suspend_queues_args suspend_queues;
    struct kfd_ioctl_dbg_trap_resume_queues_args resume_queues;
    struct kfd_ioctl_dbg_trap_set_node_address_watch_args set_node_address_watch;
    struct kfd_ioctl_dbg_trap_clear_node_address_watch_args clear_node_address_watch;
    struct kfd_ioctl_dbg_trap_set_flags_args set_flags;
    struct kfd_ioctl_dbg_trap_query_debug_event_args query_debug_event;
    struct kfd_ioctl_dbg_trap_query_exception_info_args query_exception_info;
    struct kfd_ioctl_dbg_trap_queue_snapshot_args queue_snapshot;
    struct kfd_ioctl_dbg_trap_device_snapshot_args device_snapshot;
  };
};
#define AMDKFD_IOCTL_BASE 'K'
#define AMDKFD_IO(nr) _IO(AMDKFD_IOCTL_BASE, nr)
#define AMDKFD_IOR(nr,type) _IOR(AMDKFD_IOCTL_BASE, nr, type)
#define AMDKFD_IOW(nr,type) _IOW(AMDKFD_IOCTL_BASE, nr, type)
#define AMDKFD_IOWR(nr,type) _IOWR(AMDKFD_IOCTL_BASE, nr, type)
#define AMDKFD_IOC_GET_VERSION AMDKFD_IOR(0x01, struct kfd_ioctl_get_version_args)
#define AMDKFD_IOC_CREATE_QUEUE AMDKFD_IOWR(0x02, struct kfd_ioctl_create_queue_args)
#define AMDKFD_IOC_DESTROY_QUEUE AMDKFD_IOWR(0x03, struct kfd_ioctl_destroy_queue_args)
#define AMDKFD_IOC_SET_MEMORY_POLICY AMDKFD_IOW(0x04, struct kfd_ioctl_set_memory_policy_args)
#define AMDKFD_IOC_GET_CLOCK_COUNTERS AMDKFD_IOWR(0x05, struct kfd_ioctl_get_clock_counters_args)
#define AMDKFD_IOC_GET_PROCESS_APERTURES AMDKFD_IOR(0x06, struct kfd_ioctl_get_process_apertures_args)
#define AMDKFD_IOC_UPDATE_QUEUE AMDKFD_IOW(0x07, struct kfd_ioctl_update_queue_args)
#define AMDKFD_IOC_CREATE_EVENT AMDKFD_IOWR(0x08, struct kfd_ioctl_create_event_args)
#define AMDKFD_IOC_DESTROY_EVENT AMDKFD_IOW(0x09, struct kfd_ioctl_destroy_event_args)
#define AMDKFD_IOC_SET_EVENT AMDKFD_IOW(0x0A, struct kfd_ioctl_set_event_args)
#define AMDKFD_IOC_RESET_EVENT AMDKFD_IOW(0x0B, struct kfd_ioctl_reset_event_args)
#define AMDKFD_IOC_WAIT_EVENTS AMDKFD_IOWR(0x0C, struct kfd_ioctl_wait_events_args)
#define AMDKFD_IOC_DBG_REGISTER_DEPRECATED AMDKFD_IOW(0x0D, struct kfd_ioctl_dbg_register_args)
#define AMDKFD_IOC_DBG_UNREGISTER_DEPRECATED AMDKFD_IOW(0x0E, struct kfd_ioctl_dbg_unregister_args)
#define AMDKFD_IOC_DBG_ADDRESS_WATCH_DEPRECATED AMDKFD_IOW(0x0F, struct kfd_ioctl_dbg_address_watch_args)
#define AMDKFD_IOC_DBG_WAVE_CONTROL_DEPRECATED AMDKFD_IOW(0x10, struct kfd_ioctl_dbg_wave_control_args)
#define AMDKFD_IOC_SET_SCRATCH_BACKING_VA AMDKFD_IOWR(0x11, struct kfd_ioctl_set_scratch_backing_va_args)
#define AMDKFD_IOC_GET_TILE_CONFIG AMDKFD_IOWR(0x12, struct kfd_ioctl_get_tile_config_args)
#define AMDKFD_IOC_SET_TRAP_HANDLER AMDKFD_IOW(0x13, struct kfd_ioctl_set_trap_handler_args)
#define AMDKFD_IOC_GET_PROCESS_APERTURES_NEW AMDKFD_IOWR(0x14, struct kfd_ioctl_get_process_apertures_new_args)
#define AMDKFD_IOC_ACQUIRE_VM AMDKFD_IOW(0x15, struct kfd_ioctl_acquire_vm_args)
#define AMDKFD_IOC_ALLOC_MEMORY_OF_GPU AMDKFD_IOWR(0x16, struct kfd_ioctl_alloc_memory_of_gpu_args)
#define AMDKFD_IOC_FREE_MEMORY_OF_GPU AMDKFD_IOW(0x17, struct kfd_ioctl_free_memory_of_gpu_args)
#define AMDKFD_IOC_MAP_MEMORY_TO_GPU AMDKFD_IOWR(0x18, struct kfd_ioctl_map_memory_to_gpu_args)
#define AMDKFD_IOC_UNMAP_MEMORY_FROM_GPU AMDKFD_IOWR(0x19, struct kfd_ioctl_unmap_memory_from_gpu_args)
#define AMDKFD_IOC_SET_CU_MASK AMDKFD_IOW(0x1A, struct kfd_ioctl_set_cu_mask_args)
#define AMDKFD_IOC_GET_QUEUE_WAVE_STATE AMDKFD_IOWR(0x1B, struct kfd_ioctl_get_queue_wave_state_args)
#define AMDKFD_IOC_GET_DMABUF_INFO AMDKFD_IOWR(0x1C, struct kfd_ioctl_get_dmabuf_info_args)
#define AMDKFD_IOC_IMPORT_DMABUF AMDKFD_IOWR(0x1D, struct kfd_ioctl_import_dmabuf_args)
#define AMDKFD_IOC_ALLOC_QUEUE_GWS AMDKFD_IOWR(0x1E, struct kfd_ioctl_alloc_queue_gws_args)
#define AMDKFD_IOC_SMI_EVENTS AMDKFD_IOWR(0x1F, struct kfd_ioctl_smi_events_args)
#define AMDKFD_IOC_SVM AMDKFD_IOWR(0x20, struct kfd_ioctl_svm_args)
#define AMDKFD_IOC_SET_XNACK_MODE AMDKFD_IOWR(0x21, struct kfd_ioctl_set_xnack_mode_args)
#define AMDKFD_IOC_CRIU_OP AMDKFD_IOWR(0x22, struct kfd_ioctl_criu_args)
#define AMDKFD_IOC_AVAILABLE_MEMORY AMDKFD_IOWR(0x23, struct kfd_ioctl_get_available_memory_args)
#define AMDKFD_IOC_EXPORT_DMABUF AMDKFD_IOWR(0x24, struct kfd_ioctl_export_dmabuf_args)
#define AMDKFD_IOC_RUNTIME_ENABLE AMDKFD_IOWR(0x25, struct kfd_ioctl_runtime_enable_args)
#define AMDKFD_IOC_DBG_TRAP AMDKFD_IOWR(0x26, struct kfd_ioctl_dbg_trap_args)
#define AMDKFD_COMMAND_START 0x01
#define AMDKFD_COMMAND_END 0x27
#endif
