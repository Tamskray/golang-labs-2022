package functions

func MinElement(a int, b int, c int) int {
	var array []int
	array = append(array, a, b, c)

	var m int = array[0]
	for i := 0; i < len(array); i++ {
		if array[i] < m {
			m = array[i]
		}
	}
	//fmt.Printf("Min -> %d", m)
	return m
}

func Average(numbers ...int) int {
	var sum = 0
	var i = 0
	for _, number := range numbers {
		sum += number
		i++
	}
	//fmt.Println("\naverage = ", sum/i)
	return sum / i
}

func Equation() {
	//
}
