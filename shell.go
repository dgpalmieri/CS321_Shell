// shell.go
// Dylan Palmieri
// Custom shell for CS321 Operating Systems

package main

import (
    "bufio"
    "fmt"
    "strings"
    "os"
    "os/exec"
)

func changeDirectory(wordArray []string) (error) {
    var err (error) = nil

    if len(wordArray) == 1 {
        err = os.Chdir("/home")
    }
    if len(wordArray) == 2 {
        err = os.Chdir(wordArray[1])
    }

    return err

}

func executeInput(inputArray [][]string) (error) {
    var err (error) = nil
    var andIndex = 0

    for _, slice := range(inputArray) {
        fmt.Println("slice:", slice)
        if slice[0] == "cd" {
            err = changeDirectory(slice)
            if err != nil {
                break
            }
            continue
        }

        command := exec.Command(slice[0], slice[1:]...)
        command.Stdout = os.Stdout
        command.Stderr = os.Stdout

        if slice[len(slice)-1] == "|" {
        }
        if slice[len(slice)-1] == "&" {
            command.Stdout = nil
            if len(slice) > 2 {
                command = exec.Command(slice[0], slice[1:len(slice)-2]...)
            } else {
                command = exec.Command(slice[0], slice[1:len(slice)-1]...)
            }
        }
        for index, item := range(slice) {
            if item == "&&" {
                andIndex = index
            }
        }

        if andIndex != 0 {
            command_one := exec.Command(slice[0], slice[1:andIndex]...)
            fmt.Println("cmd:", slice[0], "args:", slice[1:andIndex], "index:", andIndex)
            command_one.Stdout = os.Stdout
            command_one.Stderr = os.Stdout
            command_two := exec.Command(slice[andIndex + 1], slice[andIndex + 2:]...)
            fmt.Println("cmd:", slice[andIndex + 1], "args:", slice[andIndex + 2:])
            command_two.Stdout = os.Stdout
            command_two.Stderr = os.Stdout
            err = command_one.Run()
            if err == nil {
                err = command_two.Run()
            }
            andIndex = 0
        } else {
            err = command.Run()
        }

    }

    return err
}

func parseInput(input string) (error, bool) {
    var exit = false
    var err (error) = nil
    var startIndex = 0
    var parsedSlice = make([][]string, 0, 5)
    inputArray := strings.Fields(input)

    for index, field := range(inputArray) {
        if field == "exit" {
            exit = true
            break
        }
        if field == "cd" {
            parsedSlice = append(parsedSlice, inputArray[index: index + 2])
        }
        if field == "|" || field == "&" {
            parsedSlice = append(parsedSlice,inputArray[startIndex:index + 1])
            startIndex = index + 1
        }
    }

    if len(parsedSlice) == 0 {
        parsedSlice = append(parsedSlice, inputArray)
    }

    fmt.Println("parsedSlice:", parsedSlice)

    err = executeInput(parsedSlice)

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

func prompt() {
    pwdCommand := exec.Command("pwd")
    byteSlice, _ := pwdCommand.Output()
    workdir := string(byteSlice)
    workdir = strings.TrimSuffix(workdir, "\n")

    fmt.Println("-----", workdir, "-----")
    fmt.Print("Hello, user: ")
}

func getInput(reader *bufio.Reader) (string) {
    input, _ := reader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")
    return input
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        prompt()
        input := getInput(reader)
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
