package main

import (
	"fmt"
)

func main() {

	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue"}

	fmt.Println(sammy)
	fmt.Println(sammy["color"])

	//Accessing map items
	fmt.Println(sammy["color"])

	//Iteration in a map
	for key, value := range sammy {
		fmt.Printf("%q is the key for the value %q\n", key, value)
	}

	//To get just the keys
	keys := []string{}
	for key := range sammy {
		keys = append(keys, key)
	}
	fmt.Printf("%q", keys)

	//Efficient way to retrieve items

	items := make([]string, len(sammy))
	var i int
	for _, v := range sammy {
		items[i] = v
		i++
	}
	fmt.Printf("%q", items)

	//Check existence of a key
	count, ok := counts["sammy"]
	if ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy was not found")
	}
}
