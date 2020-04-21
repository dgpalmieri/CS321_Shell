// shell.go
// Dylan Palmieri
// 4/20/20
// Custom shell for CS321 Operating Systems

package main

import (
    "bufio"
    "fmt"
    "strings"
    "os"
    "os/exec"
)

func parseInput(input string) (error, bool) {
    var exit = false
    var err (error) = nil

    if (input == "exit"){
        exit = true
    } else {
        command := exec.Command(input)
        command.Stdout = os.Stdout
        command.Stderr = os.Stdout
        err = command.Run()
    }

    return err, exit
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("----- /home/dgpalmieri/Documents/Operating_Systems/Project_Shell ------")
        fmt.Print("Hello, user: ")

        input, _ := reader.ReadString('\n')
        input = strings.TrimSuffix(input, "\n")

        err, exit := parseInput(input)

         if exit {
            break
        }

        if err != nil {
            fmt.Println("Unable to parse/execute input.")
            fmt.Println(err)
        }
   }

    fmt.Println("Exiting.")
}
