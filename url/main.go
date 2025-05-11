package main

import (
	"fmt"
	"net/url"
)
func main(){
	fmt.Println("Started url handling in go")

	var urlString="https://jsonplaceholder.typicode.com/todos/1"
	var url,err=url.Parse(urlString);
	if err != nil{
		fmt.Println("Error parsing url",err)
		return
	}
	fmt.Println(url.Scheme)
	fmt.Println(url.Host)
	fmt.Println(url.Path)
	fmt.Println(url.RawQuery)

	url.Path="/users";
	url.Host="nitesh";
	fmt.Println(url.String())
	
}