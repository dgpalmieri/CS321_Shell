// shell.go
// Dylan Palmieri
// Basic shell for CS321 Operating Systems

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
    var andIndexList = []int{}
    andIndexList = append(andIndexList, -1)
    var pipeIndexList = []int{}

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
                andIndexList = append(andIndexList, index)
            }
            if item == "|" {
                pipeIndexList = append(pipeIndexList, index)
            }
        }

        if len(andIndexList) > 1 {
            var end = 0
            var args = 0
            for index, andIndex := range(andIndexList) {
                if end < len(andIndexList) - 1{
                    end = andIndexList[index + 1]
                } else {
                    end = len(slice)
                }
                args = andIndex + 2
                if args > len(slice) - 1 {
                    args = end
                }
                command := exec.Command(slice[andIndex + 1], slice[args:end]...)
                err = command.Run()
                if err != nil {
                    break
                }
            }
        } else if len(pipeIndexList) > 0 {
            if len(pipeIndexList) == 1 {
                cmd_one := exec.Command(slice[0],slice[1:pipeIndexList[0]]...)
                cmd_two := exec.Command(slice[pipeIndexList[0] + 1],slice[pipeIndexList[0]+2:]...)
                fmt.Println(slice[0],slice[1:pipeIndexList[0]])
                fmt.Println(slice[pipeIndexList[0] + 1],slice[pipeIndexList[0]+2:])

                cmd_two.Stdin, _ = cmd_one.StdoutPipe()
                cmd_two.Stdout = os.Stdout

                cmd_one.Start()
                cmd_two.Start()
                cmd_one.Wait()
                cmd_two.Wait()
            }
            if len(pipeIndexList) == 2 {
                cmd_one := exec.Command(slice[0],slice[1:pipeIndexList[0]]...)
                cmd_two := exec.Command(slice[pipeIndexList[0] + 1],
                                        slice[pipeIndexList[0]+2:pipeIndexList[1]]...)
                cmd_three := exec.Command(slice[pipeIndexList[1] + 1],
                                          slice[pipeIndexList[1]+2:]...)
                fmt.Println(slice[0],slice[1:pipeIndexList[0]])
                fmt.Println(slice[pipeIndexList[0] + 1],
                            slice[pipeIndexList[0]+2:pipeIndexList[1]])
                fmt.Println(slice[pipeIndexList[1] + 1],slice[pipeIndexList[1]+2:])

                cmd_two.Stdin, _ = cmd_one.StdoutPipe()
                cmd_three.Stdin, _ = cmd_two.StdoutPipe()
                cmd_three.Stdout = os.Stdout

                cmd_one.Start()
                cmd_two.Start()
                cmd_three.Start()
                cmd_one.Wait()
                cmd_two.Wait()
                cmd_three.Wait()

            }
            if len(pipeIndexList) == 3 {
                cmd_one := exec.Command(slice[0],slice[1:pipeIndexList[0]]...)
                cmd_two := exec.Command(slice[pipeIndexList[0] + 1],
                                        slice[pipeIndexList[0]+2:pipeIndexList[1]]...)
                cmd_three := exec.Command(slice[pipeIndexList[1] + 1],
                                        slice[pipeIndexList[1]+2:pipeIndexList[2]]...)
                cmd_four := exec.Command(slice[pipeIndexList[2] + 1],
                                          slice[pipeIndexList[2]+2:]...)
                fmt.Println(slice[0],slice[1:pipeIndexList[0]])
                fmt.Println(slice[pipeIndexList[0] + 1],
                            slice[pipeIndexList[0]+2:pipeIndexList[1]])
                fmt.Println(slice[pipeIndexList[1] + 1],
                            slice[pipeIndexList[1]+2:pipeIndexList[2]])
                fmt.Println(slice[pipeIndexList[2] + 1],slice[pipeIndexList[2]+2:])

                cmd_two.Stdin, _ = cmd_one.StdoutPipe()
                cmd_three.Stdin, _ = cmd_two.StdoutPipe()
                cmd_four.Stdin, _ = cmd_three.StdoutPipe()
                cmd_four.Stdout = os.Stdout

                cmd_one.Start()
                cmd_two.Start()
                cmd_three.Start()
                cmd_four.Start()
                cmd_one.Wait()
                cmd_two.Wait()
                cmd_three.Wait()
                cmd_four.Wait()
            }

//            After some consideration, I finally figured out what to do, and
//            will be implementing a slice of exec commands to operate on next
//            This will be done sometime in the near future

/*            var end = 0
 *            var args = 0
 *            reader, writer := io.Pipe()
 *            for index, pipeIndex := range(pipeIndexList) {
 *
 *                if end < len(pipeIndexList) - 1{
 *                    end = pipeIndexList[index + 1]
 *                } else {
 *                    end = len(slice)
 *                }
 *                args = pipeIndex + 2
 *                if args > len(slice) - 1 {
 *                    args = end
 *                }
 *
 *                command := exec.Command(slice[pipeIndex + 1], slice[args:end]...)
 *                command.Stdout = writer
 *                if index > 0 {
 *                    command.Stdin = reader
 *                }
 *                if index == len(pipeIndexList) - 1 {
 *                    command.Stdout = os.Stdout
 *                }
 *                err = command.Start()
 *                err = command.Wait()
 *                err = writer.Close()
 *                fmt.Println("reader:", reader, "writer:", writer)
 *                if index > 0 {
 *                    reader.Close()
 *                }
 *
 *                if err != nil {
 *                    break
 *                }
 *           }
 */
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
        if field == "&" {
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
