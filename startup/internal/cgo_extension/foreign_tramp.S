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
 * Foreign function trampoline with TLS switching (multi-threaded).
 *
 * When calling from a static musl binary into dynamically loaded glibc
 * libraries, we switch to glibc's TLS before calling, then restore
 * musl's TLS afterwards. This is a minimal trampoline that:
 *
 *   1. Saves native TLS and return address to a per-thread TLS stack
 *   2. Switches to this thread's foreign TLS from the pool
 *   3. Calls the foreign function
 *   4. Pops from TLS stack and restores native TLS
 *
 * Each thread has its own slot in the TLS pool (assigned by the stub
 * before jumping here). The pool contains glibc TLS pointers and
 * tramp_ctx addresses for each slot.
 *
 * On entry:
 *   x86_64: %rax = function pointer to call
 *           stack: [slot] [return_addr]  (slot pushed by stub)
 *   aarch64: x8 = function pointer to call, x9 = slot index
 *
 * The __foreign struct layout (must match dlopen.c):
 *   offset 0:  atomic_uint once
 *   offset 8:  struct tls_pool *pool
 *   offset 16: void *foreign_tls (main thread's, for compat)
 *   offset 24: void *native_tls (main thread's)
 *   ...
 *
 * The tls_pool struct layout:
 *   offset 0:    void *tls_ptrs[64]    (glibc TLS pointers, 512 bytes)
 *   offset 512:  void *tramp_ctxs[64]  (per-thread __tramp_ctx addresses)
 *
 * The __tramp_ctx struct layout (per-thread, in glibc TLS):
 *   offset 0:  long sp (stack pointer, 0-15)
 *   offset 8:  void *stack[32] (pairs of native_tls, return_addr)
 *
 * This file is only compiled when targeting musl (not glibc).
 */

#ifndef __GLIBC__

	.text

	/* Declare external symbol */
	.extern	__foreign

#ifdef __x86_64__

	/* Make symbols hidden (private) to avoid conflicts */
	.hidden	foreign_tramp
	.hidden	__set_tls_internal

	.globl	foreign_tramp
	.type	foreign_tramp, @function
foreign_tramp:
	/*
	 * TLS switch with stack-based restore for nested calls.
	 * We push native_tls and return_addr onto a per-thread TLS stack,
	 * then jmp to the foreign function. When it returns, we pop from the stack.
	 *
	 * On entry:
	 *   %rax = function pointer to call
	 *   stack: [slot] [return_addr]  (slot was pushed by stub)
	 *
	 * __foreign layout:
	 *   offset 8:  pool pointer (struct tls_pool *)
	 *
	 * tls_pool layout:
	 *   offset 0:    tls_ptrs[64]    (512 bytes)
	 *   offset 512:  tramp_ctxs[64]  (512 bytes)
	 *
	 * __tramp_ctx layout (per-thread):
	 *   offset 0:  long sp (0-15)
	 *   offset 8:  void *stack[32] (pairs: native_tls, ret_addr)
	 */

	/*
	 * On entry: stack is [slot][ret_addr], rsp points to slot
	 * After both pops: rsp points to normal position with red zone below
	 *
	 * Red zone layout (all offsets from final rsp):
	 *   -8:  rdi (arg1)
	 *   -16: rsi (arg2)
	 *   -24: rdx (arg3)
	 *   -32: rcx (arg4)
	 *   -40: r8  (arg5)
	 *   -48: r9  (arg6)
	 *   -56: rax (function pointer)
	 *   -64: slot
	 *   -72: return address
	 *   -80: native TLS
	 */

	/* Pop slot and return address FIRST into scratch registers */
	pop	%r10			/* slot */
	pop	%r11			/* return address */

	/* Now rsp is in standard position - save everything with consistent offsets */
	mov	%rdi,-8(%rsp)
	mov	%rsi,-16(%rsp)
	mov	%rdx,-24(%rsp)
	mov	%rcx,-32(%rsp)
	mov	%r8,-40(%rsp)
	mov	%r9,-48(%rsp)
	mov	%rax,-56(%rsp)		/* function pointer */
	mov	%r10,-64(%rsp)		/* slot */
	mov	%r11,-72(%rsp)		/* return address */

	/* Save native TLS */
	mov	%fs:0,%r10
	mov	%r10,-80(%rsp)		/* native TLS */

	/* Load pool pointer */
	mov	__foreign+8(%rip),%r11	/* pool = __foreign.pool */

	/* Calculate foreign TLS: pool->tls_ptrs[slot] = pool + slot*8 */
	mov	-64(%rsp),%r10		/* reload slot */
	movslq	%r10d,%rcx		/* sign-extend slot to 64-bit */
	mov	(%r11,%rcx,8),%rsi	/* foreign_tls = pool->tls_ptrs[slot] */

	/* Switch to foreign TLS */
	mov	$0x1002,%edi
	mov	$158,%eax
	syscall

	/* Now on foreign TLS - get tramp_ctx for this slot */
	/* tramp_ctx = pool->tramp_ctxs[slot] = pool + 512 + slot*8 */
	mov	__foreign+8(%rip),%r11	/* reload pool */
	mov	-64(%rsp),%r10		/* reload slot */
	movslq	%r10d,%rcx
	lea	512(%r11,%rcx,8),%r11	/* &pool->tramp_ctxs[slot] */
	mov	(%r11),%r11		/* tramp_ctx = pool->tramp_ctxs[slot] */

	/* Push to TLS stack: load sp, store native_tls and ret_addr */
	mov	(%r11),%rcx		/* sp = pair index (0-15) */
	shl	$4,%rcx			/* sp * 16 (each pair is 16 bytes) */
	add	$8,%rcx			/* offset to stack[0] */
	/* Now store: native_tls at tramp_ctx+rcx, ret_addr at tramp_ctx+rcx+8 */
	mov	-80(%rsp),%r10		/* native TLS */
	mov	%r10,(%r11,%rcx,1)	/* store native_tls */
	mov	-72(%rsp),%r10		/* original return address */
	mov	%r10,8(%r11,%rcx,1)	/* store return_addr */
	/* Increment sp */
	mov	(%r11),%rcx		/* reload sp */
	inc	%rcx
	mov	%rcx,(%r11)		/* store sp+1 */

	/* Restore ALL args and function pointer */
	mov	-56(%rsp),%rax		/* function pointer */
	mov	-8(%rsp),%rdi
	mov	-16(%rsp),%rsi
	mov	-24(%rsp),%rdx
	mov	-32(%rsp),%rcx
	mov	-40(%rsp),%r8
	mov	-48(%rsp),%r9

	/* Push our restore handler as the return address */
	lea	.Lrestore(%rip),%r11
	push	%r11

	/* Tail call to foreign function - stack alignment preserved! */
	jmp	*%rax

.Lrestore:
	/* Foreign function returned here. rax/rdx have return values. */
	/* We're still in foreign TLS, with this thread's tramp_ctx. */
	/* We need to find the slot to look up tramp_ctx again... */
	/* The slot was stored in the TLS stack entry. */

	/* Save return values in red zone */
	mov	%rax,-8(%rsp)
	mov	%rdx,-16(%rsp)

	/* We need to find our tramp_ctx. On entry, we stored the slot in the
	 * TLS stack. But we need the slot to find the tramp_ctx!
	 *
	 * Solution: we're in foreign TLS, so we can read fs:0 to get foreign TLS,
	 * then search the pool to find our slot. But that's expensive.
	 *
	 * Better: we stored slot in the stack entry at offset 16. We need to
	 * find the tramp_ctx first using... the same chicken-and-egg problem.
	 *
	 * Alternative: store tramp_ctx address in the red zone before call.
	 * But we already used red zone and the foreign function may have clobbered it.
	 *
	 * Solution: Store tramp_ctx address at a fixed location that survives.
	 * We'll use fs:-8 (8 bytes before TLS base) which is typically safe.
	 * Actually, that's not guaranteed safe.
	 *
	 * Real solution: The foreign TLS pointer is unique to each slot.
	 * We can read fs:0 (our foreign TLS), then search pool->tls_ptrs[] for it.
	 * This is O(n) but n=64 and we only search once per return.
	 */

	/* Read current foreign TLS */
	mov	%fs:0,%r10

	/* Load pool pointer */
	mov	__foreign+8(%rip),%r11

	/* Search for our slot: find i where pool->tls_ptrs[i] == r10 */
	xor	%ecx,%ecx		/* i = 0 */
.Lsearch:
	cmp	(%r11,%rcx,8),%r10	/* pool->tls_ptrs[i] == foreign_tls? */
	je	.Lfound
	inc	%ecx
	cmp	$64,%ecx
	jb	.Lsearch
	/* Not found - this shouldn't happen, fall back to slot 0 */
	xor	%ecx,%ecx
.Lfound:
	/* ecx = slot */

	/* Get tramp_ctx for this slot */
	lea	512(%r11,%rcx,8),%r11	/* &pool->tramp_ctxs[slot] */
	mov	(%r11),%r11		/* tramp_ctx = pool->tramp_ctxs[slot] */

	/* Pop from TLS stack: decrement sp, load native_tls and ret_addr */
	mov	(%r11),%rcx		/* load sp */
	dec	%rcx			/* sp-- */
	mov	%rcx,(%r11)		/* store decremented sp */
	shl	$4,%rcx			/* sp * 16 */
	add	$8,%rcx			/* offset to stack[sp] */
	mov	8(%r11,%rcx,1),%r10	/* load return_addr from stack[sp]+8 */
	mov	%r10,-24(%rsp)		/* save to red zone */
	mov	(%r11,%rcx,1),%rsi	/* load native_tls for syscall */

	/* Restore native TLS */
	mov	$0x1002,%edi
	mov	$158,%eax
	syscall

	/* Restore return values */
	mov	-8(%rsp),%rax
	mov	-16(%rsp),%rdx

	/* Jump to original return address */
	mov	-24(%rsp),%r11
	jmp	*%r11
	.size	foreign_tramp, . - foreign_tramp

/*
 * Internal TLS setter - uses arch_prctl syscall on Linux x86_64
 * Input: %rdi = new TLS pointer
 */
	.globl	__set_tls_internal
	.type	__set_tls_internal, @function
__set_tls_internal:
	mov	%rdi,%rsi	/* TLS pointer as second arg */
	mov	$0x1002,%edi	/* ARCH_SET_FS as first arg */
	mov	$158,%eax	/* __NR_arch_prctl */
	syscall
	ret
	.size	__set_tls_internal, . - __set_tls_internal

#elif defined(__aarch64__)

	.globl	foreign_tramp
	.type	foreign_tramp, @function
foreign_tramp:
	/*
	 * TLS switch with stack-based restore for nested calls.
	 *
	 * On entry:
	 *   x8 = function pointer to call
	 *   x9 = slot index (which thread's TLS to use)
	 *
	 * __foreign layout:
	 *   offset 8:  pool pointer (struct tls_pool *)
	 *
	 * tls_pool layout:
	 *   offset 0:    tls_ptrs[64]    (512 bytes)
	 *   offset 512:  tramp_ctxs[64]  (512 bytes)
	 *
	 * __tramp_ctx layout (per-thread):
	 *   offset 0:  long sp (0-15)
	 *   offset 8:  void *stack[32] (pairs: native_tls, ret_addr)
	 */

	/* Save callee-saved registers we'll use */
	stp	x19,x20,[sp,#-48]!
	stp	x21,x22,[sp,#16]
	stp	x23,x30,[sp,#32]

	/* Save native TLS, function pointer, and slot */
	mrs	x19,tpidr_el0		/* native TLS */
	mov	x20,x8			/* function pointer */
	mov	x22,x9			/* slot */

	/* Load __foreign base address and pool */
	adrp	x10,__foreign
	add	x10,x10,:lo12:__foreign
	ldr	x10,[x10,8]		/* pool = __foreign.pool */

	/* Calculate foreign TLS: pool->tls_ptrs[slot] = pool + slot*8 */
	ldr	x11,[x10,x22,lsl #3]	/* foreign_tls = pool->tls_ptrs[slot] */
	msr	tpidr_el0,x11		/* switch to foreign TLS */

	/* Get tramp_ctx for this slot: pool->tramp_ctxs[slot] = pool + 512 + slot*8 */
	add	x21,x10,#512		/* &pool->tramp_ctxs[0] */
	ldr	x21,[x21,x22,lsl #3]	/* tramp_ctx = pool->tramp_ctxs[slot] */

	/* Push to TLS stack */
	ldr	x11,[x21]		/* sp */
	lsl	x12,x11,#4		/* sp * 16 */
	add	x12,x12,#8		/* offset to stack[sp] */
	add	x12,x21,x12		/* &tramp_ctx->stack[sp] */
	str	x19,[x12]		/* store native_tls */
	str	x30,[x12,#8]		/* store return address */
	add	x11,x11,#1		/* sp++ */
	str	x11,[x21]		/* store incremented sp */

	/* Call foreign function (x0-x7 args preserved) */
	blr	x20

	/* Pop from TLS stack - x21 still has tramp_ctx */
	ldr	x11,[x21]		/* load sp */
	sub	x11,x11,#1		/* sp-- */
	str	x11,[x21]		/* store decremented sp */
	lsl	x12,x11,#4		/* sp * 16 */
	add	x12,x12,#8		/* offset to stack[sp] */
	add	x12,x21,x12		/* &tramp_ctx->stack[sp] */
	ldr	x19,[x12]		/* load native_tls */

	/* Restore native TLS */
	msr	tpidr_el0,x19

	/* Restore callee-saved registers and return */
	ldp	x23,x30,[sp,#32]
	ldp	x21,x22,[sp,#16]
	ldp	x19,x20,[sp],#48
	ret
	.size	foreign_tramp, . - foreign_tramp

#endif

#endif /* __GLIBC__ */
