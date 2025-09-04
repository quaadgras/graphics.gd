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

#pragma once

/**
 * @file time.h
 * @brief Clock and timer functionality.
 */

#include <sys/cdefs.h>
#include <sys/time.h>
#include <xlocale.h>

__BEGIN_DECLS

/* If we just use void* in the typedef, the compiler exposes that in error messages. */
struct __timezone_t;

/**
 * The `timezone_t` type that represents a timezone.
 *
 * To use this with std::unique_ptr you'll want something like
 * `std::unique_ptr<std::remove_pointer_t<timezone_t>, decltype(&tzfree)> tz{tzalloc("Asia/Seoul"), tzfree};`
 * to remove the pointer.
 */
typedef struct __timezone_t* timezone_t;

/** Divisor to compute seconds from the result of a call to clock(). */
#define CLOCKS_PER_SEC 1000000

/**
 * The name of the current timezone's non-daylight savings (`tzname[0]`) and
 * daylight savings (`tzname[1]`) variants. See tzset().
 */
extern char* _Nonnull tzname[];

/** Whether the current timezone ever uses daylight savings time. See tzset(). */
extern int daylight;

/** The difference in seconds between UTC and the current timezone. See tzset(). */
extern long int timezone;

struct sigevent;

/**
 * A "broken-down" time, useful for parsing/formatting times for human consumption.
 */
struct tm {
  /** Seconds, 0-60. (60 is a leap second.) */
  int tm_sec;
  /** Minutes, 0-59. */
  int tm_min;
  /** Hours, 0-23. */
  int tm_hour;
  /** Day of month, 1-31. */
  int tm_mday;
  /** Month of year, 0-11. (Not 1-12!) */
  int tm_mon;
  /** Years since 1900. (So 2023 is 123, not 2023!) */
  int tm_year;
  /** Day of week, 0-6. (Sunday is 0, Saturday is 6.) */
  int tm_wday;
  /** Day of year, 0-365. */
  int tm_yday;
  /** Daylight savings flag, positive for DST in effect, 0 for DST not in effect, and -1 for unknown. */
  int tm_isdst;
  /** Offset from UTC (GMT) in seconds for this time. */
  long int tm_gmtoff;
  /** Name of the timezone for this time. */
  const char* _Nullable tm_zone;
};

/** Alternative name for `tm_zone` in `struct tm`. */
#define TM_ZONE tm_zone

/**
 * [time(2)](https://man7.org/linux/man-pages/man2/time.2.html) returns
 * the number of seconds since the Unix epoch (1970-01-01 00:00:00 +0000).
 *
 * Returns the time in seconds on success, and returns -1 and sets `errno` on failure.
 */
time_t time(time_t* _Nullable __t);

/**
 * [nanosleep(2)](https://man7.org/linux/man-pages/man2/nanosleep.2.html) sleeps
 * for at least the given time (or until a signal arrives).
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure. If the sleep
 * was interrupted by a signal, `errno` will be `EINTR` and `remainder` will be
 * the amount of time remaining.
 */
int nanosleep(const struct timespec* _Nonnull __duration, struct timespec* _Nullable __remainder);

/**
 * [asctime(3)](https://man7.org/linux/man-pages/man3/asctime.3p.html) formats
 * the time `tm` as a string.
 *
 * Returns a pointer to a string on success, and returns NULL on failure.
 *
 * That string will be overwritten by later calls to this function.
 *
 * New code should prefer strftime().
 */
char* _Nullable asctime(const struct tm* _Nonnull __tm);

/**
 * [asctime_r(3)](https://man7.org/linux/man-pages/man3/asctime_r.3p.html) formats
 * the time `tm` as a string in the given buffer `buf`.
 *
 * Returns a pointer to a string on success, and returns NULL on failure.
 *
 * New code should prefer strftime().
 */
char* _Nullable asctime_r(const struct tm* _Nonnull __tm, char* _Nonnull __buf);

/**
 * [difftime(3)](https://man7.org/linux/man-pages/man3/difftime.3.html) returns
 * the difference between two times.
 *
 * Returns the difference in seconds.
 */
double difftime(time_t __lhs, time_t __rhs);

/**
 * [mktime(3)](https://man7.org/linux/man-pages/man3/mktime.3p.html) converts
 * broken-down time `tm` into the number of seconds since the Unix epoch.
 *
 * See tzset() for details of how the timezone is set, and mktime_rz()
 * for an alternative.
 *
 * Returns the time in seconds on success, and returns -1 and sets `errno` on failure.
 */
time_t mktime(struct tm* _Nonnull __tm);

