package main

import (
	"fmt"
	"time"
    "math"
)

// MyError implemented error
type MyError struct {
	When time.Time
	What string
}
func (e *MyError) Error() string {
	return fmt.Sprintf("[ERROR][%v] %s", e.When, e.What)
}
func throwError(msg string) error {
	return &MyError{ time.Now(), msg }
}
func throwBadAssError() error {
	return throwError("Wtf??? It didn't work.")
}

// ErrNegativeSqrt implemented error
type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
    msg := fmt.Sprintf("Cannot Sqrt negative number: %f", float64(e))
    return throwError(msg).Error()
}

func Sqrt(x float64) (float64, error) {
    if x >= 0 {
        return math.Sqrt(x), nil
    } else {
        return 0, ErrNegativeSqrt(x)
    }
}

func main() {
	if err := throwBadAssError(); err != nil {
		fmt.Println(err)
	}
    result, err := Sqrt(2);
    if err == nil {
        fmt.Println("Success. Sqrt of", 2, "is", result)
    } else {
        fmt.Println(err)
    }
    result, err = Sqrt(-2);
    if err == nil {
        fmt.Println("Success. Sqrt of", -2, "is", result)
    } else {
        fmt.Println(err)
    }
}
