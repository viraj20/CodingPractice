package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 10000 * 11
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	n, _ = strconv.Atoi((scanner.Text()))
	for i := 0; i < n; i++ {
		scanner.Scan()
		scanner.Scan()
		count := 0
		var currentMax int
		var currentValue int

		carsMaxSpeed := scanner.Text()
		carsMaxSpeedArr := strings.Split(carsMaxSpeed, " ")
		for i, v := range carsMaxSpeedArr {
			currentValue, _ = strconv.Atoi(v)
			if i == 0 {
				currentMax = currentValue
				count++
			} else if currentValue < currentMax {
				currentMax = currentValue
				count++
			}
		}
		fmt.Println(count)
	}
}
