package features

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"data1", "data2", "data3", "data4", "data5"}
var results []string

// DemonstrateGoroutines demonstrates the difference between synchronous and asynchronous
// execution of database calls using goroutines. It performs the following steps:
//
//  1. Executes database calls synchronously and measures the time taken.
//  2. Clears the results and prepares for asynchronous execution.
//  3. Executes database calls asynchronously using goroutines and a WaitGroup to wait for all
//     goroutines to complete, then measures the time taken.
//
// The function prints the results of the database calls and the time taken for both synchronous
// and asynchronous executions.
func DemonstrateGoroutines() {
	// Demonstrating goroutines
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i, false)
	}
	fmt.Println("The results of db calls are: ", results)
	fmt.Printf("Time taken for synchronous calls: %v\n", time.Since(t0))

	clear(results)
	results = []string{}

	t1 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i, true)
	}
	wg.Wait()
	fmt.Println("The results of db calls are: ", results)
	fmt.Printf("Time taken for asynchronous calls: %v\n", time.Since(t1))
}

// dbCall simulates a database call with a delay, saves the result, and logs the operation.
// If the async parameter is true, it signals the wait group that the goroutine is done.
//
// Parameters:
//   - i: The index of the data to be saved.
//   - async: A boolean indicating whether the call is asynchronous.
func dbCall(i int, async bool) {
	var delay float32 = 250
	time.Sleep(time.Duration(delay) * time.Millisecond)
	saveResult(dbData[i])
	log()
	if async {
		wg.Done()
	}
}

// saveResult appends a result string to the results slice in a thread-safe manner.
// It locks the mutex before appending to the slice and unlocks it afterward to ensure
// that concurrent access to the results slice does not cause data races.
//
// Parameters:
// - result: The result string to be appended to the results slice.
func saveResult(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

// log prints the current results of database calls to the standard output.
// It acquires a read lock before accessing the shared resource 'results' to ensure
// thread-safe read operations and releases the lock after printing.
func log() {
	m.RLock()
	fmt.Println("The current results of db calls are: ", results)
	m.RUnlock()
}
