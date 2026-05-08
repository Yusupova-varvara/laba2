package main

import (
    "fmt"
    "strconv"
)

func compress(s string) string {
    result := ""
    count := 1

    for i := 0; i < len(s); i++ {
        if i+1 < len(s) && s[i] == s[i+1] {
            count++
        } else {
            result += string(s[i]) + strconv.Itoa(count)
            count = 1
        }
    }

    if len(result) < len(s) {
        return result
    }
    return s
}

func main() {
    var input string
    fmt.Print("Введите строку: ")
    fmt.Scan(&input)
    fmt.Println(compress(input))
}

