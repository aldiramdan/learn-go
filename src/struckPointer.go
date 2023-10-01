package src

type bilang interface {
	Prima() []int
	Ganjil() []int
	Genap() []int
	Fibonacci() []int
}

type Deret interface {
	bilang
}

type Number struct {
	Limit *int
}

func (n *Number) Prima() []int {
	var result []int
	for i := 0; i <= int(*n.Limit); i++ {
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
	return result
}

func (n *Number) Ganjil() []int {
	var result []int
	for i := 0; i <= int(*n.Limit); i++ {
		if i%2 != 0 {
			result = append(result, i)
		}
	}
	return result
}

func (n *Number) Genap() []int {
	var result []int
	for i := 0; i <= int(*n.Limit); i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	return result
}

func (n *Number) Fibonacci() []int {
	var result []int
	var previous, current, next int = 0, 1, 1
	for i := 0; i < int(*n.Limit); i++ {
		result = append(result, previous)
		next = previous + current
		previous = current
		current = next
	}
	return result
}
