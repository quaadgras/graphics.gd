/*
  FastLZ - lightning-fast lossless compression library

  Copyright (C) 2007 Ariya Hidayat (ariya@kde.org)
  Copyright (C) 2006 Ariya Hidayat (ariya@kde.org)
  Copyright (C) 2005 Ariya Hidayat (ariya@kde.org)

  Permission is hereby granted, free of charge, to any person obtaining a copy
  of this software and associated documentation files (the "Software"), to deal
  in the Software without restriction, including without limitation the rights
  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
  copies of the Software, and to permit persons to whom the Software is
  furnished to do so, subject to the following conditions:

  The above copyright notice and this permission notice shall be included in
  all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
  THE SOFTWARE.
*/

package Packed

import (
	"math"
	"unsafe"
)

/*
 * Always check for bound when decompressing.
 * Generally it is best to leave it defined.
 */
// FASTLZ_SAFE - will leave it

/*
 * Give hints to the compiler for branch prediction optimization.
 */
// FASTLZ_EXPECT_CONDITIONAL - golang does not support

/*
 * Use inlined functions for supported systems.
 */
// FASTLZ_INLINE - golang does not support

/*
 * Prevent accessing more than 8-bit at once, except on x86 architectures.
 */
// FASTLZ_STRICT_ALIGN - will ONLY support x64 for golang !!!

/*
 * FIXME: use preprocessor magic to set this on different platforms!
 */
//type flzuint8 uint8
//type flzuint16 uint16
//type flzuint32 uint32

/* prototypes */
//int fastlz_compress(const void* input, int length, void* output);
//int fastlz_compress_level(int level, const void* input, int length, void* output);
//int fastlz_decompress(const void* input, int length, void* output, int maxout);

const fastlz_MAX_COPY = 32
const fastlz_MAX_LEN = 264        /* 256 + 8 */
const fastlz_MAX_DISTANCE1 = 8192 // MAX_DISTANCE

// FASTLZ_READU16 -  will expand into the code

const fastlz_HASH_LOG = 13
const fastlz_HASH_SIZE = (1 << fastlz_HASH_LOG)
const fastlz_HASH_MASK = (fastlz_HASH_SIZE - 1)

// HASH_FUNCTION - will expand into the code

const fastlz_MAX_DISTANCE2 = 8191 // MAX_DISTANCE
const fastlz_MAX_FARDISTANCE = (65535 + fastlz_MAX_DISTANCE2 - 1)

/*
*

	Compress a block of data in the input buffer and returns the size of
	compressed block. The size of input buffer is specified by length. The
	minimum input buffer size is 16.

	The output buffer must be at least 5% larger than the input buffer
	and can not be smaller than 66 bytes.

	If the input is not compressible, the return value might be larger than
	length (input buffer size).

	The input buffer and the output buffer can not overlap.
*/
func fastlz_compress(input []byte, output []byte) int {
	length := len(input)
	/* for short block, choose fastlz1 */
	if length < 65536 {
		return fastlz1_compress(input, length, output)
	}
	/* else... */
	return fastlz2_compress(input, length, output)
}

/*
*

	Decompress a block of compressed data and returns the size of the
	decompressed block. If error occurs, e.g. the compressed data is
	corrupted or the output buffer is not large enough, then 0 (zero)
	will be returned instead.

	The input buffer and the output buffer can not overlap.

	Decompression is memory safe and guaranteed not to write the output buffer
	more than what is specified in maxout.
*/
func fastlz_decompress(input []byte, output []byte) int {
	length := len(input)
	maxout := len(output)
	/* magic identifier for compression level */
	level := ((*(*uint8)(unsafe.Pointer(&input[0]))) >> 5) + 1

	if level == 1 {
		return fastlz1_decompress(input, length, output, maxout)
	}
	if level == 2 {
		return fastlz2_decompress(input, length, output, maxout)
	}
	/* unknown level, trigger error */
	return 0
}

