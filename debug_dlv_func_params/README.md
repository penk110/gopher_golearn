
#### 测试代码
```go
package main

import "strconv"

func task(id int) {
	ids := strconv.Itoa(id)
	println("task_" + ids + " is running ...")
}

func add(x, y int) (int, bool) {
	var z = x+y
	println(z)
	return z, true
}

func main() {
	add(1, 2)
	task(10)
}
```

#### set 修改变量值
```shell
➜  debug_dlv_func_params git:(master) ✗ dlv debug params.go
Type 'help' for list of commands.
(dlv) b main.main
Breakpoint 1 (enabled) set at 0x106cf8f for main.main() ./params.go:16
(dlv) c
> main.main() ./params.go:16 (hits goroutine(1):1 total:1) (PC: 0x106cf8f)
    11:         var z = x+y
    12:         println(z)
    13:         return z, true
    14: }
    15: 
=>  16: func main() {
    17:         add(1, 2)
    18:         task(10)
    19: }
(dlv) s
> main.main() ./params.go:17 (PC: 0x106cf9d)
    12:         println(z)
    13:         return z, true
    14: }
    15: 
    16: func main() {
=>  17:         add(1, 2)
    18:         task(10)
    19: }
(dlv) args
(no args)
(dlv) s
> main.add() ./params.go:10 (PC: 0x106cf0f)
     5: func task(id int) {
     6:         ids := strconv.Itoa(id)
     7:         println("task_" + ids + " is running ...")
     8: }
     9: 
=>  10: func add(x, y int) (int, bool) {
    11:         var z = x+y
    12:         println(z)
    13:         return z, true
    14: }
    15: 
(dlv) args
x = 1
y = 2
~r2 = 824634163200
~r3 = true
(dlv) set x=10000
(dlv) args
x = 10000
y = 2
~r2 = 824634163200
~r3 = true
(dlv) s
> main.add() ./params.go:11 (PC: 0x106cf2b)
     6:         ids := strconv.Itoa(id)
     7:         println("task_" + ids + " is running ...")
     8: }
     9: 
    10: func add(x, y int) (int, bool) {
=>  11:         var z = x+y
    12:         println(z)
    13:         return z, true
    14: }
    15: 
    16: func main() {
(dlv) s
> main.add() ./params.go:12 (PC: 0x106cf3a)
     7:         println("task_" + ids + " is running ...")
     8: }
     9: 
    10: func add(x, y int) (int, bool) {
    11:         var z = x+y
=>  12:         println(z)
    13:         return z, true
    14: }
    15: 
    16: func main() {
    17:         add(1, 2)
(dlv) s
10002
> main.add() ./params.go:13 (PC: 0x106cf57)
     8: }
     9: 
    10: func add(x, y int) (int, bool) {
    11:         var z = x+y
    12:         println(z)
=>  13:         return z, true
    14: }
    15: 
    16: func main() {
    17:         add(1, 2)
    18:         task(10)
(dlv) 

```