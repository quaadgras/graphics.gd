/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _ASM_X86_A_OUT_H
#define _ASM_X86_A_OUT_H
struct exec {
  unsigned int a_info;
  unsigned a_text;
  unsigned a_data;
  unsigned a_bss;
  unsigned a_syms;
  unsigned a_entry;
  unsigned a_trsize;
  unsigned a_drsize;
};
#define N_TRSIZE(a) ((a).a_trsize)
#define N_DRSIZE(a) ((a).a_drsize)
#define N_SYMSIZE(a) ((a).a_syms)
#endif
