package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	for index, strir := range str {
		if index == len(str)-1 {
			fmt.Printf("%s", string(strir))
			break
		}
		fmt.Printf("%s*", string(strir))
	}
}
