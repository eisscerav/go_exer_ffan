package helper

import "fmt"

// Name of function should start with capital letter when using as module in go
func MyPrint(name string, age int) {
	fmt.Printf("Hello %v you are %v\n", name, age)
}