/**
 * mktime_z(3) converts broken-down time `tm` into the number of seconds
 * since the Unix epoch, assuming the given timezone.
 *
 * Returns the time in seconds on success, and returns -1 and sets `errno` on failure.
 *
 * Available since API level 35.
 */

#if __BIONIC_AVAILABILITY_GUARD(35)
time_t mktime_z(timezone_t _Nonnull __tz, struct tm* _Nonnull __tm) __INTRODUCED_IN(35);
#endif /* __BIONIC_AVAILABILITY_GUARD(35) */


/**
 * [localtime(3)](https://man7.org/linux/man-pages/man3/localtime.3p.html) converts
 * the number of seconds since the Unix epoch in `t` to a broken-down time, taking
 * the device's timezone into account.
 *
 * That broken-down time will be overwritten by later calls to this function.
 *
 * Returns a pointer to a broken-down time on success, and returns null and sets `errno` on failure.
 */
struct tm* _Nullable localtime(const time_t* _Nonnull __t);

/**
 * [localtime_r(3)](https://man7.org/linux/man-pages/man3/localtime_r.3p.html) converts
 * the number of seconds since the Unix epoch in `t` to a broken-down time.
 * That broken-down time will be written to the given struct `tm`.
 *
 * See tzset() for details of how the timezone is set, and localtime_rz()
 * for an alternative.
 *
 * Returns a pointer to a broken-down time on success, and returns null and sets `errno` on failure.
 */
struct tm* _Nullable localtime_r(const time_t* _Nonnull __t, struct tm* _Nonnull __tm);

/**
 * localtime_rz(3) converts the number of seconds since the Unix epoch in
 * `t` to a broken-down time, assuming the given timezone. That broken-down
 * time will be written to the given struct `tm`.
 *
 * Returns a pointer to a broken-down time on success, and returns null and sets `errno` on failure.
 *
 * Available since API level 35.
 */

#if __BIONIC_AVAILABILITY_GUARD(35)
struct tm* _Nullable localtime_rz(timezone_t _Nonnull __tz, const time_t* _Nonnull __t, struct tm* _Nonnull __tm) __INTRODUCED_IN(35);
#endif /* __BIONIC_AVAILABILITY_GUARD(35) */


/**
 * Inverse of localtime().
 */
time_t timelocal(struct tm* _Nonnull __tm);

/**
 * [gmtime(3)](https://man7.org/linux/man-pages/man3/gmtime.3p.html) converts
 * the number of seconds since the Unix epoch in `t` to a broken-down time, using
 * UTC (historically also known as GMT).
 *
 * That broken-down time will be overwritten by later calls to this function.
 *
 * Returns a pointer to a broken-down time on success, and returns null and sets `errno` on failure.
 */
struct tm* _Nullable gmtime(const time_t* _Nonnull __t);

/**
 * [gmtime_r(3)](https://man7.org/linux/man-pages/man3/gmtime_r.3p.html) converts
 * the number of seconds since the Unix epoch in `t` to a broken-down time, using
 * UTC (historically also known as GMT).
 *
 * That broken-down time will be written to the provided struct `tm`.
 *
 * Returns a pointer to a broken-down time on success, and returns null and sets `errno` on failure.
 */
struct tm* _Nullable gmtime_r(const time_t* _Nonnull __t, struct tm* _Nonnull __tm);

/**
 * Inverse of gmtime().
 */
time_t timegm(struct tm* _Nonnull __tm);

/**
 * [strptime(3)](https://man7.org/linux/man-pages/man3/strptime.3.html) parses
 * a string `s` assuming format `fmt` into broken-down time `tm`.
 *
 * Returns a pointer to the first character _not_ parsed, or null if no characters were parsed.
 */
char* _Nullable strptime(const char* _Nonnull __s, const char* _Nonnull __fmt, struct tm* _Nonnull __tm) __strftimelike(2);

/**
 * Equivalent to strptime() on Android where only C/POSIX locales are available.
 */
char* _Nullable strptime_l(const char* _Nonnull __s, const char* _Nonnull __fmt, struct tm* _Nonnull __tm, locale_t _Nonnull __l) __strftimelike(2) __RENAME(strptime);

/**
 * [strftime(3)](https://man7.org/linux/man-pages/man3/strftime.3.html) formats
 * a broken-down time `tm` into the buffer `buf` using format `fmt`.
 *
 * Returns a pointer to the first character _not_ parsed, or null if no characters were parsed.
 */
size_t strftime(char* _Nonnull __buf, size_t __n, const char* _Nonnull __fmt, const struct tm* _Nullable __tm) __strftimelike(3);

/**
 * Equivalent to strftime() on Android where only C/POSIX locales are available.
 */
