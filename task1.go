package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func isPangram(s string) string {
    s = strings.ToLower(s)
    letters := make(map[rune]bool)

    for _, c := range s {
        if c >= 'a' && c <= 'z' {
            letters[c] = true
        }
    }

    if len(letters) == 26 {
        return "Yes"
    }
    return "No"
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Введите строку: ")
    input, _ := reader.ReadString('\n')
    fmt.Println(isPangram(input))
}

