package main

import (
	"fmt"
	"math"
)

func main() {
	var k float64
	var p float64
	var v float64
	fmt.Scan(&k)
	fmt.Scan(&p)
	fmt.Scan(&v)

	var m float64
	m = M(p, v)
	var w float64
	w = W(m, k)
	var t float64
	t = T(w)
	fmt.Println(t)

}

func M(p, v float64) float64 {
	return p * v
}

func W(m, k float64) float64 {
	return math.Sqrt(k / m)
}

func T(w float64) float64 {
	return 6 / w
}
