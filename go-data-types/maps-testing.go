package main

import (
    "fmt"
)

func main() {
    aMap := make(map[string]int)
	aMap["one"] = 1
	// Prints "1"
	fmt.Println(aMap["one"])
}