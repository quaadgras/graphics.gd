/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI__ASM_MMAN_H
#define _UAPI__ASM_MMAN_H
#include <asm-generic/mman.h>
#define PROT_BTI 0x10
#define PROT_MTE 0x20
#define PKEY_DISABLE_EXECUTE 0x4
#define PKEY_DISABLE_READ 0x8
#undef PKEY_ACCESS_MASK
#define PKEY_ACCESS_MASK (PKEY_DISABLE_ACCESS | PKEY_DISABLE_WRITE | PKEY_DISABLE_READ | PKEY_DISABLE_EXECUTE)
#endif
