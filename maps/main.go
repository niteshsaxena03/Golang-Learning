package main

import "fmt"

func main() {
	fmt.Println("learning maps in go")

	var stringMap = make(map[string]int)

	stringMap["key1"] = 1
	stringMap["key2"] = 2

	fmt.Println(stringMap)
	fmt.Println(stringMap["key1"])

	delete(stringMap, "key1")
	fmt.Println(stringMap)

	//check if a key exists
	var marks, exists = stringMap["key1"]
	fmt.Println(marks, exists)

	for key, value := range stringMap {
		fmt.Println(key, value)
	}

	var newMap = map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	fmt.Println(newMap)

}
