package main

import "fmt"

func check_type(i interface{}) {
    switch v_type, ok := i.(type); ok {
    case string:
        fmt.Println("string")
    default:


    }
}
func main() {

}
