package main

import "fmt"

func main(){
	//defer statement maintains a stack of functions to be called after the function returns
	fmt.Println("Starting the defer keyword");
	fmt.Println("This is the first statement");
	defer fmt.Println("This is the second statement");
	defer fmt.Println("Another defer statement")
	fmt.Println("This is the third statement");
}