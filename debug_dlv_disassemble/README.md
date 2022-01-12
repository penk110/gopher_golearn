

#### [参考](https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-09-debug.html)

#### dlv disassemble
```text
➜  debug_dlv_disassemble git:(master) ✗ dlv exec ./dlv_disassemble 
Type 'help' for list of commands.
(dlv) 
(dlv) 
(dlv) b main.main
Breakpoint 1 set at 0x10a2e0f for main.main() ./dlv_disassemble.go:9
(dlv) c
> main.main() ./dlv_disassemble.go:9 (hits goroutine(1):1 total:1) (PC: 0x10a2e0f)
Warning: debugging optimized function
     4: 
     5: func disassembleFunc() {
     6:         fmt.Println("dlv disassemble")
     7: }
     8: 
=>   9: func main() {
    10:         disassembleFunc()
    11: }
(dlv) disassemble
TEXT main.main(SB) /Users/zhang/gopath/src/zyphub/gopher_golearn/debug_dlv_disassemble/dlv_disassemble.go
        dlv_disassemble.go:9    0x10a2e00       65488b0c2530000000      mov rcx, qword ptr gs:[0x30]
        dlv_disassemble.go:9    0x10a2e09       483b6110                cmp rsp, qword ptr [rcx+0x10]
        dlv_disassemble.go:9    0x10a2e0d       7671                    jbe 0x10a2e80
=>      dlv_disassemble.go:9    0x10a2e0f*      4883ec58                sub rsp, 0x58
        dlv_disassemble.go:9    0x10a2e13       48896c2450              mov qword ptr [rsp+0x50], rbp
        dlv_disassemble.go:9    0x10a2e18       488d6c2450              lea rbp, ptr [rsp+0x50]
        dlv_disassemble.go:10   0x10a2e1d       0f57c0                  xorps xmm0, xmm0
        dlv_disassemble.go:6    0x10a2e20       0f11442440              movups xmmword ptr [rsp+0x40], xmm0
        dlv_disassemble.go:6    0x10a2e25       488d05d4af0000          lea rax, ptr [rip+0xafd4]
        dlv_disassemble.go:6    0x10a2e2c       4889442440              mov qword ptr [rsp+0x40], rax
        dlv_disassemble.go:6    0x10a2e31       488d05b8360400          lea rax, ptr [rip+0x436b8]
        dlv_disassemble.go:6    0x10a2e38       4889442448              mov qword ptr [rsp+0x48], rax
        print.go:274            0x10a2e3d       488b05f4ba0b00          mov rax, qword ptr [os.Stdout]
        print.go:274            0x10a2e44       488d0d3d4d0400          lea rcx, ptr [rip+0x44d3d]
        print.go:274            0x10a2e4b       48890c24                mov qword ptr [rsp], rcx
        print.go:274            0x10a2e4f       4889442408              mov qword ptr [rsp+0x8], rax
        print.go:274            0x10a2e54       488d442440              lea rax, ptr [rsp+0x40]
        print.go:274            0x10a2e59       4889442410              mov qword ptr [rsp+0x10], rax
        print.go:274            0x10a2e5e       48c744241801000000      mov qword ptr [rsp+0x18], 0x1
        print.go:274            0x10a2e67       48c744242001000000      mov qword ptr [rsp+0x20], 0x1
        print.go:274            0x10a2e70       e88b9affff              call $fmt.Fprintln
        dlv_disassemble.go:6    0x10a2e75       488b6c2450              mov rbp, qword ptr [rsp+0x50]
        dlv_disassemble.go:6    0x10a2e7a       4883c458                add rsp, 0x58
        dlv_disassemble.go:6    0x10a2e7e       c3                      ret
        dlv_disassemble.go:9    0x10a2e7f       90                      nop
        dlv_disassemble.go:9    0x10a2e80       e8fbf5fbff              call $runtime.morestack_noctxt
        .:0                     0x10a2e85       e976ffffff              jmp $main.main
(dlv) 

```

#### 可对某个函数进行打断点，然后查看寄存器状态等，如果可以直接拿到某个变量的地址，可直接print *(类型断言、转换之类的)，然后把变量当前值打印出来
    
    参考这种方式可对某个方法、变量进行调试、查看等

```text
(dlv) break main.disassembleFunc
Breakpoint 2 set at 0x10a2e20 for main.main() ./dlv_disassemble.go:6
(dlv) c
> main.main() ./dlv_disassemble.go:6 (hits goroutine(1):1 total:1) (PC: 0x10a2e20)
Warning: debugging optimized function
     1: package main
     2: 
     3: import "fmt"
     4: 
     5: func disassembleFunc() {
=>   6:         fmt.Println("dlv disassemble")
     7: }
     8: 
     9: func main() {
    10:         disassembleFunc()
    11: }
(dlv) regs
   Rip = 0x00000000010a2e20
   Rsp = 0x000000c000092f28
   Rax = 0x00000000010a2e00
   Rbx = 0x0000000000000000
   Rcx = 0x000000c000000180
   Rdx = 0x00000000010cf0d0
   Rsi = 0x0000000000000001
   Rdi = 0x000000000114bdb8
   Rbp = 0x000000c000092f78
    R8 = 0x0000000000000001
    R9 = 0x0000000000000008
   R10 = 0x00000000010cee00
   R11 = 0x000000c000096210
   R12 = 0x0000000000000001
   R13 = 0x0000000000000001
   R14 = 0x0000000000000022
   R15 = 0xffffffffffffffff
Rflags = 0x0000000000000216     [PF AF IF IOPL=0]
    Cs = 0x000000000000002b
    Fs = 0x0000000000000000
    Gs = 0x0000000000000000

```