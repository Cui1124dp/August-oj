package main

import "fmt"

func margeSort(q []int, l int, r int) {

	if l >= r {
		return
	}
	mid := (l + r) >> 1
	margeSort(q, l, mid)
	margeSort(q, mid+1, r)
	temp := make([]int, r-l+1)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] < q[j] {
			temp[k] = q[i]
			k++
			i++
		} else {
			temp[k] = q[j]
			k++
			j++
		}
	}
	for i <= mid {
		temp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		temp[k] = q[j]
		k++
		j++
	}
	j = 0
	for i := l; i <= r; i++ {
		q[i] = temp[j]
		j++
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	q := make([]int, n)
	for i := range q {
		fmt.Scan(&q[i])
	}
	margeSort(q, 0, n-1)
	for _, i := range q {
		fmt.Print(i, " ")
	}

}
