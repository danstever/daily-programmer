package main

import (
    "fmt"
)

func prefix(n int) string {
    switch {
    case n == 11:
        return "th"
    case n == 12:
        return "th"
    case n == 13:
        return "th"
    case n == 1 || (n > 10 && n % 10 == 1):
        return "st"
    case n == 2 || (n > 10 && n % 10 == 2):
        return "nd"
    case n == 3 || (n > 10 && n % 10 == 3):
        return "rd"
    }
    return "th"
}

func main() {
    var place int

    fmt.Print("What place did your dog finish? ")
    _, err := fmt.Scanf("%d", &place)
    if err != nil {
        fmt.Println("not a number")
    }

    for i := 1; i < 151; i++ {
        if i == place {
            continue
        }
        prefix := prefix(i)
        fmt.Printf("%d%s, ", i, prefix)
    }

}
