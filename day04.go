package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

type duration struct {
	start time.Time
	end   time.Time
	len   int
}

// mostCommonSleepInterval gets the most common minute
// to be asleep for a given list of naps
func mostCommonSleepInterval(naps []duration) (int, int) {

	var frequencies [60]int
	for _, n := range naps {
		for j := n.start.Minute(); j < n.end.Minute(); j++ {
			frequencies[j]++
		}
	}

	modeMinute := 0
	valMinute := 0
	for i, val := range frequencies {
		if val > valMinute {
			valMinute = val
			modeMinute = i
		}
	}

	return modeMinute, valMinute
}

// findMostReliableGuard gets the guard who is most frequently
// asleep during a certain minute period
func findMostReliableGuard(napProfiles map[int][]duration) int {
	mostCommonNapTimes := make(map[int]int)
	for k, v := range napProfiles {
		_, res := mostCommonSleepInterval(v)
		mostCommonNapTimes[k] = res
	}

	var mostReliableGuard int
	var mostReliableFreq int

	for k, v := range mostCommonNapTimes {
		if v > mostReliableFreq {
			mostReliableGuard = k
			mostReliableFreq = v
		}
	}

	return mostReliableGuard
}

func main() {
	dat, err := ioutil.ReadFile("day4.txt")

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	sleepPeriods := make(map[int][]duration)

	lines := strings.Split(string(dat), "\n")
	sort.Strings(lines)

	var currentGuard int
	var year, month, day, hr, min int
	var status string
	var start, end time.Time

	for _, line := range lines {
		if strings.Contains(line, "begins shift") {
			fmt.Sscanf(line, "[%d-%d-%d %d:%d] Guard #%d begins shift", &year, &month, &day, &hr, &min, &currentGuard)
		} else {
			fmt.Sscanf(line, "[%d-%d-%d %d:%d] %s", &year, &month, &day, &hr, &min, &status)

			if strings.Contains(line, "wakes up") {
				end = time.Date(year, time.Month(month), day, hr, min, 0, 0, time.UTC)
				newPeriod := duration{start, end, int(end.Sub(start).Minutes())}
				if val, ok := sleepPeriods[currentGuard]; ok {
					sleepPeriods[currentGuard] = append(val, newPeriod)
				} else {
					sleepPeriods[currentGuard] = []duration{newPeriod}
				}
			} else {
				start = time.Date(year, time.Month(month), day, hr, min, 0, 0, time.UTC)
			}
		}
	}

	tiredGuard := 0
	highestValue := 0
	for k, v := range sleepPeriods {
		sum := 0
		for _, p := range v {
			sum += p.len
		}

		if sum > highestValue {
			highestValue = sum
			tiredGuard = k
		}
	}

	sleepiestGuardMostCommonMinute, sleepiestGuardCountTimes := mostCommonSleepInterval(sleepPeriods[tiredGuard])

	fmt.Println("The most common minute to be asleep at is:", sleepiestGuardMostCommonMinute)
	fmt.Println("The times the guard was asleep then was:", sleepiestGuardCountTimes)

	fmt.Println("The answer to strategy 1 is:", sleepiestGuardMostCommonMinute*tiredGuard)

	mostReliableGuard := findMostReliableGuard(sleepPeriods)

	mostReliableMin, _ := mostCommonSleepInterval(sleepPeriods[mostReliableGuard])
	fmt.Println("The most frequent time any one guard was asleep was:", mostReliableMin)
	fmt.Println("The answer to strategy 2 is:", mostReliableMin*mostReliableGuard)
}
