package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"sort"
	"strconv"
)

var partB = flag.Bool("partB", false, "Use Part B logic?")
var file = flag.String("input", "inputs/day9.txt", "what file to use")

type marble struct {
	val  int
	next *marble
	prev *marble
}

func main() {

	contents, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println("Error reading the file", err)
		return
	}

	expr := regexp.MustCompile("(\\d+) players; last marble is worth (\\d+) points")
	res := expr.FindSubmatch(contents)
	players, _ := strconv.Atoi(string(res[1]))
	lastMarble, _ := strconv.Atoi(string(res[2]))

	currentPlayer := 1
	currentMarble := &marble{0, nil, nil}
	currentMarble.next = currentMarble
	currentMarble.prev = currentMarble

	fmt.Println(currentMarble)

	scores := make([]int, players)

	for i := 1; i <= lastMarble; i++ {
		if math.Mod(float64(i), 23) == 0 {
			scores[currentPlayer] += i
			marbleToRemove := currentMarble
			for j := 0; j < 7; j++ {
				marbleToRemove = marbleToRemove.prev
			}

			marbleToRemove.prev.next = marbleToRemove.next
			marbleToRemove.next.prev = marbleToRemove.prev

			scores[currentPlayer] += marbleToRemove.val
			currentMarble = marbleToRemove.next
		} else {
			newMarble := &marble{i, currentMarble.next.next, currentMarble.next}
			currentMarble.next.next.prev = newMarble
			currentMarble.next.next = newMarble
			currentMarble = newMarble
		}
		currentPlayer = int(math.Mod(float64(currentPlayer+1), float64(players)))
	}

	sort.Ints(scores)
	fmt.Printf("%d\n", scores[len(scores)-1])
}
