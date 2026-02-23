//go:build !precision_double

#include "textflag.h"

// Trig constants (float64)
DATA invpio2<>+0x00(SB)/8, $0x3FE45F306DC9C883 // 2/π
DATA pio2_hi<>+0x00(SB)/8, $0x3FF921FB50000000 // π/2 high
DATA pio2_lo<>+0x00(SB)/8, $0x3E5110B4611A6263 // π/2 low

// Sin polynomial coefficients
DATA s1<>+0x00(SB)/8, $0xBFC5555554CBAC77
DATA s2<>+0x00(SB)/8, $0x3F811110896EFBB2
DATA s3<>+0x00(SB)/8, $0xBF2A01A019BFDF03
DATA s4<>+0x00(SB)/8, $0x3EC6CD878C3B46A7

// Cos polynomial coefficients
DATA c0<>+0x00(SB)/8, $0xBFDFFFFFFD0C5E81
DATA c1<>+0x00(SB)/8, $0x3FA55553E1053A42
DATA c2<>+0x00(SB)/8, $0xBF56C087E80F1E48
DATA c3<>+0x00(SB)/8, $0x3EF99342E0EE5069

DATA one<>+0x00(SB)/8, $0x3FF0000000000000

GLOBL invpio2<>(SB), RODATA|NOPTR, $8
GLOBL pio2_hi<>(SB), RODATA|NOPTR, $8
GLOBL pio2_lo<>(SB), RODATA|NOPTR, $8
GLOBL s1<>(SB), RODATA|NOPTR, $8
GLOBL s2<>(SB), RODATA|NOPTR, $8
GLOBL s3<>(SB), RODATA|NOPTR, $8
GLOBL s4<>(SB), RODATA|NOPTR, $8
GLOBL c0<>(SB), RODATA|NOPTR, $8
GLOBL c1<>(SB), RODATA|NOPTR, $8
GLOBL c2<>(SB), RODATA|NOPTR, $8
GLOBL c3<>(SB), RODATA|NOPTR, $8
GLOBL one<>(SB), RODATA|NOPTR, $8

// func sin32(x float32) float32
TEXT ·sin32(SB),NOSPLIT,$0-12
	// Load x as uint32
	MOVWU	x+0(FP), R0        // R0 = x bits
	MOVW	R0, R9              // R9 = save original bits for sign
	ANDW	$0x7FFFFFFF, R0     // R0 = |x| bits

	// NaN/Inf check: |x| >= 0x7F800000
	MOVW	$0x7F800000, R1
	CMPW	R1, R0
	BGE	sin32_nan

	// Move |x| to float register, convert to float64
	FMOVS	R0, F0              // F0 = |x| as float32
	FCVTSD	F0, F0              // F0 = |x| as float64

	// Range reduction: n = round(|x| * 2/π)
	MOVD	invpio2<>(SB), R3
	FMOVD	R3, F4
	FMULD	F4, F0, F1          // F1 = |x| * 2/π
	FRINTND	F1, F1              // F1 = round(|x| * 2/π)
	FCVTZSDW	F1, R2      // R2 = n as int32
	SCVTFWD	R2, F1              // F1 = float64(n)

	// r = |x| - n*pio2_hi - n*pio2_lo
	MOVD	pio2_hi<>(SB), R3
	FMOVD	R3, F4
	FMULD	F1, F4, F4          // F4 = n * pio2_hi
	FSUBD	F4, F0, F0          // F0 = |x| - n*pio2_hi
	MOVD	pio2_lo<>(SB), R3
	FMOVD	R3, F4
	FMULD	F1, F4, F4          // F4 = n * pio2_lo
	FSUBD	F4, F0, F0          // F0 = r

	// r2 = r*r
	FMULD	F0, F0, F1          // F1 = r2

	// sinpoly = r + r*r2*(S1 + r2*(S2 + r2*(S3 + r2*S4)))
	MOVD	s4<>(SB), R3
	FMOVD	R3, F2              // F2 = S4
	MOVD	s3<>(SB), R3
	FMOVD	R3, F4              // F4 = S3
	FMADDD	F1, F2, F4, F2      // F2 = S3 + r2*S4
	MOVD	s2<>(SB), R3
	FMOVD	R3, F4              // F4 = S2
	FMADDD	F1, F2, F4, F2      // F2 = S2 + r2*(S3+r2*S4)
	MOVD	s1<>(SB), R3
	FMOVD	R3, F4              // F4 = S1
	FMADDD	F1, F2, F4, F2      // F2 = S1 + r2*(S2+...)
	FMULD	F1, F2, F2          // F2 = r2*(S1+...)
	FMADDD	F0, F2, F0, F2      // F2 = r + r*(r2*(S1+...)) = sinpoly

	// cospoly = 1 + r2*(C0 + r2*(C1 + r2*(C2 + r2*C3)))
	MOVD	c3<>(SB), R3
	FMOVD	R3, F3              // F3 = C3
	MOVD	c2<>(SB), R3
	FMOVD	R3, F4              // F4 = C2
	FMADDD	F1, F3, F4, F3      // F3 = C2 + r2*C3
	MOVD	c1<>(SB), R3
	FMOVD	R3, F4              // F4 = C1
	FMADDD	F1, F3, F4, F3      // F3 = C1 + r2*(C2+r2*C3)
	MOVD	c0<>(SB), R3
	FMOVD	R3, F4              // F4 = C0
	FMADDD	F1, F3, F4, F3      // F3 = C0 + r2*(C1+...)
	MOVD	one<>(SB), R3
	FMOVD	R3, F4              // F4 = 1.0
	FMADDD	F1, F3, F4, F3      // F3 = 1.0 + r2*(C0+...) = cospoly

	// Select: if n&1==0, result=sinpoly(F2); else result=cospoly(F3)
	ANDW	$1, R2, R3
	CBNZW	R3, sin32_use_cos
	FMOVD	F2, F5
	B	sin32_sign
