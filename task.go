// #1
//package main

//import (
    //"fmt"
    //"strconv"
//)

//func compress(s string) string {
    //result := ""
    //count := 1
    
    //for i := 0; i < len(s); i++ {
        //if i+1 < len(s) && s[i] == s[i+1] {
            //count++
        //} else {
            //result += string(s[i]) + strconv.Itoa(count)
            //count = 1
        //}
    //}
    
    //if len(result) < len(s) {
        //return result
    //}
    //return s
//}

//func main() {
    //var input string
    //fmt.Print("Введите строку: ")
    //fmt.Scan(&input)
    //fmt.Println(compress(input))
//}

//#2

//package main

//import (
    //"bufio"
    //"fmt"
    //"os"
    //"strings"
//)

//func isPangram(s string) string {
    //s = strings.ToLower(s)
    //letters := make(map[rune]bool)

    //for _, c := range s {
        //if c >= 'a' && c <= 'z' {
            //letters[c] = true
        //}
    //}

    //if len(letters) == 26 {
        //return "Yes"
    //}
    //return "No"
//}

//func main() {
    //reader := bufio.NewReader(os.Stdin)
    //fmt.Print("Введите строку: ")
    //input, _ := reader.ReadString('\n')
    //fmt.Println(isPangram(input))
//} 

//#3

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


