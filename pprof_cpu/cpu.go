package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

func generator(n int) string {
	var buf bytes.Buffer
	for i := 0; i < 100000; i++ {
		buf.WriteString(strconv.Itoa(n))
	}
	sum := sha256.Sum256(buf.Bytes())

	var b []byte
	for i := 0; i < int(sum[0]); i++ {
		x := sum[(i*7+1)%len(sum)] ^ sum[(i*5+3)%len(sum)]
		c := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 10)[x]
		b = append(b, c)
	}
	return string(b)
}

func main() {
	var (
		cpuFile *os.File
		err     error
	)

	if cpuFile, err = os.Create("cpu.pprof"); err != nil {
		panic(err)
	}

	if err = pprof.StartCPUProfile(cpuFile); err != nil {
		panic(err)
	}
	defer func(cpuFile *os.File) {
		if err = cpuFile.Close(); err != nil {
			log.Printf("close cup file failed, err: %v\n", err.Error())
		}
	}(cpuFile)
	defer pprof.StopCPUProfile()

	// ensure function output is accurate
	if generator(12345) == "aajmtxaattdzsxnukawxwhmfotnm" {
		fmt.Println("Test PASS")
	} else {
		fmt.Println("Test FAIL")
	}

	for i := 0; i < 100; i++ {
		generator(rand.Int())
	}
}

/*

go tool pprof cpu.pprof


(pprof) top
Showing nodes accounting for 1040ms, 89.66% of 1160ms total
Showing top 10 nodes out of 65
      flat  flat%   sum%        cum   cum%
     480ms 41.38% 41.38%      480ms 41.38%  crypto/sha256.block
     190ms 16.38% 57.76%      300ms 25.86%  strconv.formatBits
      80ms  6.90% 64.66%       80ms  6.90%  runtime.kevent
      60ms  5.17% 69.83%      170ms 14.66%  bytes.(*Buffer).WriteString
      50ms  4.31% 74.14%      160ms 13.79%  runtime.mallocgc
      50ms  4.31% 78.45%       50ms  4.31%  runtime.memclrNoHeapPointers
      40ms  3.45% 81.90%       40ms  3.45%  runtime.memmove
      40ms  3.45% 85.34%       40ms  3.45%  runtime.pthread_cond_wait
      30ms  2.59% 87.93%      990ms 85.34%  main.generator
      20ms  1.72% 89.66%       20ms  1.72%  bytes.(*Buffer).tryGrowByReslice (inline)

(pprof) pdf
Generating report in profile001.pdf


*/
