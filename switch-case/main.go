package main

import "fmt";

func main(){
	var day=3;

	switch day{
	case 1:
		fmt.Println("Monday");
	case 2:
		fmt.Println("Tuesday");
	case 3:
		fmt.Println("Wednesday");
	default:
		fmt.Println("Invalid day");
	}

	var temperature=25;

	switch {
	case temperature<0:
		fmt.Println("It is too cold");
	case temperature>0 && temperature<15:
		fmt.Println("It is cool");
	case temperature>15 && temperature<25:
		fmt.Println("It is warm");
	default:
		fmt.Println("It is hot");
	}
}