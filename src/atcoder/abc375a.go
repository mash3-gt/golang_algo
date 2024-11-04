package main

import (
	"fmt"
)

func countSeats(S string) int {
	count := 0
	for i := 0; i < len(S)-2; i++ {
		if S[i] == '#' && S[i+1] == '.' && S[i+2] == '#' {
			count++
		}
	}
	return count
}

func main() {
	var N int
	var S string
	fmt.Scan(&N)
	fmt.Scan(&S)

	result := countSeats(S)
	fmt.Println(result)
}
