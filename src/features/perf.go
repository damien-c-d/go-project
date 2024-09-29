package features

import (
	"fmt"
	"time"
)

// DemonstrateSlicePerformance measures and compares the performance of appending elements
// to a slice without preallocation versus with preallocation. It initializes two slices:
// one without preallocation and one with preallocation for a specified size, then prints
// the total time taken for each operation.
func DemonstrateSlicePerformance() {
	const size = 1_000_000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, size)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, size))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, size))
}

// timeLoop appends elements to a slice until it reaches the specified size.
// It returns the duration taken to complete this operation.
//
// Parameters:
//   - slice: The initial slice of integers.
//   - size: The target size of the slice.
//
// Returns:
//   - time.Duration: The duration taken to append elements to the slice until it reaches the target size.
func timeLoop(slice []int, size int) time.Duration {
	var start = time.Now()
	for len(slice) < size {
		slice = append(slice, 1)
	}
	return time.Since(start)
}
