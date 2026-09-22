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
 * @file bits/bionic_multibyte_result.h
 * @brief Named values for the magic number return values of multibyte
 * conversion APIs defined by C.
 */

#include <sys/cdefs.h>

#include <stddef.h>

__BEGIN_DECLS

/**
 * @brief The error values defined by C for multibyte conversion APIs.
 *
 * Refer to C23 7.30.1 Restartable multibyte/wide character conversion functions
 * for more details.
 */
enum : size_t {
  /// @brief An encoding error occurred. The bytes read are not a valid unicode
  /// character, nor are they a partially valid character.
  BIONIC_MULTIBYTE_RESULT_ILLEGAL_SEQUENCE = -1UL,
#define BIONIC_MULTIBYTE_RESULT_ILLEGAL_SEQUENCE BIONIC_MULTIBYTE_RESULT_ILLEGAL_SEQUENCE

  /// @brief The bytes read may produce a valid unicode character, but the
  /// sequence is incomplete. Future calls may complete the character.
  BIONIC_MULTIBYTE_RESULT_INCOMPLETE_SEQUENCE = -2UL,
#define BIONIC_MULTIBYTE_RESULT_INCOMPLETE_SEQUENCE BIONIC_MULTIBYTE_RESULT_INCOMPLETE_SEQUENCE

  /// @brief The output of the call was the result of a previous successful
  /// decoding. No new bytes were consumed.
  ///
  /// The common case for this return value is when mbrtoc16 returns the low
  /// surrogate of a pair.
  BIONIC_MULTIBYTE_RESULT_NO_BYTES_CONSUMED = -3UL,
#define BIONIC_MULTIBYTE_RESULT_NO_BYTES_CONSUMED BIONIC_MULTIBYTE_RESULT_NO_BYTES_CONSUMED
};

__END_DECLS
