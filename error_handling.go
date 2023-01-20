package main

import (
	"fmt"
)

// func panic(v interface{})
// func recover() interface

func p(s string, i int) {
	if i > 2 {
		panic(s)
	}
}

func r() {
	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println("Recovered from panic")
	}
}

func main() {

	defer r()
	p("runtime error: enter panic state", 3)

	fmt.Println("End!")
}
