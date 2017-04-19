	.text
	.file	"address.c"
	.globl	main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# BB#0:                                 # %entry
	pushq	%rbp
.Lcfi0:
	.cfi_def_cfa_offset 16
.Lcfi1:
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
.Lcfi2:
	.cfi_def_cfa_register %rbp
	subq	$32, %rsp
	movabsq	$.L.str, %rdi
	leaq	-8(%rbp), %rax
	movl	$0, -4(%rbp)
	movl	$15, -8(%rbp)
	movq	%rax, -16(%rbp)
	movq	-16(%rbp), %rsi
	movq	-16(%rbp), %rdx
	movq	-16(%rbp), %r8
	movq	%rax, %rcx
	movb	$0, %al
	callq	printf
	xorl	%r9d, %r9d
	movl	%eax, -20(%rbp)         # 4-byte Spill
	movl	%r9d, %eax
	addq	$32, %rsp
	popq	%rbp
	retq
.Lfunc_end0:
	.size	main, .Lfunc_end0-main
	.cfi_endproc

	.type	.L.str,@object          # @.str
	.section	.rodata.str1.1,"aMS",@progbits,1
.L.str:
	.asciz	"%p %p %p %p"
	.size	.L.str, 12


	.ident	"clang version 5.0.0 (trunk 292058)"
	.section	".note.GNU-stack","",@progbits
