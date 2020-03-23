package main

import (
    "fmt"
)

func main() {
    aSlice := []int{1, 2, 3, 4, 5}
    fmt.Println("aSlice:\t\t", aSlice)
    integers := make([]int, 2)
    fmt.Println("integers:\t", integers)
    integers = nil
    fmt.Println("integers:\t", integers)
    fmt.Println()

    anArray := [5]int{-1, -2, -3, -4, -5}
    sliceFromArray := anArray[:]
    fmt.Println("anArray:\t",anArray)
    fmt.Println("sliceFromArray:\t", sliceFromArray)
    anArray[4] = -100
    fmt.Println("Updating fourth element of sliceFromArray...")
    fmt.Println("sliceFromArray:\t", sliceFromArray)
    fmt.Println("anArray:\t", anArray)
    fmt.Println()

    s := make([]byte, 5)
    fmt.Println("Byte slice:\t", s)
    fmt.Println()

    twoDSlice := make([][]int, 3)
    fmt.Println("twoDSlice:\t",twoDSlice)
    fmt.Println("Adding elements to twoDSlice...")
    for i := 0; i < len(twoDSlice); i++ {
        for j := 0; j < 2; j++ {
            twoDSlice[i] = append(twoDSlice[i], i*j)
        }
    }
    for _, secondLevelSlice := range twoDSlice {
        fmt.Println(secondLevelSlice)
    }
}