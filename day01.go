package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main ()  {
    frequency := 0

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        i, err := strconv.Atoi(line)

        if err != nil {
            fmt.Fprintln(os.Stderr, "could not process number:", line)
        }

        frequency += i;

        fmt.Printf("The current frequency is %d, after a change of %d\n", frequency, i);
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}
