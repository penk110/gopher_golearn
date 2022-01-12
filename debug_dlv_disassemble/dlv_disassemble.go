package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func disassembleFunc() {

	var ss []int
	_cap := rand.Intn(10)
	for i := 0; i <= _cap; i++ {
		ss = append(ss, i)
	}
	fmt.Println("dlv disassemble")
}
func Parallel() {
	signal := make(chan struct{})
	parallelCount := 9
	wg := &sync.WaitGroup{}

	for i := 1; i <= parallelCount; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, single chan struct{}, i int) {
			<-single
			fmt.Printf("[Paralleler] index: %v  time: %v\n", i, time.Now().Format("2006-01-02 15:04:05.0000"))
			wg.Done()
		}(wg, signal, i)
	}
	fmt.Println("waiting signal")
	fmt.Println("sleep 2s ...")
	go func() {
		i := 9
		for ; i > 0; {
			i--
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println(" trigger signal ...")
	close(signal)
	wg.Wait()
}

func main() {
	disassembleFunc()

	Parallel()
}
