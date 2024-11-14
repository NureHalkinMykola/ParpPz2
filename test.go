package main

import (
	"runtime"
	"runtime/pprof"
	"sync"
)

var wait sync.WaitGroup
var threads = pprof.Lookup("threadcreate")

func test() {
	runtime.LockOSThread()
	defer wait.Done()
	var k int

	for i := 0; i < 2e9; i++ {
		k = k + i
	}

	runtime.UnlockOSThread()
}

func main() {
	runtime.GOMAXPROCS(1)
	count := 5

	wait.Add(count)
	println("Start threads -", threads.Count())

	for i := 0; i < count; i++ {
		go test()
	}

	wait.Wait()
	println("End threads -", threads.Count())
}
