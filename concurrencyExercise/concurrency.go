package main

import (
	"fmt"
	"sync"
)

var pl = fmt.Println

// func writeEven() {
// 	if (data+1)%2 == 1 {
// 		data++
// 	}
// }

// func writeOdd() {
// 	if (data+1)%2 == 0 {
// 		data++
// 	}
// }

//	func raceCondition() {
//		for i := 0; i < 20; i++ {
//			go writeEven()
//			pl(data)
//			go writeOdd()
//			pl(data)
//		}
//	}

func useChannel(nums int) {
	dataChannel := make(chan int, 1)
	done := make(chan bool)
	var wg sync.WaitGroup
	dataChannel <- 0
	for i := 0; i < nums; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := <-dataChannel
			if data%2 != 0 {
				data++
			}
			pl("Even ", data)
			dataChannel <- data
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := <-dataChannel
			if data%2 == 0 {
				data++
			}
			pl("Odd ", data)
			dataChannel <- data
		}()
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done
	pl("All goroutines completed")
}
func main() {
	//raceCondition()
	useChannel(10)
	//MutexFunction()
	//time.Sleep(1 * time.Second)
}
