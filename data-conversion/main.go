package main

import (
	"fmt"
	"strconv"
)

func main(){
	var num int=42;
	fmt.Printf("The type is %T\n",num);

	var num1 float64=float64(num);
	fmt.Printf("The type is %T\n",num1);

    var s1 string=strconv.Itoa(num);
	fmt.Printf("The type is %T\n",s1);
	
	var numberString string="123";
	var num2,_=strconv.Atoi(numberString);
	
	fmt.Println("The number is ",num2);
	
	var num3,_=strconv.ParseFloat(numberString,64);
	fmt.Println("The float number is ",num3);
}