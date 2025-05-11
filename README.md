# Go Programming - Revision Notes

## Introduction to Go

Go (or Golang) is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It was announced in 2009 and released as an open source project in 2012. Go was created to address criticisms of other languages while maintaining their positive characteristics.

Key features that make Go unique:

- Simplicity and readability
- Fast compilation
- Built-in concurrency with goroutines and channels
- Strong but flexible type system
- Garbage collection with low latency
- Static linking (produces standalone binaries)
- Cross-platform support

## Basics

The foundation of any Go program includes package declarations, import statements, and the main function as an entry point for executables.

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

### Variable Declaration

Go supports multiple ways to declare variables:

1. Using `var` keyword with explicit type: `var name string`
2. Using `var` with initialization: `var name = "John"` (type is inferred)
3. Short declaration (inside functions): `name := "John"`
4. Multiple variables: `var a, b, c int = 1, 2, 3`
5. Block declaration:
   ```go
   var (
       name string
       age int
       isActive bool
   )
   ```

### Constants

Constants in Go are declared using the `const` keyword. They can be character, string, boolean, or numeric values that cannot be changed after declaration. Go also provides an `iota` identifier that can create a sequence of related constants.

```go
const (
    Sunday = iota // 0
    Monday        // 1
    Tuesday       // 2
    Wednesday     // 3
)
```

## Data Types

Go has a rich set of built-in data types to represent different kinds of data efficiently.

### Basic Types

Go's basic types include numeric types, strings, and booleans.

```go
// Numbers
var intVal int = 42          // Platform dependent (32 or 64 bit)
var floatVal float64 = 3.14  // 64-bit floating point

// String
var str string = "hello"

// Boolean
var isTrue bool = true
```

#### Numeric Types

Go provides various numeric types:

- Integers: `int8`, `int16`, `int32`, `int64`, `int`, `uint8`, `uint16`, `uint32`, `uint64`, `uint`
- Floating point: `float32`, `float64`
- Complex: `complex64`, `complex128`
- Byte: Alias for `uint8`
- Rune: Alias for `int32` (represents a Unicode code point)

The default `int` and `uint` types are platform-dependent (32-bit on 32-bit systems, 64-bit on 64-bit systems).

#### String Type

Strings in Go are immutable sequences of bytes, typically representing UTF-8 encoded text. Go provides extensive support for string manipulation through the `strings` package.

#### Boolean Type

Boolean type in Go is represented by the keyword `bool` with possible values `true` and `false`. They are commonly used in conditional statements and as flags.

### Arrays and Slices

Arrays and slices store sequences of elements. Arrays have fixed size while slices are dynamic.

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

#### Arrays

An array in Go has a fixed size defined at declaration time. The size is part of the type, which means `[5]int` and `[10]int` are different types.

Key characteristics of arrays:

- Fixed size
- Size is part of the type
- Zero-based indexing
- Arrays are value types (copying an array creates a complete copy)
- Elements are initialized to their zero values by default

#### Slices

Slices are more flexible than arrays and are used more frequently in Go code. A slice is a descriptor of an array segment, providing access to a numbered sequence of elements from an underlying array.

Key characteristics of slices:

- Dynamic size
- Reference to an underlying array
- Three components: pointer to the array, length, and capacity
- Created using `make()` function or slice literals
- Zero value is `nil`
- Common operations: `append()`, `copy()`, `len()`, `cap()`

### Maps

Maps are Go's built-in associative data type that maps keys to values (similar to dictionaries or hash tables in other languages).

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

Key characteristics of maps:

- Dynamic size
- Reference type (when you assign a map to a new variable, both variables refer to the same map)
- Keys must be comparable (can be tested for equality)
- Zero value is `nil`
- Writing to a nil map causes a panic
- Not safe for concurrent use without proper synchronization

## Control Structures

Control structures in Go determine the flow of execution in a program. Go keeps it simple with fewer control structures than other languages but provides flexibility.

### If-Else

Conditional execution is handled using `if`, `else if`, and `else` statements.

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

