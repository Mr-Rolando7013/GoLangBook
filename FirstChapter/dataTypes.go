package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//Convert data types
	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)
	fmt.Println(bigIndex)

	var big int64 = 64
	var little int8
	little = int8(big)
	fmt.Println(little)

	//Convert integer to floats
	var x int64 = 57
	var y float64 = float64(x)
	fmt.Printf("%.2f\n", y)

	//Convert float to integers
	var f float64 = 390.8
	var i int = int(f)
	fmt.Printf("f = %.2f\n", f)
	fmt.Printf("i = %d\n", i)

	//Convert numbers to string
	a := strconv.Itoa(12)
	fmt.Printf("%q\n", a)

	user := "Sammy"
	lines := 50
	fmt.Println("Congratulations, " + user + "! You just wrote " +
		strconv.Itoa(lines) + " lines of code.")

	//Float to string
	fmt.Println(fmt.Sprint(421.034))
	t := 5524.53
	fmt.Println(fmt.Sprint(t))

	//Convert strings to number
	lines_yesterday := "50"
	lines_today := "108"
	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
	}
	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}
	lines_more := today - yesterday
	fmt.Println(lines_more)

	z := "not a number"
	b, err := strconv.Atoi(z)
	fmt.Println(b)
	fmt.Println(err)

	//Strings and bytes
	o := "my string"
	p := []byte(o)
	c := string(p)
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(c)
}
