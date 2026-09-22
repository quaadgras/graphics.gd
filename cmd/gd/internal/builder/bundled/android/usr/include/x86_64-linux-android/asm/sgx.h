/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_ASM_X86_SGX_H
#define _UAPI_ASM_X86_SGX_H
#include <linux/types.h>
#include <linux/ioctl.h>
enum sgx_page_flags {
  SGX_PAGE_MEASURE = 0x01,
};
#define SGX_MAGIC 0xA4
#define SGX_IOC_ENCLAVE_CREATE _IOW(SGX_MAGIC, 0x00, struct sgx_enclave_create)
#define SGX_IOC_ENCLAVE_ADD_PAGES _IOWR(SGX_MAGIC, 0x01, struct sgx_enclave_add_pages)
#define SGX_IOC_ENCLAVE_INIT _IOW(SGX_MAGIC, 0x02, struct sgx_enclave_init)
#define SGX_IOC_ENCLAVE_PROVISION _IOW(SGX_MAGIC, 0x03, struct sgx_enclave_provision)
#define SGX_IOC_VEPC_REMOVE_ALL _IO(SGX_MAGIC, 0x04)
#define SGX_IOC_ENCLAVE_RESTRICT_PERMISSIONS _IOWR(SGX_MAGIC, 0x05, struct sgx_enclave_restrict_permissions)
#define SGX_IOC_ENCLAVE_MODIFY_TYPES _IOWR(SGX_MAGIC, 0x06, struct sgx_enclave_modify_types)
#define SGX_IOC_ENCLAVE_REMOVE_PAGES _IOWR(SGX_MAGIC, 0x07, struct sgx_enclave_remove_pages)
struct sgx_enclave_create {
  __u64 src;
};
struct sgx_enclave_add_pages {
  __u64 src;
  __u64 offset;
  __u64 length;
  __u64 secinfo;
  __u64 flags;
  __u64 count;
};
struct sgx_enclave_init {
  __u64 sigstruct;
};
struct sgx_enclave_provision {
  __u64 fd;
};
struct sgx_enclave_restrict_permissions {
  __u64 offset;
  __u64 length;
  __u64 permissions;
  __u64 result;
  __u64 count;
};
struct sgx_enclave_modify_types {
  __u64 offset;
  __u64 length;
  __u64 page_type;
  __u64 result;
  __u64 count;
};
struct sgx_enclave_remove_pages {
  __u64 offset;
  __u64 length;
  __u64 count;
};
struct sgx_enclave_run;
typedef int(* sgx_enclave_user_handler_t) (long rdi, long rsi, long rdx, long rsp, long r8, long r9, struct sgx_enclave_run * run);
struct sgx_enclave_run {
  __u64 tcs;
  __u32 function;
  __u16 exception_vector;
  __u16 exception_error_code;
  __u64 exception_addr;
  __u64 user_handler;
  __u64 user_data;
  __u8 reserved[216];
};
typedef int(* vdso_sgx_enter_enclave_t) (unsigned long rdi, unsigned long rsi, unsigned long rdx, unsigned int function, unsigned long r8, unsigned long r9, struct sgx_enclave_run * run);
#endif