Special features of Go's if statements:

- No parentheses required around conditions
- Braces `{}` are mandatory
- Can include a short statement before the condition
  ```go
  if err := someFunction(); err != nil {
      // Handle error
  }
  ```
- The variables declared in the short statement are only in scope until the end of the if statement

### Switch-Case

Switch statements provide a cleaner way to express complex conditionals.

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

Key features of Go's switch statement:

- Cases evaluate from top to bottom, stopping when a case succeeds
- No fall-through by default (unlike C, C++, and Java)
- Can use the `fallthrough` keyword to explicitly fall through to the next case
- Can have multiple expressions in a single case (comma-separated)
- Expression is optional (acts like if-else chain)
- Type switches allow checking types of interface values

### Loops

Go simplifies looping by using only the `for` loop construct, which can be used in various ways.

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

Key features of Go's `for` loops:

- Three components: init statement, condition expression, post statement
- All components are optional (creating various loop types)
- The init statement and post statement can be multiple variable operations separated by commas
- `break` and `continue` statements for controlling loop execution
- The `range` keyword for iterating over data structures
- No while or do-while loops (use for loops instead)

## Functions

Functions are central to Go's organization and reusability of code. They serve as the primary mechanism for code organization and reuse.

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

Key features of Go functions:

- First-class citizens (can be assigned to variables and passed as arguments)
- Can return multiple values (commonly used for returning results and errors)
- Support for named return values
- Can be defined inside other functions (closures)
- Variadic functions (variable number of arguments)
- Function types and function literals
- No default parameters or method overloading

### Function Parameters

Go passes function parameters by value. This means when a function is called, each argument is copied to the corresponding parameter. If you want to modify the original values, you need to use pointers.

### Multiple Return Values

One of Go's distinctive features is the ability for functions to return multiple values. This is commonly used to return a result and an error.

```go
func getUserInfo(id int) (string, int, error) {
    // Logic to retrieve user info
    return "John", 30, nil
}

name, age, err := getUserInfo(42)
if err != nil {
    // Handle error
}
```

### Variadic Functions

Go supports functions that can take a variable number of arguments.

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Call with individual arguments
sum(1, 2, 3)

// Call with a slice
numbers := []int{1, 2, 3, 4}
sum(numbers...)
```

## Structs

Structs in Go are composite types that group together variables (fields) under a single name. They're used to represent records, objects, and custom data types.

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

Key aspects of Go structs:

- Fields can be of any type, including other structs
- Fields can have tags (metadata used by reflection)
- Zero value is a struct with all fields set to their zero values
- Can be compared with `==` if all fields are comparable
- Exported fields (capitalized) are accessible outside the package
- Support for embedding (composition over inheritance)
- Can have associated methods

### Methods

Methods are functions associated with a specific type. In Go, you define methods by specifying a receiver parameter.

```go
type Rectangle struct {
    width, height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.width * r.height
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.width *= factor
    r.height *= factor
}

