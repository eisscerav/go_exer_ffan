package helper

import "fmt"

// export variable
var MyVar = 40

// Name of function should start with capital letter when using as module in go
func MyPrint(name string, age int) {
	fmt.Printf("Hello %v you are %v years old\n", name, age)
}
