package src

import (
	"fmt"
	"sync"
)

var mx = sync.Mutex{}

func WgBChannel() {
	a := make(chan []int, 3)
	b := 20

	wg.Add(3)

	mx.Lock()
	go dertGanjil(&b, a, &wg)
	mx.Lock()
	go deretGenap(&b, a, &wg)
	mx.Lock()
	go deretPrima(&b, a, &wg)

	for v := range a {
		fmt.Println(v)
	}
	wg.Wait()
}

func dertGanjil(n *int, ch chan []int, wg *sync.WaitGroup) {
	defer func() {
		mx.Unlock()
		wg.Done()
	}()

	var resultGanjil []int
	for i := 0; i < int(*n); i++ {
		if i%2 != 0 {
			resultGanjil = append(resultGanjil, i)
		}
	}
	ch <- resultGanjil
}

func deretGenap(n *int, ch chan []int, wg *sync.WaitGroup) {
	defer func() {
		mx.Unlock()
		wg.Done()
	}()

	var resultGenap []int
	for i := 0; i < int(*n); i++ {
		if i%2 == 0 {
			resultGenap = append(resultGenap, i)
		}
	}
	ch <- resultGenap
}

func deretPrima(n *int, ch chan []int, wg *sync.WaitGroup) {
	defer func() {
		mx.Unlock()
		close(ch)
		wg.Done()
	}()
	var resultPrima []int
	for i := 0; i <= int(*n); i++ {
		j := 0
		for k := 1; k <= i; k++ {
			if i%k == 0 {
				j++
			}
		}
		if (j == 2) && (i != 1) {
			resultPrima = append(resultPrima, i)
		}
	}
	ch <- resultPrima
}
