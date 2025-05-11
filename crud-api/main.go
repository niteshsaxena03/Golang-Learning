package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	var response, err = http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error making request", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		fmt.Println("Request successful")
	} else {
		fmt.Println("Request failed", response.StatusCode)
	}

	// var data, error = io.ReadAll(response.Body)
	// if error != nil {
	// 	fmt.Println("Error reading response", error)
	// 	return
	// }
	// fmt.Println(string(data))

	var todo Todo
	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding response", err)
		return
	}
	fmt.Println(todo)
}

func performPostRequest() {
	var todo = Todo{
		UserId:    1,
		ID:        1,
		Title:     "Buy groceries",
		Completed: false,
	}
	var jsonData, err = json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling json", err)
		return
	}
	var jsonString = string(jsonData)
	var jsonReader = strings.NewReader(jsonString)

	var myURL = "https://jsonplaceholder.typicode.com/todos"
	var response, error = http.Post(myURL, "application/json", jsonReader)
	if error != nil {
		fmt.Println("Error making request", error)
		return
	}
	defer response.Body.Close()

	var data, _ = io.ReadAll(response.Body)
	fmt.Println(string(data))
}

func performUpdateRequest() {
	var todo = Todo{
		UserId:    2,
		ID:        2,
		Title:     "Buy temp",
		Completed: true,
	}
	var jsonData, _ = json.Marshal(todo)

	var jsonString = string(jsonData)
	var jsonReader = strings.NewReader(jsonString)

	var myURL = "https://jsonplaceholder.typicode.com/todos/2"
	var request, error = http.NewRequest("PUT", myURL, jsonReader)
	if error != nil {
		fmt.Println("Error making request", error)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	defer request.Body.Close()

	var client = http.Client{}
	var res, err = client.Do(request)
	if err != nil {
		fmt.Println("Error making request", err)
		return
	}
	defer res.Body.Close()

	var data, _ = io.ReadAll(res.Body)
	fmt.Println(string(data))
}

func performDeleteRequest() {
	var myURL = "https://jsonplaceholder.typicode.com/todos/2"
	var request, err = http.NewRequest("DELETE", myURL, nil)
	if err != nil {
		fmt.Println("Error making request", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	var client = http.Client{}
	var res, error = client.Do(request)
	if error != nil {
		fmt.Println("Error making request", error)
		return
	}
	defer res.Body.Close()

	fmt.Println("Status code", res.Status)
}
func main() {
	fmt.Println("Started crud api in go")
	//performGetRequest()
	//performPostRequest()
	//performUpdateRequest()
	performDeleteRequest()
}
