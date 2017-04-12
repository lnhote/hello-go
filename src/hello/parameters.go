package main

import "fmt"

func greeting(prefix string, names ...string) {
	fmt.Println(prefix)
	for _, name := range names {
		fmt.Print(name, ", ")
	}
}
func main() {
	greeting("Hello: ", "alice", "bob", "John")
}
