package main

import (
    "bufio"
    "os"
    "fmt"
    "sort"
    "errors"
)

func findMutualLetters(str1 string, str2 string) string {

    fmt.Printf("The shared strings are %s and %s\n", str1, str2)
    joint := ""
    for i, _ := range str1 {

        if str1[i] == str2[i] {
            joint += string(str1[i])
        }
    }

    return joint
}
func findTheBox(ids []string, idLength int) (string, error){
    // This is doable in O(n^2) time.
    // Can we do it O(nlog(n)) time?
    // Sorting a list is n log n
    // And then from each operation
    // compare to next
    // mutual number of starting similar letters
    // mutual number of ending similar letters
    // add those two, if it is len(id) - 1, they are the pairs.
    // Then go row by row, and if the letters are the same, add them to the list
    // Return the list of similar letters

    sort.Strings(ids)
    n := len(ids)

    for i, _ := range ids[0:(n-1)] {

        var startSharedChars, endSharedChars int

        for j := 0; j < idLength; j++ {
            if ids[i][j] != ids[i+1][j] {
                startSharedChars = j
                break
            }
        }

        for k := 0; k < idLength; k++ {
            if ids[i][idLength-(k+1)] != ids[i+1][idLength-(k+1)] {
                endSharedChars = k
                break
            }
        }

        if (startSharedChars + endSharedChars) == (idLength - 1) {
            return findMutualLetters(ids[i], ids[i+1]), nil
        }
    }

    return "", errors.New("Couldn't find it :(")
}

func main () {

    scanner := bufio.NewScanner(os.Stdin)

    ids := make([]string, 0)
    for scanner.Scan() {

        line := scanner.Text()

        if line == "search!" {
            break
        }
        ids = append(ids, line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("There was an error reading the input: ", err)
        return
    }

    result, err := findTheBox(ids, len(ids[0]))

    if err != nil {
        fmt.Println("There was an error")
    }

    fmt.Println("The common letters are: ", result)


}
