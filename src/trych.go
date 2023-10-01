package src

import "fmt"

func Try() {
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	c := make(chan int, 2)

	wg.Add(2)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	wg.Wait()
}

func sum(d []int, ch chan int) {
	defer wg.Done()
	result := 0
	for _, v := range d {
		result += v
	}
	ch <- result
	fmt.Println("Ouput		:", <-ch)
}
