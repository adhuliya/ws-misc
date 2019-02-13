	.text
	.file	"main1.c"
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %entry
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	subq	$32, %rsp
	movabsq	$.L.str, %rax
	# movl	%edi, -4(%rbp)
	# movq	%rsi, -16(%rbp)
	# movl	-4(%rbp), %esi
	# movq	-16(%rbp), %rcx
	# movq	(%rcx), %rdx
  movabsq $.L.str1, %rdx
	movl	%edi, %esi
	movq	%rax, %rdi
	# movb	$0, %al
	callq	printf
	movl	%eax, -20(%rbp)         # 4-byte Spill
	addq	$32, %rsp
	popq	%rbp
	retq
.Lfunc_end0:
	.size	main, .Lfunc_end0-main
	.cfi_endproc
                                        # -- End function
	.type	.L.str,@object          # @.str
	.section	.rodata.str1.1,"aMS",@progbits,1
.L.str:
	.asciz	"%d, %s\n"
	.size	.L.str, 8

.L.str1:
	.asciz	"Hello"
	.size	.L.str, 8


	.ident	"clang version 6.0.1 (tags/RELEASE_601/final)"
	.section	".note.GNU-stack","",@progbits