sin32_use_cos:
	FMOVD	F3, F5

sin32_sign:
	// Sign correction: XOR of input sign and quadrant sign
	ANDW	$2, R2, R3          // R3 = n&2 (0 or 2)
	LSL	$62, R3, R3         // R3 = 0 or 0x8000000000000000
	LSR	$31, R9, R4         // R4 = sign bit of x (0 or 1)
	LSL	$63, R4, R4         // R4 = 0 or 0x8000000000000000
	EOR	R3, R4, R3          // R3 = combined sign mask
	FMOVD	F5, R4
	EOR	R3, R4, R4
	FMOVD	R4, F5

	// Convert to float32 and return
	FCVTDS	F5, F5
	FMOVS	F5, ret+8(FP)
	RET

sin32_nan:
	FMOVS	R0, F0
	FSUBS	F0, F0, F0
	FMOVS	F0, ret+8(FP)
	RET

// func cos32(x float32) float32
TEXT ·cos32(SB),NOSPLIT,$0-12
	// Load x as uint32
	MOVWU	x+0(FP), R0
	ANDW	$0x7FFFFFFF, R0     // |x| (cos is even)

	// NaN/Inf check
	MOVW	$0x7F800000, R1
	CMPW	R1, R0
	BGE	cos32_nan

	// Move |x| to float register, convert to float64
	FMOVS	R0, F0
	FCVTSD	F0, F0

	// Range reduction
	MOVD	invpio2<>(SB), R3
	FMOVD	R3, F4
	FMULD	F4, F0, F1
	FRINTND	F1, F1
	FCVTZSDW	F1, R2
	SCVTFWD	R2, F1

	MOVD	pio2_hi<>(SB), R3
	FMOVD	R3, F4
	FMULD	F1, F4, F4
	FSUBD	F4, F0, F0
	MOVD	pio2_lo<>(SB), R3
	FMOVD	R3, F4
	FMULD	F1, F4, F4
	FSUBD	F4, F0, F0          // F0 = r

	FMULD	F0, F0, F1          // F1 = r2

	// sinpoly
	MOVD	s4<>(SB), R3
	FMOVD	R3, F2
	MOVD	s3<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F2, F4, F2
	MOVD	s2<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F2, F4, F2
	MOVD	s1<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F2, F4, F2
	FMULD	F1, F2, F2
	FMADDD	F0, F2, F0, F2      // F2 = sinpoly

	// cospoly
	MOVD	c3<>(SB), R3
	FMOVD	R3, F3
	MOVD	c2<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F3, F4, F3
	MOVD	c1<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F3, F4, F3
	MOVD	c0<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F3, F4, F3
	MOVD	one<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F3, F4, F3      // F3 = cospoly

	// For cos: if n&1==0, result=cospoly; else result=sinpoly (swapped)
	ANDW	$1, R2, R3
	CBNZW	R3, cos32_use_sin
	FMOVD	F3, F5
	B	cos32_sign
