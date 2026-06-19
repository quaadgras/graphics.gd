/*-*- mode:unix-assembly; indent-tabs-mode:t; tab-width:8; coding:utf-8     -*-│
│ vi: set et ft=asm ts=8 tw=8 fenc=utf-8                                   :vi │
╞══════════════════════════════════════════════════════════════════════════════╡
│ Copyright 2024 Justine Alexandra Roberts Tunney                              │
│                                                                              │
│ Permission to use, copy, modify, and/or distribute this software for         │
│ any purpose with or without fee is hereby granted, provided that the         │
│ above copyright notice and this permission notice appear in all copies.      │
│                                                                              │
│ THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL                │
│ WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED                │
│ WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE             │
│ AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL         │
│ DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR        │
│ PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER               │
│ TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR             │
│ PERFORMANCE OF THIS SOFTWARE.                                                │
╚─────────────────────────────────────────────────────────────────────────────*/

/*
 * Foreign function trampoline with TLS switching.
 *
 * A static musl binary calls into dynamically loaded glibc libraries through
 * stubs (see foreign_wrap in dlopen.c) that jump here with the real function
 * pointer in a scratch register. We:
 *
 *   1. Save the integer + SSE/FP argument registers (and the variadic vector
 *      count / struct-return pointer).
 *   2. Resolve THIS thread's glibc TLS via get_thread_foreign_tls() -- keyed by
 *      kernel tid, so it is correct no matter which TLS is currently active.
 *   3. Save the current TLS, switch to the glibc TLS, COPY the caller's
 *      stack-passed arguments into our frame, call the real function, then
 *      restore the saved TLS.
 *
 * This is an ordinary stack-framed function, so each (possibly nested)
 * invocation keeps its saved-TLS and return address in its OWN machine-stack
 * frame and nesting to any depth just works -- no per-thread nesting stack, no
 * slot index, no search on return.
 *
 * Building a frame relocates the caller's stack-passed arguments (the 7th+
 * integer arg, 9th+ SSE arg), so we copy a fixed window of them down to where
 * the real function expects them ([rsp+8] upward). STACK_ARG_SLOTS=16 covers
 * any realistic foreign function (XCreateWindow, glTexImage2D, ...); extra
 * copied words past the real argument count are simply ignored by the callee.
 *
 * This file is only compiled when targeting musl (not glibc).
 */

#ifndef __GLIBC__

#define STACK_ARG_SLOTS 16

	.text
	.extern	get_thread_foreign_tls

#ifdef __x86_64__

	.hidden	foreign_tramp
	.globl	foreign_tramp
	.type	foreign_tramp, @function
