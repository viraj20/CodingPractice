package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	t, _ := strconv.Atoi(sc.Text())

	for i := 0; i < t; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		divisor := 5
		var numberOfZeros int
		for num >= divisor {
			numberOfZeros = numberOfZeros + (num / divisor)
			divisor = divisor * 5
		}
		fmt.Println(numberOfZeros)
	}
}
