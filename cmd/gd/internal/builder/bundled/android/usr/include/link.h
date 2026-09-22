/*
 * Copyright (C) 2012 The Android Open Source Project
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *  * Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in
 *    the documentation and/or other materials provided with the
 *    distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
 * OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED
 * AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT
 * OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */

#pragma once

/**
 * @file link.h
 * @brief Extra dynamic linker functionality (see also <dlfcn.h>).
 */

#include <sys/cdefs.h>

#include <stdint.h>
#include <sys/types.h>

#include <elf.h>

__BEGIN_DECLS

#if defined(__LP64__)
/** Convenience macro to get the appropriate 32-bit or 64-bit <elf.h> type for the caller's bitness. */
#define ElfW(type) Elf64_ ## type
#else
/** Convenience macro to get the appropriate 32-bit or 64-bit <elf.h> type for the caller's bitness. */
#define ElfW(type) Elf32_ ## type
#endif

/**
 * Information passed by dl_iterate_phdr() to the callback.
 */
struct dl_phdr_info {
  /** The address of the shared object. */
  ElfW(Addr) dlpi_addr;
  /** The name of the shared object. */
  const char* _Nullable dlpi_name;
  /** Pointer to the shared object's program headers. */
  const ElfW(Phdr)* _Nullable dlpi_phdr;
  /** Number of program headers pointed to by `dlpi_phdr`. */
  ElfW(Half) dlpi_phnum;

  /**
   * The total number of library load events at the time dl_iterate_phdr() was
   * called.
   *
   * This field is only available since API level 30; you can use the size
   * passed to the callback to determine whether you have the full struct,
   * or just the fields up to and including `dlpi_phnum`.
   */
  unsigned long long dlpi_adds;
  /**
   * The total number of library unload events at the time dl_iterate_phdr() was
   * called.
   *
   * This field is only available since API level 30; you can use the size
   * passed to the callback to determine whether you have the full struct,
   * or just the fields up to and including `dlpi_phnum`.
   */
  unsigned long long dlpi_subs;
  /**
   * The module ID for TLS relocations in this shared object.
   *
   * This field is only available since API level 30; you can use the size
   * passed to the callback to determine whether you have the full struct,
   * or just the fields up to and including `dlpi_phnum`.
   */
  size_t dlpi_tls_modid;
  /**
   * The caller's TLS data for this shared object.
   *
   * This field is only available since API level 30; you can use the size
   * passed to the callback to determine whether you have the full struct,
   * or just the fields up to and including `dlpi_phnum`.
   */
  void* _Nullable dlpi_tls_data;
};

/**
 * [dl_iterate_phdr(3)](https://man7.org/linux/man-pages/man3/dl_iterate_phdr.3.html)
 * calls the given callback once for every loaded shared object. The size
 * argument to the callback lets you determine whether you have a smaller
 * `dl_phdr_info` from before API level 30, or the newer full one.
 * The data argument to the callback is whatever you pass as the data argument
 * to dl_iterate_phdr().
 *
 * Returns the value returned by the final call to the callback.
 */
int dl_iterate_phdr(int (* _Nonnull __callback)(struct dl_phdr_info* _Nonnull __info, size_t __size, void* _Nullable __data), void* _Nullable __data);

#ifdef __arm__
typedef uintptr_t _Unwind_Ptr;
_Unwind_Ptr dl_unwind_find_exidx(_Unwind_Ptr, int* _Nonnull);
#endif

/** Used by the dynamic linker to communicate with the debugger. */
struct link_map {
  ElfW(Addr) l_addr;
  char* _Nullable l_name;
  ElfW(Dyn)* _Nullable l_ld;
  struct link_map* _Nullable l_next;
  struct link_map* _Nullable l_prev;
};

/** Used by the dynamic linker to communicate with the debugger. */
struct r_debug {
  int32_t r_version;
  struct link_map* _Nullable r_map;
  ElfW(Addr) r_brk;
  enum {
    RT_CONSISTENT,
    RT_ADD,
    RT_DELETE
  } r_state;
  ElfW(Addr) r_ldbase;
};

__END_DECLS
