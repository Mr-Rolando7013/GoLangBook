package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "Rolando Enriquez"

	fmt.Println(strings.ToLower(ss))
	fmt.Println(strings.ToUpper(ss))

	//strings.HasPrefix Searches the string from the begging.
	//strings.HasSuffix Searches the string from the end.
	//strings.Contains Searches anywhere in the string.
	//strings.Count Count how many time the string appears.
	fmt.Println(strings.Contains(ss, "Rola"))
	fmt.Println(len(ss))

	//Join is to combine single strings into one.
	fmt.Println(strings.Join([]string{"sharks", "crustaceans", "plankton"}, ","))

	//Split funciton.
	balloon := "Sammy has a balloon."
	s := strings.Split(balloon, " ")
	fmt.Println(s)

	//To properly see the output we gotta do:
	fmt.Printf("%q", s)

	//Fields is similar to split but will only ignore all whitespaces.
	data := " username password email date"
	fields := strings.Fields(data)
	fmt.Printf("%q", fields)

	//Replace all function
	fmt.Println(strings.ReplaceAll(balloon, "has", "had"))

}