/*
*

	Compress a block of data in the input buffer and returns the size of
	compressed block. The size of input buffer is specified by length. The
	minimum input buffer size is 16.

	The output buffer must be at least 5% larger than the input buffer
	and can not be smaller than 66 bytes.

	If the input is not compressible, the return value might be larger than
	length (input buffer size).

	The input buffer and the output buffer can not overlap.

	Compression level can be specified in parameter level. At the moment,
	only level 1 and level 2 are supported.
	Level 1 is the fastest compression and generally useful for short data.
	Level 2 is slightly slower but it gives better compression ratio.

	Note that the compressed data, regardless of the level, can always be
	decompressed using the function fastlz_decompress above.
*/
func fastlz_compress_level(level int, input []byte, length int, output []byte) int {
	if len(output) < int(math.Max(66, float64(length)*1.05)) {
		return 0
	}
	if level == 1 {
		return fastlz1_compress(input, length, output)
	}
	if level == 2 {
		return fastlz2_compress(input, length, output)
	}
	return 0
}

func fastlz1_compress(input []byte, length int, output []byte) int {
	var ip uint = 0
	var ip_bound uint = uint(length - 2)
	var ip_limit uint = uint(length - 12)
	var op uint = 0

	var htab [fastlz_HASH_SIZE]uint
	var hslot uint
	var hval uint

	var copy uint

	/* sanity check */
	if length < 4 {
		if length != 0 {
			/* create literal copy only */
			output[op] = byte(length - 1)
			op++
			ip_bound++
			for ip <= ip_bound {
				output[op] = input[ip]
				op++
				ip++
			}
			return length + 1
		} else {
			return 0
		}
	}

	/* initializes hash table */
	// do nothing

	/* we start with literal copy */
	copy = 2
	output[op] = fastlz_MAX_COPY - 1
	op++
	output[op] = input[ip]
	op++
	ip++
	output[op] = input[ip]
	op++
	ip++

	/* main loop */
	for ip < ip_limit {
		var ref uint
		var distance uint

		/* minimum match length */
		var len uint = 3

		/* comparison starting-point */
		anchor := ip

		/* check for a run */
		// do nothing

		/* find potential match */
		hval = (uint(input[ip]) | uint(input[ip+1])<<8)
		hval ^= (uint(input[ip+1]) | uint(input[ip+2])<<8) ^ (hval >> (16 - fastlz_HASH_LOG))
		hval &= fastlz_HASH_MASK

		hslot = hval
		ref = htab[hval]

		/* calculate distance to the match */
		distance = anchor - ref

		/* update hash table */
		htab[hslot] = anchor

		/* is this a match? check the first 3 bytes */
		if distance == 0 ||
			(distance >= fastlz_MAX_DISTANCE1) ||
			input[ref] != input[ip] || input[ref+1] != input[ip+1] || input[ref+2] != input[ip+2] {
			goto literal
		}

		/* last matched byte */
		ref += len
		ip = anchor + len

		/* distance is biased */
		distance--

		if distance == 0 {
			/* zero distance means a run */
			x := input[ip-1]
			for ip < ip_bound {
				if input[ref] != x {
					break
				} else {
					ip++
				}
				ref++
			}
		} else {
			for {
				/* safe because the outer check against ip limit */
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++

				for ip < ip_bound {
					if input[ref] != input[ip] {
						break
					}
					ref++
					ip++
				}
				break
			}
			ref++
			ip++
		}
		/* if we have copied something, adjust the copy count */
		if copy != 0 {
			/* copy is biased, '0' means 1 byte copy */
			output[op-copy-1] = byte(copy - 1)
		} else {
			/* back, to overwrite the copy count */
			op--
		}
		/* reset literal counter */
		copy = 0

		/* length is biased, '1' means a match of 3 bytes */
		ip -= 3
		len = ip - anchor

		/* encode the match */
		for len > fastlz_MAX_LEN-2 {
			output[op] = byte((7 << 5) + (distance >> 8))
			op++
			output[op] = fastlz_MAX_LEN - 2 - 7 - 2
			op++
			output[op] = byte(distance & 255)
			op++
			len -= fastlz_MAX_LEN - 2
		}

		if len < 7 {
			output[op] = byte((len << 5) + (distance >> 8))
			op++
			output[op] = byte(distance & 255)
			op++
		} else {
			output[op] = byte((7 << 5) + (distance >> 8))
			op++
			output[op] = byte(len - 7)
			op++
			output[op] = byte(distance & 255)
			op++
		}

		/* update the hash at match boundary */
		hval = (uint(input[ip]) | uint(input[ip+1])<<8)
		hval ^= (uint(input[ip+1]) | uint(input[ip+2])<<8) ^ (hval >> (16 - fastlz_HASH_LOG))
		hval &= fastlz_HASH_MASK
		htab[hval] = ip
		ip++
		hval = (uint(input[ip]) | uint(input[ip+1])<<8)
		hval ^= (uint(input[ip+1]) | uint(input[ip+2])<<8) ^ (hval >> (16 - fastlz_HASH_LOG))
		hval &= fastlz_HASH_MASK
		htab[hval] = ip
		ip++

		/* assuming literal copy */
		output[op] = fastlz_MAX_COPY - 1
		op++

		continue

	literal:
		output[op] = input[anchor]
		op++
		anchor++
		ip = anchor
		copy++
		if copy == fastlz_MAX_COPY {
			copy = 0
			output[op] = fastlz_MAX_COPY - 1
			op++
		}
	}

	/* left-over as literal copy */
	ip_bound++
	for ip <= ip_bound {
		output[op] = input[ip]
		op++
		ip++
		copy++
		if copy == fastlz_MAX_COPY {
			copy = 0
			output[op] = fastlz_MAX_COPY - 1
			op++
		}
	}

	/* if we have copied something, adjust the copy length */
	if copy != 0 {
		output[op-copy-1] = byte(copy - 1)
	} else {
		op--
	}

	return int(op)
}

