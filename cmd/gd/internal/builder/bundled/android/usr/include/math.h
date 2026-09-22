/*
 * ====================================================
 * Copyright (C) 1993 by Sun Microsystems, Inc. All rights reserved.
 *
 * Developed at SunPro, a Sun Microsystems, Inc. business.
 * Permission to use, copy, modify, and distribute this
 * software is freely granted, provided that this notice
 * is preserved.
 * ====================================================
 */

/*
 * Originally based on fdlibm.h 5.1 via FreeBSD.
 */

#pragma once

#include <sys/cdefs.h>
#include <limits.h>

__BEGIN_DECLS

/* C11. */

typedef double __double_t;
typedef __double_t double_t;
typedef float __float_t;
typedef __float_t float_t;

#define HUGE_VAL __builtin_huge_val()
#define HUGE_VALF __builtin_huge_valf()
#define HUGE_VALL __builtin_huge_vall()

#define INFINITY __builtin_inff()

#define NAN __builtin_nanf("")

#define FP_INFINITE 0x01
#define FP_NAN 0x02
#define FP_NORMAL 0x04
#define FP_SUBNORMAL 0x08
#define FP_ZERO 0x10

#if defined(__FP_FAST_FMA)
#define FP_FAST_FMA 1
#endif
#if defined(__FP_FAST_FMAF)
#define FP_FAST_FMAF 1
#endif
#if defined(__FP_FAST_FMAL)
#define FP_FAST_FMAL 1
#endif

#define FP_ILOGB0 (-INT_MAX)
#define FP_ILOGBNAN INT_MAX

#define MATH_ERRNO 1
#define MATH_ERREXCEPT 2
#define math_errhandling MATH_ERREXCEPT

#define fpclassify(x) __builtin_fpclassify(FP_NAN, FP_INFINITE, FP_NORMAL, FP_SUBNORMAL, FP_ZERO, x)

#define isfinite(x) __builtin_isfinite(x)

#define isinf(x) __builtin_isinf(x)

#define isnan(x) __builtin_isnan(x)

#define isnormal(x) __builtin_isnormal(x)

#define signbit(x) __builtin_signbit(x)

double acos(double __x);
float acosf(float __x);
long double acosl(long double __x);

double asin(double __x);
float asinf(float __x);
long double asinl(long double __x);

double atan(double __x);
float atanf(float __x);
long double atanl(long double __x);

double atan2(double __y, double __x);
float atan2f(float __y, float __x);
long double atan2l(long double __y, long double __x);

double cos(double __x);
float cosf(float __x);
long double cosl(long double __x);

double sin(double __x);
float sinf(float __x);
long double sinl(long double __x);

double tan(double __x);
float tanf(float __x);
long double tanl(long double __x);

double acosh(double __x);
float acoshf(float __x);
long double acoshl(long double __x);

double asinh(double __x);
float asinhf(float __x);
long double asinhl(long double __x);

double atanh(double __x);
float atanhf(float __x);
long double atanhl(long double __x);

double cosh(double __x);
float coshf(float __x);
long double coshl(long double __x);

double sinh(double __x);
float sinhf(float __x);
long double sinhl(long double __x);

double tanh(double __x);
float tanhf(float __x);
long double tanhl(long double __x);

double exp(double __x);
float expf(float __x);
long double expl(long double __x);

double exp2(double __x);
float exp2f(float __x);
long double exp2l(long double __x);

double expm1(double __x);
float expm1f(float __x);
long double expm1l(long double __x);

double frexp(double __x, int* _Nonnull __exponent);
float frexpf(float __x, int* _Nonnull __exponent);
long double frexpl(long double __x, int* _Nonnull __exponent);

int ilogb(double __x) __attribute_const__;
int ilogbf(float __x) __attribute_const__;
int ilogbl(long double __x) __attribute_const__;

double ldexp(double __x, int __exponent);
float ldexpf(float __x, int __exponent);
long double ldexpl(long double __x, int __exponent);

double log(double __x);
float logf(float __x);
long double logl(long double __x);

double log10(double __x);
float log10f(float __x);
long double log10l(long double __x);

double log1p(double __x);
float log1pf(float __x);
long double log1pl(long double __x);

double log2(double __x);
float log2f(float __x);
long double log2l(long double __x);