foreign_tramp:
	/*
	 * On entry: %r11 = real function pointer; integer args in
	 * %rdi/%rsi/%rdx/%rcx/%r8/%r9; SSE args in %xmm0-7; %al = number of
	 * vector registers used (variadic); return address at (%rsp); any
	 * stack-passed args at 8(%rsp), 16(%rsp), ...
	 *
	 * Frame (offsets from %rbp, 16-byte aligned):
	 *   -8   real_func        -48  rcx         -96..-208  xmm0..xmm7
	 *   -16  rax (al / ret)   -56  r8          -384..-264 copied stack args
	 *   -24  rdi (/ ret rdx)  -64  r9
	 *   -32  rsi              -72  foreign_tls
	 *   -40  rdx              -80  entry_tls
	 */
	push	%rbp
	mov	%rsp,%rbp
	sub	$384,%rsp

	mov	%r11,-8(%rbp)		/* real_func */
	mov	%rax,-16(%rbp)		/* variadic vector count (%al) */
	mov	%rdi,-24(%rbp)
	mov	%rsi,-32(%rbp)
	mov	%rdx,-40(%rbp)
	mov	%rcx,-48(%rbp)
	mov	%r8,-56(%rbp)
	mov	%r9,-64(%rbp)
	movaps	%xmm0,-96(%rbp)
	movaps	%xmm1,-112(%rbp)
	movaps	%xmm2,-128(%rbp)
	movaps	%xmm3,-144(%rbp)
	movaps	%xmm4,-160(%rbp)
	movaps	%xmm5,-176(%rbp)
	movaps	%xmm6,-192(%rbp)
	movaps	%xmm7,-208(%rbp)

	/* Block Go's async signals (SIGURG/SIGPROF) for the whole switched window so
	 * their handlers can't run while %fs points at the glibc TLS -- they read g
	 * from %fs-relative TLS and would deref garbage. Also covers the brief switch
	 * inside get_thread_foreign_tls' on-demand path. Old mask saved at -224.
	 * rt_sigprocmask(SIG_BLOCK, {SIGURG,SIGPROF}, &old, 8). */
	mov	$0x4400000,%eax		/* (1<<22 SIGURG)|(1<<26 SIGPROF) */
	mov	%rax,-216(%rbp)
	xor	%edi,%edi		/* SIG_BLOCK */
	lea	-216(%rbp),%rsi
	lea	-224(%rbp),%rdx
	mov	$8,%r10d
	mov	$14,%eax			/* rt_sigprocmask */
	syscall

	/* entry_tls = %fs base (self pointer at %fs:0); pass it as the lookup key */
	mov	%fs:0,%rdi
	mov	%rdi,-80(%rbp)

	/* foreign_tls = get_thread_foreign_tls(entry_tls) -- no syscall on hot path */
	call	get_thread_foreign_tls
	mov	%rax,-72(%rbp)

	/* arch_prctl(ARCH_SET_FS, foreign_tls) */
	mov	-72(%rbp),%rsi
	mov	$0x1002,%edi
	mov	$158,%eax
	syscall

	/* Copy the caller's stack args [rbp+16 ..] down to [rbp-384 ..] so the
	 * real function finds them at (rsp+8) once we move rsp there. */
	mov	16(%rbp),%rax  ; mov %rax,-384(%rbp)
	mov	24(%rbp),%rax  ; mov %rax,-376(%rbp)
	mov	32(%rbp),%rax  ; mov %rax,-368(%rbp)
	mov	40(%rbp),%rax  ; mov %rax,-360(%rbp)
	mov	48(%rbp),%rax  ; mov %rax,-352(%rbp)
	mov	56(%rbp),%rax  ; mov %rax,-344(%rbp)
	mov	64(%rbp),%rax  ; mov %rax,-336(%rbp)
	mov	72(%rbp),%rax  ; mov %rax,-328(%rbp)
	mov	80(%rbp),%rax  ; mov %rax,-320(%rbp)
	mov	88(%rbp),%rax  ; mov %rax,-312(%rbp)
	mov	96(%rbp),%rax  ; mov %rax,-304(%rbp)
	mov	104(%rbp),%rax ; mov %rax,-296(%rbp)
	mov	112(%rbp),%rax ; mov %rax,-288(%rbp)
	mov	120(%rbp),%rax ; mov %rax,-280(%rbp)
	mov	128(%rbp),%rax ; mov %rax,-272(%rbp)
	mov	136(%rbp),%rax ; mov %rax,-264(%rbp)

	/* restore args + variadic vector count */
	mov	-24(%rbp),%rdi
	mov	-32(%rbp),%rsi
	mov	-40(%rbp),%rdx
	mov	-48(%rbp),%rcx
	mov	-56(%rbp),%r8
	mov	-64(%rbp),%r9
	movaps	-96(%rbp),%xmm0
	movaps	-112(%rbp),%xmm1
	movaps	-128(%rbp),%xmm2
	movaps	-144(%rbp),%xmm3
	movaps	-160(%rbp),%xmm4
	movaps	-176(%rbp),%xmm5
	movaps	-192(%rbp),%xmm6
	movaps	-208(%rbp),%xmm7
	mov	-16(%rbp),%rax		/* %al */
	mov	-8(%rbp),%r11		/* real_func */

	/* point rsp at the copied stack args (16-byte aligned) and call */
	lea	-384(%rbp),%rsp
	call	*%r11

	/* stash integer return values across the restore syscall */
	mov	%rax,-16(%rbp)
	mov	%rdx,-24(%rbp)

	/* arch_prctl(ARCH_SET_FS, entry_tls) */
	mov	-80(%rbp),%rsi
	mov	$0x1002,%edi
	mov	$158,%eax
	syscall

	/* Restore the signal mask: rt_sigprocmask(SIG_SETMASK, &old, NULL, 8). */
	mov	$2,%edi			/* SIG_SETMASK */
	lea	-224(%rbp),%rsi
	xor	%edx,%edx
	mov	$8,%r10d
	mov	$14,%eax
	syscall

	mov	-16(%rbp),%rax
	mov	-24(%rbp),%rdx

	mov	%rbp,%rsp
	pop	%rbp
	ret
	.size	foreign_tramp, . - foreign_tramp

#elif defined(__aarch64__)

	.globl	foreign_tramp
	.type	foreign_tramp, @function
