package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    // s, sep := "", ""
    for idx, arg := range os.Args[1:] {
        // s += sep + arg
        // sep = " "
        fmt.Println(idx)
        s := string(idx) + " " + arg
        fmt.Println(s)
    }
    join()
    fmt.Println(os.Args[1:])
}

func join() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}