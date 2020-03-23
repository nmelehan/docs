package main

import (
    "fmt"
)

func main() {
    oneDimension := [5]int{-1, -2, -3, -4, -5}
    twoDimension := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {-13, -14, -15, -16}}
    threeDimension := [2][2][2]int{ {{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}} }

    fmt.Println("The length of", oneDimension, "is", len(oneDimension))
    fmt.Println("The first element of twoDimension (", twoDimension, ") is", twoDimension[0][0])
    fmt.Println("The length of threeDimension (", threeDimension, ") is", len(threeDimension))

	fmt.Println("\nIterating through threeDimension: ")
    for _, topLevelArray := range threeDimension {
        for _, secondLevelArray := range topLevelArray {
            for _, thirdLevelArray := range secondLevelArray {
                fmt.Print(thirdLevelArray, " ")
            }
        }
        fmt.Println()
    }
}