func fastlz2_compress(input []byte, length int, output []byte) int {
	var ip uint = 0
	var ip_bound uint = uint(length - 2)
	var ip_limit uint = uint(length - 12)
	var op uint = 0

	var htab [fastlz_HASH_SIZE]uint
	var hslot uint
	var hval uint

	var copy uint

	/* sanity check */
	if length < 4 {
		if length != 0 {
			/* create literal copy only */
			output[op] = byte(length - 1)
			op++
			ip_bound++
			for ip <= ip_bound {
				output[op] = input[ip]
				op++
				ip++
			}
			return length + 1
		} else {
			return 0
		}
	}

	/* initializes hash table */
	// do nothing

	/* we start with literal copy */
	copy = 2
	output[op] = fastlz_MAX_COPY - 1
	op++
	output[op] = input[ip]
	op++
	ip++
	output[op] = input[ip]
	op++
	ip++

	/* main loop */
	for ip < ip_limit {
		var ref uint
		var distance uint

		/* minimum match length */
		var len uint = 3

		/* comparison starting-point */
		anchor := ip

		/* check for a run */
		if input[ip] == input[ip-1] && (uint(input[ip-1])|uint(input[ip])<<8) == (uint(input[ip+1])|uint(input[ip+2])<<8) {
			distance = 1
			ip += 3
			ref = anchor - 1 + 3
			goto match
		}

		/* find potential match */
		hval = (uint(input[ip]) | uint(input[ip+1])<<8)
		hval ^= (uint(input[ip+1]) | uint(input[ip+2])<<8) ^ (hval >> (16 - fastlz_HASH_LOG))
		hval &= fastlz_HASH_MASK

		hslot = hval
		ref = htab[hval]

		/* calculate distance to the match */
		distance = anchor - ref

		/* update hash table */
		htab[hslot] = anchor

		/* is this a match? check the first 3 bytes */
		if distance == 0 ||
			(distance >= fastlz_MAX_FARDISTANCE) ||
			input[ref] != input[ip] || input[ref+1] != input[ip+1] || input[ref+2] != input[ip+2] {
			goto literal
		}

		/* far, needs at least 5-byte match */
		if distance >= fastlz_MAX_DISTANCE2 {
			if input[ref+3] != input[ip+3] || input[ref+4] != input[ip+4] {
				goto literal
			}
			len += 2
		}

		ref += len

	match:

		/* last matched byte */
		ip = anchor + len

		/* distance is biased */
		distance--

		if distance == 0 {
			/* zero distance means a run */
			x := input[ip-1]
			for ip < ip_bound {
				if input[ref] != x {
					break
				} else {
					ip++
				}
				ref++
			}
		} else {
			for {
				/* safe because the outer check against ip limit */
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++
				if input[ref] != input[ip] {
					break
				}
				ref++
				ip++

				for ip < ip_bound {
					if input[ref] != input[ip] {
						break
					}
					ref++
					ip++
				}
				break
			}
			ref++
			ip++
		}
		/* if we have copied something, adjust the copy count */
		if copy != 0 {
			/* copy is biased, '0' means 1 byte copy */
			output[op-copy-1] = byte(copy - 1)
		} else {
			/* back, to overwrite the copy count */
			op--
		}
		/* reset literal counter */
		copy = 0

		/* length is biased, '1' means a match of 3 bytes */
		ip -= 3
		len = ip - anchor

		/* encode the match */
		if distance < fastlz_MAX_DISTANCE2 {
			if len < 7 {
				output[op] = byte((len << 5) + (distance >> 8))
				op++
				output[op] = byte(distance & 255)
				op++
			} else {
				output[op] = byte((7 << 5) + (distance >> 8))
				op++
				for len -= 7; len >= 255; len -= 255 {
					output[op] = 255
					op++
				}
				output[op] = byte(len)
				op++
				output[op] = byte(distance & 255)
				op++
			}
		} else {
			/* far away, but not yet in the another galaxy... */
			if len < 7 {
				distance -= fastlz_MAX_DISTANCE2
				output[op] = byte((len << 5) + 31)
				op++
				output[op] = 255
				op++
				output[op] = byte(distance >> 8)
				op++
				output[op] = byte(distance & 255)
				op++
			} else {
				distance -= fastlz_MAX_DISTANCE2
				output[op] = (7 << 5) + 31
				op++
				for len -= 7; len >= 255; len -= 255 {
					output[op] = 255
					op++
				}
				output[op] = byte(len)
				op++
				output[op] = 255
				op++
				output[op] = byte(distance >> 8)
				op++
				output[op] = byte(distance & 255)
				op++
			}
		}

		/* update the hash at match boundary */
		hval = (uint(input[ip]) | uint(input[ip+1])<<8)
		hval ^= (uint(input[ip+1]) | uint(input[ip+2])<<8) ^ (hval >> (16 - fastlz_HASH_LOG))
		hval &= fastlz_HASH_MASK
		htab[hval] = ip
		ip++
		hval = (uint(input[ip]) | uint(input[ip+1])<<8)
		hval ^= (uint(input[ip+1]) | uint(input[ip+2])<<8) ^ (hval >> (16 - fastlz_HASH_LOG))
		hval &= fastlz_HASH_MASK
		htab[hval] = ip
		ip++

		/* assuming literal copy */
		output[op] = fastlz_MAX_COPY - 1
		op++

		continue

	literal:
		output[op] = input[anchor]
		op++
		anchor++
		ip = anchor
		copy++
		if copy == fastlz_MAX_COPY {
			copy = 0
			output[op] = fastlz_MAX_COPY - 1
			op++
		}
	}

	/* left-over as literal copy */
	ip_bound++
	for ip <= ip_bound {
		output[op] = input[ip]
		op++
		ip++
		copy++
		if copy == fastlz_MAX_COPY {
			copy = 0
			output[op] = fastlz_MAX_COPY - 1
			op++
		}
	}

	/* if we have copied something, adjust the copy length */
	if copy != 0 {
		output[op-copy-1] = byte(copy - 1)
	} else {
		op--
	}

	/* marker for fastlz2 */
	output[0] |= (1 << 5)

	return int(op)
}

