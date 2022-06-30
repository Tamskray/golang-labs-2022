package main

import (
	"fmt"
	"math"
)

func lcm(x0, a, c, m, k int) []int {
	results := []int{}

	for i := 1; i <= k; i++ {
		x0 = (a*x0 + c) % m
		results = append(results, x0%250)
	}
	return results
}

func checked(arr []int, k, num int) float64 {
	var freq float64 = 0
	for i := 0; i < k; i++ {
		if arr[i] == num {
			freq++
		}
	}
	return freq
}

func main() {
	x0 := 1
	var a, c, m, k int
	a = 22695477
	c = 1
	m = 4294967296 // 2^32
	//k = 400
	k = 400

	randNums := lcm(x0, a, c, m, k)

	for i, value := range randNums {
		fmt.Println("Value", i+1, "->", value)
	}

	// імовірність появи
	frequency := [400]float64{}
	fmt.Println("\nЧастота інтервалів появи випадкових чисел")

	for i := 0; i < k; i++ {
		frequency[i] = checked(randNums, k, randNums[i])
		fmt.Println("Value №", i+1, "=", frequency[i])
	}

	// Статистична імовірність
	fmt.Println("\nСтатистична імовірність")
	stat := [400]float64{}
	for i := 0; i < k; i++ {
		stat[i] = (frequency[i] / 400)
	}

	for i, value := range stat {
		fmt.Println("Value №", i+1, "->", value)
	}

	// математичне сподівання
	var mathExp float64 = 0
	for i := 0; i < k; i++ {
		mathExp += float64(i) * stat[i]
	}
	fmt.Println("Математичне сподівання ->", mathExp)

	// дисперсія
	var despersion float64 = 0
	for i := 0; i < k; i++ {
		despersion += math.Pow((float64(i)-mathExp), 2) * stat[i]
	}
	fmt.Println("Дисперсія ->", despersion)

	// середньоквадратичне відхилення
	fmt.Println("Середньоквадратичне відхилення ->", math.Sqrt(despersion))
}
