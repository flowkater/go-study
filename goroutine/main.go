package main

import (
	"time"
)

func multiply(inChan, outChan chan int) {
	defer close(outChan)
	for in := range inChan {
		outChan <- in * 2
	}
}

func input(inChan chan int) {
	defer close(inChan)
	for i := 0; i < 5; i++ {
		inChan <- i
	}
}

func main() {
	done := make(chan struct{})

	go func() {
		time.Sleep(10 * time.Second)
		done <- struct{}{}
	}()
	<-done

	// inChan := make(chan int, 5)
	// outChan := make(chan int)

	// go multiply(inChan, outChan)
	// go input(inChan)

	// for {
	// 	out, more := <-outChan
	// 	if !more {
	// 		return
	// 	}
	// 	log.Println(out)
	// }

	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		log.Println(i)
	// 	}(i)
	// }

	// time.Sleep(time.Second)
}
