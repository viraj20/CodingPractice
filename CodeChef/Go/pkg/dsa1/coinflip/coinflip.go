package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var g int
		fmt.Scanf("%d", &g)
		for i := 0; i < g; i++ {
			var i, n, q int
			fmt.Scanf("%d %d %d", &i, &n, &q)
			if n%2 == 0 {
				fmt.Println(n / 2)
			} else {
				if q == 1 {
					if i == 1 {
						fmt.Println(n / 2)
					} else {
						fmt.Println((n / 2) + 1)
					}
				} else {
					if i == 2 {
						fmt.Println(n / 2)
					} else {
						fmt.Println((n / 2) + 1)
					}

				}
			}
		}
	}
}
