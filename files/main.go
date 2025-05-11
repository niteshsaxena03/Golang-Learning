package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("learning file handling in go")

	// var file,err = os.Create("test.txt")
	// if err != nil{
	// 	fmt.Println("Error creating file",err)
	// 	return
	// }
	// fmt.Println("File created successfully",file)
	// defer file.Close();

	// var text = "Hello, World!";
	// file.WriteString(text);
	// fmt.Println("File written successfully");

	var file, err = os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	fmt.Println("File opened successfully", file)

	var buffer = make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}

	var content, error = os.ReadFile("test.txt")

	if error != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println(string(content))
}
