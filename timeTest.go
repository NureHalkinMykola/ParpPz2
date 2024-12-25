package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	threads := 1
	runtime.GOMAXPROCS(threads)
	size := 1000000000
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = 1
	}

	results := make([]int64, threads)
	chunkSize := size / threads
	startTime := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == threads-1 {
			end = size
		}
		go func(start int, end int, index int) {
			defer wg.Done()
			sum := int64(0)
			for j := start; j < end; j++ {
				sum += int64(data[j])
			}
			results[index] = sum
		}(start, end, i)
	}

	wg.Wait()
	var sum int64
	for _, result := range results {
		sum += result
	}
	endTime := time.Now()

	fmt.Printf("Threads: %d\n", threads)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Time: %v\n", endTime.Sub(startTime))
}
