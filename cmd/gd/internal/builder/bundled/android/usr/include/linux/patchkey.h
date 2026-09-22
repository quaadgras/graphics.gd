/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_PATCHKEY_H_INDIRECT
#error "patchkey.h included directly"
#endif
#ifndef _UAPI_LINUX_PATCHKEY_H
#define _UAPI_LINUX_PATCHKEY_H
#include <endian.h>
#ifdef __BYTE_ORDER
#if __BYTE_ORDER == __BIG_ENDIAN
#define _PATCHKEY(id) (0xfd00 | id)
#elif __BYTE_ORDER==__LITTLE_ENDIAN
#define _PATCHKEY(id) ((id << 8) | 0x00fd)
#else
#error "could not determine byte order"
#endif
#endif
#endif
