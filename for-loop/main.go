package main

import "fmt"

func main() {
	var i int;
	for i=0;i<10;i++{
		fmt.Println(i)
	}

	var counter=0;
	for{
		fmt.Println("Infinite loop")
		counter++
	}
	
	var numbers=[]int{1,2,3,4,5};
	for index,value:=range numbers{
		fmt.Println(index,value)
	}
	
}
