package main

import "fmt"

func main() {
	fmt.Println("Hello.This is a sample program")
	a := 1
	fmt.Println("value is", a)

	for i := 0; i < 3; i++ {
		fmt.Println("value is", i)
	}
}
