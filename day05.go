package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

// H/T Liz Fong-Jones' (@lizthegrey) livestreams for teaching me about flags.
var partB = flag.Bool("partb", false, "whether to use part b logic")
var file = flag.String("file", "day5.txt", "What file to point to")

func findLengthCollapseChain(s string) int {
	var cur rune
	loose := false
	chain := make([]rune, 0)
	for _, val := range s {
		if loose == false {
			cur = val
			loose = true
			// unicode codepoints ftw!
		} else if math.Abs(float64(val)-float64(cur)) == 32 {
			if len(chain) == 0 {
				loose = false
			} else {
				cur = chain[len(chain)-1]
				chain = chain[:len(chain)-1]
			}
		} else {
			chain = append(chain, cur)
			cur = val
		}
	}
	return len(chain)
}
func main() {
	flag.Parse()
	s, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if *partB != true {
		res := findLengthCollapseChain(string(s))
		fmt.Println("The answer is:", res)
	} else {
		collapsedLengths := make(map[rune]int)
		for i := 0; i < 26; i++ {
			char := rune('A' + i)
			replacedStringUpper := strings.Replace(string(s), string(char), "", -1)
			replacedStringLower := strings.Replace(string(replacedStringUpper), string(rune('A'+i+32)), "", -1)
			collapsedLengths[char] = findLengthCollapseChain(replacedStringLower)
		}

		lowest := len(s)
		for _, v := range collapsedLengths {
			if v < lowest {
				lowest = v
			}
		}
		fmt.Println("The answer is:", lowest)
	}
}
