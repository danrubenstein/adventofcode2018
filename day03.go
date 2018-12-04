package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type claim struct {
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
		claims = append(claims, claim{b, c, d, e})

		totalArea += (d * e)
	}
	return claims, nil
}

func identifyOverlap(claims []claim) (int, error) {

	overlap := 0
	seen := make(map[int]int)
	for _, c := range claims {
		for j := 0; j < c.width; j++ {
			for i := 0; i < c.height; i++ {
				place := (c.distanceFromLeftEdge+j)*100000 + (c.distanceFromTopEdge + i)
				if val, ok := seen[place]; ok {
					seen[place] = val + 1
				} else {
					seen[place] = 1
				}
			}
		}
	}

	for _, v := range seen {
		if v > 1 {
			overlap++
		}
	}
	return overlap, nil
}
func main() {
	claims, err := readClaimsFromFile("tmp.txt")
	if err != nil {
		return
	}

	overlap, err := identifyOverlap(claims)

	fmt.Println("There is much overlap: ", overlap)
}
