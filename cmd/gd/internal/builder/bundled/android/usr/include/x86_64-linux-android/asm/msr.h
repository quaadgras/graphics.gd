/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_ASM_X86_MSR_H
#define _UAPI_ASM_X86_MSR_H
#ifndef __ASSEMBLY__
#include <linux/types.h>
#include <linux/ioctl.h>
#define X86_IOC_RDMSR_REGS _IOWR('c', 0xA0, __u32[8])
#define X86_IOC_WRMSR_REGS _IOWR('c', 0xA1, __u32[8])
#endif
#endif