double logb(double __x);
float logbf(float __x);
long double logbl(long double __x);

double modf(double __x, double* _Nonnull __integral_part);
float modff(float __x, float* _Nonnull __integral_part);
long double modfl(long double __x, long double* _Nonnull __integral_part);

double scalbn(double __x, int __exponent);
float scalbnf(float __x, int __exponent);
long double scalbnl(long double __x, int __exponent);

double scalbln(double __x, long __exponent);
float scalblnf(float __x, long __exponent);
long double scalblnl(long double __x, long __exponent);

double cbrt(double __x);
float cbrtf(float __x);
long double cbrtl(long double __x);

double fabs(double __x) __attribute_const__;
float fabsf(float __x) __attribute_const__;
long double fabsl(long double __x) __attribute_const__;

double hypot(double __x, double __y);
float hypotf(float __x, float __y);
long double hypotl(long double __x, long double __y);

double pow(double __x, double __y);
float powf(float __x, float __y);
long double powl(long double __x, long double __y);

double sqrt(double __x);
float sqrtf(float __x);
long double sqrtl(long double __x);

double erf(double __x);
float erff(float __x);
long double erfl(long double __x);

double erfc(double __x);
float erfcf(float __x);
long double erfcl(long double __x);

double lgamma(double __x);
float lgammaf(float __x);
long double lgammal(long double __x);

double tgamma(double __x);
float tgammaf(float __x);
long double tgammal(long double __x);

double ceil(double __x);
float ceilf(float __x);
long double ceill(long double __x);

double floor(double __x);
float floorf(float __x);
long double floorl(long double __x);

double nearbyint(double __x);
float nearbyintf(float __x);
long double nearbyintl(long double __x);

double rint(double __x);
float rintf(float __x);
long double rintl(long double __x);

long lrint(double __x);
long lrintf(float __x);
long lrintl(long double __x);

long long llrint(double __x);
long long llrintf(float __x);
long long llrintl(long double __x);

double round(double __x);
float roundf(float __x);
long double roundl(long double __x);

long lround(double __x);
long lroundf(float __x);
long lroundl(long double __x);

long long llround(double __x);
long long llroundf(float __x);
long long llroundl(long double __x);

double trunc(double __x);
float truncf(float __x);
long double truncl(long double __x);

double fmod(double __x, double __y);
float fmodf(float __x, float __y);
long double fmodl(long double __x, long double __y);

double remainder(double __x, double __y);
float remainderf(float __x, float __y);
long double remainderl(long double __x, long double __y);

double remquo(double __x, double __y, int* _Nonnull __quotient_bits);
float remquof(float __x, float __y, int* _Nonnull __quotient_bits);
long double remquol(long double __x, long double __y, int* _Nonnull __quotient_bits);

double copysign(double __value, double __sign) __attribute_const__;
float copysignf(float __value, float __sign) __attribute_const__;
long double copysignl(long double __value, long double __sign) __attribute_const__;

double nan(const char* _Nonnull __kind) __attribute_const__;
float nanf(const char* _Nonnull __kind) __attribute_const__;
long double nanl(const char* _Nonnull __kind) __attribute_const__;

double nextafter(double __x, double __y);
float nextafterf(float __x, float __y);
long double nextafterl(long double __x, long double __y);

double nexttoward(double __x, long double __y);
float nexttowardf(float __x, long double __y);
long double nexttowardl(long double __x, long double __y);

double fdim(double __x, double __y);
float fdimf(float __x, float __y);
long double fdiml(long double __x, long double __y);

double fmax(double __x, double __y) __attribute_const__;
float fmaxf(float __x, float __y) __attribute_const__;
long double fmaxl(long double __x, long double __y) __attribute_const__;

double fmin(double __x, double __y) __attribute_const__;
float fminf(float __x, float __y) __attribute_const__;
long double fminl(long double __x, long double __y) __attribute_const__;

double fma(double __x, double __y, double __z);
float fmaf(float __x, float __y, float __z);
long double fmal(long double __x, long double __y, long double __z);

#define isgreater(x, y) __builtin_isgreater((x), (y))
#define isgreaterequal(x, y) __builtin_isgreaterequal((x), (y))
#define isless(x, y) __builtin_isless((x), (y))
#define islessequal(x, y) __builtin_islessequal((x), (y))
#define islessgreater(x, y) __builtin_islessgreater((x), (y))
#define isunordered(x, y) __builtin_isunordered((x), (y))

