package routines

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func Main() {
	fmt.Printf("\n------------------\nRoutine.go\n------------------\n")
	s := "Hello"

	wg.Add(3)

	go say(s)
	go func(s string) {
		fmt.Printf("Saying %s from an anonymous function!\n", s)
		wg.Done()
	}(s)
	go func() {
		fmt.Printf("Saying %s from an anonymous function that is tightly coupled to the main routine!\n", s)
		wg.Done()
	}()
	s = "Bye"

	wg.Wait()
	fmt.Printf("\n------------------\n")
}

func say(s string) {
	fmt.Printf("Saying %s from a function!\n", s)
	wg.Done()
}

var counter int = 0
var m = sync.RWMutex{}

func SynchedMain() {
	fmt.Printf("\n------------------\nSynchedMain\n------------------\n")

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go increment()
		go say(fmt.Sprintf("%v", counter))
	}

	wg.Wait()
	fmt.Printf("\n------------------\n")
}

func increment() {
	counter++
	wg.Done()
}

func ChanneledMain() {
	fmt.Printf("\n------------------\nChanneledMain\n------------------\n")
	ch := make(chan int) // Strongly typed messages (in this case int)

	for x := 0; x < 5; x++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Printf("Received %v from the channel!\n", i)
			wg.Done()
		}()

		go func() {
			i := 42
			fmt.Printf("Sending %v to the channel!\n", i)
			ch <- i
			wg.Done()
		}()
	}
	wg.Wait()

	// We can also set the direction we want data to flow in
	fmt.Printf("One way channel\n")
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Printf("Received %v from the channel!\n", i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		i := 42
		fmt.Printf("Sending %v to the channel!\n", i)
		ch <- i
		wg.Done()
	}(ch)
	wg.Wait()
}