func fastlz1_decompress(input []byte, length int, output []byte, maxout int) int {
	var ip uint = 0
	var ip_limit uint = uint(length)
	var op uint = 0
	var op_limit uint = uint(maxout)
	var ctrl uint = uint(input[ip] & 31)
	ip++
	loop := true

	for loop {
		ref := op
		len := ctrl >> 5
		ofs := (ctrl & 31) << 8

		if ctrl >= 32 {
			len--
			ref -= ofs
			if len == 7-1 {
				len += uint(input[ip])
				ip++
			}
			ref -= uint(input[ip])
			ip++

			if op+len+3 > op_limit {
				return 0
			}
			if int(ref-1) < 0 {
				return 0
			}
			if ip < ip_limit {
				ctrl = uint(input[ip])
				ip++
			} else {
				loop = false
			}
			if ref == op {
				/* optimize copy for a run */
				b := output[ref-1]
				output[op] = b
				op++
				output[op] = b
				op++
				output[op] = b
				op++
				for ; len != 0; len-- {
					output[op] = b
					op++
				}
			} else {
				/* copy from reference */
				ref--

				output[op] = output[ref]
				op++
				ref++
				output[op] = output[ref]
				op++
				ref++
				output[op] = output[ref]
				op++
				ref++

				for ; len != 0; len-- {
					output[op] = output[ref]
					op++
					ref++
				}
			}
		} else {
			ctrl++
			if op+ctrl > op_limit {
				return 0
			}
			if ip+ctrl > ip_limit {
				return 0
			}
			output[op] = input[ip]
			op++
			ip++
			for ctrl--; ctrl != 0; ctrl-- {
				output[op] = input[ip]
				op++
				ip++
			}
			loop = (ip < ip_limit)
			if loop {
				ctrl = uint(input[ip])
				ip++
			}
		}
	}

	return int(op)
}

