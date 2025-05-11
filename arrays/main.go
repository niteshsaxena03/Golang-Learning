package main

import "fmt"

func main() {
	fmt.Println("Started arrays")

	var array [5]string

	array[0] = "nitesh"
	array[1] = "saxena"

	fmt.Println("The array is", array)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("The array of numbers is ", numbers)

	fmt.Println("The length of the array is ", len(numbers))
}
