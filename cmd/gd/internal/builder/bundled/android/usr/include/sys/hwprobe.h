/*
 * Copyright (C) 2023 The Android Open Source Project
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

#if __riscv

/**
 * @file sys/hwprobe.h
 * @brief RISC-V hardware probing.
 */

#include <sys/cdefs.h>
#include <sys/types.h>

/* Pull in struct riscv_hwprobe and corresponding constants. */
#include <asm/hwprobe.h>

__BEGIN_DECLS

/**
 * [__riscv_hwprobe(2)](https://docs.kernel.org/riscv/hwprobe.html)
 * queries hardware characteristics.
 *
 * A `__cpu_count` of 0 and null `__cpus` means "all online cpus".
 *
 * Returns 0 on success and returns an error number on failure.
 */
int __riscv_hwprobe(struct riscv_hwprobe* _Nonnull __pairs, size_t __pair_count, size_t __cpu_count, unsigned long* _Nullable __cpus, unsigned __flags);

/**
 * The type of the second argument passed to riscv64 ifunc resolvers.
 * This argument allows riscv64 ifunc resolvers to call __riscv_hwprobe()
 * without worrying about whether that relocation is resolved before
 * the ifunc resolver is called.
 */
typedef int (*__riscv_hwprobe_t)(struct riscv_hwprobe* _Nonnull __pairs, size_t __pair_count, size_t __cpu_count, unsigned long* _Nullable __cpus, unsigned __flags);

__END_DECLS

#endif