rect := Rectangle{width: 10, height: 5}
area := rect.Area()
rect.Scale(2)
```

Value receivers operate on a copy of the value, while pointer receivers can modify the actual value.

### Struct Tags

Struct tags provide metadata about struct fields and are commonly used for encoding/decoding data.

```go
type User struct {
    Name     string `json:"name" xml:"name" validate:"required"`
    Age      int    `json:"age" xml:"age" validate:"gte=0,lte=130"`
    Email    string `json:"email,omitempty" validate:"required,email"`
    Password string `json:"-"` // Hyphen means field will be ignored
}
```

Tags are interpreted by reflection and used by packages like `encoding/json`, `encoding/xml`, and validation libraries.

## Pointers

Pointers are variables that store the memory address of another variable. They're useful for sharing data across functions without copying.

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

Key pointer concepts in Go:

- Zero value is `nil`
- `&` operator obtains the address of a variable
- `*` operator dereferences a pointer (accesses the value it points to)
- Pointer arithmetic is not allowed
- Go is safe with pointers (no pointer arithmetic, garbage collection)
- Common uses: modifying arguments, efficiency for large data structures, implementing data structures

## Error Handling

Go takes a different approach to error handling compared to many other languages. Instead of exceptions, Go functions return error values.

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

Key aspects of Go error handling:

- Errors are values (typically interface implementations)
- Functions return errors as additional return values
- Explicit error checking with if statements
- The `errors` package for creating simple errors
- `fmt.Errorf()` for formatted error messages
- Custom error types can be created by implementing the `error` interface
- Sentinel errors (predefined error values) can be checked with `==`

## File Operations

Go provides robust file handling capabilities through the `os` and `io` packages.

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

Common file operations in Go:

- Creating files with `os.Create()`
- Opening files with `os.Open()` (read-only) or `os.OpenFile()` (with options)
- Reading entire files with `os.ReadFile()`
- Writing entire files with `os.WriteFile()`
- Buffered reading with `bufio.Reader`
- Buffered writing with `bufio.Writer`
- File information with `os.Stat()`
- Directory operations with `os.Mkdir()`, `os.MkdirAll()`, `os.ReadDir()`

## Defer Statement

The `defer` statement schedules a function call to be executed immediately before the surrounding function returns. It's commonly used for cleanup operations.

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

Key aspects of defer statements:

- Deferred calls are executed in LIFO (Last In, First Out) order
- Arguments to deferred functions are evaluated when the defer statement is encountered, not when the function is called
- Commonly used for resource cleanup (closing files, network connections, etc.)
- Can be used with anonymous functions
- Deferred functions can access and modify named return values
- Executed even if the function panics

## User Input

Go provides several methods for accepting user input through the standard library.

```go
var x int
fmt.Println("Enter a number")
fmt.Scan(&x)
fmt.Println("Number is", x)
```

Methods for handling user input:

- `fmt.Scan()` - Scans space-separated values
- `fmt.Scanf()` - Scans formatted input
- `fmt.Scanln()` - Scans a line of input
- `bufio.Scanner` - For more complex input parsing

## JSON Handling

Go has excellent support for working with JSON through the `encoding/json` package.

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

Key features of JSON handling in Go:

- Marshal (struct to JSON) with `json.Marshal()` or `json.MarshalIndent()`
- Unmarshal (JSON to struct) with `json.Unmarshal()`
- Streaming encoding with `json.Encoder`
- Streaming decoding with `json.Decoder`
- Struct tags control field names and behavior
- Tag options: `omitempty` (omit empty fields), `-` (ignore field)
- Works with custom types by implementing the `json.Marshaler` and `json.Unmarshaler` interfaces

## HTTP Requests

Go's `net/http` package provides comprehensive support for HTTP clients and servers.

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

Key aspects of HTTP client operations:

- Simple requests with `http.Get()`, `http.Post()`, `http.Head()`
- Custom requests with `http.NewRequest()` and `http.Client`
- Custom headers, cookies, and request parameters
- Timeouts and cancellation with contexts
- Response handling with `response.Body` (implements `io.ReadCloser`)
- Always close the response body with `defer response.Body.Close()`

## URL Handling

The `net/url` package provides functions for parsing and manipulating URLs.

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

Key features of URL handling:

- Parse URLs with `url.Parse()`
- Access URL components (scheme, host, path, query, etc.)
- URL encoding and decoding
- Query parameter handling with `url.Values`
- Relative URL resolution
- Building URLs with `url.URL` struct

## Time and Date

Go's `time` package provides comprehensive support for handling time and dates.

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

Key features of time handling in Go:

- Time representation with `time.Time` struct
- Current time with `time.Now()`
- Time zones with `time.LoadLocation()` and `time.Time.In()`
- Duration with `time.Duration` type
- Time arithmetic (add, subtract, compare)
- Formatting with the reference time: "Mon Jan 2 15:04:05 MST 2006"
- Parsing with `time.Parse()` and `time.ParseInLocation()`
- Timers and tickers for scheduled operations

## Type Conversion

Go provides several methods for converting between different data types.

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

Type conversion mechanisms in Go:

- Basic type conversion with T(v) syntax (e.g., `float64(42)`)
- String conversions with the `strconv` package
- Numeric parsing with `ParseInt()`, `ParseFloat()`, `ParseUint()`
- Numeric formatting with `FormatInt()`, `FormatFloat()`, `FormatUint()`
- Boolean conversions with `ParseBool()` and `FormatBool()`
- Type assertions for interface values
- Type switches for multiple type checks

## Custom Packages

Go's package system allows you to organize your code into reusable components.

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

Key aspects of Go packages:

- Each source file belongs to a package
- Package name should match the directory name (with exceptions)
- Exported identifiers start with an uppercase letter
- Unexported identifiers start with a lowercase letter
- Import paths are based on module paths
- Package initialization with `init()` functions
- Circular imports are not allowed
- Standard library packages are imported without a domain

## Concurrency in Go

Go provides built-in support for concurrency with goroutines and channels.

### Goroutines

Goroutines are lightweight threads managed by the Go runtime. They allow concurrent execution with minimal resource overhead.

```go
func printNumbers() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Start a goroutine
    go printNumbers()

    // Continue execution in main goroutine
    for i := 0; i < 5; i++ {
        fmt.Println("Main:", i)
        time.Sleep(150 * time.Millisecond)
    }

    // Wait to see output
    time.Sleep(time.Second)
}
```

### Channels

Channels provide a way for goroutines to communicate and synchronize.

```go
func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // Send sum to channel
}

