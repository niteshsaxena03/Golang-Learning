# Go Programming - Revision Notes

## Basics

```go
// Package declaration
package main

// Import statements
import (
    "fmt"
    "strconv"
)

// Main function - entry point
func main() {
    fmt.Println("Hello World")
}

// Variables
var number = 189.98   // Type inference
const name = "nitesh" // Constants
boolean := true       // Short declaration (only inside functions)
```

## Data Types

### Basic Types

```go
// Numbers
var intVal int = 42          // Platform dependent (32 or 64 bit)
var floatVal float64 = 3.14  // 64-bit floating point

// String
var str string = "hello"

// Boolean
var isTrue bool = true
```

### Arrays and Slices

```go
// Arrays - fixed size
var array [5]string
array[0] = "nitesh"
array[1] = "saxena"
numbers := [5]int{1, 2, 3, 4, 5}

// Slices - dynamic arrays
slice := []int{1, 2, 3, 4, 5}
slice = append(slice, 2, 3, 4, 5)
fmt.Println("The length is", len(slice))
```

### Maps

```go
// Declaration and initialization
stringMap := make(map[string]int)
stringMap["key1"] = 1
stringMap["key2"] = 2

// Predefined maps
newMap := map[string]int{
    "key1": 1,
    "key2": 2,
}

// Checking if key exists
marks, exists := stringMap["key1"]

// Deleting keys
delete(stringMap, "key1")

// Iteration
for key, value := range stringMap {
    fmt.Println(key, value)
}
```

## Control Structures

### If-Else

```go
if x > 5 {
    fmt.Println("x is greater than 5")
} else if x < 5 {
    fmt.Println("x is less than 5")
} else {
    fmt.Println("x is equal to 5")
}

// Compound conditions
if x == 5 && y == 10 {
    fmt.Println("Condition is true")
}
```

### Switch-Case

```go
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
default:
    fmt.Println("Invalid day")
}

// Switch with no expression (like if-else)
switch {
case temperature < 0:
    fmt.Println("It is too cold")
case temperature > 0 && temperature < 15:
    fmt.Println("It is cool")
default:
    fmt.Println("It is hot")
}
```

### Loops

```go
// Traditional for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-like loop
counter := 0
for {
    fmt.Println("Infinite loop")
    counter++
    if counter > 5 {
        break
    }
}

// Range-based loop
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Println(index, value)
}
```

## Functions

```go
// Basic function
func add(x int, y int) int {
    return x + y
}

// Multiple return values
func divide(x float64, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("divisor cant be 0")
    }
    return x / y, nil
}

// Using functions
sum := add(10, 20)
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
}
```

## Structs

```go
// Basic struct
type Person struct {
    firstName string
    lastName  string
    age       int
}

// Nested structs
type Employee struct {
    person  Person
    contact Contact
    address Address
}

// Using structs
var person1 Person
person1.firstName = "Nitesh"
person1.lastName = "Saxena"
person1.age = 22

// Struct with JSON tags
type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    IsAdult bool   `json:"is_adult"`
}
```

## Pointers

```go
// Declaring pointers
var num int = 10
var ptr *int
ptr = &num

fmt.Println("Number is ", num)     // Value
fmt.Println("Pointer is ", ptr)     // Memory address
fmt.Println("Value at pointer ", *ptr) // Dereferencing

// Using pointers in functions
func modifyPointer(ptr *int) {
    *ptr = *ptr * 2
}
```

## Error Handling

```go
// Error checking pattern
if err != nil {
    fmt.Println("Error:", err)
    return
}

// Creating custom errors
func divide(x float64, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("divisor cant be 0")
    }
    return x / y, nil
}
```

## File Operations

```go
// Creating and writing to files
file, err := os.Create("test.txt")
if err != nil {
    fmt.Println("Error creating file", err)
    return
}
defer file.Close() // Important for resource cleanup

text := "Hello, World!"
file.WriteString(text)

// Reading files
content, err := os.ReadFile("test.txt")
if err != nil {
    fmt.Println("Error reading file", err)
    return
}
fmt.Println(string(content))

// Reading chunks
file, err := os.Open("test.txt")
buffer := make([]byte, 1024)
for {
    n, err := file.Read(buffer)
    if err == io.EOF {
        break
    }
    fmt.Println(string(buffer[:n]))
}
```

