package main

import "fmt"

func minElement(a int, b int, c int) {
	var array []int
	array = append(array, a, b, c)

	var m int = array[0]
	for i := 0; i < len(array); i++ {
		if array[i] < m {
			m = array[i]
		}
	}
	fmt.Printf("Min -> %d", m)
}

func average(numbers ...int) {
	var sum = 0
	var i = 0
	for _, number := range numbers {
		sum += number
		i++
	}
	fmt.Println("\naverage = ", sum/i)
}
