/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_ASM_X86_E820_H
#define _UAPI_ASM_X86_E820_H
#define E820MAP 0x2d0
#define E820MAX 128
#define E820_X_MAX E820MAX
#define E820NR 0x1e8
#define E820_RAM 1
#define E820_RESERVED 2
#define E820_ACPI 3
#define E820_NVS 4
#define E820_UNUSABLE 5
#define E820_PMEM 7
#define E820_PRAM 12
#define E820_RESERVED_KERN 128
#ifndef __ASSEMBLY__
#include <linux/types.h>
struct e820entry {
  __u64 addr;
  __u64 size;
  __u32 type;
} __attribute__((packed));
struct e820map {
  __u32 nr_map;
  struct e820entry map[E820_X_MAX];
};
#define ISA_START_ADDRESS 0xa0000
#define ISA_END_ADDRESS 0x100000
#define BIOS_BEGIN 0x000a0000
#define BIOS_END 0x00100000
#define BIOS_ROM_BASE 0xffe00000
#define BIOS_ROM_END 0xffffffff
#endif
#endif
