package src

import "fmt"

func RevPyramid(nilai int) {
	for i := 0; i < nilai; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf(" ")
		}

		for k := 0; k < 2*(nilai-i)-1; k++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func Pyramid(nilai int) {
	for i := 1; i <= nilai; i++ {
		for j := 0; j < nilai-i; j++ {
			fmt.Printf(" ")
		}

		for k := 0; k < 2*i-1; k++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
