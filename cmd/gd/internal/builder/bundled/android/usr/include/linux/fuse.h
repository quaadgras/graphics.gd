/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_FUSE_H
#define _LINUX_FUSE_H
#include <stdint.h>
#define FUSE_KERNEL_VERSION 7
#define FUSE_KERNEL_MINOR_VERSION 41
#define FUSE_ROOT_ID 1
struct fuse_attr {
  uint64_t ino;
  uint64_t size;
  uint64_t blocks;
  uint64_t atime;
  uint64_t mtime;
  uint64_t ctime;
  uint32_t atimensec;
  uint32_t mtimensec;
  uint32_t ctimensec;
  uint32_t mode;
  uint32_t nlink;
  uint32_t uid;
  uint32_t gid;
  uint32_t rdev;
  uint32_t blksize;
  uint32_t flags;
};
struct fuse_sx_time {
  int64_t tv_sec;
  uint32_t tv_nsec;
  int32_t __reserved;
};
struct fuse_statx {
  uint32_t mask;
  uint32_t blksize;
  uint64_t attributes;
  uint32_t nlink;
  uint32_t uid;
  uint32_t gid;
  uint16_t mode;
  uint16_t __spare0[1];
  uint64_t ino;
  uint64_t size;
  uint64_t blocks;
  uint64_t attributes_mask;
  struct fuse_sx_time atime;
  struct fuse_sx_time btime;
  struct fuse_sx_time ctime;
  struct fuse_sx_time mtime;
  uint32_t rdev_major;
  uint32_t rdev_minor;
  uint32_t dev_major;
  uint32_t dev_minor;
  uint64_t __spare2[14];
};
struct fuse_kstatfs {
  uint64_t blocks;
  uint64_t bfree;
  uint64_t bavail;
  uint64_t files;
  uint64_t ffree;
  uint32_t bsize;
  uint32_t namelen;
  uint32_t frsize;
  uint32_t padding;
  uint32_t spare[6];
};
struct fuse_file_lock {
  uint64_t start;
  uint64_t end;
  uint32_t type;
  uint32_t pid;
};
#define FATTR_MODE (1 << 0)
#define FATTR_UID (1 << 1)
#define FATTR_GID (1 << 2)
#define FATTR_SIZE (1 << 3)
#define FATTR_ATIME (1 << 4)
#define FATTR_MTIME (1 << 5)
#define FATTR_FH (1 << 6)
#define FATTR_ATIME_NOW (1 << 7)
#define FATTR_MTIME_NOW (1 << 8)
#define FATTR_LOCKOWNER (1 << 9)
#define FATTR_CTIME (1 << 10)
#define FATTR_KILL_SUIDGID (1 << 11)
#define FOPEN_DIRECT_IO (1 << 0)
#define FOPEN_KEEP_CACHE (1 << 1)
#define FOPEN_NONSEEKABLE (1 << 2)
#define FOPEN_CACHE_DIR (1 << 3)
#define FOPEN_STREAM (1 << 4)
#define FOPEN_NOFLUSH (1 << 5)
#define FOPEN_PARALLEL_DIRECT_WRITES (1 << 6)
#define FOPEN_PASSTHROUGH (1 << 7)
#define FUSE_ASYNC_READ (1 << 0)
#define FUSE_POSIX_LOCKS (1 << 1)
#define FUSE_FILE_OPS (1 << 2)
#define FUSE_ATOMIC_O_TRUNC (1 << 3)
#define FUSE_EXPORT_SUPPORT (1 << 4)
#define FUSE_BIG_WRITES (1 << 5)
#define FUSE_DONT_MASK (1 << 6)
#define FUSE_SPLICE_WRITE (1 << 7)
#define FUSE_SPLICE_MOVE (1 << 8)
#define FUSE_SPLICE_READ (1 << 9)
#define FUSE_FLOCK_LOCKS (1 << 10)
#define FUSE_HAS_IOCTL_DIR (1 << 11)
#define FUSE_AUTO_INVAL_DATA (1 << 12)
#define FUSE_DO_READDIRPLUS (1 << 13)
#define FUSE_READDIRPLUS_AUTO (1 << 14)
#define FUSE_ASYNC_DIO (1 << 15)
#define FUSE_WRITEBACK_CACHE (1 << 16)
#define FUSE_NO_OPEN_SUPPORT (1 << 17)
#define FUSE_PARALLEL_DIROPS (1 << 18)
#define FUSE_HANDLE_KILLPRIV (1 << 19)
#define FUSE_POSIX_ACL (1 << 20)
#define FUSE_ABORT_ERROR (1 << 21)
#define FUSE_MAX_PAGES (1 << 22)
#define FUSE_CACHE_SYMLINKS (1 << 23)
#define FUSE_NO_OPENDIR_SUPPORT (1 << 24)
#define FUSE_EXPLICIT_INVAL_DATA (1 << 25)
#define FUSE_MAP_ALIGNMENT (1 << 26)
#define FUSE_SUBMOUNTS (1 << 27)
#define FUSE_HANDLE_KILLPRIV_V2 (1 << 28)
#define FUSE_SETXATTR_EXT (1 << 29)
#define FUSE_INIT_EXT (1 << 30)
#define FUSE_INIT_RESERVED (1 << 31)
#define FUSE_SECURITY_CTX (1ULL << 32)
#define FUSE_HAS_INODE_DAX (1ULL << 33)
#define FUSE_CREATE_SUPP_GROUP (1ULL << 34)
#define FUSE_HAS_EXPIRE_ONLY (1ULL << 35)
#define FUSE_DIRECT_IO_ALLOW_MMAP (1ULL << 36)
#define FUSE_PASSTHROUGH (1ULL << 37)
#define FUSE_NO_EXPORT_SUPPORT (1ULL << 38)
#define FUSE_HAS_RESEND (1ULL << 39)
#define FUSE_DIRECT_IO_RELAX FUSE_DIRECT_IO_ALLOW_MMAP
#define FUSE_ALLOW_IDMAP (1ULL << 40)
#define CUSE_UNRESTRICTED_IOCTL (1 << 0)
#define FUSE_RELEASE_FLUSH (1 << 0)
#define FUSE_RELEASE_FLOCK_UNLOCK (1 << 1)
#define FUSE_GETATTR_FH (1 << 0)
#define FUSE_LK_FLOCK (1 << 0)
#define FUSE_WRITE_CACHE (1 << 0)
#define FUSE_WRITE_LOCKOWNER (1 << 1)
#define FUSE_WRITE_KILL_SUIDGID (1 << 2)
#define FUSE_WRITE_KILL_PRIV FUSE_WRITE_KILL_SUIDGID
#define FUSE_READ_LOCKOWNER (1 << 1)
#define FUSE_IOCTL_COMPAT (1 << 0)
#define FUSE_IOCTL_UNRESTRICTED (1 << 1)
#define FUSE_IOCTL_RETRY (1 << 2)
#define FUSE_IOCTL_32BIT (1 << 3)
#define FUSE_IOCTL_DIR (1 << 4)
#define FUSE_IOCTL_COMPAT_X32 (1 << 5)
#define FUSE_IOCTL_MAX_IOV 256
#define FUSE_POLL_SCHEDULE_NOTIFY (1 << 0)
#define FUSE_FSYNC_FDATASYNC (1 << 0)
#define FUSE_ATTR_SUBMOUNT (1 << 0)
#define FUSE_ATTR_DAX (1 << 1)
#define FUSE_OPEN_KILL_SUIDGID (1 << 0)
#define FUSE_SETXATTR_ACL_KILL_SGID (1 << 0)
#define FUSE_EXPIRE_ONLY (1 << 0)
enum fuse_ext_type {
  FUSE_MAX_NR_SECCTX = 31,
  FUSE_EXT_GROUPS = 32,
};
enum fuse_opcode {
  FUSE_LOOKUP = 1,
  FUSE_FORGET = 2,
  FUSE_GETATTR = 3,
  FUSE_SETATTR = 4,
  FUSE_READLINK = 5,
  FUSE_SYMLINK = 6,
  FUSE_MKNOD = 8,
  FUSE_MKDIR = 9,
  FUSE_UNLINK = 10,
  FUSE_RMDIR = 11,
  FUSE_RENAME = 12,
  FUSE_LINK = 13,
  FUSE_OPEN = 14,
  FUSE_READ = 15,
  FUSE_WRITE = 16,
  FUSE_STATFS = 17,
  FUSE_RELEASE = 18,
  FUSE_FSYNC = 20,
  FUSE_SETXATTR = 21,
  FUSE_GETXATTR = 22,
  FUSE_LISTXATTR = 23,
  FUSE_REMOVEXATTR = 24,
  FUSE_FLUSH = 25,
  FUSE_INIT = 26,
  FUSE_OPENDIR = 27,
  FUSE_READDIR = 28,
  FUSE_RELEASEDIR = 29,
  FUSE_FSYNCDIR = 30,
  FUSE_GETLK = 31,
  FUSE_SETLK = 32,
  FUSE_SETLKW = 33,
  FUSE_ACCESS = 34,
  FUSE_CREATE = 35,
  FUSE_INTERRUPT = 36,
  FUSE_BMAP = 37,
  FUSE_DESTROY = 38,
  FUSE_IOCTL = 39,
  FUSE_POLL = 40,
  FUSE_NOTIFY_REPLY = 41,
  FUSE_BATCH_FORGET = 42,
  FUSE_FALLOCATE = 43,
  FUSE_READDIRPLUS = 44,
  FUSE_RENAME2 = 45,
  FUSE_LSEEK = 46,
  FUSE_COPY_FILE_RANGE = 47,
  FUSE_SETUPMAPPING = 48,
  FUSE_REMOVEMAPPING = 49,
  FUSE_SYNCFS = 50,
  FUSE_TMPFILE = 51,
  FUSE_STATX = 52,
  FUSE_CANONICAL_PATH = 2016,
  CUSE_INIT = 4096,
  CUSE_INIT_BSWAP_RESERVED = 1048576,
  FUSE_INIT_BSWAP_RESERVED = 436207616,
};
enum fuse_notify_code {
  FUSE_NOTIFY_POLL = 1,
  FUSE_NOTIFY_INVAL_INODE = 2,
  FUSE_NOTIFY_INVAL_ENTRY = 3,
  FUSE_NOTIFY_STORE = 4,
  FUSE_NOTIFY_RETRIEVE = 5,
  FUSE_NOTIFY_DELETE = 6,
  FUSE_NOTIFY_RESEND = 7,
  FUSE_NOTIFY_CODE_MAX,
};
#define FUSE_MIN_READ_BUFFER 8192
#define FUSE_COMPAT_ENTRY_OUT_SIZE 120
struct fuse_entry_out {
  uint64_t nodeid;
  uint64_t generation;
  uint64_t entry_valid;
  uint64_t attr_valid;
  uint32_t entry_valid_nsec;
  uint32_t attr_valid_nsec;
  struct fuse_attr attr;
};
struct fuse_forget_in {
  uint64_t nlookup;
};
struct fuse_forget_one {
  uint64_t nodeid;
  uint64_t nlookup;
};
struct fuse_batch_forget_in {
  uint32_t count;
  uint32_t dummy;
};
struct fuse_getattr_in {
  uint32_t getattr_flags;
  uint32_t dummy;
  uint64_t fh;
};
#define FUSE_COMPAT_ATTR_OUT_SIZE 96
struct fuse_attr_out {
  uint64_t attr_valid;
  uint32_t attr_valid_nsec;
  uint32_t dummy;
  struct fuse_attr attr;
};
struct fuse_statx_in {
  uint32_t getattr_flags;
  uint32_t reserved;
  uint64_t fh;
  uint32_t sx_flags;
  uint32_t sx_mask;
};
struct fuse_statx_out {
  uint64_t attr_valid;
  uint32_t attr_valid_nsec;
  uint32_t flags;
  uint64_t spare[2];
  struct fuse_statx stat;
};
#define FUSE_COMPAT_MKNOD_IN_SIZE 8
struct fuse_mknod_in {
  uint32_t mode;
  uint32_t rdev;
  uint32_t umask;
  uint32_t padding;
};
struct fuse_mkdir_in {
  uint32_t mode;
  uint32_t umask;
};
struct fuse_rename_in {
  uint64_t newdir;
};
struct fuse_rename2_in {
  uint64_t newdir;
  uint32_t flags;
  uint32_t padding;
};
struct fuse_link_in {
  uint64_t oldnodeid;
};
struct fuse_setattr_in {
  uint32_t valid;
  uint32_t padding;
  uint64_t fh;
  uint64_t size;
  uint64_t lock_owner;
  uint64_t atime;
  uint64_t mtime;
  uint64_t ctime;
  uint32_t atimensec;
  uint32_t mtimensec;
  uint32_t ctimensec;
  uint32_t mode;
  uint32_t unused4;
  uint32_t uid;
  uint32_t gid;
  uint32_t unused5;
};
struct fuse_open_in {
  uint32_t flags;
  uint32_t open_flags;
};
struct fuse_create_in {
  uint32_t flags;
  uint32_t mode;
  uint32_t umask;
  uint32_t open_flags;
};
struct fuse_open_out {
  uint64_t fh;
  uint32_t open_flags;
  int32_t backing_id;
};
struct fuse_release_in {
  uint64_t fh;
  uint32_t flags;
  uint32_t release_flags;
  uint64_t lock_owner;
};
struct fuse_flush_in {
  uint64_t fh;
  uint32_t unused;
  uint32_t padding;
  uint64_t lock_owner;
};
struct fuse_read_in {
  uint64_t fh;
  uint64_t offset;
  uint32_t size;
  uint32_t read_flags;
  uint64_t lock_owner;
  uint32_t flags;
  uint32_t padding;
};
#define FUSE_COMPAT_WRITE_IN_SIZE 24
struct fuse_write_in {
  uint64_t fh;
  uint64_t offset;
  uint32_t size;
  uint32_t write_flags;
  uint64_t lock_owner;
  uint32_t flags;
  uint32_t padding;
};
struct fuse_write_out {
  uint32_t size;
  uint32_t padding;
};
#define FUSE_COMPAT_STATFS_SIZE 48
struct fuse_statfs_out {
  struct fuse_kstatfs st;
};
struct fuse_fsync_in {
  uint64_t fh;
  uint32_t fsync_flags;
  uint32_t padding;
};
#define FUSE_COMPAT_SETXATTR_IN_SIZE 8
struct fuse_setxattr_in {
  uint32_t size;
  uint32_t flags;
  uint32_t setxattr_flags;
  uint32_t padding;
};
struct fuse_getxattr_in {
  uint32_t size;
  uint32_t padding;
};
struct fuse_getxattr_out {
  uint32_t size;
  uint32_t padding;
};
struct fuse_lk_in {
  uint64_t fh;
  uint64_t owner;
  struct fuse_file_lock lk;
  uint32_t lk_flags;
  uint32_t padding;
};
struct fuse_lk_out {
  struct fuse_file_lock lk;
};
struct fuse_access_in {
  uint32_t mask;
  uint32_t padding;
};
struct fuse_init_in {
  uint32_t major;
  uint32_t minor;
  uint32_t max_readahead;
  uint32_t flags;
  uint32_t flags2;
  uint32_t unused[11];
};
#define FUSE_COMPAT_INIT_OUT_SIZE 8
#define FUSE_COMPAT_22_INIT_OUT_SIZE 24
struct fuse_init_out {
  uint32_t major;
  uint32_t minor;
  uint32_t max_readahead;
  uint32_t flags;
  uint16_t max_background;
  uint16_t congestion_threshold;
  uint32_t max_write;
  uint32_t time_gran;
  uint16_t max_pages;
  uint16_t map_alignment;
  uint32_t flags2;
  uint32_t max_stack_depth;
  uint32_t unused[6];
};
#define CUSE_INIT_INFO_MAX 4096
struct cuse_init_in {
  uint32_t major;
  uint32_t minor;
  uint32_t unused;
  uint32_t flags;
};
struct cuse_init_out {
  uint32_t major;
  uint32_t minor;
  uint32_t unused;
  uint32_t flags;
  uint32_t max_read;
  uint32_t max_write;
  uint32_t dev_major;
  uint32_t dev_minor;
  uint32_t spare[10];
};
struct fuse_interrupt_in {
  uint64_t unique;
};
struct fuse_bmap_in {
  uint64_t block;
  uint32_t blocksize;
  uint32_t padding;
};
struct fuse_bmap_out {
  uint64_t block;
};
struct fuse_ioctl_in {
  uint64_t fh;
  uint32_t flags;
  uint32_t cmd;
  uint64_t arg;
  uint32_t in_size;
  uint32_t out_size;
};
struct fuse_ioctl_iovec {
  uint64_t base;
  uint64_t len;
};
struct fuse_ioctl_out {
  int32_t result;
  uint32_t flags;
  uint32_t in_iovs;
  uint32_t out_iovs;
};
struct fuse_poll_in {
  uint64_t fh;
  uint64_t kh;
  uint32_t flags;
  uint32_t events;
};
struct fuse_poll_out {
  uint32_t revents;
  uint32_t padding;
};
struct fuse_notify_poll_wakeup_out {
  uint64_t kh;
};
struct fuse_fallocate_in {
  uint64_t fh;
  uint64_t offset;
  uint64_t length;
  uint32_t mode;
  uint32_t padding;
};
#define FUSE_UNIQUE_RESEND (1ULL << 63)
#define FUSE_INVALID_UIDGID ((uint32_t) (- 1))
struct fuse_in_header {
  uint32_t len;
  uint32_t opcode;
  uint64_t unique;
  uint64_t nodeid;
  uint32_t uid;
  uint32_t gid;
  uint32_t pid;
  uint16_t total_extlen;
  uint16_t padding;
};
struct fuse_out_header {
  uint32_t len;
  int32_t error;
  uint64_t unique;
};
struct fuse_dirent {
  uint64_t ino;
  uint64_t off;
  uint32_t namelen;
  uint32_t type;
  char name[];
};
#define FUSE_REC_ALIGN(x) (((x) + sizeof(uint64_t) - 1) & ~(sizeof(uint64_t) - 1))
#define FUSE_NAME_OFFSET offsetof(struct fuse_dirent, name)
#define FUSE_DIRENT_ALIGN(x) FUSE_REC_ALIGN(x)
#define FUSE_DIRENT_SIZE(d) FUSE_DIRENT_ALIGN(FUSE_NAME_OFFSET + (d)->namelen)
struct fuse_direntplus {
  struct fuse_entry_out entry_out;
  struct fuse_dirent dirent;
};
#define FUSE_NAME_OFFSET_DIRENTPLUS offsetof(struct fuse_direntplus, dirent.name)
#define FUSE_DIRENTPLUS_SIZE(d) FUSE_DIRENT_ALIGN(FUSE_NAME_OFFSET_DIRENTPLUS + (d)->dirent.namelen)
struct fuse_notify_inval_inode_out {
  uint64_t ino;
  int64_t off;
  int64_t len;
};
struct fuse_notify_inval_entry_out {
  uint64_t parent;
  uint32_t namelen;
  uint32_t flags;
};
struct fuse_notify_delete_out {
  uint64_t parent;
  uint64_t child;
  uint32_t namelen;
  uint32_t padding;
};
struct fuse_notify_store_out {
  uint64_t nodeid;
  uint64_t offset;
  uint32_t size;
  uint32_t padding;
};
struct fuse_notify_retrieve_out {
  uint64_t notify_unique;
  uint64_t nodeid;
  uint64_t offset;
  uint32_t size;
  uint32_t padding;
};
struct fuse_notify_retrieve_in {
  uint64_t dummy1;
  uint64_t offset;
  uint32_t size;
  uint32_t dummy2;
  uint64_t dummy3;
  uint64_t dummy4;
};
struct fuse_backing_map {
  int32_t fd;
  uint32_t flags;
  uint64_t padding;
};
#define FUSE_DEV_IOC_MAGIC 229
#define FUSE_DEV_IOC_CLONE _IOR(FUSE_DEV_IOC_MAGIC, 0, uint32_t)
#define FUSE_DEV_IOC_BACKING_OPEN _IOW(FUSE_DEV_IOC_MAGIC, 1, struct fuse_backing_map)
#define FUSE_DEV_IOC_BACKING_CLOSE _IOW(FUSE_DEV_IOC_MAGIC, 2, uint32_t)
struct fuse_lseek_in {
  uint64_t fh;
  uint64_t offset;
  uint32_t whence;
  uint32_t padding;
};
struct fuse_lseek_out {
  uint64_t offset;
};
struct fuse_copy_file_range_in {
  uint64_t fh_in;
  uint64_t off_in;
  uint64_t nodeid_out;
  uint64_t fh_out;
  uint64_t off_out;
  uint64_t len;
  uint64_t flags;
};
#define FUSE_SETUPMAPPING_FLAG_WRITE (1ull << 0)
#define FUSE_SETUPMAPPING_FLAG_READ (1ull << 1)
struct fuse_setupmapping_in {
  uint64_t fh;
  uint64_t foffset;
  uint64_t len;
  uint64_t flags;
  uint64_t moffset;
};
struct fuse_removemapping_in {
  uint32_t count;
};
struct fuse_removemapping_one {
  uint64_t moffset;
  uint64_t len;
};
#define FUSE_REMOVEMAPPING_MAX_ENTRY (PAGE_SIZE / sizeof(struct fuse_removemapping_one))
struct fuse_syncfs_in {
  uint64_t padding;
};
struct fuse_secctx {
  uint32_t size;
  uint32_t padding;
};
struct fuse_secctx_header {
  uint32_t size;
  uint32_t nr_secctx;
};
struct fuse_ext_header {
  uint32_t size;
  uint32_t type;
};
struct fuse_supp_groups {
  uint32_t nr_groups;
  uint32_t groups[];
};
#endif
