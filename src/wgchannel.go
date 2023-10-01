package src

import (
	"fmt"
	"sync"
)

func WgChannel() {
	a := make(chan []int)
	c := 20

	wg.Add(2)

	go deretFibonacci(&c, a, &wg)
	mx.Lock()
	go filterGanjilGenap(a, &wg)

	wg.Wait()
}

func deretFibonacci(n *int, ch chan<- []int, wg *sync.WaitGroup) {
	defer func() {
		mx.Unlock()
		wg.Done()
	}()
	var result []int
	var previous, current, next int = 0, 1, 1
	for i := 0; i < int(*n); i++ {
		result = append(result, previous)
		next = previous + current
		previous = current
		current = next
	}
	ch <- result
	fmt.Println("Fibonacci	:", result)
}

func filterGanjilGenap(ch <-chan []int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	var resultGanjil, resultGenap []int
	a := <-ch
	for _, data := range a {
		if data%2 != 0 {
			resultGanjil = append(resultGanjil, data)
		} else {
			resultGenap = append(resultGenap, data)
		}
	}
	fmt.Println("Fib Ganjil	:", resultGanjil)
	fmt.Println("Fib Genap	:", resultGenap)
}
