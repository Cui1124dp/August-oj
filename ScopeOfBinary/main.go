package main

import "fmt"

func scopeOfBinary(num []int, a int) (int, int) {
	l, r := 0, len(num)-1
	for l < r {
		mid := (l + r) >> 1
		if num[mid] < a {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if num[l] != a {
		return -1, -1
	}
	re1 := l
	r = len(num) - 1
	for l < r {
		mid := (l + r + 1) >> 1
		if num[mid] <= a {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return re1, l

}

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	num, que := make([]int, n), make([]int, q)
	for i := range num {
		fmt.Scan(&num[i])
	}
	for i := range que {
		fmt.Scan(&que[i])
	}
	for _, i := range que {
		fmt.Println(scopeOfBinary(num, i))
	}
}
