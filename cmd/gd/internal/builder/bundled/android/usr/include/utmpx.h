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

/**
 * @file utmpx.h
 * @brief No-op implementation of POSIX login records.
 */

#include <sys/cdefs.h>
#include <sys/types.h>
#include <time.h>

#define EMPTY         0
#define RUN_LVL       1
#define BOOT_TIME     2
#define NEW_TIME      3
#define OLD_TIME      4
#define INIT_PROCESS  5
#define LOGIN_PROCESS 6
#define USER_PROCESS  7
#define DEAD_PROCESS  8
#define ACCOUNTING    9

struct utmpx {
  short ut_type;
  pid_t ut_pid;
  char ut_line[32];
  char ut_id[4];
  char ut_user[32];
  char ut_host[256];

  struct {
    short e_termination;
    short e_exit;
  } ut_exit;

  long ut_session;
  struct timeval ut_tv;

  int32_t ut_addr_v6[4];
  char unused[20];
};

__BEGIN_DECLS

/**
 * Does nothing.
 */
void setutxent(void) __RENAME(setutent);

/**
 * Does nothing and returns null.
 */
struct utmpx* _Nullable getutxent(void) __RENAME(getutent);

/**
 * Does nothing and returns null.
 */
struct utmpx* _Nullable getutxid(const struct utmpx* _Nonnull __entry) __RENAME(getutent);

/**
 * Does nothing and returns null.
 */
struct utmpx* _Nullable getutxline(const struct utmpx* _Nonnull __entry) __RENAME(getutent);

/**
 * Does nothing and returns null.
 */
struct utmpx* _Nullable pututxline(const struct utmpx* _Nonnull __entry) __RENAME(pututline);

/**
 * Does nothing.
 */
void endutxent(void) __RENAME(endutent);

__END_DECLS
