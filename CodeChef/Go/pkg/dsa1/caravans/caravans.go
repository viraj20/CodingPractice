package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ = strconv.Atoi((scanner.Text()))
	for i := 0; i < n; i++ {
		var numberOfCars int
		scanner.Scan()
		scanner.Scan()
		numberOfCars, _ = strconv.Atoi(scanner.Text())
		fmt.Println(numberOfCars, "---")
		count := 1
		//var currentMax int
		var currentValue int
		for i := 0; i < numberOfCars; i++ {
			scanner.Scan()
			currentValue, _ = strconv.Atoi((scanner.Text()))
			fmt.Println(currentValue, "***")
			//if i == 0 {
			//	currentMax = currentValue
			//}
			//if currentValue < currentMax {
			//	currentMax = currentValue
			//	count++
			//}
		}
		fmt.Println(count)
	}
}
