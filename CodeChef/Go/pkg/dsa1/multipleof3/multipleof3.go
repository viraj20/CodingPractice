package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n, d0, d1 int
		fmt.Scanf("%d %d %d", &n, &d0, &d1)
		sum := d0
		sum = sum + d1
		n = n - 2
		for n != 0 {
			n--
			if sum%10 == 8 {
				sum = sum + (sum % 10)
				break
			} else if sum%10 == 0 {
				break
			} else {
				sum = sum + (sum % 10)
			}
		}
		if n != 0 && sum%10 != 0 {
			sum = sum + (20 * (n / 4))
			if n%4 == 1 {
				sum = sum + 6
			} else if n%4 == 2 {
				sum = sum + 8
			} else if n%4 == 3 {
				sum = sum + 12
			}
		}
		if sum%3 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
