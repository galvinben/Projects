package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if length, _ := strconv.Atoi(os.Args[2]); length > 48 {
		fmt.Println("Don't be silly (48 max)")
	} else {
		switch os.Args[1] {
		case "pi":
			fmt.Println(fmt.Sprintf("%."+os.Args[2]+"f", math.Pi))
		case "e":
			fmt.Println(fmt.Sprintf("%."+os.Args[2]+"f", math.E))
		default:
			fmt.Println("please choose pi or e")
		}
	}
}
