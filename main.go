package main

import (
	"fmt"

	"github.com/aldiramdan/golang-introduction/src"
)

func main() {
	// task 1
	src.RevPyramid(5)
	src.Pyramid(5)

	// task 2
	q := src.GenPass("abcd", "strong")
	fmt.Printf("Generate Password : %s \n", q)

	w := src.PassGen("aldi", "strong")
	fmt.Printf("Generate Password : %s \n", w)

	// task 3
	var flight int = 7
	var data = []int{1, 7, 3, 4}
	e := src.FlightRange(flight, data)
	fmt.Println(e)

	r := src.FlightLoop(flight, data)
	fmt.Println(r)
}
