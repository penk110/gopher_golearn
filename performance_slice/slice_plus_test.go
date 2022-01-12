package performance_slice

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"testing"
)

// https://github.com/askuy/gopherlearn/blob/master/2018wuhan_meetup/1%E7%BC%93%E5%AD%98%E6%95%B0%E6%8D%AE%E4%BC%98%E5%8C%96/14GO%E5%A4%9A%E6%95%B0%E6%8D%AE%E4%BC%98%E5%8C%96/slice_test.go

// https://golang.org/pkg/testing/

//
/*
go test -bench=BenchmarkDefaultSlice
go test -bench=.

goos: darwin
goarch: amd64
pkg: zyphub/gopher_golearn/performance_slice
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
								 循环次数			  每次循环需要的时间		每次分配的内存Byte 产生了多少次内存分配操作
BenchmarkDefaultSlice
BenchmarkDefaultSlice-12          326876              7373 ns/op            3343 B/op        200 allocs/op
BenchmarkPreAllocSlice
BenchmarkPreAllocSlice-12        1000000              1070 ns/op             782 B/op        100 allocs/op
BenchmarkSyncPoolSlice
BenchmarkSyncPoolSlice-12        1000000              1493 ns/op            1291 B/op        100 allocs/op
PASS
ok      zyphub/gopher_golearn/performance_slice 5.953s

*/

type Item struct {
	roomId  int
	roomTag string
}

func TestMain(m *testing.M) {
	fmt.Println("关于切片的压测例子")

	// 必须显示调用 m.Run()
	os.Exit(m.Run())
}

func BenchmarkDefaultSlice(b *testing.B) {
	b.ReportAllocs()
	//b.ReportMetric(10, "metric")

	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			for i := 0; i < 100; i++ {
				// 模拟频繁创建
				output := make([]Item, 0)
				output = append(output, Item{
					roomId:  i,
					roomTag: "tag_" + strconv.Itoa(i),
				})
			}
			wg.Done()
		}(wg)
	}

	wg.Wait()

}

func BenchmarkPreAllocSlice(b *testing.B) {
	b.ReportAllocs()

	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			output := make([]Item, 0, 100)
			for i := 0; i < 100; i++ {
				output = append(output, Item{
					roomId:  i,
					roomTag: "tag_" + strconv.Itoa(i),
				})
			}
			wg.Done()
		}(wg)
	}

	wg.Wait()

}

func BenchmarkSyncPoolSlice(b *testing.B) {
	b.ReportAllocs()
	wg := &sync.WaitGroup{}

	// 详情看 sync.Pool
	var SomethingPool = sync.Pool{
		New: func() interface{} {
			b := make([]Item, 100)
			return &b
		},
	}
	for i := 0; i < b.N; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			obj := SomethingPool.Get().(*[]Item)
			for i := 0; i < 100; i++ {
				some := *obj
				some[i].roomId = i
				some[i].roomTag = "tag_" + strconv.Itoa(i)
			}
			SomethingPool.Put(obj)
			wg.Done()
		}(wg)
	}
	wg.Wait()
}
