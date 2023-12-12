package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please enter your name")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	//The strings.TrimSpace function removes any space characters,
	//including new lines, from the start and end of a string.
	fmt.Printf("Hi, %s! I'm GO!", name)
	/*
		Everything here
		will be considered
		a block comment
	*/
	a := `Say "hello" to Go!`
	fmt.Println(a)

	b := `This string is on
	multiple lines
	within a single back
	quote on either side.`
	fmt.Println(b)

}
