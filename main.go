package main

import (
	"fmt"
	"os"
	"strconv"
)

func solveHanoi(t, from, to, work int) {
	if t == 1 {
		fmt.Printf("from %v to %v\n", from, to)
		return
	}
	solveHanoi(t-1, from, work, to)
	fmt.Printf("from %v to %v\n", from, to)
	solveHanoi(t-1, work, to, from)
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// from 1 to 2 using 3 as work space
	solveHanoi(num, 1, 2, 3)
}
