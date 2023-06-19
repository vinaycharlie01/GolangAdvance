// package main

// import "fmt"

// func divide(x, y float64) (float64, error) {
// 	if y == 0 {
// 		return 0, fmt.Errorf("division by zero")
// 	}
// 	return x / y, nil
// }

// func main() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from panic:", r)
// 		}
// 		fmt.Println("Finally  completed")
// 		// return ""
// 	}()

// 	result, err := divide(10, 0.0)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Result:", result)
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func Work(a string, ch chan string, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("%v index %v", a, i)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go Work("This is my 1st work", ch1, &wg)
	go Work("This is my 2st work", ch2, &wg)

	for {
		select {

		case val, ok := <-ch1:
			if !ok {
				ch1 = nil
			}
			fmt.Println(val)
		case val, ok := <-ch2:
			if !ok {
				ch2 = nil
			}
			fmt.Println(val)

		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}

	go func() {
		wg.Wait()
	}()

}
