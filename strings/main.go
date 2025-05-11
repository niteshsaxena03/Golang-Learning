package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = "apple,orange,banana"
	var parts = strings.Split(data, ",")
	fmt.Println(parts)

	var data2 = "one one one two two two"
	fmt.Println(strings.Count(data2, "one"))

	var extra string = "   hello    "
	fmt.Println(strings.TrimSpace(extra))

	var firstName string = "Nitesh"
	var lastName string = "Saxena"
	var joined = strings.Join([]string{firstName, lastName}, ",")
	fmt.Println(joined)

}
