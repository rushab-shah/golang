package main

import "fmt"

func main() {
	fmt.Println("Checking how to use variables!")

	var test string = "Hello"
	test2 := "World!"

	fmt.Println(test + " " + test2)

	var flag bool = false
	opp_flag := true

	if flag != opp_flag {
		fmt.Println("Boolean values compared")
	}

	
}
