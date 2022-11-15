package main

import (
	"fmt"
	"time"
)

// If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.
//
// {20,48,52}, {24,45,51}, {30,40,50}
//
// For which value of p ≤ 1000, is the number of solutions maximised?

const MIN_SIDE_LENGTH = 3
const MIN_PERIMETER = 3 + 4 + 5

func solutionGivenP(p int) int {
	count := 0
	for a := MIN_SIDE_LENGTH; a < p/2; a++ {
		for b := a; b < p/2; b++ {
			c := p - a - b
			a2 := a * a
			b2 := b * b
			c2 := c * c
			if a2+b2 == c2 {
				count++
			} else if a2+b2 > c*c {
				break
			}
		}
	}
	return count
}

func solution1() (int, int) {
	var maxCount int
	var maxP int
	for p := MIN_PERIMETER; p <= 1000; p++ {
		count := solutionGivenP(p)
		if count > maxCount {
			maxCount = count
			maxP = p
		}
	}
	return maxP, maxCount
}

func main() {
	start := time.Now()
	P, count := solution1()
	elapse := time.Since(start)

	fmt.Printf("P=%d, count=%d, elapse=%s\n", P, count, elapse)
}