/* POSIX extensions. */

extern int signgam;

double j0(double __x);
double j1(double __x);
double jn(int __n, double __x);
double y0(double __x);
double y1(double __x);
double yn(int __n, double __x);

#define M_E		2.7182818284590452354	/* e */
#define M_LOG2E		1.4426950408889634074	/* log 2e */
#define M_LOG10E	0.43429448190325182765	/* log 10e */
#define M_LN2		0.69314718055994530942	/* log e2 */
#define M_LN10		2.30258509299404568402	/* log e10 */
#define M_PI		3.14159265358979323846	/* pi */
#define M_PI_2		1.57079632679489661923	/* pi/2 */
#define M_PI_4		0.78539816339744830962	/* pi/4 */
#define M_1_PI		0.31830988618379067154	/* 1/pi */
#define M_2_PI		0.63661977236758134308	/* 2/pi */
#define M_2_SQRTPI	1.12837916709551257390	/* 2/sqrt(pi) */
#define M_SQRT2		1.41421356237309504880	/* sqrt(2) */
#define M_SQRT1_2	0.70710678118654752440	/* 1/sqrt(2) */

#define MAXFLOAT	((float)3.40282346638528860e+38)

/* BSD extensions. */

#if defined(__USE_BSD)
#define HUGE MAXFLOAT
#endif

/* Extensions in both BSD and GNU. */

#if defined(__USE_BSD) || defined(__USE_GNU)
double gamma(double __x);
double scalb(double __x, double __exponent);
double drem(double __x, double __y);
int finite(double __x) __attribute_const__;
int isinff(float __x) __attribute_const__;
int isnanf(float __x) __attribute_const__;
double gamma_r(double __x, int* _Nonnull __sign);
double lgamma_r(double __x, int* _Nonnull __sign);
double significand(double __x);

#if __BIONIC_AVAILABILITY_GUARD(23)
long double lgammal_r(long double __x, int* _Nonnull __sign) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */

long double significandl(long double __x);
float dremf(float __x, float __y);
int finitef(float __x) __attribute_const__;
float gammaf(float __x);
float j0f(float __x);
float j1f(float __x);
float jnf(int __n, float __x);
float scalbf(float __x, float __exponent);
float y0f(float __x);
float y1f(float __x);
float ynf(int __n, float __x);
float gammaf_r(float __x, int* _Nonnull __sign);
float lgammaf_r(float __x, int* _Nonnull __sign);
float significandf(float __x);
void sincos(double __x, double* _Nonnull __sin, double* _Nonnull __cos);
void sincosf(float __x, float* _Nonnull __sin, float* _Nonnull __cos);
void sincosl(long double __x, long double* _Nonnull __sin, long double* _Nonnull __cos);
#endif

/* GNU extensions. */

#if defined(__USE_GNU)
#define M_El            2.718281828459045235360287471352662498L /* e */
#define M_LOG2El        1.442695040888963407359924681001892137L /* log 2e */
#define M_LOG10El       0.434294481903251827651128918916605082L /* log 10e */
#define M_LN2l          0.693147180559945309417232121458176568L /* log e2 */
#define M_LN10l         2.302585092994045684017991454684364208L /* log e10 */
#define M_PIl           3.141592653589793238462643383279502884L /* pi */
#define M_PI_2l         1.570796326794896619231321691639751442L /* pi/2 */
#define M_PI_4l         0.785398163397448309615660845819875721L /* pi/4 */
#define M_1_PIl         0.318309886183790671537767526745028724L /* 1/pi */
#define M_2_PIl         0.636619772367581343075535053490057448L /* 2/pi */
#define M_2_SQRTPIl     1.128379167095512573896158903121545172L /* 2/sqrt(pi) */
#define M_SQRT2l        1.414213562373095048801688724209698079L /* sqrt(2) */
#define M_SQRT1_2l      0.707106781186547524400844362104849039L /* 1/sqrt(2) */
int isinfl(long double __x) __attribute_const__;
int isnanl(long double __x) __attribute_const__;
#endif

__END_DECLS
