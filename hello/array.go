package main

import "fmt"

func main() {
	fmt.Println("===== array =====")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// This is an array literal
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println("===== slice =====")
	// slice is just reference, it doesn't store any data
	// start, end
	slice := primes[1:4]
	fmt.Println(slice)

	slice2 := primes[0:2]
	fmt.Println(slice2)

	slice[0] = 17
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(primes)

	// slice literal
	// And dog creates the same array as above, then builds a slice that references it:
	q := []int{1, 3, 5, 7, 9}
	fmt.Println(q)

	users := []struct {
		i    int
		name string
	}{
		{100, "Alice"},
		{101, "Bob"},
		{102, "John"},
	}
	fmt.Println(users)
}
