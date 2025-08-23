package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
	//Output: [{<nil> 2} {0xc0000a8040 865} {0xc0000b0010 236} {0xc0000aa0e0 864} {0xc0000aa0f0 342} {0xc0000b2040 247}]
}

func largestPrimeFactor(n int) int {
	var possiblePrime int
	debug := 0
OuterLoop:
	for i := 2; i < n; i++ {
		if n%i == 0 {
			possiblePrime = n / i
			for j := 2; j < int(math.Sqrt(float64(possiblePrime))); j++ {
				if possiblePrime%j == 0 {
					continue OuterLoop
				}
			}
			return possiblePrime
		}
	}
	return debug
}
