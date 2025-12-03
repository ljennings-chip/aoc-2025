package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	solve("day_2/input.txt")
}

func solve(filename string) {
	data, _ := os.ReadFile(filename)

	start := time.Now()

	var silverTotal int64
	var goldTotal int64

	lastPos := 0
	for i := 0; i <= len(data); {
		if i < len(data) && data[i] != ',' {
			i++
			continue
		}

		s := string(data[lastPos:i])
		t := strings.Split(s, "-")

		start, _ := strconv.ParseInt(t[0], 10, 64)
		end, _ := strconv.ParseInt(t[1], 10, 64)

		for i := start; i <= end; i++ {
			id := strconv.FormatInt(i, 10)
			if isTwiceRepeating(id) {
				silverTotal = silverTotal + i
			}
			if isPeriodic(id) {
				goldTotal = goldTotal + i
			}
		}

		lastPos = i + 1
		i++
	}

	fmt.Println("Time elapsed:", time.Since(start))
	fmt.Println(silverTotal, goldTotal)
}

// isPeriodic uses a neat trick where we can confirm if a string is periodic
// by checking if the base string is contained within the middle of a combined
// string.
//
// If we strip the first and last characters from the combined string, we can
// check if the base string is contained within the middle, which confirms
// that the string is periodic.
//
// Example:
// s = "6464"
// combined = "64646464"
// substring = "464646"
// strings.Contains(middle, s) == true
func isPeriodic(s string) bool {
	combined := s + s
	middle := combined[1 : len(combined)-1]
	return strings.Contains(middle, s)
}

// isTwiceRepeating checks if a string is twice repeated within itself.
//
// We can confirm this intuitively by checking if the first half of the string
// is equal to the second half.
func isTwiceRepeating(id string) bool {
	n := len(id)
	if n%2 != 0 {
		return false
	}

	half := n / 2
	return id[:half] == id[half:]
}
