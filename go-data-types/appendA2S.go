package main

import (
    "fmt"
)

func main() {
    anArray := [3]int{14, 15, 16}
	aSlice := []int{-1, -2, -3}
	fmt.Println("aSlice:\t\t\t\t", aSlice)
	fmt.Println("anArray:\t\t\t", anArray)

	sliceFromArray := anArray[:]
	fmt.Println("sliceFromArray:\t\t\t", sliceFromArray)
	fmt.Println()

    aSlicePlusSliceFromArray := append(aSlice, sliceFromArray...)
	fmt.Println("aSlice + sliceFromArray:\t", aSlicePlusSliceFromArray)

	aSlicePlusSameSlice := append(aSlice, aSlice...)
    fmt.Println("aSlice + aSlice:\t\t", aSlicePlusSameSlice)
}