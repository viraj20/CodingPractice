package main

import (
	"fmt"
	"math"
)

func main() {
	var n, i int
	fmt.Scanf("%d", &n)
	for i = 0; i < n; i++ {
		var number int
		fmt.Scanf("%d", &number)
		fmt.Println(reverseNumber(number, numberOfDigits(number)))

	}

}

func reverseNumber(number int, numberOfDigits int) int {
	var response int
	for i := numberOfDigits - 1; i >= 0; i-- {
		lastDigit := number % 10
		number = number / 10
		response = response + lastDigit*int(math.Pow10(i))
	}
	return response

}

func numberOfDigits(number int) int {
	count := 1
	for number/10 != 0 {
		number = number / 10
		count++
	}
	return count
}

func reverse(x int) int {
	temp := 0
	for x > 0 {
		temp = temp*10 + (x % 10)
		x = x / 10
	}
	return temp
}
