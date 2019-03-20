	.file	"optimizeTest.c"
	.text
	.section	.rodata.str1.1,"aMS",@progbits,1
.LC0:
	.string	"%d"
.LC1:
	.string	"%d\n"
	.section	.text.startup,"ax",@progbits
	.p2align 4,,15
	.globl	main
	.type	main, @function
main:
.LFB23:
	.cfi_startproc
	subq	$40, %rsp
	.cfi_def_cfa_offset 48
	leaq	.LC0(%rip), %rdi
	movq	%fs:40, %rax
	movq	%rax, 24(%rsp)
	xorl	%eax, %eax
	leaq	20(%rsp), %rsi
	movl	$111, 8(%rsp)
	call	__isoc99_scanf@PLT
	movl	20(%rsp), %edx
	testl	%edx, %edx
	jle	.L12
	movl	$111, %eax
	leaq	12(%rsp), %rdi
	leaq	16(%rsp), %rsi
	movl	%eax, %ecx
	shrl	$31, %ecx
	addl	%ecx, %eax
	andl	$1, %eax
	cmpl	%ecx, %eax
	jne	.L4
	.p2align 4,,10
	.p2align 3
.L16:
	subl	$1, %edx
	movl	$166, 12(%rsp)
	je	.L15
	movq	%rsi, %rax
	movl	(%rax), %eax
.L17:
	movl	%eax, %ecx
	shrl	$31, %ecx
	addl	%ecx, %eax
	andl	$1, %eax
	cmpl	%ecx, %eax
	je	.L16
.L4:
	subl	$1, %edx
	movl	$155, 12(%rsp)
	je	.L10
	movq	%rdi, %rax
	movl	(%rax), %eax
	jmp	.L17
.L15:
	movl	$166, %edx
.L6:
	movl	$0, 20(%rsp)
	.p2align 4,,10
	.p2align 3
.L3:
	leaq	.LC1(%rip), %rsi
	movl	$1, %edi
	xorl	%eax, %eax
	call	__printf_chk@PLT
	xorl	%eax, %eax
	movq	24(%rsp), %rdi
	xorq	%fs:40, %rdi
	jne	.L18
	addq	$40, %rsp
	.cfi_remember_state
	.cfi_def_cfa_offset 8
	ret
.L10:
	.cfi_restore_state
	movl	$155, %edx
	jmp	.L6
.L12:
	movl	12(%rsp), %edx
	jmp	.L3
.L18:
	call	__stack_chk_fail@PLT
	.cfi_endproc
.LFE23:
	.size	main, .-main
	.ident	"GCC: (Ubuntu 7.3.0-27ubuntu1~18.04) 7.3.0"
	.section	.note.GNU-stack,"",@progbits
