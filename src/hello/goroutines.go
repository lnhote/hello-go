package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
    fmt.Println("start")
	go say("1")
	go say("2")
	go say("3")
	go say("4")
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("end")
}
