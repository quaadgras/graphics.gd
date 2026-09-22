/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPILINUX_KEXEC_H
#define _UAPILINUX_KEXEC_H
#include <linux/types.h>
#define KEXEC_ON_CRASH 0x00000001
#define KEXEC_PRESERVE_CONTEXT 0x00000002
#define KEXEC_UPDATE_ELFCOREHDR 0x00000004
#define KEXEC_CRASH_HOTPLUG_SUPPORT 0x00000008
#define KEXEC_ARCH_MASK 0xffff0000
#define KEXEC_FILE_UNLOAD 0x00000001
#define KEXEC_FILE_ON_CRASH 0x00000002
#define KEXEC_FILE_NO_INITRAMFS 0x00000004
#define KEXEC_FILE_DEBUG 0x00000008
#define KEXEC_ARCH_DEFAULT (0 << 16)
#define KEXEC_ARCH_386 (3 << 16)
#define KEXEC_ARCH_68K (4 << 16)
#define KEXEC_ARCH_PARISC (15 << 16)
#define KEXEC_ARCH_X86_64 (62 << 16)
#define KEXEC_ARCH_PPC (20 << 16)
#define KEXEC_ARCH_PPC64 (21 << 16)
#define KEXEC_ARCH_IA_64 (50 << 16)
#define KEXEC_ARCH_ARM (40 << 16)
#define KEXEC_ARCH_S390 (22 << 16)
#define KEXEC_ARCH_SH (42 << 16)
#define KEXEC_ARCH_MIPS_LE (10 << 16)
#define KEXEC_ARCH_MIPS (8 << 16)
#define KEXEC_ARCH_AARCH64 (183 << 16)
#define KEXEC_ARCH_RISCV (243 << 16)
#define KEXEC_ARCH_LOONGARCH (258 << 16)
#define KEXEC_SEGMENT_MAX 16
struct kexec_segment {
  const void * buf;
  __kernel_size_t bufsz;
  const void * mem;
  __kernel_size_t memsz;
};
#endif
