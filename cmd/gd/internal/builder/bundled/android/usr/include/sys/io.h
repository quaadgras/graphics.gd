/*
 * Copyright (C) 2024 The Android Open Source Project
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
 * @file sys/io.h
 * @brief The x86/x86-64 I/O port functions.
 */

#include <sys/cdefs.h>

#include <errno.h>
#include <sys/syscall.h>
#include <unistd.h>

__BEGIN_DECLS

/**
 * [iopl(2)](https://man7.org/linux/man-pages/man2/iopl.2.html) changes the I/O
 * privilege level for all x86/x8-64 I/O ports, for the calling thread.
 *
 * New callers should use ioperm() instead.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 *
 * Only available for x86/x86-64.
 */
#if defined(__NR_iopl)
__attribute__((__deprecated__("use ioperm() instead"))) static __inline int iopl(int __level) {
  return syscall(__NR_iopl, __level);
}
#endif

/**
 * [ioperm(2)](https://man7.org/linux/man-pages/man2/ioperm.2.html) sets the I/O
 * permissions for the given number of x86/x86-64 I/O ports, starting at the
 * given port.
 *
 * Returns 0 on success, and returns -1 and sets `errno` on failure.
 *
 * Only available for x86/x86-64.
 */
#if defined(__NR_iopl)
static __inline int ioperm(unsigned long __from, unsigned long __n, int __enabled) {
  return syscall(__NR_ioperm, __from, __n, __enabled);
}
#endif

/**
 * [inb(2)](https://man7.org/linux/man-pages/man2/inb.2.html)
 * reads a byte from the given x86/x86-64 I/O port.
 *
 * Only available for x86/x86-64.
 */
#if defined(__i386__) || defined(__x86_64__)
static __inline unsigned char inb(unsigned short __port) {
  unsigned char __value;
  __asm__ __volatile__("inb %1, %0" : "=a"(__value) : "dN"(__port));
  return __value;
}
#endif

/**
 * [inw(2)](https://man7.org/linux/man-pages/man2/inw.2.html)
 * reads a 16-bit "word" from the given x86/x86-64 I/O port.
 *
 * Only available for x86/x86-64.
 */
#if defined(__i386__) || defined(__x86_64__)
static __inline unsigned short inw(unsigned short __port) {
  unsigned short __value;
  __asm__ __volatile__("inw %1, %0" : "=a"(__value) : "dN"(__port));
  return __value;
}
#endif

/**
 * [inl(2)](https://man7.org/linux/man-pages/man2/inl.2.html)
 * reads a 32-bit "long word" from the given x86/x86-64 I/O port.
 *
 * Only available for x86/x86-64.
 */
#if defined(__i386__) || defined(__x86_64__)
static __inline unsigned int inl(unsigned short __port) {
  unsigned int __value;
  __asm__ __volatile__("inl %1, %0" : "=a"(__value) : "dN"(__port));
  return __value;
}
#endif

/**
 * [outb(2)](https://man7.org/linux/man-pages/man2/outb.2.html)
 * writes the given byte to the given x86/x86-64 I/O port.
 *
 * Only available for x86/x86-64.
 */
#if defined(__i386__) || defined(__x86_64__)
static __inline void outb(unsigned char __value, unsigned short __port) {
  __asm__ __volatile__("outb %0, %1" : : "a"(__value), "dN"(__port));
}
#endif

/**
 * [outw(2)](https://man7.org/linux/man-pages/man2/outw.2.html)
 * writes the given 16-bit "word" to the given x86/x86-64 I/O port.
 *
 * Only available for x86/x86-64.
 */
#if defined(__i386__) || defined(__x86_64__)
static __inline void outw(unsigned short __value, unsigned short __port) {
  __asm__ __volatile__("outw %0, %1" : : "a"(__value), "dN"(__port));
}
#endif

/**
 * [outl(2)](https://man7.org/linux/man-pages/man2/outl.2.html)
 * writes the given 32-bit "long word" to the given x86/x86-64 I/O port.
 *
 * Only available for x86/x86-64.
 */
#if defined(__i386__) || defined(__x86_64__)
static __inline void outl(unsigned int __value, unsigned short __port) {
  __asm__ __volatile__("outl %0, %1" : : "a"(__value), "dN"(__port));
}
#endif

__END_DECLS
