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

extern void __cxa_finalize(void *);
extern void *__dso_handle;

__attribute__((destructor))
static void __on_dlclose(void) {
  __cxa_finalize(&__dso_handle);
}

/* Define a weak stub function here that will be overridden if the solib uses
 * emutls. The function needs to be a definition, not just a declaration,
 * because gold has a bug where it outputs weak+hidden symbols into the .dynsym
 * table. */
__attribute__((weak,visibility("hidden")))
void __emutls_unregister_key(void) {
}

/* Use a priority of 0 to run after any ordinary destructor function. The
 * priority setting moves the function towards the front of the .fini_array
 * section. */
__attribute__((destructor(0)))
static void __on_dlclose_late(void) {
  __emutls_unregister_key();
}

/* CRT_LEGACY_WORKAROUND should only be defined when building
 * this file as part of the platform's C library.
 *
 * The C library already defines a function named 'atexit()'
 * for backwards compatibility with older NDK-generated binaries.
 *
 * For newer ones, 'atexit' is actually embedded in the C
 * runtime objects that are linked into the final ELF
 * binary (shared library or executable), and will call
 * __cxa_atexit() in order to un-register any atexit()
 * handler when a library is unloaded.
 *
 * This function must be global *and* hidden. Only the
 * code inside the same ELF binary should be able to access it.
 */

#ifdef CRT_LEGACY_WORKAROUND
# include "__dso_handle.h"
#else
# include "__dso_handle_so.h"
# include "atexit.h"
#endif
#include "pthread_atfork.h"
