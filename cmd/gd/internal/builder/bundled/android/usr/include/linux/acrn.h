/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_ACRN_H
#define _UAPI_ACRN_H
#include <linux/types.h>
#define ACRN_IO_REQUEST_MAX 16
#define ACRN_IOREQ_STATE_PENDING 0
#define ACRN_IOREQ_STATE_COMPLETE 1
#define ACRN_IOREQ_STATE_PROCESSING 2
#define ACRN_IOREQ_STATE_FREE 3
#define ACRN_IOREQ_TYPE_PORTIO 0
#define ACRN_IOREQ_TYPE_MMIO 1
#define ACRN_IOREQ_TYPE_PCICFG 2
#define ACRN_IOREQ_DIR_READ 0
#define ACRN_IOREQ_DIR_WRITE 1
struct acrn_mmio_request {
  __u32 direction;
  __u32 reserved;
  __u64 address;
  __u64 size;
  __u64 value;
};
struct acrn_pio_request {
  __u32 direction;
  __u32 reserved;
  __u64 address;
  __u64 size;
  __u32 value;
};
struct acrn_pci_request {
  __u32 direction;
  __u32 reserved[3];
  __u64 size;
  __u32 value;
  __u32 bus;
  __u32 dev;
  __u32 func;
  __u32 reg;
};
struct acrn_io_request {
  __u32 type;
  __u32 completion_polling;
  __u32 reserved0[14];
  union {
    struct acrn_pio_request pio_request;
    struct acrn_pci_request pci_request;
    struct acrn_mmio_request mmio_request;
    __u64 data[8];
  } reqs;
  __u32 reserved1;
  __u32 kernel_handled;
  __u32 processed;
} __attribute__((aligned(256)));
struct acrn_io_request_buffer {
  union {
    struct acrn_io_request req_slot[ACRN_IO_REQUEST_MAX];
    __u8 reserved[4096];
  };
};
struct acrn_ioreq_notify {
  __u16 vmid;
  __u16 reserved;
  __u32 vcpu;
};
struct acrn_vm_creation {
  __u16 vmid;
  __u16 reserved0;
  __u16 vcpu_num;
  __u16 reserved1;
  __u8 uuid[16];
  __u64 vm_flag;
  __u64 ioreq_buf;
  __u64 cpu_affinity;
};
struct acrn_gp_regs {
  __le64 rax;
  __le64 rcx;
  __le64 rdx;
  __le64 rbx;
  __le64 rsp;
  __le64 rbp;
  __le64 rsi;
  __le64 rdi;
  __le64 r8;
  __le64 r9;
  __le64 r10;
  __le64 r11;
  __le64 r12;
  __le64 r13;
  __le64 r14;
  __le64 r15;
};
struct acrn_descriptor_ptr {
  __le16 limit;
  __le64 base;
  __le16 reserved[3];
} __attribute__((__packed__));
struct acrn_regs {
  struct acrn_gp_regs gprs;
  struct acrn_descriptor_ptr gdt;
  struct acrn_descriptor_ptr idt;
  __le64 rip;
  __le64 cs_base;
  __le64 cr0;
  __le64 cr4;
  __le64 cr3;
  __le64 ia32_efer;
  __le64 rflags;
  __le64 reserved_64[4];
  __le32 cs_ar;
  __le32 cs_limit;
  __le32 reserved_32[3];
  __le16 cs_sel;
  __le16 ss_sel;
  __le16 ds_sel;
  __le16 es_sel;
  __le16 fs_sel;
  __le16 gs_sel;
  __le16 ldt_sel;
  __le16 tr_sel;
};
struct acrn_vcpu_regs {
  __u16 vcpu_id;
  __u16 reserved[3];
  struct acrn_regs vcpu_regs;
};
#define ACRN_MEM_ACCESS_RIGHT_MASK 0x00000007U
#define ACRN_MEM_ACCESS_READ 0x00000001U
#define ACRN_MEM_ACCESS_WRITE 0x00000002U
#define ACRN_MEM_ACCESS_EXEC 0x00000004U
#define ACRN_MEM_ACCESS_RWX (ACRN_MEM_ACCESS_READ | ACRN_MEM_ACCESS_WRITE | ACRN_MEM_ACCESS_EXEC)
#define ACRN_MEM_TYPE_MASK 0x000007C0U
#define ACRN_MEM_TYPE_WB 0x00000040U
#define ACRN_MEM_TYPE_WT 0x00000080U
#define ACRN_MEM_TYPE_UC 0x00000100U
#define ACRN_MEM_TYPE_WC 0x00000200U
#define ACRN_MEM_TYPE_WP 0x00000400U
#define ACRN_MEMMAP_RAM 0
#define ACRN_MEMMAP_MMIO 1
struct acrn_vm_memmap {
  __u32 type;
  __u32 attr;
  __u64 user_vm_pa;
  union {
    __u64 service_vm_pa;
    __u64 vma_base;
  };
  __u64 len;
};
#define ACRN_PTDEV_IRQ_INTX 0
#define ACRN_PTDEV_IRQ_MSI 1
#define ACRN_PTDEV_IRQ_MSIX 2
struct acrn_ptdev_irq {
  __u32 type;
  __u16 virt_bdf;
  __u16 phys_bdf;
  struct {
    __u32 virt_pin;
    __u32 phys_pin;
    __u32 is_pic_pin;
  } intx;
};
#define ACRN_PTDEV_QUIRK_ASSIGN (1U << 0)
#define ACRN_MMIODEV_RES_NUM 3
#define ACRN_PCI_NUM_BARS 6
struct acrn_pcidev {
  __u32 type;
  __u16 virt_bdf;
  __u16 phys_bdf;
  __u8 intr_line;
  __u8 intr_pin;
  __u32 bar[ACRN_PCI_NUM_BARS];
};
struct acrn_mmiodev {
  __u8 name[8];
  struct {
    __u64 user_vm_pa;
    __u64 service_vm_pa;
    __u64 size;
    __u64 mem_type;
  } res[ACRN_MMIODEV_RES_NUM];
};
struct acrn_vdev {
  union {
    __u64 value;
    struct {
      __le16 vendor;
      __le16 device;
      __le32 legacy_id;
    } fields;
  } id;
  __u64 slot;
  __u32 io_addr[ACRN_PCI_NUM_BARS];
  __u32 io_size[ACRN_PCI_NUM_BARS];
  __u8 args[128];
};
struct acrn_msi_entry {
  __u64 msi_addr;
  __u64 msi_data;
};
struct acrn_acpi_generic_address {
  __u8 space_id;
  __u8 bit_width;
  __u8 bit_offset;
  __u8 access_size;
  __u64 address;
} __attribute__((__packed__));
struct acrn_cstate_data {
  struct acrn_acpi_generic_address cx_reg;
  __u8 type;
  __u32 latency;
  __u64 power;
};
struct acrn_pstate_data {
  __u64 core_frequency;
  __u64 power;
  __u64 transition_latency;
  __u64 bus_master_latency;
  __u64 control;
  __u64 status;
};
#define PMCMD_TYPE_MASK 0x000000ff
enum acrn_pm_cmd_type {
  ACRN_PMCMD_GET_PX_CNT,
  ACRN_PMCMD_GET_PX_DATA,
  ACRN_PMCMD_GET_CX_CNT,
  ACRN_PMCMD_GET_CX_DATA,
};
#define ACRN_IOEVENTFD_FLAG_PIO 0x01
#define ACRN_IOEVENTFD_FLAG_DATAMATCH 0x02
#define ACRN_IOEVENTFD_FLAG_DEASSIGN 0x04
struct acrn_ioeventfd {
  __u32 fd;
  __u32 flags;
  __u64 addr;
  __u32 len;
  __u32 reserved;
  __u64 data;
};
#define ACRN_IRQFD_FLAG_DEASSIGN 0x01
struct acrn_irqfd {
  __s32 fd;
  __u32 flags;
  struct acrn_msi_entry msi;
};
#define ACRN_IOCTL_TYPE 0xA2
#define ACRN_IOCTL_CREATE_VM _IOWR(ACRN_IOCTL_TYPE, 0x10, struct acrn_vm_creation)
#define ACRN_IOCTL_DESTROY_VM _IO(ACRN_IOCTL_TYPE, 0x11)
#define ACRN_IOCTL_START_VM _IO(ACRN_IOCTL_TYPE, 0x12)
#define ACRN_IOCTL_PAUSE_VM _IO(ACRN_IOCTL_TYPE, 0x13)
#define ACRN_IOCTL_RESET_VM _IO(ACRN_IOCTL_TYPE, 0x15)
#define ACRN_IOCTL_SET_VCPU_REGS _IOW(ACRN_IOCTL_TYPE, 0x16, struct acrn_vcpu_regs)
#define ACRN_IOCTL_INJECT_MSI _IOW(ACRN_IOCTL_TYPE, 0x23, struct acrn_msi_entry)
#define ACRN_IOCTL_VM_INTR_MONITOR _IOW(ACRN_IOCTL_TYPE, 0x24, unsigned long)
#define ACRN_IOCTL_SET_IRQLINE _IOW(ACRN_IOCTL_TYPE, 0x25, __u64)
#define ACRN_IOCTL_NOTIFY_REQUEST_FINISH _IOW(ACRN_IOCTL_TYPE, 0x31, struct acrn_ioreq_notify)
#define ACRN_IOCTL_CREATE_IOREQ_CLIENT _IO(ACRN_IOCTL_TYPE, 0x32)
#define ACRN_IOCTL_ATTACH_IOREQ_CLIENT _IO(ACRN_IOCTL_TYPE, 0x33)
#define ACRN_IOCTL_DESTROY_IOREQ_CLIENT _IO(ACRN_IOCTL_TYPE, 0x34)
#define ACRN_IOCTL_CLEAR_VM_IOREQ _IO(ACRN_IOCTL_TYPE, 0x35)
#define ACRN_IOCTL_SET_MEMSEG _IOW(ACRN_IOCTL_TYPE, 0x41, struct acrn_vm_memmap)
#define ACRN_IOCTL_UNSET_MEMSEG _IOW(ACRN_IOCTL_TYPE, 0x42, struct acrn_vm_memmap)
#define ACRN_IOCTL_SET_PTDEV_INTR _IOW(ACRN_IOCTL_TYPE, 0x53, struct acrn_ptdev_irq)
#define ACRN_IOCTL_RESET_PTDEV_INTR _IOW(ACRN_IOCTL_TYPE, 0x54, struct acrn_ptdev_irq)
#define ACRN_IOCTL_ASSIGN_PCIDEV _IOW(ACRN_IOCTL_TYPE, 0x55, struct acrn_pcidev)
#define ACRN_IOCTL_DEASSIGN_PCIDEV _IOW(ACRN_IOCTL_TYPE, 0x56, struct acrn_pcidev)
#define ACRN_IOCTL_ASSIGN_MMIODEV _IOW(ACRN_IOCTL_TYPE, 0x57, struct acrn_mmiodev)
#define ACRN_IOCTL_DEASSIGN_MMIODEV _IOW(ACRN_IOCTL_TYPE, 0x58, struct acrn_mmiodev)
#define ACRN_IOCTL_CREATE_VDEV _IOW(ACRN_IOCTL_TYPE, 0x59, struct acrn_vdev)
#define ACRN_IOCTL_DESTROY_VDEV _IOW(ACRN_IOCTL_TYPE, 0x5A, struct acrn_vdev)
#define ACRN_IOCTL_PM_GET_CPU_STATE _IOWR(ACRN_IOCTL_TYPE, 0x60, __u64)
#define ACRN_IOCTL_IOEVENTFD _IOW(ACRN_IOCTL_TYPE, 0x70, struct acrn_ioeventfd)
#define ACRN_IOCTL_IRQFD _IOW(ACRN_IOCTL_TYPE, 0x71, struct acrn_irqfd)
#endif
