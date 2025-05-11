package main

import "fmt"

func main() {
	fmt.Println("Learning slices")

	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println("The slice is", slice)

	slice = append(slice, 2, 3, 4, 5)
	fmt.Println("The new slice is", slice)

	var slice2 = []string{}
	slice2 = append(slice2, "nitesh", "saxena")
	fmt.Println(slice2)
	fmt.Println("The length of the slice is", len(slice2))
}
