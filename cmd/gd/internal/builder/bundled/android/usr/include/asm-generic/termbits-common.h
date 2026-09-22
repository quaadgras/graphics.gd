/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __ASM_GENERIC_TERMBITS_COMMON_H
#define __ASM_GENERIC_TERMBITS_COMMON_H
typedef unsigned char cc_t;
typedef unsigned int speed_t;
#define IGNBRK 0x001
#define BRKINT 0x002
#define IGNPAR 0x004
#define PARMRK 0x008
#define INPCK 0x010
#define ISTRIP 0x020
#define INLCR 0x040
#define IGNCR 0x080
#define ICRNL 0x100
#define IXANY 0x800
#define OPOST 0x01
#define OCRNL 0x08
#define ONOCR 0x10
#define ONLRET 0x20
#define OFILL 0x40
#define OFDEL 0x80
#define B0 0x00000000
#define B50 0x00000001
#define B75 0x00000002
#define B110 0x00000003
#define B134 0x00000004
#define B150 0x00000005
#define B200 0x00000006
#define B300 0x00000007
#define B600 0x00000008
#define B1200 0x00000009
#define B1800 0x0000000a
#define B2400 0x0000000b
#define B4800 0x0000000c
#define B9600 0x0000000d
#define B19200 0x0000000e
#define B38400 0x0000000f
#define EXTA B19200
#define EXTB B38400
#define ADDRB 0x20000000
#define CMSPAR 0x40000000
#define CRTSCTS 0x80000000
#define IBSHIFT 16
#define TCOOFF 0
#define TCOON 1
#define TCIOFF 2
#define TCION 3
#define TCIFLUSH 0
#define TCOFLUSH 1
#define TCIOFLUSH 2
#endif