func fastlz2_decompress(input []byte, length int, output []byte, maxout int) int {
	var ip uint = 0
	var ip_limit uint = uint(length)
	var op uint = 0
	var op_limit uint = uint(maxout)
	var ctrl uint = uint(input[ip] & 31)
	ip++
	loop := true

	for loop {
		ref := op
		len := ctrl >> 5
		ofs := (ctrl & 31) << 8

		if ctrl >= 32 {
			var code byte
			len--
			ref -= ofs
			if len == 7-1 {
				code = 255
				for code == 255 {
					code = input[ip]
					ip++
					len += uint(code)
				}
			}
			code = input[ip]
			ip++
			ref -= uint(code)

			/* match from 16-bit distance */
			if code == 255 {
				if ofs == (31 << 8) {
					ofs = uint(input[ip]) << 8
					ip++
					ofs += uint(input[ip])
					ip++
					ref = op - ofs - fastlz_MAX_DISTANCE2
				}
			}
			if op+len+3 > op_limit {
				return 0
			}
			if int(ref-1) < 0 {
				return 0
			}
			if ip < ip_limit {
				ctrl = uint(input[ip])
				ip++
			} else {
				loop = false
			}
			if ref == op {
				/* optimize copy for a run */
				b := output[ref-1]
				output[op] = b
				op++
				output[op] = b
				op++
				output[op] = b
				op++
				for ; len != 0; len-- {
					output[op] = b
					op++
				}
			} else {
				/* copy from reference */
				ref--

				output[op] = output[ref]
				op++
				ref++
				output[op] = output[ref]
				op++
				ref++
				output[op] = output[ref]
				op++
				ref++

				for ; len != 0; len-- {
					output[op] = output[ref]
					op++
					ref++
				}
			}
		} else {
			ctrl++
			if op+ctrl > op_limit {
				return 0
			}
			if ip+ctrl > ip_limit {
				return 0
			}
			output[op] = input[ip]
			op++
			ip++
			for ctrl--; ctrl != 0; ctrl-- {
				output[op] = input[ip]
				op++
				ip++
			}
			loop = (ip < ip_limit)
			if loop {
				ctrl = uint(input[ip])
				ip++
			}
		}
	}

	return int(op)
}
