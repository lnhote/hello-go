package main
import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s) // len = 4, cap = 6

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len = 4, cap = 4

    // s = s[0:6] // error
    // printSlice(s)

    // dynamically-sized arrays
    fmt.Println("===== dynamically-sized arrays =====")
    a := make([]int, 5)  // len(a)=5
    printSlice(a)

    b := make([]int, 5, 5) // len=5, cap=5
	printSlice(b)

    c := append(b, 2, 3, 4) // len=5+3=8, cap=5*2=10
    printSlice(c)

    // range
    fmt.Println("===== range =====")
    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
    for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
    for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
