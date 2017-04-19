
; nasm assembly
; Program to print Hello World

global _start

section .text
_start:
    ; call write syscall
    ; for list of syscalls,
    ; look at http://syscalls.kernelgrok.com/
    mov eax, 4
    mov ebx, 1
    mov ecx, string
    mov edx, length
    int 80h

    ; call exit syscall
    mov eax, 1
    mov ebx, 0
    int 80h

section .data
string: db 'Hello World', 0Ah
length: equ 5


