package main

import "fmt";

func main(){
	var x=5;
	var y=10;

	if(x>5){
		fmt.Println("x is greater than 5");
	}else if(x<5){
		fmt.Println("x is less than 5");
	}else{
		fmt.Println("x is equal to 5");
	}

	if(x==5 && y==10){
		fmt.Println("Condition is true");
	}
}