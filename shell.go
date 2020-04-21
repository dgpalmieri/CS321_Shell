// shell.go
// Dylan Palmieri
// 4/20/20
// Custom shell for CS321 Operating Systems

package main

import (
    "bufio"
    "fmt"
    "os"
)

func parseInput(input string) (bool, bool) {
    var err, exit = false, false

    if (input == "exit\n"){
        exit = true
    }

    return err, exit
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    for true {
        fmt.Println("----- /home/dgpalmieri/Documents/Operating_Systems/Project_Shell ------")
        fmt.Print("Hello, user: ")
        input, _ := reader.ReadString('\n')
        err, exit := parseInput(input)

        if err {
            fmt.Println("Unable to parse input.")
        }

        if exit {
            break
        }
    }

    fmt.Println("Exiting.")
}
