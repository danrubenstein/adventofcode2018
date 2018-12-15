package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

var partB = flag.Bool("partB", false, "Use part B logic")
var file = flag.String("file", "day7.txt", "What file to use")

type depMap map[string]map[string]bool

func findDependencies(instructions string) depMap {
	expr := regexp.MustCompile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")
	matches := expr.FindAllStringSubmatch(instructions, -1)
	dependencies := make(depMap)
	for _, val := range matches {
		upstream := val[1]
		downstream := val[2]

		if _, ok := dependencies[downstream]; ok {
			dependencies[downstream][upstream] = true
		} else {
			dependencies[downstream] = map[string]bool{upstream: true}
		}

		if _, ok := dependencies[upstream]; !ok {
			dependencies[upstream] = map[string]bool{}
		}
	}
	return dependencies
}

// getNoDeps returns a sorted list from a dependency map
func getNoDeps(deps depMap) []string {
	options := make([]string, 0)
	for k, v := range depMap {
		if len(v) == 0 {
			options = append(options, k)
		}
	}
	if len(options) == 0 {
		break allocating
	}
	sort.Strings(options)

	return options
}

func main() {

	flag.Parse()

	contents, err := ioutil.ReadFile(*file)

	if err != nil {
		fmt.Println("Error reading the file:", err)
	}

	dependencies := findDependencies(string(contents))
	if !*partB {
		list := ""
	outer:
		for {
			options := getNoDeps(dependencies)

			if len(options) == 0 {
				break outer
			}

			next := options[0]
			delete(dependencies, next)

			for k, v := range dependencies {
				delete(v, next)
				dependencies[k] = v
			}

			list = strings.Join([]string{list, next}, "")
		}
	} else {
		activeWork := make(map[string]int64)
		time := 0
	partBOuter:
		for {

			fmt.Println(activeWork)
			for k, v := range activeWork {
				activeWork[k] = v - 1
				// Remove the blockers freed up by this works
				if activeWork[k] == 0 {
					list = strings.Join([]string{list, k}, "")
					delete(activeWork, k)
					for kD, vD := range dependencies {
						delete(vD, k)
						dependencies[kD] = vD
					}
					delete(activeWork, k)
				}
			}

			// Hard code the things
			availableSlots := 5 - len(activeWork)

		allocating:
			for i := 0; i < availableSlots; i++ {
				options := getNoDeps(dependencies)
				if len(options) == 0 {
					break allocating
				}
				next := options[0]
				delete(dependencies, next)
				activeWork[next] = int64(next[0]) - int64('A') + 61
			}
			if len(activeWork) == 0 {
				break partBOuter
			}
			time++
		}
		fmt.Println(time)
	}
}
