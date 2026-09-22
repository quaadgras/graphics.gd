/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI__ADB_H
#define _UAPI__ADB_H
#define ADB_BUSRESET 0
#define ADB_FLUSH(id) (0x01 | ((id) << 4))
#define ADB_WRITEREG(id,reg) (0x08 | (reg) | ((id) << 4))
#define ADB_READREG(id,reg) (0x0C | (reg) | ((id) << 4))
#define ADB_DONGLE 1
#define ADB_KEYBOARD 2
#define ADB_MOUSE 3
#define ADB_TABLET 4
#define ADB_MODEM 5
#define ADB_MISC 7
#define ADB_RET_OK 0
#define ADB_RET_TIMEOUT 3
#define ADB_PACKET 0
#define CUDA_PACKET 1
#define ERROR_PACKET 2
#define TIMER_PACKET 3
#define POWER_PACKET 4
#define MACIIC_PACKET 5
#define PMU_PACKET 6
#define ADB_QUERY 7
#define ADB_QUERY_GETDEVINFO 1
#endif
