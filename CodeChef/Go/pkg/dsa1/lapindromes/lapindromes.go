package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var countArray [26]int
		var word string
		fmt.Scanf("%s", &word)
		//fmt.Println(word)
		length := len(word)
		//try with two counters
		for j, v := range word {
			//fmt.Printf("Ascii Value of %c = %d\n", v, int(v))

			if length%2 != 0 && j == int(math.Ceil(float64(length/2.0))) {
				continue
			} else if j < (length / 2) {
				countArray[int(v)-97]++
			} else {
				countArray[int(v)-97]--
			}
		}
		var flag bool = true
		for _, v := range countArray {
			if v != 0 {
				flag = false
				break
			}
		}
		if flag == true {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
