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
/*
 * Copyright (c) 1982, 1986, 1988, 1993
 *	The Regents of the University of California.  All rights reserved.
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

#pragma once

#include <sys/cdefs.h>

#include <stdio.h>
#include <stdarg.h>

__BEGIN_DECLS

/** Corresponds to the Android ERROR log priority. */
#define LOG_EMERG 0
/** Corresponds to the Android ERROR log priority. */
#define LOG_ALERT 1
/** Corresponds to the Android ERROR log priority. */
#define LOG_CRIT 2
/** Corresponds to the Android ERROR log priority. */
#define LOG_ERR 3
/** Corresponds to the Android WARN log priority. */
#define LOG_WARNING 4
/** Corresponds to the Android INFO log priority. */
#define LOG_NOTICE 5
/** Corresponds to the Android INFO log priority. */
#define LOG_INFO 6
/** Corresponds to the Android DEBUG log priority. */
#define LOG_DEBUG 7

#define LOG_PRIMASK 7
#define LOG_PRI(x) ((x) & LOG_PRIMASK)
#define LOG_MAKEPRI(fac, pri) ((fac) | (pri))

/** Currently ignored on Android. */
#define LOG_KERN     (0<<3)
/** Currently ignored on Android. */
#define LOG_USER     (1<<3)
/** Currently ignored on Android. */
#define LOG_MAIL     (2<<3)
/** Currently ignored on Android. */
#define LOG_DAEMON   (3<<3)
/** Currently ignored on Android. */
#define LOG_AUTH     (4<<3)
/** Currently ignored on Android. */
#define LOG_SYSLOG   (5<<3)
/** Currently ignored on Android. */
#define LOG_LPR      (6<<3)
/** Currently ignored on Android. */
#define LOG_NEWS     (7<<3)
/** Currently ignored on Android. */
#define LOG_UUCP     (8<<3)
/** Currently ignored on Android. */
#define LOG_CRON     (9<<3)
/** Currently ignored on Android. */
#define LOG_AUTHPRIV (10<<3)
/** Currently ignored on Android. */
#define LOG_FTP      (11<<3)
/** Currently ignored on Android. */
#define LOG_LOCAL0   (16<<3)
/** Currently ignored on Android. */
#define LOG_LOCAL1   (17<<3)
/** Currently ignored on Android. */
#define LOG_LOCAL2   (18<<3)
/** Currently ignored on Android. */
#define LOG_LOCAL3   (19<<3)
/** Currently ignored on Android. */
#define LOG_LOCAL4   (20<<3)
/** Currently ignored on Android. */
#define LOG_LOCAL5   (21<<3)
/** Currently ignored on Android. */
#define LOG_LOCAL6   (22<<3)
/** Currently ignored on Android. */
#define LOG_LOCAL7   (23<<3)

#define LOG_NFACILITIES 24
#define LOG_FACMASK 0x3f8
#define LOG_FAC(x) (((x) >> 3) & (LOG_FACMASK >> 3))

/**
 * Converts a log priority into a mask enabling that single priority,
 * for use with setlogmask().
 */
#define LOG_MASK(pri) (1 << (pri))

/**
 * Converts a log priority into a mask enabling that priority and all lower
 * priorities, for use with setlogmask().
 */
#define LOG_UPTO(pri) ((1 << ((pri)+1)) - 1)

/** openlog() option ignored on Android. */
#define LOG_PID    0x01
/** openlog() option ignored on Android. */
#define LOG_CONS   0x02
/** openlog() option ignored on Android. */
#define LOG_ODELAY 0x04
/** openlog() option ignored on Android. */
#define LOG_NDELAY 0x08
/** openlog() option ignored on Android. */
#define LOG_NOWAIT 0x10
/**
 * openlog() option to log to stderr as well as the system log.
 *
 * Available since API level 34 (ignored before then).
 */
#define LOG_PERROR 0x20

#if defined(SYSLOG_NAMES)
/** A mapping from name to value, used by `facilitynames` and `prioritynames`. */
typedef struct _code {
  char* c_name;
  int c_val;
} CODE;
/* A bogus facility value for "mark". */
#define INTERNAL_MARK LOG_MAKEPRI((LOG_NFACILITIES<<3), 0)
/** A table mapping facility names to values. */
static const CODE facilitynames[] = {
  { "auth", LOG_AUTH, },
  { "authpriv", LOG_AUTHPRIV, },
  { "cron", LOG_CRON, },
  { "daemon", LOG_DAEMON, },
  { "ftp", LOG_FTP, },
  { "kern", LOG_KERN, },
  { "lpr", LOG_LPR, },
  { "mail", LOG_MAIL, },
  { "mark", INTERNAL_MARK, },
  { "news", LOG_NEWS, },
  { "security", LOG_AUTH, },
  { "syslog", LOG_SYSLOG, },
  { "user", LOG_USER, },
  { "uucp", LOG_UUCP, },
  { "local0", LOG_LOCAL0, },
  { "local1", LOG_LOCAL1, },
  { "local2", LOG_LOCAL2, },
  { "local3", LOG_LOCAL3, },
  { "local4", LOG_LOCAL4, },
  { "local5", LOG_LOCAL5, },
  { "local6", LOG_LOCAL6, },
  { "local7", LOG_LOCAL7, },
  { NULL, -1, },
};
/* A bogus priority value for "none". */
#define INTERNAL_NOPRI 8
/** A table mapping priority names to values. */
static const CODE prioritynames[] = {
  { "alert", LOG_ALERT, },
  { "crit", LOG_CRIT, },
  { "debug", LOG_DEBUG, },
  { "emerg", LOG_EMERG, },
  { "err", LOG_ERR, },
  { "error", LOG_ERR, },
  { "info", LOG_INFO, },
  { "none", INTERNAL_NOPRI, },
  { "notice", LOG_NOTICE, },
  { "panic", LOG_EMERG, },
  { "warn", LOG_WARNING, },
  { "warning", LOG_WARNING, },
  { NULL, -1, },
};
#endif

/**
 * [closelog(3)](https://man7.org/linux/man-pages/man3/closelog.3.html) does
 * nothing on Android.
 */
void closelog(void);

/**
 * [openlog(3)](https://man7.org/linux/man-pages/man3/openlog.3.html) sets
 * the log tag to `__prefix`, which can be NULL to return to the default of
 * getprogname(). On Android, the other two arguments are ignored.
 */
void openlog(const char* _Nullable __prefix, int __option, int __facility);

/**
 * [setlogmask(3)](https://man7.org/linux/man-pages/man3/setlogmask.3.html)
 * sets which log priorities will actually be logged. See `LOG_MASK` and
 * `LOG_UPTO`.
 */
int setlogmask(int __mask);

/**
 * [syslog(3)](https://man7.org/linux/man-pages/man3/syslog.3.html) formats
 * the printf()-like message and logs it with the given priority, unless
 * suppressed by setlogmask(). On Android, the output goes to logcat.
 */
void syslog(int __priority, const char* _Nonnull __fmt, ...) __printflike(2, 3);

/**
 * [vsyslog(3)](https://man7.org/linux/man-pages/man3/vsyslog.3.html) formats
 * the vprintf()-like message and logs it with the given priority, unless
 * suppressed by setlogmask(). On Android, the output goes to logcat.
 */
void vsyslog(int __priority, const char* _Nonnull __fmt, va_list __args) __printflike(2, 0);

__END_DECLS
