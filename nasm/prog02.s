section .data
    hello   db  "Hello, World\n"
    msg1    db  "Message 1\n"
    msg2    db  "Message 2\n"

section .text
    global _start

_start:
    mov rax, 2
    cmp rax, 2
    jne not_two

    mov rax, 1
    mov rdi, 1
    mov rsi, msg1
    mov rdx, 10
    syscall
    jmp exit_ok

not_two:
    mov rax, 1
    mov rdi, 1
    mov rsi, msg2
    mov rdx, 10
    syscall

exit_not_ok:
    mov rax, 60
    mov rdi, 1
    syscall

exit_ok:
    mov rax, 60
    mov rdi, 0
    syscall