size_t strftime_l(char* _Nonnull __buf, size_t __n, const char* _Nonnull __fmt, const struct tm* _Nullable __tm, locale_t _Nonnull __l) __strftimelike(3);

/**
 * [ctime(3)](https://man7.org/linux/man-pages/man3/ctime.3p.html) formats
 * the time `tm` as a string.
 *
 * Returns a pointer to a string on success, and returns NULL on failure.
 *
 * That string will be overwritten by later calls to this function.
 *
 * New code should prefer strftime().
 */
char* _Nullable ctime(const time_t* _Nonnull __t);

/**
 * [ctime_r(3)](https://man7.org/linux/man-pages/man3/ctime_r.3p.html) formats
 * the time `tm` as a string in the given buffer `buf`.
 *
 * Returns a pointer to a string on success, and returns NULL on failure.
 *
 * New code should prefer strftime().
 */
char* _Nullable ctime_r(const time_t* _Nonnull __t, char* _Nonnull __buf);

/**
 * [tzset(3)](https://man7.org/linux/man-pages/man3/tzset.3.html) tells
 * libc that the timezone has changed.
 *
 * tzset() on Android looks at both the system property
 * `persist.sys.timezone` and the environment variable `TZ`. The former is
 * the device's current timezone as shown in Settings, while the latter is
 * usually unset but can be used to override the global setting. This is a
 * bad idea outside of unit tests or single-threaded programs because it's
 * inherently thread-unsafe. See tzalloc(), localtime_rz(), mktime_z(),
 * and tzfree() for an alternative.
 */
void tzset(void);

/**
 * tzalloc(3) allocates a timezone corresponding to the given Olson ID.
 *
 * A null `id` returns the system timezone (as seen in Settings) from
 * the system property `persist.sys.timezone`, ignoring `$TZ`. Although
 * tzset() honors `$TZ`, callers of tzalloc() can use `$TZ` themselves if
 * that's the (thread-unsafe) behavior they want, but by ignoring `$TZ`
 * tzalloc() is thread safe (though obviously the system timezone can
 * change, especially if your mobile device is actually mobile!).
 *
 * To use this with std::unique_ptr you'll want something like
 * `std::unique_ptr<std::remove_pointer_t<timezone_t>, decltype(&tzfree)> tz{tzalloc("Asia/Seoul"), tzfree};`
 * to remove the pointer.
 *
 * Returns a timezone object on success, and returns NULL and sets `errno` on failure.
 *
 * Available since API level 35.
 */

#if __BIONIC_AVAILABILITY_GUARD(35)
timezone_t _Nullable tzalloc(const char* _Nullable __id) __INTRODUCED_IN(35);

/**
 * tzfree(3) frees a timezone object returned by tzalloc().
 *
 * To use this with std::unique_ptr you'll want something like
 * `std::unique_ptr<std::remove_pointer_t<timezone_t>, decltype(&tzfree)> tz{tzalloc("Asia/Seoul"), tzfree};`
 * to remove the pointer.
 *
 * Available since API level 35.
 */
void tzfree(timezone_t _Nullable __tz) __INTRODUCED_IN(35);
#endif /* __BIONIC_AVAILABILITY_GUARD(35) */


/**
 * [clock(3)](https://man7.org/linux/man-pages/man3/clock.3.html)
 * returns an approximation of CPU time used, equivalent to
 * `clock_gettime(CLOCK_PROCESS_CPUTIME_ID)` but with more confusing
 * units. Use `CLOCKS_PER_SEC` to convert the result to seconds.
 *
 * Returns the time in seconds on success, and returns -1 and sets `errno` on failure.
 *
 * New code should prefer `clock_gettime(CLOCK_PROCESS_CPUTIME_ID)`.
 */
clock_t clock(void);

/**
 * [clock_getcpuclockid(3)](https://man7.org/linux/man-pages/man3/clock_getcpuclockid.3.html)
 * gets the clock ID of the cpu-time clock for the given `pid`.
 *
 * Returns 0 on success, and returns -1 and returns an error number on failure.
 */

#if __BIONIC_AVAILABILITY_GUARD(23)
int clock_getcpuclockid(pid_t __pid, clockid_t* _Nonnull __clock) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


/**
 * [clock_getres(2)](https://man7.org/linux/man-pages/man2/clock_getres.2.html)
 * gets the resolution of the given clock.
 *
 * Returns 0 on success, and returns -1 and returns an error number on failure.
 */
int clock_getres(clockid_t __clock, struct timespec* _Nullable __resolution);

