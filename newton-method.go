package main

import (
	"fmt"
	"math"
)

const (
	EPSILON = 0.001
)

func main() {
	fmt.Println(newtonRaphson(float64(4)))
}

func baseFunc(x float64) float64 {
	return x*x*x - x*x - 2
}

func derivFunc(x float64) float64 {
	return 3*x*x - x*2
}

func newtonRaphson(x float64) float64 {
	for h := baseFunc(x) / derivFunc(x); math.Abs(h) >= EPSILON; {
		h = baseFunc(x) / derivFunc(x)
		x = x - h
	}
	return x
}
