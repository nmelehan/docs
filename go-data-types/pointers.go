package main

import (
    "fmt"
)

func triple(n *int) {
    *n = *n * 3
}

func returnTripleOf(n int) *int {
    v := n * 3
    return &v
}

func main() {
    i := -10
    j := 25

    pI := &i
    pJ := &j

    fmt.Println("pI memory address:\t", pI)
    fmt.Println("pJ memory address:\t", pJ)
    fmt.Println("pI value:\t\t", *pI)
    fmt.Println("pJ value:\t\t", *pJ)

    *pI = 123456
    *pI--
    fmt.Println("i:\t\t\t", i)

    triple(pJ)
    fmt.Println("j:\t\t\t", j)

    k := returnTripleOf(12)
    fmt.Println("Value of k:\t\t", *k)
    fmt.Println("Memory address of k:\t", k)
}