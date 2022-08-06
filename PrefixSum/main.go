package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	num := make([]int, n+1)
	pre := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&num[i])
		pre[i] = pre[i-1] + num[i]
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Println(pre[b] - pre[a-1])
	}
}
