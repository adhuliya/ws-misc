	.text
	.file	"optimizeTest.c"
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %entry
	pushq	%rax
	.cfi_def_cfa_offset 16
	movq	%rsp, %rsi
	movl	$.L.str, %edi
	xorl	%eax, %eax
	callq	__isoc99_scanf
	movl	(%rsp), %eax
	testl	%eax, %eax
	jle	.LBB0_1
# %bb.2:                                # %while.body.lr.ph
	cmpl	$1, %eax
	jne	.LBB0_4
# %bb.3:
	addl	$-1, %eax
	movl	$155, %esi
	jmp	.LBB0_7
.LBB0_1:
                                        # implicit-def: %esi
	jmp	.LBB0_8
.LBB0_4:                                # %while.body.while.body_crit_edge.preheader
	movl	4(%rsp), %esi
	xorl	%ecx, %ecx
	.p2align	4, 0x90
.LBB0_5:                                # %while.body.while.body_crit_edge
                                        # =>This Inner Loop Header: Depth=1
	testb	$1, %sil
	sete	%dl
	andb	%dl, %cl
	addl	$-1, %eax
	cmpl	$1, %eax
	jg	.LBB0_5
# %bb.6:                                # %while.body.while.cond.while.end_crit_edge_crit_edge
	testb	%cl, %cl
	movl	$166, %ecx
	movl	$155, %esi
	cmovnel	%ecx, %esi
	addl	$-1, %eax
.LBB0_7:                                # %while.cond.while.end_crit_edge
	movl	%eax, (%rsp)
.LBB0_8:                                # %while.end
	movl	$.L.str.1, %edi
	xorl	%eax, %eax
	callq	printf
	xorl	%eax, %eax
	popq	%rcx
	retq
.Lfunc_end0:
	.size	main, .Lfunc_end0-main
	.cfi_endproc
                                        # -- End function
	.type	.L.str,@object          # @.str
	.section	.rodata.str1.1,"aMS",@progbits,1
.L.str:
	.asciz	"%d"
	.size	.L.str, 3

	.type	.L.str.1,@object        # @.str.1
.L.str.1:
	.asciz	"%d\n"
	.size	.L.str.1, 4


	.ident	"clang version 6.0.1 (tags/RELEASE_601/final)"
	.section	".note.GNU-stack","",@progbits
