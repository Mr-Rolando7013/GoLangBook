package main

//Package is a Go keyword that defines which code bundle this file
//belongs to. There can only be one package per folder

import (
	"fmt"
)

//The fmt package provides
//formatting and printing functions

func main() {
	fmt.Println("Hello, World!")

	//An array looks like this:
	coral := [3]string{"blue coral", "staghorn coral", "pillar coral"}
	fmt.Println(coral)

	/*

		A slice is an ordered sequence of elements that can change in length. Slices can increase their size dynamically.
		When you add new items to a slice, if the slice does not have enough memory to store the new items, it will
		request more memory from the system as needed. Because a slice can be expanded to add more elements when
		needed, they are more commonly used than arrays.

	*/
	seaCreatures := []string{"shark", "cuttlefish", "squid", "mantis shrimp"}
	fmt.Println(seaCreatures)

	seaCreatures = append(seaCreatures, "seahorse") //Append
	fmt.Println(seaCreatures)

}
