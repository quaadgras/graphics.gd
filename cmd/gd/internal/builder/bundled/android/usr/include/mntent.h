/*
 * Copyright (C) 2008 The Android Open Source Project
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

#ifndef _MNTENT_H_
#define _MNTENT_H_

#include <sys/cdefs.h>

#include <stdio.h>
#include <paths.h>  /* for _PATH_MOUNTED */

#define MOUNTED _PATH_MOUNTED

#define MNTTYPE_IGNORE "ignore"
#define MNTTYPE_NFS "nfs"
#define MNTTYPE_SWAP "swap"

#define MNTOPT_DEFAULTS "defaults"
#define MNTOPT_NOAUTO "noauto"
#define MNTOPT_NOSUID "nosuid"
#define MNTOPT_RO "ro"
#define MNTOPT_RW "rw"
#define MNTOPT_SUID "suid"

struct mntent {
  char* _Nullable mnt_fsname;
  char* _Nullable mnt_dir;
  char* _Nullable mnt_type;
  char* _Nullable mnt_opts;
  int mnt_freq;
  int mnt_passno;
};

__BEGIN_DECLS

int endmntent(FILE* _Nullable __fp);
struct mntent* _Nullable getmntent(FILE* _Nonnull __fp);
struct mntent* _Nullable getmntent_r(FILE* _Nonnull __fp, struct mntent* _Nonnull __entry, char* _Nonnull __buf, int __size);
FILE* _Nullable setmntent(const char* _Nonnull __filename, const char* _Nonnull __type);

#if __BIONIC_AVAILABILITY_GUARD(26)
char* _Nullable hasmntopt(const struct mntent* _Nonnull __entry, const char* _Nonnull __option) __INTRODUCED_IN(26);
#endif /* __BIONIC_AVAILABILITY_GUARD(26) */


__END_DECLS

#endif
