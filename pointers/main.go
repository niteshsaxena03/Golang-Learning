package main

import "fmt"

func modifyPointer(ptr *int) {
	*ptr = *ptr * 2
}

func main() {
	var num int = 10

	var ptr *int
	if ptr == nil {
		fmt.Println("Pointer is nil")
	}
	ptr = &num

	fmt.Println("Number is ", num)
	fmt.Println("Pointer is ", ptr)
	fmt.Println("Value at pointer is ", *ptr)

	modifyPointer(ptr)
	fmt.Println("Number is ", num)
}
