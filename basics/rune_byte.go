package main

import (
    "fmt"
)

func main() {

    msg := "🐘 🦥 🐋"

    data := []rune(msg)
    fmt.Println(data)

    data2 := []byte(msg)
    fmt.Println(data2)
}
