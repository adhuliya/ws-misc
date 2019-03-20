	.text
	.file	"optimizeTest1.c"
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %entry
	pushq	%rax
	.cfi_def_cfa_offset 16
	leaq	4(%rsp), %rsi
	movl	$.L.str, %edi
	xorl	%eax, %eax
	callq	__isoc99_scanf
	movl	4(%rsp), %ecx
	testl	%ecx, %ecx
	je	.LBB0_1
# %bb.2:                                # %while.body.lr.ph
	movl	$155, (%rsp)
	movl	%ecx, %eax
	addl	$-1, %eax
	movl	%eax, 4(%rsp)
	je	.LBB0_3
# %bb.4:                                # %while.body.while.body_crit_edge.preheader
                                        # implicit-def: %esi
	testb	$1, %al
	je	.LBB0_6
# %bb.5:                                # %while.body.while.body_crit_edge.prol
	cmpl	$0, (%rsp)
	movl	$166, %eax
	movl	$155, %esi
	cmovel	%eax, %esi
	movl	%esi, (%rsp)
	leal	-2(%rcx), %eax
	movl	%eax, 4(%rsp)
.LBB0_6:                                # %while.body.while.body_crit_edge.prol.loopexit
	cmpl	$2, %ecx
	je	.LBB0_13
# %bb.7:                                # %while.body.while.body_crit_edge.preheader.new
	addl	$-1, %eax
	movq	%rsp, %r8
	movq	%r8, %rdx
	.p2align	4, 0x90
.LBB0_8:                                # %while.body.while.body_crit_edge
                                        # =>This Inner Loop Header: Depth=1
	movl	(%rsp), %edi
	movl	$166, %esi
	movl	$166, %ecx
	testl	%edi, %edi
	je	.LBB0_10
# %bb.9:                                # %while.body.while.body_crit_edge
                                        #   in Loop: Header=BB0_8 Depth=1
	movl	$155, %ecx
.LBB0_10:                               # %while.body.while.body_crit_edge
                                        #   in Loop: Header=BB0_8 Depth=1
	testl	%ecx, %ecx
	je	.LBB0_12
# %bb.11:                               # %while.body.while.body_crit_edge
                                        #   in Loop: Header=BB0_8 Depth=1
	movl	$155, %esi
.LBB0_12:                               # %while.body.while.body_crit_edge
                                        #   in Loop: Header=BB0_8 Depth=1
	orl	%edi, %ecx
	cmovneq	%r8, %rdx
	movl	%esi, (%rsp)
	leal	-1(%rax), %ecx
	movl	%ecx, 4(%rsp)
	addl	$-2, %eax
	cmpl	$-1, %eax
	jne	.LBB0_8
	jmp	.LBB0_13
.LBB0_1:
                                        # implicit-def: %esi
	jmp	.LBB0_13
.LBB0_3:
	movl	$155, %esi
.LBB0_13:                               # %while.end
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