## Defer Statement

```go
// Defer executes when the surrounding function returns
func main() {
    fmt.Println("Starting the defer keyword")
    fmt.Println("This is the first statement")
    defer fmt.Println("This is the second statement")
    defer fmt.Println("Another defer statement")
    fmt.Println("This is the third statement")
}
// Output order: first, third, another defer, second
```

## User Input

```go
var x int
fmt.Println("Enter a number")
fmt.Scan(&x)
fmt.Println("Number is", x)
```

## JSON Handling

```go
// Struct with JSON tags
type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    IsAdult bool   `json:"is_adult"`
}

// Converting struct to JSON
jsonData, err := json.Marshal(person)
if err != nil {
    fmt.Println("Error marshalling json", err)
    return
}
fmt.Println(string(jsonData))

// Converting JSON to struct
var personData Person
err = json.Unmarshal(jsonData, &personData)
if err != nil {
    fmt.Println("Error unmarshalling json", err)
    return
}
```

## HTTP Requests

```go
// GET request
response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
if err != nil {
    fmt.Println("Error making request", err)
    return
}
defer response.Body.Close()

data, error := io.ReadAll(response.Body)
fmt.Println(string(data))

// POST request
todo := Todo{
    UserId:    1,
    ID:        1,
    Title:     "Buy groceries",
    Completed: false,
}
jsonData, _ := json.Marshal(todo)
jsonString := string(jsonData)
jsonReader := strings.NewReader(jsonString)

response, err := http.Post(url, "application/json", jsonReader)

// PUT request
request, err := http.NewRequest("PUT", url, jsonReader)
request.Header.Set("Content-Type", "application/json")

client := http.Client{}
res, err := client.Do(request)

// DELETE request
request, err := http.NewRequest("DELETE", url, nil)
client := http.Client{}
res, err := client.Do(request)
```

## URL Handling

```go
urlString := "https://jsonplaceholder.typicode.com/todos/1"
url, err := url.Parse(urlString)
if err != nil {
    fmt.Println("Error parsing url", err)
    return
}

fmt.Println(url.Scheme)  // https
fmt.Println(url.Host)    // jsonplaceholder.typicode.com
fmt.Println(url.Path)    // /todos/1

// Modifying URL
url.Path = "/users"
fmt.Println(url.String())
```

## Time and Date

```go
// Current time
currentTime := time.Now()
fmt.Println(currentTime)

// Formatting time
formattedTime := currentTime.Format("2006-01-02 15:04:05")
fmt.Println(formattedTime)

newFormattedTime := currentTime.Format("2006/01/02, Monday, 3:04 PM")
fmt.Println(newFormattedTime)

// Parsing time
parsedTime, _ := time.Parse("2006-01-02 15:04:05", formattedTime)
fmt.Println(parsedTime)

// Adding time
newTime := currentTime.AddDate(0, 0, 1) // Add 1 day
fmt.Println(newTime)
```

## Type Conversion

```go
// Integer to float
num := 42
num1 := float64(num)

// Integer to string
s1 := strconv.Itoa(num)

// String to integer
numberString := "123"
num2, _ := strconv.Atoi(numberString)

// String to float
num3, _ := strconv.ParseFloat(numberString, 64)
```

## Custom Packages

```go
// In myutil/myutil.go
package myutil

import "fmt"

func PrintMessage(message string) {
    fmt.Println(message)
}

// In main.go
import "myproject/myutil"

func main() {
    myutil.PrintMessage("Hello from custom package")
}
```

## Best Practices

1. Always check errors
2. Use `defer` for cleanup operations
3. Keep functions small and focused
4. Use meaningful variable names
5. Follow Go naming conventions (CamelCase for exported, camelCase for unexported)
6. Use proper error handling patterns
7. Format your code with `go fmt`
