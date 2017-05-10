package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func add2(a, b int) int {
	return a + b
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var i, j int = 1, 2

// var (
//     x int = 1
//     y int = 2
//     z int = 3
// )
func main() {
	fmt.Println(add(7, 6))
	fmt.Println(add2(7, 6))
	a, b := swap("World", "Hello")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(c, python, java)
	fmt.Println(i, j)
	v := 42 // change me!
	fmt.Printf("v is of type %T, value = %v\n", v, v)
	// fmt.Println(x, y, z)
}
