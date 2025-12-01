package main

import (
	"fmt"
	"os"
)

func main() {
	direct, indirect := solve("day_1/day_1.txt")
	fmt.Println(direct, indirect)
}

func solve(filename string) (int, int) {
	data, _ := os.ReadFile(filename)

	curPos := 50
	directZeroCount := 0
	indirectZeroCount := 0

	// We already know that the lock has 100 positions.
	// We can use this knowledge to skip actually creating the array
	// and operating on the array directly. Instead, we can just
	// use some math tricks to calculate the position of the lock
	// and simulate a circular array using modulo operations.
	const size = 100

	// We can also skip reading the file line by line and instead throw it in memory all at once.
	// This reduces the amount of time spent on disk I/O.
	//
	// We can treat the input as a single string and move the cursor forward as we need to perform
	// a single pass.
	for i := 0; i < len(data); {
		// At the start of each loop we can assume we have correctly moved the cursor forward
		// to the next instruction.
		// Example:
		// The first iteration
		//
		// L46\nR26\nL18\nR18
		// ^
		//
		// The second iteration after parsing the number and accounting for the newline would move
		// the cursor to the next character after the newline.
		//
		// L46\nR26\nL18\nR18
		//      ^
		dir := data[i]
		i++

		// Parse number
		num := 0
		for i < len(data) && data[i] >= '0' && data[i] <= '9' {
			num = num*10 + int(data[i]-'0')
			i++
		}

		// Skip the newline if present
		if i < len(data) && data[i] == '\n' {
			i++
		}

		var loops int

		// For the actual movement, we can use modulo arithmetic to simulate a circular array and
		// use branchless statements for a bit of flare.
		if dir == 'R' {
			// Moving right is a lot more simple than moving left, as we don't need to worry about
			// dealing with negative numbers.
			loops = (curPos + num - 1) / size
			curPos = (curPos + num) % size
		} else { // 'L'
			// Moving left is trickier, as we need to counteract Go division truncating towards zero.
			// We handle this by calculating two terms and subtracting them.
			// The first term (curPos-size)/size handles moving left from 0
			// The second term (curPos-num-size+1)/size handles a negative target
			//
			// Example
			// loops = (curPos-size)/size - (curPos-num-size+1)/size
			//       = (10-100)/100       - (10-18-100+1)/100
			//       = (-90)/100          - (-107)/100
			//       = 0                  - (-1) // Go division truncates towards zero
			//       = 1
			loops = (curPos-size)/size - (curPos-num-size+1)/size
			// Example
			// curPos = (curPos - num%size + size) % size
			//        = (10 - 18 + 100) % 100
			//        = (92) % 100
			//        = 92
			curPos = (curPos - num%size + size) % size
		}

		if curPos == 0 {
			directZeroCount++
		}
		indirectZeroCount += loops
	}

	// directZero count is the answer for part 1
	// indirectZeroCount is the answer for part 2
	return directZeroCount, indirectZeroCount + directZeroCount
}
