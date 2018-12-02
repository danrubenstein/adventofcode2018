package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main () {

    freqs := make(map[int]bool)
    inputs := make([]int, 0)
    frequency := 0
    freqs[frequency] = true

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {

        text := scanner.Text()
        if strings.Compare(text, "search") == 0 {
            fmt.Println("Ok!")
            break
        }
        i, err := strconv.Atoi(text)

        if err != nil {
            fmt.Println("There was an error reading the number", err)
            continue
        }

        inputs = append(inputs, i)

        frequency += i
        _, ok := freqs[frequency]
        if ok {
            fmt.Printf("We've seen %d as a frequency before.\n", frequency)
            return
        }

        freqs[frequency] = true
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("There was an error reading the input", err)
        os.Exit(1)
    }

    for {
        for _, element := range inputs {
            frequency += element
            if _, ok := freqs[frequency]; ok {
                fmt.Printf("We've seen %d as a frequency before.\n", frequency)
                return
            }
            freqs[frequency] = true
        }
        fmt.Println("Looping again!")
    }
}
