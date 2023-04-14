package main

import (
        "strings"
)

func LeftPad(str string, length int, char string) string {
        if len(str) >= length {
                return str
        }
        return strings.Repeat(char, length-len(str)) + str
}

func main() {
        padded := LeftPad("hello", 10, "*")
        println(padded) // *****hello
}
