package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	//Задание.
	//1. Вывести украинскую букву 'Ї'
	fmt.Printf("ї\n")

	//2. Пояснить назначение типа "rune"
	// тип данных для хранения одного символа в кодировке
	var letter rune = 'Ї'
	fmt.Printf("%c - %d\n", letter, letter)
	fmt.Println(letter)
}
