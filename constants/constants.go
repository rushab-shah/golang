package main

import (
	"fmt"
	"math"
)

const test string = "Hello World!"

func main() {
	fmt.Println("Testing how to use constants! " + test)

	const n = 3e10
	const d = 2e5
	r := n / d
	fmt.Printf("Division result float %f\n", r)
	fmt.Printf("Division result int %d\n", int64(r))

	fmt.Printf("Sin value %f\n", math.Sin(r))
}
