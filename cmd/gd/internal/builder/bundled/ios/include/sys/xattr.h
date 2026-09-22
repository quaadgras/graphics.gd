// sys/xattr.h declarations for graphics.gd's iOS SDK, written from the
// xattr(2) manual pages. Not derived from Apple's SDK headers.
#ifndef GD_SYS_XATTR_H
#define GD_SYS_XATTR_H

#include <sys/types.h>

#define XATTR_NOFOLLOW 0x0001
#define XATTR_CREATE 0x0002
#define XATTR_REPLACE 0x0004
#define XATTR_NOSECURITY 0x0008
#define XATTR_NODEFAULT 0x0010
#define XATTR_SHOWCOMPRESSION 0x0020
#define XATTR_MAXNAMELEN 127

#if defined(__cplusplus)
extern "C" {
#endif

ssize_t getxattr(const char *path, const char *name, void *value, size_t size, u_int32_t position, int options);
ssize_t fgetxattr(int fd, const char *name, void *value, size_t size, u_int32_t position, int options);
int setxattr(const char *path, const char *name, const void *value, size_t size, u_int32_t position, int options);
int fsetxattr(int fd, const char *name, const void *value, size_t size, u_int32_t position, int options);
int removexattr(const char *path, const char *name, int options);
int fremovexattr(int fd, const char *name, int options);
ssize_t listxattr(const char *path, char *namebuff, size_t size, int options);
ssize_t flistxattr(int fd, char *namebuff, size_t size, int options);

#if defined(__cplusplus)
}
#endif

#endif
