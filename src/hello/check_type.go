package main

import "fmt"

func check_type(i interface{}) {
    switch v_type := i.(type) {
    case string:
        fmt.Printf("string %v\n", v_type)
    case int:
        fmt.Printf("int %d\n", v_type)
    case bool:
        fmt.Printf("bool %v\n", v_type)
    case float64:
        fmt.Printf("float64 %v\n", v_type)
    default:
        fmt.Println("wtf?\n")

    }
}
func main() {
    var i interface{} = "Hunter"
    val, ok := i.(string)
    fmt.Printf("%s %v\n", val, ok)
    val2, ok := i.(int)
    fmt.Printf("%d %v\n", val2, ok)
    check_type("Alice")
    check_type(123)
}