foreign_tramp:
	/*
	 * On entry: x9 = real function pointer; integer args in x0-x7; FP/SIMD
	 * args in v0-v7; x8 = indirect result location; x30 = return address;
	 * stack-passed args at [sp], [sp+8], ...
	 *
	 * Frame (sp = bottom, 384 bytes):
	 *   [sp+0   .. sp+127]  copied stack args (rsp at blr points here)
	 *   [sp+128 .. sp+191]  x0-x7
	 *   [sp+192]            x8
	 *   [sp+208 .. sp+335]  v0-v7
	 * real_func / foreign_tls / entry_tls live in callee-saved x19/x20/x21.
	 */
	stp	x29,x30,[sp,#-16]!	/* x29 = caller stack args at [x29+16] */
	mov	x29,sp
	stp	x19,x20,[sp,#-16]!
	stp	x21,x22,[sp,#-16]!
	sub	sp,sp,#384

	stp	x0,x1,[sp,#128]
	stp	x2,x3,[sp,#144]
	stp	x4,x5,[sp,#160]
	stp	x6,x7,[sp,#176]
	str	x8,[sp,#192]
	stp	q0,q1,[sp,#208]
	stp	q2,q3,[sp,#240]
	stp	q4,q5,[sp,#272]
	stp	q6,q7,[sp,#304]

	mov	x19,x9			/* real_func (survives the calls) */

	/* Block Go's async signals (SIGURG/SIGPROF) for the switched window so their
	 * handlers can't run while tpidr_el0 points at the glibc TLS. x9 is free now.
	 * rt_sigprocmask(SIG_BLOCK, {SIGURG,SIGPROF}, &old@sp+344, 8). */
	movz	x9,#0x440,lsl #16	/* 0x4400000 = (1<<22 SIGURG)|(1<<26 SIGPROF) */
	str	x9,[sp,#336]
	mov	x0,#0			/* SIG_BLOCK */
	add	x1,sp,#336
	add	x2,sp,#344
	mov	x3,#8
	mov	x8,#135			/* rt_sigprocmask */
	svc	#0

	mrs	x21,tpidr_el0		/* entry_tls */
	mov	x0,x21			/* pass entry_tls as the lookup key */
	bl	get_thread_foreign_tls
	mov	x20,x0			/* foreign_tls */
	msr	tpidr_el0,x20		/* switch to foreign TLS */

	ldp	x0,x1,[sp,#128]
	ldp	x2,x3,[sp,#144]
	ldp	x4,x5,[sp,#160]
	ldp	x6,x7,[sp,#176]
	ldr	x8,[sp,#192]
	ldp	q0,q1,[sp,#208]
	ldp	q2,q3,[sp,#240]
	ldp	q4,q5,[sp,#272]
	ldp	q6,q7,[sp,#304]

	/* copy caller stack args [x29+16 ..] -> [sp+0 ..] */
	add	x10,x29,#16
	ldp	x12,x13,[x10,#0]   ; stp x12,x13,[sp,#0]
	ldp	x12,x13,[x10,#16]  ; stp x12,x13,[sp,#16]
	ldp	x12,x13,[x10,#32]  ; stp x12,x13,[sp,#32]
	ldp	x12,x13,[x10,#48]  ; stp x12,x13,[sp,#48]
	ldp	x12,x13,[x10,#64]  ; stp x12,x13,[sp,#64]
	ldp	x12,x13,[x10,#80]  ; stp x12,x13,[sp,#80]
	ldp	x12,x13,[x10,#96]  ; stp x12,x13,[sp,#96]
	ldp	x12,x13,[x10,#112] ; stp x12,x13,[sp,#112]

	blr	x19			/* sp points at copied stack args */

	str	x0,[sp,#352]		/* stash int return values across the syscall */
	str	x1,[sp,#360]		/* (v0/v1 FP returns survive a syscall) */
	msr	tpidr_el0,x21		/* restore entry TLS */

	/* Restore the signal mask: rt_sigprocmask(SIG_SETMASK, &old@sp+344, NULL, 8). */
	mov	x0,#2			/* SIG_SETMASK */
	add	x1,sp,#344
	mov	x2,#0
	mov	x3,#8
	mov	x8,#135
	svc	#0
	ldr	x0,[sp,#352]
	ldr	x1,[sp,#360]

	add	sp,sp,#384
	ldp	x21,x22,[sp],#16
	ldp	x19,x20,[sp],#16
	ldp	x29,x30,[sp],#16
	ret
	.size	foreign_tramp, . - foreign_tramp

#endif

#endif /* __GLIBC__ */
