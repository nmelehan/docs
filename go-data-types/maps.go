package main

import (
    "fmt"
)

func main() {
    mapA := make(map[string]int)
    mapA["k1"] = 12
    mapA["k2"] = 13
	fmt.Println("mapA:", mapA)
	fmt.Println()

    mapB := map[string]int{
        "k1": 12,
        "k2": 13,
	}

	fmt.Println("mapB:", mapB)
	fmt.Println("Deleting key `k1` from mapB...")
    delete(mapB, "k1")
    delete(mapB, "k1")
	fmt.Println("mapB:", mapB)
	fmt.Println()

	fmt.Println("Checking if the `doesItExist` key exists in mapA...")
    _, ok := mapA["doesItExist"]
    if ok {
        fmt.Println("Exists!")
    } else {
        fmt.Println("Does NOT exist")
	}
	fmt.Println()

	fmt.Println("Iterating through keys and values of mapA...")
    for key, value := range mapA {
        fmt.Println(key, ":", value)
    }
}