package main

import "fmt"

func add(x int,y int)int{
	return x+y;
}
func multiply(x int,y int)int {
	return x*y;
}

func main(){
    fmt.Println("learning functions")	

	fmt.Println("The sum of 2 numbers is ",add(10,20))
	var temp=multiply(10,20);
	fmt.Println("The multiply is",temp)
}