cos32_use_sin:
	FMOVD	F2, F5

cos32_sign:
	// Negate if (n+1)&2 != 0
	ADDW	$1, R2, R3
	ANDW	$2, R3, R3
	LSL	$62, R3, R3
	FMOVD	F5, R4
	EOR	R3, R4, R4
	FMOVD	R4, F5

	FCVTDS	F5, F5
	FMOVS	F5, ret+8(FP)
	RET

cos32_nan:
	FMOVS	R0, F0
	FSUBS	F0, F0, F0
	FMOVS	F0, ret+8(FP)
	RET

// func sincos32(x float32) (float32, float32)
TEXT ·sincos32(SB),NOSPLIT,$0-16
	// Load x as uint32
	MOVWU	x+0(FP), R0
	MOVW	R0, R9              // save original bits
	ANDW	$0x7FFFFFFF, R0

	// NaN/Inf check
	MOVW	$0x7F800000, R1
	CMPW	R1, R0
	BGE	sincos32_nan

	// Move |x| to float register, convert to float64
	FMOVS	R0, F0
	FCVTSD	F0, F0

	// Range reduction
	MOVD	invpio2<>(SB), R3
	FMOVD	R3, F4
	FMULD	F4, F0, F1
	FRINTND	F1, F1
	FCVTZSDW	F1, R2
	SCVTFWD	R2, F1

	MOVD	pio2_hi<>(SB), R3
	FMOVD	R3, F4
	FMULD	F1, F4, F4
	FSUBD	F4, F0, F0
	MOVD	pio2_lo<>(SB), R3
	FMOVD	R3, F4
	FMULD	F1, F4, F4
	FSUBD	F4, F0, F0          // F0 = r

	FMULD	F0, F0, F1          // F1 = r2

	// sinpoly
	MOVD	s4<>(SB), R3
	FMOVD	R3, F2
	MOVD	s3<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F2, F4, F2
	MOVD	s2<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F2, F4, F2
	MOVD	s1<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F2, F4, F2
	FMULD	F1, F2, F2
	FMADDD	F0, F2, F0, F2      // F2 = sinpoly

	// cospoly
	MOVD	c3<>(SB), R3
	FMOVD	R3, F3
	MOVD	c2<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F3, F4, F3
	MOVD	c1<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F3, F4, F3
	MOVD	c0<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F3, F4, F3
	MOVD	one<>(SB), R3
	FMOVD	R3, F4
	FMADDD	F1, F3, F4, F3      // F3 = cospoly

	// --- Sin result ---
	ANDW	$1, R2, R3
	CBNZW	R3, sincos32_sin_use_cos
	FMOVD	F2, F5              // sin = sinpoly
	B	sincos32_sin_sign
sincos32_sin_use_cos:
	FMOVD	F3, F5              // sin = cospoly

sincos32_sin_sign:
	ANDW	$2, R2, R3
	LSL	$62, R3, R3
	LSR	$31, R9, R4
	LSL	$63, R4, R4
	EOR	R3, R4, R3
	FMOVD	F5, R4
	EOR	R3, R4, R4
	FMOVD	R4, F5              // F5 = sin result

	// --- Cos result ---
	ANDW	$1, R2, R3
	CBNZW	R3, sincos32_cos_use_sin
	FMOVD	F3, F6              // cos = cospoly
	B	sincos32_cos_sign
sincos32_cos_use_sin:
	FMOVD	F2, F6              // cos = sinpoly

sincos32_cos_sign:
	ADDW	$1, R2, R3
	ANDW	$2, R3, R3
	LSL	$62, R3, R3
	FMOVD	F6, R4
	EOR	R3, R4, R4
	FMOVD	R4, F6              // F6 = cos result

	// Convert and return
	FCVTDS	F5, F5
	FCVTDS	F6, F6
	FMOVS	F5, ret+8(FP)
	FMOVS	F6, ret1+12(FP)
	RET

sincos32_nan:
	FMOVS	R0, F0
	FSUBS	F0, F0, F0
	FMOVS	F0, ret+8(FP)
	FMOVS	F0, ret1+12(FP)
	RET
