//go:build !precision_double

#include "textflag.h"

// Trig constants (float64)
DATA invpio2<>+0x00(SB)/8, $0x3FE45F306DC9C883 // 2/π = 0.6366197723675814
DATA pio2_hi<>+0x00(SB)/8, $0x3FF921FB50000000 // π/2 high = 1.5707963109016418
DATA pio2_lo<>+0x00(SB)/8, $0x3E5110B4611A6263 // π/2 low  = 1.5893254773528197e-8

// Sin polynomial coefficients (float64)
DATA s1<>+0x00(SB)/8, $0xBFC5555554CBAC77 // S1 = -0.16666666641626524
DATA s2<>+0x00(SB)/8, $0x3F811110896EFBB2 // S2 =  0.008333329385889463
DATA s3<>+0x00(SB)/8, $0xBF2A01A019BFDF03 // S3 = -1.9841269829589539e-4
DATA s4<>+0x00(SB)/8, $0x3EC6CD878C3B46A7 // S4 =  2.7183114939898219e-6

// Cos polynomial coefficients (float64)
DATA c0<>+0x00(SB)/8, $0xBFDFFFFFFD0C5E81 // C0 = -0.499999997251031
DATA c1<>+0x00(SB)/8, $0x3FA55553E1053A42 // C1 =  0.04166662332373906
DATA c2<>+0x00(SB)/8, $0xBF56C087E80F1E48 // C2 = -1.388676377461e-3
DATA c3<>+0x00(SB)/8, $0x3EF99342E0EE5069 // C3 =  2.439044879627741e-5

DATA one<>+0x00(SB)/8, $0x3FF0000000000000 // 1.0
DATA signmask<>+0x00(SB)/8, $0x8000000000000000 // sign bit mask (float64)

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
GLOBL signmask<>(SB), RODATA|NOPTR, $8

// func sin32(x float32) float32
TEXT ·sin32(SB),NOSPLIT,$0-12
	// Load x as uint32 for bit manipulation
	MOVL	x+0(FP), AX
	MOVL	AX, DX            // DX = save original bits for sign extraction
	ANDL	$0x7FFFFFFF, AX   // AX = |x| as uint32

	// NaN/Inf check: |x| >= 0x7F800000
	CMPL	AX, $0x7F800000
	JAE	sin32_nan

	// Move |x| into XMM as float32, then convert to float64
	MOVQ	AX, X0
	CVTSS2SD	X0, X0    // X0 = |x| as float64

	// Range reduction: n = round(|x| * 2/π)
	MOVSD	invpio2<>(SB), X1
	MULSD	X0, X1            // X1 = |x| * 2/π
	CVTSD2SL	X1, AX    // AX = n = round(|x| * 2/π) as int32
	CVTSL2SD	AX, X1    // X1 = float64(n)

	// r = |x| - n*pio2_hi - n*pio2_lo
	MOVSD	pio2_hi<>(SB), X2
	MULSD	X1, X2            // X2 = n * pio2_hi
	SUBSD	X2, X0            // X0 = |x| - n*pio2_hi
	MOVSD	pio2_lo<>(SB), X2
	MULSD	X1, X2            // X2 = n * pio2_lo
	SUBSD	X2, X0            // X0 = r

	// r2 = r*r
	MOVSD	X0, X1
	MULSD	X0, X1            // X1 = r2

	// sinpoly = r + r*r2*(S1 + r2*(S2 + r2*(S3 + r2*S4)))
	MOVSD	s4<>(SB), X2
	MULSD	X1, X2            // r2*S4
	ADDSD	s3<>(SB), X2      // S3 + r2*S4
	MULSD	X1, X2            // r2*(S3 + r2*S4)
	ADDSD	s2<>(SB), X2      // S2 + ...
	MULSD	X1, X2            // r2*(S2 + ...)
	ADDSD	s1<>(SB), X2      // S1 + ...
	MULSD	X1, X2            // r2*(S1 + ...)
	MULSD	X0, X2            // r*r2*(S1 + ...)
	ADDSD	X0, X2            // X2 = sinpoly

	// cospoly = 1 + r2*(C0 + r2*(C1 + r2*(C2 + r2*C3)))
	MOVSD	c3<>(SB), X3
	MULSD	X1, X3            // r2*C3
	ADDSD	c2<>(SB), X3      // C2 + r2*C3
	MULSD	X1, X3            // r2*(C2 + ...)
	ADDSD	c1<>(SB), X3      // C1 + ...
	MULSD	X1, X3            // r2*(C1 + ...)
	ADDSD	c0<>(SB), X3      // C0 + ...
	MULSD	X1, X3            // r2*(C0 + ...)
	ADDSD	one<>(SB), X3     // X3 = cospoly

	// Select: if n&1==0, result=sinpoly; else result=cospoly
	TESTL	$1, AX
	JNZ	sin32_use_cos
	MOVSD	X2, X4            // result = sinpoly
	JMP	sin32_sign
