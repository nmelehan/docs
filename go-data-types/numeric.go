package main

import (
    "fmt"
)

func main() {
    c1 := -12 + 2i
    c2 := complex(5, 7)
    fmt.Println("Type of c1:\t", c1)
    fmt.Println("Type of c2:\t", c2)

    var c3 complex64 = complex64(c1 + c2)
    fmt.Println("c3:\t\t", c3)
    fmt.Println("Type of c3:\t", c3)

    cZero := c3 - c3
    fmt.Println("cZero:\t\t", cZero)

    x := 12
    k := 7
    fmt.Println("Type of x:\t", x)

    div := x / k
    fmt.Println("Integer division generates an integer:", div)
    fmt.Println("Convert to float64:", float64(x)/float64(k))

    var m, n float64
    m = 1.223
    fmt.Println("m, n:\t\t", m, n)

    y := 4 / 2.3
    fmt.Println("y:\t\t", y)

    divFloat := float64(x) / float64(k)
    fmt.Println("divFloat:\t", divFloat)
    fmt.Println("Type of divFloat:", divFloat)
}