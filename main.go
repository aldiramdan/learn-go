package main

import (
	"fmt"

	"github.com/aldiramdan/golang-basic/src"
)

func main() {
	r := 4.37
	s := 4.32
	t := 4.324

	var u uint = 1
	fmt.Println("Output	Loop	:", src.RoundLooping(&r, &u))
	fmt.Println("Output	Loop	:", src.RoundLooping(&s, &u))
	fmt.Println("Output	Loop	:", src.RoundLooping(&t, &u))

	fmt.Println("Output	Float	:", src.RoundFloat(&r, &u))
	fmt.Println("Output	Float	:", src.RoundFloat(&s, &u))
	fmt.Println("Output	Float	:", src.RoundFloat(&t, &u))

	q := 20
	var deret src.Deret = &src.Number{Limit: &q}
	fmt.Println("Prima		:", deret.Prima())
	fmt.Println("Ganjil		:", deret.Ganjil())
	fmt.Println("Genap		:", deret.Genap())
	fmt.Println("Fibonacci	:", deret.Fibonacci())

	p := 6.
	var bangunRuang src.Hitung = &src.Kubus{Sisi: &p}
	fmt.Println("Luas		:", bangunRuang.Luas())
	fmt.Println("Keliling	:", bangunRuang.Keliling())
	fmt.Println("Volume		:", bangunRuang.Volume())
}
