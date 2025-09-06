/*-
 * Copyright (c) 1989, 1993
 *    The Regents of the University of California.  All rights reserved.
 * (c) UNIX System Laboratories, Inc.
 * All or some portions of this file are derived from material licensed
 * to the University of California by American Telephone and Telegraph
 * Co. or Unix System Laboratories, Inc. and are reproduced herein with
 * the permission of UNIX System Laboratories, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */

#ifndef _GRP_H_
#define _GRP_H_

#include <sys/cdefs.h>
#include <sys/types.h>

struct group {
  char* _Nullable gr_name; /* group name */
  char* _Nullable gr_passwd; /* group password */
  gid_t gr_gid; /* group id */
  char* _Nullable * _Nullable gr_mem; /* group members */
};

__BEGIN_DECLS

struct group* _Nullable getgrgid(gid_t __gid);
struct group* _Nullable getgrnam(const char* _Nonnull __name);

/* Note: Android has thousands and thousands of ids to iterate through. */

#if __BIONIC_AVAILABILITY_GUARD(26)
struct group* _Nullable getgrent(void) __INTRODUCED_IN(26);

void setgrent(void) __INTRODUCED_IN(26);
void endgrent(void) __INTRODUCED_IN(26);
#endif /* __BIONIC_AVAILABILITY_GUARD(26) */


#if __BIONIC_AVAILABILITY_GUARD(24)
int getgrgid_r(gid_t __gid, struct group* __BIONIC_COMPLICATED_NULLNESS __group, char* _Nonnull __buf, size_t __n, struct group* _Nullable * _Nonnull __result) __INTRODUCED_IN(24);
int getgrnam_r(const char* _Nonnull __name, struct group* __BIONIC_COMPLICATED_NULLNESS __group, char* _Nonnull __buf, size_t __n, struct group* _Nullable *_Nonnull __result) __INTRODUCED_IN(24);
#endif /* __BIONIC_AVAILABILITY_GUARD(24) */

int getgrouplist(const char* _Nonnull __user, gid_t __group, gid_t* __BIONIC_COMPLICATED_NULLNESS __groups, int* _Nonnull __group_count);
int initgroups(const char* _Nonnull __user, gid_t __group);

__END_DECLS

#endif
