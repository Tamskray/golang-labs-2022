package main

import "fmt"

func lcm(x0, a, c, m, k int) []int {
	results := []int{}

	for i := 1; i <= k; i++ {
		x0 = (a*x0 + c) % m
		results = append(results, x0%25000)
	}
	return results
}

func main() {
	x0 := 1
	var a, c, m, k int
	a = 22695477
	c = 1
	m = 4294967296 // 2^32
	k = 400

	randNums := lcm(x0, a, c, m, k)
	randNumsFloat := [400]float64{}
	for i := 0; i < k; i++ {
		randNumsFloat[i] = float64(randNums[i]) / 100
	}

	for i, value := range randNumsFloat {
		fmt.Println("Value", i+1, "->", value)
	}
}
