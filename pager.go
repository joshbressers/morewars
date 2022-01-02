package main

import (
    "bufio"
    "fmt"
    "os"
    "github.com/rthornton128/goncurses"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    file, err := os.Open("/tmp/dat")
    check(err)

    lines := bufio.NewScanner(file)
    for lines.Scan() {
        fmt.Println(lines.Text())
    }

    fmt.Println("")

    s, err := goncurses.Init()
    check(err)
    defer goncurses.End()

    s.Move(5,2)
    s.Print("XXX TEST XXX")
    s.MovePrint(10,10, "XXX 10,10 XXX")
    s.Refresh()

    s.GetChar()
}
