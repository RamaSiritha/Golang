package main

import (
	"fmt"
)

func generateNumbers(numbers chan<- int) {
	for i := 2; i < 100; i++ {
		numbers <- i
	}
}
func generatePrimeNumbers(in <-chan int, out chan<- int, prime int) {
	for {
		num := <-in
		if num%prime != 0 {
			out <- num
		}
	}
}

func main() {
	numbers := make(chan int)
	go generateNumbers(numbers)
	for i := 0; i < 25; i++ {
		prime := <-numbers
		fmt.Println(prime)
		newNumbers := make(chan int)
		go generatePrimeNumbers(numbers, newNumbers, prime)
		numbers = newNumbers
	}
}
