package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Scan(&a)
	a_len := len(a)
	a_max := 0
	a_num, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("stop")
		return
	}
	for i := 0; i < a_len; i++ {
		a_t := a_num % 10
		if a_t > a_max {
			a_max = a_t
		}
		a_num /= 10
	}
	fmt.Println(a_max)
}