func main() {
    s := []int{7, 2, 8, 9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)

    x, y := <-c, <-c // Receive from channel

    fmt.Println(x, y, x+y)
}
```

### Select Statement

The select statement allows a goroutine to wait on multiple communication operations.

```go
func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}
```

## Best Practices

1. Always check errors

   - Go's error handling requires explicit checks
   - Avoid the `_` blank identifier for errors unless you're sure they can be ignored

2. Use `defer` for cleanup operations

   - Ensures resources are properly released even if errors occur
   - Commonly used for closing files, network connections, etc.

3. Keep functions small and focused

   - Each function should do one thing well
   - Makes code easier to test, debug, and maintain

4. Use meaningful variable names

   - Short but descriptive names
   - Common Go idioms: `i` for indexes, `v` for values, `k` for keys, `err` for errors

5. Follow Go naming conventions

   - CamelCase for exported (public) names: `ExportedName`
   - camelCase for unexported (private) names: `unexportedName`
   - Acronyms should be in the same case: `HTTPServer` or `httpServer`

6. Use proper error handling patterns

   - Return errors rather than panicking
   - Wrap errors with context when appropriate
   - Don't use `panic` for normal error handling

7. Format your code with `go fmt`

   - Standardizes code formatting
   - Eliminates style debates

8. Use `go vet` and `golint`

   - Catch common mistakes and style issues
   - Part of a good CI/CD pipeline

9. Write tests with the `testing` package

   - Test files should end with `_test.go`
   - Test functions should be named `TestXxx`

10. Document your code
    - Package comments should describe the package
    - Exported functions, types, and variables should have comments
    - Use godoc format for documentation

## Common Packages

- `fmt`: Formatted I/O with functions like `Println`, `Printf`
- `io`: Basic I/O interfaces
- `os`: Operating system functionality like file operations
- `strconv`: String conversions
- `strings`: String manipulation functions
- `time`: Time and date functionality
- `encoding/json`: JSON encoding and decoding
- `net/http`: HTTP client and server
- `net/url`: URL parsing and manipulation
- `sync`: Synchronization primitives
- `log`: Logging functionality
- `flag`: Command-line flag parsing
- `math`: Mathematical functions
- `sort`: Sorting algorithms
- `reflect`: Runtime reflection, allowing inspection of types
- `regexp`: Regular expression matching
- `crypto`: Cryptographic functions
- `database/sql`: Database interface
- `html/template`: HTML templating
- `context`: For cancellation, deadlines, and request-scoped values
