package main

import "lab2/functions"

func main() {
	//functions.MinElement(3, -2, 5)
	minEl := functions.MinElement(-3, -2, 5)
	print("Min -> ", minEl)

	averageEl := functions.Average(15, -4, 13)
	print("\naverage -> ", averageEl)
	//functions.Average(15, -4, 13)

	//print("\n\nРівняння першого порядку\ny' = y/cos^2x\n")
	functions.Equation()
	//
}
