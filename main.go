package main

import (
	"fmt"
)

func calculateSum(numbers []int, result chan<- int) {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	result <- sum
}
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make(chan int)
	divide := len(numbers) / 2
	go calculateSum(numbers[:divide], result)
	go calculateSum(numbers[divide:], result)
	firstSum := <-result
	secondSum := <-result
	totalSum := firstSum + secondSum
	fmt.Println("Total Sum:", totalSum)
}
