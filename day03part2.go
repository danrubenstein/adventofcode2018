package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type claim struct {
	id                   int
	distanceFromLeftEdge int
	distanceFromTopEdge  int
	width                int
	height               int
}

// readClaimsFromFile reads an input string and factors the
// lines into elf claims on fabric.
func readClaimsFromFile(filename string) ([]claim, error) {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)

	totalArea := 0
	claims := make([]claim, 0)
	for s.Scan() {
		line := s.Text()
		var (
			a, b, c, d, e int
		)
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &a, &b, &c, &d, &e)
		claims = append(claims, claim{a, b, c, d, e})

		totalArea += (d * e)
	}
	return claims, nil
}

func identifySingleClaim(claims []claim) int {

	seen := make(map[int]int)
	overlaps := make(map[int]bool)
	for _, c := range claims {
		for j := 0; j < c.width; j++ {
			for i := 0; i < c.height; i++ {
				place := (c.distanceFromLeftEdge+j)*100000 + (c.distanceFromTopEdge + i)
				if val, ok := seen[place]; ok {
					overlaps[val] = true
					overlaps[c.id] = true
				} else {
					seen[place] = c.id
				}
			}
		}
	}

	for _, c := range claims {
		if _, ok := overlaps[c.id]; !ok {
			return c.id
		}
	}
	return -1
}
func main() {
	claims, err := readClaimsFromFile("tmp.txt")
	if err != nil {
		return
	}

	claim := identifySingleClaim(claims)

	fmt.Println("This is the lone claim!", claim)
}
