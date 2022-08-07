package main

import "fmt"

func main() {
	var n, m, q int
	fmt.Scanf("%d%d%d", &n, &m, &q)
	matrix := [1010][1010]int{}
	pre := [1010][1010]int{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scanf("%d", &matrix[i][j])
			pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + matrix[i][j]
		}
	}
	for i := 0; i < q; i++ {
		var x1, y1, x2, y2 int
		fmt.Scanf("%d%d%d%d", &x1, &y1, &x2, &y2)
		fmt.Printf("%d\n", pre[x2][y2]-pre[x1-1][y2]-pre[x2][y1-1]+pre[x1-1][y1-1])
	}

}
