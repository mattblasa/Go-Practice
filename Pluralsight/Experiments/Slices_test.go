package Experiments

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //find primes from index 1 to 4. 3 - 11
	fmt.Println(s)
}
