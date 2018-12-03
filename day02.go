package main

import (
    "fmt"
    "os"
    "bufio"
)

func main () {

    scanner := bufio.NewScanner(os.Stdin)

    twos := 0
    threes := 0

    for scanner.Scan() {
        line := scanner.Text()

        if line == "compute" {
            break
        }
        freqs := make(map[rune]int)

        for _, char := range line {

            if val, ok := freqs[char]; ok {
                freqs[char] = val + 1
            } else {
                freqs[char] = 1
            }
        }

        for _, v := range freqs {
            if v == 2 {
                twos++
                break
            }
        }

        for _, v := range freqs {
            if v == 3 {
                threes++
                break
            }
        }
    }

    fmt.Println("The checksum is: ", twos * threes)


    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input: ", err)
    }
}
