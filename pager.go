package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "time"
    "github.com/rthornton128/goncurses"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func draw_screen(s *goncurses.Window, lines []string) {

    for i := range lines {
        s.Move(i,0)
        s.Print(lines[i])
    }

    s.Refresh()
}

func main() {

    var filename string

    if len(os.Args) == 2 {
        filename = os.Args[1]
    } else {
        os.Exit(1)
    }

    file, err := os.Open(filename)
    check(err)

    s, err := goncurses.Init()
    check(err)
    defer goncurses.End()

    row, col := s.MaxYX()

    all_lines := make([]string, row)

    // There's probably a better way to do this
    for i := 0; i < row; i++ {
        all_lines[i] = ""
    }

    lines := bufio.NewScanner(file)
    for lines.Scan() {
        new_lines := make([]string, row)

        // First shift all the strings up
        for i := 0; i < row - 1; i++ {
            // We need to pull out two characters in the middle
            if len(all_lines[i]) == 0 {
                new_lines[i] = all_lines[i+1]
            } else {
                var middle int = len(all_lines[i])/2
                new_lines[i] = fmt.Sprintf(" %s%s ", all_lines[i+1][:middle-1], all_lines[i+1][middle+1:])
            }
        }

        // Now add our new string
        one_line := lines.Text()
        if len(one_line) > col {
            one_line = one_line[:row]
        }
        var pad_len int= (col-len(one_line))/2
        // Pad the string and store it in a new string.
        // We want row - strlen / 2
        one_line = fmt.Sprintf("%s%s", strings.Repeat(" ", pad_len), one_line)
        one_line = fmt.Sprintf("%s%s", one_line, strings.Repeat(" ", pad_len))
        new_lines[row-1] = one_line
        all_lines = new_lines
        draw_screen(s, all_lines)
        time.Sleep(100000000)
    }

    s.GetChar()


}
