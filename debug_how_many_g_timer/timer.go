package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

func main() {
	debug.SetTraceback("system")
	// TODO: 用于确认没有启动timer前有几个goroutine，并且可查看对应goroutine是用于处理什么的
	if len(os.Args) == 1 {
		panic("before timers")
	}
	for i := 0; i < 10; i++ {
		time.AfterFunc(time.Duration(5*time.Second), func() {
			fmt.Println("Hello, Gopher!")
		})
	}
	// TODO: 启用timer之后
	panic("after timers")
}

/*
https://blog.gopheracademy.com/advent-2016/go-timers/

日志复现:
	无论有多少timer，都只有5个goroutine

备注：
	当 timer 足够多时，出现多 goroutine 51 [GC worker (idle)]，相关源码部分：/Users/zhang/sdk/go1.16/src/runtime/mgc.go:1835 +0x37



go run timer.go 2>&1 | grep "^goroutine" | wc -l

and after spawning 10k timers:
go run timer.go after 2>&1 | grep "^goroutine" | wc -l


➜  debug_how_many_g_timer git:(master) ✗ go run timer.go
panic: before timers

goroutine 1 [running]:
panic(0x10ae4c0, 0x10e6eb0)
        /Users/zhang/sdk/go1.16/src/runtime/panic.go:1065 +0x565 fp=0xc000072f58 sp=0xc000072e90 pc=0x1032065
main.main()
        /Users/zhang/gopath/src/zyphub/gopher_golearn/debug_how_many_g_timer/timer.go:14 +0xb4 fp=0xc000072f88 sp=0xc000072f58 pc=0x10a33d4
runtime.main()
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:225 +0x256 fp=0xc000072fe0 sp=0xc000072f88 pc=0x1034c96
runtime.goexit()
        /Users/zhang/sdk/go1.16/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc000072fe8 sp=0xc000072fe0 pc=0x1064341

goroutine 2 [force gc (idle)]:
runtime.gopark(0x10cfae0, 0x115ff20, 0x1411, 0x1)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:336 +0xe5 fp=0xc000048fb0 sp=0xc000048f90 pc=0x10350c5
runtime.goparkunlock(...)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:342
runtime.forcegchelper()
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:276 +0xc5 fp=0xc000048fe0 sp=0xc000048fb0 pc=0x1034f25
runtime.goexit()
        /Users/zhang/sdk/go1.16/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc000048fe8 sp=0xc000048fe0 pc=0x1064341
created by runtime.init.6
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:264 +0x35

goroutine 3 [GC sweep wait]:
runtime.gopark(0x10cfae0, 0x1160040, 0x140c, 0x1)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:336 +0xe5 fp=0xc0000497a8 sp=0xc000049788 pc=0x10350c5
runtime.goparkunlock(...)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:342
runtime.bgsweep(0xc00005e000)
        /Users/zhang/sdk/go1.16/src/runtime/mgcsweep.go:163 +0x9e fp=0xc0000497d8 sp=0xc0000497a8 pc=0x102173e
runtime.goexit()
        /Users/zhang/sdk/go1.16/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc0000497e0 sp=0xc0000497d8 pc=0x1064341
created by runtime.gcenable
        /Users/zhang/sdk/go1.16/src/runtime/mgc.go:217 +0x5c

goroutine 4 [GC scavenge wait]:
runtime.gopark(0x10cfae0, 0x11600e0, 0x140d, 0x1)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:336 +0xe5 fp=0xc000049f78 sp=0xc000049f58 pc=0x10350c5
runtime.goparkunlock(...)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:342
runtime.bgscavenge(0xc00005e000)
        /Users/zhang/sdk/go1.16/src/runtime/mgcscavenge.go:265 +0xd2 fp=0xc000049fd8 sp=0xc000049f78 pc=0x101f792
runtime.goexit()
        /Users/zhang/sdk/go1.16/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc000049fe0 sp=0xc000049fd8 pc=0x1064341
created by runtime.gcenable
        /Users/zhang/sdk/go1.16/src/runtime/mgc.go:218 +0x7e

goroutine 5 [finalizer wait]:
runtime.gopark(0x10cfae0, 0x118d338, 0x10a1410, 0x1)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:336 +0xe5 fp=0xc000048758 sp=0xc000048738 pc=0x10350c5
runtime.goparkunlock(...)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:342
runtime.runfinq()
        /Users/zhang/sdk/go1.16/src/runtime/mfinal.go:175 +0xa9 fp=0xc0000487e0 sp=0xc000048758 pc=0x1016a89
runtime.goexit()
        /Users/zhang/sdk/go1.16/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc0000487e8 sp=0xc0000487e0 pc=0x1064341
created by runtime.createfing
        /Users/zhang/sdk/go1.16/src/runtime/mfinal.go:156 +0x65
exit status 2
➜  debug_how_many_g_timer git:(master) ✗ go run timer.go after
panic: after timers

goroutine 1 [running]:
panic(0x10ae4c0, 0x10e6ec0)
        /Users/zhang/sdk/go1.16/src/runtime/panic.go:1065 +0x565 fp=0xc000072f58 sp=0xc000072e90 pc=0x1032065
main.main()
        /Users/zhang/gopath/src/zyphub/gopher_golearn/debug_how_many_g_timer/timer.go:22 +0x98 fp=0xc000072f88 sp=0xc000072f58 pc=0x10a33b8
runtime.main()
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:225 +0x256 fp=0xc000072fe0 sp=0xc000072f88 pc=0x1034c96
runtime.goexit()
        /Users/zhang/sdk/go1.16/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc000072fe8 sp=0xc000072fe0 pc=0x1064341

goroutine 2 [force gc (idle)]:
runtime.gopark(0x10cfae0, 0x115ff20, 0x1411, 0x1)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:336 +0xe5 fp=0xc000048fb0 sp=0xc000048f90 pc=0x10350c5
runtime.goparkunlock(...)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:342
runtime.forcegchelper()
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:276 +0xc5 fp=0xc000048fe0 sp=0xc000048fb0 pc=0x1034f25
runtime.goexit()
        /Users/zhang/sdk/go1.16/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc000048fe8 sp=0xc000048fe0 pc=0x1064341
created by runtime.init.6
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:264 +0x35

goroutine 3 [GC sweep wait]:
runtime.gopark(0x10cfae0, 0x1160040, 0x140c, 0x1)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:336 +0xe5 fp=0xc0000497a8 sp=0xc000049788 pc=0x10350c5
runtime.goparkunlock(...)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:342
runtime.bgsweep(0xc000060000)
        /Users/zhang/sdk/go1.16/src/runtime/mgcsweep.go:163 +0x9e fp=0xc0000497d8 sp=0xc0000497a8 pc=0x102173e
runtime.goexit()
        /Users/zhang/sdk/go1.16/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc0000497e0 sp=0xc0000497d8 pc=0x1064341
created by runtime.gcenable
        /Users/zhang/sdk/go1.16/src/runtime/mgc.go:217 +0x5c

goroutine 4 [GC scavenge wait]:
runtime.gopark(0x10cfae0, 0x11600e0, 0x140d, 0x1)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:336 +0xe5 fp=0xc000049f78 sp=0xc000049f58 pc=0x10350c5
runtime.goparkunlock(...)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:342
runtime.bgscavenge(0xc000060000)
        /Users/zhang/sdk/go1.16/src/runtime/mgcscavenge.go:265 +0xd2 fp=0xc000049fd8 sp=0xc000049f78 pc=0x101f792
runtime.goexit()
        /Users/zhang/sdk/go1.16/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc000049fe0 sp=0xc000049fd8 pc=0x1064341
created by runtime.gcenable
        /Users/zhang/sdk/go1.16/src/runtime/mgc.go:218 +0x7e

goroutine 5 [finalizer wait]:
runtime.gopark(0x10cfae0, 0x118d338, 0x10a1410, 0x1)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:336 +0xe5 fp=0xc000048758 sp=0xc000048738 pc=0x10350c5
runtime.goparkunlock(...)
        /Users/zhang/sdk/go1.16/src/runtime/proc.go:342
runtime.runfinq()
        /Users/zhang/sdk/go1.16/src/runtime/mfinal.go:175 +0xa9 fp=0xc0000487e0 sp=0xc000048758 pc=0x1016a89
runtime.goexit()
        /Users/zhang/sdk/go1.16/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc0000487e8 sp=0xc0000487e0 pc=0x1064341
created by runtime.createfing
        /Users/zhang/sdk/go1.16/src/runtime/mfinal.go:156 +0x65
exit status 2

*/
