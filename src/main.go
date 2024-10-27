package main

import (
    "fmt"
    "calc_project/calc"
)


func main() {
    expr := "3 + 5 * (2 - 8)"
    result, err := calc.Calc(expr)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
