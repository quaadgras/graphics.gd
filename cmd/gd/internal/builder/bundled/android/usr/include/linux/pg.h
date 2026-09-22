/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_PG_H
#define _UAPI_LINUX_PG_H
#define PG_MAGIC 'P'
#define PG_RESET 'Z'
#define PG_COMMAND 'C'
#define PG_MAX_DATA 32768
struct pg_write_hdr {
  char magic;
  char func;
  int dlen;
  int timeout;
  char packet[12];
};
struct pg_read_hdr {
  char magic;
  char scsi;
  int dlen;
  int duration;
  char pad[12];
};
#endif
