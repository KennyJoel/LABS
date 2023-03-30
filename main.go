package main

import "fmt"

func main() {
    x := 5
    y := &x // y is a pointer to x

    fmt.Println("x before:", x)
    changeValue(x)
    fmt.Println("x after pass by value:", x)

    fmt.Println("x before:", x)
    changePointer(y)
    fmt.Println("x after pass by reference:", x)
}

func changeValue(val int) {
    val = val + 10
}

func changePointer(ptr *int) {
    *ptr = *ptr + 10
}
