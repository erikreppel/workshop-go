package main

import "fmt"

// TODO: Make a function that adds two numbers and returns the result
func add(a int, b int) int {
    sum := a + b
    return sum
}

func main() {
    fmt.Println(add(4, 5))
}