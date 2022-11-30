package helpers

import (
	"fmt"
	"time"
)

// func to calculate and print execution time
func ExeTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s execution time: %v\n", name, time.Since(start))
	}
}
