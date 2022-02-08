package main

import "fmt"

func main() {
	//Инициализация переменных
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 //Такой вариант инициализации также возможен

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	//Краткая запись объявления переменной
	//только для новых переменных
	intVar := 10

	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	//Задание.
	//1. Вывести типы всех переменных
	fmt.Printf("Variables types -> %T, %T, %T, %T\n", userinit8, userinit16, userinit64, userautoinit)

	//2. Присвоить переменной intVar переменные userinit16 и userautoinit. Результат вывести.
	intVar = userautoinit
	fmt.Printf("intVar = %d\n", intVar)

	intVar = int(userinit16)
	fmt.Printf("intVar = %d\n", intVar)
}
