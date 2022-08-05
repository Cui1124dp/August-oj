package main

import "fmt"

func quickSort(q []int, l int, r int) {
	if l >= r {
		return
	}
	i, j, x := l-1, r+1, q[(l+r)>>1]
	for i < j {
		i++
		for q[i] < x {
			i++
		}
		j--
		for q[j] > x {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

func main() {
	var n int
	fmt.Scan(&n)
	q := make([]int, n)
	for i, _ := range q {
		fmt.Scan(&q[i])
	}
	quickSort(q, 0, n-1)
	for _, i := range q {
		fmt.Print(i, " ")
	}
}
