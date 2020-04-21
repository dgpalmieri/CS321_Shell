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

func appendHistory(str string) {
    fstream, err := os.OpenFile(".histfile", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error:", err)
    }

    defer fstream.Close()

    if _, err := fstream.WriteString(str + "\n"); err != nil {
        fmt.Println("Error:", err)
    }

}

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        pwdCommand := exec.Command("pwd")
        byteSlice, _ := pwdCommand.Output()
        workdir := string(byteSlice)
        workdir = strings.TrimSuffix(workdir, "\n")

        fmt.Println("-----", workdir, "-----")
        fmt.Print("Hello, user: ")

        input, _ := reader.ReadString('\n')
        input = strings.TrimSuffix(input, "\n")

        appendHistory(input)
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
