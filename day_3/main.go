package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var data []byte

func main() {
	var total int
	// We know that each line of the input is 101 characters long (including the newline).
	// So we can cheat a bit and skip checking for a new line and avoid a
	// conditional branch. It's a trivial change, but if we can make the code as
	// predictable as possible, the compiler can optimise more effectively.
	for off := 0; off+100 <= len(data); off += 100 + 1 {
		total += TopKDigitsOrdered(data[off : off+100])
	}
	fmt.Println("Solution:", total)
}

// TopKDigitsOrdered acts as a sort of min-heap, but I couldn't be bothered to write a min-heap,
// but I assume the overhead of the heap is more than this solution.
//
// This sort of question can usually be solved with a Top K elements algorithm, but the key
// difference here is we need to preserve the order of the digits.
func TopKDigitsOrdered(b []byte) int {
	n := len(b)
	// To prevent boundary checking, we can again cheat because we know that
	// for the gold solution we only need the top 12 digits. This is obviously
	// not the best practice, but I'm going for speed here.
	rem := n - 12
	var st [12]byte
	top := 0

	for i := 0; i < n; i++ {
		d := b[i]
		for rem > 0 && top > 0 && st[top-1] < d {
			top--
			rem--
		}
		if top < 12 {
			st[top] = d
			top++
		} else if rem > 0 {
			rem--
		}
	}

	var r int
	// Convert the stack to a single int
	// Example: st = [1, 2, 3] -> r = 123
	for i := 0; i < 12; i++ {
		r = r*10 + int(st[i]-'0')
	}
	return r
}
