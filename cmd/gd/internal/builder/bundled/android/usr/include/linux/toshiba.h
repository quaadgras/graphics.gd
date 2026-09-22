/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_TOSHIBA_H
#define _UAPI_LINUX_TOSHIBA_H
#define TOSH_PROC "/proc/toshiba"
#define TOSH_DEVICE "/dev/toshiba"
#define TOSHIBA_ACPI_PROC "/proc/acpi/toshiba"
#define TOSHIBA_ACPI_DEVICE "/dev/toshiba_acpi"
typedef struct {
  unsigned int eax;
  unsigned int ebx __attribute__((packed));
  unsigned int ecx __attribute__((packed));
  unsigned int edx __attribute__((packed));
  unsigned int esi __attribute__((packed));
  unsigned int edi __attribute__((packed));
} SMMRegisters;
#define TOSH_SMM _IOWR('t', 0x90, SMMRegisters)
#define TOSHIBA_ACPI_SCI _IOWR('t', 0x91, SMMRegisters)
#endif
