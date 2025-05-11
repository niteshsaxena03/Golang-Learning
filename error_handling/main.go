package main

import "fmt";

func divide(x float64,y float64)(float64,error){
	if(y==0){
		return 0,fmt.Errorf("divisor cant be 0")
	}
	return x/y,nil;
}

func main(){
	fmt.Println("Started error handling");

	var temp,_=divide(10,0);
	fmt.Println(temp)
}