/*
 * Copyright (C) 2020 The Android Open Source Project
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
 * @file thread_properties.h
 * @brief Thread properties API.
 *
 * https://sourceware.org/glibc/wiki/ThreadPropertiesAPI
 * API for querying various properties of the current thread, used mostly by
 * the sanitizers.
 *
 * Available since API level 31.
 *
 */

#include <sys/cdefs.h>
#include <unistd.h>

__BEGIN_DECLS

/**
 * Gets the bounds of static TLS for the current thread.
 *
 * Available since API level 31.
 */

#if __BIONIC_AVAILABILITY_GUARD(31)
void __libc_get_static_tls_bounds(void* _Nonnull * _Nonnull __static_tls_begin,
                                  void* _Nonnull * _Nonnull __static_tls_end) __INTRODUCED_IN(31);


/**
 * Registers callback to be called right before the thread is dead.
 * The callbacks are chained, they are called in the order opposite to the order
 * they were registered.
 *
 * The callbacks must be registered only before any threads were created.
 * No signals may arrive during the calls to these callbacks.
 * The callbacks may not access the thread's dynamic TLS because they will have
 * been freed by the time these callbacks are invoked.
 *
 * Available since API level 31.
 */
void __libc_register_thread_exit_callback(void (* _Nonnull __cb)(void)) __INTRODUCED_IN(31);

/**
 * Iterates over all dynamic TLS chunks for the given thread.
 * The thread should have been suspended. It is undefined-behaviour if there is concurrent
 * modification of the target thread's dynamic TLS.
 *
 * Available since API level 31.
 */
void __libc_iterate_dynamic_tls(pid_t __tid,
                                void (* _Nonnull __cb)(void* _Nonnull __dynamic_tls_begin,
                                             void* _Nonnull __dynamic_tls_end,
                                             size_t __dso_id,
                                             void* _Nullable __arg),
                                void* _Nullable __arg) __INTRODUCED_IN(31);

/**
 * Register on_creation and on_destruction callbacks, which will be called after a dynamic
 * TLS creation and before a dynamic TLS destruction, respectively.
 *
 * Available since API level 31.
 */
void __libc_register_dynamic_tls_listeners(
    void (* _Nonnull __on_creation)(void* _Nonnull __dynamic_tls_begin,
                          void* _Nonnull __dynamic_tls_end),
    void (* _Nonnull __on_destruction)(void* _Nonnull __dynamic_tls_begin,
                             void* _Nonnull __dynamic_tls_end)) __INTRODUCED_IN(31);
#endif /* __BIONIC_AVAILABILITY_GUARD(31) */


__END_DECLS
