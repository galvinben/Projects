package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if num <= 1 || err != nil {
		fmt.Println("Please send a number greater than 1")
		return
	}
	steps := checkSteps(num)
	fmt.Printf("Number of steps: %v\n", steps)
}

func checkSteps(num int) int {
	steps := 0
	one := false
	for one == false {
		if num%2 == 0 {
			num = num / 2
			steps, one = checkIfOne(num, steps)
		} else {
			num = (num * 3) + 1
			steps, one = checkIfOne(num, steps)
		}
		fmt.Println(num)
	}
	return steps
}

func checkIfOne(num int, steps int) (int, bool) {
	steps++
	if num == 1 {
		return steps, true
	}
	return steps, false
}
