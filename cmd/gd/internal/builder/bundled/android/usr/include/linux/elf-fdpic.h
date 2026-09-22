/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_ELF_FDPIC_H
#define _UAPI_LINUX_ELF_FDPIC_H
#include <linux/elf.h>
#define PT_GNU_STACK (PT_LOOS + 0x474e551)
struct elf32_fdpic_loadseg {
  Elf32_Addr addr;
  Elf32_Addr p_vaddr;
  Elf32_Word p_memsz;
};
struct elf32_fdpic_loadmap {
  Elf32_Half version;
  Elf32_Half nsegs;
  struct elf32_fdpic_loadseg segs[];
};
#define ELF32_FDPIC_LOADMAP_VERSION 0x0000
struct elf64_fdpic_loadseg {
  Elf64_Addr addr;
  Elf64_Addr p_vaddr;
  Elf64_Word p_memsz;
};
struct elf64_fdpic_loadmap {
  Elf64_Half version;
  Elf64_Half nsegs;
  struct elf64_fdpic_loadseg segs[];
};
#define ELF64_FDPIC_LOADMAP_VERSION 0x0000
#endif
