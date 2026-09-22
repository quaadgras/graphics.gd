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

#include <sys/cdefs.h>
#include <stdint.h>

__BEGIN_DECLS

struct tcphdr {
  __extension__ union {
    struct {
      uint16_t th_sport;
      uint16_t th_dport;
      uint32_t th_seq;
      uint32_t th_ack;
      uint8_t th_x2:4;
      uint8_t th_off:4;
      uint8_t th_flags;
      uint16_t th_win;
      uint16_t th_sum;
      uint16_t th_urp;
    };
    struct {
      uint16_t source;
      uint16_t dest;
      uint32_t seq;
      uint32_t ack_seq;
      uint16_t res1:4;
      uint16_t doff:4;
      uint16_t fin:1;
      uint16_t syn:1;
      uint16_t rst:1;
      uint16_t psh:1;
      uint16_t ack:1;
      uint16_t urg:1;
      uint16_t res2:2;
      uint16_t window;
      uint16_t check;
      uint16_t urg_ptr;
    };
  };
};

__END_DECLS
