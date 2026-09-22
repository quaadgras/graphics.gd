/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_HIDRAW_H
#define _UAPI_HIDRAW_H
#include <linux/hid.h>
#include <linux/types.h>
struct hidraw_report_descriptor {
  __u32 size;
  __u8 value[HID_MAX_DESCRIPTOR_SIZE];
};
struct hidraw_devinfo {
  __u32 bustype;
  __s16 vendor;
  __s16 product;
};
#define HIDIOCGRDESCSIZE _IOR('H', 0x01, int)
#define HIDIOCGRDESC _IOR('H', 0x02, struct hidraw_report_descriptor)
#define HIDIOCGRAWINFO _IOR('H', 0x03, struct hidraw_devinfo)
#define HIDIOCGRAWNAME(len) _IOC(_IOC_READ, 'H', 0x04, len)
#define HIDIOCGRAWPHYS(len) _IOC(_IOC_READ, 'H', 0x05, len)
#define HIDIOCSFEATURE(len) _IOC(_IOC_WRITE | _IOC_READ, 'H', 0x06, len)
#define HIDIOCGFEATURE(len) _IOC(_IOC_WRITE | _IOC_READ, 'H', 0x07, len)
#define HIDIOCGRAWUNIQ(len) _IOC(_IOC_READ, 'H', 0x08, len)
#define HIDIOCSINPUT(len) _IOC(_IOC_WRITE | _IOC_READ, 'H', 0x09, len)
#define HIDIOCGINPUT(len) _IOC(_IOC_WRITE | _IOC_READ, 'H', 0x0A, len)
#define HIDIOCSOUTPUT(len) _IOC(_IOC_WRITE | _IOC_READ, 'H', 0x0B, len)
#define HIDIOCGOUTPUT(len) _IOC(_IOC_WRITE | _IOC_READ, 'H', 0x0C, len)
#define HIDIOCREVOKE _IOW('H', 0x0D, int)
#define HIDRAW_FIRST_MINOR 0
#define HIDRAW_MAX_DEVICES 64
#define HIDRAW_BUFFER_SIZE 64
#endif