sin32_use_cos:
	MOVSD	X3, X4            // result = cospoly

sin32_sign:
	// Negate result if n&2 != 0
	TESTL	$2, AX
	JZ	sin32_input_sign
	MOVQ	X4, BX
	MOVQ	signmask<>(SB), CX
	XORQ	CX, BX
	MOVQ	BX, X4

sin32_input_sign:
	// XOR with sign of original x (bit 31 → bit 63)
	SHRL	$31, DX           // DX = 0 or 1
	SHLQ	$63, DX           // DX = 0 or 0x8000000000000000
	MOVQ	X4, BX
	XORQ	DX, BX
	MOVQ	BX, X4

	// Convert to float32 and return
	CVTSD2SS	X4, X4
	MOVSS	X4, ret+8(FP)
	RET

sin32_nan:
	// Return NaN: x - x
	MOVSS	x+0(FP), X0
	SUBSS	X0, X0
	MOVSS	X0, ret+8(FP)
	RET

// func cos32(x float32) float32
TEXT ·cos32(SB),NOSPLIT,$0-12
	// Load x as uint32 for bit manipulation
	MOVL	x+0(FP), AX
	ANDL	$0x7FFFFFFF, AX   // AX = |x| (cos is even, sign doesn't matter)

	// NaN/Inf check
	CMPL	AX, $0x7F800000
	JAE	cos32_nan

	// Move |x| into XMM as float32, then convert to float64
	MOVQ	AX, X0
	CVTSS2SD	X0, X0    // X0 = |x| as float64

	// Range reduction: n = round(|x| * 2/π)
	MOVSD	invpio2<>(SB), X1
	MULSD	X0, X1
	CVTSD2SL	X1, AX    // AX = n
	CVTSL2SD	AX, X1    // X1 = float64(n)

	// r = |x| - n*pio2_hi - n*pio2_lo
	MOVSD	pio2_hi<>(SB), X2
	MULSD	X1, X2
	SUBSD	X2, X0
	MOVSD	pio2_lo<>(SB), X2
	MULSD	X1, X2
	SUBSD	X2, X0            // X0 = r

	// r2 = r*r
	MOVSD	X0, X1
	MULSD	X0, X1            // X1 = r2

	// sinpoly (same as sin32)
	MOVSD	s4<>(SB), X2
	MULSD	X1, X2
	ADDSD	s3<>(SB), X2
	MULSD	X1, X2
	ADDSD	s2<>(SB), X2
	MULSD	X1, X2
	ADDSD	s1<>(SB), X2
	MULSD	X1, X2
	MULSD	X0, X2
	ADDSD	X0, X2            // X2 = sinpoly

	// cospoly (same as sin32)
	MOVSD	c3<>(SB), X3
	MULSD	X1, X3
	ADDSD	c2<>(SB), X3
	MULSD	X1, X3
	ADDSD	c1<>(SB), X3
	MULSD	X1, X3
	ADDSD	c0<>(SB), X3
	MULSD	X1, X3
	ADDSD	one<>(SB), X3     // X3 = cospoly

	// For cos: if n&1==0, result=cospoly; else result=sinpoly (swapped vs sin)
	TESTL	$1, AX
	JNZ	cos32_use_sin
	MOVSD	X3, X4            // result = cospoly
	JMP	cos32_sign
cos32_use_sin:
	MOVSD	X2, X4            // result = sinpoly

cos32_sign:
	// For cos: negate if (n+1)&2 != 0
	ADDL	$1, AX
	TESTL	$2, AX
	JZ	cos32_done
	MOVQ	X4, BX
	MOVQ	signmask<>(SB), CX
	XORQ	CX, BX
	MOVQ	BX, X4

cos32_done:
	CVTSD2SS	X4, X4
	MOVSS	X4, ret+8(FP)
	RET

