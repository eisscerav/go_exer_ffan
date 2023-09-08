package main

import (
	"fmt"
)

func main()  {
	name := "ffan"
	var age int = 40
	const total int = 100
	var books [10]string // define array
	var s []int
	books[0] = "cuda"
	s = append(s, 100)  // demo append a slice
	var length int = len(books)
	fmt.Println("hello go")
	fmt.Printf("name = %v, age = %v, length = %v\n", name, age, length)
}
