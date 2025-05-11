package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2 * time.Second)
	fmt.Println("Hello function executed")
}
func sayHi() {
	fmt.Println("Hi")
	time.Sleep(1 * time.Second)
	fmt.Println("Hi function executed")
}

func main() {
	fmt.Println("Learning Go Routines")
	go sayHello()
	go sayHi()

	time.Sleep(2 * time.Second)
}
