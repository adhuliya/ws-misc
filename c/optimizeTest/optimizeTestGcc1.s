	.file	"optimizeTest1.c"
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
	movl	$111, 12(%rsp)
	call	__isoc99_scanf@PLT
	movl	20(%rsp), %eax
	movl	16(%rsp), %edx
	testl	%eax, %eax
	je	.L3
	movl	$111, %edx
	leaq	12(%rsp), %rcx
	leaq	16(%rsp), %rsi
	testl	%edx, %edx
	jne	.L4
	.p2align 4,,10
	.p2align 3
.L15:
	subl	$1, %eax
	movl	$166, 16(%rsp)
	je	.L14
	movl	(%rcx), %edx
.L16:
	testl	%edx, %edx
	je	.L15
.L4:
	subl	$1, %eax
	movl	$155, 16(%rsp)
	je	.L9
	movq	%rsi, %rcx
	movl	(%rcx), %edx
	jmp	.L16
.L14:
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
	jne	.L17
	addq	$40, %rsp
	.cfi_remember_state
	.cfi_def_cfa_offset 8
	ret
.L9:
	.cfi_restore_state
	movl	$155, %edx
	jmp	.L6
.L17:
	call	__stack_chk_fail@PLT
	.cfi_endproc
.LFE23:
	.size	main, .-main
	.ident	"GCC: (Ubuntu 7.3.0-27ubuntu1~18.04) 7.3.0"
	.section	.note.GNU-stack,"",@progbits
