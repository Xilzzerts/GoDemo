// er
package main

import (
	"fmt"
)

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, item := range values {
		sum += item
	}
	resultChan <- sum
}

func main() {
	fmt.Println("Hello World!")
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 2)
	length := len(values)
	go sum(values[:length/2], resultChan)
	go sum(values[length/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan
	fmt.Println("Res: ", sum1, sum2, sum1+sum2)
}