cos32_nan:
	MOVSS	x+0(FP), X0
	SUBSS	X0, X0
	MOVSS	X0, ret+8(FP)
	RET

// func sincos32(x float32) (float32, float32)
TEXT ·sincos32(SB),NOSPLIT,$0-16
	// Load x as uint32
	MOVL	x+0(FP), AX
	MOVL	AX, DX            // DX = save original bits for sign
	ANDL	$0x7FFFFFFF, AX   // AX = |x|

	// NaN/Inf check
	CMPL	AX, $0x7F800000
	JAE	sincos32_nan

	// Move |x| into XMM as float32, then convert to float64
	MOVQ	AX, X0
	CVTSS2SD	X0, X0    // X0 = |x| as float64

	// Range reduction
	MOVSD	invpio2<>(SB), X1
	MULSD	X0, X1
	CVTSD2SL	X1, AX    // AX = n
	CVTSL2SD	AX, X1    // X1 = float64(n)

	// r = |x| - n*pio2_hi - n*pio2_lo
	MOVSD	pio2_hi<>(SB), X2
	MULSD	X1, X2
	SUBSD	X2, X0
	MOVSD	pio2_lo<>(SB), X2
	MULSD	X1, X2
	SUBSD	X2, X0            // X0 = r

	// r2 = r*r
	MOVSD	X0, X1
	MULSD	X0, X1            // X1 = r2

	// sinpoly
	MOVSD	s4<>(SB), X2
	MULSD	X1, X2
	ADDSD	s3<>(SB), X2
	MULSD	X1, X2
	ADDSD	s2<>(SB), X2
	MULSD	X1, X2
	ADDSD	s1<>(SB), X2
	MULSD	X1, X2
	MULSD	X0, X2
	ADDSD	X0, X2            // X2 = sinpoly

	// cospoly
	MOVSD	c3<>(SB), X3
	MULSD	X1, X3
	ADDSD	c2<>(SB), X3
	MULSD	X1, X3
	ADDSD	c1<>(SB), X3
	MULSD	X1, X3
	ADDSD	c0<>(SB), X3
	MULSD	X1, X3
	ADDSD	one<>(SB), X3     // X3 = cospoly

	// --- Compute sin result ---
	// Select: if n&1==0, sin_result=sinpoly; else sin_result=cospoly
	TESTL	$1, AX
	JNZ	sincos32_sin_use_cos
	MOVSD	X2, X4            // sin_result = sinpoly
	JMP	sincos32_sin_sign
sincos32_sin_use_cos:
	MOVSD	X3, X4            // sin_result = cospoly

sincos32_sin_sign:
	// Negate sin_result if n&2 != 0
	TESTL	$2, AX
	JZ	sincos32_sin_input_sign
	MOVQ	X4, BX
	MOVQ	signmask<>(SB), CX
	XORQ	CX, BX
	MOVQ	BX, X4

sincos32_sin_input_sign:
	// XOR with sign of original x
	MOVL	DX, SI            // SI = original bits
	SHRL	$31, SI
	SHLQ	$63, SI
	MOVQ	X4, BX
	XORQ	SI, BX
	MOVQ	BX, X4            // X4 = sin result (float64)

	// --- Compute cos result ---
	// Select: if n&1==0, cos_result=cospoly; else cos_result=sinpoly
	TESTL	$1, AX
	JNZ	sincos32_cos_use_sin
	MOVSD	X3, X5            // cos_result = cospoly
	JMP	sincos32_cos_sign
sincos32_cos_use_sin:
	MOVSD	X2, X5            // cos_result = sinpoly

sincos32_cos_sign:
	// Negate cos_result if (n+1)&2 != 0
	LEAL	1(AX), CX         // CX = n+1 (doesn't modify AX)
	TESTL	$2, CX
	JZ	sincos32_done
	MOVQ	X5, BX
	MOVQ	signmask<>(SB), CX
	XORQ	CX, BX
	MOVQ	BX, X5

sincos32_done:
	// Convert and return: ret+8 = sin, ret1+12 = cos
	CVTSD2SS	X4, X4
	CVTSD2SS	X5, X5
	MOVSS	X4, ret+8(FP)
	MOVSS	X5, ret1+12(FP)
	RET

sincos32_nan:
	MOVSS	x+0(FP), X0
	SUBSS	X0, X0
	MOVSS	X0, ret+8(FP)
	MOVSS	X0, ret1+12(FP)
	RET
