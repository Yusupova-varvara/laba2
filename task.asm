section .text
    global _start

_start:
    call read_int
    mov ecx, eax

read_loop:
    push ecx
    call read_int
    pop ecx

    cmp eax, 10
    jl skip

    mov ebx, eax
    mov esi, 10
    xor edx, edx
    div esi
    push edx

check:
    cmp eax, 0
    je print
    xor edx, edx
    div esi
    pop edi
    cmp edx, edi
    jne skip
    push edi
    jmp check

print:
    pop edi
    mov eax, ebx
    call print_int
    mov al, ' '
    call print_char

skip:
    dec ecx
    jnz read_loop

    mov al, 10
    call print_char
    mov eax, 1
    xor ebx, ebx
    int 0x80

read_int:
    xor eax, eax
    xor ebx, ebx
.read:
    mov edx, 3
    mov ecx, buf
    mov ebx, 0
    mov eax, 3
    int 0x80
    mov al, [buf]
    cmp al, 10
    je .done
    sub al, '0'
    imul ebx, ebx, 10
    add ebx, eax
    jmp .read
.done:
    mov eax, ebx
    ret

print_int:
    push eax
    push ebx
    push ecx
    push edx

    mov ecx, 10
    mov ebx, outbuf
    add ebx, 9
    mov byte [ebx], 0

.conv:
    xor edx, edx
    div ecx
    add dl, '0'
    dec ebx
    mov [ebx], dl
    test eax, eax
    jnz .conv

    mov eax, 4
    mov ecx, ebx
    mov edx, outbuf
    add edx, 9
    sub edx, ebx
    mov ebx, 1
    int 0x80

    pop edx
    pop ecx
    pop ebx
    pop eax
    ret

print_char:
    pusha
    mov [char], al
    mov eax, 4
    mov ebx, 1
    mov ecx, char
    mov edx, 1
    int 0x80
    popa
    ret

section .bss
    buf resb 1
    outbuf resb 10
    char resb 1
