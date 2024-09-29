package features

import (
	"fmt"
	"math/rand"
	"time"
)

// DemonstrateChannels showcases various channel operations in Go.
// It includes demonstrations of:
// - Simple channel usage
// - Looping with channels (both non-buffered and buffered)
// - Buffered channels
// - Advanced channel techniques
func DemonstrateChannels() {
	demonstrateSimpleChannel()
	demonstrateLoopChannel(false)
	demonstrateLoopChannel(true)
	demonstrateBufferedChannel()
	demonstrateAdvancedChannel()
}

// demonstrateSimpleChannel demonstrates the usage of a simple channel in Go.
// It creates an unbuffered channel of type int, starts a goroutine that sends
// a value into the channel, and then receives and prints that value from the channel.
func demonstrateSimpleChannel() {
	// Simple channel
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println("The value received from the simple channel is: ", <-ch)
}

// demonstrateLoopChannel demonstrates the use of a loop to receive values from a channel.
// It starts a goroutine that sends values to the channel, either using the process or processDefer function,
// based on the callProcessDefer parameter.
//
// Parameters:
// - callProcessDefer: A boolean flag that determines whether to call the processDefer function (true) or the process function (false).
//
// The function prints the values received from the channel along with an indication of whether the processDefer function was used.
func demonstrateLoopChannel(callProcessDefer bool) {
	// Goroutine with channel
	c := make(chan int)
	if callProcessDefer {
		go processDefer(c)
	} else {
		go process(c)
	}
	// Loop to receive values from the channel
	for i := range c {
		var loopType string = "without defer"
		if callProcessDefer {
			loopType = "with defer"
		}
		fmt.Printf("\nThe value received from the loop channel %v is: %v", loopType, i)
	}
}

// process sends a series of integers to the provided channel and then closes the channel.
// It sends the integers 0 through 4 to the channel in a loop.
//
// Parameters:
//
//	c - A channel of integers to which the function will send values.
func process(c chan int) {
	// Loop to send values to the channel
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c) // Closing the channel to avoid deadlock
}

// processDefer sends a sequence of integers to the provided channel and then closes the channel.
// The function uses a defer statement to ensure the channel is closed right before the function returns,
// preventing potential deadlocks.
//
// Parameters:
//
//	c - A channel of integers to which the function will send values.
func processDefer(c chan int) {
	defer close(c) // Closing the channel to avoid deadlock, defer is called right before the function returns
	// Loop to send values to the channel
	for i := 0; i < 5; i++ {
		c <- i
	}
}

// demonstrateBufferedChannel demonstrates the usage of a buffered channel in Go.
// It creates a buffered channel with a capacity of 5 and starts a goroutine to process the buffer.
// The function then enters a loop to receive values from the channel and prints them with a delay of 250 milliseconds.
func demonstrateBufferedChannel() {
	var c = make(chan int, 5)
	go processBuffer(c)
	// Loop to receive values from the channel
	for i := range c {
		fmt.Printf("\nThe value received from the buffered channel is: %v", i)
		time.Sleep(time.Millisecond * 250)
	}
}

// processBuffer sends a series of integers to the provided channel and then closes the channel.
// It sends the integers 0 through 4 to the channel in a loop.
// After sending all values, it prints a message indicating the goroutine is exiting.
//
// Parameters:
//
//	c - A channel of integers to which the function will send values.
func processBuffer(c chan int) {
	defer close(c)
	// Loop to send values to the channel
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("\nExiting the processBuffer goroutine")
}

const maxChickenPrice float32 = 5.0
const maxBeefPrice float32 = 10.0

// demonstrateAdvancedChannel concurrently checks chicken and beef prices from a list of websites.
// It creates two channels, chickenChannel and beefChannel, to receive the prices.
// For each website in the list, it starts two goroutines: one to check chicken prices and another to check beef prices.
// Finally, it sends the collected messages from both channels to the sendMessage function.
func demonstrateAdvancedChannel() {
	chickenChannel := make(chan string)
	beefChannel := make(chan string)
	websites := []string{"woolworths.com.au", "coles.com.au", "aldi.com.au", "iga.com.au"}
	for i := range websites {
		go checkChickenPrice(websites[i], chickenChannel)
		go checkBeefPrice(websites[i], beefChannel)
	}
	sendMessage(chickenChannel, beefChannel)
}

// checkChickenPrice continuously checks the price of chicken from a given website
// and sends a message to the provided channel if the price is less than or equal
// to the specified maximum price.
//
// Parameters:
//   - website: A string representing the website to check the chicken price from.
//   - ch: A channel of strings to send the price message to.
//
// The function runs indefinitely, checking the price every 500 milliseconds until
// the price is less than or equal to the maximum price, at which point it sends
// a formatted message to the channel and exits.
func checkChickenPrice(website string, ch chan string) {
	for {
		time.Sleep(time.Millisecond * 500)
		chickenPrice := rand.Float32() * 30
		if chickenPrice <= maxChickenPrice {
			ch <- fmt.Sprintf("The price of chicken is $%.2f at %s", chickenPrice, website)
			break
		}
	}
}

// checkBeefPrice continuously checks the price of beef on a given website
// at intervals of 500 milliseconds. If the price is less than or equal to
// maxBeefPrice, it sends a formatted string with the price and website
// to the provided channel and then exits.
//
// Parameters:
//   - website: The URL of the website to check the beef price.
//   - ch: A channel to send the formatted price string.
//
// Note: This function runs an infinite loop until the condition is met,
// so it should be run as a goroutine to avoid blocking.
func checkBeefPrice(website string, ch chan string) {
	for {
		time.Sleep(time.Millisecond * 500)
		beefPrice := rand.Float32() * 30
		if beefPrice <= maxBeefPrice {
			ch <- fmt.Sprintf("The price of beef is $%.2f at %s", beefPrice, website)
			break
		}
	}
}

// sendMessage listens on two channels, chickenChannel and beefChannel, and prints a message
// based on which channel receives a value first. If chickenChannel receives a value, it prints
// a message about a deal on chicken. If beefChannel receives a value, it prints a message about
// a deal on beef.
//
// Parameters:
//   - chickenChannel: a channel of strings that sends messages about chicken deals.
//   - beefChannel: a channel of strings that sends messages about beef deals.
func sendMessage(chickenChannel chan string, beefChannel chan string) {
	select {
	case chicken := <-chickenChannel:
		fmt.Printf("\nText Sent: Found a deal on chicken! %s\n", chicken)
	case beef := <-beefChannel:
		fmt.Printf("\nEmail Sent: Found a deal on beef! %s\n", beef)
	}
}
