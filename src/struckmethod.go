package src

import (
	"fmt"
	"sync"
)

type number struct {
	limit *int
}

var wg = sync.WaitGroup{}

func WgStruck() {
	b := 20
	var deret = number{limit: &b}

	wg.Add(4)

	go deret.Prima()
	go deret.Ganjil()
	go deret.Genap()
	go deret.Fibonacci()

	wg.Wait()
}

func (n *number) Prima() {
	defer wg.Done()
	var result []int
	for i := 0; i <= int(*n.limit); i++ {
		j := 0
		for k := 1; k <= i; k++ {
			if i%k == 0 {
				j++
			}
		}
		if (j == 2) && (i != 1) {
			result = append(result, i)
		}
	}
	fmt.Println("Prima		:", result)
}

func (n *number) Ganjil() {
	defer wg.Done()
	var result []int
	for i := 0; i <= int(*n.limit); i++ {
		if i%2 != 0 {
			result = append(result, i)
		}
	}
	fmt.Println("Ganjil		:", result)
}

func (n *number) Genap() {
	defer wg.Done()
	var result []int
	for i := 0; i <= int(*n.limit); i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	fmt.Println("Genap		:", result)
}

func (n *number) Fibonacci() {
	defer wg.Done()
	var result []int
	var previous, current, next int = 0, 1, 1
	for i := 0; i < int(*n.limit); i++ {
		result = append(result, previous)
		next = previous + current
		previous = current
		current = next
	}
	fmt.Println("Fibonacci	:", result)
}