/**
 * [clock_gettime(2)](https://man7.org/linux/man-pages/man2/clock_gettime.2.html)
 * gets the time according to the given clock.
 *
 * Returns 0 on success, and returns -1 and returns an error number on failure.
 */
int clock_gettime(clockid_t __clock, struct timespec* _Nonnull __ts);

/**
 * [clock_nanosleep(2)](https://man7.org/linux/man-pages/man2/clock_nanosleep.2.html)
 * sleeps for the given time (or until the given time if the TIMER_ABSTIME flag
 * is used), as measured by the given clock.
 *
 * Returns 0 on success, and returns -1 and returns an error number on failure.
 * If the sleep was interrupted by a signal, the return value will be `EINTR`
 * and `remainder` will be the amount of time remaining.
 */
int clock_nanosleep(clockid_t __clock, int __flags, const struct timespec* _Nonnull __time, struct timespec* _Nullable __remainder);

/**
 * [clock_settime(2)](https://man7.org/linux/man-pages/man2/clock_settime.2.html)
 * sets the time for the given clock.
 *
 * Returns 0 on success, and returns -1 and returns an error number on failure.
 */
int clock_settime(clockid_t __clock, const struct timespec* _Nonnull __ts);

/**
 * [timer_create(2)](https://man7.org/linux/man-pages/man2/timer_create.2.html)
 * creates a POSIX timer.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int timer_create(clockid_t __clock, struct sigevent* _Nullable __event, timer_t _Nonnull * _Nonnull __timer_ptr);

/**
 * [timer_delete(2)](https://man7.org/linux/man-pages/man2/timer_delete.2.html)
 * destroys a POSIX timer.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int timer_delete(timer_t _Nonnull __timer);

/**
 * [timer_settime(2)](https://man7.org/linux/man-pages/man2/timer_settime.2.html)
 * starts or stops a POSIX timer.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int timer_settime(timer_t _Nonnull __timer, int __flags, const struct itimerspec* _Nonnull __new_value, struct itimerspec* _Nullable __old_value);

/**
 * [timer_gettime(2)](https://man7.org/linux/man-pages/man2/timer_gettime.2.html)
 * gets the time until the given timer next fires.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 */
int timer_gettime(timer_t _Nonnull _timer, struct itimerspec* _Nonnull __ts);

/**
 * [timer_getoverrun(2)](https://man7.org/linux/man-pages/man2/timer_getoverrun.2.html)
 * gets the overrun count (the number of times the timer should have fired, but
 * didn't) for the last time the timer fired.
 *
 * Returns the overrun count on success, and returns -1 and sets `errno` on failure.
 */
int timer_getoverrun(timer_t _Nonnull __timer);

/**
 * The timebase for timespec_get() and timespec_getres() corresponding to CLOCK_REALTIME.
 *
 * Available since API level 29.
 */
#define TIME_UTC (CLOCK_REALTIME+1)

/**
 * The timebase for timespec_get() and timespec_getres() corresponding to CLOCK_MONOTONIC.
 *
 * Available since API level 35.
 */
#define TIME_MONOTONIC (CLOCK_MONOTONIC+1)

/**
 * The timebase for timespec_get() and timespec_getres() corresponding to CLOCK_PROCESS_CPUTIME_ID.
 *
 * Available since API level 35.
 */
#define TIME_ACTIVE (CLOCK_PROCESS_CPUTIME_ID+1)

/**
 * The timebase for timespec_get() and timespec_getres() corresponding to CLOCK_THREAD_CPUTIME_ID.
 *
 * Available since API level 35.
 */
#define TIME_THREAD_ACTIVE (CLOCK_THREAD_CPUTIME_ID+1)

/**
 * timespec_get(3) is equivalent to clock_gettime() for the clock corresponding to the given base.
 *
 * Returns the base on success and returns 0 on failure.
 *
 * Available since API level 29 for TIME_UTC; other bases arrived later.
 * Code for Android should prefer clock_gettime().
 */

#if __BIONIC_AVAILABILITY_GUARD(29)
int timespec_get(struct timespec* _Nonnull __ts, int __base) __INTRODUCED_IN(29);
#endif /* __BIONIC_AVAILABILITY_GUARD(29) */


/**
 * timespec_getres(3) is equivalent to clock_getres() for the clock corresponding to the given base.
 *
 * Returns the base on success and returns 0 on failure.
 *
 * Available since API level 35.
 * Code for Android should prefer clock_gettime().
 */

#if __BIONIC_AVAILABILITY_GUARD(35)
int timespec_getres(struct timespec* _Nonnull __ts, int __base) __INTRODUCED_IN(35);
#endif /* __BIONIC_AVAILABILITY_GUARD(35) */


__END_DECLS
