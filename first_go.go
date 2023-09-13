package main

import (
	"fmt"
	"hello_go/helper" // import other module file
	"math/rand"
	"strconv"
)

func greeting(first_name string, last_name string) (int, int) {
	fmt.Printf("hello %v %v\n", first_name, last_name)
	return 1, 3
}

func main() {
	name := "ffan" // declare var in shortcut version
	name = getName()
	var age int = rand.Intn(100)
	var age_2 int = rand.Intn(100)
	const total int = 100
	var room = helper.MyVar
	room += 1
	// map in go
	var user = make(map[string]string)
	user["ffan"] = strconv.Itoa(age)
	user["dili"] = strconv.Itoa(age_2)
	var books [10]string                     // define array
	var s []int                              // slice
	var users = make([]map[string]string, 0) // list of maps
	users[0]["ffan"] = "good"
	books[0] = "ffan"
	books[1] = "cuda"
	s = append(s, 100) // demo append a slice
	var length int = len(books)
	helper.MyPrint(name, age)
	if age > 30 {
		// do something
	} else {
		//do other things
		fmt.Println("age < 30")
	}
	//for _, book := range books {  // placeholder symbol _
	for index, book := range books { // loop in go
		fmt.Printf("index: %v, book: %v\n", index, book)
	}
	greeting(books[0], books[1])
	fmt.Println("hello go")
	fmt.Printf("name = %v, age = %v, length = %v\n", name, age, length)

	return
}
