package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	customersBudget := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &customersBudget[i])
	}

	sort.Ints(customersBudget)
	var max int
	var length int = len(customersBudget)
	for _, v := range customersBudget {
		if max < v*length {
			max = v * length
		}
		length--
	}
	fmt.Println(max)
}
