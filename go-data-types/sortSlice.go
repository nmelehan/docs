package main

import (
    "fmt"
    "sort"
)

func main() {
    mySlice := make([]int, 0)
    mySlice = append(mySlice, 90)
    mySlice = append(mySlice, 45)
    mySlice = append(mySlice, 45)
    mySlice = append(mySlice, 50)
    mySlice = append(mySlice, 0)

    fmt.Println("Unsorted:\t\t", mySlice)

    sort.Slice(mySlice, func(i, j int) bool {
        return mySlice[i] < mySlice[j]
    })
    fmt.Println("Ascending order:\t", mySlice)

    sort.Slice(mySlice, func(i, j int) bool {
        return mySlice[i] > mySlice[j]
    })
    fmt.Println("Descending order:\t", mySlice)
}