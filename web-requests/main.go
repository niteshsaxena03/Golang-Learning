package main

import (
	"fmt"
	"io"
	"net/http"
)
func main(){
	fmt.Println("Started web requests in go")

	var response,err=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil{
		fmt.Println("Error making request",err)
		return
	}
	defer response.Body.Close();
	
	var data,error=io.ReadAll(response.Body);
	if error != nil{
		fmt.Println("Error reading response",error)
		return
	}
	fmt.Println(string(data))
}