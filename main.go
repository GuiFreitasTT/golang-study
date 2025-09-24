package main

import "fmt"

func Somar(a, b int) int {
    return a + b
}

func main() {
    resultado := Somar(2, 2)
    fmt.Println("2 + 2 =", resultado)
}
