package main

import (
	"fmt"
	"time"
)
func main(){
	currentTime:=time.Now();
	fmt.Println(currentTime);

	formattedTime:=currentTime.Format("2006-01-02 15:04:05");
	fmt.Println(formattedTime);
	
	var newFormattedTime=currentTime.Format("2006/01/02,Monday, 3:04 PM");
	fmt.Println(newFormattedTime);

	parsedTime,_:=time.Parse("2006-01-02 15:04:05",formattedTime);
	
	fmt.Println(parsedTime);
	
	//add 1 more day
	newTime:=currentTime.AddDate(0,0,1);
	fmt.Println(newTime);

		
}