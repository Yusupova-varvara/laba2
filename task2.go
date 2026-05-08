package main

import "fmt"

func main() {
    var n, x int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Scan(&x)

        if x < 10 {
            continue
        }

        lastDigit := x % 10
        temp := x
        good := true

        for temp > 0 {
            if temp%10 != lastDigit {
                good = false
                break
            }
            temp = temp / 10
        }

        if good {
            fmt.Print(x, " ")
        }
    }
    fmt.Println()
}

