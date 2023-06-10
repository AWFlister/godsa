package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello World!")

	// Variable definitions
	var x int = 5 // Full
	var y = 6     // Auto type
	z := 7        // Shorthand
	p, q := 3, 4  // Multiple values

	fmt.Println(p, q, x, y, z)

	// Arrays and Slices
	var a = [5]int{2, 3, 4, 5, 6}                     // Like new Array()
	b := []string{"Sakti", "Cantas", "Puan", "Kinar"} // Slice: dynamic allocation; vector
	c := a[0:3]                                       // Slice slicing, pythonic
	fmt.Println(a, b, c)

	// strings and sort modules
	d := []int{23, 4, 54, 6, 543, 5, 6543, 76, 879, 9, 854, 56}
	b = append(b, "Sigit", "Nina", "Sri", "Bambang")
	sort.Ints(d)
	sort.Strings(b)
	fmt.Println(d)
	fmt.Println(b)
}
